// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package collectors

import (
	"github.com/andrerrcosta2/gtools/core/gtools"
	"github.com/andrerrcosta2/gtools/core/gtools/functions"
	"sync"
)

// SliceCollector returns a Collector that collects all values from the given Stream into a slice.
type SliceCollector[T any] Collector[T, []T]

func Slice[T any]() SliceCollector[T] {
	return &sliceCollector[T]{}
}

type sliceCollector[T any] struct {
	once sync.Once
}

// Collect implements the Collector interface.
// It collects all values from the given Stream into a slice.
func (c *sliceCollector[T]) Collect(stream gtools.Stream[T]) []T {
	out := make([]T, 0)
	// Only run the collector once
	c.once.Do(func() {
		// Iterate over the stream
		for item := range stream {
			// Add the item to the output slice
			out = append(out, item)
		}
	})
	return out
}

func (c *sliceCollector[T]) Errors() []error {
	return nil
}

var _ Collector[int, []int] = (*sliceCollector[int])(nil)
var _ SliceCollector[int] = (*sliceCollector[int])(nil)

// Set returns a Collector that collects all values from the given Stream into a slice of unique values.
func Set[T comparable]() SliceCollector[T] {
	return &uniqueCollector[T]{}
}

type uniqueCollector[T comparable] struct {
	once sync.Once
}

// Collect implements the Collector interface.
// It collects all values from the given Stream into an array of unique values.
func (c *uniqueCollector[T]) Collect(stream gtools.Stream[T]) []T {
	out := make([]T, 0)
	// Only run the collector once
	c.once.Do(func() {
		set := make(map[T]struct{})
		// Iterate over the stream
		for item := range stream {
			// Check if the item is already in the set
			if _, exists := set[item]; !exists {
				// If not, add it to the set and the output slice
				out = append(out, item)
				set[item] = struct{}{}
			}
		}
	})
	return out
}

func (c *uniqueCollector[T]) Errors() []error {
	return nil
}

var _ Collector[int, []int] = (*uniqueCollector[int])(nil)
var _ SliceCollector[int] = (*uniqueCollector[int])(nil)

// PipedSliceCollector returns a Collector that collects all values from the given Stream into a slice after applying the given flt.
type PipedSliceCollector[I any, O any] Piped[I, O, []O]

// PipedSlice returns a Piped that collects all values from the given Stream,
// applies the given flt to each value and collects the results into a slice.
func PipedSlice[I any, O any](pipe functions.Function[I, O]) Piped[I, O, []O] {
	return &pipedSliceCollector[I, O, []O]{pipe: pipe}
}

type pipedSliceCollector[I any, O any, C ~[]O] struct {
	pipe functions.Function[I, O]
}

// Collect implements the Collector interface.
// It collects all values from the given Stream into a slice by applying the given flt to each value.
// It's thread-safe and can be used concurrently.
func (c *pipedSliceCollector[I, O, C]) Collect(stream gtools.Stream[I]) C {
	// Create a slice with the same length as the stream
	result := make(C, len(stream))
	// Iterate over the stream
	for item := range stream {
		// Apply the given flt to each value and append the result to the slice
		result = append(result, c.pipe(item))
	}
	return result
}

// Pipe returns a Piped that collects all values from the given Stream,
// applies the given flt to each value and collects the results into a slice.
func (c *pipedSliceCollector[I, O, C]) Pipe(f functions.Function[I, O]) Piped[I, O, C] {
	// Unique the given flt as the next one to be applied to the stream
	c.pipe = f
	// Return itself as a Piped
	return c
}

func (c *pipedSliceCollector[I, O, C]) Errors() []error {
	return nil
}

var _ Collector[float64, []float64] = (*pipedSliceCollector[float64, float64, []float64])(nil)
var _ Piped[int, float64, []float64] = (*pipedSliceCollector[int, float64, []float64])(nil)
