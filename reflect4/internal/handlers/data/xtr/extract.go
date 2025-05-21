// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package xtr

import (
	"github.com/andrerrcosta2/gtools/reflect4/internal/values"
	"reflect"
	"unsafe"
)

// Fields extracts both exported and unexported fields from a struct.
func Fields(obj any) []values.FieldInfo {
	fi := values.FieldInterfaceOf(obj)
	v := fi.Field()
	t := v.Type()

	if v.Kind() != reflect.Struct {
		panic("reflect4: Fields() expects a struct or pointer to struct")
	}

	var fields []values.FieldInfo
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldValue := v.Field(i)

		// If the field is unexported, use unsafe to access it
		if !fieldValue.CanInterface() {
			fieldValue = accessUnexportedField(v, field)
		}

		fields = append(fields, values.FieldInfo{
			Name:  field.Name,
			Value: fieldValue.Interface(),
		})
	}
	return fields
}

// Interface extract all interfaces from a value until it finds a non interface kind
func Interface(v any) (interfaces []reflect.Value, value reflect.Value) {
	rv := reflect.ValueOf(v)
	for rv.Kind() == reflect.Interface {
		interfaces = append(interfaces, rv.Elem())
		rv = interfaces[len(interfaces)-1]
	}
	return interfaces, rv
}

// accessUnexportedField uses unsafe to retrieve unexported field values.
func accessUnexportedField(v reflect.Value, field reflect.StructField) reflect.Value {
	ptr := unsafe.Pointer(v.UnsafeAddr() + field.Offset)
	return reflect.NewAt(field.Type, ptr).Elem()
}
