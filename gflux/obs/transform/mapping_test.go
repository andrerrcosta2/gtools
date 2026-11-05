// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package transform

import (
	"errors"
	"fmt"
	"strconv"
	"sync/atomic"
	"testing"
	"time"

	"github.com/andrerrcosta2/gtools/core/testlite/assertlite"
	"github.com/andrerrcosta2/gtools/gflux/obs"
	"github.com/andrerrcosta2/gtools/gflux/obs/internal/observers"
	"github.com/andrerrcosta2/gtools/gflux/obs/internal/subscr"
)

// TestConcatMap_Sync_Sync
//
// Should assert:
//  1. Values are mapped and emitted in strict order
//  2. Inner observables are subscribed sequentially
//  3. Next inner starts only after previous completes
//  4. Source completion waits for last inner completion
//  5. Source error propagates immediately
//  6. Inner error propagates immediately
//  7. No values are emitted after termination
//  8. Completion happens exactly once
//  9. State is local per subscription
func TestConcatMap_Sync_Sync(t *testing.T) {
	t.Run("dicts and flattens sequentially preserving order", func(t *testing.T) {
		src := obs.Of(1, 2, 3)

		out := ConcatMap(func(v int) obs.Observable[string] {
			return obs.Of(
				fmt.Sprintf("A%d", v),
				fmt.Sprintf("B%d", v),
			)
		})(src)

		rec, o := observers.NewRecorder[string](t)
		_, err := out.Sub((*obs.Observer[string])(o))
		assertlite.NoError(t, err)

		rec.AwaitCompleted(time.Second)

		assertlite.Equals(t, []string{
			"A1", "B1",
			"A2", "B2",
			"A3", "B3",
		}, rec.Values())
	})

	t.Run("completes only after source and last inner complete", func(t *testing.T) {
		calls := atomic.Int32{}

		src := obs.Of(1, 2)

		out := ConcatMap(func(v int) obs.Observable[int] {
			return obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
				o.Next(v)
				o.Complete()
				calls.Add(1)
				return subscr.Noop{}, nil
			})
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		rec.AwaitCompleted(time.Second)

		assertlite.Equals(t, []int{1, 2}, rec.Values())
		assertlite.Equals(t, int32(2), calls.Load())
	})

	t.Run("inner observables are never concurrent", func(t *testing.T) {
		var active atomic.Int32
		var maximum atomic.Int32

		src := obs.Of(1, 2, 3)

		out := ConcatMap(func(v int) obs.Observable[int] {
			return obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
				n := active.Add(1)
				for {
					old := maximum.Load()
					if n > old && maximum.CompareAndSwap(old, n) {
						break
					}
					if n <= old {
						break
					}
				}

				o.Next(v)
				active.Add(-1)
				o.Complete()
				return subscr.Noop{}, nil
			})
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		rec.AwaitCompleted(time.Second)

		assertlite.Equals(t, []int{1, 2, 3}, rec.Values())
		assertlite.Equals(t, int32(1), maximum.Load())
	})

	t.Run("propagates inner error immediately", func(t *testing.T) {
		src := obs.Of(1, 2, 3)

		expErr := errors.New("fail")

		out := ConcatMap(func(v int) obs.Observable[int] {
			if v == 2 {
				return obs.Error[int](expErr)
			}
			return obs.Of(v)
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		assertlite.IsErrorOf(t, expErr, rec.AwaitError(time.Second))
		assertlite.Equals(t, []int{1}, rec.Values())
	})

	t.Run("propagates source error immediately and cancels inner", func(t *testing.T) {
		innerUnsub := atomic.Bool{}
		srcErr := errors.New("src-fail")

		src := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			o.Next(1)
			o.Error(srcErr)
			return subscr.Noop{}, nil
		})

		out := ConcatMap(func(v int) obs.Observable[int] {
			return obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
				o.Next(v)
				return subscr.Func(func() {
					innerUnsub.Store(true)
				}), nil
			})
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		assertlite.IsErrorOf(t, srcErr, rec.AwaitError(time.Second))
		assertlite.Equals(t, []int{1}, rec.Values())
		assertlite.True(t, innerUnsub.Load())
	})

	t.Run("empty source completes immediately", func(t *testing.T) {
		src := obs.Empty[int]()

		out := ConcatMap(func(v int) obs.Observable[int] {
			t.Error("mapping function should not be called")
			return obs.Empty[int]()
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		rec.AwaitCompleted(time.Second)
		assertlite.IsNil(t, rec.Values())
	})
}

// TestConcatMap_Sync_Async
//
// Source is synchronous.
// Inner observables emit asynchronously.
//
// Should assert:
//  1. Values are emitted in strict order
//  2. Inner observables do not overlap
//  3. Next inner starts only after previous completes
//  4. Source completion waits for last inner completion
//  5. Inner error propagates immediately
//  6. Source error propagates immediately and cancels inner
//  7. No values are emitted after termination
//  8. Completion happens exactly once
func TestConcatMap_Sync_Async(t *testing.T) {
	t.Run("dicts and flattens sequentially with async inners", func(t *testing.T) {
		src := obs.Of(1, 2, 3)

		out := ConcatMap(func(v int) obs.Observable[string] {
			return obs.AsyncIntervalOf[string](10*time.Millisecond, "A"+strconv.Itoa(v), "B"+strconv.Itoa(v))
		})(src)

		rec, o := observers.NewRecorder[string](t)
		_, err := out.Sub((*obs.Observer[string])(o))
		assertlite.NoError(t, err)

		rec.AwaitCompleted(time.Second)

		assertlite.EqualSlices(t, []string{
			"A1", "B1",
			"A2", "B2",
			"A3", "B3",
		}, rec.Values())
	})

	t.Run("inners are never concurrent", func(t *testing.T) {
		var active atomic.Int32
		var maximum atomic.Int32

		src := obs.Of(1, 2, 3)

		out := ConcatMap(func(v int) obs.Observable[int] {
			return obs.Async[int](func(o *obs.Observer[int]) {
				n := active.Add(1)
				for {
					old := maximum.Load()
					if n > old && maximum.CompareAndSwap(old, n) {
						break
					}
					if n <= old {
						break
					}
				}

				time.Sleep(20 * time.Millisecond)
				o.Next(v)
				active.Add(-1)
				o.Complete()
			})
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		rec.AwaitCompleted(time.Second)

		assertlite.EqualSlices(t, []int{1, 2, 3}, rec.Values())
		assertlite.True(t, int32(1) == maximum.Load())
	})

	t.Run("propagates inner error immediately", func(t *testing.T) {
		src := obs.Of(1, 2, 3)

		expErr := errors.New("inner-fail")

		out := ConcatMap(func(v int) obs.Observable[int] {
			return obs.Async[int](func(o *obs.Observer[int]) {
				time.Sleep(10 * time.Millisecond)
				if v == 2 {
					o.Error(expErr)
					return
				}
				o.Next(v)
				o.Complete()
			})
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		assertlite.IsErrorOf(t, expErr, rec.AwaitError(time.Second))
		assertlite.EqualSlices(t, []int{1}, rec.Values())
	})

	t.Run("propagates source error immediately and cancels async inner", func(t *testing.T) {
		innerCancelled := atomic.Bool{}
		srcErr := errors.New("src-fail")
		cancel := make(chan struct{})
		cancelled := make(chan struct{})

		src := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			o.Next(1)
			o.Error(srcErr)
			return subscr.Noop{}, nil
		})

		out := ConcatMap(func(v int) obs.Observable[int] {
			return obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
				go func() {
					select {
					case <-time.After(100 * time.Millisecond):
						o.Next(v)
						o.Complete()
					case <-cancel:
						innerCancelled.Store(true)
					}
				}()
				return subscr.Func(func() {
					close(cancel)
					time.Sleep(10 * time.Millisecond)
					close(cancelled)
				}), nil
			})
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		assertlite.IsErrorOf(t, srcErr, rec.AwaitError(time.Second))
		assertlite.IsNil(t, rec.Values())
		<-cancelled
		assertlite.True(t, innerCancelled.Load())
	})

	t.Run("source completion waits for last async inner", func(t *testing.T) {
		src := obs.Of(1)
		done := atomic.Bool{}

		out := ConcatMap(func(v int) obs.Observable[int] {
			return obs.Async[int](func(o *obs.Observer[int]) {
				time.Sleep(50 * time.Millisecond)
				o.Next(v)
				o.Complete()
				done.Store(true)
			})
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		rec.AwaitCompleted(time.Second)

		assertlite.True(t, done.Load())
		assertlite.Equals(t, []int{1}, rec.Values())
	})
}

// TestConcatMap_Async_Sync
//
// Source emits asynchronously.
// Inner observables emit synchronously.
//
// Should assert:
//  1. Values are emitted in strict order
//  2. Inner observables do not overlap
//  3. Next inner starts only after previous completes
//  4. Source completion waits for last inner completion
//  5. Inner error propagates immediately
//  6. Source error propagates immediately and cancels inner
//  7. No values are emitted after termination
//  8. Completion happens exactly once
func TestConcatMap_Async_Sync(t *testing.T) {
	t.Run("dicts and flattens sequentially with async source", func(t *testing.T) {
		src := obs.Async[int](func(o *obs.Observer[int]) {
			o.Next(1)
			o.Next(2)
			o.Next(3)
			o.Complete()
		})

		out := ConcatMap(func(v int) obs.Observable[string] {
			return obs.Of(
				"A"+strconv.Itoa(v),
				"B"+strconv.Itoa(v),
			)
		})(src)

		rec, o := observers.NewRecorder[string](t)
		_, err := out.Sub((*obs.Observer[string])(o))
		assertlite.NoError(t, err)

		rec.AwaitCompleted(time.Second)

		assertlite.EqualSlices(t, []string{
			"A1", "B1",
			"A2", "B2",
			"A3", "B3",
		}, rec.Values())
	})

	t.Run("inners are never concurrent even with async source", func(t *testing.T) {
		var active atomic.Int32
		var maximum atomic.Int32

		src := obs.Async[int](func(o *obs.Observer[int]) {
			o.Next(1)
			o.Next(2)
			o.Next(3)
			o.Complete()
		})

		out := ConcatMap(func(v int) obs.Observable[int] {
			return obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
				n := active.Add(1)
				for {
					old := maximum.Load()
					if n > old && maximum.CompareAndSwap(old, n) {
						break
					}
					if n <= old {
						break
					}
				}

				// sync emission
				o.Next(v)
				active.Add(-1)
				o.Complete()
				return subscr.Noop{}, nil
			})
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		rec.AwaitCompleted(time.Second)

		assertlite.EqualSlices(t, []int{1, 2, 3}, rec.Values())
		assertlite.True(t, int32(1) == maximum.Load())
	})

	t.Run("propagates inner error immediately", func(t *testing.T) {
		src := obs.Async[int](func(o *obs.Observer[int]) {
			o.Next(1)
			o.Next(2)
			o.Next(3)
			o.Complete()
		})

		expErr := errors.New("inner-fail")

		out := ConcatMap(func(v int) obs.Observable[int] {
			return obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
				if v == 2 {
					o.Error(expErr)
					return subscr.Noop{}, nil
				}
				o.Next(v)
				o.Complete()
				return subscr.Noop{}, nil
			})
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		assertlite.IsErrorOf(t, expErr, rec.AwaitError(time.Second))
		assertlite.EqualSlices(t, []int{1}, rec.Values())
	})

	t.Run("propagates source error immediately", func(t *testing.T) {
		srcErr := errors.New("src-fail")

		src := obs.Async[int](func(o *obs.Observer[int]) {
			o.Next(1)
			o.Error(srcErr)
		})

		out := ConcatMap(func(v int) obs.Observable[int] {
			return obs.Of(v)
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		assertlite.IsErrorOf(t, srcErr, rec.AwaitError(time.Second))
		assertlite.EqualSlices(t, []int{1}, rec.Values())
	})

	t.Run("source completion waits for last sync inner", func(t *testing.T) {
		src := obs.Async[int](func(o *obs.Observer[int]) {
			o.Next(1)
			o.Complete()
		})

		done := atomic.Bool{}

		out := ConcatMap(func(v int) obs.Observable[int] {
			return obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
				o.Next(v)
				o.Complete()
				done.Store(true)
				return subscr.Noop{}, nil
			})
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		rec.AwaitCompleted(time.Second)

		assertlite.True(t, done.Load())
		assertlite.Equals(t, []int{1}, rec.Values())
	})
}

// TestConcatMap_Async_Async
//
// Source emits asynchronously.
// Inner observables emit asynchronously.
//
// Should assert:
//  1. Values are emitted in strict order
//  2. Inner observables do not overlap
//  3. Next inner starts only after previous completes
//  4. Source completion waits for last inner completion
//  5. Inner error propagates immediately
//  6. Source error propagates immediately and cancels inner
//  7. No values are emitted after termination
//  8. Completion happens exactly once
func TestConcatMap_Async_Async(t *testing.T) {
	t.Run("dicts and flattens sequentially with async source and async inners", func(t *testing.T) {
		src := obs.Async[int](func(o *obs.Observer[int]) {
			o.Next(1)
			o.Next(2)
			o.Next(3)
			o.Complete()
		})

		out := ConcatMap(func(v int) obs.Observable[string] {
			return obs.AsyncIntervalOf[string](
				10*time.Millisecond,
				"A"+strconv.Itoa(v),
				"B"+strconv.Itoa(v),
			)
		})(src)

		rec, o := observers.NewRecorder[string](t)
		_, err := out.Sub((*obs.Observer[string])(o))
		assertlite.NoError(t, err)

		rec.AwaitCompleted(time.Second)

		assertlite.EqualSlices(t, []string{
			"A1", "B1",
			"A2", "B2",
			"A3", "B3",
		}, rec.Values())
	})

	t.Run("inners are never concurrent even under full async pressure", func(t *testing.T) {
		var active atomic.Int32
		var maximum atomic.Int32

		src := obs.Async[int](func(o *obs.Observer[int]) {
			o.Next(1)
			o.Next(2)
			o.Next(3)
			o.Complete()
		})

		out := ConcatMap(func(v int) obs.Observable[int] {
			return obs.Async[int](func(o *obs.Observer[int]) {
				n := active.Add(1)
				for {
					old := maximum.Load()
					if n > old && maximum.CompareAndSwap(old, n) {
						break
					}
					if n <= old {
						break
					}
				}

				time.Sleep(20 * time.Millisecond)
				o.Next(v)
				active.Add(-1)
				o.Complete()
			})
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		rec.AwaitCompleted(time.Second)

		assertlite.EqualSlices(t, []int{1, 2, 3}, rec.Values())
		assertlite.True(t, int32(1) == maximum.Load())
	})

	t.Run("propagates inner error immediately", func(t *testing.T) {
		src := obs.Async[int](func(o *obs.Observer[int]) {
			o.Next(1)
			o.Next(2)
			o.Next(3)
			o.Complete()
		})

		expErr := errors.New("inner-fail")

		out := ConcatMap(func(v int) obs.Observable[int] {
			return obs.Async[int](func(o *obs.Observer[int]) {
				time.Sleep(10 * time.Millisecond)
				if v == 2 {
					o.Error(expErr)
					return
				}
				o.Next(v)
				o.Complete()
			})
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		assertlite.IsErrorOf(t, expErr, rec.AwaitError(time.Second))
		assertlite.EqualSlices(t, []int{1}, rec.Values())
	})

	t.Run("propagates source error immediately and cancels async inner", func(t *testing.T) {
		innerCancelled := atomic.Bool{}
		srcErr := errors.New("src-fail")
		cancel := make(chan struct{})
		cancelAck := make(chan struct{})

		src := obs.Async[int](func(o *obs.Observer[int]) {
			o.Next(1)
			o.Error(srcErr)
		})

		out := ConcatMap(func(v int) obs.Observable[int] {
			return obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
				go func() {
					select {
					case <-time.After(500 * time.Millisecond):
						o.Next(v)
						o.Complete()
					case <-cancel:
						innerCancelled.Store(true)
					}
				}()
				return subscr.Func(func() {
					close(cancel)
					time.Sleep(10 * time.Millisecond)
					close(cancelAck)
				}), nil
			})
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		assertlite.IsErrorOf(t, srcErr, rec.AwaitError(time.Second))
		<-cancelAck // we must wait the unsubscription has no relation with error or complete emissions
		assertlite.IsNil(t, rec.Values())
		assertlite.True(t, innerCancelled.Load())
	})

	t.Run("source completion waits for last async inner", func(t *testing.T) {
		src := obs.Async[int](func(o *obs.Observer[int]) {
			o.Next(1)
			o.Complete()
		})

		done := atomic.Bool{}

		out := ConcatMap(func(v int) obs.Observable[int] {
			return obs.Async[int](func(o *obs.Observer[int]) {
				time.Sleep(50 * time.Millisecond)
				o.Next(v)
				o.Complete()
				done.Store(true)
			})
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		rec.AwaitCompleted(time.Second)

		assertlite.True(t, done.Load())
		assertlite.Equals(t, []int{1}, rec.Values())
	})
}

// TestExhaustMap_Sync_Sync
//
// Source is synchronous.
// Inner observables emit synchronously.
//
// Should assert:
//  1. Only one inner is active at a time
//  2. Source values emitted during active inner are ignored
//  3. Values from inner are forwarded synchronously
//  4. Inner completion allows next source value
//  5. Inner error propagates immediately
//  6. Source error propagates immediately
//  7. Completion waits for active inner
//  8. Completion happens exactly once
func TestExhaustMap_Sync_Sync(t *testing.T) {
	t.Run("dicts and flattens synchronously", func(t *testing.T) {
		src := obs.Of(1, 2, 3)

		out := ExhaustMap(func(v int) obs.Observable[string] {
			return obs.Of(
				"A"+strconv.Itoa(v),
				"B"+strconv.Itoa(v),
			)
		})(src)

		rec, o := observers.NewRecorder[string](t)
		_, err := out.Sub((*obs.Observer[string])(o))
		assertlite.NoError(t, err)

		rec.AwaitCompleted(time.Second)

		assertlite.EqualSlices(t, []string{
			"A1", "B1",
			"A2", "B2",
			"A3", "B3",
		}, rec.Values())
	})

	t.Run("schedules sequentially on synchronous flow", func(t *testing.T) {
		src := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			o.Next(1)
			o.Next(2)
			o.Next(3)
			o.Complete()
			return subscr.Noop{}, nil
		})

		out := ExhaustMap(func(v int) obs.Observable[int] {
			return obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
				o.Next(v)
				o.Complete()
				return subscr.Noop{}, nil
			})
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		rec.AwaitCompleted(time.Second)

		assertlite.EqualSlices(t, []int{1, 2, 3}, rec.Values())
	})

	t.Run("allows next source value after inner completes", func(t *testing.T) {
		var innerCount atomic.Int32

		src := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			o.Next(1)
			o.Next(2)
			o.Complete()
			return subscr.Noop{}, nil
		})

		out := ExhaustMap(func(v int) obs.Observable[int] {
			return obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
				innerCount.Add(1)
				o.Next(v)
				o.Complete()
				return subscr.Noop{}, nil
			})
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		rec.AwaitCompleted(time.Second)

		assertlite.EqualSlices(t, []int{1, 2}, rec.Values())
		assertlite.Equals(t, int32(2), innerCount.Load())
	})

	t.Run("propagates inner error immediately", func(t *testing.T) {
		src := obs.Of(1, 2)

		expErr := errors.New("inner-fail")

		out := ExhaustMap(func(v int) obs.Observable[int] {
			if v == 1 {
				return obs.Error[int](expErr)
			}
			return obs.Of(v)
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		assertlite.IsErrorOf(t, expErr, rec.AwaitError(time.Second))
		assertlite.IsNil(t, rec.Values())
	})

	t.Run("propagates source error immediately", func(t *testing.T) {
		srcErr := errors.New("src-fail")

		src := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			o.Next(1)
			o.Error(srcErr)
			return subscr.Noop{}, nil
		})

		out := ExhaustMap(func(v int) obs.Observable[int] {
			return obs.Of(v)
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		assertlite.IsErrorOf(t, srcErr, rec.AwaitError(time.Second))
		assertlite.EqualSlices(t, []int{1}, rec.Values())
	})

	t.Run("completion waits for active inner", func(t *testing.T) {
		var innerCompleted atomic.Bool

		src := obs.Of(1)

		out := ExhaustMap(func(v int) obs.Observable[int] {
			return obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
				o.Next(v)
				innerCompleted.Store(true)
				o.Complete()
				return subscr.Noop{}, nil
			})
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		rec.AwaitCompleted(time.Second)

		assertlite.True(t, innerCompleted.Load())
		assertlite.EqualSlices(t, []int{1}, rec.Values())
	})

	t.Run("completion happens exactly once", func(t *testing.T) {
		var completes atomic.Int32

		src := obs.Of(1, 2, 3)

		out := ExhaustMap(func(v int) obs.Observable[int] {
			return obs.Of(v)
		})(src)

		_, err := out.Sub(obs.NewObserver(
			func(int) {},
			func(error) {},
			func() {
				completes.Add(1)
			},
		))
		assertlite.NoError(t, err)

		assertlite.Equals(t, int32(1), completes.Load())
	})
}

// TestExhaustMap_Sync_Async

// Source emits synchronously
// Inner observables emit asynchronously
//
// we should assert:
//  1. Emit values from the first inner
//  2. Ignore source values emitted while the inner is active
//  3. Accept the next source value after inner completes
//  4. Propagate inner errors immediately
//  5. Propagate source errors immediately
//  6. Wait for active inner before completing
//  7. Complete exactly once
//  8. Never overlap inners
func TestExhaustMap_Sync_Async(t *testing.T) {
	t.Run("emits first inner and ignores subsequent source values until completion", func(t *testing.T) {
		src := obs.Of(1, 2, 3)

		out := ExhaustMap(func(v int) obs.Observable[int] {
			return obs.Async[int](func(o *obs.Observer[int]) {
				time.Sleep(50 * time.Millisecond)
				o.Next(v)
				o.Complete()
			})
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		rec.AwaitCompleted(time.Second)

		// Only the first value should be processed
		assertlite.EqualSlices(t, []int{1}, rec.Values())
	})

	t.Run("accepts next source value after inner completes", func(t *testing.T) {
		src := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			o.Next(1)
			time.Sleep(80 * time.Millisecond) // allow inner to complete
			o.Next(2)
			o.Complete()
			return subscr.Noop{}, nil
		})

		out := ExhaustMap(func(v int) obs.Observable[int] {
			return obs.Async[int](func(o *obs.Observer[int]) {
				time.Sleep(30 * time.Millisecond)
				o.Next(v)
				o.Complete()
			})
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		rec.AwaitCompleted(time.Second)

		assertlite.EqualSlices(t, []int{1, 2}, rec.Values())
	})

	t.Run("inners are never concurrent", func(t *testing.T) {
		var active atomic.Int32
		var max atomic.Int32

		src := obs.Of(1, 2, 3)

		out := ExhaustMap(func(v int) obs.Observable[int] {
			return obs.Async[int](func(o *obs.Observer[int]) {
				n := active.Add(1)
				for {
					old := max.Load()
					if n > old && max.CompareAndSwap(old, n) {
						break
					}
					if n <= old {
						break
					}
				}

				time.Sleep(40 * time.Millisecond)
				o.Next(v)
				active.Add(-1)
				o.Complete()
			})
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		rec.AwaitCompleted(time.Second)

		assertlite.EqualSlices(t, []int{1}, rec.Values())
		assertlite.Equals(t, int32(1), max.Load())
	})

	t.Run("propagates inner error immediately", func(t *testing.T) {
		expErr := errors.New("inner-fail")

		src := obs.Of(1, 2)

		out := ExhaustMap(func(v int) obs.Observable[int] {
			return obs.Async[int](func(o *obs.Observer[int]) {
				time.Sleep(20 * time.Millisecond)
				o.Error(expErr)
			})
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		assertlite.IsErrorOf(t, expErr, rec.AwaitError(time.Second))
		assertlite.IsNil(t, rec.Values())
	})

	t.Run("propagates source error immediately and cancels inner", func(t *testing.T) {
		innerCancelled := atomic.Bool{}
		srcErr := errors.New("src-fail")
		cancel := make(chan struct{})
		cancelAck := make(chan struct{})

		src := obs.Async[int](func(o *obs.Observer[int]) {
			o.Next(1)
			o.Error(srcErr)
		})

		out := ExhaustMap(func(v int) obs.Observable[int] {
			return obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
				go func() {
					select {
					case <-time.After(500 * time.Millisecond): // it won't actually wait that much
						o.Next(v)
						o.Complete()
					case <-cancel:
						innerCancelled.Store(true)
					}
				}()

				return subscr.Func(func() {
					close(cancel)
					time.Sleep(10 * time.Millisecond)
					close(cancelAck)
				}), nil
			})
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		assertlite.IsErrorOf(t, srcErr, rec.AwaitError(time.Second))
		assertlite.IsNil(t, rec.Values())
		<-cancelAck // waiting for the unsub
		assertlite.True(t, innerCancelled.Load())
	})

	t.Run("source completion waits for active inner", func(t *testing.T) {
		done := atomic.Bool{}

		src := obs.Of(1)

		out := ExhaustMap(func(v int) obs.Observable[int] {
			return obs.Async[int](func(o *obs.Observer[int]) {
				time.Sleep(50 * time.Millisecond)
				o.Next(v)
				o.Complete()
				done.Store(true)
			})
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		rec.AwaitCompleted(time.Second)

		assertlite.True(t, done.Load())
		assertlite.EqualSlices(t, []int{1}, rec.Values())
	})
}

// TestExhaustMap_Async_Sync
//
// Source emits asynchronously (single goroutine)
// Inner observables emit synchronously
//
// We should assert:
//  1. Source values are processed sequentially
//  2. Range inner runs to completion before the next source value is handled
//  3. No overlap between inners
//  4. Inner errors propagate immediately
//  5. Source errors propagate immediately
//  6. Completion happens exactly once
func TestExhaustMap_Async_Sync(t *testing.T) {
	t.Run("accepts next source value after inner completes", func(t *testing.T) {
		src := obs.Async[int](func(o *obs.Observer[int]) {
			o.Next(1)
			time.Sleep(50 * time.Millisecond)
			o.Next(2)
			o.Complete()
		})

		out := ExhaustMap(func(v int) obs.Observable[int] {
			return obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
				o.Next(v)
				o.Complete()
				return subscr.Noop{}, nil
			})
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		rec.AwaitCompleted(time.Second)
		assertlite.EqualSlices(t, []int{1, 2}, rec.Values())
	})

	t.Run("propagates inner error immediately", func(t *testing.T) {
		innerErr := errors.New("inner-fail")

		src := obs.Async[int](func(o *obs.Observer[int]) {
			o.Next(1)
			o.Next(2)
		})

		out := ExhaustMap(func(v int) obs.Observable[int] {
			return obs.Error[int](innerErr)
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		assertlite.IsErrorOf(t, innerErr, rec.AwaitError(time.Second))
		assertlite.IsNil(t, rec.Values())
	})

	t.Run("propagates source error immediately", func(t *testing.T) {
		srcErr := errors.New("src-fail")

		src := obs.Async[int](func(o *obs.Observer[int]) {
			o.Next(1)
			o.Error(srcErr)
		})

		out := ExhaustMap(func(v int) obs.Observable[int] {
			return obs.Just(v)
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		assertlite.IsErrorOf(t, srcErr, rec.AwaitError(time.Second))
		assertlite.EqualSlices(t, []int{1}, rec.Values())
	})

	t.Run("waits for active inner before completing", func(t *testing.T) {
		src := obs.Async[int](func(o *obs.Observer[int]) {
			o.Next(1)
			o.Complete()
		})

		out := ExhaustMap(func(v int) obs.Observable[int] {
			return obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
				go func() {
					time.Sleep(50 * time.Millisecond)
					o.Next(v)
					o.Complete()
				}()
				return subscr.Noop{}, nil
			})
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		rec.AwaitCompleted(time.Second)
		assertlite.EqualSlices(t, []int{1}, rec.Values())
	})

}

// TestExhaustMap_Async_Async
//
// Source emits asynchronously
// Inner observables emit asynchronously
//
// We should assert:
//  1. Emit values from the first inner
//  2. Ignore source values emitted while the inner is active
//  3. Accept the next source value after inner completes
//  4. Propagate inner errors immediately
//  5. Propagate source errors immediately
//  6. Wait for active inner before completing
//  7. Complete exactly once
//  8. Never overlap inners
func TestExhaustMap_Async_Async(t *testing.T) {
	t.Run("Async_Async: ignores source values while async inner is active", func(t *testing.T) {
		src := obs.Async[int](func(o *obs.Observer[int]) {
			o.Next(1)
			o.Next(2)
			o.Next(3)
			o.Complete()
		})

		out := ExhaustMap(func(v int) obs.Observable[int] {
			return obs.Async[int](func(o *obs.Observer[int]) {
				time.Sleep(50 * time.Millisecond)
				o.Next(v)
				o.Complete()
			})
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		rec.AwaitCompleted(time.Second)
		assertlite.EqualSlices(t, []int{1}, rec.Values())
	})

	t.Run("Async_Async: source completes but waits for async inner", func(t *testing.T) {
		src := obs.Async[int](func(o *obs.Observer[int]) {
			o.Next(1)
			o.Complete()
		})

		out := ExhaustMap(func(v int) obs.Observable[int] {
			return obs.Async[int](func(o *obs.Observer[int]) {
				time.Sleep(100 * time.Millisecond)
				o.Next(v)
				o.Complete()
			})
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)
		rec.AwaitCompleted(time.Second)

		assertlite.EqualSlices(t, []int{1}, rec.Values())
	})

	t.Run("Async_Async: source error cancels async inner immediately", func(t *testing.T) {
		innerCancelled := atomic.Bool{}
		cancel := make(chan struct{})
		cancelAck := make(chan struct{})
		srcErr := errors.New("src-fail")

		src := obs.Async[int](func(o *obs.Observer[int]) {
			o.Next(1)
			o.Error(srcErr)
		})

		out := ExhaustMap(func(v int) obs.Observable[int] {
			return obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
				go func() {
					select {
					case <-time.After(500 * time.Millisecond):
						o.Next(v)
						o.Complete()
					case <-cancel:
						innerCancelled.Store(true)
					}
				}()

				return subscr.Func(func() {
					close(cancel)
					time.Sleep(10 * time.Millisecond)
					close(cancelAck)
				}), nil
			})
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		assertlite.IsErrorOf(t, srcErr, rec.AwaitError(time.Second))
		<-cancelAck

		assertlite.IsNil(t, rec.Values())
		assertlite.True(t, innerCancelled.Load())
	})

	t.Run("Async_Async: inner error terminates stream immediately", func(t *testing.T) {
		innerErr := errors.New("inner-fail")

		src := obs.Async[int](func(o *obs.Observer[int]) {
			o.Next(1)
			o.Next(2)
			o.Complete()
		})

		out := ExhaustMap(func(v int) obs.Observable[int] {
			return obs.Async[int](func(o *obs.Observer[int]) {
				time.Sleep(50 * time.Millisecond)
				o.Error(innerErr)
			})
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		assertlite.IsErrorOf(t, innerErr, rec.AwaitError(time.Second))
		assertlite.IsNil(t, rec.Values())
	})

	t.Run("Async_Async: unsubscribe cancels async inner", func(t *testing.T) {
		innerCancelled := atomic.Bool{}
		cancel := make(chan struct{})
		cancelAck := make(chan struct{})

		src := obs.Async[int](func(o *obs.Observer[int]) {
			o.Next(1)
		})

		out := ExhaustMap(func(v int) obs.Observable[int] {
			return obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
				go func() {
					select {
					case <-time.After(500 * time.Millisecond):
						o.Next(v)
						o.Complete()
					case <-cancel:
						innerCancelled.Store(true)
					}
				}()

				return subscr.Func(func() {
					close(cancel)
					time.Sleep(10 * time.Millisecond)
					close(cancelAck)
				}), nil
			})
		})(src)

		sub, err := out.Sub(&obs.Observer[int]{})
		assertlite.NoError(t, err)

		time.Sleep(50 * time.Millisecond)
		sub.Unsub()

		<-cancelAck
		assertlite.True(t, innerCancelled.Load())
	})

	t.Run("Async_Async: multiple subscriptions are independent", func(t *testing.T) {
		makeSrc := func() obs.Observable[int] {
			return obs.Async[int](func(o *obs.Observer[int]) {
				o.Next(1)
				o.Next(2)
				o.Complete()
			})
		}

		out := ExhaustMap(func(v int) obs.Observable[int] {
			return obs.Async[int](func(o *obs.Observer[int]) {
				time.Sleep(50 * time.Millisecond)
				o.Next(v)
				o.Complete()
			})
		})

		rec1, o1 := observers.NewRecorder[int](t)
		rec2, o2 := observers.NewRecorder[int](t)

		_, err := out(makeSrc()).Sub((*obs.Observer[int])(o1))
		assertlite.NoError(t, err)

		_, err = out(makeSrc()).Sub((*obs.Observer[int])(o2))
		assertlite.NoError(t, err)

		rec1.AwaitCompleted(time.Second)
		rec2.AwaitCompleted(time.Second)
		assertlite.EqualSlices(t, []int{1}, rec1.Values())
		assertlite.EqualSlices(t, []int{1}, rec2.Values())
	})
}

// TestMap_Sync
//
// Source emits synchronously.
//
// We should assert:
//  1. Range source value is transformed by the mapping function
//  2. Values are emitted in the same order
//  3. Source errors are forwarded immediately
//  4. Source completion is forwarded
//  5. No values are emitted after termination
//  6. Completion happens exactly once
//  7. Mapping function is called exactly once per value
func TestMap_Sync(t *testing.T) {
	t.Run("dicts values synchronously preserving order", func(t *testing.T) {
		src := obs.Of(1, 2, 3)

		out := Map(func(v int) string {
			return strconv.Itoa(v)
		})(src)

		rec, o := observers.NewRecorder[string](t)
		_, err := out.Sub((*obs.Observer[string])(o))
		assertlite.NoError(t, err)

		rec.AwaitCompleted(time.Second)
		assertlite.EqualSlices(t, []string{"1", "2", "3"}, rec.Values())
	})

	t.Run("calls mapping function exactly once per value", func(t *testing.T) {
		var calls atomic.Int32

		src := obs.Of(10, 20, 30)

		out := Map(func(v int) int {
			calls.Add(1)
			return v * 2
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		rec.AwaitCompleted(time.Second)

		assertlite.True(t, int32(3) == calls.Load())
		assertlite.EqualSlices(t, []int{20, 40, 60}, rec.Values())
	})

	t.Run("propagates source error immediately", func(t *testing.T) {
		expErr := errors.New("error")

		src := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			o.Next(1)
			o.Error(expErr)
			o.Next(2) // must not be forwarded
			return subscr.Noop{}, nil
		})

		out := Map(func(v int) int {
			return v + 1
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		assertlite.IsErrorOf(t, expErr, rec.AwaitError(time.Second))
		assertlite.EqualSlices(t, []int{2}, rec.Values())
	})

	t.Run("forwards completion exactly once", func(t *testing.T) {
		var completes atomic.Int32

		src := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			o.Next(1)
			o.Complete()
			o.Complete() // ignored by source contract, but we still assert downstream behavior
			return subscr.Noop{}, nil
		})

		out := Map(func(v int) int {
			return v
		})(src)

		_, err := out.Sub(&obs.Observer[int]{
			Complete: func() {
				completes.Add(1)
			},
		})

		assertlite.NoError(t, err)
		time.Sleep(50 * time.Millisecond)
		assertlite.True(t, int32(1) == completes.Load())
	})

	t.Run("emits no values after completion", func(t *testing.T) {
		src := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			o.Next(1)
			o.Complete()
			o.Next(2) // must be ignored
			return subscr.Noop{}, nil
		})

		out := Map(func(v int) int {
			return v * 10
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		rec.AwaitCompleted(time.Second)

		assertlite.EqualSlices(t, []int{10}, rec.Values())
	})
}

// TestMap_Async
//
// Source emits asynchronously.
//
// We should assert:
//  1. Range value is mapped correctly
//  2. Emission order is preserved
//  3. Source error is propagated immediately
//  4. Source completion is propagated
//  5. No values are emitted after termination
//  6. Completion happens exactly once
//  7. Mapping function is called exactly once per value
func TestMap_Async(t *testing.T) {
	t.Run("dicts async source values preserving order", func(t *testing.T) {
		src := obs.Async[int](func(o *obs.Observer[int]) {
			o.Next(1)
			o.Next(2)
			o.Next(3)
			o.Complete()
		})

		out := Map(func(v int) string {
			return strconv.Itoa(v)
		})(src)

		rec, o := observers.NewRecorder[string](t)
		_, err := out.Sub((*obs.Observer[string])(o))
		assertlite.NoError(t, err)

		rec.AwaitCompleted(time.Second)

		assertlite.EqualSlices(t, []string{"1", "2", "3"}, rec.Values())
	})

	t.Run("calls mapping function exactly once per async value", func(t *testing.T) {
		var calls atomic.Int32

		src := obs.Async[int](func(o *obs.Observer[int]) {
			for i := 0; i < 5; i++ {
				o.Next(i)
			}
			o.Complete()
		})

		out := Map(func(v int) int {
			calls.Add(1)
			return v * 2
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		rec.AwaitCompleted(time.Second)

		assertlite.True(t, int32(5) == calls.Load())
		assertlite.EqualSlices(t, []int{0, 2, 4, 6, 8}, rec.Values())
	})

	t.Run("propagates async source error immediately", func(t *testing.T) {
		expErr := errors.New("async-fail")

		src := obs.Async[int](func(o *obs.Observer[int]) {
			o.Next(1)
			o.Error(expErr)
			o.Next(2) // must be ignored
		})

		out := Map(func(v int) int {
			return v + 1
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		assertlite.IsErrorOf(t, expErr, rec.AwaitError(time.Second))
		assertlite.EqualSlices(t, []int{2}, rec.Values())
	})

	t.Run("forwards async completion exactly once", func(t *testing.T) {
		var completes atomic.Int32

		src := obs.Async[int](func(o *obs.Observer[int]) {
			o.Next(1)
			o.Complete()
			o.Complete() // ignored
		})

		out := Map(func(v int) int {
			return v
		})(src)

		rec, o := observers.NewRecorderFor(t, &observers.Observer[int]{
			Complete: func() {
				completes.Add(1)
			},
		})

		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)
		rec.AwaitCompleted(time.Second)
		assertlite.True(t, int32(1) == completes.Load())
	})

	t.Run("emits no values after async completion", func(t *testing.T) {
		src := obs.Async[int](func(o *obs.Observer[int]) {
			o.Next(1)
			o.Complete()
			o.Next(2) // must be ignored
		})

		out := Map(func(v int) int {
			return v * 10
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		rec.AwaitCompleted(time.Second)

		assertlite.EqualSlices(t, []int{10}, rec.Values())
	})
}

// TestMergeMap_Sync_Sync
//
// we must assert in Sync_Sync
//  1. Projects all source values
//  2. Subscribes to all inner observables
//  3. Forwards all inner values
//  4. Does NOT guarantee ordering
//  5. Completes only after source + all inners complete
//  6. Propagates errors immediately
//  7. Stops further emissions after error
func TestMergeMap_Sync_Sync(t *testing.T) {
	t.Run("merges all inner observables synchronously", func(t *testing.T) {
		src := obs.Of(1, 2, 3)

		out := MergeMap(func(v int) obs.Observable[int] {
			return obs.Of(v*10, v*10+1)
		})(src)

		rec, o := observers.NewRecorder[int](t)

		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		rec.AwaitCompleted(time.Second)

		values := rec.Values()

		// Order is NOT guaranteed by MergeMap
		assertlite.Len(t, values, 6)
		assertlite.ElementsMatch(t,
			[]int{10, 11, 20, 21, 30, 31},
			values,
		)
	})

	t.Run("completes only after all inner observables complete", func(t *testing.T) {
		var completed atomic.Bool

		src := obs.Of(1, 2)

		out := MergeMap(func(v int) obs.Observable[int] {
			return obs.FromFunc[int](func(o *obs.Observer[int]) {
				o.Next(v)
				o.Complete()
			})
		})(src)

		_, err := out.Sub(obs.CompleteObserver[int](func() { completed.Store(true) }))
		assertlite.NoError(t, err)

		assertlite.True(t, completed.Load())
	})

	t.Run("propagates inner error immediately", func(t *testing.T) {
		expErr := errors.New("error")

		src := obs.Of(1, 2, 3)
		pipe := MergeMap(func(v int) obs.Observable[int] {
			if v == 2 {
				return obs.Error[int](expErr)
			}
			return obs.Of(v)
		})
		out := pipe(src)

		rec, o := observers.NewRecorder[int](t)

		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		assertlite.IsErrorOf(t, expErr, rec.AwaitError(time.Second))

		// Must not emit after error
		values := rec.Values()
		assertlite.True(t, len(values) >= 1)
		assertlite.True(t, len(values) <= 2)
	})
}

// TestMergeMap_Sync_Async
//
// Source emits synchronously
// Inner observables emit asynchronously
//
// We should assert:
//  1. Subscribe to all inner observables immediately
//  2. Forward values from all inners as they arrive
//  3. Allow interleaving between inners
//  4. Complete only after source completes AND all inners complete
//  5. Propagate inner errors immediately
//  6. Cancel source and all active inners on error
//  7. Never emit after termination
func TestMergeMap_Sync_Async(t *testing.T) {
	t.Run("merges async inner observables", func(t *testing.T) {
		src := obs.Of(1, 2, 3)

		out := MergeMap(func(v int) obs.Observable[int] {
			return obs.Async(func(o *obs.Observer[int]) {
				time.Sleep(time.Duration(10*v) * time.Millisecond)
				o.Next(v)
				o.Complete()
			})
		})(src)

		rec, o := observers.NewRecorder[int](t)

		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		rec.AwaitCompleted(time.Second)

		values := rec.Values()

		// Order is NOT guaranteed
		assertlite.Len(t, values, 3)
		assertlite.ElementsMatch(t, []int{1, 2, 3}, values)
	})

	t.Run("waits for all active inners before completing", func(t *testing.T) {
		src := obs.Of(1, 2)

		var completed atomic.Bool

		out := MergeMap(func(v int) obs.Observable[int] {
			return obs.Async(func(o *obs.Observer[int]) {
				time.Sleep(50 * time.Millisecond)
				o.Next(v)
				o.Complete()
			})
		})(src)

		rec, o := observers.NewRecorderFor(t, observers.NewComplete[int](func() { completed.Store(true) }))

		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		rec.AwaitNext(time.Second)
		assertlite.False(t, completed.Load(), "must not complete early")

		rec.AwaitNext(time.Second)
		rec.AwaitCompleted(time.Second)
	})

	t.Run("propagates inner error and cancels others", func(t *testing.T) {
		expErr := errors.New("inner error")

		src := obs.Of(1, 2, 3)

		out := MergeMap(func(v int) obs.Observable[int] {
			if v == 2 {
				return obs.Async(func(o *obs.Observer[int]) {
					time.Sleep(20 * time.Millisecond)
					o.Error(expErr)
				})
			}

			return obs.Async(func(o *obs.Observer[int]) {
				time.Sleep(100 * time.Millisecond)
				o.Next(v)
				o.Complete()
			})
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := out.Sub((*obs.Observer[int])(o))

		assertlite.NoError(t, err)
		assertlite.IsErrorOf(t, expErr, rec.AwaitError(time.Second))
		assertlite.IsNil(t, rec.Values())
	})
}

// TestMergeMap_Async_Sync
//
// Source emits asynchronously
// Inner observables emit synchronously
//
// We should assert:
//  1. Subscribe to inners as source emits
//  2. Forward inner values immediately
//  3. Allow interleaving driven by source timing
//  4. Complete only after source completes AND all inners complete
//  5. Propagate inner errors immediately
//  6. Propagate source errors immediately
//  7. Cancel all active subscriptions on error
//  8. Never emit after termination
func TestMergeMap_Async_Sync(t *testing.T) {
	t.Run("merges sync inners from async source", func(t *testing.T) {
		src := obs.Async(func(o *obs.Observer[int]) {
			o.Next(1)
			time.Sleep(10 * time.Millisecond)
			o.Next(2)
			time.Sleep(10 * time.Millisecond)
			o.Next(3)
			o.Complete()
		})

		out := MergeMap(func(v int) obs.Observable[int] {
			// synchronous inner
			return obs.Of(v*10, v*10+1)
		})(src)

		rec, o := observers.NewRecorder[int](t)

		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		rec.AwaitCompleted(time.Second)

		values := rec.Values()
		assertlite.Len(t, values, 6)

		// Order is NOT guaranteed
		assertlite.ElementsMatch(t,
			[]int{10, 11, 20, 21, 30, 31},
			values,
		)
	})

	t.Run("waits for source completion before completing", func(t *testing.T) {
		var completed atomic.Bool

		src := obs.Async(func(o *obs.Observer[int]) {
			o.Next(1)
			time.Sleep(50 * time.Millisecond)
			o.Complete()
		})

		out := MergeMap(func(v int) obs.Observable[int] {
			return obs.Of(v)
		})(src)

		rec, o := observers.NewRecorderFor(t,
			observers.NewComplete[int](func() {
				completed.Store(true)
			}),
		)

		start := time.Now()

		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		rec.AwaitCompleted(time.Second)

		elapsed := time.Since(start)
		assertlite.True(t, elapsed >= 50*time.Millisecond)
		assertlite.True(t, completed.Load())
	})

	t.Run("propagates inner error immediately", func(t *testing.T) {
		expErr := errors.New("inner error")

		src := obs.Async(func(o *obs.Observer[int]) {
			o.Next(1)
			o.Next(2)
			o.Next(3)
			o.Complete()
		})

		out := MergeMap(func(v int) obs.Observable[int] {
			if v == 2 {
				return obs.Error[int](expErr)
			}
			return obs.Of(v)
		})(src)

		rec, o := observers.NewRecorder[int](t)

		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		assertlite.IsErrorOf(t, expErr, rec.AwaitError(time.Second))

		// Must not emit after error
		values := rec.Values()
		assertlite.True(t, len(values) <= 1)
	})

	t.Run("propagates source error and cancels inners", func(t *testing.T) {
		expErr := errors.New("source error")

		src := obs.Async(func(o *obs.Observer[int]) {
			o.Next(1)
			o.Error(expErr)
		})

		out := MergeMap(func(v int) obs.Observable[int] {
			return obs.Of(v, v+1)
		})(src)

		rec, o := observers.NewRecorder[int](t)

		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		assertlite.IsErrorOf(t, expErr, rec.AwaitError(time.Second))
		assertlite.False(t, rec.Completed())
	})
}

// TestMergeMap_Async_Async
//
// Source emits asynchronously
// Inner observables emit asynchronously
//
// We should assert:
//  1. Subscribe to all inners concurrently
//  2. Merge values as they arrive (no ordering guarantees)
//  3. Complete only after source completes AND all inners complete
//  4. Propagate inner errors immediately
//  5. Propagate source errors immediately
//  6. Cancel all active subscriptions on error
//  7. Never emit after termination
func TestMergeMap_Async_Async(t *testing.T) {
	t.Run("merges values from concurrent async inners", func(t *testing.T) {
		src := obs.Async(func(o *obs.Observer[int]) {
			o.Next(1)
			o.Next(2)
			o.Next(3)
			o.Complete()
		})

		out := MergeMap(func(v int) obs.Observable[string] {
			return obs.AsyncDelayedOf(func(_ int, _ string) time.Duration {
				return time.Duration(10*v) * time.Millisecond
			},
				fmt.Sprintf("%d-a", v),
				fmt.Sprintf("%d-b", v),
			)
		})(src)

		rec, o := observers.NewRecorder[string](t)

		_, err := out.Sub((*obs.Observer[string])(o))
		assertlite.NoError(t, err)

		rec.AwaitCompleted(2 * time.Second)

		values := rec.Values()
		assertlite.Len(t, values, 6)

		assertlite.ElementsMatch(t, []string{
			"1-a", "1-b",
			"2-a", "2-b",
			"3-a", "3-b",
		}, values)
	})

	t.Run("waits for async inners after source completes", func(t *testing.T) {
		src := obs.Async(func(o *obs.Observer[int]) {
			o.Next(1)
			o.Next(2)
			o.Complete()
		})

		out := MergeMap(func(v int) obs.Observable[int] {
			return obs.AsyncIntervalOf(50*time.Millisecond, v)
		})(src)

		rec, o := observers.NewRecorder[int](t)

		start := time.Now()

		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		rec.AwaitCompleted(time.Second)
		elapsed := time.Since(start)
		assertlite.True(t, elapsed >= 50*time.Millisecond)
		assertlite.ElementsMatch(t, []int{1, 2}, rec.Values())
	})

	t.Run("propagates inner error and cancels others", func(t *testing.T) {
		expErr := errors.New("inner error")

		src := obs.Async(func(o *obs.Observer[int]) {
			o.Next(1)
			o.Next(2)
			o.Next(3)
		})

		out := MergeMap(func(v int) obs.Observable[int] {
			if v == 2 {
				return obs.Async(func(o *obs.Observer[int]) {
					time.Sleep(20 * time.Millisecond)
					o.Error(expErr)
				})
			}
			return obs.AsyncIntervalOf(100*time.Millisecond, v)
		})(src)

		rec, o := observers.NewRecorder[int](t)

		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		assertlite.IsErrorOf(t, expErr, rec.AwaitError(time.Second))
		// No values should be received by the observer
		// since the delay before the error is too big
		assertlite.IsNil(t, rec.Values())
		// shouldn't complete
		assertlite.False(t, rec.Completed())
	})

	t.Run("propagates source error and cancels inners", func(t *testing.T) {
		expErr := errors.New("source error")

		src := obs.Async(func(o *obs.Observer[int]) {
			o.Next(1)
			o.Next(2)
			o.Error(expErr)
		})

		out := MergeMap(func(v int) obs.Observable[int] {
			return obs.AsyncIntervalOf(100*time.Millisecond, v)
		})(src)

		rec, o := observers.NewRecorder[int](t)

		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		assertlite.IsErrorOf(t, expErr, rec.AwaitError(time.Second))
		assertlite.False(t, rec.Completed())
		// The error shouldn't be piped
		assertlite.IsNil(t, rec.Values())
	})
}

// TestSwitchMap_Sync_Sync
//
// Source emits synchronously
// Inner observables emit synchronously
//
// In this axis there is NO overlap between inners.
// SwitchMap degenerates to a simple map + flatten.
//
// We should assert:
//  1. All source values are projected
//  2. All inner values are forwarded
//  3. No cancellation occurs
//  4. Completion happens after source and last inner complete
//  5. Inner errors propagate immediately
//  6. Source errors propagate immediately
func TestSwitchMap_Sync_Sync(t *testing.T) {
	t.Run("projects and forwards all inner values synchronously", func(t *testing.T) {
		src := obs.Of(1, 2, 3)

		out := SwitchMap(func(v int) obs.Observable[int] {
			return obs.Of(v*10, v*10+1)
		})(src)

		rec, o := observers.NewRecorder[int](t)

		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)
		values := rec.Values()

		assertlite.EqualSlices(t,
			[]int{10, 11, 20, 21, 30, 31},
			values,
		)
	})

	t.Run("completes after source and last inner complete", func(t *testing.T) {
		var completed atomic.Bool

		src := obs.Of(1)

		out := SwitchMap(func(v int) obs.Observable[int] {
			return obs.Of(v)
		})(src)

		_, err := out.Sub(&obs.Observer[int]{
			Complete: func() {
				completed.Store(true)
			},
		})

		assertlite.NoError(t, err)
		assertlite.True(t, completed.Load())
	})

	t.Run("propagates inner error immediately", func(t *testing.T) {
		expErr := errors.New("inner error")

		src := obs.Of(1, 2)

		out := SwitchMap(func(v int) obs.Observable[int] {
			if v == 2 {
				return obs.Error[int](expErr)
			}
			return obs.Of(v)
		})(src)

		rec, o := observers.NewRecorder[int](t)

		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)
		assertlite.IsErrorOf(t, expErr, rec.Err())
	})

	t.Run("propagates source error immediately", func(t *testing.T) {
		expErr := errors.New("source error")

		src := obs.FromFunc[int](func(o *obs.Observer[int]) {
			o.Next(1)
			o.Error(expErr)
		})

		out := SwitchMap(func(v int) obs.Observable[int] {
			return obs.Of(v)
		})(src)

		rec, o := observers.NewRecorder[int](t)

		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)
		assertlite.IsErrorOf(t, expErr, rec.Err())
		assertlite.False(t, rec.Completed())
	})
}

// TestSwitchMap_Sync_Async
//
// Source emits synchronously
// Inner observables emit asynchronously
//
// We should assert:
//  1. Cancel previous inner on new source emission
//  2. Forward values only from the latest inner
//  3. Discard values from cancelled inners
//  4. Complete only after source completes AND latest inner completes
//  5. Propagate inner errors immediately
//  6. Propagate source errors immediately
//  7. Never emit after termination
func TestSwitchMap_Sync_Async(t *testing.T) {

	t.Run("switches to latest async inner", func(t *testing.T) {
		src := obs.Of(1, 2, 3)

		out := SwitchMap(func(v int) obs.Observable[int] {
			return obs.Async(func(o *obs.Observer[int]) {
				time.Sleep(20 * time.Millisecond)
				o.Next(v)
				o.Complete()
			})
		})(src)

		rec, o := observers.NewRecorder[int](t)

		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		rec.AwaitCompleted(time.Second)

		// Only the last inner should be processed
		assertlite.EqualSlices(t, []int{3}, rec.Values())
	})

	t.Run("cancels previous inners before they emit", func(t *testing.T) {
		src := obs.Of(1, 2)

		out := SwitchMap(func(v int) obs.Observable[int] {
			return obs.Async(func(o *obs.Observer[int]) {
				time.Sleep(50 * time.Millisecond)
				o.Next(v)
				o.Complete()
			})
		})(src)

		rec, o := observers.NewRecorder[int](t)

		start := time.Now()

		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		rec.AwaitCompleted(time.Second)

		elapsed := time.Since(start)

		// Only the second inner should be processed
		assertlite.EqualSlices(t, []int{2}, rec.Values())

		// Completion must wait for the last inner
		assertlite.True(t, elapsed >= 50*time.Millisecond)
	})

	t.Run("propagates inner error immediately", func(t *testing.T) {
		expErr := errors.New("inner error")

		src := obs.Of(1, 2)

		out := SwitchMap(func(v int) obs.Observable[int] {
			if v == 2 {
				return obs.Async(func(o *obs.Observer[int]) {
					time.Sleep(20 * time.Millisecond)
					o.Error(expErr)
				})
			}

			return obs.Async(func(o *obs.Observer[int]) {
				time.Sleep(100 * time.Millisecond)
				o.Next(v)
				o.Complete()
			})
		})(src)

		rec, o := observers.NewRecorder[int](t)

		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		assertlite.IsErrorOf(t, expErr, rec.AwaitError(time.Second))

		// Cancelled inners must not emit
		assertlite.IsNil(t, rec.Values())
		assertlite.False(t, rec.Completed())
	})

	t.Run("propagates source error and cancels active inner", func(t *testing.T) {
		expErr := errors.New("source error")

		src := obs.FromFunc[int](func(o *obs.Observer[int]) {
			o.Next(1)
			o.Error(expErr)
		})

		out := SwitchMap(func(v int) obs.Observable[int] {
			return obs.Async(func(o *obs.Observer[int]) {
				time.Sleep(100 * time.Millisecond)
				o.Next(v)
				o.Complete()
			})
		})(src)

		rec, o := observers.NewRecorder[int](t)

		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		assertlite.IsErrorOf(t, expErr, rec.AwaitError(time.Second))
		assertlite.IsNil(t, rec.Values())
		assertlite.False(t, rec.Completed())
	})
}

// TestSwitchMap_Async_Sync
//
// Source emits asynchronously
// Inner observables emit synchronously
//
// In this axis, inners do not overlap, but the flow and the test runs in parallel.
// SwitchMap degenerates to map + flatten.
//
// We should assert:
//  1. Project all source values
//  2. Forward all inner values
//  3. No cancellation affects emissions
//  4. Complete only after source completes
//  5. Propagate inner errors immediately
//  6. Propagate source errors immediately
func TestSwitchMap_Async_Sync(t *testing.T) {
	t.Run("projects and forwards all inner values", func(t *testing.T) {
		// should propagate all because despite the asynchronous source,
		// the following operation still blocking the asynchronous thread
		src := obs.AsyncIntervalOf(10*time.Millisecond, 1, 2, 3)

		out := SwitchMap(func(v int) obs.Observable[int] {
			return obs.Of(v*10, v*10+1)
		})(src)

		rec, o := observers.NewRecorder[int](t)

		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		rec.AwaitCompleted(time.Second)

		values := rec.Values()

		assertlite.EqualSlices(t,
			[]int{10, 11, 20, 21, 30, 31},
			values,
		)
	})

	t.Run("completes only after async source completes", func(t *testing.T) {
		var completed atomic.Bool

		src := obs.Async(func(o *obs.Observer[int]) {
			o.Next(1)
			time.Sleep(50 * time.Millisecond)
			o.Complete()
		})

		out := SwitchMap(func(v int) obs.Observable[int] {
			return obs.Of(v)
		})(src)

		rec, o := observers.NewRecorderFor(t,
			observers.NewComplete[int](func() {
				completed.Store(true)
			}),
		)

		start := time.Now()

		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		// already asserts completion
		rec.AwaitCompleted(time.Second)
		elapsed := time.Since(start)
		assertlite.True(t, elapsed >= 50*time.Millisecond)
	})

	t.Run("propagates inner error immediately", func(t *testing.T) {
		expErr := errors.New("inner error")

		src := obs.Async(func(o *obs.Observer[int]) {
			o.Next(1)
			o.Next(2)
			o.Complete()
		})

		out := SwitchMap(func(v int) obs.Observable[int] {
			if v == 2 {
				return obs.Error[int](expErr)
			}
			return obs.Of(v)
		})(src)

		rec, o := observers.NewRecorder[int](t)

		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		assertlite.IsErrorOf(t, expErr, rec.AwaitError(time.Second))
		// the second value errors
		assertlite.EqualSlices(t, []int{1}, rec.Values())
	})

	t.Run("propagates source error immediately", func(t *testing.T) {
		expErr := errors.New("source error")

		src := obs.Async(func(o *obs.Observer[int]) {
			o.Next(1)
			o.Error(expErr)
		})

		out := SwitchMap(func(v int) obs.Observable[int] {
			return obs.Of(v)
		})(src)

		rec, o := observers.NewRecorder[int](t)

		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		assertlite.IsErrorOf(t, expErr, rec.AwaitError(time.Second))
		assertlite.False(t, rec.Completed())
	})
}

// TestSwitchMap_Async_Async
//
// Source emits asynchronously
// Inner observables emit asynchronously
//
// This is the full SwitchMap semantics axis.
//
// We should assert:
//  1. Cancel previous inners on new source emission
//  2. Forward values only from the latest inner
//  3. Discard values from cancelled inners
//  4. Complete only after source completes AND latest inner completes
//  5. Propagate inner errors immediately
//  6. Propagate source errors immediately
//  7. Never emit after termination
func TestSwitchMap_Async_Async(t *testing.T) {
	t.Run("switches between overlapping async inners", func(t *testing.T) {
		src := obs.AsyncIntervalOf(20*time.Millisecond, 1, 2, 3)

		out := SwitchMap(func(v int) obs.Observable[string] {
			return obs.Async(func(o *obs.Observer[string]) {
				o.Next(fmt.Sprintf("%d-a", v))
				time.Sleep(40 * time.Millisecond)
				o.Next(fmt.Sprintf("%d-b", v))
				o.Complete()
			})
		})(src)

		rec, o := observers.NewRecorder[string](t)

		_, err := out.Sub((*obs.Observer[string])(o))
		assertlite.NoError(t, err)

		rec.AwaitCompleted(time.Second)

		values := rec.Values()

		// Only values from the LAST inner (v=3) will be fully processed
		assertlite.EqualSlices(t,
			[]string{"1-a", "2-a", "3-a", "3-b"},
			values,
		)
	})

	t.Run("defers completion until latest inner completes", func(t *testing.T) {
		src := obs.Async(func(o *obs.Observer[int]) {
			o.Next(1)
			time.Sleep(10 * time.Millisecond)
			o.Complete()
		})

		out := SwitchMap(func(v int) obs.Observable[int] {
			return obs.Async(func(o *obs.Observer[int]) {
				time.Sleep(50 * time.Millisecond)
				o.Next(v)
				o.Complete()
			})
		})(src)

		rec, o := observers.NewRecorder[int](t)

		start := time.Now()

		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		rec.AwaitCompleted(time.Second)

		elapsed := time.Since(start)

		assertlite.True(t, elapsed >= 50*time.Millisecond)
		assertlite.EqualSlices(t, []int{1}, rec.Values())
	})

	t.Run("propagates inner error and cancels previous inners", func(t *testing.T) {
		expErr := errors.New("inner error")

		src := obs.Async(func(o *obs.Observer[int]) {
			o.Next(1)
			time.Sleep(10 * time.Millisecond)
			o.Next(2)
		})

		out := SwitchMap(func(v int) obs.Observable[int] {
			if v == 2 {
				return obs.Async(func(o *obs.Observer[int]) {
					time.Sleep(20 * time.Millisecond)
					o.Error(expErr)
				})
			}

			return obs.Async(func(o *obs.Observer[int]) {
				time.Sleep(100 * time.Millisecond)
				o.Next(v)
				o.Complete()
			})
		})(src)

		rec, o := observers.NewRecorder[int](t)

		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		assertlite.IsErrorOf(t, expErr, rec.AwaitError(time.Second))

		// Cancelled inners must not emit
		assertlite.IsNil(t, rec.Values())
		assertlite.False(t, rec.Completed())
	})

	t.Run("propagates source error and cancels active inner", func(t *testing.T) {
		expErr := errors.New("source error")

		src := obs.Async(func(o *obs.Observer[int]) {
			o.Next(1)
			time.Sleep(20 * time.Millisecond)
			o.Error(expErr)
		})

		out := SwitchMap(func(v int) obs.Observable[int] {
			return obs.Async(func(o *obs.Observer[int]) {
				time.Sleep(100 * time.Millisecond)
				o.Next(v)
				o.Complete()
			})
		})(src)

		rec, o := observers.NewRecorder[int](t)

		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		assertlite.IsErrorOf(t, expErr, rec.AwaitError(time.Second))
		assertlite.IsNil(t, rec.Values())
		assertlite.False(t, rec.Completed())
	})
}
