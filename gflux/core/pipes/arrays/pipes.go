// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package arrays

import (
	"github.com/andrerrcosta2/gtools/core/domain/constraints/prim"
	"github.com/andrerrcosta2/gtools/core/domain/function"
)

// InsertAt inserts the element at the specified index from the slice and
// returns true if the index is reachable, false otherwise
func InsertAt[S ~[]E, E any](arr S, value E, index int) (S, bool) {
	if index >= 0 && index < len(arr) {
		return append(arr[:index],
			append(S{value}, arr[index:]...)...), true
	}
	return S{}, false
}

// RemoveByIndex removes the element at the specified index from the slice.
//
// Parameters:
// - arr: the input slice
// - index: the index of the element to remove
//
// Returns:
// - []T: the new slice with the element removed
func RemoveByIndex[S ~[]E, E any](arr S, index int) (S, bool) {
	// Check if the index is within the bounds of the slice
	if index >= 0 && index < len(arr) {
		// Use the append function to remove the element at the specified index
		return append(arr[:index], arr[index+1:]...), true
	}
	// Return the original slice if the index is out of bounds
	return arr, false
}

// Reverse reverses the order of elements in a slice.
//
// Parameters:
// - arr: The slice to be reversed.
//
// Returns:
// - The reversed slice.
func Reverse[S ~[]E, E any](arr S) S {
	// Use two ptrs, one starting from the beginning and the other from the end.
	// Swap the elements at these ptrs until they meet in the middle.
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
func Unique[S ~[]E, E comparable](arr S) S {
	// Create a map to keep track of the seen elements
	keys := make(map[E]bool)

	// Create an empty slice to store the unique elements
	var list []E

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
func UniqueBy[S ~[]E, E any, K comparable](arr S, f function.Function[E, K]) S {
	// Create a map to keep track of the seen keys
	keys := make(map[K]bool)

	// Create an empty slice to store the unique elements
	var list []E

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
func Last[S ~[]E, E any](arr S) E {
	// Check if the slice is empty
	if len(arr) == 0 {
		panic("slice is empty")
	}

	// Return the last element of the slice
	return arr[len(arr)-1]
}

// First returns the first element of the given slice.
// If the slice is empty, it returns the zero value of the element type.
func First[S ~[]E, E any](arr S) E {
	if len(arr) == 0 {
		var zeroValue E
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
func LastIndexOf[S ~[]E, E comparable](arr S, val E) int {
	// Iterate over the slice in reverse order
	for i := len(arr) - 1; i >= 0; i-- {
		// Check if the current element is compare to the target value
		if arr[i] == val {
			// If it is, return the index
			return i
		}
	}
	// If the value is not found, return -1
	return -1
}

// LastIndexOfBy returns the index of the last element in the slice that satisfies the condition defined by the provided function.
// If no elements satisfy the condition, it returns -1.
// It has a complexity of O(n) where n is the length of the slice.
//
// Parameters:
// - arr: the input slice
// - val: the value to search for
// - f: a function that takes an element of type T and returns a boolean
//
// Returns:
// - int: the index of the last element that satisfies the condition, or -1 if no elements satisfy the condition
func LastIndexOfBy[S ~[]E, E any, K any](arr S, val K, f function.BiPredicate[E, K]) int {
	for i := len(arr) - 1; i >= 0; i-- {
		if f(arr[i], val) {
			return i
		}
	}
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
func IndexOf[S ~[]E, E comparable](arr S, val E) int {
	for i, v := range arr {
		if v == val {
			return i
		}
	}
	return -1
}

// IndexOfBy returns the index of the first element in the slice that satisfies the condition defined by the provided function.
// If no elements satisfy the condition, it returns -1.
// It has a complexity of O(n) where n is the length of the slice.
//
// Parameters:
// - arr: the input slice
// - val: the value to search for
// - f: a function that takes an element of type T and returns a boolean
//
// Returns:
// - int: the index of the first element that satisfies the condition, or -1 if no elements satisfy the condition
func IndexOfBy[S ~[]E, E any, K any](arr S, val K, f function.BiPredicate[E, K]) int {
	for i, v := range arr {
		if f(v, val) {
			return i
		}
	}
	return -1
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
func Find[S ~[]E, E any](arr S, f function.Function[E, bool]) int {
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
func FindAll[S ~[]E, E any](arr S, f function.Function[E, bool]) []int {
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
func Compare[S1 ~[]T, T any, S2 ~[]R, R any](arr1 S1, arr2 S2, compare function.BiFunction[T, R, bool]) []int {
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

// Fold applies a bi-function to each element of a slice, starting from an initial value,
// and returns the final result.
//
// Parameters:
// - arr: the input slice
// - f: the bi-function to apply to each element
// - initial: the initial value
//
// Returns:
// - R: the final result
func Fold[S ~[]E, E any, R any](arr S, initial R, f function.BiFunction[R, E, R]) R {
	// Initialize the result with the initial value
	result := initial

	// Iterate over each element in the slice
	for _, v := range arr {
		// Apply the bi-function to the current element and the previous result
		result = f(result, v)
	}

	// Return the final result
	return result
}

// FoldRight applies a bi-function to each element of a slice in reverse order,
// starting from an initial value, and returns the final result.
//
// Parameters:
// - arr: the input slice
// - f: the bi-function to apply to each element
// - initial: the initial value
//
// Returns:
// - R: the final result
func FoldRight[S ~[]E, E any, R any](arr S, initial R, f function.BiFunction[E, R, R]) R {
	// Initialize the result with the initial value
	result := initial

	// Iterate over the elements of the slice in reverse order
	for i := len(arr) - 1; i >= 0; i-- {
		// Apply the bi-function to the current element and the previous result
		result = f(arr[i], result)
	}

	// Return the final result
	return result
}

// Higher returns the highest value from a slice of prim.Ordered elements.
//
// Parameters:
// - arr: a slice of elements that are ordered.
//
// Returns:
// - T: the highest value from the slice.
func Higher[S ~[]E, E prim.Ordered](arr S) E {
	// If the slice is empty, return the zero value of the element type.
	if len(arr) == 0 {
		var zero E
		return zero
	}

	// Initialize the highest value with the first element of the slice.
	out := arr[0]

	// Iterate over the rest of the elements in the slice.
	for _, v := range arr[1:] {
		// If the current element is higher than the highest value seen so far,
		// update the highest value.
		if v > out {
			out = v
		}
	}

	// Return the highest value.
	return out
}

// HigherBy returns the highest value from a slice of elements using a BiFunction function.
//
// The function iterates over the elements of the slice and applies the BiFunction
// function to each element. If the predicate function returns true for the current element,
// the function returns true. Otherwise, it returns false.
//
// Parameters:
// - arr: the input slice
// - f: the BiFunction function that takes an element of the slice and a value of type T and returns a boolean
// - T: the type of the value
//
// Returns:
// - bool: true if the value is found in the slice, false otherwise
func HigherBy[S ~[]E, E any](arr S, f function.BiFunction[E, E, E]) E {
	// If the slice is empty, return the zero value of the element type.
	if len(arr) == 0 {
		var zero E
		return zero
	}

	// Initialize the highest value with the first element of the slice.
	out := arr[0]

	// Iterate over the rest of the elements in the slice.
	for _, v := range arr[1:] {
		// If the current element is higher than the highest value seen so far,
		// update the highest value.
		out = f(v, out)
	}

	// Return the highest value.
	return out
}

// Lower returns the lowest value from a slice of ordered elements.
//
// Parameters:
// - arr: a slice of elements that are ordered.
//
// Returns:
// - T: the lowest value from the slice.
func Lower[S ~[]E, E prim.Ordered](arr S) E {
	// If the slice is empty, return the zero value of the element type.
	if len(arr) == 0 {
		var zero E
		return zero
	}

	// Initialize the highest value with the first element of the slice.
	out := arr[0]

	// Iterate over the rest of the elements in the slice.
	for _, v := range arr[1:] {
		// If the current element is higher than the highest value seen so far,
		// update the lowest value.
		if v < out {
			out = v
		}
	}

	// Return the lowest value.
	return out
}

// LowerBy returns the lowest value from a slice of elements using a BiFunction function.
func LowerBy[S ~[]E, E any](arr S, f function.BiFunction[E, E, E]) E {
	// If the slice is empty, return the zero value of the element type.
	if len(arr) == 0 {
		var zero E
		return zero
	}

	// Initialize the highest value with the first element of the slice.
	out := arr[0]

	// Iterate over the rest of the elements in the slice.
	for _, v := range arr[1:] {
		// If the current element is higher than the highest value seen so far,
		// update the lowest value.
		out = f(v, out)
	}

	// Return the lowest value.
	return out
}

// Kadane returns the maximum sum of a contiguous subarray in an array of integers.
// If the array is empty, it returns 0.
// It has a complexity of O(n) where n is the length of the array.
//
// Example:
//
// Input: [-2, 1, -3, 4, -1, 2, 1, -5, 4]
// Output: 6 is within { 4, -1, 2, 1 } = 6
//
// Parameters:
// - arr: the input array
//
// Returns:
// - int: the maximum sum of a contiguous subarray in the array
func Kadane[S ~[]E, E prim.Ordered](arr S) (maxSoFar E) {
	// If the array is empty, return 0
	if len(arr) == 0 {
		return
	}
	// Initialize maxSoFar and maxEndingHere to the first element
	maxSoFar = arr[0]
	maxEndingHere := arr[0]
	for _, v := range arr[1:] { // start iterating from the second element
		// Update maxEndingHere to be either the current element itself
		// or the sum of maxEndingHere and the current element
		maxEndingHere = max(v, maxEndingHere+v)
		// Update maxSoFar if maxEndingHere is greater
		maxSoFar = max(maxSoFar, maxEndingHere)
	}
	// Return the maximum sum of the array
	return
}

// Majority returns the element that appears more than half the time in an array.
// If the array is empty, it returns the zero value of the element type.
// It uses the Boyer-Moore Majority Vote Algorithm
// It has a time complexity of O(n) where n is the length of the array and space complexity of O(1).
//
// Parameters:
// - arr: the input array
//
// Returns:
// - T: the element that appears more than half the time
func Majority[S ~[]E, E comparable](arr S) E {
	var zero E
	var k E
	var count int

	for _, e := range arr {
		if count == 0 {
			k = e
		}
		if e == k {
			count++
		} else {
			count--
		}
	}

	// Check if the element that appears more than half the time is the majority.
	// Without this loop, the function would be unsafe in cases where you need
	// to guarantee that the element returned is indeed the majority element
	// because this algorithm only points that if there's any majority element
	// must be the candidate.
	count = 0
	for _, e := range arr {
		if e == k {
			count++
		}
	}

	if count > len(arr)/2 {
		return k
	}

	return zero
}
