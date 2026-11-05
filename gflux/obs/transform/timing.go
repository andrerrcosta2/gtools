// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package transform

import (
	"sync/atomic"
	"time"

	"github.com/andrerrcosta2/gtools/conc/actors/mailboxes"
	"github.com/andrerrcosta2/gtools/core/domain/contracts/timing/scheduler"
	"github.com/andrerrcosta2/gtools/datastr/buffers"

	"github.com/andrerrcosta2/gtools/core/domain/contracts/chrono"
	"github.com/andrerrcosta2/gtools/gflux/core/runners"
	"github.com/andrerrcosta2/gtools/gflux/events"
	"github.com/andrerrcosta2/gtools/gflux/obs"
	"github.com/andrerrcosta2/gtools/gflux/obs/internal/subscr"
)

// Audit emits the most recent value from the source Observable
// after a silence period of duration d that starts when the first
// value in a window is received.
//
// At a value arrives and no window is active, Audit:
//
//  1. Starts a timer of duration d.
//  2. Suppresses all further emissions during that period.
//  3. Remembers the most recent value received while the timer is active.
//
// At the timer expires, the latest remembered value (if any)
// is emitted downstream, and the window closes.
//
// A new window does NOT begin until another source value arrives.
// Windows are event-driven, not globally time-aligned.
//
// In other words:
//
//   - The timer starts on the first value of a window.
//   - The timer is NOT reset by subsequent values.
//   - No values are emitted immediately.
//   - Exactly one value (the most recent) is emitted per window.
//
// Completion behavior:
//
//   - If the source completes while a window is active and a value
//     has been observed in that window, the latest value is emitted
//     before completion.
//   - Completion is then forwarded downstream.
//   - If no window is active, completion is forwarded immediately.
//
// Error behavior:
//
//   - Errors from the source are propagated immediately.
//   - Any active window is canceled and no pending value is emitted.
//
// Diagram:
//
//	source:  --1--2--3-----4--5------|
//	audit:   ---------3--------5----|
//	window:    [----d----]   [---d--]
//
// Explanation:
//
//   - Value 1 starts a window.
//   - Values 2 and 3 arrive while the window is active.
//   - After d, value 3 (the latest) is emitted.
//   - Value 4 starts a new window.
//   - Value 5 replaces 4 as the latest value.
//   - After d, value 5 is emitted.
//   - Completion is then forwarded.
//
// Example:
//
//	src := obs.SlotSubject[int]()
//
//	audited := Audit[int](100 * time.Millisecond)(src)
//
//	audited.Sub(obs.NewObserver(
//	    func(v int) { fmt.Println(v) },
//	    nil,
//	    func() { fmt.Println("done") },
//	))
//
//	src.Next(1)
//	time.Sleep(30 * time.Millisecond)
//	src.Next(2)
//	time.Sleep(30 * time.Millisecond)
//	src.Next(3)
//
//	time.Sleep(120 * time.Millisecond)
//
//	src.Next(4)
//	src.Next(5)
//	src.Complete()
//
// Output:
//
//	3
//	5
//	done
//
// Notes:
//
//   - Audit is useful for emitting the trailing value of bursts of commands.
//   - Unlike Throttle (leading), Audit never emits immediately.
//   - Unlike Debounce, Audit does not wait for a quiet period;
//     it emits exactly d after the first value in a window.
//   - Windows are started by commands, not by wall-clock alignment.
func Audit[T any](d time.Duration) obs.TransformPipe[T, T] {
	if d < 0 {
		panic("audit duration must be >= 0")
	}
	if d == 0 {
		return func(src obs.Observable[T]) obs.Observable[T] {
			return src
		}
	}

	return func(src obs.Observable[T]) obs.Observable[T] {
		return obs.Create[T](func(o *obs.Observer[T]) (obs.Subscription, error) {

			var (
				next        T
				hasVal      bool
				timerActive bool
				stopped     bool
				m           events.Machine[auditEvent[T]]
			)

			startTimer := func() {
				time.AfterFunc(d, func() {
					m.Send(auditEvent[T]{timer: true})
				})
			}

			cleanup := func() {
				stopped = true
			}

			// Audit uses a FIFO buffer mailbox because transport-level coalescing is unsafe:
			// every upstream value must reach the states machine, even though only the latest
			// is retained semantically.
			//
			// Coalescing happens strictly at the STATE level, not in the mailbox.
			//
			// Invariants:
			//
			// 1. Next commands (unbounded producer):
			//    - Upstream may emit arbitrarily fast
			//    - All Next commands must be delivered (no dropping)
			//    - State overwrites previous values avoiding accumulation in states
			//
			// 2. Timer commands:
			//    - At most 1 active timer at any time
			//    - Timer may race with Next and terminal commands
			//    - +1 concurrent slot
			//
			// 3. Terminal commands (Error, Complete):
			//    - Must never be dropped
			//    - May arrive concurrently with Next or timer
			//    - +1 slot
			//
			// 4. Unsubscription:
			//    - External teardown may race with all other commands
			//    - +1 slot
			//
			// 5. Scheduler / goroutine skew:
			//    - Producers (Next), timer callback, and consumer run independently
			//    - Short bursts may enqueue multiple commands before draining
			//    - Requires small buffering margin
			//
			// Worst-case concurrent in-flight commands:
			//
			//   - N Next commands (burst, bounded by backpressure)
			//   - 1 timer event
			//   - 1 terminal event
			//   - 1 unsubscribe event
			//
			// The buffer provides BACKPRESSURE, so N is bounded by buffer size.
			//
			// Therefore:
			//
			//   structural minimum: 3–4 (Next + timer + terminal [+ unsub])
			//   practical buffer:   6–8 (absorbs burst + scheduler skew)
			//
			// We choose 8 as a balanced trade-off:
			//   - avoids excessive producer blocking during bursts
			//   - keeps latency low
			//   - preserves strict ordering and correctness
			//
			// Larger buffers do NOT improve throughput meaningfully and may hide
			// backpressure issues upstream.
			m = events.WeakMachine(
				events.AsyncScheduler(),
				mailboxes.Ring[auditEvent[T]](8),
				func(e auditEvent[T]) {
					if stopped {
						return
					}

					switch {
					case e.unsub:
						cleanup()
						m.Stop()
						return

					case e.err != nil:
						cleanup()
						o.Error(e.err)
						m.Stop()
						return

					case e.complete:
						if hasVal {
							o.Next(next)
							hasVal = false
						}
						o.Complete()
						cleanup()
						m.Stop()
						return

					case e.timer:
						if hasVal {
							o.Next(next)
							hasVal = false
						}
						timerActive = false
						return

					case e.hasValue:
						// STATE-LEVEL COALESCING
						next = e.value
						hasVal = true

						if !timerActive {
							timerActive = true
							startTimer()
						}
					}
				},
			)

			sub, err := src.Sub(&obs.Observer[T]{
				Next: func(v T) {
					m.Send(auditEvent[T]{
						value:    v,
						hasValue: true,
					})
				},
				Error: func(err error) {
					m.Send(auditEvent[T]{err: err})
				},
				Complete: func() {
					m.Send(auditEvent[T]{complete: true})
				},
			})

			if err != nil {
				m.Stop()
				return nil, err
			}

			return subscr.Func(func() {
				m.Send(auditEvent[T]{unsub: true})
				sub.Unsub()
			}), nil
		})
	}
}

