// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package trees

import (
	"github.com/andrerrcosta2/gtools/core/data/comparators"
	"github.com/andrerrcosta2/gtools/core/seeders/random"
	"github.com/andrerrcosta2/gtools/gtests/testingtools"
	"github.com/andrerrcosta2/gtools/gtests/testingtools/config/testlogs"
	"sync"
	"testing"
)

// TestConcRedBlack_InsertConcurrently
func TestConcRedBlack_InsertConcurrently(t *testing.T) {
	tt := testingtools.ConcLite(t, testlogs.OnFailure)
	cmp := comparators.Ordered[int]{}

	// Test 1: random values inserted concurrently
	t.Run("random values inserted concurrently", func(t *testing.T) {
		tree := ConcRedBlack[int](cmp)
		ins := sync.Map{}
		random.Int(100, 10, 1000).
			Parallel(func(i, v int) {
				tree.Insert(v)
				if _, ok := ins.Load(v); !ok {
					ins.Store(v, true)
					tt.RegisterCalls(1, "inserted")
				}
				if !tree.Contains(v) {
					t.Errorf("tree should contain '%d'", v)
				}
			}, 6)

		tt.AssertCallsTo("inserted", tree.Size(), "tree size should be '%d', but got '%d",
			tt.CallsTo("inserted"), tree.Size())
	})

}
