// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package unwrap

import (
	"github.com/andrerrcosta2/gtools/reflect4/internal/types"
	"github.com/andrerrcosta2/gtools/reflect4/internal/values"
	"reflect"
)

func FromInterfacesToValue(value any) reflect.Value {
	v := reflect.ValueOf(value)
	for v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	return v
}

func FromPointersToValue(value any) reflect.Value {
	v := reflect.ValueOf(value)
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	return v
}

func ToValue(value any) reflect.Value {
	return values.Unwrap(reflect.ValueOf(value))
}

func ToTypeValue(value any) reflect.Type {
	return types.Unwrap(reflect.TypeOf(value))
}