// Debounce emits a value from the source Observable only after
// a specified duration has passed without another value being emitted.
//
// Debounce suppresses intermediate values that arrive within the
// debounce duration and emits only the most recent value once
// the source becomes silent for the configured period.
//
// If the source emits multiple values rapidly, only the last value
// in each burst is forwarded downstream.
//
// At the source Observable completes, any pending debounced value
// is emitted immediately before the completion signal is forwarded.
//
// Errors from the source are propagated immediately and cancel any
// pending debounced emission.
//
// Debounce is a stateful, time-based transformation and introduces
// asynchronous behavior. It is safe to use with both hot and cold
// Observables.
//
// Diagram (debounce = 3 time units):
//
//	source:    --1--2--3------4----|
//	debounce: -----------3------4-|
//
// Example:
//
//	src := obs.SlotSubject[int]()
//
//	debounced := Debounce[int](300 * time.Millisecond)(src)
//
//	debounced.Sub(obs.NewObserver(
//		func(v int) {
//			fmt.Println(v)
//		},
//		func(err error) {},
//		func() {
//			fmt.Println("done")
//		},
//	))
//
//	src.Next(1)
//	src.Next(2)
//	src.Next(3)
//
//	time.Sleep(400 * time.Millisecond)
//
//	src.Next(4)
//	src.Complete()
//
// Output:
//
//	3
//	4
//	done
//
// Panics if the debounce duration is negative.
func Debounce[T any](d time.Duration, ts scheduler.TimerFactory) obs.TransformPipe[T, T] {
	if d < 0 {
		panic("debounce duration must be >= 0")
	}
	return func(src obs.Observable[T]) obs.Observable[T] {
		return obs.Create[T](func(o *obs.Observer[T]) (obs.Subscription, error) {
			var (
				stopped bool
				version uint64
				timer   chrono.Timer
				last    T
				hasLast bool

				m events.Machine[debounceEvent[T]]
			)

			// single callback instance (ZERO allocations after init)
			timer = ts.NewTimer(func() {
				v := atomic.LoadUint64(&version)
				m.Send(debounceEvent[T]{kind: onTimeout, version: v})
			})

			stopTimer := func() {
				if timer != nil {
					timer.Stop()
				}
			}

			resetTimer := func() {
				timer.Reset(chrono.FromStdDuration(d))
			}

			// We use a cancellable mailbox with a small bounded buffer.
			// The size here isn't about throughput, but about handling transient
			// scheduling skew between producers (Next/Timeout/terminal commands) and
			// the single-threaded machine consumer.
			//
			// Invariants:
			//
			// 1. Cancellable commands collapse by kind:
			//    - onNext:    only the latest matters (states replacement)
			//    - onTimeout: only the latest matters (versioned)
			//    - at most 1 onNext + 1 onTimeout in-flight
			//
			// 2. Terminal commands (onError, onComplete) are NOT cancellable:
			//    - they must never be dropped or replaced
			//    - +1 slot required
			//
			// 3. Async scheduling + timer races:
			//    - A timer may enqueue onTimeout while a newer onNext is already queued
			//    - Producers may enqueue before the machine drains the mailbox
			//    - transient coexistence of:
			//       [onNext(latest), onTimeout(latest), terminal]
			//
			// 4. Teardown race (unsubscribe):
			//    - onUnsub can briefly overlap with terminal or in-flight commands
			//    - +1 defensive slot
			//
			// Therefore:
			//   theoretical minimum: 3 (onNext + onTimeout + terminal)
			//   defensive bound:     4 (accounts for teardown/scheduler skew)
			//
			// If this buffer is exceeded, it indicates a bug in cancellable semantics
			// more likely commands not properly collapsing, not a capacity issue.
			mb := mailboxes.Buffered(buffers.OverwritableRing[debounceEvent[T]](2))
			m = events.WeakMachine(
				events.AsyncScheduler(),
				mb,
				func(e debounceEvent[T]) {

					if stopped {
						return
					}

					switch e.kind {

					case onNext:
						last = e.val
						hasLast = true

						atomic.AddUint64(&version, 1)
						resetTimer()

					case onTimeout:
						// version check = stale timer protection
						if e.version != atomic.LoadUint64(&version) {
							return
						}

						if hasLast {
							o.Next(last)
							hasLast = false
						}
					}
				},
			)

			sub, err := src.Sub(&obs.Observer[T]{
				Next: func(v T) {
					m.Send(debounceEvent[T]{kind: onNext, val: v})
				},
				Error: func(err error) {
					stopTimer()
					mb.Shutdown().Wait()
					stopped = true
					o.Error(err)
					m.Stop()
				},
				Complete: func() {
					stopTimer()
					mb.Shutdown().Wait()
					stopped = true
					o.Complete()
					m.Stop()
				},
			})

			if err != nil {
				timer.Stop()
				m.Stop()
				return nil, err
			}

			return subscr.Func(func() {
				stopped = true
				stopTimer()
				timer.Stop()
				m.Stop()
				sub.Unsub()
			}), nil
		})
	}
}

