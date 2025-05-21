// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package trees

import (
	"github.com/andrerrcosta2/gtools/core/data/comparators"
	"github.com/andrerrcosta2/gtools/core/data/str/iterables"
	"github.com/andrerrcosta2/gtools/core/data/str/nodes"
	"testing"
)

// TestRedBlack_Creation tests the creation of a RBT
//
// after being created the tree should:
//   - contain the same number of nodes as the number of elements passed
func TestRedBlack_Creation(t *testing.T) {
	cmp := comparators.Ordered[int]{}
	tree := RedBlack[int](cmp)
	if tree == nil {
		t.Errorf("tree is nil")
	}
	if tree.Size() != 0 {
		t.Errorf("tree size is not 0")
	}

	for i := 0; i < 10; i++ {
		tree.Insert(i)
		if tree.Size() != i+1 {
			t.Errorf("tree size is not %d", i+1)
		}
	}
}

// TestRedBlack_Ceiling tests the Ceiling method
//
// after it is called this test must assert:
//   - the value retrieved is the smallest value greater than
//     or equal to the given value
func TestRedBlack_Ceiling(t *testing.T) {
	cmp := comparators.Ordered[int]{}

	// Test 1: Only positive values
	t.Run("only positive values", func(t *testing.T) {
		tree := RedBlack[int](cmp)
		it := iterables.OfSlice(1, 2, 3, 5, 6, 7)
		it.EachN(func(i, v int) {
			tree.Insert(v)
			ceiling, ok := tree.Ceiling(v)
			if !ok {
				t.Errorf("tree should find ceiling of '%d'", v)
			}
			if ceiling != v {
				t.Errorf("ceiling of '%d' should be '%d' but got '%d'", v, v, ceiling)
			}
			ceiling, ok = tree.Ceiling(v + 1)
			if ok {
				t.Errorf("tree shouldn't find ceiling of '%d'", v+1)
			}
			if ceiling != 0 {
				t.Errorf("ceiling of 'value + 1' should be zero value but got '%d'", ceiling)
			}
			if i > 0 {
				if i == 3 {
					ceiling, ok = tree.Ceiling(v - 1)
					if !ok {
						t.Errorf("tree should find ceiling of '%d'", v-1)
					}
					if ceiling != v {
						t.Errorf("ceiling of '%d' should be '%d' but got '%d'", v-1, v, ceiling)
					}
				} else {
					ceiling, ok = tree.Ceiling(v - 1)
					if !ok {
						t.Errorf("tree should find ceiling of '%d'", v-1)
					}
					if ceiling != v-1 {
						t.Errorf("ceiling of '%d' should be '%d' but got '%d'", v-1, v, ceiling)
					}
				}
			} else {
				ceiling, ok = tree.Ceiling(v - 1)
				if !ok {
					t.Errorf("tree should find ceiling of '%d'", v-1)
				}
				if ceiling != it.At(i) {
					t.Errorf("ceiling of '%d' should be '%d' but got '%d'", v-1, it.At(i), ceiling)
				}
			}
		})
	})

	// Test 2: Only negative values
	t.Run("only positive values", func(t *testing.T) {
		tree := RedBlack[int](cmp)
		it := iterables.OfSlice(-7, -6, -5, -3, -2, -1)
		it.EachN(func(i, v int) {
			tree.Insert(v)
			ceiling, ok := tree.Ceiling(v)
			if !ok {
				t.Errorf("tree should find ceiling of '%d'", v)
			}
			if ceiling != v {
				t.Errorf("ceiling of '%d' should be '%d' but got '%d'", v, v, ceiling)
			}
			ceiling, ok = tree.Ceiling(v + 1)
			if ok {
				t.Errorf("tree shouldn't find ceiling of '%d'", v+1)
			}
			if ceiling != 0 {
				t.Errorf("ceiling of 'value + 1' should be zero value but got '%d'", ceiling)
			}
			if i > 0 {
				if i == 3 {
					ceiling, ok = tree.Ceiling(v - 1)
					if !ok {
						t.Errorf("tree should find ceiling of '%d'", v-1)
					}
					if ceiling != v {
						t.Errorf("ceiling of '%d' should be '%d' but got '%d'", v-1, v, ceiling)
					}
				} else {
					ceiling, ok = tree.Ceiling(v - 1)
					if !ok {
						t.Errorf("tree should find ceiling of '%d'", v-1)
					}
					if ceiling != v-1 {
						t.Errorf("ceiling of '%d' should be '%d' but got '%d'", v-1, v, ceiling)
					}
				}
			} else {
				ceiling, ok = tree.Ceiling(v - 1)
				if !ok {
					t.Errorf("tree should find ceiling of '%d'", v-1)
				}
				if ceiling != it.At(i) {
					t.Errorf("ceiling of '%d' should be '%d' but got '%d'", v-1, it.At(i), ceiling)
				}
			}
		})
	})

	// Test 1: Positive and negative values
	t.Run("positive and negative values", func(t *testing.T) {
		tree := RedBlack[int](cmp)
		it := iterables.OfSlice(-7, -6, -5, -3, -2, -1, 1, 2, 3, 5, 6, 7)
		it.EachN(func(i, v int) {
			tree.Insert(v)
			ceiling, ok := tree.Ceiling(v)
			if !ok {
				t.Errorf("tree should find ceiling of '%d'", v)
			}
			if ceiling != v {
				t.Errorf("ceiling of '%d' should be '%d' but got '%d'", v, v, ceiling)
			}
			ceiling, ok = tree.Ceiling(v + 1)
			if ok {
				t.Errorf("tree shouldn't find ceiling of '%d'", v+1)
			}
			if ceiling != 0 {
				t.Errorf("ceiling of 'value + 1' should be zero value but got '%d'", ceiling)
			}
			if i > 0 {
				if i == 3 || i == 6 || i == 9 {
					ceiling, ok = tree.Ceiling(v - 1)
					if !ok {
						t.Errorf("tree should find ceiling of '%d'", v-1)
					}
					if ceiling != v {
						t.Errorf("ceiling of '%d' should be '%d' but got '%d'", v-1, v, ceiling)
					}
				} else {
					ceiling, ok = tree.Ceiling(v - 1)
					if !ok {
						t.Errorf("tree should find ceiling of '%d'", v-1)
					}
					if ceiling != v-1 {
						t.Errorf("ceiling of '%d' should be '%d' but got '%d'", v-1, v-1, ceiling)
					}
				}
			} else {
				ceiling, ok = tree.Ceiling(v - 1)
				if !ok {
					t.Errorf("tree should find ceiling of '%d'", v-1)
				}
				if ceiling != it.At(i) {
					t.Errorf("ceiling of '%d' should be '%d' but got '%d'", v-1, it.At(i), ceiling)
				}
			}
		})
	})
}

