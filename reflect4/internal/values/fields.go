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
	ErrNotStruct       = errors.New("reflect4.GetField - target is not a struct")
	ErrFieldNotFound   = errors.New("reflect4.GetField - field not found")
	ErrUnexportedField = errors.New("reflect4.GetField - field is not exported")
)

// AccessFieldByName accesses a struct field by the provided function.
func AccessFieldByName(v reflect.Value, name string, fn functions.Consumer[reflect.Value]) error {
	if v.Kind() != reflect.Struct {
		return fmx.Errorf("%s: %s", ErrNotStruct.Error(), v.Type().String())
	}

	_, ok := v.Type().FieldByName(name)
	if !ok {
		return fmx.Errorf("%s: '%s'", ErrFieldNotFound.Error(), name)
	}

	fn(v.FieldByName(name))
	return nil
}

// AccessFields accesses all exported struct fields by the provided function
// retrieving each name and reflect.Value
func AccessFields(v reflect.Value, fn functions.BiConsumer[string, reflect.Value]) error {
	if v.Kind() != reflect.Struct {
		return fmx.Errorf("%s: %s", ErrNotStruct.Error(), v.Type().String())
	}
	for i := 0; i < v.NumField(); i++ {
		fn(v.Type().Field(i).Name, v.Field(i))
	}
	return nil
}

// Field returns a field from a reflect.Value by its name
func Field(target reflect.Value, name string) (value reflect.Value, err error) {
	err = AccessFieldByName(target, name, func(field reflect.Value) {
		value = field
	})
	return
}

// Fields returns all exported fields within a struct
func Fields(v reflect.Value) (fields map[string]reflect.Value, err error) {
	if v.Kind() != reflect.Struct {
		return nil, fmx.Errorf("%s: %s", ErrNotStruct.Error(), v.Type().String())
	}
	fields = make(map[string]reflect.Value, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		fields[v.Type().Field(i).Name] = v.Field(i)
	}
	return
}

