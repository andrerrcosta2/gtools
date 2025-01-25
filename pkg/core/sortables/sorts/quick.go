// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sorts

import (
	"github.com/andrerrcosta2/gtools/core/domain/constraints/prim"
	"github.com/andrerrcosta2/gtools/core/domain/gtools"
)

// Quick sorts an array of elements of type T constraints.Ordered using the quicksort algorithm.
// The function recursively partitions the array into two sub-arrays based on a pivot element,
// and then sorts each sub-array separately.
// The pivot element is chosen as the middle element of the array.
// The function returns the sorted array.
func Quick[T prim.Ordered](arr []T) {
	// Base case: if the array has less than 2 elements, it is already sorted
	if len(arr) < 2 {
		return
	}

	// Choose the middle element as the pivot
	pivotIndex := len(arr) / 2
	pivot := (arr)[pivotIndex]

	// Partition the array into two sub-arrays based on the pivot
	left, right := 0, len(arr)-1
	for left <= right {
		for (arr)[left] < pivot {
			left++
		}
		for (arr)[right] > pivot {
			right--
		}
		if left <= right {
			(arr)[left], (arr)[right] = (arr)[right], (arr)[left]
			left++
			right--
		}
	}

	// Recursively sort the sub-arrays
	if right > 0 {
		leftPart := (arr)[:right+1]
		Quick(leftPart)
	}
	if left < len(arr) {
		rightPart := (arr)[left:]
		Quick(rightPart)
	}
}

// QuickOf sorts an array of elements of type T core.SortableOf using the quicksort algorithm.
// The function recursively partitions the array into two sub-arrays based on a pivot element,
// and then sorts each sub-array separately.
// The pivot element is chosen as the middle element of the array.
// The function returns the sorted array.
func QuickOf[T gtools.SortableOf](arr []T) {
	// Base case: if the array has less than 2 elements, it is already sorted
	if len(arr) < 2 {
		return
	}

	// Choose the middle element as the pivot
	pivotIndex := len(arr) / 2
	pivot := (arr)[pivotIndex]

	// Partition the array into two sub-arrays based on the pivot
	left, right := 0, len(arr)-1
	for left <= right {
		for (arr)[left].Less(pivot) {
			left++
		}
		for pivot.Less((arr)[right]) {
			right--
		}
		if left <= right {
			(arr)[left], (arr)[right] = (arr)[right], (arr)[left]
			left++
			right--
		}
	}

	// Recursively sort the sub-arrays
	if right > 0 {
		leftPart := (arr)[:right+1]
		QuickOf(leftPart)
	}
	if left < len(arr) {
		rightPart := (arr)[left:]
		QuickOf(rightPart)
	}
}
