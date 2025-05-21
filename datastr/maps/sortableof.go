// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package maps

import (
	"fmt"
	"github.com/andrerrcosta2/gtools/core/data/comparators"
	"github.com/andrerrcosta2/gtools/core/data/str"
	"github.com/andrerrcosta2/gtools/core/domain/gtools"
	"github.com/andrerrcosta2/gtools/core/sortables"
	"github.com/andrerrcosta2/gtools/core/sortables/sorters"
	"sort"
	"strings"
	"sync"
)

// SortableOf returns a new instance of sortableOfMap.
// It creates a new map with a comparator for the given key type K.
func SortableOf[K gtools.SortableOf, V any](e ...EntrySetOf[K, V]) str.Map[K, V] {
	// Create a new instance of sortableOfMap with an empty map and a comparator.
	sm := &sortableOfMap[K, V]{
		// Initialize the map with a string key type.
		data: make(map[string]str.Entry[K, V]),
		// Create a comparator for the given key type K.
		comparator: sortables.ComparatorOf[K](),
	}

	// Addf each entry to the map.
	for _, es := range e {
		for _, entry := range es.set {
			sm.data[sm.comparator.Hash(entry.Key())] = NewAnyEntry(entry.Key(), entry.Value())
		}
	}

	// Return the sortableOfMap instance.
	return sm
}

type sortableOfMap[K gtools.SortableOf, V any] struct {
	data       map[string]str.Entry[K, V]
	comparator comparators.KeyTyped[K, string]
}

// Put adds a new key-value pair to the map.
// If the key already exists, the old value is replaced.
func (m *sortableOfMap[K, V]) Put(key K, value V) {
	hash := m.comparator.Hash(key)
	m.data[hash] = NewAnyEntry(key, value)
}

func (m *sortableOfMap[K, V]) Get(key K) (V, bool) {
	v, ok := m.data[m.comparator.Hash(key)]
	if !ok {
		var zero V
		return zero, false
	}
	return v.Value(), ok
}

func (m *sortableOfMap[K, V]) Delete(key K) {
	delete(m.data, m.comparator.Hash(key))
}

func (m *sortableOfMap[K, V]) Contains(key K) bool {
	_, ok := m.data[m.comparator.Hash(key)]
	return ok
}

func (m *sortableOfMap[K, V]) Len() int {
	return len(m.data)
}

func (m *sortableOfMap[K, V]) Clear() {
	m.data = make(map[string]str.Entry[K, V])
}

func (m *sortableOfMap[K, V]) Keys() []K {
	out := make([]K, 0, len(m.data))
	for _, entry := range m.data {
		out = append(out, entry.Key())
	}
	return out
}

func (m *sortableOfMap[K, V]) Values() []V {
	out := make([]V, 0, len(m.data))
	for _, entry := range m.data {
		out = append(out, entry.Value())
	}
	return out
}

// Iterator the variadic parameter is just a trick to allow to use the iterator without requiring parameters.
// its presence indicates the keys must be sorted.
func (m *sortableOfMap[K, V]) Iterator(comparator ...comparators.Functional[K]) str.MapIterator[K, V] {
	keys := make([]K, 0, len(m.data))
	for _, v := range m.data {
		keys = append(keys, v.Key())
	}

	if comparator != nil && len(comparator) > 0 {
		sorter := sorters.Quick[K, []K](comparator[0])
		sorter.Sort(&keys)
	}

	return &SortableOfMapIterator[K, V]{
		m:       m,
		keys:    keys,
		current: 0,
	}
}

var _ StructMap[gtools.SortableOf, string] = (*sortableOfMap[gtools.SortableOf, string])(nil)

type SortableOfMapIterator[K gtools.SortableOf, V any] struct {
	m       *sortableOfMap[K, V]
	keys    []K
	current int
}

func (it *SortableOfMapIterator[K, V]) Next() (key K, value V, ok bool) {
	if it.current >= len(it.keys) {
		return
	}

	mapKey := it.keys[it.current]
	entry, ok := it.m.data[it.m.comparator.Hash(mapKey)]

	if !ok {
		panic(fmt.Sprintf("this iterator or its map is corrupt, key '%v' was expected but"+
			"was not found on available next position", mapKey))
	}

	key = entry.Key()
	value = entry.Value()
	ok = true
	it.current++

	return
}

