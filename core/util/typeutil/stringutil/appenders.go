// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package stringutil

import (
	"strconv"
	"unsafe"

	"github.com/andrerrcosta2/gtools/core/domain/constraints/prim/nums"
)

// AppendNum appends a numeric value to a prefix string.
func AppendNum[N nums.Any](prefix string, v N) string {
	switch x := any(v).(type) {
	case int:
		return AppendInt(prefix, x)
	case int8:
		return AppendInt8(prefix, x)
	case int16:
		return AppendInt16(prefix, x)
	case int32:
		return AppendInt32(prefix, x)
	case int64:
		return AppendInt64(prefix, x)

	case uint:
		return AppendUint(prefix, x)
	case uint8:
		return AppendUint8(prefix, x)
	case uint16:
		return AppendUint16(prefix, x)
	case uint32:
		return AppendUint32(prefix, x)
	case uint64:
		return AppendUint64(prefix, x)
	case uintptr:
		return AppendUint64(prefix, uint64(x)) // uintptr → uint64

	case float32:
		return AppendFloat32(prefix, x)
	case float64:
		return AppendFloat64(prefix, x)

	case complex64:
		return AppendComplex64(prefix, x)
	case complex128:
		return AppendComplex128(prefix, x)

	default:
		// Should never happen due to nums.Any constraint
		panic("unsupported numeric type")
	}
}

// PrependNum prepends a numeric value to a suffix string.
func PrependNum[N nums.Any](v N, suffix string) string {
	switch x := any(v).(type) {
	case int:
		return PrependInt(x, suffix)
	case int8:
		return PrependInt8(x, suffix)
	case int16:
		return PrependInt16(x, suffix)
	case int32:
		return PrependInt32(x, suffix)
	case int64:
		return PrependInt64(x, suffix)

	case uint:
		return PrependUint(x, suffix)
	case uint8:
		return PrependUint8(x, suffix)
	case uint16:
		return PrependUint16(x, suffix)
	case uint32:
		return PrependUint32(x, suffix)
	case uint64:
		return PrependUint64(x, suffix)
	case uintptr:
		return PrependUint64(uint64(x), suffix)

	case float32:
		return PrependFloat32(x, suffix)
	case float64:
		return PrependFloat64(x, suffix)

	case complex64:
		return PrependComplex64(x, suffix)
	case complex128:
		return PrependComplex128(x, suffix)

	default:
		panic("unsupported numeric type")
	}
}

// AppendIntUnsafe avoids the memory copy of the final string.
func AppendIntUnsafe(prefix string, i int) string {
	b := make([]byte, 0, len(prefix)+20)
	b = append(b, prefix...)
	b = strconv.AppendInt(b, int64(i), 10)

	// In Go 1.20+, this is the preferred way to do zero-copy
	return unsafe.String(unsafe.SliceData(b), len(b))
}

// ---- Signed Integers ----

func AppendInt8(prefix string, v int8) string {
	b := make([]byte, 0, len(prefix)+4)
	b = append(b, prefix...)
	b = strconv.AppendInt(b, int64(v), 10)
	return string(b)
}

func PrependInt8(v int8, suffix string) string {
	b := make([]byte, 0, len(suffix)+4)
	b = strconv.AppendInt(b, int64(v), 10)
	b = append(b, suffix...)
	return string(b)
}

func AppendInt16(prefix string, v int16) string {
	b := make([]byte, 0, len(prefix)+6)
	b = append(b, prefix...)
	b = strconv.AppendInt(b, int64(v), 10)
	return string(b)
}

func PrependInt16(v int16, suffix string) string {
	b := make([]byte, 0, len(suffix)+6)
	b = strconv.AppendInt(b, int64(v), 10)
	b = append(b, suffix...)
	return string(b)
}

func AppendInt32(prefix string, v int32) string {
	b := make([]byte, 0, len(prefix)+11)
	b = append(b, prefix...)
	b = strconv.AppendInt(b, int64(v), 10)
	return string(b)
}

func PrependInt32(v int32, suffix string) string {
	b := make([]byte, 0, len(suffix)+11)
	b = strconv.AppendInt(b, int64(v), 10)
	b = append(b, suffix...)
	return string(b)
}

func AppendInt64(prefix string, v int64) string {
	b := make([]byte, 0, len(prefix)+20)
	b = append(b, prefix...)
	b = strconv.AppendInt(b, v, 10)
	return string(b)
}

func PrependInt64(v int64, suffix string) string {
	b := make([]byte, 0, len(suffix)+20)
	b = strconv.AppendInt(b, v, 10)
	b = append(b, suffix...)
	return string(b)
}

func AppendInt(prefix string, v int) string {
	return AppendInt64(prefix, int64(v))
}

func PrependInt(v int, suffix string) string {
	return PrependInt64(int64(v), suffix)
}

