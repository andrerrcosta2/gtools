// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package arrays

import (
	"github.com/andrerrcosta2/gtools/pkg/constraints"
	"github.com/andrerrcosta2/gtools/pkg/functions"
	"sort"
)

// Sorted sorts a slice of ordered elements in ascending order.
// It uses the sort.Slice function from the standard library to perform the sorting.
// The function returns the sorted slice.
func Sorted[T constraints.Ordered](arr []T) SortedSlice[T] {
	// Sort the slice using the provided less function
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	// Return the sorted slice
	return arr
}

func SortedBy[T constraints.Ordered, K constraints.Ordered](arr []T, f functions.Function[T, K]) SortedSlice[T] {
	sort.Slice(arr, func(i, j int) bool {
		return f(arr[i]) < f(arr[j])
	})
	return arr
}

type SortedSlice[T constraints.Ordered] []T

func (a SortedSlice[T]) Len() int           { return len(a) }
func (a SortedSlice[T]) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortedSlice[T]) Less(i, j int) bool { return a[i] < a[j] }
