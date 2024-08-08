// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package arrays

import (
	"github.com/andrerrcosta2/gtools/pkg/constraints"
	"github.com/andrerrcosta2/gtools/pkg/functions"
	"sort"
)

// Reverse reverses the order of elements in a slice.
//
// Parameters:
// - arr: The slice to be reversed.
//
// Returns:
// - The reversed slice.
func Reverse[T any](arr []T) []T {
	// Use two pointers, one starting from the beginning and the other from the end.
	// Swap the elements at these pointers until they meet in the middle.
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	// Return the reversed slice.
	return arr
}

// Unique returns a new slice containing only the unique elements from the input slice.
// The elements in the input slice must be comparable.
//
// Parameters:
// - arr: the input slice
//
// Returns:
// - []T: a new slice containing only the unique elements from the input slice
func Unique[T comparable](arr []T) []T {
	// Create a map to keep track of the seen elements
	keys := make(map[T]bool)

	// Create an empty slice to store the unique elements
	var list []T

	// Iterate over the input slice
	for _, entry := range arr {
		// Check if the element has already been seen
		if _, value := keys[entry]; !value {
			// If not seen, mark it as seen and add it to the list
			keys[entry] = true
			list = append(list, entry)
		}
	}

	// Return the list of unique elements
	return list
}

// UniqueBy returns a new slice containing only the unique elements from the input slice, based on the provided function.
// The elements in the input slice must be comparable.
//
// Parameters:
// - arr: the input slice
// - f: a function that takes an element of type T and returns a key of type K
//
// Returns:
// - []T: a new slice containing only the unique elements from the input slice
func UniqueBy[T any, K comparable](arr []T, f functions.Function[T, K]) []T {
	// Create a map to keep track of the seen keys
	keys := make(map[K]bool)

	// Create an empty slice to store the unique elements
	var list []T

	// Iterate over the input slice
	for _, entry := range arr {
		// Get the key for the current entry
		key := f(entry)

		// Check if the key has already been seen
		if _, value := keys[key]; !value {
			// If not seen, mark it as seen and add the entry to the list
			keys[key] = true
			list = append(list, entry)
		}
	}

	return list
}

// Last returns the last element of a given slice.
// It panics if the slice is empty.
func Last[T any](arr []T) T {
	// Check if the slice is empty
	if len(arr) == 0 {
		panic("slice is empty")
	}

	// Return the last element of the slice
	return arr[len(arr)-1]
}

// First returns the first element of the given slice.
// If the slice is empty, it returns the zero value of the element type.
func First[T any](arr []T) T {
	if len(arr) == 0 {
		var zeroValue T
		return zeroValue
	}
	return arr[0]
}

// LastIndexOf returns the index of the last occurrence of a given value in a slice.
// If the value is not found, it returns -1.
// It has a complexity of O(n) where n is the length of the slice.
//
// Parameters:
// - arr: the input slice
// - val: the value to search for
//
// Returns:
// - int: the index of the last occurrence of the value, or -1 if not found
func LastIndexOf[T comparable](arr []T, val T) int {
	// Iterate over the slice in reverse order
	for i := len(arr) - 1; i >= 0; i-- {
		// Check if the current element is equal to the target value
		if arr[i] == val {
			// If it is, return the index
			return i
		}
	}
	// If the value is not found, return -1
	return -1
}

