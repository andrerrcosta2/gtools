// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package dicts

import (
	"sync"

	"github.com/andrerrcosta2/gtools/core/data/comparator"
	"github.com/andrerrcosta2/gtools/core/data/str/dict"
	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"github.com/andrerrcosta2/gtools/core/format/fmx"
	"github.com/andrerrcosta2/gtools/core/format/sprints"
	"github.com/andrerrcosta2/gtools/core/sortables/sorters"
	"github.com/andrerrcosta2/gtools/datastr/internal/hash"
)

// Comparable returns a new instance of a comparable map.
// It creates a new map with a comparator for the given key type K.
func Comparable[K comparable, V any](entries ...dict.Entry[K, V]) dict.Mutable[K, V] {
	m := &cmpMap[K, V]{
		table: make(hash.FlatStore[K, V]),
	}
	for _, entry := range entries {
		m.table[entry.Key()] = entry.Value()
	}
	return m
}

type cmpMap[K comparable, V any] struct {
	table hash.FlatStore[K, V]
}

func (m *cmpMap[K, V]) Put(key K, value V) {
	m.table.Store(key, value)
}

func (m *cmpMap[K, V]) Get(key K) (V, bool) {
	return m.table.Get(key)
}

func (m *cmpMap[K, V]) Delete(key K) {
	m.table.TryRemove(key)
}

func (m *cmpMap[K, V]) Contains(key K) bool {
	return m.table.HasHash(key)
}

func (m *cmpMap[K, V]) Len() int {
	return len(m.table)
}

func (m *cmpMap[K, V]) Clear() {
	m.table.Clear()
}

func (m *cmpMap[K, V]) Keys() []K {
	return m.table.Keys()
}

func (m *cmpMap[K, V]) Values() []V {
	return m.table.Values()
}

func (m *cmpMap[K, V]) Iterator(_ ...comparator.EqComparer[K]) dict.MapIterator[K, V] {
	snap := make([]dict.Entry[K, V], 0, m.Len())
	for key, value := range m.table {
		snap = append(snap, NewComparableEntry(key, value))
	}

	return &cmpMapIter[K, V]{
		m:       m,
		entries: snap,
		current: 0,
	}
}

func (m *cmpMap[K, V]) String() string {
	return sprints.Map[map[K]V](indent.Zero(), m.table)
}

type cmpMapIter[K comparable, V any] struct {
	m       *cmpMap[K, V]
	entries []dict.Entry[K, V]
	current int
}

func (it *cmpMapIter[K, V]) Next() (key K, value V, ok bool) {
	if it.current >= len(it.entries) {
		return
	}
	key, value, ok = it.entries[it.current].Key(), it.entries[it.current].Value(), true
	it.current++
	return
}

// ConcComparable returns a new instance of concCmp.
func ConcComparable[K comparable, V any](entries ...dict.Entry[K, V]) dict.Mutable[K, V] {
	m := &concCmp[K, V]{
		mtx:   sync.RWMutex{},
		table: make(hash.FlatStore[K, V]),
	}
	for _, entry := range entries {
		m.table[entry.Key()] = entry.Value()
	}
	return m
}

type concCmp[K comparable, V any] struct {
	mtx   sync.RWMutex
	table hash.FlatStore[K, V]
}

func (m *concCmp[K, V]) Put(key K, value V) {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	m.table.TryStore(key, value)
}

func (m *concCmp[K, V]) Get(key K) (V, bool) {
	m.mtx.RLock()
	defer m.mtx.RUnlock()
	return m.table.Get(key)
}

func (m *concCmp[K, V]) Delete(key K) {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	m.table.TryRemove(key)
}

func (m *concCmp[K, V]) Contains(key K) bool {
	m.mtx.RLock()
	defer m.mtx.RUnlock()
	return m.table.HasHash(key)
}

func (m *concCmp[K, V]) Len() int {
	m.mtx.RLock()
	defer m.mtx.RUnlock()
	return m.table.Len()
}

func (m *concCmp[K, V]) Clear() {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	m.table.Clear()
}

func (m *concCmp[K, V]) Keys() []K {
	m.mtx.RLock()
	defer m.mtx.RUnlock()
	return m.table.Keys()
}

func (m *concCmp[K, V]) Values() []V {
	m.mtx.RLock()
	defer m.mtx.RUnlock()
	return m.table.Values()
}

func (m *concCmp[K, V]) Iterator(comparator ...comparator.EqComparer[K]) dict.MapIterator[K, V] {
	m.mtx.RLock()
	defer m.mtx.RUnlock()

	keys := m.table.Keys()

	if len(comparator) > 0 {
		sorter := sorters.Quick[[]K](comparator[0])
		sorter.Sort(&keys)
	}

	return &concCmpMapIter[K, V]{
		m:       m,
		keys:    keys,
		current: 0,
	}
}

func (m *concCmp[K, V]) String() string {
	m.mtx.RLock()
	defer m.mtx.RUnlock()
	return m.sprint(indent.Zero())
}

func (m *concCmp[K, V]) Sprint(tab indent.Tab) string {
	m.mtx.RLock()
	defer m.mtx.RUnlock()
	return m.sprint(tab)
}

func (m *concCmp[K, V]) sprint(tab indent.Tab) string {
	return sprints.Map(tab, m.table)
}

type concCmpMapIter[K comparable, V any] struct {
	m       *concCmp[K, V]
	keys    []K
	current int
}

func (it *concCmpMapIter[K, V]) Next() (key K, value V, ok bool) {
	it.m.mtx.RLock()
	defer it.m.mtx.RUnlock()

	if it.current >= len(it.keys) {
		return
	}

	key = it.keys[it.current]
	value, ok = it.m.table[key]

	if !ok {
		panic(fmx.Sprintf("this iterator or its map is corrupt, key '%v' was expected but"+
			"was not found on available next position", key))
	}

	it.current++
	return
}
