// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sets

import (
	"github.com/andrerrcosta2/gtools/core/data/comparators"
	"github.com/andrerrcosta2/gtools/core/data/str/iterables"
	"github.com/andrerrcosta2/gtools/core/generics"
	"math/rand"
	"testing"
)

// TestInsSet_Creation tests the insertion set states after its creation
//
// after the creation the set should:
//   - have the correct number of elements
//   - have the correct number of indexes (ordered by insertion)
func TestInsSet_Creation(t *testing.T) {
	// Test empty set
	cmp := comparators.KeyOrdered[int]{}
	set := Insertion[int, int](cmp)

	if set.Len() != 0 {
		t.Errorf("Expected set length 0, but got %d", set.Len())
	}

	// test set with elements
	set = Insertion[int, int](cmp, 1, 2, 3, 5, 6, 7)

	if set.Len() != 6 {
		t.Errorf("Expected length 6, but got %d", set.Len())
	}

	for i, v := range set.Values() {
		if i < 3 {
			if v != i+1 {
				t.Errorf("Expected value '%d' at idx %d, but got %d", i+1, i, v)
			}
		} else if v != i+2 {
			t.Errorf("Expected set value '%d' at idx %d, but got %d", i+2, i, v)
		}
	}

	// test with duplicate elements
	set = Insertion[int, int](cmp, 1, 1, 2, 2, 3, 3, 5, 5, 6, 6, 7, 7)
	if set.Len() != 6 {
		t.Errorf("Expected set length 6, but got %d", set.Len())
	}

	for i, v := range set.Values() {
		if i < 3 {
			if v != i+1 {
				t.Errorf("Expected set value '%d' at idx %d, but got %d", i+1, i, v)
			}
		} else if v != i+2 {
			t.Errorf("Expected set value '%d' at idx %d, but got %d", i+2, i, v)
		}
	}
}

// TestInsSet_Add tests the insertion set states after adding elements
//
// after adding elements the set should:
//   - have the correct number of elements
//   - have the correct number of indexes (ordered by insertion)
//   - not include duplicates
func TestInsSet_Add(t *testing.T) {
	cmp := comparators.KeyOrdered[int]{}
	// Test with empty set
	set := Insertion[int, int](cmp)

	if set.Len() != 0 {
		t.Errorf("Expected set length 0, but got %d", set.Len())
	}

	it := iterables.OfSlice(1, 2, 3, 5, 6, 7)
	it.EachN(func(i int, v int) {
		set.Add(v)
		if set.Len() != i+1 {
			t.Errorf("Expected set length %d after adding '%d' "+
				"elements, but got %d", i+1, i, set.Len())
		}
	})

	// Test with non-empty set
	set = Insertion[int, int](cmp, 1, 2, 3, 5, 6, 7)

	if set.Len() != 6 {
		t.Errorf("Expected set length 6, but got %d", set.Len())
	}

	it.EachN(func(i int, v int) {
		initSize := set.Len()
		set.Add(v)
		if it.At(i) == 4 {
			if set.Len() != 7 {
				t.Errorf("Expected set length %d after adding missing element "+
					"'%d' element, but got %d", 7, 4, set.Len())
			}
		} else {
			if set.Len() != initSize {
				t.Errorf("Expected set length remain the same '%d' after adding duplicate "+
					"element '%d' element, but got %d", initSize, 4, set.Len())
			}
		}
	})
}

// TestInsSet_Clear tests the insertion set states after clearing.
//
// after calling 'Clear()' the set should:
//   - have no elements
//   - return true when calling 'IsEmpty()'
func TestInsSet_Clear(t *testing.T) {
	cmp := comparators.KeyOrdered[int]{}
	set := Insertion[int, int](cmp, 1, 2, 3, 5, 6, 7)

	if set.Len() != 6 {
		t.Errorf("Expected set length 6, but got %d", set.Len())
	}

	set.Clear()

	if set.Len() != 0 {
		t.Errorf("Expected set length 0 after clear, but got %d", set.Len())
	}
}