// TestRedBlack_Clear tests the Clear method
//
// after it is called this test must assert:
//   - the size of the tree is 0
func TestRedBlack_Clear(t *testing.T) {
	cmp := comparators.Ordered[int]{}
	tree := RedBlack[int](cmp)
	iterables.OfSlice(1, 2, 3, 5, 6, 7).EachN(func(i, v int) {
		tree.Insert(v)
		if tree.Size() != 1 {
			t.Errorf("expected tree size of 1 but got %d", tree.Size())
		}
		tree.Clear()
		if tree.Size() != 0 {
			t.Errorf("expected tree size of 0 but got %d", tree.Size())
		}
	})
}

func TestRedBlack_ColorOf(t *testing.T) {
	cmp := comparators.Ordered[int]{}
	tree := RedBlack[int](cmp, 1, 2, 3, 5, 6, 7)
	// Expected tree
	//
	c, ok := tree.ColorOf(1)
	if !ok {
		t.Errorf("tree should contain 1")
	}
	if c != nodes.Black {
		t.Errorf("node '1' should be 'black' but got '%s'", c)
	}
	c, ok = tree.ColorOf(2)
	if !ok {
		t.Errorf("tree should contain 2")
	}
	if c != nodes.Black {
		t.Errorf("node '2' should be 'black' but got '%s'", c)
	}
	c, ok = tree.ColorOf(3)
	if !ok {
		t.Errorf("tree should contain 3")
	}
	if c != nodes.Black {
		t.Errorf("node '3' should be 'black' but got '%s'", c)
	}
	c, ok = tree.ColorOf(4)
	if ok {
		t.Errorf("tree should not contain 4")
	}
	c, ok = tree.ColorOf(5)
	if !ok {
		t.Errorf("tree should contain 5")
	}
	if c != nodes.Red {
		t.Errorf("node '5' should be 'red' but got '%s'", c)
	}
	c, ok = tree.ColorOf(6)
	if !ok {
		t.Errorf("tree should contain 6")
	}
	if c != nodes.Black {
		t.Errorf("node '6' should be 'black' but got '%s'", c)
	}
	c, ok = tree.ColorOf(7)
	if !ok {
		t.Errorf("tree should contain 7")
	}
	if c != nodes.Red {
		t.Errorf("node '7' should be 'black' but got '%s'", c)
	}
}

// TestRedBlack_Contains tests the Contains method
//
// after it is called this test must assert:
//   - the element is found if it was inserted
//   - the element is not found if it was not inserted
func TestRedBlack_Contains(t *testing.T) {
	cmp := comparators.Ordered[int]{}
	tree := RedBlack[int](cmp)
	if tree.Contains(0) {
		t.Errorf("tree should not contain 0")
	}

	iterables.OfSlice(1, 2, 3, 5, 6, 7).EachN(func(i, v int) {
		tree.Insert(v)
		if !tree.Contains(v) {
			t.Errorf("tree should contain %d", v)
		}
		if tree.Contains(v + 1) {
			t.Errorf("tree should not contain %d", v+1)
		}
	})
}

// TestRedBlack_Delete tests the Delete method
//
// after each deletion the tree state should assert
//   - its size is decrease
//   - the method Contains returns false for the deleted value
//   - the iteration cannot find the deleted value
func TestRedBlack_Delete(t *testing.T) {
	cmp := comparators.Ordered[int]{}

	// Test 1: delete from empty tree
	t.Run("empty tree", func(t *testing.T) {
		tree := RedBlack[int](cmp)
		if tree.Size() != 0 {
			t.Errorf("expected tree size 0 but got %d", tree.Size())
		}
		tree.Delete(1)
		if tree.Size() != 0 {
			t.Errorf("expected tree size 0 but got %d", tree.Size())
		}
	})

	// Test 2: after inserting on tree created empty
	t.Run("after inserting on tree created empty", func(t *testing.T) {
		tree := RedBlack[int](cmp)
		cast, ok := tree.(*redBlack[int])
		if !ok {
			t.Errorf("expected tree to be of type redBlack but got %T", tree)
		}
		iterables.OfSlice(1, 2, 3, 5, 6, 7).EachN(func(i, v int) {
			tree.Insert(v)
			if cast._size != 1 {
				t.Errorf("expected tree._size '1', but got '%d'", cast._size)
			}
			if cast._root.Value() != v {
				t.Errorf("expected tree._root.Value() '%d', but got '%d'", v, cast._root.Value())
			}
			if tree.Size() != 1 {
				t.Errorf("expected tree size '1', but got %d", tree.Size())
			}
			if !tree.Contains(v) {
				t.Errorf("tree should contain %d", v)
			}

			tree.Delete(v)
			if cast._size != 0 {
				t.Errorf("expected tree._size '0', but got '%d'", cast._size)
			}
			if cast._root != nil {
				t.Errorf("expected tree._root to be 'nil', but got \n%v", cast._root)
			}
			if tree.Size() != 0 {
				t.Errorf("expected tree size '0', but got %d", tree.Size())
			}
			if tree.Contains(v) {
				t.Errorf("tree should not contain %d", v)
			}
		})
	})

	// Test 3: Deleting from full tree
	t.Run("deleting from full tree", func(t *testing.T) {
		tree := RedBlack[int](cmp, 1, 2, 3, 5, 6, 7, 9, 10, 11, 13, 14, 15)
		cast, ok := tree.(*redBlack[int])
		if !ok {
			t.Errorf("expected tree to be of type redBlack but got %T", tree)
		}
		it := iterables.OfSlice(1, 2, 3, 5, 6, 7, 9, 10, 11, 13, 14, 15)
		for i := 0; i < 12; i++ {
			v, exists := it.First()
			if !exists {
				t.Errorf("iterable shouldn't be empty")
			}

			// assert value can be found
			if !tree.Contains(v) {
				t.Errorf("tree should contain '%d'", v)
			}
			// Assert both sizes are equal
			if cast._size != it.Len() {
				t.Errorf("expected tree._size '%d', but got '%d'", it.Len(), cast._size)
			}
			if tree.Size() != it.Len() {
				t.Errorf("expected tree size '%d', but got %d", it.Len(), tree.Size())
			}

			tree.Delete(v)
			// assert value cannot be found
			if tree.Contains(v) {
				t.Errorf("tree should not contain '%d'", v)
			}
			// remove from iterable
			_, ok := it.Remove(v, cmp.Compare)
			if !ok {
				t.Errorf("iterable should contain '%d'", v)
			}

			// assert sizes still equal
			if cast._size != it.Len() {
				t.Errorf("expected tree._size '%d', but got '%d'", it.Len(), cast._size)
			}
			if tree.Size() != it.Len() {
				t.Errorf("expected tree size '%d', but got %d", it.Len(), tree.Size())
			}
		}
	})
}

