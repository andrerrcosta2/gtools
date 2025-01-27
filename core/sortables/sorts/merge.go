// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sorts

import (
	"github.com/andrerrcosta2/gtools/core/domain/constraints/prim"
	"github.com/andrerrcosta2/gtools/core/domain/gtools"
)

// Merge sorts an array of elements using the merge sort algorithm.
//
// The function recursively divides the input array into two halves until each half has one or zero elements.
// Then, it merges the halves back together in sorted order.
//
// Parameters:
// - arr: the input array of comparable to be sorted.
//
// Returns:
// - []T: the sorted array of core.SortableOf.
func Merge[T prim.Ordered](arr *[]T) {
	if len(*arr) <= 1 {
		return
	}

	mid := len(*arr) / 2

	left := (*arr)[:mid]
	right := (*arr)[mid:]

	// Recursively sort the left and right halves
	Merge(&left)
	Merge(&right)

	// Merge the sorted halves back into the original slice
	merged := merge(left, right)
	copy(*arr, merged)
}

// merge combines two sorted slices into a single sorted slice.
//
// The function takes two slices of ordered elements as input, left and right.
// It returns a new slice that contains all elements from both input slices in sorted order.
func merge[T prim.Ordered](left, right []T) []T {
	result := make([]T, 0, len(left)+len(right))

	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

// MergeOf sorts an array of elements using the merge sort algorithm.
//
// The function recursively divides the input array into two halves until each half has one or zero elements.
// Then, it merges the halves back together in sorted order.
//
// Parameters:
// - arr: the input array of comparable to be sorted.
//
// Returns:
// - []T: the sorted array of core.SortableOf.
func MergeOf[T gtools.SortableOf](arr *[]T) {
	if len(*arr) <= 1 {
		return
	}

	mid := len(*arr) / 2

	left := (*arr)[:mid]
	right := (*arr)[mid:]

	// Recursively sort the left and right halves
	MergeOf(&left)
	MergeOf(&right)

	// Merge the sorted halves back into the original slice
	merged := mergeOf(left, right)
	copy(*arr, merged)
}

// mergeOf combines two sorted slices into a single sorted slice.
//
// The function takes two slices of ordered elements as input, left and right.
// It returns a new slice that contains all elements from both input slices in sorted order.
func mergeOf[T gtools.SortableOf](left, right []T) []T {
	result := make([]T, 0, len(left)+len(right))

	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i].Less(right[j]) {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}
