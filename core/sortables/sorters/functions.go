// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sorters

import (
	"github.com/andrerrcosta2/gtools/core/data/comparators"
	"github.com/andrerrcosta2/gtools/core/domain/constraints/prim"
)

// Get returns a new instance of the Sorter interface using the default comparator.
//
// The implementation will depend on the type parameter T.
// If T is a mergeSorter, Quick or binarySorter, the respective Sorter instance will be returned.
// Otherwise, the Native Sorter will be returned.
//
// Note: The default comparator is an instance of comparables.Ordered[S]{},
// which is the default comparator for the given type S.
func Get[T Sorter[E, S], E prim.Ordered, S ~[]E]() Sorter[E, S] {
	var sort T
	switch any(sort).(type) {
	case mergeSorter[E, S]:
		return Merge[E, S](comparators.Ordered[E]{})
	case quickSorter[E, S]:
		return Quick[E, S](comparators.Ordered[E]{})
	case binarySorter[E, S]:
		return Binary[E, S](comparators.Ordered[E]{})
	default:
		return Native[E, S](comparators.Ordered[E]{})
	}
}