// Delay shifts all values emitted by the source Observable forward in time
// by the specified duration.
//
// Range value emitted by the source is scheduled to be emitted downstream
// after the delay duration elapses. The relative order of values is preserved.
//
// Completion behavior:
//   - Completion is also delayed until all scheduled values have been emitted.
//   - The completion signal is emitted after the last delayed value.
//
// Error behavior:
//   - Errors from the source are propagated immediately.
//   - Pending delayed values are discarded when an error occurs.
//
// Timing behavior:
//   - Delay does not affect the rate at which the source produces values.
//   - Values may accumulate internally if the source emits faster than
//     the delay duration.
//
// Delay is a time-based transformation and relies on timers.
// It does not modify or resubscribe to the source Observable.
//
// Diagram (delay = 2s):
//
//	source:  --1--2--3--|
//	delay:   ----1--2--3--|
//
// Example:
//
//	src := obs.SlotSubject[int]()
//
//	delayed := Delay(2 * time.Second)(src)
//
//	delayed.Sub(obs.NewObserver(
//		func(v int) {
//			fmt.Println(v, time.AtNow())
//		},
//		func(err error) {},
//		func() {
//			fmt.Println("done")
//		},
//	))
//
//	src.Next(1)
//	src.Next(2)
//	src.Next(3)
//	src.Complete()
//
// Output (timestamps approximate):
//
//	1  +2s
//	2  +2s
//	3  +2s
//	done
//
// Notes:
//
//   - Delay preserves ordering but shifts emissions in time.
//   - Errors are not delayed.
//   - Delay is useful for temporal alignment, UI effects,
//     and simulating latency.
func Delay[T any](delay time.Duration, bufferSize int) obs.TransformPipe[T, T] {
	if delay < 0 {
		panic("delay must be >= 0")
	}

	return func(src obs.Observable[T]) obs.Observable[T] {
		return obs.Create[T](func(o *obs.Observer[T]) (obs.Subscription, error) {

			start := time.Now()
			ss := subscr.SetOf()
			gate := runners.AsyncSerialGate(bufferSize)

			schedule := func(fn func()) {
				gate.Run(func(done func()) {
					if ss.Closed() {
						done()
						return
					}

					elapsed := time.Since(start)
					remaining := delay - elapsed

					if remaining > 0 {
						time.Sleep(remaining)
					}

					if !ss.Closed() {
						fn()
					}

					done()
				})
			}

			sub, err := src.Sub(&obs.Observer[T]{
				Next: func(v T) {
					val := v
					schedule(func() {
						o.Next(val)
					})
				},

				Error: func(err error) { // should be immediate
					if ss.Closed() {
						return
					}

					o.Error(err)
					gate.Stop()
					ss.Unsub()
				},

				Complete: func() {
					schedule(func() {
						o.Complete()
						gate.Stop()
						ss.Unsub()
					})
				},
			})

			if err != nil {
				return nil, err
			}

			ss.Add(sub)

			// External unsubscribe
			ss.Add(subscr.Func(func() {
				gate.Stop()
			}))

			return ss, nil
		})
	}
}

