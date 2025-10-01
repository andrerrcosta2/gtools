// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package values

import (
	"reflect"
	"unicode"
)

func CanNil(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.Ptr, reflect.Slice, reflect.Map, reflect.Chan, reflect.Func, reflect.Interface:
		return true
	default:
		return false
	}
}

func IsExportedField(v reflect.Value, idx int) bool {
	field := v.Type().Field(idx)
	return unicode.IsUpper(rune(field.Name[0]))
}

func IsUnexportedField(v reflect.Value, idx int) bool {
	field := v.Type().Field(idx)
	return unicode.IsLower(rune(field.Name[0]))
}
