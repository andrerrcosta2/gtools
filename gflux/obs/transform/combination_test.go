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

// TestConcat_Sync_Sync tests the Concat pipe using
// synchronous sources and synchronous pipe observables
//
// For purely synchronous sources, we should assert:
//  1. Concatenates multiple sources in order
//  2. Works with a single source
//  3. Skip empty sources correctly
//  4. Handles all sources empty
//  5. Stops on error and does not subscribe to next
//  6. Complete synchronously
//  7. Does not emit after completion
func TestConcat_Sync_Sync(t *testing.T) {
	t.Run("concatenates multiple sync sources in order", func(t *testing.T) {
		src1 := obs.Of(1, 2)
		src2 := obs.Of(3, 4)
		src3 := obs.Of(5)

		concat := Concat(src1, src2, src3)

		rec, o := observers.NewRecorder[int](t)
		_, err := concat.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		assertlite.EqualSlices(t, []int{1, 2, 3, 4, 5}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("works with a single sync source", func(t *testing.T) {
		src := obs.Of(1, 2, 3)

		concat := Concat(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = concat.Sub((*obs.Observer[int])(o))

		assertlite.EqualSlices(t, []int{1, 2, 3}, rec.Values())
		assertlite.True(t, rec.Completed())
	})

	t.Run("skips empty sync sources", func(t *testing.T) {
		src1 := obs.Of[int]()
		src2 := obs.Of(1, 2)
		src3 := obs.Of[int]()
		src4 := obs.Of(3)

		concat := Concat(src1, src2, src3, src4)

		rec, o := observers.NewRecorder[int](t)
		_, _ = concat.Sub((*obs.Observer[int])(o))

		assertlite.EqualSlices(t, []int{1, 2, 3}, rec.Values())
		assertlite.True(t, rec.Completed())
	})

	t.Run("completes immediately if all sources are empty", func(t *testing.T) {
		src1 := obs.Of[int]()
		src2 := obs.Of[int]()

		concat := Concat(src1, src2)

		rec, o := observers.NewRecorder[int](t)
		_, _ = concat.Sub((*obs.Observer[int])(o))

		assertlite.Equals(t, 0, len(rec.Values()))
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("stops on error and does not subscribe to next source", func(t *testing.T) {
		failure := errors.New("failure")

		src1 := obs.Of(1)
		src2 := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			o.Error(failure)
			return subscr.Noop{}, nil
		})
		src3 := obs.Of(999) // must never be reached

		concat := Concat(src1, src2, src3)

		rec, o := observers.NewRecorder[int](t)
		_, _ = concat.Sub((*obs.Observer[int])(o))

		assertlite.EqualSlices(t, []int{1}, rec.Values())
		assertlite.IsErrorOf(t, rec.Err(), failure)
		assertlite.False(t, rec.Completed())
	})

	t.Run("completes synchronously", func(t *testing.T) {
		src1 := obs.Of(1)
		src2 := obs.Of(2)

		concat := Concat(src1, src2)

		done := false
		o := &obs.Observer[int]{
			Next: func(int) {},
			Complete: func() {
				done = true
			},
		}

		_, _ = concat.Sub(o)

		assertlite.True(t, done)
	})
}

// TestConcat_Sync_Async tests the Concat pipe using
// synchronous sources and asynchronous pipe observables
//
// We must assert:
//  1. Sync source fully emits before async source starts emitting
//  2. Values are not interleaved
//  3. Async source still emits correctly
//  4. Completion happens only after async completes
//  5. Error in async stops concat
//  6. Sync completion triggers async subscription (not eagerly)
//
// No timing, only ordering and completeness.
func TestConcat_Sync_Async(t *testing.T) {
	t.Run("concatenates sync source followed by async source", func(t *testing.T) {
		src1 := obs.Of(1, 2)
		src2 := obs.AsyncOf(3, 4, 5)

		concat := Concat(src1, src2)

		rec, o := observers.NewRecorder[int](t)
		_, err := concat.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		rec.AwaitTermination(200 * time.Millisecond)

		assertlite.EqualSlices(t, []int{1, 2, 3, 4, 5}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("does not subscribe to async source before sync completes", func(t *testing.T) {
		var asyncSubscribed bool

		src1 := obs.Of(1, 2)

		src2 := obs.Async[int](func(o *obs.Observer[int]) {
			asyncSubscribed = true
			o.Next(3)
			o.Complete()
		})

		concat := Concat(src1, src2)

		rec, o := observers.NewRecorder[int](t)
		_, _ = concat.Sub((*obs.Observer[int])(o))

		rec.AwaitTermination(200 * time.Millisecond)

		assertlite.True(t, asyncSubscribed)
		assertlite.EqualSlices(t, []int{1, 2, 3}, rec.Values())
	})

	t.Run("propagates async error and stops concatenation", func(t *testing.T) {
		failure := errors.New("failure")

		src1 := obs.Of(1)

		src2 := obs.Async[int](func(o *obs.Observer[int]) {
			o.Error(failure)
		})

		src3 := obs.Of(999) // must never be reached

		concat := Concat(src1, src2, src3)

		rec, o := observers.NewRecorder[int](t)
		_, _ = concat.Sub((*obs.Observer[int])(o))

		rec.AwaitTermination(200 * time.Millisecond)

		assertlite.EqualSlices(t, []int{1}, rec.Values())
		assertlite.IsErrorOf(t, rec.Err(), failure)
		assertlite.False(t, rec.Completed())
	})

	t.Run("sync source completes synchronously before async starts", func(t *testing.T) {
		var seen []int

		src1 := obs.Of(1, 2)

		src2 := obs.Async[int](func(o *obs.Observer[int]) {
			o.Next(3)
			o.Complete()
		})

		concat := Concat(src1, src2)

		o := &obs.Observer[int]{
			Next: func(v int) {
				seen = append(seen, v)
			},
		}

		_, _ = concat.Sub(o)

		// At this point, sync values must already be present
		assertlite.True(t, len(seen) >= 2)
		assertlite.Equals(t, []int{1, 2}, seen[:2])
	})
}

// TestConcat_Async_Sync tests the Concat pipe using
// asynchronous sources followed by synchronous sources
//
// We must assert:
//  1. Async source fully emits before sync source is subscribed
//  2. Values are emitted in strict source order (no interleaving)
//  3. Sync source emits only after async completion
//  4. Completion happens only after both async and sync sources complete
//  5. Error in async source stops concat and skips sync sources
//  6. Sync source subscription is gated by async completion
//
// No timing assumptions, only sequencing, ordering, and termination semantics.
func TestConcat_Async_Sync(t *testing.T) {
	t.Run("concatenates async source followed by sync source", func(t *testing.T) {
		src1 := obs.AsyncOf(1, 2, 3)
		src2 := obs.Of(4, 5)

		concat := Concat(src1, src2)

		rec, o := observers.NewRecorder[int](t)
		_, err := concat.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		rec.AwaitTermination(200 * time.Millisecond)

		assertlite.EqualSlices(t, []int{1, 2, 3, 4, 5}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("does not subscribe to sync source before async completes", func(t *testing.T) {
		var syncSubscribed bool

		src1 := obs.Async[int](func(o *obs.Observer[int]) {
			o.Next(1)
			time.Sleep(20 * time.Millisecond)
			o.Complete()
		})

		src2 := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			syncSubscribed = true
			o.Next(2)
			o.Complete()
			return subscr.Noop{}, nil
		})

		concat := Concat(src1, src2)

		rec, o := observers.NewRecorder[int](t)
		_, _ = concat.Sub((*obs.Observer[int])(o))

		// give async source time to run but not complete yet
		time.Sleep(5 * time.Millisecond)

		assertlite.False(t, syncSubscribed)

		rec.AwaitTermination(200 * time.Millisecond)

		assertlite.True(t, syncSubscribed)
		assertlite.EqualSlices(t, []int{1, 2}, rec.Values())
	})

	t.Run("propagates async error and skips sync source", func(t *testing.T) {
		failure := errors.New("failure")

		src1 := obs.Async[int](func(o *obs.Observer[int]) {
			o.Next(1)
			o.Error(failure)
		})

		src2 := obs.Of(999) // must not be subscribed

		concat := Concat(src1, src2)

		rec, o := observers.NewRecorder[int](t)
		_, _ = concat.Sub((*obs.Observer[int])(o))

		rec.AwaitTermination(200 * time.Millisecond)

		assertlite.EqualSlices(t, []int{1}, rec.Values())
		assertlite.IsErrorOf(t, rec.Err(), failure)
		assertlite.False(t, rec.Completed())
	})

	t.Run("completes only after async then sync", func(t *testing.T) {
		src1 := obs.AsyncOf(1)
		src2 := obs.Of(2)

		concat := Concat(src1, src2)

		rec, o := observers.NewRecorder[int](t)
		_, _ = concat.Sub((*obs.Observer[int])(o))

		rec.AwaitTermination(200 * time.Millisecond)

		assertlite.True(t, rec.Completed())
		assertlite.EqualSlices(t, []int{1, 2}, rec.Values())
	})
}

// TestConcat_Async_Async tests the Concat pipe using
// asynchronouns sources and pipe observables
//
// With two (or more) async sources, we must assert:
//  1. second async source is not subscribed until the first completes
//  2. emissions are not interleaved
//  3. values preserve source order
//  4. completion happens only after the last async source completes
//  5. error in any async source short-circuits the chain
//  6. later async sources are never subscribed after an error
//
// No reliance on scheduling quirks — only observable semantics.
func TestConcat_Async_Async(t *testing.T) {
	t.Run("concatenates async sources sequentially without interleaving", func(t *testing.T) {
		src1 := obs.AsyncOf(1, 2)
		src2 := obs.AsyncOf(3, 4)

		concat := Concat(src1, src2)

		rec, o := observers.NewRecorder[int](t)
		_, err := concat.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		rec.AwaitTermination(200 * time.Millisecond)

		assertlite.EqualSlices(t, []int{1, 2, 3, 4}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("does not subscribe to second async source before first completes", func(t *testing.T) {
		var secondSubscribed bool

		src1 := obs.Async[int](func(o *obs.Observer[int]) {
			o.Next(1)
			time.Sleep(20 * time.Millisecond)
			o.Complete()
		})

		src2 := obs.Async[int](func(o *obs.Observer[int]) {
			secondSubscribed = true
			o.Next(2)
			o.Complete()
		})

		concat := Concat(src1, src2)

		rec, o := observers.NewRecorder[int](t)
		_, _ = concat.Sub((*obs.Observer[int])(o))

		time.Sleep(5 * time.Millisecond)
		assertlite.False(t, secondSubscribed)

		rec.AwaitTermination(200 * time.Millisecond)

		assertlite.True(t, secondSubscribed)
		assertlite.EqualSlices(t, []int{1, 2}, rec.Values())
	})

	t.Run("propagates error from first async source and skips second", func(t *testing.T) {
		failure := errors.New("failure")

		src1 := obs.Async[int](func(o *obs.Observer[int]) {
			o.Next(1)
			o.Error(failure)
		})

		src2 := obs.AsyncOf(999) // must never be subscribed

		concat := Concat(src1, src2)

		rec, o := observers.NewRecorder[int](t)
		_, _ = concat.Sub((*obs.Observer[int])(o))

		rec.AwaitTermination(200 * time.Millisecond)

		assertlite.EqualSlices(t, []int{1}, rec.Values())
		assertlite.IsErrorOf(t, rec.Err(), failure)
		assertlite.False(t, rec.Completed())
	})

	t.Run("error in later async source stops concat", func(t *testing.T) {
		failure := errors.New("failure")

		src1 := obs.AsyncOf(1, 2)

		src2 := obs.Async[int](func(o *obs.Observer[int]) {
			o.Error(failure)
		})

		concat := Concat(src1, src2)

		rec, o := observers.NewRecorder[int](t)
		_, _ = concat.Sub((*obs.Observer[int])(o))

		rec.AwaitTermination(200 * time.Millisecond)

		assertlite.EqualSlices(t, []int{1, 2}, rec.Values())
		assertlite.IsErrorOf(t, rec.Err(), failure)
		assertlite.False(t, rec.Completed())
	})
}

// TestMerge_Sync_Sync tests the Merge pipe using synchronous sources
// followed by synchronous observables
//
// With only synchronous sources:
//
//  1. all sources are subscribed eagerly
//  2. values are emitted during subscription
//  3. ordering is deterministic but source-order–dependent
//  4. Completion happens only after all sources complete
//  5. Error in any source short-circuits immediately
//  6. Later sources may never be subscribed if error happens synchronously
//
// Unlike Concat, interleaving is impossible here, everything happens inline.
func TestMerge_Sync_Sync(t *testing.T) {
	t.Run("merges multiple sync sources eagerly", func(t *testing.T) {
		src1 := obs.Of(1, 2)
		src2 := obs.Of(3, 4)

		merged := Merge(src1, src2)

		rec, o := observers.NewRecorder[int](t)
		_, err := merged.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		assertlite.EqualSlices(t, []int{1, 2, 3, 4}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("subscribes to all sync sources eagerly", func(t *testing.T) {
		var subscribed []int

		src := func(id int, values ...int) obs.Observable[int] {
			return obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
				subscribed = append(subscribed, id)
				for _, v := range values {
					o.Next(v)
				}
				o.Complete()
				return subscr.Noop{}, nil
			})
		}

		merged := Merge(
			src(1, 1),
			src(2, 2),
			src(3, 3),
		)

		rec, o := observers.NewRecorder[int](t)
		_, _ = merged.Sub((*obs.Observer[int])(o))

		assertlite.EqualSlices(t, []int{1, 2, 3}, subscribed)
		assertlite.EqualSlices(t, []int{1, 2, 3}, rec.Values())
		assertlite.True(t, rec.Completed())
	})

	t.Run("completes only after all sync sources complete", func(t *testing.T) {
		var completed []int

		src := func(id int) obs.Observable[int] {
			return obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
				o.Next(id)
				o.Complete()
				completed = append(completed, id)
				return subscr.Noop{}, nil
			})
		}

		merged := Merge(src(1), src(2), src(3))

		rec, o := observers.NewRecorder[int](t)
		_, _ = merged.Sub((*obs.Observer[int])(o))

		assertlite.EqualSlices(t, []int{1, 2, 3}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.EqualSlices(t, []int{1, 2, 3}, completed)
	})

	t.Run("fails fast on synchronous error", func(t *testing.T) {
		failure := errors.New("failure")
		var subscribed []int

		src := func(id int, err error) obs.Observable[int] {
			return obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
				subscribed = append(subscribed, id)
				if err != nil {
					o.Error(err)
					return subscr.Noop{}, nil
				}
				o.Next(id)
				o.Complete()
				return subscr.Noop{}, nil
			})
		}

		merged := Merge(
			src(1, nil),
			src(2, failure),
			src(3, nil), // must never be subscribed
		)

		rec, o := observers.NewRecorder[int](t)
		_, _ = merged.Sub((*obs.Observer[int])(o))

		assertlite.EqualSlices(t, []int{1}, rec.Values())
		assertlite.IsErrorOf(t, rec.Err(), failure)
		assertlite.False(t, rec.Completed())
		assertlite.EqualSlices(t, []int{1, 2}, subscribed)
	})

	t.Run("merge with no sources completes immediately", func(t *testing.T) {
		merged := Merge[int]()

		rec, o := observers.NewRecorder[int](t)
		_, err := merged.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
		assertlite.IsNil(t, rec.Values())
	})
}

// TestMerge_Sync_Async tests the Merge pipe where sources
// are subscribed synchronously, but at least one source emits asynchronously.
//
// We must assert:
//  1. All sources are subscribed eagerly
//  2. Sync sources emit immediately
//  3. Async sources emit later
//  4. Values may interleave (but no loss / duplication)
//     5.Completion happens only after async completes
//  6. Error from async fails fast and unsubscribes sync
//
// We do not assert exact ordering across sources, only relative correctness.
func TestMerge_Sync_Async(t *testing.T) {
	t.Run("sync emits immediately, async emits later", func(t *testing.T) {
		asyncCh := make(chan int)

		syncSrc := obs.Of(1, 2)
		asyncSrc := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			go func() {
				for v := range asyncCh {
					o.Next(v)
				}
				o.Complete()
			}()
			return subscr.Noop{}, nil
		})

		merged := Merge(syncSrc, asyncSrc)

		rec, o := observers.NewRecorder[int](t)
		_, err := merged.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		// Sync values must already be there
		assertlite.EqualSlices(t, []int{1, 2}, rec.Values())
		assertlite.False(t, rec.Completed())

		asyncCh <- 3
		asyncCh <- 4
		close(asyncCh)

		rec.AwaitCompleted(100 * time.Millisecond)

		assertlite.EqualSlices(t, []int{1, 2, 3, 4}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("async error fails merge and stops sync", func(t *testing.T) {
		failure := errors.New("failure")
		asyncErr := make(chan struct{})

		syncSrc := obs.Of(1, 2, 3)

		asyncSrc := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			go func() {
				<-asyncErr
				o.Error(failure)
			}()
			return subscr.Noop{}, nil
		})

		merged := Merge(syncSrc, asyncSrc)

		rec, o := observers.NewRecorder[int](t)
		_, _ = merged.Sub((*obs.Observer[int])(o))

		// sync emits immediately
		assertlite.EqualSlices(t, []int{1, 2, 3}, rec.Values())

		close(asyncErr)
		_ = rec.AwaitError(100 * time.Millisecond)

		assertlite.IsErrorOf(t, rec.Err(), failure)
		assertlite.False(t, rec.Completed())
	})

	t.Run("completion waits for async source", func(t *testing.T) {
		done := make(chan struct{})

		syncSrc := obs.Of(1)

		asyncSrc := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			go func() {
				<-done
				o.Next(2)
				o.Complete()
			}()
			return subscr.Noop{}, nil
		})

		merged := Merge(syncSrc, asyncSrc)

		rec, o := observers.NewRecorder[int](t)
		_, _ = merged.Sub((*obs.Observer[int])(o))

		assertlite.EqualSlices(t, []int{1}, rec.Values())
		assertlite.False(t, rec.Completed())

		close(done)
		rec.AwaitCompleted(100 * time.Millisecond)

		assertlite.EqualSlices(t, []int{1, 2}, rec.Values())
		assertlite.True(t, rec.Completed())
	})
}

// TestMerge_Async_Sync tests Merge with asynchronous sources
// followed by synchronous sources.
//
// We assert:
//  1. Both sources are subscribed eagerly
//  2. Sync source emits immediately even if async is first
//  3. Async values arrive later
//  4. Completion waits for async
//  5. Error in sync fails async immediately
//  6. Error in async fails sync even if sync already emitted
//
// Ordering is NOT guaranteed — only presence and termination rules.
func TestMerge_Async_Sync(t *testing.T) {
	t.Run("sync emits immediately even when async is first", func(t *testing.T) {
		ch := make(chan int)

		asyncSrc := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			go func() {
				for v := range ch {
					o.Next(v)
				}
				o.Complete()
			}()
			return subscr.Noop{}, nil
		})

		syncSrc := obs.Of(2, 3)

		merged := Merge(asyncSrc, syncSrc)

		rec, o := observers.NewRecorder[int](t)
		_, err := merged.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		// Sync emits immediately
		assertlite.EqualSlices(t, []int{2, 3}, rec.Values())
		assertlite.False(t, rec.Completed())

		ch <- 1
		close(ch)

		rec.AwaitCompleted(100 * time.Millisecond)

		assertlite.ElementsMatch(t, []int{1, 2, 3}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("completion waits for async even if sync completes first", func(t *testing.T) {
		done := make(chan struct{})

		asyncSrc := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			go func() {
				<-done
				o.Next(1)
				o.Complete()
			}()
			return subscr.Noop{}, nil
		})

		syncSrc := obs.Of(2)

		merged := Merge(asyncSrc, syncSrc)

		rec, o := observers.NewRecorder[int](t)
		_, _ = merged.Sub((*obs.Observer[int])(o))

		assertlite.EqualSlices(t, []int{2}, rec.Values())
		assertlite.False(t, rec.Completed())

		close(done)
		rec.AwaitCompleted(100 * time.Millisecond)

		assertlite.ElementsMatch(t, []int{1, 2}, rec.Values())
		assertlite.True(t, rec.Completed())
	})

	t.Run("sync error fails async immediately", func(t *testing.T) {
		failure := errors.New("failure")

		asyncStarted := make(chan struct{})
		asyncStopped := make(chan struct{})

		asyncSrc := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			close(asyncStarted)
			go func() {
				defer close(asyncStopped)
				select {}
			}()
			return subscr.Func(func() {
				close(asyncStopped)
			}), nil
		})

		syncSrc := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			o.Error(failure)
			return subscr.Noop{}, nil
		})

		merged := Merge(asyncSrc, syncSrc)

		rec, o := observers.NewRecorder[int](t)
		_, _ = merged.Sub((*obs.Observer[int])(o))

		<-asyncStarted
		<-asyncStopped

		assertlite.IsErrorOf(t, rec.Err(), failure)
		assertlite.False(t, rec.Completed())
	})

	t.Run("async error fails merge even if sync already emitted", func(t *testing.T) {
		failure := errors.New("failure")
		trigger := make(chan struct{})

		asyncSrc := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			go func() {
				<-trigger
				o.Error(failure)
			}()
			return subscr.Noop{}, nil
		})

		syncSrc := obs.Of(1, 2)

		merged := Merge(asyncSrc, syncSrc)

		rec, o := observers.NewRecorder[int](t)
		_, _ = merged.Sub((*obs.Observer[int])(o))

		assertlite.EqualSlices(t, []int{1, 2}, rec.Values())

		close(trigger)
		_ = rec.AwaitError(100 * time.Millisecond)

		assertlite.IsErrorOf(t, rec.Err(), failure)
		assertlite.False(t, rec.Completed())
	})
}

