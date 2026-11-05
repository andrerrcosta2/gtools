// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package transform

import (
	"errors"
	"sync/atomic"
	"testing"
	"time"

	"github.com/andrerrcosta2/gtools/conc/testing/testschedulers"
	"github.com/andrerrcosta2/gtools/core/chronos"
	"github.com/andrerrcosta2/gtools/core/testlite/assertlite"
	"github.com/andrerrcosta2/gtools/gflux/obs"
	"github.com/andrerrcosta2/gtools/gflux/obs/internal/observers"
	"github.com/andrerrcosta2/gtools/gflux/obs/internal/subscr"
)

const (
	zero = 0 * time.Millisecond
)

// TestAudit_Sync
//
// Source emits synchronously.
//
// We should assert:
//  1. Values are not emitted immediately
//  2. Only the latest value in the window is emitted
//  3. Pending value is flushed on source completion
//  4. Source errors are forwarded immediately
//  5. No values are emitted after termination
//  6. Completion is forwarded exactly once
//  7. Audit never emits more than one value per window
func TestAudit_Sync(t *testing.T) {
	t.Run("emits latest value on completion", func(t *testing.T) {
		src := obs.Of(1, 2, 3)

		out := Audit[int](50 * time.Millisecond)(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		rec.AwaitCompleted(time.Second)

		assertlite.EqualSlices(t, []int{3}, rec.Values())
	})

	t.Run("emits single value on completion", func(t *testing.T) {
		src := obs.Just(42)

		out := Audit[int](10 * time.Millisecond)(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		rec.AwaitCompleted(time.Second)

		assertlite.EqualSlices(t, []int{42}, rec.Values())
	})

	t.Run("emits nothing for empty source", func(t *testing.T) {
		src := obs.Empty[int]()

		out := Audit[int](100 * time.Millisecond)(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		rec.AwaitCompleted(time.Second)

		assertlite.EqualSlices(t, nil, rec.Values())
	})

	t.Run("propagates source error immediately", func(t *testing.T) {
		expErr := errors.New("error.")

		src := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			o.Next(1)
			o.Error(expErr)
			o.Next(2) // must be ignored
			return subscr.Noop{}, nil
		})

		out := Audit[int](time.Second)(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		assertlite.IsErrorOf(t, expErr, rec.AwaitError(time.Second))
		assertlite.EqualSlices(t, nil, rec.Values())
	})

	t.Run("forwards completion exactly once", func(t *testing.T) {
		var completes atomic.Int32

		src := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			o.Next(1)
			o.Complete()
			o.Complete() // ignored by source contract
			return subscr.Noop{}, nil
		})

		out := Audit[int](10 * time.Millisecond)(src)
		rec, o := observers.NewRecorderFor(t, observers.NewComplete[int](func() {
			completes.Add(1)
		}))

		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)
		rec.AwaitTermination(time.Second)
		assertlite.True(t, completes.Load() == 1, "expected complete once but got '%v'", completes.Load())
	})

	t.Run("emits no values after completion", func(t *testing.T) {
		src := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			o.Next(1)
			o.Complete()
			o.Next(2) // must be ignored
			return subscr.Noop{}, nil
		})

		out := Audit[int](10 * time.Millisecond)(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		rec.AwaitCompleted(time.Second)

		assertlite.EqualSlices(t, []int{1}, rec.Values())
	})
}

// TestAudit_Sync_ZeroDuration Source emits synchronously with zero audit duration.
//
// We should assert:
//  1. All values are eventually emitted
//  2. Values preserve source order
//  3. Completion is forwarded
//  4. No values are emitted after completion
//  5. Completion happens exactly once
func TestAudit_Sync_ZeroDuration(t *testing.T) {
	t.Run("emits all values preserving order", func(t *testing.T) {
		src := obs.Of(1, 2, 3)

		out := Audit[int](zero)(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		rec.AwaitCompleted(time.Second)

		assertlite.EqualSlices(t, []int{1, 2, 3}, rec.Values())
	})

	t.Run("propagates completion", func(t *testing.T) {
		var completed atomic.Bool

		src := obs.Of(10, 20)

		out := Audit[int](zero)(src)

		_, err := out.Sub(&obs.Observer[int]{
			Complete: func() {
				completed.Store(true)
			},
		})
		assertlite.NoError(t, err)

		time.Sleep(50 * time.Millisecond)
		assertlite.True(t, completed.Load())
	})

	t.Run("emits no values after completion", func(t *testing.T) {
		src := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			o.Next(1)
			o.Complete()
			o.Next(2) // must be ignored
			return subscr.Noop{}, nil
		})

		out := Audit[int](zero)(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		rec.AwaitCompleted(time.Second)

		assertlite.EqualSlices(t, []int{1}, rec.Values())
	})

	t.Run("forwards source error immediately", func(t *testing.T) {
		expErr := errors.New("audit zero error")

		src := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			o.Next(1)
			o.Error(expErr)
			o.Next(2) // must be ignored
			return subscr.Noop{}, nil
		})

		out := Audit[int](zero)(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		assertlite.IsErrorOf(t, expErr, rec.AwaitError(time.Second))
		assertlite.EqualSlices(t, []int{1}, rec.Values())
	})

	t.Run("forwards completion exactly once", func(t *testing.T) {
		var completes atomic.Int32

		src := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			o.Next(1)
			o.Complete()
			o.Complete() // ignored by source contract
			return subscr.Noop{}, nil
		})

		out := Audit[int](zero)(src)

		_, err := out.Sub(&obs.Observer[int]{
			Complete: func() {
				completes.Add(1)
			},
		})
		assertlite.NoError(t, err)

		time.Sleep(50 * time.Millisecond)
		assertlite.True(t, completes.Load() == 1)
	})
}

