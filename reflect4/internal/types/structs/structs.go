// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package structs

import (
	"github.com/andrerrcosta2/gtools/reflect4/internal/reflect4"
	"path"
	"reflect"
)

func Name(t reflect.Type) (string, error) {
	if t.Kind() != reflect.Struct {
		return "", reflect4.ErrNotStruct
	}
	if t.Name() == "" {
		return t.String(), nil // Anonymous struct
	}

	pkgPath := t.PkgPath()
	if pkgPath != "" {
		// Extract the last segment of the package path
		short := path.Base(pkgPath)
		return short + "." + t.Name(), nil
	}
	return t.Name(), nil
}

// QfName returns the fully qualified name of a struct
func QfName(t reflect.Type) (string, error) {
	if t.Kind() != reflect.Struct {
		return "", reflect4.ErrNotStruct
	}
	if t.Name() == "" {
		return t.String(), nil // Anonymous struct
	}

	pkgPath := t.PkgPath()
	if pkgPath != "" {
		// full package path
		return pkgPath + "." + t.Name(), nil
	}
	return t.Name(), nil
}
