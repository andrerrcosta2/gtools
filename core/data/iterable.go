// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package data

// Iterable interface for pagination
type Iterable[D any] interface {
	// Iterator Returns an iterator for the elements
	Iterator() Iterator[D]
}

// Iterator interface to traverse through the elements
type Iterator[D any] interface {
	// HasNext Checks if there's a next element
	HasNext() bool

	// Next Returns the next element
	Next() D
}

type Traverser[D any] interface {
	Iterator(g Iterable[D], start D) Iterator[D]
}
