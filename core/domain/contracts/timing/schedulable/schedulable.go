// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package schedulable

import "github.com/andrerrcosta2/gtools/core/domain/constraints"

type Event[T constraints.Sortable[T]] interface {
	At() T
}

type Recurring[T constraints.Sortable[T]] interface {
	Next(from T) (next T)
}
