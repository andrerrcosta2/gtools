// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sorts

import (
	"github.com/andrerrcosta2/gtools/pkg/comparables"
	"github.com/andrerrcosta2/gtools/pkg/constraints"
	"github.com/andrerrcosta2/gtools/pkg/gtools"
)

func NewMergeSort[T constraints.Ordered](comparator comparables.Comparator[T]) *MergeSort[T] {
	return &MergeSort[T]{
		comparator: comparator,
	}
}

// MergeSort is a classic divide-and-conquer algorithm used for sorting.
//
// Merge sort is particularly useful in situations where stability is important, and where
// a predictable O(n log n) time complexity is needed regardless of the input. It’s also effective
// when working with large datasets that don’t fit entirely in memory, as merge sort is amenable
// to external sorting techniques.
//
// Best Case: O(n log n) when the array is already sorted.
// Worst Case: O(n log n) when the array is completely unsorted.
// Average Case: O(n log n) when the array is partially sorted.
// Space: Worst Case: O(n)
// Recursion Depth: O(log n)
type MergeSort[T any] struct {
	comparator comparables.Comparator[T]
}

func (m *MergeSort[T]) Sort(arr *[]T) {
	m.merge(arr)
}

func (m *MergeSort[T]) merge(arr *[]T) {
	// Base case: If the array has one or zero elements, it is already sorted.
	if len(*arr) <= 1 {
		return
	}

	// Find the middle index of the array.
	mid := len(*arr) / 2

	// Divide the array into two halves by creating slices pointing to the original array.
	left := (*arr)[:mid]
	right := (*arr)[mid:]

	// Recursively sort the left and right halves.
	m.merge(&left)
	m.merge(&right)

	// Merge the sorted halves back into the original array.
	*arr = m.join(left, right)
}

func (m *MergeSort[T]) join(left, right []T) []T {
	// Initialize the result slice with a capacity equal to the total length of the input slices.
	result := make([]T, 0, len(left)+len(right))

	// Initialize indices for the left and right slices.
	i, j := 0, 0

	// Merge smaller elements first.
	for i < len(left) && j < len(right) {
		c := m.comparator.Compare(left[i], right[j])
		// Compare the current elements of the left and right slices.
		if c == 1 {
			// If the left element is smaller, append it to the result slice and move to the next element in the left slice.
			result = append(result, left[i])
			i++
		} else {
			// If the right element is smaller, append it to the result slice and move to the next element in the right slice.
			result = append(result, right[j])
			j++
		}
	}

	// Append any remaining elements from the left slice.
	result = append(result, left[i:]...)

	// Append any remaining elements from the right slice.
	result = append(result, right[j:]...)

	return result
}

var _ Sort[any] = (*MergeSort[any])(nil)
var _ Sort[string] = (*MergeSort[string])(nil)

// Merge sorts an array of elements using the merge sort algorithm.
//
// The function recursively divides the input array into two halves until each half has one or zero elements.
// Then, it merges the halves back together in sorted order.
//
// Parameters:
// - arr: the input array of comparable to be sorted.
//
// Returns:
// - []T: the sorted array of gtools.SortableOf.
func Merge[T constraints.Ordered](arr *[]T) {
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
func merge[T constraints.Ordered](left, right []T) []T {
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

// Merge sorts an array of elements using the merge sort algorithm.
//
// The function recursively divides the input array into two halves until each half has one or zero elements.
// Then, it merges the halves back together in sorted order.
//
// Parameters:
// - arr: the input array of comparable to be sorted.
//
// Returns:
// - []T: the sorted array of gtools.SortableOf.
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
