// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package maps

import (
	"github.com/andrerrcosta2/gtools/core/data/comparators"
	"github.com/andrerrcosta2/gtools/core/data/str"
	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"github.com/andrerrcosta2/gtools/core/format/fmx"
	"github.com/andrerrcosta2/gtools/core/format/sprints"
	"github.com/andrerrcosta2/gtools/core/sortables/sorters"
	"sync"
)

// Comparable returns a new instance of a comparable map.
// It creates a new map with a comparator for the given key type K.
func Comparable[K comparable, V any](e ...EntrySet[K, V]) str.Map[K, V] {
	m := &cmp[K, V]{
		data: make(map[K]V),
	}
	for _, es := range e {
		for _, entry := range es.entries {
			m.data[entry.Key()] = entry.Value()
		}
	}
	return m
}

type cmp[K comparable, V any] struct {
	data map[K]V
}

func (m *cmp[K, V]) Put(key K, value V) {
	m.data[key] = value
}

func (m *cmp[K, V]) Get(key K) (V, bool) {
	v, ok := m.data[key]
	return v, ok
}

func (m *cmp[K, V]) Delete(key K) {
	delete(m.data, key)
}

func (m *cmp[K, V]) Contains(key K) bool {
	_, ok := m.data[key]
	return ok
}

func (m *cmp[K, V]) Len() int {
	return len(m.data)
}

func (m *cmp[K, V]) Clear() {
	m.data = make(map[K]V)
}

func (m *cmp[K, V]) Keys() []K {
	keys := make([]K, 0, m.Len())
	for key := range m.data {
		keys[len(keys)] = key
	}
	return keys
}

func (m *cmp[K, V]) Values() []V {
	values := make([]V, 0, m.Len())
	for _, value := range m.data {
		values[len(values)] = value
	}
	return values
}

func (m *cmp[K, V]) Iterator(comparator ...comparators.Functional[K]) str.MapIterator[K, V] {
	keys := m.Keys()

	if comparator != nil && len(comparator) > 0 {
		sorter := sorters.Quick[K, []K](comparator[0])
		sorter.Sort(&keys)
	}

	return &cmpMapIter[K, V]{
		m:       m,
		keys:    keys,
		current: 0,
	}
}

func (m *cmp[K, V]) String() string {
	return sprints.Map[map[K]V](indent.Zero(), m.data)
}

type cmpMapIter[K comparable, V any] struct {
	m       *cmp[K, V]
	keys    []K
	current int
}

func (it *cmpMapIter[K, V]) Next() (key K, value V, ok bool) {
	if it.current >= len(it.keys) {
		return
	}

	key = it.keys[it.current]
	value, ok = it.m.data[key]

	if !ok {
		panic(fmx.Sprintf("this iterator or its map is corrupt, key '%v' was expected but"+
			"was not found on available next position", key))
	}

	it.current++
	return
}

// ConcComparable returns a new instance of concCmp.
func ConcComparable[K comparable, V any](e ...EntrySet[K, V]) str.Map[K, V] {
	m := &concCmp[K, V]{
		data: make(map[K]V),
	}
	for _, es := range e {
		for _, entry := range es.entries {
			m.data[entry.Key()] = entry.Value()
		}
	}
	return m
}

type concCmp[K comparable, V any] struct {
	mtx  sync.RWMutex
	data map[K]V
}

func (m *concCmp[K, V]) Put(key K, value V) {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	m.data[key] = value
}

func (m *concCmp[K, V]) Get(key K) (V, bool) {
	m.mtx.RLock()
	defer m.mtx.RUnlock()
	v, ok := m.data[key]
	return v, ok
}

func (m *concCmp[K, V]) Delete(key K) {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	delete(m.data, key)
}

func (m *concCmp[K, V]) Contains(key K) bool {
	m.mtx.RLock()
	defer m.mtx.RUnlock()
	_, ok := m.data[key]
	return ok
}

func (m *concCmp[K, V]) Len() int {
	m.mtx.RLock()
	defer m.mtx.RUnlock()
	return len(m.data)
}

func (m *concCmp[K, V]) Clear() {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	m.data = make(map[K]V)
}

func (m *concCmp[K, V]) Keys() []K {
	m.mtx.RLock()
	defer m.mtx.RUnlock()

	keys := make([]K, 0, m.Len())
	for key := range m.data {
		keys[len(keys)] = key
	}
	return keys
}

func (m *concCmp[K, V]) Values() []V {
	m.mtx.RLock()
	defer m.mtx.RUnlock()

	values := make([]V, 0, m.Len())
	for _, value := range m.data {
		values[len(values)] = value
	}
	return values
}

func (m *concCmp[K, V]) Iterator(comparator ...comparators.Functional[K]) str.MapIterator[K, V] {
	m.mtx.RLock()
	defer m.mtx.RUnlock()

	keys := m.Keys()

	if comparator != nil && len(comparator) > 0 {
		sorter := sorters.Quick[K, []K](comparator[0])
		sorter.Sort(&keys)
	}

	return &concCmpMapIter[K, V]{
		m:       m,
		keys:    keys,
		current: 0,
	}
}

func (m *concCmp[K, V]) String() string {
	return sprints.Map(indent.Zero(), m.data)
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
	value, ok = it.m.data[key]

	if !ok {
		panic(fmx.Sprintf("this iterator or its map is corrupt, key '%v' was expected but"+
			"was not found on available next position", key))
	}

	it.current++
	return
}
