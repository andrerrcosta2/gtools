// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package differ

import (
	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"github.com/andrerrcosta2/gtools/core/format/differs"
	"github.com/andrerrcosta2/gtools/reflect4/internal/tracker"
	"reflect"
)

// Struct returns the difference between two structs
func Struct(tab indent.Tab, a, b reflect.Value) (diff string, isDiff bool, err error) {
	if a.Kind() != reflect.Struct {
		return differs.TypesMismatch(tab, a.Kind().String(), reflect.Struct.String()), true, nil
	}
	if b.Kind() != reflect.Struct {
		return differs.TypesMismatch(tab, b.Kind().String(), reflect.Struct.String()), true, nil
	}
	differ := st(tab, a, b, tracker.Diff())
	message := differs.Append(tab, differs.StructFields(tab, a.Type().Field(0).Name), differ.Message)
	return differs.Message(message, tab.Sprint(differ.Diff)), differ.Equals, differ.Err
}
