// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package transform

import (
	"errors"
	"sync/atomic"
	"testing"
	"time"

	"github.com/andrerrcosta2/gtools/core/domain/function/predicates"
	"github.com/andrerrcosta2/gtools/core/testlite/assertlite"
	"github.com/andrerrcosta2/gtools/gflux/obs"
	"github.com/andrerrcosta2/gtools/gflux/obs/internal/observers"
	"github.com/andrerrcosta2/gtools/gflux/obs/internal/subscr"
)

// TestDistinct_Sync tests Distinct using synchronous sources.
//
// For synchronous sources, we must assert:
//
//  1. First occurrences are forwarded
//  2. Duplicate values are suppressed
//  3. Order of first occurrences is preserved
//  4. Error from source is propagated immediately
//  5. Completion from source is forwarded
//  6. No values are emitted after termination
//  7. State is local to the subscription
//  8. Completes synchronously when source completes synchronously
func TestDistinct_Sync(t *testing.T) {
	t.Run("forwards only first occurrences and preserves order", func(t *testing.T) {
		src := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			o.Next(1)
			o.Next(2)
			o.Next(1)
			o.Next(3)
			o.Next(2)
			o.Complete()
			return subscr.Noop{}, nil
		})

		distinct := Distinct[int]()(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = distinct.Sub((*obs.Observer[int])(o))

		assertlite.EqualSlices(t, []int{1, 2, 3}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("does nothing when all values are unique", func(t *testing.T) {
		src := obs.Of(1, 2, 3)

		distinct := Distinct[int]()(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = distinct.Sub((*obs.Observer[int])(o))

		assertlite.EqualSlices(t, []int{1, 2, 3}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("propagates source error immediately", func(t *testing.T) {
		failure := errors.New("failure")

		src := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			o.Next(1)
			o.Next(1)
			o.Error(failure)
			o.Next(2) // must be ignored
			return subscr.Noop{}, nil
		})

		distinct := Distinct[int]()(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = distinct.Sub((*obs.Observer[int])(o))

		assertlite.EqualSlices(t, []int{1}, rec.Values())
		assertlite.IsErrorOf(t, rec.Err(), failure)
		assertlite.False(t, rec.Completed())
	})

	t.Run("does not share states between subscriptions", func(t *testing.T) {
		src := obs.Of(1, 1, 2)

		distinct := Distinct[int]()(src)

		rec1, o1 := observers.NewRecorder[int](t)
		rec2, o2 := observers.NewRecorder[int](t)

		_, _ = distinct.Sub((*obs.Observer[int])(o1))
		_, _ = distinct.Sub((*obs.Observer[int])(o2))

		assertlite.EqualSlices(t, []int{1, 2}, rec1.Values())
		assertlite.EqualSlices(t, []int{1, 2}, rec2.Values())
	})

	t.Run("completes synchronously when source completes synchronously", func(t *testing.T) {
		done := false

		src := obs.Of(1, 1)

		distinct := Distinct[int]()(src)

		o := &obs.Observer[int]{
			Next: func(int) {},
			Complete: func() {
				done = true
			},
		}

		_, _ = distinct.Sub(o)

		assertlite.True(t, done)
	})
}

// TestDistinctUntilChanged_Sync tests DistinctUntilChanged using synchronous sources.
//
// For synchronous sources, we must assert:
//
//  1. Consecutive duplicates are suppressed
//  2. Non-consecutive duplicates are forwarded
//  3. Order is preserved
//  4. Error from source is propagated immediately
//  5. Completion from source is forwarded
//  6. No values are emitted after termination
//  7. State is local to the subscription
//  8. Completes synchronously when source completes synchronously
func TestDistinctUntilChanged_Sync(t *testing.T) {

	t.Run("suppresses only consecutive duplicates", func(t *testing.T) {
		src := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			o.Next(1)
			o.Next(1)
			o.Next(2)
			o.Next(2)
			o.Next(1)
			o.Next(1)
			o.Next(3)
			o.Complete()
			return subscr.Noop{}, nil
		})

		filtered := DistinctUntilChanged[int]()(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = filtered.Sub((*obs.Observer[int])(o))

		assertlite.EqualSlices(t, []int{1, 2, 1, 3}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("forwards all values when no consecutive duplicates exist", func(t *testing.T) {
		src := obs.Of(1, 2, 1, 2, 3)

		filtered := DistinctUntilChanged[int]()(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = filtered.Sub((*obs.Observer[int])(o))

		assertlite.EqualSlices(t, []int{1, 2, 1, 2, 3}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("propagates source error immediately", func(t *testing.T) {
		failure := errors.New("failure")

		src := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			o.Next(1)
			o.Next(1)
			o.Error(failure)
			o.Next(2) // must be ignored
			return subscr.Noop{}, nil
		})

		filtered := DistinctUntilChanged[int]()(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = filtered.Sub((*obs.Observer[int])(o))

		assertlite.EqualSlices(t, []int{1}, rec.Values())
		assertlite.IsErrorOf(t, rec.Err(), failure)
		assertlite.False(t, rec.Completed())
	})

	t.Run("does not share states between subscriptions", func(t *testing.T) {
		src := obs.Of(1, 1, 2)

		filtered := DistinctUntilChanged[int]()(src)

		rec1, o1 := observers.NewRecorder[int](t)
		rec2, o2 := observers.NewRecorder[int](t)

		_, _ = filtered.Sub((*obs.Observer[int])(o1))
		_, _ = filtered.Sub((*obs.Observer[int])(o2))

		assertlite.EqualSlices(t, []int{1, 2}, rec1.Values())
		assertlite.EqualSlices(t, []int{1, 2}, rec2.Values())
	})

	t.Run("completes synchronously when source completes synchronously", func(t *testing.T) {
		done := false

		src := obs.Of(1, 1)

		filtered := DistinctUntilChanged[int]()(src)

		o := &obs.Observer[int]{
			Next: func(int) {},
			Complete: func() {
				done = true
			},
		}

		_, _ = filtered.Sub(o)

		assertlite.True(t, done)
	})
}

// TestFilter_Sync tests Filter using synchronous sources.
//
// For synchronous sources, we must assert:
//
//  1. Values satisfying predicate are forwarded
//  2. Values failing predicate are suppressed
//  3. Order of values is preserved
//  4. Predicate is evaluated for each value until termination
//  5. Error from source is propagated immediately
//  6. Completion from source is forwarded
//  7. No values are emitted after termination
//  8. Predicate is not called after termination
//  9. State is local to the subscription
//  10. Completes synchronously when source completes synchronously
func TestFilter_Sync(t *testing.T) {
	t.Run("forwards only values that satisfy predicate", func(t *testing.T) {
		src := obs.Of(1, 2, 3, 4, 5)

		filtered := Filter(func(v int) bool {
			return v%2 == 0
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = filtered.Sub((*obs.Observer[int])(o))

		assertlite.EqualSlices(t, []int{2, 4}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("propagates source error immediately", func(t *testing.T) {
		failure := errors.New("failure")

		src := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			o.Next(1)
			o.Next(2)
			o.Error(failure)
			o.Next(4) // must be ignored
			return subscr.Noop{}, nil
		})

		filtered := Filter(func(v int) bool {
			return v%2 == 0
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = filtered.Sub((*obs.Observer[int])(o))

		assertlite.EqualSlices(t, []int{2}, rec.Values())
		assertlite.IsErrorOf(t, rec.Err(), failure)
		assertlite.False(t, rec.Completed())
	})

	t.Run("does not evaluate predicate after termination", func(t *testing.T) {
		failure := errors.New("failure")
		calls := 0

		src := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			o.Next(1)
			o.Error(failure)
			o.Next(2)
			return subscr.Noop{}, nil
		})

		filtered := Filter(func(v int) bool {
			calls++
			return true
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = filtered.Sub((*obs.Observer[int])(o))

		assertlite.Equals(t, 1, calls)
		assertlite.IsErrorOf(t, rec.Err(), failure)
	})

	t.Run("does not share states between subscriptions", func(t *testing.T) {
		src := obs.Of(1, 2, 3, 4)

		filtered := Filter(func(v int) bool {
			return v > 2
		})(src)

		rec1, o1 := observers.NewRecorder[int](t)
		rec2, o2 := observers.NewRecorder[int](t)

		_, _ = filtered.Sub((*obs.Observer[int])(o1))
		_, _ = filtered.Sub((*obs.Observer[int])(o2))

		assertlite.EqualSlices(t, []int{3, 4}, rec1.Values())
		assertlite.EqualSlices(t, []int{3, 4}, rec2.Values())
	})

	t.Run("completes synchronously when source completes synchronously", func(t *testing.T) {
		done := false

		src := obs.Of(1, 2)

		filtered := Filter(func(int) bool { return true })(src)

		o := &obs.Observer[int]{
			Next: func(int) {},
			Complete: func() {
				done = true
			},
		}

		_, _ = filtered.Sub(o)

		assertlite.True(t, done)
	})
}

// TestSkip_Sync tests Skip using synchronous sources.
//
// For synchronous sources, we must assert:
//
//  1. First n values are suppressed
//  2. Remaining values are forwarded unchanged
//  3. Order is preserved
//  4. If source completes before n, no values are emitted
//  5. Errors are propagated immediately
//  6. n <= 0 behaves as identity
//  7. State is local to the subscription
//  8. Completes synchronously when source completes synchronously
func TestSkip_Sync(t *testing.T) {
	t.Run("skips first n values and forwards the rest", func(t *testing.T) {
		src := obs.Of(1, 2, 3, 4, 5)

		skipped := Skip[int](2)(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = skipped.Sub((*obs.Observer[int])(o))

		assertlite.EqualSlices(t, []int{3, 4, 5}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("emits nothing if source completes before n", func(t *testing.T) {
		src := obs.Of(1, 2)

		skipped := Skip[int](2)(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = skipped.Sub((*obs.Observer[int])(o))

		assertlite.IsNil(t, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("propagates source error immediately", func(t *testing.T) {
		failure := errors.New("failure")

		src := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			o.Next(1)
			o.Next(2)
			o.Error(failure)
			o.Next(3) // must be ignored
			return subscr.Noop{}, nil
		})

		skipped := Skip[int](2)(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = skipped.Sub((*obs.Observer[int])(o))

		assertlite.IsNil(t, rec.Values())
		assertlite.IsErrorOf(t, rec.Err(), failure)
		assertlite.False(t, rec.Completed())
	})

	t.Run("n <= 0 forwards all values unchanged", func(t *testing.T) {
		src := obs.Of(1, 2, 3)

		skipped := Skip[int](-1)(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = skipped.Sub((*obs.Observer[int])(o))

		assertlite.EqualSlices(t, []int{1, 2, 3}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("does not share states between subscriptions", func(t *testing.T) {
		src := obs.Of(1, 2, 3)

		skipped := Skip[int](1)(src)

		rec1, o1 := observers.NewRecorder[int](t)
		rec2, o2 := observers.NewRecorder[int](t)

		_, _ = skipped.Sub((*obs.Observer[int])(o1))
		_, _ = skipped.Sub((*obs.Observer[int])(o2))

		assertlite.EqualSlices(t, []int{2, 3}, rec1.Values())
		assertlite.EqualSlices(t, []int{2, 3}, rec2.Values())
	})

	t.Run("completes synchronously when source completes synchronously", func(t *testing.T) {
		done := false

		src := obs.Of(1, 2)

		skipped := Skip[int](0)(src)

		o := &obs.Observer[int]{
			Next: func(int) {},
			Complete: func() {
				done = true
			},
		}

		_, _ = skipped.Sub(o)

		assertlite.True(t, done)
	})
}

// TestSkip_Async tests Skip using asynchronous sources.
//
// For asynchronous sources, we must assert:
//
//  1. First n values are suppressed
//  2. Remaining values are forwarded unchanged
//  3. Order is preserved
//  4. If source completes before n, no values are emitted
//  5. Errors are propagated immediately
//  6. n <= 0 behaves as identity
//  7. State is local to the subscription
//  8. Completion happens asynchronously (after subscription returns)
func TestSkip_Async(t *testing.T) {
	t.Run("skips first n values and forwards the rest", func(t *testing.T) {
		src := obs.AsyncIntervalOf(2*time.Millisecond, 1, 2, 3, 4, 5)

		skipped := Skip[int](2)(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = skipped.Sub((*obs.Observer[int])(o))

		rec.AwaitCompleted(1 * time.Second)

		assertlite.EqualSlices(t, []int{3, 4, 5}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("emits nothing if source completes before n", func(t *testing.T) {
		src := obs.AsyncIntervalOf(2*time.Millisecond, 1, 2)

		skipped := Skip[int](2)(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = skipped.Sub((*obs.Observer[int])(o))

		rec.AwaitCompleted(1 * time.Second)

		assertlite.IsNil(t, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("propagates source error immediately", func(t *testing.T) {
		failure := errors.New("failure")

		src := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			go func() {
				o.Next(1)
				o.Next(2)
				o.Error(failure)
				o.Next(3) // must be ignored
			}()
			return subscr.Noop{}, nil
		})

		skipped := Skip[int](2)(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = skipped.Sub((*obs.Observer[int])(o))

		assertlite.IsErrorOf(t, rec.AwaitError(1*time.Second), failure)
		assertlite.IsNil(t, rec.Values())
		assertlite.False(t, rec.Completed())
	})

	t.Run("n <= 0 forwards all values unchanged", func(t *testing.T) {
		src := obs.AsyncIntervalOf(2*time.Millisecond, 1, 2, 3)

		skipped := Skip[int](0)(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = skipped.Sub((*obs.Observer[int])(o))

		rec.AwaitCompleted(1 * time.Second)

		assertlite.EqualSlices(t, []int{1, 2, 3}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("does not share states between subscriptions", func(t *testing.T) {
		src := obs.AsyncIntervalOf(2*time.Millisecond, 1, 2, 3)

		skipped := Skip[int](1)(src)

		rec1, o1 := observers.NewRecorder[int](t)
		rec2, o2 := observers.NewRecorder[int](t)

		_, _ = skipped.Sub((*obs.Observer[int])(o1))
		_, _ = skipped.Sub((*obs.Observer[int])(o2))

		rec1.AwaitCompleted(1 * time.Second)
		rec2.AwaitCompleted(1 * time.Second)

		assertlite.EqualSlices(t, []int{2, 3}, rec1.Values())
		assertlite.EqualSlices(t, []int{2, 3}, rec2.Values())
	})

	t.Run("does not complete synchronously when source is async", func(t *testing.T) {
		src := obs.AsyncIntervalOf(2*time.Millisecond, 1, 2)

		skipped := Skip[int](0)(src)

		rec, o := observers.NewRecorder[int](t)

		_, _ = skipped.Sub((*obs.Observer[int])(o))

		// Must NOT be completed immediately
		assertlite.False(t, rec.Completed())
		rec.AwaitCompleted(1 * time.Second)
	})
}

// TestSkipUntil_Sync_Sync
//
// For synchronous source + synchronous signal, we want to assert:
//  1. Source values before signal fires are suppressed
//  2. Values after signal fires are forwarded
//  3. Wait.Next unlocks emissions
//  4. Wait.Complete also unlocks emissions
//  5. Wait.Error propagates immediately
//  6. Source.Error propagates immediately
//  7. If source completes before signal fires, no values are emitted
//  8. Wait subscription is canceled once the gate opens
//  9. State is local per subscription
func TestSkipUntil_Sync_Sync(t *testing.T) {
	t.Run("skips source values until signal emits", func(t *testing.T) {
		src := obs.Of(1, 2, 3, 4, 5)
		sig := obs.Of[any](nil) // fires immediately after subscription

		out := SkipUntil[int](sig)(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		// signal fires before any source emission
		assertlite.EqualSlices(t, []int{1, 2, 3, 4, 5}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("suppresses values until signal fires mid-stream", func(t *testing.T) {
		src := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			o.Next(1)
			o.Next(2)
			return subscr.Noop{}, nil
		})

		sig := obs.Of[any](nil)

		out := SkipUntil[int](sig)(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		// signal fires before src emits (sync)
		assertlite.EqualSlices(t, []int{1, 2}, rec.Values())
	})

	t.Run("emits nothing if source completes before signal", func(t *testing.T) {
		src := obs.Of(1, 2, 3, 4, 5)
		sig := obs.Empty[any]()

		out := SkipUntil[int](sig)(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		assertlite.EqualSlices(t, []int{1, 2, 3, 4, 5}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("signal completion unlocks emissions", func(t *testing.T) {
		src := obs.Of(1, 2, 3)
		sig := obs.Create[any](func(o *obs.Observer[any]) (obs.Subscription, error) {
			o.Complete()
			return subscr.Noop{}, nil
		})

		out := SkipUntil[int](sig)(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		assertlite.EqualSlices(t, []int{1, 2, 3}, rec.Values())
		assertlite.True(t, rec.Completed())
	})

	t.Run("propagates signal error immediately", func(t *testing.T) {
		failure := errors.New("signal failed")

		src := obs.Of(1, 2, 3)
		sig := obs.Error[any](failure)

		out := SkipUntil[int](sig)(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		assertlite.IsErrorOf(t, rec.Err(), failure)
		assertlite.IsNil(t, rec.Values())
		assertlite.False(t, rec.Completed())
	})

	t.Run("propagates source error immediately", func(t *testing.T) {
		failure := errors.New("source failed")

		src := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			o.Next(1)
			o.Error(failure)
			o.Next(2) // must be ignored
			return subscr.Noop{}, nil
		})

		sig := obs.Of[any](nil)

		out := SkipUntil[int](sig)(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		assertlite.IsErrorOf(t, rec.Err(), failure)
		assertlite.EqualSlices(t, []int{1}, rec.Values())
		assertlite.False(t, rec.Completed())
	})

	t.Run("does not share states between subscriptions", func(t *testing.T) {
		src := obs.Of(1, 2, 3)
		sig := obs.Of[any](nil)

		out := SkipUntil[int](sig)(src)

		rec1, o1 := observers.NewRecorder[int](t)
		rec2, o2 := observers.NewRecorder[int](t)

		_, _ = out.Sub((*obs.Observer[int])(o1))
		_, _ = out.Sub((*obs.Observer[int])(o2))

		assertlite.EqualSlices(t, []int{1, 2, 3}, rec1.Values())
		assertlite.EqualSlices(t, []int{1, 2, 3}, rec2.Values())
	})
}

// TestSkipUntil_Sync_Sync
//
// For sync source + async signal, we care about:
//  1. All source emissions before signal fires are suppressed
//  2. Source completion before signal fires it completes with no values
//  3. If signal fires after subscription but before source emits the values pass
//  4. Wait error propagates immediately
//  5. Wait firing after source completion has no effect
//  6. No race leaks (no values sneaking through)
func TestSkipUntil_Sync_Async(t *testing.T) {
	t.Run("suppresses all values if async signal fires after source completes", func(t *testing.T) {
		sig := obs.AsyncIntervalOf[any](100*time.Millisecond, nil) // fires later
		src := obs.Of(1, 2, 3)
		out := SkipUntil[int](sig)(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		// source completed before signal
		assertlite.IsNil(t, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("unlocks emissions if async signal fires before source emits", func(t *testing.T) {
		start := make(chan struct{})

		sig := obs.Async[any](func(o *obs.Observer[any]) {
			o.Next(nil)
			close(start)
		})

		src := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			<-start
			o.Next(1)
			o.Next(2)
			o.Complete()
			return subscr.Noop{}, nil
		})

		out := SkipUntil[int](sig)(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))
		rec.AwaitCompleted(time.Second)

		assertlite.EqualSlices(t, []int{1, 2}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("does not emit values if source emits before async signal fires", func(t *testing.T) {
		start := make(chan struct{})

		sig := obs.Async[any](func(o *obs.Observer[any]) {
			<-start
			o.Next(nil)
		})

		src := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			o.Next(1)
			o.Next(2)
			o.Complete()
			close(start)
			return subscr.Noop{}, nil
		})

		out := SkipUntil[int](sig)(src)
		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		rec.AwaitCompleted(time.Second)

		assertlite.IsNil(t, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("propagates async signal error immediately", func(t *testing.T) {
		failure := errors.New("signal failed")
		stop := make(chan struct{})

		sig := obs.Async[any](func(o *obs.Observer[any]) {
			time.Sleep(5 * time.Millisecond)
			o.Error(failure)
			close(stop)
		})

		src := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			o.Next(1)
			o.Next(2)
			<-stop // block for signal
			o.Next(3)
			o.Complete()
			return subscr.Noop{}, nil
		})

		out := SkipUntil[int](sig)(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		assertlite.IsErrorOf(t, rec.AwaitError(1*time.Second), failure)
		assertlite.IsNil(t, rec.Values())
		assertlite.False(t, rec.Completed())
	})
}

// TestSkipUntil_Async_Sync
//
// we must assert For Async source + Sync signal:
//  1. Wait fires immediately → no values are skipped
//  2. All async source values are forwarded
//  3. Completion is forwarded
//  4. Wait completion also unlocks
//  5. Wait error propagates immediately
//  6. Wait subscription is cancelled after firing
func TestSkipUntil_Async_Sync(t *testing.T) {
	t.Run("sync signal unlocks async source immediately", func(t *testing.T) {
		sig := obs.Of[any](nil)

		src := obs.AsyncOf(1, 2, 3)

		out := SkipUntil[int](sig)(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		rec.AwaitTermination(time.Second)

		assertlite.EqualSlices(t, []int{1, 2, 3}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("sync signal completion unlocks async source", func(t *testing.T) {
		sig := obs.Empty[any]() // completes immediately

		src := obs.AsyncOf(10, 20)

		out := SkipUntil[int](sig)(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		rec.AwaitTermination(time.Second)

		assertlite.EqualSlices(t, []int{10, 20}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("signal error propagates immediately and cancels async source", func(t *testing.T) {
		failure := errors.New("err")

		sig := obs.Error[any](failure)

		src := obs.AsyncOf(1, 2, 3)

		out := SkipUntil[int](sig)(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		assertlite.IsErrorOf(t, rec.AwaitError(time.Second), failure)
		assertlite.IsNil(t, rec.Values())
		assertlite.False(t, rec.Completed())
	})

	t.Run("signal fires before async source emits any value", func(t *testing.T) {
		start := make(chan struct{})

		sig := obs.Create[any](func(o *obs.Observer[any]) (obs.Subscription, error) {
			o.Next(1)
			close(start)
			return subscr.Noop{}, nil
		})

		src := obs.Async[int](func(o *obs.Observer[int]) {
			<-start
			o.Next(8)
			o.Complete()
		})

		out := SkipUntil[int](sig)(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		rec.AwaitCompleted(time.Second)

		assertlite.EqualSlices(t, []int{8}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})
}

// TestSkipUntil_Async_Async must prove
//
//  1. Both source and signal are async, so we must validate:
//  2. Values before signal fires are skipped
//  3. Values after signal fires are forwarded
//  4. Order is preserved
//  5. Wait Next unlocks
//  6. Wait Complete unlocks
//  7. Wait Error errors immediately
//  8. Source Error errors immediately
//  9. Source completion before signal has no emissions
//  10. Wait subscription is cancelled after unlock
//  11. No goroutine leaks / races
func TestSkipUntil_Async_Async(t *testing.T) {
	t.Run("skips values until async signal fires", func(t *testing.T) {
		srcCh := make(chan int)
		sigCh := make(chan struct{})
		sigFired := make(chan struct{})

		src := obs.Async[int](func(o *obs.Observer[int]) {
			for v := range srcCh {
				o.Next(v)
			}
			o.Complete()
		})

		sig := obs.Async[any](func(o *obs.Observer[any]) {
			<-sigCh
			o.Next(nil)
			// avoid source emit x signal next race
			close(sigFired)
			o.Complete()
		})

		out := SkipUntil[int](sig)(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		// before signal
		srcCh <- 1
		srcCh <- 2
		srcCh <- 3

		// unlock
		close(sigCh)
		<-sigFired

		// after signal
		srcCh <- 4
		srcCh <- 5

		close(srcCh)

		rec.AwaitTermination(time.Second)

		assertlite.EqualSlices(t, []int{4, 5}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("signal completion unlocks async source", func(t *testing.T) {
		srcCh := make(chan int)
		src := obs.AsyncChan[int](srcCh)
		sig := obs.AsyncEmpty[any]()

		out := SkipUntil[int](sig)(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		// this delay will allow the signal to complete
		time.Sleep(10 * time.Millisecond)
		srcCh <- 10
		srcCh <- 20
		close(srcCh)

		rec.AwaitCompleted(time.Second)

		assertlite.EqualSlices(t, []int{10, 20}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("signal error propagates immediately", func(t *testing.T) {
		failure := errors.New("signal failure")
		srcCh := make(chan int)

		src := obs.AsyncChan[int](srcCh)
		sig := obs.AsyncError[any](failure)
		out := SkipUntil[int](sig)(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		srcCh <- 1
		srcCh <- 2

		assertlite.IsErrorOf(t, rec.AwaitError(time.Second), failure)
		assertlite.IsNil(t, rec.Values())
		assertlite.False(t, rec.Completed())
	})

	t.Run("source error propagates immediately even before signal", func(t *testing.T) {
		failure := errors.New("source failure")

		src := obs.Async[int](func(o *obs.Observer[int]) {
			o.Next(1)
			o.Error(failure)
		})

		sig := obs.Async[any](func(o *obs.Observer[any]) {
			// noop
		})

		out := SkipUntil[int](sig)(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		assertlite.IsErrorOf(t, rec.AwaitError(time.Second), failure)
		assertlite.IsNil(t, rec.Values())
		assertlite.False(t, rec.Completed())
	})

	t.Run("source completes before signal emits", func(t *testing.T) {
		src := obs.Async[int](func(o *obs.Observer[int]) {
			o.Next(1)
			o.Next(2)
			o.Complete()
		})

		sig := obs.Async[any](func(o *obs.Observer[any]) {
			// noop
		})

		out := SkipUntil[int](sig)(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		rec.AwaitTermination(time.Second)

		assertlite.IsNil(t, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})
}

// TestSkipWhile_Sync
//
// For synchronous source, we must assert:
//
//  1. Values are skipped while pred(v) returns true
//  2. The first value for which pred(v) == false is emitted
//  3. Once skipping ends, all subsequent values are forwarded
//  4. Predicate is not evaluated after skipping ends
//  5. If source completes while pred stays true, no values are emitted
//  6. Source completion is forwarded once skipping ends
//  7. Source error is propagated immediately
//  8. Predicate is evaluated in source emission order
//  9. State is local per subscription
func TestSkipWhile_Sync(t *testing.T) {
	t.Run("skips values while predicate is true", func(t *testing.T) {
		src := obs.Of(1, 2, 3, 4, 5)

		out := SkipWhile(func(v int) bool {
			return v < 4
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		rec.AwaitTermination(time.Second)

		assertlite.EqualSlices(t, []int{4, 5}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("emits first value when predicate becomes false", func(t *testing.T) {
		src := obs.Of(10, 20, 30)

		out := SkipWhile(func(v int) bool {
			return v < 20
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		rec.AwaitTermination(time.Second)

		assertlite.EqualSlices(t, []int{20, 30}, rec.Values())
	})

	t.Run("completes without values if source completes while skipping", func(t *testing.T) {
		src := obs.Of(1, 2, 3)

		out := SkipWhile(func(v int) bool {
			return true
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		rec.AwaitTermination(time.Second)

		assertlite.IsNil(t, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("predicate is not evaluated after skipping ends", func(t *testing.T) {
		evals := 0

		src := obs.Of(1, 2, 3, 4, 5)

		out := SkipWhile(func(v int) bool {
			evals++
			return v < 3
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		rec.AwaitTermination(time.Second)

		assertlite.Equals(t, 3, evals) // only 1,2,3 evaluated
		assertlite.EqualSlices(t, []int{3, 4, 5}, rec.Values())
	})

	t.Run("propagates source error immediately", func(t *testing.T) {
		err := errors.New("err")

		src := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			o.Next(1)
			o.Error(err)
			return subscr.Noop{}, nil
		})

		out := SkipWhile(func(v int) bool {
			return true
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		rec.AwaitTermination(time.Second)

		assertlite.IsNil(t, rec.Values())
		assertlite.False(t, rec.Completed())
		assertlite.IsErrorOf(t, err, rec.Err())
	})

	t.Run("states is local per subscription", func(t *testing.T) {
		src := obs.Of(1, 2, 3, 4)

		out := SkipWhile(func(v int) bool {
			return v < 3
		})(src)

		rec1, o1 := observers.NewRecorder[int](t)
		rec2, o2 := observers.NewRecorder[int](t)

		_, _ = out.Sub((*obs.Observer[int])(o1))
		_, _ = out.Sub((*obs.Observer[int])(o2))

		rec1.AwaitTermination(time.Second)
		rec2.AwaitTermination(time.Second)

		assertlite.EqualSlices(t, []int{3, 4}, rec1.Values())
		assertlite.EqualSlices(t, []int{3, 4}, rec2.Values())
	})
}

// TestSkipWhile_Async
//
// For asynchronous source, we must assert:
//
//  1. Values emitted before pred(v) becomes false are skipped
//  2. The first value for which pred(v) == false is emitted
//  3. All subsequent values are forwarded
//  4. Order is preserved
//  5. Predicate is evaluated only until it returns false
//  6. Source completion is forwarded correctly
//  7. Source error is propagated immediately
//  8. Unsubscription stops source emission
//  9. State is local per subscription
//
// 10. No races or goroutine leaks
func TestSkipWhile_Async(t *testing.T) {
	t.Run("skips values until predicate becomes false", func(t *testing.T) {
		srcCh := make(chan int)

		src := obs.AsyncChan(srcCh)

		out := SkipWhile(func(v int) bool {
			return v < 4
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		srcCh <- 1
		srcCh <- 2
		srcCh <- 3
		srcCh <- 4
		srcCh <- 5

		close(srcCh)

		rec.AwaitTermination(time.Second)

		assertlite.EqualSlices(t, []int{4, 5}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("completes without values if source completes while skipping", func(t *testing.T) {
		srcCh := make(chan int)

		src := obs.AsyncChan(srcCh)

		out := SkipWhile(func(v int) bool {
			return true
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		srcCh <- 1
		srcCh <- 2
		srcCh <- 3
		close(srcCh)

		rec.AwaitTermination(time.Second)

		assertlite.IsNil(t, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("predicate is evaluated only until it returns false", func(t *testing.T) {
		var evals atomic.Int32

		srcCh := make(chan int)

		src := obs.AsyncChan(srcCh)

		out := SkipWhile(func(v int) bool {
			evals.Add(1)
			return v < 3
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		srcCh <- 1
		srcCh <- 2
		srcCh <- 3
		srcCh <- 4
		srcCh <- 5
		close(srcCh)

		rec.AwaitTermination(time.Second)

		assertlite.Equals(t, int32(3), evals.Load())
		assertlite.EqualSlices(t, []int{3, 4, 5}, rec.Values())
	})

	t.Run("propagates source error immediately", func(t *testing.T) {
		err := errors.New("err")

		src := obs.Async[int](func(o *obs.Observer[int]) {
			o.Next(1)
			o.Error(err)
		})

		out := SkipWhile(predicates.True[int])(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		rec.AwaitTermination(time.Second)

		assertlite.IsNil(t, rec.Values())
		assertlite.False(t, rec.Completed())
		assertlite.IsErrorOf(t, err, rec.Err())
	})

	t.Run("states is local per subscription", func(t *testing.T) {
		srcCh := make(chan int)

		src := obs.AsyncMulticast(srcCh)

		out := SkipWhile(func(v int) bool {
			return v < 3
		})(src)

		rec1, o1 := observers.NewRecorder[int](t)
		rec2, o2 := observers.NewRecorder[int](t)

		_, _ = out.Sub((*obs.Observer[int])(o1))
		_, _ = out.Sub((*obs.Observer[int])(o2))

		srcCh <- 1
		srcCh <- 2
		srcCh <- 3
		srcCh <- 4
		close(srcCh)

		rec1.AwaitTermination(time.Second)
		rec2.AwaitTermination(time.Second)

		assertlite.EqualSlices(t, []int{3, 4}, rec1.Values())
		assertlite.EqualSlices(t, []int{3, 4}, rec2.Values())
	})
}

// TestTake_Sync
//
// For synchronous sources, we must assert:
//  1. Emits exactly the first n values
//  2. Completes immediately after n values
//  3. Preserves order
//  4. Does not emit values after completion
//  5. Forwards early source completion correctly
//  6. Propagates source error immediately
//  7. Does not rely on async behavior
//  8. Handles n <= 0 correctly
//  9. State is local per subscription
func TestTake_Sync(t *testing.T) {
	t.Run("takes exactly n values and completes", func(t *testing.T) {
		src := obs.Of(1, 2, 3, 4, 5)

		out := Take[int](3)(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		rec.AwaitCompleted(time.Second)

		assertlite.EqualSlices(t, []int{1, 2, 3}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("preserves order", func(t *testing.T) {
		src := obs.Of(5, 4, 3, 2, 1)

		out := Take[int](3)(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		rec.AwaitCompleted(time.Second)

		assertlite.EqualSlices(t, []int{5, 4, 3}, rec.Values())
	})

	t.Run("completes early if source completes before n", func(t *testing.T) {
		src := obs.Of(1, 2)

		out := Take[int](3)(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		rec.AwaitCompleted(time.Second)

		assertlite.EqualSlices(t, []int{1, 2}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("propagates source error immediately", func(t *testing.T) {
		err := errors.New("fail")

		src := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			o.Next(1)
			o.Next(2)
			o.Error(err)
			return subscr.Noop{}, nil
		})

		out := Take[int](3)(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		rec.AwaitTermination(time.Second)

		assertlite.EqualSlices(t, []int{1, 2}, rec.Values())
		assertlite.False(t, rec.Completed())
		assertlite.IsErrorOf(t, err, rec.Err())
	})

	t.Run("skip source error immediately after n is reached", func(t *testing.T) {
		err := errors.New("fail")

		src := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			o.Next(1)
			o.Next(2)
			o.Error(err)
			return subscr.Noop{}, nil
		})

		out := Take[int](2)(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		rec.AwaitTermination(time.Second)

		assertlite.EqualSlices(t, []int{1, 2}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("n == 0 completes immediately without subscribing", func(t *testing.T) {
		called := atomic.Bool{}

		src := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			called.Store(true)
			o.Next(1)
			return subscr.Noop{}, nil
		})

		out := Take[int](0)(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		rec.AwaitCompleted(time.Second)

		assertlite.IsNil(t, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
		assertlite.False(t, called.Load())
	})

	t.Run("n < 0 completes immediately without subscribing", func(t *testing.T) {
		called := atomic.Bool{}

		src := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			called.Store(true)
			o.Next(1)
			return subscr.Noop{}, nil
		})

		out := Take[int](-5)(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		rec.AwaitCompleted(time.Second)

		assertlite.IsNil(t, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
		assertlite.False(t, called.Load())
	})

	t.Run("states is local per subscription", func(t *testing.T) {
		src := obs.Of(1, 2, 3, 4, 5)

		out := Take[int](2)(src)

		rec1, o1 := observers.NewRecorder[int](t)
		rec2, o2 := observers.NewRecorder[int](t)

		_, _ = out.Sub((*obs.Observer[int])(o1))
		_, _ = out.Sub((*obs.Observer[int])(o2))

		rec1.AwaitCompleted(time.Second)
		rec2.AwaitCompleted(time.Second)

		assertlite.EqualSlices(t, []int{1, 2}, rec1.Values())
		assertlite.EqualSlices(t, []int{1, 2}, rec2.Values())
	})
}

// TestTake_Async
//
// For Take(n) + async source:
//  1. Emits at most n values
//  2. Emits exactly n values if source produces ≥ n
//  3. Completes immediately after the nth value
//  4. Preserves order
//  5. Does not forward errors after n
//  6. Forwards errors if they occur before n
//  7. Unsubscription stops further emissions
//  8. State is local per subscription
//  9. Works with single-cast async (AsyncChan)
//  10. Works with multicast async (AsyncMulticast)
//  11. No goroutine leaks / deadlocks
func TestTake_Async(t *testing.T) {
	t.Run("takes first n values and completes", func(t *testing.T) {
		srcCh := make(chan int)

		src := obs.AsyncChan(srcCh)
		out := Take[int](3)(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		go func() {
			srcCh <- 1
			srcCh <- 2
			srcCh <- 3
			srcCh <- 4
			srcCh <- 5
			close(srcCh)
		}()

		rec.AwaitTermination(time.Second)

		assertlite.EqualSlices(t, []int{1, 2, 3}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("completes early if source completes before n", func(t *testing.T) {
		srcCh := make(chan int)

		src := obs.AsyncChan(srcCh)
		out := Take[int](5)(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		go func() {
			srcCh <- 1
			srcCh <- 2
			close(srcCh)
		}()

		rec.AwaitTermination(time.Second)

		assertlite.EqualSlices(t, []int{1, 2}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("propagates source error before n", func(t *testing.T) {
		err := errors.New("fail")

		src := obs.Async[int](func(o *obs.Observer[int]) {
			o.Next(1)
			o.Error(err)
		})

		out := Take[int](3)(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		rec.AwaitTermination(time.Second)

		assertlite.EqualSlices(t, []int{1}, rec.Values())
		assertlite.False(t, rec.Completed())
		assertlite.IsErrorOf(t, err, rec.Err())
	})

	t.Run("ignores source error after n is reached", func(t *testing.T) {
		err := errors.New("fail")

		src := obs.Async[int](func(o *obs.Observer[int]) {
			o.Next(1)
			o.Next(2)
			o.Error(err)
		})

		out := Take[int](2)(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		rec.AwaitTermination(time.Second)

		assertlite.EqualSlices(t, []int{1, 2}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("states is local per subscription (multicast)", func(t *testing.T) {
		srcCh := make(chan int)

		src := obs.AsyncMulticast(srcCh)
		out := Take[int](2)(src)

		rec1, o1 := observers.NewRecorder[int](t)
		rec2, o2 := observers.NewRecorder[int](t)

		_, _ = out.Sub((*obs.Observer[int])(o1))
		_, _ = out.Sub((*obs.Observer[int])(o2))

		go func() {
			srcCh <- 1
			srcCh <- 2
			srcCh <- 3
			close(srcCh)
		}()

		rec1.AwaitTermination(time.Second)
		rec2.AwaitTermination(time.Second)

		assertlite.EqualSlices(t, []int{1, 2}, rec1.Values())
		assertlite.EqualSlices(t, []int{1, 2}, rec2.Values())
	})
}

// TestTakeUntil_Sync_Sync
//
// Should assert:
//  1. Values before the signal fires are forwarded
//  2. Wait Next stops the source and completes
//  3. Wait Complete stops the source and completes
//  4. Wait Error errors immediately
//  5. Source Error errors immediately
//  6. Source completion before signal is forwarded
//  7. Wait firing before source emits produces no values
//  8. Completion happens at most once
//  9. State is local per subscription
func TestTakeUntil_Sync_Sync(t *testing.T) {
	t.Run("emits values until signal emits Next", func(t *testing.T) {
		src := obs.Of(1, 2, 3, 4, 5)

		signal := obs.FromFunc[any](func(o *obs.Observer[any]) {
			o.Next(nil)
		})

		out := TakeUntil[int](signal)(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		rec.AwaitCompleted(time.Second)

		assertlite.IsNil(t, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("signal Complete stops the source", func(t *testing.T) {
		src := obs.Of(1, 2, 3, 4)

		signal := obs.FromFunc[any](func(o *obs.Observer[any]) {
			o.Complete()
		})

		out := TakeUntil[int](signal)(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		rec.AwaitCompleted(time.Second)

		assertlite.IsNil(t, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("signal Error propagates immediately", func(t *testing.T) {
		err := errors.New("signal fail")

		src := obs.Of(1, 2, 3)

		signal := obs.Error[any](err)

		out := TakeUntil[int](signal)(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		rec.AwaitTermination(time.Second)

		assertlite.IsNil(t, rec.Values())
		assertlite.False(t, rec.Completed())
		assertlite.IsErrorOf(t, err, rec.Err())
	})

	t.Run("source Error propagates immediately", func(t *testing.T) {
		err := errors.New("src fail")

		src := obs.Error[int](err)
		signal := obs.Never[any]()

		out := TakeUntil[int](signal)(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		rec.AwaitTermination(time.Second)

		assertlite.IsNil(t, rec.Values())
		assertlite.False(t, rec.Completed())
		assertlite.IsErrorOf(t, err, rec.Err())
	})

	t.Run("source completion before signal is forwarded", func(t *testing.T) {
		src := obs.Of(1, 2)
		signal := obs.Never[any]()

		out := TakeUntil[int](signal)(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		rec.AwaitCompleted(time.Second)

		assertlite.EqualSlices(t, []int{1, 2}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("states is local per subscription", func(t *testing.T) {
		signal := obs.Of[any](nil)

		src := obs.Of(1, 2, 3)

		out := TakeUntil[int](signal)(src)

		rec1, o1 := observers.NewRecorder[int](t)
		rec2, o2 := observers.NewRecorder[int](t)

		_, _ = out.Sub((*obs.Observer[int])(o1))
		_, _ = out.Sub((*obs.Observer[int])(o2))

		rec1.AwaitCompleted(time.Second)
		rec2.AwaitCompleted(time.Second)

		assertlite.IsNil(t, rec1.Values())
		assertlite.IsNil(t, rec2.Values())
	})
}

// TestTakeUntil_Sync_Async
//
// Sync source + Async signal semantics.
//
// In this configuration, the source Observable emits synchronously
// during subscription, while the signal Observable emits asynchronously
// from a separate goroutine.
//
// As a consequence, synchronous source emissions are not preemptible
// by the signal: all source values emitted during Sub will be observed
// before any asynchronous signal event can occur.
//
// This test suite asserts the following guarantees:
//
//  1. Wait Error propagates immediately if it occurs before termination.
//     Once an error is observed, all subscriptions are cancelled.
//
//  2. Source Error propagates immediately and takes precedence over the signal.
//
//  3. Source completion before the signal fires is forwarded downstream
//     with all emitted values preserved.
//
//  4. Completion happens at most once, regardless of the ordering or
//     timing of asynchronous signal commands.
//
//  5. State is local per subscription: each subscription maintains its
//     own termination states and independently observes the signal.
//
// Non-guarantees (by design):
//
//   - The signal cannot prevent or preempt synchronous source emissions.
//   - Wait Next or Complete occurring after subscription may arrive
//     after the source has already completed.
//   - No ordering guarantees are made between asynchronous signal commands
//     and synchronous source emissions.
//
// Notes:
//
//   - Sleeping or emitting from the signal before Sub has no effect;
//     both source and signal are cold and begin execution only upon
//     subscription.
//   - This behavior establishes a clear happens-before boundary where
//     synchronous sources dominate asynchronous signals.
//   - Tests asserting preemption or early termination by the signal
//     belong in the Async_Async suite, not here.
func TestTakeUntil_Sync_Async(t *testing.T) {
	t.Run("signal Error propagates immediately", func(t *testing.T) {
		err := errors.New("fail")

		signal := obs.AsyncError[any](err)

		src := obs.Range(1, 1_000_000, 1)

		out := TakeUntil[int](signal)(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		rec.AwaitTermination(time.Second)

		assertlite.False(t, rec.Completed())
		assertlite.IsErrorOf(t, err, rec.Err())
	})

	t.Run("source Error propagates immediately", func(t *testing.T) {
		err := errors.New("src fail")

		src := obs.Error[int](err)
		signal := obs.Never[any]()

		out := TakeUntil[int](signal)(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		rec.AwaitTermination(time.Second)

		assertlite.False(t, rec.Completed())
		assertlite.IsErrorOf(t, err, rec.Err())
	})

	t.Run("source completion before signal is forwarded", func(t *testing.T) {
		signal := obs.AsyncIntervalOf[any](50*time.Millisecond, nil)

		src := obs.Of(1, 2)

		out := TakeUntil[int](signal)(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		rec.AwaitCompleted(time.Second)

		assertlite.EqualSlices(t, []int{1, 2}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("completion happens at most once", func(t *testing.T) {
		signal := obs.AsyncIntervalOf[any](time.Millisecond)

		src := obs.Of(1, 2)

		out := TakeUntil[int](signal)(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		rec.AwaitCompleted(time.Second)

		assertlite.True(t, rec.Completed())
	})

	t.Run("states is local per subscription", func(t *testing.T) {
		signal := obs.AsyncIntervalOf[any](50 * time.Millisecond)
		time.Sleep(2 * time.Millisecond)

		src := obs.Of(1, 2, 3)

		out := TakeUntil[int](signal)(src)

		rec1, o1 := observers.NewRecorder[int](t)
		rec2, o2 := observers.NewRecorder[int](t)

		_, _ = out.Sub((*obs.Observer[int])(o1))
		_, _ = out.Sub((*obs.Observer[int])(o2))

		rec1.AwaitCompleted(time.Second)
		rec2.AwaitCompleted(time.Second)

		assertlite.True(t, rec1.Completed())
		assertlite.True(t, rec2.Completed())
	})
}

// TestTakeUntil_Async_Sync
//
// Async source + Sync signal semantics.
//
// The signal Observable emits synchronously during subscription,
// while the source emits asynchronously from another goroutine.
//
// In this configuration, the signal can deterministically preempt
// the source before any values are emitted.
//
// Should assert:
//
//  1. Wait Next before source emits produces no values
//  2. Wait Complete before source emits produces no values
//  3. Wait Error propagates immediately
//  4. Source Error propagates if it occurs before signal
//  5. Source completion before signal is forwarded
//  6. Completion happens at most once
//  7. State is local per subscription
//  8. Async source is cancelled when signal fires
func TestTakeUntil_Async_Sync(t *testing.T) {
	t.Run("signal Next before source emits produces no values", func(t *testing.T) {
		signal := obs.Of[any](nil) // fires synchronously

		src := obs.AsyncIntervalOf[int](10*time.Millisecond, 1, 2, 3)

		out := TakeUntil[int](signal)(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		rec.AwaitCompleted(time.Second)

		assertlite.IsNil(t, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("signal Complete before source emits produces no values", func(t *testing.T) {
		signal := obs.Empty[any]() // completes synchronously

		src := obs.AsyncIntervalOf[int](10*time.Millisecond, 1, 2, 3)

		out := TakeUntil[int](signal)(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		rec.AwaitCompleted(time.Second)

		assertlite.IsNil(t, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("signal Error propagates immediately", func(t *testing.T) {
		err := errors.New("signal fail")

		signal := obs.Error[any](err)

		src := obs.AsyncIntervalOf[int](10*time.Millisecond, 1, 2, 3)

		out := TakeUntil[int](signal)(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		rec.AwaitTermination(time.Second)

		assertlite.False(t, rec.Completed())
		assertlite.IsErrorOf(t, err, rec.Err())
	})

	t.Run("source Error propagates if it occurs before signal", func(t *testing.T) {
		err := errors.New("src fail")

		src := obs.Async[int](func(o *obs.Observer[int]) {
			time.Sleep(5 * time.Millisecond)
			o.Error(err)
		})

		signal := obs.Never[any]()

		out := TakeUntil[int](signal)(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		rec.AwaitTermination(time.Second)

		assertlite.False(t, rec.Completed())
		assertlite.IsErrorOf(t, err, rec.Err())
	})

	t.Run("source completion before signal is forwarded", func(t *testing.T) {
		src := obs.AsyncIntervalOf[int](5*time.Millisecond, 1, 2)

		signal := obs.Never[any]()

		out := TakeUntil[int](signal)(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		rec.AwaitCompleted(time.Second)

		assertlite.EqualSlices(t, []int{1, 2}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("completion happens at most once", func(t *testing.T) {
		signal := obs.Of[any](nil)

		src := obs.AsyncIntervalOf[int](time.Millisecond, 1, 2, 3)

		out := TakeUntil[int](signal)(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		rec.AwaitCompleted(time.Second)

		assertlite.True(t, rec.Completed())
	})

	t.Run("states is local per subscription", func(t *testing.T) {
		signal := obs.Of[any](nil)

		src := obs.AsyncIntervalOf[int](10*time.Millisecond, 1, 2, 3)

		out := TakeUntil[int](signal)(src)

		rec1, o1 := observers.NewRecorder[int](t)
		rec2, o2 := observers.NewRecorder[int](t)

		_, _ = out.Sub((*obs.Observer[int])(o1))
		_, _ = out.Sub((*obs.Observer[int])(o2))

		rec1.AwaitCompleted(time.Second)
		rec2.AwaitCompleted(time.Second)

		assertlite.IsNil(t, rec1.Values())
		assertlite.IsNil(t, rec2.Values())
	})
}

// TestTakeUntil_Async_Async
//
// Should assert:
//  1. Values emitted by the source before the signal fires are forwarded
//  2. Wait Next stops the source and completes
//  3. Wait Complete stops the source and completes
//  4. Wait Error errors immediately
//  5. Source Error errors immediately
//  6. Source completion before the signal fires is forwarded
//  7. Wait firing before source emits produces no values
//  8. Completion happens at most once under concurrent races
//  9. State is local per subscription
//
// Notes:
//   - Source and signal may emit concurrently
//   - Races between source, signal, and terminal commands must be handled safely
//   - No value, completion, or error may be delivered after termination
func TestTakeUntil_Async_Async(t *testing.T) {
	t.Run("values before signal fires are forwarded", func(t *testing.T) {
		signal := obs.AsyncIntervalOf[any](30 * time.Millisecond)
		src := obs.AsyncIntervalOf(10*time.Millisecond, 1, 2, 3, 4, 5)

		out := TakeUntil[int](signal)(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		rec.AwaitCompleted(time.Second)

		values := rec.Values()
		assertlite.True(t, len(values) > 0, "expected to find at least one value, got none")
		assertlite.True(t, len(values) < 5, "expected to find at least 5 values, got %d", len(values))
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("signal Next stops source and completes", func(t *testing.T) {
		signal := obs.AsyncOf[any](nil)
		src := obs.AsyncIntervalOf(10*time.Millisecond, 1, 2, 3)

		out := TakeUntil[int](signal)(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		rec.AwaitCompleted(time.Second)

		assertlite.IsNil(t, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("signal Complete stops source and completes", func(t *testing.T) {
		signal := obs.AsyncEmpty[any]()
		src := obs.AsyncIntervalOf(10*time.Millisecond, 1, 2, 3)

		out := TakeUntil[int](signal)(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		rec.AwaitCompleted(time.Second)

		assertlite.IsNil(t, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("signal Error errors immediately", func(t *testing.T) {
		err := errors.New("signal fail")

		signal := obs.AsyncError[any](err)
		src := obs.AsyncIntervalOf(10*time.Millisecond, 1, 2, 3)

		out := TakeUntil[int](signal)(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		rec.AwaitTermination(time.Second)

		assertlite.False(t, rec.Completed())
		assertlite.IsErrorOf(t, err, rec.Err())
	})

	t.Run("source Error errors immediately", func(t *testing.T) {
		err := errors.New("src fail")

		signal := obs.AsyncIntervalOf[any](50 * time.Millisecond)
		src := obs.AsyncError[int](err)

		out := TakeUntil[int](signal)(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		rec.AwaitTermination(time.Second)

		assertlite.False(t, rec.Completed())
		assertlite.IsErrorOf(t, err, rec.Err())
	})

	t.Run("source completion before signal is forwarded", func(t *testing.T) {
		signal := obs.AsyncIntervalOf[any](50 * time.Millisecond)
		src := obs.AsyncOf(1, 2)

		out := TakeUntil[int](signal)(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		rec.AwaitCompleted(time.Second)

		assertlite.EqualSlices(t, []int{1, 2}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("signal firing before source emits produces no values", func(t *testing.T) {
		signal := obs.AsyncOf[any](nil)
		src := obs.AsyncIntervalOf(50*time.Millisecond, 1, 2, 3)

		out := TakeUntil[int](signal)(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		rec.AwaitCompleted(time.Second)

		assertlite.IsNil(t, rec.Values())
		assertlite.True(t, rec.Completed())
	})

	t.Run("completion happens at most once", func(t *testing.T) {
		signal := obs.AsyncIntervalOf[any](10 * time.Millisecond)
		src := obs.AsyncIntervalOf(10*time.Millisecond, 1, 2)

		out := TakeUntil[int](signal)(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		rec.AwaitCompleted(time.Second)

		assertlite.True(t, rec.Completed())
	})

	t.Run("states is local per subscription", func(t *testing.T) {
		signal := obs.AsyncIntervalOf[any](30 * time.Millisecond)
		src := obs.AsyncIntervalOf(10*time.Millisecond, 1, 2, 3)

		out := TakeUntil[int](signal)(src)

		rec1, o1 := observers.NewRecorder[int](t)
		rec2, o2 := observers.NewRecorder[int](t)

		_, _ = out.Sub((*obs.Observer[int])(o1))
		_, _ = out.Sub((*obs.Observer[int])(o2))

		rec1.AwaitCompleted(time.Second)
		rec2.AwaitCompleted(time.Second)

		assertlite.True(t, rec1.Completed())
		assertlite.True(t, rec2.Completed())
	})
}

// TestTakeWhile_Sync
//
// Should assert:
//  1. Values are forwarded while the predicate returns true
//  2. The first value failing the predicate is NOT emitted
//  3. Predicate failure completes the stream immediately
//  4. Source is unsubscribed when the predicate fails
//  5. Source completion before predicate failure is forwarded
//  6. Source Error propagates immediately
//  7. Completion happens at most once
//  8. Predicate is evaluated synchronously and in order
//  9. State is local per subscription
func TestTakeWhile_Sync(t *testing.T) {
	t.Run("values are forwarded while predicate is true", func(t *testing.T) {
		src := obs.Of(1, 2, 3)

		out := TakeWhile(func(v int) bool {
			return v < 4
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		rec.AwaitCompleted(time.Second)

		assertlite.EqualSlices(t, []int{1, 2, 3}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("failing value is not emitted and completes immediately", func(t *testing.T) {
		src := obs.Of(1, 2, 3, 4, 5)

		out := TakeWhile(func(v int) bool {
			return v < 4
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		rec.AwaitCompleted(time.Second)

		assertlite.EqualSlices(t, []int{1, 2, 3}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("source completion before predicate failure is forwarded", func(t *testing.T) {
		src := obs.Of(1, 2)

		out := TakeWhile(func(v int) bool {
			return v < 10
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		rec.AwaitCompleted(time.Second)

		assertlite.EqualSlices(t, []int{1, 2}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("source error propagates immediately", func(t *testing.T) {
		err := errors.New("fail")

		src := obs.Error[int](err)

		out := TakeWhile(func(v int) bool {
			return true
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		rec.AwaitTermination(time.Second)

		assertlite.False(t, rec.Completed())
		assertlite.IsErrorOf(t, err, rec.Err())
	})

	t.Run("completion happens at most once", func(t *testing.T) {
		src := obs.Of(1, 2, 3)

		out := TakeWhile(func(v int) bool {
			return v < 2
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		rec.AwaitCompleted(time.Second)

		assertlite.EqualSlices(t, []int{1}, rec.Values())
		assertlite.True(t, rec.Completed())
	})

	t.Run("states is local per subscription", func(t *testing.T) {
		src := obs.Of(1, 2, 3)

		out := TakeWhile(func(v int) bool {
			return v < 3
		})(src)

		rec1, o1 := observers.NewRecorder[int](t)
		rec2, o2 := observers.NewRecorder[int](t)

		_, _ = out.Sub((*obs.Observer[int])(o1))
		_, _ = out.Sub((*obs.Observer[int])(o2))

		rec1.AwaitCompleted(time.Second)
		rec2.AwaitCompleted(time.Second)

		assertlite.EqualSlices(t, []int{1, 2}, rec1.Values())
		assertlite.EqualSlices(t, []int{1, 2}, rec2.Values())
		assertlite.True(t, rec1.Completed())
		assertlite.True(t, rec2.Completed())
	})
}

// TestTakeWhile_Async
//
// Should assert:
//  1. Values are forwarded while the predicate returns true
//  2. The first value failing the predicate is NOT emitted
//  3. Predicate failure completes the stream immediately
//  4. Source emissions after predicate failure are ignored
//  5. Source Error propagates immediately
//  6. Source completion before predicate failure is forwarded
//  7. Completion happens at most once
//  8. Predicate evaluation remains synchronous per value
//  9. State is local per subscription
func TestTakeWhile_Async(t *testing.T) {
	t.Run("values are forwarded while predicate is true", func(t *testing.T) {
		src := obs.AsyncIntervalOf(10*time.Millisecond, 1, 2, 3)

		out := TakeWhile(func(v int) bool {
			return v < 4
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		rec.AwaitCompleted(time.Second)

		assertlite.EqualSlices(t, []int{1, 2, 3}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("failing value is not emitted and completes immediately", func(t *testing.T) {
		src := obs.AsyncIntervalOf(10*time.Millisecond, 1, 2, 3, 4, 5)

		out := TakeWhile(func(v int) bool {
			return v < 4
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		rec.AwaitCompleted(time.Second)

		assertlite.EqualSlices(t, []int{1, 2, 3}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("source completion before predicate failure is forwarded", func(t *testing.T) {
		src := obs.AsyncIntervalOf(10*time.Millisecond, 1, 2)

		out := TakeWhile(func(v int) bool {
			return v < 10
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		rec.AwaitCompleted(time.Second)

		assertlite.EqualSlices(t, []int{1, 2}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("source error propagates immediately", func(t *testing.T) {
		err := errors.New("fail")

		src := obs.AsyncError[int](err)

		out := TakeWhile(func(v int) bool {
			return true
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		rec.AwaitTermination(time.Second)

		assertlite.False(t, rec.Completed())
		assertlite.IsErrorOf(t, err, rec.Err())
	})

	t.Run("completion happens at most once", func(t *testing.T) {
		src := obs.AsyncIntervalOf(10*time.Millisecond, 1, 2, 3)

		out := TakeWhile(func(v int) bool {
			return v < 2
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		rec.AwaitCompleted(time.Second)

		assertlite.EqualSlices(t, []int{1}, rec.Values())
		assertlite.True(t, rec.Completed())
	})

	t.Run("states is local per subscription", func(t *testing.T) {
		src := obs.AsyncIntervalOf(10*time.Millisecond, 1, 2, 3)

		out := TakeWhile(func(v int) bool {
			return v < 3
		})(src)

		rec1, o1 := observers.NewRecorder[int](t)
		rec2, o2 := observers.NewRecorder[int](t)

		_, _ = out.Sub((*obs.Observer[int])(o1))
		_, _ = out.Sub((*obs.Observer[int])(o2))

		rec1.AwaitCompleted(time.Second)
		rec2.AwaitCompleted(time.Second)

		assertlite.EqualSlices(t, []int{1, 2}, rec1.Values())
		assertlite.EqualSlices(t, []int{1, 2}, rec2.Values())
		assertlite.True(t, rec1.Completed())
		assertlite.True(t, rec2.Completed())
	})
}
