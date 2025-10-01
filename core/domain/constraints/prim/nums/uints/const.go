// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package uints

const (
	Min8 = 0        // smallest 8 bitwise unsigned integer
	Max8 = 1<<8 - 1 // highest 8 bitwise unsigned integer

	Min16 = 0         // smallest 16 bitwise unsigned integer
	Max16 = 1<<16 - 1 // highest 16 bitwise unsigned integer

	Min32 = 0         // smallest 32 bitwise unsigned integer
	Max32 = 1<<32 - 1 // highest 32 bitwise unsigned integer

	Min64 = 0         // smallest 64 bitwise unsigned integer
	Max64 = 1<<64 - 1 // highest 64 bitwise unsigned integer

	Min = 0        // smallest architecture size bitwise unsigned integer
	Max = ^uint(0) // highest architecture size bitwise unsigned integer
)
