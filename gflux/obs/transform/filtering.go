// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package transform

import (
	"sync"
	"sync/atomic"

	"github.com/andrerrcosta2/gtools/gflux/obs"
	"github.com/andrerrcosta2/gtools/gflux/obs/internal/subscr"
)

// Distinct suppresses duplicate values emitted by the source Observable.
//
// Range value is compared against all previously seen values using
// Go's equality operator (==). If the value has already been observed,
// it is ignored; otherwise, it is forwarded downstream.
//
// Distinct preserves the order of first occurrences and maintains
// infra states tracking all distinct values seen so far.
//
// Completion:
//   - Completion from the source is forwarded after all values are processed.
//
// Error handling:
//   - Errors from the source are propagated immediately.
//
// Memory considerations:
//   - Distinct keeps all previously seen values in memory.
//   - It is best suited for finite or reasonably bounded streams.
//
// Diagram:
//
//	source:    --1--2--1--3--2--|
//	distinct:  --1--2-----3----|
//
// Example:
//
//	src := obs.Of(1, 2, 1, 3, 2, 4)
//
//	distinct := Distinct[int]()
//
//	distinct.Sub(obs.NewObserver(
//		func(v int) { fmt.Println(v) },
//		nil,
//		func() { fmt.Println("done") },
//	))(src)
//
// Output:
//
//	1
//	2
//	3
//	4
//	done
func Distinct[T comparable]() obs.TransformPipe[T, T] {
	return func(src obs.Observable[T]) obs.Observable[T] {
		return obs.Create[T](func(o *obs.Observer[T]) (obs.Subscription, error) {
			seen := make(map[T]struct{})
			terminated := atomic.Bool{}
			sub, err := src.Sub(&obs.Observer[T]{
				Next: func(v T) {
					if terminated.Load() {
						return
					}
					if _, ok := seen[v]; ok {
						return
					}
					seen[v] = struct{}{}
					o.Next(v)
				},

				Error: func(err error) {
					if terminated.CompareAndSwap(false, true) {
						o.Error(err)
					}
				},

				Complete: func() {
					if terminated.CompareAndSwap(false, true) {
						o.Complete()
					}
				},
			})

			if err != nil {
				return nil, err
			}

			return sub, nil
		})
	}
}

// DistinctUntilChanged suppresses consecutive duplicate values emitted
// by the source Observable.
//
// Range emitted value is compared only with the previously emitted value.
// If the value is equal to the previous one, it is ignored; otherwise,
// it is forwarded downstream.
//
// Unlike Distinct, this operator does NOT track all previously seen
// values and uses constant memory.
//
// Completion:
//   - Completion from the source is forwarded normally.
//
// Error handling:
//   - Errors from the source are propagated immediately.
//
// Diagram:
//
//	source:    --1--1--2--2--1--1--3--|
//	distinct:  --1-----2-----1-----3--|
//
// Example:
//
//	src := obs.Of(1, 1, 2, 2, 1, 1, 3)
//
//	filtered := DistinctUntilChanged[int]()
//
//	filtered.Sub(obs.NewObserver(
//		func(v int) { fmt.Println(v) },
//		nil,
//		func() { fmt.Println("done") },
//	))
//
// Output:
//
//	1
//	2
//	1
//	3
//	done
func DistinctUntilChanged[T comparable]() obs.TransformPipe[T, T] {
	return func(src obs.Observable[T]) obs.Observable[T] {
		return obs.Create[T](func(o *obs.Observer[T]) (obs.Subscription, error) {

			var (
				prev       T
				set        bool
				terminated = atomic.Bool{}
			)

			sub, err := src.Sub(&obs.Observer[T]{
				Next: func(v T) {
					if terminated.Load() || (set && v == prev) {
						return
					}
					prev = v
					set = true
					o.Next(v)
				},

				Error: func(err error) {
					if terminated.CompareAndSwap(false, true) {
						o.Error(err)
					}
				},

				Complete: func() {
					if terminated.CompareAndSwap(false, true) {
						o.Complete()
					}
				},
			})

			if err != nil {
				return nil, err
			}

			return sub, nil
		})
	}
}

