// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package prim

func AssignableTo[T Any](v any) bool {
	_, ok := v.(T)
	return ok
}

func ConvertibleTo[T Any](v any) bool {
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
