// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package prng

import (
	"github.com/andrerrcosta2/gtools/core/seeders/random/internal/prng/lgc"
	"github.com/andrerrcosta2/gtools/core/seeders/random/internal/seeds"
	"math"
)

func Bool() bool {
	return seeds.Clock()&1 == 1
}

func Bytes(minMax ...int) []byte {
	// Generate the random length of the byte slice
	randomBytes := make([]byte, Int(minMax...))

	// Fill the byte slice with random values
	for i := range randomBytes {
		rand := randInt(lgc.Int8(seeds.Clock()), 0, 255)
		randomBytes[i] = byte(rand)
	}

	return randomBytes
}

func Int(minMax ...int) int {
	var m, M int64
	switch len(minMax) {
	case 0:
		m, M = math.MinInt, math.MaxInt
	case 1:
		m, M = int64(minMax[0]), math.MaxInt
	default:
		m, M = int64(minMax[0]), int64(minMax[1])
	}
	if m == M {
		return int(m)
	}
	if m > M {
		m, M = M, m
	}
	clock := seeds.Clock()
	seed := lgc.Int(clock)
	return int(randInt(seed, m, M))
}

func Int8(minMax ...int8) int8 {
	var m, M int64
	switch len(minMax) {
	case 0:
		m, M = math.MinInt8, math.MaxInt8
	case 1:
		m, M = int64(minMax[0]), math.MaxInt8
	default:
		m, M = int64(minMax[0]), int64(minMax[1])
	}
	if m == M {
		return int8(m)
	}
	if m > M {
		m, M = M, m
	}
	clock := seeds.Clock()
	seed := lgc.Int8(clock)
	return int8(randInt(seed, m, M))
}

func Int16(minMax ...int16) int16 {
	var m, M int64
	switch len(minMax) {
	case 0:
		m, M = math.MinInt16, math.MaxInt16
	case 1:
		m, M = int64(minMax[0]), math.MaxInt16
	default:
		m, M = int64(minMax[0]), int64(minMax[1])
	}
	if m == M {
		return int16(m)
	}
	if m > M {
		m, M = M, m
	}
	seed := lgc.Int16(seeds.Clock())
	return int16(randInt(seed, m, M))
}

func Int32(minMax ...int32) int32 {
	var m, M int64
	switch len(minMax) {
	case 0:
		m, M = math.MinInt32, math.MaxInt32
	case 1:
		m, M = int64(minMax[0]), math.MaxInt32
	default:
		m, M = int64(minMax[0]), int64(minMax[1])
	}
	if m == M {
		return int32(m)
	}
	if m > M {
		m, M = M, m
	}
	seed := lgc.Int32(seeds.Clock())
	return int32(randInt(seed, m, M))
}

func Int64(minMax ...int64) int64 {
	var m, M int64
	switch len(minMax) {
	case 0:
		m, M = math.MinInt64, math.MaxInt64
	case 1:
		m, M = minMax[0], math.MaxInt64
	default:
		m, M = minMax[0], minMax[1]
	}
	if m == M {
		return m
	}
	if m > M {
		m, M = M, m
	}
	seed := lgc.Int64(seeds.Clock())
	return randInt(seed, m, M)
}

func randInt(randomSeed uint64, m, M int64) int64 {
	if uint64(M-m) == math.MaxUint64 {
		return int64(randomSeed)
	}
	o := int64(randomSeed%uint64(M-m+1)) + m
	return o
}

func Uint(minMax ...uint) uint {
	var m, M uint64
	switch len(minMax) {
	case 0:
		m, M = 0, math.MaxUint
	case 1:
		m, M = uint64(minMax[0]), math.MaxUint
	default:
		m, M = uint64(minMax[0]), uint64(minMax[1])
	}
	if m == M {
		return uint(m)
	}
	if m > M {
		m, M = M, m
	}
	clock := seeds.Clock()
	seed := lgc.Uint(clock)
	return uint(randUint(seed, m, M))
}

