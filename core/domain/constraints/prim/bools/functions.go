// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package bools

// FromBytes converts a byte slice to a bool.
// It assumes that the byte slice is a valid binary representation of a boolean value.
// It returns the bool value represented by the byte slice.
func FromBytes(b []byte) bool {
	// We use the first byte of the slice as the boolean value.
	// 0 is false, any other value is true.
	return b[0] == 1
}

// ToBytes converts a bool to a byte slice.
// It returns a byte slice containing the binary representation of the bool.
func ToBytes(b bool) []byte {
	if b {
		return []byte{1}
	}
	return []byte{0}
}
