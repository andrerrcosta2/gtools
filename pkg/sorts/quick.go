// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sorts

import (
	"github.com/andrerrcosta2/gtools/pkg/comparables"
	"github.com/andrerrcosta2/gtools/pkg/constraints"
	"github.com/andrerrcosta2/gtools/pkg/gtools"
)

// NewQuicksort creates a new instance of the Quicksort struct with the given comparator.
// The comparator is used to compare elements in the array during the sorting process.
//
// Parameters:
// - comparator: A function that takes two elements of type T and returns a boolean indicating their order.
//
// Returns:
// - *Quicksort[T]: A pointer to the newly created Quicksort struct.
func NewQuicksort[T any](comparator comparables.Comparator[T]) *Quicksort[T] {
	return &Quicksort[T]{
		comparator: comparator,
	}
}

// Quicksort is a divide-and-conquer sorting algorithm that works by partitioning the array
// into smaller sub-arrays and then sorting those sub-arrays.
//
// Best case: O(n log(n)) when the pivot divides the array into roughly equal parts.
// Worst case: O(n^2) when the pivot divides the array into roughly unequal parts.
// Average case: O(n log(n)) with good pivot selection.
//
// Space Complexity: O(log(n)) due to recursion stack space.
type Quicksort[T any] struct {
	comparator comparables.Comparator[T]
}

func (s *Quicksort[T]) Sort(arr *[]T) {
	s.quick(arr)
}

func (s *Quicksort[T]) quick(arr *[]T) {
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

	// Recursively sort the left and right sub-arrays
	leftPart := (*arr)[:leftIndex]
	rightPart := (*arr)[leftIndex+1:]

	s.quick(&leftPart)
	s.quick(&rightPart)
}

var _ Sort[any] = (*Quicksort[any])(nil)
var _ Sort[string] = (*Quicksort[string])(nil)

// Quick sorts an array of elements of type T constraints.Ordered using the quicksort algorithm.
// The function recursively partitions the array into two sub-arrays based on a pivot element,
// and then sorts each sub-array separately.
// The pivot element is chosen as the middle element of the array.
// The function returns the sorted array.
func Quick[T constraints.Ordered](arr *[]T) {
	// Base case: if the array has less than 2 elements, it is already sorted
	if len(*arr) < 2 {
		return
	}

	// Choose the middle element as the pivot
	pivotIndex := len(*arr) / 2
	pivot := (*arr)[pivotIndex]

	// Partition the array into two sub-arrays based on the pivot
	left, right := 0, len(*arr)-1
	for left <= right {
		for (*arr)[left] < pivot {
			left++
		}
		for (*arr)[right] > pivot {
			right--
		}
		if left <= right {
			(*arr)[left], (*arr)[right] = (*arr)[right], (*arr)[left]
			left++
			right--
		}
	}

	// Recursively sort the sub-arrays
	if right > 0 {
		leftPart := (*arr)[:right+1]
		Quick(&leftPart)
	}
	if left < len(*arr) {
		rightPart := (*arr)[left:]
		Quick(&rightPart)
	}
}

// QuickOf sorts an array of elements of type T gtools.SortableOf using the quicksort algorithm.
// The function recursively partitions the array into two sub-arrays based on a pivot element,
// and then sorts each sub-array separately.
// The pivot element is chosen as the middle element of the array.
// The function returns the sorted array.
func QuickOf[T gtools.SortableOf](arr *[]T) {
	// Base case: if the array has less than 2 elements, it is already sorted
	if len(*arr) < 2 {
		return
	}

	// Choose the middle element as the pivot
	pivotIndex := len(*arr) / 2
	pivot := (*arr)[pivotIndex]

	// Partition the array into two sub-arrays based on the pivot
	left, right := 0, len(*arr)-1
	for left <= right {
		for (*arr)[left].Less(pivot) {
			left++
		}
		for pivot.Less((*arr)[right]) {
			right--
		}
		if left <= right {
			(*arr)[left], (*arr)[right] = (*arr)[right], (*arr)[left]
			left++
			right--
		}
	}

	// Recursively sort the sub-arrays
	if right > 0 {
		leftPart := (*arr)[:right+1]
		QuickOf(&leftPart)
	}
	if left < len(*arr) {
		rightPart := (*arr)[left:]
		QuickOf(&rightPart)
	}
}
