// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package lists

import (
	"testing"

	"github.com/andrerrcosta2/gtools/core/data/comparators"
	"github.com/andrerrcosta2/gtools/core/data/str/iterables"
	"github.com/andrerrcosta2/gtools/core/seeders/random"
	"github.com/andrerrcosta2/gtools/core/util/typeutil/generics"
	"github.com/andrerrcosta2/gtools/gtests/testingtools"
	"github.com/andrerrcosta2/gtools/gtests/testingtools/config/testlogs"
	"github.com/andrerrcosta2/gtools/gtests/testingtools/themes"
)

// TestConcHashLinked_Concurrency_Add tests the add method
//
// This method is used to add elements to the list
// After it is called this test must assert:
//   - the element is added to the list
//   - the list size is increased
//   - the head value points to the first value inserted
//   - the tail value points to the last value inserted
//   - the index map is updated and can be used to retrieve the inserted value by its hash
func TestConcHashLinked_Concurrency_Add(t *testing.T) {
	// Helper
	tt := testingtools.ConcLitetm(t, testlogs.OnFailure, themes.Default)
	// Test empty
	cmp := comparators.StringOrdered[int]{}
	list := ConcHashLinked[int, string](cmp)
	cast := list.(*concHashLinked[int, string])
	if cast.size() != 0 {
		t.Errorf("Expected size to be 0, but it is not")
	}

	// Test adding elements and asserts its related states
	it := iterables.OfSlice(1, 2, 3, 5, 6, 7)
	it.Parallel(func(i, v int) { // Assert real time states
		hash := cast.hash(v)
		idx := list.Add(v)
		tt.RegisterConst(idx, hash)
		tt.RegisterCalls(1, "add")
		if idx == 0 {
			tt.RegisterConst(v, "head")
		}
		// assert the size is compare to the number of calls
		if cast.size() < tt.CallsTo("add") {
			t.Errorf("Expected size to be %d, but received %d", tt.CallsTo("add"), cast.size())
		}
		// assert the head still always the same
		if cast.head().Value() != tt.Const("head").(int) {
			t.Errorf("Expected head value to be first inserted '%d' at '%d', but received '%d'", cast.head().Value(), i, cast.head().Value())
		}
		// assert the hash map contains the same values
		if cast.idx[cast.hash(v)][0].Value() != v {
			t.Errorf("Expected index value to be '%d' at '%d', but received '%d'", v, i, cast.idx[cast.hash(v)][0].Value())
		}
	}, 6).
		// Assert final states
		EachN(func(i, v int) {
			if cast.idx[cast.hash(v)][0].Value() != v {
				t.Errorf("Expected index value to be '%d' at '%d', but received '%d'", v, i, cast.idx[cast.hash(v)][0].Value())
			}
		})

	if cast.size() != it.Len() {
		t.Errorf("Expected size to be %d, but received %d", it.Len(), cast.size())
	}

	if len(cast.idx) != it.Len() {
		t.Errorf("Expected index length to be %d, but received %d", it.Len(), len(cast.idx))
	}

	head := tt.Const("head").(int)
	if cast.head().Value() != head {
		t.Errorf("Expected head value to be '%d', but received '%d'", head, cast.head().Value())
	}
}

// TestConcHashLinked_Concurrency_Contains tests the contains method
//
// This method is used to check if an element is in the list
// After it is called this test must assert:
//   - the element is found if it was inserted
//   - the element is not found if it was not inserted
func TestConcHashLinked_Concurrency_Contains(t *testing.T) {
	cmp := comparators.StringOrdered[int]{}
	list := ConcHashLinked[int, string](cmp)
	cast := list.(*concHashLinked[int, string])

	it := iterables.OfSlice(2, 4, 6, 8, 10, 12)
	it.Parallel(func(i, v int) {
		list.Add(v)
		cast = list.(*concHashLinked[int, string])
		if cast.Contains(v - 1) {
			t.Errorf("Expected list to not contain '%d', but it does", i)
		}
		if !cast.Contains(v) {
			t.Errorf("Expected list to contain '%d', but it does not", v)
		}
	}, 3)
}

// TestHashLinked_Get tests the get method
//
// This method is used to retrieve an element from the list
// After it is called this test must assert:
//   - the element is found if it was inserted
//   - the element is not found if it was not inserted
//   - the element is not found if the index is negative
func TestConcHashLinked_Concurrency_Get(t *testing.T) {
	tt := testingtools.ConcLite(t, testlogs.Default)
	cmp := comparators.StringOrdered[int]{}
	list := ConcHashLinked[int, string](cmp)

	it := iterables.OfSlice(2, 4, 6, 8, 10, 12)

	it.Parallel(func(i, v int) {
		hash := cmp.Hash(v)
		idx := list.Add(v)
		tt.RegisterConst(idx, hash)
		if val, ok := list.Get(idx); !ok {
			t.Errorf("Expected list to contain inserted index, but received %d", val)
		}
	}, 3)
}

