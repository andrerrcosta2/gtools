// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package observables

import (
	"errors"
	"testing"

	"github.com/andrerrcosta2/gtools/core/testlite/assertlite"
	"github.com/andrerrcosta2/gtools/gflux/obs/internal/observers"
)

// TestSlotSubject_Sync tests the SlotSubject, which is a hot,
// multicasting subject that directly owns its observers.
//
// For synchronous usage, we assert:
//  1. Next emits values to all active subscribers
//  2. Error terminates and notifies all subscribers
//  3. Complete terminates and notifies all subscribers
//  4. Late subscribers receive terminal signals immediately
//  5. UnsubAll closes all active subscriptions
//  6. Values emitted after termination are ignored
func TestSlotSubject_Sync(t *testing.T) {
	t.Run("next emits values to all active subscribers", func(t *testing.T) {
		subj := SlotSubject[int]()

		recA, oA := observers.NewRecorder[int](t)
		recB, oB := observers.NewRecorder[int](t)

		_, _ = subj.Sub((*observers.Observer[int])(oA))
		_, _ = subj.Sub((*observers.Observer[int])(oB))

		subj.Next(1)
		subj.Next(2)

		assertlite.EqualSlices(t, []int{1, 2}, recA.Values())
		assertlite.EqualSlices(t, []int{1, 2}, recB.Values())
		assertlite.False(t, recA.Completed())
		assertlite.False(t, recB.Completed())
	})

	t.Run("error terminates all active subscribers", func(t *testing.T) {
		subj := SlotSubject[int]()
		recA, oA := observers.NewRecorder[int](t)
		recB, oB := observers.NewRecorder[int](t)
		_, _ = subj.Sub((*observers.Observer[int])(oA))
		_, _ = subj.Sub((*observers.Observer[int])(oB))

		failure := errors.New("failure")
		subj.Error(failure)

		assertlite.IsErrorOf(t, recA.Err(), failure)
		assertlite.IsErrorOf(t, recB.Err(), failure)
		assertlite.False(t, recA.Completed())
		assertlite.False(t, recB.Completed())
	})

	t.Run("complete terminates all active subscribers", func(t *testing.T) {
		subj := SlotSubject[int]()
		recA, oA := observers.NewRecorder[int](t)
		recB, oB := observers.NewRecorder[int](t)
		_, _ = subj.Sub((*observers.Observer[int])(oA))
		_, _ = subj.Sub((*observers.Observer[int])(oB))

		subj.Complete()

		assertlite.True(t, recA.Completed())
		assertlite.True(t, recB.Completed())
		assertlite.NoError(t, recA.Err())
		assertlite.NoError(t, recB.Err())
	})

	t.Run("late subscribers receive terminal signals immediately", func(t *testing.T) {
		subj := SlotSubject[int]()
		subj.Complete() // terminate before subscribing

		rec, o := observers.NewRecorder[int](t)
		_, _ = subj.Sub((*observers.Observer[int])(o))

		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("unsubAll closes all active subscriptions", func(t *testing.T) {
		subj := SlotSubject[int]()
		recA, oA := observers.NewRecorder[int](t)
		recB, oB := observers.NewRecorder[int](t)
		subA, _ := subj.Sub((*observers.Observer[int])(oA))
		subB, _ := subj.Sub((*observers.Observer[int])(oB))

		subj.UnsubAll()

		// further emissions are ignored
		subj.Next(42)
		subj.Complete()

		assertlite.Equals(t, 0, len(recA.Values()))
		assertlite.Equals(t, 0, len(recB.Values()))
		assertlite.False(t, recA.Completed())
		assertlite.False(t, recB.Completed())

		// individual subscriptions report closed
		assertlite.True(t, subA.Closed())
		assertlite.True(t, subB.Closed())
	})

	t.Run("values emitted after termination are ignored", func(t *testing.T) {
		subj := SlotSubject[int]()
		rec, o := observers.NewRecorder[int](t)
		_, _ = subj.Sub((*observers.Observer[int])(o))

		subj.Complete()
		subj.Next(100)
		subj.Error(errors.New("fail"))

		assertlite.Equals(t, 0, len(rec.Values()))
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})
}

// TestAsyncSubject_Sync tests AsyncSubject, which only emits the
// last value upon completion.
//
// For synchronous usage, we assert:
//  1. Only the last Next value is emitted on Complete
//  2. Error terminates immediately and discards any stored value
//  3. Late subscribers receive last value if completed successfully
//  4. Late subscribers receive error if terminated with error
//  5. Values emitted after termination are ignored
//  6. UnsubAll closes all active subscriptions
func TestAsyncSubject_Sync(t *testing.T) {
	t.Run("emits only last value on complete", func(t *testing.T) {
		subj := AsyncSubject[int]()
		rec, o := observers.NewRecorder[int](t)
		_, _ = subj.Sub((*observers.Observer[int])(o))

		subj.Next(1)
		subj.Next(2)
		subj.Next(3)
		subj.Complete()

		assertlite.EqualSlices(t, []int{3}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("error terminates immediately and discards last value", func(t *testing.T) {
		subj := AsyncSubject[int]()
		rec, o := observers.NewRecorder[int](t)
		_, _ = subj.Sub((*observers.Observer[int])(o))

		subj.Next(42)
		errFail := errors.New("fail")
		subj.Error(errFail)

		assertlite.Equals(t, 0, len(rec.Values()))
		assertlite.IsErrorOf(t, rec.Err(), errFail)
		assertlite.False(t, rec.Completed())
	})

	t.Run("late subscribers receive last value if completed successfully", func(t *testing.T) {
		subj := AsyncSubject[int]()
		subj.Next(1)
		subj.Next(2)
		subj.Next(3)
		subj.Complete()

		rec, o := observers.NewRecorder[int](t)
		_, _ = subj.Sub((*observers.Observer[int])(o))

		assertlite.EqualSlices(t, []int{3}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("late subscribers receive error if terminated with error", func(t *testing.T) {
		subj := AsyncSubject[int]()
		errFail := errors.New("fail")
		subj.Error(errFail)

		rec, o := observers.NewRecorder[int](t)
		_, _ = subj.Sub((*observers.Observer[int])(o))

		assertlite.Equals(t, 0, len(rec.Values()))
		assertlite.IsErrorOf(t, rec.Err(), errFail)
		assertlite.False(t, rec.Completed())
	})

	t.Run("values emitted after termination are ignored", func(t *testing.T) {
		subj := AsyncSubject[int]()
		rec, o := observers.NewRecorder[int](t)
		_, _ = subj.Sub((*observers.Observer[int])(o))

		subj.Complete()
		subj.Next(999)
		subj.Error(errors.New("fail"))

		assertlite.Equals(t, 0, len(rec.Values()))
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("unsubAll closes all active subscriptions", func(t *testing.T) {
		subj := AsyncSubject[int]()
		rec, o := observers.NewRecorder[int](t)
		sub, _ := subj.Sub((*observers.Observer[int])(o))

		subj.UnsubAll()

		// further emissions are ignored
		subj.Next(100)
		subj.Complete()

		assertlite.Equals(t, 0, len(rec.Values()))
		assertlite.False(t, rec.Completed())
		assertlite.NoError(t, rec.Err())

		assertlite.True(t, sub.Closed())
	})
}

// TestBehaviorSubject_Sync tests BehaviorSubject, which always holds the latest value.
//
// For synchronous usage, we assert:
//  1. New subscribers immediately receive the current value
//  2. All subsequent Next calls are broadcast to all subscribers
//  3. Error terminates the subject and is broadcast to all
//  4. Complete terminates the subject and is broadcast to all
//  5. Late subscribers receive the last value followed by the terminal signal
//  6. Values emitted after termination are ignored
//  7. UnsubAll closes all active subscriptions
func TestBehaviorSubject_Sync(t *testing.T) {
	t.Run("new subscriber receives current value", func(t *testing.T) {
		subj := BehaviorSubject(0)
		rec, o := observers.NewRecorder[int](t)
		_, _ = subj.Sub((*observers.Observer[int])(o))

		assertlite.EqualSlices(t, []int{0}, rec.Values())
		assertlite.False(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("next calls are broadcast to all subscribers", func(t *testing.T) {
		subj := BehaviorSubject(0)
		recA, oA := observers.NewRecorder[int](t)
		recB, oB := observers.NewRecorder[int](t)
		_, _ = subj.Sub((*observers.Observer[int])(oA))
		_, _ = subj.Sub((*observers.Observer[int])(oB))

		subj.Next(1)
		subj.Next(2)

		assertlite.EqualSlices(t, []int{0, 1, 2}, recA.Values())
		assertlite.EqualSlices(t, []int{0, 1, 2}, recB.Values())
	})

	t.Run("error terminates immediately and broadcasts to all", func(t *testing.T) {
		subj := BehaviorSubject(0)
		recA, oA := observers.NewRecorder[int](t)
		recB, oB := observers.NewRecorder[int](t)
		_, _ = subj.Sub((*observers.Observer[int])(oA))
		_, _ = subj.Sub((*observers.Observer[int])(oB))

		errFail := errors.New("fail")
		subj.Error(errFail)

		assertlite.IsErrorOf(t, recA.Err(), errFail)
		assertlite.IsErrorOf(t, recB.Err(), errFail)
		assertlite.False(t, recA.Completed())
		assertlite.False(t, recB.Completed())
	})

	t.Run("complete terminates and broadcasts to all", func(t *testing.T) {
		subj := BehaviorSubject(0)
		recA, oA := observers.NewRecorder[int](t)
		recB, oB := observers.NewRecorder[int](t)
		_, _ = subj.Sub((*observers.Observer[int])(oA))
		_, _ = subj.Sub((*observers.Observer[int])(oB))

		subj.Next(1)
		subj.Complete()

		assertlite.True(t, recA.Completed())
		assertlite.True(t, recB.Completed())
		assertlite.NoError(t, recA.Err())
		assertlite.NoError(t, recB.Err())
		assertlite.EqualSlices(t, []int{0, 1}, recA.Values())
		assertlite.EqualSlices(t, []int{0, 1}, recB.Values())
	})

	t.Run("late subscribers receive last value followed by terminal signal", func(t *testing.T) {
		subj := BehaviorSubject(0)
		subj.Next(1)
		subj.Next(2)
		subj.Complete()

		rec, o := observers.NewRecorder[int](t)
		_, _ = subj.Sub((*observers.Observer[int])(o))

		assertlite.EqualSlices(t, []int{2}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("values emitted after termination are ignored", func(t *testing.T) {
		subj := BehaviorSubject(0)
		rec, o := observers.NewRecorder[int](t)
		_, _ = subj.Sub((*observers.Observer[int])(o))

		subj.Complete()
		subj.Next(999)
		subj.Error(errors.New("fail"))

		assertlite.EqualSlices(t, []int{0}, rec.Values())
		assertlite.True(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
	})

	t.Run("unsubAll closes all active subscriptions", func(t *testing.T) {
		subj := BehaviorSubject(0)
		rec, o := observers.NewRecorder[int](t)
		sub, _ := subj.Sub((*observers.Observer[int])(o))

		subj.UnsubAll()
		subj.Next(100)
		subj.Complete()

		assertlite.EqualSlices(t, []int{0}, rec.Values())
		assertlite.False(t, rec.Completed())
		assertlite.NoError(t, rec.Err())
		assertlite.True(t, sub.Closed())
	})
}

// TestDirectSubject tests the DirectSubject implementation.
//
// DirectSubject is a direct, slotless, multicasting Subject:
// 1. Immediately emits values to all active subscribers.
// 2. Terminates properly on Complete or Error.
// 3. Allows unsubscribing individual subscribers via Close.
// 4. Supports late subscribers receiving terminal signals only.
//
// We will test:
//   - Emitting values to multiple subscribers
//   - Emission order and completeness
//   - Error propagation
//   - Late subscriptions after termination
//   - Individual Unsub behavior
func TestDirectSubject(t *testing.T) {
	t.Run("emits values to multiple subscribers", func(t *testing.T) {
		subj := DirectSubject[int]()

		recA, oA := observers.NewRecorder[int](t)
		recB, oB := observers.NewRecorder[int](t)

		_, _ = subj.Sub(oA)
		_, _ = subj.Sub(oB)

		subj.Next(1)
		subj.Next(2)
		subj.Complete()

		assertlite.EqualSlices(t, []int{1, 2}, recA.Values())
		assertlite.EqualSlices(t, []int{1, 2}, recB.Values())
		assertlite.True(t, recA.Completed())
		assertlite.True(t, recB.Completed())
	})

	t.Run("stops emitting after termination", func(t *testing.T) {
		subj := DirectSubject[int]()

		rec, o := observers.NewRecorder[int](t)
		_, _ = subj.Sub(o)

		subj.Next(1)
		subj.Complete()
		subj.Next(2) // must be ignored

		assertlite.EqualSlices(t, []int{1}, rec.Values())
		assertlite.True(t, rec.Completed())
	})

	t.Run("propagates error to all subscribers", func(t *testing.T) {
		subj := DirectSubject[int]()

		recA, oA := observers.NewRecorder[int](t)
		recB, oB := observers.NewRecorder[int](t)

		_, _ = subj.Sub(oA)
		_, _ = subj.Sub(oB)

		errFail := errors.New("fail")
		subj.Next(1)
		subj.Error(errFail)
		subj.Next(2) // ignored

		assertlite.EqualSlices(t, []int{1}, recA.Values())
		assertlite.EqualSlices(t, []int{1}, recB.Values())
		assertlite.IsErrorOf(t, recA.Err(), errFail)
		assertlite.IsErrorOf(t, recB.Err(), errFail)
	})

	t.Run("late subscribers receive terminal signals only", func(t *testing.T) {
		subj := DirectSubject[int]()

		subj.Next(1)
		subj.Complete()

		rec, o := observers.NewRecorder[int](t)
		_, _ = subj.Sub(o)

		assertlite.IsNil(t, rec.Values())
		assertlite.True(t, rec.Completed())
	})

	t.Run("individual unsubscribe stops receiving values", func(t *testing.T) {
		subj := DirectSubject[int]()

		recA, oA := observers.NewRecorder[int](t)
		recB, oB := observers.NewRecorder[int](t)

		subA, _ := subj.Sub(oA)
		_, _ = subj.Sub(oB)

		subj.Next(1)
		subA.Unsub() // unsubscribe A
		subj.Next(2)
		subj.Complete()

		assertlite.EqualSlices(t, []int{1}, recA.Values()) // only received first
		assertlite.EqualSlices(t, []int{1, 2}, recB.Values())
		assertlite.True(t, recB.Completed())
	})
}

// TestManySubject tests ManySubject, which allows a bounded number of subscribers.
//
// ManySubject multicasts values to all active subscribers, enforces a subscription limit,
// and handles termination correctly. This suite verifies:
// 1. Emission to multiple subscribers
// 2. Enforcing the maximum subscriber limit
// 3. Unsubscribing frees capacity
// 4. Terminal signals (Complete and Error) propagate to current and late subscribers
// 5. Ignoring further Next calls after termination
func TestManySubject(t *testing.T) {
	t.Run("emits values to multiple subscribers within capacity", func(t *testing.T) {
		subj := ManySubject[int](3)

		recA, oA := observers.NewRecorder[int](t)
		recB, oB := observers.NewRecorder[int](t)

		_, _ = subj.Sub(oA)
		_, _ = subj.Sub(oB)

		subj.Next(1)
		subj.Next(2)
		subj.Complete()

		assertlite.EqualSlices(t, []int{1, 2}, recA.Values())
		assertlite.EqualSlices(t, []int{1, 2}, recB.Values())
		assertlite.True(t, recA.Completed())
		assertlite.True(t, recB.Completed())
	})

	t.Run("rejects subscribers beyond max capacity", func(t *testing.T) {
		subj := ManySubject[int](1)

		_, _ = subj.Sub(&observers.Observer[int]{})
		_, _ = subj.Sub(&observers.Observer[int]{})

		subC := &observers.Observer[int]{}
		sub, err := subj.Sub(subC)

		assertlite.IsNil(t, sub)
		assertlite.True(t, err.Error() == "maximum subscribers reached")
	})

	t.Run("unsub frees capacity for new subscribers", func(t *testing.T) {
		subj := ManySubject[int](2)

		recA, oA := observers.NewRecorder[int](t)
		recB, oB := observers.NewRecorder[int](t)
		recC, oC := observers.NewRecorder[int](t)

		// Fill capacity
		subA, err := subj.Sub(oA)
		assertlite.NoError(t, err)

		_, err = subj.Sub(oB)
		assertlite.NoError(t, err)

		// Len must be exhausted
		_, err = subj.Sub(oC)
		assertlite.IsErrorOf(t, err, MaxSubsError)

		// Emit while full
		subj.Next(1)

		// Free one slot
		subA.Unsub()

		// AtNow capacity must be available
		_, err = subj.Sub(oC)
		assertlite.NoError(t, err)

		// Emit again
		subj.Next(2)
		subj.Complete()

		assertlite.EqualSlices(t, []int{1}, recA.Values())
		assertlite.EqualSlices(t, []int{1, 2}, recB.Values())
		assertlite.EqualSlices(t, []int{2}, recC.Values())
	})

	t.Run("propagates error to all active subscribers", func(t *testing.T) {
		subj := ManySubject[int](3)

		recA, oA := observers.NewRecorder[int](t)
		recB, oB := observers.NewRecorder[int](t)

		_, _ = subj.Sub(oA)
		_, _ = subj.Sub(oB)

		errFail := errors.New("fail")
		subj.Next(1)
		subj.Error(errFail)
		subj.Next(2) // ignored

		assertlite.EqualSlices(t, []int{1}, recA.Values())
		assertlite.EqualSlices(t, []int{1}, recB.Values())
		assertlite.IsErrorOf(t, recA.Err(), errFail)
		assertlite.IsErrorOf(t, recB.Err(), errFail)
	})

	t.Run("late subscribers receive terminal signals", func(t *testing.T) {
		subj := ManySubject[int](2)

		subj.Next(1)
		subj.Complete()

		rec, o := observers.NewRecorder[int](t)
		_, _ = subj.Sub(o)

		assertlite.IsNil(t, rec.Values())
		assertlite.True(t, rec.Completed())
	})
}

func TestMonoSubject(t *testing.T) {
	t.Run("single subscriber receives values and completion", func(t *testing.T) {
		subj := MonoSubject[int]()

		rec, o := observers.NewRecorder[int](t)
		sub, err := subj.Sub(o)
		assertlite.NoError(t, err)
		assertlite.NotNil(t, sub)

		subj.Next(1)
		subj.Next(2)
		subj.Complete()

		assertlite.EqualSlices(t, []int{1, 2}, rec.Values())
		assertlite.True(t, rec.Completed())
	})

	t.Run("rejects additional subscribers while active", func(t *testing.T) {
		subj := MonoSubject[int]()

		_, _ = subj.Sub(&observers.Observer[int]{})
		sub, err := subj.Sub(&observers.Observer[int]{})
		assertlite.IsNil(t, sub)
		assertlite.True(t, err.Error() == "mono subject allows only one subscriber")
	})

	t.Run("late subscriber receives terminal signal", func(t *testing.T) {
		subj := MonoSubject[int]()

		subj.Next(1)
		subj.Complete()

		rec, o := observers.NewRecorder[int](t)
		_, _ = subj.Sub(o)

		assertlite.IsNil(t, rec.Values())
		assertlite.True(t, rec.Completed())
	})

	t.Run("unsub frees slot for new subscriber if not terminated", func(t *testing.T) {
		subj := MonoSubject[int]()

		recA, oA := observers.NewRecorder[int](t)
		_, _ = subj.Sub(oA)
		subj.UnsubAll() // frees slot

		recB, oB := observers.NewRecorder[int](t)
		subB, err := subj.Sub(oB)
		assertlite.NoError(t, err)
		assertlite.NotNil(t, subB)

		subj.Next(42)
		subj.Complete()

		assertlite.IsNil(t, recA.Values())
		assertlite.EqualSlices(t, []int{42}, recB.Values())
	})
}
