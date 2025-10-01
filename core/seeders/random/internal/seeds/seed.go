// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package seeds

import (
	"time"
)

var gc uint64

func init() {
	// Initialize global counter
	now := uint64(time.Now().UnixNano())
	o := ((now >> 3) % 0x9E3779B97F4A7C15) + 2
	gc = (now >> 3) * o & 0xFFFFFFFFFFFFFFFF
}

// Clock returns a seed based on the current time and a global counter
// that is incremented every time the function is called
// This strategy is just to avoid stuttering seeds, it doesn't add any
// security in terms of entropy. It still very easily predictable.
func Clock() uint64 {
	nanos := uint64(time.Now().UnixNano())
	// Update the global counter with a simple XOR-shift op
	// to avoid stuttering seeds
	gc ^= gc << 13
	gc ^= gc >> 7
	gc ^= gc << 17

	return (gc ^ nanos) * 0x9E3779B97F4A7C15
}
