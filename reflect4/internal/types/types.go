// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package types

import (
	"log"
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
	return ValidValueName(t)
}

// Unwrap unwraps all ptrs and interfaces until it reaches a value
func Unwrap(v reflect.Type) reflect.Type {
	for v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	return v
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

func ValidValueName(t reflect.Type) string {
	log.Printf("TYPE DEBUG: kind=%s name=%q pkg=%q string=%q\n",
		t.Kind(), t.Name(), t.PkgPath(), t.String())
	if t.Name() != "" {
		if pkgPath := t.PkgPath(); pkgPath != "" {
			return pkgPath + "." + t.Name()
		}
		return t.Name()
	}
	return t.String()
}
