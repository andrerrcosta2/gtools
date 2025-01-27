// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package iterators

// DefaultIterator is a simple iterator implementation for a slice of elements
type DefaultIterator[T any] struct {
	elements []T
	index    int
}

// Default creates a new DefaultIterator for the given slice
func Default[T any](elements []T) *DefaultIterator[T] {
	return &DefaultIterator[T]{elements: elements, index: 0}
}

// HasNext checks if there are more elements in the iterator
func (it *DefaultIterator[T]) HasNext() bool {
	return it.index < len(it.elements)
}

// Next returns the next element in the iterator
func (it *DefaultIterator[T]) Next() T {
	if !it.HasNext() {
		var zeroValue T
		return zeroValue // Return the zero value if there are no more elements
	}
	element := it.elements[it.index]
	it.index++
	return element
}
