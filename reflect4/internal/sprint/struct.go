// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sprint

import (
	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"github.com/andrerrcosta2/gtools/core/format/sprints"
	"github.com/andrerrcosta2/gtools/reflect4/internal/reflect4"
	"github.com/andrerrcosta2/gtools/reflect4/internal/tracker"
	"github.com/andrerrcosta2/gtools/reflect4/internal/values"
	"reflect"
	"strings"
	"unsafe"
)

func Struct(tab indent.Tab, v reflect.Value) (string, error) {
	if v.Kind() != reflect.Struct {
		return sprints.Error(indent.Zero(), reflect4.ErrNotStruct.Error()), reflect4.ErrNotStruct
	}
	return sprintStruct(tab, v, tracker.Sprint())
}

// Fields returns a formatted sprint of the given struct fields
func Fields(tab indent.Tab, v reflect.Value) (string, error) {
	tv := values.Unwrap(v)
	if tv.Kind() != reflect.Struct {
		return "", reflect4.ErrNotStruct
	}
	if !tv.CanAddr() {
		tv = values.OfUnaddr(tv)
	}
	sb := strings.Builder{}
	for i := 0; i < tv.NumField(); i++ {
		field := tv.Field(i)
		name := tv.Type().Field(i).Name
		if !field.IsValid() {
			sb.WriteString(tab.Inc().Sprintf("%s <Invalid>", name))
			if i < tv.NumField()-1 {
				sb.WriteString("\n")
			}
			continue
		}
		ptr := unsafe.Pointer(field.UnsafeAddr())
		field = reflect.NewAt(field.Type(), ptr).Elem()
		sb.WriteString(tab.Inc().Sprintf("%s %s", name, field.Type().String()))
		if i < tv.NumField()-1 {
			sb.WriteString("\n")
		}
	}
	return sb.String(), nil
}
