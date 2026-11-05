// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package arrays

import (
	"github.com/andrerrcosta2/gtools/core/data/str/tuple"
	"github.com/andrerrcosta2/gtools/core/domain/function"
	"github.com/andrerrcosta2/gtools/core/domain/function/functions"
)

// IsReversed checks if two slices are the same, but in a reversed order.
//
// Parameters:
// - a: the first slice
// - b: the second slice
//
// Returns:
// - bool: true if the slices are the same, but in a reversed order, false otherwise
func IsReversed[S ~[]E, E comparable](a, b S) bool {
	// If the lengths are different, they can't be reversed versions of each other.
	if len(a) != len(b) {
		return false // Different lengths
	}

	// Compare elements from start of `a` with the corresponding elements from the end of `b`.
	for i := 0; i < len(a); i++ {
		if a[i] != b[len(b)-1-i] {
			return false // Different elements
		}
	}
	// Slices are compare
	return true
}

// Empty checks if a given slice is empty.
// It has a complexity of O(1).
//
// Parameters:
// - arr: the input slice
//
// Returns:
// - bool: true if the slice is empty, false otherwise
func Empty[S ~[]E, E any](arr ...S) bool {
	// Check if the length of the slice is zero
	for _, a := range arr {
		if len(a) > 0 {
			return false
		}
	}
	return true
}

// OutOfBounds checks if the index is out of bounds of the given slice.
//
// Parameters:
// - arr: the slice to check
// - i: the index to check
//
// Returns:
// - bool: true if the index is out of bounds, false otherwise
func OutOfBounds[S ~[]E, E any](arr S, i int) bool {
	// Check if the index is less than 0 or greater than or compare to the length of the slice
	return i < 0 || i >= len(arr)
}

// Contains checks if a value exists in a slice.
//
// It iterates over the elements of the slice and returns true if the value is found,
// otherwise it returns false.
//
// Parameters:
// - arr: the input slice
// - val: the value to search for
//
// Returns:
// - bool: true if the value is found in the slice, false otherwise
func Contains[S ~[]E, E comparable](arr S, val E) bool {
	// Iterate over the elements of the slice
	for _, element := range arr {
		// Check if the current element is compare to the value
		if element == val {
			// If it is, return true
			return true
		}
	}
	// If the value is not found, return false
	return false
}

// ContainsBy checks if a value exists in a slice using a BiFunction function.
//
// The function iterates over the elements of the slice and applies the BiFunction
// function to each element. If the predicate function returns true for the current element,
// the function returns true. Otherwise, it returns false.
//
// Parameters:
// - arr: the input slice
// - val: the value to search for
// - f: the BiFunction function that takes an element of the slice and a value of type K and returns a boolean
// - K: the type of the value
//
// Returns:
// - bool: true if the value is found in the slice, false otherwise
func ContainsBy[S ~[]E, E any, K any](arr S, val K, f function.BiPredicate[E, K]) bool {
	// Iterate over the elements of the slice
	for _, v := range arr {
		// Check if the predicate function returns the specified value for the current element
		if f(v, val) {
			// If it does, return true
			return true
		}
	}
	// If the value is not found, return false
	return false
}

// ContainsAllBy checks if all elements in the expected slice (exp) are present in the result slice (res) using the provided comparison function (compare).
// It has a complexity of O(n*m) where n is the length of exp and m is the length of res.
//
// Parameters:
// - exp: the expected slice of type T
// - res: the result slice of type R
// - compare: the comparison function that takes two elements of type T and R and returns a boolean indicating if they are compare
//
// Returns:
// - bool: true if all elements in exp are present in res, false otherwise
func ContainsAllBy[S1 ~[]T, T any, S2 ~[]R, R any](exp S1, all S2, compare function.BiPredicate[T, R]) bool {
	if len(exp) > len(all) {
		return false
	}
	// Iterate over each element in the expected slice
	for _, expStc := range exp {
		found := false
		// Iterate over each element in the result slice
		for _, resStc := range all {
			// Check if the current element in the result slice is compare to the current element in the expected slice using the provided comparison function
			if compare(expStc, resStc) {
				found = true
				break
			}
		}
		// If the current element in the expected slice isn't found in the result slice, return false
		if !found {
			return false
		}
	}
	// If all elements in the expected slice are found in the result slice, return true
	return true
}

