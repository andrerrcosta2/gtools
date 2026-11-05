// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package observables

import (
	"errors"
	"testing"
	"time"

	"github.com/andrerrcosta2/gtools/core/testlite/assertlite"
	"github.com/andrerrcosta2/gtools/gflux/obs/internal/observers"
)

// TestAsyncObs tests the asynchronous Async[T any] observable.
//
// Async wraps a user-provided function to emit values asynchronously.
// We want to assert:
//  1. Values are emitted correctly
//  2. Completion is signaled after emissions
//  3. Supports multiple asynchronous emissions
//  4. Errors propagate correctly
//  5. Subscription returns immediately while emissions happen asynchronously
func TestAsyncObs(t *testing.T) {
	t.Run("emits values asynchronously and completes", func(t *testing.T) {
		src := Async(func(o *observers.Observer[int]) {
			o.Next(1)
			o.Next(2)
			o.Next(3)
			o.Complete()
		})

		rec, o := observers.NewRecorder[int](t)
		_, _ = src.Sub(o)

		// Wait a little to ensure goroutine completes
		time.Sleep(10 * time.Millisecond)

		assertlite.EqualSlices(t, []int{1, 2, 3}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("propagates errors correctly", func(t *testing.T) {
		failure := errors.New("failure")
		src := Async(func(o *observers.Observer[int]) {
			o.Next(42)
			o.Error(failure)
		})

		rec, o := observers.NewRecorder[int](t)
		_, _ = src.Sub(o)

		time.Sleep(10 * time.Millisecond)

		assertlite.EqualSlices(t, []int{42}, rec.Values())
		assertlite.IsErrorOf(t, rec.Err(), failure)
		assertlite.False(t, rec.Completed())
	})

	t.Run("subscription returns immediately while emissions happen asynchronously", func(t *testing.T) {
		src := Async(func(o *observers.Observer[int]) {
			time.Sleep(20 * time.Millisecond)
			o.Next(1)
			o.Complete()
		})

		rec, o := observers.NewRecorder[int](t)
		subStart := time.Now()
		_, _ = src.Sub(o)
		subEnd := time.Now()

		assertlite.True(t, subEnd.Sub(subStart) < 10*time.Millisecond, "subscription should return immediately")

		rec.AwaitCompleted(time.Second)
		assertlite.EqualSlices(t, []int{1}, rec.Values())
		assertlite.True(t, rec.Completed())
	})
}

// TestAsyncChanObs tests the asynchronous AsyncChan[T any] observable.
//
// AsyncChan creates a cold Observable from a Go channel. We want to assert:
//  1. Values sent to the channel are emitted to the subscriber
//  2. Completion occurs when the channel is closed
//  3. Range subscription gets its own goroutine reading the channel
//  4. Subscription returns immediately
func TestAsyncChanObs(t *testing.T) {
	t.Run("emits values from channel and completes", func(t *testing.T) {
		ch := make(chan int)
		src := AsyncChan(ch)

		rec, o := observers.NewRecorder[int](t)
		_, _ = src.Sub(o)

		go func() {
			ch <- 1
			ch <- 2
			ch <- 3
			close(ch)
		}()

		time.Sleep(10 * time.Millisecond)

		assertlite.EqualSlices(t, []int{1, 2, 3}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("subscription returns immediately even if channel emits later", func(t *testing.T) {
		ch := make(chan int)
		src := AsyncChan(ch)

		rec, o := observers.NewRecorder[int](t)
		start := time.Now()
		_, _ = src.Sub(o)
		subEnd := time.Now()

		assertlite.True(t, subEnd.Sub(start) < 5*time.Millisecond,
			"subscription should return immediately")

		go func() {
			ch <- 42
			close(ch)
		}()

		time.Sleep(10 * time.Millisecond)
		assertlite.EqualSlices(t, []int{42}, rec.Values())
		assertlite.True(t, rec.Completed())
	})

	t.Run("completes immediately if channel is already closed", func(t *testing.T) {
		ch := make(chan int)
		close(ch)

		src := AsyncChan(ch)

		rec, o := observers.NewRecorder[int](t)
		_, _ = src.Sub(o)

		time.Sleep(5 * time.Millisecond)
		assertlite.Equals(t, 0, len(rec.Values()))
		assertlite.True(t, rec.Completed())
	})
}

// TestAsyncDelayedOfObs tests AsyncDelayedOf[T any], which emits values
// asynchronously with a per-value delay.
//
// We want to assert:
//  1. Values are emitted in order after their respective delays
//  2. Completion occurs after the last delayed value
//  3. Multiple asynchronous emissions are handled correctly
//  4. Subscription returns immediately
func TestAsyncDelayedOfObs(t *testing.T) {
	t.Run("emits values with per-value delays and completes", func(t *testing.T) {
		src := AsyncDelayedOf(func(i int, v int) time.Duration {
			return time.Duration(i*5) * time.Millisecond
		}, 1, 2, 3)

		rec, o := observers.NewRecorder[int](t)
		start := time.Now()
		_, _ = src.Sub(o)

		time.Sleep(20 * time.Millisecond) // ensure all goroutines finish

		assertlite.EqualSlices(t, []int{1, 2, 3}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
		assertlite.True(t, time.Since(start) >= 10*time.Millisecond, "delays should have been applied")
	})

	t.Run("subscription returns immediately", func(t *testing.T) {
		src := AsyncDelayedOf(func(i int, v int) time.Duration {
			return 50 * time.Millisecond
		}, 42)

		rec, o := observers.NewRecorder[int](t)
		start := time.Now()
		_, _ = src.Sub(o)
		subEnd := time.Now()

		assertlite.True(t, subEnd.Sub(start) < 5*time.Millisecond,
			"subscription should return immediately")

		time.Sleep(60 * time.Millisecond) // wait for emission
		assertlite.EqualSlices(t, []int{42}, rec.Values())
		assertlite.True(t, rec.Completed())
	})
}

// TestAsyncIntervalOfObs tests AsyncIntervalOf[T any], which emits values
// asynchronously at fixed intervals.
//
// We want to assert:
//  1. Values are emitted in order with the specified interval
//  2. Completion is delayed according to the interval
//  3. Works even when no values are provided
//  4. Subscription returns immediately
func TestAsyncIntervalOfObs(t *testing.T) {
	t.Run("emits values at fixed intervals and completes", func(t *testing.T) {
		src := AsyncIntervalOf(5*time.Millisecond, 1, 2, 3)

		rec, o := observers.NewRecorder[int](t)
		start := time.Now()
		_, _ = src.Sub(o)

		time.Sleep(25 * time.Millisecond) // ensure all emissions occur

		assertlite.EqualSlices(t, []int{1, 2, 3}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
		assertlite.True(t, time.Since(start) > 15*time.Millisecond, "intervals should have been applied")
	})

	t.Run("subscription returns immediately", func(t *testing.T) {
		src := AsyncIntervalOf(10*time.Millisecond, 42)

		rec, o := observers.NewRecorder[int](t)
		start := time.Now()
		_, _ = src.Sub(o)
		subEnd := time.Now()

		assertlite.True(t, subEnd.Sub(start) < 5*time.Millisecond, "subscription should return immediately")
		rec.AwaitCompleted(time.Second)
		assertlite.EqualSlices(t, []int{42}, rec.Values())
	})

	t.Run("completes after interval even with no values", func(t *testing.T) {
		src := AsyncIntervalOf[int](10 * time.Millisecond)

		rec, o := observers.NewRecorder[int](t)
		start := time.Now()
		_, _ = src.Sub(o)

		time.Sleep(15 * time.Millisecond)
		assertlite.Equals(t, 0, len(rec.Values()))
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
		assertlite.True(t, time.Since(start) > 10*time.Millisecond,
			"completion delayed by initial interval")
	})
}

// TestAsyncEmptyObs tests AsyncEmpty[T any], which completes asynchronously
// without emitting any values.
//
// We want to assert:
//  1. No values are emitted
//  2. Completion is signaled asynchronously
//  3. Subscription returns immediately
func TestAsyncEmptyObs(t *testing.T) {
	t.Run("completes asynchronously without emitting", func(t *testing.T) {
		src := AsyncEmpty[int]()

		rec, o := observers.NewRecorder[int](t)
		start := time.Now()
		_, _ = src.Sub(o)

		// immediately after subscription, nothing should have happened
		assertlite.Equals(t, 0, len(rec.Values()))
		assertlite.False(t, rec.Completed())

		// wait a tiny bit to allow async completion
		time.Sleep(5 * time.Millisecond)

		assertlite.Equals(t, 0, len(rec.Values()))
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
		assertlite.True(t, time.Since(start) >= 0, "should complete asynchronously")
	})

	t.Run("subscription returns immediately", func(t *testing.T) {
		src := AsyncEmpty[int]()
		rec, o := observers.NewRecorder[int](t)

		start := time.Now()
		_, _ = src.Sub(o)
		subEnd := time.Now()

		assertlite.True(t, subEnd.Sub(start) < 2*time.Millisecond, "subscription should return immediately")

		time.Sleep(5 * time.Millisecond)
		assertlite.True(t, rec.Completed())
	})
}

// TestAsyncErrorObs tests AsyncError[T any], which asynchronously emits
// a terminal error without producing any values.
//
// We want to assert:
//  1. No values are emitted
//  2. Error is emitted asynchronously
//  3. Subscription returns immediately
func TestAsyncErrorObs(t *testing.T) {
	t.Run("emits an error asynchronously without emitting values", func(t *testing.T) {
		failure := errors.New("test failure")
		src := AsyncError[int](failure)

		rec, o := observers.NewRecorder[int](t)
		start := time.Now()
		_, _ = src.Sub(o)

		// immediately after subscription, no error yet
		assertlite.Equals(t, 0, len(rec.Values()))
		assertlite.False(t, rec.Completed())
		assertlite.NoError(t, rec.Err())

		// wait a tiny bit to allow async error propagation
		time.Sleep(5 * time.Millisecond)

		assertlite.Equals(t, 0, len(rec.Values()))
		assertlite.False(t, rec.Completed())
		assertlite.IsErrorOf(t, rec.Err(), failure)
		assertlite.True(t, time.Since(start) >= 0, "error should be delivered asynchronously")
	})

	t.Run("subscription returns immediately", func(t *testing.T) {
		failure := errors.New("immediate test")
		src := AsyncError[int](failure)
		rec, o := observers.NewRecorder[int](t)

		start := time.Now()
		_, _ = src.Sub(o)
		subEnd := time.Now()

		assertlite.True(t, subEnd.Sub(start) < 2*time.Millisecond, "subscription should return immediately")

		time.Sleep(5 * time.Millisecond)
		assertlite.IsErrorOf(t, rec.Err(), failure)
	})
}

// TestAsyncJustObs tests AsyncJust[T any], which asynchronously emits a
// single value and completes.
//
// We want to assert:
//  1. Exactly one value is emitted
//  2. Completion occurs asynchronously after the value
//  3. Subscription returns immediately
func TestAsyncJustObs(t *testing.T) {
	t.Run("emits one value asynchronously and completes", func(t *testing.T) {
		src := AsyncJust(42)

		rec, o := observers.NewRecorder[int](t)
		start := time.Now()
		_, _ = src.Sub(o)

		// immediately after subscription, no values yet
		assertlite.Equals(t, 0, len(rec.Values()))
		assertlite.False(t, rec.Completed())

		// wait a bit for async emission
		time.Sleep(5 * time.Millisecond)

		assertlite.EqualSlices(t, []int{42}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
		assertlite.True(t, time.Since(start) >= 0, "value and completion should happen asynchronously")
	})

	t.Run("subscription returns immediately", func(t *testing.T) {
		src := AsyncJust("hello")
		rec, o := observers.NewRecorder[string](t)

		start := time.Now()
		_, _ = src.Sub((*observers.Observer[string])(o))
		subEnd := time.Now()

		assertlite.True(t, subEnd.Sub(start) < 2*time.Millisecond, "subscription should return immediately")

		time.Sleep(5 * time.Millisecond)
		assertlite.EqualSlices(t, []string{"hello"}, rec.Values())
		assertlite.True(t, rec.Completed())
	})
}

// TestAsyncMulticastObs tests AsyncMulticast[T any], which multicasts
// values from a channel to all subscribers asynchronously.
//
// We want to assert:
//  1. Values emitted on the channel are received by all active subscribers
//  2. Completion occurs after the channel is closed
//  3. Subscribers added after some values are emitted only get future values
func TestAsyncMulticastObs(t *testing.T) {
	t.Run("broadcasts channel values to multiple subscribers", func(t *testing.T) {
		ch := make(chan int)
		src := AsyncMulticast(ch)

		recA, oA := observers.NewRecorder[int](t)
		recB, oB := observers.NewRecorder[int](t)

		_, _ = src.Sub(oA)
		_, _ = src.Sub(oB)

		go func() {
			ch <- 1
			ch <- 2
			ch <- 3
			close(ch)
		}()

		// wait for async propagation
		time.Sleep(10 * time.Millisecond)

		assertlite.EqualSlices(t, []int{1, 2, 3}, recA.Values())
		assertlite.EqualSlices(t, []int{1, 2, 3}, recB.Values())
		assertlite.True(t, recA.Completed())
		assertlite.True(t, recB.Completed())
		assertlite.NoError(t, recA.Err())
		assertlite.NoError(t, recB.Err())
	})

	t.Run("late subscribers only receive future values", func(t *testing.T) {
		ch := make(chan int)
		src := AsyncMulticast(ch)

		// Early subscriber
		recA, oA := observers.NewRecorder[int](t)
		_, _ = src.Sub(oA)

		// Producer
		go func() {
			ch <- 1
			time.Sleep(1 * time.Millisecond)
			ch <- 2
			time.Sleep(1 * time.Millisecond)
			ch <- 3
			close(ch)
		}()

		// Wait until A actually RECEIVES the first value
		recA.AwaitNext(50 * time.Millisecond)

		// Late subscriber
		recB, oB := observers.NewRecorder[int](t)
		_, _ = src.Sub(oB)

		// Wait for termination
		recA.AwaitCompleted(50 * time.Millisecond)
		recB.AwaitCompleted(50 * time.Millisecond)

		assertlite.EqualSlices(t, []int{1, 2, 3}, recA.Values())
		assertlite.EqualSlices(t, []int{2, 3}, recB.Values())
	})
}

// TestAsyncOf_Sync verifies the AsyncOf observable emits values asynchronously
// and completes correctly.
//
// We aim to assert:
//  1. Emits all provided values in order
//  2. Completes after emitting all values
//  3. Supports single and multiple subscribers
//  4. Async behavior — emissions do not block subscription
func TestAsyncOf(t *testing.T) {
	t.Run("emits all values in order and completes", func(t *testing.T) {
		src := AsyncOf(1, 2, 3)

		rec, o := observers.NewRecorder[int](t)
		_, _ = src.Sub(o)

		// wait for async propagation
		time.Sleep(5 * time.Millisecond)

		assertlite.EqualSlices(t, []int{1, 2, 3}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("works with a single value", func(t *testing.T) {
		src := AsyncOf(42)

		rec, o := observers.NewRecorder[int](t)
		_, _ = src.Sub(o)

		time.Sleep(5 * time.Millisecond)

		assertlite.EqualSlices(t, []int{42}, rec.Values())
		assertlite.True(t, rec.Completed())
	})

	t.Run("works with zero values", func(t *testing.T) {
		src := AsyncOf[int]()

		rec, o := observers.NewRecorder[int](t)
		_, _ = src.Sub(o)

		time.Sleep(2 * time.Millisecond)

		assertlite.Equals(t, 0, len(rec.Values()))
		assertlite.True(t, rec.Completed())
	})

	t.Run("multiple subscribers receive independent async emissions", func(t *testing.T) {
		src := AsyncOf(10, 20, 30)

		recA, oA := observers.NewRecorder[int](t)
		recB, oB := observers.NewRecorder[int](t)

		_, _ = src.Sub(oA)
		_, _ = src.Sub(oB)

		time.Sleep(5 * time.Millisecond)

		assertlite.EqualSlices(t, []int{10, 20, 30}, recA.Values())
		assertlite.EqualSlices(t, []int{10, 20, 30}, recB.Values())
		assertlite.True(t, recA.Completed())
		assertlite.True(t, recB.Completed())
	})
}

// TestAsyncRange_Sync tests AsyncRange for asynchronous emission of numeric ranges.
//
// We aim to assert:
//  1. Emits the full range in correct order for positive delta
//  2. Emits the full range in correct order for negative delta
//  3. Panics on invalid delta (zero or mismatched signs)
//  4. Completes after emitting all values
//  5. Supports single and multiple subscribers
func TestAsyncRange(t *testing.T) {
	t.Run("emits positive range correctly", func(t *testing.T) {
		src := AsyncRange(1, 5, 1) // 1,2,3,4,5

		rec, o := observers.NewRecorder[int](t)
		_, _ = src.Sub(o)

		time.Sleep(5 * time.Millisecond)

		assertlite.EqualSlices(t, []int{1, 2, 3, 4, 5}, rec.Values())
		assertlite.True(t, rec.Completed())
	})

	t.Run("emits negative range correctly", func(t *testing.T) {
		src := AsyncRange(5, 1, -1) // 5,4,3,2,1

		rec, o := observers.NewRecorder[int](t)
		_, _ = src.Sub(o)

		time.Sleep(5 * time.Millisecond)

		assertlite.EqualSlices(t, []int{5, 4, 3, 2, 1}, rec.Values())
		assertlite.True(t, rec.Completed())
	})

	t.Run("panics on zero delta", func(t *testing.T) {
		assertlite.Panics(t, func() {
			_ = AsyncRange(1, 10, 0)
		})
	})

	t.Run("panics on positive delta with decreasing range", func(t *testing.T) {
		assertlite.Panics(t, func() {
			_ = AsyncRange(10, 1, 1)
		})
	})

	t.Run("panics on negative delta with increasing range", func(t *testing.T) {
		assertlite.Panics(t, func() {
			_ = AsyncRange(1, 10, -1)
		})
	})

	t.Run("supports multiple subscribers independently", func(t *testing.T) {
		src := AsyncRange(1, 3, 1) // 1,2,3

		recA, oA := observers.NewRecorder[int](t)
		recB, oB := observers.NewRecorder[int](t)

		_, _ = src.Sub(oA)
		_, _ = src.Sub(oB)

		time.Sleep(5 * time.Millisecond)

		assertlite.EqualSlices(t, []int{1, 2, 3}, recA.Values())
		assertlite.EqualSlices(t, []int{1, 2, 3}, recB.Values())
		assertlite.True(t, recA.Completed())
		assertlite.True(t, recB.Completed())
	})
}
