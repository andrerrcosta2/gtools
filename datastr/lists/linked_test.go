// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package lists

import (
	"github.com/andrerrcosta2/gtools/core/data/comparators"
	"github.com/andrerrcosta2/gtools/core/data/str/iterables"
	"github.com/andrerrcosta2/gtools/core/generics"
	"github.com/andrerrcosta2/gtools/core/seeders/random"
	"testing"
)

// TestHashLinked tests the hashLinked creation states
//
// This method is used to check if the hashLinked is created correctly
// After it is called this test must assert:
//   - the head is nil if it is created with no elements
//   - the head is not nil and is different from the tail if it is created with elements
//   - the tail is nil if it is created with no elements
//   - the tail is not nil and is different from the head if it is created with elements
//   - the size is 0 if it is created with no elements
//   - the size is not 0 if it is created with elements
//   - the index map is empty if it is created with no elements
//   - the index map is not empty if it is created with elements and can be used to retrieve the inserted value by its hash
func TestHashLinked(t *testing.T) {
	cmp := comparators.StringOrdered[int]{}
	// Test empty
	list := HashLinked[int, string](cmp)
	cast := list.(*hashLinked[int, string])
	if cast.head() != nil {
		t.Errorf("Expected head to be nil, but it is not")
	}
	if cast.tail() != nil {
		t.Errorf("Expected tail to be nil, but it is not")
	}
	if cast.size() != 0 {
		t.Errorf("Expected size to be 0, but it is not")
	}
	if len(cast.idx) != 0 {
		t.Errorf("Expected idx length to be 0, but it is not")
	}
	if list.Size() != 0 {
		t.Errorf("Expected list.Size() to be 0, but it is not")
	}

	// Test with elements
	list = HashLinked[int, string](cmp, 1, 2, 3, 5, 6, 7)
	cast = list.(*hashLinked[int, string])
	if cast.head() == nil {
		t.Errorf("Expected head to not be nil, but it is")
	}
	if cast.head().Value() != 1 {
		t.Errorf("Expected head value to be 1, but it is not")
	}
	if cast.tail() == nil {
		t.Errorf("Expected tail to not be nil, but it is")
	}
	if cast.tail().Value() != 7 {
		t.Errorf("Expected tail value to be 7, but it is not: %v", cast.tail().Value())
	}
	if cast.size() != 6 {
		t.Errorf("Expected size to be 6, but it is not")
	}
	if len(cast.idx) != 6 {
		t.Errorf("Expected idx length to be 6, but it is not")
	}
	if list.Size() != 6 {
		t.Errorf("Expected list.Size() to be 6, but it is not")
	}
	iterables.OfSlice(1, 2, 3, 5, 6, 7).EachN(func(i, v int) {
		hash := cast.hash(v)
		if cast.idx[hash] == nil {
			t.Errorf("Expected idx[%d] to not be nil, but it is", i)
		}
		if cast.idx[hash][0].Value() != v {
			t.Errorf("Expected idx[%d] value to be %d, but it is not", i, i+1)
		}
	})
}

// TestHashLinked_Add tests the add method
//
// This method is used to add elements to the list
// After it is called this test must assert:
//   - the element is added to the list
//   - the list size is increased
//   - the head value points to the first value inserted
//   - the tail value points to the last value inserted
//   - the index map is updated and can be used to retrieve the inserted value by its hash
func TestHashLinked_Add(t *testing.T) {
	// Test empty
	cmp := comparators.StringOrdered[int]{}
	list := HashLinked[int, string](cmp)
	cast := list.(*hashLinked[int, string])
	if cast.size() != 0 {
		t.Errorf("Expected size to be 0, but it is not")
	}

	// Test adding elements and asserts its related states
	it := iterables.OfSlice(1, 2, 3, 5, 6, 7)
	it.EachN(func(i, v int) {
		list.Add(v)
		cast = list.(*hashLinked[int, string])
		if cast.size() != i+1 {
			t.Errorf("Expected size to be %d, but received %d", v, cast.size())
		}
		if cast.head().Value() != it.At(0) {
			t.Errorf("Expected head value to be '%d' at '%d', but received '%d'", cast.head().Value(), i, cast.head().Value())
		}
		if cast.tail().Value() != v {
			t.Errorf("Expected tail value to be '%d' at '%d', but received '%d'", it.At(i), i, cast.tail().Value())
		}
		if len(cast.idx) != i+1 {
			t.Errorf("Expected index length to be %d, but received %d", i+1, len(cast.idx))
		}
		if cast.idx[cast.hash(v)][0].Value() != v {
			t.Errorf("Expected index value to be '%d' at '%d', but received '%d'", v, i, cast.idx[cast.hash(v)][0].Value())
		}
	})
}

// TestHashLinked_Contains tests the contains method
//
// This method is used to check if an element is in the list
// After it is called this test must assert:
//   - the element is found if it was inserted
//   - the element is not found if it was not inserted
func TestHashLinked_Contains(t *testing.T) {
	cmp := comparators.StringOrdered[int]{}
	list := HashLinked[int, string](cmp)
	cast := list.(*hashLinked[int, string])

	it := iterables.OfSlice(2, 4, 6, 8, 10, 12)
	it.EachN(func(i, v int) {
		list.Add(v)
		cast = list.(*hashLinked[int, string])
		if cast.Contains(v - 1) {
			t.Errorf("Expected list to not contain '%d', but it does", i)
		}
		if !cast.Contains(v) {
			t.Errorf("Expected list to contain '%d', but it does not", v)
		}
	})
}

// TestHashLinked_Get tests the get method
//
// This method is used to retrieve an element from the list
// After it is called this test must assert:
//   - the element is found if it was inserted
//   - the element is not found if it was not inserted
//   - the element is not found if the index is negative
func TestHashLinked_Get(t *testing.T) {
	cmp := comparators.StringOrdered[int]{}
	list := HashLinked[int, string](cmp)

	it := iterables.OfSlice(2, 4, 6, 8, 10, 12)
	it.EachN(func(i, v int) {
		list.Add(v)
		if val, ok := list.Get(i + 1); ok {
			t.Errorf("Expected list to not contain found data, but received %d", val)
		}
		if val, ok := list.Get(i); !ok {
			t.Errorf("Expected list to contain found data, but received %d", val)
		}
	})

	// Test negative index
	if val, ok := list.Get(-1); ok {
		t.Errorf("Expected list to not contain found data, but received %d", val)
	}
}