// TestInsSet_Delete tests the method Delete(idx)
//
// after calling 'Delete(idx)' the set should:
//   - have the correct number of elements
//   - have the correct number of indexes (ordered by insertion)
//   - do not include the deleted element
func TestInsSet_Delete(t *testing.T) {
	cmp := comparators.KeyOrdered[int]{}
	// Test 1: empty set
	set := Insertion[int, int](cmp)

	if set.Len() != 0 {
		t.Errorf("Expected set length 0, but got %d", set.Len())
	}
	set.Delete(1)
	if set.Len() != 0 {
		t.Errorf("Expected set length 0, but got %d", set.Len())
	}

	// Test 2: non-empty set, delete after each insertion
	it := iterables.OfSlice(1, 2, 3, 5, 6, 7)
	it.EachN(func(i int, v int) {
		set.Add(v)
		if set.Len() != 1 {
			t.Errorf("Expected set length '%d', but got '%d'", 1, set.Len())
		}
		set.Delete(0)
		if set.Len() != 0 {
			t.Errorf("Expected set length return to 0 after delete inserted element '%d', "+
				"but got %d", v, set.Len())
		}
	})

	if set.Len() != 0 {
		t.Errorf("Expected set length 0, but got %d", set.Len())
	}

	// Test 3: non-empty set, delete after all insertions
	it.EachN(func(i int, v int) {
		set.Add(v)
		if set.Len() != i+1 {
			t.Errorf("Expected set length %d, but got %d", i+1, set.Len())
		}
		idx := set.IndexOf(v)
		if idx != i {
			t.Errorf("Expected element idx of '%d' to be on the same insertion idx "+
				"'%d', but got %d", v, i, idx)
		}
	})

	// assert after deletion the elements cannot be found
	initSize := set.Len()
	for i := initSize - 1; i > 0; i-- {
		set.Delete(i)
		it.RemoveAt(i)
		// assert that the element was deleted
		val, ok := set.Get(i)
		if ok {
			t.Errorf("Expected set to not contain deleted element '%d', but received '%d, "+
				"'%t'", i, val, ok)
		}

		if set.Len() != initSize-(initSize-i) {
			t.Errorf("Expected set length %d, but got %d", initSize-(initSize-i), set.Len())
		}

		// assert each idx is valid
		it.EachN(func(ii int, v int) {
			if idx := set.IndexOf(v); idx != ii { // assert new indices
				t.Errorf("Expected element idx of '%d' to replace idx of deleted "+
					"'%d', but got %d", v, i, idx)
			}
			if val, ok := set.Get(ii); !ok || val != v { // assert not deleted set still valid
				t.Errorf("expected set to contain element '%d', but received '%d, '%t'", v, val, ok)
			}
		})
	}
}

// TestInsSet_Equals tests the equality check between two sets.
//
// this method should state that two sets are compare:
//   - if they have the same number of elements on the same order and value
//   - if they have the same memory address
func TestInsSet_Equals(t *testing.T) {
	cmp := comparators.KeyOrdered[int]{}
	// Test 1: same memory address
	setA := Insertion[int, int](cmp)
	setB := setA

	if !setA.Equals(setB) {
		t.Errorf("Expected sets to be compare, but they are not")
	}

	// Test 2: Both empty, different memory address
	setA = Insertion[int, int](cmp)
	setB = Insertion[int, int](cmp)

	if !setA.Equals(setB) {
		t.Errorf("Expected sets to be not compare, but they are")
	}

	// Test 3: One empty, other nil
	setA = Insertion[int, int](cmp, 1, 2, 3, 5, 6, 7)
	setB = nil

	if setA.Equals(setB) {
		t.Errorf("Expected sets to be not compare, but they are")
	}

	// Test 4: Different number of elements
	setA = Insertion[int, int](cmp, 1, 2, 3, 5, 6, 7)
	setB = Insertion[int, int](cmp, 1, 2, 3, 5, 6, 7, 8)

	if setA.Equals(setB) {
		t.Errorf("Expected sets to be not compare, but they are")
	}

	// Test 5: Different elements
	setA = Insertion[int, int](cmp, 1, 2, 3, 5, 6, 7)
	setB = Insertion[int, int](cmp, 1, 2, 3, 5, 6, 8)

	if setA.Equals(setB) {
		t.Errorf("Expected sets to be not compare, but they are")
	}

	// Test 6: Same elements, different order
	setA = Insertion[int, int](cmp, 1, 2, 3, 5, 6, 7)
	setB = Insertion[int, int](cmp, 5, 6, 7, 1, 2, 3)

	if setA.Equals(setB) {
		t.Errorf("Expected sets to be not compare, but they are")
	}

	// Test 7: Same elements, same order
	setA = Insertion[int, int](cmp, 1, 2, 3, 5, 6, 7)
	setB = Insertion[int, int](cmp, 1, 2, 3, 5, 6, 7)

	if !setA.Equals(setB) {
		t.Errorf("Expected sets to be compare, but they are not")
	}

	// Test 8: Same set after removing element of one of them
	setB.Remove(1)
	if setA.Equals(setB) {
		t.Errorf("Expected sets to be not compare, but they are")
	}
}

