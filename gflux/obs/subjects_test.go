// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package obs

import (
	"errors"
	"sync"
	"testing"

	"github.com/andrerrcosta2/gtools/core/testlite/assertlite"
	"github.com/andrerrcosta2/gtools/gflux/obs/internal/observers"
)

func newNopObserver[T any]() *Observer[T] {
	return &Observer[T]{
		Next:     func(t T) {},
		Error:    func(e error) {},
		Complete: func() {},
	}
}

func TestAsyncSubject(t *testing.T) {
	t.Run("emits last value on subscribe", func(t *testing.T) {
		subj := AsyncSubject[int]()

		rec, o := observers.NewRecorder[int](t)
		_, _ = subj.Sub((*Observer[int])(o))

		subj.Next(1)
		subj.Next(2)
		subj.Next(3)
		subj.Complete()

		assertlite.EqualSlices(t, rec.Values(), []int{3})
		assertlite.True(t, rec.Completed())
		assertlite.IsNil(t, rec.Err())
	})

	t.Run("completed without values", func(t *testing.T) {
		subj := AsyncSubject[int]()

		rec, o := observers.NewRecorder[int](t)
		_, _ = subj.Sub((*Observer[int])(o))

		subj.Complete()

		assertlite.EqualSlices(t, rec.Values(), []int{})
		assertlite.True(t, rec.Completed())
		assertlite.IsNil(t, rec.Err())
	})

	t.Run("error should discard values", func(t *testing.T) {
		subj := AsyncSubject[int]()

		rec, o := observers.NewRecorder[int](t)
		_, _ = subj.Sub((*Observer[int])(o))

		subj.Next(42)
		subj.Error(errors.New("failure"))

		assertlite.EqualSlices(t, rec.Values(), []int{})
		assertlite.False(t, rec.Completed())
		assertlite.NotNil(t, rec.Err())
	})

	t.Run("late subscription after complete", func(t *testing.T) {
		subj := AsyncSubject[int]()

		subj.Next(7)
		subj.Next(9)
		subj.Complete()

		rec, o := observers.NewRecorder[int](t)
		_, _ = subj.Sub((*Observer[int])(o))

		assertlite.EqualSlices(t, rec.Values(), []int{9})
		assertlite.True(t, rec.Completed())
		assertlite.IsNil(t, rec.Err())
	})

	t.Run("late subscription after error", func(t *testing.T) {
		subj := AsyncSubject[int]()

		subj.Next(10)
		subj.Error(errors.New("failed"))

		rec, o := observers.NewRecorder[int](t)
		_, _ = subj.Sub((*Observer[int])(o))

		assertlite.EqualSlices(t, rec.Values(), []int{})
		assertlite.False(t, rec.Completed())
		assertlite.NotNil(t, rec.Err())
	})

	t.Run("multicast should send same value", func(t *testing.T) {
		subj := AsyncSubject[int]()

		recA, oA := observers.NewRecorder[int](t)
		recB, oB := observers.NewRecorder[int](t)

		_, _ = subj.Sub((*Observer[int])(oA))
		_, _ = subj.Sub((*Observer[int])(oB))

		subj.Next(1)
		subj.Next(2)
		subj.Complete()

		assertlite.EqualSlices(t, recA.Values(), []int{2})
		assertlite.EqualSlices(t, recB.Values(), []int{2})
	})

	t.Run("next after termination should be ignored", func(t *testing.T) {
		subj := AsyncSubject[int]()

		rec, o := observers.NewRecorder[int](t)
		_, _ = subj.Sub((*Observer[int])(o))

		subj.Next(1)
		subj.Complete()
		subj.Next(999)

		assertlite.EqualSlices(t, rec.Values(), []int{1})
	})

}

