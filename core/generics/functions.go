// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package generics

func Zero[T any]() T {
	var zero T
	return zero
}

func IsZero[T comparable](value T) bool {
	var zero T
	return value == zero
}