// TestInsSet_Get tests the method Get()
//
// after calling 'Get()' the set should:
//   - return the correct element
//   - return true if the element was found
//   - return false if the element was not found
func TestInsSet_Get(t *testing.T) {
	cmp := comparators.KeyOrdered[int]{}

	// Test 1: empty set
	set := Insertion[int, int](cmp)
	if set.Len() != 0 {
		t.Errorf("Expected set length 0, but got %d", set.Len())
	}
	rnd := rand.Int()
	if _, ok := set.Get(rnd); ok {
		t.Errorf("Expected set to not contain %d, but it does", rnd)
	}

	// Test 2: non-empty set
	iterables.OfSlice(1, 2, 3, 5, 6, 7).EachN(func(i int, v int) {
		set.Add(v)
		got, ok := set.Get(i)
		if !ok {
			t.Errorf("Expected set to contain %d, but it does not", v)
		}
		if got != v {
			t.Errorf("Expected retrieved value to be compare inserted '%d' but "+
				"got '%d'", v, got)
		}

		got, ok = set.Get(i + 1)
		if ok {
			t.Errorf("Expected set to not contain %d, but it does", v+1)
		}
		if !generics.IsZero(got) {
			t.Errorf("Expected retrieved value to be 'zero', but got %d", got)
		}
	})
}

// TestInsSet_Has tests the methods:
//
// after calling 'Has()' the set should:
//   - return true if the element was inserted
//   - return false if the element was not inserted
func TestInsSet_Has(t *testing.T) {
	cmp := comparators.KeyOrdered[int]{}

	// Test 1: empty set
	set := Insertion[int, int](cmp)
	if set.Len() != 0 {
		t.Errorf("Expected set length 0, but got %d", set.Len())
	}
	v := rand.Int()
	if set.Has(v) {
		t.Errorf("Expected set to not contain %d, but it does", v)
	}

	// Test 2: non-empty set
	iterables.OfSlice(1, 2, 3, 5, 6, 7).EachN(func(i int, v int) {
		set.Add(v)
		if !set.Has(v) {
			t.Errorf("Expected set to contain %d, but it does not", v)
		}
		if set.Has(v + 1) {
			t.Errorf("Expected set to not contain %d, but it does", v+1)
		}
	})
}

// TestInsSet_IndexOf tests the method IndexOf()
//
// after calling 'IndexOf()' the set should:
//   - return the correct idx if the element should be found
//   - return -1 if the element should not be found
func TestIns_IndexOf(t *testing.T) {
	cmp := comparators.KeyOrdered[int]{}

	// Test 1: empty set
	set := Insertion[int, int](cmp)
	if set.Len() != 0 {
		t.Errorf("Expected set length 0, but got %d", set.Len())
	}
	v := rand.Int()
	if set.IndexOf(v) != -1 {
		t.Errorf("Expected set to not contain %d, but it does", v)
	}

	// Test 2: non-empty set
	iterables.OfSlice(1, 2, 3, 5, 6, 7).EachN(func(i int, v int) {
		set.Add(v)
		if set.IndexOf(v) != i {
			t.Errorf("Expected set to contain %d at idx %d, but it does not", v, i)
		}
		if set.IndexOf(v+1) != -1 {
			t.Errorf("Expected set to not contain %d, but it does", v+1)
		}
	})

	// Test 3: non-empty set since creation
	set = Insertion[int, int](cmp, 1, 2, 3, 5, 6, 7)
	for i := 0; i < 6; i++ {
		if i < 3 {
			if set.IndexOf(i+1) != i {
				t.Errorf("Expected set to contain %d at idx %d, but it does not", i+1, i)
			}
		} else if set.IndexOf(i+2) != i {
			t.Errorf("Expected set to contain %d at idx %d, but it does not", i+2, i)
		}
	}
}

// TestInsSet_IsEmpty tests the methods:
//
// after calling 'IsEmpty()' the set should:
//   - return true if the set is empty
//   - return false if the set is not empty
func TestInsSet_IsEmpty(t *testing.T) {
	cmp := comparators.KeyOrdered[int]{}
	// Test 1: empty set
	set := Insertion[int, int](cmp)

	if !set.IsEmpty() {
		t.Errorf("Expected set to be empty, but it is not")
	}

	// Test 2: non-empty set
	iterables.OfSlice(1, 2, 3, 5, 6, 7).EachN(func(i int, v int) {
		set.Add(v)
		if set.IsEmpty() {
			t.Errorf("Expected set to not be empty, but it is")
		}
		set.Remove(v)
		if !set.IsEmpty() {
			t.Errorf("Expected set to be empty after removing %d, but it is not", v)
		}
	})
}

// TestInsSet_Len tests the methods Len()
//
// after calling 'Len()' the set should:
//   - return the correct number of elements
func TestInsSet_Len(t *testing.T) {
	cmp := comparators.KeyOrdered[int]{}
	// Test 1: empty set
	set := Insertion[int, int](cmp)
	if set.Len() != 0 {
		t.Errorf("Expected set length 0, but got %d", set.Len())
	}

	// Test 2: non-empty set
	iterables.OfSlice(1, 2, 3, 5, 6, 7).EachN(func(i int, v int) {
		set.Add(v)
		if set.Len() != i+1 {
			t.Errorf("Expected set length %d, but got %d", i+1, set.Len())
		}
	})

	// Test 3: non-empty set since creation
	set = Insertion[int, int](cmp, 1, 2, 3, 5, 6, 7)
	if set.Len() != 6 {
		t.Errorf("Expected set length 6, but got %d", set.Len())
	}
}