func TestBehaviorSubj(t *testing.T) {
	t.Run("emits initial value on subscribe", func(t *testing.T) {
		subj := BehaviorSubject(10)

		rec, o := observers.NewRecorder[int](t)
		_, _ = subj.Sub((*Observer[int])(o))

		assertlite.EqualSlices(t, rec.Values(), []int{10})
		assertlite.False(t, rec.Completed())
		assertlite.IsNil(t, rec.Err())
	})

	t.Run("multicast and update value", func(t *testing.T) {
		subj := BehaviorSubject(1)

		r1, o1 := observers.NewRecorder[int](t)
		r2, o2 := observers.NewRecorder[int](t)

		_, _ = subj.Sub((*Observer[int])(o1))
		subj.Next(2)
		_, _ = subj.Sub((*Observer[int])(o2))
		subj.Next(3)

		assertlite.EqualSlices(t, r1.Values(), []int{1, 2, 3})
		assertlite.EqualSlices(t, r2.Values(), []int{2, 3})
	})

	t.Run("late subscriber receives last value", func(t *testing.T) {
		subj := BehaviorSubject(0)

		subj.Next(5)
		subj.Next(8)

		rec, o := observers.NewRecorder[int](t)
		_, _ = subj.Sub((*Observer[int])(o))

		assertlite.EqualSlices(t, rec.Values(), []int{8})
	})

	t.Run("late subscriber after complete", func(t *testing.T) {
		subj := BehaviorSubject(0)

		subj.Next(2)
		subj.Complete()

		rec, o := observers.NewRecorder[int](t)
		_, _ = subj.Sub((*Observer[int])(o))

		assertlite.EqualSlices(t, rec.Values(), []int{2})
		assertlite.True(t, rec.Completed())
		assertlite.IsNil(t, rec.Err())
	})

	t.Run("late subscriber after error", func(t *testing.T) {
		subj := BehaviorSubject(0)
		expected := errors.New("failure")

		subj.Next(3)
		subj.Error(expected)

		rec, o := observers.NewRecorder[int](t)
		_, _ = subj.Sub((*Observer[int])(o))

		assertlite.EqualSlices(t, rec.Values(), []int{3})
		assertlite.IsErrorOf(t, rec.Err(), expected)
		assertlite.False(t, rec.Completed())
	})

	t.Run("ignores next after termination", func(t *testing.T) {
		subj := BehaviorSubject(1)

		subj.Complete()
		subj.Next(2)

		rec, o := observers.NewRecorder[int](t)
		_, _ = subj.Sub((*Observer[int])(o))

		assertlite.EqualSlices(t, rec.Values(), []int{1})
	})
}

func TestConcSubj(t *testing.T) {
	t.Run("multicast", func(t *testing.T) {
		subj := ConcSlotSubject[int]()

		r1, o1 := observers.NewRecorder[int](t)
		r2, o2 := observers.NewRecorder[int](t)

		_, _ = subj.Sub((*Observer[int])(o1))
		_, _ = subj.Sub((*Observer[int])(o2))

		subj.Next(1)
		subj.Next(2)

		assertlite.EqualSlices(t, r1.Values(), []int{1, 2})
		assertlite.EqualSlices(t, r2.Values(), []int{1, 2})
	})

	t.Run("concurrent next", func(t *testing.T) {
		subj := ConcSlotSubject[int]()

		rec, o := observers.NewRecorder[int](t)
		_, _ = subj.Sub((*Observer[int])(o))

		const N = 1000
		var wg sync.WaitGroup
		wg.Add(N)

		for i := 0; i < N; i++ {
			go func(v int) {
				defer wg.Done()
				subj.Next(v)
			}(i)
		}

		wg.Wait()

		assertlite.Equals(t, len(rec.Values()), N)
	})

	t.Run("concurrent sub", func(t *testing.T) {
		subj := ConcSlotSubject[int]()
		const N = 100
		var wg sync.WaitGroup
		wg.Add(N)
		for i := 0; i < N; i++ {
			go func() {
				defer wg.Done()
				_, _ = subj.Sub(NewObserver[int](func(int) {}, nil, nil))
			}()
		}
		wg.Wait()
		subj.Next(1) // should not panic
	})

	t.Run("subscribe while emitting", func(t *testing.T) {
		subj := ConcSlotSubject[int]()
		var wg sync.WaitGroup

		for i := 0; i < 50; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				_, _ = subj.Sub(NewObserver[int](func(int) {}, nil, nil))
			}()
		}

		for i := 0; i < 50; i++ {
			wg.Add(1)
			go func(v int) {
				defer wg.Done()
				subj.Next(v)
			}(i)
		}

		wg.Wait()
	})

	t.Run("concurrent termination", func(t *testing.T) {
		subj := ConcSlotSubject[int]()

		rec, o := observers.NewRecorder[int](t)
		_, _ = subj.Sub((*Observer[int])(o))

		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			defer wg.Done()
			subj.Error(errors.New("failed"))
		}()

		go func() {
			defer wg.Done()
			subj.Next(1)
		}()

		wg.Wait()

		assertlite.NotNil(t, rec.Err())
		// At most one value may be delivered
		assertlite.True(t, len(rec.Values()) == 0 || len(rec.Values()) == 1)
	})

	t.Run("subscriber after termination", func(t *testing.T) {
		subj := ConcSlotSubject[int]()

		subj.Complete()

		rec, o := observers.NewRecorder[int](t)
		_, _ = subj.Sub((*Observer[int])(o))

		assertlite.True(t, rec.Completed())
		assertlite.EqualSlices(t, rec.Values(), []int{})
	})
}

