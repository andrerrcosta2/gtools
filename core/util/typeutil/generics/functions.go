// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package generics

// Zero returns a zero value of the given type
func Zero[T any]() T {
	var zero T
	return zero
}

// IsZero reports whether a value is zero
func IsZero[T comparable](value T) bool {
	var zero T
	return value == zero
}
