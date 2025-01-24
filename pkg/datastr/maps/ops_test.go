// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package maps

import (
	"github.com/andrerrcosta2/gtools/core/data/str"
	"github.com/andrerrcosta2/gtools/core/gtools/functions"
	"github.com/andrerrcosta2/gtools/gtests"
)

func shouldFindKey[K comparable, V any](t gtests.Loggable, m str.Map[K, V], key K) {
	t.Helper()
	if _, ok := m.Get(key); !ok {
		t.Errorf("Expected a value, nothing found")
	}
}

func shouldFindExactKey[K comparable, V any](t gtests.Loggable, m str.Map[K, V], key K, value V, comp functions.BiPredicate[V, V]) {
	t.Helper()
	if got, ok := m.Get(key); !ok || !comp(got, value) {
		t.Errorf("Expected %v, got %v", value, got)
	}
}

func shouldNotFindKey[K comparable, V any](t gtests.Loggable, m str.Map[K, V], key K) {
	t.Helper()
	if _, ok := m.Get(key); ok {
		t.Errorf("Expected key %v to not exist", key)
	}
}

func shouldDeleteKey[K comparable, V any](t gtests.Loggable, m str.Map[K, V], key K) {
	t.Helper()
	m.Delete(key)
	if _, ok := m.Get(key); ok {
		t.Errorf("Expected key %v to be deleted", key)
	}
}

func shouldHaveLength[K comparable, V any](t gtests.Loggable, m str.Map[K, V], length int) {
	t.Helper()
	if m.Len() != length {
		t.Errorf("Expected length %d, got %d", length, m.Len())
	}
}

func shouldFindAllUsingIteratorBy[K comparable, V any](t gtests.Loggable, it str.MapIterator[K, V], expected map[K]V, compare functions.BiPredicate[V, V]) {
	t.Helper()
	var counter int
	for k, v, hn := it.Next(); hn; k, v, hn = it.Next() {
		counter++
		exp, ok := expected[k]
		if !ok {
			t.Errorf("key '%v' existed but not expected", k)
			continue
		}

		if !compare(v, exp) {
			t.Errorf("Expected key '%s' to have value '%v, got %v", k, exp, v)
		}
	}

	if counter != len(expected) {
		t.Errorf("Expected %d keys, got %d", len(expected), counter)
	}
}