func TestRedBlack_Floor(t *testing.T) {
	cmp := comparators.Ordered[int]{}

	// Test 1: Only positive values
	t.Run("only positive values", func(t *testing.T) {
		tree := RedBlack[int](cmp)
		it := iterables.OfSlice(1, 2, 3, 5, 6, 7)
		it.EachN(func(i, v int) {
			tree.Insert(v)
			floor, ok := tree.Floor(v)
			if !ok {
				t.Errorf("tree should find floor of '%d'", v)
			}
			if floor != v {
				t.Errorf("floor should be '%d' but got '%d'", v, floor)
			}
			floor, ok = tree.Floor(v + 1)
			if !ok {
				t.Errorf("tree should find floor of '%d'", v+1)
			}
			if floor != v {
				t.Errorf("floor should be '%d' but got '%d'", v, floor)
			}

			if i > 0 {
				floor, ok = tree.Floor(v - 1)
				if !ok {
					t.Errorf("tree should find floor of '%d'", v-1)
				}
				if floor != it.At(i-1) {
					t.Errorf("tree should contain %d", v-1)
				}
			} else {
				floor, ok = tree.Floor(v - 1)
				if ok {
					t.Errorf("tree should not find floor of '%d'", v-1)
				}
				if floor != 0 {
					t.Errorf("floor should be zero value but got '%d'", floor)
				}
			}
		})
	})

	// Test 2: Only negative values
	t.Run("only positive values", func(t *testing.T) {
		tree := RedBlack[int](cmp)
		it := iterables.OfSlice(-7, -6, -5, -3, -2, -1)
		it.EachN(func(i, v int) {
			tree.Insert(v)
			floor, ok := tree.Floor(v)
			if !ok {
				t.Errorf("tree should find floor of '%d'", v)
			}
			if floor != v {
				t.Errorf("floor should be '%d' but got '%d'", v, floor)
			}
			floor, ok = tree.Floor(v + 1)
			if !ok {
				t.Errorf("tree should find floor of '%d'", v+1)
			}
			if floor != v {
				t.Errorf("floor should be '%d' but got '%d'", v, floor)
			}

			if i > 0 {
				floor, ok = tree.Floor(v - 1)
				if !ok {
					t.Errorf("tree should find floor of '%d'", v-1)
				}
				if floor != it.At(i-1) {
					t.Errorf("tree should contain %d", v-1)
				}
			} else {
				floor, ok = tree.Floor(v - 1)
				if ok {
					t.Errorf("tree should not find floor of '%d'", v-1)
				}
				if floor != 0 {
					t.Errorf("floor should be zero value but got '%d'", floor)
				}
			}
		})
	})

	// Test 1: Positive and negative values
	t.Run("positive and negative values", func(t *testing.T) {
		tree := RedBlack[int](cmp)
		it := iterables.OfSlice(-7, -6, -5, -3, -2, -1, 1, 2, 3, 5, 6, 7)
		it.EachN(func(i, v int) {
			tree.Insert(v)
			floor, ok := tree.Floor(v)
			if !ok {
				t.Errorf("tree should find floor of '%d'", v)
			}
			if floor != v {
				t.Errorf("floor should be '%d' but got '%d'", v, floor)
			}
			floor, ok = tree.Floor(v + 1)
			if !ok {
				t.Errorf("tree should find floor of '%d'", v+1)
			}
			if floor != v {
				t.Errorf("floor should be '%d' but got '%d'", v, floor)
			}

			if i > 0 {
				floor, ok = tree.Floor(v - 1)
				if !ok {
					t.Errorf("tree should find floor of '%d'", v-1)
				}
				if floor != it.At(i-1) {
					t.Errorf("tree should contain %d", v-1)
				}
			} else {
				floor, ok = tree.Floor(v - 1)
				if ok {
					t.Errorf("tree should not find floor of '%d'", v-1)
				}
				if floor != 0 {
					t.Errorf("floor should be zero value but got '%d'", floor)
				}
			}
		})
	})
}

// TestRedBlack_Insert tests the Insert method
//
// after each insertion the tree state should assert
//   - its size is increase
//   - the method Contains returns true for the inserted value
//   - the method Contains returns false for the not inserted value
//   - the iteration can find the inserted value
func TestRedBlack_Insert(t *testing.T) {
	cmp := comparators.Ordered[int]{}
	tree := RedBlack[int](cmp)
	iterables.OfSlice(1, 2, 3, 5, 6, 7).EachN(func(i, v int) {
		tree.Insert(v)
		if tree.Size() != i+1 {
			t.Errorf("expected tree size of '%d', but got '%d'", i+1, tree.Size())
		}
		if !tree.Contains(v) {
			t.Errorf("tree should contain '%d'", v)
		}
		if tree.Contains(v + 1) {
			t.Errorf("tree should not contain '%d'", v+1)
		}
	})
}

