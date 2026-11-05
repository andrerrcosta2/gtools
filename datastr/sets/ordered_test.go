// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sets

import (
	"cmp"
	"testing"

	"github.com/andrerrcosta2/gtools/core/data/keyer/keyers"
	"github.com/andrerrcosta2/gtools/core/data/str/iterables"
	"github.com/andrerrcosta2/gtools/gflux/core/pipes/arrays"
)

// TestOrdered_Creation tests the creation of an ordered backingMap.
//
// after creating the backingMap it must assert:
//   - It is not nil
//   - Its size is the same as the number of inserted elements
func TestOrdered_Creation(t *testing.T) {
	// Test 1: empty backingMap
	t.Run("create empty backingMap", func(t *testing.T) {
		set := Ordered[int]()
		if set == nil {
			t.Errorf("Expected backingMap to be not nil, but it is")
		}
		if set.Len() != 0 {
			t.Errorf("Expected backingMap length 0, but got %d", set.Len())
		}
	})

	// Test 2: backingMap with elements
	t.Run("create backingMap with elements", func(t *testing.T) {
		set := Ordered[int](1, 2, 3, 5, 6, 7)

		if set == nil {
			t.Errorf("Expected backingMap to be not nil, but it is")
		}
		if set.Len() != 6 {
			t.Errorf("Expected backingMap length 6, but got %d", set.Len())
		}

		iterables.OfSlice(1, 2, 3, 5, 6, 7).
			EachN(func(i int, v int) {
				if set.IndexOf(v) != i {
					t.Errorf("Expected backingMap to contain '%d' at index '%d', but it "+
						"doesn't", v, i)
				}
			})
	})

	// Test 3: backingMap with unordered elements
	t.Run("create backingMap with unordered elements", func(t *testing.T) {
		set := Ordered[int](7, 6, 5, 3, 2, 1)
		if set.Len() != 6 {
			t.Errorf("Expected backingMap length 6, but got %d", set.Len())
		}
		iterables.OfSlice(1, 2, 3, 5, 6, 7).
			EachN(func(i int, v int) {
				if set.IndexOf(v) != i {
					t.Errorf("Expected backingMap to contain '%d' at index '%d', but it "+
						"doesn't", v, i)
				}
			})
	})

	// Test 4: backingMap with duplicated unordered elements
	t.Run("create backingMap with duplicated unordered elements", func(t *testing.T) {
		set := Ordered[int](7, 7, 6, 6, 5, 5, 3, 3, 2, 2, 1, 1)
		if set.Len() != 6 {
			t.Errorf("Expected backingMap length 6, but got %d", set.Len())
		}
		iterables.OfSlice(1, 2, 3, 5, 6, 7).
			EachN(func(i int, v int) {
				if set.IndexOf(v) != i {
					t.Errorf("Expected backingMap to contain '%d' at index '%d', but it "+
						"doesn't", v, i)
				}
			})
	})
}

// TestOrdered_Add tests the add method of an ordered backingMap.
//
// after adding an element to the backingMap it must assert:
//   - The element is in the backingMap slice
//   - The element is in the idx map
//   - The backingMap size is the same as the number of inserted elements
//   - The order of elements is the same as the natural comparable order
func TestOrdered_Add(t *testing.T) {
	// Test 1: start empty, add ordered
	t.Run("start empty, add ordered", func(t *testing.T) {
		set := Ordered[int]()
		if set.Len() != 0 {
			t.Errorf("Expected backingMap length 0, but got %d", set.Len())
		}
		iterables.OfSlice(1, 2, 3, 5, 6, 7).
			EachN(func(i int, v int) {
				set.Add(v)
				if set.Len() != i+1 {
					t.Errorf("Expected backingMap length %d, but got %d", i+1, set.Len())
				}
			}).
			EachN(func(i int, v int) {
				if set.IndexOf(v) != i {
					t.Errorf("Expected backingMap to contain '%d' at index '%d', but it "+
						"doesn't", v, i)
				}
			})
	})

	// Test 2: start empty, add unordered
	t.Run("start empty, add unordered", func(t *testing.T) {
		set := Ordered[int]()
		if set.Len() != 0 {
			t.Errorf("Expected backingMap length 0, but got %d", set.Len())
		}
		it := iterables.OfSlice(7, 6, 5, 3, 2, 1)
		it.EachN(func(i int, v int) {
			set.Add(v)
			if set.Len() != i+1 {
				t.Errorf("Expected backingMap length %d, but got %d", i+1, set.Len())
			}
		})
		it.Sort(cmp.Less[int]).
			EachN(func(i int, v int) {
				if set.IndexOf(v) != i {
					t.Errorf("Expected backingMap to contain '%d' at index '%d', but it "+
						"doesn't", v, i)
				}
			})
	})

	// Test 3: start empty, add unordered with duplicates
	t.Run("start empty, add unordered with duplicates", func(t *testing.T) {
		set := Ordered[int]()
		if set.Len() != 0 {
			t.Errorf("Expected backingMap length 0, but got %d", set.Len())
		}
		it := iterables.OfSlice(7, 7, 6, 6, 5, 5, 3, 3, 2, 2, 1, 1)
		it.EachN(func(i int, v int) {
			set.Add(v)
			if set.Len() != (i)/2+1 {
				t.Errorf("Expected backingMap length '%d' at index '%d', but got %d", (i)/2+1, i, set.Len())
			}
		})
		it.ToSet(keyers.FingerPrint[int]()).
			Sort(cmp.Less[int]).
			EachN(func(i int, v int) {
				if idx := set.IndexOf(v); idx != i {
					t.Errorf("Expected backingMap to contain '%d' at index '%d', but it "+
						"returned %d", v, i, idx)
				}
			})
	})
}