// Filter emits only the values from the source Observable that satisfy
// the given predicate function.
//
// For each value emitted by the source Observable, Filter evaluates
// the predicate:
//
//   - If the predicate returns true, the value is forwarded downstream.
//   - If the predicate returns false, the value is ignored.
//
// Errors and completion signals from the source Observable are always
// forwarded downstream without modification.
//
// Filter does not alter subscription or completion timing; it only
// controls which values are emitted.
//
// Unsubscribing from the resulting Observable unsubscribes from the
// source Observable.
//
// Example:
//
//	src := obs.SlotSubject[int]()
//
//	evens := Filter(func(v int) bool {
//		return v%2 == 0
//	})(src)
//
//	evens.Sub(obs.NewObserver(
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
//	src.Next(4)
//	src.Complete()
//
// Output:
//
//	2
//	4
//	done
func Filter[T any](pred func(T) bool) obs.TransformPipe[T, T] {
	return func(src obs.Observable[T]) obs.Observable[T] {
		return obs.Create[T](func(o *obs.Observer[T]) (obs.Subscription, error) {
			return src.Sub(&obs.Observer[T]{
				Next: func(v T) {
					if pred(v) {
						o.Next(v)
					}
				},
				Error:    o.Error,
				Complete: o.Complete,
			})
		})
	}
}

// Skip suppresses the first `n` values emitted by the source Observable
// and forwards all subsequent values downstream.
//
// Skip counts values starting from the moment the source is subscribed.
// Once `n` values have been skipped, all remaining values are emitted
// unchanged.
//
// Completion semantics:
//   - If the source completes before `n` values are emitted, Skip
//     completes without emitting any values.
//   - Once skipping is finished, Skip behaves transparently and
//     forwards completion immediately when the source completes.
//
// Error semantics:
//   - Errors from the source are propagated immediately, even if
//     skipping has not completed.
//
// Unsubscription:
//   - Unsubscribing from the output Observable unsubscribes from
//     the source Observable.
//
// Skip is a stateful transformation and is safe to use with
// both hot and cold Observables.
//
// Diagram (n = 2):
//
//	source:  --1--2--3--4--5--|
//	skip:    --------3--4--5--|
//
// Example:
//
//	src := obs.FromSlice([]int{1, 2, 3, 4, 5})
//
//	skipped := Skip(2)
//
//	skipped.Sub(obs.NewObserver(
//		func(v int) {
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
// Output:
//
//	3
//	4
//	5
//	done
//
// Notes:
//
//   - If n <= 0, Skip forwards all values unchanged.
//   - Skip does not buffer values and does not delay emissions.
func Skip[T any](n int) obs.TransformPipe[T, T] {
	return func(src obs.Observable[T]) obs.Observable[T] {
		return obs.Create[T](func(o *obs.Observer[T]) (obs.Subscription, error) {
			if n <= 0 {
				return src.Sub(o)
			}

			var skipped int

			return src.Sub(&obs.Observer[T]{
				Next: func(v T) {
					if skipped < n {
						skipped++
						return
					}
					o.Next(v)
				},
				Error:    o.Error,
				Complete: o.Complete,
			})
		})
	}
}

