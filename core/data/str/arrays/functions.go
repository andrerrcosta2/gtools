// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package arrays

// SetAt sets the value at the given index in the given array.
//
// It takes the array 'a', the index idx and the value data as arguments.
// If the index is out of bounds of the array, it resizes the array to fit the index and sets the value at the index.
// It returns the modified array.
func SetAt[E ~[]T, T any](a E, idx int, data T) E {
	if idx >= len(a) {
		// If the index is out of bounds, resize the array to fit the index
		space := make(E, idx+1)
		// Copy the original array to the new one
		copy(space, a)
		// ToSet the original array to the new resized array
		a = space
	}
	// ToSet the value at the given index in the array
	a[idx] = data
	// Return the modified array
	return a
}

// Grow grows the array by adding 'rows' empty rows at the end.
//
// It takes the array 'a' and the number of rows 'rows' as arguments.
// It returns the modified array.
func Grow[E ~[]T, T any](a *E, rows int) {
	// Create a new array with the correct number of rows
	newRows := make(E, len(*a)+rows)
	// Copy all existing rows to the new array
	copy(newRows, *a)
	// Replace the original array with the new one
	*a = newRows
}

// GrowIfLessThan grows the array by adding 'size' empty rows at the end only if the current length of the array is less than 'size'.
//
// It takes the array 'a' and the desired length 'size' as arguments.
// If the length of the array is less than 'size', it calls Grow to add the required number of empty elements at the end.
// It returns the modified array.
func GrowIfLessThan[E ~[]T, T any](a *E, size int) {
	if len(*a) < size {
		Grow(a, size-len(*a))
	}
}