// TestOrdered_Clear tests the clear method of an ordered backingMap
//
// after calling 'Clear()' the backingMap should:
//   - have no elements
//   - return true when calling 'IsEmpty()'
func TestOrdered_Clear(t *testing.T) {
	// Test 1: clear empty backingMap
	t.Run("clear empty backingMap", func(t *testing.T) {
		set := Ordered[int]()
		if set.Len() != 0 {
			t.Errorf("Expected backingMap length 0, but got %d", set.Len())
		}
		if !set.IsEmpty() {
			t.Errorf("Expected backingMap to be empty, but it is not")
		}
		set.Clear()
		if set.Len() != 0 {
			t.Errorf("Expected backingMap length 0, but got %d", set.Len())
		}
		if !set.IsEmpty() {
			t.Errorf("Expected backingMap to be empty, but it is not")
		}
	})

	// Test 2: clear after add backingMap
	t.Run("clear after add backingMap", func(t *testing.T) {
		set := Ordered[int]()
		iterables.OfSlice(1, 2, 3, 5, 6, 7).
			EachN(func(i int, v int) {
				set.Add(v)
				if set.Len() != i+1 {
					t.Errorf("Expected backingMap length %d, but got %d", i+1, set.Len())
				}
			})
		if set.Len() != 6 {
			t.Errorf("Expected backingMap length 6, but got %d", set.Len())
		}
		if set.IsEmpty() {
			t.Errorf("Expected backingMap to not be empty, but it is")
		}
		set.Clear()
		if set.Len() != 0 {
			t.Errorf("Expected backingMap length 0, but got %d", set.Len())
		}
		if !set.IsEmpty() {
			t.Errorf("Expected backingMap to be empty, but it is not")
		}
	})

	// Test 3: clear after each add
	t.Run("clear after each add", func(t *testing.T) {
		set := Ordered[int]()
		iterables.OfSlice(1, 2, 3, 5, 6, 7).
			EachN(func(i int, v int) {
				set.Add(v)
				if set.IsEmpty() {
					t.Errorf("Expected backingMap to not be empty, but it is")
				}
				if set.Len() != 1 {
					t.Errorf("Expected backingMap length 1, but got %d", set.Len())
				}
				set.Clear()
				if set.Len() != 0 {
					t.Errorf("Expected backingMap length 0, but got %d", set.Len())
				}
				if !set.IsEmpty() {
					t.Errorf("Expected backingMap to be empty, but it is")
				}
			})
	})

	// Test 4: clear not empty backingMap
	t.Run("clear not empty backingMap", func(t *testing.T) {
		set := Ordered[int](1, 2, 3, 5, 6, 7)
		if set.Len() != 6 {
			t.Errorf("Expected backingMap length 6, but got %d", set.Len())
		}
		if set.IsEmpty() {
			t.Errorf("Expected backingMap to not be empty, but it is")
		}
		set.Clear()
		if set.Len() != 0 {
			t.Errorf("Expected backingMap length 0, but got %d", set.Len())
		}
		if !set.IsEmpty() {
			t.Errorf("Expected backingMap to be empty, but it is")
		}
	})
}

// TestOrdered_Delete tests the delete method of an ordered backingMap
//
// after calling 'Delete()' the backingMap should:
//   - have one less element
//   - return false when calling 'Has()'
//   - not find the value on the underlying map
func TestOrdered_Delete(t *testing.T) {
	// Test 1: delete from empty backingMap
	t.Run("deleting from empty backingMap", func(t *testing.T) {
		set := Ordered[int]()
		cast, ok := set.(*ordered[int, int])
		if !ok {
			t.Errorf("Expected backingMap to be of type '%T', but got '%T'", cast, set)
		}
		if set.Len() != 0 {
			t.Errorf("Expected backingMap length 0, but got %d", set.Len())
		}
		if !set.IsEmpty() {
			t.Errorf("Expected backingMap to be empty, but it is not")
		}
		if len(cast.table) != set.Len() {
			t.Errorf("Expected slice length 0, but got %d", len(cast.table))
		}
		set.Delete(1)
		if set.Len() != 0 {
			t.Errorf("Expected backingMap length 0, but got %d", set.Len())
		}
		if !set.IsEmpty() {
			t.Errorf("Expected backingMap to be empty, but it is not")
		}
		if len(cast.table) != set.Len() {
			t.Errorf("Expected slice length 0, but got %d", len(cast.table))
		}
	})

	// Test 2: delete from non-empty backingMap
	t.Run("deleting from non-empty backingMap", func(t *testing.T) {
		set := Ordered[int](1, 2, 3, 5, 6, 7)
		cast, ok := set.(*ordered[int, int])
		if !ok {
			t.Errorf("Expected backingMap to be of type '%T', but got '%T'", cast, set)
		}
		if set.Len() != 6 {
			t.Errorf("Expected backingMap length 6, but got %d", set.Len())
		}
		if set.IsEmpty() {
			t.Errorf("Expected backingMap to not be empty, but it is")
		}
		if len(cast.table) != set.Len() {
			t.Errorf("Expected slice length 6, but got %d", len(cast.table))
		}

		it := iterables.OfSlice(1, 2, 3, 5, 6, 7)
		it.EachN(func(i int, v int) {
			if !set.Delete(0) {
				t.Errorf("Loop[%d]: expected to delete %d, but it wasn't", i, v)
			}
			it.RemoveAt(0)
			if set.Len() != it.Len() {
				t.Errorf("Loop[%d]: Expected backingMap length %d, but got %d", i, it.Len(), set.Len())
			}
			if len(cast.table) != set.Len() {
				t.Errorf("Loop[%d]: Expected slice length %d, but got %d", i, it.Len(), len(cast.table))
			}
			if cast.tree.Size() != set.Len() {
				t.Errorf("Loop[%d]: Expected tree size %d, but got %d", i, it.Len(), cast.tree.Size())
			}
		})
	})

}

// TestOrdered_Equals tests the compare method of an ordered backingMap
//
// after calling 'Equals()' the backingMap should:
//   - return true when the sets have the same address
//   - return true when the sets contain the same backingMap (order-agnostic)
//   - return false when the sets do not contain the same backingMap
func TestOrdered_Equals(t *testing.T) {
	// Test 1: compare address
	t.Run("compare address", func(t *testing.T) {
		set1 := Ordered[int](1, 2, 3, 5, 6, 7)
		set2 := set1
		if !set1.Equals(set2) {
			t.Errorf("Expected sets to be compare, but they are not")
		}
	})

	// Test 2: compare backingMap
	t.Run("compare backingMap", func(t *testing.T) {
		set1 := Ordered[int](1, 2, 3, 5, 6, 7)
		set2 := Ordered[int](1, 2, 3, 5, 6, 7)
		if !set1.Equals(set2) {
			t.Errorf("Expected sets to be compare, but they are not")
		}
	})

	// Test 3: not compare backingMap
	t.Run("not compare backingMap", func(t *testing.T) {
		set1 := Ordered[int](1, 2, 3, 5, 6, 7)
		set2 := Ordered[int](1, 2, 3, 5, 6, 8)
		if set1.Equals(set2) {
			t.Errorf("Expected sets to be not compare, but they are")
		}
	})

	// Test 4: compare backingMap, different order
	t.Run("compare backingMap, different order", func(t *testing.T) {
		set1 := Ordered[int](1, 2, 3, 5, 6, 7)
		set2 := Ordered[int](3, 5, 6, 7, 1, 2)
		if !set1.Equals(set2) {
			t.Errorf("Expected sets to be compare, but they are not")
		}
	})
}

