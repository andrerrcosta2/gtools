// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package arrays

import "github.com/andrerrcosta2/gtools/core/domain/functions"

// IsReversed checks if two slices are the same, but in a reversed order.
//
// Parameters:
// - a: the first slice
// - b: the second slice
//
// Returns:
// - bool: true if the slices are the same, but in a reversed order, false otherwise
func IsReversed[T comparable](a, b []T) bool {
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
	// Slices are equal
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
func Empty[T any](arr []T) bool {
	// Check if the length of the slice is zero
	return len(arr) == 0
}

// OutOfBounds checks if the index is out of bounds of the given slice.
//
// Parameters:
// - arr: the slice to check
// - i: the index to check
//
// Returns:
// - bool: true if the index is out of bounds, false otherwise
func OutOfBounds[T any](arr []T, i int) bool {
	// Check if the index is less than 0 or greater than or equal to the length of the slice
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
func Contains[T comparable](arr []T, val T) bool {
	// Iterate over the elements of the slice
	for _, element := range arr {
		// Check if the current element is equal to the value
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
func ContainsBy[T any, K any](arr []T, val K, f functions.BiPredicate[T, K]) bool {
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
// - compare: the comparison function that takes two elements of type T and R and returns a boolean indicating if they are equal
//
// Returns:
// - bool: true if all elements in exp are present in res, false otherwise
func ContainsAllBy[T any, R any](exp []T, all []R, compare functions.BiPredicate[T, R]) bool {
	if len(exp) > len(all) {
		return false
	}
	// Iterate over each element in the expected slice
	for _, expStc := range exp {
		found := false
		// Iterate over each element in the result slice
		for _, resStc := range all {
			// Check if the current element in the result slice is equal to the current element in the expected slice using the provided comparison function
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

// Equals checks if two slices of comparable elements are equal.
//
// Parameters:
// - a: the first slice
// - b: the second slice
//
// Returns:
// - bool: true if the slices are equal, false otherwise
func Equals[T comparable](a, b []T) bool {
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

	// Slices are equal
	return true
}

// EqualsBy checks if two slices of elements are equal based on a given function.
// The function f is used to compare elements of the slices.
// The function returns true if the slices are equal, false otherwise.
//
// Parameters:
// - a: the first slice
// - b: the second slice
// - f: the function used to compare elements of the slices
//
// Returns:
// - bool: true if the slices are equal, false otherwise
func EqualsBy[A any](a, b []A, f functions.BiPredicate[A, A]) bool {
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

	// Slices are equal
	return true
}

// EqualsByHash checks if two slices of elements are equal based on a given function.
// The function f is used to compare elements of the slices.
// The function returns true if the slices are equal, false otherwise.
//
// Parameters:
// - a: the first slice
// - b: the second slice
// - f: the function used to compare elements of the slices
//
// Returns:
// - bool: true if the slices are equal, false otherwise
func EqualsByHash[T any, K comparable](a, b []T, f functions.Function[T, K]) bool {
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

	// Slices are equal
	return true
}

// SortedEqualsBy checks if two sorted slices of elements are equal based on a given equality function.
// The function eq is used to compare elements of the slices.
// The function returns true if the slices are equal, false otherwise.
//
// Parameters:
// - a: the first slice
// - b: the second slice
// - eq: the function used to compare elements of the slices
//
// Returns:
// - bool: true if the slices are equal, false otherwise
func SortedEqualsBy[T any](a, b []T, eq functions.BiPredicate[T, T]) bool {
	// Check if the slices have different lengths
	if len(a) != len(b) {
		// If lengths are different, slices cannot be equal
		return false
	}

	// Compare each element of the slices using the function eq
	for i := range a {
		// If any pair of elements is not equal, return false
		if !eq(a[i], b[i]) {
			return false
		}
	}

	// If all elements are equal, return true
	return true
}
