// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package lgc

import (
	"github.com/andrerrcosta2/gtools/core/seeders/random/internal/mod"
	"github.com/andrerrcosta2/gtools/core/seeders/random/internal/prng/inc"
	"github.com/andrerrcosta2/gtools/core/seeders/random/internal/prng/mult"
	"math"
)

func Int(seed uint64) uint64 { return (mult.Int64*seed + inc.Int64) & mod.Uint64 }

func Int8(seed uint64) uint64 {
	return (mult.Int8*seed + inc.Int8) & mod.Uint64
}

func Int16(seed uint64) uint64 {
	return (mult.Int16*seed + inc.Int16) & mod.Uint64
}

func Int32(seed uint64) uint64 {
	return (mult.Int32*seed + inc.Int32) & mod.Uint64
}

func Int64(seed uint64) uint64 {
	return (mult.Int64*seed + inc.Int64) & mod.Uint64
}

func Uint(seed uint64) uint64 {
	return (mult.Uint64*seed + inc.Uint64) & mod.Uint64
}

func Uint8(seed uint64) uint64 {
	return (mult.Uint8*seed + inc.Uint8) & mod.Uint64
}

func Uint16(seed uint64) uint64 {
	return (mult.Uint16*seed + inc.Uint16) & mod.Uint64
}

func Uint32(seed uint64) uint64 {
	return (mult.Uint32*seed + inc.Uint32) & mod.Uint64
}

func Uint64(seed uint64) uint64 {
	return (mult.Uint64*seed + inc.Uint64) & mod.Uint64
}

func Float32(s1, s2 uint64) (is float64, fs uint64) {
	ls1 := (mult.Uint64*s1 + inc.Uint64) & 0xFFFFFFFFFFFFFFFF
	fs = (mult.Uint64*s2 + inc.Uint64) & 0xFFFFFFFFFFFFFFFF

	if ls1 == math.MaxUint64 {
		return 0.0, fs
	}

	// Scale ls into the float64 range
	//fmt.Printf("ls1: %d\n", ls1)
	ss1 := float64(ls1) * (math.MaxFloat64 / math.MaxUint64)
	//fmt.Printf("ss1: %f\n", float64(ls1)*(math.MaxFloat64/math.MaxUint64))

	// Calculate overflow parameters
	s1to := math.MaxFloat64 - ss1  // Remaining space to overflow
	s1or := math.MaxFloat64 / s1to // Scaling ratio to simulate overflow
	// this value may require adjustment. in theory s1or
	// may be lower than this constant
	s1ov := 95384372348.0 / s1or // Simulated overflow value

	// Extract fractional overflow
	s1fo := s1ov - math.Floor(s1ov)

	// Calculate the final adjusted value
	is = ss1 * s1fo

	return
}

func Float64(s1, s2 uint64) (is float64, fs uint64) {
	ls1 := (mult.Uint64*s1 + inc.Uint64) & 0xFFFFFFFFFFFFFFFF
	fs = (mult.Uint64*s2 + inc.Uint64) & 0xFFFFFFFFFFFFFFFF

	if ls1 == math.MaxUint64 {
		return 0.0, fs
	}

	// Scale ls into the float64 range
	ss1 := float64(ls1) * (math.MaxFloat64 / math.MaxUint64)

	// Calculate overflow parameters
	s1to := math.MaxFloat64 - ss1  // Remaining space to overflow
	s1or := math.MaxFloat64 / s1to // Scaling ratio to simulate overflow
	// this value may require adjustment. in theory s1or
	// may be lower than this constant
	s1ov := 95384372348.0 / s1or // Simulated overflow value

	// Extract fractional overflow
	s1fo := s1ov - math.Floor(s1ov)

	// Calculate the final adjusted value
	is = ss1 * s1fo

	return
}
