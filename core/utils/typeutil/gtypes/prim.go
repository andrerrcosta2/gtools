// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package gtypes

import (
	"github.com/andrerrcosta2/gtools/core/domain/gtools"
)

func ComparableOf[T comparable](of T) gtools.SortableOf {
	return compOf[T]{
		of: of,
	}
}

type compOf[T comparable] struct {
	of T
}

func (c compOf[T]) Equal(o interface{}) bool {
	if other, ok := o.(T); ok {
		return c.of == other
	}
	if other, ok := o.(compOf[T]); ok {
		return c.of == other.of
	}
	return false
}

func (c compOf[T]) Less(o interface{}) bool {
	return false
}
