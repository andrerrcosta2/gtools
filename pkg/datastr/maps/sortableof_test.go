// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package maps

import (
	"github.com/andrerrcosta2/gtools/pkg/arrays"
	"github.com/andrerrcosta2/gtools/pkg/datastr/iterables"
	"github.com/andrerrcosta2/gtools/pkg/testdata/testsortables"
	"testing"
)

func TestSortableOfMap_PutAndGet(t *testing.T) {
	mapa := SortableOf[testsortables.TestNode, testsortables.TestNode]()

	// Test Inserting
	sortables := testsortables.RandomTestNodes(20, "node").Each(func(n testsortables.TestNode) {
		mapa.Put(n, n)
	})

	// Test retrieving
	sortables.Each(func(n testsortables.TestNode) {
		if got, ok := mapa.Get(n); !ok || got != n {
			t.Errorf("Expected %v, got %v", n, got)
		}
	})

	// Test retrieving a non-existent key
	nonExistentKey := testsortables.TestNode("non-existent")
	if _, ok := mapa.Get(nonExistentKey); ok {
		t.Errorf("Expected key %v to not exist", nonExistentKey)
	}
}

func TestSortableOfMap_Delete(t *testing.T) {
	mapa := SortableOf[testsortables.TestNode, testsortables.TestNode]()

	sortables := testsortables.RandomTestNodes(20, "node").Each(func(n testsortables.TestNode) {
		mapa.Put(n, n)
	})

	// Delete a random key
	deleted := sortables.Rand()
	mapa.Delete(deleted)

	// Check that the key was deleted
	if _, ok := mapa.Get(deleted); ok {
		t.Errorf("Expected key %v to be deleted", deleted)
	}
}

func TestSortableOfMap_Contains(t *testing.T) {
	mapa := SortableOf[testsortables.TestNode, testsortables.TestNode]()

	sortables := testsortables.RandomTestNodes(10, "node").Each(func(n testsortables.TestNode) {
		mapa.Put(n, n)
	})

	// Test for existing key
	sortables.Each(func(n testsortables.TestNode) {
		if !mapa.Contains(n) {
			t.Errorf("Expected map to contain key %v", n)
		}
	})

	// Test for non-existent key
	nonExistentKey := testsortables.TestNode("non-existent")
	if mapa.Contains(nonExistentKey) {
		t.Errorf("Expected map to not contain key %v", nonExistentKey)
	}
}

func TestSortableOfMap_LenAndClear(t *testing.T) {
	mapa := SortableOf[testsortables.TestNode, testsortables.TestNode]()

	testsortables.RandomTestNodes(10, "node").Each(func(n testsortables.TestNode) {
		mapa.Put(n, n)
	})

	if mapa.Len() != 10 {
		t.Errorf("Expected map length to be 10, got %d", mapa.Len())
	}

	mapa.Put(testsortables.TestNode("non-existent"), testsortables.TestNode("non-existent"))

	if mapa.Len() != 11 {
		t.Errorf("Expected map length to be 11, got %d", mapa.Len())
	}

	mapa.Delete(testsortables.TestNode("non-existent"))

	if mapa.Len() != 10 {
		t.Errorf("Expected map length to be 10, got %d", mapa.Len())
	}

	mapa.Clear()

	if mapa.Len() != 0 {
		t.Errorf("Expected map length to be 0 after Clear, got %d", mapa.Len())
	}
}

func TestSortableOfMap_Iterator(t *testing.T) {
	mapa := SortableOf[testsortables.TestNode, testsortables.TestNode]()

	exp := make([]Entry[testsortables.TestNode, testsortables.TestNode], 20)

	sortables := testsortables.RandomTestNodes(20, "node").
		Operation(func(i int, n *iterables.Slice[testsortables.TestNode]) {
			it, that := n.At(i), n.Rand()
			mapa.Put(it, that)
			exp[i] = NewAnyEntry(it, that)
		}).Values()

	iter := mapa.Iterator()

	count := 0
	for key, value, ok := iter.Next(); ok; key, value, ok = iter.Next() {
		var entry Entry[testsortables.TestNode, testsortables.TestNode] = NewAnyEntry(key, value)
		if !arrays.ContainsBy(&exp, entry, func(a, b Entry[testsortables.TestNode, testsortables.TestNode]) bool {
			return a.Key().Equal(b.Key()) && a.Value().Equal(b.Value())
		}) {
			t.Errorf("Expected key-value pair %v to be in the map", entry)
		}
		count++
	}

	if count != len(sortables) {
		t.Errorf("Expected to iterate over 10 elements, iterated over %d", count)
	}
}
