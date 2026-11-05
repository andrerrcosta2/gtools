// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package observables

import (
	"errors"
	"testing"
	"time"

	"github.com/andrerrcosta2/gtools/core/data/iterators"
	"github.com/andrerrcosta2/gtools/core/testlite/assertlite"
	"github.com/andrerrcosta2/gtools/gflux/obs/internal/observers"
	"github.com/andrerrcosta2/gtools/gflux/obs/internal/subscr"
)

// TestColdObs tests the core behavior of ColdObs[T any] observables.
//
// ColdObs wraps a subscription function and exposes Sub.
// We want to assert:
//  1. Emits all values synchronously in order
//  2. Calls Complete after the last value
//  3. Forwards errors correctly
//  4. Supports unsubscription (though for sync sources, unsub cannot interrupt)
//  5. Works with empty sequences
func TestColdObs(t *testing.T) {
	t.Run("emits all values in order and completes", func(t *testing.T) {
		values := []int{1, 2, 3}

		src := &ColdObs[int]{
			sub: func(o *observers.Observer[int]) (subscr.Subscription, error) {
				for _, v := range values {
					o.Next(v)
				}
				o.Complete()
				return subscr.Noop{}, nil
			},
		}

		rec, o := observers.NewRecorder[int](t)
		_, err := src.Sub(o)
		assertlite.NoError(t, err)
		assertlite.EqualSlices(t, values, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("works with empty sequence", func(t *testing.T) {
		src := &ColdObs[int]{
			sub: func(o *observers.Observer[int]) (subscr.Subscription, error) {
				o.Complete()
				return subscr.Noop{}, nil
			},
		}

		rec, o := observers.NewRecorder[int](t)
		_, _ = src.Sub(o)

		assertlite.Equals(t, 0, len(rec.Values()))
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("unsub returns subscription and closes slot", func(t *testing.T) {
		called := false
		src := &ColdObs[int]{
			sub: func(o *observers.Observer[int]) (subscr.Subscription, error) {
				called = true
				return subscr.Noop{}, nil
			},
		}

		sub, err := src.Sub(&observers.Observer[int]{})
		assertlite.NoError(t, err)
		assertlite.True(t, called)

		sub.Unsub() // should be safe to call even if sync
	})
}

// TestReplayObs tests the ReplayObs[T any] behavior.
//
// ReplayObs wraps a source Observable and buffers all emitted values.
// We want to assert:
//  1. The first subscriber triggers the source
//  2. Values are replayed to subsequent subscribers in order
//  3. Completion is signaled to all subscribers
//  4. Errors are propagated and replayed
//  5. No values are emitted after completion/error
//  6. Works correctly with empty sources
func TestReplayObs(t *testing.T) {
	t.Run("first subscriber triggers source", func(t *testing.T) {
		called := false
		src := Func[int](func(o *observers.Observer[int]) (subscr.Subscription, error) {
			called = true
			o.Next(42)
			o.Complete()
			return subscr.Noop{}, nil
		})

		replay := &ReplayObs[int]{src: src}
		rec, o := observers.NewRecorder[int](t)
		_, _ = replay.Sub(o)

		assertlite.True(t, called)
		assertlite.EqualSlices(t, []int{42}, rec.Values())
		assertlite.True(t, rec.Completed())
	})

	t.Run("replays values to subsequent subscribers", func(t *testing.T) {
		src := Of(1, 2, 3)
		replay := &ReplayObs[int]{src: src}

		rec1, o1 := observers.NewRecorder[int](t)
		_, _ = replay.Sub(o1)
		assertlite.EqualSlices(t, []int{1, 2, 3}, rec1.Values())

		// second subscriber should get same values
		rec2, o2 := observers.NewRecorder[int](t)
		_, _ = replay.Sub(o2)
		assertlite.EqualSlices(t, []int{1, 2, 3}, rec2.Values())
		assertlite.True(t, rec2.Completed())
	})

	t.Run("propagates error to all subscribers", func(t *testing.T) {
		failure := errors.New("failure")
		src := Func[int](func(o *observers.Observer[int]) (subscr.Subscription, error) {
			o.Next(1)
			o.Error(failure)
			return subscr.Noop{}, nil
		})

		replay := &ReplayObs[int]{src: src}

		rec1, o1 := observers.NewRecorder[int](t)
		_, _ = replay.Sub(o1)
		assertlite.EqualSlices(t, []int{1}, rec1.Values())
		assertlite.IsErrorOf(t, rec1.Err(), failure)
		assertlite.False(t, rec1.Completed())

		// late subscriber receives replayed values + error
		rec2, o2 := observers.NewRecorder[int](t)
		_, _ = replay.Sub(o2)
		assertlite.EqualSlices(t, []int{1}, rec2.Values())
		assertlite.IsErrorOf(t, rec2.Err(), failure)
		assertlite.False(t, rec2.Completed())
	})

	t.Run("replays completion to late subscribers", func(t *testing.T) {
		src := Of(1, 2)
		replay := &ReplayObs[int]{src: src}

		rec1, o1 := observers.NewRecorder[int](t)
		_, _ = replay.Sub(o1)
		assertlite.True(t, rec1.Completed())

		// late subscriber gets same values + completion
		rec2, o2 := observers.NewRecorder[int](t)
		_, _ = replay.Sub(o2)
		assertlite.EqualSlices(t, []int{1, 2}, rec2.Values())
		assertlite.True(t, rec2.Completed())
	})

	t.Run("does not emit after completion", func(t *testing.T) {
		src := Of(1)
		replay := &ReplayObs[int]{src: src}

		rec, o := observers.NewRecorder[int](t)
		_, _ = replay.Sub(o)
		_, _ = replay.Sub(o) // late subscriber

		assertlite.EqualSlices(t, []int{1}, rec.Values())
		assertlite.True(t, rec.Completed())
	})

	t.Run("works with empty source", func(t *testing.T) {
		src := Of[int]()
		replay := &ReplayObs[int]{src: src}

		rec, o := observers.NewRecorder[int](t)
		_, _ = replay.Sub(o)

		assertlite.Equals(t, 0, len(rec.Values()))
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})
}

// TestChanObs tests the Chan observable wrapper.
//
// Chan wraps a Go channel as a Coldsubscr. We want to assert:
//  1. Emits all values sent on the channel
//  2. Completes when the channel closes
//  3. Does not emit after completion
//  4. Supports cancellation via subscription
//  5. Works with empty channels
func TestChanObs(t *testing.T) {
	t.Run("emits all values and completes", func(t *testing.T) {
		ch := make(chan int)
		src := Chan(ch)

		rec, o := observers.NewRecorder[int](t)
		go func() {
			_, _ = src.Sub(o)
		}()

		go func() {
			ch <- 1
			ch <- 2
			ch <- 3
			close(ch)
		}()

		time.Sleep(10 * time.Millisecond) // allow goroutines to process

		assertlite.EqualSlices(t, []int{1, 2, 3}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("completes immediately if channel is already closed", func(t *testing.T) {
		ch := make(chan int)
		close(ch)
		src := Chan(ch)

		rec, o := observers.NewRecorder[int](t)
		_, _ = src.Sub(o)

		assertlite.Equals(t, 0, len(rec.Values()))
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("works with empty channel without sending values", func(t *testing.T) {
		ch := make(chan int)
		src := Chan(ch)

		rec, o := observers.NewRecorder[int](t)
		go func() { _, _ = src.Sub(o) }()

		close(ch)
		time.Sleep(5 * time.Millisecond)

		assertlite.Equals(t, 0, len(rec.Values()))
		assertlite.True(t, rec.Completed())
	})
}

// TestDeferObs tests the Defer[T any] observable.
//
// Defer delays the creation of the actual observable until Sub is called.
// We want to assert:
//  1. The factory function is only invoked on subscription
//  2. Values from the deferred observable are emitted correctly
//  3. Completion is forwarded to subscribers
//  4. Errors are forwarded to subscribers
//  5. Range subscription creates a fresh observable
func TestDeferObs(t *testing.T) {
	t.Run("factory function is only called on subscription", func(t *testing.T) {
		called := false
		src := Defer(func() Observable[int] {
			called = true
			return Of(1, 2)
		})

		// not called yet
		assertlite.False(t, called)

		rec, o := observers.NewRecorder[int](t)
		_, _ = src.Sub(o)

		assertlite.True(t, called)
		assertlite.EqualSlices(t, []int{1, 2}, rec.Values())
		assertlite.True(t, rec.Completed())
	})

	t.Run("emits values and completes correctly", func(t *testing.T) {
		src := Defer(func() Observable[int] {
			return Of(10, 20, 30)
		})

		rec, o := observers.NewRecorder[int](t)
		_, _ = src.Sub(o)

		assertlite.EqualSlices(t, []int{10, 20, 30}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("forwards error from deferred observable", func(t *testing.T) {
		failure := errors.New("deferred error")
		src := Defer(func() Observable[int] {
			return Error[int](failure)
		})

		rec, o := observers.NewRecorder[int](t)
		_, _ = src.Sub(o)

		assertlite.Equals(t, 0, len(rec.Values()))
		assertlite.IsErrorOf(t, rec.Err(), failure)
		assertlite.False(t, rec.Completed())
	})

	t.Run("each subscription triggers a fresh observable", func(t *testing.T) {
		counter := 0
		src := Defer(func() Observable[int] {
			counter++
			return Of(counter)
		})

		rec1, o1 := observers.NewRecorder[int](t)
		_, _ = src.Sub(o1)
		assertlite.EqualSlices(t, []int{1}, rec1.Values())

		rec2, o2 := observers.NewRecorder[int](t)
		_, _ = src.Sub(o2)
		assertlite.EqualSlices(t, []int{2}, rec2.Values())
	})
}

// TestDelayedOfObs tests DelayedOf[T any].
//
// DelayedOf emits each value with a delay determined by the IndexedDelayFn.
// We want to assert:
//  1. Values are emitted in the correct order
//  2. Completion occurs after all values
//  3. Works correctly when no values are provided
func TestDelayedOfObs(t *testing.T) {
	t.Run("emits values in order with delays", func(t *testing.T) {
		delays := []time.Duration{1, 2, 1} // in ms
		src := DelayedOf(func(i int, v int) time.Duration {
			return delays[i] * time.Millisecond
		}, 10, 20, 30)

		start := time.Now()
		rec, o := observers.NewRecorder[int](t)
		_, _ = src.Sub(o)

		values := rec.Values()
		assertlite.EqualSlices(t, []int{10, 20, 30}, values)
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())

		elapsed := time.Since(start)
		expected := delays[0] + delays[1] + delays[2]
		assertlite.True(t, elapsed >= time.Duration(expected)*time.Millisecond, "total elapsed time should respect sum of delays")
	})

	t.Run("completes immediately if no values", func(t *testing.T) {
		src := DelayedOf(func(i int, v int) time.Duration { return 0 })

		rec, o := observers.NewRecorder[int](t)
		_, _ = src.Sub(o)

		assertlite.Equals(t, 0, len(rec.Values()))
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})
}

// TestErrorObs tests the synchronous Error[T any] observable.
//
// Error immediately signals the provided error to the observer.
// We want to assert:
//  1. Error is received correctly
//  2. No values are emitted
//  3. Complete is not called
func TestErrorObs(t *testing.T) {
	t.Run("emits error immediately", func(t *testing.T) {
		failure := errors.New("immediate failure")
		src := Error[int](failure)

		rec, o := observers.NewRecorder[int](t)
		_, _ = src.Sub(o)

		assertlite.Equals(t, 0, len(rec.Values()))
		assertlite.IsErrorOf(t, rec.Err(), failure)
		assertlite.False(t, rec.Completed())
	})
}

// TestEmptyObs tests the synchronous Empty[T any] observable.
//
// Empty immediately signals completion to the observer without emitting any values.
// We want to assert:
//  1. No values are emitted
//  2. Completion is signaled immediately
//  3. No error is emitted
func TestEmptyObs(t *testing.T) {
	t.Run("completes immediately with no values", func(t *testing.T) {
		src := Empty[int]()

		rec, o := observers.NewRecorder[int](t)
		_, _ = src.Sub(o)

		assertlite.Equals(t, 0, len(rec.Values()))
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})
}

// TestFromFuncObs tests the FromFunc[T any] observable.
//
// FromFunc executes a user-provided function immediately on subscription,
// passing an observer for emitting values. We want to assert:
//  1. Values emitted by the function are forwarded correctly
//  2. Completion is signaled if the function calls Complete
//  3. Errors are signaled if the function calls Error
//  4. Works with multiple values or no values
func TestFromFuncObs(t *testing.T) {
	t.Run("emits multiple values and completes", func(t *testing.T) {
		src := FromFunc[int](func(o *observers.Observer[int]) {
			o.Next(1)
			o.Next(2)
			o.Complete()
		})

		rec, o := observers.NewRecorder[int](t)
		_, _ = src.Sub(o)

		assertlite.EqualSlices(t, []int{1, 2}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("emits error correctly", func(t *testing.T) {
		failure := errors.New("function error")
		src := FromFunc[int](func(o *observers.Observer[int]) {
			o.Error(failure)
		})

		rec, o := observers.NewRecorder[int](t)
		_, _ = src.Sub(o)

		assertlite.Equals(t, 0, len(rec.Values()))
		assertlite.IsErrorOf(t, rec.Err(), failure)
		assertlite.False(t, rec.Completed())
	})

	t.Run("handles no values and immediate completion", func(t *testing.T) {
		src := FromFunc[int](func(o *observers.Observer[int]) {
			o.Complete()
		})

		rec, o := observers.NewRecorder[int](t)
		_, _ = src.Sub(o)

		assertlite.IsNil(t, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})
}

// TestFromIteratorObs tests an observable built from an iterator.
//
// This observable emits all values from the iterator synchronously and
// completes when exhausted. We want to assert:
//  1. Values are emitted in order
//  2. Completion is signaled after the last value
//  3. Works with empty iterators
func TestFromIter(t *testing.T) {
	t.Run("emits all values in order and completes", func(t *testing.T) {
		data := []int{1, 2, 3, 4}
		it := iterators.Default(data)

		src := FromIter[int](it)

		rec, o := observers.NewRecorder[int](t)
		_, err := src.Sub(o)

		assertlite.NoError(t, err)
		assertlite.EqualSlices(t, data, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("completes immediately if iterator is empty", func(t *testing.T) {
		it := iterators.Default([]int{})

		src := FromIter[int](it)

		rec, o := observers.NewRecorder[int](t)
		_, err := src.Sub(o)

		assertlite.NoError(t, err)
		assertlite.IsNil(t, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})
}

// TestIntervalOfObs tests the synchronous IntervalOf[T any] observable.
//
// IntervalOf emits each value at a fixed delay. We want to assert:
//  1. Values are emitted in the correct order
//  2. Completion occurs after the last value
//  3. Handles empty values correctly
//  4. Total elapsed time respects the interval delays
func TestIntervalOfObs(t *testing.T) {
	t.Run("emits values in order with fixed intervals", func(t *testing.T) {
		interval := 10 * time.Millisecond
		src := IntervalOf(interval, 1, 2, 3)

		start := time.Now()
		rec, o := observers.NewRecorder[int](t)
		_, _ = src.Sub(o)

		assertlite.EqualSlices(t, []int{1, 2, 3}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())

		elapsed := time.Since(start)
		expected := interval * time.Duration(len([]int{1, 2, 3})+1) // initial + per-value delays
		assertlite.True(t, elapsed >= expected, "total elapsed time should respect interval schedule")
	})

	t.Run("completes after a single interval if no values", func(t *testing.T) {
		interval := 5 * time.Millisecond
		src := IntervalOf[int](interval)

		start := time.Now()
		rec, o := observers.NewRecorder[int](t)
		_, _ = src.Sub(o)

		assertlite.Equals(t, 0, len(rec.Values()))
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())

		elapsed := time.Since(start)
		assertlite.True(t, elapsed >= interval, "completion should be delayed by a single interval even if no values")
	})
}

// TestJustObs tests the synchronous Just[T any] observable.
//
// Just emits exactly one value and completes immediately. We want to assert:
//  1. The single value is emitted
//  2. Completion is signaled immediately after the value
//  3. No error is emitted
func TestJustObs(t *testing.T) {
	t.Run("emits the single value and completes", func(t *testing.T) {
		src := Just(42)

		rec, o := observers.NewRecorder[int](t)
		_, _ = src.Sub(o)

		assertlite.EqualSlices(t, []int{42}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("does not emit additional values after completion", func(t *testing.T) {
		src := Just("hello")

		rec, o := observers.NewRecorder[string](t)
		_, _ = src.Sub((*observers.Observer[string])(o))

		assertlite.EqualSlices(t, []string{"hello"}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})
}

// TestOfObs tests the synchronous Of[T any] observable.
//
// Of emits all provided values in order and completes immediately.
// We want to assert:
//  1. Multiple values are emitted in order
//  2. Completion occurs after the last value
//  3. Works with a single value
//  4. Works with no values
func TestOfObs(t *testing.T) {
	t.Run("emits multiple values in order and completes", func(t *testing.T) {
		src := Of(1, 2, 3)

		rec, o := observers.NewRecorder[int](t)
		_, _ = src.Sub(o)

		assertlite.EqualSlices(t, []int{1, 2, 3}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("emits a single value correctly", func(t *testing.T) {
		src := Of("hello")

		rec, o := observers.NewRecorder[string](t)
		_, _ = src.Sub((*observers.Observer[string])(o))

		assertlite.EqualSlices(t, []string{"hello"}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("completes immediately if no values are provided", func(t *testing.T) {
		src := Of[int]()

		rec, o := observers.NewRecorder[int](t)
		_, _ = src.Sub(o)

		assertlite.Equals(t, 0, len(rec.Values()))
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})
}

// TestRangeObs tests the synchronous Range[T any] observable.
//
// Range emits a sequence of ordered values from `from` to `to` with a step `delta`.
// We want to assert:
//  1. Values are emitted in correct order for positive delta
//  2. Values are emitted in correct order for negative delta
//  3. Handles single-value ranges
//  4. Panics on invalid delta (zero or wrong sign)
//  5. Completes after the last value
func TestRangeObs(t *testing.T) {
	t.Run("emits increasing values correctly", func(t *testing.T) {
		src := Range(1, 5, 2) // 1,3,5

		rec, o := observers.NewRecorder[int](t)
		_, _ = src.Sub(o)

		assertlite.EqualSlices(t, []int{1, 3, 5}, rec.Values())
		assertlite.True(t, rec.Completed())
	})

	t.Run("emits decreasing values correctly", func(t *testing.T) {
		src := Range(5, 1, -2) // 5,3,1

		rec, o := observers.NewRecorder[int](t)
		_, _ = src.Sub(o)

		assertlite.EqualSlices(t, []int{5, 3, 1}, rec.Values())
		assertlite.True(t, rec.Completed())
	})

	t.Run("handles single-value range", func(t *testing.T) {
		src := Range(5, 5, 1) // only 5

		rec, o := observers.NewRecorder[int](t)
		_, _ = src.Sub(o)

		assertlite.EqualSlices(t, []int{5}, rec.Values())
		assertlite.True(t, rec.Completed())
	})

	t.Run("panics if delta is zero", func(t *testing.T) {
		assertlite.Panics(t, func() {
			_ = Range(1, 5, 0)
		})
	})

	t.Run("panics if delta sign does not match range direction", func(t *testing.T) {
		assertlite.Panics(t, func() {
			_ = Range(1, 5, -1) // negative delta but increasing
		})
		assertlite.Panics(t, func() {
			_ = Range(5, 1, 1) // positive delta but decreasing
		})
	})
}
