// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package values

import (
	"errors"
	"reflect"
	"strings"
	"unsafe"

	"github.com/andrerrcosta2/gtools/core/domain/functions"
	"github.com/andrerrcosta2/gtools/reflect4/internal/types"
)

var (
	ErrFieldNotFound = func(caller, field string) error {
		return errors.New(caller + ": field '" + field + "' not found")
	}
	ErrInvalidValue = func(caller string) error {
		return errors.New(caller + ": the value sent as argument is invalid")
	}
	ErrUnexportedField = func(caller, field string) error {
		return errors.New(caller + ": field '" + field + "' is not exported")
	}

	ErrFieldsNotFound = func(caller string, fields ...string) error {
		sb := strings.Builder{}
		sb.WriteString(caller + ": the field4")
		for _, field := range fields {
			sb.WriteString(" '" + field + "', ")
		}
		sb.WriteString("were not found")
		return errors.New(sb.String())
	}
	ErrFieldTypeMismatch = func(caller, exp, rec string) error {
		return errors.New(caller + ": field type mismatch. expected '" + exp +
			"' got '" + rec + "'")
	}
	ErrUnaddressable = func(caller string) error {
		return errors.New(caller + ": the struct is unaddressable")
	}
)

// AccessFieldByName accesses a struct field by the provided function.
func AccessFieldByName(v reflect.Value, name string, fn functions.Consumer[reflect.Value]) error {
	_, ok := v.Type().FieldByName(name)
	if !ok {
		return ErrFieldNotFound("reflect4.AccessFieldByName", name)
	}

	fn(v.FieldByName(name))
	return nil
}

// Field returns a field from a reflect.Value by its name
func Field(target reflect.Value, name string) (value reflect.Value, err error) {
	_, ok := target.Type().FieldByName(name)
	if !ok {
		return Empty, ErrFieldNotFound("reflect4.Field", name)
	}
	return target.FieldByName(name), nil
}

func FieldExp(v reflect.Value, name string) (reflect.Value, error) {
	sf, ok := v.Type().FieldByName(name)
	if !ok {
		return Empty, ErrFieldNotFound("reflect4.FieldExp", name)
	}
	if !sf.IsExported() {
		return Empty, ErrUnexportedField("reflect4.FieldExp", name)
	}
	return v.FieldByName(name), nil
}

// FieldName returns a field name of a reflect value by its index. It panics if the type is not a struct.
func FieldName(target reflect.Value, idx int) string {
	return target.Type().Field(idx).Name
}

// Fields returns all exported field4 within a struct
// if the reflect value isn't a struct it panics.
func Fields(v reflect.Value) (fields map[string]reflect.Value) {
	fields = make(map[string]reflect.Value, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		fields[v.Type().Field(i).Name] = v.Field(i)
	}
	return
}

// UnsafeFieldValueByIndex retrieve a field value by index using its unsafe
// address
func UnsafeFieldValueByIndex(v reflect.Value, idx int) reflect.Value {
	field := v.Field(idx)
	if field.CanInterface() {
		return field
	}
	ptr := unsafe.Pointer(field.UnsafeAddr())
	return reflect.NewAt(field.Type(), ptr).Elem()
}

// NilFields returns all exported nil field4 within a
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

// NoNilFields returns all non-nil field4 within a struct
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

// RideExportedFieldsByIdx rides a struct exported field4 with a bi-predicate giving its index and value
func RideExportedFieldsByIdx(v reflect.Value, fn functions.BiPredicate[int, reflect.Value]) {
	for i, n := 0, v.NumField(); i < n; i++ {
		if v.Type().Field(i).IsExported() {
			if !fn(i, v.Field(i)) {
				break
			}
		}
	}
}

// RideExportedFieldsByName rides a struct exported field4 with a bi-predicate giving its name and value.
// It panics if the reflect.Value isn't a struct
func RideExportedFieldsByName(v reflect.Value, fn functions.BiPredicate[string, reflect.Value]) {
	for i, n := 0, v.NumField(); i < n; i++ {
		ft := v.Type().Field(i)
		if ft.IsExported() {
			if !fn(ft.Name, v.Field(i)) {
				break
			}
		}
	}
}

// RideFields rides a struct field4 with a bi-predicate giving its index and value
func RideFields(v reflect.Value, fn functions.BiPredicate[int, reflect.Value]) {
	for i, n := 0, v.NumField(); i < n; i++ {
		if !fn(i, v.Field(i)) {
			break
		}
	}
}

// RideFieldsByName accesses all struct field4 by the provided function
// retrieving each name and reflect.Value
func RideFieldsByName(v reflect.Value, fn functions.BiPredicate[string, reflect.Value]) {
	for i := 0; i < v.NumField(); i++ {
		if !fn(v.Type().Field(i).Name, v.Field(i)) {
			break
		}
	}
}

