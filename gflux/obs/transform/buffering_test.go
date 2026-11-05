// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package transform

import (
	"errors"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/andrerrcosta2/gtools/core/testlite/assertlite"
	"github.com/andrerrcosta2/gtools/gflux/obs"
	"github.com/andrerrcosta2/gtools/gflux/obs/internal/observers"
)

// TestBuffer_Sync
//
// Source emits synchronously.
//
// We should assert:
//  1. Values are grouped into fixed-size buffers
//  2. Final partial buffer is flushed on completion
//  3. Empty sources emit no buffers
//  4. Errors are propagated and incomplete buffers are discarded
//  5. Invalid buffer size panics
func TestBuffer_Sync(t *testing.T) {
	t.Run("buffers values into exact-sized slices", func(t *testing.T) {
		src := obs.Of(1, 2, 3, 4, 5, 6)

		// size = 3: two full buffers
		buffered := Buffer[int](3)(src)

		rec, o := observers.NewRecorder[[]int](t)
		_, err := buffered.Sub((*obs.Observer[[]int])(o))

		assertlite.NoError(t, err)
		assertlite.EqualSlices(t, rec.Values(), [][]int{
			{1, 2, 3},
			{4, 5, 6},
		})
		assertlite.True(t, rec.Completed())
		assertlite.IsNil(t, rec.Err())
	})

	t.Run("emits final partial buffer on completion", func(t *testing.T) {
		src := obs.Of(1, 2, 3, 4, 5)

		// size = 3 one full, one partial
		buffered := Buffer[int](3)(src)

		rec, o := observers.NewRecorder[[]int](t)
		_, _ = buffered.Sub((*obs.Observer[[]int])(o))

		assertlite.EqualSlices(t, rec.Values(), [][]int{
			{1, 2, 3},
			{4, 5},
		})
		assertlite.True(t, rec.Completed())
		assertlite.IsNil(t, rec.Err())
	})

	t.Run("does not emit anything for empty source", func(t *testing.T) {
		src := obs.Of[int]()

		// size is irrelevant here, but must be > 0
		buffered := Buffer[int](3)(src)

		rec, o := observers.NewRecorder[[]int](t)
		_, _ = buffered.Sub((*obs.Observer[[]int])(o))

		assertlite.EqualSlices(t, rec.Values(), nil)
		assertlite.True(t, rec.Completed())
		assertlite.IsNil(t, rec.Err())
	})

	t.Run("buffer size 1 emits each value as its own slice", func(t *testing.T) {
		src := obs.Of(10, 20, 30)

		// size = 1: every value flushes immediately
		buffered := Buffer[int](1)(src)

		rec, o := observers.NewRecorder[[]int](t)
		_, _ = buffered.Sub((*obs.Observer[[]int])(o))

		assertlite.EqualSlices(t, rec.Values(), [][]int{
			{10},
			{20},
			{30},
		})
		assertlite.True(t, rec.Completed())
	})

	t.Run("propagates error and discards incomplete buffer", func(t *testing.T) {
		src := obs.FromFunc(func(o *obs.Observer[int]) {
			o.Next(1)
			o.Next(2)
			o.Error(errors.New("failure"))
		})

		// size = 3: buffer never fills
		buffered := Buffer[int](3)(src)

		rec, o := observers.NewRecorder[[]int](t)
		_, _ = buffered.Sub((*obs.Observer[[]int])(o))

		assertlite.EqualSlices(t, rec.Values(), nil)
		assertlite.False(t, rec.Completed())
		assertlite.NotNil(t, rec.Err())
	})

	t.Run("emits full buffers before propagating error", func(t *testing.T) {
		src := obs.FromFunc(func(o *obs.Observer[int]) {
			o.Next(1)
			o.Next(2)
			o.Next(3) // fills buffer
			o.Next(4)
			o.Error(errors.New("failure"))
		})

		// size = 3: first buffer emitted, second discarded
		buffered := Buffer[int](3)(src)

		rec, o := observers.NewRecorder[[]int](t)
		_, _ = buffered.Sub((*obs.Observer[[]int])(o))

		assertlite.EqualSlices(t, rec.Values(), [][]int{
			{1, 2, 3},
		})
		assertlite.False(t, rec.Completed())
		assertlite.NotNil(t, rec.Err())
	})

	t.Run("completes synchronously", func(t *testing.T) {
		src := obs.Of(1, 2)

		// size = 2 exactly one buffer
		buffered := Buffer[int](2)(src)

		done := false
		_, o := observers.NewRecorder[[]int](t)
		o.Complete = func() {
			done = true
		}

		_, _ = buffered.Sub((*obs.Observer[[]int])(o))

		assertlite.True(t, done)
	})

	t.Run("panics when buffer size is non-positive", func(t *testing.T) {
		assertlite.Panics(t, func() {
			_ = Buffer[any](-1)
		})
	})

	t.Run("fuzz sizes", func(t *testing.T) {
		values := []int{1, 2, 3, 4, 5}

		for size := 1; size <= 10; size++ {
			t.Run(fmt.Sprintf("size=%d", size), func(t *testing.T) {
				src := obs.Of(values...)

				buffered := Buffer[int](size)(src)

				rec, o := observers.NewRecorder[[]int](t)
				_, _ = buffered.Sub((*obs.Observer[[]int])(o))

				var flat []int
				for _, b := range rec.Values() {
					flat = append(flat, b...)
				}

				assertlite.EqualSlices(t, flat, values)
				assertlite.True(t, rec.Completed())
			})
		}
	})
}

