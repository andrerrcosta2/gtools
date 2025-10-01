// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package differ

import (
	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"github.com/andrerrcosta2/gtools/core/format/differs"
	"reflect"
)

// Struct returns the defaultStringDiff between two structs
func Struct(tab indent.Tab, a, b reflect.Value, s *Strategy) (diff string, isDiff bool, err error) {
	if a.Kind() != reflect.Struct {
		return differs.TypesMismatch(tab, a.Kind().String(), reflect.Struct.String()), true, nil
	}
	if b.Kind() != reflect.Struct {
		return differs.TypesMismatch(tab, b.Kind().String(), reflect.Struct.String()), true, nil
	}
	if !a.CanAddr() {
		if !a.CanInterface() {
			return "", false, ErrUnexportedFieldStruct
		}
		ptr := reflect.New(a.Type())
		ptr.Elem().Set(a)
		a = ptr.Elem()
	}
	if !b.CanAddr() {
		if !b.CanInterface() {
			return "", false, ErrUnexportedFieldStruct
		}
		ptr := reflect.New(b.Type())
		ptr.Elem().Set(b)
		b = ptr.Elem()
	}
	differ := diffStructs(tab, a, b, s)
	message := differs.Append(tab, differs.StructFields(tab, a.Type().Field(0).Name), differ.Message)
	return differs.Message(message, tab.Sprint(differ.Diff)), differ.Equals, differ.Err
}
