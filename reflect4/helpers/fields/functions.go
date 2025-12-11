// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package fields

import (
	"errors"
	"reflect"

	"github.com/andrerrcosta2/gtools/core/format/fmx"
	"github.com/andrerrcosta2/gtools/reflect4/helpers/unwrap"
	"github.com/andrerrcosta2/gtools/reflect4/internal/types"
	"github.com/andrerrcosta2/gtools/reflect4/internal/types/structs"
	"github.com/andrerrcosta2/gtools/reflect4/internal/values"
)

var ErrNotStruct = errors.New("fields.GetValue - target is not a struct")
var ErrFieldNotFound = errors.New("fields.GetValue - field not found")
var ErrUnexportedField = errors.New("fields.GetValue - field is not exported")

// Each Iterates over exported fields and applies a function.
// It returns an error if the target isn't a struct
// It ignores unexported fields
//
// If you need to access unexported fields, use UnsafeEach function
func Each(target any, fn func(name string, value any)) error {
	v := unwrap.ToValue(target)
	if v.Kind() != reflect.Struct {
		return fmx.Errorf("%s: %s", ErrNotStruct.Error(), v.Type().String())
	}

	for i := 0; i < v.NumField(); i++ {
		ft := v.Type().Field(i)
		if ft.IsExported() {
			fn(ft.Name, v.Field(i).Interface())
		}
	}
	return nil
}

// From returns a slice of the values of all the fields of a struct
// It returns an error if the target isn't a struct
// It ignores unexported fields
//
// If you need to access unexported fields, use UnsafeFrom function
func From(target any) ([]any, error) {
	v := unwrap.ToValue(target)
	if v.Kind() != reflect.Struct {
		return nil, fmx.Errorf("%s: %s", ErrNotStruct.Error(), v.Type().String())
	}
	fm := values.Fields(v)

	fields := make([]any, len(fm))
	for _, field := range fm {
		fields = append(fields, field.Interface())
	}
	return fields, nil
}

// Get returns the value of a struct field by name
// It returns an error if the target isn't a struct.
//
// This method only works with exported fields. If you need to access unexported fields,
// use UnsafeGet function
func Get(target any, name string) (any, error) {
	v := values.Unwrap(reflect.ValueOf(target))
	value, err := values.Field(v, name)
	if err != nil {
		return nil, err
	}
	return value.Interface(), nil
}

// Has returns true if the struct has a field with the given name.
// It returns an error if the target is not a struct
func Has(target any, name string) (bool, error) {
	t := unwrap.ToTypeValue(target)
	if t.Kind() != reflect.Struct {
		return false, fmx.Errorf("%s: %s", ErrNotStruct.Error(), t.String())
	}
	if t.Kind() != reflect.Struct {
		return false, ErrNotStruct
	}
	_, ok := t.FieldByName(name)
	return ok, nil
}

// Names Returns the list of exported field names.
func Names(target any) ([]string, error) {
	v := unwrap.ToValue(target)
	if v.Kind() != reflect.Struct {
		return nil, fmx.Errorf("%s: %s", ErrNotStruct.Error(), v.Type().String())
	}
	fields := values.Fields(v)

	names := make([]string, len(fields))
	for _, field := range fields {
		names = append(names, field.Type().Name())
	}
	return names, nil
}

// Nil returns all nil field names. If the target isn't a struct, an error is returned.
func Nil(target any) (names []string, err error) {
	v := unwrap.ToValue(target)
	if v.Kind() != reflect.Struct {
		return nil, fmx.Errorf("%s: %s", ErrNotStruct.Error(), v.Type().String())
	}
	names = make([]string, 0, v.NumField())
	values.AccessFields(v, func(_ string, field reflect.Value) {
		if values.IsNil(field) {
			names = append(names, field.Type().Name())
		}
	})
	return
}

// NotNil returns all exported not nil fields within a struct
func NotNil(target any) (fields []any, err error) {
	v := unwrap.ToValue(target)
	if v.Kind() != reflect.Struct {
		return nil, fmx.Errorf("%s: %s", ErrNotStruct.Error(), v.Type().String())
	}
	fields = make([]any, 0, v.NumField())
	values.AccessFields(v, func(_ string, field reflect.Value) {
		if !values.IsNil(field) {
			fields = append(fields, field.Interface())
		}
	})
	return

}

// NumExported returns the number of exported fields. If the target isn't a struct, an error is returned.
func NumExported(target any) (int, error) {
	t := types.Unwrap(reflect.TypeOf(target))
	return structs.NumExportedFields(t)
}

// NumUnexported returns the number of unexported fields. If the target isn't a struct, an error is returned.
func NumUnexported(target any) (int, error) {
	t := types.Unwrap(reflect.TypeOf(target))
	return structs.NumUnexportedFields(t)
}

// Set sets the field with the given name to the given value. If the target isn't a struct, an error is returned.
// This method only works with exported fields. For unexported fields use UnsafeSet
func Set(target any, name string, value any) error {
	return values.SetField(reflect.ValueOf(target), name, reflect.ValueOf(value))
}

// Tags returns a map of fieldName -> tagValue for a given struct tag key (e.g., "json", "db").
// It returns an error if the target isn't struct
func Tags(target any, key string) (map[string]string, error) {
	t := types.Unwrap(reflect.TypeOf(target))
	return structs.Tags(t, key)
}

// Zero Sets the specified fields to their zero value.
// It returns an error if the target isn't struct
func Zero(target any, names ...string) error {
	v := reflect.ValueOf(target)
	if v.Kind() != reflect.Pointer || v.Elem().Kind() != reflect.Struct {
		return ErrNotStruct
	}
	v = v.Elem()

	for _, name := range names {
		field := v.FieldByName(name)
		if !field.IsValid() {
			return ErrFieldNotFound
		}
		if !field.CanSet() {
			return ErrUnexportedField
		}
		field.Set(reflect.Zero(field.Type()))
	}
	return nil
}
