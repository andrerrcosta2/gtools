// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package transform

import (
	"sync"
	"sync/atomic"

	"github.com/andrerrcosta2/gtools/core/domain/function"
	"github.com/andrerrcosta2/gtools/core/domain/function/runnables"
	"github.com/andrerrcosta2/gtools/gflux/core/runners"
	"github.com/andrerrcosta2/gtools/gflux/obs"
	"github.com/andrerrcosta2/gtools/gflux/obs/internal/subscr"
)

// ConcatMap dicts each value from the source Observable into an inner
// Observable and subscribes to them sequentially, one at a time.
//
// Range source value v is mapped using fn(v) into an inner Observable.
// ConcatMap subscribes to the next inner Observable only after the
// previous one has completed.
//
// Inner Observables are never active concurrently.
// Ordering of values is strictly preserved.
//
// Completion behavior:
//   - The resulting Observable completes only after:
//     1) the source Observable completes, and
//     2) the currently active inner Observable (if any) completes.
//
// Error behavior:
//   - If the source Observable errors, the error is propagated immediately.
//   - If any inner Observable errors, the error is propagated immediately.
//   - Upon error, all active subscriptions are cancelled.
//
// Diagram:
//
//	source:     --1----2----3----|
//	inner(1):      ---a--b--|
//	inner(2):               ----c--d--|
//	inner(3):                           --e--|
//
//	concatMap:  -----a--b------c--d------e--|
//
// Example:
//
//	src := obs.SlotSubject[int]()
//
//	mapped := ConcatMap(func(v int) obs.Observable[string] {
//		return obs.Of(
//			fmt.Sprintf("A%d", v),
//			fmt.Sprintf("B%d", v),
//		)
//	})(src)
//
//	mapped.Sub(obs.NewObserver(
//		func(v string) { fmt.Println(v) },
//		nil,
//		func() { fmt.Println("done") },
//	))
//
//	src.Next(1)
//	src.Next(2)
//	src.Next(3)
//	src.Complete()
//
// Output:
//
//	A1
//	B1
//	A2
//	B2
//	A3
//	B3
//	done
//
// Notes:
//
//   - ConcatMap preserves order and guarantees sequential execution.
//   - Unlike MergeMap, ConcatMap never runs inner Observables concurrently.
//   - Unlike SwitchMap, ConcatMap never cancels inner Observables.
//   - ConcatMap is ideal for modeling queues, ordered workflows,
//     and rate-limited or stateful operations.
func ConcatMap[I, O any](fn func(I) obs.Observable[O]) obs.TransformPipe[I, O] {
	return func(src obs.Observable[I]) obs.Observable[O] {
		return obs.Create[O](func(o *obs.Observer[O]) (obs.Subscription, error) {

			gate := runners.ConcSerialGate()

			var (
				srcDone  atomic.Bool
				ended    atomic.Bool
				srcSub   obs.Subscription
				innerSub atomic.Value // obs.Subscription
			)

			fulStop := func(callback func()) {
				if ended.CompareAndSwap(false, true) {
					gate.Stop()
					callback()
					if srcSub != nil {
						srcSub.Unsub()
					}
					if s, ok := innerSub.Load().(obs.Subscription); ok && s != nil {
						s.Unsub()
					}
				}
			}

			// Subscribe to source
			sub, err := src.Sub(&obs.Observer[I]{
				Next: func(v I) {
					gate.Run(func(done func()) {
						innerObservable := fn(v)
						s, err := innerObservable.Sub(&obs.Observer[O]{
							Next: o.Next,
							Error: func(err error) {
								fulStop(func() {
									o.Error(err)
								})
							},
							Complete: func() {
								done() // done goes to the next

								// If source already completed and this was the last inner observable
								if srcDone.Load() {
									gate.Run(func(_ func()) {
										fulStop(o.Complete)
									})
								}
							},
						})

						if err != nil {
							fulStop(func() {
								o.Error(err)
							})
							done()
							return
						}

						innerSub.Store(s)
					})
				},

				Error: func(err error) {
					fulStop(func() {
						o.Error(err)
					})
				},

				Complete: func() {
					srcDone.Store(true)

					// If no inner is running or queued, complete immediately
					gate.Run(func(_ func()) {
						fulStop(o.Complete)
					})
				},
			})

			if err != nil {
				return nil, err
			}

			srcSub = sub

			return subscr.Func(func() {
				fulStop(runnables.Void)
			}), nil
		})
	}
}

