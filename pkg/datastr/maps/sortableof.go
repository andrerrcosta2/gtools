// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package maps

import (
	"fmt"
	"github.com/andrerrcosta2/gtools/pkg/comparables"
	"github.com/andrerrcosta2/gtools/pkg/datastr/iterables"
	"github.com/andrerrcosta2/gtools/pkg/gtools"
	"github.com/andrerrcosta2/gtools/pkg/sortables"
	"github.com/andrerrcosta2/gtools/pkg/sorts"
	"sort"
	"strings"
)

// SortableOf returns a new instance of SortableOfMap.
// It creates a new map with a comparator for the given key type K.
func SortableOf[K gtools.SortableOf, V any]() *SortableOfMap[K, V] {
	// Create a new instance of SortableOfMap with an empty map and a comparator.
	return &SortableOfMap[K, V]{
		// Initialize the map with a string key type.
		data: make(map[string]Entry[K, V]),
		// Create a comparator for the given key type K.
		comparator: sortables.ComparatorOf[K](),
	}
}

type SortableOfMap[K gtools.SortableOf, V any] struct {
	data       map[string]Entry[K, V]
	comparator comparables.KeyComparator[K, string]
}

// Put adds a new key-value pair to the map.
// If the key already exists, the old value is replaced.
func (m *SortableOfMap[K, V]) Put(key K, value V) {
	hash := m.comparator.Hash(key)
	m.data[hash] = NewAnyEntry(key, value)
}

func (m *SortableOfMap[K, V]) Get(key K) (V, bool) {
	v, ok := m.data[m.comparator.Hash(key)]
	if !ok {
		var zero V
		return zero, false
	}
	return v.Value(), ok
}

func (m *SortableOfMap[K, V]) Delete(key K) {
	delete(m.data, m.comparator.Hash(key))
}

func (m *SortableOfMap[K, V]) Contains(key K) bool {
	_, ok := m.data[m.comparator.Hash(key)]
	return ok
}

func (m *SortableOfMap[K, V]) Len() int {
	return len(m.data)
}

func (m *SortableOfMap[K, V]) Clear() {
	m.data = make(map[string]Entry[K, V])
}

func (m *SortableOfMap[K, V]) Keys() []K {
	out := make([]K, 0, len(m.data))
	for _, entry := range m.data {
		out = append(out, entry.Key())
	}
	return out
}

func (m *SortableOfMap[K, V]) Values() []V {
	out := make([]V, 0, len(m.data))
	for _, entry := range m.data {
		out = append(out, entry.Value())
	}
	return out
}

// Iterator the variadic parameter is just a trick to allow to use the iterator without requiring parameters.
// its presence indicates the keys must be sorted.
func (m *SortableOfMap[K, V]) Iterator(comparator ...comparables.FunctionalComparator[K]) iterables.MapIterator[K, V] {
	keys := make([]K, 0, len(m.data))
	for _, v := range m.data {
		keys = append(keys, v.Key())
	}

	if comparator != nil && len(comparator) > 0 {
		sortables.Sort[K](&keys, sorts.NewQuicksort[K](comparator[0]))
	}

	return &SortableOfMapIterator[K, V]{
		m:       m,
		keys:    keys,
		current: 0,
	}
}

var _ StructMap[gtools.SortableOf, string] = (*SortableOfMap[gtools.SortableOf, string])(nil)

type SortableOfMapIterator[K gtools.SortableOf, V any] struct {
	m       *SortableOfMap[K, V]
	keys    []K
	current int
}

func (it *SortableOfMapIterator[K, V]) Next() (key K, value V, ok bool) {
	if it.current >= len(it.keys) {
		return
	}

	mapKey := it.keys[it.current]
	entry := it.m.data[it.m.comparator.Hash(mapKey)]

	key = entry.Key()
	value = entry.Value()
	ok = true
	it.current++

	return
}

func (m *SortableOfMap[K, V]) String() string {
	var keys []string
	for key := range m.data {
		keys = append(keys, key)
	}

	// Sort keys to maintain a consistent order
	sort.Strings(keys)

	var sb strings.Builder
	for i, key := range keys {
		entry := m.data[key]
		sb.WriteString(fmt.Sprintf("%d: %s\n", i, entry.String()))
	}
	return sb.String()
}
