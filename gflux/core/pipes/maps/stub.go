// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package maps

import (
	"fmt"

	"github.com/andrerrcosta2/gtools/core/data/str/dict"
)

func newTestEntry[K, V any](key K, value V) dict.Entry[K, V] {
	return &testEntry[K, V]{
		key:   key,
		value: value,
	}
}

type testEntry[K, V any] struct {
	key   K
	value V
}

func (t *testEntry[K, V]) Key() K {
	return t.key
}

func (t *testEntry[K, V]) Value() V {
	return t.value
}

func (t *testEntry[K, V]) String() string {
	return fmt.Sprintf("{key: %v, value: %v}", t.key, t.value)
}
