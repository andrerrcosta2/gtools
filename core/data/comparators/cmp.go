// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package comparators

import (
	"github.com/andrerrcosta2/gtools/core/domain/constraints/prim"
)

type TypedKeyOrdered[T prim.Ordered, H prim.Ordered] struct {
	Ordered[T]
}

type KeyOrdered[T prim.Ordered] struct {
	Ordered[T]
}

func (d TypedKeyOrdered[T, H]) Hash(t T) H {
	return Hash[T, H](t)
}

func (d KeyOrdered[T]) Hash(t T) T {
	return t
}

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

type StringOrdered[T prim.Ordered] struct {
	Ordered[T]
}

func (c StringOrdered[T]) Hash(t T) string {
	return StringHash(t)
}