// ExhaustMap dicts each value emitted by the source Observable to an inner
// Observable, subscribing to it and emitting its values downstream.
//
// While an inner Observable is active, any new values emitted by the source
// are ignored. Once the active inner Observable completes, the next source
// value is accepted and mapped again.
//
// Behavior:
//   - Only one inner Observable can be active at a time.
//   - Source values emitted while an inner Observable is running are dropped.
//   - Values from the active inner Observable are forwarded downstream.
//   - Errors from either the source or the active inner Observable
//     are propagated immediately.
//   - Completion occurs when the source completes AND
//     there is no active inner Observable.
//
// ExhaustMap is useful when overlapping inner subscriptions must be avoided,
// such as preventing duplicate requests or ignoring rapid repeated triggers.
//
// Diagram:
//
//	source:      --A---B---C------D---|
//	inner(A):        ---1---2---|
//	inner(B):            (ignored)
//	inner(C):                (ignored)
//	inner(D):                    ---3---|
//
//	exhaustMap:  ------1---2----------3--|
//
// Example:
//
//	clicks := obs.SlotSubject[int]()
//
//	result := ExhaustMap(func(v int) obs.Observable[string] {
//			return fakeRequest(v)
//		})(clicks)
//
//	result.Sub(obs.NewObserver(
//		func(v string) {
//			fmt.Println(v)
//		},
//		func(err error) {},
//		func() {
//			fmt.Println("done")
//		},
//	))
//
//	clicks.Next(1)
//	clicks.Next(2) // ignored
//	clicks.Next(3) // ignored
//	// inner(1) completes
//	clicks.Next(4)
//
// Output (example):
//
//	response-1
//	response-4
//
// Notes:
//
//   - ExhaustMap differs from MergeMap by ignoring new values
//     while an inner Observable is active.
//   - Unlike SwitchMap, ExhaustMap does not cancel the active
//     inner Observable.
//   - Unlike ConcatMap, ExhaustMap does not queues values.
func ExhaustMap[I, O any](fn func(I) obs.Observable[O]) obs.TransformPipe[I, O] {
	return func(src obs.Observable[I]) obs.Observable[O] {
		return obs.Create[O](func(o *obs.Observer[O]) (obs.Subscription, error) {

			gate := runners.ConcExclusiveGate()

			var (
				srcSub      obs.Subscription
				innerSub    atomic.Value // obs.Subscription
				completedCb atomic.Value // func()
				ended       atomic.Bool  // marks whether we forwarded complete/error
			)

			fulStop := func(callback func()) {
				if ended.CompareAndSwap(false, true) {
					gate.Stop() // permanently stops the gate
					callback()
					if srcSub != nil {
						srcSub.Unsub()
					}
					if s, ok := innerSub.Load().(obs.Subscription); ok && s != nil {
						s.Unsub()
					}
				}
			}

			sub, err := src.Sub(&obs.Observer[I]{
				Next: func(v I) {
					gate.Run(func(done func()) {
						inner := fn(v)

						s, err := inner.Sub(&obs.Observer[O]{
							Next: func(val O) {
								o.Next(val)
								// If source called complete while this inner observable was running, run it now
								if cb, ok := completedCb.Load().(function.Runnable); ok && cb != nil {
									fulStop(cb)
								}
							},
							Error: func(err error) { fulStop(func() { o.Error(err) }) },
							Complete: func() {
								done()
								// Also check if there is a pending source complete
								if cb, ok := completedCb.Load().(function.Runnable); ok && cb != nil {
									fulStop(cb)
								}
							},
						})

						if err != nil {
							fulStop(func() { o.Error(err) })
							done()
							return
						}

						innerSub.Store(s)
					})
				},
				Error: func(err error) {
					// Rx law: Error is immediate, Complete is cooperative.
					fulStop(func() { o.Error(err) })
				},
				Complete: func() {
					// Try to run complete immediately via gate
					if !gate.Run(func(_ func()) { fulStop(o.Complete) }) {
						// Gate busy: store for later
						completedCb.Store(o.Complete)
					}
				},
			})

			if err != nil {
				return nil, err
			}

			srcSub = sub

			return subscr.Func(func() { fulStop(runnables.Void) }), nil
		})
	}
}