func TestRedBlack_IsEmpty(t *testing.T) {
	cmp := comparators.Ordered[int]{}

	// Case 1: Empty tree
	t.Run("empty tree", func(t *testing.T) {
		tree := RedBlack[int](cmp)
		cast, ok := tree.(*redBlack[int])
		if !ok {
			t.Errorf("tree should be of type redBlack")
		}
		if !tree.IsEmpty() {
			t.Errorf("tree should be empty")
		}
		if cast._size != 0 {
			t.Errorf("tree size should be 0")
		}
		if cast._root != nil {
			t.Errorf("tree root should be nil")
		}

		tree.Insert(1)
		if tree.IsEmpty() {
			t.Errorf("tree should not be empty")
		}
		if cast._size != 1 {
			t.Errorf("tree size should be 1")
		}
		if cast._root.Value() != 1 {
			t.Errorf("tree root should be 1")
		}
	})

	// Case 2: Non-empty tree
	t.Run("non-empty tree", func(t *testing.T) {
		tree := RedBlack[int](cmp, 1, 2, 3, 5, 6, 7)
		cast, ok := tree.(*redBlack[int])
		if !ok {
			t.Errorf("tree should be of type redBlack")
		}
		if tree.IsEmpty() {
			t.Errorf("tree should not be empty")
		}
		if cast._size != 6 {
			t.Errorf("tree size should be '6' but got '%d'", cast._size)
		}
		if cast._root.Value() != 2 {
			t.Errorf("tree root should be '2' but got '%d'", cast._root.Value())
		}

		it := iterables.OfSlice(1, 2, 3, 5, 6, 7)
		for i := 0; i < 6; i++ {
			v, ok := it.First()
			if !ok {
				t.Errorf("iterable should not be empty")
			}
			if tree.IsEmpty() {
				t.Errorf("tree should not be empty")
			}
			if cast._size != it.Len() {
				t.Errorf("tree size should be '%d' but got '%d'", it.Len(), cast._size)
			}
			tree.Delete(v)
			it.Remove(v, cmp.Compare)
		}
		if !tree.IsEmpty() {
			t.Errorf("tree should be empty")
		}
	})
}

// TestRedBlack_Iterator tests the rbtIterator.
//
// the iterator should be able to iterate over the tree
// from its minimum to its maximum value
func TestRedBlack_Iterator(t *testing.T) {
	cmp := comparators.Ordered[int]{}

	// Test 1: Empty tree
	t.Run("empty tree", func(t *testing.T) {
		tree := RedBlack[int](cmp)
		it := tree.Iterator()
		if it.HasNext() {
			t.Errorf("iterator should be empty")
		}
		if it.Next() != 0 {
			t.Errorf("iterator should return a zero value as next value when the tree is empty")
		}
	})

	// Test 2: Non-empty tree
	t.Run("non-empty tree", func(t *testing.T) {
		tree := RedBlack[int](cmp, 1, 2, 3, 5, 6, 7)
		it := tree.Iterator()
		iterables.OfSlice(1, 2, 3, 5, 6, 7).Each(func(v int) {
			if !it.HasNext() {
				t.Errorf("iterator should not be empty")
			}
			if it.Next() != v {
				t.Errorf("iterator should return '%d' as next value", v)
			}
		})
		if it.HasNext() {
			t.Errorf("iterator should be empty")
		}
	})
}

// TestRedBlack_Max tests the Max method
//
// after each insertion the tree state should assert
//   - the method Max returns the max value of the tree
func TestRedBlack_Max(t *testing.T) {
	cmp := comparators.Ordered[int]{}

	// Test 1: Empty tree
	t.Run("max on empty tree", func(t *testing.T) {
		tree := RedBlack[int](cmp)
		mxm, ok := tree.Max()
		if ok {
			t.Errorf("tree should be empty but got max value '%d'", mxm)
		}
		if mxm != 0 {
			t.Errorf("tree max should be zero value but got '%d'", mxm)
		}
	})

	// Test 2: after inserting on tree created empty
	t.Run("after inserting on tree created empty", func(t *testing.T) {
		tree := RedBlack[int](cmp)
		iterables.OfSlice(1, 2, 3, 5, 6, 7).EachN(func(i, v int) {
			tree.Insert(v)
			mxm, ok := tree.Max()
			if !ok {
				t.Errorf("tree should find max after inserting '%d'", v)
			}
			if mxm != v {
				t.Errorf("tree max should be '%d' but got '%d'", v, mxm)
			}
		})
	})

	// Test 3: should return always the same value for created fetched tree
	t.Run("should return always the same value for created fetched tree", func(t *testing.T) {
		tree := RedBlack[int](cmp, 1, 2, 3, 5, 6, 7)
		mxm, ok := tree.Max()
		if !ok {
			t.Errorf("tree should find max after inserting '%d'", 7)
		}
		if mxm != 7 {
			t.Errorf("tree max should be '%d' but got '%d'", 7, mxm)
		}
	})
}

// TestRedBlack_Min tests the Min method
//
// after each insertion the tree state should assert
//   - the method Min returns the min value of the tree
func TestRedBlack_Min(t *testing.T) {
	cmp := comparators.Ordered[int]{}

	// Test 1: empty tree
	t.Run("min on empty tree", func(t *testing.T) {
		tree := RedBlack[int](cmp)
		min, ok := tree.Min()
		if ok {
			t.Errorf("tree should be empty but got min value '%d'", min)
		}
		if min != 0 {
			t.Errorf("tree min should be zero value but got '%d'", min)
		}
	})

	// Test 2: after inserting on tree created empty
	t.Run("after inserting on tree created empty", func(t *testing.T) {
		tree := RedBlack[int](cmp)
		iterables.OfSlice(1, 2, 3, 5, 6, 7).EachN(func(i, v int) {
			tree.Insert(v)
			mnm, ok := tree.Min()
			if !ok {
				t.Errorf("tree should find min after inserting '%d'", v)
			}
			if mnm != 1 {
				t.Errorf("tree min should be '1' but got '%d'", mnm)
			}
		})
	})

	// Test 3: should return always the same value for created fetched tree
	t.Run("should return always the same value for created fetched tree", func(t *testing.T) {
		tree := RedBlack[int](cmp, 1, 2, 3, 5, 6, 7)
		mnm, ok := tree.Min()
		if !ok {
			t.Errorf("tree should find min after inserting '%d'", 1)
		}
		if mnm != 1 {
			t.Errorf("tree min should be '1' but got '%d'", mnm)
		}
	})
}

