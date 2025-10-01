// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package prng

import (
	"github.com/andrerrcosta2/gtools/core/domain/constraints/prim/nums/uints"
	"github.com/andrerrcosta2/gtools/core/seeders/random/internal/prng/lgc"
	"github.com/andrerrcosta2/gtools/core/seeders/random/internal/seeds"
	"math"
)

// Bool generates a pseudo-random bool
func Bool() bool {
	return seeds.Clock()&1 == 1
}

// Bytes generates a pseudo-random byte slice between the given range
func Bytes(m, M int) []byte {
	// Generate the random length of the byte slice
	randomBytes := make([]byte, Int(m, M))

	// Fill the byte slice with random values
	for i := range randomBytes {
		rand := randInt(lgc.Int8(seeds.Clock()), 0, 255)
		randomBytes[i] = byte(rand)
	}

	return randomBytes
}

// Complex64 generates a pseudo-random complex64 between the given range
func Complex64(rm, rM, im, iM float32) complex64 {
	rs1, rs2 := lgc.Float32(seeds.Clock(), seeds.Clock())
	is1, is2 := lgc.Float32(seeds.Clock(), seeds.Clock())

	var rp, ip float32
	if rm == rM {
		rp = rm
	} else {
		rp = f32(rs1, rs2, rm, rM)
	}
	if im == iM {
		ip = im
	} else {
		ip = f32(is1, is2, im, iM)
	}
	return complex(rp, ip)
}

// Complex128 generates a pseudo-random complex128 between the given range
func Complex128(rm, rM, im, iM float64) complex128 {
	rs1, rs2 := lgc.Float64(seeds.Clock(), seeds.Clock())
	is1, is2 := lgc.Float64(seeds.Clock(), seeds.Clock())

	var rp, ip float64
	if rm == rM {
		rp = rm
	} else {
		rp = f64(rs1, rs2, rm, rM)
	}
	if im == iM {
		ip = im
	} else {
		ip = f64(is1, is2, im, iM)
	}
	return complex(rp, ip)
}

// Float32 generates a pseudo-random float32 between the given range
func Float32(m, M float32) float32 {
	s1, s2 := lgc.Float32(seeds.Clock(), seeds.Clock())
	return f32(s1, s2, m, M)
}

func f32(s1, s2 uint64, m, M float32) float32 {
	const K = 1.0 / float32(uints.Max32)
	base := float32(s1) * K //  [0,1)
	inc := float32(s2&((1<<52)-1)) * (1.0 / (1 << 52))
	norm := base + inc
	norm -= float32(uint32(norm)) //[0,1)
	return m*(1.0-norm) + M*norm
}

// Float64 generates a pseudo-random float64 between the given range
func Float64(m, M float64) float64 {
	s1, s2 := lgc.Float64(seeds.Clock(), seeds.Clock())
	return f64(s1, s2, m, M)
}

func f64(s1, s2 uint64, m, M float64) float64 {
	const mantissaBits = 52
	const mask = (1 << mantissaBits) - 1
	const scale = 1.0 / float64(1<<mantissaBits)

	base := float64(s1) / float64(math.MaxUint64) // ∈ [0, 1)
	inc := float64(s2&mask) * scale               // jitter
	norm := base + inc
	norm -= float64(uint32(norm)) // wrap to [0, 1)

	return m*(1.0-norm) + M*norm
}

func f64t(seed uint64, m, M float64) float64 {
	const mantissaBits = 52
	const mask = (1 << mantissaBits) - 1
	const scale = 1.0 / float64(1<<mantissaBits)

	// Extract mantissas from m and M, scaled into [0,1)
	tm := float64(math.Float64bits(m)&mask) * scale
	tM := float64(math.Float64bits(M)&mask) * scale

	// Generate factor from seed
	norm := float64(seed&mask) * scale

	// Interpolate proportionally in the [tm, tM) domain
	return tm + norm*(tM-tm)
}

// Int generates a pseudo-random int between the given range
func Int(m, M int) int {
	seed := lgc.Int(seeds.Clock())
	return int(randInt(seed, int64(m), int64(M)))
}

// Int8 generates a pseudo-random int8 between the given range
func Int8(m, M int8) int8 {
	seed := lgc.Int8(seeds.Clock())
	return int8(randInt(seed, int64(m), int64(M)))
}

// Int16 generates a pseudo-random int16 between the given range
func Int16(m, M int16) int16 {
	seed := lgc.Int16(seeds.Clock())
	return int16(randInt(seed, int64(m), int64(M)))
}

// Int32 generates a pseudo-random int32 between the given range
func Int32(m, M int32) int32 {
	seed := lgc.Int32(seeds.Clock())
	return int32(randInt(seed, int64(m), int64(M)))
}

// Int64 generates a pseudo-random int64 between the given range
func Int64(m, M int64) int64 {
	seed := lgc.Int64(seeds.Clock())
	return randInt(seed, m, M)
}

func randInt(randomSeed uint64, m, M int64) int64 {
	//if uint64(M-m) == math.MaxUint64 {
	//	return int64(randomSeed)
	//}
	o := int64(randomSeed%uint64(M-m+1)) + m
	return o
}

// Uint generates a pseudo-random uint between the given range
func Uint(m, M uint) uint {
	seed := lgc.Uint(seeds.Clock())
	return uint(randUint(seed, uint64(m), uint64(M)))
}

// Uint8 generates a pseudo-random uint8 between the given range
func Uint8(m, M uint8) uint8 {
	seed := lgc.Uint8(seeds.Clock())
	return uint8(randUint(seed, uint64(m), uint64(M)))
}

// Uint16 generates a pseudo-random uint16 between the given range
func Uint16(m, M uint16) uint16 {
	seed := lgc.Uint16(seeds.Clock())
	return uint16(randUint(seed, uint64(m), uint64(M)))
}

// Uint32 generates a pseudo-random uint32 between the given range
func Uint32(m, M uint32) uint32 {
	seed := lgc.Uint32(seeds.Clock())
	return uint32(randUint(seed, uint64(m), uint64(M)))
}

// Uint64 generates a pseudo-random uint64 between the given range
func Uint64(m, M uint64) uint64 {
	seed := lgc.Uint64(seeds.Clock())
	return randUint(seed, m, M)
}

func randUint(randomSeed, m, M uint64) uint64 {
	//if M-m == uints.Max64 {
	//	return randomSeed
	//}
	return randomSeed%(M-m+1) + m
}
