// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package queues

import (
	"github.com/andrerrcosta2/gtools/core/data/comparator"
	"github.com/andrerrcosta2/gtools/core/data/comparator/comparators"
	"github.com/andrerrcosta2/gtools/core/data/str/queue"
	"github.com/andrerrcosta2/gtools/core/domain/constraints"
)

func Priority[T any](cmp comparator.Comparer[T]) queue.MutablePriority[T] {
	return &ordered[T]{
		cmp:    cmp,
		buffer: make([]T, 0),
	}
}

func PriorityOf[T constraints.Ordered[T]]() queue.MutablePriority[T] {
	return &ordered[T]{
		cmp: comparators.OrderedOf[T](),
	}
}

type ordered[T any] struct {
	cmp    comparator.Comparer[T]
	buffer []T
}

func (o *ordered[T]) Peek() (T, bool) {
	if len(o.buffer) > 0 {
		return o.buffer[0], len(o.buffer) > 0
	}
	var zero T
	return zero, false
}

func (o *ordered[T]) Pop() (T, bool) {
	if len(o.buffer) == 0 {
		var zero T
		return zero, false
	}

	e := o.buffer[0]
	o.buffer = o.buffer[1:]
	return e, true
}

func (o *ordered[T]) Push(t T) {
	low := 0
	high := len(o.buffer)

	for low < high {
		mid := int(uint(low+high) >> 1)
		if o.cmp.Compare(o.buffer[mid], t) == -1 {
			low = mid + 1
		} else {
			high = mid
		}
	}

	o.writeAt(t, low)
}

func (o *ordered[T]) writeAt(t T, idx int) {
	o.buffer = append(o.buffer[:idx], append([]T{t}, o.buffer[idx:]...)...)
}

func (o *ordered[T]) Clear() {
	o.buffer = make([]T, 0)
}

func (o *ordered[T]) Len() int {
	return len(o.buffer)
}
