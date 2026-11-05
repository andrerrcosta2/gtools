// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sorters

import (
	"github.com/andrerrcosta2/gtools/core/data/comparator"
)

// Quick creates a new instance of Sorter that uses the quickSorter with the given comparator.
// The comparator is used to compare elements in the array during the sorting process.
//
// Parameters:
// - comparator: A function that takes two elements of type T and returns a boolean indicating their order.
//
// Returns:
// - *quickSorter[T]: A pointer to the newly created quickSorter struct.
func Quick[S ~[]E, E any](comparator comparator.EqComparer[E]) Sorter[S, E] {
	return &quickSorter[S, E]{
		comparator: comparator,
	}
}

// quickSorter is a divide-and-conquer sorting algorithm that works by partitioning the array
// into smaller sub-array and then sorting those sub-array.
//
// Best case: O(n log(n)) when the pivot divides the array into roughly compare parts.
// Worst case: O(n^2) when the pivot divides the array into roughly unequal parts.
// Average case: O(n log(n)) with good pivot selection.
//
// Space Complexity: O(log(n)) due to recursion stack space.
type quickSorter[S ~[]E, E any] struct {
	comparator comparator.EqComparer[E]
}

func (s *quickSorter[S, E]) Sort(arr *S) {
	s.quick(arr)
}

func (s *quickSorter[S, E]) quick(arr *S) {
	if len(*arr) < 2 {
		return
	}

	// Choose the last element as the pivot
	pivotIndex := len(*arr) - 1
	pivot := (*arr)[pivotIndex]

	// Partition the array into two parts around the pivot
	leftIndex := 0
	for i := 0; i < pivotIndex; i++ {
		if s.comparator.Compare((*arr)[i], pivot) < 0 {
			// Swap elements to place smaller elements before the pivot
			(*arr)[i], (*arr)[leftIndex] = (*arr)[leftIndex], (*arr)[i]
			leftIndex++
		}
	}

	// Move the pivot to its correct position
	(*arr)[leftIndex], (*arr)[pivotIndex] = (*arr)[pivotIndex], (*arr)[leftIndex]

	// Recursively sort the left and right sub-array
	leftPart := (*arr)[:leftIndex]
	rightPart := (*arr)[leftIndex+1:]

	s.quick(&leftPart)
	s.quick(&rightPart)
}

var _ Sorter[[]any, any] = (*quickSorter[[]any, any])(nil)
var _ Sorter[[]string, string] = (*quickSorter[[]string, string])(nil)
