// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package dicts

import (
	"fmt"
	"testing"

	"github.com/andrerrcosta2/gtools/core/data/comparator/comparators"
	"github.com/andrerrcosta2/gtools/core/data/keyer/keyers"
	"github.com/andrerrcosta2/gtools/core/domain/function/functions"
	"github.com/andrerrcosta2/gtools/core/seeders/random"
	"github.com/andrerrcosta2/gtools/datastr/internal/tests"
)

func TestComparableOfMap_Constructor(t *testing.T) {
	type N = *tests.SortableNode

	n1 := tests.NewSortableNode("a")
	n2 := tests.NewSortableNode("b")

	m := ComparableOf(
		NewAnyEntry(n1, n1),
		NewAnyEntry(n2, n2),
	)

	shouldFindExactKey(t, m, n1, n1, functions.ComparableEquality[N])
	shouldFindExactKey(t, m, n2, n2, functions.ComparableEquality[N])
	shouldHaveLength(t, m, 2)
}

// TestComparableOfMap_PutAndGet tests the Put and Get methods of the sortableOfMap
func TestComparableOfMap_PutAndGet(t *testing.T) {
	// std
	// types
	type N = *tests.SortableNode

	// Create map
	m := ComparableOf[N, N]()

	// Test Inserting
	random.Struct[N](20).
		// Insert
		Each(func(n N) { m.Put(n, n) }).

		// Test retrieving
		Each(func(n N) {
			shouldFindExactKey(t, m, n, n, functions.ComparableEquality[N])
		})

	// Test retrieving a non-existent key
	shouldNotFindKey(t, m, tests.NewSortableNode("non-existent"))

}

func TestComparableOfMap_KeysAndValues(t *testing.T) {
	type N = *tests.SortableNode

	m := ComparableOf[N, N]()
	expected := make(map[string]N)

	random.Struct[N](10).Each(func(n N) {
		m.Put(n, n)
		expected[n.ID()] = n
	})

	keys := m.Keys()
	if len(keys) != len(expected) {
		t.Errorf("Expected %d keys, got %d", len(expected), len(keys))
	}

	for _, k := range keys {
		if _, ok := expected[k.ID()]; !ok {
			t.Errorf("Unexpected key: %v", k)
		}
	}

	values := m.Values()
	if len(values) != len(expected) {
		t.Errorf("Expected %d values, got %d", len(expected), len(values))
	}
}

// TestComparableOfMap_Delete tests the Delete method of the sortableOfMap
func TestComparableOfMap_Delete(t *testing.T) {
	// std
	// types
	type N = *tests.SortableNode

	// Create map
	m := ComparableOf[N, N]()

	random.Struct[N](20).
		Each(func(n N) { m.Put(n, n) }).
		Some(5).
		// Delete 5 random keys
		Each(func(n N) {
			shouldDeleteKey(t, m, n)
		})

}

// TestComparableOfMap_Contains tests the Contains method of the sortableOfMap
func TestComparableOfMap_Contains(t *testing.T) {
	// std
	// types
	type N = *tests.SortableNode

	// Create map
	m := ComparableOf[N, N]()

	random.Struct[N](20).Each(func(n N) { m.Put(n, n) }).
		// Test for existing key
		Each(func(n *tests.SortableNode) {
			shouldFindKey(t, m, n)
		})

	// Test for non-existent key
	shouldNotFindKey(t, m, tests.NewSortableNode("non-existent"))

}

// TestComparableOfMap_LenAndClear tests the Len and Clear methods of the sortableOfMap
func TestComparableOfMap_LenAndClear(t *testing.T) {
	// std
	// types
	type N = *tests.SortableNode

	// Create map
	m := ComparableOf[N, N]()

	random.Struct[N](20).
		Each(func(n N) { m.Put(n, n) })

	shouldHaveLength(t, m, 20)

	m.Put(tests.NewSortableNode("non-existent"), tests.NewSortableNode("non-existent"))

	shouldHaveLength(t, m, 21)

	m.Delete(tests.NewSortableNode("non-existent"))

	shouldHaveLength(t, m, 20)

	m.Clear()

	shouldHaveLength(t, m, 0)

}