// Map transforms each value emitted by the source Observable by applying
// the provided mapping function.
//
// For every value v emitted by the source Observable, Map applies f(v)
// and emits the resulting value downstream.
//
// Errors and completion signals from the source Observable are forwarded
// downstream without modification.
//
// Map does not introduce buffering, concurrency, or reordering.
// Unsubscribing from the resulting Observable unsubscribes from the
// source Observable.
//
// Example:
//
//	src := obs.SlotSubject[int]()
//
//	mapped := Map(func(v int) string {
//		return fmt.Sprintf("value=%d", v)
//	})(src)
//
//	mapped.Sub(obs.NewObserver(
//		func(v string) {
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
//	src.Complete()
//
// Output:
//
//	value=1
//	value=2
//	done
func Map[I, O any](f func(I) O) obs.TransformPipe[I, O] {
	return func(src obs.Observable[I]) obs.Observable[O] {
		return obs.Create[O](func(o *obs.Observer[O]) (obs.Subscription, error) {
			return src.Sub(&obs.Observer[I]{
				Next: func(v I) {
					o.Next(f(v))
				},
				Error:    o.Error,
				Complete: o.Complete,
			})
		})
	}
}

// MergeMap projects each value emitted by the source Observable into an
// inner Observable and merges all inner Observables concurrently
// into a single output Observable.
//
// For every value received from the source, the projection function `fn`
// is applied to create a new inner Observable. All active inner Observables
// are subscribed to immediately, and their emissions are forwarded
// downstream as they arrive.
//
// MergeMap does not preserve ordering between inner Observables.
// Values are emitted as soon as they are produced by any active inner
// Observable.
//
// Completion semantics:
//   - The output Observable completes only after:
//     1. the source Observable completes, and
//     2. all active inner Observables have completed.
//   - If the source completes while inner Observables are still active,
//     MergeMap waits for all of them to finish.
//
// Error semantics:
//   - If the source Observable errors, the error is propagated immediately
//     and all active inner subscriptions are cancelled.
//   - If any inner Observable errors, the error is propagated immediately
//     and all active subscriptions (source and inners) are cancelled.
//
// Unsubscription:
//   - Unsubscribing from the output Observable cancels the source
//     subscription and all active inner subscriptions.
//
// MergeMap is a concurrent, stateful transformation and is safe to use
// with both hot and cold Observables.
//
// Diagram:
//
//	source:    --1----2--------3--|
//	project:   fn(1) fn(2)    fn(3)
//
//	fn(1):     ----a----b--|
//	fn(2):         ----c------d--|
//	fn(3):                   --e--|
//
//	merge:     ----a--c--b--d--e--|
//
// Example:
//
//	src := obs.FromSlice([]int{1, 2, 3})
//
//	result := MergeMap(func(v int) obs.Observable[string] {
//		return obs.FromSlice([]string{
//			fmt.Sprintf("%d-a", v),
//			fmt.Sprintf("%d-b", v),
//		})
//	}(src)
//
//
//	result.Sub(obs.NewObserver(
//		func(v string) {
//			fmt.Println(v)
//		},
//		func(err error) {
//			fmt.Println("error:", err)
//		},
//		func() {
//			fmt.Println("done")
//		},
//	))
//
// Possible output (order not guaranteed):
//
//	1-a
//	2-a
//	1-b
//	3-a
//	2-b
//	3-b
//	done
//
// Notes:
//
//   - MergeMap is also known as "flatMap" in other reactive libraries.
//   - Use ConcatMap to preserve ordering between inner Observables.
//   - Use SwitchMap to cancel previous inner Observables on new emissions.
//   - Use ExhaustMap to ignore new source values while an inner Observable
//     is active.
func MergeMap[I, O any](fn func(I) obs.Observable[O]) obs.TransformPipe[I, O] {
	return func(src obs.Observable[I]) obs.Observable[O] {
		return obs.Create[O](func(o *obs.Observer[O]) (obs.Subscription, error) {

			var (
				mu      sync.Mutex
				active  int
				srcDone bool
				ss      = subscr.SetOf()
			)

			checkComplete := func() {
				if srcDone && active == 0 {
					o.Complete()
				}
			}

			srcSub, err := src.Sub(&obs.Observer[I]{
				Next: func(v I) {
					mu.Lock()
					active++
					mu.Unlock()

					innerSub, err := fn(v).Sub(&obs.Observer[O]{
						Next: o.Next,
						Error: func(err error) {
							o.Error(err)
							ss.Unsub()
						},
						Complete: func() {
							mu.Lock()
							active--
							checkComplete()
							mu.Unlock()
						},
					})

					if err != nil {
						o.Error(err)
						ss.Unsub()
						return
					}

					ss.Add(innerSub)
				},
				Error: func(err error) {
					o.Error(err)
					ss.Unsub()
				},
				Complete: func() {
					mu.Lock()
					srcDone = true
					checkComplete()
					mu.Unlock()
				},
			})

			if err != nil {
				return nil, err
			}

			ss.Add(srcSub)
			return ss, nil
		})
	}
}