// TestBuffer_Async
//
// Source emits asynchronously.
//
// We should assert:
//  1. Values are grouped into fixed-size buffers
//  2. Final partial buffer is flushed on completion
//  3. Empty sources emit no buffers
//  4. Errors are propagated and incomplete buffers are discarded
//  5. Concurrent emissions remain safe and all values are accounted for
func TestBuffer_Async(t *testing.T) {
	t.Run("buffers async values into exact-sized slices", func(t *testing.T) {
		src := obs.AsyncOf(1, 2, 3, 4, 5, 6)

		// size = 3: two full buffers
		buffered := Buffer[int](3)(src)

		rec, o := observers.NewRecorder[[]int](t)

		sub, err := buffered.Sub((*obs.Observer[[]int])(o))
		assertlite.NoError(t, err)
		assertlite.NotNil(t, sub)

		rec.AwaitTermination(100 * time.Millisecond)

		assertlite.EqualSlices(t, rec.Values(), [][]int{
			{1, 2, 3},
			{4, 5, 6},
		})
		assertlite.True(t, rec.Completed())
		assertlite.IsNil(t, rec.Err())
	})

	t.Run("emits final partial buffer on async completion", func(t *testing.T) {
		src := obs.AsyncOf(1, 2, 3, 4, 5)

		// size = 3 one full, one partial
		buffered := Buffer[int](3)(src)

		rec, o := observers.NewRecorder[[]int](t)
		_, _ = buffered.Sub((*obs.Observer[[]int])(o))

		rec.AwaitTermination(100 * time.Millisecond)

		assertlite.EqualSlices(t, rec.Values(), [][]int{
			{1, 2, 3},
			{4, 5},
		})
		assertlite.True(t, rec.Completed())
		assertlite.IsNil(t, rec.Err())
	})

	t.Run("empty async source emits nothing and completes", func(t *testing.T) {
		src := obs.AsyncOf[int]()

		buffered := Buffer[int](3)(src)

		rec, o := observers.NewRecorder[[]int](t)
		_, _ = buffered.Sub((*obs.Observer[[]int])(o))

		rec.AwaitTermination(100 * time.Millisecond)

		assertlite.EqualSlices(t, rec.Values(), nil)
		assertlite.True(t, rec.Completed())
		assertlite.IsNil(t, rec.Err())
	})

	t.Run("buffer size 1 flushes on every async value", func(t *testing.T) {
		src := obs.AsyncOf(10, 20, 30)

		buffered := Buffer[int](1)(src)

		rec, o := observers.NewRecorder[[]int](t)
		_, _ = buffered.Sub((*obs.Observer[[]int])(o))

		rec.AwaitTermination(100 * time.Millisecond)

		assertlite.EqualSlices(t, rec.Values(), [][]int{
			{10},
			{20},
			{30},
		})
		assertlite.True(t, rec.Completed())
	})

	t.Run("propagates async error and discards incomplete buffer", func(t *testing.T) {
		failure := errors.New("failure")

		src := obs.Async[int](func(o *obs.Observer[int]) {
			o.Next(1)
			o.Next(2)
			o.Error(failure)
		})

		// size = 3 never fills
		buffered := Buffer[int](3)(src)

		rec, o := observers.NewRecorder[[]int](t)
		_, _ = buffered.Sub((*obs.Observer[[]int])(o))

		rec.AwaitTermination(100 * time.Millisecond)

		assertlite.EqualSlices(t, rec.Values(), nil)
		assertlite.False(t, rec.Completed())
		assertlite.IsErrorOf(t, rec.Err(), failure)
	})

	t.Run("emits full buffers before async error", func(t *testing.T) {
		failure := errors.New("failure")

		src := obs.Async[int](func(o *obs.Observer[int]) {
			o.Next(1)
			o.Next(2)
			o.Next(3) // fills buffer
			o.Next(4)
			o.Error(failure)
		})

		// size = 3
		buffered := Buffer[int](3)(src)

		rec, o := observers.NewRecorder[[]int](t)
		_, _ = buffered.Sub((*obs.Observer[[]int])(o))

		rec.AwaitTermination(100 * time.Millisecond)

		assertlite.EqualSlices(t, rec.Values(), [][]int{
			{1, 2, 3},
		})
		assertlite.False(t, rec.Completed())
		assertlite.IsErrorOf(t, rec.Err(), failure)
	})

	t.Run("concurrent source", func(t *testing.T) {
		src := obs.Async[int](func(o *obs.Observer[int]) {
			var wg sync.WaitGroup

			for g := 0; g < 4; g++ {
				wg.Add(1)
				go func(base int) {
					defer wg.Done()
					for i := 0; i < 25; i++ {
						o.Next(base*100 + i)
					}
				}(g)
			}

			wg.Wait()
			o.Complete()
		})

		buffered := Buffer[int](10)(src)

		rec, o := observers.NewRecorder[[]int](t)
		_, _ = buffered.Sub((*obs.Observer[[]int])(o))

		rec.AwaitTermination(200 * time.Millisecond)

		var count int
		for _, b := range rec.Values() {
			count += len(b)
		}

		assertlite.Equals(t, 100, count)
	})
}

