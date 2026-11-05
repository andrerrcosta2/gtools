// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sorters

import (
	"slices"

	"github.com/andrerrcosta2/gtools/core/data/comparator"
)

// Native returns a Sorter that uses the native sort function from the standard
// library.
func Native[S ~[]E, E any](comparator comparator.EqComparer[E]) Sorter[S, E] {
	// Return a new instance of pdqSort with the given comparator.
	return &pdqSort[S, E]{
		comparator: comparator,
	}
}

type pdqSort[S ~[]E, E any] struct {
	comparator comparator.EqComparer[E]
}

// Sort sorts the given slice in ascending order using the provided comparator.
func (s *pdqSort[S, E]) Sort(slice *S) {
	// Sorter the slice using the given comparator.
	slices.SortFunc(*slice, s.comparator.Compare)
}

var _ Sorter[[]int, int] = (*pdqSort[[]int, int])(nil)
