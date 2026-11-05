// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package collectors

import (
	"github.com/andrerrcosta2/gtools/conc/core/collector"
	"github.com/andrerrcosta2/gtools/conc/streams"
	"github.com/andrerrcosta2/gtools/core/domain/contracts/stream"
	"github.com/andrerrcosta2/gtools/core/domain/function"
)

// SliceCollector returns a Collector that collects all values from the given Stream into a slice.
type SliceCollector[T any] collector.Collector[T, []T]

func Slice[T any]() SliceCollector[T] {
	return &sliceCollector[T]{}
}

type sliceCollector[T any] struct{}

// Collect implements the Collector interface.
// It collects all values from the given Stream into a slice.
func (c *sliceCollector[T]) Collect(src stream.Readable[T]) []T {
	out := make([]T, 0)
	streams.Range(src, func(item T) bool {
		out = append(out, item)
		return true
	})
	return out
}

func (c *sliceCollector[T]) Errors() []error {
	return nil
}

var _ collector.Collector[int, []int] = (*sliceCollector[int])(nil)
var _ SliceCollector[int] = (*sliceCollector[int])(nil)

// Set returns a Collector that collects all values from the given Stream into a slice of unique values.
func Set[T comparable]() SliceCollector[T] {
	return &uniqueCollector[T]{}
}

type uniqueCollector[T comparable] struct{}

// Collect implements the Collector interface.
// It collects all values from the given Stream into an array of unique values.
func (c *uniqueCollector[T]) Collect(src stream.Readable[T]) []T {
	out := make([]T, 0)
	set := make(map[T]struct{})
	streams.Range(src, func(item T) bool {
		// Check if the item is already in the set
		if _, exists := set[item]; !exists {
			// If not, add it to the set and the output slice
			out = append(out, item)
			set[item] = struct{}{}
		}
		return true
	})
	return out
}

func (c *uniqueCollector[T]) Errors() []error {
	return nil
}

var _ collector.Collector[int, []int] = (*uniqueCollector[int])(nil)
var _ SliceCollector[int] = (*uniqueCollector[int])(nil)

// PipedSliceCollector returns a Collector that collects all values from the given Stream into a slice after applying the given flt.
type PipedSliceCollector[I any, O any] collector.Piped[I, O, []O]

// PipedSlice returns a Piped that collects all values from the given Stream,
// applies the given flt to each value and collects the results into a slice.
func PipedSlice[I any, O any](pipe function.Function[I, O]) collector.Piped[I, O, []O] {
	return &pipedSliceCollector[I, O, []O]{pipe: pipe}
}

type pipedSliceCollector[I any, O any, C ~[]O] struct {
	pipe function.Function[I, O]
}

// Collect implements the Collector interface.
// It collects all values from the given Stream into a slice by applying the given flt to each value.
// It's thread-safe and can be used concurrently.
func (c *pipedSliceCollector[I, O, C]) Collect(src stream.Readable[I]) C {
	result := make(C, tryLen(src))
	// Iterate over the src
	streams.Range(src, func(item I) bool {
		result = append(result, c.pipe(item))
		return true
	})
	return result
}

// Pipe returns a Piped that collects all values from the given Stream,
// applies the given flt to each value and collects the results into a slice.
func (c *pipedSliceCollector[I, O, C]) Pipe(f function.Function[I, O]) collector.Piped[I, O, C] {
	// Unique the given flt as the next one to be applied to the stream
	c.pipe = f
	// Return itself as a Piped
	return c
}

func (c *pipedSliceCollector[I, O, C]) Errors() []error {
	return nil
}

var _ collector.Collector[float64, []float64] = (*pipedSliceCollector[float64, float64, []float64])(nil)
var _ collector.Piped[int, float64, []float64] = (*pipedSliceCollector[int, float64, []float64])(nil)
