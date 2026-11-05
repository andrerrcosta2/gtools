// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package transform

import (
	"sync/atomic"

	"github.com/andrerrcosta2/gtools/gflux/obs"
	"github.com/andrerrcosta2/gtools/gflux/obs/internal/subscr"
)

// Catch intercepts an error from the source Observable and replaces it
// with a fallback Observable returned by the provided function.
//
// At the source Observable emits an error, Catch:
//
//  1. Cancels the source subscription.
//  2. Calls the provided function with the error.
//  3. Subscribes to the returned fallback Observable.
//  4. Forwards all values, errors, and completion from the fallback.
//
// Catch switches to the fallback at most once.
// If the fallback Observable emits an error, that error is propagated
// downstream and the stream terminates.
//
// If the source Observable completes normally, Catch has no effect
// and the completion is forwarded downstream.
//
// Unsubscribing from the resulting Observable cancels both the
// source subscription (if active) and the fallback subscription
// (if active).
//
// Catch is useful for error recovery, fallback strategies,
// and graceful degradation of streams.
//
// Example:
//
//	src := obs.SlotSubject[int]()
//
//	recovered := Catch(func(err error) obs.Observable[int] {
//		// Replace error with a fallback sequence
//		return obs.Of(100, 200, 300)
//	})(src)
//
//	recovered.Sub(obs.NewObserver(
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
//	src.Error(errors.New("failure"))
//
// Output:
//
//	1
//	2
//	100
//	200
//	300
//	done
func Catch[T any](fn func(err error) obs.Observable[T]) obs.TransformPipe[T, T] {
	return func(src obs.Observable[T]) obs.Observable[T] {
		return obs.Create[T](func(o *obs.Observer[T]) (obs.Subscription, error) {
			var (
				srcSub     obs.Subscription
				innerSub   obs.Subscription
				terminated atomic.Bool
				switched   atomic.Bool
			)

			terminate := func() bool {
				return terminated.CompareAndSwap(false, true)
			}

			sub, err := src.Sub(&obs.Observer[T]{
				Next: func(v T) {
					if terminated.Load() {
						return
					}
					o.Next(v)
				},

				Error: func(err error) {
					// Ignore all errors after termination
					if terminated.Load() {
						return
					}

					// First error: switch
					if switched.CompareAndSwap(false, true) {
						if srcSub != nil {
							srcSub.Unsub()
						}

						fallback := fn(err)
						is, err2 := fallback.Sub(&obs.Observer[T]{
							Next: func(v T) {
								if terminated.Load() {
									return
								}
								o.Next(v)
							},
							Error: func(e error) {
								if terminate() {
									o.Error(e)
								}
							},
							Complete: func() {
								if terminate() {
									o.Complete()
								}
							},
						})

						if err2 != nil {
							if terminate() {
								o.Error(err2)
							}
							return
						}

						innerSub = is
						return
					}

					// Subsequent source errors are ignored
				},

				Complete: func() {
					if terminate() {
						o.Complete()
					}
				},
			})

			if err != nil {
				return nil, err
			}

			srcSub = sub

			return subscr.Func(func() {
				terminated.Store(true)
				if srcSub != nil {
					srcSub.Unsub()
				}
				if innerSub != nil {
					innerSub.Unsub()
				}
			}), nil
		})
	}
}

// Repeat resubscribes to the source Observable each time it completes,
// repeating the sequence up to the specified count.
//
// The source Observable is subscribed to immediately. At it completes,
// Repeat resubscribes again, until the total number of completed runs
// reaches count + 1 (the initial subscription plus count repeats).
//
// If the source Observable emits an error, the error is forwarded
// downstream and repetition stops immediately.
//
// Repeat does not buffer or replay values itself; it simply re-executes
// the source Observable by resubscribing.
//
// Unsubscribing from the resulting Observable stops any further
// resubscriptions and unsubscribes from the current source subscription.
//
// Example:
//
//	src := obs.Of(1, 2)
//
//	repeated := Repeat(2)(src)
//
//	repeated.Sub(obs.NewObserver(
//		func(v int) {
//			fmt.Println(v)
//		},
//		func(err error) {},
//		func() {
//			fmt.Println("done")
//		},
//	))
//
// Output:
//
//	1
//	2
//	1
//	2
//	1
//	2
//	done
func Repeat[T any](count int) obs.TransformPipe[T, T] {
	if count < 0 {
		panic("repeat can't be less than zero")
	}
	return func(src obs.Observable[T]) obs.Observable[T] {
		return obs.Create[T](func(o *obs.Observer[T]) (obs.Subscription, error) {

			var (
				done   int
				srcSub obs.Subscription
				closed atomic.Bool
			)

			var subscribe func()

			subscribe = func() {
				if closed.Load() {
					return
				}

				sub, err := src.Sub(&obs.Observer[T]{
					Next: o.Next,

					Error: func(err error) {
						o.Error(err)
					},

					Complete: func() {
						done++
						if done <= count {
							subscribe() // repeat
							return
						}
						o.Complete()
					},
				})

				if err != nil {
					o.Error(err)
					return
				}

				srcSub = sub
			}

			subscribe()

			return subscr.Func(func() {
				closed.Store(true)
				if srcSub != nil {
					srcSub.Unsub()
				}
			}), nil
		})
	}
}

