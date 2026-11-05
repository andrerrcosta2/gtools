// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sets

import (
	"maps"

	"github.com/andrerrcosta2/gtools/core/data/comparator"
	"github.com/andrerrcosta2/gtools/core/data/str/set"
	"github.com/andrerrcosta2/gtools/datastr/internal/hash"
)

func eqComparable[T comparable](m map[T]struct{}, other set.Readable[T]) bool {
	if len(m) != other.Len() {
		return false
	}

	if internal, ok := other.(mapProvider[T, struct{}]); ok {
		return maps.Equal(m, internal.backingMap())
	}

	out := true
	other.Range(func(value T) bool {
		if _, ok := m[value]; !ok {
			out = false
			return false
		}
		return true
	})

	return out
}

func eqHashable[H comparable, V any](
	m hash.Store[H, V],
	other set.Readable[V],
	cmp comparator.EqKeyer[H, V]) bool {

	if m.Len() != other.Len() {
		return false
	}

	if internal, ok := other.(hashable[H, V]); ok {
		return internal.table().Equals(m, cmp.Equals)
	}
	out := true
	other.Range(func(value V) bool {
		if !m.HasHash(cmp.Key(value)) {
			out = false
			return false
		}
		return true
	})
	return out
}

// eqHashableOrdered makes no sense in deny equality by order.
// order in sets is a typed property which would lead us to
// require the same type to have the same order with the same elements...
func eqHashableOrdered[H comparable, V any](
	m hash.Store[H, struct{}],
	other set.Readable[V],
	cmp comparator.EqKeyer[H, V]) bool {

	if m.Len() != other.Len() {
		return false
	}
	out := true
	other.Range(func(value V) bool {
		if !m.HasHash(cmp.Key(value)) {
			out = false
			return false
		}
		return true
	})
	return out
}
