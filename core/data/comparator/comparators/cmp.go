// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package comparators

import (
	"github.com/andrerrcosta2/gtools/core/data/comparator"
	"github.com/andrerrcosta2/gtools/core/data/keyer"
	"github.com/andrerrcosta2/gtools/core/data/keyer/keyers"
	"github.com/andrerrcosta2/gtools/core/domain/constraints"
	"github.com/andrerrcosta2/gtools/core/domain/constraints/prim"
)

func TypedOrderedKeyer[T prim.Ordered, H prim.Ordered]() comparator.EqCmpKeyer[T, H] {
	return typedOrderedKeyer[T, H]{}
}

type typedOrderedKeyer[T prim.Ordered, H prim.Ordered] struct {
	ordered[T]
}

func (d typedOrderedKeyer[T, H]) Key(t T) H {
	return keyers.ComparableTyped[T, H](t)
}

func IdEqCmpKeyer[T prim.Ordered]() comparator.EqCmpKeyer[T, T] {
	return idEqCmpKeyer[T]{}
}

type idEqCmpKeyer[T prim.Ordered] struct {
	ordered[T]
}

func (d idEqCmpKeyer[T]) Key(t T) T {
	return t
}

func OrderedOf[T constraints.Ordered[T]]() comparator.Comparer[T] {
	return &orderedOf[T]{}
}

type orderedOf[T constraints.Ordered[T]] struct{}

func (d orderedOf[T]) Compare(a, b T) int {
	if a.Less(b) {
		return -1
	} else if b.Less(a) {
		return 1
	}
	return 0
}

func Ordered[T prim.Ordered]() comparator.EqComparer[T] {
	return ordered[T]{}
}

type ordered[T prim.Ordered] struct {
}

func (d ordered[T]) Compare(a, b T) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

func (d ordered[T]) Equals(a, b T) bool {
	return a == b
}

func StringOrdered[T prim.Ordered]() comparator.EqCmpKeyer[T, string] {
	return stringOrdered[T]{}
}

type stringOrdered[T prim.Ordered] struct {
	ordered[T]
}

func (c stringOrdered[T]) Key(t T) string {
	return keyers.CmpToString(t)
}

type comparableOf[N constraints.Comparable[N]] struct{}

func ComparableOf[N constraints.Comparable[N]]() comparator.Equals[N] {
	return comparableOf[N]{}
}

func (c comparableOf[N]) Equals(a, b N) bool {
	return a.Equal(b)
}

func KeySortableOf[O constraints.Sortable[O], H comparable](k keyer.Comparable[O, H]) comparator.EqCmpKeyer[O, H] {
	return keySortableOf[O, H]{
		key: k,
	}
}

func SortableOf[O constraints.Sortable[O]]() comparator.EqComparer[O] {
	return keySortableOf[O, any]{}
}

type keySortableOf[O constraints.Sortable[O], H comparable] struct {
	key keyer.Comparable[O, H]
}

func (c keySortableOf[O, H]) Compare(a, b O) int {
	if a.Less(b) {
		return -1
	}
	if b.Less(a) {
		return 1
	}
	return 0
}

func (c keySortableOf[O, H]) Equals(a, b O) bool {
	return a.Equal(b)
}

func (c keySortableOf[O, H]) Key(t O) H {
	return c.key.Key(t)
}

func Entities[T constraints.Entity[T]]() comparator.EqCmpKeyer[T, string] {
	return entity[T]{}
}

type entity[T constraints.Entity[T]] struct{}

func (c entity[T]) Equals(a, b T) bool {
	return a.Equal(b)
}

func (c entity[T]) Key(t T) string {
	return t.ID()
}

func (c entity[T]) Compare(a, b T) int {
	if a.Less(b) {
		return -1
	}
	if b.Less(a) {
		return 1
	}
	return 0
}

func Identity[T comparable]() comparator.EqKeyer[T, T] {
	return identity[T]{}
}

type identity[T comparable] struct{}

func (k identity[T]) Equals(a, b T) bool {
	return a == b
}

func (k identity[T]) Key(t T) T {
	return t
}