func Uint8(minMax ...uint8) uint8 {
	var m, M uint64
	switch len(minMax) {
	case 0:
		m, M = 0, math.MaxUint8
	case 1:
		m, M = uint64(minMax[0]), math.MaxUint8
	default:
		m, M = uint64(minMax[0]), uint64(minMax[1])
	}
	if m == M {
		return uint8(m)
	}
	if m > M {
		m, M = M, m
	}
	seed := lgc.Uint8(seeds.Clock())
	return uint8(randUint(seed, m, M))
}

func Uint16(minMax ...uint16) uint16 {
	var m, M uint64
	switch len(minMax) {
	case 0:
		m, M = 0, math.MaxUint16
	case 1:
		m, M = uint64(minMax[0]), math.MaxUint16
	default:
		m, M = uint64(minMax[0]), uint64(minMax[1])
	}
	if m == M {
		return uint16(m)
	}
	if m > M {
		m, M = M, m
	}
	clock := seeds.Clock()
	seed := lgc.Uint16(clock)
	return uint16(randUint(seed, m, M))
}

func Uint32(minMax ...uint32) uint32 {
	var m, M uint64
	switch len(minMax) {
	case 0:
		m, M = 0, math.MaxUint32
	case 1:
		m, M = uint64(minMax[0]), math.MaxUint32
	default:
		m, M = uint64(minMax[0]), uint64(minMax[1])
	}
	if m == M {
		return uint32(m)
	}
	if m > M {
		m, M = M, m
	}
	seed := lgc.Uint32(seeds.Clock())
	return uint32(randUint(seed, m, M))
}

func Uint64(minMax ...uint64) uint64 {
	var m, M uint64
	switch len(minMax) {
	case 0:
		m, M = 0, math.MaxUint64
	case 1:
		m, M = minMax[0], math.MaxUint64
	default:
		m, M = minMax[0], minMax[1]
	}
	if m == M {
		return m
	}
	if m > M {
		m, M = M, m
	}
	seed := lgc.Uint64(seeds.Clock())
	return randUint(seed, m, M)
}

func randUint(randomSeed uint64, m, M uint64) uint64 {
	if M-m == math.MaxUint64 {
		return randomSeed
	}
	return randomSeed%(M-m+1) + m
}

func Float32(minMax ...float32) float32 {
	var m, M float32
	switch len(minMax) {
	case 0:
		m, M = -math.MaxFloat32, math.MaxFloat32
	case 1:
		m, M = minMax[0], math.MaxFloat32
	default:
		m, M = minMax[0], minMax[1]
	}
	if m == M {
		return m
	}
	if m > M {
		m, M = M, m
	}
	exp, mnt := lgc.Float32(seeds.Clock(), seeds.Clock())
	return randF32(exp, mnt, m, M)
}

func randF32(exp float64, mnt uint64, m, M float32) float32 {
	// Extract integer parts of m and M
	Mi := float32(math.Floor(float64(M)))
	mi := float32(math.Floor(float64(m)))

	// Int range
	rng := Mi - mi
	//fmt.Printf("Before range of %f\nValue is: %f\n", rng, exp)
	//fmt.Printf("Exp > math.MaxFloat32? %t\n", float32(exp) > float32(math.MaxFloat32))
	factor := float32(math.Mod(exp, math.MaxFloat32)) / float32(math.MaxFloat32)
	//fmt.Printf("factor: %f\n", factor)
	exn := factor*rng + mi
	//fmt.Printf("After exn: %f\n\n", exn)

	// Fractional part
	Mf := M - Mi
	mnn := float32(mnt) / float32(math.MaxUint64) * Mf

	return exn + mnn
}

//func uint64ToFraction(num uint64) float64 {
//	// Calculate the number of digits in the number
//	digits := math.Floor(math.Log10(float64(num)) + 1)
//	// Scale the number to create a fractional value
//	return float64(num) / math.Pow(10, digits)
//}

