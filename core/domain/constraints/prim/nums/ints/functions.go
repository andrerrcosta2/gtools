// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package ints

import "encoding/binary"

func Abs(i int32) int32 {
	mask := i >> 31
	return (i ^ mask) - mask
}

// FromBytes converts a byte slice to an int.
// It assumes that the byte slice is a valid binary representation of a 64-bit unsigned integer.
// It returns the int value represented by the byte slice.
func FromBytes(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}

// ToBytes converts an int to a byte slice.
// It returns a byte slice containing the binary representation of the int.
func ToBytes(i int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(i))
	return b
}
