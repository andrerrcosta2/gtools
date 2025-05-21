// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package arrays

import (
	"github.com/andrerrcosta2/gtools/reflect4/internal/reflect4"
	"reflect"
)

func Name(t reflect.Type) (string, error) {
	if t.Kind() != reflect.Array {
		return "", reflect4.ErrNotArray
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