// TestHashLinked_HardClear tests the hardClear method
//
// This method is used to remove all elements from the list
// After it is called this test must assert:
//   - the list size is 0
//   - the list head is nil
//   - the list tail is nil
//   - the index map is empty
func TestHashLinked_HardClear(t *testing.T) {
	cmp := comparators.StringOrdered[int]{}
	list := HashLinked[int, string](cmp, 1, 2, 3, 5, 6, 7)
	cast := list.(*hashLinked[int, string])

	if cast.size() != 6 {
		t.Errorf("Expected *doubleLinked.size() to be 6, but received %d", cast.size())
	}

	if cast._size != 6 {
		t.Errorf("Expected *doubleLinked._size to be 6, but received %d", cast._size)
	}

	if len(cast.idx) != 6 {
		t.Errorf("Expected *hashLinked.idx length to be 6, but received %d", len(cast.idx))
	}

	if list.Size() != 6 {
		t.Errorf("Expected list.Size() to be 6, but received %d", list.Size())
	}

	if cast.head() == nil {
		t.Errorf("Expected *doubleLinked.head() to not be nil, but received %v", cast.head())
	}

	if cast.tail() == nil {
		t.Errorf("Expected *doubleLinked.tail() to not be nil, but received %v", cast.tail())
	}

	if cast.head() == cast.tail() {
		t.Errorf("Expected *doubleLinked.head() to not be compare to *doubleLinked.tail(), but received %t",
			cast.head() == cast.tail())
	}

	list.HardClear()
	cast = list.(*hashLinked[int, string])

	if cast.size() != 0 {
		t.Errorf("Expected *doubleLinked.size() to be 0, but received %d", cast.size())
	}

	if cast._size != 0 {
		t.Errorf("Expected *doubleLinked._size to be 0, but received %d", cast._size)
	}

	if len(cast.idx) != 0 {
		t.Errorf("Expected idx *hashLinked.idx length to be 0, but received %d", len(cast.idx))
	}

	if cast.head() != nil {
		t.Errorf("Expected head to be *doubleLinked.head() to be nil, but received %v", cast.head())
	}

	if cast.tail() != nil {
		t.Errorf("Expected *doubleLinked.tail() to be nil, but received %v", cast.tail())
	}

	if list.Size() != 0 {
		t.Errorf("Expected list.Size() to be 0, but received %d", list.Size())
	}
}

// TestHashLinked_IndexOf tests the indexOf method
//
// after each insertion the list state should assert
//   - the index map can be used to retrieve the inserted value
//   - the IndexOf can retrieve the same index the value was inserted
//   - not inserted values must retrieve '-1'
func TestHashLinked_IndexOf(t *testing.T) {
	cmp := comparators.StringOrdered[int]{}
	list := HashLinked[int, string](cmp)
	cast := list.(*hashLinked[int, string])

	it := iterables.OfSlice(2, 4, 6, 8, 10, 12)
	it.EachN(func(i, v int) {
		list.Add(v)
		if idx := cast.IndexOf(v - 1); idx != -1 {
			t.Errorf("Expected list to not contain '%d', but received its index at %d", i, idx)
		}
		if idx := cast.IndexOf(v); idx != i {
			t.Errorf("Expected list to have value at the same insertion index '%d', but received %d", i, idx)
		}
	})
}

// TestHashLinked_Insert tests the insert method
//
// after each insertion the list state should assert
//   - its size is increase
//   - the insertion index can be used to retrieve the inserted value
//   - the index map can be used to retrieve the inserted value
func TestHashLinked_Insert(t *testing.T) {
	cmp := comparators.StringOrdered[int]{}
	list := HashLinked[int, string](cmp, 1, 3, 5, 7, 9, 11)
	cast := list.(*hashLinked[int, string])

	initialSize := cast.size()
	if initialSize != 6 {
		t.Errorf("Expected *doubleLinked.size() to be 6, but received %d", initialSize)
	}

	it := iterables.OfSlice(2, 4, 6, 8, 10, 12)
	it.EachN(func(i, v int) {
		at := random.Int(0, 0, it.Len()).At(0)
		if !list.Insert(at, v) {
			t.Errorf("Expected list to insert '%d' at index '%d', but it did not", v, i)
		}
		initialSize++
		get, ok := list.Get(at)
		if !ok {
			t.Errorf("Expected list to insert '%d' at index '%d', but it did not", v, i)
		}

		if cast.idx[cmp.Hash(v)][0].Value() != get {
			t.Errorf("Expected tail to have value at the same insertion index '%d', "+
				"but received \n%v != \n%v", i, cast.idx[cmp.Hash(v)], cast.tail())
		}

		if cast.IndexOf(v) != at {
			t.Errorf("Expected list to have value at the same insertion index '%d', but received %d", i,
				cast.IndexOf(v))
		}

		if len(cast.idx) != initialSize {
			t.Errorf("Expected *hashLinked.idx length to be %d, but received %d", initialSize, len(cast.idx))
		}

		if cast.size() != initialSize {
			t.Errorf("Expected *hashLinked.size to be %d, but received %d", initialSize, cast.Size())
		}
	})
}

// TestHashLinked_IsEmpty tests the isEmpty method
//
// after each insertion the list state should assert
//   - the list is empty if its size is 0
//   - the list is not empty if its size is greater than 0
func TestHashLinked_IsEmpty(t *testing.T) {
	cmp := comparators.StringOrdered[int]{}
	list := HashLinked[int, string](cmp)
	cast := list.(*hashLinked[int, string])

	if !list.IsEmpty() {
		t.Errorf("Expected *doubleLinked.IsEmpty() to be true, but received false")
	}

	if cast.size() != 0 {
		t.Errorf("Expected *doubleLinked.size() to be 0, but received %d", cast.size())
	}

	list.Add(1)
	if list.IsEmpty() {
		t.Errorf("Expected *doubleLinked.IsEmpty() to be false, but received true")
	}
}

