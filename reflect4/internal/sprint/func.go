// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sprint

import (
	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"github.com/andrerrcosta2/gtools/core/format/sprints"
	"github.com/andrerrcosta2/gtools/reflect4/internal/reflect4"
	"reflect"
)

func Func(tab indent.Tab, v reflect.Value) (string, error) {
	if v.Kind() != reflect.Func {
		return sprints.Error(indent.Zero(), reflect4.ErrNotFunc.Error()), reflect4.ErrNotFunc
	}
	return fn(tab, v), nil
}