// Retry resubscribes to the source Observable when an error occurs,
// up to the specified maximum number of retries.
//
// The source Observable is subscribed to immediately. If it emits an
// error, Retry will resubscribe and restart the sequence, until the
// number of retries exceeds maxRetries.
//
// If the source Observable completes successfully, Retry completes
// immediately and no further retries occur.
//
// If an error occurs after the maximum number of retries has been
// exhausted, the error is forwarded downstream.
//
// Unsubscribing from the resulting Observable stops any further retry
// attempts and unsubscribes from the current source subscription.
//
// Diagram (maxRetries = 2):
//
//	source:   --1--2--X
//	             |
//	retry #1:     --1--2--X
//	                     |
//	retry #2:             --1--2--|
//
//	output:   --1--2--1--2--1--2--|
//
// Diagram (maxRetries = 1, failure persists):
//
//	source:   --1--2--X
//	             |
//	retry #1:     --1--2--X
//
//	output:   --1--2--1--2--X
//
// Example:
//
//	src := obs.ObservableThatFailsSometimes()
//
//	retried := Retry(src)
//
//	retried.Sub(obs.NewObserver(
//		func(v int) {
//			fmt.Println(v)
//		},
//		func(err error) {
//			fmt.Println("failed:", err)
//		},
//		func() {
//			fmt.Println("done")
//		},
//	))
//
// Semantics:
//
//	maxRetries = 0 → no retries, fail immediately
//	maxRetries = 1 → one retry after first failure
//	maxRetries = N → up to N retries on error
//
// Notes:
//
//   - Retry restarts the source sequence from the beginning on each retry.
//   - Values emitted before an error are not buffered or replayed;
//     they are simply re-emitted by the new subscription.
//   - Retry is a stateful operator.
func Retry[T any](maxRetries int) obs.TransformPipe[T, T] {
	if maxRetries < 0 {
		panic("retry maxRetries must be >= 0")
	}

	return func(src obs.Observable[T]) obs.Observable[T] {
		return obs.Create[T](func(o *obs.Observer[T]) (obs.Subscription, error) {
			var (
				attempt int
				srcSub  obs.Subscription
				term    = &termination{}
			)

			var subscribe func()

			subscribe = func() {
				// No work after unsubscribe or termination
				if term.IsClosedOrTerminated() {
					return
				}

				sub, err := src.Sub(&obs.Observer[T]{
					Next: func(v T) {
						if term.IsClosedOrTerminated() {
							return
						}
						o.Next(v)
					},

					Error: func(err error) {
						if term.IsClosedOrTerminated() {
							return
						}

						attempt++

						if attempt <= maxRetries {
							// cancel current attempt before retry
							if srcSub != nil {
								srcSub.Unsub()
							}
							subscribe()
							return
						}

						// terminal error (once)
						if term.TerminateIfActive() {
							o.Error(err)
						}
					},

					Complete: func() {
						if term.TerminateIfActive() {
							o.Complete()
						}
					},
				})

				if err != nil {
					if term.TerminateIfActive() {
						o.Error(err)
					}
					return
				}

				srcSub = sub
			}

			// First attempt
			subscribe()

			return subscr.Func(func() {
				if term.CloseIfOpen() {
					if srcSub != nil {
						srcSub.Unsub()
					}
				}
			}), nil
		})
	}
}