// TestHashLinked_Remove tests the remove method
//
// after each removal the list state should assert
//   - its size is decrease
//   - the index map can't be used to retrieve the removed value
//   - the removed value is not present in the list
//   - the removed value cannot be reached by the list iteration
func TestHashLinked_Remove(t *testing.T) {
	cmp := comparators.StringOrdered[int]{}
	list := HashLinked[int, string](cmp)
	cast := list.(*hashLinked[int, string])

	it := iterables.OfSlice(0, 3, 6, 9, 12, 15, 18, 21, 24)
	it.EachN(func(i, v int) {
		initSize := cast.size()
		list.Add(v)

		if initSize != cast.size()-1 {
			t.Errorf("Expected *doubleLinked.size() to grow from '%d' to '%d', but received %d", initSize,
				cast.size(), initSize)
		}
		remove := i * 2
		removed := list.Remove(remove)
		if remove%3 != 0 {
			if removed {
				t.Errorf("Expected list to not remove not inserted value '%d', but received '%t'",
					remove, removed)
			}
			if cast.Size() != initSize+1 {
				t.Errorf("Expected list to keep its size '%d' after trying to remove not inserted value '%d',"+
					" but received '%d'", initSize, remove, cast.Size())
			}
			if len(cast.idx) != initSize+1 {
				t.Errorf("Expected *hashLinked.idx length to be '%d' after trying to remove not inserted "+
					"value '%d', but received '%d'", initSize, remove, len(cast.idx))
			}
			hash := cmp.Hash(it.At(i))
			if cast.idx[hash] == nil {
				t.Errorf("Expected *hashLinked.idx to have value '%d', but received nil", remove)
			}
			if cast.idx[hash][0].Value() != it.At(i) {
				t.Errorf("Expected *hashLinked.idx to have value '%d', but received '%d'", it.At(i),
					cast.idx[hash][0].Value())
			}
		} else {
			if !removed {
				t.Errorf("Expected list to remove inserted value '%d', but received '%t'", remove, removed)
			}
			if cast.Size() != initSize {
				t.Errorf("Expected list to have size '%d' after removing inserted value '%d',"+
					"but received '%d'", initSize, remove, cast.Size())
			}
			if len(cast.idx) != initSize {
				t.Errorf("Expected *hashLinked.idx length to be '%d' after removing inserted value '%d',"+
					"but received '%d'", initSize, remove, len(cast.idx))
			}
			hash := cmp.Hash(remove)
			nodes, ok := cast.idx[hash]
			if ok || nodes != nil {
				t.Errorf("Expected *hashLinked.idx to not have value '%d', but it retrieved '%v' when "+
					"its hash '%s' was queried: ", remove, nodes, hash)
			}
			val, ok := list.Get(i)
			if ok && val == remove { // must be the pair because the list shrinks
				t.Errorf("Expected value '%d' to not be present in the list, but it is", remove)
			}
		}
	})
}

// TestHashLinked_RemoveAt tests the removeAt method
//
// after each removal the list state should assert
//   - the removed value is the same as inserted at the same index
//   - the removed value cannot be reached by the list iteration
//   - the index map can't be used to retrieve the removed value
func TestHashLinked_RemoveAt(t *testing.T) {
	cmp := comparators.StringOrdered[int]{}
	list := HashLinked[int, string](cmp)
	cast := list.(*hashLinked[int, string])

	it := iterables.OfSlice(0, 3, 6, 9, 12, 15, 18, 21, 24)
	it.EachN(func(i, v int) {
		initSize := cast.size()
		list.Add(v)

		if initSize != cast.size()-1 {
			t.Errorf("Expected *doubleLinked.size() to grow from '%d' to '%d', but received %d", initSize,
				cast.size(), initSize)
		}
		rid := random.Int(1, 0, it.Len()).At(0)
		val, exists := list.Get(rid)
		removedValue, wasRemoved := list.RemoveAt(rid)

		if !exists {
			if wasRemoved {
				t.Errorf("Loop[%d]: Expected list to not remove at not existing index '%d', "+
					"but received '%t' and value '%d'", i, rid, wasRemoved, removedValue)
			}
			if !generics.IsZero(removedValue) {
				t.Errorf("Loop[%d]: Expected response of RemoveAt to be zero value when trying "+
					"to remove not existing index '%d', but received '%d'", i, rid, removedValue)
			}
			if cast.Size() != initSize+1 {
				t.Errorf("Loop[%d]: Expected list to keep its size '%d' after trying to remove not existing "+
					" index '%d', but received '%d'", i, initSize, rid, cast.Size())
			}
			if len(cast.idx) != initSize+1 {
				t.Errorf("Loop[%d]: Expected *hashLinked.idx length to be '%d' after trying to remove not existing "+
					"index '%d', but received '%d'", i, initSize, rid, len(cast.idx))
			}
		} else {
			if !wasRemoved {
				t.Errorf("Loop[%d]: Expected list to remove inserted value '%d', at index '%d', "+
					"but received '%t'", i, val, rid, wasRemoved)
			}
			if cast.Size() != initSize {
				t.Errorf("Loop[%d]: Expected list to have size '%d' after removing inserted value '%d', "+
					"but received '%d'", i, initSize, val, cast.Size())
			}
			if len(cast.idx) != initSize {
				t.Errorf("Loop[%d]: Expected *hashLinked.idx length to be '%d' after removing "+
					"inserted value '%d', but received '%d'", i, initSize, val, len(cast.idx))
			}
			hash := cmp.Hash(val)
			nodes, ok := cast.idx[hash]
			if ok || nodes != nil {
				t.Errorf("Loop[%d]: Expected *hashLinked.idx to not have deleted value [%d]: '%d', "+
					"but it retrieved '%v' when its hash '%s' was queried: ", i, rid, val, nodes, hash)
			}
		}
	})
}

