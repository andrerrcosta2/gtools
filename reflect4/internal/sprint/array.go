// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sprint

import (
	"reflect"

	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"github.com/andrerrcosta2/gtools/core/format/sprints"
	"github.com/andrerrcosta2/gtools/reflect4/internal"
	"github.com/andrerrcosta2/gtools/reflect4/internal/reflect4"
	"github.com/andrerrcosta2/gtools/reflect4/internal/values"
)

func Array[O internal.Option](v reflect.Value, o ...O) (string, error) {
	if v.Kind() != reflect.Array {
		return sprints.Error(indent.Zero(), reflect4.ErrNotArray.Error()), reflect4.ErrNotArray
	}
	return defaultArray(indent.Zero(), values.ForceOfUnaddr(v), NewStrategy(o...)), nil
}
