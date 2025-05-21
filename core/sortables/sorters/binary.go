// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sorters

import (
	"github.com/andrerrcosta2/gtools/core/data/comparators"
)

func Binary[T any, S ~[]T](comparator comparators.Typed[T]) Sorter[T, S] {
	return &binarySorter[T, S]{
		comparator: comparator,
	}
}

// binarySorter (Binary Insertion Sorter) is a variation of insertion sort that uses Binary search to find the
// correct position for inserting an element.
//
// Best case: O(n log(n)) when the list is already sorted or nearly sorted.
// Worst case: O(n^2) in the worst case when the list is in reverse order.
// Average case: O(n^2)) due to the insertion process.
//
// Space Complexity: O(1) due to no additional data structures used.
// Stability: Stable - Preserves the relative order of equal elements.
type binarySorter[T any, S ~[]T] struct {
	comparator comparators.Typed[T]
}

func (b *binarySorter[T, S]) Sort(arr *S) {
	work := *arr
	for i := 1; i < len(work); i++ {
		// Get the current element
		key := work[i]

		// Find the correct position to insert the element using Binary search
		pos := b.obs(work, key, 0, i)

		// Shift elements to the right to make space for the new element
		copy(work[pos+1:i+1], work[pos:i])

		// Insert the element at the correct position
		work[pos] = key
	}
}

func (b *binarySorter[T, S]) obs(arr []T, key T, lo, hi int) int {
	for lo <= hi {
		mid := lo + (hi-lo)/2
		c := b.comparator.Compare(key, arr[mid])
		if c == 0 {
			return mid
		} else if c == 1 {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return lo
}

var _ Sorter[any, []any] = (*binarySorter[any, []any])(nil)
var _ Sorter[string, []string] = (*binarySorter[string, []string])(nil)