func TestHashLinked_Set(t *testing.T) {
	cmp := comparators.StringOrdered[int]{}
	list := HashLinked[int, string](cmp, 2, 4, 6, 8, 10)
	cast := list.(*hashLinked[int, string])

	it := iterables.OfSlice(1, 3, 5, 7, 9)
	it.EachN(func(i, v int) {
		if !list.Contains(v + 1) {
			t.Errorf("Expected list to contain '%d' at index '%d', but it did not", v+1, i)
		}
		list.Set(i, v)
		if _, ok := list.Get(i); !ok {
			t.Errorf("Expected list to contain '%d' at index '%d', but it did not", v, i)
		}
		if node, ok := cast.idx[cmp.Hash(v+1)]; ok || node != nil {
			t.Errorf("Expected list to not contain '%d' after replacing it by '%d', "+
				"but found a node '%v", v+1, v, node)
		}
		if cast.idx[cmp.Hash(v)][0].Value() != v {
			t.Errorf("Expected tail to have value at the same insertion index '%d', "+
				"but received '%d'", i, v)
		}
		if list.Contains(v + 1) {
			t.Errorf("Expected list to not contain '%d' after replacing it by '%d'", v+1, v)
		}
	})
	if list.Size() != 5 {
		t.Errorf("Expected list size to be 5, but received %d", list.Size())
	}
}

func TestHashLinked_Size(t *testing.T) {
	cmp := comparators.StringOrdered[int]{}
	list := HashLinked[int, string](cmp)
	cast := list.(*hashLinked[int, string])
	if list.Size() != 0 {
		t.Errorf("Expected list size to be 0, but received %d", list.Size())
	}
	initSize := cast.size()
	// test with different elements
	iterables.OfSlice(1, 3, 5, 7, 9).EachN(func(i, v int) {
		list.Add(v)
		if list.Size() != initSize+i+1 {
			t.Errorf("Expected list size to grow from %d to %d, but received %d", initSize, initSize+i+1, list.Size())
		}
		if cast.size() != initSize+i+1 {
			t.Errorf("Expected *hashLinked.size() to grow from %d to %d, but received %d", initSize, initSize+i+1, cast.size())
		}
		if cast._size != initSize+i+1 {
			t.Errorf("Expected *hashLinked._size to grow from %d to %d, but received %d", initSize, initSize+i+1, cast._size)
		}
		if len(cast.idx) != initSize+i+1 {
			t.Errorf("Expected *hashLinked.idx length to grow from %d to %d, but received %d", initSize, initSize+i+1, len(cast.idx))
		}
	})

	list.SoftClear()
	initSize = cast.size()
	// test with repeated elements
	iterables.OfSlice(1, 1, 3, 3, 5, 5, 7, 7, 9, 9).EachN(func(i, v int) {
		list.Add(v)
		if list.Size() != initSize+i+1 {
			t.Errorf("Expected list size to grow from %d to %d, but received %d", initSize, initSize+i+1, list.Size())
		}
		if cast.size() != initSize+i+1 {
			t.Errorf("Expected *hashLinked.size() to grow from %d to %d, but received %d", initSize, initSize+i+1, cast.size())
		}
		if cast._size != initSize+i+1 {
			t.Errorf("Expected *hashLinked._size to grow from %d to %d, but received %d", initSize, initSize+i+1, cast._size)
		}

		if len(cast.idx) != initSize+(i+2)/2 {
			t.Errorf("Expected *hashLinked.idx length to grow from %d to %d, but received %d", initSize, initSize+(i+1)/2, len(cast.idx))
		}
	})
}

func TestHashLinked_SoftClear(t *testing.T) {
	cmp := comparators.StringOrdered[int]{}
	list := HashLinked[int, string](cmp)
	cast := list.(*hashLinked[int, string])
	list.SoftClear()
	if list.Size() != 0 {
		t.Errorf("Expected list size to be 0, but received %d", list.Size())
	}
	if cast.size() != 0 {
		t.Errorf("Expected *hashLinked.size() to be 0, but received %d", cast.size())
	}
	if cast._size != 0 {
		t.Errorf("Expected *hashLinked._size to be 0, but received %d", cast._size)
	}
	if len(cast.idx) != 0 {
		t.Errorf("Expected *hashLinked.idx length to be 0, but received %d", len(cast.idx))
	}
	iterables.OfSlice(1, 3, 5, 7, 9).Each(func(v int) { _ = list.Add(v) })
	if list.Size() != 5 {
		t.Errorf("Expected list size to be 5, but received %d", list.Size())
	}
	if cast.size() != 5 {
		t.Errorf("Expected *hashLinked.size() to be 5, but received %d", cast.size())
	}
	if cast._size != 5 {
		t.Errorf("Expected *hashLinked._size to be 5, but received %d", cast._size)
	}
	if len(cast.idx) != 5 {
		t.Errorf("Expected *hashLinked.idx length to be 5, but received %d", len(cast.idx))
	}
	list.SoftClear()
	if list.Size() != 0 {
		t.Errorf("Expected list size to be 0, but received %d", list.Size())
	}
	if cast.size() != 0 {
		t.Errorf("Expected *hashLinked.size() to be 0, but received %d", cast.size())
	}
	if cast._size != 0 {
		t.Errorf("Expected *hashLinked._size to be 0, but received %d", cast._size)
	}
	if len(cast.idx) != 0 {
		t.Errorf("Expected *hashLinked.idx length to be 0, but received %d", len(cast.idx))
	}
}

// This method is used to convert the list to a slice
// After it is called this test must assert:
//   - the slice is the same as the list
//   - the slice is in the same order as the list
//   - the slice is the same size as the list
func TestHashLinked_ToSlice(t *testing.T) {
	cmp := comparators.StringOrdered[int]{}
	list := HashLinked[int, string](cmp)
	cast, ok := list.(*hashLinked[int, string])
	if !ok {
		t.Errorf("Expected list to be of type *hashLinked, but received %T", list)
	}

	var present = make([]int, 0, 5)
	iterables.OfSlice(1, 3, 5, 7, 9).EachN(func(i, v int) {
		list.Add(v)
		present = append(present, v)
		slice := list.ToSlice()
		if len(slice) != list.Size() {
			t.Errorf("Expected slice length to be %d, but received %d", list.Size(), len(slice))
		}
		if len(slice) != len(cast.idx) {
			t.Errorf("Expected slice length to be %d, but received %d", len(cast.idx), len(slice))
		}
		for ii, p := range present {
			if slice[ii] != p {
				t.Errorf("Expected slice[%d] to be '%d', but received '%d': %v", i, p, slice[i], present)
			}
		}
	})
}

