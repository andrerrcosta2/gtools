// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package cat

import "reflect"

func CastMethod(t reflect.Type) RandomKind {
	switch t.Kind() {
	case reflect.Ptr, reflect.Chan, reflect.Map, reflect.Func, reflect.UnsafePointer:
		return Reference
	case reflect.Interface:
		if t.NumMethod() == 0 {
			return Any
		}
		return Injectable
	case reflect.Struct:
		return Struct
	}
	return Value
}

type RandomKind uint8

const (
	Injectable RandomKind = iota
	Value
	Struct
	Reference
	Any
)
