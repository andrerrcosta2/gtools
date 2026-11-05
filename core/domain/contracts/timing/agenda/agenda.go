// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package agenda

import (
	"github.com/andrerrcosta2/gtools/core/domain/contracts/timing/chrono"
	"github.com/andrerrcosta2/gtools/core/domain/function"
)

type Chrono[T any] interface {
	Push(T)
	Due(now chrono.Time) []T
	// RangeDue invokes fn for every task due since the last call, stopping early
	// if fn returns false. Returns false if iteration was stopped early — any
	// undrained periods remain pending and will be retried on the next call.
	RangeDue(ow chrono.Time, fn function.Predicate[T]) bool
	// Drain retrieve all pending events in the agenda using a range function
	// and returns how many events were drained
	Drain(fn function.Predicate[T]) int
}