// TestIns_Remove tests the methods Remove() that removes a value from the set
//
// after calling 'Remove()' the set should:
//   - remove the value from the set
//   - return the correct number of elements
//   - not be able to find the removed value anymore
//   - retrieve the new correct indexes of the elements
func TestIns_Remove(t *testing.T) {
	cmp := comparators.KeyOrdered[int]{}

	// Test 1: empty set
	set := Insertion[int, int](cmp)
	if set.Len() != 0 {
		t.Errorf("Expected set length 0, but got %d", set.Len())
	}
	set.Remove(1)
	if set.Len() != 0 {
		t.Errorf("Expected set length 0, but got %d", set.Len())
	}

	// Test 2: non-empty set
	it := iterables.OfSlice(1, 2, 3, 5, 6, 7)
	it.EachN(func(i int, v int) {
		set.Add(v)
		if set.Len() != 1 {
			t.Errorf("Expected set length %d, but got %d", i+1, set.Len())
		}
		// expect to find element at idx 0
		e, ok := set.Get(0)
		if !ok {
			t.Errorf("Expected set to contain element '%d' at idx 0, but it does not", v)
		}
		if e != v {
			t.Errorf("Expected retrieved value at idx 0 to be compare inserted '%d' but "+
				"got '%d'", v, e)
		}
		set.Remove(v)
		// expect to not found idx of element
		idx := set.IndexOf(v)
		if idx != -1 {
			t.Errorf("Expected set to retrieve idx 0 for element '%d', but got %d", v, idx)
		}
		if set.Len() != 0 {
			t.Errorf("Expected set length to be 0 after removing '%d', but got %d", i,
				set.Len())
		}
		// Expect to not find removed element at idx 0
		e, ok = set.Get(0)
		if ok {
			t.Errorf("Expected set to not contain element '%d' at idx 0, but it retrieved '%t", v, ok)
		}
		if !generics.IsZero(e) {
			t.Errorf("Expected retrieved value at idx 0 to be compare zero but got '%d'", e)
		}
		// expect to retrieve idx -1 for removed element
		idx = set.IndexOf(v)
		if idx != -1 {
			t.Errorf("Expected set to retrieve idx -1 for element '%d', but got %d", v, idx)
		}
	})

	// Test 3: non-empty set since creation
	set = Insertion[int, int](cmp, 1, 2, 3, 5, 6, 7)
	if set.Len() != 6 {
		t.Errorf("Expected set length 6, but got %d", set.Len())
	}

	// Expect to find each element on its updated positions
	for !it.IsEmpty() {
		v := it.At(0)
		// expect to find element at the same insertion position
		idx := set.IndexOf(v)
		if idx != 0 {
			t.Errorf("Expected set to contain element '%d' at idx 0, but it does not", v)
		}
		// expect to find element at idx 0
		e, ok := set.Get(0)
		if !ok {
			t.Errorf("Expected set to contain element '%d' at idx 0, but it does not", v)
		}
		if e != v {
			t.Errorf("Expected retrieved value at idx 0 to be compare inserted '%d' but "+
				"got '%d'", v, e)
		}
		set.Remove(v)
		it.RemoveAt(0)

		// expect size to decrease
		if set.Len() != it.Len() {
			t.Errorf("Expected set length %d, but got %d", it.Len(), set.Len())
		}

		// expect other elements to keep their positions
		it.EachN(func(i int, v int) {
			e, ok := set.Get(i)
			if !ok {
				t.Errorf("Expected set to contain element '%d' at idx %d, but it does not",
					v, i)
			}
			if e != v {
				t.Errorf("Expected retrieved value to be compare inserted '%d' but "+
					"got '%d'", v, e)
			}

			idx := set.IndexOf(v)
			if idx != i {
				t.Errorf("Expected set to retrieve idx %d for element '%d', but got %d", i, v, idx)
			}
		})
	}
}

// TestInsSet_Values tests the method Values() that retrieves all elements on the set
func TestInsSet_Values(t *testing.T) {
	cmp := comparators.KeyOrdered[int]{}

	// Test empty set
	set := Insertion[int, int](cmp)
	values := set.Values()
	if len(values) != 0 {
		t.Errorf("Expected set length 0, but got %d", len(values))
	}

	// Test non-empty set
	it := iterables.OfSlice(1, 2, 3, 5, 6, 7)
	it.EachN(func(i int, v int) {
		set.Add(v)
		values := set.Values()
		if len(values) != i+1 {
			t.Errorf("Expected set length %d, but got %d", i+1, len(values))
		}
		subset := it.Before(i + 1)
		for j, e := range values {
			if e != subset.At(j) {
				t.Errorf("Expected set value '%d' at idx %d, but got %d", subset.At(j),
					j, e)
			}
		}
	})
}