func TestDirectSubject(t *testing.T) {
	t.Run("multicast", func(t *testing.T) {
		subj := DirectSubject[int]()

		r1, o1 := observers.NewRecorder[int](t)
		r2, o2 := observers.NewRecorder[int](t)

		_, _ = subj.Sub((*Observer[int])(o1))
		_, _ = subj.Sub((*Observer[int])(o2))

		subj.Next(1)
		subj.Next(2)

		assertlite.EqualSlices(t, r1.Values(), []int{1, 2})
		assertlite.EqualSlices(t, r2.Values(), []int{1, 2})
	})

	t.Run("unsub stops receiving immediately", func(t *testing.T) {
		subj := DirectSubject[int]()

		r1, o1 := observers.NewRecorder[int](t)
		r2, o2 := observers.NewRecorder[int](t)

		s1, _ := subj.Sub((*Observer[int])(o1))
		_, _ = subj.Sub((*Observer[int])(o2))

		subj.Next(1)
		s1.Unsub()
		subj.Next(2)

		assertlite.EqualSlices(t, r1.Values(), []int{1})
		assertlite.EqualSlices(t, r2.Values(), []int{1, 2})
	})

	t.Run("error is broadcast and terminates", func(t *testing.T) {
		subj := DirectSubject[int]()
		expected := errors.New("failure")

		r1, o1 := observers.NewRecorder[int](t)
		r2, o2 := observers.NewRecorder[int](t)

		_, _ = subj.Sub((*Observer[int])(o1))
		_, _ = subj.Sub((*Observer[int])(o2))

		subj.Error(expected)

		assertlite.IsErrorOf(t, r1.Err(), expected)
		assertlite.IsErrorOf(t, r2.Err(), expected)
		assertlite.False(t, r1.Completed())
		assertlite.False(t, r2.Completed())

		subj.Next(1)
		assertlite.EqualSlices(t, r1.Values(), []int{})
	})

	t.Run("complete is broadcast and terminates", func(t *testing.T) {
		subj := DirectSubject[int]()

		r1, o1 := observers.NewRecorder[int](t)
		r2, o2 := observers.NewRecorder[int](t)

		_, _ = subj.Sub((*Observer[int])(o1))
		_, _ = subj.Sub((*Observer[int])(o2))

		subj.Complete()

		assertlite.True(t, r1.Completed())
		assertlite.True(t, r2.Completed())
		assertlite.IsNil(t, r1.Err())
		assertlite.IsNil(t, r2.Err())

		subj.Next(1)
		assertlite.EqualSlices(t, r1.Values(), []int{})
	})

	t.Run("late subscriber after complete completes immediately", func(t *testing.T) {
		subj := DirectSubject[int]()
		subj.Complete()

		rec, o := observers.NewRecorder[int](t)
		_, _ = subj.Sub((*Observer[int])(o))

		assertlite.True(t, rec.Completed())
		assertlite.EqualSlices(t, rec.Values(), []int{})
	})

	t.Run("late subscriber after error receives error immediately", func(t *testing.T) {
		subj := DirectSubject[int]()
		expected := errors.New("fail")

		subj.Error(expected)

		rec, o := observers.NewRecorder[int](t)
		_, _ = subj.Sub((*Observer[int])(o))

		assertlite.IsErrorOf(t, rec.Err(), expected)
		assertlite.False(t, rec.Completed())
	})

	t.Run("unsub all closes all subscriptions", func(t *testing.T) {
		subj := DirectSubject[int]()

		r, o := observers.NewRecorder[int](t)
		sub, err := subj.Sub((*Observer[int])(o))
		assertlite.NoError(t, err)
		subj.UnsubAll()
		subj.Next(1)

		assertlite.EqualSlices(t, r.Values(), []int{})
		assertlite.True(t, sub.Closed())
	})
}

