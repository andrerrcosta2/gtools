// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package dicts

import (
	"github.com/andrerrcosta2/gtools/core/data/str/dict"
	"github.com/andrerrcosta2/gtools/core/domain/function"
)

func PutIfAbsent[K any, V any](m dict.Mutable[K, V], key K, value V) {
	if !m.Contains(key) {
		m.Put(key, value)
	}
}

func GetOrPut[K comparable, V any](m dict.Mutable[K, V], key K, supplier function.Supplier[V]) V {
	if v, ok := m.Get(key); ok {
		return v
	}

	v := supplier()
	m.Put(key, v)
	return v
}