// SkipUntil suppresses values emitted by the source Observable until a
// signal Observable emits or completes.
//
// Until the signal Observable produces its first Next or Complete event,
// all values from the source are ignored. Once the signal fires, all
// subsequent source values are forwarded downstream unchanged.
//
// Trigger semantics:
//   - signal.Next(...)     → stop skipping, start emitting source values
//   - signal.Complete()    → stop skipping, start emitting source values
//   - signal.Error(err)    → error is propagated downstream immediately
//
// Completion semantics:
//   - If the source completes before the signal fires, SkipUntil
//     completes without emitting any values.
//   - After the signal fires, completion is forwarded from the source
//     normally.
//
// Error semantics:
//   - Errors from the source are always propagated immediately.
//   - Errors from the signal are propagated immediately.
//
// Unsubscription:
//   - At skipping ends, the signal subscription is cancelled.
//   - Unsubscribing from the output Observable cancels both the
//     source and signal subscriptions.
//
// SkipUntil is a stateful transformation and is safe to use with
// both hot and cold Observables.
//
// Diagram:
//
//	source:  --1--2--3--4--5--6--|
//	signal:  -----------s-------|
//	skip:    -----------4--5--6--|
//
// Example:
//
//	src := obs.SlotSubject[int]()
//	sig := obs.SlotSubject[any]()
//
//	out := SkipUntil(sig)
//
//	out.Sub(obs.NewObserver(
//		func(v int) {
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
//	src.Next(1)
//	src.Next(2)
//	src.Next(3)
//
//	sig.Next(nil) // unlock emissions
//
//	src.Next(4)
//	src.Next(5)
//	src.Complete()
//
// Output:
//
//	4
//	5
//	done
//
// Notes:
//
//   - The signal Observable is used only as a trigger; its values
//     are ignored.
//   - SkipUntil does not buffer or replay skipped values.
func SkipUntil[T any](signal obs.Observable[any]) obs.TransformPipe[T, T] {
	return func(src obs.Observable[T]) obs.Observable[T] {
		return obs.Create[T](func(o *obs.Observer[T]) (obs.Subscription, error) {

			var (
				srcSub obs.Subscription
				sigSub obs.Subscription
				open   atomic.Bool
				once   sync.Once
			)

			sub := subscr.Func(func() {
				if srcSub != nil {
					srcSub.Unsub()
				}
				if sigSub != nil {
					sigSub.Unsub()
				}
			})

			openGate := func() {
				once.Do(func() {
					open.Store(true)
					if sigSub != nil {
						sigSub.Unsub()
					}
				})
			}

			sigSub, _ = signal.Sub(&obs.Observer[any]{
				Next:     func(any) { openGate() },
				Complete: openGate,
				Error:    o.Error,
			})

			srcSub, err := src.Sub(&obs.Observer[T]{
				Next: func(v T) {
					if open.Load() {
						o.Next(v)
					}
				},
				Error: func(err error) {
					o.Error(err)
					sub.Unsub()
				},
				Complete: o.Complete,
			})

			if err != nil {
				return nil, err
			}

			return sub, nil
		})
	}
}

// SkipWhile suppresses values emitted by the source Observable while a
// predicate function returns true.
//
// Values are skipped as long as pred(value) evaluates to true.
// Once pred returns false for the first time, SkipWhile stops skipping
// and forwards that value and all subsequent values downstream.
//
// Predicate semantics:
//   - pred(v) == true  → value is skipped
//   - pred(v) == false → value is emitted and skipping permanently ends
//
// Completion semantics:
//   - If the source completes while pred continues to return true,
//     SkipWhile completes without emitting any values.
//   - Once skipping ends, completion is forwarded normally.
//
// Error semantics:
//   - Errors from the source are propagated immediately.
//
// Unsubscription:
//   - Unsubscribing from the resulting Observable stops the source
//     subscription immediately.
//
// SkipWhile is a stateful transformation and is safe to use with
// both hot and cold Observables.
//
// Diagram:
//
//	source:   --1--2--3--4--5--6--|
//	pred:     T   T   T   F   F
//	skip:     ------------4--5--6--|
//
// Example:
//
//	src := obs.Of(1, 2, 3, 4, 5, 6)
//
//	out := SkipWhile(func(v int) bool {
//		return v < 4
//	}),
//
//	out.Sub(obs.NewObserver(
//		func(v int) {
//			fmt.Println(v)
//		},
//		func(err error) {},
//		func() {
//			fmt.Println("done")
//		},
//	))(src)
//
// Output:
//
//	4
//	5
//	6
//	done
//
// Notes:
//
//   - The predicate is evaluated only until it returns false once.
//   - Skipped values are discarded and never replayed.
//   - SkipWhile complements TakeWhile by inverting the gating condition.
func SkipWhile[T any](pred func(T) bool) obs.TransformPipe[T, T] {
	return func(src obs.Observable[T]) obs.Observable[T] {
		return obs.Create[T](func(o *obs.Observer[T]) (obs.Subscription, error) {

			skipping := true

			return src.Sub(&obs.Observer[T]{
				Next: func(v T) {
					if skipping {
						if pred(v) {
							return
						}
						skipping = false
					}
					o.Next(v)
				},
				Error:    o.Error,
				Complete: o.Complete,
			})
		})
	}
}

