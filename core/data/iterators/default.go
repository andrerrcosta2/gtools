// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package iterators

import "github.com/andrerrcosta2/gtools/core/data"

// Default creates a new default implementation of data.Iterator[T] for a given slice
func Default[T any](elements []T) data.Iterator[T] {
	return &defaultIter[T]{elements: elements, index: 0}
}

// defaultIter is a simple iterator implementation for a slice of elements
type defaultIter[T any] struct {
	elements []T
	index    int
}

// HasNext checks if there are more elements in the iterator
func (it *defaultIter[T]) HasNext() bool {
	return it.index < len(it.elements)
}

// Next returns the next element in the iterator
func (it *defaultIter[T]) Next() T {
	if !it.HasNext() {
		var zeroValue T
		return zeroValue // Return the zero value if there are no more elements
	}
	element := it.elements[it.index]
	it.index++
	return element
}

var _ data.Iterator[any] = (*defaultIter[any])(nil)