// TestAudit_Async
//
// Source emits asynchronously over time.
//
// We should assert:
//  1. Values are NOT emitted immediately
//  2. Only the most recent value in each time window is emitted
//  3. Windows are fixed-duration and do NOT reset on each emission
//  4. Multiple windows emit independently
//  5. If the source completes while a window is active:
//     - the last observed value is emitted
//     - completion is then forwarded
//  6. Source errors are propagated immediately
//  7. Active timers are cancelled on error
//  8. No values are emitted after termination
//  9. Completion is forwarded exactly once
func TestAudit_Async(t *testing.T) {
	t.Run("emits latest value at end of each window", func(t *testing.T) {
		src := obs.Async(func(o *obs.Observer[int]) {
			o.Next(1)
			time.Sleep(20 * time.Millisecond)
			o.Next(2)
			time.Sleep(20 * time.Millisecond)
			o.Next(3)
			time.Sleep(80 * time.Millisecond) // window closes here
			o.Complete()
		})

		out := Audit[int](50 * time.Millisecond)(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		rec.AwaitCompleted(time.Second)

		// All values fall into the same window → only latest emitted
		assertlite.EqualSlices(t, []int{3}, rec.Values())
	})

	t.Run("emits multiple windows independently", func(t *testing.T) {
		src := obs.Async(func(o *obs.Observer[int]) {
			o.Next(1)
			time.Sleep(30 * time.Millisecond)
			o.Next(2)

			time.Sleep(80 * time.Millisecond) // first window closes

			o.Next(3)
			time.Sleep(30 * time.Millisecond)
			o.Next(4)

			time.Sleep(80 * time.Millisecond) // second window closes
			o.Complete()
		})

		out := Audit[int](50 * time.Millisecond)(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		rec.AwaitCompleted(time.Second)

		assertlite.EqualSlices(t, []int{2, 4}, rec.Values())
	})

	t.Run("flushes pending value on completion", func(t *testing.T) {
		src := obs.Async(func(o *obs.Observer[int]) {
			o.Next(10)
			time.Sleep(20 * time.Millisecond)
			o.Next(20)
			o.Complete() // window still active
		})

		out := Audit[int](100 * time.Millisecond)(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		rec.AwaitCompleted(time.Second)

		assertlite.EqualSlices(t, []int{20}, rec.Values())
	})

	t.Run("propagates source error immediately", func(t *testing.T) {
		expErr := errors.New("audit async error")

		src := obs.Async(func(o *obs.Observer[int]) {
			o.Next(1)
			time.Sleep(20 * time.Millisecond)
			o.Error(expErr)
			o.Next(2) // must be ignored
		})

		out := Audit[int](50 * time.Millisecond)(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		assertlite.IsErrorOf(t, expErr, rec.AwaitError(time.Second))
		assertlite.EqualSlices(t, nil, rec.Values())
	})

	t.Run("emits no values after completion", func(t *testing.T) {
		src := obs.Async(func(o *obs.Observer[int]) {
			o.Next(1)
			time.Sleep(20 * time.Millisecond)
			o.Complete()
			o.Next(2) // must be ignored
		})

		out := Audit[int](50 * time.Millisecond)(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		rec.AwaitCompleted(time.Second)

		assertlite.EqualSlices(t, []int{1}, rec.Values())
	})

	t.Run("forwards completion exactly once", func(t *testing.T) {
		var completes atomic.Int32

		src := obs.Async(func(o *obs.Observer[int]) {
			o.Next(1)
			o.Complete()
			o.Complete() // ignored by source contract
		})

		out := Audit[int](50 * time.Millisecond)(src)

		_, err := out.Sub(&obs.Observer[int]{
			Complete: func() {
				completes.Add(1)
			},
		})
		assertlite.NoError(t, err)

		time.Sleep(200 * time.Millisecond)
		assertlite.True(t, completes.Load() == 1)
	})
}

var testScheduler = testschedulers.New(chronos.VirtualClock())

// TestDebounce_Sync
//
// Source emits synchronously.
//
// We should assert:
//  1. Values are NOT emitted immediately
//  2. Intermediate values within the debounce duration are suppressed
//  3. Only the last value in a burst is emitted
//  4. If the source completes with a pending value, it is emitted immediately
//  5. Source errors are propagated immediately
//  6. Pending debounced emissions are cancelled on error
//  7. No values are emitted after termination
//  8. Completion is forwarded exactly once
func TestDebounce_Sync(t *testing.T) {
	t.Run("suppresses intermediate synchronous values", func(t *testing.T) {
		src := obs.Of(1, 2, 3)

		out := Debounce[int](50*time.Millisecond, testScheduler)(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		rec.AwaitCompleted(time.Second)

		// All values are within the same burst → only last survives
		assertlite.EqualSlices(t, []int{3}, rec.Values())
	})

	t.Run("emits pending value immediately on completion", func(t *testing.T) {
		src := obs.Of(10, 20, 30)

		out := Debounce[int](100*time.Millisecond, testScheduler)(src)

		rec, o := observers.NewRecorder[int](t)
		start := time.Now()

		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		rec.AwaitCompleted(time.Second)
		elapsed := time.Since(start)

		// Must not wait for debounce duration after completion
		assertlite.True(t, elapsed < 100*time.Millisecond)
		assertlite.EqualSlices(t, []int{30}, rec.Values())
	})

	t.Run("propagates source error immediately", func(t *testing.T) {
		expErr := errors.New("debounce error")

		src := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			o.Next(1)
			o.Next(2)
			o.Error(expErr)
			o.Next(3) // must be ignored
			return subscr.Noop{}, nil
		})

		out := Debounce[int](50*time.Millisecond, testScheduler)(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		assertlite.IsErrorOf(t, expErr, rec.AwaitError(time.Second))
		assertlite.IsNil(t, rec.Values())
	})

	t.Run("does not emit after error", func(t *testing.T) {
		expErr := errors.New("error")

		src := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			o.Next(1)
			o.Error(expErr)
			time.Sleep(100 * time.Millisecond)
			return subscr.Noop{}, nil
		})

		out := Debounce[int](20*time.Millisecond, testScheduler)(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		assertlite.IsErrorOf(t, expErr, rec.AwaitError(time.Second))
		assertlite.IsNil(t, rec.Values())
	})

	t.Run("forwards completion exactly once", func(t *testing.T) {
		var completes atomic.Int32

		src := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			o.Next(1)
			o.Complete()
			o.Complete() // ignored by source contract
			return subscr.Noop{}, nil
		})

		out := Debounce[int](10*time.Millisecond, testScheduler)(src)

		_, err := out.Sub(&obs.Observer[int]{
			Complete: func() {
				completes.Add(1)
			},
		})
		assertlite.NoError(t, err)

		time.Sleep(100 * time.Millisecond)
		assertlite.True(t, completes.Load() == 1)
	})
}

// TestDebounce_Sync_ZeroDuration
//
// Source emits synchronously and debounce duration is zero.
//
// We should assert:
//  1. Values are not emitted immediately during synchronous emission
//  2. Intermediate values are suppressed
//  3. Only the last value is emitted
//  4. The pending value is flushed on completion
//  5. Source errors are propagated immediately
//  6. No values are emitted after termination
//  7. Completion is forwarded exactly once
func TestDebounce_Sync_ZeroDuration(t *testing.T) {
	t.Run("emits only the last synchronous value", func(t *testing.T) {
		src := obs.Of(1, 2, 3)

		out := Debounce[int](zero, testScheduler)(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		rec.AwaitCompleted(time.Second)

		assertlite.EqualSlices(t, []int{3}, rec.Values())
	})

	t.Run("does not emit during synchronous execution", func(t *testing.T) {
		var emitted atomic.Bool

		src := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			o.Next(1)
			o.Next(2)
			o.Next(3)

			assertlite.False(t, emitted.Load(), "must not emit synchronously")
			o.Complete()
			return subscr.Noop{}, nil
		})

		out := Debounce[int](zero, testScheduler)(src)

		_, err := out.Sub(&obs.Observer[int]{
			Next: func(v int) {
				emitted.Store(true)
			},
		})
		assertlite.NoError(t, err)
	})

	t.Run("flushes last value on completion", func(t *testing.T) {
		src := obs.Of(10, 20, 30)

		out := Debounce[int](zero, testScheduler)(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		rec.AwaitCompleted(time.Second)

		assertlite.EqualSlices(t, []int{30}, rec.Values())
	})

	t.Run("propagates source error immediately", func(t *testing.T) {
		expErr := errors.New("zero debounce error")

		src := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			o.Next(1)
			o.Error(expErr)
			o.Next(2)
			return subscr.Noop{}, nil
		})

		out := Debounce[int](zero, testScheduler)(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		assertlite.IsErrorOf(t, expErr, rec.AwaitError(time.Second))
		assertlite.IsNil(t, rec.Values())
	})

	t.Run("forwards completion exactly once", func(t *testing.T) {
		var completes atomic.Int32

		src := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			o.Next(1)
			o.Complete()
			o.Complete()
			return subscr.Noop{}, nil
		})

		out := Debounce[int](zero, testScheduler)(src)

		_, err := out.Sub(&obs.Observer[int]{
			Complete: func() {
				completes.Add(1)
			},
		})
		assertlite.NoError(t, err)

		time.Sleep(50 * time.Millisecond)
		assertlite.True(t, completes.Load() == 1)
	})

	// this is the worse case for simultaneous commands inside the buffer
	t.Run("bounded under heavy sync emission", func(t *testing.T) {
		src := obs.Range(0, 1_000_000, 1)

		out := Debounce[int](zero, testScheduler)(src)

		rec, o := observers.NewRecorder[int](t)

		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		rec.AwaitCompleted(time.Second)

		assertlite.EqualSlices(t, []int{1_000_000}, rec.Values())
	})

}

