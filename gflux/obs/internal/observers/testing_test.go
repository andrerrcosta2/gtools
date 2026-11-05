// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package observers

import (
	"testing"
	"time"

	"github.com/andrerrcosta2/gtools/core/testlite/assertlite"
)

func TestRecorder_RecordsValues(t *testing.T) {
	r, o := NewRecorder[int](t)

	o.Next(1)
	o.Next(2)

	assertlite.EqualSlices(t, []int{1, 2}, r.Values())
}

func TestRecorder_CompleteSetsFlags(t *testing.T) {
	r, o := NewRecorder[int](t)

	o.Complete()

	assertlite.True(t, r.Completed())
	assertlite.True(t, r.Terminated())
	assertlite.NoError(t, r.Err())
}

func TestRecorder_ErrorSetsFlags(t *testing.T) {
	r, o := NewRecorder[int](t)

	o.Error(TestError)

	assertlite.False(t, r.Completed())
	assertlite.True(t, r.Terminated())
	assertlite.IsErrorOf(t, TestError, r.Err())
}

func TestRecorder_AwaitTermination(t *testing.T) {
	r, o := NewRecorder[int](t)

	go func() {
		time.Sleep(10 * time.Millisecond)
		o.Complete()
	}()

	r.AwaitTermination(1 * time.Second)

	assertlite.True(t, r.Completed())
}

func TestRecorder_AwaitNext(t *testing.T) {
	r, o := NewRecorder[int](t)

	go func() {
		time.Sleep(10 * time.Millisecond)
		o.Next(42)
	}()

	r.AwaitNext(1 * time.Second)

	assertlite.EqualSlices(t, []int{42}, r.Values())
}

func TestNewRecorderFor_PanicsOnNilCustom(t *testing.T) {
	assertlite.Panics(t, func() {
		NewRecorderFor[int](t, nil)
	})
}

func TestNewRecorderFor_ForwardsNext(t *testing.T) {
	customCalled := 0

	custom := &Observer[int]{
		Next: func(v int) {
			customCalled++
		},
	}

	r, o := NewRecorderFor[int](t, custom)

	o.Next(10)

	assertlite.True(t, 1 == customCalled)
	assertlite.EqualSlices(t, []int{10}, r.Values())
}
