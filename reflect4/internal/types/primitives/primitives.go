// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package primitives

import (
	"github.com/andrerrcosta2/gtools/reflect4/internal/reflect4"
	"path"
	"reflect"
)

func Name(t reflect.Type) (string, error) {
	if t.Kind() > reflect.Complex128 && t.Kind() != reflect.String {
		return "", reflect4.ErrNotPrimitive
	}
	if t.Name() == "" {
		return t.String(), nil
	}
	pkgPath := t.PkgPath()
	if pkgPath != "" {
		short := path.Base(pkgPath)
		return short + "." + t.Name(), nil
	}
	return t.Name(), nil
}

func QfName(t reflect.Type) (string, error) {
	if t.Kind() > reflect.Complex128 && t.Kind() != reflect.String {
		return "", reflect4.ErrNotPrimitive
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
