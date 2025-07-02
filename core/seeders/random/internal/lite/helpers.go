// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package lite

import (
	"fmt"
	"github.com/andrerrcosta2/gtools/core/format/fmx"
	"github.com/andrerrcosta2/gtools/core/seeders/random/internal/prng"
	"github.com/andrerrcosta2/gtools/core/seeders/random/internal/reflectutils"
	"github.com/andrerrcosta2/gtools/core/seeders/random/internal/reflectutils/reflectrand"
	"github.com/andrerrcosta2/gtools/core/util/typeutil/charsets"
	"github.com/google/uuid"
	"math/rand"
	"reflect"
	"time"
	"unicode"
	"unicode/utf8"
)

// RandAny generates a random value of any type except structs.
func RandAny() any {
	kind := RandGiven(reflectutils.InterfaceKinds)
	switch kind {
	case reflect.Bool:
		return prng.Bool()
	case reflect.Int:
		return prng.Int()
	case reflect.Int8:
		return prng.Int8()
	case reflect.Int16:
		return prng.Int16()
	case reflect.Int32:
		return prng.Int32()
	case reflect.Int64:
		return prng.Int64()
	case reflect.Uint, reflect.Uintptr:
		return prng.Uint()
	case reflect.Uint8:
		return prng.Uint8()
	case reflect.Uint16:
		return prng.Uint16()
	case reflect.Uint32:
		return prng.Uint32()
	case reflect.Uint64:
		return prng.Uint64()
	case reflect.Float32:
		return prng.Float32()
	case reflect.Float64:
		return prng.Float64()
	case reflect.Complex64:
		return prng.Complex64()
	case reflect.Complex128:
		return prng.Complex128()
	case reflect.Array:
		return RandArray()
	case reflect.Chan:
		return RandChan()
	case reflect.Func:
		return RandFunc()
	case reflect.Map:
		return RandMap()
	case reflect.Ptr:
		return RandPointer()
	case reflect.Slice:
		return RandSlice()
	case reflect.String:
		// TODO: This method lacks constraints control. Finish validation library
		// TODO: Add its interface to the core package
		return RandString(prng.Int(1, 50), charsets.AlphaNumeric)
	default:
		// Unreachable
		panic(fmx.Sprintf("RandCmp: unsupported comparable kind: '%v'\n", kind))
		return nil
	}
}

// RandArray generates a random array of a random type
func RandArray() any {
	return reflectrand.ArrayOf(reflectrand.Array()).Interface()
}

// RandArrayOf generates a random array of the given type
func RandArrayOf(t reflect.Type) any {
	return reflectrand.ArrayOf(t).Interface()
}

// RandChan generates a random channel of a random type
func RandChan() any {
	return reflectrand.ChanOf(reflectrand.Chan()).Interface()
}

// RandChanOf generates a random channel from a given type
func RandChanOf(t reflect.Type) any {
	return reflectrand.ChanOf(t).Interface()
}

// RandCmp generates a random value of a random comparable type
func RandCmp() any {
	kind := reflectrand.CmpKind()
	reflectrand.Cmp()
	switch kind {
	case reflect.Bool:
		return prng.Bool()
	case reflect.Int:
		return prng.Int()
	case reflect.Int8:
		return prng.Int8()
	case reflect.Int16:
		return prng.Int16()
	case reflect.Int32:
		return prng.Int32()
	case reflect.Int64:
		return prng.Int64()
	case reflect.Uint, reflect.Uintptr:
		return prng.Uint()
	case reflect.Uint8:
		return prng.Uint8()
	case reflect.Uint16:
		return prng.Uint16()
	case reflect.Uint32:
		return prng.Uint32()
	case reflect.Uint64:
		return prng.Uint64()
	case reflect.String:
		return RandString(prng.Int(1, 50), charsets.AlphaNumeric)
	case reflect.Float32:
		return prng.Float32()
	case reflect.Float64:
		return prng.Float64()
	case reflect.Complex64:
		return prng.Complex64()
	case reflect.Complex128:
		return prng.Complex128()
	case reflect.Array:
		return RandArrayOf(reflectrand.CmpArray())
	case reflect.Ptr:
		return RandPointer()
	default:
		// Unreachable
		panic(fmx.Sprintf("RandCmp: unsupported comparable kind: '%v'\n", kind))
		return nil
	}
}

