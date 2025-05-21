// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package pointers

import (
	"github.com/andrerrcosta2/gtools/reflect4/internal/reflect4"
	"reflect"
)

func Name(value reflect.Type) (string, error) {
	if value.Kind() != reflect.Ptr {
		return "", reflect4.ErrNotPointer
	}
	return value.String(), nil
}

func Of(value reflect.Value) (uintptr, error) {
	if !value.IsValid() {
		// Handle invalid values gracefully
		return 0, reflect4.ErrInvalidValue
	}

	switch value.Kind() {
	// These types support the Pointer() method
	case reflect.Ptr, reflect.UnsafePointer, reflect.Slice, reflect.Map, reflect.Chan, reflect.Func:
		return value.Pointer(), nil

	default:
		if value.CanAddr() {
			return value.Addr().Pointer(), nil
		}
		// Create a pointer to the value using reflect.NewAt
		return reflect.ValueOf(&value).Pointer(), nil
	}
}