// TestConcCmpSet_Creation tests the creation of a concurrent insertion set
//
// after the creation the set should:
//   - have the correct number of elements
//   - have the correct number of indexes (ordered by insertion)
func TestConcCmpSet_Creation(t *testing.T) {
	// Test empty set
	cmp := comparators.KeyOrdered[int]{}
	set := ConcInsertion[int, int](cmp)

	if set.Len() != 0 {
		t.Errorf("Expected set length 0, but got %d", set.Len())
	}

	// test set with elements
	set = ConcInsertion[int, int](cmp, 1, 2, 3, 5, 6, 7)

	if set.Len() != 6 {
		t.Errorf("Expected set length 6, but got %d", set.Len())
	}

	for i, v := range set.Values() {
		if i < 3 {
			if v != i+1 {
				t.Errorf("Expected value '%d' at idx %d, but got %d", i+1, i, v)
			}
		} else if v != i+2 {
			t.Errorf("Expected set value '%d' at idx %d, but got %d", i+2, i, v)
		}
	}

	// test set with duplicate elements
	set = ConcInsertion[int, int](cmp, 1, 1, 2, 2, 3, 3, 5, 5, 6, 6, 7, 7)
	if set.Len() != 6 {
		t.Errorf("Expected set length 6, but got %d", set.Len())
	}

	for i, v := range set.Values() {
		if i < 3 {
			if v != i+1 {
				t.Errorf("Expected set value '%d' at idx %d, but got %d", i+1, i, v)
			}
		} else if v != i+2 {
			t.Errorf("Expected set value '%d' at idx %d, but got %d", i+2, i, v)
		}
	}
}

// TestConcIns_Add tests the insertion set states after adding elements
//
// after adding elements the set should:
//   - have the correct number of elements
//   - have the correct number of indexes (ordered by insertion)
//   - not include duplicates
func TestConcIns_Add(t *testing.T) {
	cmp := comparators.KeyOrdered[int]{}
	// Test with empty set
	set := ConcInsertion[int, int](cmp)

	if set.Len() != 0 {
		t.Errorf("Expected set length 0, but got %d", set.Len())
	}

	it := iterables.OfSlice(1, 2, 3, 5, 6, 7)
	it.EachN(func(i int, v int) {
		set.Add(v)
		if set.Len() != i+1 {
			t.Errorf("Expected set length %d after adding '%d' "+
				"elements, but got %d", i+1, i, set.Len())
		}
	})

	// Test with non-empty set
	set = ConcInsertion[int, int](cmp, 1, 2, 3, 5, 6, 7)

	if set.Len() != 6 {
		t.Errorf("Expected set length 6, but got %d", set.Len())
	}

	it.EachN(func(i int, v int) {
		initSize := set.Len()
		set.Add(v)
		if it.At(i) == 4 {
			if set.Len() != 7 {
				t.Errorf("Expected set length %d after adding missing element "+
					"'%d' element, but got %d", 7, 4, set.Len())
			}
		} else {
			if set.Len() != initSize {
				t.Errorf("Expected set length remain the same '%d' after adding duplicate "+
					"element '%d' element, but got %d", initSize, 4, set.Len())
			}
		}
	})
}

// TestConcIns_Clear tests the insertion set states after clearing elements
//
// after calling 'Clear()' the set should:
//   - have no elements
//   - return true when calling 'IsEmpty()'
func TestConcIns_Clear(t *testing.T) {
	cmp := comparators.KeyOrdered[int]{}
	set := ConcInsertion[int, int](cmp, 1, 2, 3, 5, 6, 7)

	if set.Len() != 6 {
		t.Errorf("Expected set length 6, but got %d", set.Len())
	}

	set.Clear()

	if set.Len() != 0 {
		t.Errorf("Expected set length 0 after clear, but got %d", set.Len())
	}
}