// RandFunc generates a random function
func RandFunc() any {
	return reflectrand.FuncOf(reflectrand.Func()).Interface()
}

// RandFuncOf generates a random function from a given type
func RandFuncOf(t reflect.Type) any {
	return reflectrand.FuncOf(t).Interface()
}

// RandGiven returns a random value from a given slice
func RandGiven[S ~[]E, E any](s S) E {
	return s[prng.Int(0, len(s)-1)]
}

// RandMap generates a random map from random types
func RandMap() any {
	return reflectrand.MapOf(reflectrand.Map()).Interface()
}

// RandMapOf generates a random map from a given type
func RandMapOf(t reflect.Type) any {
	return reflectrand.MapOf(t).Interface()
}

// RandOf generates a single random value
func RandOf(t reflect.Type) any {
	if t == nil {
		return nil
	}
	kind := t.Kind()
	switch kind {
	case reflect.Bool:
		return prng.Bool()
	case reflect.Int:
		return prng.Int()
	case reflect.Int8:
		return prng.Int8()
	case reflect.Int16:
		return prng.Int16()
	case reflect.Int32:
		return prng.Int32()
	case reflect.Int64:
		return prng.Int64()
	case reflect.Uint, reflect.Uintptr:
		return prng.Uint()
	case reflect.Uint8:
		return prng.Uint8()
	case reflect.Uint16:
		return prng.Uint16()
	case reflect.Uint32:
		return prng.Uint32()
	case reflect.Uint64:
		return prng.Uint64()
	case reflect.Float32:
		return prng.Float32()
	case reflect.Float64:
		return prng.Float64()
	case reflect.Complex64:
		return prng.Complex64()
	case reflect.Complex128:
		return prng.Complex128()
	case reflect.Array:
		return RandArrayOf(t)
	case reflect.Chan:
		return RandChanOf(t)
	case reflect.Func:
		return RandFuncOf(t)
	case reflect.Map:
		return RandMapOf(t)
	case reflect.Ptr:
		return RandPointerOf(t)
	case reflect.Slice:
		return RandSliceOf(t)
	case reflect.Struct:
		return RandStructOf(t)
	case reflect.Interface:
		if t.NumMethod() == 0 {
			return RandAny()
		}
		return nil

	case reflect.String:
		// TODO: This method lacks constraints control. Finish validation library
		// TODO: Add its interface to the core package
		return RandString(prng.Int(1, 50), charsets.AlphaNumeric)
	default:
		panic(fmt.Sprintf("gtools:random: nil or invalid type: '%T' passed to be randomized", t))
	}
}

// RandPointer generates a random pointer to a random value from a random type
func RandPointer() any {
	return reflectrand.PointerOf(reflectrand.Pointer()).Interface()
}

// RandPointerOf generates a random pointer to a random value of the given type
func RandPointerOf(t reflect.Type) any {
	return reflectrand.PointerOf(t).Interface()
}

// RandRune generates a random Unicode rune that is valid and graphic.
// It continues to generate runes until one satisfies these conditions.
func RandRune() rune {
	for {
		// Generate a random rune within the full Unicode range (up to 0x10FFFF)
		r := rune(prng.Uint(0, 0x10FFFF))
		// Check if the rune is valid and a graphic character (i.e. not a control character like '\n', '\t', etc.)
		if utf8.ValidRune(r) && unicode.IsGraphic(r) {
			return r
		}
	}
}

// RandSlice generates a random slice of random value from a random type
func RandSlice() any {
	return reflectrand.SliceOf(reflectrand.Slice()).Interface()
}

// RandSliceOf generates a random slice of random value from a given type
func RandSliceOf(t reflect.Type) any {
	return reflectrand.SliceOf(t).Interface()
}

func RandString(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

// RandStructOf generates a new instance of the given struct type with random values for its fields.
// It uses reflection to dynamically create an instance and set the fields.
func RandStructOf(t reflect.Type) any {
	return reflectrand.StructOf(t).Interface()
}

func RandTimestamp(from time.Time, diff time.Duration) time.Time {
	randomDuration := time.Duration(rand.Int63n(int64(diff)))
	return from.Add(randomDuration)
}

func RandUuid() string {
	return uuid.New().String()
}