// Take emits only the first `n` values from the source Observable
// and then completes immediately.
//
// After emitting `n` values:
//   - The downstream observer receives Complete
//   - No further values, errors, or completion signals are forwarded
//   - The operator attempts to cancel the upstream subscription
//
// If the source completes before emitting `n` values, all emitted
// values are forwarded and completion is propagated normally.
//
// If the source errors before `n` values are emitted, the error is
// forwarded immediately.
//
// Special cases:
//
//   - n <= 0 → completes immediately without subscribing to the source
//
// Diagram:
//
//	source: --1--2--3--4--5--|
//	Take(3)
//	output: --1--2--3--|
//
// Early completion:
//
//	source: --1--2--|
//	Take(5)
//	output: --1--2--|
//
// Error propagation:
//
//	source: --1--2--X
//	Take(3)
//	output: --1--2--X
//
// Example:
//
//	src := obs.Of(1, 2, 3, 4, 5).Pipe(
//		Take(3),
//	)
//
//	src.Sub(obs.NewObserver(
//		func(v int) {
//			fmt.Println(v)
//		},
//		func(err error) {
//			fmt.Println("err:", err)
//		},
//		func() {
//			fmt.Println("done")
//		},
//	))
//
// Output:
//
//	1
//	2
//	3
//	done
//
// Semantics:
//
//   - Take preserves the order of values.
//   - Take does not buffer or delay values.
//   - Take introduces early completion.
//   - Take does not spawn goroutines.
//
// Subscription behavior:
//
//   - At possible, Take cancels the upstream subscription after the
//     value limit is reached.
//   - For synchronous sources that emit during Sub, cancellation may
//     occur after the source has already finished emitting.
//
// Notes:
//
//   - Take is a terminal-style operator with respect to the value count.
//   - Cleanup logic in synchronous sources must not rely on Unsub being
//     invoked to prevent emissions.
func Take[T any](n int) obs.TransformPipe[T, T] {
	return func(src obs.Observable[T]) obs.Observable[T] {
		return obs.Create[T](func(o *obs.Observer[T]) (obs.Subscription, error) {
			if n <= 0 {
				o.Complete()
				return subscr.Noop{}, nil
			}

			var (
				count      int
				terminated atomic.Bool
				sub        obs.Subscription
			)

			s, err := src.Sub(&obs.Observer[T]{
				Next: func(v T) {
					if terminated.Load() {
						return
					}

					o.Next(v)
					count++
					if count == n {
						if terminated.CompareAndSwap(false, true) {
							if sub != nil {
								sub.Unsub()
							}
							o.Complete()
						}
					}
				},
				Error: func(err error) {
					if terminated.CompareAndSwap(false, true) {
						o.Error(err)
					}
				},
				Complete: func() {
					if terminated.CompareAndSwap(false, true) {
						o.Complete()
					}
				},
			})

			if err != nil {
				return nil, err
			}

			sub = s

			// downstream unsubscribe
			return subscr.Func(func() {
				if terminated.CompareAndSwap(false, true) {
					sub.Unsub()
				}
			}), nil
		})
	}
}

