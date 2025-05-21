// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package str

// MapIterator is the default iterator used to loop over maps
//
// Most of the implementations implements an "occasionally consistent" behaviour
// which is standard for iterators in concurrent environments.
// It is a common trade-off in concurrent data structures to balance performance.
// simplicity, and correctness
type MapIterator[T any, U any] interface {
	// Next returns the next value from the iterator.
	//
	// It takes no parameters.
	// Returns the value of type K, the value of type U, and a boolean indicating whether the iteration is complete.
	Next() (T, U, bool)
}
