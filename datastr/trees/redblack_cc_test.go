// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package trees

import (
	"sync"
	"testing"

	"github.com/andrerrcosta2/gtools/core/data/comparator/comparators"
	"github.com/andrerrcosta2/gtools/core/seeders/random"
	"github.com/andrerrcosta2/gtools/core/testlite"
)

// TestConcRedBlack_InsertConcurrently
func TestConcRedBlack_InsertConcurrently(t *testing.T) {
	cmp := comparators.Ordered[int]()
	counter := testlite.Map(t)

	// Test 1: random values inserted concurrently
	t.Run("random values inserted concurrently", func(t *testing.T) {
		tree := ConcRedBlack[int](cmp)
		ins := sync.Map{}
		random.Int(100, 10, 1000).
			Parallel(func(i, v int) {
				tree.Insert(v)
				if _, ok := ins.Load(v); !ok {
					ins.Store(v, true)
					counter.Add(1, "inserted")
				}
				if !tree.Contains(v) {
					t.Errorf("tree should contain '%d'", v)
				}
			}, 6)

		counter.AssertCountTo("inserted", tree.Size(), "tree size should be '%d', but got '%d",
			counter.Get("inserted"), tree.Size())
	})

}
