// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package lgc

import (
	"github.com/andrerrcosta2/gtools/core/seeders/random/internal/mod"
	"github.com/andrerrcosta2/gtools/core/seeders/random/internal/prng/inc"
	"github.com/andrerrcosta2/gtools/core/seeders/random/internal/prng/mult"
)

func Float32(s1, s2 uint64) (uint64, uint64) {
	ss1 := (mult.Uint64*s1 + inc.Uint64) & mod.Uint64
	ss2 := (mult.Uint64*s2 + inc.Uint64) & mod.Uint64
	return ss1, ss2
}

func Float64(s1, s2 uint64) (uint64, uint64) {
	ss1 := (mult.Uint64*s1 + inc.Uint64) & mod.Uint64
	ss2 := (mult.Uint64*s2 + inc.Uint64) & mod.Uint64
	return ss1, ss2
}

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
