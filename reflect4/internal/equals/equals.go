// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package equals

import (
	"reflect"

	"github.com/andrerrcosta2/gtools/core/format/fmx"
	"github.com/andrerrcosta2/gtools/reflect4/internal"
)

func Deep[O internal.Option](a, b reflect.Value, o ...O) bool {
	return deep(a, b, NewStrategy(o...))
}

func deep(a, b reflect.Value, differ *Strategy) bool {
	if !a.IsValid() || !b.IsValid() {
		return a.IsValid() == b.IsValid()
	}
	if a.Type() != b.Type() {
		return false
	}
	switch a.Kind() {
	case reflect.Array:
		return differ.arrays(a, b, differ)
	case reflect.Slice:
		return differ.slices(a, b, differ)
	case reflect.Interface:
		return differ.interfaces(a, b, differ)
	case reflect.Pointer:
		return differ.pointers(a, b, differ)
	case reflect.Struct:
		return differ.structs(a, b, differ)
	case reflect.Map:
		return differ.maps(a, b, differ)
	case reflect.Func:
		return differ.functions(a, b)
	case reflect.String:
		return differStrings(a, b, differ)
	case reflect.Chan:
		return differ.channels(a, b)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return a.Int() == b.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return a.Uint() == b.Uint()
	case reflect.Bool:
		return a.Bool() == b.Bool()
	case reflect.Float32, reflect.Float64:
		return a.Float() == b.Float()
	case reflect.Complex64, reflect.Complex128:
		return a.Complex() == b.Complex()
	case reflect.UnsafePointer:
		return differ.unsafe(a, b)
	default:
		// shouldn't reach this
		panic(fmx.Sprintf("unreachable kind %v", a.Kind()))
		return false
	}
}
