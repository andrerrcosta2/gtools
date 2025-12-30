// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package structs

import (
	"path"
	"reflect"
)

// Name returns the struct name
// It panics if the type isn't struct
func Name(t reflect.Type) string {
	if t.Name() == "" {
		return t.String() // Anonymous struct
	}

	pkgPath := t.PkgPath()
	if pkgPath != "" {
		// Extract the last segment of the package path
		short := path.Base(pkgPath)
		return short + "." + t.Name()
	}
	return t.Name()
}

// NumExportedFields returns the number of exported fields within a struct
// It panics if the type isn't struct
func NumExportedFields(t reflect.Type) int {
	var count int
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).IsExported() {
			count++
		}
	}
	return count
}

// NumUnexportedFields returns the number of unexported fields within a struct
// It panics if the type isn't struct
func NumUnexportedFields(t reflect.Type) int {
	var count int
	for i := 0; i < t.NumField(); i++ {
		if !t.Field(i).IsExported() {
			count++
		}
	}
	return count
}

// Tags returns all tags within a struct
// It panics if the type isn't a struct
func Tags(t reflect.Type, key string) map[string]string {
	result := make(map[string]string)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		result[field.Name] = field.Tag.Get(key)
	}
	return result
}