// TestComparableOfMap_Iterator tests the Iterator method of the sortableOfMap
func TestComparableOfMap_Iterator(t *testing.T) {
	// std
	// types
	type N = *tests.SortableNode

	// Create map
	sm := ComparableOf[N, N]()
	mm := make(map[N]N)

	random.Struct[N](20).
		EachN(func(i int, n N) {
			fmt.Printf("%d) %v\n", i, n)
			sm.Put(n, n)
			mm[n] = n
		}).
		Values()

	shouldFindAllUsingIteratorBy(t, sm.Iterator(), mm, functions.ComparableEquality[N])
}

func TestComparableOfMap_SortedIterator(t *testing.T) {
	type N = *tests.SortableNode

	m := ComparableOf[N, N]()

	// Insert in non-sorted order
	keys := []*tests.SortableNode{
		tests.NewSortableNode("c"),
		tests.NewSortableNode("a"),
		tests.NewSortableNode("b"),
	}

	for _, k := range keys {
		m.Put(k, k)
	}

	// Get sorted iterator
	comparator := comparators.KeySortableOf(keyers.FingerPrint[N]())
	it := m.Iterator(comparator)

	var results []string
	for k, _, hn := it.Next(); hn; k, _, hn = it.Next() {
		results = append(results, k.ID())
	}

	expected := []string{"a", "b", "c"}
	for i, r := range results {
		if r != expected[i] {
			t.Errorf("Expected %v at position %d, got %v", expected[i], i, r)
		}
	}
}

// Test empty map behavior
func TestComparableOfMap_EmptyMap(t *testing.T) {
	type N = *tests.SortableNode

	m := ComparableOf[N, N]()

	shouldHaveLength(t, m, 0)

	// Get from empty map
	val, ok := m.Get(tests.NewSortableNode("nonexistent"))
	if ok {
		t.Error("Expected no value from empty map")
	}
	if val != nil {
		t.Error("Expected nil value from empty map")
	}

	// Delete from empty map (should not panic)
	m.Delete(tests.NewSortableNode("nonexistent"))

	// Iterator on empty map
	count := 0
	for _, _, hn := m.Iterator().Next(); hn; _, _, hn = m.Iterator().Next() {
		count++
	}
	if count != 0 {
		t.Errorf("Expected empty iterator, got %d items", count)
	}

	// Keys/Values on empty map
	if len(m.Keys()) != 0 {
		t.Error("Expected empty keys")
	}
	if len(m.Values()) != 0 {
		t.Error("Expected empty values")
	}
}

// Test nil value handling
func TestComparableOfMap_NilValues(t *testing.T) {
	type N = *tests.SortableNode

	m := ComparableOf[N, N]()
	key := tests.NewSortableNode("key")

	// Insert nil value
	m.Put(key, nil)

	val, ok := m.Get(key)
	if !ok {
		t.Error("Expected to find key with nil value")
	}
	if val != nil {
		t.Errorf("Expected nil value, got %v", val)
	}

	// Overwrite nil with actual value
	m.Put(key, key)
	val, ok = m.Get(key)
	if !ok || val != key {
		t.Error("Expected updated value")
	}
}

func TestComparableOfMap_DuplicateKeys(t *testing.T) {
	type N = *tests.SortableNode

	m := ComparableOf[N, N]()
	key := tests.NewSortableNode("key")
	val1 := tests.NewSortableNode("val1")
	val2 := tests.NewSortableNode("val2")

	m.Put(key, val1)
	m.Put(key, val2) // Should overwrite

	shouldFindExactKey(t, m, key, val2, functions.ComparableEquality[N])
	shouldHaveLength(t, m, 1)
}
