// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package structs

import (
	"github.com/andrerrcosta2/gtools/core/format/fmx"
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

// NumExportedFields returns the number of exported fields within a struct
func NumExportedFields(t reflect.Type) (int, error) {
	if t.Kind() != reflect.Struct {
		return -1, fmx.Errorf("%s: %s", reflect4.ErrNotStruct.Error(), t.String())
	}
	var count int
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).IsExported() {
			count++
		}
	}
	return count, nil
}

// NumUnexportedFields returns the number of unexported fields within a struct
func NumUnexportedFields(t reflect.Type) (int, error) {
	if t.Kind() != reflect.Struct {
		return -1, fmx.Errorf("%s: %s", reflect4.ErrNotStruct.Error(), t.String())
	}
	var count int
	for i := 0; i < t.NumField(); i++ {
		if !t.Field(i).IsExported() {
			count++
		}
	}
	return count, nil
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

func Tags(t reflect.Type, key string) (map[string]string, error) {
	if t.Kind() != reflect.Struct {
		return nil, reflect4.ErrNotStruct
	}

	result := make(map[string]string)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.IsExported() {
			result[field.Name] = field.Tag.Get(key)
		}
	}
	return result, nil
}
