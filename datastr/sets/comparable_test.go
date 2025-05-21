// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sets

import (
	"github.com/andrerrcosta2/gtools/core/data/str/iterables"
	"testing"
)

// TestComparableSet_Creation tests the creation of a cmpSet
//   - expects a non-nil set with the correct number of elements
func TestComparableSet_Creation(t *testing.T) {
	set := Comparable(1, 2, 3, 5, 6, 7)

	if set == nil {
		t.Errorf("Expected a non-nil set, but got nil")
	}

	if len(set.(*cmpSet[int]).set) != 6 {
		t.Errorf("Expected set length 6, but got %d", set.Len())
	}
}

// TestComparableSet_Len tests the len method of a cmpSet
//   - expects the set to have the correct number of elements
func TestComparableSet_Len(t *testing.T) {
	// Test 1: empty set
	set := Comparable[int]()

	if set.Len() != 0 {
		t.Errorf("Expected set length 0, but got %d", set.Len())
	}

	// Test 2: after each insertion
	it := iterables.OfSlice(1, 2, 3, 5, 6, 7)
	it.EachN(func(i int, v int) {
		set.Add(v)
		if set.Len() != i+1 {
			t.Errorf("Expected set length %d, but got %d", i+1, set.Len())
		}
	})

	// Test 3: after creation
	set = Comparable(1, 2, 3, 5, 6, 7)

	if set.Len() != 6 {
		t.Errorf("Expected set length 6, but got %d", set.Len())
	}

	// Test 4: after removal
	it.EachN(func(i int, v int) {
		set.Remove(v)
		if set.Len() != 6-i-1 {
			t.Errorf("Expected set length %d, but got %d", 6-i-1, set.Len())
		}
	})
}

// TestComparableSet_Add tests the addition of elements to a cmpSet
//   - expects the set to have the correct number of elements
//   - expects the set to have added only elements that doesn't already exist
func TestComparableSet_Add(t *testing.T) {
	set := Comparable(1, 2, 3, 5, 6, 7)
	// Add unique elements
	set.Add(4)

	if set.Len() != 7 {
		t.Errorf("Expected set length 7, but got %d", set.Len())
	}

	// Add existing elements
	set.Add(5)

	if set.Len() != 7 {
		t.Errorf("Expected set length 7, but got %d", set.Len())
	}
}

// TestComparableSet_Remove tests the removal of elements from a cmpSet
//   - expects the set to have the correct number of elements
//   - expects the set to have removed only elements that exist
func TestComparableSet_Remove(t *testing.T) {
	set := Comparable(1, 2, 3, 5, 6, 7)
	// Remove existing elements
	set.Remove(5)

	if set.Len() != 5 {
		t.Errorf("Expected set length 5, but got %d", set.Len())
	}

	// Remove non-existing elements
	set.Remove(8)

	if set.Len() != 5 {
		t.Errorf("Expected set length 5, but got %d", set.Len())
	}
}

func TestComparableSet_Contains(t *testing.T) {
	set := Comparable(1, 2, 3, 5, 6, 7)

	if !set.Has(5) {
		t.Errorf("Expected set to contain 5, but it doesn't")
	}

	if set.Has(8) {
		t.Errorf("Expected set to not contain 8, but it does")
	}
}

func TestComparableSet_Values(t *testing.T) {
	set := Comparable(1, 2, 3, 5, 6, 7)

	values := set.Values()

	if len(values) != 6 {
		t.Errorf("Expected set length 6, but got %d", len(values))
	}

	for _, value := range values {
		if !set.Has(value) {
			t.Errorf("Expected set to contain %d, but it doesn't", value)
		}
	}

	// remove an element and check if it's not in the set
	set.Remove(5)
	if set.Has(5) {
		t.Errorf("Expected set to not contain 5, but it does")
	}
}

func TestComparableSet_Clear(t *testing.T) {
	set := Comparable(1, 2, 3, 5, 6, 7)

	if set.Len() != 6 {
		t.Errorf("Expected set length 6, but got %d", set.Len())
	}

	set.Clear()

	if set.Len() != 0 {
		t.Errorf("Expected set length 0, but got %d", set.Len())
	}
}

func TestComparableSet_Equals(t *testing.T) {
	set1 := Comparable(1, 2, 3, 5, 6, 7)
	set2 := Comparable(1, 2, 3, 5, 6, 7)
	set3 := Comparable(1, 2, 3, 5, 6, 8)

	if !set1.Equals(set2) {
		t.Errorf("Expected sets to be equal, but they are not")
	}

	if set1.Equals(set3) {
		t.Errorf("Expected sets to not be equal, but they are")
	}
}

