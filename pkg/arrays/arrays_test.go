// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package arrays

import (
	"testing"
)

func TestReverse(t *testing.T) {

	arr := []int{1, 2, 3, 4, 5}
	Reverse(arr)
	if arr[0] != 5 || arr[1] != 4 || arr[2] != 3 || arr[3] != 2 || arr[4] != 1 {
		t.Errorf("Reverse() = %v, want %v", arr, []int{5, 4, 3, 2, 1})
	}
}

func TestIndexOf(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	index := IndexOf(arr, 3)
	if index != 2 {
		t.Errorf("IndexOf() = %v, want %v", index, 2)
	}
}

func TestLastIndexOf(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	index := LastIndexOf(arr, 3)
	if index != 2 {
		t.Errorf("LastIndexOf() = %v, want %v", index, 2)
	}
}

func TestFind(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	index := Find(arr, func(v int) bool {
		return v == 3
	})
	if index != 2 {
		t.Errorf("Find() = %v, want %v", index, 2)
	}
}

func TestFindAll(t *testing.T) {
	arr := []int{1, 4, 9, 16, 25, 36, 49, 64, 81, 100}

	isPerfectSquare := func(x int) bool {
		for i := 1; i*i <= x; i++ {
			if i*i == x {
				return true
			}
		}
		return false
	}

	isGreaterThan20 := func(x int) bool {
		return x > 20
	}

	equal := func(a, b []int) bool {
		if len(a) != len(b) {
			return false
		}
		for i := range a {
			if a[i] != b[i] {
				return false
			}
		}
		return true
	}

	result := FindAll(arr, isPerfectSquare)
	expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	if !equal(result, expected) {
		t.Errorf("FindAll(arr, isPerfectSquare) = %v; want %v", result, expected)
	}

	result = FindAll(arr, isGreaterThan20)
	expected = []int{4, 5, 6, 7, 8, 9}

	if !equal(result, expected) {
		t.Errorf("FindAll(arr, isGreaterThan20) = %v; want %v", result, expected)
	}
}

func TestUnique(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	unique := Unique(arr)
	if unique[0] != 1 || unique[1] != 2 || unique[2] != 3 || unique[3] != 4 || unique[4] != 5 {
		t.Errorf("Unique() = %v, want %v", unique, []int{1, 2, 3, 4, 5})
	}
}

func TestUniqueBy(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	unique := UniqueBy(arr, func(v int) int {
		return v
	})
	if unique[0] != 1 || unique[1] != 2 || unique[2] != 3 || unique[3] != 4 || unique[4] != 5 {
		t.Errorf("UniqueBy() = %v, want %v", unique, []int{1, 2, 3, 4, 5})
	}
}

func TestEmpty(t *testing.T) {
	notEmpty := []int{1, 2, 3, 4, 5}
	empty := []int{}

	if Empty(notEmpty) {
		t.Errorf("Empty() = %v, want %v", notEmpty, empty)
	}
	if !Empty(empty) {
		t.Errorf("Empty() = %v, want %v", empty, notEmpty)
	}
}

func TestOutOfBounds(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	if !OutOfBounds(arr, 6) {
		t.Errorf("OutOfBounds() = %v, want %v", true, false)
	}
	if OutOfBounds(arr, 4) {
		t.Errorf("OutOfBounds() = %v, want %v", false, true)
	}
}

func TestContains(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	if !Contains(arr, 3) {
		t.Errorf("Contains() = %v, want %v", false, true)
	}
	if Contains(arr, 6) {
		t.Errorf("Contains() = %v, want %v", true, false)
	}
}

func TestContainsBy(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	if !ContainsBy(arr, 3, func(v int) int {
		return v
	}) {
		t.Errorf("ContainsBy() = %v, want %v", false, true)
	}
	if ContainsBy(arr, 6, func(v int) int {
		return v
	}) {
		t.Errorf("ContainsBy() = %v, want %v", true, false)
	}
}

func TestSorted(t *testing.T) {
	arr := []int{5, 4, 3, 2, 1}
	sorted := Sorted(arr)
	if sorted[0] != 1 || sorted[1] != 2 || sorted[2] != 3 || sorted[3] != 4 || sorted[4] != 5 {
		t.Errorf("Sorted() = %v, want %v", sorted, []int{1, 2, 3, 4, 5})
	}
}

func TestSortedBy(t *testing.T) {
	arr := []int{5, 4, 3, 2, 1}
	sorted := SortedBy(arr, func(v int) int {
		return v
	})
	if sorted[0] != 1 || sorted[1] != 2 || sorted[2] != 3 || sorted[3] != 4 || sorted[4] != 5 {
		t.Errorf("SortedBy() = %v, want %v", sorted, []int{1, 2, 3, 4, 5})
	}
}

func TestFold(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}

	cmt1 := Fold(arr, 0, add)
	if cmt1 != 15 {
		t.Errorf("Fold() = %v, want %v", cmt1, 15)
	}
	cmt2 := FoldRight(arr, 0, add)
	if cmt2 != cmt1 {
		t.Errorf("FoldRight() = %v, want %v", cmt2, cmt1)
	}

	ncm1 := Fold(arr, 1, div)
	if ncm1 != 0 {
		t.Errorf("Fold() = %v, want %v", ncm1, 0)
	}

	ncm2 := FoldRight(arr, 1, div)
	if ncm2 == ncm1 {
		t.Errorf("FoldRight() = %v, want different from %v", ncm2, ncm1)
	}

}

func add(v int, acc int) int {
	return acc + v
}

func div(acc, v int) int {
	if v == 0 {
		return acc
	}
	return acc / v
}
