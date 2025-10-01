// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package fields

import (
	"github.com/andrerrcosta2/gtools/core/domain/functions"
	"github.com/andrerrcosta2/gtools/core/format/fmx"
	"github.com/andrerrcosta2/gtools/reflect4/helpers/unwrap"
	"github.com/andrerrcosta2/gtools/reflect4/internal/values"
	"reflect"
	"unsafe"
)

// UnsafeEach Iterates over all fields, exported and unexported and applies a function.
// It returns an error if the target isn't a struct
func UnsafeEach(target any, fn functions.BiConsumer[string, any]) error {
	v := unwrap.ToValue(target)
	if v.Kind() != reflect.Struct {
		return fmx.Errorf("%s: %s", ErrNotStruct.Error(), v.Type().String())
	}
	if !v.CanAddr() {
		v = values.UnsafeOfUnaddr(v)
	}
	for i := 0; i < v.NumField(); i++ {
		fn(v.Type().Field(i).Name, unsafe.Pointer(v.Field(i).UnsafeAddr()))
	}
	return nil
}

// UnsafeSet sets the value of a struct field by name
// It returns an error if the target is not a struct or the field is not found
func UnsafeSet(target any, name string, value any) error {
	v := values.Unwrap(reflect.ValueOf(target))
	return values.UnsafeFieldAccess(v, name, func(field reflect.Value, ptr unsafe.Pointer) {
		addr := reflect.NewAt(field.Type(), ptr).Elem()
		addr.Set(reflect.ValueOf(value))
	})
}

// UnsafeGet returns the value of a struct field by name
// It returns an error if the target is not a struct or the field is not found
//
// to avoid GC issues, the returned value is a clone of the original value
func UnsafeGet(target any, name string) (value any, err error) {
	err = values.UnsafeFieldAccess(reflect.ValueOf(target), name, func(field reflect.Value, ptr unsafe.Pointer) {
		// Use a read-only pointer to avoid GC issues.
		addr := reflect.NewAt(field.Type(), ptr)
		value = addr.Elem().Interface()
	})
	return value, err
}

// UnsafeGetAll returns all unexported fields of a struct
// It returns an error if the target is not a struct
//
// to avoid GC issues, the returned value is a clone of the original value
func UnsafeGetAll(target any) (map[string]any, error) {
	out := make(map[string]any)
	m, err := values.UnsafeGetAllFields(reflect.ValueOf(target))
	if err != nil {
		return nil, err
	}
	for k, v := range m {
		out[k] = v.Interface()
	}
	return out, nil
}

// UnsafeGetf returns a map of values of struct fields by names
// It returns an error if the target isn't a struct
//
// to avoid GC issues, the returned values are copies of the original values
func UnsafeGetf(target any, names ...string) (map[string]any, error) {
	return values.UnsafeGetFields(reflect.ValueOf(target), names...)
}

// UnsafeAccess accesses a struct field by name using the unsafe package
func UnsafeAccess(target any, name string, fn functions.BiConsumer[any, unsafe.Pointer]) error {
	return values.UnsafeFieldAccess(reflect.ValueOf(target), name, func(field reflect.Value, ptr unsafe.Pointer) {
		fn(field.Interface(), ptr)
	})
}
