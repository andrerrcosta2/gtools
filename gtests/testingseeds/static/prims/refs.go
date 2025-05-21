// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package prims

import (
	"math/rand"
	"slices"
	"time"
)

func zeroAsRef() []interface{} {
	var p1 bool
	var p2 int
	var p3 int8
	var p4 int16
	var p5 int32
	var p6 int64
	var p7 uint
	var p8 uint8
	var p9 uint16
	var p10 uint32
	var p11 uint64
	var p12 uintptr
	var p13 float32
	var p14 float64
	var p15 complex64
	var p16 complex128
	var p17 string
	return []interface{}{&p1, &p2, &p3, &p4, &p5, &p6, &p7, &p8, &p9, &p10, &p11, &p12, &p13, &p14, &p15, &p16, &p17}
}

func randomAsRef() []any {
	p1 := rand.Intn(2) == 1
	p2 := rand.Intn(1000000)
	p3 := int8(rand.Intn(255))
	p4 := int16(rand.Intn(1000000))
	p5 := int32(rand.Intn(1000000))
	p6 := int64(rand.Intn(1000000))
	p7 := uint(rand.Intn(1000000))
	p8 := uint8(rand.Intn(255))
	p9 := uint16(rand.Intn(1000000))
	p10 := uint32(rand.Intn(1000000))
	p11 := uint64(rand.Intn(1000000))
	p12 := uintptr(rand.Intn(1000000))
	p13 := rand.Float32()
	p14 := rand.Float64()
	p15 := complex(rand.Float32(), rand.Float32())
	p16 := complex(rand.Float64(), rand.Float64())
	p17 := randomString(10)
	return []any{&p1, &p2, &p3, &p4, &p5, &p6, &p7, &p8, &p9, &p10, &p11, &p12, &p13, &p14, &p15, &p16, &p17}
}

func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func ofRefs() []interface{} {
	return slices.Concat(zeroAsRef(), randomAsRef())
}