// TestRedBlack_Creation tests the creation of a RBT
//
// after being created the tree should:
//   - contain the same number of nodes as the number of elements passed
func TestConcRedBlack_Creation(t *testing.T) {
	cmp := comparators.Ordered[int]{}
	tree := ConcRedBlack[int](cmp)
	if tree == nil {
		t.Errorf("tree is nil")
	}
	if tree.Size() != 0 {
		t.Errorf("tree size is not 0")
	}

	for i := 0; i < 10; i++ {
		tree.Insert(i)
		if tree.Size() != i+1 {
			t.Errorf("tree size is not %d", i+1)
		}
	}
}

// TestRedBlack_Ceiling tests the Ceiling method
//
// after it is called this test must assert:
//   - the value retrieved is the smallest value greater than
//     or equal to the given value
func TestConcRedBlack_Ceiling(t *testing.T) {
	cmp := comparators.Ordered[int]{}

	// Test 1: Only positive values
	t.Run("only positive values", func(t *testing.T) {
		tree := ConcRedBlack[int](cmp)
		it := iterables.OfSlice(1, 2, 3, 5, 6, 7)
		it.EachN(func(i, v int) {
			tree.Insert(v)
			ceiling, ok := tree.Ceiling(v)
			if !ok {
				t.Errorf("tree should find ceiling of '%d'", v)
			}
			if ceiling != v {
				t.Errorf("ceiling of '%d' should be '%d' but got '%d'", v, v, ceiling)
			}
			ceiling, ok = tree.Ceiling(v + 1)
			if ok {
				t.Errorf("tree shouldn't find ceiling of '%d'", v+1)
			}
			if ceiling != 0 {
				t.Errorf("ceiling of 'value + 1' should be zero value but got '%d'", ceiling)
			}
			if i > 0 {
				if i == 3 {
					ceiling, ok = tree.Ceiling(v - 1)
					if !ok {
						t.Errorf("tree should find ceiling of '%d'", v-1)
					}
					if ceiling != v {
						t.Errorf("ceiling of '%d' should be '%d' but got '%d'", v-1, v, ceiling)
					}
				} else {
					ceiling, ok = tree.Ceiling(v - 1)
					if !ok {
						t.Errorf("tree should find ceiling of '%d'", v-1)
					}
					if ceiling != v-1 {
						t.Errorf("ceiling of '%d' should be '%d' but got '%d'", v-1, v, ceiling)
					}
				}
			} else {
				ceiling, ok = tree.Ceiling(v - 1)
				if !ok {
					t.Errorf("tree should find ceiling of '%d'", v-1)
				}
				if ceiling != it.At(i) {
					t.Errorf("ceiling of '%d' should be '%d' but got '%d'", v-1, it.At(i), ceiling)
				}
			}
		})
	})

	// Test 2: Only negative values
	t.Run("only positive values", func(t *testing.T) {
		tree := ConcRedBlack[int](cmp)
		it := iterables.OfSlice(-7, -6, -5, -3, -2, -1)
		it.EachN(func(i, v int) {
			tree.Insert(v)
			ceiling, ok := tree.Ceiling(v)
			if !ok {
				t.Errorf("tree should find ceiling of '%d'", v)
			}
			if ceiling != v {
				t.Errorf("ceiling of '%d' should be '%d' but got '%d'", v, v, ceiling)
			}
			ceiling, ok = tree.Ceiling(v + 1)
			if ok {
				t.Errorf("tree shouldn't find ceiling of '%d'", v+1)
			}
			if ceiling != 0 {
				t.Errorf("ceiling of 'value + 1' should be zero value but got '%d'", ceiling)
			}
			if i > 0 {
				if i == 3 {
					ceiling, ok = tree.Ceiling(v - 1)
					if !ok {
						t.Errorf("tree should find ceiling of '%d'", v-1)
					}
					if ceiling != v {
						t.Errorf("ceiling of '%d' should be '%d' but got '%d'", v-1, v, ceiling)
					}
				} else {
					ceiling, ok = tree.Ceiling(v - 1)
					if !ok {
						t.Errorf("tree should find ceiling of '%d'", v-1)
					}
					if ceiling != v-1 {
						t.Errorf("ceiling of '%d' should be '%d' but got '%d'", v-1, v, ceiling)
					}
				}
			} else {
				ceiling, ok = tree.Ceiling(v - 1)
				if !ok {
					t.Errorf("tree should find ceiling of '%d'", v-1)
				}
				if ceiling != it.At(i) {
					t.Errorf("ceiling of '%d' should be '%d' but got '%d'", v-1, it.At(i), ceiling)
				}
			}
		})
	})

	// Test 1: Positive and negative values
	t.Run("positive and negative values", func(t *testing.T) {
		tree := ConcRedBlack[int](cmp)
		it := iterables.OfSlice(-7, -6, -5, -3, -2, -1, 1, 2, 3, 5, 6, 7)
		it.EachN(func(i, v int) {
			tree.Insert(v)
			ceiling, ok := tree.Ceiling(v)
			if !ok {
				t.Errorf("tree should find ceiling of '%d'", v)
			}
			if ceiling != v {
				t.Errorf("ceiling of '%d' should be '%d' but got '%d'", v, v, ceiling)
			}
			ceiling, ok = tree.Ceiling(v + 1)
			if ok {
				t.Errorf("tree shouldn't find ceiling of '%d'", v+1)
			}
			if ceiling != 0 {
				t.Errorf("ceiling of 'value + 1' should be zero value but got '%d'", ceiling)
			}
			if i > 0 {
				if i == 3 || i == 6 || i == 9 {
					ceiling, ok = tree.Ceiling(v - 1)
					if !ok {
						t.Errorf("tree should find ceiling of '%d'", v-1)
					}
					if ceiling != v {
						t.Errorf("ceiling of '%d' should be '%d' but got '%d'", v-1, v, ceiling)
					}
				} else {
					ceiling, ok = tree.Ceiling(v - 1)
					if !ok {
						t.Errorf("tree should find ceiling of '%d'", v-1)
					}
					if ceiling != v-1 {
						t.Errorf("ceiling of '%d' should be '%d' but got '%d'", v-1, v-1, ceiling)
					}
				}
			} else {
				ceiling, ok = tree.Ceiling(v - 1)
				if !ok {
					t.Errorf("tree should find ceiling of '%d'", v-1)
				}
				if ceiling != it.At(i) {
					t.Errorf("ceiling of '%d' should be '%d' but got '%d'", v-1, it.At(i), ceiling)
				}
			}
		})
	})
}

