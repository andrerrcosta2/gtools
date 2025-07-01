// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package fuzz

import (
	"github.com/andrerrcosta2/gtools/core/seeders/random"
	"math/rand"
)

func valuesSet() []interface{} {
	return []interface{}{
		true, rand.Intn(10000), int8(rand.Intn(255)), int16(rand.Intn(10000)), int32(rand.Intn(10000)),
		int64(rand.Intn(10000)), uint(rand.Intn(10000)), uint8(rand.Intn(255)), uint16(rand.Intn(10000)),
		uint32(rand.Intn(10000)), uint64(rand.Intn(10000)), uintptr(rand.Intn(10000)), rand.Float32(),
		rand.Float64(), complex(rand.Float32(), rand.Float32()), complex(rand.Float64(), rand.Float64()),
		random.SingleOf[string](),
	}
}
