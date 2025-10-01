// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package reflectutils

import (
	"github.com/andrerrcosta2/gtools/core/seeders/random/internal/prng"
	"reflect"
)

type KindSlice []reflect.Kind

func (x KindSlice) Len() int { return len(x) }

func (x KindSlice) Rand() reflect.Kind {
	return x[prng.Int(0, x.Len()-1)]
}

var (
	AllKinds = KindSlice{
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
		reflect.UnsafePointer,
	}

	CmpKinds = KindSlice{
		reflect.Bool,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
		reflect.Float32, reflect.Float64,
		reflect.Complex64, reflect.Complex128,
		reflect.String,
		reflect.Array,
		reflect.Ptr,
		reflect.UnsafePointer,
	}

	HashableKinds = KindSlice{
		reflect.Bool,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
		reflect.Float32, reflect.Float64,
		reflect.Complex64, reflect.Complex128,
		reflect.Array,
		reflect.Chan,
		reflect.Ptr,
		reflect.String,
		reflect.UnsafePointer,
	}

	HashableInterfaceKinds = KindSlice{
		reflect.Bool,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
		reflect.Float32, reflect.Float64,
		reflect.Complex64, reflect.Complex128,
		reflect.Array,
		reflect.Chan,
		reflect.Ptr,
		reflect.String,
		reflect.Struct,
		reflect.UnsafePointer,
	}

	InterfaceKinds = KindSlice{
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
		reflect.UnsafePointer,
	}

	RandomizableKinds = KindSlice{
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
		reflect.UnsafePointer,
	}
)

func isHashable(t reflect.Type) bool {
	switch t.Kind() {
	case reflect.Slice, reflect.Map, reflect.Func:
		return false

	case reflect.Array:
		return isHashable(t.Elem())

	case reflect.Struct:
		for i := 0; i < t.NumField(); i++ {
			if !isHashable(t.Field(i).Type) {
				return false
			}
		}
		return true

	case reflect.Interface:
		// The interface itself is hashable only if the assigned value is hashable
		// But for generation: assume we'll assign hashable values
		// So: interface{} is okay as a key type, as long as we don't put unhashable things in it
		return true

	default:
		return true // includes: bool, numbers, string, ptr, chan, unsafe.Pointer
	}
}

func IsAnyInterface(t reflect.Type) bool {
	return t.Kind() == reflect.Interface && t.NumMethod() == 0
}

func IsNullable(t reflect.Type) bool {
	switch t.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return true
	default:
		return false
	}
}

func TrackerOf(maxDepth ...int) *Tracker {
	if len(maxDepth) == 0 {
		return &Tracker{
			visit:    make(map[reflect.Type]int),
			maxDepth: 3,
		}
	}
	return &Tracker{
		visit:    make(map[reflect.Type]int),
		maxDepth: maxDepth[0],
	}
}

type Tracker struct {
	visit    map[reflect.Type]int
	maxDepth int
}

func (tt *Tracker) Next(t reflect.Type) bool {
	if tt.visit[t] < tt.maxDepth {
		tt.visit[t]++
		return true
	}
	return false
}
