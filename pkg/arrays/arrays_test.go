// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package arrays

import (
	"fmt"
	"github.com/andrerrcosta2/gtools/pkg/functions"
	"github.com/andrerrcosta2/gtools/pkg/testdata/testcomparables"
	"reflect"
	"strings"
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
	index := IndexOf(&arr, 3)
	if index != 2 {
		t.Errorf("IndexOf() = %v, want %v", index, 2)
	}

	arr2 := testcomparables.RandomStructs(5).Values()
	index2 := IndexOf[testcomparables.ComparableStruct](&arr2, arr2[2])
	if index2 != 2 {
		t.Errorf("IndexOf() = %v, want %v", index2, 2)
	}
}

func TestLastIndexOf(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	index := LastIndexOf(&arr, 3)
	if index != 2 {
		t.Errorf("LastIndexOf() = %v, want %v", index, 2)
	}
}

func TestFind(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	index := Find(&arr, func(v int) bool {
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

	result := FindAll(&arr, isPerfectSquare)
	expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	if !Equals(&result, &expected) {
		t.Errorf("FindAll(arr, isPerfectSquare) = %v; want %v", result, expected)
	}

	result = FindAll(&arr, isGreaterThan20)
	expected = []int{4, 5, 6, 7, 8, 9}

	if !Equals(&result, &expected) {
		t.Errorf("FindAll(arr, isGreaterThan20) = %v; want %v", result, expected)
	}
}

func TestCompare(t *testing.T) {
	arr1 := []int{2, 3, 4, 6, 7, 9, 12, 15, 18, 21, 24, 27}
	arr2 := []int{2, 3, 5}

	isMultiple := func(a int, b int) bool {
		return a%b == 0
	}

	result := Compare(&arr1, &arr2, isMultiple)
	expected := []int{0, 1, 2, 3, 5, 6, 7, 8, 9, 10, 11}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestUnique(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	unique := Unique(&arr)
	if unique[0] != 1 || unique[1] != 2 || unique[2] != 3 || unique[3] != 4 || unique[4] != 5 {
		t.Errorf("Unique() = %v, want %v", unique, []int{1, 2, 3, 4, 5})
	}
}

func TestUniqueBy(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	unique := UniqueBy(&arr, func(v int) int {
		return v
	})
	if unique[0] != 1 || unique[1] != 2 || unique[2] != 3 || unique[3] != 4 || unique[4] != 5 {
		t.Errorf("UniqueBy() = %v, want %v", unique, []int{1, 2, 3, 4, 5})
	}
}

func TestEmpty(t *testing.T) {
	notEmpty := []int{1, 2, 3, 4, 5}
	empty := []int{}

	if Empty(&notEmpty) {
		t.Errorf("Empty() = %v, want %v", notEmpty, empty)
	}
	if !Empty(&empty) {
		t.Errorf("Empty() = %v, want %v", empty, notEmpty)
	}
}

func TestOutOfBounds(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	if !OutOfBounds(&arr, 6) {
		t.Errorf("OutOfBounds() = %v, want %v", true, false)
	}
	if OutOfBounds(&arr, 4) {
		t.Errorf("OutOfBounds() = %v, want %v", false, true)
	}
}

func TestContains(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	if !Contains(&arr, 3) {
		t.Errorf("Contains() = %v, want %v", false, true)
	}
	if Contains(&arr, 6) {
		t.Errorf("Contains() = %v, want %v", true, false)
	}
}

func TestContainsBy(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	if !ContainsBy(&arr, 3, func(v, w int) bool {
		return v == w
	}) {
		t.Errorf("ContainsBy() = %v, want %v", false, true)
	}
	if ContainsBy(&arr, 6, func(v, w int) bool {
		return v == w
	}) {
		t.Errorf("ContainsBy() = %v, want %v", true, false)
	}
}

// TestContainsAllBy_Success tests that ContainsAllBy returns true when all elements are found
func TestContainsAllBy_Success(t *testing.T) {
	exp := []Plant{
		{Nm: "Tree"},
		{Nm: "Shrub"},
		{Nm: "Flower"},
	}

	res := []Plant{
		{Nm: "Tree"},
		{Nm: "Shrub"},
		{Nm: "Flower"},
		{Nm: "Grass"},
	}

	result := ContainsAllBy(&exp, &res, func(a, b Plant) bool {
		return a.Nm == b.Nm
	})

	if !result {
		t.Errorf("Expected true, but got false")
	}
}

// TestContainsAllBy_Failure tests that ContainsAllBy returns false when not all elements are found
func TestContainsAllBy_Failure(t *testing.T) {
	exp := []Plant{
		{Nm: "Tree"},
		{Nm: "Shrub"},
		{Nm: "Flower"},
	}

	res := []Plant{
		{Nm: "Tree"},
		{Nm: "Shrub"},
		{Nm: "Grass"},
	}

	result := ContainsAllBy(&exp, &res, func(a, b Plant) bool {
		return a.Nm == b.Nm
	})

	if result {
		t.Errorf("Expected false, but got true")
	}
}

// TestContainsAllBy_EmptyExp tests that ContainsAllBy returns true when the expected slice is empty
func TestContainsAllBy_EmptyExp(t *testing.T) {
	exp := []Plant{}

	res := []Plant{
		{Nm: "Tree"},
		{Nm: "Shrub"},
		{Nm: "Flower"},
		{Nm: "Grass"},
	}

	result := ContainsAllBy(&exp, &res, func(a, b Plant) bool {
		return a.Nm == b.Nm
	})

	if !result {
		t.Errorf("Expected true, but got false")
	}
}

// TestContainsAllBy_EmptyRes tests that ContainsAllBy returns false when the result slice is empty
func TestContainsAllBy_EmptyRes(t *testing.T) {
	exp := []Plant{
		{Nm: "Tree"},
		{Nm: "Shrub"},
	}

	res := []Plant{}

	result := ContainsAllBy(&exp, &res, func(a, b Plant) bool {
		return a.Nm == b.Nm
	})

	if result {
		t.Errorf("Expected false, but got true")
	}
}

// TestContainsAllBy_PartialMatch tests that ContainsAllBy returns false when only some elements are found
func TestContainsAllBy_PartialMatch(t *testing.T) {
	exp := []Plant{
		{Nm: "Tree"},
		{Nm: "Shrub"},
		{Nm: "Flower"},
	}

	res := []Plant{
		{Nm: "Tree"},
		{Nm: "Shrub"},
		{Nm: "Grass"},
		{Nm: "Flower"},
	}

	result := ContainsAllBy(&exp, &res, func(a, b Plant) bool {
		return a.Nm == b.Nm
	})

	if !result {
		t.Errorf("Expected true, but got false")
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

	cmt1 := Fold(&arr, 0, add)
	if cmt1 != 15 {
		t.Errorf("Fold() = %v, want %v", cmt1, 15)
	}
	cmt2 := FoldRight(&arr, 0, add)
	if cmt2 != cmt1 {
		t.Errorf("FoldRight() = %v, want %v", cmt2, cmt1)
	}

	ncm1 := Fold(&arr, 1, div)
	if ncm1 != 0 {
		t.Errorf("Fold() = %v, want %v", ncm1, 0)
	}

	ncm2 := FoldRight(&arr, 1, div)
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

func TestEqual(t *testing.T) {
	// Test Case 1: Equals integer slices
	a1 := []int{1, 2, 3}
	b1 := []int{1, 2, 3}
	if !Equals(&a1, &b1) {
		t.Errorf("Test Case 1 Failed: expected true, got false")
	}

	// Test Case 2: Unequal integer slices (different lengths)
	a2 := []int{1, 2, 3}
	b2 := []int{1, 2, 3, 4}
	if Equals(&a2, &b2) {
		t.Errorf("Test Case 2 Failed: expected false, got true")
	}

	// Test Case 3: Unequal integer slices (different elements)
	a3 := []int{1, 2, 3}
	b3 := []int{1, 2, 4}
	if Equals(&a3, &b3) {
		t.Errorf("Test Case 3 Failed: expected false, got true")
	}

	// Test Case 4: Equals string slices
	a4 := []string{"apple", "banana", "cherry"}
	b4 := []string{"apple", "banana", "cherry"}
	if !Equals(&a4, &b4) {
		t.Errorf("Test Case 4 Failed: expected true, got false")
	}

	// Test Case 5: Unequal string slices (different elements)
	a5 := []string{"apple", "banana", "cherry"}
	b5 := []string{"apple", "banana", "date"}
	if Equals(&a5, &b5) {
		t.Errorf("Test Case 5 Failed: expected false, got true")
	}

	// Test Case 6: Empty slices
	a6 := []int{}
	b6 := []int{}
	if !Equals(&a6, &b6) {
		t.Errorf("Test Case 6 Failed: expected true, got false")
	}
}

func TestEqualBy(t *testing.T) {
	// Test Case 1: Compare integer slices by identity (should behave like Equals)
	a1 := []int{1, 2, 3}
	b1 := []int{1, 2, 3}
	if !EqualsBy(&a1, &b1, functions.Identity[int]) {
		t.Errorf("Test Case 1 Failed: expected true, got false")
	}

	// Test Case 2: Compare integer slices by a custom function (modulus)
	a2 := []int{1, 2, 3}
	b2 := []int{4, 5, 6} // All elements are congruent modulo 3
	if !EqualsBy(&a2, &b2, func(v int) int { return v % 3 }) {
		t.Errorf("Test Case 2 Failed: expected true, got false")
	}

	// Test Case 3: Compare string slices by length
	a3 := []string{"apple", "banana", "melon"}
	b3 := []string{"grape", "orange", "lemon"} // All elements have the same lengths
	if !EqualsBy(&a3, &b3, func(v string) int { return len(v) }) {
		t.Errorf("Test Case 3 Failed: expected true, got false")
	}

	// Test Case 4: Compare string slices by first letter
	a4 := []string{"apple", "banana", "cherry"}
	b4 := []string{"apricot", "blueberry", "cranberry"} // All elements start with the same letters
	if !EqualsBy(&a4, &b4, func(v string) string { return strings.ToLower(string(v[0])) }) {
		t.Errorf("Test Case 4 Failed: expected true, got false")
	}

	// Test Case 5: Unequal slices by the given function
	a5 := []string{"apple", "banana", "cherry"}
	b5 := []string{"date", "elderberry", "fig"} // Different lengths
	if EqualsBy(&a5, &b5, func(v string) int { return len(v) }) {
		t.Errorf("Test Case 5 Failed: expected false, got true")
	}

	// Test Case 6: Compare with empty slices
	a6 := []string{}
	b6 := []string{}
	if !EqualsBy(&a6, &b6, func(v string) string { return v }) {
		t.Errorf("Test Case 6 Failed: expected true, got false")
	}
}