func TestOrdered_Get(t *testing.T) {
	// Test 1: empty backingMap
	t.Run("empty backingMap", func(t *testing.T) {
		set := Ordered[int]()
		_, ok := set.Get(1)
		if ok {
			t.Errorf("Expected backingMap to not contain '1', but it does")
		}
	})

	// Test 2: non-empty backingMap
	t.Run("non-empty backingMap", func(t *testing.T) {
		set := Ordered[int](1, 2, 3, 5, 6, 7)
		for i := 1; i < 6; i++ {
			v, ok := set.Get(i)
			if !ok {
				t.Errorf("Expected backingMap to contain '%d', but it doesn't", i)
			}
			if i >= 3 {
				if v != i+2 {
					t.Errorf("Expected backingMap to contain '%d' at index '%d', but it "+
						"returned '%d'", i+2, i, v)
				}
			} else if v != i+1 {
				t.Errorf("Expected backingMap to contain '%d' at index '%d', but it "+
					"returned '%d'", i+2, i, v)
			}
		}
	})
}

func TestOrdered_Has(t *testing.T) {
	// Test 1: empty backingMap
	t.Run("empty backingMap", func(t *testing.T) {
		set := Ordered[int]()
		if set.Has(1) {
			t.Errorf("Expected backingMap to not contain '1', but it does")
		}
	})

	// Test 2: non-empty backingMap
	t.Run("non-empty backingMap", func(t *testing.T) {
		set := Ordered[int](1, 2, 3, 5, 6, 7)
		for i := 1; i <= 7; i++ {
			if i != 4 {
				if !set.Has(i) {
					t.Errorf("Expected backingMap to contain '%d', but it doesn't", i)
				}
			} else if set.Has(i) {
				t.Errorf("Expected backingMap to not contain '%d', but it does", i)
			}
		}
	})
}

// TestOrdered_IsEmpty tests the isEmpty method of an ordered backingMap
//
// after calling 'IsEmpty()' the backingMap should:
//   - return true if the backingMap is empty
//   - return false if the backingMap is not empty
func TestOrdered_IsEmpty(t *testing.T) {
	// Test 1: empty backingMap
	t.Run("empty backingMap", func(t *testing.T) {
		set := Ordered[int]()
		if !set.IsEmpty() {
			t.Errorf("Expected backingMap to be empty, but it is not")
		}
	})

	// Test 2: non-empty backingMap
	t.Run("non-empty backingMap", func(t *testing.T) {
		set := Ordered[int](1, 2, 3, 5, 6, 7)
		if set.IsEmpty() {
			t.Errorf("Expected backingMap to not be empty, but it is")
		}
	})

	// Test 3: after each element is removed
	t.Run("after each element removed", func(t *testing.T) {
		set := Ordered[int]()
		iterables.OfSlice(1, 2, 3, 5, 6, 7).
			Each(func(v int) {
				set.Add(v)
				if set.IsEmpty() {
					t.Errorf("Expected backingMap to not be empty, but it is")
				}
				set.Remove(v)
				if !set.IsEmpty() {
					t.Errorf("Expected backingMap to be empty, but it is not")
				}
			})
	})
}

// TestOrdered_Len tests the len method of an ordered backingMap
//
// after calling 'Len()' the backingMap should:
//   - return the correct number of elements
//   - contain the underlying map compare its length
func TestOrdered_Len(t *testing.T) {
	// Test 1: empty backingMap
	t.Run("empty backingMap", func(t *testing.T) {
		set := Ordered[int]()
		cast, ok := set.(*ordered[int, int])
		if !ok {
			t.Errorf("Expected backingMap to be of type ordered, but it is not")
		}
		if set.Len() != 0 {
			t.Errorf("Expected backingMap length 0, but got %d", set.Len())
		}
		if len(cast.table) != set.Len() {
			t.Errorf("Expected slice length 0, but got %d", len(cast.table))
		}
	})

	// Test 2: non-empty backingMap
	t.Run("non-empty backingMap", func(t *testing.T) {
		set := Ordered[int](1, 2, 3, 5, 6, 7)
		cast, ok := set.(*ordered[int, int])
		if !ok {
			t.Errorf("Expected backingMap to be of type ordered, but it is not")
		}
		if set.Len() != 6 {
			t.Errorf("Expected backingMap length 6, but got %d", set.Len())
		}
		if len(cast.table) != set.Len() {
			t.Errorf("Expected slice length 6, but got %d", len(cast.table))
		}
	})

	// Test 3: after each element is added
	t.Run("after each element added", func(t *testing.T) {
		set := Ordered[int]()
		cast, ok := set.(*ordered[int, int])
		if !ok {
			t.Errorf("Expected backingMap to be of type ordered, but it is not")
		}
		iterables.OfSlice(1, 2, 3, 5, 6, 7).
			EachN(func(i, v int) {
				set.Add(v)
				if set.Len() != i+1 {
					t.Errorf("Expected backingMap length %d, but got %d", v, set.Len())
				}
				if len(cast.table) != set.Len() {
					t.Errorf("Expected slice length %d, but got %d", v, len(cast.table))
				}
			})
	})

	// Test 4: after each element is removed
	t.Run("after each element removed", func(t *testing.T) {
		set := Ordered[int]()
		cast, ok := set.(*ordered[int, int])
		if !ok {
			t.Errorf("Expected backingMap to be of type ordered, but it is not")
		}
		iterables.OfSlice(1, 2, 3, 5, 6, 7).
			EachN(func(i, v int) {
				set.Add(v)
			}).
			EachN(func(i, v int) {
				set.Remove(v)
				if set.Len() != 6-i-1 {
					t.Errorf("Expected backingMap length %d, but got %d", 6-i-1, set.Len())
				}
				if len(cast.table) != set.Len() {
					t.Errorf("Expected slice length %d, but got %d", 6-i-1, len(cast.table))
				}
			})
	})
}

// TestOrdered_Remove tests the remove method of an ordered backingMap
//
// after calling 'Remove()' the backingMap should:
//   - remove the element from the backingMap slice
//   - remove the element from the idx map
//   - return the correct number of elements
func TestOrdered_Remove(t *testing.T) {
	// Test 1: empty backingMap
	t.Run("empty backingMap", func(t *testing.T) {
		set := Ordered[int]()
		cast, ok := set.(*ordered[int, int])
		if !ok {
			t.Errorf("Expected backingMap to be of type ordered, but it is not")
		}
		if set.Len() != 0 {
			t.Errorf("Expected backingMap length 0, but got %d", set.Len())
		}
		if len(cast.table) != set.Len() {
			t.Errorf("Expected slice length 0, but got %d", len(cast.table))
		}
		set.Remove(1)
		if set.Len() != 0 {
			t.Errorf("Expected backingMap length 0, but got %d", set.Len())
		}
		if len(cast.table) != set.Len() {
			t.Errorf("Expected slice length 0, but got %d", len(cast.table))
		}
	})

	// Test 2: non-empty backingMap
	t.Run("non-empty backingMap", func(t *testing.T) {
		set := Ordered[int](1, 2, 3, 5, 6, 7)
		cast, ok := set.(*ordered[int, int])
		if !ok {
			t.Errorf("Expected backingMap to be of type ordered, but it is not")
		}
		if set.Len() != 6 {
			t.Errorf("Expected backingMap length 6, but got %d", set.Len())
		}
		if len(cast.table) != set.Len() {
			t.Errorf("Expected slice length 6, but got %d", len(cast.table))
		}
		for i, v := range []int{1, 2, 3, 5, 6, 7} {
			set.Remove(v)
			if set.Len() != 6-i-1 {
				t.Errorf("Expected backingMap length %d, but got %d", 6-1, set.Len())
			}
			if len(cast.table) != set.Len() {
				t.Errorf("Expected slice length %d, but got %d", 6-1, len(cast.table))
			}
		}
	})
}

