// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sets

import (
	"github.com/andrerrcosta2/gtools/core/data/comparators"
	"github.com/andrerrcosta2/gtools/core/domain/constraints/prim"
	"github.com/andrerrcosta2/gtools/core/domain/functions"
	"github.com/andrerrcosta2/gtools/core/sortables/sorters"
)

type hashSet[T any, H prim.Hashable] interface {
	hash(t T) H
}

type insertionSet[T any, H prim.Hashable] interface {
	hashSet[T, H]
	contains(t T) bool
	size() int
	setHash(h H, idx int)
	append(t T)
}

func Filter[T any, H prim.Hashable](cmp comparators.KeyTyped[T, H], e ...T) []T {
	values := make(map[H]bool)
	output := make([]T, 0, len(e))

	for _, v := range e {
		hash := cmp.Hash(v)
		if _, ok := values[hash]; !ok {
			values[hash] = true
			output = append(output, v)
		}
	}
	return output
}

func CmpFilter[T comparable](e ...T) []T {
	values := make(map[T]bool)
	output := make([]T, 0, len(e))
	for _, v := range e {
		if _, ok := values[v]; !ok {
			values[v] = true
			output = append(output, v)
		}
	}
	return output
}

func OrderedFilter[T prim.Ordered](e []T, fn functions.Consumer[T]) {
	sorters.Quick[T, []T](comparators.Ordered[T]{}).Sort(&e)
	var prev T
	first := true
	for _, value := range e {
		// just skip duplicates
		if first || value != prev {
			fn(value)
			prev = value
			first = false
		}
	}
}
