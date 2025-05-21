// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package data

import (
	"fmt"
	"github.com/andrerrcosta2/gtools/reflect4/internal/values"
	"reflect"
)

type Opt uint8

const (
	Default                Opt = 0
	IgnoreUnexportedFields Opt = 1 << iota
	AllowNilVsEmpty
)

func DeepEqual(v1, v2 reflect.Value, opt Opt, path string, diffs *[]string) {
	if !v1.IsValid() || !v2.IsValid() {
		if v1.IsValid() != v2.IsValid() {
			*diffs = append(*diffs, fmt.Sprintf("%s: one is nil", path))
		}
		return
	}

	if v1.Type() != v2.Type() {
		*diffs = append(*diffs, fmt.Sprintf("%s: type mismatch (%v vs %v)", path, v1.Type(), v2.Type()))
		return
	}

	// Extract wrapped interface values
	fi1 := values.FieldInterfaceOf(v1.Interface())
	fi2 := values.FieldInterfaceOf(v2.Interface())

	v1 = fi1.Field()
	v2 = fi2.Field()

	switch v1.Kind() {
	case reflect.Ptr, reflect.Interface:
		DeepEqual(v1.Elem(), v2.Elem(), opt, path, diffs)
	case reflect.Struct:
		//for _, field := range values.Fields(v1.Interface()) {
		//	fieldPath := fmt.Sprintf("%s.%s", path, field.Name)
		//	DeepEqual(reflect.ValueOf(field.Value), reflect.ValueOf(values.Fields(v2.Interface())[0].Value), opt, fieldPath, diffs)
		//}
	case reflect.Slice, reflect.Array:
		if v1.Len() != v2.Len() {
			*diffs = append(*diffs, fmt.Sprintf("%s: length mismatch (%d vs %d)", path, v1.Len(), v2.Len()))
			return
		}
		for i := 0; i < v1.Len(); i++ {
			DeepEqual(v1.Index(i), v2.Index(i), opt, fmt.Sprintf("%s[%d]", path, i), diffs)
		}
	case reflect.Map:
		if v1.Len() != v2.Len() {
			*diffs = append(*diffs, fmt.Sprintf("%s: map length mismatch (%d vs %d)", path, v1.Len(), v2.Len()))
			return
		}
		for _, key := range v1.MapKeys() {
			keyStr := fmt.Sprintf("%v", key.Interface())
			DeepEqual(v1.MapIndex(key), v2.MapIndex(key), opt, fmt.Sprintf("%s[%s]", path, keyStr), diffs)
		}
	default:
		if v1.Interface() != v2.Interface() {
			*diffs = append(*diffs, fmt.Sprintf("%s: %v != %v", path, v1.Interface(), v2.Interface()))
		}
	}
}