// TestOrdered_Values tests the backingMap method of an ordered backingMap
//
// after calling 'Values()' the backingMap should:
//   - return the correct number of elements
//   - return the correct elements
func TestOrdered_Values(t *testing.T) {
	// Test 1: Empty
	t.Run("empty backingMap", func(t *testing.T) {
		set := Ordered[int]()
		if set.Len() != 0 {
			t.Errorf("Expected backingMap length 0, but got %d", set.Len())
		}
		if len(set.Elements()) != 0 {
			t.Errorf("Expected backingMap length 0, but got %d", len(set.Elements()))
		}
	})

	// Test 2: Non-empty
	t.Run("non-empty backingMap", func(t *testing.T) {
		values := map[int]bool{1: true, 2: true, 3: true, 5: true, 6: true, 7: true}
		set := Ordered[int](1, 2, 3, 5, 6, 7)
		if set.Len() != 6 {
			t.Errorf("Expected backingMap length 6, but got %d", set.Len())
		}
		if len(set.Elements()) != 6 {
			t.Errorf("Expected backingMap length 6, but got %d", len(set.Elements()))
		}
		for _, v := range set.Elements() {
			if !values[v] {
				t.Errorf("Expected backingMap to contain %d, but it does not", v)
			}
		}
		for v := range values {
			if !arrays.Contains(set.Elements(), v) {
				t.Errorf("Expected backingMap to contain %d, but it does not", v)
			}
		}
	})
}

// TestOrdered_Creation tests the creation of an ordered backingMap.
//
// after creating the backingMap it must assert:
//   - It is not nil
//   - Its size is the same as the number of inserted elements
func TestConcOrdered_Creation(t *testing.T) {
	// Test 1: empty backingMap
	t.Run("create empty backingMap", func(t *testing.T) {
		set := ConcOrdered[int]()
		cast, ok := set.(*concOrdered[int, int])
		if !ok {
			t.Errorf("Expected backingMap to be of type ConcOrdered, but it is not")
		}
		if len(cast.table) != 0 {
			t.Errorf("Expected map length 0, but got %d", len(cast.table))
		}
		if cast.tree.Size() != 0 {
			t.Errorf("Expected tree size 0, but got %d", cast.tree.Size())
		}
		if set == nil {
			t.Errorf("Expected backingMap to be not nil, but it is")
		}
		if set.Len() != 0 {
			t.Errorf("Expected backingMap length 0, but got %d", set.Len())
		}
	})

	// Test 2: backingMap with elements
	t.Run("create backingMap with elements", func(t *testing.T) {
		set := ConcOrdered[int](1, 2, 3, 5, 6, 7)
		cast, ok := set.(*concOrdered[int, int])
		if !ok {
			t.Errorf("Expected backingMap to be of type ConcOrdered, but it is not")
		}
		if len(cast.table) != 6 {
			t.Errorf("Expected map length 6, but got %d", len(cast.table))
		}
		if cast.tree.Size() != 6 {
			t.Errorf("Expected tree size 6, but got %d", cast.tree.Size())
		}
		if set == nil {
			t.Errorf("Expected backingMap to be not nil, but it is")
		}
		if set.Len() != 6 {
			t.Errorf("Expected backingMap length 6, but got %d", set.Len())
		}

		iterables.OfSlice(1, 2, 3, 5, 6, 7).
			EachN(func(i int, v int) {
				if set.IndexOf(v) != i {
					t.Errorf("Expected backingMap to contain '%d' at index '%d', but it "+
						"doesn't", v, i)
				}
			})
	})

	// Test 3: backingMap with unordered elements
	t.Run("create backingMap with unordered elements", func(t *testing.T) {
		set := ConcOrdered[int](7, 6, 5, 3, 2, 1)
		if set.Len() != 6 {
			t.Errorf("Expected backingMap length 6, but got %d", set.Len())
		}
		iterables.OfSlice(1, 2, 3, 5, 6, 7).
			EachN(func(i int, v int) {
				if set.IndexOf(v) != i {
					t.Errorf("Expected backingMap to contain '%d' at index '%d', but it "+
						"doesn't", v, i)
				}
			})
	})

	// Test 4: backingMap with duplicated unordered elements
	t.Run("create backingMap with duplicated unordered elements", func(t *testing.T) {
		set := ConcOrdered[int](7, 7, 6, 6, 5, 5, 3, 3, 2, 2, 1, 1)
		if set.Len() != 6 {
			t.Errorf("Expected backingMap length 6, but got %d", set.Len())
		}
		iterables.OfSlice(1, 2, 3, 5, 6, 7).
			EachN(func(i int, v int) {
				if set.IndexOf(v) != i {
					t.Errorf("Expected backingMap to contain '%d' at index '%d', but it "+
						"doesn't", v, i)
				}
			})
	})
}

