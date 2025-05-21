// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sorters

import (
	"github.com/andrerrcosta2/gtools/core/data/comparators"
)

// Quick creates a new instance of Sorter that uses the quickSorter with the given comparator.
// The comparator is used to compare elements in the array during the sorting process.
//
// Parameters:
// - comparator: A function that takes two elements of type T and returns a boolean indicating their order.
//
// Returns:
// - *quickSorter[T]: A pointer to the newly created quickSorter struct.
func Quick[T any, S ~[]T](comparator comparators.Typed[T]) Sorter[T, S] {
	return &quickSorter[T, S]{
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
type quickSorter[T any, S ~[]T] struct {
	comparator comparators.Typed[T]
}

func (s *quickSorter[T, S]) Sort(arr *S) {
	s.quick(arr)
}

func (s *quickSorter[T, S]) quick(arr *S) {
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

var _ Sorter[any, []any] = (*quickSorter[any, []any])(nil)
var _ Sorter[string, []string] = (*quickSorter[string, []string])(nil)
