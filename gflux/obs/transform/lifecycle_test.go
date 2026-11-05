// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package transform

import (
	"errors"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/andrerrcosta2/gtools/core/testlite/assertlite"
	"github.com/andrerrcosta2/gtools/gflux/obs"
	"github.com/andrerrcosta2/gtools/gflux/obs/internal/observers"
	"github.com/andrerrcosta2/gtools/gflux/obs/internal/subscr"
)

// TestDoOnComplete_Sync
//
// Should assert:
//  1. Values from the source are forwarded unchanged
//  2. The side-effect runs exactly once on normal completion
//  3. The side-effect runs before downstream completion
//  4. Errors do NOT trigger the side-effect
//  5. Source error is propagated immediately
//  6. Unsubscription before completion does NOT trigger the side-effect
//  7. Completion happens at most once
//  8. State is local per subscription
func TestDoOnComplete_Sync(t *testing.T) {
	t.Run("values are forwarded unchanged", func(t *testing.T) {
		called := 0
		src := obs.Of(1, 2, 3)
		out := DoOnComplete[int](func() {
			called++
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		rec.AwaitCompleted(time.Second)

		assertlite.EqualSlices(t, []int{1, 2, 3}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
		assertlite.True(t, 1 == called)
	})

	t.Run("side-effect runs before downstream completion", func(t *testing.T) {
		var order []string

		src := obs.Of(1)

		out := DoOnComplete[int](func() {
			order = append(order, "fn")
		})(src)

		_, _ = out.Sub(obs.NewObserver(
			func(v int) {},
			func(err error) {},
			func() {
				order = append(order, "complete")
			},
		))

		assertlite.EqualSlices(t, []string{"fn", "complete"}, order)
	})

	t.Run("side-effect is not called on error", func(t *testing.T) {
		err := errors.New("fail")
		src := obs.Error[int](err)

		out := DoOnComplete[int](func() {
			t.Error("should not be called on error")
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		rec.AwaitTermination(time.Second)

		assertlite.False(t, rec.Completed())
		assertlite.IsErrorOf(t, err, rec.Err())
	})

	t.Run("unsubscription before completion does not trigger side-effect", func(t *testing.T) {
		src := obs.FromFunc[int](func(o *obs.Observer[int]) {
			o.Next(1)
			// never completes
		})

		out := DoOnComplete[int](func() {
			t.Error("should not be called if it doesn't complete")
		})(src)

		sub, _ := out.Sub(&obs.Observer[int]{})
		sub.Unsub()
		time.Sleep(10 * time.Millisecond)
	})

	t.Run("completion happens at most once", func(t *testing.T) {
		called := 0

		src := obs.FromFunc[int](func(o *obs.Observer[int]) {
			o.Complete()
			o.Complete()
		})

		out := DoOnComplete[int](func() {
			called++
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		rec.AwaitCompleted(time.Second)

		assertlite.True(t, rec.Completed())
		assertlite.True(t, 1 == called)
	})

	t.Run("states is local per subscription", func(t *testing.T) {
		var calls1, calls2 int

		src := obs.Of(1, 2)

		out := DoOnComplete[int](func() {
			calls1++
		})(src)

		out2 := DoOnComplete[int](func() {
			calls2++
		})(src)

		rec1, o1 := observers.NewRecorder[int](t)
		rec2, o2 := observers.NewRecorder[int](t)

		_, _ = out.Sub((*obs.Observer[int])(o1))
		_, _ = out2.Sub((*obs.Observer[int])(o2))

		rec1.AwaitCompleted(time.Second)
		rec2.AwaitCompleted(time.Second)

		assertlite.True(t, 1 == calls1)
		assertlite.True(t, 1 == calls2)
	})
}

// TestDoOnComplete_Async
//
// Should assert:
//  1. Values from an async source are forwarded unchanged
//  2. The side-effect runs exactly once on async completion
//  3. The side-effect runs before downstream completion
//  4. Errors do NOT trigger the side-effect
//  5. Source error propagates immediately
//  6. Unsubscription before async completion does NOT trigger the side-effect
//  7. Completion happens at most once under concurrency
//  8. State is local per subscription
func TestDoOnComplete_Async(t *testing.T) {
	t.Run("values are forwarded unchanged", func(t *testing.T) {
		called := atomic.Int32{}

		src := obs.AsyncIntervalOf[int](5*time.Millisecond, 1, 2, 3)

		out := DoOnComplete[int](func() {
			called.Add(1)
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		rec.AwaitCompleted(time.Second)

		assertlite.EqualSlices(t, []int{1, 2, 3}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
		assertlite.Equals(t, int32(1), called.Load())
	})

	t.Run("side-effect runs before downstream completion", func(t *testing.T) {
		var order []string

		src := obs.AsyncIntervalOf[int](5*time.Millisecond, 1)

		out := DoOnComplete[int](func() {
			order = append(order, "fn")
		})(src)

		_, _ = out.Sub(&obs.Observer[int]{
			Complete: func() {
				order = append(order, "complete")
			},
		})

		time.Sleep(50 * time.Millisecond)

		assertlite.EqualSlices(t, []string{"fn", "complete"}, order)
	})

	t.Run("side-effect is not called on async error", func(t *testing.T) {
		err := errors.New("async fail")
		src := obs.AsyncError[int](err)

		out := DoOnComplete[int](func() {
			t.Error("should not be called on error")
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		rec.AwaitTermination(time.Second)

		assertlite.False(t, rec.Completed())
		assertlite.IsErrorOf(t, err, rec.Err())
	})

	t.Run("unsubscription before completion does not trigger side-effect", func(t *testing.T) {
		src := obs.AsyncIntervalOf[int](50*time.Millisecond, 1, 2, 3)

		out := DoOnComplete[int](func() {
			t.Error("should not be called after unsubscription")
		})(src)

		sub, _ := out.Sub(&obs.Observer[int]{})
		time.Sleep(10 * time.Millisecond)
		sub.Unsub()
		time.Sleep(30 * time.Millisecond)
	})

	t.Run("completion happens at most once under concurrency", func(t *testing.T) {
		called := atomic.Int32{}

		src := obs.FromFunc[int](func(o *obs.Observer[int]) {
			go o.Complete()
			go o.Complete()
		})

		out := DoOnComplete[int](func() {
			called.Add(1)
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		rec.AwaitCompleted(time.Second)

		assertlite.True(t, rec.Completed())
		assertlite.Equals(t, int32(1), called.Load())
	})

	t.Run("states is local per subscription", func(t *testing.T) {
		c1 := atomic.Int32{}
		c2 := atomic.Int32{}

		src := obs.AsyncIntervalOf[int](5*time.Millisecond, 1, 2)

		out := DoOnComplete[int](func() {
			c1.Add(1)
		})(src)

		out2 := DoOnComplete[int](func() {
			c2.Add(1)
		})(src)

		rec1, o1 := observers.NewRecorder[int](t)
		rec2, o2 := observers.NewRecorder[int](t)

		_, _ = out.Sub((*obs.Observer[int])(o1))
		_, _ = out2.Sub((*obs.Observer[int])(o2))

		rec1.AwaitCompleted(time.Second)
		rec2.AwaitCompleted(time.Second)

		assertlite.Equals(t, int32(1), c1.Load())
		assertlite.Equals(t, int32(1), c2.Load())
	})
}

// TestDoOnError_Sync
//
// Should assert:
//  1. Values before an error are forwarded unchanged
//  2. The side-effect runs exactly once when the source errors
//  3. The side-effect runs BEFORE the downstream error is observed
//  4. The original error is forwarded unchanged
//  5. Completion does NOT trigger the side-effect
//  6. Side-effect panic propagates and terminates the subscription
//  7. State is local per subscription
func TestDoOnError_Sync(t *testing.T) {
	t.Run("values before error are forwarded unchanged", func(t *testing.T) {
		err := errors.New("fail")
		called := atomic.Int32{}

		src := obs.FromFunc(func(o *obs.Observer[int]) {
			o.Next(1)
			o.Next(2)
			o.Error(err)
		})

		out := DoOnError[int](func(error) {
			called.Add(1)
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		rec.AwaitTermination(time.Second)

		assertlite.EqualSlices(t, []int{1, 2}, rec.Values())
		assertlite.False(t, rec.Completed())
		assertlite.IsErrorOf(t, err, rec.Err())
		assertlite.Equals(t, int32(1), called.Load())
	})

	t.Run("side-effect runs before downstream error", func(t *testing.T) {
		err := errors.New("fail")

		var order []string

		src := obs.Error[int](err)

		out := DoOnError[int](func(error) {
			order = append(order, "fn")
		})(src)

		_, _ = out.Sub(&obs.Observer[int]{
			Error: func(err error) {
				order = append(order, "error")
			},
		})

		assertlite.EqualSlices(t, []string{"fn", "error"}, order)
	})

	t.Run("completion does not trigger side-effect", func(t *testing.T) {
		called := atomic.Int32{}

		src := obs.Of(1, 2, 3)

		out := DoOnError[int](func(error) {
			called.Add(1)
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		rec.AwaitCompleted(time.Second)

		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
		assertlite.Equals(t, int32(0), called.Load())
	})

	t.Run("panic inside side-effect propagates", func(t *testing.T) {
		err := errors.New("fail")
		src := obs.Error[int](err)
		out := DoOnError[int](func(error) {
			panic("panic")
		})(src)

		assertlite.Panics(t, func() {
			_, _ = out.Sub(&obs.Observer[int]{})
		})
	})

	t.Run("states is local per subscription", func(t *testing.T) {
		err := errors.New("fail")

		c1 := atomic.Int32{}
		c2 := atomic.Int32{}

		src := obs.Error[int](err)

		out1 := DoOnError[int](func(error) {
			c1.Add(1)
		})(src)

		out2 := DoOnError[int](func(error) {
			c2.Add(1)
		})(src)

		rec1, o1 := observers.NewRecorder[int](t)
		rec2, o2 := observers.NewRecorder[int](t)

		_, _ = out1.Sub((*obs.Observer[int])(o1))
		_, _ = out2.Sub((*obs.Observer[int])(o2))

		rec1.AwaitTermination(time.Second)
		rec2.AwaitTermination(time.Second)

		assertlite.Equals(t, int32(1), c1.Load())
		assertlite.Equals(t, int32(1), c2.Load())
	})
}

// TestDoOnError_Async
//
// Should assert:
//  1. Values emitted before an async error are forwarded
//  2. The side-effect runs exactly once on async error
//  3. The side-effect runs BEFORE the downstream error
//  4. Completion does NOT trigger the side-effect
//  5. Unsubscription before async error prevents the side-effect
//  6. Panics inside the side-effect propagates
//  7. Completion happens at most once
//  8. State is local per subscription
func TestDoOnError_Async(t *testing.T) {
	t.Run("values before async error are forwarded", func(t *testing.T) {
		err := errors.New("fail")
		called := atomic.Int32{}

		src := obs.Async(func(o *obs.Observer[int]) {
			o.Next(1)
			o.Next(2)
			o.Error(err)
		})

		out := DoOnError[int](func(error) {
			called.Add(1)
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		rec.AwaitTermination(time.Second)

		assertlite.EqualSlices(t, []int{1, 2}, rec.Values())
		assertlite.False(t, rec.Completed())
		assertlite.IsErrorOf(t, err, rec.Err())
		assertlite.Equals(t, int32(1), called.Load())
	})

	t.Run("side-effect runs before downstream async error", func(t *testing.T) {
		err := errors.New("fail")

		var orderMu sync.Mutex
		var order []string

		src := obs.AsyncError[int](err)

		out := DoOnError[int](func(error) {
			orderMu.Lock()
			order = append(order, "fn")
			orderMu.Unlock()
		})(src)

		rec, o := observers.NewRecorderFor[int](t, &observers.Observer[int]{
			Error: func(err error) {
				orderMu.Lock()
				order = append(order, "error")
				orderMu.Unlock()
			},
		})

		_, _ = out.Sub((*obs.Observer[int])(o))
		rec.AwaitTermination(time.Second)
		assertlite.EqualSlices(t, []string{"fn", "error"}, order)
	})

	t.Run("completion does not trigger side-effect", func(t *testing.T) {
		src := obs.AsyncOf(1, 2, 3)

		out := DoOnError[int](func(error) {
			t.Error("should not be called on completion")
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		rec.AwaitCompleted(time.Second)
		assertlite.NoError(t, rec.Err())
	})

	t.Run("unsubscription before async error prevents side-effect", func(t *testing.T) {
		src := obs.Async[int](func(o *obs.Observer[int]) {
			time.Sleep(20 * time.Millisecond)
			o.Error(errors.New("late error"))
		})

		out := DoOnError[int](func(error) {
			t.Error("shouldn't call pipe after unsubscription")
		})(src)

		_, o := observers.NewRecorder[int](t)
		sub, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)
		sub.Unsub()
		time.Sleep(30 * time.Millisecond)
	})

	t.Run("completion happens at most once", func(t *testing.T) {
		called := atomic.Int32{}

		src := obs.AsyncError[int](errors.New("fail"))

		out := DoOnError[int](func(error) {
			called.Add(1)
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		rec.AwaitTermination(time.Second)

		assertlite.Equals(t, int32(1), called.Load())
	})

	t.Run("states is local per subscription", func(t *testing.T) {
		err := errors.New("fail")

		c1 := atomic.Int32{}
		c2 := atomic.Int32{}

		src := obs.AsyncError[int](err)

		out := DoOnError[int](func(error) {
			c1.Add(1)
		})(src)

		out2 := DoOnError[int](func(error) {
			c2.Add(1)
		})(src)

		rec1, o1 := observers.NewRecorder[int](t)
		rec2, o2 := observers.NewRecorder[int](t)

		_, _ = out.Sub((*obs.Observer[int])(o1))
		_, _ = out2.Sub((*obs.Observer[int])(o2))

		rec1.AwaitTermination(time.Second)
		rec2.AwaitTermination(time.Second)

		assertlite.Equals(t, int32(1), c1.Load())
		assertlite.Equals(t, int32(1), c2.Load())
	})
}

// TestFinalize_Sync
//
// Should assert:
//  1. Values are forwarded unchanged
//  2. Cleanup runs on normal completion
//  3. Cleanup runs on error
//  4. Cleanup runs on explicit unsubscription
//  5. Cleanup runs exactly once
//  6. Cleanup runs before downstream termination
//  7. Cleanup runs if source.Sub returns an error
//  8. Cleanup is not run prematurely
func TestFinalize_Sync(t *testing.T) {
	t.Run("values are forwarded unchanged", func(t *testing.T) {
		called := atomic.Int32{}

		src := obs.Of(1, 2, 3)

		out := Finalize[int](func() {
			called.Add(1)
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		rec.AwaitCompleted(time.Second)

		assertlite.EqualSlices(t, []int{1, 2, 3}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
		assertlite.Equals(t, int32(1), called.Load())
	})

	t.Run("cleanup runs on sync completion", func(t *testing.T) {
		called := atomic.Int32{}

		src := obs.Of(1)

		out := Finalize[int](func() {
			called.Add(1)
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		rec.AwaitCompleted(time.Second)

		assertlite.Equals(t, int32(1), called.Load())
	})

	t.Run("cleanup runs on sync error", func(t *testing.T) {
		err := errors.New("fail")
		called := atomic.Int32{}

		src := obs.Error[int](err)

		out := Finalize[int](func() {
			called.Add(1)
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		rec.AwaitTermination(time.Second)

		assertlite.False(t, rec.Completed())
		assertlite.IsErrorOf(t, err, rec.Err())
		assertlite.Equals(t, int32(1), called.Load())
	})

	t.Run("cleanup runs on explicit unsubscription", func(t *testing.T) {
		called := atomic.Int32{}

		src := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			o.Next(1)
			return subscr.Noop{}, nil
		})

		out := Finalize[int](func() {
			called.Add(1)
		})(src)

		sub, _ := out.Sub(&obs.Observer[int]{})

		sub.Unsub()

		assertlite.Equals(t, int32(1), called.Load())
	})

	t.Run("cleanup runs exactly once", func(t *testing.T) {
		called := atomic.Int32{}

		src := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			o.Error(errors.New("fail"))
			o.Complete()
			return subscr.Noop{}, nil
		})

		out := Finalize[int](func() {
			called.Add(1)
		})(src)

		rec, o := observers.NewRecorder[int](t)
		sub, _ := out.Sub((*obs.Observer[int])(o))

		sub.Unsub()
		rec.AwaitTermination(time.Second)

		assertlite.Equals(t, int32(1), called.Load())
	})

	t.Run("cleanup runs before downstream termination", func(t *testing.T) {
		var order []string

		src := obs.Error[int](errors.New("fail"))

		out := Finalize[int](func() {
			order = append(order, "cleanup")
		})(src)

		_, _ = out.Sub(&obs.Observer[int]{
			Error: func(err error) {
				order = append(order, "error")
			},
			Complete: func() {
				order = append(order, "complete")
			},
		})

		assertlite.EqualSlices(t, []string{"cleanup", "error"}, order)
	})

	t.Run("cleanup runs if source.Sub returns error", func(t *testing.T) {
		called := atomic.Int32{}
		errSub := errors.New("sub fail")

		src := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			return nil, errSub
		})

		out := Finalize[int](func() {
			called.Add(1)
		})(src)

		_, err := out.Sub(&obs.Observer[int]{})

		assertlite.IsErrorOf(t, err, errSub)
		assertlite.Equals(t, int32(1), called.Load())
	})
}

// TestFinalize_Async
//
// Should assert:
//  1. Values are forwarded unchanged
//  2. Cleanup runs on normal completion
//  3. Cleanup runs on error
//  4. Cleanup runs on explicit unsubscription
//  5. Cleanup runs exactly once
//  6. Cleanup runs before downstream termination
//  7. Cleanup runs if source.Sub returns an error
//  8. Cleanup is not run prematurely
func TestFinalize_Async(t *testing.T) {
	t.Run("values are forwarded unchanged", func(t *testing.T) {
		called := atomic.Int32{}
		src := obs.Async[int](func(o *obs.Observer[int]) {
			o.Next(1)
			o.Next(2)
			o.Next(3)
			o.Complete()
		})

		out := Finalize[int](func() {
			called.Add(1)
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		rec.AwaitCompleted(time.Second)

		assertlite.EqualSlices(t, []int{1, 2, 3}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
		assertlite.Equals(t, int32(1), called.Load())
	})

	t.Run("cleanup runs on async completion", func(t *testing.T) {
		called := atomic.Int32{}

		src := obs.Async[int](func(o *obs.Observer[int]) {
			o.Next(1)
			o.Complete()
		})

		out := Finalize[int](func() {
			called.Add(1)
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		rec.AwaitCompleted(time.Second)

		assertlite.Equals(t, int32(1), called.Load())
	})

	t.Run("cleanup runs on async error", func(t *testing.T) {
		fail := errors.New("fail")
		called := atomic.Int32{}

		src := obs.AsyncError[int](fail)

		out := Finalize[int](func() {
			called.Add(1)
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		rec.AwaitTermination(time.Second)

		assertlite.False(t, rec.Completed())
		assertlite.IsErrorOf(t, fail, rec.Err())
		assertlite.Equals(t, int32(1), called.Load())
	})

	t.Run("cleanup runs on explicit unsubscription", func(t *testing.T) {
		called := atomic.Int32{}
		block := make(chan struct{})

		src := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			go func() {
				o.Next(1)
				<-block
			}()
			return subscr.Func(func() {
				close(block)
			}), nil
		})

		out := Finalize[int](func() {
			called.Add(1)
		})(src)

		sub, _ := out.Sub(&obs.Observer[int]{})
		sub.Unsub()

		assertlite.Equals(t, int32(1), called.Load())
	})

	t.Run("cleanup runs exactly once", func(t *testing.T) {
		called := atomic.Int32{}

		src := obs.Async[int](func(o *obs.Observer[int]) {
			o.Error(errors.New("fail"))
			o.Complete()
		})

		out := Finalize[int](func() {
			called.Add(1)
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		rec.AwaitTermination(time.Second)

		assertlite.Equals(t, int32(1), called.Load())
	})

	t.Run("cleanup runs before downstream termination", func(t *testing.T) {
		var order []string

		src := obs.AsyncError[int](errors.New("fail"))
		out := Finalize[int](func() {
			order = append(order, "cleanup")
		})(src)

		rec, o := observers.NewRecorderFor[int](t, &observers.Observer[int]{
			Error: func(err error) {
				order = append(order, "error")
			},
			Complete: func() {
				order = append(order, "complete")
			},
		})

		_, _ = out.Sub((*obs.Observer[int])(o))
		rec.AwaitTermination(time.Second)
		assertlite.EqualSlices(t, []string{"cleanup", "error"}, order)
	})

	t.Run("cleanup runs if source subscribe returns error", func(t *testing.T) {
		called := atomic.Int32{}
		errSub := errors.New("sub fail")

		src := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			go func() {
				o.Next(1)
			}()
			return nil, errSub
		})

		out := Finalize[int](func() {
			called.Add(1)
		})(src)

		_, err := out.Sub(&obs.Observer[int]{})

		assertlite.IsErrorOf(t, err, errSub)
		assertlite.True(t, int32(1) == called.Load())
	})

	t.Run("cleanup is not run before complete", func(t *testing.T) {
		called := atomic.Int32{}
		block := make(chan struct{})

		src := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			go func() {
				o.Next(1)
				o.Next(2)
				<-block
				o.Complete()
			}()
			return subscr.Func(func() {
				close(block)
			}), nil
		})

		out := Finalize[int](func() {
			called.Add(1)
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		rec.AwaitNext(time.Second) // one next is enough

		assertlite.Equals(t, int32(0), called.Load())

		close(block)
		rec.AwaitCompleted(time.Second)

		assertlite.Equals(t, int32(1), called.Load())
	})
}

// TestTap_Sync
//
// Should assert:
//  1. Values are forwarded unchanged
//  2. Tap function is called for each value
//  3. Tap is called before downstream receives the value
//  4. Errors are forwarded unchanged
//  5. Completion is forwarded unchanged
//  6. Tap is not called on error
//  7. Tap is not called on completion
//  8. Panics inside Tap propagate
func TestTap_Sync(t *testing.T) {
	t.Run("values are forwarded unchanged", func(t *testing.T) {
		var seen []int

		src := obs.Of(1, 2, 3)

		out := Tap(func(v int) {
			seen = append(seen, v)
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		rec.AwaitCompleted(time.Second)

		assertlite.EqualSlices(t, []int{1, 2, 3}, seen)
		assertlite.EqualSlices(t, []int{1, 2, 3}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("tap is called before downstream receives value", func(t *testing.T) {
		var order []string

		src := obs.Of(1)

		out := Tap(func(v int) {
			order = append(order, "tap")
		})(src)

		_, _ = out.Sub(&obs.Observer[int]{
			Next: func(v int) {
				order = append(order, "next")
			},
			Complete: func() {
				order = append(order, "complete")
			},
		})

		assertlite.EqualSlices(t, []string{"tap", "next", "complete"}, order)
	})

	t.Run("errors are forwarded unchanged", func(t *testing.T) {
		err := errors.New("fail")

		src := obs.Error[int](err)
		out := Tap(func(v int) {
			t.Error("tab shouldn't be called on error")
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		assertlite.IsErrorOf(t, err, rec.AwaitError(time.Second))
		assertlite.False(t, rec.Completed())
	})

	t.Run("completion is forwarded unchanged", func(t *testing.T) {
		src := obs.Empty[int]()

		out := Tap(func(v int) {
			t.Error("tap should not be called on completion")
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		rec.AwaitCompleted(time.Second)

		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("panics inside tap propagate", func(t *testing.T) {
		src := obs.Of(1)

		out := Tap(func(v int) {
			panic("fail")
		})(src)

		assertlite.Panics(t, func() {
			_, _ = out.Sub(&obs.Observer[int]{})
		})
	})
}

// TestTap_Async
//
// Should assert:
//  1. Values are forwarded unchanged
//  2. Tap function is called for each value
//  3. Tap is called before downstream receives the value
//  4. Errors are forwarded unchanged
//  5. Completion is forwarded unchanged
//  6. Tap is not called on error
//  7. Tap is not called on completion
//  8. Panics inside Tap propagate
func TestTap_Async(t *testing.T) {
	t.Run("values are forwarded unchanged", func(t *testing.T) {
		var seen []int

		src := obs.Async[int](func(o *obs.Observer[int]) {
			o.Next(1)
			o.Next(2)
			o.Next(3)
			o.Complete()
		})

		out := Tap(func(v int) {
			seen = append(seen, v)
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		rec.AwaitCompleted(time.Second)

		assertlite.EqualSlices(t, []int{1, 2, 3}, seen)
		assertlite.EqualSlices(t, []int{1, 2, 3}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("tap is called before downstream receives value", func(t *testing.T) {
		var order []string

		src := obs.Async[int](func(o *obs.Observer[int]) {
			o.Next(1)
			o.Complete()
		})

		out := Tap(func(v int) {
			order = append(order, "tap")
		})(src)

		rec, o := observers.NewRecorderFor[int](t, &observers.Observer[int]{
			Next: func(v int) {
				order = append(order, "next")
			},
			Complete: func() {
				order = append(order, "complete")
			},
		})

		_, _ = out.Sub((*obs.Observer[int])(o))

		rec.AwaitCompleted(time.Second)
		assertlite.EqualSlices(t, []string{"tap", "next", "complete"}, order)
	})

	t.Run("errors are forwarded unchanged", func(t *testing.T) {
		err := errors.New("fail")

		src := obs.AsyncError[int](err)

		out := Tap(func(v int) {
			t.Error("tap should not be called on error")
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		assertlite.IsErrorOf(t, err, rec.AwaitError(time.Second))
		assertlite.False(t, rec.Completed())
	})

	t.Run("completion is forwarded unchanged", func(t *testing.T) {
		src := obs.Async[int](func(o *obs.Observer[int]) {
			o.Complete()
		})

		out := Tap(func(v int) {
			t.Error("tab shouldn't be called on complete")
		})(src)

		rec, o := observers.NewRecorder[int](t)
		_, _ = out.Sub((*obs.Observer[int])(o))

		rec.AwaitCompleted(time.Second)

		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})
}
