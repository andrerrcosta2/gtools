// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sprint

import (
	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"github.com/andrerrcosta2/gtools/core/format/sprints"
	"github.com/andrerrcosta2/gtools/reflect4/internal/reflect4"
	"github.com/andrerrcosta2/gtools/reflect4/internal/tracker"
	"reflect"
)

// Pointer returns the string representation of a pointer value
func Pointer(tab indent.Tab, v reflect.Value) (string, error) {
	if v.Kind() != reflect.Ptr {
		return sprints.Error(indent.Zero(), reflect4.ErrNotPointer.Error()), reflect4.ErrNotPointer
	}
	return defaultPointer(tab, v, tracker.Sprint())
}
