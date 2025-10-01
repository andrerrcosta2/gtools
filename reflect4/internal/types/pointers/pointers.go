// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package pointers

import (
	"github.com/andrerrcosta2/gtools/reflect4/internal/reflect4"
	"reflect"
)

// Name returns the kind string representation
func Name(value reflect.Type) (string, error) {
	if value.Kind() != reflect.Ptr {
		return "", reflect4.ErrNotPointer
	}
	return value.String(), nil
}

// UnsafeNewOf creates a new pointer with a given reflect.Value as element.
//
// ⚠️ This function does not check for v.IsValid(), and the address returned
// for non-addressable values is from a clone
func UnsafeNewOf(x reflect.Value) reflect.Value {
	ptrCopy := reflect.New(x.Type())
	ptrCopy.Elem().Set(x)
	return ptrCopy
}