// TestPairwise_Sync
//
// Source emits synchronously.
//
// We should assert:
//  1. Consecutive pairs are emitted in order
//  2. Sources with fewer than two values emit nothing
//  3. Completion is forwarded synchronously
func TestPairwise_Sync(t *testing.T) {
	t.Run("pairs consecutive values", func(t *testing.T) {
		src := obs.Of(1, 2, 3, 4)

		pairwise := Pairwise[int]()(src)

		rec, o := observers.NewRecorder[[2]int](t)
		_, _ = pairwise.Sub((*obs.Observer[[2]int])(o))

		assertlite.EqualSlices(t, [][2]int{
			{1, 2},
			{2, 3},
			{3, 4},
		}, rec.Values())

		assertlite.True(t, rec.Completed())
	})

	t.Run("emits nothing for single value", func(t *testing.T) {
		src := obs.Of(42)

		pairwise := Pairwise[int]()(src)

		rec, o := observers.NewRecorder[[2]int](t)
		_, _ = pairwise.Sub((*obs.Observer[[2]int])(o))

		assertlite.Equals(t, 0, len(rec.Values()))
		assertlite.True(t, rec.Completed())
	})

	t.Run("emits nothing for empty source", func(t *testing.T) {
		src := obs.Of[int]()

		pairwise := Pairwise[int]()(src)

		rec, o := observers.NewRecorder[[2]int](t)
		_, _ = pairwise.Sub((*obs.Observer[[2]int])(o))

		assertlite.Equals(t, 0, len(rec.Values()))
		assertlite.True(t, rec.Completed())
	})

	t.Run("completes synchronously", func(t *testing.T) {
		src := obs.Of(1, 2)

		pairwise := Pairwise[int]()(src)

		done := false
		o := &obs.Observer[[2]int]{
			Complete: func() {
				done = true
			},
		}

		_, _ = pairwise.Sub(o)

		assertlite.True(t, done)
	})
}

