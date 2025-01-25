// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package async

import (
	"github.com/andrerrcosta2/gtools/core/domain/functions"
	"sync"
)

func MapOf[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{
		data:  make(map[K]*asyncData[V]),
		async: make(map[K]*asyncData[V]),
	}
}

// Map is a thread-safe map of key-value pairs of asynchronous values.
// The path to optimization here seems to be related between the use of linear data structures
// with the actual values, and maps with asynchronous data. For now, I'm not working on it.
type Map[K comparable, V any] struct {
	mtx   sync.RWMutex
	async map[K]*asyncData[V]
	data  map[K]*asyncData[V]
}

func (m *Map[K, V]) Put(k K, v V) {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	m.data[k] = &asyncData[V]{data: v}
}

func (m *Map[K, V]) PutIfAbsent(k K, v V) {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	if _, ok := m.data[k]; ok {
		return
	}
	m.data[k] = &asyncData[V]{data: v}
}

func (m *Map[K, V]) PutAsync(k K, supplier functions.Supplier2[V, error]) {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	pair := &asyncData[V]{sync: supplier}
	m.data[k] = pair
	m.async[k] = pair
}

func (m *Map[K, V]) PutIfAbsentAsync(k K, supplier functions.Supplier2[V, error]) {
	if _, ok := m.data[k]; ok {
		return
	}
	pair := &asyncData[V]{sync: supplier}
	m.data[k] = pair
	m.async[k] = pair
}

func (m *Map[K, V]) Sync() (errs []error) {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	for k, v := range m.async {
		err := v.Sync()
		if err != nil {
			errs = append(errs, err)
		} else {
			delete(m.async, k)
		}
	}
	return
}

func (m *Map[K, V]) SyncValue(k K) (exists bool, err error) {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	if pair, ok := m.data[k]; ok {
		if err = pair.Sync(); err == nil {
			delete(m.async, k)
		}
		return true, err
	}
	return false, nil
}

func (m *Map[K, V]) Get(k K) (data V, exists, sync bool) {
	m.mtx.RLock()
	defer m.mtx.RUnlock()
	if pair, ok := m.data[k]; ok {
		data, sync = pair.Value()
		return data, true, sync
	}
	return data, false, false
}

func (m *Map[K, V]) Remove(k K) {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	delete(m.data, k)
	delete(m.async, k)
}

func (m *Map[K, V]) Keys() []K {
	m.mtx.RLock()
	defer m.mtx.RUnlock()
	out := make([]K, 0, len(m.data))
	for k := range m.data {
		out = append(out, k)
	}
	return out
}

func (m *Map[K, V]) Values() []V {
	m.mtx.RLock()
	defer m.mtx.RUnlock()
	out := make([]V, 0, len(m.data))
	for _, v := range m.data {
		if val, exists := v.Value(); exists {
			out = append(out, val)
		}
	}
	return out
}