// TestConcOrdered_Add tests the add method of an ordered backingMap.
//
// after adding an element to the backingMap it must assert:
//   - The element is in the backingMap map
//   - The backingMap size is the same as the number of inserted elements
//   - The order of elements is the same as the natural comparable order
func TestConcOrdered_Add(t *testing.T) {
	// Test 1: start empty, add ordered
	t.Run("start empty, add ordered", func(t *testing.T) {
		set := ConcOrdered[int]()
		if set.Len() != 0 {
			t.Errorf("Expected backingMap length 0, but got %d", set.Len())
		}
		iterables.OfSlice(1, 2, 3, 5, 6, 7).
			EachN(func(i int, v int) {
				set.Add(v)
				if set.Len() != i+1 {
					t.Errorf("Expected backingMap length %d, but got %d", i+1, set.Len())
				}
			}).
			EachN(func(i int, v int) {
				if set.IndexOf(v) != i {
					t.Errorf("Expected backingMap to contain '%d' at index '%d', but it "+
						"doesn't", v, i)
				}
			})
	})

	// Test 2: start empty, add unordered
	t.Run("start empty, add unordered", func(t *testing.T) {
		set := ConcOrdered[int]()
		if set.Len() != 0 {
			t.Errorf("Expected backingMap length 0, but got %d", set.Len())
		}
		it := iterables.OfSlice(7, 6, 5, 3, 2, 1)
		it.EachN(func(i int, v int) {
			set.Add(v)
			if set.Len() != i+1 {
				t.Errorf("Expected backingMap length %d, but got %d", i+1, set.Len())
			}
		})
		it.Sort(cmp.Less[int]).
			EachN(func(i int, v int) {
				if set.IndexOf(v) != i {
					t.Errorf("Expected backingMap to contain '%d' at index '%d', but it "+
						"doesn't", v, i)
				}
			})
	})

	// Test 3: start empty, add unordered with duplicates
	t.Run("start empty, add unordered with duplicates", func(t *testing.T) {
		set := ConcOrdered[int]()
		if set.Len() != 0 {
			t.Errorf("Expected backingMap length 0, but got %d", set.Len())
		}
		it := iterables.OfSlice(7, 7, 6, 6, 5, 5, 3, 3, 2, 2, 1, 1)
		it.EachN(func(i int, v int) {
			set.Add(v)
			if set.Len() != (i)/2+1 {
				t.Errorf("Expected backingMap length '%d' at index '%d', but got %d", (i)/2+1, i, set.Len())
			}
		})
		it.ToSet(keyers.FingerPrint[int]()).
			Sort(cmp.Less[int]).
			EachN(func(i int, v int) {
				if idx := set.IndexOf(v); idx != i {
					t.Errorf("Expected backingMap to contain '%d' at index '%d', but it "+
						"returned %d", v, i, idx)
				}
			})
	})
}

// TestConcOrdered_Clear tests the clear method of an ordered backingMap
//
// after calling 'Clear()' the backingMap should:
//   - have no elements
//   - return true when calling 'IsEmpty()'
func TestConcOrdered_Clear(t *testing.T) {
	// Test 1: clear empty backingMap
	t.Run("clear empty backingMap", func(t *testing.T) {
		set := ConcOrdered[int]()
		cast, ok := set.(*concOrdered[int, int])
		if !ok {
			t.Errorf("Expected backingMap to be of type ConcOrdered, but it is not")
		}
		if set.Len() != 0 {
			t.Errorf("Expected backingMap length 0, but got %d", set.Len())
		}
		if len(cast.table) != 0 {
			t.Errorf("Expected slice length 0, but got %d", len(cast.table))
		}
		if !set.IsEmpty() {
			t.Errorf("Expected backingMap to be empty, but it is not")
		}
		set.Clear()
		if set.Len() != 0 {
			t.Errorf("Expected backingMap length 0, but got %d", set.Len())
		}
		if !set.IsEmpty() {
			t.Errorf("Expected backingMap to be empty, but it is not")
		}
	})

	// Test 2: clear after add backingMap
	t.Run("clear after add backingMap", func(t *testing.T) {
		set := ConcOrdered[int]()
		cast, ok := set.(*concOrdered[int, int])
		if !ok {
			t.Errorf("Expected backingMap to be of type ConcOrdered, but it is not")
		}
		iterables.OfSlice(1, 2, 3, 5, 6, 7).
			EachN(func(i int, v int) {
				set.Add(v)
				if set.Len() != i+1 {
					t.Errorf("Expected backingMap length %d, but got %d", i+1, set.Len())
				}
			})
		if set.Len() != 6 {
			t.Errorf("Expected backingMap length 6, but got %d", set.Len())
		}
		if len(cast.table) != 6 {
			t.Errorf("Expected slice length 6, but got %d", len(cast.table))
		}
		if set.IsEmpty() {
			t.Errorf("Expected backingMap to not be empty, but it is")
		}
		if cast.tree.Size() != 6 {
			t.Errorf("Expected tree size 6, but got %d", cast.tree.Size())
		}
		set.Clear()
		if set.Len() != 0 {
			t.Errorf("Expected backingMap length 0, but got %d", set.Len())
		}
		if !set.IsEmpty() {
			t.Errorf("Expected backingMap to be empty, but it is not")
		}
		if len(cast.table) != 0 {
			t.Errorf("Expected slice length 0, but got %d", len(cast.table))
		}
		if cast.tree.Size() != 0 {
			t.Errorf("Expected tree size 0, but got %d", cast.tree.Size())
		}
	})

	// Test 3: clear after each add
	t.Run("clear after each add", func(t *testing.T) {
		set := ConcOrdered[int]()
		cast, ok := set.(*concOrdered[int, int])
		if !ok {
			t.Errorf("Expected backingMap to be of type ConcOrdered, but it is not")
		}
		iterables.OfSlice(1, 2, 3, 5, 6, 7).
			EachN(func(i int, v int) {
				set.Add(v)
				if set.IsEmpty() {
					t.Errorf("Expected backingMap to not be empty, but it is")
				}
				if set.Len() != 1 {
					t.Errorf("Expected backingMap length 1, but got %d", set.Len())
				}
				if len(cast.table) != 1 {
					t.Errorf("Expected slice length 1, but got %d", len(cast.table))
				}
				if cast.tree.Size() != 1 {
					t.Errorf("Expected tree size 1, but got %d", cast.tree.Size())
				}
				set.Clear()
				if set.Len() != 0 {
					t.Errorf("Expected backingMap length 0, but got %d", set.Len())
				}
				if !set.IsEmpty() {
					t.Errorf("Expected backingMap to be empty, but it is")
				}
				if len(cast.table) != 0 {
					t.Errorf("Expected slice length 0, but got %d", len(cast.table))
				}
				if cast.tree.Size() != 0 {
					t.Errorf("Expected tree size 0, but got %d", cast.tree.Size())
				}
			})
	})

	// Test 4: clear not empty backingMap
	t.Run("clear not empty backingMap", func(t *testing.T) {
		set := ConcOrdered[int](1, 2, 3, 5, 6, 7)
		cast, ok := set.(*concOrdered[int, int])
		if !ok {
			t.Errorf("Expected backingMap to be of type ConcOrdered, but it is not")
		}
		if set.Len() != 6 {
			t.Errorf("Expected backingMap length 6, but got %d", set.Len())
		}
		if set.IsEmpty() {
			t.Errorf("Expected backingMap to not be empty, but it is")
		}
		if len(cast.table) != 6 {
			t.Errorf("Expected slice length 6, but got %d", len(cast.table))
		}
		if cast.tree.Size() != 6 {
			t.Errorf("Expected tree size 6, but got %d", cast.tree.Size())
		}
		set.Clear()
		if set.Len() != 0 {
			t.Errorf("Expected backingMap length 0, but got %d", set.Len())
		}
		if !set.IsEmpty() {
			t.Errorf("Expected backingMap to be empty, but it is")
		}
		if len(cast.table) != 0 {
			t.Errorf("Expected slice length 0, but got %d", len(cast.table))
		}
		if cast.tree.Size() != 0 {
			t.Errorf("Expected tree size 0, but got %d", cast.tree.Size())
		}
	})
}