// NilFields returns all exported nil fields within a
func NilFields(v reflect.Value) (fields map[string]reflect.Value, err error) {
	if v.Kind() != reflect.Struct {
		return nil, fmx.Errorf("%s: %s", ErrNotStruct.Error(), v.Type().String())
	}
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
func NoNilFields(v reflect.Value) (fields map[string]reflect.Value, err error) {
	if v.Kind() != reflect.Struct {
		return nil, fmx.Errorf("%s: %s", ErrNotStruct.Error(), v.Type().String())
	}
	fields = make(map[string]reflect.Value, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if !IsNil(field) {
			fields[v.Type().Field(i).Name] = v.Field(i)
		}
	}
	return
}

// RideFields rides a struct exported fields with a bi-consumer giving its name and value
func RideFields(v reflect.Value, fn functions.BiConsumer[string, reflect.Value]) error {
	if v.Kind() != reflect.Struct {
		return fmx.Errorf("%s: %s", ErrNotStruct.Error(), v.Type().String())
	}
	for i := 0; i < v.NumField(); i++ {
		fn(v.Type().Field(i).Name, v.Field(i))
	}

	return nil
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
	if v.Kind() != reflect.Struct {
		return fmx.Errorf("%s: %s", ErrNotStruct.Error(), v.Type().String())
	}
	if !v.CanAddr() {
		v = OfUnaddr(v)
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
func UnsafeFieldsAccess(v reflect.Value, fn functions.BiConsumer[reflect.Value, unsafe.Pointer]) error {
	if v.Kind() != reflect.Struct {
		return fmx.Errorf("%s: %s", ErrNotStruct.Error(), v.Type().String())
	}
	if !v.CanAddr() {
		v = OfUnaddr(v)
	}

	for i := 0; i < v.NumField(); i++ {
		fn(v.Field(i), unsafe.Pointer(v.Field(i).UnsafeAddr()))
	}
	return nil
}

// UnsafeGetAllFields returns all unexported fields of a struct
//
// to avoid GC issues, the returned value is a copy of the original value
func UnsafeGetAllFields(v reflect.Value) (fields map[string]reflect.Value, err error) {
	if v.Kind() != reflect.Struct {
		return nil, ErrNotStruct
	}

	if !v.CanAddr() {
		v = OfUnaddr(v)
	}
	fields = make(map[string]reflect.Value)
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		name := v.Type().Field(i).Name
		if !field.IsValid() {
			return nil, ErrFieldNotFound
		}
		ptr := unsafe.Pointer(field.UnsafeAddr())
		fields[name] = reflect.NewAt(field.Type(), ptr).Elem()
	}
	return
}

// UnsafeGetField returns a struct field by name using the unsafe package
//
// to avoid GC issues, the returned value is a copy of the original value
func UnsafeGetField(v reflect.Value, name string) (value reflect.Value, err error) {
	err = UnsafeFieldAccess(v, name, func(field reflect.Value, ptr unsafe.Pointer) {
		value = reflect.NewAt(field.Type(), ptr).Elem()
	})
	return
}

// UnsafeGetFields returns a map of values of struct fields by names
//
// to avoid GC issues the returned values are copies of the original values
func UnsafeGetFields(v reflect.Value, names ...string) (fields map[string]any, err error) {
	if v.Kind() != reflect.Struct {
		return nil, fmx.Errorf("%s: %s", ErrNotStruct.Error(), v.Type().String())
	}
	if !v.CanAddr() {
		v = OfUnaddr(v)
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
	if v.Kind() != reflect.Struct {
		return nil, fmx.Errorf("%s: %s", ErrNotStruct.Error(), v.Type().String())
	}
	fields = make([]reflect.Value, 0, v.NumField())

	if !v.CanAddr() {
		v = OfUnaddr(v)
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
func UnsafeNotNilFields(v reflect.Value) (fields []reflect.Value, err error) {
	if v.Kind() != reflect.Struct {
		return nil, fmx.Errorf("%s: %s", ErrNotStruct.Error(), v.Type().String())
	}

	fields = make([]reflect.Value, 0, v.NumField())

	if !v.CanAddr() {
		v = OfUnaddr(v)
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
func UnsafeRideFields(v reflect.Value, fn functions.BiConsumer[string, unsafe.Pointer]) error {
	if v.Kind() != reflect.Struct {
		return fmx.Errorf("%s: %s", ErrNotStruct.Error(), v.Type().String())
	}
	if !v.CanAddr() {
		v = OfUnaddr(v)
	}

	for i := 0; i < v.NumField(); i++ {
		name := v.Type().Field(i).Name
		fn(name, unsafe.Pointer(v.Field(i).UnsafeAddr()))
	}
	return nil
}

// UnsafeDeepRideFields rides a struct deeply using the unsafe package to allow access to
// its unexported fields.
func UnsafeDeepRideFields(v reflect.Value, fn functions.BiConsumer[string, unsafe.Pointer]) error {
	if v.Kind() != reflect.Struct {
		return fmx.Errorf("%s: %s", ErrNotStruct.Error(), v.Type().String())
	}
	if !v.CanAddr() {
		v = OfUnaddr(v)
	}
	unsafeDeepRideFields(v, fn)
	return nil
}

// unsafeDeepRideFields should be used only under favorable conditions where none on its operations
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
	if target.Kind() != reflect.Struct {
		return fmx.Errorf("%s: %s", ErrNotStruct.Error(), target.Type().String())
	}
	if !target.CanAddr() {
		return fmx.Errorf("target must be addressable")
	}
	field := target.FieldByName(name)
	if !field.IsValid() {
		return fmx.Errorf("%s: '%s'", ErrFieldNotFound.Error(), name)
	}
	if field.Kind() != value.Kind() {
		return fmx.Errorf("error. mismatched kinds between target field '%s' and value '%s'",
			field.Kind(), value.Kind())
	}
	reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).
		Elem().Set(value)
	return nil
}

// UnsafeSetFieldByIndex sets a struct field by index using the unsafe package to allow access
// to unexported fields.
func UnsafeSetFieldByIndex(target, value reflect.Value, idx ...int) error {
	if target.Kind() != reflect.Struct {
		return fmx.Errorf("%s: %s", ErrNotStruct.Error(), target.Type().String())
	}
	if !target.CanAddr() {
		return fmx.Errorf("target must be addressable")
	}
	field := target.FieldByIndex(idx)
	if !field.IsValid() {
		return fmx.Errorf("%s: '%v'", ErrFieldNotFound.Error(), idx)
	}
	if field.Kind() != value.Kind() {
		return fmx.Errorf("error. mismatched kinds between target field '%s' and value '%s'",
			field.Kind(), value.Kind())
	}
	reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).
		Elem().Set(value)
	return nil
}