// TestConcHashLinked_Concurrency_IndexOf tests the indexOf method
//
// after each insertion the list state should assert
//   - the index map can be used to retrieve the inserted value
//   - the IndexOf can retrieve the same index the value was inserted
//   - not inserted values must retrieve '-1'
func TestConcHashLinked_Concurrency_IndexOf(t *testing.T) {
	cmp := comparators.StringOrdered[int]{}
	list := ConcHashLinked[int, string](cmp)
	cast := list.(*concHashLinked[int, string])

	it := iterables.OfSlice(2, 4, 6, 8, 10, 12)
	it.Parallel(func(i, v int) {
		list.Add(v)
		if idx := cast.IndexOf(v - 1); idx != -1 {
			t.Errorf("Expected list to not contain '%d', but received its index at %d", i, idx)
		}
		if idx := cast.IndexOf(v); idx == -1 {
			t.Errorf("Expected list to retrieve a valid index for inserted value '%d', but "+
				"received %d", v, idx)
		}
	}, 3)
}

// TestConcHashLinked_Concurrency_Insert tests the insert method
//
// after each insertion the list state should assert
//   - its size is increase
//   - the insertion index can be used to retrieve the inserted value
//   - the index map can be used to retrieve the inserted value
func TestConcHashLinked_Concurrency_Insert(t *testing.T) {
	cmp := comparators.StringOrdered[int]{}
	list := ConcHashLinked[int, string](cmp, 1, 3, 5, 7, 9, 11)
	cast := list.(*concHashLinked[int, string])

	initialSize := cast.size()
	if initialSize != 6 {
		t.Errorf("Expected *doubleLinked.size() to be 6, but received %d", initialSize)
	}

	it := iterables.OfSlice(2, 4, 6, 8, 10, 12)
	it.Parallel(func(i, v int) {
		at := (i + 1) % 6
		if !list.Insert(at, v) {
			t.Errorf("Expected list to insert '%d' at index '%d', but it did not", v, i)
		}
	}, 3)

	if cast.size() != 12 {
		t.Errorf("Expected *doubleLinked.size() to be 12, but received %d", cast.size())
	}

}

// TestConcHashLinked_Concurrency_Remove tests the remove method
//
// after each removal the list state should assert
//   - its size is decrease
//   - the index map can't be used to retrieve the removed value
//   - the removed value is not present in the list
//   - the removed value cannot be reached by the list iteration
func TestConcHashLinked_Concurrency_Remove(t *testing.T) {
	// Helper
	tt := testingtools.ConcLite(t, testlogs.OnFailure)
	cmp := comparators.StringOrdered[int]{}
	list := ConcHashLinked[int, string](cmp)
	cast := list.(*concHashLinked[int, string])

	it := iterables.OfSlice(0, 3, 6, 9, 12, 15, 18, 21, 24)
	tt.RegisterCalls(cast.size(), "current-size")
	tt.AssertCalls(0, "current-size")

	it.Parallel(func(i, v int) {
		hash := cmp.Hash(v)
		idx := list.Add(v)
		tt.RegisterConst(idx, hash)
		tt.RegisterCalls(1, "add", "current-size")
		rv := i * 2
		rhash := cmp.Hash(rv)
		shouldRemove := tt.Const(rhash) != nil
		wasRemoved := list.Remove(i * 2)
		if shouldRemove {
			if !wasRemoved {
				t.Errorf("Expected list to remove value at index '%d', but received '%t'", rv, wasRemoved)
			}

			// Clear registers
			tt.RegisterCalls(1, "removed")
			tt.RegisterCalls(-1, "current-size")
			tt.RegisterConst(nil, rhash)
			it.RemoveAt(idx)
		} else if wasRemoved {
			t.Errorf("Expected list to not remove value at index '%d', but received '%t'", rv, wasRemoved)
		}
	}, 3)

	tt.AssertCalls(tt.CallsTo("add")-tt.CallsTo("removed"), "current-size")
	tt.AssertCalls(9, "add")

	if cast.size() != tt.CallsTo("current-size") {
		t.Errorf("Expected *concHashLinked.size() to end up from '%d' to '%d', but received %d", it.Len(),
			tt.CallsTo("current-size"), cast.size())
	}

	if len(cast.idx) != tt.CallsTo("current-size") {
		t.Errorf("Expected *concHashLinked.idx length to decrease from '%d' to '%d', but received %d", it.Len(),
			tt.CallsTo("current-size"), len(cast.idx))
	}

	// Can't know which value was inserted first.
	// But can assert it must keep sequential after deletion
	it.EachN(func(i, v int) {
		v, ok := list.Get(i)
		if !ok {
			t.Errorf("Expected list to have value at index '%d', but received '%t'", i, ok)
		}
		// Deletions can only decrease the index
		idx := tt.Const(cmp.Hash(v)).(int)
		if idx < i {
			t.Errorf("Expected list to have value '%d' at index '%d', but received '%d'", v, i, idx)
		}
	})

}

