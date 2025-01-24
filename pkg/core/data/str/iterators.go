// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package str

type Loopable[T any] interface {
	// Loop returns a read-only channel that yields values of type T.
	//
	// No parameters.
	// Returns a receive-only channel of type T.
	Loop() <-chan T
}

type MapIterator[T any, U any] interface {
	// Next returns the next value from the iterator.
	//
	// It takes no parameters.
	// Returns the value of type T, the value of type U, and a boolean indicating whether the iteration is complete.
	Next() (T, U, bool)
}
