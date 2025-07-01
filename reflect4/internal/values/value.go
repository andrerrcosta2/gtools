// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package values

import (
	"reflect"
)

var Empty = reflect.Value{}

// FieldInfo holds metadata and value of a struct field.
type FieldInfo struct {
	Name  string
	Value any
}

type FieldInterface struct {
	i     reflect.Value
	field reflect.Value // First interface wrapper, if any
}

func (fi FieldInterface) Interface() reflect.Value {
	return fi.i
}

// Field returns the first interface wrapper (if any).
func (fi FieldInterface) Field() reflect.Value {
	return fi.field
}

func FieldInterfaceOf(v any) FieldInterface {
	// Unwrap interface layers
	i, field := xint(v)

	// Ensure addressability
	if !field.CanAddr() {
		ptr := reflect.New(field.Type()) // Create pointer
		ptr.Elem().Set(field)            // Copy value into pointer
		field = ptr.Elem()
	}

	if len(i) < 1 {
		return FieldInterface{i: Empty, field: field}
	}
	return FieldInterface{i: i[0], field: field}
}

func IsNil(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.Pointer, reflect.Map, reflect.Slice, reflect.Chan, reflect.Func:
		return value.IsNil()
	default:
		return false
	}
}

// OfUnaddr wraps reflect.NewAt() to create a new reflect.Value at an arbitrary memory address.
func OfUnaddr(value reflect.Value) reflect.Value {
	// Create a new pointer to the value's type
	ptr := reflect.New(value.Type())
	// Copy the original value into the pointer
	ptr.Elem().Set(value)
	// Return the addressable value
	return ptr.Elem()
}

func xint(v any) (i []reflect.Value, value reflect.Value) {
	value = reflect.ValueOf(v)
	for value.Kind() == reflect.Interface {
		i = append(i, value.Elem())
		value = i[len(i)-1]
	}
	return
}
