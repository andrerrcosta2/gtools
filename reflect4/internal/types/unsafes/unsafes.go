// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package unsafes

import (
	"github.com/andrerrcosta2/gtools/reflect4/internal/reflect4"
	"reflect"
)

// Addr returns the address of the unsafe.Pointer or an error if it isn't an unsafe.Pointer
func Addr(value reflect.Value) (uintptr, error) {
	if value.Kind() != reflect.UnsafePointer {
		return 0, reflect4.ErrNotUnsafePtr
	}
	return value.Pointer(), nil
}

func Name(value reflect.Type) (string, error) {
	if value.Kind() != reflect.UnsafePointer {
		return "", reflect4.ErrNotUnsafePtr
	}
	return value.String(), nil
}
