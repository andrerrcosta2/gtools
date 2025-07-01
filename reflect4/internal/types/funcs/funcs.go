// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package funcs

import (
	"github.com/andrerrcosta2/gtools/reflect4/internal/reflect4"
	"reflect"
)

// Addr returns the address of the function, or an error if the type is not a function
func Addr(t reflect.Value) (uintptr, error) {
	if t.Kind() != reflect.Func {
		return 0, reflect4.ErrNotFunc
	}
	return t.Pointer(), nil
}

// Name returns the name of the type, or an error if the type is not a function
func Name(t reflect.Type) (string, error) {
	if t.Kind() != reflect.Func {
		return "", reflect4.ErrNotFunc
	}
	if t.Name() != "" {
		if pkgPath := t.PkgPath(); pkgPath != "" {
			return pkgPath + "." + t.Name(), nil
		}
		return t.Name(), nil
	}
	return t.String(), nil
}
