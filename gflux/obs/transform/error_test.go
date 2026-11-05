// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package transform

import (
	"errors"
	"sync/atomic"
	"testing"
	"time"

	"github.com/andrerrcosta2/gtools/core/testlite/assertlite"
	"github.com/andrerrcosta2/gtools/gflux/obs"
	"github.com/andrerrcosta2/gtools/gflux/obs/internal/observers"
	"github.com/andrerrcosta2/gtools/gflux/obs/internal/subscr"
)

// TestCatch_Sync tests Catch using synchronous sources.
//
// For synchronous sources, we must assert:
//
//  1. Values before error are forwarded
//  2. Error triggers fallback subscription
//  3. Fallback values are forwarded
//  4. Completion comes from fallback
//  5. If source completes normally, fallback is never used
//  6. Error in fallback is propagated
//  7. Catch switches at most once
//  8. Completes synchronously when possible
func TestCatch_Sync(t *testing.T) {
	t.Run("recovers from sync error with fallback", func(t *testing.T) {
		failure := errors.New("failure")

		src := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			o.Next(1)
			o.Next(2)
			o.Error(failure)
			return subscr.Noop{}, nil
		})

		recovered := Catch(func(err error) obs.Observable[int] {
			assertlite.IsErrorOf(t, err, failure)
			return obs.Of(100, 200)
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = recovered.Sub((*obs.Observer[int])(o))

		assertlite.EqualSlices(t, []int{1, 2, 100, 200}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("does nothing if source completes normally", func(t *testing.T) {
		src := obs.Of(1, 2, 3)

		recovered := Catch(func(err error) obs.Observable[int] {
			t.Fatalf("fallback must not be called")
			return obs.Of[int]()
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = recovered.Sub((*obs.Observer[int])(o))

		assertlite.EqualSlices(t, []int{1, 2, 3}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("fallback error is propagated", func(t *testing.T) {
		failure := errors.New("failure")
		fallbackErr := errors.New("fallback failed")

		src := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			o.Next(1)
			o.Error(failure)
			return subscr.Noop{}, nil
		})

		recovered := Catch(func(err error) obs.Observable[int] {
			assertlite.IsErrorOf(t, err, failure)
			return obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
				o.Error(fallbackErr)
				return subscr.Noop{}, nil
			})
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = recovered.Sub((*obs.Observer[int])(o))

		assertlite.EqualSlices(t, []int{1}, rec.Values())
		assertlite.IsErrorOf(t, rec.Err(), fallbackErr)
		assertlite.False(t, rec.Completed())
	})

	t.Run("switches to fallback only once", func(t *testing.T) {
		failure := errors.New("failure")
		calls := 0

		src := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			o.Error(failure)
			o.Error(errors.New("ignored"))
			return subscr.Noop{}, nil
		})

		recovered := Catch(func(err error) obs.Observable[int] {
			calls++
			return obs.Of(1)
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = recovered.Sub((*obs.Observer[int])(o))

		assertlite.EqualSlices(t, []int{1}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.Equals(t, 1, calls)
	})

	t.Run("completes synchronously when fallback completes synchronously", func(t *testing.T) {
		failure := errors.New("failure")

		src := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			o.Error(failure)
			return subscr.Noop{}, nil
		})

		done := false

		recovered := Catch(func(err error) obs.Observable[int] {
			return obs.Of(42)
		})(src)

		o := &obs.Observer[int]{
			Next: func(int) {},
			Complete: func() {
				done = true
			},
		}
		_, _ = recovered.Sub(o)
		assertlite.True(t, done)
	})
}

// TestCatch_Async tests Catch using asynchronous sources and/or fallbacks.
//
// We must assert:
//
//  1. Async source values before error are forwarded
//  2. Async error triggers fallback subscription
//  3. Async fallback values are forwarded
//  4. Completion comes from fallback
//  5. If async source completes normally, fallback is never used
//  6. Error in async fallback is propagated
//  7. Catch switches at most once
//  8. Unsubscribe cancels active source or fallback
func TestCatch_Async(t *testing.T) {
	t.Run("recovers from async error with async fallback", func(t *testing.T) {
		failure := errors.New("failure")

		src := obs.Async[int](func(o *obs.Observer[int]) {
			o.Next(1)
			o.Next(2)
			o.Error(failure)
		})

		recovered := Catch(func(err error) obs.Observable[int] {
			assertlite.IsErrorOf(t, err, failure)
			return obs.AsyncOf(100, 200)
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = recovered.Sub((*obs.Observer[int])(o))

		rec.AwaitTermination(200 * time.Millisecond)

		assertlite.EqualSlices(t, []int{1, 2, 100, 200}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("does nothing if async source completes normally", func(t *testing.T) {
		src := obs.AsyncOf(1, 2, 3)

		recovered := Catch(func(err error) obs.Observable[int] {
			t.Fatalf("fallback must not be called")
			return obs.Of[int]()
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = recovered.Sub((*obs.Observer[int])(o))

		rec.AwaitCompleted(200 * time.Millisecond)

		assertlite.EqualSlices(t, []int{1, 2, 3}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("async fallback error is propagated", func(t *testing.T) {
		failure := errors.New("failure")
		fallbackErr := errors.New("fallback failed")

		src := obs.Async[int](func(o *obs.Observer[int]) {
			o.Next(1)
			o.Error(failure)
		})

		recovered := Catch(func(err error) obs.Observable[int] {
			return obs.Async[int](func(o *obs.Observer[int]) {
				o.Error(fallbackErr)
			})
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = recovered.Sub((*obs.Observer[int])(o))

		err := rec.AwaitError(200 * time.Millisecond)

		assertlite.EqualSlices(t, []int{1}, rec.Values())
		assertlite.IsErrorOf(t, err, fallbackErr)
		assertlite.False(t, rec.Completed())
	})

	t.Run("switches to fallback only once for async source", func(t *testing.T) {
		failure := errors.New("failure")
		calls := atomic.Int32{}

		src := obs.Async[int](func(o *obs.Observer[int]) {
			o.Error(failure)
			o.Error(errors.New("ignored"))
		})

		recovered := Catch(func(err error) obs.Observable[int] {
			calls.Add(1)
			return obs.AsyncOf(42)
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = recovered.Sub((*obs.Observer[int])(o))

		rec.AwaitCompleted(200 * time.Millisecond)

		assertlite.EqualSlices(t, []int{42}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.Equals(t, int32(1), calls.Load())
	})

	t.Run("unsubscribe cancels async fallback", func(t *testing.T) {
		failure := errors.New("failure")

		// Source errors asynchronously
		src := obs.Async[int](func(o *obs.Observer[int]) {
			o.Error(failure)
		})

		first := make(chan struct{})
		second := make(chan struct{})

		recovered := Catch(func(err error) obs.Observable[int] {
			return obs.Async[int](func(o *obs.Observer[int]) {
				// First blocking point
				<-first
				o.Next(1)

				// Second blocking point
				<-second
				o.Next(2) // must NOT be observed
			})
		})(src)

		rec, o := observers.NewRecorder[int](t)
		sub, err := recovered.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		// Allow first emission
		close(first)
		rec.AwaitNext(200 * time.Millisecond)

		// Unsubscribe before second emission
		sub.Unsub()

		// Unblock second emission attempt
		close(second)

		// Give async goroutine time to attempt emission
		time.Sleep(100 * time.Millisecond)

		// Assertions
		assertlite.EqualSlices(t, []int{1}, rec.Values())
		assertlite.False(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})
}

// TestRepeat_Sync tests the Repeat pipe using synchronous sources.
//
// For synchronous sources, we must assert:
//  1. Repeats the source sequence the correct number of times
//  2. count = 0 results in a single execution
//  3. Works with empty sources
//  4. Errors stop repetition immediately
//  5. Completes synchronously
//  6. Does not emit after completion
func TestRepeat_Sync(t *testing.T) {
	t.Run("repeats source sequence count times", func(t *testing.T) {
		src := obs.Of(1, 2)

		repeated := Repeat[int](2)(src) // initial + 2 repeats = 3 runs

		rec, o := observers.NewRecorder[int](t)
		_, err := repeated.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		assertlite.EqualSlices(t,
			[]int{1, 2, 1, 2, 1, 2},
			rec.Values(),
		)
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("count zero executes source only once", func(t *testing.T) {
		src := obs.Of(1, 2, 3)

		repeated := Repeat[int](0)(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = repeated.Sub((*obs.Observer[int])(o))

		assertlite.EqualSlices(t, []int{1, 2, 3}, rec.Values())
		assertlite.True(t, rec.Completed())
	})

	t.Run("works with single value source", func(t *testing.T) {
		src := obs.Of(42)

		repeated := Repeat[int](3)(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = repeated.Sub((*obs.Observer[int])(o))

		assertlite.EqualSlices(t, []int{42, 42, 42, 42}, rec.Values())
		assertlite.True(t, rec.Completed())
	})

	t.Run("empty source repeats empty and completes", func(t *testing.T) {
		src := obs.Of[int]()

		repeated := Repeat[int](5)(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = repeated.Sub((*obs.Observer[int])(o))

		assertlite.Equals(t, 0, len(rec.Values()))
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("error stops repetition immediately", func(t *testing.T) {
		failure := errors.New("failure")

		runs := 0
		src := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			runs++
			o.Next(runs)
			o.Error(failure)
			return subscr.Noop{}, nil
		})

		repeated := Repeat[int](5)(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = repeated.Sub((*obs.Observer[int])(o))

		assertlite.EqualSlices(t, []int{1}, rec.Values())
		assertlite.IsErrorOf(t, rec.Err(), failure)
		assertlite.False(t, rec.Completed())
	})

	t.Run("completes synchronously", func(t *testing.T) {
		src := obs.Of(1)

		repeated := Repeat[int](2)(src)

		done := false
		o := &obs.Observer[int]{
			Next: func(int) {},
			Complete: func() {
				done = true
			},
		}

		_, _ = repeated.Sub(o)

		assertlite.True(t, done)
	})

	t.Run("sync reentrancy", func(t *testing.T) {
		var (
			subscribeDepth int
			maxDepth       int
		)

		src := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			subscribeDepth++
			if subscribeDepth > maxDepth {
				maxDepth = subscribeDepth
			}

			o.Next(subscribeDepth)
			o.Complete()

			subscribeDepth--
			return subscr.Noop{}, nil
		})

		repeated := Repeat[int](3)(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := repeated.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		// 1 initial + 3 repeats = 4 total
		assertlite.EqualSlices(t, []int{1, 2, 3, 4}, rec.Values())
		assertlite.True(t, rec.Completed())

		// Reentrancy must have occurred
		assertlite.True(t, maxDepth > 1)
	})
}

// TestRepeat_Async tests the Repeat pipe using asynchronous sources.
//
// For asynchronous sources, we must assert:
//  1. repeats the source sequence the correct number of times
//  2. count = 0 results in a single execution
//  3. works with empty async sources
//  4. errors stop repetition immediately
//  5. completes asynchronously
//  6. does not emit after completion
//  7. unsubscribe stops further repetition
func TestRepeat_Async(t *testing.T) {
	t.Run("repeats async source sequence count times", func(t *testing.T) {
		src := obs.AsyncOf(1, 2)

		repeated := Repeat[int](2)(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := repeated.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		rec.AwaitTermination(100 * time.Millisecond)

		assertlite.EqualSlices(t,
			[]int{1, 2, 1, 2, 1, 2},
			rec.Values(),
		)
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("count zero executes async source only once", func(t *testing.T) {
		src := obs.AsyncOf(1, 2, 3)

		repeated := Repeat[int](0)(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = repeated.Sub((*obs.Observer[int])(o))

		rec.AwaitTermination(100 * time.Millisecond)

		assertlite.EqualSlices(t, []int{1, 2, 3}, rec.Values())
		assertlite.True(t, rec.Completed())
	})

	t.Run("works with single value async source", func(t *testing.T) {
		src := obs.AsyncOf(42)

		repeated := Repeat[int](3)(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = repeated.Sub((*obs.Observer[int])(o))

		rec.AwaitTermination(100 * time.Millisecond)

		assertlite.EqualSlices(t, []int{42, 42, 42, 42}, rec.Values())
		assertlite.True(t, rec.Completed())
	})

	t.Run("empty async source repeats empty and completes", func(t *testing.T) {
		src := obs.AsyncOf[int]()

		repeated := Repeat[int](2)(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = repeated.Sub((*obs.Observer[int])(o))

		rec.AwaitTermination(100 * time.Millisecond)

		assertlite.Equals(t, 0, len(rec.Values()))
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("async error stops repetition immediately", func(t *testing.T) {
		failure := errors.New("failure")

		runs := 0
		src := obs.Async[int](func(o *obs.Observer[int]) {
			runs++
			o.Next(runs)
			o.Error(failure)
		})

		repeated := Repeat[int](5)(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = repeated.Sub((*obs.Observer[int])(o))

		rec.AwaitTermination(100 * time.Millisecond)

		assertlite.EqualSlices(t, []int{1}, rec.Values())
		assertlite.IsErrorOf(t, rec.Err(), failure)
		assertlite.False(t, rec.Completed())
	})

	t.Run("completes asynchronously", func(t *testing.T) {
		src := obs.AsyncOf(1)

		repeated := Repeat[int](2)(src)

		completed := false
		done := make(chan struct{})

		o := &obs.Observer[int]{
			Next: func(int) {},
			Complete: func() {
				completed = true
				close(done)
			},
		}

		_, _ = repeated.Sub(o)

		// async completion must not happen immediately
		assertlite.False(t, completed)

		<-done
		assertlite.True(t, completed)
	})

	t.Run("unsubscribe stops further repetition", func(t *testing.T) {
		started := make(chan struct{})
		block := make(chan struct{})

		src := obs.Async[int](func(o *obs.Observer[int]) {
			close(started)
			o.Next(1)
			<-block
			o.Complete()
		})

		repeated := Repeat[int](5)(src)

		rec, o := observers.NewRecorder[int](t)
		sub, _ := repeated.Sub((*obs.Observer[int])(o))

		<-started
		sub.Unsub()
		close(block)

		// Allow time for any illegal emission
		time.Sleep(100 * time.Millisecond)

		assertlite.EqualSlices(t, []int{1}, rec.Values())
		assertlite.False(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})
}

// TestRetry_Sync tests the Retry pipe using synchronous sources.
//
// For synchronous sources, we must assert:
//  1. Retries the source sequence up to maxRetries times on error
//  2. maxRetries = 0 fails immediately
//  3. Completes immediately if source completes successfully
//  4. Errors after retries are exhausted
//  5. Values are re-emitted from the beginning on each retry
//  6. Does not emit after terminal error
//  7. Supports synchronous reentrancy
func TestRetry_Sync(t *testing.T) {
	t.Run("retries source up to maxRetries and then succeeds", func(t *testing.T) {
		var runs int

		src := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			runs++
			o.Next(1)
			o.Next(2)

			if runs < 3 {
				o.Error(errors.New("fail"))
				return subscr.Noop{}, nil
			}

			o.Complete()
			return subscr.Noop{}, nil
		})

		retried := Retry[int](2)(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := retried.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		// 2 failures + 1 success
		assertlite.EqualSlices(t,
			[]int{1, 2, 1, 2, 1, 2},
			rec.Values(),
		)
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("maxRetries zero fails immediately", func(t *testing.T) {
		failure := errors.New("failure")

		src := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			o.Next(1)
			o.Error(failure)
			return subscr.Noop{}, nil
		})

		retried := Retry[int](0)(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = retried.Sub((*obs.Observer[int])(o))

		assertlite.EqualSlices(t, []int{1}, rec.Values())
		assertlite.IsErrorOf(t, rec.Err(), failure)
		assertlite.False(t, rec.Completed())
	})

	t.Run("successful completion stops retries", func(t *testing.T) {
		var runs int

		src := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			runs++
			o.Next(runs)
			o.Complete()
			return subscr.Noop{}, nil
		})

		retried := Retry[int](5)(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = retried.Sub((*obs.Observer[int])(o))

		assertlite.EqualSlices(t, []int{1}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("error after retries exhausted is forwarded", func(t *testing.T) {
		failure := errors.New("persistent failure")
		attempts := 0

		src := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			attempts++
			o.Next(attempts)
			o.Error(failure)
			return subscr.Noop{}, nil
		})

		retried := Retry[int](2)(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = retried.Sub((*obs.Observer[int])(o))

		// initial + 2 retries = 3 attempts
		assertlite.EqualSlices(t, []int{1, 2, 3}, rec.Values())
		assertlite.IsErrorOf(t, rec.Err(), failure)
		assertlite.False(t, rec.Completed())
	})

	t.Run("does not emit after terminal error", func(t *testing.T) {
		failure := errors.New("fail")

		src := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			o.Next(1)
			o.Error(failure)
			o.Next(999) // illegal
			return subscr.Noop{}, nil
		})

		retried := Retry[int](0)(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = retried.Sub((*obs.Observer[int])(o))

		assertlite.EqualSlices(t, []int{1}, rec.Values())
		assertlite.IsErrorOf(t, rec.Err(), failure)
	})

	t.Run("sync reentrancy", func(t *testing.T) {
		var (
			subscribeDepth int
			maxDepth       int
			failure        = errors.New("fail")
		)

		src := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			subscribeDepth++
			if subscribeDepth > maxDepth {
				maxDepth = subscribeDepth
			}

			o.Next(subscribeDepth)
			o.Error(failure)

			subscribeDepth--
			return subscr.Noop{}, nil
		})

		retried := Retry[int](3)(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = retried.Sub((*obs.Observer[int])(o))

		// initial + 3 retries = 4 attempts
		assertlite.EqualSlices(t, []int{1, 2, 3, 4}, rec.Values())
		assertlite.IsErrorOf(t, failure, rec.Err())
		assertlite.False(t, rec.Completed())

		// Reentrancy must have occurred
		assertlite.True(t, maxDepth > 1)
	})
}

// TestRetry_Async tests the Retry pipe using asynchronous sources.
//
// For asynchronous sources, we must assert:
//  1. Retries occur on async error
//  2. maxRetries = 0 fails immediately
//  3. Successful completion stops retries
//  4. Error after retries exhausted is forwarded
//  5. Values are re-emitted from the beginning on each retry
//  6. Unsubscribe stops further retries
//  7. No emissions after termination
func TestRetry_Async(t *testing.T) {
	t.Run("retries async source up to maxRetries and then succeeds", func(t *testing.T) {
		var runs atomic.Int32

		src := obs.Async[int](func(o *obs.Observer[int]) {
			n := runs.Add(1)
			o.Next(int(n))

			if n < 3 {
				o.Error(errors.New("fail"))
				return
			}

			o.Complete()
		})

		// 2 is enough for 3 but it should complete on second retry anyway
		retried := Retry[int](3)(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := retried.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		rec.AwaitCompleted(1 * time.Second)

		assertlite.EqualSlices(t, []int{1, 2, 3}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("maxRetries zero fails immediately", func(t *testing.T) {
		failure := errors.New("failure")

		src := obs.Async[int](func(o *obs.Observer[int]) {
			o.Next(1)
			o.Error(failure)
		})

		retried := Retry[int](0)(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = retried.Sub((*obs.Observer[int])(o))

		_ = rec.AwaitError(1 * time.Second)

		assertlite.EqualSlices(t, []int{1}, rec.Values())
		assertlite.IsErrorOf(t, rec.Err(), failure)
		assertlite.False(t, rec.Completed())
	})

	t.Run("successful async completion stops retries", func(t *testing.T) {
		var runs atomic.Int32

		src := obs.Async[int](func(o *obs.Observer[int]) {
			o.Next(int(runs.Add(1)))
			o.Complete()
		})

		retried := Retry[int](5)(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = retried.Sub((*obs.Observer[int])(o))

		rec.AwaitCompleted(1 * time.Second)

		assertlite.EqualSlices(t, []int{1}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("error after retries exhausted is forwarded", func(t *testing.T) {
		failure := errors.New("persistent failure")
		var attempts atomic.Int32

		src := obs.Async[int](func(o *obs.Observer[int]) {
			o.Next(int(attempts.Add(1)))
			o.Error(failure)
		})

		retried := Retry[int](2)(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = retried.Sub((*obs.Observer[int])(o))

		_ = rec.AwaitError(1 * time.Second)

		// initial + 2 retries = 3 attempts
		assertlite.EqualSlices(t, []int{1, 2, 3}, rec.Values())
		assertlite.IsErrorOf(t, rec.Err(), failure)
		assertlite.False(t, rec.Completed())
	})

	t.Run("unsubscribe stops further retries", func(t *testing.T) {
		started := make(chan struct{})
		block := make(chan struct{})

		var attempts atomic.Int32

		src := obs.Async[int](func(o *obs.Observer[int]) {
			attempts.Add(1)
			close(started)
			o.Next(1)
			<-block
			o.Error(errors.New("fail"))
		})

		retried := Retry[int](3)(src)

		rec, o := observers.NewRecorder[int](t)
		sub, _ := retried.Sub((*obs.Observer[int])(o))

		<-started
		sub.Unsub()
		close(block)

		time.Sleep(100 * time.Millisecond)

		// Only first emission must be observed
		assertlite.EqualSlices(t, []int{1}, rec.Values())
		// no error, no complete
		assertlite.False(t, rec.Terminated())
	})

	t.Run("no emission after terminal error", func(t *testing.T) {
		failure := errors.New("fail")

		src := obs.Async[int](func(o *obs.Observer[int]) {
			o.Next(1)
			o.Error(failure)
			o.Next(999) // illegal
		})

		retried := Retry[int](0)(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = retried.Sub((*obs.Observer[int])(o))

		_ = rec.AwaitError(1 * time.Second)

		assertlite.EqualSlices(t, []int{1}, rec.Values())
		assertlite.IsErrorOf(t, rec.Err(), failure)
	})
}

// TestRetryWhen_Sync_Sync tests RetryWhen using synchronous source
// and synchronous control observables.
//
// For synchronous execution, we assert:
//  1. Control Next triggers retry
//  2. Control Complete propagates original source error
//  3. Control Error propagates control error
//  4. Multiple control Nexts cause multiple retries
//  5. Successful source completion stops retries
//  6. No retry without control emission
func TestRetryWhen_Sync_Sync(t *testing.T) {
	t.Run("control next triggers retry", func(t *testing.T) {
		attempts := 0
		srcErr := errors.New("fail")

		src := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			attempts++
			if attempts < 2 {
				o.Error(srcErr)
				return subscr.Noop{}, nil
			}
			o.Next(42)
			o.Complete()
			return subscr.Noop{}, nil
		})

		retried := RetryWhen[int](func(err error) obs.Observable[any] {
			assertlite.IsErrorOf(t, err, srcErr)
			return obs.Of[any](nil) // trigger retry
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := retried.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		assertlite.EqualSlices(t, []int{42}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("control complete propagates original source error", func(t *testing.T) {
		srcErr := errors.New("source failure")

		src := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			o.Next(1)
			o.Error(srcErr)
			return subscr.Noop{}, nil
		})

		retried := RetryWhen[int](func(err error) obs.Observable[any] {
			return obs.Of[any]() // completes immediately
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = retried.Sub((*obs.Observer[int])(o))

		assertlite.EqualSlices(t, []int{1}, rec.Values())
		assertlite.IsErrorOf(t, rec.Err(), srcErr)
		assertlite.False(t, rec.Completed())
	})

	t.Run("control error propagates control error", func(t *testing.T) {
		srcErr := errors.New("source failure")
		ctrlErr := errors.New("control failure")

		src := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			o.Next(1)
			o.Error(srcErr)
			return subscr.Noop{}, nil
		})

		retried := RetryWhen[int](func(err error) obs.Observable[any] {
			return obs.Error[any](ctrlErr)
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = retried.Sub((*obs.Observer[int])(o))

		assertlite.EqualSlices(t, []int{1}, rec.Values())
		assertlite.IsErrorOf(t, rec.Err(), ctrlErr)
	})

	t.Run("multiple control nexts trigger multiple retries", func(t *testing.T) {
		attempts := 0
		fail := errors.New("fail")

		src := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			attempts++
			o.Next(attempts)
			o.Error(fail)
			return subscr.Noop{}, nil
		})

		retried := RetryWhen[int](func(err error) obs.Observable[any] {
			if attempts < 3 {
				return obs.Of[any](nil)
			}
			return obs.Error[any](err)
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = retried.Sub((*obs.Observer[int])(o))

		assertlite.EqualSlices(t, []int{1, 2, 3}, rec.Values())
		assertlite.IsErrorOf(t, rec.Err(), fail)
	})

	t.Run("successful source completion bypasses control", func(t *testing.T) {
		called := false

		src := obs.Of(1, 2, 3)

		retried := RetryWhen[int](func(err error) obs.Observable[any] {
			called = true
			return obs.Of[any](nil)
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = retried.Sub((*obs.Observer[int])(o))

		assertlite.EqualSlices(t, []int{1, 2, 3}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.False(t, called)
	})
}

// TestRetryWhen_Sync_Async tests RetryWhen using a synchronous source
// and an asynchronous control observable.
//
// For sync source + async control, we assert:
//  1. Control Next (async) triggers retry
//  2. Control Complete (async) propagates original source error
//  3. Control Error (async) propagates control error
//  4. Multiple control Nexts cause multiple retries
//  5. Successful source completion bypasses control
//  6. No retry occurs before control emission
func TestRetryWhen_Sync_Async(t *testing.T) {

	t.Run("control next triggers retry", func(t *testing.T) {
		attempts := 0
		srcErr := errors.New("fail")
		trigger := make(chan struct{})

		src := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			attempts++
			if attempts < 2 {
				o.Error(srcErr)
				return subscr.Noop{}, nil
			}
			o.Next(42)
			o.Complete()
			return subscr.Noop{}, nil
		})

		retried := RetryWhen[int](func(err error) obs.Observable[any] {
			assertlite.IsErrorOf(t, err, srcErr)
			return obs.Async[any](func(o *obs.Observer[any]) {
				<-trigger
				o.Next(nil)
				o.Complete()
			})
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = retried.Sub((*obs.Observer[int])(o))

		// ensure no retry before async signal
		time.Sleep(20 * time.Millisecond)
		assertlite.Equals(t, 1, attempts)
		assertlite.IsNil(t, rec.Values())

		close(trigger)
		rec.AwaitTermination(time.Second)

		assertlite.EqualSlices(t, []int{42}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("control complete propagates original source error", func(t *testing.T) {
		srcErr := errors.New("source failure")
		trigger := make(chan struct{})

		src := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			o.Next(1)
			o.Error(srcErr)
			return subscr.Noop{}, nil
		})

		retried := RetryWhen[int](func(err error) obs.Observable[any] {
			return obs.Async[any](func(o *obs.Observer[any]) {
				<-trigger
				o.Complete()
			})
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = retried.Sub((*obs.Observer[int])(o))

		time.Sleep(20 * time.Millisecond)
		assertlite.NoError(t, rec.Err())

		close(trigger)
		rec.AwaitTermination(time.Second)

		assertlite.EqualSlices(t, []int{1}, rec.Values())
		assertlite.IsErrorOf(t, rec.Err(), srcErr)
		assertlite.False(t, rec.Completed())
	})

	t.Run("control error propagates control error", func(t *testing.T) {
		srcErr := errors.New("source failure")
		ctrlErr := errors.New("control failure")
		trigger := make(chan struct{})

		src := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			o.Next(1)
			o.Error(srcErr)
			return subscr.Noop{}, nil
		})

		retried := RetryWhen[int](func(err error) obs.Observable[any] {
			return obs.Async[any](func(o *obs.Observer[any]) {
				<-trigger
				o.Error(ctrlErr)
			})
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = retried.Sub((*obs.Observer[int])(o))

		close(trigger)
		rec.AwaitTermination(time.Second)

		assertlite.EqualSlices(t, []int{1}, rec.Values())
		assertlite.IsErrorOf(t, rec.Err(), ctrlErr)
	})

	t.Run("multiple control nexts trigger multiple retries", func(t *testing.T) {
		attempts := 0
		fail := errors.New("fail")

		src := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			attempts++
			o.Next(attempts)
			o.Error(fail)
			return subscr.Noop{}, nil
		})

		retried := RetryWhen[int](func(err error) obs.Observable[any] {
			return obs.Async[any](func(o *obs.Observer[any]) {
				if attempts < 3 {
					o.Next(nil)
				}
				o.Complete()
			})
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = retried.Sub((*obs.Observer[int])(o))

		rec.AwaitTermination(time.Second)

		assertlite.EqualSlices(t, []int{1, 2, 3}, rec.Values())
		assertlite.IsErrorOf(t, rec.Err(), fail)
	})

	t.Run("successful source completion bypasses control", func(t *testing.T) {
		called := false

		src := obs.Of(1, 2, 3)

		retried := RetryWhen[int](func(err error) obs.Observable[any] {
			called = true
			return obs.Async[any](func(o *obs.Observer[any]) {
				o.Next(nil)
				o.Complete()
			})
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = retried.Sub((*obs.Observer[int])(o))

		assertlite.EqualSlices(t, []int{1, 2, 3}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.False(t, called)
	})
}

// TestRetryWhen_Async_Sync tests RetryWhen using an asynchronous source
// and a synchronous control observable.
//
// For async source + sync control, we assert:
//  1. Sync control Next triggers retry after async error
//  2. Sync control Complete propagates original source error
//  3. Sync control Error propagates control error
//  4. Multiple sync control Nexts cause multiple retries
//  5. Successful async source completion bypasses control
//  6. Async source errors do not retry without control emission
func TestRetryWhen_Async_Sync(t *testing.T) {
	t.Run("control next triggers retry", func(t *testing.T) {
		attempts := 0
		srcErr := errors.New("fail")

		src := obs.Async[int](func(o *obs.Observer[int]) {
			attempts++
			if attempts < 2 {
				o.Error(srcErr)
				t.Logf("Will RETURN at attempt %d", attempts)
				return
			}
			t.Logf("will NEXT at attempt %d", attempts)
			o.Next(42)
			o.Complete()
		})

		retried := RetryWhen[int](func(err error) obs.Observable[any] {
			assertlite.IsErrorOf(t, err, srcErr, "expected sane error from source but got '%v'", err)
			t.Logf("Will RETRY at attempt %d", attempts)
			return obs.Of[any](nil) // sync retry signal
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = retried.Sub((*obs.Observer[int])(o))

		rec.AwaitTermination(time.Second)

		assertlite.NoError(t, rec.Err())
		assertlite.True(t, rec.Completed())
		assertlite.EqualSlices(t, []int{42}, rec.Values())
	})

	t.Run("control complete propagates original source error", func(t *testing.T) {
		srcErr := errors.New("source failure")

		src := obs.Async[int](func(o *obs.Observer[int]) {
			o.Next(1)
			o.Error(srcErr)
		})

		retried := RetryWhen[int](func(err error) obs.Observable[any] {
			return obs.Of[any]() // completes immediately
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = retried.Sub((*obs.Observer[int])(o))

		rec.AwaitTermination(time.Second)

		assertlite.EqualSlices(t, []int{1}, rec.Values())
		assertlite.IsErrorOf(t, rec.Err(), srcErr)
		assertlite.False(t, rec.Completed())
	})

	t.Run("control error propagates control error", func(t *testing.T) {
		srcErr := errors.New("source failure")
		ctrlErr := errors.New("control failure")

		src := obs.Async[int](func(o *obs.Observer[int]) {
			o.Next(1)
			o.Error(srcErr)
		})

		retried := RetryWhen[int](func(err error) obs.Observable[any] {
			return obs.Error[any](ctrlErr)
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = retried.Sub((*obs.Observer[int])(o))

		rec.AwaitTermination(time.Second)

		assertlite.EqualSlices(t, []int{1}, rec.Values())
		assertlite.IsErrorOf(t, rec.Err(), ctrlErr)
	})

	t.Run("multiple control nexts trigger multiple retries", func(t *testing.T) {
		attempts := 0
		fail := errors.New("fail")

		src := obs.Async[int](func(o *obs.Observer[int]) {
			attempts++
			o.Next(attempts)
			o.Error(fail)
		})

		retried := RetryWhen[int](func(err error) obs.Observable[any] {
			if attempts < 3 {
				return obs.Of[any](nil)
			}
			return obs.Error[any](err)
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = retried.Sub((*obs.Observer[int])(o))

		rec.AwaitTermination(time.Second)

		assertlite.EqualSlices(t, []int{1, 2, 3}, rec.Values())
		assertlite.IsErrorOf(t, rec.Err(), fail)
	})

	t.Run("successful async source completion bypasses control", func(t *testing.T) {
		called := false

		src := obs.Async[int](func(o *obs.Observer[int]) {
			o.Next(1)
			o.Next(2)
			o.Next(3)
			o.Complete()
		})

		retried := RetryWhen[int](func(err error) obs.Observable[any] {
			called = true
			return obs.Of[any](nil)
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = retried.Sub((*obs.Observer[int])(o))

		rec.AwaitTermination(time.Second)

		assertlite.EqualSlices(t, []int{1, 2, 3}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.False(t, called)
	})

	t.Run("no retry without control emission", func(t *testing.T) {
		srcErr := errors.New("fail")

		src := obs.Async[int](func(o *obs.Observer[int]) {
			o.Error(srcErr)
		})

		retried := RetryWhen[int](func(err error) obs.Observable[any] {
			return obs.Empty[any]()
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = retried.Sub((*obs.Observer[int])(o))

		time.Sleep(50 * time.Millisecond)

		assertlite.IsNil(t, rec.Values())
		assertlite.IsErrorOf(t, srcErr, rec.Err())
		assertlite.False(t, rec.Completed())
	})
}

// TestRetryWhen_Async_Async tests RetryWhen using asynchronous source
// and asynchronous control observables.
//
// For async source + async control, we assert:
//  1. Async control Next triggers retry after async error
//  2. Async control Complete propagates original source error
//  3. Async control Error propagates control error
//  4. Multiple async control Nexts cause multiple retries
//  5. Successful async source completion bypasses control
//  6. Unsubscribe cancels both source and control
func TestRetryWhen_Async_Async(t *testing.T) {
	t.Run("async control next triggers retry", func(t *testing.T) {
		attempts := 0
		srcErr := errors.New("fail")

		src := obs.Async[int](func(o *obs.Observer[int]) {
			attempts++
			if attempts < 2 {
				o.Error(srcErr)
				return
			}
			o.Next(42)
			o.Complete()
		})

		retried := RetryWhen[int](func(err error) obs.Observable[any] {
			return obs.Async[any](func(o *obs.Observer[any]) {
				o.Next(nil) // async retry trigger
				o.Complete()
			})
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = retried.Sub((*obs.Observer[int])(o))

		rec.AwaitTermination(time.Second)

		assertlite.EqualSlices(t, []int{42}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("async control complete propagates original source error", func(t *testing.T) {
		srcErr := errors.New("source failure")

		src := obs.Async[int](func(o *obs.Observer[int]) {
			o.Next(1)
			o.Error(srcErr)
		})

		retried := RetryWhen[int](func(err error) obs.Observable[any] {
			return obs.Async[any](func(o *obs.Observer[any]) {
				o.Complete() // async complete, no retry
			})
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = retried.Sub((*obs.Observer[int])(o))

		rec.AwaitTermination(time.Second)

		assertlite.EqualSlices(t, []int{1}, rec.Values())
		assertlite.IsErrorOf(t, rec.Err(), srcErr)
		assertlite.False(t, rec.Completed())
	})

	t.Run("async control error propagates control error", func(t *testing.T) {
		srcErr := errors.New("source failure")
		ctrlErr := errors.New("control failure")

		src := obs.Async[int](func(o *obs.Observer[int]) {
			o.Next(1)
			o.Error(srcErr)
		})

		retried := RetryWhen[int](func(err error) obs.Observable[any] {
			return obs.Async[any](func(o *obs.Observer[any]) {
				o.Error(ctrlErr)
			})
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = retried.Sub((*obs.Observer[int])(o))

		rec.AwaitTermination(time.Second)

		assertlite.EqualSlices(t, []int{1}, rec.Values())
		assertlite.IsErrorOf(t, rec.Err(), ctrlErr)
	})

	t.Run("multiple async control nexts trigger multiple retries", func(t *testing.T) {
		attempts := 0
		fail := errors.New("fail")

		src := obs.Async[int](func(o *obs.Observer[int]) {
			attempts++
			o.Next(attempts)
			o.Error(fail)
		})

		retried := RetryWhen[int](func(err error) obs.Observable[any] {
			return obs.Async[any](func(o *obs.Observer[any]) {
				if attempts < 3 {
					o.Next(nil)
					o.Complete()
					return
				}
				o.Error(err)
			})
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = retried.Sub((*obs.Observer[int])(o))

		rec.AwaitTermination(time.Second)

		assertlite.EqualSlices(t, []int{1, 2, 3}, rec.Values())
		assertlite.IsErrorOf(t, rec.Err(), fail)
	})

	t.Run("successful async source completion bypasses async control", func(t *testing.T) {
		called := atomic.Bool{}

		src := obs.Async[int](func(o *obs.Observer[int]) {
			o.Next(1)
			o.Next(2)
			o.Next(3)
			o.Complete()
		})

		retried := RetryWhen[int](func(err error) obs.Observable[any] {
			called.Store(true)
			return obs.AsyncOf[any](nil)
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = retried.Sub((*obs.Observer[int])(o))

		rec.AwaitTermination(time.Second)

		assertlite.EqualSlices(t, []int{1, 2, 3}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.False(t, called.Load())
	})

	t.Run("unsubscribe cancels source and control", func(t *testing.T) {
		srcStarted := make(chan struct{})
		ctrlStarted := make(chan struct{})
		ctrlBlocked := make(chan struct{})

		src := obs.Async[int](func(o *obs.Observer[int]) {
			close(srcStarted)
			o.Error(errors.New("fail"))
		})

		retried := RetryWhen[int](func(err error) obs.Observable[any] {
			return obs.Async[any](func(o *obs.Observer[any]) {
				close(ctrlStarted)
				<-ctrlBlocked // block forever unless unsubscribed
				o.Next(nil)
			})
		})(src)

		rec, o := observers.NewRecorder[int](t)
		sub, _ := retried.Sub((*obs.Observer[int])(o))

		<-srcStarted
		<-ctrlStarted

		sub.Unsub()
		close(ctrlBlocked)

		time.Sleep(50 * time.Millisecond)

		assertlite.IsNil(t, rec.Values())
		assertlite.False(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})
}
