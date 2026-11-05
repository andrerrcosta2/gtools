// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sorts

import (
	"github.com/andrerrcosta2/gtools/core/domain/constraints"
	"github.com/andrerrcosta2/gtools/core/domain/constraints/prim"
)

// Quick sorts an array of elements of type T prim.Ordered using the quicksort algorithm.
func Quick[S ~[]T, T prim.Ordered](s S) {
	// Base case: if the array has less than 2 elements, it is already sorted
	if len(s) < 2 {
		return
	}

	// Choose the middle element as the pivot
	pivotIndex := len(s) / 2
	pivot := (s)[pivotIndex]

	// Partition the say into two sub-says based on the pivot
	left, right := 0, len(s)-1
	for left <= right {
		for (s)[left] < pivot {
			left++
		}
		for (s)[right] > pivot {
			right--
		}
		if left <= right {
			(s)[left], (s)[right] = (s)[right], (s)[left]
			left++
			right--
		}
	}

	// Recursively sort the sub-array
	if right > 0 {
		leftPart := (s)[:right+1]
		Quick(leftPart)
	}
	if left < len(s) {
		rightPart := (s)[left:]
		Quick(rightPart)
	}
}

// QuickOf sorts an array of elements of type T core.SortableOf using the quicksort algorithm.
func QuickOf[S ~[]T, T constraints.Sortable[T]](s S) {
	// Base case: if the array has less than 2 elements, it is already sorted
	if len(s) < 2 {
		return
	}

	// Choose the middle element as the pivot
	pivotIndex := len(s) / 2
	pivot := (s)[pivotIndex]

	// Partition the array into two sub-array based on the pivot
	left, right := 0, len(s)-1
	for left <= right {
		for (s)[left].Less(pivot) {
			left++
		}
		for pivot.Less((s)[right]) {
			right--
		}
		if left <= right {
			(s)[left], (s)[right] = (s)[right], (s)[left]
			left++
			right--
		}
	}

	// Recursively sort the sub-array
	if right > 0 {
		leftPart := (s)[:right+1]
		QuickOf(leftPart)
	}
	if left < len(s) {
		rightPart := (s)[left:]
		QuickOf(rightPart)
	}
}

// QuickFunc sorts an array of elements of any type T using a function for comparison for the quicksort algorithm.
func QuickFunc[S ~[]T, T any](s S, less func(T, T) bool) {
	// Base case: if the array has less than 2 elements, it is already sorted
	if len(s) < 2 {
		return
	}

	// Choose the middle element as the pivot
	pivotIndex := len(s) / 2
	pivot := s[pivotIndex]

	// Partition the array into two sub-array based on the pivot
	left, right := 0, len(s)-1
	for left <= right {
		for less((s)[left], pivot) {
			left++
		}
		for less(pivot, (s)[right]) {
			right--
		}
		if left <= right {
			(s)[left], (s)[right] = (s)[right], (s)[left]
			left++
			right--
		}
	}

	// Recursively sort the sub-array
	if right > 0 {
		leftPart := (s)[:right+1]
		QuickFunc(leftPart, less)
	}
	if left < len(s) {
		rightPart := (s)[left:]
		QuickFunc(rightPart, less)
	}
}
