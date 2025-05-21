// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package prim

// IsAssignableTo checks if the provided value can be assigned to the type T.
func IsAssignableTo[T Any](v any) bool {
	_, ok := v.(T)
	return ok
}

// IsConvertibleTo checks if the provided value can be converted to the type T.
func IsConvertibleTo[T Any](v any) bool {
	var t T
	switch any(t).(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr, float32, float64:
		switch v.(type) {
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr, float32, float64:
			return true
		default:
			return false
		}

	case string:
		switch v.(type) {
		case string:
			return true
		default:
			return false
		}

	case complex64, complex128:
		switch v.(type) {
		case complex64, complex128:
			return true
		default:
			return false
		}

	case bool:
		switch v.(type) {
		case bool:
			return true
		default:
			return false
		}

	default:
		return false
	}
}

// IsPrimitive checks if the provided value is a primitive type.
func IsPrimitive(value any) bool {
	switch value.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr, float32, float64, string, bool, complex64, complex128:
		return true
	default:
		return false
	}
}

// IsComparable checks if all provided values are of comparable types.
//
// It takes a variable number of arguments of any type and returns a boolean indicating
// if all values are of types that can be compared for equality, such as numeric types,
// strings, and booleans.
func IsComparable(value ...any) (is bool) {
	// Iterate over each value provided
	for _, v := range value {
		// Check the type of each value
		switch v.(type) {
		// If the value is of a comparable type, set the return value to true
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr, float32,
			float64, string, bool, complex64, complex128, *int, *int8, *int16, *int32, *int64, *uint,
			*uint8, *uint16, *uint32, *uint64, *uintptr, *float32, *float64, *string, *bool, *complex64,
			*complex128:
			return true
		// If any value is not of a comparable type, return false immediately
		default:
			return false
		}
	}
	// Return false if there is no value
	return
}

// IsPointer checks if the provided value is a pointer to a primitive type.
//
// It returns true if the value is a pointer to a recognized primitive type,
// and false otherwise.
func IsPointer(v any) bool {
	if v == nil {
		// A nil value cannot be a pointer
		return false
	}

	// Check if the type of the value is a pointer to a primitive type
	switch v.(type) {
	case *int, *int8, *int16, *int32, *int64, *uint, *uint8, *uint16, *uint32, *uint64, *uintptr,
		*float32, *float64, *string, *bool, *complex64, *complex128:
		return true
	default:
		return false
	}
}

func IsZeroOf[T Any](value any) bool {
	var zero T
	return value == zero
}