// TestConcHashLinked_Concurrency_RemoveAt tests the removeAt method
//
// after each removal the list state should assert
//   - the removed value is the same as inserted at the same index
//   - the removed value cannot be reached by the list iteration
//   - the index map can't be used to retrieve the removed value
func TestConcHashLinked_Concurrency_RemoveAt(t *testing.T) {
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

// TestConcHashLinked_Concurrency_Set tests the ToSet method
//
// after the method set is called the list state should assert
//   - the set value is the same as inserted at the same index
//   - the set value can be reached by the list iteration
//   - if an earlier value was set, the index map can't be used to retrieve the removed value
//   - the index map should be able to retrieve the current one at the same address (performance)
func TestConcHashLinked_Concurrency_Set(t *testing.T) {
	cmp := comparators.StringOrdered[int]{}
	list := ConcHashLinked[int, string](cmp, 2, 4, 6, 8, 10)
	cast := list.(*concHashLinked[int, string])

	it := iterables.OfSlice(1, 3, 5, 7, 9)
	it.Parallel(func(i, v int) {
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
	}, 3)
	if list.Size() != 5 {
		t.Errorf("Expected list size to be 5, but received %d", list.Size())
	}
}

// TestConcHashLinked_Concurrency_Size tests the Size method under concurrent access
//
// after each insertion the list size should be increased
// after each removal the list size should be decreased
func TestConcHashLinked_Concurrency_Size(t *testing.T) {
	// std
	tt := testingtools.ConcLitetm(t, testlogs.OnFailure, themes.Default)

	cmp := comparators.StringOrdered[int]{}
	list := ConcHashLinked[int, string](cmp)
	cast := list.(*concHashLinked[int, string])
	if list.Size() != 0 {
		t.Errorf("Expected list size to be 0, but received %d", list.Size())
	}
	initSize := cast.size()
	// test with different elements
	iterables.OfSlice(1, 3, 5, 7, 9).Parallel(func(i, v int) {
		list.Add(v)
		tt.RegisterCalls(1, "add", "current-size")
	}, 3)

	// assertions
	if list.Size() != tt.CallsTo("current-size") {
		t.Errorf("Expected list size to grow from %d to %d, but received %d", initSize,
			tt.CallsTo("current-size"), list.Size())
	}
	if cast.size() != tt.CallsTo("current-size") {
		t.Errorf("Expected *concHashLinked.size() to grow from %d to %d, but received %d", initSize,
			tt.CallsTo("current-size"), cast.size())
	}
	if cast._size != tt.CallsTo("current-size") {
		t.Errorf("Expected *concHashLinked._size to grow from %d to %d, but received %d", initSize,
			tt.CallsTo("current-size"), cast._size)
	}
	if len(cast.idx) != tt.CallsTo("current-size") {
		t.Errorf("Expected *concHashLinked.idx length to grow from %d to %d, but received %d", initSize,
			tt.CallsTo("current-size"), len(cast.idx))
	}

	list.SoftClear()
	tt.Clear()
	initSize = cast.size()
	tt.RegisterCalls(initSize, "current-size")
	// test with repeated elements
	iterables.OfSlice(1, 1, 3, 3, 5, 5, 7, 7, 9, 9).Parallel(func(i, v int) {
		list.Add(v)

		// register assertion informations
		hash := cmp.Hash(v)
		tt.RegisterCalls(1, "add", "current-size")
		if tt.Const(hash) == nil {
			tt.RegisterCalls(1, "hash-size")
		}
		tt.RegisterConst(v, cmp.Hash(v))
	}, 3)

	// assertions
	if list.Size() != tt.CallsTo("current-size") {
		t.Errorf("Expected list size to grow from %d to %d, but received %d", initSize,
			tt.CallsTo("current-size"), list.Size())
	}
	if cast.size() != tt.CallsTo("current-size") {
		t.Errorf("Expected *concHashLinked.size() to grow from %d to %d, but received %d",
			initSize, tt.CallsTo("current-size"), cast.size())
	}
	if cast._size != tt.CallsTo("current-size") {
		t.Errorf("Expected *concHashLinked._size to grow from %d to %d, but received %d",
			initSize, tt.CallsTo("current-size"), cast._size)
	}

	if len(cast.idx) != tt.CallsTo("hash-size") {
		t.Errorf("Expected *concHashLinked.idx length to grow from %d to %d, but received %d",
			initSize, tt.CallsTo("hash-size"), len(cast.idx))
	}
}
