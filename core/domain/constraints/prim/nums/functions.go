// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package nums

import (
	"errors"
	"fmt"
	"math"
)

func ToNumeric[T Any](value any) (T, error) {
	switch v := value.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr, float32, float64, complex64, complex128:
		var result T
		if value, ok := v.(T); ok {
			result = value
		} else {
			return result, errors.New("couldn't convert to T\n")
		}
		return result, nil

	default:
		var zero T
		return zero, errors.New("not a numeric type\n")
	}
}

func IsNaN(v any) bool {
	switch x := v.(type) {
	case float32:
		return math.IsNaN(float64(x))
	case float64:
		return math.IsNaN(x)
	default:
		return false
	}
}

func IsReal(value any) bool {
	switch value.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr, float32, float64:
		return true
	default:
		return false
	}
}

func Less[T Real](a, b T) bool {
	return a < b
}

func TryLess(a, b any) (bool, error) {
	aFloat, aErr := toFloat64(a)
	if aErr != nil {
		return false, fmt.Errorf("first argument is not a number: %w", aErr)
	}

	bFloat, bErr := toFloat64(b)
	if bErr != nil {
		return false, fmt.Errorf("second argument is not a number: %w", bErr)
	}

	return aFloat < bFloat, nil
}

func Greater[T Real](a, b T) bool {
	return a > b
}

// ToFloat32 tries to convert any number to float32
func toFloat32(val any) (float32, error) {
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
func toFloat64(val any) (float64, error) {
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

func TryGreater(a, b any) (bool, error) {
	aFloat, aErr := toFloat64(a)
	if aErr != nil {
		return false, fmt.Errorf("first argument is not a number: %w", aErr)
	}

	bFloat, bErr := toFloat64(b)
	if bErr != nil {
		return false, fmt.Errorf("second argument is not a number: %w", bErr)
	}

	return aFloat > bFloat, nil
}