// TestConcRedBlack_Clear tests the Clear method
//
// after it is called this test must assert:
//   - the size of the tree is 0
func TestConcRedBlack_Clear(t *testing.T) {
	cmp := comparators.Ordered[int]{}
	tree := ConcRedBlack[int](cmp)
	iterables.OfSlice(1, 2, 3, 5, 6, 7).EachN(func(i, v int) {
		tree.Insert(v)
		if tree.Size() != 1 {
			t.Errorf("expected tree size of 1 but got %d", tree.Size())
		}
		tree.Clear()
		if tree.Size() != 0 {
			t.Errorf("expected tree size of 0 but got %d", tree.Size())
		}
	})
}

func TestConcRedBlack_ColorOf(t *testing.T) {
	cmp := comparators.Ordered[int]{}
	tree := ConcRedBlack[int](cmp, 1, 2, 3, 5, 6, 7)
	// Expected tree
	//
	c, ok := tree.ColorOf(1)
	if !ok {
		t.Errorf("tree should contain 1")
	}
	if c != nodes.Black {
		t.Errorf("node '1' should be 'black' but got '%s'", c)
	}
	c, ok = tree.ColorOf(2)
	if !ok {
		t.Errorf("tree should contain 2")
	}
	if c != nodes.Black {
		t.Errorf("node '2' should be 'black' but got '%s'", c)
	}
	c, ok = tree.ColorOf(3)
	if !ok {
		t.Errorf("tree should contain 3")
	}
	if c != nodes.Black {
		t.Errorf("node '3' should be 'black' but got '%s'", c)
	}
	c, ok = tree.ColorOf(4)
	if ok {
		t.Errorf("tree should not contain 4")
	}
	c, ok = tree.ColorOf(5)
	if !ok {
		t.Errorf("tree should contain 5")
	}
	if c != nodes.Red {
		t.Errorf("node '5' should be 'red' but got '%s'", c)
	}
	c, ok = tree.ColorOf(6)
	if !ok {
		t.Errorf("tree should contain 6")
	}
	if c != nodes.Black {
		t.Errorf("node '6' should be 'black' but got '%s'", c)
	}
	c, ok = tree.ColorOf(7)
	if !ok {
		t.Errorf("tree should contain 7")
	}
	if c != nodes.Red {
		t.Errorf("node '7' should be 'black' but got '%s'", c)
	}
}

// TestConcRedBlack_Contains tests the Contains method
//
// after it is called this test must assert:
//   - the element is found if it was inserted
//   - the element is not found if it was not inserted
func TestConcRedBlack_Contains(t *testing.T) {
	cmp := comparators.Ordered[int]{}
	tree := ConcRedBlack[int](cmp)
	if tree.Contains(0) {
		t.Errorf("tree should not contain 0")
	}

	iterables.OfSlice(1, 2, 3, 5, 6, 7).EachN(func(i, v int) {
		tree.Insert(v)
		if !tree.Contains(v) {
			t.Errorf("tree should contain %d", v)
		}
		if tree.Contains(v + 1) {
			t.Errorf("tree should not contain %d", v+1)
		}
	})
}

// TestConcRedBlack_Delete tests the Delete method
//
// after each deletion the tree state should assert
//   - its size is decrease
//   - the method Contains returns false for the deleted value
//   - the iteration cannot find the deleted value
func TestConcRedBlack_Delete(t *testing.T) {
	cmp := comparators.Ordered[int]{}

	// Test 1: delete from empty tree
	t.Run("empty tree", func(t *testing.T) {
		tree := ConcRedBlack[int](cmp)
		if tree.Size() != 0 {
			t.Errorf("expected tree size 0 but got %d", tree.Size())
		}
		tree.Delete(1)
		if tree.Size() != 0 {
			t.Errorf("expected tree size 0 but got %d", tree.Size())
		}
	})

	// Test 2: after inserting on tree created empty
	t.Run("after inserting on tree created empty", func(t *testing.T) {
		tree := ConcRedBlack[int](cmp)
		cast, ok := tree.(*concRedBlack[int])
		if !ok {
			t.Errorf("expected tree to be of type redBlack but got %T", tree)
		}
		iterables.OfSlice(1, 2, 3, 5, 6, 7).EachN(func(i, v int) {
			tree.Insert(v)
			if cast._size != 1 {
				t.Errorf("expected tree._size '1', but got '%d'", cast._size)
			}
			if cast._root.Value() != v {
				t.Errorf("expected tree._root.Value() '%d', but got '%d'", v, cast._root.Value())
			}
			if tree.Size() != 1 {
				t.Errorf("expected tree size '1', but got %d", tree.Size())
			}
			if !tree.Contains(v) {
				t.Errorf("tree should contain %d", v)
			}

			tree.Delete(v)
			if cast._size != 0 {
				t.Errorf("expected tree._size '0', but got '%d'", cast._size)
			}
			if cast._root != nil {
				t.Errorf("expected tree._root to be 'nil', but got \n%v", cast._root)
			}
			if tree.Size() != 0 {
				t.Errorf("expected tree size '0', but got %d", tree.Size())
			}
			if tree.Contains(v) {
				t.Errorf("tree should not contain %d", v)
			}
		})
	})

	// Test 3: Deleting from full tree
	t.Run("deleting from full tree", func(t *testing.T) {
		tree := ConcRedBlack[int](cmp, 1, 2, 3, 5, 6, 7, 9, 10, 11, 13, 14, 15)
		cast, ok := tree.(*concRedBlack[int])
		if !ok {
			t.Errorf("expected tree to be of type redBlack but got %T", tree)
		}
		it := iterables.OfSlice(1, 2, 3, 5, 6, 7, 9, 10, 11, 13, 14, 15)
		for i := 0; i < 12; i++ {
			v, exists := it.First()
			if !exists {
				t.Errorf("iterable shouldn't be empty")
			}

			// assert value can be found
			if !tree.Contains(v) {
				t.Errorf("tree should contain '%d'", v)
			}
			// Assert both sizes are equal
			if cast._size != it.Len() {
				t.Errorf("expected tree._size '%d', but got '%d'", it.Len(), cast._size)
			}
			if tree.Size() != it.Len() {
				t.Errorf("expected tree size '%d', but got %d", it.Len(), tree.Size())
			}

			tree.Delete(v)
			// assert value cannot be found
			if tree.Contains(v) {
				t.Errorf("tree should not contain '%d'", v)
			}
			// remove from iterable
			_, ok := it.Remove(v, cmp.Compare)
			if !ok {
				t.Errorf("iterable should contain '%d'", v)
			}

			// assert sizes still equal
			if cast._size != it.Len() {
				t.Errorf("expected tree._size '%d', but got '%d'", it.Len(), cast._size)
			}
			if tree.Size() != it.Len() {
				t.Errorf("expected tree size '%d', but got %d", it.Len(), tree.Size())
			}
		}
	})
}

