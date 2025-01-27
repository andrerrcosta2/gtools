// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package pointers

func IsPrimitive(v any) bool {
	if v == nil {
		return false
	}
	switch v.(type) {
	case *int, *int8, *int16, *int32, *int64, *uint, *uint8, *uint16, *uint32, *uint64, *uintptr, *float32, *float64, *string, *bool, *complex64, *complex128:
		return true
	default:
		return false
	}
}
