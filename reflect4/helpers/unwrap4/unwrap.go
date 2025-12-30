// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package unwrap4

import (
	"reflect"

	"github.com/andrerrcosta2/gtools/reflect4/internal/reflect4"
	"github.com/andrerrcosta2/gtools/reflect4/internal/types"
	"github.com/andrerrcosta2/gtools/reflect4/internal/values"
)

// ToValue
func ToValue(value any) (reflect.Value, error) {
	if value == nil {
		return reflect.Value{}, reflect4.ErrNilInterface("unwrap4.ToValue")
	}
	v := values.Unwrap(reflect.ValueOf(value))
	if !v.IsValid() {
		return reflect.Value{}, reflect4.ErrNilInterface("unwrap4.ToValue")
	}
	return v, nil
}

func ToTypeValue(value any) (reflect.Type, error) {
	if value == nil {
		return nil, reflect4.ErrNilInterface("unwrap4.ToTypeValue")
	}
	return types.Unwrap(reflect.TypeOf(value)), nil
}
