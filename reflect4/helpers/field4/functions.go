// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package field4

import (
	"errors"
	"reflect"

	"github.com/andrerrcosta2/gtools/reflect4/internal/reflect4"
	"github.com/andrerrcosta2/gtools/reflect4/internal/types"
	"github.com/andrerrcosta2/gtools/reflect4/internal/types/structs"
	"github.com/andrerrcosta2/gtools/reflect4/internal/values"
)

var (
	ErrNotStruct = func(caller, typ string) error {
		return errors.New(caller + ": target is not a struct, it is a '" + typ + "'")
	}
	ErrFieldNotFound = func(caller, field string) error {
		return errors.New(caller + ": field '" + field + "' was not found")
	}
	ErrUnexportedField = func(caller, field string) error {
		return errors.New(caller + ": field '" + field + "' is not exported")
	}
)

// EachExp Iterates over exported field4 and applies a function.
// It returns an error if the target isn't a struct
// It ignores unexported field4
//
// If you need to access unexported field4, use UnsafeEach function
func EachExp(target any, fn func(name string, value any) bool) error {
	v := values.Unwrap(reflect.ValueOf(target))
	if !v.IsValid() || v.Kind() != reflect.Struct {
		return ErrNotStruct("field4.EachExp", values.Name(v))
	}
	values.RideExportedFieldsByName(v, func(name string, value reflect.Value) bool {
		return fn(name, value.Interface())
	})
	return nil
}

// FromExp returns a slice of the values of all the exported field4 of a struct
// It returns an error if the target isn't a struct
// It ignores unexported field4
//
// If you need to access unexported field4, use UnsafeFrom function
func FromExp(target any) ([]any, error) {
	v := values.Unwrap(reflect.ValueOf(target))
	if !v.IsValid() || v.Kind() != reflect.Struct {
		return nil, ErrNotStruct("field4.FromExp", values.Name(v))
	}
	fields := make([]any, 0, v.NumField())
	values.RideExportedFieldsByName(v, func(_ string, value reflect.Value) bool {
		fields = append(fields, value.Interface())
		return true
	})
	return fields, nil
}

// GetExp returns the value of an exported struct field by name
// It returns an error if the target isn't a struct.
//
// This method only works with exported field4. If you need to access unexported field4,
// use UnsafeGet function
func GetExp(target any, name string) (any, error) {
	v := values.Unwrap(reflect.ValueOf(target))
	if !v.IsValid() || v.Kind() != reflect.Struct {
		return nil, ErrNotStruct("field4.GetExp", values.Name(v))
	}
	value, err := values.FieldExp(v, name)
	if err != nil {
		return nil, err
	}
	return value.Interface(), nil
}

// Has returns true if the struct has a field with the given name.
// It returns an error if the target is not a struct
func Has(target any, name string) (bool, error) {
	t := reflect.TypeOf(target)
	if t == nil {
		return false, reflect4.ErrNilInterface("field4.Has")
	}
	t = types.Unwrap(t)
	if t.Kind() != reflect.Struct {
		return false, ErrNotStruct("field4.Has", types.Name(t))
	}
	_, ok := t.FieldByName(name)
	return ok, nil
}

// Names returns the list of all field names.
func Names(target any) ([]string, error) {
	v := values.Unwrap(reflect.ValueOf(target))
	if !v.IsValid() || v.Kind() != reflect.Struct {
		return nil, ErrNotStruct("field4.Names", values.Name(v))
	}
	t := v.Type()
	names := make([]string, 0, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		names = append(names, t.Field(i).Name)
	}
	return names, nil
}

// Nil returns all field names whose values are nil.
// Only field4 of nil-able type4 (ptr, slice, map, chan, func, interface) can be nil.
// Returns an error if target is not a struct.
func Nil(target any) ([]string, error) {
	v := values.Unwrap(reflect.ValueOf(target))
	if !v.IsValid() || v.Kind() != reflect.Struct {
		return nil, ErrNotStruct("field4.Nil", values.Name(v))
	}
	names := make([]string, 0, v.NumField())
	values.RideFieldsByName(v, func(name string, field reflect.Value) bool {
		if values.IsNil(field) {
			names = append(names, name)
		}
		return true
	})
	return names, nil
}

// NotNilExp returns all exported not nil field4 within a struct
func NotNilExp(target any) (fields []any, err error) {
	v := values.Unwrap(reflect.ValueOf(target))
	if !v.IsValid() || v.Kind() != reflect.Struct {
		return nil, ErrNotStruct("field4.NotNilExp", values.Name(v))
	}
	fields = make([]any, 0, v.NumField())
	values.RideExportedFieldsByName(v, func(_ string, field reflect.Value) bool {
		if !values.IsNil(field) {
			fields = append(fields, field.Interface())
		}
		return true
	})
	return
}

// NumExported returns the number of exported field4. If the target isn't a struct, an error is returned.
func NumExported(target any) (int, error) {
	t := reflect.TypeOf(target)
	if t == nil {
		return 0, reflect4.ErrNilInterface("field4.NumExported")
	}
	t = types.Unwrap(t)
	if t.Kind() != reflect.Struct {
		return 0, ErrNotStruct("field4.NumExported", types.Name(t))
	}
	return structs.NumExportedFields(t), nil
}

// NumUnexported returns the number of unexported field4. If the target isn't a struct, an error is returned.
func NumUnexported(target any) (int, error) {
	t := reflect.TypeOf(target)
	if t == nil {
		return 0, reflect4.ErrNilInterface("field4.NumUnexported")
	}
	t = types.Unwrap(t)
	if t.Kind() != reflect.Struct {
		return 0, ErrNotStruct("field4.NumUnexported", types.Name(t))
	}
	return structs.NumUnexportedFields(t), nil
}

// SetExp sets the field with the given name to the given value. If the target isn't a struct, an error is returned.
// This method only works with exported field4. For unexported field4 use UnsafeSet
func SetExp(target any, name string, value any) error {
	v := values.Unwrap(reflect.ValueOf(target))
	if !v.IsValid() || v.Kind() != reflect.Struct {
		return ErrNotStruct("field4.SetExp", values.Name(v))
	}
	return values.SetExpFieldByName(v, name, reflect.ValueOf(value))
}

// Tags returns a map of fieldName -> tagValue for a given struct tag key (e.g., "json", "db").
// It returns an error if the target isn't struct
func Tags(target any, key string) (map[string]string, error) {
	t := reflect.TypeOf(target)
	if t == nil {
		return nil, reflect4.ErrNilInterface("field4.Tags")
	}
	t = types.Unwrap(t)
	if t.Kind() != reflect.Struct {
		return nil, ErrNotStruct("field4.Tags", types.Name(t))
	}
	return structs.Tags(t, key), nil
}
