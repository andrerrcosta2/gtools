// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package comparables

import "github.com/andrerrcosta2/gtools/core/gtools/constraints/prim"

type Ordered[T prim.Ordered] struct {
}

func (d Ordered[T]) Compare(a, b T) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

func (d Ordered[T]) Equals(a, b T) bool {
	return a == b
}