// TestConcComparable tests the creation of a concCmp
//   - expects the set to be not nil
//   - expects the set to have the correct number of elements
func TestConcComparable(t *testing.T) {
	set := ConcComparable(1, 2, 3, 5, 6, 7)

	if set == nil {
		t.Errorf("Expected set to be not nil, but it is")
	}

	if len(set.(*concCmp[int]).set) != 6 {
		t.Errorf("Expected set length 6, but got %d", set.Len())
	}
}

func TestConcComparableSet_Len(t *testing.T) {
	set := ConcComparable(1, 2, 3, 5, 6, 7)

	if len(set.(*concCmp[int]).set) != 6 {
		t.Errorf("Expected set length 6, but got %d", set.Len())
	}

	if set.Len() != 6 {
		t.Errorf("Expected set length 6, but got %d", set.Len())
	}

	// delete a value and check if the length is correct
	delete(set.(*concCmp[int]).set, 5)
	if set.Len() != 5 {
		t.Errorf("Expected set length 5, but got %d", set.Len())
	}
}

// TestConcComparableSet_Add tests the addition of elements to a concCmp
//   - expects the set to have the correct number of elements after adding
func TestConcComparableSet_Add(t *testing.T) {
	set := ConcComparable(1, 2, 3, 5, 6, 7)

	if set.Len() != 6 {
		t.Errorf("Expected set length 6, but got %d", set.Len())
	}

	// Add unique elements
	set.Add(4)
	set.Add(8)

	if len(set.(*concCmp[int]).set) != 8 {
		t.Errorf("Expected set length 8, but got %d", set.Len())
	}
}

func TestConcCmpSet_Has(t *testing.T) {
	set := ConcComparable(1, 2, 3, 5, 6, 7)

	if !set.Has(5) {
		t.Errorf("Expected set to contain 5, but it doesn't")
	}

	if set.Has(8) {
		t.Errorf("Expected set to not contain 8, but it does")
	}

	// Add unique elements
	set.Add(4)
	set.Add(8)

	if !set.Has(4) {
		t.Errorf("Expected set to contain 4, but it doesn't")
	}

	if !set.Has(8) {
		t.Errorf("Expected set to contain 8, but it doesn't")
	}

	// delete a value and check if it's not in the set
	delete(set.(*concCmp[int]).set, 5)
	if set.Has(5) {
		t.Errorf("Expected set to not contain 5, but it does")
	}
}

// TestConcComparableSet_Remove tests the removal of elements from a concCmp
//   - expects the set to have the correct number of elements after removing
func TestConcComparableSet_Remove(t *testing.T) {
	set := ConcComparable(1, 2, 3, 5, 6, 7)

	if set.Len() != 6 {
		t.Errorf("Expected set length 6, but got %d", set.Len())
	}

	// Remove existing elements
	set.Remove(5)
	set.Remove(7)

	if len(set.(*concCmp[int]).set) != 4 {
		t.Errorf("Expected set length 4, but got %d", set.Len())
	}

	if set.Has(5) {
		t.Errorf("Expected set to not contain 5, but it does")
	}

	if set.Has(7) {
		t.Errorf("Expected set to not contain 7, but it does")
	}
}

func TestConcCmpSet_Values(t *testing.T) {
	set := ConcComparable(1, 2, 3, 5, 6, 7)

	values := set.Values()

	if len(values) != 6 {
		t.Errorf("Expected set length 6, but got %d", len(values))
	}

	for _, value := range values {
		if !set.Has(value) {
			t.Errorf("Expected set to contain %d, but it doesn't", value)
		}
	}

	// remove an element and check if it's not in the set
	set.Remove(5)

	values = set.Values()

	for _, value := range values {
		if value == 5 {
			t.Errorf("Expected set to not contain 5, but it does")
		}
	}
}

func TestConcCmpSet_Equals(t *testing.T) {
	set1 := ConcComparable(1, 2, 3, 5, 6, 7)
	set2 := ConcComparable(1, 2, 3, 5, 6, 7)

	if !set1.Equals(set2) {
		t.Errorf("Expected sets to be equal, but they are not")
	}

	set3 := ConcComparable(1, 2, 3, 5, 6)
	if set1.Equals(set3) {
		t.Errorf("Expected sets to be not equal, but they are")
	}
}
