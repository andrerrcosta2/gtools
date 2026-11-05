// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package transform

import (
	"sync"
	"sync/atomic"

	"github.com/andrerrcosta2/gtools/gflux/obs"
	"github.com/andrerrcosta2/gtools/gflux/obs/internal/subscr"
)

// DoOnComplete registers a side-effect function that is executed
// when the source Observable completes successfully.
//
// The provided function is invoked exactly once, immediately
// before the completion signal is forwarded downstream.
//
// Behavior:
//   - Values emitted by the source are forwarded unchanged.
//   - The side-effect runs only on normal completion.
//   - Errors do NOT trigger the side-effect.
//   - The completion signal is always forwarded after the side-effect executes.
//
// DoOnComplete does not modify the data flow and does not
// affect subscription or unsubscription behavior.
//
// Diagram:
//
//	source:        --1--2--3--|
//	doOnComplete:  --1--2--3--(fn)--|
//
// Example:
//
//	src := obs.SlotSubject[int]()
//
//	withSideEffect := DoOnComplete(func() {
//		fmt.Println("completed")
//	})(src)
//
//	withSideEffect.Sub(obs.NewObserver(
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
//	src.Complete()
//
// Output:
//
//	1
//	2
//	3
//	completed
//	done
//
// Notes:
//
//   - Use DoOnComplete for logging, metrics, or lifecycle hooks.
//   - For cleanup logic that must run on completion, error,
//     or unsubscription, prefer Finalize / Finally.
func DoOnComplete[T any](fn func()) obs.TransformPipe[T, T] {
	return func(src obs.Observable[T]) obs.Observable[T] {
		return obs.Create[T](func(o *obs.Observer[T]) (obs.Subscription, error) {

			var (
				sub  obs.Subscription
				done atomic.Bool
			)

			s, err := src.Sub(&obs.Observer[T]{
				Next: func(v T) {
					if done.Load() {
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
						fn()
						o.Complete()
					}
				},
			})

			if err != nil {
				return nil, err
			}

			sub = s

			return subscr.Func(func() {
				if done.CompareAndSwap(false, true) {
					sub.Unsub()
				}
			}), nil
		})
	}
}