// TestManySubj tests the manySubj:
//
// Semantics:
//  1. behaves like a normal multicast subject.
//  2. accepts at most capacity active subscriptions.
//  3. a subscription slot counts only if it is alive.
//  4. If the subject is terminated
//     - late subscribers receive the terminal signal immediately.
//     - capacity is irrelevant after termination.
//  5. If capacity is exceeded:
//     -Sub returns an error
//     - The observer is not registered
func TestManySubj(t *testing.T) {
	t.Run("accepts up to capacity", func(t *testing.T) {
		subj := ManySubject[int](2)

		_, err1 := subj.Sub(newNopObserver[int]())
		_, err2 := subj.Sub(newNopObserver[int]())

		assertlite.IsNil(t, err1)
		assertlite.IsNil(t, err2)
	})

	t.Run("rejects when capacity exceeded", func(t *testing.T) {
		subj := ManySubject[int](1)

		_, err1 := subj.Sub(newNopObserver[int]())
		_, err2 := subj.Sub(newNopObserver[int]())

		assertlite.IsNil(t, err1)
		assertlite.NotNil(t, err2)
	})

	t.Run("unsubscribe frees capacity", func(t *testing.T) {
		subj := ManySubject[int](1)

		sub, err := subj.Sub(newNopObserver[int]())
		assertlite.IsNil(t, err)

		sub.Unsub()

		_, err = subj.Sub(newNopObserver[int]())
		assertlite.IsNil(t, err)
	})

	t.Run("only active subscribers receive values", func(t *testing.T) {
		subj := ManySubject[int](2)

		rec1, o1 := observers.NewRecorder[int](t)
		rec2, o2 := observers.NewRecorder[int](t)

		sub1, _ := subj.Sub((*Observer[int])(o1))
		_, _ = subj.Sub((*Observer[int])(o2))

		sub1.Unsub()

		subj.Next(42)

		assertlite.EqualSlices(t, rec1.Values(), []int{})
		assertlite.EqualSlices(t, rec2.Values(), []int{42})
	})

	t.Run("termination ignores capacity", func(t *testing.T) {
		subj := ManySubject[int](0)

		subj.Complete()

		rec, o := observers.NewRecorder[int](t)
		_, err := subj.Sub((*Observer[int])(o))

		assertlite.IsNil(t, err)
		assertlite.True(t, rec.Completed())
	})

	t.Run("late subscriber after error receives error", func(t *testing.T) {
		subj := ManySubject[int](1)

		errX := errors.New("failure")
		subj.Error(errX)

		rec, o := observers.NewRecorder[int](t)
		_, err := subj.Sub((*Observer[int])(o))

		assertlite.IsNil(t, err)
		assertlite.Equals(t, rec.Err(), errX)
	})

	t.Run("capacity failure does not register observer", func(t *testing.T) {
		subj := ManySubject[int](1)

		_, _ = subj.Sub(newNopObserver[int]())

		rec, o := observers.NewRecorder[int](t)
		_, err := subj.Sub((*Observer[int])(o))

		assertlite.NotNil(t, err)

		subj.Next(1)
		assertlite.EqualSlices(t, rec.Values(), []int{})
	})
}

