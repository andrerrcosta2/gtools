// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package slices

import (
	"github.com/andrerrcosta2/gtools/reflect4/internal/reflect4"
	"reflect"
)

// Addr returns the address of the slice
func Addr(t reflect.Value) (uintptr, error) {
	if t.Kind() != reflect.Slice {
		return 0, reflect4.ErrNotSlice
	}
	return t.Pointer(), nil
}

func Name(t reflect.Type) (string, error) {
	if t.Kind() != reflect.Slice {
		return "", reflect4.ErrNotSlice
	}
	if t.Name() == "" {
		return t.String(), nil
	}
	pkgPath := t.PkgPath()
	if pkgPath != "" {
		return pkgPath + "." + t.Name(), nil
	}
	return t.Name(), nil
}