// DoOnError registers a side-effect function that is invoked when the
// source Observable emits an error.
//
// DoOnError does NOT modify the stream:
//   - Values (Next) are forwarded unchanged.
//   - The error is forwarded unchanged after the side-effect executes.
//   - Completion is forwarded normally.
//
// The provided function `fn` is executed exactly once per subscription,
// at the moment the error is observed.
//
// Error semantics:
//   - The side-effect runs BEFORE the error is propagated downstream.
//   - If the side-effect panics, the panic propagates and terminates
//     the subscription.
//
// Unsubscription:
//   - Unsubscribing before an error prevents the side-effect from running.
//
// DoOnError is typically used for logging, metrics, tracing, or debugging,
// without affecting the behavior of the Observable pipeline.
//
// Diagram:
//
//	source:   --1--2--X
//	doOnErr: --1--2--X
//	               ^
//	               fn(error) is called here
//
// Example:
//
//	src := obs.Throw[int](errors.New("failure"))
//
//	out := 	DoOnError(func(err error) {
//		fmt.Println("observed error:", err)
//	})(src)
//
//	out.Sub(obs.NewObserver(
//		func(v int) {
//			fmt.Println("value:", v)
//		},
//		func(err error) {
//			fmt.Println("downstream error:", err)
//		},
//		func() {
//			fmt.Println("done")
//		},
//	))
//
// Output:
//
//	observed error: failure
//	downstream error: failure
//
// Notes:
//
//   - DoOnError is for side effects only.
//   - It must not be used to recover from errors.
//   - For error recovery or retry logic, use Catch, Retry, or RetryWhen.
func DoOnError[T any](fn func(error)) obs.TransformPipe[T, T] {
	return func(src obs.Observable[T]) obs.Observable[T] {
		return obs.Create[T](func(o *obs.Observer[T]) (obs.Subscription, error) {

			var (
				sub  obs.Subscription
				done atomic.Bool
			)

			s, err := src.Sub(&obs.Observer[T]{
				Next: func(v T) {
					if done.Load() {
						return
					}
					o.Next(v)
				},

				Error: func(err error) {
					if done.CompareAndSwap(false, true) {
						fn(err)
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

			sub = s

			return subscr.Func(func() {
				if done.CompareAndSwap(false, true) {
					sub.Unsub()
				}
			}), nil
		})
	}
}

// Finalize registers a cleanup function that is executed exactly once
// when the Observable terminates, regardless of how termination occurs.
//
// The cleanup function is invoked when the downstream subscription
// is terminated by:
//   - normal completion,
//   - an error,
//   - or explicit unsubscription.
//
// Behavior:
//   - Values emitted by the source are forwarded unchanged.
//   - Errors are propagated downstream after cleanup is scheduled.
//   - Completion is forwarded after cleanup is scheduled.
//   - The cleanup function is guaranteed to run at most once.
//
// Finalize is typically used to release resources, stop timers,
// close channels, or perform side effects that must occur on teardown.
//
// Diagram:
//
//	source:    --1--2--3--|
//	finalize: --1--2--3--|
//	                 cleanup()
//
//	source:    --1--2--#
//	finalize: --1--2--#
//	                 cleanup()
//
//	source:    --1--2--3--X
//	finalize: --1--2--3
//	                 cleanup()
//
// Example:
//
//	src := obs.SlotSubject[int]()
//
//	result := Finalize(func() {
//		fmt.Println("cleanup called")
//	})(src)
//
//	sub, _ := result.Sub(obs.NewObserver(
//		func(v int) {
//			fmt.Println(v)
//		},
//		func(err error) {
//			fmt.Println("error")
//		},
//		func() {
//			fmt.Println("done")
//		},
//	))
//
//	src.Next(1)
//	src.Next(2)
//	sub.Unsub()
//
// Output:
//
//	1
//	2
//	cleanup called
//
// Notes:
//
//   - Finalize is equivalent to a "finally" or "defer" semantics
//     for Observables.
//   - The cleanup function must be side-effect only and should
//     not emit values.
//   - Finalize does not alter the timing or ordering of emissions.
func Finalize[T any](cleanup func()) obs.TransformPipe[T, T] {
	return func(src obs.Observable[T]) obs.Observable[T] {
		return obs.Create[T](func(o *obs.Observer[T]) (obs.Subscription, error) {
			var once sync.Once
			run := func() {
				once.Do(cleanup)
			}

			sub, err := src.Sub(&obs.Observer[T]{
				Next: o.Next,

				Error: func(err error) {
					run()
					o.Error(err)
				},

				Complete: func() {
					run()
					o.Complete()
				},
			})

			if err != nil {
				run()
				return nil, err
			}

			return subscr.Func(func() {
				sub.Unsub()
				run()
			}), nil
		})
	}
}

// Tap performs a side effect for each value emitted by the source Observable
// without modifying, filtering, or delaying the stream.
//
// Tap is a transparent operator: it forwards all values, errors, and
// completion signals exactly as received from the source.
//
// The provided function `fn` is called synchronously for each value
// before the value is forwarded downstream.
//
// Typical use cases include:
//   - Logging
//   - Debugging
//   - Metrics collection
//   - Tracing
//
// Tap must NOT be used to:
//   - Modify values
//   - Control flow
//   - Trigger retries or cancellations
//
// Diagram:
//
//	source:  --1--2--3--|
//	tap(fn)
//	output:  --1--2--3--|
//
// Side effects:
//
//	fn(1)
//	fn(2)
//	fn(3)
//
// Error handling:
//
//   - Errors from the source are forwarded downstream unchanged.
//   - The Tap function is NOT called for error or completion signals.
//
// Completion:
//
//   - Completion is forwarded downstream unchanged.
//
// Example:
//
//	src := obs.Of(1, 2, 3).Pipe(
//		Tap(func(v int) {
//			fmt.Println("seen:", v)
//		}),
//	)
//
//	src.Sub(obs.NewObserver(
//		func(v int) {
//			fmt.Println("next:", v)
//		},
//		func(err error) {},
//		func() {
//			fmt.Println("done")
//		},
//	))
//
// Output:
//
//	seen: 1
//	next: 1
//	seen: 2
//	next: 2
//	seen: 3
//	next: 3
//	done
//
// Notes:
//
//   - Tap does not introduce concurrency or buffering.
//   - The Tap function executes in the same call stack as the source emission.
//   - Panics inside `fn` will propagate and may terminate the stream.
func Tap[T any](fn func(T)) obs.TransformPipe[T, T] {
	return func(src obs.Observable[T]) obs.Observable[T] {
		return obs.Create[T](func(o *obs.Observer[T]) (obs.Subscription, error) {
			return src.Sub(&obs.Observer[T]{
				Next: func(v T) {
					fn(v)
					o.Next(v)
				},
				Error:    o.Error,
				Complete: o.Complete,
			})
		})
	}
}