func TestMonoSubject(t *testing.T) {
	t.Run("single subscriber receives values", func(t *testing.T) {
		subj := MonoSubject[int]()

		rec, o := observers.NewRecorder[int](t)
		_, err := subj.Sub((*Observer[int])(o))
		assertlite.IsNil(t, err)

		subj.Next(1)
		subj.Next(2)

		assertlite.EqualSlices(t, rec.Values(), []int{1, 2})
		assertlite.IsNil(t, rec.Err())
		assertlite.False(t, rec.Completed())
	})

	t.Run("rejects second subscriber", func(t *testing.T) {
		subj := MonoSubject[int]()

		_, o1 := observers.NewRecorder[int](t)
		_, o2 := observers.NewRecorder[int](t)

		_, err1 := subj.Sub((*Observer[int])(o1))
		_, err2 := subj.Sub((*Observer[int])(o2))

		assertlite.IsNil(t, err1)
		assertlite.NotNil(t, err2)
	})

	t.Run("allows resubscribe after unsubscribe", func(t *testing.T) {
		subj := MonoSubject[int]()

		rec1, o1 := observers.NewRecorder[int](t)
		sub1, _ := subj.Sub((*Observer[int])(o1))

		subj.Next(1)
		sub1.Unsub()

		rec2, o2 := observers.NewRecorder[int](t)
		_, err := subj.Sub((*Observer[int])(o2))
		assertlite.IsNil(t, err)

		subj.Next(2)

		assertlite.EqualSlices(t, rec1.Values(), []int{1})
		assertlite.EqualSlices(t, rec2.Values(), []int{2})
	})

	t.Run("error terminates subject", func(t *testing.T) {
		subj := MonoSubject[int]()

		rec, o := observers.NewRecorder[int](t)
		_, _ = subj.Sub((*Observer[int])(o))

		err := errors.New("fail")
		subj.Error(err)

		assertlite.IsErrorOf(t, rec.Err(), err)
		assertlite.False(t, rec.Completed())

		subj.Next(1) // ignored
		assertlite.EqualSlices(t, rec.Values(), nil)
	})

	t.Run("complete terminates subject", func(t *testing.T) {
		subj := MonoSubject[int]()

		rec, o := observers.NewRecorder[int](t)
		subj.Sub((*Observer[int])(o))

		subj.Complete()

		assertlite.True(t, rec.Completed())

		subj.Next(1)
		assertlite.EqualSlices(t, rec.Values(), nil)
	})

	t.Run("late subscriber after error", func(t *testing.T) {
		subj := MonoSubject[int]()

		subj.Error(errors.New("fail"))

		rec, o := observers.NewRecorder[int](t)
		_, _ = subj.Sub((*Observer[int])(o))

		assertlite.NotNil(t, rec.Err())
		assertlite.False(t, rec.Completed())
	})

	t.Run("late subscriber after completion", func(t *testing.T) {
		subj := MonoSubject[int]()

		subj.Complete()

		rec, o := observers.NewRecorder[int](t)
		_, _ = subj.Sub((*Observer[int])(o))

		assertlite.True(t, rec.Completed())
		assertlite.IsNil(t, rec.Err())
	})

	t.Run("next before subscribe is no-op", func(t *testing.T) {
		subj := MonoSubject[int]()

		assertlite.NoPanics(t, func() {
			subj.Next(1)
		})
	})
}

func TestReplaySubject(t *testing.T) {
	t.Run("replays all past values (unbounded)", func(t *testing.T) {
		subj := ReplaySubject[int]()

		subj.Next(1)
		subj.Next(2)
		subj.Next(3)

		rec, o := observers.NewRecorder[int](t)
		_, _ = subj.Sub((*Observer[int])(o))

		assertlite.EqualSlices(t, rec.Values(), []int{1, 2, 3})
		assertlite.False(t, rec.Completed())
		assertlite.IsNil(t, rec.Err())
	})

	t.Run("replays only last N values (bounded)", func(t *testing.T) {
		subj := ReplaySubject[int]()

		subj.Next(1)
		subj.Next(2)
		subj.Next(3)

		rec, o := observers.NewRecorder[int](t)
		_, _ = subj.Sub((*Observer[int])(o))

		assertlite.EqualSlices(t, rec.Values(), []int{2, 3})
	})

	t.Run("replay then live values", func(t *testing.T) {
		subj := ReplaySubject[int]()

		subj.Next(1)
		subj.Next(2)

		rec, o := observers.NewRecorder[int](t)
		_, _ = subj.Sub((*Observer[int])(o))

		subj.Next(3)

		assertlite.EqualSlices(t, rec.Values(), []int{1, 2, 3})
	})

	t.Run("multiple subscribers receive same replay", func(t *testing.T) {
		subj := ReplaySubject[int]()

		subj.Next(1)
		subj.Next(2)
		subj.Next(3)

		r1, o1 := observers.NewRecorder[int](t)
		r2, o2 := observers.NewRecorder[int](t)

		_, _ = subj.Sub((*Observer[int])(o1))
		_, _ = subj.Sub((*Observer[int])(o2))

		assertlite.EqualSlices(t, r1.Values(), []int{1, 2, 3})
		assertlite.EqualSlices(t, r2.Values(), []int{1, 2, 3})
	})

	t.Run("late subscriber receives replay then complete", func(t *testing.T) {
		subj := ReplaySubject[int]()

		subj.Next(1)
		subj.Next(2)
		subj.Complete()

		rec, o := observers.NewRecorder[int](t)
		_, _ = subj.Sub((*Observer[int])(o))

		assertlite.EqualSlices(t, rec.Values(), []int{1, 2})
		assertlite.True(t, rec.Completed())
		assertlite.IsNil(t, rec.Err())
	})

	t.Run("late subscriber receives replay then error", func(t *testing.T) {
		subj := ReplaySubject[int]()
		expected := errors.New("failure")

		subj.Next(1)
		subj.Next(2)
		subj.Next(3)
		subj.Error(expected)

		rec, o := observers.NewRecorder[int](t)
		_, _ = subj.Sub((*Observer[int])(o))

		assertlite.EqualSlices(t, rec.Values(), []int{2, 3})
		assertlite.IsErrorOf(t, rec.Err(), expected)
		assertlite.False(t, rec.Completed())
	})

	t.Run("no emissions after termination", func(t *testing.T) {
		subj := ReplaySubject[int]()

		subj.Next(1)
		subj.Complete()
		subj.Next(2)

		rec, o := observers.NewRecorder[int](t)
		_, _ = subj.Sub((*Observer[int])(o))

		assertlite.EqualSlices(t, rec.Values(), []int{1})
		assertlite.True(t, rec.Completed())
	})

	t.Run("bounded replay overwrites oldest values", func(t *testing.T) {
		subj := ReplaySubject[int]()

		subj.Next(1)
		subj.Next(2)
		subj.Next(3)

		rec, o := observers.NewRecorder[int](t)
		_, _ = subj.Sub((*Observer[int])(o))

		assertlite.EqualSlices(t, rec.Values(), []int{3})
	})
}

