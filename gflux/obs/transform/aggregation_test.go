// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package transform

import (
	"errors"
	"testing"
	"time"

	"github.com/andrerrcosta2/gtools/core/testlite/assertlite"
	"github.com/andrerrcosta2/gtools/gflux/obs"
	"github.com/andrerrcosta2/gtools/gflux/obs/internal/observers"
	"github.com/andrerrcosta2/gtools/gflux/obs/internal/subscr"
)

// TestReduce_Sync
//
// Source emits synchronously.
//
// We should assert:
//  1. Reduction emits a single accumulated value on completion
//  2. Single-value and empty sources behave correctly (seed handling)
//  3. No intermediate accumulator states are emitted
//  4. Completion is forwarded synchronously
func TestReduce_Sync(t *testing.T) {
	t.Run("reduces multiple values", func(t *testing.T) {
		src := obs.Of(1, 2, 3)

		reduced := Reduce(0, func(acc, v int) int {
			return acc + v
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := reduced.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		assertlite.EqualSlices(t, []int{6}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("reduces single value", func(t *testing.T) {
		src := obs.Of(5)

		reduced := Reduce(10, func(acc, v int) int {
			return acc + v
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = reduced.Sub((*obs.Observer[int])(o))

		assertlite.EqualSlices(t, []int{15}, rec.Values())
		assertlite.True(t, rec.Completed())
	})

	t.Run("emits seed on empty source", func(t *testing.T) {
		src := obs.Of[int]()

		reduced := Reduce(42, func(acc, v int) int {
			return acc + v
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = reduced.Sub((*obs.Observer[int])(o))

		assertlite.EqualSlices(t, []int{42}, rec.Values())
		assertlite.True(t, rec.Completed())
	})

	t.Run("emits exactly once", func(t *testing.T) {
		src := obs.Of(1, 2, 3, 4)

		reduced := Reduce(0, func(acc, v int) int {
			return acc + v
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = reduced.Sub((*obs.Observer[int])(o))

		assertlite.True(t, len(rec.Values()) == 1)
	})

	t.Run("completes synchronously", func(t *testing.T) {
		src := obs.Of(1, 2)

		reduced := Reduce(0, func(acc, v int) int {
			return acc + v
		})(src)

		_, o := observers.NewRecorder[int](t)

		done := false
		o.Complete = func() {
			done = true
		}

		_, _ = reduced.Sub((*obs.Observer[int])(o))

		assertlite.True(t, done)
	})
}

// TestReduce_Async
//
// Source emits asynchronously.
//
// We should assert:
//  1. Reduction emits exactly one final accumulated value
//  2. Intermediate accumulator states are not emitted
//  3. Empty source emits the seed then completes
//  4. Errors are propagated and terminate accumulation
func TestReduce_Async(t *testing.T) {
	t.Run("reduces async source into single emission", func(t *testing.T) {
		src := obs.AsyncOf(1, 2, 3)

		reduced := Reduce(0, func(acc, v int) int {
			return acc + v
		})(src)

		rec, o := observers.NewRecorder[int](t)

		done, err := reduced.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)
		assertlite.NotNil(t, done)

		rec.AwaitTermination(100 * time.Millisecond)

		assertlite.EqualSlices(t, rec.Values(), []int{6})
		assertlite.True(t, rec.Completed())
		assertlite.IsNil(t, rec.Err())
	})

	t.Run("does not emit intermediate values", func(t *testing.T) {
		src := obs.AsyncOf(1, 2, 3, 4)

		reduced := Reduce(1, func(acc, v int) int {
			return acc * v
		})(src)

		rec, o := observers.NewRecorder[int](t)

		done, err := reduced.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)
		assertlite.NotNil(t, done)

		rec.AwaitTermination(100 * time.Millisecond)

		// only final result
		assertlite.EqualSlices(t, rec.Values(), []int{24})
	})

	t.Run("empty async source emits seed and completes", func(t *testing.T) {
		src := obs.AsyncOf[int]()

		reduced := Reduce(42, func(acc, v int) int {
			return acc + v
		})(src)

		rec, o := observers.NewRecorder[int](t)

		done, err := reduced.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)
		assertlite.NotNil(t, done)

		rec.AwaitTermination(100 * time.Millisecond)

		assertlite.EqualSlices(t, rec.Values(), []int{42})
		assertlite.True(t, rec.Completed())
	})

	t.Run("propagates error and does not emit", func(t *testing.T) {
		failure := errors.New("failure")

		src := obs.AsyncError[int](failure)

		reduced := Reduce(0, func(acc, v int) int {
			return acc + v
		})(src)

		rec, o := observers.NewRecorder[int](t)

		sub, err := reduced.Sub((*obs.Observer[int])(o))

		assertlite.NoError(t, err)
		assertlite.NotNil(t, sub)

		rec.AwaitTermination(100 * time.Millisecond)

		assertlite.EqualSlices(t, rec.Values(), nil)
		assertlite.False(t, rec.Completed())
		assertlite.IsErrorOf(t, rec.Err(), failure)
	})

	t.Run("stops accumulation after error", func(t *testing.T) {
		src := obs.Async[int](func(o *obs.Observer[int]) {
			o.Next(1)
			o.Next(2)
			o.Error(errors.New("fail"))
			o.Next(999) // must be ignored
		})

		reduced := Reduce(0, func(acc, v int) int {
			return acc + v
		})(src)

		rec, o := observers.NewRecorder[int](t)

		done, err := reduced.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)
		assertlite.NotNil(t, done)

		rec.AwaitTermination(100 * time.Millisecond)

		assertlite.EqualSlices(t, rec.Values(), nil)
		assertlite.False(t, rec.Completed())
		assertlite.NotNil(t, rec.Err())
	})
}

// TestScan_Sync
//
// Source emits synchronously.
//
// We should assert:
//  1. Range step of accumulation is emitted in order
//  2. Seed is not emitted by itself for empty sources
//  3. Completion is forwarded normally
//  4. Errors are propagated immediately and stop further emissions
func TestScan_Sync(t *testing.T) {
	t.Run("accumulates and emits each step", func(t *testing.T) {
		src := obs.Of(1, 2, 3)

		scanned := Scan(0, func(acc, v int) int {
			return acc + v
		})(src)

		rec, o := observers.NewRecorder[int](t)

		sub, err := scanned.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)
		assertlite.NotNil(t, sub)

		assertlite.EqualSlices(t, rec.Values(), []int{1, 3, 6})
		assertlite.True(t, rec.Completed())
		assertlite.IsNil(t, rec.Err())
	})

	t.Run("does not emit seed by itself", func(t *testing.T) {
		src := obs.Of[int]()

		scanned := Scan(42, func(acc, v int) int {
			return acc + v
		})(src)

		rec, o := observers.NewRecorder[int](t)

		sub, err := scanned.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)
		assertlite.NotNil(t, sub)

		assertlite.EqualSlices(t, rec.Values(), nil)
		assertlite.True(t, rec.Completed())
	})

	t.Run("works with single emission", func(t *testing.T) {
		src := obs.Of(5)

		scanned := Scan(10, func(acc, v int) int {
			return acc + v
		})(src)

		rec, o := observers.NewRecorder[int](t)

		sub, err := scanned.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)
		assertlite.NotNil(t, sub)

		assertlite.EqualSlices(t, rec.Values(), []int{15})
		assertlite.True(t, rec.Completed())
	})

	t.Run("propagates error immediately", func(t *testing.T) {
		failure := errors.New("failure")

		src := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			o.Next(1)
			o.Error(failure)
			o.Next(999) // must be ignored
			return subscr.Noop{}, nil
		})

		scanned := Scan(0, func(acc, v int) int {
			return acc + v
		})(src)

		rec, o := observers.NewRecorder[int](t)

		sub, err := scanned.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)
		assertlite.NotNil(t, sub)

		assertlite.EqualSlices(t, rec.Values(), []int{1})
		assertlite.False(t, rec.Completed())
		assertlite.IsErrorOf(t, rec.Err(), failure)
	})

	t.Run("stops emitting after error", func(t *testing.T) {
		src := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			o.Next(1)
			o.Next(2)
			o.Error(errors.New("fail"))
			o.Next(999)
			return subscr.Noop{}, nil
		})

		scanned := Scan(0, func(acc, v int) int {
			return acc + v
		})(src)

		rec, o := observers.NewRecorder[int](t)

		sub, err := scanned.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)
		assertlite.NotNil(t, sub)

		assertlite.EqualSlices(t, rec.Values(), []int{1, 3})
		assertlite.False(t, rec.Completed())
		assertlite.NotNil(t, rec.Err())
	})
}

