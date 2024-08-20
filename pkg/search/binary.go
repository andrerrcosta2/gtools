// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package search

import (
	"github.com/andrerrcosta2/gtools/pkg/comparables"
	"github.com/andrerrcosta2/gtools/pkg/constraints"
	"github.com/andrerrcosta2/gtools/pkg/gtools"
)

// NewBinarySearch returns a new instance of BinarySearch with the given comparator.
// This function is used to create a new binary search algorithm with a custom comparator.
func NewBinarySearch[T constraints.Ordered](comparator comparables.Comparator[T]) Search[T] {
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
	comparator comparables.Comparator[T]
}

var _ Search[any] = (*BinarySearch[any])(nil)

// Search finds the position of a target value within a sorted array or slice.
// It uses a binary search algorithm to achieve this in O(log(n)) time complexity.
func (b *BinarySearch[T]) Search(arr []T, t T) int {
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

	// Return the final position of the target value
	return low
}

// Binary performs a binary search on a sorted array to find the position of a target value.
// It returns the index of the target value if found, or the index where it should be inserted to maintain sorted order.
func Binary[T constraints.Ordered](arr []T, key T) int {
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

	// Return the final position of the target value
	return low
}

// BinaryOf performs a binary search on a sorted array to find the position of a target value.
// It returns the index of the target value if found, or the index where it should be inserted to maintain sorted order.
func BinaryOf[T gtools.SortableOf](arr []T, key T) int {
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

	// Return the final position of the target value
	return low
}
