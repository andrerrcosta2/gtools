// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package values

import (
	"reflect"
	"unsafe"

	"github.com/andrerrcosta2/gtools/core/format/fmx"
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
	case reflect.Pointer, reflect.Map, reflect.Slice, reflect.Chan, reflect.Func, reflect.Interface, reflect.UnsafePointer:
		return value.IsNil()
	default:
		return false
	}
}

// OfUnaddr forces a refkect.Value addressability. It returns an error if the value
// is invalid, or if the value was obtained from an unexported field.
func OfUnaddr(value reflect.Value) (reflect.Value, error) {
	if !value.IsValid() {
		return Empty, fmx.Errorf("value is invalid")
	}
	if !value.CanInterface() {
		return Empty, fmx.Errorf("cannot interface unexported value of type '%v'", value.Type())
	}
	ptr := reflect.New(value.Type())
	ptr.Elem().Set(value)
	return ptr.Elem(), nil
}

// ForceOfUnaddr makes a reflect.Value addressable without safety checks.
// ⚠️ It will panic if 'value' is invalid or obtained from an unexported field.
// Use only when you are sure the input is safe to clone.
func ForceOfUnaddr(value reflect.Value) reflect.Value {
	// Create a new pointer to the value's type
	ptr := reflect.New(value.Type())
	// Copy the original value into the pointer
	ptr.Elem().Set(value)
	// Return the addressable value
	return ptr.Elem()
}

// UnsafeForceOfUnaddr makes a reflect.Value addressable without safety checks
// using unsafe pointer operations.
func UnsafeForceOfUnaddr(v reflect.Value) reflect.Value {
	return reflect.NewAt(v.Type(), unsafe.Pointer(v.UnsafeAddr())).Elem()
}

// UnsafeSet force-set a target value using the unsafe package
func UnsafeSet(target, value reflect.Value) error {
	if target.Kind() != value.Kind() {
		return fmx.Errorf("mismatched kinds between target '%s' and value '%s'",
			target.Kind(), value.Kind())
	}
	if !target.CanAddr() {
		target = ForceOfUnaddr(target)
	}
	reflect.NewAt(target.Type(), unsafe.Pointer(target.UnsafeAddr())).
		Elem().Set(value)
	return nil
}

func xint(v any) (i []reflect.Value, value reflect.Value) {
	value = reflect.ValueOf(v)
	for value.Kind() == reflect.Interface {
		i = append(i, value.Elem())
		value = i[len(i)-1]
	}
	return
}
