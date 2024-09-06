// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package util

func ArraysEquals[T comparable](a, b []T) bool {
	wa, wb := a, b
	// Check if the slices have different lengths
	if len(wa) != len(wb) {
		return false
	}

	// Compare each element of the slices
	for i := range wa {
		if wa[i] != wb[i] {
			return false
		}
	}

	// Slices are equal
	return true
}