func TestConcRedBlack_Floor(t *testing.T) {
	cmp := comparators.Ordered[int]{}

	// Test 1: Only positive values
	t.Run("only positive values", func(t *testing.T) {
		tree := ConcRedBlack[int](cmp)
		it := iterables.OfSlice(1, 2, 3, 5, 6, 7)
		it.EachN(func(i, v int) {
			tree.Insert(v)
			floor, ok := tree.Floor(v)
			if !ok {
				t.Errorf("tree should find floor of '%d'", v)
			}
			if floor != v {
				t.Errorf("floor should be '%d' but got '%d'", v, floor)
			}
			floor, ok = tree.Floor(v + 1)
			if !ok {
				t.Errorf("tree should find floor of '%d'", v+1)
			}
			if floor != v {
				t.Errorf("floor should be '%d' but got '%d'", v, floor)
			}

			if i > 0 {
				floor, ok = tree.Floor(v - 1)
				if !ok {
					t.Errorf("tree should find floor of '%d'", v-1)
				}
				if floor != it.At(i-1) {
					t.Errorf("tree should contain %d", v-1)
				}
			} else {
				floor, ok = tree.Floor(v - 1)
				if ok {
					t.Errorf("tree should not find floor of '%d'", v-1)
				}
				if floor != 0 {
					t.Errorf("floor should be zero value but got '%d'", floor)
				}
			}
		})
	})

	// Test 2: Only negative values
	t.Run("only positive values", func(t *testing.T) {
		tree := ConcRedBlack[int](cmp)
		it := iterables.OfSlice(-7, -6, -5, -3, -2, -1)
		it.EachN(func(i, v int) {
			tree.Insert(v)
			floor, ok := tree.Floor(v)
			if !ok {
				t.Errorf("tree should find floor of '%d'", v)
			}
			if floor != v {
				t.Errorf("floor should be '%d' but got '%d'", v, floor)
			}
			floor, ok = tree.Floor(v + 1)
			if !ok {
				t.Errorf("tree should find floor of '%d'", v+1)
			}
			if floor != v {
				t.Errorf("floor should be '%d' but got '%d'", v, floor)
			}

			if i > 0 {
				floor, ok = tree.Floor(v - 1)
				if !ok {
					t.Errorf("tree should find floor of '%d'", v-1)
				}
				if floor != it.At(i-1) {
					t.Errorf("tree should contain %d", v-1)
				}
			} else {
				floor, ok = tree.Floor(v - 1)
				if ok {
					t.Errorf("tree should not find floor of '%d'", v-1)
				}
				if floor != 0 {
					t.Errorf("floor should be zero value but got '%d'", floor)
				}
			}
		})
	})

	// Test 1: Positive and negative values
	t.Run("positive and negative values", func(t *testing.T) {
		tree := ConcRedBlack[int](cmp)
		it := iterables.OfSlice(-7, -6, -5, -3, -2, -1, 1, 2, 3, 5, 6, 7)
		it.EachN(func(i, v int) {
			tree.Insert(v)
			floor, ok := tree.Floor(v)
			if !ok {
				t.Errorf("tree should find floor of '%d'", v)
			}
			if floor != v {
				t.Errorf("floor should be '%d' but got '%d'", v, floor)
			}
			floor, ok = tree.Floor(v + 1)
			if !ok {
				t.Errorf("tree should find floor of '%d'", v+1)
			}
			if floor != v {
				t.Errorf("floor should be '%d' but got '%d'", v, floor)
			}

			if i > 0 {
				floor, ok = tree.Floor(v - 1)
				if !ok {
					t.Errorf("tree should find floor of '%d'", v-1)
				}
				if floor != it.At(i-1) {
					t.Errorf("tree should contain %d", v-1)
				}
			} else {
				floor, ok = tree.Floor(v - 1)
				if ok {
					t.Errorf("tree should not find floor of '%d'", v-1)
				}
				if floor != 0 {
					t.Errorf("floor should be zero value but got '%d'", floor)
				}
			}
		})
	})
}

// TestConcRedBlack_Insert tests the Insert method
//
// after each insertion the tree state should assert
//   - its size is increase
//   - the method Contains returns true for the inserted value
//   - the method Contains returns false for the not inserted value
//   - the iteration can find the inserted value
func TestConcRedBlack_Insert(t *testing.T) {
	cmp := comparators.Ordered[int]{}
	tree := ConcRedBlack[int](cmp)
	iterables.OfSlice(1, 2, 3, 5, 6, 7).EachN(func(i, v int) {
		tree.Insert(v)
		if tree.Size() != i+1 {
			t.Errorf("expected tree size of '%d', but got '%d'", i+1, tree.Size())
		}
		if !tree.Contains(v) {
			t.Errorf("tree should contain '%d'", v)
		}
		if tree.Contains(v + 1) {
			t.Errorf("tree should not contain '%d'", v+1)
		}
	})
}

