// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sorters

import (
	"github.com/andrerrcosta2/gtools/core/data/comparables"
	"slices"
)

// Native returns a Sorter that uses the native sort function from the standard
// library.
func Native[T any](comparator comparables.Comparator[T]) Sorter[T] {
	// Return a new instance of pdqSort with the given comparator.
	return &pdqSort[T]{
		comparator: comparator,
	}
}

type pdqSort[T any] struct {
	comparator comparables.Comparator[T]
}

// Sort sorts the given slice in ascending order using the provided comparator.
func (s *pdqSort[T]) Sort(slice *[]T) {
	// Sorter the slice using the given comparator.
	slices.SortFunc(*slice, func(a, b T) int {
		// Compare the two elements and return an integer indicating their order.
		return s.comparator.Compare(a, b)
	})
}

// SortP sorts the given slice of pointers in ascending order using the provided comparator.
func (s *pdqSort[T]) SortP(arr *[]*T) {
	// Sorter the slice using the given comparator.
	slices.SortFunc(*arr, func(a, b *T) int {
		// Compare the two elements and return an integer indicating their order.
		return s.comparator.Compare(*a, *b)
	})
}

var _ Sorter[int] = (*pdqSort[int])(nil)
