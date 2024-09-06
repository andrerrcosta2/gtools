// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sorts

import (
	"github.com/andrerrcosta2/gtools/comparables"
	"github.com/andrerrcosta2/gtools/core/constraints/prim"
)

func Get[T Sort[S], S prim.Ordered]() Sort[S] {
	var sort T
	switch any(sort).(type) {
	case MergeSort[S]:
		return NewMergeSort[S](comparables.Ordered[S]{})
	case Quicksort[S]:
		return NewQuicksort[S](comparables.Ordered[S]{})
	case BinarySort[S]:
		return NewBinarySort[S](comparables.Ordered[S]{})
	default:
		return nil
	}
}
