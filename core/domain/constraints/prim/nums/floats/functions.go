// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package floats

import (
	"encoding/binary"
	"fmt"
	"math"
)

// Abs32 returns the absolute value for a given float32
func Abs32(x float32) float32 {
	return math.Float32frombits(math.Float32bits(x) &^ (1 << 31))
}

// Abs64 returns the absolute value for a given float64
func Abs64(x float64) float64 {
	return math.Float64frombits(math.Float64bits(x) &^ (1 << 63))
}

// FromBytes converts a byte slice to a float32 or float64 value.
// It uses the binary package to decode the value.
// It returns an error if the value is not a supported type.
func FromBytes(b []byte) (any, error) {
	switch len(b) {
	case 4:
		// 4 bytes for float32
		return math.Float32frombits(binary.BigEndian.Uint32(b)), nil
	case 8:
		// 8 bytes for float64
		return math.Float64frombits(binary.BigEndian.Uint64(b)), nil
	default:
		return nil, fmt.Errorf("invalid byte slice length %d", len(b))
	}
}

// Max32Fraction takes a float64 input, computes its absolute value as
// a fraction of Max32, and applies the original sign to return a value
// in the range [-1.0, 1.0].
func Max32Fraction(f float32) float32 {
	const scale = 1.0 / Max32
	return f * scale
}

// Max64Fraction takes a float64 input, computes its absolute value as
// a fraction of Max64, and applies the original sign to return a value
// in the range [-1.0, 1.0].
func Max64Fraction(f float64) float64 {
	const scale = 1.0 / Max64
	return f * scale
}

// NextAfter32 returns the next float32 after a given value
func NextAfter32(x float32) float32 {
	return math.Float32frombits(math.Float32bits(x) + 1)
}

// NextAfter64 returns the next float64 after a given value
func NextAfter64(x float64) float64 {
	return math.Float64frombits(math.Float64bits(x) + 1)
}

// Spacing32 calculate the spacing between two adjacent float32 values
//
//	Example:
//	fmt.Println(Spacing32(16_777_216)) // Output: 2
//	fmt.Println(Spacing32(33_554_432)) // Output: 4
func Spacing32(x float32) float32 {
	return NextAfter32(x) - x
}

// Spacing64 calculate the spacing between two adjacent float64 values
//
// Example
//
//	fmt.Println(Spacing32(16_777_216)) // Output: 2
//	fmt.Println(Spacing32(33_554_432)) // Output: 4
func Spacing64(x float64) float64 { return NextAfter64(x) - x }

// ToBytes converts a float32 or float64 value to its byte representation.
// It uses the binary package to encode the value.
// It returns an error if the value is not a supported type.
func ToBytes(value any) ([]byte, error) {
	switch v := value.(type) {
	case float32:
		// 4 bytes for float32
		b := make([]byte, 4)
		binary.BigEndian.PutUint32(b, math.Float32bits(v))
		return b, nil
	case float64:
		// 8 bytes for float64
		b := make([]byte, 8)
		binary.BigEndian.PutUint64(b, math.Float64bits(v))
		return b, nil
	default:
		return nil, fmt.Errorf("unsupported type %T", v)
	}
}

// ToFloat32 tries to convert any number to float32
func ToFloat32(val any) (float32, error) {
	switch v := val.(type) {
	case float64:
		return float32(v), nil
	case float32:
		return v, nil
	case int:
		return float32(v), nil
	case int8:
		return float32(v), nil
	case int16:
		return float32(v), nil
	case int32:
		return float32(v), nil
	case int64:
		return float32(v), nil
	case uint:
		return float32(v), nil
	case uint8:
		return float32(v), nil
	case uint16:
		return float32(v), nil
	case uint32:
		return float32(v), nil
	case uint64:
		return float32(v), nil
	default:
		return 0, fmt.Errorf("type not convertible to float32: %T", v)
	}
}

// ToFloat64 tris to convert any number to float64
func ToFloat64(val any) (float64, error) {
	switch v := val.(type) {
	case float64:
		return v, nil
	case float32:
		return float64(v), nil
	case int:
		return float64(v), nil
	case int8:
		return float64(v), nil
	case int16:
		return float64(v), nil
	case int32:
		return float64(v), nil
	case int64:
		return float64(v), nil
	case uint:
		return float64(v), nil
	case uint8:
		return float64(v), nil
	case uint16:
		return float64(v), nil
	case uint32:
		return float64(v), nil
	case uint64:
		return float64(v), nil
	default:
		return 0, fmt.Errorf("type not convertible to float64: %T", v)
	}
}

// ToFraction converts a float64 to its fraction representation [0, 1]
func ToFraction(f float64) float64 {
	const mantissaBits = 52
	const mask = (1 << mantissaBits) - 1
	const scale = 1.0 / float64(1<<mantissaBits)
	return float64(math.Float64bits(f)&mask) * scale
}
