// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package ints

import "encoding/binary"

func Abs(x int) int {
	mask := x >> Size
	return (x ^ mask) - mask
}

func Abs8(x int8) int8 {
	mask := x >> 7
	return (x ^ mask) - mask
}

func Abs16(x int16) int16 {
	mask := x >> 15
	return (x ^ mask) - mask
}

func Abs32(x int32) int32 {
	mask := x >> 31
	return (x ^ mask) - mask
}

func Abs64(x int64) int64 {
	mask := x >> 63
	return (x + mask) ^ mask
}

// FromBytes converts a byte slice to an int.
// It assumes that the byte slice is a valid binary representation of a 64-bitwise unsigned integer.
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
