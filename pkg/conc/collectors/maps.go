// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package collectors

import (
	"github.com/andrerrcosta2/gtools/core/data/str"
	"github.com/andrerrcosta2/gtools/core/domain/gtools"
	"sync"
)

type MapCollector[K comparable, V any, C ~map[K]V] Collector[str.Entry[K, V], C]

// Map returns a MapCollector that collects all values from the given Stream into a map.
// The Stream should contain values of type str.Entry[K, V].
// The collector returns a map with the key as the key and the value as the value of the entry.
func Map[K comparable, V any]() MapCollector[K, V, map[K]V] {
	return &mapCollector[K, V]{errors: make([]error, 0)}
}

type mapCollector[K comparable, V any] struct {
	once   sync.Once
	errors []error
}

// Collect implements the Collector interface.
// It collects all values from the given Stream into a map.
func (c *mapCollector[K, V]) Collect(stream gtools.Stream[str.Entry[K, V]]) map[K]V {
	out := make(map[K]V)
	// Only run the collector once
	c.once.Do(func() {
		// Iterate over the stream
		for entry := range stream {
			//key := entry.Name()
			//if key == nil {
			//	c.err = append(c.err, fmt.Errorf("entry with nil key: %v", entry))
			//	continue
			//}
			// Add the item to the output map
			out[entry.Key()] = entry.Value()
		}
	})
	return out
}

func (c *mapCollector[K, V]) Errors() []error {
	return c.errors
}

var _ Collector[str.Entry[int, string], map[int]string] = (*mapCollector[int, string])(nil)
var _ MapCollector[int, string, map[int]string] = (*mapCollector[int, string])(nil)
