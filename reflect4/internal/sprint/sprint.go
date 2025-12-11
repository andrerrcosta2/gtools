// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sprint

import (
	"reflect"

	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"github.com/andrerrcosta2/gtools/core/format/sprints"
	"github.com/andrerrcosta2/gtools/reflect4/internal"
)

func Of[O internal.Option](tab indent.Indentor, value reflect.Value, o ...O) string {
	return sprintOf(tab, value, NewStrategy(o...))
}

func sprintOf(tab indent.Indentor, value reflect.Value, s *Strategy) string {
	if !value.IsValid() {
		return "invalid type: <invalid>"
	}

	if value.Kind() <= reflect.Complex128 || value.Kind() == reflect.String {
		return sprintPrimitive(tab, value)
	}

	switch value.Kind() {
	case reflect.Array:
		return defaultArray(tab, value, s)
	case reflect.Chan:
		return defaultChan(tab, value)
	case reflect.Func:
		return defaultFunc(tab, value)
	case reflect.Interface:
		return defaultInterface(tab, value, s)
	case reflect.Map:
		return defaultMap(tab, value, s)
	case reflect.Ptr:
		return defaultPointer(tab, value, s)
	case reflect.Slice:
		return defaultSlice(tab, value, s)
	case reflect.Struct:
		return defaultStruct(tab, value, s)
	case reflect.UnsafePointer:
		return defaultUnsafe(tab, value)
	default:
		return sprints.Errorf(tab, "invalid type: %s", value.Kind().String())
	}
}
