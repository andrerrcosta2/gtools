// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package gtypes

// IndexableSlice is a safe slice for direct
// index storage
type IndexableSlice[T any] []T

// Store stores a value at the given index
// and resizes the slice if needed
func (s *IndexableSlice[T]) Store(i uint, t T) {
	idx := int(i)
	if idx >= len(*s) {
		newCap := cap(*s)
		if newCap == 0 {
			newCap = 2
		}
		for newCap <= idx {
			newCap *= 2 // Exponential growth
		}
		newSlice := make([]T, idx+1, newCap)
		copy(newSlice, *s)
		*s = newSlice
	}
	(*s)[idx] = t
}

// Get returns the value at the given index
func (s *IndexableSlice[T]) Get(i uint) (t T, ok bool) {
	if i >= uint(len(*s)) {
		return
	}
	return (*s)[i], true
}

// Len returns the length of the slice
func (s *IndexableSlice[T]) Len() int {
	return len(*s)
}

// Cap returns the capacity of the slice
func (s *IndexableSlice[T]) Cap() int {
	return cap(*s)
}
