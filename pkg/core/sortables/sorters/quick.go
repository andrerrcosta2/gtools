// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sorters

import (
	"github.com/andrerrcosta2/gtools/core/comparables"
)

// Quick creates a new instance of Sorter that uses the quickSorter with the given comparator.
// The comparator is used to compare elements in the array during the sorting process.
//
// Parameters:
// - comparator: A function that takes two elements of type T and returns a boolean indicating their order.
//
// Returns:
// - *quickSorter[T]: A pointer to the newly created quickSorter struct.
func Quick[T any](comparator comparables.Comparator[T]) Sorter[T] {
	return &quickSorter[T]{
		comparator: comparator,
	}
}

// quickSorter is a divide-and-conquer sorting algorithm that works by partitioning the array
// into smaller sub-arrays and then sorting those sub-arrays.
//
// Best case: O(n log(n)) when the pivot divides the array into roughly equal parts.
// Worst case: O(n^2) when the pivot divides the array into roughly unequal parts.
// Average case: O(n log(n)) with good pivot selection.
//
// Space Complexity: O(log(n)) due to recursion stack space.
type quickSorter[T any] struct {
	comparator comparables.Comparator[T]
}

func (s *quickSorter[T]) Sort(arr *[]T) {
	s.quick(arr)
}

func (s *quickSorter[T]) SortP(arr *[]*T) {
	s.quickp(arr)
}

func (s *quickSorter[T]) quick(arr *[]T) {
	work := *arr
	if len(work) < 2 {
		return
	}

	// Choose the last element as the pivot
	pivotIndex := len(work) - 1
	pivot := work[pivotIndex]

	// Partition the array into two parts around the pivot
	leftIndex := 0
	for i := 0; i < pivotIndex; i++ {
		if s.comparator.Compare((work)[i], pivot) < 0 {
			// Swap elements to place smaller elements before the pivot
			work[i], work[leftIndex] = work[leftIndex], work[i]
			leftIndex++
		}
	}

	// Move the pivot to its correct position
	work[leftIndex], work[pivotIndex] = work[pivotIndex], work[leftIndex]

	// Recursively sort the left and right sub-arrays
	leftPart := (work)[:leftIndex]
	rightPart := (work)[leftIndex+1:]

	s.quick(&leftPart)
	s.quick(&rightPart)
}

func (s *quickSorter[T]) quickp(arr *[]*T) {
	work := *arr
	if len(work) < 2 {
		return
	}

	// Choose the last element as the pivot
	pivotIndex := len(work) - 1
	pivot := (work)[pivotIndex]

	// Partition the array into two parts around the pivot
	leftIndex := 0
	for i := 0; i < pivotIndex; i++ {
		if s.comparator.Compare(*work[i], *pivot) < 0 {
			// Swap elements to place smaller elements before the pivot
			(work)[i], (work)[leftIndex] = (work)[leftIndex], (work)[i]
			leftIndex++
		}
	}

	// Move the pivot to its correct position
	work[leftIndex], work[pivotIndex] = work[pivotIndex], work[leftIndex]

	// Recursively sort the left and right sub-arrays
	leftPart := work[:leftIndex]
	rightPart := work[leftIndex+1:]

	s.quickp(&leftPart)
	s.quickp(&rightPart)
}

var _ Sorter[any] = (*quickSorter[any])(nil)
var _ Sorter[string] = (*quickSorter[string])(nil)
