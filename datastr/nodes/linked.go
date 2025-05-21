// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package nodes

import (
	"fmt"
	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"github.com/andrerrcosta2/gtools/core/format/sprints"
)

func NewTypedDoubleLinked[T any](prev, next *TypedDoubleLinked[T], value T) *TypedDoubleLinked[T] {
	return &TypedDoubleLinked[T]{
		value: value,
		prev:  prev,
		next:  next,
	}
}

type TypedDoubleLinked[T any] struct {
	prev  *TypedDoubleLinked[T]
	next  *TypedDoubleLinked[T]
	value T
}

func (n *TypedDoubleLinked[T]) Next() *TypedDoubleLinked[T] {
	return n.next
}

func (n *TypedDoubleLinked[T]) Prev() *TypedDoubleLinked[T] {
	return n.prev
}

func (n *TypedDoubleLinked[T]) Value() T {
	return n.value
}

func (n *TypedDoubleLinked[T]) SetNext(next *TypedDoubleLinked[T]) {
	n.next = next
}

func (n *TypedDoubleLinked[T]) SetPrev(prev *TypedDoubleLinked[T]) {
	n.prev = prev
}

func (n *TypedDoubleLinked[T]) SetValue(value T) {
	n.value = value
}

func (n *TypedDoubleLinked[T]) String() string {
	var prev, next = "..." + nilDbl, nilDbl + "..."
	if n.prev != nil {
		prev = sprints.Lfield(indent.Tab(1), "...prev", fmt.Sprintf(dbl, n.prev.value))
	}
	if n.next != nil {
		next = sprints.Lfield(indent.Tab(1), "next", fmt.Sprintf(dbl+"...", n.next.value))
	}
	return "TypedDoubleLinked{" + prev +
		sprints.Lfield(indent.Tab(1), "value", sprints.Valuef(n.value)) +
		next + "\n}"
}

const nilDbl = "*TypedDoubleLinked<nil>"
const dbl = "*TypedDoubleLinked{\n\tvalue: %v\n}"
