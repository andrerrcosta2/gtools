// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package prim

import (
	"fmt"
)

// ToOrdered casts the provided value to the type T and returns it.
// T must be a primitive ordered type.
// If the value is not of type G, an error is returned.
func ToOrdered[T Ordered](value any) (T, error) {
	// Check if the value is of a primitive ordered type
	switch v := value.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr, float32, float64, string:
		// If the value is of a primitive ordered type, return it
		var result T
		if value, ok := any(v).(T); ok {
			result = value
		} else {
			// If the value is not of the correct type, return an error
			return result, fmt.Errorf("couldn't convert '%v' to Ordered\n", value)
		}
		return result, nil

	default:
		// If the value is not of a primitive ordered type, return an error
		var zero T
		return zero, fmt.Errorf("'%T' is not an ordered type\n", value)
	}
}

// ToHashable casts the provided value to the type T and returns it.
// T must be a primitive hashable type.
// If the value is not of type G, an error is returned.
func ToHashable[T Hashable](value any) (T, error) {
	switch v := value.(type) {
	// If the value is of a primitive hashable type, return it
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr, float32, float64, string:
		var result T
		if value, ok := any(v).(T); ok {
			result = value
		} else {
			return result, fmt.Errorf("couldn't convert '%v' to Hashable\n", value)
		}
		return result, nil
	default:
		// If the value is not of a primitive hashable type, return an error
		var zero T
		return zero, fmt.Errorf("'%T' is not a primitive hashable type\n", value)
	}
}

// ToPrimitive casts the provided value to the type T and returns it.
// T must be a primitive type.
// If the value is not of type G, an error is returned.
func ToPrimitive[T Any](value any) (T, error) {
	switch v := value.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr, float32, float64, string, bool, complex64, complex128:
		// If the value is of the correct type, return it
		var result T
		if value, ok := v.(T); ok {
			result = value
		} else {
			return result, fmt.Errorf("couldn't convert '%v' to primitive\n", value)
		}
		return result, nil

	default:
		// If the value is not of a primitive type, return an error
		var zero T
		return zero, fmt.Errorf("'%T' is not a primitive type\n", value)
	}
}

// ToComparable casts the provided value to the type T and returns it.
// T must be a primitive comparable type.
// If the value is not of type G, an error is returned.
func ToComparable[T Comparable](value any) (T, error) {
	switch v := value.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr, float32, float64, string, bool, complex64, complex128:
		// If the value is of the correct type, return it
		var result T
		if value, ok := v.(T); ok {
			result = value
		} else {
			// If the value is not of the correct type, return an error
			return result, fmt.Errorf("couldn't convert '%v' to comparable\n", value)
		}
		return result, nil

	default:
		// If the value is not of a primitive comparable type, return an error
		var zero T
		return zero, fmt.Errorf("'%T' is not a comparable type\n", value)
	}
}

func Less[T Ordered](a, b T) bool {
	return a < b
}

func TryLess(a, b any) (bool, error) {
	switch a := a.(type) {
	case int:
		if b, ok := b.(int); ok {
			return a < b, nil
		}
	case int8:
		if b, ok := b.(int8); ok {
			return a < b, nil
		}
	case int16:
		if b, ok := b.(int16); ok {
			return a < b, nil
		}
	case int32:
		if b, ok := b.(int32); ok {
			return a < b, nil
		}
	case int64:
		if b, ok := b.(int64); ok {
			return a < b, nil
		}
	case uint:
		if b, ok := b.(uint); ok {
			return a < b, nil
		}
	case uint8:
		if b, ok := b.(uint8); ok {
			return a < b, nil
		}
	case uint16:
		if b, ok := b.(uint16); ok {
			return a < b, nil
		}
	case uint32:
		if b, ok := b.(uint32); ok {
			return a < b, nil
		}
	case uint64:
		if b, ok := b.(uint64); ok {
			return a < b, nil
		}
	case float32:
		if b, ok := b.(float32); ok {
			return a < b, nil
		}
	case float64:
		if b, ok := b.(float64); ok {
			return a < b, nil
		}
	case string:
		if b, ok := b.(string); ok {
			return a < b, nil
		}
	case bool:
		if b, ok := b.(bool); ok {
			return !a && b, nil
		}
	default:
		return false, fmt.Errorf("types do not match or are not comparable\n")
	}
	return false, nil
}

func Greater[T Ordered](a, b T) bool {
	return a > b
}

func TryGreater(a, b any) (bool, error) {
	less, err := TryLess(b, a)
	return !less, err
}

// Compare checks if two values are equal.
//
// It takes two parameters, `x` and `y`, which are compared for equality.
// If the values are not comparable, an error is returned.
// If the values are comparable, but not equal, the function returns false.
// If the values are equal, the function returns true.
func Compare(x, y any) (bool, error) {
	// Check if the values are comparable
	if !IsComparable(x, y) {
		return false, fmt.Errorf("values are not comparable types")
	}

	// Compare the values
	switch x := x.(type) {
	case int:
		return x == y.(int), nil
	case int8:
		return x == y.(int8), nil
	case int16:
		return x == y.(int16), nil
	case int32:
		return x == y.(int32), nil
	case int64:
		return x == y.(int64), nil
	case uint:
		return x == y.(uint), nil
	case uint8:
		return x == y.(uint8), nil
	case uint16:
		return x == y.(uint16), nil
	case uint32:
		return x == y.(uint32), nil
	case uint64:
		return x == y.(uint64), nil
	case uintptr:
		return x == y.(uintptr), nil
	case float32:
		return x == y.(float32), nil
	case float64:
		return x == y.(float64), nil
	case string:
		return x == y.(string), nil
	case bool:
		return x == y.(bool), nil
	case complex64:
		return x == y.(complex64), nil
	case complex128:
		return x == y.(complex128), nil
	default:
		return false, fmt.Errorf("values are not from the same comparable type")
	}
}
