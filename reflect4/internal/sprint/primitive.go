// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sprint

import (
	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"github.com/andrerrcosta2/gtools/core/format/sprints"
	"github.com/andrerrcosta2/gtools/reflect4/internal/reflect4"
	"reflect"
)

// SprintPrimitive returns the string representation of a primitive value
func SprintPrimitive(tab indent.Tab, v reflect.Value) (string, error) {
	if v.Kind() == reflect.Invalid {
		return sprints.Error(indent.Zero(), reflect4.ErrInvalidValue.Error()), reflect4.ErrInvalidValue
	}
	if v.Kind() > reflect.Complex128 && v.Kind() != reflect.String {
		return sprints.Error(indent.Zero(), reflect4.ErrNotPrimitive.Error()), reflect4.ErrNotPrimitive
	}
	return pr(tab, v), nil
}