// TestConcOrdered_Delete tests the delete method of an ordered backingMap
//
// after calling 'Delete()' the backingMap should:
//   - have one less element
//   - return false when calling 'Has()'
//   - not find the value on the underlying map
func TestConcOrdered_Delete(t *testing.T) {
	// Test 1: delete from empty backingMap
	t.Run("deleting from empty backingMap", func(t *testing.T) {
		set := ConcOrdered[int]()
		cast, ok := set.(*concOrdered[int, int])
		if !ok {
			t.Errorf("Expected backingMap to be of type ConcOrdered, but it is not")
		}
		if set.Len() != 0 {
			t.Errorf("Expected backingMap length '0', but got '%d'", set.Len())
		}
		if !set.IsEmpty() {
			t.Errorf("Expected backingMap to be empty, but it is not")
		}
		if len(cast.table) != 0 {
			t.Errorf("Expected slice length '0', but got '%d'", len(cast.table))
		}
		if cast.tree.Size() != 0 {
			t.Errorf("Expected tree size '0', but got '%d'", cast.tree.Size())
		}
		set.Delete(1)
		if set.Len() != 0 {
			t.Errorf("Expected backingMap length '0', but got '%d'", set.Len())
		}
		if !set.IsEmpty() {
			t.Errorf("Expected backingMap to be empty, but it is not")
		}
		if len(cast.table) != set.Len() {
			t.Errorf("Expected slice length '0', but got '%d'", len(cast.table))
		}
		if cast.tree.Size() != set.Len() {
			t.Errorf("Expected tree size '0', but got '%d'", cast.tree.Size())
		}
	})

	// Test 2: delete from non-empty backingMap
	t.Run("deleting from non-empty backingMap", func(t *testing.T) {
		set := ConcOrdered[int](1, 2, 3, 5, 6, 7)
		cast, ok := set.(*concOrdered[int, int])
		if !ok {
			t.Errorf("Expected backingMap to be of type 'ConcOrdered', but it is not")
		}
		if !ok {
			t.Errorf("Expected backingMap to be of type '%T', but got '%T'", cast, set)
		}
		if set.Len() != 6 {
			t.Errorf("Expected backingMap length '6', but got '%d'", set.Len())
		}
		if set.IsEmpty() {
			t.Errorf("Expected backingMap to not be empty, but it is")
		}
		if len(cast.table) != 6 {
			t.Errorf("Expected slice length '6', but got '%d'", len(cast.table))
		}
		if cast.tree.Size() != 6 {
			t.Errorf("Expected tree size '6', but got '%d'", cast.tree.Size())
		}

		it := iterables.OfSlice(1, 2, 3, 5, 6, 7)
		it.EachN(func(i int, v int) {
			if !set.Delete(0) {
				t.Errorf("Loop[%d]: expected to delete '%d', but it didn't", i, v)
			}
			it.RemoveAt(0)
			if set.Len() != it.Len() {
				t.Errorf("Loop[%d]: Expected backingMap length '%d', but got '%d'", i, it.Len(), set.Len())
			}
			if len(cast.table) != set.Len() {
				t.Errorf("Loop[%d]: Expected slice length '%d', but got '%d'", i, it.Len(), len(cast.table))
			}
			if cast.tree.Size() != set.Len() {
				t.Errorf("Loop[%d]: Expected tree size '%d', but got '%d'", i, it.Len(), cast.tree.Size())
			}
		})
	})

}

// TestConcOrdered_Equals tests the compare method of an ordered backingMap
//
// after calling 'Equals()' the backingMap should:
//   - return true when the sets have the same address
//   - return true when the sets contain the same backingMap (order-agnostic)
//   - return false when the sets do not contain the same backingMap
func TestConcOrdered_Equals(t *testing.T) {
	// Test 1: compare address
	t.Run("compare address", func(t *testing.T) {
		set1 := ConcOrdered[int](1, 2, 3, 5, 6, 7)
		set2 := set1
		if !set1.Equals(set2) {
			t.Errorf("Expected sets to be compare, but they are not")
		}
	})

	// Test 2: compare backingMap
	t.Run("compare backingMap", func(t *testing.T) {
		set1 := ConcOrdered[int](1, 2, 3, 5, 6, 7)
		set2 := ConcOrdered[int](1, 2, 3, 5, 6, 7)
		if !set1.Equals(set2) {
			t.Errorf("Expected sets to be compare, but they are not")
		}
	})

	// Test 3: not compare backingMap
	t.Run("not compare backingMap", func(t *testing.T) {
		set1 := ConcOrdered[int](1, 2, 3, 5, 6, 7)
		set2 := ConcOrdered[int](1, 2, 3, 5, 6, 8)
		if set1.Equals(set2) {
			t.Errorf("Expected sets to be not compare, but they are")
		}
	})

	// Test 4: compare backingMap, different order
	t.Run("compare backingMap, different order", func(t *testing.T) {
		set1 := ConcOrdered[int](1, 2, 3, 5, 6, 7)
		set2 := ConcOrdered[int](3, 5, 6, 7, 1, 2)
		if !set1.Equals(set2) {
			t.Errorf("Expected sets to be compare, but they are not")
		}
	})
}

// TestConcOrdered_Get tests the get method of an ordered backingMap
//
// after calling 'Get()' the backingMap should:
//   - return true when the backingMap contains the value
//   - return the value when the backingMap contains the value
//   - assert the value exists on the map backingMap
//   - return false when the backingMap does not contain the value
func TestConcOrdered_Get(t *testing.T) {
	// Test 1: empty backingMap
	t.Run("empty backingMap", func(t *testing.T) {
		set := ConcOrdered[int]()
		cast, ok := set.(*concOrdered[int, int])
		if !ok {
			t.Errorf("Expected backingMap to be of type ConcOrdered, but it is not")
		}
		_, ok = set.Get(1)
		if ok {
			t.Errorf("Expected backingMap to not contain '1', but it does")
		}
		if _, ok = cast.table[1]; ok {
			t.Errorf("Expected map backingMap to not contain '1', but it does")
		}
	})

	// Test 2: non-empty backingMap
	t.Run("non-empty backingMap", func(t *testing.T) {
		set := ConcOrdered[int](1, 2, 3, 5, 6, 7)
		cast, ok := set.(*concOrdered[int, int])
		if !ok {
			t.Errorf("Expected backingMap to be of type ConcOrdered, but it is not")
		}
		for i := 1; i < 6; i++ {
			v, ok := set.Get(i)
			if !ok {
				t.Errorf("Expected backingMap to contain '%d', but it doesn't", i)
			}
			if _, ok = cast.table[v]; !ok {
				t.Errorf("Expected map backingMap to contain '%d', but it doesn't", i)
			}
			if i >= 3 {
				if v != i+2 {
					t.Errorf("Expected backingMap to contain '%d' at index '%d', but it "+
						"returned '%d'", i+2, i, v)
				}
			} else if v != i+1 {
				t.Errorf("Expected backingMap to contain '%d' at index '%d', but it "+
					"returned '%d'", i+2, i, v)
			}
		}
	})
}

