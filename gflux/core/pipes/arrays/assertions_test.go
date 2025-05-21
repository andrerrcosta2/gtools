// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package arrays

import (
	"github.com/andrerrcosta2/gtools/core/domain/functions"
	"github.com/andrerrcosta2/gtools/core/seeders/random"
	"github.com/andrerrcosta2/gtools/gflux/core/pipes/internal/tests"
	"strings"
	"testing"
)

func TestEmpty(t *testing.T) {
	notEmpty := []int{1, 2, 3, 4, 5}
	var empty []int

	if Empty(notEmpty) {
		t.Errorf("IsEmpty() = %v, want %v", notEmpty, empty)
	}
	if !Empty(empty) {
		t.Errorf("IsEmpty() = %v, want %v", empty, notEmpty)
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
	if !ContainsBy(arr, 3, func(v, w int) bool {
		return v == w
	}) {
		t.Errorf("ContainsBy() = %v, want %v", false, true)
	}
	if ContainsBy(arr, 6, func(v, w int) bool {
		return v == w
	}) {
		t.Errorf("ContainsBy() = %v, want %v", true, false)
	}
}

// TestContainsAllBy_Success tests that ContainsAllBy returns true when all elements are found
func TestContainsAllBy_Success(t *testing.T) {
	dat, dup := random.Struct[tests.Comparable](10).Duplicate()

	// Expects a part to be contained
	if !ContainsAllBy(dat.Some(7).Values(), dat.Values(), functions.Equality[tests.Comparable]) {
		t.Errorf("expected all elements to be found, but got false\n")
	}

	// Expects all elements to be contained
	if !ContainsAllBy(dat.Values(), dup.Values(), functions.Equality[tests.Comparable]) {
		t.Errorf("expected all elements to be found, but got false\n")
	}

	// Expects false after addition
	dup.Append(random.Struct[tests.Comparable](2).Values()...)
	if ContainsAllBy(dup.Values(), dat.Values(), functions.Equality[tests.Comparable]) {
		t.Errorf("expected false, but got true\n")
	}

	// Expected true for empty exp
	if !ContainsAllBy([]tests.Comparable{}, dat.Values(), functions.Equality[tests.Comparable]) {
		t.Errorf("expected true, but got false\n")
	}
}

func TestEqual(t *testing.T) {
	// Test Case 1: Equals integer slices
	a1 := []int{1, 2, 3}
	b1 := []int{1, 2, 3}
	if !Equals(a1, b1) {
		t.Errorf("Test Case 1 Failed: expected true, got false")
	}

	// Test Case 2: Unequal integer slices (different lengths)
	a2 := []int{1, 2, 3}
	b2 := []int{1, 2, 3, 4}
	if Equals(a2, b2) {
		t.Errorf("Test Case 2 Failed: expected false, got true")
	}

	// Test Case 3: Unequal integer slices (different elements)
	a3 := []int{1, 2, 3}
	b3 := []int{1, 2, 4}
	if Equals(a3, b3) {
		t.Errorf("Test Case 3 Failed: expected false, got true")
	}

	// Test Case 4: Equals string slices
	a4 := []string{"apple", "banana", "cherry"}
	b4 := []string{"apple", "banana", "cherry"}
	if !Equals(a4, b4) {
		t.Errorf("Test Case 4 Failed: expected true, got false")
	}

	// Test Case 5: Unequal string slices (different elements)
	a5 := []string{"apple", "banana", "cherry"}
	b5 := []string{"apple", "banana", "date"}
	if Equals(a5, b5) {
		t.Errorf("Test Case 5 Failed: expected false, got true")
	}

	// Test Case 6: IsEmpty slices
	a6 := []int{}
	b6 := []int{}
	if !Equals(a6, b6) {
		t.Errorf("Test Case 6 Failed: expected true, got false")
	}
}

func TestEqualBy(t *testing.T) {
	// Test Case 1: Compare integer slices by identity (should behave like Equals)
	a1 := []int{1, 2, 3}
	b1 := []int{1, 2, 3}
	if !EqualsByHash(a1, b1, functions.Identity[int]) {
		t.Errorf("Test Case 1 Failed: expected true, got false")
	}

	// Test Case 2: Compare integer slices by a custom function (modulus)
	a2 := []int{1, 2, 3}
	b2 := []int{4, 5, 6} // WhenAllCancels elements are congruent modulo 3
	if !EqualsByHash(a2, b2, func(v int) int { return v % 3 }) {
		t.Errorf("Test Case 2 Failed: expected true, got false")
	}

	// Test Case 3: Compare string slices by length
	a3 := []string{"apple", "banana", "melon"}
	b3 := []string{"grape", "orange", "lemon"} // WhenAllCancels elements have the same lengths
	if !EqualsByHash(a3, b3, func(v string) int { return len(v) }) {
		t.Errorf("Test Case 3 Failed: expected true, got false")
	}

	// Test Case 4: Compare string slices by first letter
	a4 := []string{"apple", "banana", "cherry"}
	b4 := []string{"apricot", "blueberry", "cranberry"} // WhenAllCancels elements start with the same letters
	if !EqualsByHash(a4, b4, func(v string) string { return strings.ToLower(string(v[0])) }) {
		t.Errorf("Test Case 4 Failed: expected true, got false")
	}

	// Test Case 5: Unequal slices by the given function
	a5 := []string{"apple", "banana", "cherry"}
	b5 := []string{"date", "elderberry", "fig"} // Different lengths
	if EqualsByHash(a5, b5, func(v string) int { return len(v) }) {
		t.Errorf("Test Case 5 Failed: expected false, got true")
	}

	// Test Case 6: Compare with empty slices
	var a6 []string
	var b6 []string
	if !EqualsByHash(a6, b6, func(v string) string { return v }) {
		t.Errorf("Test Case 6 Failed: expected true, got false")
	}
}
