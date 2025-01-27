// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package mod

const (
	// ^uint(0) can be 0xFFFFFFFFFFFFFFFF or 0xFFFFFFFF. After shifting it by 63 bits
	// only the most significant bit remains.
	intSize = 32 << (^uint(0) >> 63) // This multiplication can be 32 or 64.

	Uint            = 1<<intSize - 1
	Uint8           = 0xFF
	Uint16          = 0xFFFF
	Uint32          = 0xFFFFFFFF
	Uint64          = 0xFFFFFFFFFFFFFFFF
	Float32Fraction = 1<<23 - 1
)
