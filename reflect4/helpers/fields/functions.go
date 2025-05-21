// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package fields

import (
	"errors"
	"github.com/andrerrcosta2/gtools/reflect4/internal/values"
	"reflect"
)

var ErrNotStruct = errors.New("fields.GetValue - target is not a struct")
var ErrFieldNotFound = errors.New("fields.GetValue - field not found")
var ErrUnexportedField = errors.New("fields.GetValue - field is not exported")

// From returns a slice of the values of all the fields of a struct
// It returns an error if the target is not a struct
// It ignores unexported fields
//
// If you need to access unexported fields, use UnsafeFrom function
func From(target any) ([]any, error) {
	fields, err := values.Fields(reflect.ValueOf(target))
	if err != nil {
		return nil, err
	}

	values := make([]any, len(fields))
	for i, field := range fields {
		values[i] = field.Interface()
	}
	return values, nil
}

// Get returns the value of a struct field by name
// It returns an error if the target is not a struct or the field is not found
// or the field is not exported
//
// If you need to access unexported fields, use UnsafeGet function
func Get(target any, name string) (any, error) {
	value, err := values.Field(reflect.ValueOf(target), name)
	if err != nil {
		return nil, err
	}
	return value.Interface(), nil
}

func Set(target any, name string, value any) error {
	return values.SetField(reflect.ValueOf(target), name, reflect.ValueOf(value))
}
