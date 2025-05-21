// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package fields

import (
	"github.com/andrerrcosta2/gtools/core/domain/functions"
	"github.com/andrerrcosta2/gtools/reflect4/internal/values"
	"reflect"
	"unsafe"
)

// UnsafeSet sets the value of a struct field by name
// It returns an error if the target is not a struct or the field is not found
func UnsafeSet(target any, name string, value any) error {
	return values.UnsafeFieldAccess(reflect.ValueOf(target), name, func(field reflect.Value, ptr unsafe.Pointer) {
		addr := reflect.NewAt(field.Type(), ptr).Elem()
		addr.Set(reflect.ValueOf(value))
	})
}

// UnsafeGet returns the value of a struct field by name
// It returns an error if the target is not a struct or the field is not found
//
// In order to avoid GC issues, the returned value is a copy of the original value
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
// In order to avoid GC issues, the returned value is a copy of the original value
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
// It returns an error if the target is not a struct
//
// In order to avoid GC issues, the returned values are copies of the original values
func UnsafeGetf(target any, names ...string) (map[string]any, error) {
	return values.UnsafeGetFields(reflect.ValueOf(target), names...)
}

// UnsafeAccess accesses a struct field by name using the unsafe package
func UnsafeAccess(target any, name string, fn functions.BiConsumer[any, unsafe.Pointer]) error {
	return values.UnsafeFieldAccess(reflect.ValueOf(target), name, func(field reflect.Value, ptr unsafe.Pointer) {
		fn(field.Interface(), ptr)
	})
}