// TestDebounce_Async
//
// Source emits asynchronously over time.
//
// We should assert:
//  1. Values are emitted only after the source becomes silent
//  2. Intermediate values reset the debounce timer
//  3. Only the last value in each burst is emitted
//  4. Multiple bursts emit independently
//  5. If the source completes with a pending value, it is emitted immediately
//  6. Source errors are propagated immediately
//  7. Pending debounced emissions are cancelled on error
//  8. No values are emitted after termination
//  9. Completion is forwarded exactly once
func TestDebounce_Async(t *testing.T) {
	t.Run("emits value only after silence", func(t *testing.T) {
		src := obs.Async(func(o *obs.Observer[int]) {
			o.Next(1)
			time.Sleep(20 * time.Millisecond)
			o.Next(2)
			time.Sleep(20 * time.Millisecond)
			o.Next(3)

			// silence longer than debounce
			time.Sleep(80 * time.Millisecond)
			o.Complete()
		})

		out := Debounce[int](50*time.Millisecond, testScheduler)(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		rec.AwaitCompleted(time.Second)

		assertlite.EqualSlices(t, []int{3}, rec.Values())
	})

	t.Run("resets debounce timer on each emission", func(t *testing.T) {
		src := obs.Async(func(o *obs.Observer[int]) {
			o.Next(1)
			time.Sleep(30 * time.Millisecond)
			o.Next(2)
			time.Sleep(30 * time.Millisecond)
			o.Next(3)

			// only now silence starts
			time.Sleep(80 * time.Millisecond)
			o.Complete()
		})

		out := Debounce[int](50*time.Millisecond, testScheduler)(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		rec.AwaitCompleted(time.Second)

		assertlite.EqualSlices(t, []int{3}, rec.Values())
	})

	t.Run("emits values from multiple bursts", func(t *testing.T) {
		src := obs.Async(func(o *obs.Observer[int]) {
			// first burst
			o.Next(1)
			o.Next(2)
			time.Sleep(80 * time.Millisecond)

			// second burst
			o.Next(3)
			o.Next(4)
			time.Sleep(80 * time.Millisecond)

			o.Complete()
		})

		out := Debounce[int](50*time.Millisecond, testScheduler)(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		rec.AwaitCompleted(time.Second)

		assertlite.EqualSlices(t, []int{2, 4}, rec.Values())
	})

	t.Run("flushes pending value immediately on completion", func(t *testing.T) {
		src := obs.Async(func(o *obs.Observer[int]) {
			o.Next(10)
			time.Sleep(20 * time.Millisecond)
			o.Next(20)
			o.Complete() // silence not reached
		})

		out := Debounce[int](100*time.Millisecond, testScheduler)(src)

		rec, o := observers.NewRecorder[int](t)
		start := time.Now()

		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		rec.AwaitCompleted(time.Second)
		elapsed := time.Since(start)

		// Must not wait for debounce duration after completion
		assertlite.True(t, elapsed < 100*time.Millisecond)
		assertlite.EqualSlices(t, []int{20}, rec.Values())
	})

	t.Run("propagates source error immediately", func(t *testing.T) {
		expErr := errors.New("debounce async error")

		src := obs.Async(func(o *obs.Observer[int]) {
			o.Next(1)
			time.Sleep(20 * time.Millisecond)
			o.Error(expErr)
			o.Next(2) // must be ignored
		})

		out := Debounce[int](50*time.Millisecond, testScheduler)(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		assertlite.IsErrorOf(t, expErr, rec.AwaitError(time.Second))
		assertlite.IsNil(t, rec.Values())
	})

	t.Run("does not emit after termination", func(t *testing.T) {
		src := obs.Async(func(o *obs.Observer[int]) {
			o.Next(1)
			o.Complete()
			time.Sleep(100 * time.Millisecond)
		})

		out := Debounce[int](20*time.Millisecond, testScheduler)(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		rec.AwaitCompleted(time.Second)

		assertlite.EqualSlices(t, []int{1}, rec.Values())
	})

	t.Run("forwards completion exactly once", func(t *testing.T) {
		var completes atomic.Int32

		src := obs.Async(func(o *obs.Observer[int]) {
			o.Next(1)
			o.Complete()
			o.Complete()
		})

		out := Debounce[int](20*time.Millisecond, testScheduler)(src)

		_, err := out.Sub(&obs.Observer[int]{
			Complete: func() {
				completes.Add(1)
			},
		})
		assertlite.NoError(t, err)

		time.Sleep(200 * time.Millisecond)
		assertlite.True(t, completes.Load() == 1)
	})
}

// TestDelay_Sync
//
// Source emits synchronously.
//
// We should assert:
//  1. Values are NOT emitted immediately (unless delay = 0)
//  2. All values are emitted (none are suppressed)
//  3. Order is preserved
//  4. Completion is delayed until after the last value
//  5. Source errors are propagated immediately
//  6. Pending delayed values are discarded on error
//  7. No values are emitted after termination
//  8. Completion is forwarded exactly once
func TestDelay_Sync(t *testing.T) {

	t.Run("delays all synchronous values", func(t *testing.T) {
		src := obs.Of(1, 2, 3)

		out := Delay[int](50*time.Millisecond, 5)(src)

		rec, o := observers.NewRecorder[int](t)
		start := time.Now()

		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		rec.AwaitCompleted(time.Second)
		elapsed := time.Since(start)

		// Must wait at least the delay duration
		assertlite.True(t, elapsed >= 50*time.Millisecond)
		assertlite.EqualSlices(t, []int{1, 2, 3}, rec.Values())
	})

	t.Run("preserves order of delayed values", func(t *testing.T) {
		src := obs.Of(10, 20, 30, 40)

		out := Delay[int](30*time.Millisecond, 5)(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		rec.AwaitCompleted(time.Second)

		assertlite.EqualSlices(t, []int{10, 20, 30, 40}, rec.Values())
	})

	t.Run("delays completion until after last value", func(t *testing.T) {
		src := obs.Of(1, 2, 3)

		out := Delay[int](40*time.Millisecond, 5)(src)

		rec, o := observers.NewRecorder[int](t)
		start := time.Now()

		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		rec.AwaitCompleted(time.Second)
		elapsed := time.Since(start)

		// Completion must not be immediate
		assertlite.True(t, elapsed >= 40*time.Millisecond)
		assertlite.EqualSlices(t, []int{1, 2, 3}, rec.Values())
	})

	t.Run("propagates source error immediately", func(t *testing.T) {
		expErr := errors.New("delay error")

		src := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			o.Next(1)
			o.Next(2)
			o.Error(expErr)
			o.Next(3) // must be ignored
			return subscr.Noop{}, nil
		})

		out := Delay[int](100*time.Millisecond, 5)(src)

		rec, o := observers.NewRecorder[int](t)
		start := time.Now()

		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		errGot := rec.AwaitError(time.Second)
		elapsed := time.Since(start)

		// Error must not be delayed
		assertlite.True(t, elapsed < 100*time.Millisecond)
		assertlite.IsErrorOf(t, expErr, errGot)
		assertlite.IsNil(t, rec.Values())
	})

	t.Run("does not emit delayed values after error", func(t *testing.T) {
		expErr := errors.New("error")

		src := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			o.Next(1)
			o.Error(expErr)
			time.Sleep(100 * time.Millisecond)
			return subscr.Noop{}, nil
		})

		out := Delay[int](50*time.Millisecond, 5)(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		assertlite.IsErrorOf(t, expErr, rec.AwaitError(time.Second))
		assertlite.IsNil(t, rec.Values())
	})

	t.Run("forwards completion exactly once", func(t *testing.T) {
		var completes atomic.Int32

		src := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			o.Next(1)
			o.Complete()
			o.Complete() // ignored by source contract
			return subscr.Noop{}, nil
		})

		out := Delay[int](10*time.Millisecond, 5)(src)

		_, err := out.Sub(&obs.Observer[int]{
			Complete: func() {
				completes.Add(1)
			},
		})
		assertlite.NoError(t, err)

		time.Sleep(100 * time.Millisecond)
		assertlite.True(t, completes.Load() == 1)
	})
}

// TestDelay_Sync_ZeroDuration
//
// Source emits synchronously and delay duration is zero.
//
// We should assert:
//  1. Values are emitted immediately
//  2. No values are suppressed
//  3. Order is preserved
//  4. Completion is forwarded immediately
//  5. Errors are propagated immediately
func TestDelay_Sync_ZeroDuration(t *testing.T) {

	t.Run("emits values immediately", func(t *testing.T) {
		src := obs.Of(1, 2, 3)

		out := Delay[int](zero, 5)(src)

		rec, o := observers.NewRecorder[int](t)
		start := time.Now()

		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		rec.AwaitCompleted(time.Second)
		elapsed := time.Since(start)

		assertlite.True(t, elapsed < 10*time.Millisecond)
		assertlite.EqualSlices(t, []int{1, 2, 3}, rec.Values())
	})

	t.Run("forwards completion immediately", func(t *testing.T) {
		src := obs.Of(42)

		out := Delay[int](zero, 5)(src)

		rec, o := observers.NewRecorder[int](t)
		start := time.Now()

		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		rec.AwaitCompleted(time.Second)
		elapsed := time.Since(start)

		assertlite.True(t, elapsed < 10*time.Millisecond)
		assertlite.EqualSlices(t, []int{42}, rec.Values())
	})

	t.Run("propagates error immediately", func(t *testing.T) {
		expErr := errors.New("delay zero error")

		src := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			o.Error(expErr)
			return subscr.Noop{}, nil
		})

		out := Delay[int](zero, 5)(src)

		rec, o := observers.NewRecorder[int](t)
		start := time.Now()

		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		errGot := rec.AwaitError(time.Second)
		elapsed := time.Since(start)

		assertlite.True(t, elapsed < 10*time.Millisecond)
		assertlite.IsErrorOf(t, expErr, errGot)
		assertlite.IsNil(t, rec.Values())
	})
}

