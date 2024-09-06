// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package arrays

import (
	"fmt"
	"github.com/andrerrcosta2/gtools/pkg/sorts"
	"github.com/andrerrcosta2/gtools/pkg/testdata/testcomparables"
	"reflect"
	"testing"
)

type Plant struct {
	Nm string
}

func TestReverse(t *testing.T) {
	// Testing with primitives
	arr := []int{1, 2, 3, 4, 5}
	Reverse(arr)
	fmt.Printf("Reverse() = %v\n", arr)
	if arr[0] != 5 || arr[1] != 4 || arr[2] != 3 || arr[3] != 2 || arr[4] != 1 {
		t.Errorf("Reverse() = %v, want %v", arr, []int{5, 4, 3, 2, 1})
	}

	// Testing with Objects
	arr2, ref := testcomparables.RandomStructs(5).Duplicate()
	Reverse(arr2)
	fmt.Printf("Reverse() = %v\n", arr2)
	if !IsReversed(arr2, ref) {
		t.Errorf("Reverse() = %v, want %v", arr2, testcomparables.RandomStructs(5))
	}

}

func TestIndexOf(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	index := IndexOf(arr, 3)
	if index != 2 {
		t.Errorf("IndexOf() = %v, want %v", index, 2)
	}

	arr2 := testcomparables.RandomStructs(5).Values()
	index2 := IndexOf[testcomparables.ComparableStruct](arr2, arr2[2])
	if index2 != 2 {
		t.Errorf("IndexOf() = %v, want %v", index2, 2)
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

	result := FindAll(arr, isPerfectSquare)
	expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	if !Equals(result, expected) {
		t.Errorf("FindAll(arr, isPerfectSquare) = %v; want %v", result, expected)
	}

	result = FindAll(arr, isGreaterThan20)
	expected = []int{4, 5, 6, 7, 8, 9}

	if !Equals(result, expected) {
		t.Errorf("FindAll(arr, isGreaterThan20) = %v; want %v", result, expected)
	}
}

func TestCompare(t *testing.T) {
	arr1 := []int{2, 3, 4, 6, 7, 9, 12, 15, 18, 21, 24, 27}
	arr2 := []int{2, 3, 5}

	isMultiple := func(a int, b int) bool {
		return a%b == 0
	}

	result := Compare(arr1, arr2, isMultiple)
	expected := []int{0, 1, 2, 3, 5, 6, 7, 8, 9, 10, 11}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
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

func TestOutOfBounds(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	if !OutOfBounds(arr, 6) {
		t.Errorf("OutOfBounds() = %v, want %v", true, false)
	}
	if OutOfBounds(arr, 4) {
		t.Errorf("OutOfBounds() = %v, want %v", false, true)
	}
}

func TestSorted(t *testing.T) {
	arr := []int{5, 4, 3, 2, 1}
	sorted := Sorted[*sorts.Quicksort[int]](arr)
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

func TestHigher(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"Single element", []int{5}, 5},
		{"Multiple elements", []int{1, 3, 2, 7, 4}, 7},
		{"Negative numbers", []int{-1, -5, -3, -4}, -1},
		{"All elements the same", []int{2, 2, 2, 2}, 2},
		{"Empty slice", []int{}, 0}, // Adjust this according to the zero value for T
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Higher(tt.input)
			if result != tt.expected {
				t.Errorf("Higher(%v) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestLower(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"Single element", []int{5}, 5},
		{"Multiple elements", []int{1, 3, 2, 7, 4}, 1},
		{"Negative numbers", []int{-1, -5, -3, -4}, -5},
		{"All elements the same", []int{2, 2, 2, 2}, 2},
		{"Empty slice", []int{}, 0}, // Adjust this according to the zero value for T
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Lower(tt.input)
			if result != tt.expected {
				t.Errorf("Lower(%v) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestKadane(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"Single element", []int{5}, 5},
		{"All positive numbers", []int{1, 2, 3, 4}, 10},
		{"All negative numbers", []int{-1, -2, -3, -4}, -1},
		{"Mixed numbers", []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{"Empty array", []int{}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Kadane(tt.input)
			if result != tt.expected {
				t.Errorf("Kadane(%v) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestMajority(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"Majority element", []int{1, 2, 1, 1, 3, 1}, 1},
		{"No majority element", []int{1, 2, 3, 4, 5}, 0},
		{"Empty array", []int{}, 0},
		{"All elements the same", []int{2, 2, 2, 2}, 2},
		{"Multiple elements, no majority", []int{1, 2, 3, 2, 2, 3, 3, 3}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Majority(tt.input)
			if result != tt.expected {
				t.Errorf("Majority(%v) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}
