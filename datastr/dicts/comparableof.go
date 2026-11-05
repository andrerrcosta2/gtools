// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package dicts

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/andrerrcosta2/gtools/core/data/comparator"
	"github.com/andrerrcosta2/gtools/core/data/keyer"
	"github.com/andrerrcosta2/gtools/core/data/keyer/keyers"
	"github.com/andrerrcosta2/gtools/core/data/str/dict"
	"github.com/andrerrcosta2/gtools/core/domain/constraints"
	"github.com/andrerrcosta2/gtools/core/sortables/sorters"
)

// ComparableOf returns a new instance of cmpOf.
// It creates a new map with a comparator for the given key type K.
func ComparableOf[K constraints.Comparable[K], V any](entries ...dict.Entry[K, V]) dict.Mutable[K, V] {
	keyer := keyers.FingerPrint[K]()
	m := &cmpOf[K, V]{
		// Initialize the map with a string key type.
		data:  make(map[string]dict.Entry[K, V]),
		keyer: keyer,
	}

	// Add each entry to the map.
	for _, entry := range entries {
		// Last entry wins
		hash := keyer.Key(entry.Key())
		m.data[hash] = entry
	}

	// Return the cmpOf instance.
	return m
}

type cmpOf[K constraints.Comparable[K], V any] struct {
	data  map[string]dict.Entry[K, V]
	keyer keyer.Comparable[K, string]
}

// Put adds a new key-value pair to the map.
// If the key already exists, the old value is replaced.
func (m *cmpOf[K, V]) Put(key K, value V) {
	m.data[m.hash(key)] = NewAnyEntry(key, value)
}

func (m *cmpOf[K, V]) Get(key K) (V, bool) {
	v, ok := m.data[m.hash(key)]
	if !ok {
		var zero V
		return zero, false
	}
	return v.Value(), ok
}

func (m *cmpOf[K, V]) Clear() {
	m.data = make(map[string]dict.Entry[K, V])
}

func (m *cmpOf[K, V]) Contains(key K) bool {
	_, ok := m.data[m.hash(key)]
	return ok
}

func (m *cmpOf[K, V]) Delete(key K) {
	delete(m.data, m.hash(key))
}

func (m *cmpOf[K, V]) hash(key K) string {
	return m.keyer.Key(key)
}

func (m *cmpOf[K, V]) Keys() []K {
	out := make([]K, 0, len(m.data))
	for _, entry := range m.data {
		out = append(out, entry.Key())
	}
	return out
}

func (m *cmpOf[K, V]) Len() int {
	return len(m.data)
}

func (m *cmpOf[K, V]) Values() []V {
	out := make([]V, 0, len(m.data))
	for _, entry := range m.data {
		out = append(out, entry.Value())
	}
	return out
}

// Iterator the variadic parameter is just a trick to allow to use the iterator without requiring parameters.
// its presence indicates the keys must be sorted.
func (m *cmpOf[K, V]) Iterator(comparator ...comparator.EqComparer[K]) dict.MapIterator[K, V] {
	keys := make([]K, 0, len(m.data))
	for _, v := range m.data {
		keys = append(keys, v.Key())
	}

	if len(comparator) > 0 {
		sorter := sorters.Quick[[]K](comparator[0])
		sorter.Sort(&keys)
	}

	return &cmpOfIter[K, V]{
		m:       m,
		keys:    keys,
		current: 0,
	}
}

func (m *cmpOf[K, V]) String() string {
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

type cmpOfIter[K constraints.Comparable[K], V any] struct {
	m       *cmpOf[K, V]
	keys    []K
	current int
}

func (it *cmpOfIter[K, V]) Next() (key K, value V, ok bool) {
	if it.current >= len(it.keys) {
		return
	}

	mapKey := it.keys[it.current]
	entry, ok := it.m.data[it.m.hash(mapKey)]

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

func ConcComparableOf[K constraints.Sortable[K], V any](entries ...dict.Entry[K, V]) dict.Mutable[K, V] {
	keyer := keyers.FingerPrint[K]()
	m := &concCmpOf[K, V]{
		mtx:   sync.RWMutex{},
		data:  make(map[string]dict.Entry[K, V]),
		keyer: keyer,
	}

	for _, entry := range entries {
		m.data[keyer.Key(entry.Key())] = NewAnyEntry(entry.Key(), entry.Value())
	}
	return m
}

type concCmpOf[K constraints.Sortable[K], V any] struct {
	mtx   sync.RWMutex
	data  map[string]dict.Entry[K, V]
	keyer keyer.Comparable[K, string]
}

func (m *concCmpOf[K, V]) Clear() {
	m.mtx.Lock()
	m.data = make(map[string]dict.Entry[K, V])
	m.mtx.Unlock()
}

func (m *concCmpOf[K, V]) Contains(key K) bool {
	m.mtx.RLock()
	_, ok := m.data[m.hash(key)]
	m.mtx.RUnlock()
	return ok
}

func (m *concCmpOf[K, V]) Delete(key K) {
	m.mtx.Lock()
	delete(m.data, m.hash(key))
	m.mtx.Unlock()
}

func (m *concCmpOf[K, V]) Get(key K) (value V, exists bool) {
	m.mtx.RLock()
	v, ok := m.data[m.hash(key)]
	m.mtx.RUnlock()
	if !ok {
		return value, false
	}
	return v.Value(), ok
}

func (m *concCmpOf[K, V]) hash(key K) string {
	return m.keyer.Key(key)
}

func (m *concCmpOf[K, V]) Keys() []K {
	m.mtx.RLock()
	defer m.mtx.RUnlock()
	out := make([]K, 0, len(m.data))
	for _, entry := range m.data {
		out = append(out, entry.Key())
	}
	return out
}

func (m *concCmpOf[K, V]) Len() int {
	m.mtx.RLock()
	defer m.mtx.RUnlock()
	return len(m.data)
}

func (m *concCmpOf[K, V]) Put(key K, value V) {
	m.mtx.Lock()
	m.data[m.hash(key)] = NewAnyEntry(key, value)
	m.mtx.Unlock()
}

func (m *concCmpOf[K, V]) Values() []V {
	m.mtx.RLock()
	defer m.mtx.RUnlock()
	out := make([]V, 0, len(m.data))
	for _, entry := range m.data {
		out = append(out, entry.Value())
	}
	return out
}

func (m *concCmpOf[K, V]) String() string {
	m.mtx.RLock()
	defer m.mtx.RUnlock()
	var sb strings.Builder
	for _, entry := range m.data {
		sb.WriteString(fmt.Sprintf("%s\n", entry.String()))
	}
	return sb.String()
}

func (m *concCmpOf[K, V]) Iterator(comparator ...comparator.EqComparer[K]) dict.MapIterator[K, V] {
	keys := make([]K, 0, len(m.data))
	for _, v := range m.data {
		keys = append(keys, v.Key())
	}

	if len(comparator) > 0 {
		sorter := sorters.Quick[[]K](comparator[0])
		sorter.Sort(&keys)
	}

	return &concCmpOfIter[K, V]{
		m:       m,
		keys:    keys,
		current: 0,
	}
}

type concCmpOfIter[K constraints.Sortable[K], V any] struct {
	m       *concCmpOf[K, V]
	keys    []K
	current int
}

func (it *concCmpOfIter[K, V]) Next() (key K, value V, ok bool) {
	it.m.mtx.RLock()
	defer it.m.mtx.RUnlock()

	if it.current >= len(it.keys) {
		return
	}

	mapKey := it.keys[it.current]
	entry, ok := it.m.data[it.m.hash(mapKey)]

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