// TestConcHashLinked tests the hashLinked construction states
//
// This method is used to check if the hashLinked is created correctly
// After it is called this test must assert:
//   - the head is nil if it is created with no elements
//   - the head is not nil and is different from the tail if it is created with elements
//   - the tail is nil if it is created with no elements
//   - the tail is not nil and is different from the head if it is created with elements
//   - the size is 0 if it is created with no elements
//   - the size is not 0 if it is created with elements
//   - the index map is empty if it is created with no elements
//   - the index map is not empty if it is created with elements and can be used to retrieve the inserted value by its hash
func TestConcHashLinked(t *testing.T) {
	cmp := comparators.StringOrdered[int]{}
	// Test empty
	list := ConcHashLinked[int, string](cmp)
	cast := list.(*concHashLinked[int, string])
	if cast.head() != nil {
		t.Errorf("Expected head to be nil, but it is not")
	}
	if cast.tail() != nil {
		t.Errorf("Expected tail to be nil, but it is not")
	}
	if cast.size() != 0 {
		t.Errorf("Expected size to be 0, but it is not")
	}
	if len(cast.idx) != 0 {
		t.Errorf("Expected idx length to be 0, but it is not")
	}

	// Test with elements
	list = ConcHashLinked[int, string](cmp, 1, 2, 3, 5, 6, 7)
	cast = list.(*concHashLinked[int, string])
	if cast.head() == nil {
		t.Errorf("Expected head to not be nil, but it is")
	}
	if cast.head().Value() != 1 {
		t.Errorf("Expected head value to be 1, but it is not")
	}
	if cast.tail() == nil {
		t.Errorf("Expected tail to not be nil, but it is")
	}
	if cast.tail().Value() != 7 {
		t.Errorf("Expected tail value to be 7, but it is not: %v", cast.tail().Value())
	}
	if cast.size() != 6 {
		t.Errorf("Expected size to be 6, but it is not")
	}
	if len(cast.idx) != 6 {
		t.Errorf("Expected idx length to be 6, but it is not")
	}
	iterables.OfSlice(1, 2, 3, 5, 6, 7).EachN(func(i, v int) {
		hash := cast.hash(v)
		if cast.idx[hash] == nil {
			t.Errorf("Expected idx[%d] to not be nil, but it is", i)
		}
		if cast.idx[hash][0].Value() != v {
			t.Errorf("Expected idx[%d] value to be %d, but it is not", i, i+1)
		}
	})
}

// TestConcHashLinked_Add tests the add method
//
// This method is used to add elements to the list
// After it is called this test must assert:
//   - the element is added to the list
//   - the list size is increased
//   - the head value points to the first value inserted
//   - the tail value points to the last value inserted
//   - the index map is updated and can be used to retrieve the inserted value by its hash
func TestConcHashLinked_Add(t *testing.T) {
	// Test empty
	cmp := comparators.StringOrdered[int]{}
	list := ConcHashLinked[int, string](cmp)
	cast := list.(*concHashLinked[int, string])
	if cast.size() != 0 {
		t.Errorf("Expected size to be 0, but it is not")
	}

	// Test adding elements and asserts its related states
	it := iterables.OfSlice(1, 2, 3, 5, 6, 7)
	it.EachN(func(i, v int) {
		list.Add(v)
		cast = list.(*concHashLinked[int, string])
		if cast.size() != i+1 {
			t.Errorf("Expected size to be %d, but received %d", v, cast.size())
		}
		if cast.head().Value() != it.At(0) {
			t.Errorf("Expected head value to be '%d' at '%d', but received '%d'", cast.head().Value(), i, cast.head().Value())
		}
		if cast.tail().Value() != v {
			t.Errorf("Expected tail value to be '%d' at '%d', but received '%d'", it.At(i), i, cast.tail().Value())
		}
		if len(cast.idx) != i+1 {
			t.Errorf("Expected index length to be %d, but received %d", i+1, len(cast.idx))
		}
		if cast.idx[cast.hash(v)][0].Value() != v {
			t.Errorf("Expected index value to be '%d' at '%d', but received '%d'", v, i, cast.idx[cast.hash(v)][0].Value())
		}
	})
}

// TestConcHashLinked_Contains tests the contains method
//
// This method is used to check if an element is in the list
// After it is called this test must assert:
//   - the element is found if it was inserted
//   - the element is not found if it was not inserted
func TestConcHashLinked_Contains(t *testing.T) {
	cmp := comparators.StringOrdered[int]{}
	list := ConcHashLinked[int, string](cmp)
	cast := list.(*concHashLinked[int, string])

	it := iterables.OfSlice(2, 4, 6, 8, 10, 12)
	it.EachN(func(i, v int) {
		list.Add(v)
		cast = list.(*concHashLinked[int, string])
		if cast.Contains(v - 1) {
			t.Errorf("Expected list to not contain '%d', but it does", i)
		}
		if !cast.Contains(v) {
			t.Errorf("Expected list to contain '%d', but it does not", v)
		}
	})
}

// TestHashLinked_Get tests the get method
//
// This method is used to retrieve an element from the list
// After it is called this test must assert:
//   - the element is found if it was inserted
//   - the element is not found if it was not inserted
//   - the element is not found if the index is negative
func TestConcHashLinked_Get(t *testing.T) {
	cmp := comparators.StringOrdered[int]{}
	list := ConcHashLinked[int, string](cmp)

	it := iterables.OfSlice(2, 4, 6, 8, 10, 12)
	it.EachN(func(i, v int) {
		list.Add(v)
		if val, ok := list.Get(i + 1); ok {
			t.Errorf("Expected list to not contain found data, but received %d", val)
		}
		if val, ok := list.Get(i); !ok {
			t.Errorf("Expected list to contain found data, but received %d", val)
		}
	})

	// Test negative index
	if val, ok := list.Get(-1); ok {
		t.Errorf("Expected list to not contain found data, but received %d", val)
	}
}