// ---- Unsigned Integers ----

func AppendUint8(prefix string, v uint8) string {
	b := make([]byte, 0, len(prefix)+3)
	b = append(b, prefix...)
	b = strconv.AppendUint(b, uint64(v), 10)
	return string(b)
}

func PrependUint8(v uint8, suffix string) string {
	b := make([]byte, 0, len(suffix)+3)
	b = strconv.AppendUint(b, uint64(v), 10)
	b = append(b, suffix...)
	return string(b)
}

func AppendUint16(prefix string, v uint16) string {
	b := make([]byte, 0, len(prefix)+5)
	b = append(b, prefix...)
	b = strconv.AppendUint(b, uint64(v), 10)
	return string(b)
}

func PrependUint16(v uint16, suffix string) string {
	b := make([]byte, 0, len(suffix)+5)
	b = strconv.AppendUint(b, uint64(v), 10)
	b = append(b, suffix...)
	return string(b)
}

func AppendUint32(prefix string, v uint32) string {
	b := make([]byte, 0, len(prefix)+10)
	b = append(b, prefix...)
	b = strconv.AppendUint(b, uint64(v), 10)
	return string(b)
}

func PrependUint32(v uint32, suffix string) string {
	b := make([]byte, 0, len(suffix)+10)
	b = strconv.AppendUint(b, uint64(v), 10)
	b = append(b, suffix...)
	return string(b)
}

func AppendUint64(prefix string, v uint64) string {
	b := make([]byte, 0, len(prefix)+20)
	b = append(b, prefix...)
	b = strconv.AppendUint(b, v, 10)
	return string(b)
}

func PrependUint64(v uint64, suffix string) string {
	b := make([]byte, 0, len(suffix)+20)
	b = strconv.AppendUint(b, v, 10)
	b = append(b, suffix...)
	return string(b)
}

func AppendUint(prefix string, v uint) string {
	return AppendUint64(prefix, uint64(v))
}

func PrependUint(v uint, suffix string) string {
	return PrependUint64(uint64(v), suffix)
}

// ---- Floats ----

const (
	maxFloat32StrLen = 47
	maxFloat64StrLen = 325
)

func AppendFloat32(prefix string, v float32) string {
	b := make([]byte, 0, len(prefix)+maxFloat32StrLen)
	b = append(b, prefix...)
	b = strconv.AppendFloat(b, float64(v), 'G', -1, 32)
	return string(b)
}

func PrependFloat32(v float32, suffix string) string {
	b := make([]byte, 0, len(suffix)+maxFloat32StrLen)
	b = strconv.AppendFloat(b, float64(v), 'G', -1, 32)
	b = append(b, suffix...)
	return string(b)
}

func AppendFloat64(prefix string, v float64) string {
	b := make([]byte, 0, len(prefix)+maxFloat64StrLen)
	b = append(b, prefix...)
	b = strconv.AppendFloat(b, v, 'G', -1, 64)
	return string(b)
}

func PrependFloat64(v float64, suffix string) string {
	b := make([]byte, 0, len(suffix)+maxFloat64StrLen)
	b = strconv.AppendFloat(b, v, 'G', -1, 64)
	b = append(b, suffix...)
	return string(b)
}

func AppendFloat(prefix string, v float64) string {
	return AppendFloat64(prefix, v)
}

func PrependFloat(v float64, suffix string) string {
	return PrependFloat64(v, suffix)
}

// ---- Complex Numbers ----

const (
	maxComplex64StrLen  = 100 // (a+b*i) with float32 extremes
	maxComplex128StrLen = 700 // conservative for float64 extremes
)

func AppendComplex64(prefix string, v complex64) string {
	b := make([]byte, 0, len(prefix)+maxComplex64StrLen)
	b = append(b, prefix...)
	b = strconv.AppendQuote(b, strconv.FormatComplex(complex128(v), 'G', -1, 64))
	// Remove quotes added by FormatComplex via AppendQuote
	if len(b) >= 2 && b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	return string(b)
}

func PrependComplex64(v complex64, suffix string) string {
	s := strconv.FormatComplex(complex128(v), 'G', -1, 64)
	b := make([]byte, 0, len(s)+len(suffix))
	b = append(b, s...)
	b = append(b, suffix...)
	return string(b)
}

func AppendComplex128(prefix string, v complex128) string {
	s := strconv.FormatComplex(v, 'G', -1, 128)
	b := make([]byte, 0, len(prefix)+len(s))
	b = append(b, prefix...)
	b = append(b, s...)
	return string(b)
}

func PrependComplex128(v complex128, suffix string) string {
	s := strconv.FormatComplex(v, 'G', -1, 128)
	b := make([]byte, 0, len(s)+len(suffix))
	b = append(b, s...)
	b = append(b, suffix...)
	return string(b)
}