// TestDelay_Async
//
// Source emits asynchronously.
//
// We should assert:
//  1. Values are delayed by at least the configured duration
//  2. Relative ordering of values is preserved
//  3. Completion is delayed until all pending values are emitted
//  4. Errors are propagated immediately
//  5. Pending delayed values are discarded on error
//  6. No values are emitted after termination
//  7. Completion is forwarded exactly once
func TestDelay_Async(t *testing.T) {

	t.Run("delays each value by at least the configured duration", func(t *testing.T) {
		src := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			go func() {
				o.Next(1)
				o.Next(2)
				o.Next(3)
				o.Complete()
			}()
			return subscr.Noop{}, nil
		})

		delay := 50 * time.Millisecond
		out := Delay[int](delay, 5)(src)

		rec, o := observers.NewRecorder[int](t)
		start := time.Now()

		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		rec.AwaitCompleted(time.Second)
		elapsed := time.Since(start)

		// Completion must be delayed until after the last value
		assertlite.True(t, elapsed >= delay)
		assertlite.EqualSlices(t, []int{1, 2, 3}, rec.Values())
	})

	t.Run("preserves ordering of asynchronously emitted values", func(t *testing.T) {
		src := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			go func() {
				o.Next(10)
				time.Sleep(10 * time.Millisecond)
				o.Next(20)
				time.Sleep(10 * time.Millisecond)
				o.Next(30)
				o.Complete()
			}()
			return subscr.Noop{}, nil
		})

		out := Delay[int](30*time.Millisecond, 5)(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		rec.AwaitCompleted(time.Second)
		assertlite.EqualSlices(t, []int{10, 20, 30}, rec.Values())
	})

	t.Run("propagates error immediately and cancels pending values", func(t *testing.T) {
		expErr := errors.New("delay async error")

		src := obs.Async[int](func(o *obs.Observer[int]) {
			o.Next(1)
			time.Sleep(10 * time.Millisecond)
			o.Error(expErr)
			o.Next(2) // must not be delivered
		})

		out := Delay[int](50*time.Millisecond, 5)(src)

		rec, o := observers.NewRecorder[int](t)
		start := time.Now()

		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		errGot := rec.AwaitError(time.Second)
		elapsed := time.Since(start)

		// Error must not wait for delay
		assertlite.True(t, elapsed < 50*time.Millisecond)
		assertlite.IsErrorOf(t, expErr, errGot)
		assertlite.IsNil(t, rec.Values())
	})

	t.Run("does not emit values after error", func(t *testing.T) {
		expErr := errors.New("error")

		src := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			go func() {
				o.Next(1)
				o.Error(expErr)
				time.Sleep(100 * time.Millisecond)
				o.Next(2) // must be ignored
			}()
			return subscr.Noop{}, nil
		})

		out := Delay[int](20*time.Millisecond, 5)(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		assertlite.IsErrorOf(t, expErr, rec.AwaitError(time.Second))
		assertlite.IsNil(t, rec.Values())
	})

	t.Run("forwards completion exactly once after delayed emissions", func(t *testing.T) {
		var completes atomic.Int32

		src := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			go func() {
				o.Next(1)
				o.Complete()
				o.Complete() // ignored by source contract
			}()
			return subscr.Noop{}, nil
		})

		out := Delay[int](20*time.Millisecond, 5)(src)

		_, err := out.Sub(&obs.Observer[int]{
			Complete: func() {
				completes.Add(1)
			},
		})
		assertlite.NoError(t, err)

		time.Sleep(100 * time.Millisecond)
		assertlite.True(t, completes.Load() == 1)
	})
}

