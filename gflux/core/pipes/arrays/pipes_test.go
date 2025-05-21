// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package arrays

import (
	"github.com/andrerrcosta2/gtools/core/seeders/random"
	"github.com/andrerrcosta2/gtools/gflux/core/pipes/internal/tests"
	"github.com/andrerrcosta2/gtools/gtests/testingtools"
	"github.com/andrerrcosta2/gtools/gtests/testingtools/config/testlogs"
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	// helper
	tt := testingtools.LoggersLite(t, testlogs.OnFailure)

	// Testing with primitives
	arr := []int{1, 2, 3, 4, 5}
	Reverse(arr)
	tt.StackLogf("Reverse() = %v\n", arr)
	if arr[0] != 5 || arr[1] != 4 || arr[2] != 3 || arr[3] != 2 || arr[4] != 1 {
		tt.Errorf("Reverse() = %v, want %v", arr, []int{5, 4, 3, 2, 1})
	}

	// Testing with Objects
	arr2, cp := random.Struct[tests.Comparable](5).Duplicate()
	Reverse[tests.Comparable](arr2.Values())
	tt.StackLogf("Reverse() = %v\n", arr2)
	if !IsReversed(arr2.Values(), cp.Values()) {
		tt.Errorf("Reverse() = %v, want %v", arr2, cp)
	}

	tt.PrintLogStack()

}

func TestIndexOf(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	index := IndexOf(arr, 3)
	if index != 2 {
		t.Errorf("IndexOf() = %v, want %v", index, 2)
	}

	arr2 := random.Struct[tests.Comparable](5).Values()
	index2 := IndexOf[tests.Comparable](arr2, arr2[2])
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
	// helper
	tt := testingtools.LoggersLite(t, testlogs.OnFailure)

	tcs := []struct {
		name     string
		input    []int
		expected int
	}{
		{"Single element", []int{5}, 5},
		{"Multiple elements", []int{1, 3, 2, 7, 4}, 7},
		{"Negative numbers", []int{-1, -5, -3, -4}, -1},
		{"WhenAllCancels elements the same", []int{2, 2, 2, 2}, 2},
		{"IsEmpty slice", []int{}, 0}, // Adjust this according to the zero value for T
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			tt.StackLogf("Higher(%v)", tc.input)
			result := Higher(tc.input)
			if result != tc.expected {
				tt.Errorf("Higher(%v) = %v; want %v", tc.input, result, tc.expected)
			} else {
				tt.StackLogf("Higher(%v) equals expected\n", tc.input)
			}
		})
	}

	tt.PrintLogStack()
}

func TestLower(t *testing.T) {
	// helper
	tt := testingtools.LoggersLite(t, testlogs.OnFailure)

	tcs := []struct {
		name     string
		input    []int
		expected int
	}{
		{"Single element", []int{5}, 5},
		{"Multiple elements", []int{1, 3, 2, 7, 4}, 1},
		{"Negative numbers", []int{-1, -5, -3, -4}, -5},
		{"WhenAllCancels elements the same", []int{2, 2, 2, 2}, 2},
		{"IsEmpty slice", []int{}, 0}, // Adjust this according to the zero value for T
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			tt.StackLogf("Lower(%v)", tc.input)
			result := Lower(tc.input)
			if result != tc.expected {
				t.Errorf("Lower(%v) = %v; want %v", tc.input, result, tc.expected)
			} else {
				tt.StackLogf("Lower(%v) equals expected\n", tc.input)
			}
		})
	}

	tt.PrintLogStack()
}

func TestKadane(t *testing.T) {
	// helper
	tt := testingtools.LoggersLite(t, testlogs.OnFailure)

	tcs := []struct {
		name     string
		input    []int
		expected int
	}{
		{"Single element", []int{5}, 5},
		{"WhenAllCancels positive numbers", []int{1, 2, 3, 4}, 10},
		{"WhenAllCancels negative numbers", []int{-1, -2, -3, -4}, -1},
		{"Mixed numbers", []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{"IsEmpty array", []int{}, 0},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			tt.StackLogf("Kadane(%v)", tc.input)
			result := Kadane(tc.input)
			if result != tc.expected {
				t.Errorf("Kadane(%v) = %v; want %v", tc.input, result, tc.expected)
			} else {
				tt.StackLogf("Kadane(%v) equals expected\n", tc.input)
			}
		})
	}

	tt.PrintLogStack()
}

func TestMajority(t *testing.T) {
	// helper
	tt := testingtools.LoggersLite(t, testlogs.OnFailure)

	tcs := []struct {
		name     string
		input    []int
		expected int
	}{
		{"Majority element", []int{1, 2, 1, 1, 3, 1}, 1},
		{"No majority element", []int{1, 2, 3, 4, 5}, 0},
		{"IsEmpty array", []int{}, 0},
		{"WhenAllCancels elements the same", []int{2, 2, 2, 2}, 2},
		{"Multiple elements, no majority", []int{1, 2, 3, 2, 2, 3, 3, 3}, 0},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			tt.StackLogf("Majority(%v)", tc.input)
			result := Majority(tc.input)
			if result != tc.expected {
				t.Errorf("Majority(%v) = %v; want %v", tc.input, result, tc.expected)
			} else {
				tt.StackLogf("Majority(%v) equals expected\n", tc.input)
			}
		})
	}

	tt.PrintLogStack()
}
