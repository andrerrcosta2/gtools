// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sprint

import (
	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"github.com/andrerrcosta2/gtools/core/format/sprints"
	"github.com/andrerrcosta2/gtools/reflect4/internal/reflect4"
	"github.com/andrerrcosta2/gtools/reflect4/internal/tracker"
	"reflect"
)

func Slice(tab indent.Tab, v reflect.Value) (string, error) {
	if v.Kind() != reflect.Slice {
		return sprints.Error(indent.Zero(), reflect4.ErrNotSlice.Error()), reflect4.ErrNotSlice
	}
	return sprintSlice(tab, v, tracker.Sprint())
}