// TestPairwise_Async
//
// Source emits asynchronously.
//
// We should assert:
//  1. Consecutive pairs are emitted in order
//  2. Sources with fewer than two values emit nothing
//  3. Completion is forwarded after async termination
func TestPairwise_Async(t *testing.T) {

	t.Run("pairs consecutive values", func(t *testing.T) {
		src := obs.AsyncOf(1, 2, 3, 4)

		pairwise := Pairwise[int]()(src)

		rec, o := observers.NewRecorder[[2]int](t)
		_, _ = pairwise.Sub((*obs.Observer[[2]int])(o))

		rec.AwaitTermination(100 * time.Millisecond)

		assertlite.EqualSlices(t, [][2]int{
			{1, 2},
			{2, 3},
			{3, 4},
		}, rec.Values())

		assertlite.True(t, rec.Completed())
	})

	t.Run("emits nothing for single value", func(t *testing.T) {
		src := obs.AsyncOf(42)

		pairwise := Pairwise[int]()(src)

		rec, o := observers.NewRecorder[[2]int](t)
		_, _ = pairwise.Sub((*obs.Observer[[2]int])(o))

		rec.AwaitTermination(100 * time.Millisecond)

		assertlite.Equals(t, 0, len(rec.Values()))
		assertlite.True(t, rec.Completed())
	})

	t.Run("emits nothing for empty source", func(t *testing.T) {
		src := obs.AsyncOf[int]()

		pairwise := Pairwise[int]()(src)

		rec, o := observers.NewRecorder[[2]int](t)
		_, _ = pairwise.Sub((*obs.Observer[[2]int])(o))

		rec.AwaitTermination(100 * time.Millisecond)

		assertlite.Equals(t, 0, len(rec.Values()))
		assertlite.True(t, rec.Completed())
	})

	t.Run("completes asynchronously", func(t *testing.T) {
		src := obs.AsyncOf(1, 2)

		pairwise := Pairwise[int]()(src)

		rec, o := observers.NewRecorder[[2]int](t)
		_, _ = pairwise.Sub((*obs.Observer[[2]int])(o))

		rec.AwaitTermination(100 * time.Millisecond)

		assertlite.True(t, rec.Completed())
	})
}

// TestWindow_Sync
//
// Source emits synchronously.
//
// We should assert:
//  1. Source is split into windows of the configured size
//  2. Final partial window is emitted on completion
//  3. Empty source emits no windows
//  4. Window and outer stream errors are propagated consistently
func TestWindow_Sync(t *testing.T) {
	t.Run("splits source into windows of exact size", func(t *testing.T) {
		const size = 3

		src := obs.Of(1, 2, 3, 4, 5, 6)

		windowed := Window[int](size)(src)

		rec, o := observers.NewRecorder[obs.Observable[int]](t)
		_, err := windowed.Sub((*obs.Observer[obs.Observable[int]])(o))
		assertlite.NoError(t, err)

		wins := rec.Values()
		assertlite.Equals(t, 2, len(wins))

		for i, w := range wins {
			r, ro := observers.NewRecorder[int](t)
			_, _ = w.Sub((*obs.Observer[int])(ro))

			assertlite.Equals(t, size, len(r.Values()), "window %d size mismatch", i)
			assertlite.True(t, r.Completed())
		}

		assertlite.EqualSlices(t, []int{1, 2, 3}, winsValues(t, wins[0]))
		assertlite.EqualSlices(t, []int{4, 5, 6}, winsValues(t, wins[1]))

		assertlite.True(t, rec.Completed())
	})

	t.Run("emits final partial window smaller than size", func(t *testing.T) {
		const size = 3

		src := obs.Of(1, 2, 3, 4, 5)

		windowed := Window[int](size)(src)

		rec, o := observers.NewRecorder[obs.Observable[int]](t)
		_, _ = windowed.Sub((*obs.Observer[obs.Observable[int]])(o))

		wins := rec.Values()
		assertlite.Equals(t, 2, len(wins))

		assertlite.Equals(t, size, len(winsValues(t, wins[0])))
		assertlite.Equals(t, 2, len(winsValues(t, wins[1])))

		assertlite.EqualSlices(t, []int{1, 2, 3}, winsValues(t, wins[0]))
		assertlite.EqualSlices(t, []int{4, 5}, winsValues(t, wins[1]))

		assertlite.True(t, rec.Completed())
	})

	t.Run("single value creates one window of size 1", func(t *testing.T) {
		const size = 5

		src := obs.Of(42)

		windowed := Window[int](size)(src)

		rec, o := observers.NewRecorder[obs.Observable[int]](t)
		_, _ = windowed.Sub((*obs.Observer[obs.Observable[int]])(o))

		assertlite.Equals(t, 1, len(rec.Values()))

		assertlite.EqualSlices(t, []int{42}, winsValues(t, rec.Values()[0]))
	})

	t.Run("empty source emits no windows", func(t *testing.T) {
		src := obs.Of[int]()

		windowed := Window[int](5)(src)

		rec, o := observers.NewRecorder[obs.Observable[int]](t)
		_, _ = windowed.Sub((*obs.Observer[obs.Observable[int]])(o))

		assertlite.Equals(t, 0, len(rec.Values()))
		assertlite.True(t, rec.Completed())
	})

	t.Run("error closes current window and propagates", func(t *testing.T) {
		const size = 3
		failure := errors.New("failure")

		src := obs.FromFunc(func(o *obs.Observer[int]) {
			o.Next(1)
			o.Next(2)
			o.Error(failure)
		})

		windowed := Window[int](size)(src)

		rec, o := observers.NewRecorder[obs.Observable[int]](t)
		_, _ = windowed.Sub((*obs.Observer[obs.Observable[int]])(o))

		wins := rec.Values()
		assertlite.Equals(t, 1, len(wins))

		r, ro := observers.NewRecorder[int](t)
		_, _ = wins[0].Sub((*obs.Observer[int])(ro))

		assertlite.EqualSlices(t, []int{1, 2}, r.Values())
		assertlite.IsErrorOf(t, r.Err(), failure)

		assertlite.IsErrorOf(t, rec.Err(), failure)
	})
}

