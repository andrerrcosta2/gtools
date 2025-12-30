// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package maps

import (
	"github.com/andrerrcosta2/gtools/core/domain/functions"
	"github.com/andrerrcosta2/gtools/core/seeders/random"
	"github.com/andrerrcosta2/gtools/datastr/internal/tests"
	"github.com/andrerrcosta2/gtools/gtests/testingtools"
	"github.com/andrerrcosta2/gtools/gtests/testingtools/config/testlogs"
	"testing"
)

// TestSortableOfMap_PutAndGet tests the Put and Get methods of the sortableOfMap
func TestSortableOfMap_PutAndGet(t *testing.T) {
	// std
	tt := testingtools.LoggersLite(t, testlogs.OnFailure)
	// types
	type N = *tests.SortableNode

	// Create map
	m := SortableOf[N, N]()

	// Test Inserting
	random.Struct[N](20).
		// Insert
		Each(func(n N) { m.Put(n, n) }).

		// Test retrieving
		Each(func(n N) {
			shouldFindExactKey[N, N](tt, m, n, n, functions.Equality[N])
		})

	// Test retrieving a non-existent key
	shouldNotFindKey[N, N](tt, m, tests.NewSortableNode("non-existent"))

	tt.PrintLogStack()
}

// TestSortableOfMap_Delete tests the Delete method of the sortableOfMap
func TestSortableOfMap_Delete(t *testing.T) {
	// std
	tt := testingtools.LoggersLite(t, testlogs.OnFailure)
	// types
	type N = *tests.SortableNode

	// Create map
	m := SortableOf[N, N]()

	random.Struct[N](20).
		Each(func(n N) { m.Put(n, n) }).
		Some(5).
		// Delete 5 random keys
		Each(func(n N) {
			shouldDeleteKey[N, N](tt, m, n)
		})

	tt.PrintLogStack()
}

// TestSortableOfMap_Contains tests the Contains method of the sortableOfMap
func TestSortableOfMap_Contains(t *testing.T) {
	// std
	tt := testingtools.LoggersLite(t, testlogs.OnFailure)
	// types
	type N = *tests.SortableNode

	// Create map
	m := SortableOf[N, N]()

	random.Struct[N](20).Each(func(n N) { m.Put(n, n) }).
		// Test for existing key
		Each(func(n *tests.SortableNode) {
			shouldFindKey[N, N](tt, m, n)
		})

	// Test for non-existent key
	shouldNotFindKey[*tests.SortableNode, *tests.SortableNode](tt, m, tests.NewSortableNode("non-existent"))

	tt.PrintLogStack()
}

// TestSortableOfMap_LenAndClear tests the Len and Clear methods of the sortableOfMap
func TestSortableOfMap_LenAndClear(t *testing.T) {
	// std
	tt := testingtools.LoggersLite(t, testlogs.OnFailure)
	// types
	type N = *tests.SortableNode

	// Create map
	m := SortableOf[N, N]()

	random.Struct[N](20).
		Each(func(n N) { m.Put(n, n) })

	shouldHaveLength[N, N](tt, m, 20)

	m.Put(tests.NewSortableNode("non-existent"), tests.NewSortableNode("non-existent"))

	shouldHaveLength[N, N](tt, m, 21)

	m.Delete(tests.NewSortableNode("non-existent"))

	shouldHaveLength[N, N](tt, m, 20)

	m.Clear()

	shouldHaveLength[N, N](tt, m, 0)

	tt.PrintLogStack()
}

// TestSortableOfMap_Iterator tests the Iterator method of the sortableOfMap
func TestSortableOfMap_Iterator(t *testing.T) {
	// std
	tt := testingtools.LoggersLite(t, testlogs.OnFailure)
	// types
	type N = *tests.SortableNode

	// Create map
	sm := SortableOf[N, N]()
	mm := make(map[N]N)

	random.Struct[N](20).
		Each(func(n N) {
			sm.Put(n, n)
			mm[n] = n
		}).
		Values()

	shouldFindAllUsingIteratorBy(tt, sm.Iterator(), mm, functions.Equality[N])

	tt.PrintLogStack()
}