func TestSubject(t *testing.T) {
	t.Run("multicast", func(t *testing.T) {
		subj := SlotSubject[int]()

		r1, o1 := observers.NewRecorder[int](t)
		r2, o2 := observers.NewRecorder[int](t)

		_, _ = subj.Sub((*Observer[int])(o1))
		_, _ = subj.Sub((*Observer[int])(o2))

		subj.Next(1)
		subj.Next(2)

		assertlite.EqualSlices(t, r1.Values(), []int{1, 2})
		assertlite.EqualSlices(t, r2.Values(), []int{1, 2})
	})

	t.Run("unsub stops receiving", func(t *testing.T) {
		subj := SlotSubject[int]()

		r1, o1 := observers.NewRecorder[int](t)
		r2, o2 := observers.NewRecorder[int](t)

		s1, _ := subj.Sub((*Observer[int])(o1))
		_, _ = subj.Sub((*Observer[int])(o2))

		subj.Next(1)
		s1.Unsub()
		subj.Next(2)

		assertlite.EqualSlices(t, r1.Values(), []int{1})
		assertlite.EqualSlices(t, r2.Values(), []int{1, 2})
	})

	t.Run("error is broadcast", func(t *testing.T) {
		subj := SlotSubject[int]()
		expected := errors.New("fail")

		r1, o1 := observers.NewRecorder[int](t)
		r2, o2 := observers.NewRecorder[int](t)

		_, _ = subj.Sub((*Observer[int])(o1))
		_, _ = subj.Sub((*Observer[int])(o2))

		subj.Error(expected)

		assertlite.IsErrorOf(t, r1.Err(), expected)
		assertlite.IsErrorOf(t, r2.Err(), expected)
		assertlite.False(t, r1.Completed())
		assertlite.False(t, r2.Completed())
	})

	t.Run("complete is broadcast", func(t *testing.T) {
		subj := SlotSubject[int]()

		r1, o1 := observers.NewRecorder[int](t)
		r2, o2 := observers.NewRecorder[int](t)

		_, _ = subj.Sub((*Observer[int])(o1))
		_, _ = subj.Sub((*Observer[int])(o2))

		subj.Complete()

		assertlite.True(t, r1.Completed())
		assertlite.True(t, r2.Completed())
		assertlite.IsNil(t, r1.Err())
		assertlite.IsNil(t, r2.Err())
	})

	t.Run("unsub all", func(t *testing.T) {
		subj := SlotSubject[int]()

		r, o := observers.NewRecorder[int](t)
		_, _ = subj.Sub((*Observer[int])(o))

		subj.UnsubAll()
		subj.Next(1)

		assertlite.EqualSlices(t, r.Values(), []int{})
	})

	t.Run("late subscriber completes immediately", func(t *testing.T) {
		subj := SlotSubject[int]()
		subj.Complete()

		rec, o := observers.NewRecorder[int](t)
		_, _ = subj.Sub((*Observer[int])(o))

		assertlite.True(t, rec.Completed())
		assertlite.EqualSlices(t, rec.Values(), []int{})
	})

	t.Run("no emission after error", func(t *testing.T) {
		subj := SlotSubject[int]()
		subj.Error(errors.New("failure"))

		rec, o := observers.NewRecorder[int](t)
		_, err := subj.Sub((*Observer[int])(o))
		assertlite.NoError(t, err)
		subj.Next(1)

		assertlite.EqualSlices(t, rec.Values(), []int{})
	})
}