// TestScan_Async
//
// Source emits asynchronously.
//
// We should assert:
//  1. Accumulated values are emitted step-by-step in order
//  2. Completion is forwarded and no values are emitted after termination
//  3. Errors are propagated and stop further emissions
//  4. Empty source completes without emitting values
func TestScan_Async(t *testing.T) {
	t.Run("emits accumulated values from async source", func(t *testing.T) {
		src := obs.AsyncOf(1, 2, 3)

		scanned := Scan(0, func(acc, v int) int {
			return acc + v
		})(src)

		rec, o := observers.NewRecorder[int](t)

		sub, err := scanned.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)
		assertlite.NotNil(t, sub)

		rec.AwaitCompleted(100 * time.Millisecond)

		assertlite.EqualSlices(t, rec.Values(), []int{1, 3, 6})
		assertlite.True(t, rec.Completed())
		assertlite.IsNil(t, rec.Err())
	})

	t.Run("does not emit after completion", func(t *testing.T) {
		src := obs.Async[int](func(o *obs.Observer[int]) {
			o.Next(1)
			o.Complete()
			o.Next(999) // must be ignored
		})

		scanned := Scan(0, func(acc, v int) int {
			return acc + v
		})(src)

		rec, o := observers.NewRecorder[int](t)

		sub, err := scanned.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)
		assertlite.NotNil(t, sub)

		rec.AwaitCompleted(100 * time.Millisecond)

		assertlite.EqualSlices(t, rec.Values(), []int{1})
		assertlite.True(t, rec.Completed())
	})

	t.Run("propagates error and stops emitting", func(t *testing.T) {
		failure := errors.New("failure")

		src := obs.Async[int](func(o *obs.Observer[int]) {
			o.Next(1)
			o.Next(2)
			o.Error(failure)
			o.Next(999) // must be ignored
		})

		scanned := Scan(0, func(acc, v int) int {
			return acc + v
		})(src)

		rec, o := observers.NewRecorder[int](t)

		sub, err := scanned.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)
		assertlite.NotNil(t, sub)

		_ = rec.AwaitError(100 * time.Millisecond)

		assertlite.EqualSlices(t, rec.Values(), []int{1, 3})
		assertlite.False(t, rec.Completed())
		assertlite.IsErrorOf(t, rec.Err(), failure)
	})

	t.Run("empty async source completes without emitting", func(t *testing.T) {
		src := obs.AsyncOf[int]()

		scanned := Scan(0, func(acc, v int) int {
			return acc + v
		})(src)

		rec, o := observers.NewRecorder[int](t)

		sub, err := scanned.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)
		assertlite.NotNil(t, sub)

		rec.AwaitCompleted(100 * time.Millisecond)

		assertlite.EqualSlices(t, rec.Values(), nil)
		assertlite.True(t, rec.Completed())
		assertlite.IsNil(t, rec.Err())
	})
}
