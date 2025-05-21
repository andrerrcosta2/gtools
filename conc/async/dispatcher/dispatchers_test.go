// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package dispatcher

import (
	"github.com/andrerrcosta2/gtools/core/format/fmx"
	"github.com/andrerrcosta2/gtools/gtests/testingtools"
	"github.com/andrerrcosta2/gtools/gtests/testingtools/config/testlogs"
	"testing"
)

//var handler = func(i int) error {
//	time.Sleep(10 * time.Millisecond)
//	tt.Flag(true, comparators.StringHash(i))
//	if i == 50 { return fmt.Errorf("fail on 50")	}
//	return nil
//}

func TestHotDispatcher_Creation(t *testing.T) {
	tt := testingtools.ConcLite(t, testlogs.OnFailure)

	// Case 1: Create a single-routine dispatcher
	t.Run("Create a single-routine dispatcher", func(t *testing.T) {
		d := Hot[int](func(i int) error { return nil }, func(i int, err error) {}, 1)
		cast, ok := d.(*hotDispatcher[int])
		tt.Condition(ok, "expected dispatcher to be of type hotDispatcher")
		tt.Condition(d != nil, "expected dispatcher to be created")
		tt.Condition(!d.IsClosed(), "expected dispatcher to not be closed")
		tt.Condition(cast.maxWorkers == 1, "expected dispatcher to have 1 worker")
	})

	// Case 2: Create a multi-routine dispatcher
	t.Run("Create a multi-routine dispatcher", func(t *testing.T) {
		d := Hot[int](func(i int) error { return nil }, func(i int, err error) {}, 10)
		cast, ok := d.(*hotDispatcher[int])
		tt.Condition(ok, "expected dispatcher to be of type hotDispatcher")
		tt.Condition(d != nil, "expected dispatcher to be created")
		tt.Condition(!d.IsClosed(), "expected dispatcher to not be closed")
		tt.Condition(cast.maxWorkers == 10, "expected dispatcher to have 10 workers")
	})
}

func TestHotDispatcher_Dispatch(t *testing.T) {
	tt := testingtools.ConcLite(t, testlogs.OnFailure)

	// Test 1: Single thread, shouldn't fail
	t.Run("single thread shouldn't fail", func(t *testing.T) {
		d := Hot[int](func(_ int) error { return nil }, func(_ int, _ error) {}, 1)
		for i := 0; i < 100; i++ {
			err := d.Dispatch(i)
			tt.Condition(err == nil, "expected dispatcher to not fail")
		}
	})

	// Test 2: Single thread, should fail on 30
	t.Run("single thread, should fail on 30", func(t *testing.T) {
		handler := func(i int) error {
			fmx.Printf("i: %d\n", i)
			if i == 30 {
				fmx.Redf("failing on %d", i)
				return fmx.Errorf("fail on 30")
			}
			return nil
		}

		listener := func(i int, err error) {
			if err != nil {
				tt.RegisterCalls(i, err.Error())
			} else {
				tt.RegisterCalls(i, "ok")
			}
		}

		d := Hot[int](handler, listener, 1)

		for i := 0; i < 100; i++ {
			err := d.Dispatch(i)
			tt.Condition(err == nil, "expected dispatcher to not fail on queuing item but got %v", err)
		}

		d.Close()
		tt.Condition(d.IsClosed(), "expected dispatcher to be closed")
		tt.AssertCalls(30, "fail on 30")
	})
}
