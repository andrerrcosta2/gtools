// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package search

import (
	"fmt"
	"github.com/andrerrcosta2/gtools/core/data/comparators"
	"github.com/andrerrcosta2/gtools/core/domain/constraints/prim"
	"github.com/andrerrcosta2/gtools/core/domain/gtools"
)

// NewBinary returns a new instance of BinarySearch with the given comparator.
// This function is used to create a new binary search algorithm with a custom comparator.
func NewBinary[T any](comparator comparators.Typed[T]) Search[T] {
	// Return a new BinarySearch instance with the given comparator.
	return &BinarySearch[T]{comparator}
}

// BinarySearch is a classic algorithm used for finding the position of a target value within
// a sorted array or slice. It works by repeatedly dividing the search interval in half.
// In each step, it compares the target value to the middle element of the current interval.
// The complexities are generally the same regardless of the type of the elements in the slice.
//
// Best case: O(log(n))
// Worst case: O(log(n))
// Average case: O(log(n))
// Space: O(1)
type BinarySearch[T any] struct {
	comparator comparators.Typed[T]
}

var _ Search[any] = (*BinarySearch[any])(nil)

// Search finds the position of a target value within a sorted array or slice.
// It uses a binary search algorithm to achieve this in O(log(n)) time complexity.
func (b *BinarySearch[T]) Search(arr []T, t T) (int, bool) {
	// Initialize the search interval boundaries
	low, high := 0, len(arr)

	// Continue the search until the interval is empty
	for low < high {
		// Calculate the midpoint of the current interval
		mid := (low + high) / 2

		// Compare the target value to the middle element of the current interval
		c := b.comparator.Compare(t, arr[mid])

		// Adjust the search interval based on the comparison result
		if c == 1 {
			// Target value is greater than the middle element, move to the right half
			low = mid + 1
		} else {
			// Target value is less than or equal to the middle element, move to the left half
			high = mid
		}
	}

	// Check if the key was found
	if low < len(arr) && b.comparator.Equals(t, arr[low]) {
		return low, true
	}

	// KeyTyped not found
	return low, false
}

// Binary performs a binary search on a sorted array to find the position of a target value.
// It returns the index of the target value if found, or the index where it should be inserted to maintain sorted order.
func Binary[T prim.Ordered](arr []T, key T) (int, bool) {
	// Initialize the search interval boundaries
	low, high := 0, len(arr)

	// Continue the search until the interval is empty
	for low < high {
		// Calculate the midpoint of the current interval
		mid := (low + high) / 2

		// Compare the target value to the middle element of the current interval
		if arr[mid] < key {
			// Target value is greater than the middle element, move to the right half
			low = mid + 1
		} else {
			// Target value is less than or equal to the middle element, move to the left half
			high = mid
		}
	}

	// Check if the key was found
	if low < len(arr) && arr[low] == key {
		return low, true // KeyTyped found at index `low`
	}

	// KeyTyped not found, return the insertion position
	return low, false // `low` is the correct insertion index
}

// BinaryOf performs a binary search on a sorted array to find the position of a target value.
// It returns the index of the target value if found, or the index where it should be inserted to maintain sorted order.
func BinaryOf[T gtools.SortableOf](arr []T, key T) (int, bool) {
	// Initialize the search interval boundaries
	low, high := 0, len(arr)

	// Continue the search until the interval is empty
	for low < high {
		// Calculate the midpoint of the current interval
		mid := (low + high) / 2

		// Compare the target value to the middle element of the current interval
		if arr[mid].Less(key) {
			// Target value is greater than the middle element, move to the right half
			low = mid + 1
		} else {
			// Target value is less than or equal to the middle element, move to the left half
			high = mid
		}
	}

	// Check if the key was found
	if low < len(arr) && arr[low].Equal(key) {
		return low, true
	}

	// KeyTyped not found
	return low, false
}

// BinaryBy performs a binary search on a sorted array to find the position of a target value.
// It returns the index of the target value if found, or the index where it should be inserted to maintain sorted order.
func BinaryBy[T any](arr []T, value T, less func(T, T) int) (int, bool) {
	// Initialize the search interval boundaries
	low, high := 0, len(arr)

	// Continue the search until the interval is empty
	for low < high {
		// Calculate the midpoint of the current interval
		mid := (low + high) / 2
		k := less(arr[mid], value)
		switch k {
		case -1:
			low = mid + 1
		case 0:
			return mid, true
		case 1:
			high = mid
		default:
			panic(fmt.Sprintf("unexpected result from less function: '%d'; must be -1, 0, or 1", k))
		}
	}

	// KeyTyped not found
	return low, false
}
