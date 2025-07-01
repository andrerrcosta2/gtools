// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package reflectutils

import "reflect"

var (
	AllKinds = []reflect.Kind{
		reflect.Bool,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
		reflect.Float32, reflect.Float64,
		reflect.Complex64, reflect.Complex128,
		reflect.Interface,
		reflect.String,
		reflect.Array,
		reflect.Func,
		reflect.Ptr,
		reflect.Chan,
		reflect.Slice,
		reflect.Map,
	}

	CmpKinds = []reflect.Kind{
		reflect.Bool,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
		reflect.Float32, reflect.Float64,
		reflect.Complex64, reflect.Complex128,
		reflect.String,
		reflect.Array,
		reflect.Ptr,
	}

	RandomizableKinds = []reflect.Kind{
		reflect.Bool,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
		reflect.Float32, reflect.Float64,
		reflect.Complex64, reflect.Complex128,
		reflect.Interface,
		reflect.String,
		reflect.Array,
		reflect.Func,
		reflect.Ptr,
		reflect.Chan,
		reflect.Slice,
		reflect.Map,
	}

	InterfaceKinds = []reflect.Kind{
		reflect.Bool,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
		reflect.Float32, reflect.Float64,
		reflect.Complex64, reflect.Complex128,
		reflect.String,
		reflect.Array,
		reflect.Func,
		reflect.Ptr,
		reflect.Chan,
		reflect.Slice,
		reflect.Map,
	}
)

func IsNullable(t reflect.Type) bool {
	switch t.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return true
	default:
		return false
	}
}