// TestConcHashLinked_HardClear tests the hardClear method
//
// This method is used to remove all elements from the list
// After it is called this test must assert:
//   - the list size is 0
//   - the list head is nil
//   - the list tail is nil
//   - the index map is empty
func TestConcHashLinked_HardClear(t *testing.T) {
	cmp := comparators.StringOrdered[int]{}
	list := ConcHashLinked[int, string](cmp, 1, 2, 3, 5, 6, 7)
	cast := list.(*concHashLinked[int, string])

	if cast.size() != 6 {
		t.Errorf("Expected *doubleLinked.size() to be 6, but received %d", cast.size())
	}

	if cast._size != 6 {
		t.Errorf("Expected *doubleLinked._size to be 6, but received %d", cast._size)
	}

	if len(cast.idx) != 6 {
		t.Errorf("Expected *concHashLinked.idx length to be 6, but received %d", len(cast.idx))
	}

	if list.Size() != 6 {
		t.Errorf("Expected list.Size() to be 6, but received %d", list.Size())
	}

	if cast.head() == nil {
		t.Errorf("Expected *doubleLinked.head() to not be nil, but received %v", cast.head())
	}

	if cast.tail() == nil {
		t.Errorf("Expected *doubleLinked.tail() to not be nil, but received %v", cast.tail())
	}

	if cast.head() == cast.tail() {
		t.Errorf("Expected *doubleLinked.head() to not be compare to *doubleLinked.tail(), but received %t",
			cast.head() == cast.tail())
	}

	list.HardClear()
	cast = list.(*concHashLinked[int, string])

	if cast.size() != 0 {
		t.Errorf("Expected *doubleLinked.size() to be 0, but received %d", cast.size())
	}

	if cast._size != 0 {
		t.Errorf("Expected *doubleLinked._size to be 0, but received %d", cast._size)
	}

	if len(cast.idx) != 0 {
		t.Errorf("Expected idx *concHashLinked.idx length to be 0, but received %d", len(cast.idx))
	}

	if cast.head() != nil {
		t.Errorf("Expected head to be *doubleLinked.head() to be nil, but received %v", cast.head())
	}

	if cast.tail() != nil {
		t.Errorf("Expected *doubleLinked.tail() to be nil, but received %v", cast.tail())
	}

	if list.Size() != 0 {
		t.Errorf("Expected list.Size() to be 0, but received %d", list.Size())
	}
}

// TestConcHashLinked_IndexOf tests the indexOf method
//
// after each insertion the list state should assert
//   - the index map can be used to retrieve the inserted value
//   - the IndexOf can retrieve the same index the value was inserted
//   - not inserted values must retrieve '-1'
func TestConcHashLinked_IndexOf(t *testing.T) {
	cmp := comparators.StringOrdered[int]{}
	list := ConcHashLinked[int, string](cmp)
	cast := list.(*concHashLinked[int, string])

	it := iterables.OfSlice(2, 4, 6, 8, 10, 12)
	it.EachN(func(i, v int) {
		list.Add(v)
		if idx := cast.IndexOf(v - 1); idx != -1 {
			t.Errorf("Expected list to not contain '%d', but received its index at %d", i, idx)
		}
		if idx := cast.IndexOf(v); idx != i {
			t.Errorf("Expected list to have value at the same insertion index '%d', but received %d", i, idx)
		}
	})
}

// TestConcHashLinked_Insert tests the insert method
//
// after each insertion the list state should assert
//   - its size is increase
//   - the insertion index can be used to retrieve the inserted value
//   - the index map can be used to retrieve the inserted value
func TestConcHashLinked_Insert(t *testing.T) {
	cmp := comparators.StringOrdered[int]{}
	list := ConcHashLinked[int, string](cmp, 1, 3, 5, 7, 9, 11)
	cast := list.(*concHashLinked[int, string])

	initialSize := cast.size()
	if initialSize != 6 {
		t.Errorf("Expected *doubleLinked.size() to be 6, but received %d", initialSize)
	}

	it := iterables.OfSlice(2, 4, 6, 8, 10, 12)
	it.EachN(func(i, v int) {
		at := random.Int(0, 0, it.Len()).At(0)
		if !list.Insert(at, v) {
			t.Errorf("Expected list to insert '%d' at index '%d', but it did not", v, i)
		}
		initialSize++
		get, ok := list.Get(at)
		if !ok {
			t.Errorf("Expected list to insert '%d' at index '%d', but it did not", v, i)
		}

		if cast.idx[cmp.Hash(v)][0].Value() != get {
			t.Errorf("Expected tail to have value at the same insertion index '%d', "+
				"but received \n%v != \n%v", i, cast.idx[cmp.Hash(v)], cast.tail())
		}

		if cast.IndexOf(v) != at {
			t.Errorf("Expected list to have value at the same insertion index '%d', but received %d", i,
				cast.IndexOf(v))
		}

		if len(cast.idx) != initialSize {
			t.Errorf("Expected *concHashLinked.idx length to be %d, but received %d", initialSize, len(cast.idx))
		}

		if cast.size() != initialSize {
			t.Errorf("Expected *concHashLinked.size to be %d, but received %d", initialSize, cast.Size())
		}
	})
}

// TestConcHashLinked_IsEmpty tests the isEmpty method
//
// after each insertion the list state should assert
//   - the list is empty if its size is 0
//   - the list is not empty if its size is greater than 0
func TestConcHashLinked_IsEmpty(t *testing.T) {
	cmp := comparators.StringOrdered[int]{}
	list := ConcHashLinked[int, string](cmp)
	cast := list.(*concHashLinked[int, string])

	if !list.IsEmpty() {
		t.Errorf("Expected *doubleLinked.IsEmpty() to be true, but received false")
	}

	if cast.size() != 0 {
		t.Errorf("Expected *doubleLinked.size() to be 0, but received %d", cast.size())
	}

	list.Add(1)
	if list.IsEmpty() {
		t.Errorf("Expected *doubleLinked.IsEmpty() to be false, but received true")
	}
}

