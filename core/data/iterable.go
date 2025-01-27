// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package data

// Iterable interface for pagination
type Iterable[T any] interface {
	// Iterator Returns an iterator for the elements
	Iterator() Iterator[T]
}

// Iterator interface to traverse through the elements
type Iterator[T any] interface {
	// HasNext Checks if there's a next element
	HasNext() bool

	// Next Returns the next element
	Next() T
}
