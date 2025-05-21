// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package search

import (
	"github.com/andrerrcosta2/gtools/core/data/comparators"
	"github.com/andrerrcosta2/gtools/core/domain/constraints/prim"
	"github.com/andrerrcosta2/gtools/core/domain/gtools"
)

func NewQuickSelect[T any](comparator comparators.Typed[T]) *QuickSelect[T] {
	return &QuickSelect[T]{
		cmp: comparator,
	}
}

type QuickSelect[T any] struct {
	cmp comparators.Typed[T]
}

func (s *QuickSelect[T]) Search(arr []T, k int) T {
	if len(arr) == 1 {
		return arr[0]
	}
	low, high := 0, len(arr)-1
	for low < high {
		pivotIndex := s.med3(arr, low, high)
		pivotIndex = s.part(arr, low, high, pivotIndex)

		if pivotIndex == k {
			return arr[pivotIndex]
		} else if pivotIndex > k {
			high = pivotIndex - 1
		} else {
			low = pivotIndex + 1
		}
	}
	return arr[low]
}

// med3 chooses a pivot from three random points to reduce worst-case scenarios
func (s *QuickSelect[T]) med3(arr []T, low, high int) int {
	mid := (low + high) / 2
	if s.cmp.Compare(arr[mid], arr[low]) == -1 {
		arr[low], arr[mid] = arr[mid], arr[low]
	}
	if s.cmp.Compare(arr[mid], arr[high]) == -1 {
		arr[mid], arr[high] = arr[high], arr[mid]
	}
	if s.cmp.Compare(arr[mid], arr[low]) == -1 {
		arr[low], arr[mid] = arr[mid], arr[low]
	}
	return mid // Return index of median-of-three pivot
}

func (s *QuickSelect[T]) part(arr []T, low, high, pivotIndex int) int {
	arr[pivotIndex], arr[high] = arr[high], arr[pivotIndex] // Move pivot to end
	pivot := arr[high]
	i := low

	for j := low; j < high; j++ {
		if s.cmp.Compare(arr[j], pivot) == 1 {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return i
}

func Quick[T prim.Ordered](arr []T, k int) T {
	if len(arr) == 1 {
		return arr[0]
	}
	low, high := 0, len(arr)-1
	for low < high {
		pivotIndex := med3(arr, low, high)
		pivotIndex = part(arr, low, high, pivotIndex)

		if pivotIndex == k {
			return arr[pivotIndex]
		} else if pivotIndex > k {
			high = pivotIndex - 1
		} else {
			low = pivotIndex + 1
		}
	}
	return arr[low]
}

func part[T prim.Ordered](arr []T, low, high, pivotIndex int) int {
	arr[pivotIndex], arr[high] = arr[high], arr[pivotIndex] // Move pivot to end
	pivot := arr[high]
	i := low

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return i
}

// med3 chooses a pivot from three random points to reduce worst-case scenarios
func med3[T prim.Ordered](arr []T, low, high int) int {
	mid := (low + high) / 2
	if arr[low] > arr[mid] {
		arr[low], arr[mid] = arr[mid], arr[low]
	}
	if arr[mid] > arr[high] {
		arr[mid], arr[high] = arr[high], arr[mid]
	}
	if arr[low] > arr[mid] {
		arr[low], arr[mid] = arr[mid], arr[low]
	}
	return mid // Return index of median-of-three pivot
}

// QuickOf finds the k-th smallest element in an unordered list.
func QuickOf[T gtools.SortableOf](arr []T, k int) T {
	if len(arr) == 1 {
		return arr[0]
	}

	low, high := 0, len(arr)-1
	for low < high {
		pivotIndex := med3of(arr, low, high)
		pivotIndex = partof(arr, low, high, pivotIndex)

		if pivotIndex == k {
			return arr[pivotIndex]
		} else if pivotIndex > k {
			high = pivotIndex - 1
		} else {
			low = pivotIndex + 1
		}
	}
	return arr[low]
}

// med3 chooses a pivot from three random points to reduce worst-case scenarios
func med3of[T gtools.SortableOf](arr []T, low, high int) int {
	mid := (low + high) / 2
	if arr[mid].Less(arr[low]) {
		arr[low], arr[mid] = arr[mid], arr[low]
	}
	if arr[high].Less(arr[mid]) {
		arr[mid], arr[high] = arr[high], arr[mid]
	}
	if arr[mid].Less(arr[low]) {
		arr[low], arr[mid] = arr[mid], arr[low]
	}
	return mid // Return index of median-of-three pivot
}

func partof[T gtools.SortableOf](arr []T, low, high, pivotIndex int) int {
	arr[pivotIndex], arr[high] = arr[high], arr[pivotIndex] // Move pivot to end
	pivot := arr[high]
	i := low

	for j := low; j < high; j++ {
		if arr[j].Less(pivot) {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return i
}