// TestConcHashLinked_Remove tests the remove method
//
// after each removal the list state should assert
//   - its size is decrease
//   - the index map can't be used to retrieve the removed value
//   - the removed value is not present in the list
//   - the removed value cannot be reached by the list iteration
func TestConcHashLinked_Remove(t *testing.T) {
	cmp := comparators.StringOrdered[int]{}
	list := ConcHashLinked[int, string](cmp)
	cast := list.(*concHashLinked[int, string])

	it := iterables.OfSlice(0, 3, 6, 9, 12, 15, 18, 21, 24)
	it.EachN(func(i, v int) {
		initSize := cast.size()
		list.Add(v)

		if initSize != cast.size()-1 {
			t.Errorf("Expected *doubleLinked.size() to grow from '%d' to '%d', but received %d", initSize,
				cast.size(), initSize)
		}
		remove := i * 2
		removed := list.Remove(remove)
		if remove%3 != 0 {
			if removed {
				t.Errorf("Expected list to not remove not inserted value '%d', but received '%t'",
					remove, removed)
			}
			if cast.Size() != initSize+1 {
				t.Errorf("Expected list to keep its size '%d' after trying to remove not inserted value '%d',"+
					" but received '%d'", initSize, remove, cast.Size())
			}
			if len(cast.idx) != initSize+1 {
				t.Errorf("Expected *concHashLinked.idx length to be '%d' after trying to remove not inserted "+
					"value '%d', but received '%d'", initSize, remove, len(cast.idx))
			}
			hash := cmp.Hash(it.At(i))
			if cast.idx[hash] == nil {
				t.Errorf("Expected *concHashLinked.idx to have value '%d', but received nil", remove)
			}
			if cast.idx[hash][0].Value() != it.At(i) {
				t.Errorf("Expected *concHashLinked.idx to have value '%d', but received '%d'", it.At(i),
					cast.idx[hash][0].Value())
			}
		} else {
			if !removed {
				t.Errorf("Expected list to remove inserted value '%d', but received '%t'", remove, removed)
			}
			if cast.Size() != initSize {
				t.Errorf("Expected list to have size '%d' after removing inserted value '%d',"+
					"but received '%d'", initSize, remove, cast.Size())
			}
			if len(cast.idx) != initSize {
				t.Errorf("Expected *concHashLinked.idx length to be '%d' after removing inserted value '%d',"+
					"but received '%d'", initSize, remove, len(cast.idx))
			}
			hash := cmp.Hash(remove)
			nodes, ok := cast.idx[hash]
			if ok || nodes != nil {
				t.Errorf("Expected *concHashLinked.idx to not have value '%d', but it retrieved '%v' when "+
					"its hash '%s' was queried: ", remove, nodes, hash)
			}
			val, ok := list.Get(i)
			if ok && val == remove { // must be the pair because the list shrinks
				t.Errorf("Expected value '%d' to not be present in the list, but it is", remove)
			}
		}
	})
}

// TestConcHashLinked_RemoveAt tests the removeAt method
//
// after each removal the list state should assert
//   - the removed value is the same as inserted at the same index
//   - the removed value cannot be reached by the list iteration
//   - the index map can't be used to retrieve the removed value
func TestConcHashLinked_RemoveAt(t *testing.T) {
	cmp := comparators.StringOrdered[int]{}
	list := ConcHashLinked[int, string](cmp)
	cast := list.(*concHashLinked[int, string])

	it := iterables.OfSlice(0, 3, 6, 9, 12, 15, 18, 21, 24)
	it.EachN(func(i, v int) {
		initSize := cast.size()
		list.Add(v)

		if initSize != cast.size()-1 {
			t.Errorf("Expected *doubleLinked.size() to grow from '%d' to '%d', but received %d", initSize,
				cast.size(), initSize)
		}
		rid := random.Int(1, 0, it.Len()).At(0)
		val, exists := list.Get(rid)
		removedValue, wasRemoved := list.RemoveAt(rid)

		if !exists {
			if wasRemoved {
				t.Errorf("Loop[%d]: Expected list to not remove at not existing index '%d', "+
					"but received '%t' and value '%d'", i, rid, wasRemoved, removedValue)
			}
			if !generics.IsZero(removedValue) {
				t.Errorf("Loop[%d]: Expected response of RemoveAt to be zero value when trying "+
					"to remove not existing index '%d', but received '%d'", i, rid, removedValue)
			}
			if cast.Size() != initSize+1 {
				t.Errorf("Loop[%d]: Expected list to keep its size '%d' after trying to remove not existing "+
					" index '%d', but received '%d'", i, initSize, rid, cast.Size())
			}
			if len(cast.idx) != initSize+1 {
				t.Errorf("Loop[%d]: Expected *concHashLinked.idx length to be '%d' after trying to remove not existing "+
					"index '%d', but received '%d'", i, initSize, rid, len(cast.idx))
			}
		} else {
			if !wasRemoved {
				t.Errorf("Loop[%d]: Expected list to remove inserted value '%d', at index '%d', "+
					"but received '%t'", i, val, rid, wasRemoved)
			}
			if cast.Size() != initSize {
				t.Errorf("Loop[%d]: Expected list to have size '%d' after removing inserted value '%d', "+
					"but received '%d'", i, initSize, val, cast.Size())
			}
			if len(cast.idx) != initSize {
				t.Errorf("Loop[%d]: Expected *concHashLinked.idx length to be '%d' after removing "+
					"inserted value '%d', but received '%d'", i, initSize, val, len(cast.idx))
			}
			hash := cmp.Hash(val)
			nodes, ok := cast.idx[hash]
			if ok || nodes != nil {
				t.Errorf("Loop[%d]: Expected *concHashLinked.idx to not have deleted value [%d]: '%d', "+
					"but it retrieved '%v' when its hash '%s' was queried: ", i, rid, val, nodes, hash)
			}
		}
	})
}