func (m *sortableOfMap[K, V]) String() string {
	var keys []string
	for key := range m.data {
		keys = append(keys, key)
	}

	// Sorter keys to maintain a consistent order
	sort.Strings(keys)

	var sb strings.Builder
	for i, key := range keys {
		entry := m.data[key]
		sb.WriteString(fmt.Sprintf("%d: %s\n", i, entry.String()))
	}
	return sb.String()
}

func ConcSortableOf[K gtools.SortableOf, V any](e ...EntrySetOf[K, V]) str.Map[K, V] {
	sm := &concSortableOf[K, V]{
		data:       make(map[string]str.Entry[K, V]),
		comparator: sortables.ComparatorOf[K](),
	}

	for _, es := range e {
		for _, entry := range es.set {
			sm.data[sm.comparator.Hash(entry.Key())] = NewAnyEntry(entry.Key(), entry.Value())
		}
	}
	return sm
}

type concSortableOf[K gtools.SortableOf, V any] struct {
	mtx        sync.RWMutex
	data       map[string]str.Entry[K, V]
	comparator comparators.KeyTyped[K, string]
}

func (m *concSortableOf[K, V]) Put(key K, value V) {
	hash := m.comparator.Hash(key)
	m.mtx.Lock()
	m.data[hash] = NewAnyEntry(key, value)
	m.mtx.Unlock()
}

func (m *concSortableOf[K, V]) Get(key K) (value V, exists bool) {
	hash := m.comparator.Hash(key)
	m.mtx.RLock()
	v, ok := m.data[hash]
	m.mtx.RUnlock()
	if !ok {
		return value, false
	}
	return v.Value(), ok
}

func (m *concSortableOf[K, V]) Delete(key K) {
	hash := m.comparator.Hash(key)
	m.mtx.Lock()
	delete(m.data, hash)
	m.mtx.Unlock()
}

func (m *concSortableOf[K, V]) Contains(key K) bool {
	hash := m.comparator.Hash(key)
	m.mtx.RLock()
	_, ok := m.data[hash]
	m.mtx.RUnlock()
	return ok
}

func (m *concSortableOf[K, V]) Len() int {
	m.mtx.RLock()
	defer m.mtx.RUnlock()
	return len(m.data)
}

func (m *concSortableOf[K, V]) Clear() {
	m.mtx.Lock()
	m.data = make(map[string]str.Entry[K, V])
	m.mtx.Unlock()
}

func (m *concSortableOf[K, V]) Keys() []K {
	m.mtx.RLock()
	defer m.mtx.RUnlock()
	out := make([]K, 0, len(m.data))
	for _, entry := range m.data {
		out = append(out, entry.Key())
	}
	return out
}

func (m *concSortableOf[K, V]) Values() []V {
	m.mtx.RLock()
	defer m.mtx.RUnlock()
	out := make([]V, 0, len(m.data))
	for _, entry := range m.data {
		out = append(out, entry.Value())
	}
	return out
}

func (m *concSortableOf[K, V]) String() string {
	m.mtx.RLock()
	defer m.mtx.RUnlock()
	var sb strings.Builder
	for _, entry := range m.data {
		sb.WriteString(fmt.Sprintf("%s\n", entry.String()))
	}
	return sb.String()
}

func (m *concSortableOf[K, V]) Iterator(comparator ...comparators.Functional[K]) str.MapIterator[K, V] {
	keys := make([]K, 0, len(m.data))
	for _, v := range m.data {
		keys = append(keys, v.Key())
	}

	if comparator != nil && len(comparator) > 0 {
		sorter := sorters.Quick[K, []K](comparator[0])
		sorter.Sort(&keys)
	}

	return &ConcSortableOfMapIterator[K, V]{
		m:       m,
		keys:    keys,
		current: 0,
	}
}

type ConcSortableOfMapIterator[K gtools.SortableOf, V any] struct {
	m       *concSortableOf[K, V]
	keys    []K
	current int
}

func (it *ConcSortableOfMapIterator[K, V]) Next() (key K, value V, ok bool) {
	it.m.mtx.RLock()
	defer it.m.mtx.RUnlock()

	if it.current >= len(it.keys) {
		return
	}

	mapKey := it.keys[it.current]
	entry, ok := it.m.data[it.m.comparator.Hash(mapKey)]

	if !ok {
		panic(fmt.Sprintf("this iterator or its map is corrupt, key '%v' was expected but"+
			"was not found on available next position", mapKey))
	}

	key = entry.Key()
	value = entry.Value()
	ok = true
	it.current++

	return
}
