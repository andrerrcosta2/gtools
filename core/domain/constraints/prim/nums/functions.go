// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package nums

import (
	"errors"
	"fmt"
	"github.com/andrerrcosta2/gtools/core/domain/constraints/prim/nums/floats"
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
	aFloat, aErr := floats.ToFloat64(a)
	if aErr != nil {
		return false, fmt.Errorf("first argument is not a number: %w", aErr)
	}

	bFloat, bErr := floats.ToFloat64(b)
	if bErr != nil {
		return false, fmt.Errorf("second argument is not a number: %w", bErr)
	}

	return aFloat < bFloat, nil
}

func Greater[T Real](a, b T) bool {
	return a > b
}

func TryGreater(a, b any) (bool, error) {
	aFloat, aErr := floats.ToFloat64(a)
	if aErr != nil {
		return false, fmt.Errorf("first argument is not a number: %w", aErr)
	}

	bFloat, bErr := floats.ToFloat64(b)
	if bErr != nil {
		return false, fmt.Errorf("second argument is not a number: %w", bErr)
	}

	return aFloat > bFloat, nil
}