// SetExpFieldByName sets a struct field by the provided value by its name
// it returns an error if the field is not exported or not found
// it panics if the reflect.Value isn't a struct.
func SetExpFieldByName(target reflect.Value, name string, value reflect.Value) error {
	field := target.FieldByName(name)
	if !field.IsValid() {
		return ErrFieldNotFound("reflect4.SetExpFieldByName", name)
	}
	if !value.IsValid() {
		return ErrInvalidValue("reflect4.SetExpFieldByName")
	}
	if !field.CanSet() {
		return errors.New("reflect4.SetExpFieldByName: could not set. the field '" + name +
			"' is not exported or the target struct isn't addressable")
	}
	if field.Type() != value.Type() {
		return ErrFieldTypeMismatch("reflect4.SetExpFieldByName", types.ValidValueName(field.Type()),
			types.ValidValueName(value.Type()))
	}
	field.Set(value)
	return nil
}

// UnsafeFieldAccess accesses a struct field by name using the unsafe package to allow access
// unexported field4.
func UnsafeFieldAccess(v reflect.Value, name string, fn functions.BiConsumer[reflect.Value, unsafe.Pointer]) error {
	if !v.CanAddr() {
		return ErrUnaddressable("reflect4.UnsafeFieldAccess")
	}

	field := v.FieldByName(name)

	if !field.IsValid() {
		return ErrFieldNotFound("reflect4.UnsafeFieldAccess", name)
	}

	fn(field, unsafe.Pointer(field.UnsafeAddr()))
	return nil
}

// UnsafeFieldAccessIfAddr accesses a struct field by name using unsafe.
// It requires the value to be addressable; otherwise it returns an error.
func UnsafeFieldAccessIfAddr(
	v reflect.Value,
	name string,
	fn func(field reflect.Value, ptr unsafe.Pointer) error,
) error {
	if !v.CanAddr() {
		return ErrUnaddressable("reflect4.UnsafeFieldAccessIfAddr")
	}

	field := v.FieldByName(name)
	if !field.IsValid() {
		return ErrFieldNotFound("reflect4.UnsafeFieldAccessIfAddr", name)
	}

	return fn(field, unsafe.Pointer(field.UnsafeAddr()))
}

// UnsafeGetAllFields returns all unexported field4 of a struct
//
// to avoid GC issues, the returned value is a clone of the original value
func UnsafeGetAllFields(v reflect.Value) (fields map[string]reflect.Value) {
	if !v.CanAddr() {
		v = ForceOfUnaddr(v)
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
		v = ForceOfUnaddr(v)
	}
	field := v.FieldByName(name)
	if !field.IsValid() {
		return Empty, ErrFieldNotFound("reflect4.UnsafeGetField", name)
	}
	ptr := unsafe.Pointer(field.UnsafeAddr())
	return reflect.NewAt(field.Type(), ptr).Elem(), nil
}

// UnsafeGetFields returns a map of values of struct field4 by names
//
// to avoid GC issues the returned values are copies of the original values
func UnsafeGetFields(v reflect.Value, names ...string) (map[string]any, error) {
	if !v.CanAddr() {
		v = ForceOfUnaddr(v)
	}

	fields := make(map[string]any)
	var missing []string

	for _, name := range names {
		field := v.FieldByName(name)
		if !field.IsValid() {
			missing = append(missing, name)
			continue
		}

		ptr := unsafe.Pointer(field.UnsafeAddr())
		fields[name] = reflect.NewAt(field.Type(), ptr).Elem().Interface()
	}

	if len(missing) > 0 {
		return fields, ErrFieldsNotFound(
			"reflect4.UnsafeGetFields",
			missing...,
		)
	}

	return fields, nil
}

// UnsafeNilFields returns all nil field4 within a
func UnsafeNilFields(v reflect.Value) (fields []reflect.Value, err error) {
	fields = make([]reflect.Value, 0, v.NumField())

	if !v.CanAddr() {
		v = ForceOfUnaddr(v)
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

// UnsafeNotNilFields returns all non-nil field4 within a struct
func UnsafeNotNilFields(v reflect.Value) (fields []reflect.Value) {
	fields = make([]reflect.Value, 0, v.NumField())

	if !v.CanAddr() {
		v = ForceOfUnaddr(v)
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
// its unexported field4.
// it retrieves its unsafe pointer value and its reflect.StructField in order
// to provide enough information
func UnsafeRideFields(
	v reflect.Value,
	fn func(sf reflect.StructField, ptr unsafe.Pointer) bool,
) {
	if !v.CanAddr() {
		v = ForceOfUnaddr(v)
	}

	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		sf := t.Field(i)
		ptr := unsafe.Pointer(v.Field(i).UnsafeAddr())

		if !fn(sf, ptr) {
			break
		}
	}
}

// UnsafeDeepRideFields rides a struct deeply using the unsafe package to allow access to
// its unexported field4.
func UnsafeDeepRideFields(v reflect.Value, fn functions.BiConsumer[string, unsafe.Pointer]) {
	if !v.CanAddr() {
		v = ForceOfUnaddr(v)
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
// to unexported field4.
func UnsafeSetFieldByName(target, value reflect.Value, name string) error {
	field := target.FieldByName(name)
	return UnsafeSet(field, value)
}

// UnsafeSetFieldByIndex sets a struct field by index using the unsafe package to allow access
// to unexported field4.
func UnsafeSetFieldByIndex(target, value reflect.Value, idx ...int) error {
	field := target.FieldByIndex(idx)
	return UnsafeSet(field, value)
}
