// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package field4

import (
	"reflect"
	"unsafe"

	"github.com/andrerrcosta2/gtools/core/domain/functions"
	"github.com/andrerrcosta2/gtools/core/format/fmx"
	"github.com/andrerrcosta2/gtools/reflect4/internal/values"
)

// UnsafeEach Iterates over all fields, exported and unexported and applies a function.
// It returns an error if the target isn't a struct
func UnsafeEach(target any, fn functions.BiPredicate[string, any]) error {
	v := values.Unwrap(reflect.ValueOf(target))
	if !v.IsValid() || v.Kind() != reflect.Struct {
		return ErrNotStruct("field4.UnsafeEach", values.Name(v))
	}
	values.UnsafeRideFields(v, func(sf reflect.StructField, ptr unsafe.Pointer) bool {
		val := reflect.NewAt(sf.Type, ptr).Elem().Interface()
		return fn(sf.Name, val)
	})
	return nil
}

// UnsafeSet sets the value of a struct field by name
// requires the target to be addressable (typically a pointer to a struct).
// If the target is not addressable, it returns an error.
func UnsafeSet(target any, name string, value any) error {
	v := values.Unwrap(reflect.ValueOf(target))
	if !v.IsValid() || v.Kind() != reflect.Struct {
		return ErrNotStruct("field4.UnsafeSet", values.Name(v))
	}

	return values.UnsafeFieldAccessIfAddr(v, name,
		func(field reflect.Value, ptr unsafe.Pointer) error {
			ft := field.Type()
			if value == nil {
				if !values.CanNil(field) {
					return fmx.Errorf("reflect4.UnsafeSet: cannot assign nil to type '%s'", ft)
				}
				reflect.NewAt(ft, ptr).Elem().Set(reflect.Zero(ft))
				return nil
			}
			vt := reflect.TypeOf(value)
			if !vt.AssignableTo(ft) {
				return fmx.Errorf("reflect4.UnsafeSet: cannot assign '%s' to '%s'", vt, ft)
			}
			addr := reflect.NewAt(ft, ptr).Elem()
			addr.Set(reflect.ValueOf(value))
			return nil
		},
	)
}

// UnsafeGet returns the value of a struct field by name
// It returns an error if the target is not a struct or the field is not found
//
// to avoid GC issues, the returned value is a clone of the original value
func UnsafeGet(target any, name string) (value any, err error) {
	v := values.Unwrap(reflect.ValueOf(target))
	if !v.IsValid() || v.Kind() != reflect.Struct {
		return nil,
			ErrNotStruct("field4.UnsafeGet", values.Name(v))
	}
	if !v.CanAddr() {
		v = values.ForceOfUnaddr(v)
	}
	err = values.UnsafeFieldAccess(v, name, func(field reflect.Value, ptr unsafe.Pointer) {
		// Use a read-only pointer to avoid GC issues.
		addr := reflect.NewAt(field.Type(), ptr)
		value = addr.Elem().Interface()
	})
	return value, err
}

// UnsafeGetAll returns all fields of a struct
// It returns an error if the target is not a struct
//
// to avoid GC issues, the returned value is a clone of the original value
func UnsafeGetAll(target any) (map[string]any, error) {
	v := values.Unwrap(reflect.ValueOf(target))
	if v.Kind() != reflect.Struct {
		return nil, ErrNotStruct("field4.UnsafeGetAll", values.Name(v))
	}
	out := make(map[string]any)
	values.UnsafeRideFields(v, func(sf reflect.StructField, ptr unsafe.Pointer) bool {
		out[sf.Name] = reflect.NewAt(sf.Type, ptr).Elem().Interface()
		return true
	})
	return out, nil
}

// UnsafeGetf returns a map of values of struct fields by names
// It returns an error if the target isn't a struct
//
// to avoid GC issues, the returned values are copies of the original values
func UnsafeGetf(target any, names ...string) (map[string]any, error) {
	v := values.Unwrap(reflect.ValueOf(target))
	if v.Kind() != reflect.Struct {
		return nil, ErrNotStruct("field4.UnsafeGetf", values.Name(v))
	}
	return values.UnsafeGetFields(v, names...)
}

// UnsafeAccess accesses a struct field by name using the unsafe package
func UnsafeAccess(target any, name string, fn functions.BiConsumer[any, unsafe.Pointer]) error {
	v := values.Unwrap(reflect.ValueOf(target))
	if v.Kind() != reflect.Struct {
		return ErrNotStruct("field4.UnsafeAccess", values.Name(v))
	}
	return values.UnsafeFieldAccess(v, name, func(field reflect.Value, ptr unsafe.Pointer) {
		fn(reflect.NewAt(field.Type(), ptr).Elem().Interface(), ptr)
	})
}