// Throttle limits the rate of emissions from the source Observable by
// allowing at most one value to be emitted during each time window of
// duration `d`.
//
// At the source emits a value:
//   - If no throttle window is active, the value is emitted immediately
//     and a throttle window of duration `d` starts.
//   - If a throttle window is active, the value is ignored (dropped).
//
// Throttle does NOT delay values — it either emits immediately or drops
// them, depending on timing.
//
// Completion semantics:
//   - At the source completes, completion is forwarded immediately.
//   - No buffered or suppressed values are emitted on completion.
//
// Error semantics:
//   - Errors from the source are propagated immediately and cancel
//     any active throttle window.
//
// Unsubscription:
//   - Unsubscribing cancels the source subscription and any active
//     throttle timer.
//
// Throttle is a stateful, time-based transformation and is safe to use
// with both hot and cold Observables.
//
// Diagram (d = throttle duration):
//
//	source:   --1-2-3-----4--5----|
//	throttle: --1--------4-------|
//
// Example:
//
//	src := obs.SlotSubject[int]()
//
//	out := Throttle(1 * time.Second)(src)
//
//	out.Sub(obs.NewObserver(
//		func(v int) {
//			fmt.Println("value:", v)
//		},
//		func(err error) {
//			fmt.Println("error:", err)
//		},
//		func() {
//			fmt.Println("done")
//		},
//	))
//
//	src.Next(1) // emitted
//	src.Next(2) // dropped
//	src.Next(3) // dropped
//	time.Sleep(1 * time.Second)
//	src.Next(4) // emitted
//	src.Next(5) // dropped
//	src.Complete()
//
// Output:
//
//	value: 1
//	value: 4
//	done
//
// Notes:
//
//   - Throttle emits the FIRST value in each time window.
//   - Use Debounce if you want the LAST value instead.
//   - Use Audit or Sample for periodic emission strategies.
//
// Panics if d <= 0.
func Throttle[T any](d time.Duration) obs.TransformPipe[T, T] {
	if d < 0 {
		panic("throttle duration must be >= 0")
	}

	return func(src obs.Observable[T]) obs.Observable[T] {
		return obs.Create[T](func(o *obs.Observer[T]) (obs.Subscription, error) {

			gate := runners.ThrottleGate(d)

			sub, err := src.Sub(&obs.Observer[T]{
				Next: func(v T) {
					gate.Run(func(_ func()) {
						o.Next(v)
					})
				},

				Error: func(err error) {
					gate.Stop()
					o.Error(err)
				},

				Complete: func() {
					gate.Stop()
					o.Complete()
				},
			})

			if err != nil {
				return nil, err
			}

			return subscr.Func(func() {
				gate.Stop()
				sub.Unsub()
			}), nil
		})
	}
}