// RetryWhen retries the source Observable when an error occurs,
// based on a control Observable derived from the source error.
//
// At the source errors:
//
//   - The source subscription is cancelled
//   - The function `fn` is called with the source error
//   - The returned control Observable is subscribed
//
// Control Observable behavior:
//
//   - Next     → triggers a retry (resubscribe to the source)
//   - Error    → propagates the control error downstream
//   - Complete → propagates the ORIGINAL source error downstream
//
// Only one retry is triggered per control emission.
// Unsubscribing cancels both source and control subscriptions.
//
// Diagram:
//
//	source:        --1--2--X
//	                   |
//	control(err):       --•----|
//	                         |
//	retry:                   --1--2--|
//
//	output:        --1--2--1--2--|
//
// Diagram (control completes → original error):
//
//	source:        --1--X
//	                 |
//	control(err):      --|
//
//	output:        --1--X
//
// Diagram (control errors):
//
//	source:        --1--X
//	                 |
//	control(err):      --Y
//
//	output:        --1--Y
//
// Example:
//
//	attempts := 0
//
//	src := obs.Defer(func() obs.Observable[int] {
//		attempts++
//		if attempts < 3 {
//			return obs.Throw[int](errors.New("fail"))
//		}
//		return obs.Of(42)
//	}).Pipe(
//
//		RetryWhen(func(err error) obs.Observable[any] {
//			// Retry after a signal
//			return obs.Of(nil)
//		}),
//	)
//
//	src.Sub(obs.NewObserver(
//		func(v int) { fmt.Println(v) },
//		func(err error) { fmt.Println("err:", err) },
//		func() { fmt.Println("done") },
//	))
//
// Output:
//
//	42
//	done
//
// Notes:
//
//   - The control Observable does NOT receive values from the source,
//     only the source error.
//   - RetryWhen does not automatically retry; retries occur only
//     when the control Observable emits.
//   - Completion of the control Observable re-emits the original
//     source error, not a new one.
//   - RetryWhen is a stateful operator.
func RetryWhen[T any](fn func(error) obs.Observable[any]) obs.TransformPipe[T, T] {
	return func(src obs.Observable[T]) obs.Observable[T] {
		return obs.Create[T](func(o *obs.Observer[T]) (obs.Subscription, error) {
			var (
				srcSub  obs.Subscription
				ctrlSub obs.Subscription
				term    = &termination{}
				retried atomic.Bool
			)

			var subscribe func()

			// this safe error ensures only one terminal signal reaches downstream
			safeError := func(err error) {
				if term.TerminateIfActive() {
					o.Error(err)
				}
			}

			safeComplete := func() {
				if term.TerminateIfActive() {
					o.Complete()
				}
			}

			// INVARIANTS:
			// 1) At most, one source subscription is active at any time
			// 2) Retries occur only on control Next emissions
			// 3) Control completion propagates the original source error
			// 4) Unsubscription is terminal and cancels source and control
			subscribe = func() {
				if term.IsClosedOrTerminated() {
					return
				}

				sub, err := src.Sub(&obs.Observer[T]{
					Next: func(v T) {
						if !term.IsTerminated() {
							o.Next(v)
						}
					},

					Error: func(srcErr error) {
						if term.IsClosedOrTerminated() {
							return
						}

						if srcSub != nil {
							srcSub.Unsub()
						}
						retried.Store(false)

						ctrl := fn(srcErr)

						cs, _ := ctrl.Sub(&obs.Observer[any]{
							Next: func(_ any) {
								if term.IsClosedOrTerminated() {
									return
								}
								if ctrlSub != nil {
									ctrlSub.Unsub()
								}
								retried.Store(true)
								subscribe()
							},

							// Control error: propagate control error
							Error: func(err error) {
								safeError(err)
							},

							// Control complete> propagate ORIGINAL source error
							Complete: func() {
								if retried.Load() {
									return
								}
								safeError(srcErr)
							},
						})

						ctrlSub = cs
					},

					Complete: func() {
						safeComplete()
					},
				})

				if err != nil {
					safeError(err)
					return
				}

				srcSub = sub
			}

			subscribe()

			return subscr.Func(func() {
				term.TerminateAndClose()
				if srcSub != nil {
					srcSub.Unsub()
				}
				if ctrlSub != nil {
					ctrlSub.Unsub()
				}
			}), nil
		})
	}
}
