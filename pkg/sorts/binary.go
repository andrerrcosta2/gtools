// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sorts

import (
	"github.com/andrerrcosta2/gtools/pkg/comparables"
	"github.com/andrerrcosta2/gtools/pkg/constraints"
	"github.com/andrerrcosta2/gtools/pkg/gtools"
)

func NewBinarySort[T any](comparator comparables.Comparator[T]) *BinarySort[T] {
	return &BinarySort[T]{
		comparator: comparator,
	}
}

// BinarySort (Binary Insertion Sort) is a variation of insertion sort that uses binary search to find the
// correct position for inserting an element.
//
// Best case: O(n log(n)) when the list is already sorted or nearly sorted.
// Worst case: O(n^2) in the worst case when the list is in reverse order.
// Average case: O(n^2)) due to the insertion process.
//
// Space Complexity: O(1) due to no additional data structures used.
// Stability: Stable - Preserves the relative order of equal elements.
type BinarySort[T any] struct {
	comparator comparables.Comparator[T]
}

func (b *BinarySort[T]) Sort(arr *[]T) {
	for i := 1; i < len(*arr); i++ {
		// Get the current element
		key := (*arr)[i]

		// Find the correct position to insert the element using binary search
		pos := b.obs(*arr, key, 0, i)

		// Shift elements to the right to make space for the new element
		copy((*arr)[pos+1:i+1], (*arr)[pos:i])

		// Insert the element at the correct position
		(*arr)[pos] = key
	}
}

func (b *BinarySort[T]) obs(arr []T, key T, lo, hi int) int {
	for lo <= hi {
		mid := lo + (hi-lo)/2
		c := b.comparator.Compare(key, arr[mid])
		if c == 0 {
			return mid
		} else if c == 1 {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return lo
}

var _ Sort[any] = (*BinarySort[any])(nil)
var _ Sort[string] = (*BinarySort[string])(nil)

// Binary performs a binary insertion sort on the given slice of elements.
// It sorts the slice in ascending order.
//
// The time complexity of this function is O(n log n) in the average case and O(n^2) in the worst case.
// The space complexity is O(1) as no additional data structures are used.
// The function is stable, meaning it preserves the relative order of equal elements.
func Binary[T constraints.Ordered](arr *[]T) {
	for i := 1; i < len(*arr); i++ {
		key := (*arr)[i]
		pos := obs(*arr, key, 0, i)

		// Shift elements to the right to make space for the new element
		copy((*arr)[pos+1:i+1], (*arr)[pos:i])

		// Insert the element at the correct position
		(*arr)[pos] = key
	}
}

func OptimisticBinary[T constraints.Ordered](arr *[]T) {
	for i := 1; i < len(*arr); i++ {
		key := (*arr)[i]
		pos := obs(*arr, key, 0, i)

		// If the element is already in the correct position, skip
		if pos == i {
			continue
		}

		// Shift elements to make space for key
		copy((*arr)[pos+1:i+1], (*arr)[pos:i])
		(*arr)[pos] = key
	}
}

// obs performs a binary search to find the correct position to insert the given element.
// It searches the range [low, high) in the given slice.
//
// The time complexity of this function is O(log n).
func obs[T constraints.Ordered](arr []T, key T, low, high int) int {
	for low < high {
		// Calculate the midpoint of the range
		mid := (low + high) / 2

		// Compare the key with the element at the midpoint
		if key < arr[mid] {
			// If the key is less than the midpoint, search the left half
			high = mid
		} else {
			// If the key is greater than or equal to the midpoint, search the right half
			low = mid + 1
		}
	}

	return low
}

// BinaryOf performs a binary insertion sort on the given slice of elements.
// It sorts the slice in ascending order.
//
// The time complexity of this function is O(n log n) in the average case and O(n^2) in the worst case.
// The space complexity is O(1) as no additional data structures are used.
// The function is stable, meaning it preserves the relative order of equal elements.
func BinaryOf[T gtools.SortableOf](arr *[]T) {
	for i := 1; i < len(*arr); i++ {
		key := (*arr)[i]
		pos := sbs(*arr, key, 0, i)

		// Shift elements to the right to make space for the new element
		copy((*arr)[pos+1:i+1], (*arr)[pos:i])

		// Insert the element at the correct position
		(*arr)[pos] = key
	}
}

func OptimisticBinaryOf[T gtools.SortableOf](arr *[]T) {
	for i := 1; i < len(*arr); i++ {
		key := (*arr)[i]
		pos := sbs(*arr, key, 0, i)

		// If the element is already in the correct position, skip
		if pos == i {
			continue
		}

		// Shift elements to make space for key
		copy((*arr)[pos+1:i+1], (*arr)[pos:i])
		(*arr)[pos] = key
	}
}

// sbs finds the correct position for insertion in a sorted slice of SortableOf type.
// It uses a binary search algorithm to achieve this in O(log n) time complexity.
func sbs[T gtools.SortableOf](arr []T, key T, low, high int) int {
	// Continue the search until the low and high indices converge
	for low < high {
		// Calculate the midpoint of the current search range
		mid := (low + high) / 2

		// Compare the key with the element at the midpoint
		if arr[mid].Less(key) {
			// If the key is greater, move the low index to the right half
			low = mid + 1
		} else {
			// If the key is less or equal, move the high index to the left half
			high = mid
		}
	}
	// Return the final index where the key should be inserted
	return low
}