// Timeout enforces a maximum duration between emitted values from the
// source Observable.
//
// If the source Observable fails to emit a value within the specified
// duration `d`, Timeout terminates the sequence with a timeout error.
//
// Timeout resets its timer:
//   - On subscription (initial timeout window)
//   - After each successfully emitted value
//
// If a value arrives before the timeout expires, it is forwarded
// downstream and the timeout window is restarted.
//
// Completion semantics:
//   - If the source completes before a timeout occurs, completion is
//     forwarded normally.
//   - If a timeout occurs, the source subscription is cancelled and
//     an error is emitted downstream.
//
// Error semantics:
//   - Errors from the source are propagated immediately.
//   - Timeout errors are terminal and cancel the source subscription.
//
// Unsubscription:
//   - Unsubscribing cancels both the source subscription and the
//     infra timeout timer.
//
// Timeout is a stateful, time-based transformation and is safe to use
// with both hot and cold Observables.
//
// Diagram (d = timeout duration):
//
//	source:   --1----2----------3--|
//	timeout:  --1----2----X
//
// (X = timeout error)
//
// Example:
//
//	src := obs.SlotSubject[int]()
//
//	out := Timeout(2 * time.Second)(src)
//
//	out.Sub(&obs.Observer{
//		Next: func(v int) {
//			fmt.Println("value:", v)
//		},
//		Error: func(err error) {
//			fmt.Println("error:", err)
//		},
//		Complete: func() {
//			fmt.Println("done")
//		},
//	})
//
//	src.Next(1)        // ok
//	time.Sleep(1 * time.Second)
//	src.Next(2)        // ok
//	time.Sleep(3 * time.Second) // timeout triggers
//
// Output:
//
//	value: 1
//	value: 2
//	error: timeout
//
// Notes:
//
//   - Timeout measures inactivity, not total execution time.
//   - The timeout window is restarted after each emission.
//   - This operator is useful for detecting stalled streams,
//     enforcing responsiveness, or guarding against deadlocks.
//
// Panics if d <= 0.
func Timeout[T any](d time.Duration, ts scheduler.TimerFactory) obs.TransformPipe[T, T] {
	if d <= 0 {
		panic("Timeout duration must be > 0")
	}
	if ts == nil {
		panic("TimerFactory must not be nil")
	}

	return func(src obs.Observable[T]) obs.Observable[T] {
		return obs.Create[T](func(o *obs.Observer[T]) (obs.Subscription, error) {
			var (
				stopped bool
				timer   chrono.Timer
				m       events.Machine[timeoutEvent[T]]
			)

			stopTimer := func() {
				if timer != nil {
					timer.Stop()
				}
			}

			startTimer := func() {
				timer.Reset(chrono.FromStdDuration(d))
			}

			// Timeout uses a FIFO buffer mailbox because:
			//   - every upstream value must pass through immediately
			//   - no coalescing is allowed (each value resets the timer)
			//
			// Invariants:
			//
			// 1. Next commands:
			//    - unbounded producer
			//    - must NOT be dropped
			//
			// 2. Timer event:
			//    - at most 1 active timer: +1 slot
			//
			// 3. Terminal commands:
			//    - must not be dropped: +1 slot
			//
			// 4. Unsubscription:
			//    - may race with all others: +1 slot
			//
			// 5. Scheduler skew:
			//    - Next + timeout + terminal may overlap
			//
			// Therefore:
			//   structural minimum: 3–4
			//   chosen buffer:      8
			m = events.WeakMachine(
				events.AsyncScheduler(),
				mailboxes.Ring[timeoutEvent[T]](8),
				func(e timeoutEvent[T]) {

					if stopped {
						return
					}

					switch {
					case e.unsub:
						stopped = true
						stopTimer()
						m.Stop()
						return

					case e.err != nil:
						stopped = true
						stopTimer()
						o.Error(e.err)
						m.Stop()
						return

					case e.complete:
						stopped = true
						stopTimer()
						o.Complete()
						m.Stop()
						return

					case e.timeout:
						stopped = true
						o.Error(obs.TimeoutError{After: d})
						m.Stop()
						return

					case e.hasValue:
						o.Next(e.value)

						stopTimer()
						startTimer()
					}
				},
			)

			timer = ts.NewTimer(func() {
				m.Send(timeoutEvent[T]{timeout: true})
			})

			// start immediately
			startTimer()

			sub, err := src.Sub(&obs.Observer[T]{
				Next: func(v T) {
					m.Send(timeoutEvent[T]{value: v, hasValue: true})
				},
				Error: func(err error) {
					m.Send(timeoutEvent[T]{err: err})
				},
				Complete: func() {
					m.Send(timeoutEvent[T]{complete: true})
				},
			})

			if err != nil {
				stopTimer()
				m.Stop()
				return nil, err
			}

			return subscr.Func(func() {
				m.Send(timeoutEvent[T]{unsub: true})
				sub.Unsub()
			}), nil
		})
	}
}
