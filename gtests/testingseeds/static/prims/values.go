// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package prims

import "math/rand"

func zero() []interface{} {
	return []interface{}{
		false, 0, int8(0), int16(0), int32(0), int64(0), uint(0), uint8(0), uint16(0), uint32(0),
		uint64(0), uintptr(0), float32(0), float64(0), complex64(0), complex128(0), "",
	}
}

func random() []interface{} {
	return []interface{}{
		true, rand.Intn(10000), int8(rand.Intn(255)), int16(rand.Intn(10000)), int32(rand.Intn(10000)),
		int64(rand.Intn(10000)), uint(rand.Intn(10000)), uint8(rand.Intn(255)), uint16(rand.Intn(10000)),
		uint32(rand.Intn(10000)), uint64(rand.Intn(10000)), uintptr(rand.Intn(10000)), rand.Float32(),
		rand.Float64(), complex(rand.Float32(), rand.Float32()), complex(rand.Float64(), rand.Float64()),
		randomString(10),
	}
}

func ofValues() []interface{} {
	return append(zero(), random()...)
}