func Float64(minMax ...float64) float64 {
	var m, M float64
	switch len(minMax) {
	case 0:
		m, M = -math.MaxFloat64, math.MaxFloat64
	case 1:
		m, M = minMax[0], math.MaxFloat64
	default:
		m, M = minMax[0], minMax[1]
	}
	if m == M {
		return m
	}
	if m > M {
		m, M = M, m
	}
	exp, mnt := lgc.Float64(seeds.Clock(), seeds.Clock())
	return randF64(exp, mnt, m, M)
}

func randF64(exp float64, mnt uint64, m, M float64) float64 {
	// Extract integer parts of m and M
	Mi := math.Floor(M)
	mi := math.Floor(m)

	// Int range
	rngi := Mi - mi + 1
	exn := exp/math.MaxFloat64*rngi + mi

	// Fractional part
	Mf := M - Mi
	mnn := float64(mnt) / 4503599627370496.0 * Mf // 4503599627370496 is 2^52

	return exn + mnn
}

func Complex64(realMinMaxImagMinMax ...float32) complex64 {
	minMax := realMinMaxImagMinMax
	var rm, rM, im, iM float32
	switch len(realMinMaxImagMinMax) {
	case 0:
		rm, rM, im, iM = -math.MaxFloat32, math.MaxFloat32, -math.MaxFloat32, math.MaxFloat32
	case 1:
		rm, rM, im, iM = minMax[0], math.MaxFloat32, -math.MaxFloat32, math.MaxFloat32
	case 2:
		rm, rM, im, iM = minMax[0], minMax[1], -math.MaxFloat32, math.MaxFloat32
	case 3:
		rm, rM, im, iM = minMax[0], minMax[1], minMax[2], math.MaxFloat32
	default:
		rm, rM, im, iM = minMax[0], minMax[1], minMax[2], minMax[3]
	}

	if rm > rM {
		rm, rM = rM, rm
	}
	if im > iM {
		im, iM = iM, im
	}
	res, rms := lgc.Float32(seeds.Clock(), seeds.Clock())
	ies, ims := lgc.Float32(seeds.Clock(), seeds.Clock())

	var rp, ip float32
	if rm == rM {
		rp = rm
	} else {
		rp = randF32(res, rms, rm, rM)
	}
	if im == iM {
		ip = im
	} else {
		ip = randF32(ies, ims, im, iM)
	}
	return complex(rp, ip)
}

func Complex128(realMinMaxImagMinMax ...float64) complex128 {
	minMax := realMinMaxImagMinMax
	var rm, rM, im, iM float64
	switch len(realMinMaxImagMinMax) {
	case 0:
		rm, rM, im, iM = -math.MaxFloat64, math.MaxFloat64, -math.MaxFloat64, math.MaxFloat64
	case 1:
		rm, rM, im, iM = minMax[0], math.MaxFloat64, -math.MaxFloat64, math.MaxFloat64
	case 2:
		rm, rM, im, iM = minMax[0], minMax[1], -math.MaxFloat64, math.MaxFloat64
	case 3:
		rm, rM, im, iM = minMax[0], minMax[1], minMax[2], math.MaxFloat64
	default:
		rm, rM, im, iM = minMax[0], minMax[1], minMax[2], minMax[3]
	}
	if rm > rM {
		rm, rM = rM, rm
	}
	if im > iM {
		im, iM = iM, im
	}
	res, rms := lgc.Float64(seeds.Clock(), seeds.Clock())
	ies, ims := lgc.Float64(seeds.Clock(), seeds.Clock())

	var rp, ip float64
	if rm == rM {
		rp = rm
	} else {
		rp = randF64(res, rms, rm, rM)
	}
	if im == iM {
		ip = im
	} else {
		ip = randF64(ies, ims, im, iM)
	}
	return complex(rp, ip)
}
