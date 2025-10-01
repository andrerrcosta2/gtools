// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package values

import (
	"errors"
	"github.com/andrerrcosta2/gtools/core/domain/functions"
	"github.com/andrerrcosta2/gtools/core/format/fmx"
	"github.com/andrerrcosta2/gtools/reflect4/internal/types"
	"reflect"
	"unsafe"
)

var (
	ErrFieldNotFound = errors.New("reflect4.GetField - field not found")
)

// AccessFieldByName accesses a struct field by the provided function.
func AccessFieldByName(v reflect.Value, name string, fn functions.Consumer[reflect.Value]) error {
	_, ok := v.Type().FieldByName(name)
	if !ok {
		return fmx.Errorf("%s: '%s'", ErrFieldNotFound.Error(), name)
	}

	fn(v.FieldByName(name))
	return nil
}

// AccessFields accesses all exported struct fields by the provided function
// retrieving each name and reflect.Value
func AccessFields(v reflect.Value, fn functions.BiConsumer[string, reflect.Value]) {
	for i := 0; i < v.NumField(); i++ {
		fn(v.Type().Field(i).Name, v.Field(i))
	}
	return
}

// Field returns a field from a reflect.Value by its name
func Field(target reflect.Value, name string) (value reflect.Value, err error) {
	err = AccessFieldByName(target, name, func(field reflect.Value) {
		value = field
	})
	return
}

// FieldName returns a field name of a reflect value by its index. It panics if the type is not a struct.
func FieldName(target reflect.Value, idx int) string {
	return target.Type().Field(idx).Name
}

// Fields returns all exported fields within a struct
func Fields(v reflect.Value) (fields map[string]reflect.Value) {
	fields = make(map[string]reflect.Value, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		fields[v.Type().Field(i).Name] = v.Field(i)
	}
	return
}

// NilFields returns all exported nil fields within a
func NilFields(v reflect.Value) (fields map[string]reflect.Value) {
	fields = make(map[string]reflect.Value, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if IsNil(field) {
			fields[v.Type().Field(i).Name] = v.Field(i)
		}
	}
	return
}

// NoNilFields returns all non-nil fields within a struct
func NoNilFields(v reflect.Value) map[string]reflect.Value {
	fields := make(map[string]reflect.Value, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if !IsNil(field) {
			fields[v.Type().Field(i).Name] = v.Field(i)
		}
	}
	return fields
}

// RideExportedFields rides a struct exported fields with a bi-predicate giving its index and value
func RideExportedFields(v reflect.Value, fn functions.BiPredicate[int, reflect.Value]) {
	for i, n := 0, v.NumField(); i < n; i++ {
		if IsExportedField(v, i) {
			if !fn(i, v.Field(i)) {
				break
			}
		}
	}
}

// RideFields rides a struct fields with a bi-predicate giving its index and value
func RideFields(v reflect.Value, fn functions.BiPredicate[int, reflect.Value]) {
	for i, n := 0, v.NumField(); i < n; i++ {
		if !fn(i, v.Field(i)) {
			break
		}
	}
}

// SetField sets a struct field by the provided value by its name if it is exported
func SetField(target reflect.Value, name string, value reflect.Value) error {
	return AccessFieldByName(target, name, func(field reflect.Value) {
		field.Set(value)
	})
}

// UnsafeFieldAccess accesses a struct field by name using the unsafe package to allow access
// unexported fields.
func UnsafeFieldAccess(v reflect.Value, name string, fn functions.BiConsumer[reflect.Value, unsafe.Pointer]) error {
	if !v.CanAddr() {
		v = UnsafeOfUnaddr(v)
	}

	field := v.FieldByName(name)

	if !field.IsValid() {
		return fmx.Errorf("%s: '%s'", ErrFieldNotFound.Error(), name)
	}

	fn(field, unsafe.Pointer(field.UnsafeAddr()))
	return nil
}

// UnsafeFieldsAccess accesses a struct field by name using the unsafe package to allow access
// unexported fields.
func UnsafeFieldsAccess(v reflect.Value, fn functions.BiConsumer[reflect.Value, unsafe.Pointer]) {
	if !v.CanAddr() {
		v = UnsafeOfUnaddr(v)
	}

	for i := 0; i < v.NumField(); i++ {
		fn(v.Field(i), unsafe.Pointer(v.Field(i).UnsafeAddr()))
	}
	return
}

