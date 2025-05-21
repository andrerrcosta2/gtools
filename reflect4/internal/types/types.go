// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package types

import (
	"reflect"
	"strings"
)

// Name extracts the available identification of a type
func Name(t reflect.Type) string {
	if t == nil || t.Kind() == reflect.Invalid {
		return "<invalid>"
	}
	if t.Kind() == reflect.Ptr {
		return unwrapName(t)
	}
	if t.Name() != "" {
		if pkgPath := t.PkgPath(); pkgPath != "" {
			return pkgPath + "." + t.Name()
		}
		return t.Name()
	}
	return t.String()
}

func unwrapName(t reflect.Type) string {
	ptr := strings.Builder{}
	for t.Kind() == reflect.Ptr {
		ptr.WriteString("*")
		t = t.Elem()
		if t == nil || t.Kind() == reflect.Invalid {
			ptr.WriteString("<invalid>")
			return ptr.String()
		}
	}
	if t.Name() != "" {
		if pkgPath := t.PkgPath(); pkgPath != "" {
			return ptr.String() + pkgPath + "." + t.Name()
		}
		return ptr.String() + t.Name()
	}
	return ptr.String() + t.String()
}
