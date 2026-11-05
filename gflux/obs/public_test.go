// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package obs

import (
	"errors"
	"testing"

	"github.com/andrerrcosta2/gtools/core/testlite/assertlite"
	"github.com/andrerrcosta2/gtools/gflux/obs/internal/observers"
)

func TestOf(t *testing.T) {
	t.Run("emits values in order and completes", func(t *testing.T) {
		src := Of(1, 2, 3)
		rec, o := observers.NewRecorder[int](t)
		sub, err := src.Sub((*Observer[int])(o))

		assertlite.NoError(t, err)
		assertlite.NotNil(t, sub)
		assertlite.EqualSlices(t, rec.Values(), []int{1, 2, 3})
		assertlite.IsNil(t, rec.Err())
		assertlite.True(t, rec.Completed())
	})

	t.Run("empty completes immediately", func(t *testing.T) {
		src := Of[int]()
		rec, o := observers.NewRecorder[int](t)
		_, err := src.Sub((*Observer[int])(o))

		assertlite.NoError(t, err)
		assertlite.EqualSlices(t, rec.Values(), []int{})
		assertlite.True(t, rec.Completed())
	})

	t.Run("cold behaviour", func(t *testing.T) {
		src := Of(1, 2, 3)
		r1, o1 := observers.NewRecorder[int](t)
		r2, o2 := observers.NewRecorder[int](t)
		_, _ = src.Sub((*Observer[int])(o1))
		_, _ = src.Sub((*Observer[int])(o2))

		assertlite.EqualSlices(t, r1.Values(), []int{1, 2, 3})
		assertlite.EqualSlices(t, r2.Values(), []int{1, 2, 3})
		assertlite.True(t, r1.Completed())
		assertlite.True(t, r2.Completed())
	})
}

func TestDefer(t *testing.T) {
	t.Run("lazy behaviour", func(t *testing.T) {
		called := false

		src := Defer(func() Observable[int] {
			called = true
			return Of(1)
		})
		assertlite.False(t, called)
		_, _ = src.Sub(NewObserver[int](func(int) {}, nil, nil))
		assertlite.True(t, called)
	})

	t.Run("called per subscription", func(t *testing.T) {
		calls := 0

		src := Defer(func() Observable[int] {
			calls++
			return Of(calls)
		})

		r1, o1 := observers.NewRecorder[int](t)
		r2, o2 := observers.NewRecorder[int](t)

		_, _ = src.Sub((*Observer[int])(o1))
		_, _ = src.Sub((*Observer[int])(o2))

		assertlite.EqualSlices(t, r1.Values(), []int{1})
		assertlite.EqualSlices(t, r2.Values(), []int{2})
		assertlite.Equals(t, calls, 2)
	})

	t.Run("cold behaviour", func(t *testing.T) {
		src := Defer(func() Observable[int] {
			return Of(1, 2, 3)
		})

		r1, o1 := observers.NewRecorder[int](t)
		r2, o2 := observers.NewRecorder[int](t)

		_, _ = src.Sub((*Observer[int])(o1))
		_, _ = src.Sub((*Observer[int])(o2))

		assertlite.EqualSlices(t, r1.Values(), []int{1, 2, 3})
		assertlite.EqualSlices(t, r2.Values(), []int{1, 2, 3})
		assertlite.True(t, r1.Completed())
		assertlite.True(t, r2.Completed())
	})

	t.Run("error propagation", func(t *testing.T) {
		expected := errors.New("test")
		src := Defer(func() Observable[int] {
			return FromFunc(func(o *Observer[int]) {
				o.Error(expected)
			})
		})
		rec, o := observers.NewRecorder[int](t)
		_, _ = src.Sub((*Observer[int])(o))
		assertlite.IsErrorOf(t, rec.Err(), expected)
		assertlite.False(t, rec.Completed())
	})
}

func TestFromFunc(t *testing.T) {
	t.Run("emits values and completes", func(t *testing.T) {
		src := FromFunc(func(o *Observer[int]) {
			o.Next(1)
			o.Next(2)
			o.Next(3)
			o.Complete()
		})

		rec, o := observers.NewRecorder[int](t)
		_, err := src.Sub((*Observer[int])(o))

		assertlite.NoError(t, err)
		assertlite.EqualSlices(t, rec.Values(), []int{1, 2, 3})
		assertlite.IsNil(t, rec.Err())
		assertlite.True(t, rec.Completed())
	})

	t.Run("error does not complete", func(t *testing.T) {
		expected := errors.New("failure")

		src := FromFunc(func(o *Observer[int]) {
			o.Next(1)
			o.Error(expected)
		})

		rec, o := observers.NewRecorder[int](t)
		_, err := src.Sub((*Observer[int])(o))

		assertlite.NoError(t, err)
		assertlite.EqualSlices(t, rec.Values(), []int{1})
		assertlite.IsErrorOf(t, rec.Err(), expected)
		assertlite.False(t, rec.Completed())
	})

	t.Run("does not auto-complete", func(t *testing.T) {
		src := FromFunc(func(o *Observer[int]) {
			o.Next(1)
			o.Next(2)
		})

		rec, o := observers.NewRecorder[int](t)
		_, _ = src.Sub((*Observer[int])(o))

		assertlite.EqualSlices(t, rec.Values(), []int{1, 2})
		assertlite.IsNil(t, rec.Err())
		assertlite.False(t, rec.Completed())
	})

	t.Run("cold behaviour", func(t *testing.T) {
		calls := 0

		src := FromFunc(func(o *Observer[int]) {
			calls++
			o.Next(calls)
			o.Complete()
		})

		r1, o1 := observers.NewRecorder[int](t)
		r2, o2 := observers.NewRecorder[int](t)

		_, _ = src.Sub((*Observer[int])(o1))
		_, _ = src.Sub((*Observer[int])(o2))

		assertlite.EqualSlices(t, r1.Values(), []int{1})
		assertlite.EqualSlices(t, r2.Values(), []int{2})
		assertlite.Equals(t, calls, 2)
	})

	t.Run("ignores signals after termination", func(t *testing.T) {
		ff := func(o *Observer[int]) {
			o.Error(errors.New("fail"))
			o.Next(1)
			o.Complete()
		}
		src := FromFunc(ff)
		rec, o := observers.NewRecorder[int](t)
		_, _ = src.Sub((*Observer[int])(o))

		assertlite.EqualSlices(t, rec.Values(), []int{})
		assertlite.False(t, rec.Completed())
	})
}