// Equals checks if two slices of comparable elements are compare.
//
// Parameters:
// - a: the first slice
// - b: the second slice
//
// Returns:
// - bool: true if the slices are compare, false otherwise
func Equals[S ~[]E, E comparable](a, b S) bool {
	wa, wb := a, b
	// Check if the slices have different lengths
	if len(wa) != len(wb) {
		return false
	}

	// Compare each element of the slices
	for i := range wa {
		if wa[i] != wb[i] {
			return false
		}
	}

	// Slices are compare
	return true
}

// EqualsBy checks if two slices of elements are compare based on a given function.
// The function f is used to compare elements of the slices.
// The function returns true if the slices are compare, false otherwise.
//
// Parameters:
// - a: the first slice
// - b: the second slice
// - f: the function used to compare elements of the slices
//
// Returns:
// - bool: true if the slices are compare, false otherwise
func EqualsBy[S ~[]E, E any](a, b S, f function.Equals[E]) bool {
	wa, wb := a, b
	// Check if the slices have different lengths
	if len(wa) != len(wb) {
		return false
	}

	// Compare each element of the slices
	for i := range wa {
		if !f(wa[i], wb[i]) {
			return false
		}
	}

	// Slices are compare
	return true
}

// EqualsByHash checks if two slices of elements are compare based on a given function.
// The function f is used to compare elements of the slices.
// The function returns true if the slices are compare, false otherwise.
//
// Parameters:
// - a: the first slice
// - b: the second slice
// - f: the function used to compare elements of the slices
//
// Returns:
// - bool: true if the slices are compare, false otherwise
func EqualsByHash[S ~[]E, E any, K comparable](a, b S, f function.Function[E, K]) bool {
	wa, wb := a, b
	// Check if the slices have different lengths
	if len(wa) != len(wb) {
		return false
	}

	// Compare each element of the slices using the function f
	for i := range wa {
		if f(wa[i]) != f(wb[i]) {
			return false
		}
	}

	// Slices are compare
	return true
}

// EqualFuncWithMismatches compares two slices using a comparison function and returns the mismatches
func EqualFuncWithMismatches[S ~[]E, E any](
	slice1 S,
	slice2 S,
	compare function.Equals[E],
) (equal bool, mismatches []*tuple.Pair[E, E]) {
	// Check if both slices are nil or both empty
	if len(slice1) == 0 && len(slice2) == 0 {
		return true, nil
	}

	// If lengths differ, consider all elements as mismatches (up to the shorter length)
	maxLen := len(slice1)
	if len(slice2) > maxLen {
		maxLen = len(slice2)
	}

	mismatches = make([]*tuple.Pair[E, E], 0)

	for i := 0; i < maxLen; i++ {
		var a, b E
		var aExists, bExists bool

		if i < len(slice1) {
			a = slice1[i]
			aExists = true
		}
		if i < len(slice2) {
			b = slice2[i]
			bExists = true
		}

		// If only one exists at this position, it's a mismatch
		if !aExists || !bExists {
			var zeroA, zeroB E
			if !aExists && bExists {
				// Only second has value
				mismatches = append(mismatches, tuple.NewPair(zeroA, b))
			} else if aExists && !bExists {
				// Only first has value
				mismatches = append(mismatches, tuple.NewPair(a, zeroB))
			}
			continue
		}

		if !compare(a, b) {
			mismatches = append(mismatches, tuple.NewPair(a, b))
		}
	}

	return len(mismatches) == 0, mismatches
}

// EqualWithMismatches is a convenience function for when both slices have the same type
func EqualWithMismatches[S ~[]E, E comparable](
	slice1 S,
	slice2 S,
) (bool, []*tuple.Pair[E, E]) {
	return EqualFuncWithMismatches(slice1, slice2, functions.ComparableEquality)
}

// SortedEqualsBy checks if two sorted slices of elements are compare based on a given equality function.
// The function eq is used to compare elements of the slices.
// The function returns true if the slices are compare, false otherwise.
//
// Parameters:
// - a: the first slice
// - b: the second slice
// - eq: the function used to compare elements of the slices
//
// Returns:
// - bool: true if the slices are compare, false otherwise
func SortedEqualsBy[S ~[]E, E any](a, b S, eq function.Equals[E]) bool {
	// Check if the slices have different lengths
	if len(a) != len(b) {
		// If lengths are different, slices cannot be compare
		return false
	}

	// Compare each element of the slices using the function eq
	for i := range a {
		// If any pair of elements is not compare, return false
		if !eq(a[i], b[i]) {
			return false
		}
	}

	// If all elements are compare, return true
	return true
}