func winsValues[T any](t *testing.T, w obs.Observable[T]) []T {
	t.Helper()
	r, ro := observers.NewRecorder[T](t)
	_, _ = w.Sub((*obs.Observer[T])(ro))
	return r.Values()
}

// TestWindow_Async
//
// Source emits asynchronously.
//
// We should assert:
//  1. Source is split into windows of the configured size
//  2. Final partial window is emitted on completion
//  3. Empty source emits no windows
//  4. Errors are propagated to both window and outer stream
func TestWindow_Async(t *testing.T) {

	t.Run("windows async source correctly", func(t *testing.T) {
		src := obs.AsyncOf(1, 2, 3, 4, 5, 6)

		size := 3
		windowed := Window[int](size)(src)

		rec, o := observers.NewRecorder[obs.Observable[int]](t)
		_, _ = windowed.Sub((*obs.Observer[obs.Observable[int]])(o))

		rec.AwaitTermination(200 * time.Millisecond)

		assertlite.Equals(t, 2, len(rec.Values()))
		assertlite.EqualSlices(t, []int{1, 2, 3}, winsValues(t, rec.Values()[0]))
		assertlite.EqualSlices(t, []int{4, 5, 6}, winsValues(t, rec.Values()[1]))
		assertlite.True(t, rec.Completed())
	})

	t.Run("emits final partial window asynchronously", func(t *testing.T) {
		src := obs.AsyncOf(1, 2, 3, 4, 5)

		size := 3
		windowed := Window[int](size)(src)

		rec, o := observers.NewRecorder[obs.Observable[int]](t)
		_, _ = windowed.Sub((*obs.Observer[obs.Observable[int]])(o))

		rec.AwaitTermination(200 * time.Millisecond)

		assertlite.Equals(t, 2, len(rec.Values()))
		assertlite.EqualSlices(t, []int{4, 5}, winsValues(t, rec.Values()[1]))
	})

	t.Run("empty async source emits no windows", func(t *testing.T) {
		src := obs.AsyncOf[int]()

		size := 3
		windowed := Window[int](size)(src)

		rec, o := observers.NewRecorder[obs.Observable[int]](t)
		_, _ = windowed.Sub((*obs.Observer[obs.Observable[int]])(o))

		rec.AwaitTermination(200 * time.Millisecond)

		assertlite.Equals(t, 0, len(rec.Values()))
		assertlite.True(t, rec.Completed())
	})

	t.Run("propagates error asynchronously", func(t *testing.T) {
		failure := errors.New("failure")

		src := obs.Async[int](func(o *obs.Observer[int]) {
			o.Next(1)
			o.Error(failure)
		})

		size := 3
		windowed := Window[int](size)(src)

		rec, o := observers.NewRecorder[obs.Observable[int]](t)
		_, _ = windowed.Sub((*obs.Observer[obs.Observable[int]])(o))

		rec.AwaitTermination(200 * time.Millisecond)

		assertlite.Equals(t, 1, len(rec.Values()))

		r, ro := observers.NewRecorder[int](t)
		_, _ = rec.Values()[0].Sub((*obs.Observer[int])(ro))

		assertlite.IsErrorOf(t, r.Err(), failure)
		assertlite.IsErrorOf(t, rec.Err(), failure)
	})
}
