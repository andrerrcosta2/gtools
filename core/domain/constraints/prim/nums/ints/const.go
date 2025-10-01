// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package ints

const (
	Size int = 32 << (^uint(0) >> 63) // 32 or 64

	Min8 = -1 << 7  // Min8 smallest 8 bitwise integer
	Max8 = 1<<7 - 1 // Max8 highest int8

	Min16 = -1 << 15  // Min16 smallest int16
	Max16 = 1<<15 - 1 // Max16 smallest int16

	Min32 = -1 << 31  // Min8 smallest int8: -2147483648
	Max32 = 1<<31 - 1 // 2147483647

	// Min64  smallest int64
	Min64 = -1 << 63
	Max64 = 1<<63 - 1

	// Min smallest int8
	Min = -1 << (Size - 1)
	Max = 1<<(Size-1) - 1
)
