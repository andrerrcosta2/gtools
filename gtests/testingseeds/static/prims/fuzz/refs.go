// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package fuzz

import (
	"github.com/andrerrcosta2/gtools/core/seeders/random"
	"math/rand"
)

func referencesSet() []any {
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
	p17 := random.SingleOf[string]()
	return []any{&p1, &p2, &p3, &p4, &p5, &p6, &p7, &p8, &p9, &p10, &p11, &p12, &p13, &p14, &p15, &p16, &p17}
}
