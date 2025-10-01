// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package ptrs

import "unsafe"

// New creates a new pointer for a given value
func New[T any](value T) *T {
	return &value
}

// Nil creates a nil pointer of a given type
func Nil[T any]() *T {
	return (*T)(nil)
}

// Slice returns a slice-like pointer of a given type and given elements
func Slice[T ~[]E, E any](s ...E) *T {
	tmp := T(s)
	return &tmp
}

// Unsafe creates an unsafe pointer for a given value
func Unsafe[T any](value T) unsafe.Pointer {
	return unsafe.Pointer(&value)
}
