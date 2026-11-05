// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package collectors

import (
	"github.com/andrerrcosta2/gtools/conc/core/collector"
	"github.com/andrerrcosta2/gtools/conc/streams"
	"github.com/andrerrcosta2/gtools/core/data/str/dict"
	"github.com/andrerrcosta2/gtools/core/domain/contracts/stream"
)

type MapCollector[K comparable, V any, C ~map[K]V] collector.Collector[dict.Entry[K, V], C]

// Map returns a MapCollector that collects all values from the given Stream into a map.
// The Stream should contain values of type str.Entry[K, V].
// The collector returns a map with the key as the key and the value as the value of the entry.
func Map[K comparable, V any]() MapCollector[K, V, map[K]V] {
	return &mapCollector[K, V]{errors: make([]error, 0)}
}

type mapCollector[K comparable, V any] struct {
	errors []error
}

// Collect implements the Collector interface.
// It collects all values from the given Stream into a map.
func (c *mapCollector[K, V]) Collect(src stream.Readable[dict.Entry[K, V]]) map[K]V {
	out := make(map[K]V)
	streams.Range(src, func(e dict.Entry[K, V]) bool {
		out[e.Key()] = e.Value()
		return true
	})
	return out
}

func (c *mapCollector[K, V]) Errors() []error {
	return c.errors
}

var _ collector.Collector[dict.Entry[int, string], map[int]string] = (*mapCollector[int, string])(nil)
var _ MapCollector[int, string, map[int]string] = (*mapCollector[int, string])(nil)
