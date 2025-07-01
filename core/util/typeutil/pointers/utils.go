// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package pointers

// New creates a new pointer for a given value
func New[T any](value T) *T {
	return &value
}