// TestConcHashLinked_Set tests the ToSet method
//
// after the method set is called the list state should assert
//   - the set value is the same as inserted at the same index
//   - the set value can be reached by the list iteration
//   - if an earlier value was set, the index map can't be used to retrieve the removed value
//   - the index map should be able to retrieve the current one at the same address (performance)
func TestConcHashLinked_Set(t *testing.T) {
	cmp := comparators.StringOrdered[int]{}
	list := ConcHashLinked[int, string](cmp, 2, 4, 6, 8, 10)
	cast := list.(*concHashLinked[int, string])

	it := iterables.OfSlice(1, 3, 5, 7, 9)
	it.EachN(func(i, v int) {
		if !list.Contains(v + 1) {
			t.Errorf("Expected list to contain '%d' at index '%d', but it did not", v+1, i)
		}
		list.Set(i, v)
		if _, ok := list.Get(i); !ok {
			t.Errorf("Expected list to contain '%d' at index '%d', but it did not", v, i)
		}
		if node, ok := cast.idx[cmp.Hash(v+1)]; ok || node != nil {
			t.Errorf("Expected list to not contain '%d' after replacing it by '%d', "+
				"but found a node '%v", v+1, v, node)
		}
		if cast.idx[cmp.Hash(v)][0].Value() != v {
			t.Errorf("Expected tail to have value at the same insertion index '%d', "+
				"but received '%d'", i, v)
		}
		if list.Contains(v + 1) {
			t.Errorf("Expected list to not contain '%d' after replacing it by '%d'", v+1, v)
		}
	})
	if list.Size() != 5 {
		t.Errorf("Expected list size to be 5, but received %d", list.Size())
	}
}

func TestConcHashLinked_Size(t *testing.T) {
	cmp := comparators.StringOrdered[int]{}
	list := ConcHashLinked[int, string](cmp)
	cast := list.(*concHashLinked[int, string])
	if list.Size() != 0 {
		t.Errorf("Expected list size to be 0, but received %d", list.Size())
	}
	initSize := cast.size()
	// test with different elements
	iterables.OfSlice(1, 3, 5, 7, 9).EachN(func(i, v int) {
		list.Add(v)
		if list.Size() != initSize+i+1 {
			t.Errorf("Expected list size to grow from %d to %d, but received %d", initSize, initSize+i+1, list.Size())
		}
		if cast.size() != initSize+i+1 {
			t.Errorf("Expected *concHashLinked.size() to grow from %d to %d, but received %d", initSize, initSize+i+1, cast.size())
		}
		if cast._size != initSize+i+1 {
			t.Errorf("Expected *concHashLinked._size to grow from %d to %d, but received %d", initSize, initSize+i+1, cast._size)
		}
		if len(cast.idx) != initSize+i+1 {
			t.Errorf("Expected *concHashLinked.idx length to grow from %d to %d, but received %d", initSize, initSize+i+1, len(cast.idx))
		}
	})

	list.SoftClear()
	initSize = cast.size()
	// test with repeated elements
	iterables.OfSlice(1, 1, 3, 3, 5, 5, 7, 7, 9, 9).EachN(func(i, v int) {
		list.Add(v)
		if list.Size() != initSize+i+1 {
			t.Errorf("Expected list size to grow from %d to %d, but received %d", initSize, initSize+i+1, list.Size())
		}
		if cast.size() != initSize+i+1 {
			t.Errorf("Expected *concHashLinked.size() to grow from %d to %d, but received %d", initSize, initSize+i+1, cast.size())
		}
		if cast._size != initSize+i+1 {
			t.Errorf("Expected *concHashLinked._size to grow from %d to %d, but received %d", initSize, initSize+i+1, cast._size)
		}

		if len(cast.idx) != initSize+(i+2)/2 {
			t.Errorf("Expected *concHashLinked.idx length to grow from %d to %d, but received %d", initSize, initSize+(i+1)/2, len(cast.idx))
		}
	})
}

// TestConcHashLinked_SoftClear tests the SoftClear method
//
// after a SoftClear, the list state should:
//   - have a size of 0
//   - have a _size of 0
//   - have an idx length of 0
//   - have no nodes
func TestConcHashLinked_SoftClear(t *testing.T) {
	cmp := comparators.StringOrdered[int]{}
	list := ConcHashLinked[int, string](cmp)
	cast := list.(*concHashLinked[int, string])
	list.SoftClear()
	if list.Size() != 0 {
		t.Errorf("Expected list size to be 0, but received %d", list.Size())
	}
	if cast.size() != 0 {
		t.Errorf("Expected *concHashLinked.size() to be 0, but received %d", cast.size())
	}
	if cast._size != 0 {
		t.Errorf("Expected *concHashLinked._size to be 0, but received %d", cast._size)
	}
	if len(cast.idx) != 0 {
		t.Errorf("Expected *concHashLinked.idx length to be 0, but received %d", len(cast.idx))
	}
	iterables.OfSlice(1, 3, 5, 7, 9).Each(func(v int) { _ = list.Add(v) })
	if list.Size() != 5 {
		t.Errorf("Expected list size to be 5, but received %d", list.Size())
	}
	if cast.size() != 5 {
		t.Errorf("Expected *concHashLinked.size() to be 5, but received %d", cast.size())
	}
	if cast._size != 5 {
		t.Errorf("Expected *concHashLinked._size to be 5, but received %d", cast._size)
	}
	if len(cast.idx) != 5 {
		t.Errorf("Expected *concHashLinked.idx length to be 5, but received %d", len(cast.idx))
	}
	list.SoftClear()
	if list.Size() != 0 {
		t.Errorf("Expected list size to be 0, but received %d", list.Size())
	}
	if cast.size() != 0 {
		t.Errorf("Expected *concHashLinked.size() to be 0, but received %d", cast.size())
	}
	if cast._size != 0 {
		t.Errorf("Expected *concHashLinked._size to be 0, but received %d", cast._size)
	}
	if len(cast.idx) != 0 {
		t.Errorf("Expected *concHashLinked.idx length to be 0, but received %d", len(cast.idx))
	}
}
