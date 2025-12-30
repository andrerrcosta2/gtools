// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package collectors

import (
	"github.com/andrerrcosta2/gtools/core/domain/functions"
	"github.com/andrerrcosta2/gtools/core/domain/gtools"
)

type Collector[I any, O any] interface {
	// Collect implements the Collector interface.
	// It collects all values from the given Stream into the output O.
	Collect(stream gtools.Stream[I]) O
	// Errors implements the Collector interface.
	// It returns any err that occurred while collecting from the given Stream.
	Errors() []error
}

type Piped[I any, O any, C any] interface {
	Collector[I, C]
	// Pipe implements the Piped interface.
	// It returns a Piped that collects all values from the given Stream,
	// applies the given flt to each value and collects the results into the output C.
	Pipe(function functions.Function[I, O]) Piped[I, O, C]
}

type Flattener[I any, O any, C any] interface {
	Collector[I, C]
	// Pipe is a method which is applied to flat each value collected to the same output pipeline.
	// The flat function must be able to handle its own zero value type
	// since its initial value cannot be out with itself and the type
	// between collection and flattening are not guaranteed to be the same
	Pipe(function functions.BiFunction[O, I, O]) Flattener[I, O, C]
}

type Resilient[I any, O any] interface {
	Collector[I, O]
	// OnError implements the Resilient interface.
	// It returns a Collector that will call the given function for each error that occurred while collecting from the given Stream.
	// The returned Collector will collect the results of the function instead of the error.
	OnError(function functions.Function[error, O]) Resilient[I, O]
}

type ResilientPiped[I any, O any, C any] interface {
	Piped[I, O, C]
	// OnError implements the Resilient interface.
	// It returns a Collector that will call the given function for each error that occurred while collecting from the given Stream.
	// The returned Collector will collect the results of the function instead of the error.
	OnError(function functions.Function[error, O]) ResilientPiped[I, O, C]
}