// TestConcIns_Delete tests the insertion set states after deleting elements
//
// after calling 'Delete(idx)' the set should:
//   - have the correct number of elements
//   - have the correct number of indexes (ordered by insertion)
//   - do not include the deleted element
func TestConcIns_Delete(t *testing.T) {
	cmp := comparators.KeyOrdered[int]{}
	// Test 1: empty set
	set := ConcInsertion[int, int](cmp)

	if set.Len() != 0 {
		t.Errorf("Expected set length 0, but got %d", set.Len())
	}
	set.Delete(1)
	if set.Len() != 0 {
		t.Errorf("Expected set length 0, but got %d", set.Len())
	}

	// Test 2: non-empty set, delete after each insertion
	it := iterables.OfSlice(1, 2, 3, 5, 6, 7)
	it.EachN(func(i int, v int) {
		set.Add(v)
		if set.Len() != 1 {
			t.Errorf("Expected set length '%d', but got '%d'", 1, set.Len())
		}
		set.Delete(0)
		if set.Len() != 0 {
			t.Errorf("Expected set length return to 0 after delete inserted element '%d', "+
				"but got %d", v, set.Len())
		}
	})

	if set.Len() != 0 {
		t.Errorf("Expected set length 0, but got %d", set.Len())
	}

	// Test 3: non-empty set, delete after all insertions
	it.EachN(func(i int, v int) {
		set.Add(v)
		if set.Len() != i+1 {
			t.Errorf("Expected set length %d, but got %d", i+1, set.Len())
		}
		idx := set.IndexOf(v)
		if idx != i {
			t.Errorf("Expected element idx of '%d' to be on the same insertion idx "+
				"'%d', but got %d", v, i, idx)
		}
	})

	// assert after deletion the elements cannot be found
	initSize := set.Len()
	for i := initSize - 1; i > 0; i-- {
		set.Delete(i)
		it.RemoveAt(i)
		// assert that the element was deleted
		val, ok := set.Get(i)
		if ok {
			t.Errorf("Expected set to not contain deleted element '%d', but received '%d, "+
				"'%t'", i, val, ok)
		}

		if set.Len() != initSize-(initSize-i) {
			t.Errorf("Expected set length %d, but got %d", initSize-(initSize-i), set.Len())
		}

		// assert each idx is valid
		it.EachN(func(ii int, v int) {
			if idx := set.IndexOf(v); idx != ii { // assert new indices
				t.Errorf("Expected element idx of '%d' to replace idx of deleted "+
					"'%d', but got %d", v, i, idx)
			}
			if val, ok := set.Get(ii); !ok || val != v { // assert not deleted set still valid
				t.Errorf("expected set to contain element '%d', but received '%d, '%t'", v, val, ok)
			}
		})
	}
}

// TestConcIns_Equals tests the Equals method
//
// this method should state that two sets are compare:
//   - if they have the same number of elements on the same order and value
//   - if they have the same memory address
func TestConcIns_Equals(t *testing.T) {
	cmp := comparators.KeyOrdered[int]{}
	// Test 1: same memory address
	setA := ConcInsertion[int, int](cmp)
	setB := setA

	if !setA.Equals(setB) {
		t.Errorf("Expected sets to be compare, but they are not")
	}

	// Test 2: Both empty, different memory address
	setA = ConcInsertion[int, int](cmp)
	setB = ConcInsertion[int, int](cmp)

	if !setA.Equals(setB) {
		t.Errorf("Expected sets to be not compare, but they are")
	}

	// Test 3: One empty, other nil
	setA = ConcInsertion[int, int](cmp, 1, 2, 3, 5, 6, 7)
	setB = nil

	if setA.Equals(setB) {
		t.Errorf("Expected sets to be not compare, but they are")
	}

	// Test 4: Different number of elements
	setA = ConcInsertion[int, int](cmp, 1, 2, 3, 5, 6, 7)
	setB = ConcInsertion[int, int](cmp, 1, 2, 3, 5, 6, 7, 8)

	if setA.Equals(setB) {
		t.Errorf("Expected sets to be not compare, but they are")
	}

	// Test 5: Different elements
	setA = ConcInsertion[int, int](cmp, 1, 2, 3, 5, 6, 7)
	setB = ConcInsertion[int, int](cmp, 1, 2, 3, 5, 6, 8)

	if setA.Equals(setB) {
		t.Errorf("Expected sets to be not compare, but they are")
	}

	// Test 6: Same elements, different order
	setA = ConcInsertion[int, int](cmp, 1, 2, 3, 5, 6, 7)
	setB = ConcInsertion[int, int](cmp, 5, 6, 7, 1, 2, 3)

	if setA.Equals(setB) {
		t.Errorf("Expected sets to be compare, but they are not")
	}

	// Test 7: Same elements, same order
	setA = ConcInsertion[int, int](cmp, 1, 2, 3, 5, 6, 7)
	setB = ConcInsertion[int, int](cmp, 1, 2, 3, 5, 6, 7)

	if !setA.Equals(setB) {
		t.Errorf("Expected sets to be compare, but they are not")
	}

	// Test 8: Same maps after removing elements
	setB.Remove(1)
	if setA.Equals(setB) {
		t.Errorf("Expected sets to be not compare, but they are")
	}
}

// TestConcIns_Get tests the Get method
//
// after calling 'Get()' the set should:
//   - return the correct element
//   - return true if the element was found
//   - return false if the element was not found
func TestConcIns_Get(t *testing.T) {
	cmp := comparators.KeyOrdered[int]{}

	// Test 1: empty set
	set := ConcInsertion[int, int](cmp)
	if set.Len() != 0 {
		t.Errorf("Expected set length 0, but got %d", set.Len())
	}
	rnd := rand.Int()
	if _, ok := set.Get(rnd); ok {
		t.Errorf("Expected set to not contain %d, but it does", rnd)
	}

	// Test 2: non-empty set
	iterables.OfSlice(1, 2, 3, 5, 6, 7).EachN(func(i int, v int) {
		set.Add(v)
		got, ok := set.Get(i)
		if !ok {
			t.Errorf("Expected set to contain %d, but it does not", v)
		}
		if got != v {
			t.Errorf("Expected retrieved value to be compare inserted '%d' but "+
				"got '%d'", v, got)
		}

		got, ok = set.Get(i + 1)
		if ok {
			t.Errorf("Expected set to not contain %d, but it does", v+1)
		}
		if !generics.IsZero(got) {
			t.Errorf("Expected retrieved value to be 'zero', but got %d", got)
		}
	})
}

