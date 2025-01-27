// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package search

import (
	"github.com/andrerrcosta2/gtools/core/data/comparables"
	"github.com/andrerrcosta2/gtools/core/domain/constraints/prim"
	"github.com/andrerrcosta2/gtools/core/domain/gtools"
	"math/rand"
)

func NewQuickSelect[T any](comparator comparables.Comparator[T]) *QuickSelect[T] {
	return &QuickSelect[T]{
		comparator: comparator,
	}
}

type QuickSelect[T any] struct {
	comparator comparables.Comparator[T]
}

func (s *QuickSelect[T]) Search(arr []T, k int) T {
	if len(arr) == 1 {
		return arr[0]
	}

	// Pick a random pivot
	pivotIndex := rand.Intn(len(arr))
	pivot := arr[pivotIndex]

	// Partition the array around the pivot
	var left []T
	var right []T
	for i, v := range arr {
		if i != pivotIndex {
			if s.comparator.Compare(v, pivot) < 0 {
				left = append(left, v)
			} else {
				right = append(right, v)
			}
		}
	}

	if k < len(left) {
		// Recursively search the left part
		return s.Search(left, k)
	} else if k > len(left) {
		// Recursively search the right part
		return s.Search(right, k-len(left)-1)
	}

	// The pivot is the k-th smallest element
	return pivot
}

// Quick finds the k-th smallest element in an unordered list.
func Quick[T prim.Ordered](arr []T, k int) T {
	if len(arr) == 1 {
		return arr[0]
	}

	// Pick a random pivot
	pivotIndex := rand.Intn(len(arr))
	pivot := arr[pivotIndex]

	// Partition the array around the pivot
	var left []T
	var right []T
	for i, v := range arr {
		if i != pivotIndex {
			if v < pivot {
				left = append(left, v)
			} else {
				right = append(right, v)
			}
		}
	}

	if k < len(left) {
		// Recursively search the left part
		return Quick[T](left, k)
	} else if k > len(left) {
		// Recursively search the right part
		return Quick[T](right, k-len(left)-1)
	}

	// The pivot is the k-th smallest element
	return pivot
}

// QuickOf finds the k-th smallest element in an unordered list.
func QuickOf[T gtools.SortableOf](arr []T, k int) T {
	if len(arr) == 1 {
		return arr[0]
	}

	// Pick a random pivot
	pivotIndex := rand.Intn(len(arr))
	pivot := arr[pivotIndex]

	// Partition the array around the pivot
	var left []T
	var right []T
	for i, v := range arr {
		if i != pivotIndex {
			if v.Less(pivot) {
				left = append(left, v)
			} else {
				right = append(right, v)
			}
		}
	}

	if k < len(left) {
		// Recursively search the left part
		return QuickOf(left, k)
	} else if k > len(left) {
		// Recursively search the right part
		return QuickOf(right, k-len(left)-1)
	}

	// The pivot is the k-th smallest element
	return pivot
}