func TestConcRedBlack_IsEmpty(t *testing.T) {
	cmp := comparators.Ordered[int]{}

	// Case 1: Empty tree
	t.Run("empty tree", func(t *testing.T) {
		tree := ConcRedBlack[int](cmp)
		cast, ok := tree.(*concRedBlack[int])
		if !ok {
			t.Errorf("tree should be of type redBlack")
		}
		if !tree.IsEmpty() {
			t.Errorf("tree should be empty")
		}
		if cast._size != 0 {
			t.Errorf("tree size should be 0")
		}
		if cast._root != nil {
			t.Errorf("tree root should be nil")
		}

		tree.Insert(1)
		if tree.IsEmpty() {
			t.Errorf("tree should not be empty")
		}
		if cast._size != 1 {
			t.Errorf("tree size should be 1")
		}
		if cast._root.Value() != 1 {
			t.Errorf("tree root should be 1")
		}
	})

	// Case 2: Non-empty tree
	t.Run("non-empty tree", func(t *testing.T) {
		tree := ConcRedBlack[int](cmp, 1, 2, 3, 5, 6, 7)
		cast, ok := tree.(*concRedBlack[int])
		if !ok {
			t.Errorf("tree should be of type redBlack")
		}
		if tree.IsEmpty() {
			t.Errorf("tree should not be empty")
		}
		if cast._size != 6 {
			t.Errorf("tree size should be '6' but got '%d'", cast._size)
		}
		if cast._root.Value() != 2 {
			t.Errorf("tree root should be '2' but got '%d'", cast._root.Value())
		}

		it := iterables.OfSlice(1, 2, 3, 5, 6, 7)
		for i := 0; i < 6; i++ {
			v, ok := it.First()
			if !ok {
				t.Errorf("iterable should not be empty")
			}
			if tree.IsEmpty() {
				t.Errorf("tree should not be empty")
			}
			if cast._size != it.Len() {
				t.Errorf("tree size should be '%d' but got '%d'", it.Len(), cast._size)
			}
			tree.Delete(v)
			it.Remove(v, cmp.Compare)
		}
		if !tree.IsEmpty() {
			t.Errorf("tree should be empty")
		}
	})
}

// TestConcRedBlack_Iterator tests the rbtIterator.
//
// the iterator should be able to iterate over the tree
// from its minimum to its maximum value
func TestConcRedBlack_Iterator(t *testing.T) {
	cmp := comparators.Ordered[int]{}

	// Test 1: Empty tree
	t.Run("empty tree", func(t *testing.T) {
		tree := ConcRedBlack[int](cmp)
		it := tree.Iterator()
		if it.HasNext() {
			t.Errorf("iterator should be empty")
		}
		if it.Next() != 0 {
			t.Errorf("iterator should return a zero value as next value when the tree is empty")
		}
	})

	// Test 2: Non-empty tree
	t.Run("non-empty tree", func(t *testing.T) {
		tree := ConcRedBlack[int](cmp, 1, 2, 3, 5, 6, 7)
		it := tree.Iterator()
		iterables.OfSlice(1, 2, 3, 5, 6, 7).Each(func(v int) {
			if !it.HasNext() {
				t.Errorf("iterator should not be empty")
			}
			if it.Next() != v {
				t.Errorf("iterator should return '%d' as next value", v)
			}
		})
		if it.HasNext() {
			t.Errorf("iterator should be empty")
		}
	})
}

// TestConcRedBlack_Max tests the Max method
//
// after each insertion the tree state should assert
//   - the method Max returns the max value of the tree
func TestConcRedBlack_Max(t *testing.T) {
	cmp := comparators.Ordered[int]{}

	// Test 1: Empty tree
	t.Run("max on empty tree", func(t *testing.T) {
		tree := ConcRedBlack[int](cmp)
		mxm, ok := tree.Max()
		if ok {
			t.Errorf("tree should be empty but got max value '%d'", mxm)
		}
		if mxm != 0 {
			t.Errorf("tree max should be zero value but got '%d'", mxm)
		}
	})

	// Test 2: after inserting on tree created empty
	t.Run("after inserting on tree created empty", func(t *testing.T) {
		tree := ConcRedBlack[int](cmp)
		iterables.OfSlice(1, 2, 3, 5, 6, 7).EachN(func(i, v int) {
			tree.Insert(v)
			mxm, ok := tree.Max()
			if !ok {
				t.Errorf("tree should find max after inserting '%d'", v)
			}
			if mxm != v {
				t.Errorf("tree max should be '%d' but got '%d'", v, mxm)
			}
		})
	})

	// Test 3: should return always the same value for created fetched tree
	t.Run("should return always the same value for created fetched tree", func(t *testing.T) {
		tree := ConcRedBlack[int](cmp, 1, 2, 3, 5, 6, 7)
		mxm, ok := tree.Max()
		if !ok {
			t.Errorf("tree should find max after inserting '%d'", 7)
		}
		if mxm != 7 {
			t.Errorf("tree max should be '%d' but got '%d'", 7, mxm)
		}
	})
}

// TestConcRedBlack_Min tests the Min method
//
// after each insertion the tree state should assert
//   - the method Min returns the min value of the tree
func TestConcRedBlack_Min(t *testing.T) {
	cmp := comparators.Ordered[int]{}

	// Test 1: empty tree
	t.Run("min on empty tree", func(t *testing.T) {
		tree := ConcRedBlack[int](cmp)
		min, ok := tree.Min()
		if ok {
			t.Errorf("tree should be empty but got min value '%d'", min)
		}
		if min != 0 {
			t.Errorf("tree min should be zero value but got '%d'", min)
		}
	})

	// Test 2: after inserting on tree created empty
	t.Run("after inserting on tree created empty", func(t *testing.T) {
		tree := ConcRedBlack[int](cmp)
		iterables.OfSlice(1, 2, 3, 5, 6, 7).EachN(func(i, v int) {
			tree.Insert(v)
			mnm, ok := tree.Min()
			if !ok {
				t.Errorf("tree should find min after inserting '%d'", v)
			}
			if mnm != 1 {
				t.Errorf("tree min should be '1' but got '%d'", mnm)
			}
		})
	})

	// Test 3: should return always the same value for created fetched tree
	t.Run("should return always the same value for created fetched tree", func(t *testing.T) {
		tree := ConcRedBlack[int](cmp, 1, 2, 3, 5, 6, 7)
		mnm, ok := tree.Min()
		if !ok {
			t.Errorf("tree should find min after inserting '%d'", 1)
		}
		if mnm != 1 {
			t.Errorf("tree min should be '1' but got '%d'", mnm)
		}
	})
}
