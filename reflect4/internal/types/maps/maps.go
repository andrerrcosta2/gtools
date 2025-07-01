// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package maps

import (
	"github.com/andrerrcosta2/gtools/reflect4/internal/reflect4"
	"reflect"
)

// Addr returns the address of the map.
func Addr(t reflect.Value) (uintptr, error) {
	if t.Kind() != reflect.Map {
		return 0, reflect4.ErrNotMap
	}
	return t.Pointer(), nil
}

func Name(t reflect.Type) (string, error) {
	if t.Kind() != reflect.Map {
		return "", reflect4.ErrNotMap
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
