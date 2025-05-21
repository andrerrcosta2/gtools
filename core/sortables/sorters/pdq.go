// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sorters

import (
	"github.com/andrerrcosta2/gtools/core/data/comparators"
	"slices"
)

// Native returns a Sorter that uses the native sort function from the standard
// library.
func Native[T any, S ~[]T](comparator comparators.Typed[T]) Sorter[T, S] {
	// Return a new instance of pdqSort with the given comparator.
	return &pdqSort[T, S]{
		comparator: comparator,
	}
}

type pdqSort[T any, S ~[]T] struct {
	comparator comparators.Typed[T]
}

// Sort sorts the given slice in ascending order using the provided comparator.
func (s *pdqSort[T, S]) Sort(slice *S) {
	// Sorter the slice using the given comparator.
	slices.SortFunc(*slice, func(a, b T) int {
		// Compare the two elements and return an integer indicating their order.
		return s.comparator.Compare(a, b)
	})
}

var _ Sorter[int, []int] = (*pdqSort[int, []int])(nil)