// TestConcIns_Has tests the Has method
//
// after calling 'Has()' the set should:
//   - return true if the element was inserted
//   - return false if the element was not inserted
func TestConcIns_Has(t *testing.T) {
	cmp := comparators.KeyOrdered[int]{}

	// Test 1: empty set
	set := ConcInsertion[int, int](cmp)
	if set.Len() != 0 {
		t.Errorf("Expected set length 0, but got %d", set.Len())
	}
	v := rand.Int()
	if set.Has(v) {
		t.Errorf("Expected set to not contain %d, but it does", v)
	}

	// Test 2: non-empty set
	iterables.OfSlice(1, 2, 3, 5, 6, 7).EachN(func(i int, v int) {
		set.Add(v)
		if !set.Has(v) {
			t.Errorf("Expected set to contain %d, but it does not", v)
		}
		if set.Has(v + 1) {
			t.Errorf("Expected set to not contain %d, but it does", v+1)
		}
	})
}

// TestConcIns_IndexOf tests the IndexOf method
//
// after calling 'IndexOf()' the set should:
//   - return the correct idx if the element should be found
//   - return -1 if the element should not be found
func TestConcIns_IndexOf(t *testing.T) {
	cmp := comparators.KeyOrdered[int]{}

	// Test 1: empty set
	set := ConcInsertion[int, int](cmp)
	if set.Len() != 0 {
		t.Errorf("Expected set length 0, but got %d", set.Len())
	}
	v := rand.Int()
	if set.IndexOf(v) != -1 {
		t.Errorf("Expected set to not contain %d, but it does", v)
	}

	// Test 2: non-empty set
	iterables.OfSlice(1, 2, 3, 5, 6, 7).EachN(func(i int, v int) {
		set.Add(v)
		if set.IndexOf(v) != i {
			t.Errorf("Expected set to contain %d at idx %d, but it does not", v, i)
		}
		if set.IndexOf(v+1) != -1 {
			t.Errorf("Expected set to not contain %d, but it does", v+1)
		}
	})

	// Test 3: non-empty set since creation
	set = Insertion[int, int](cmp, 1, 2, 3, 5, 6, 7)
	for i := 0; i < 6; i++ {
		if i < 3 {
			if set.IndexOf(i+1) != i {
				t.Errorf("Expected set to contain %d at idx %d, but it does not", i+1, i)
			}
		} else if set.IndexOf(i+2) != i {
			t.Errorf("Expected set to contain %d at idx %d, but it does not", i+2, i)
		}
	}
}

// TestConcIns_IsEmpty tests the IsEmpty method
//
// after calling 'IsEmpty()' the set should:
//   - return true if the set is empty
//   - return false if the set is not empty
func TestConcIns_IsEmpty(t *testing.T) {
	cmp := comparators.KeyOrdered[int]{}
	// Test 1: empty set
	set := ConcInsertion[int, int](cmp)

	if !set.IsEmpty() {
		t.Errorf("Expected set to be empty, but it is not")
	}

	// Test 2: non-empty set
	iterables.OfSlice(1, 2, 3, 5, 6, 7).EachN(func(i int, v int) {
		set.Add(v)
		if set.IsEmpty() {
			t.Errorf("Expected set to not be empty, but it is")
		}
		set.Remove(v)
		if !set.IsEmpty() {
			t.Errorf("Expected set to be empty after removing %d, but it is not", v)
		}
	})
}

// TestConcIns_Len tests the Len method
//
// after calling 'Len()' the set should:
//   - return the correct number of elements
func TestConcIns_Len(t *testing.T) {
	cmp := comparators.KeyOrdered[int]{}
	// Test 1: empty set
	set := ConcInsertion[int, int](cmp)
	if set.Len() != 0 {
		t.Errorf("Expected set length 0, but got %d", set.Len())
	}

	// Test 2: non-empty set
	iterables.OfSlice(1, 2, 3, 5, 6, 7).EachN(func(i int, v int) {
		set.Add(v)
		if set.Len() != i+1 {
			t.Errorf("Expected set length %d, but got %d", i+1, set.Len())
		}
	})

	// Test 3: non-empty set since creation
	set = ConcInsertion[int, int](cmp, 1, 2, 3, 5, 6, 7)
	if set.Len() != 6 {
		t.Errorf("Expected set length 6, but got %d", set.Len())
	}
}

