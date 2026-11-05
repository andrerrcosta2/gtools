// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package dispatchers

import (
	"fmt"
	"testing"

	"github.com/andrerrcosta2/gtools/core/testlite"
	"github.com/andrerrcosta2/gtools/core/testlite/assertlite"
)

func TestHotDispatcher_Creation(t *testing.T) {
	// Case 1: Create a single-routine dispatcher
	t.Run("Create a single-routine dispatcher", func(t *testing.T) {
		d := Hot[int](func(i int) error { return nil }, func(i int, err error) {}, 1)
		cast, ok := d.(*hotDispatcher[int])
		assertlite.True(t, ok, "expected dispatcher to be of type hotDispatcher")
		assertlite.True(t, d != nil, "expected dispatcher to be created")
		assertlite.True(t, !d.IsClosed(), "expected dispatcher to not be closed")
		assertlite.True(t, cast.maxWorkers == 1, "expected dispatcher to have 1 worker")
	})

	// Case 2: Create a multi-routine dispatcher
	t.Run("Create a multi-routine dispatcher", func(t *testing.T) {
		d := Hot[int](func(i int) error { return nil }, func(i int, err error) {}, 10)
		cast, ok := d.(*hotDispatcher[int])
		assertlite.True(t, ok, "expected dispatcher to be of type hotDispatcher")
		assertlite.True(t, d != nil, "expected dispatcher to be created")
		assertlite.True(t, !d.IsClosed(), "expected dispatcher to not be closed")
		assertlite.True(t, cast.maxWorkers == 10, "expected dispatcher to have 10 workers")
	})
}

func TestHotDispatcher_Dispatch(t *testing.T) {
	// Test 1: Single thread, shouldn't fail
	t.Run("single thread shouldn't fail", func(t *testing.T) {
		d := Hot[int](func(_ int) error { return nil }, func(_ int, _ error) {}, 1)
		for i := 0; i < 100; i++ {
			err := d.Dispatch(i)
			assertlite.True(t, err == nil, "expected dispatcher to not fail")
		}
	})

	// Test 2: Single thread, should fail on 30
	t.Run("single thread, should fail on 30", func(t *testing.T) {
		counter := testlite.Map(t)
		handler := func(i int) error {
			fmt.Printf("i: %d\n", i)
			if i == 30 {
				fmt.Printf("failing on %d", i)
				return fmt.Errorf("fail on 30")
			}
			return nil
		}

		listener := func(i int, err error) {
			if err != nil {
				counter.Add(i, err.Error())
			} else {
				counter.Add(i, "ok")
			}
		}

		d := Hot[int](handler, listener, 1)

		for i := 0; i < 100; i++ {
			err := d.Dispatch(i)
			assertlite.True(t, err == nil, "expected dispatcher to not fail on queuing item but got %v", err)
		}

		d.Close()
		assertlite.True(t, d.IsClosed(), "expected dispatcher to be closed")
		f30calls := counter.Get("fail on 30")
		assertlite.True(t, 30 == f30calls)
	})
}
