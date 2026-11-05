// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sets

import (
	"github.com/andrerrcosta2/gtools/core/data/comparator"
	"github.com/andrerrcosta2/gtools/core/data/str/set"
	"github.com/andrerrcosta2/gtools/datastr/internal/hash"
)

func HashSet[T any, H comparable](cmp comparator.EqCmpKeyer[T, H], values ...T) set.Mutable[T] {
	hs := &hashset[T, H]{
		table: make(hash.FlatStore[H, T]),
		cmp:   cmp,
	}
	for _, value := range values {
		hs.table.TryStore(cmp.Key(value), value)
	}
	return hs
}

type hashset[T any, H comparable] struct {
	table hash.FlatStore[H, T]
	cmp   comparator.EqCmpKeyer[T, H]
}

func (h hashset[T, H]) IsEmpty() bool {
	return h.table.IsEmpty()
}

func (h hashset[T, H]) Equals(other set.Readable[T]) bool {
	return eqHashable(h.table, other, h.cmp)
}

func (h hashset[T, H]) Has(value T) bool {
	return h.table.HasHash(h.hash(value))
}

func (h hashset[T, H]) hash(value T) H {
	return h.cmp.Key(value)
}

func (h hashset[T, H]) Len() int {
	return h.table.Len()
}

func (h hashset[T, H]) Range(callback func(value T) bool) {
	for _, value := range h.table {
		if !callback(value) {
			return
		}
	}
}

func (h hashset[T, H]) Elements() []T {
	return h.table.Values()
}

func (h hashset[T, H]) Add(value T) {
	h.table.TryStore(h.hash(value), value)
}

func (h hashset[T, H]) Clear() {
	h.table.Clear()
}

func (h hashset[T, H]) Remove(value T) {
	h.table.TryRemove(h.hash(value))
}
