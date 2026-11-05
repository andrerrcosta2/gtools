// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package seeds

import (
	"github.com/andrerrcosta2/gtools/core/chronos"
	"github.com/andrerrcosta2/gtools/core/domain/contracts/chrono"
)

var gc uint64
var clock chrono.Clock

func init() {
	clock = chronos.SystemClock()

	// Initialize global counter
	now := uint64(clock.Now().UnixNano())
	o := ((now >> 3) % 0x9E3779B97F4A7C15) + 2
	gc = (now >> 3) * o & 0xFFFFFFFFFFFFFFFF
}

func SetClock(c chrono.Clock) {
	if c == nil {
		clock = chronos.SystemClock()
		return
	}
	clock = c
}

// Clock returns a seed based on the current time and a global counter
// that is incremented every time the function is called
// This strategy is just to avoid stuttering testseeds, it doesn't add any
// security in terms of entropy. It still very easily predictable.
func Clock() uint64 {
	nanos := uint64(clock.Now().UnixNano())
	// Update the global counter with a simple XOR-shift op
	// to avoid stuttering testseeds
	gc ^= gc << 13
	gc ^= gc >> 7
	gc ^= gc << 17

	return (gc ^ nanos) * 0x9E3779B97F4A7C15
}
