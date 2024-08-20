// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package maps

import (
	"github.com/andrerrcosta2/gtools/pkg/arrays"
	"github.com/andrerrcosta2/gtools/pkg/testdata"
	"testing"
)

func TestSortableOfMap_PutAndGet(t *testing.T) {
	m := SortableOf[testdata.TestNode, int]()

	// Test inserting and retrieving a single value
	key := testdata.TestNode("A")
	value := 1
	m.Put(key, value)

	if got, ok := m.Get(key); !ok || got != value {
		t.Errorf("Expected %v, got %v", value, got)
	}

	// Test inserting and retrieving multiple values
	keyB := testdata.TestNode("B")
	valueB := 2
	m.Put(keyB, valueB)

	if got, ok := m.Get(keyB); !ok || got != valueB {
		t.Errorf("Expected %v, got %v", valueB, got)
	}

	// Test retrieving a non-existent key
	nonExistentKey := testdata.TestNode("C")
	if _, ok := m.Get(nonExistentKey); ok {
		t.Errorf("Expected key %v to not exist", nonExistentKey)
	}
}

func TestSortableOfMap_Delete(t *testing.T) {
	m := SortableOf[testdata.TestNode, int]()

	key := testdata.TestNode("A")
	value := 1
	m.Put(key, value)

	// Ensure the key-value pair exists before deletion
	if _, ok := m.Get(key); !ok {
		t.Errorf("Expected key %v to exist", key)
	}

	// Delete the key and check that it's removed
	m.Delete(key)
	if _, ok := m.Get(key); ok {
		t.Errorf("Expected key %v to be deleted", key)
	}
}

func TestSortableOfMap_Contains(t *testing.T) {
	m := SortableOf[testdata.TestNode, int]()

	key := testdata.TestNode("A")
	value := 1
	m.Put(key, value)

	// Test for existing key
	if !m.Contains(key) {
		t.Errorf("Expected map to contain key %v", key)
	}

	// Test for non-existent key
	nonExistentKey := testdata.TestNode("B")
	if m.Contains(nonExistentKey) {
		t.Errorf("Expected map to not contain key %v", nonExistentKey)
	}
}

func TestSortableOfMap_LenAndClear(t *testing.T) {
	m := SortableOf[testdata.TestNode, int]()

	if m.Len() != 0 {
		t.Errorf("Expected map length to be 0, got %d", m.Len())
	}

	m.Put(testdata.TestNode("A"), 1)
	m.Put(testdata.TestNode("B"), 2)

	if m.Len() != 2 {
		t.Errorf("Expected map length to be 2, got %d", m.Len())
	}

	m.Clear()

	if m.Len() != 0 {
		t.Errorf("Expected map length to be 0 after Clear, got %d", m.Len())
	}
}

func TestSortableOfMap_Iterator(t *testing.T) {
	m := SortableOf[testdata.TestNode, int]()
	m.Put(testdata.TestNode("A"), 1)
	m.Put(testdata.TestNode("B"), 2)
	// Overriding some elements
	m.Put(testdata.TestNode("A"), 12)
	m.Put(testdata.TestNode("C"), 3)
	m.Put(testdata.TestNode("D"), 1)
	m.Put(testdata.TestNode("E"), 5)
	m.Put(testdata.TestNode("F"), 6)
	m.Put(testdata.TestNode("G"), 7)
	m.Put(testdata.TestNode("H"), 8)
	m.Put(testdata.TestNode("I"), 9)
	m.Put(testdata.TestNode("J"), 10)

	iter := m.Iterator()

	count := 0
	for {
		key, value, ok := iter.Next()
		if !ok {
			break
		}
		if !key.Equal(testdata.TestNode("A")) && !key.Equal(testdata.TestNode("B")) && !key.Equal(testdata.TestNode("C")) &&
			!key.Equal(testdata.TestNode("D")) && !key.Equal(testdata.TestNode("E")) && !key.Equal(testdata.TestNode("F")) &&
			!key.Equal(testdata.TestNode("G")) && !key.Equal(testdata.TestNode("H")) && !key.Equal(testdata.TestNode("I")) &&
			!key.Equal(testdata.TestNode("J")) {
			t.Errorf("Unexpected key %v", key)
		}
		if value != 1 && value != 2 && value != 12 && value != 3 && value != 5 &&
			value != 6 && value != 7 && value != 8 && value != 9 &&
			value != 10 {
			t.Errorf("Unexpected value %v", value)
		}
		count++
	}

	if count != 10 {
		t.Errorf("Expected to iterate over 10 elements, iterated over %d", count)
	}

	sameKeys := arrays.ContainsAllBy[testdata.TestNode, testdata.TestNode](m.Keys(),
		[]testdata.TestNode{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"},
		func(node testdata.TestNode, node2 testdata.TestNode) bool {
			return node.Equal(node2)
		})
	if !sameKeys {
		t.Errorf("Expected keys to be the same")
	}

	sameValues := arrays.ContainsAllBy[int, int](m.Values(),
		[]int{2, 12, 3, 1, 5, 6, 7, 8, 9, 10},
		func(value int, value2 int) bool {
			return value == value2
		})

	if !sameValues {
		t.Errorf("Expected values to be the same")
	}
}