// TakeUntil emits values from the source Observable until a signal Observable
// emits any value, completes, or errors.
//
// The resulting Observable mirrors the source until the signal triggers,
// at which point the source subscription is cancelled and the resulting
// Observable completes.
//
// Trigger conditions (any of the following):
//
//   - signal emits Next
//   - signal emits Complete
//
// Error handling:
//
//   - Errors from the source are propagated immediately.
//   - Errors from the signal are propagated immediately.
//   - Once an error is emitted, all subscriptions are cancelled.
//
// Completion rules:
//
//   - At the signal triggers, the resulting Observable completes,
//     regardless of whether the source has completed.
//   - If the source completes before the signal triggers, completion
//     is forwarded downstream.
//   - Completion is guaranteed to occur at most once.
//
// Diagram (signal emits a value):
//
//	source:  --1--2--3--4--5--|
//	signal: -----------•
//
//	output:  --1--2--3--------|
//
// Diagram (signal completes without emitting):
//
//	source:  --1--2--3--4--5--|
//	signal: -----------|
//
//	output:  --1--2--3--------|
//
// Diagram (signal errors):
//
//	source:  --1--2--3--4--5--|
//	signal: -----------X
//
//	output:  --1--2--3--------X
//
// Diagram (source completes first):
//
//	source:  --1--2--|
//	signal: -----------•
//
//	output:  --1--2--|
//
// Example:
//
//	stop := obs.SlotSubject[any]()
//
//	src := obs.Interval(time.Second).Pipe(
//		TakeUntil(stop),
//	)
//
//	sub := src.Sub(obs.NewObserver(
//		func(v int) {
//			fmt.Println(v)
//		},
//		func(err error) {
//			fmt.Println("err:", err)
//		},
//		func() {
//			fmt.Println("done")
//		},
//	))
//
//	time.Sleep(3 * time.Second)
//	stop.Next(nil)
//
// Output:
//
//	0
//	1
//	2
//	done
//
// Notes:
//
//   - TakeUntil does NOT emit any values from the signal Observable.
//   - The signal acts purely as a termination trigger.
//   - Both source and signal are subscribed eagerly.
//   - The operator is race-safe: completion is guarded by sync.Once.
//   - Unsubscribing from the resulting Observable cancels both source
//     and signal subscriptions.
func TakeUntil[T any](signal obs.Observable[any]) obs.TransformPipe[T, T] {
	return func(src obs.Observable[T]) obs.Observable[T] {
		return obs.Create[T](func(o *obs.Observer[T]) (obs.Subscription, error) {

			var (
				srcSub     obs.Subscription
				sigSub     obs.Subscription
				terminated atomic.Bool
			)

			terminate := func() {
				if terminated.CompareAndSwap(false, true) {
					if srcSub != nil {
						srcSub.Unsub()
					}
					if sigSub != nil {
						sigSub.Unsub()
					}
					o.Complete()
				}
			}

			sigSub, _ = signal.Sub(&obs.Observer[any]{
				Next: func(v any) {
					terminate()
				},
				Complete: func() {
					terminate()
				},
				Error: func(err error) {
					if terminated.CompareAndSwap(false, true) {
						o.Error(err)
					}
				},
			})

			srcSub, err := src.Sub(&obs.Observer[T]{
				Next: func(v T) {
					if terminated.Load() {
						return
					}
					o.Next(v)
				},
				Error: func(err error) {
					if terminated.CompareAndSwap(false, true) {
						o.Error(err)
					}
				},
				Complete: terminate,
			})

			if err != nil {
				return nil, err
			}

			return subscr.Func(terminate), nil
		})
	}
}

// TakeWhile emits values from the source Observable while the given
// predicate function returns true.
//
// At the predicate returns false for a value:
//   - That value is NOT emitted
//   - The resulting Observable completes immediately
//   - The source subscription is canceled
//
// 'TakeWhile' is a terminal operator based on a predicate condition
// and is commonly used to model bounded streams, guards, or early
// termination conditions.
//
// Example:
//
//	src := obs.Of(1, 2, 3, 4, 5)
//
//	taken := TakeWhile(func(v int) bool {
//		return v < 4
//	})(src)
//
//	taken.Sub(obs.NewObserver(
//		func(v int) { fmt.Println(v) },
//		func(err error) {},
//		func() { fmt.Println("done") },
//	))
//
// Output:
// 1
// 2
// 3
// done
//
// Diagram:
//
//	src:   --1--2--3--4--5--|
//	pred:   T  T  T  F
//
//	TakeWhile(v < 4):
//	       --1--2--3-------|
//
// Error handling:
//   - Errors from the source are propagated immediately.
//
// Subscription semantics:
//   - Unsubscribes from the source as soon as the predicate fails.
//   - Unsubscribing early stops further emissions.
//
// Note:
//   - The predicate is evaluated synchronously for each value.
//   - The failing value is never forwarded downstream.
func TakeWhile[T any](pred func(T) bool) obs.TransformPipe[T, T] {
	return func(src obs.Observable[T]) obs.Observable[T] {
		return obs.Create[T](func(o *obs.Observer[T]) (obs.Subscription, error) {

			var (
				srcSub obs.Subscription
				done   atomic.Bool
			)

			sw := subscr.Func(func() {
				if done.CompareAndSwap(false, true) {
					if srcSub != nil {
						srcSub.Unsub()
					}
				}
			})

			sub, err := src.Sub(&obs.Observer[T]{
				Next: func(v T) {
					if done.Load() {
						return
					}

					if !pred(v) {
						if done.CompareAndSwap(false, true) {
							o.Complete()
							sw.Unsub()
						}
						return
					}

					o.Next(v)
				},

				Error: func(err error) {
					if done.CompareAndSwap(false, true) {
						o.Error(err)
					}
				},

				Complete: func() {
					if done.CompareAndSwap(false, true) {
						o.Complete()
					}
				},
			})

			if err != nil {
				return nil, err
			}

			srcSub = sub
			return sw, nil
		})
	}
}