// TestConcIns_Remove tests the Remove method
//
// after calling 'Remove()' the set should:
//   - remove the value from the set
//   - return the correct number of elements
//   - not be able to find the removed value anymore
//   - retrieve the new correct indexes of the elements
func TestConcIns_Remove(t *testing.T) {
	cmp := comparators.KeyOrdered[int]{}

	// Test 1: empty set
	set := ConcInsertion[int, int](cmp)
	if set.Len() != 0 {
		t.Errorf("Expected set length 0, but got %d", set.Len())
	}
	set.Remove(1)
	if set.Len() != 0 {
		t.Errorf("Expected set length 0, but got %d", set.Len())
	}

	// Test 2: non-empty set
	it := iterables.OfSlice(1, 2, 3, 5, 6, 7)
	it.EachN(func(i int, v int) {
		set.Add(v)
		if set.Len() != 1 {
			t.Errorf("Expected set length %d, but got %d", i+1, set.Len())
		}
		// expect to find element at idx 0
		e, ok := set.Get(0)
		if !ok {
			t.Errorf("Expected set to contain element '%d' at idx 0, but it does not", v)
		}
		if e != v {
			t.Errorf("Expected retrieved value at idx 0 to be compare inserted '%d' but "+
				"got '%d'", v, e)
		}
		set.Remove(v)
		// expect to not found idx of element
		idx := set.IndexOf(v)
		if idx != -1 {
			t.Errorf("Expected set to retrieve idx 0 for element '%d', but got %d", v, idx)
		}
		if set.Len() != 0 {
			t.Errorf("Expected set length to be 0 after removing '%d', but got %d", i,
				set.Len())
		}
		// Expect to not find removed element at idx 0
		e, ok = set.Get(0)
		if ok {
			t.Errorf("Expected set to not contain element '%d' at idx 0, but it retrieved '%t", v, ok)
		}
		if !generics.IsZero(e) {
			t.Errorf("Expected retrieved value at idx 0 to be compare zero but got '%d'", e)
		}
		// expect to retrieve idx -1 for removed element
		idx = set.IndexOf(v)
		if idx != -1 {
			t.Errorf("Expected set to retrieve idx -1 for element '%d', but got %d", v, idx)
		}
	})

	// Test 3: non-empty set since creation
	set = ConcInsertion[int, int](cmp, 1, 2, 3, 5, 6, 7)
	if set.Len() != 6 {
		t.Errorf("Expected set length 6, but got %d", set.Len())
	}

	// Expect to find each element on its updated positions
	for !it.IsEmpty() {
		v := it.At(0)
		// expect to find element at the same insertion position
		idx := set.IndexOf(v)
		if idx != 0 {
			t.Errorf("Expected set to contain element '%d' at idx 0, but it does not", v)
		}
		// expect to find element at idx 0
		e, ok := set.Get(0)
		if !ok {
			t.Errorf("Expected set to contain element '%d' at idx 0, but it does not", v)
		}
		if e != v {
			t.Errorf("Expected retrieved value at idx 0 to be compare inserted '%d' but "+
				"got '%d'", v, e)
		}
		set.Remove(v)
		it.RemoveAt(0)

		// expect size to decrease
		if set.Len() != it.Len() {
			t.Errorf("Expected set length %d, but got %d", it.Len(), set.Len())
		}

		// expect other elements to keep their positions
		it.EachN(func(i int, v int) {
			e, ok := set.Get(i)
			if !ok {
				t.Errorf("Expected set to contain element '%d' at idx %d, but it does not",
					v, i)
			}
			if e != v {
				t.Errorf("Expected retrieved value to be compare inserted '%d' but "+
					"got '%d'", v, e)
			}

			idx := set.IndexOf(v)
			if idx != i {
				t.Errorf("Expected set to retrieve idx %d for element '%d', but got %d", i, v, idx)
			}
		})
	}
}

// TestInsSet_Values tests the method Values() that retrieves all elements on the set
func TestConcIns_Values(t *testing.T) {
	cmp := comparators.KeyOrdered[int]{}

	// Test empty set
	set := ConcInsertion[int, int](cmp)
	values := set.Values()
	if len(values) != 0 {
		t.Errorf("Expected set length 0, but got %d", len(values))
	}

	// Test non-empty set
	it := iterables.OfSlice(1, 2, 3, 5, 6, 7)
	it.EachN(func(i int, v int) {
		set.Add(v)
		values := set.Values()
		if len(values) != i+1 {
			t.Errorf("Expected set length %d, but got %d", i+1, len(values))
		}
		subset := it.Before(i + 1)
		for j, e := range values {
			if e != subset.At(j) {
				t.Errorf("Expected set value '%d' at idx %d, but got %d", subset.At(j),
					j, e)
			}
		}
	})
}