// TestConcOrdered_Has tests the has method of an ordered backingMap
//
// after calling 'Has()' the backingMap should:
//   - return true when the backingMap contains the value
//   - return false when the backingMap does not contain the value
//   - assert the same results from the map and the tree
func TestConcOrdered_Has(t *testing.T) {
	// Test 1: empty backingMap
	t.Run("empty backingMap", func(t *testing.T) {
		set := ConcOrdered[int]()
		if set.Has(1) {
			t.Errorf("Expected backingMap to not contain '1', but it does")
		}
	})

	// Test 2: non-empty backingMap
	t.Run("non-empty backingMap", func(t *testing.T) {
		set := ConcOrdered[int](1, 2, 3, 5, 6, 7)
		cast, ok := set.(*concOrdered[int, int])
		if !ok {
			t.Errorf("Expected backingMap to be of type ConcOrdered, but it is not")
		}
		for i := 1; i <= 7; i++ {
			if i != 4 {
				if !set.Has(i) {
					t.Errorf("Expected backingMap to contain '%d', but it doesn't", i)
				}
				if _, ok = cast.table[i]; !ok {
					t.Errorf("Expected map backingMap to contain '%d', but it doesn't", i)
				}
				if !cast.tree.Contains(i) {
					t.Errorf("Expected tree to contain '%d', but it doesn't", i)
				}
			} else {
				if set.Has(i) {
					t.Errorf("Expected backingMap to not contain '%d', but it does", i)
				}
				if _, ok = cast.table[i]; ok {
					t.Errorf("Expected map backingMap to not contain '%d', but it does", i)
				}
				if cast.tree.Contains(i) {
					t.Errorf("Expected tree to not contain '%d', but it does", i)
				}
			}
		}
	})
}

// TestConcOrdered_IsEmpty tests the isEmpty method of an ordered backingMap
//
// after calling 'IsEmpty()' the backingMap should:
//   - return true if the backingMap is empty
//   - return false if the backingMap is not empty
func TestConcOrdered_IsEmpty(t *testing.T) {
	// Test 1: empty backingMap
	t.Run("empty backingMap", func(t *testing.T) {
		set := ConcOrdered[int]()
		if !set.IsEmpty() {
			t.Errorf("Expected backingMap to be empty, but it is not")
		}
	})

	// Test 2: non-empty backingMap
	t.Run("non-empty backingMap", func(t *testing.T) {
		set := ConcOrdered[int](1, 2, 3, 5, 6, 7)
		cast, ok := set.(*concOrdered[int, int])
		if !ok {
			t.Errorf("Expected backingMap to be of type ConcOrdered, but it is not")
		}
		if set.IsEmpty() {
			t.Errorf("Expected backingMap to not be empty, but it is")
		}
		if cast.tree.Size() == 0 {
			t.Errorf("Expected tree size '0', but got '%d'", cast.tree.Size())
		}
		if len(cast.table) == 0 {
			t.Errorf("Expected map length '0', but got '%d'", len(cast.table))
		}
	})

	// Test 3: after each element is removed
	t.Run("after each element removed", func(t *testing.T) {
		set := ConcOrdered[int]()
		cast, ok := set.(*concOrdered[int, int])
		if !ok {
			t.Errorf("Expected backingMap to be of type ConcOrdered, but it is not")
		}
		iterables.OfSlice(1, 2, 3, 5, 6, 7).
			Each(func(v int) {
				set.Add(v)
				if set.IsEmpty() {
					t.Errorf("Expected backingMap to not be empty, but it is")
				}
				if cast.tree.Size() != 1 {
					t.Errorf("Expected tree size '1', but got '%d'", cast.tree.Size())
				}
				if cast.tree.IsEmpty() {
					t.Errorf("Expected tree to not be empty, but it is")
				}
				if len(cast.table) != 1 {
					t.Errorf("Expected map length '1', but got '%d'", len(cast.table))
				}
				set.Remove(v)
				if !set.IsEmpty() {
					t.Errorf("Expected backingMap to be empty, but it is not")
				}
				if cast.tree.Size() != 0 {
					t.Errorf("Expected tree size '0', but got '%d'", cast.tree.Size())
				}
				if !cast.tree.IsEmpty() {
					t.Errorf("Expected tree to be empty, but it is not")
				}
				if len(cast.table) != 0 {
					t.Errorf("Expected map length '0', but got '%d'", len(cast.table))
				}
			})
	})
}

