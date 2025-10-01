// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package differ

import (
	"errors"
	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"github.com/andrerrcosta2/gtools/core/format/differs"
	"github.com/andrerrcosta2/gtools/reflect4/internal"
	"reflect"
)

var (
	ErrUnexportedFieldStruct = errors.New("cannot diff between structs as unexported fields")
)

// Between returns the difference between two values
func Between[O internal.Option](value, expected reflect.Value, o ...O) (diff string, equals bool) {
	//if value.Kind() != reflect.Ptr || expected.Kind() != reflect.Ptr {
	//	panic(fmx.Sprintf("can't differ safely on non pointer values. got '%v' and '%v'",
	//		value.Kind(), expected.Kind()))
	//}
	var tab indent.Branch
	d := between(tab, value, expected, NewStrategy(o...))
	if !d.Equals {
		return differs.Message(d.Message, d.Diff), false
	}
	return "", true
}