// IndexOf returns the index of the first occurrence of a given value in a slice.
// If the value is not found, it returns -1.
// It has a complexity of O(n) where n is the length of the slice.
//
// Parameters:
// - arr: the input slice
// - val: the value to search for
//
// Returns:
// - int: the index of the first occurrence of the value, or -1 if not found
func IndexOf[T comparable](arr []T, val T) int {
	for i, v := range arr {
		if v == val {
			return i
		}
	}
	return -1
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

// Find returns the index of the first element in the slice that satisfies the condition defined by the provided function.
// If no elements satisfy the condition, it returns -1.
// It has a complexity of O(n) where n is the length of the slice.
//
// Parameters:
// - arr: the input slice
// - f: a function that takes an element of type T and returns a boolean
//
// Returns:
// - int: the index of the first element that satisfies the condition, or -1 if no elements satisfy the condition
func Find[T any](arr []T, f functions.Function[T, bool]) int {
	// Iterate over the slice
	for i, v := range arr {
		// Check if the current element satisfies the condition defined by the provided function
		if f(v) {
			// If it does, return the index
			return i
		}
	}
	// If no elements satisfy the condition, return -1
	return -1
}

// FindAll returns all the indexes of elements in the slice that satisfy the condition defined by the provided function.
// It has a complexity of O(n) where n is the length of the slice.
//
// Parameters:
// - arr: the input slice
// - f: a function that takes an element of type T and returns a boolean
//
// Returns:
// - []int: a slice of indexes of elements that satisfy the condition
func FindAll[T any](arr []T, f functions.Function[T, bool]) []int {
	var indexes []int // slice to store the indexes

	// Iterate over the slice
	for i, v := range arr {
		// Check if the current element satisfies the condition defined by the provided function
		if f(v) {
			// If it does, append the index to the slice
			indexes = append(indexes, i)
		}
	}

	// Return the slice of indexes
	return indexes
}

// Compare compares two slices and returns the indexes of elements that satisfy the comparison function.
// It has a complexity of O[m * n] where m is the length of the first slice and n is the length of the second slice.
//
// Parameters:
// - arr1: the first input slice
// - arr2: the second input slice
// - compare: a function that compares elements of type T and R and returns a boolean
//
// Returns:
// - []int: a slice of indexes of elements that satisfy the comparison function
func Compare[T any, R any](arr1 []T, arr2 []R, compare functions.BiFunction[T, R, bool]) []int {
	var indexes []int

	// Iterate over the first slice
	for i, v1 := range arr1 {
		// Iterate over the second slice
		for _, v2 := range arr2 {
			// Check if the current elements satisfy the comparison function
			if compare(v1, v2) {
				// If they do, store the index from the first slice
				indexes = append(indexes, i)
				break
			}
		}
	}

	return indexes
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
func ContainsAllBy[T any, R any](exp []T, res []R, compare functions.BiFunction[T, R, bool]) bool {
	// Iterate over each element in the expected slice
	for _, expStc := range exp {
		found := false
		// Iterate over each element in the result slice
		for _, resStc := range res {
			// Check if the current element in the result slice is equal to the current element in the expected slice using the provided comparison function
			if compare(expStc, resStc) {
				found = true
				break
			}
		}
		// If the current element in the expected slice is not found in the result slice, return false
		if !found {
			return false
		}
	}
	// If all elements in the expected slice are found in the result slice, return true
	return true
}

// ContainsBy checks if a value exists in a slice using a predicate function.
//
// The function iterates over the elements of the slice and applies the predicate
// function to each element. If the predicate function returns the specified value,
// the function returns true. Otherwise, it returns false.
//
// Parameters:
// - arr: the input slice
// - val: the value to search for
// - f: the predicate function that takes an element of the slice and returns a value
// - K: the type of the value returned by the predicate function
//
// Returns:
// - bool: true if the value is found in the slice, false otherwise
func ContainsBy[T any, K comparable](arr []T, val K, f functions.Function[T, K]) bool {
	// Iterate over the elements of the slice
	for _, v := range arr {
		// Check if the predicate function returns the specified value for the current element
		if f(v) == val {
			// If it does, return true
			return true
		}
	}
	// If the value is not found, return false
	return false
}

// Sorted sorts a slice of ordered elements in ascending order.
// It uses the sort.Slice function from the standard library to perform the sorting.
// The function returns the sorted slice.
func Sorted[T constraints.Ordered](arr []T) []T {
	// Sort the slice using the provided less function
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	// Return the sorted slice
	return arr
}

func SortedBy[T any, K constraints.Ordered](arr []T, f functions.Function[T, K]) []T {
	sort.Slice(arr, func(i, j int) bool {
		return f(arr[i]) < f(arr[j])
	})
	return arr
}

// Fold applies a binary function to each element of a slice, starting from an initial value,
// and returns the final result.
//
// Parameters:
// - arr: the input slice
// - f: the binary function to apply to each element
// - initial: the initial value
//
// Returns:
// - R: the final result
func Fold[T, R any](arr []T, initial R, f functions.BiFunction[R, T, R]) R {
	// Initialize the result with the initial value
	result := initial

	// Iterate over each element in the slice
	for _, v := range arr {
		// Apply the binary function to the current element and the previous result
		result = f(result, v)
	}

	// Return the final result
	return result
}

// FoldRight applies a binary function to each element of a slice in reverse order,
// starting from an initial value, and returns the final result.
//
// Parameters:
// - arr: the input slice
// - f: the binary function to apply to each element
// - initial: the initial value
//
// Returns:
// - R: the final result
func FoldRight[T, R any](arr []T, initial R, f functions.BiFunction[T, R, R]) R {
	// Initialize the result with the initial value
	result := initial

	// Iterate over the elements of the slice in reverse order
	for i := len(arr) - 1; i >= 0; i-- {
		// Apply the binary function to the current element and the previous result
		result = f(arr[i], result)
	}

	// Return the final result
	return result
}