// TestConcOrdered_Len tests the len method of an ordered backingMap
//
// after calling 'Len()' the backingMap should:
//   - return the correct number of elements
//   - contain the underlying map compare its length
func TestConcOrdered_Len(t *testing.T) {
	// Test 1: empty backingMap
	t.Run("empty backingMap", func(t *testing.T) {
		set := ConcOrdered[int]()
		cast, ok := set.(*concOrdered[int, int])
		if !ok {
			t.Errorf("Expected backingMap to be of type ConcOrdered, but it is not")
		}
		if set.Len() != 0 {
			t.Errorf("Expected backingMap length 0, but got %d", set.Len())
		}
		if len(cast.table) != 0 {
			t.Errorf("Expected slice length 0, but got %d", len(cast.table))
		}
		if cast.tree.Size() != 0 {
			t.Errorf("Expected tree size '0', but got '%d'", cast.tree.Size())
		}
	})

	// Test 2: non-empty backingMap
	t.Run("non-empty backingMap", func(t *testing.T) {
		set := ConcOrdered[int](1, 2, 3, 5, 6, 7)
		cast, ok := set.(*concOrdered[int, int])
		if !ok {
			t.Errorf("Expected backingMap to be of type ConcOrdered, but it is not")
		}
		if set.Len() != 6 {
			t.Errorf("Expected backingMap length 6, but got %d", set.Len())
		}
		if len(cast.table) != 6 {
			t.Errorf("Expected slice length 6, but got %d", len(cast.table))
		}
		if cast.tree.Size() != 6 {
			t.Errorf("Expected tree size '6', but got '%d'", cast.tree.Size())
		}
	})

	// Test 3: after each element is added
	t.Run("after each element added", func(t *testing.T) {
		set := ConcOrdered[int]()
		cast, ok := set.(*concOrdered[int, int])
		if !ok {
			t.Errorf("Expected backingMap to be of type ConcOrdered, but it is not")
		}
		iterables.OfSlice(1, 2, 3, 5, 6, 7).
			EachN(func(i, v int) {
				set.Add(v)
				if set.Len() != i+1 {
					t.Errorf("Expected backingMap length '%d', but got '%d'", v, set.Len())
				}
				if len(cast.table) != set.Len() {
					t.Errorf("Expected slice length '%d', but got '%d'", v, len(cast.table))
				}
				if cast.tree.Size() != set.Len() {
					t.Errorf("Expected tree size '%d', but got '%d'", v, cast.tree.Size())
				}
			})
	})

	// Test 4: after each element is removed
	t.Run("after each element removed", func(t *testing.T) {
		set := ConcOrdered[int]()
		cast, ok := set.(*concOrdered[int, int])
		if !ok {
			t.Errorf("Expected backingMap to be of type ConcOrdered, but it is not")
		}
		iterables.OfSlice(1, 2, 3, 5, 6, 7).
			EachN(func(i, v int) {
				set.Add(v)
			}).
			EachN(func(i, v int) {
				set.Remove(v)
				if set.Len() != 6-i-1 {
					t.Errorf("Expected backingMap length '%d', but got '%d'", 6-i-1, set.Len())
				}
				if len(cast.table) != set.Len() {
					t.Errorf("Expected slice length '%d', but got '%d'", 6-i-1, len(cast.table))
				}
				if cast.tree.Size() != set.Len() {
					t.Errorf("Expected tree size '%d', but got '%d'", 6-i-1, cast.tree.Size())
				}
			})
	})
}

// TestConcOrdered_Remove tests the remove method of an ordered backingMap
//
// after calling 'Remove()' the backingMap should:
//   - remove the element from the tree
//   - remove the element from the idx map
//   - return the correct number of elements
//   - the element cannot be found on the map and the tree
func TestConcOrdered_Remove(t *testing.T) {
	// Test 1: empty backingMap
	t.Run("empty backingMap", func(t *testing.T) {
		set := ConcOrdered[int]()
		cast, ok := set.(*concOrdered[int, int])
		if !ok {
			t.Errorf("Expected backingMap to be of type ConcOrdered, but it is not")
		}
		if set.Len() != 0 {
			t.Errorf("Expected backingMap length '0', but got '%d'", set.Len())
		}
		if len(cast.table) != 0 {
			t.Errorf("Expected slice length '0', but got '%d'", len(cast.table))
		}
		if cast.tree.Size() != 0 {
			t.Errorf("Expected tree size '0', but got '%d'", cast.tree.Size())
		}
		set.Remove(1)
		if set.Len() != 0 {
			t.Errorf("Expected backingMap length '0', but got '%d'", set.Len())
		}
		if len(cast.table) != 0 {
			t.Errorf("Expected slice length '0', but got '%d'", len(cast.table))
		}
		if cast.tree.Size() != 0 {
			t.Errorf("Expected tree size '0', but got '%d'", cast.tree.Size())
		}
	})

	// Test 2: non-empty backingMap
	t.Run("non-empty backingMap", func(t *testing.T) {
		set := ConcOrdered[int](1, 2, 3, 5, 6, 7)
		cast, ok := set.(*concOrdered[int, int])
		if !ok {
			t.Errorf("Expected backingMap to be of type ConcOrdered, but it is not")
		}
		if set.Len() != 6 {
			t.Errorf("Expected backingMap length '6', but got '%d'", set.Len())
		}
		if len(cast.table) != set.Len() {
			t.Errorf("Expected slice length '6', but got '%d'", len(cast.table))
		}
		if cast.tree.Size() != set.Len() {
			t.Errorf("Expected tree size '6', but got '%d'", cast.tree.Size())
		}
		for i, v := range []int{1, 2, 3, 5, 6, 7} {
			if !set.Has(v) {
				t.Errorf("Expected backingMap to contain '%d', but it doesn't", v)
			}
			if !cast.tree.Contains(v) {
				t.Errorf("Expected tree to contain '%d', but it doesn't", v)
			}
			if _, ok = cast.table[v]; !ok {
				t.Errorf("Expected map to contain '%d', but it doesn't", v)
			}
			set.Remove(v)
			expSize := 6 - i - 1
			if set.Len() != expSize {
				t.Errorf("Expected backingMap length '%d', but got '%d'", expSize, set.Len())
			}
			if len(cast.table) != expSize {
				t.Errorf("Expected slice length %d, but got %d", 6-1, len(cast.table))
			}
			if cast.tree.Size() != expSize {
				t.Errorf("Expected tree size '%d', but got '%d'", expSize, cast.tree.Size())
			}
			if set.Has(v) {
				t.Errorf("Expected backingMap to not contain '%d', but it does", v)
			}
			if cast.tree.Contains(v) {
				t.Errorf("Expected tree to not contain '%d', but it does", v)
			}
			if _, ok = cast.table[v]; ok {
				t.Errorf("Expected map to not contain '%d', but it does", v)
			}
		}
	})
}

// TestConcOrdered_Values tests the backingMap method of an ordered backingMap
//
// after calling 'Values()' the backingMap should:
//   - return the correct number of elements
//   - return the correct elements
func TestConcOrdered_Values(t *testing.T) {
	// Test 1: Empty
	t.Run("empty backingMap", func(t *testing.T) {
		set := ConcOrdered[int]()
		if set.Len() != 0 {
			t.Errorf("Expected backingMap length 0, but got %d", set.Len())
		}
		if len(set.Elements()) != 0 {
			t.Errorf("Expected backingMap length 0, but got %d", len(set.Elements()))
		}
	})

	// Test 2: Non-empty
	t.Run("non-empty backingMap", func(t *testing.T) {
		values := map[int]bool{1: true, 2: true, 3: true, 5: true, 6: true, 7: true}
		set := ConcOrdered[int](1, 2, 3, 5, 6, 7)
		if set.Len() != 6 {
			t.Errorf("Expected backingMap length '6', but got '%d'", set.Len())
		}
		if len(set.Elements()) != 6 {
			t.Errorf("Expected backingMap length '6', but got '%d'", len(set.Elements()))
		}
		for _, v := range set.Elements() {
			if !values[v] {
				t.Errorf("Expected backingMap to contain '%d', but it does not", v)
			}
		}
		for v := range values {
			if !arrays.Contains(set.Elements(), v) {
				t.Errorf("Expected backingMap to contain '%d'', but it does not", v)
			}
		}
	})
}
