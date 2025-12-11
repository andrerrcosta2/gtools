// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sprint

import (
	"reflect"

	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"github.com/andrerrcosta2/gtools/core/format/sprints"
	"github.com/andrerrcosta2/gtools/reflect4/internal"
	"github.com/andrerrcosta2/gtools/reflect4/internal/reflect4"
)

func Slice[O internal.Option](v reflect.Value, o ...O) (string, error) {
	if v.Kind() != reflect.Slice {
		return sprints.Error(indent.Zero(), reflect4.ErrNotSlice.Error()), reflect4.ErrNotSlice
	}
	return defaultSlice(indent.Zero(), v, NewStrategy(o...)), nil
}
