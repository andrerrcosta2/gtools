// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package values

import (
	"errors"
	"github.com/andrerrcosta2/gtools/core/domain/functions"
	"reflect"
	"unsafe"
)

var (
	ErrNotStruct       = errors.New("reflect4.GetField - target is not a struct")
	ErrFieldNotFound   = errors.New("reflect4.GetField - field not found")
	ErrUnexportedField = errors.New("reflect4.GetField - field is not exported")
)

func Fields(target reflect.Value) (fields []reflect.Value, err error) {
	err = AccessFields(target, func(field reflect.Value) {
		fields = append(fields, field)
	})
	return
}

func Field(target reflect.Value, name string) (value reflect.Value, err error) {
	err = AccessField(target, name, func(field reflect.Value) {
		value = field
	})
	return
}

// UnsafeGetField returns a struct field by name using the unsafe package
//
// In order to avoid GC issues, the returned value is a copy of the original value
func UnsafeGetField(v reflect.Value, name string) (value reflect.Value, err error) {
	err = UnsafeFieldAccess(v, name, func(field reflect.Value, ptr unsafe.Pointer) {
		value = reflect.NewAt(field.Type(), ptr).Elem()
	})
	return
}

// UnsafeGetAllFields returns all unexported fields of a struct
//
// In order to avoid GC issues, the returned value is a copy of the original value
func UnsafeGetAllFields(v reflect.Value) (fields map[string]reflect.Value, err error) {
	tv := Unwrap(v)
	if tv.Kind() != reflect.Struct {
		return nil, ErrNotStruct
	}

	if !tv.CanAddr() {
		tv = OfUnaddr(tv)
	}
	fields = make(map[string]reflect.Value)
	for i := 0; i < tv.NumField(); i++ {
		field := tv.Field(i)
		name := tv.Type().Field(i).Name
		if !field.IsValid() {
			return nil, ErrFieldNotFound
		}
		ptr := unsafe.Pointer(field.UnsafeAddr())
		fields[name] = reflect.NewAt(field.Type(), ptr).Elem()
	}
	return
}

// UnsafeGetFields returns a map of values of struct fields by names
//
// In order to avoid GC issues, the returned values are copies of the original values
func UnsafeGetFields(v reflect.Value, names ...string) (fields map[string]any, err error) {
	tv := Unwrap(v)
	if tv.Kind() != reflect.Struct {
		return nil, ErrNotStruct
	}

	for _, name := range names {
		field := tv.FieldByName(name)

		if !field.IsValid() {
			return nil, ErrFieldNotFound
		}

		ptr := unsafe.Pointer(field.UnsafeAddr())
		fields[name] = reflect.NewAt(field.Type(), ptr).Elem().Interface()
	}
	return
}

func SetField(target reflect.Value, name string, value reflect.Value) error {
	return AccessField(target, name, func(field reflect.Value) {
		field.Set(value)
	})
}

func AccessFields(target reflect.Value, fn functions.Consumer[reflect.Value]) error {
	tv := Unwrap(target)
	if tv.Kind() != reflect.Struct {
		return ErrNotStruct
	}
	for i := 0; i < tv.NumField(); i++ {
		fn(tv.Field(i))
	}
	return nil
}

func AccessField(target reflect.Value, name string, fn functions.Consumer[reflect.Value]) error {
	tv := Unwrap(target)
	if tv.Kind() != reflect.Struct {
		return ErrNotStruct
	}

	field := tv.FieldByName(name)

	if !field.IsValid() {
		return ErrFieldNotFound
	}

	if !field.CanSet() {
		return ErrUnexportedField
	}

	fn(field)
	return nil
}

// UnsafeFieldAccess accesses a struct field by name using the unsafe package
func UnsafeFieldAccess(target reflect.Value, name string, fn functions.BiConsumer[reflect.Value, unsafe.Pointer]) error {
	tv := Unwrap(target)
	if tv.Kind() != reflect.Struct {
		return ErrNotStruct
	}

	field := tv.FieldByName(name)

	if !field.IsValid() {
		return ErrFieldNotFound
	}

	fn(field, unsafe.Pointer(field.UnsafeAddr()))
	return nil
}

// Ride accesses a struct field by name
func Ride(target reflect.Value, fn functions.BiConsumer[string, reflect.Value]) error {
	tv := Unwrap(target)
	if tv.Kind() != reflect.Struct {
		return ErrNotStruct
	}
	for i := 0; i < tv.NumField(); i++ {
		fn(tv.Type().Field(i).Name, tv.Field(i))
	}
	return nil
}

// UnsafeRide accesses a struct field by name using the unsafe package
func UnsafeRide(target reflect.Value, fn functions.Consumer[unsafe.Pointer]) error {
	tv := Unwrap(target)
	if tv.Kind() != reflect.Struct {
		return ErrNotStruct
	}

	for i := 0; i < tv.NumField(); i++ {
		fn(unsafe.Pointer(tv.Field(i).UnsafeAddr()))
	}
	return nil
}