// UnsafeGetAllFields returns all unexported fields of a struct
//
// to avoid GC issues, the returned value is a clone of the original value
func UnsafeGetAllFields(v reflect.Value) (fields map[string]reflect.Value) {
	if !v.CanAddr() {
		v = UnsafeOfUnaddr(v)
	}
	fields = make(map[string]reflect.Value)
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		name := v.Type().Field(i).Name
		ptr := unsafe.Pointer(field.UnsafeAddr())
		fields[name] = reflect.NewAt(field.Type(), ptr).Elem()
	}
	return
}

// UnsafeGetField returns a struct field by name using the unsafe package
//
// to avoid GC issues, the returned value is a clone of the original value
func UnsafeGetField(v reflect.Value, name string) (value reflect.Value, err error) {
	if !v.CanAddr() {
		v = UnsafeOfUnaddr(v)
	}
	field := v.FieldByName(name)
	if !field.IsValid() {
		return Empty, fmx.Errorf("%s: '%s'", ErrFieldNotFound.Error(), name)
	}
	ptr := unsafe.Pointer(field.UnsafeAddr())
	return reflect.NewAt(field.Type(), ptr).Elem(), nil
}

// UnsafeGetFields returns a map of values of struct fields by names
//
// to avoid GC issues the returned values are copies of the original values
func UnsafeGetFields(v reflect.Value, names ...string) (fields map[string]any, err error) {
	if !v.CanAddr() {
		v = UnsafeOfUnaddr(v)
	}

	for _, name := range names {
		field := v.FieldByName(name)

		if !field.IsValid() {
			return nil, ErrFieldNotFound
		}

		ptr := unsafe.Pointer(field.UnsafeAddr())
		fields[name] = reflect.NewAt(field.Type(), ptr).Elem().Interface()
	}
	return
}

// UnsafeNilFields returns all nil fields within a
func UnsafeNilFields(v reflect.Value) (fields []reflect.Value, err error) {
	fields = make([]reflect.Value, 0, v.NumField())

	if !v.CanAddr() {
		v = UnsafeOfUnaddr(v)
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		elem := reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem()
		if IsNil(elem) {
			fields = append(fields, field)
		}
	}
	return
}

// UnsafeNotNilFields returns all non-nil fields within a struct
func UnsafeNotNilFields(v reflect.Value) (fields []reflect.Value) {
	fields = make([]reflect.Value, 0, v.NumField())

	if !v.CanAddr() {
		v = UnsafeOfUnaddr(v)
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		elem := reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem()
		if !IsNil(elem) {
			fields = append(fields, field)
		}
	}
	return
}

// UnsafeRideFields rides a struct using the unsafe package to allow access to
// its unexported fields.
func UnsafeRideFields(v reflect.Value, fn functions.BiConsumer[string, unsafe.Pointer]) {
	if !v.CanAddr() {
		v = UnsafeOfUnaddr(v)
	}

	for i := 0; i < v.NumField(); i++ {
		name := v.Type().Field(i).Name
		fn(name, unsafe.Pointer(v.Field(i).UnsafeAddr()))
	}
	return
}

// UnsafeDeepRideFields rides a struct deeply using the unsafe package to allow access to
// its unexported fields.
func UnsafeDeepRideFields(v reflect.Value, fn functions.BiConsumer[string, unsafe.Pointer]) {
	if !v.CanAddr() {
		v = UnsafeOfUnaddr(v)
	}
	unsafeDeepRideFields(v, fn)
	return
}

// unsafeDeepRideFields should be used only under favorable conditions where none on its ops
// could panic, like handling with unaddressable data.
func unsafeDeepRideFields(v reflect.Value, fn functions.BiConsumer[string, unsafe.Pointer]) {
	for i := 0; i < v.NumField(); i++ {
		if types.HasDepth(v.Type()) {
			unsafeDeepRideFields(v.Field(i), fn)
		}
		name := v.Type().Field(i).Name
		fn(name, unsafe.Pointer(v.Field(i).UnsafeAddr()))
	}
}

// UnsafeSetFieldByName sets a struct field by name using the unsafe package to allow access
// to unexported fields.
func UnsafeSetFieldByName(target, value reflect.Value, name string) error {
	field := target.FieldByName(name)
	return UnsafeSet(field, value)
}

// UnsafeSetFieldByIndex sets a struct field by index using the unsafe package to allow access
// to unexported fields.
func UnsafeSetFieldByIndex(target, value reflect.Value, idx ...int) error {
	field := target.FieldByIndex(idx)
	return UnsafeSet(field, value)
}