// TestThrottle_Sync
//
// Source emits synchronously.
//
// We should assert:
//  1. The first value is emitted immediately
//  2. Subsequent values within the throttle window are dropped
//  3. No values are delayed
//  4. No buffered values are emitted on completion
//  5. Source errors are propagated immediately
//  6. No values are emitted after termination
//  7. Completion is forwarded exactly once
func TestThrottle_Sync(t *testing.T) {

	t.Run("emits only the first value in a synchronous burst", func(t *testing.T) {
		src := obs.Of(1, 2, 3, 4, 5)

		out := Throttle[int](50 * time.Millisecond)(src)

		rec, o := observers.NewRecorder[int](t)
		start := time.Now()

		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		rec.AwaitCompleted(time.Second)
		elapsed := time.Since(start)

		// First value must be emitted immediately
		assertlite.True(t, elapsed < 50*time.Millisecond)
		assertlite.EqualSlices(t, []int{1}, rec.Values())
	})

	t.Run("does not emit suppressed values on completion", func(t *testing.T) {
		src := obs.Of(10, 20, 30)

		out := Throttle[int](100 * time.Millisecond)(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		rec.AwaitCompleted(time.Second)

		// Only the first value is emitted; others are dropped forever
		assertlite.EqualSlices(t, []int{10}, rec.Values())
	})

	t.Run("propagates source error immediately", func(t *testing.T) {
		expErr := errors.New("throttle error")

		src := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			o.Next(1)
			o.Next(2) // dropped
			o.Error(expErr)
			o.Next(3) // must be ignored
			return subscr.Noop{}, nil
		})

		out := Throttle[int](50 * time.Millisecond)(src)

		rec, o := observers.NewRecorder[int](t)
		start := time.Now()

		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		errGot := rec.AwaitError(time.Second)
		elapsed := time.Since(start)

		// Error must not wait for throttle window
		assertlite.True(t, elapsed < 50*time.Millisecond)
		assertlite.IsErrorOf(t, expErr, errGot)
		assertlite.EqualSlices(t, []int{1}, rec.Values())
	})

	t.Run("does not emit values after error", func(t *testing.T) {
		expErr := errors.New("error.")

		src := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			o.Next(1)
			o.Error(expErr)
			o.Next(2) // must be ignored
			return subscr.Noop{}, nil
		})

		out := Throttle[int](20 * time.Millisecond)(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		assertlite.IsErrorOf(t, expErr, rec.AwaitError(time.Second))
		assertlite.EqualSlices(t, []int{1}, rec.Values())
	})

	t.Run("forwards completion exactly once", func(t *testing.T) {
		var completes atomic.Int32

		src := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			o.Next(1)
			o.Complete()
			o.Complete() // ignored by source contract
			return subscr.Noop{}, nil
		})

		out := Throttle[int](50 * time.Millisecond)(src)

		_, err := out.Sub(&obs.Observer[int]{
			Complete: func() {
				completes.Add(1)
			},
		})
		assertlite.NoError(t, err)

		time.Sleep(100 * time.Millisecond)
		assertlite.True(t, completes.Load() == 1)
	})
}