// TestMerge_Async_Async tests the Merge operator using
// multiple asynchronous sources.
//
// Since both sources may emit concurrently, ordering is NOT guaranteed.
// We must assert:
//
//  1. All values from all async sources are eventually emitted
//  2. Values may interleave arbitrarily
//  3. Completion happens only after all async sources complete
//  4. Error from any async source fails the merge immediately
//  5. Other async sources are unsubscribed on error (fail-fast)
//
// No timing assumptions — only set equality and termination semantics.
func TestMerge_Async_Async(t *testing.T) {
	t.Run("merges async sources and emits all values", func(t *testing.T) {
		src1 := obs.AsyncOf(1, 2, 3)
		src2 := obs.AsyncOf(4, 5, 6)

		merged := Merge(src1, src2)

		rec, o := observers.NewRecorder[int](t)
		_, _ = merged.Sub((*obs.Observer[int])(o))

		rec.AwaitTermination(200 * time.Millisecond)

		assertlite.ElementsMatch(t,
			[]int{1, 2, 3, 4, 5, 6},
			rec.Values(),
		)
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("completion waits for all async sources", func(t *testing.T) {
		src1 := obs.Async[int](func(o *obs.Observer[int]) {
			time.Sleep(50 * time.Millisecond)
			o.Next(1)
			o.Complete()
		})

		src2 := obs.Async[int](func(o *obs.Observer[int]) {
			time.Sleep(100 * time.Millisecond)
			o.Next(2)
			o.Complete()
		})

		merged := Merge(src1, src2)

		rec, o := observers.NewRecorder[int](t)
		_, _ = merged.Sub((*obs.Observer[int])(o))

		rec.AwaitTermination(300 * time.Millisecond)

		assertlite.ElementsMatch(t, []int{1, 2}, rec.Values())
		assertlite.True(t, rec.Completed())
	})

	t.Run("async error fails merge immediately", func(t *testing.T) {
		failure := errors.New("failure")

		src1 := obs.Async[int](func(o *obs.Observer[int]) {
			o.Next(1)
			o.Error(failure)
		})

		src2 := obs.Async[int](func(o *obs.Observer[int]) {
			time.Sleep(100 * time.Millisecond)
			o.Next(999) // must never be observed
			o.Complete()
		})

		merged := Merge(src1, src2)

		rec, o := observers.NewRecorder[int](t)
		_, _ = merged.Sub((*obs.Observer[int])(o))

		rec.AwaitTermination(200 * time.Millisecond)

		assertlite.ElementsMatch(t, []int{1}, rec.Values())
		assertlite.IsErrorOf(t, rec.Err(), failure)
		assertlite.False(t, rec.Completed())
	})

	t.Run("error from later async source still fails merge", func(t *testing.T) {
		failure := errors.New("failure")

		src1 := obs.AsyncOf(1, 2)

		src2 := obs.Async[int](func(o *obs.Observer[int]) {
			time.Sleep(20 * time.Millisecond)
			o.Error(failure)
		})

		merged := Merge(src1, src2)

		rec, o := observers.NewRecorder[int](t)
		_, _ = merged.Sub((*obs.Observer[int])(o))

		rec.AwaitTermination(200 * time.Millisecond)

		assertlite.ElementsMatch(t, []int{1, 2}, rec.Values())
		assertlite.IsErrorOf(t, rec.Err(), failure)
	})
}
