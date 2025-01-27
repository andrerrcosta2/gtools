// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sorters

import (
	"github.com/andrerrcosta2/gtools/core/data/comparables"
	"github.com/andrerrcosta2/gtools/core/domain/constraints/prim"
)

func Merge[T prim.Ordered](comparator comparables.Comparator[T]) Sorter[T] {
	return &mergeSorter[T]{
		comparator: comparator,
	}
}

// mergeSorter is a classic divide-and-conquer algorithm used for sorting.
//
// Merge sort is particularly useful in situations where stability is important, and where
// a predictable O(n log n) time complexity is necessary regardless of the input.
// It’s also effective when working with large datasets that don’t fit entirely in memory,
// as merge sort is amenable to external sorting techniques.
//
// Best Case: O(n log n) when the array is already sorted.
// Worst Case: O(n log n) when the array is completely unsorted.
// Average Case: O(n log n) when the array is partially sorted.
// Space: Worst Case: O(n)
// Recursion Depth: O(log n)
type mergeSorter[T any] struct {
	comparator comparables.Comparator[T]
}

func (m *mergeSorter[T]) Sort(arr *[]T) {
	m.merge(arr)
}

func (m *mergeSorter[T]) merge(arr *[]T) {
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

func (m *mergeSorter[T]) join(left, right []T) []T {
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

func (m *mergeSorter[T]) SortP(arr *[]*T) {
	m.pmerge(arr)
}

func (m *mergeSorter[T]) pmerge(arr *[]*T) {
	work := *arr
	// Base case: If the array has one or zero elements, it is already sorted.
	if len(work) <= 1 {
		return
	}

	// Find the middle index of the array.
	mid := len(work) / 2

	// Divide the array into two halves by creating slices pointing to the original array.
	left := work[:mid]
	right := work[mid:]

	// Recursively sort the left and right halves.
	m.pmerge(&left)
	m.pmerge(&right)

	// Merge the sorted halves back into the original array.
	*arr = m.pjoin(left, right)
}

func (m *mergeSorter[T]) pjoin(left, right []*T) []*T {
	// Initialize the result slice with a capacity equal to the total length of the input slices.
	result := make([]*T, 0, len(left)+len(right))

	// Initialize indices for the left and right slices.
	i, j := 0, 0

	// Merge smaller elements first.
	for i < len(left) && j < len(right) {
		leftVal := left[i]
		rightVal := right[j]

		// Compare the current elements of the left and right slices.
		if m.comparator.Compare(*leftVal, *rightVal) == 1 {
			// If the left element is smaller, append it to the result slice and move to the next element in the left slice.
			result = append(result, leftVal)
			i++
		} else {
			// If the right element is smaller, append it to the result slice and move to the next element in the right slice.
			result = append(result, rightVal)
			j++
		}
	}

	// Append any remaining elements from the left slice.
	result = append(result, left[i:]...)

	// Append any remaining elements from the right slice.
	result = append(result, right[j:]...)

	return result
}

var _ Sorter[any] = (*mergeSorter[any])(nil)
var _ Sorter[string] = (*mergeSorter[string])(nil)