// SwitchMap dicts each value from the source observable into a new inner
// observable and emits values only from the most recently created inner.
//
// At the source emits a value v, SwitchMap:
//
//  1. Cancels (unsubscribes from) the previous inner observable, if any
//  2. Calls fn(v) to create a new inner observable
//  3. Subscribes to the new inner and forwards its emissions
//
// Values emitted by older inner observables are discarded once a new
// source value arrives.
//
// Diagram:
//
//	source:    --1------2----------|
//
//	fn(1):       --A1----B1----|
//	fn(2):               --A2----B2----|
//
//	SwitchMap(fn):
//	           --A1----A2----B2----|
//
// Completion rules:
//
//   - If the source completes while an inner observable is still active,
//     completion is deferred until that inner completes.
//   - If the source completes and there is no active inner observable,
//     the resulting observable completes immediately.
//
// Error handling:
//
//   - Errors from either the source or the active inner observable are
//     immediately propagated downstream.
//   - Upon error, both source and inner subscriptions are cancelled.
//
// SwitchMap is useful for modeling "latest wins" scenarios such as
// request switching, cancelable workflows, or reactive states updates
// where only the most recent computation should remain active.
func SwitchMap[I, O any](fn func(I) obs.Observable[O]) obs.TransformPipe[I, O] {
	return func(src obs.Observable[I]) obs.Observable[O] {
		return obs.Create[O](func(o *obs.Observer[O]) (obs.Subscription, error) {

			var (
				outerSub  obs.Subscription
				outerDone atomic.Bool

				gen       atomic.Uint64 // generation counter
				innerSub  obs.Subscription
				innerDone atomic.Bool
			)

			cs := subscr.Func(func() {
				if outerSub != nil {
					outerSub.Unsub()
				}
				if innerSub != nil {
					innerSub.Unsub()
				}
			})

			s, err := src.Sub(&obs.Observer[I]{
				Next: func(v I) {
					// Advance generation
					thisGen := gen.Add(1)
					innerDone.Store(false)

					// Cancel previous inner
					if innerSub != nil {
						innerSub.Unsub()
						innerSub = nil
					}

					inner := fn(v)

					sub, err := inner.Sub(&obs.Observer[O]{
						Next: func(val O) {
							if gen.Load() != thisGen || cs.Closed() {
								return
							}
							o.Next(val)
						},
						Error: func(err error) {
							if gen.Load() == thisGen && !cs.Closed() {
								o.Error(err)
								cs.Unsub()
							}
						},
						Complete: func() {
							innerDone.Store(true)

							if gen.Load() == thisGen &&
								outerDone.Load() &&
								!cs.Closed() {
								o.Complete()
								cs.Unsub()
							}
						},
					})

					if err != nil {
						if !cs.Closed() {
							o.Error(err)
							cs.Unsub()
						}
						return
					}

					innerSub = sub
				},

				Error: func(err error) {
					if !cs.Closed() {
						o.Error(err)
						cs.Unsub()
					}
				},

				Complete: func() {
					outerDone.Store(true)

					if (innerSub == nil || innerDone.Load()) &&
						!cs.Closed() {
						o.Complete()
						cs.Unsub()
					}
				},
			})

			if err != nil {
				return nil, err
			}

			outerSub = s

			return cs, nil
		})
	}
}
