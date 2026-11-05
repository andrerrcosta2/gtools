// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sorters

import (
	"github.com/andrerrcosta2/gtools/core/data/comparator/comparators"
	"github.com/andrerrcosta2/gtools/core/domain/constraints/prim"
)

// Get returns a new instance of the Sorter interface using the default comparator.
//
// The implementation will depend on the type parameter T.
// If T is a mergeSorter, Quick or binarySorter, the respective Sorter instance will be returned.
// Otherwise, the Native Sorter will be returned.
//
// Note: The default comparator is an instance of comparables.ordered[S]{},
// which is the default comparator for the given type S.
func Get[T Sorter[S, E], S ~[]E, E prim.Ordered]() Sorter[S, E] {
	var sort T
	switch any(sort).(type) {
	case mergeSorter[S, E]:
		return Merge[S](comparators.Ordered[E]())
	case quickSorter[S, E]:
		return Quick[S](comparators.Ordered[E]())
	case binarySorter[S, E]:
		return Binary[S](comparators.Ordered[E]())
	default:
		return Native[S](comparators.Ordered[E]())
	}
}
