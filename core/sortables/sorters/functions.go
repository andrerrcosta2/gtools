// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sorters

import (
	"github.com/andrerrcosta2/gtools/core/data/comparables"
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
func Get[T Sorter[S], S prim.Ordered]() Sorter[S] {
	var sort T
	switch any(sort).(type) {
	case mergeSorter[S]:
		return Merge[S](comparables.Ordered[S]{})
	case quickSorter[S]:
		return Quick[S](comparables.Ordered[S]{})
	case binarySorter[S]:
		return Binary[S](comparables.Ordered[S]{})
	default:
		return Native[S](comparables.Ordered[S]{})
	}
}