// TestThrottle_Async
//
// Source emits asynchronously.
//
// We should assert:
//  1. The first value is emitted immediately
//  2. Subsequent values within the throttle window are dropped
//  3. A new value is emitted when the throttle window reopens
//  4. Values are never delayed
//  5. Completion is forwarded immediately
//  6. Errors are propagated immediately
//  7. No values are emitted after termination
func TestThrottle_Async(t *testing.T) {
	t.Run("emits first value in each throttle window", func(t *testing.T) {
		src := obs.AsyncIntervalOf(20*time.Millisecond, 1, 2, 3, 4, 5, 6)

		out := Throttle[int](50 * time.Millisecond)(src)

		rec, o := observers.NewRecorder[int](t)
		start := time.Now()

		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		rec.AwaitCompleted(time.Second)
		elapsed := time.Since(start)

		// Expected emissions:
		// t=0ms   -> 1 emitted
		// t=20ms  -> 2 dropped
		// t=40ms  -> 3 dropped
		// t≈50ms  -> window opens
		// t=60ms  -> 4 emitted
		// t=80ms  -> 5 dropped
		// t≈100ms -> window opens
		assertlite.EqualSlices(t, []int{1, 4}, rec.Values())

		// Completion must not wait for throttle window
		assertlite.True(t, elapsed < 200*time.Millisecond)
	})

	t.Run("propagates source error immediately", func(t *testing.T) {
		expErr := errors.New("async throttle error")

		src := obs.Async(func(o *obs.Observer[int]) {
			o.Next(1)
			time.Sleep(10 * time.Millisecond)
			o.Next(2) // dropped
			time.Sleep(10 * time.Millisecond)
			o.Error(expErr)
			o.Next(3) // must be ignored
		})

		out := Throttle[int](50 * time.Millisecond)(src)

		rec, o := observers.NewRecorder[int](t)
		start := time.Now()

		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		errGot := rec.AwaitError(time.Second)
		elapsed := time.Since(start)

		// Error must not wait for throttle window
		assertlite.True(t, elapsed < 50*time.Millisecond)
		assertlite.IsErrorOf(t, expErr, errGot)

		// Only the first value should be emitted
		assertlite.EqualSlices(t, []int{1}, rec.Values())
	})

	t.Run("does not emit values after completion", func(t *testing.T) {
		src := obs.Async(func(o *obs.Observer[int]) {
			o.Next(1)
			time.Sleep(10 * time.Millisecond)
			o.Next(2) // dropped
			o.Complete()
			time.Sleep(50 * time.Millisecond)
			o.Next(3) // must be ignored
		})

		out := Throttle[int](30 * time.Millisecond)(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		rec.AwaitCompleted(time.Second)

		assertlite.EqualSlices(t, []int{1}, rec.Values())
	})
}

// TestTimeout_Sync
//
// Source emits synchronously.
//
// We should assert:
//  1. Values are forwarded immediately
//  2. Timeout does NOT trigger if source emits synchronously
//  3. Completion is forwarded normally
//  4. Source errors are propagated immediately
//  5. No timeout occurs after completion
//  6. No values are emitted after termination
func TestTimeout_Sync(t *testing.T) {

	t.Run("forwards synchronous values without timeout", func(t *testing.T) {
		src := obs.Of(1, 2, 3)

		out := Timeout[int](100*time.Millisecond, testScheduler)(src)

		rec, o := observers.NewRecorder[int](t)
		start := time.Now()

		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		rec.AwaitCompleted(time.Second)
		elapsed := time.Since(start)

		// No waiting: everything is synchronous
		assertlite.True(t, elapsed < 50*time.Millisecond)
		assertlite.EqualSlices(t, []int{1, 2, 3}, rec.Values())
	})

	t.Run("forwards completion normally", func(t *testing.T) {
		src := obs.Of(42)

		out := Timeout[int](50*time.Millisecond, testScheduler)(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		rec.AwaitCompleted(time.Second)

		assertlite.EqualSlices(t, []int{42}, rec.Values())
	})

	t.Run("propagates source error immediately", func(t *testing.T) {
		expErr := errors.New("sync error")

		src := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			o.Next(1)
			o.Error(expErr)
			o.Next(2) // must be ignored
			return subscr.Noop{}, nil
		})

		out := Timeout[int](100*time.Millisecond, testScheduler)(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		errGot := rec.AwaitError(time.Second)

		assertlite.IsErrorOf(t, expErr, errGot)
		assertlite.EqualSlices(t, []int{1}, rec.Values())
	})

	t.Run("does not emit timeout after completion", func(t *testing.T) {
		src := obs.Of(1, 2, 3)

		out := Timeout[int](10*time.Millisecond, testScheduler)(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		// Wait longer than timeout duration to ensure no late timeout
		time.Sleep(20 * time.Millisecond)

		assertlite.EqualSlices(t, []int{1, 2, 3}, rec.Values())
		assertlite.IsNil(t, rec.Err())
	})

	t.Run("forwards completion exactly once", func(t *testing.T) {
		var completes atomic.Int32

		src := obs.Create[int](func(o *obs.Observer[int]) (obs.Subscription, error) {
			o.Next(1)
			o.Complete()
			o.Complete() // ignored by source contract
			return subscr.Noop{}, nil
		})

		out := Timeout[int](20*time.Millisecond, testScheduler)(src)

		_, err := out.Sub(&obs.Observer[int]{
			Complete: func() {
				completes.Add(1)
			},
		})
		assertlite.NoError(t, err)

		time.Sleep(50 * time.Millisecond)
		assertlite.True(t, completes.Load() == 1)
	})
}

// TestTimeout_Async
//
// Source emits asynchronously.
//
// We should assert:
//  1. Values are forwarded if they arrive before the timeout
//  2. Timeout triggers if no value arrives within the duration
//  3. Timeout resets after each emission
//  4. Timeout error is terminal and cancels the source
//  5. Completion before timeout is forwarded normally
//  6. No values are emitted after a timeout error
func TestTimeout_Async(t *testing.T) {

	t.Run("forwards values that arrive before timeout", func(t *testing.T) {
		src := obs.Async(func(o *obs.Observer[int]) {
			o.Next(1)
			time.Sleep(30 * time.Millisecond)
			o.Next(2)
			time.Sleep(30 * time.Millisecond)
			o.Next(3)
			o.Complete()
		})

		out := Timeout[int](100*time.Millisecond, testScheduler)(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		rec.AwaitCompleted(time.Second)

		assertlite.EqualSlices(t, []int{1, 2, 3}, rec.Values())
		assertlite.IsNil(t, rec.Err())
	})

	t.Run("triggers timeout if source stalls", func(t *testing.T) {
		src := obs.Async(func(o *obs.Observer[int]) {
			o.Next(1)
			time.Sleep(30 * time.Millisecond)
			o.Next(2)
			// stall long enough to trigger timeout
			time.Sleep(200 * time.Millisecond)
			o.Next(3) // must be ignored
		})

		out := Timeout[int](50*time.Millisecond, testScheduler)(src)

		rec, o := observers.NewRecorder[int](t)
		start := time.Now()

		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		errGot := rec.AwaitError(time.Second)
		elapsed := time.Since(start)

		// Timeout should fire ~50ms after last value
		assertlite.True(t, elapsed >= 50*time.Millisecond)
		assertlite.True(t, elapsed < 150*time.Millisecond)

		var toErr obs.TimeoutError
		assertlite.True(t, errors.As(errGot, &toErr))

		assertlite.EqualSlices(t, []int{1, 2}, rec.Values())
	})

	t.Run("resets timeout after each emission", func(t *testing.T) {
		src := obs.Async(func(o *obs.Observer[int]) {
			o.Next(1)
			time.Sleep(40 * time.Millisecond)
			o.Next(2)
			time.Sleep(40 * time.Millisecond)
			o.Next(3)
			time.Sleep(40 * time.Millisecond)
			o.Complete()
		})

		out := Timeout[int](60*time.Millisecond, testScheduler)(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		rec.AwaitCompleted(time.Second)

		// Range value arrives before timeout window expires
		assertlite.EqualSlices(t, []int{1, 2, 3}, rec.Values())
		assertlite.IsNil(t, rec.Err())
	})

	t.Run("completion before timeout cancels timer", func(t *testing.T) {
		src := obs.Async(func(o *obs.Observer[int]) {
			o.Next(1)
			time.Sleep(20 * time.Millisecond)
			o.Complete()
		})

		out := Timeout[int](100*time.Millisecond, testScheduler)(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		rec.AwaitCompleted(time.Second)

		// Wait longer than timeout to ensure no late timeout
		time.Sleep(150 * time.Millisecond)

		assertlite.EqualSlices(t, []int{1}, rec.Values())
		assertlite.IsNil(t, rec.Err())
	})

	t.Run("does not emit values after timeout error", func(t *testing.T) {
		src := obs.Async(func(o *obs.Observer[int]) {
			o.Next(1)
			time.Sleep(100 * time.Millisecond) // triggers timeout
			o.Next(2)                          // must be ignored
		})

		timeout := 50 * time.Millisecond
		out := Timeout[int](timeout, testScheduler)(src)

		rec, o := observers.NewRecorder[int](t)
		_, err := out.Sub((*obs.Observer[int])(o))
		assertlite.NoError(t, err)

		assertlite.IsErrorOf(t, rec.AwaitError(time.Second), obs.TimeoutError{After: timeout})
		assertlite.EqualSlices(t, []int{1}, rec.Values())
	})
}
