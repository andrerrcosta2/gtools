// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package reflect4

//import (
//	"reflect"
//)
//
//// ValueOf returns a Value from the given interface value.
//// It first checks if the value implements the reflective interface, and if so,
//// it returns a Value with the reflective type.
//// If not, it checks if the value is a primitive type, and if so, it returns
//// a Value with the primitive kind.
//// If the value is not a primitive type nor a reflective interface, it returns
//// a Value with the reflected type.
//func ValueOf(v any) Value {
//	// If it's a primitive, return a Value with the primitive kind.
//	if kind, ptr := primKindOf(v); kind != reflect.Invalid && kind < reflect.Struct {
//		return fromPrimitiveType(kind, v, ptr)
//	}
//
//	// If it implements the reflective interface, return a Value with the reflective type
//	if refl, ok := asReflective(v); ok {
//		return fromReflectiveType(refl.Prototype(), ptrs.OfPointer[Reflective](refl))
//	}
//
//	// reflect to get the reflected type
//	return reflectedValueOf(v)
//}
//
//// reflectedValueOf returns a Value for the given `v` by using reflection.
//func reflectedValueOf(v any) Value {
//	r := reflect.ValueOf(v)
//	return Value{
//		reflect: r,
//		reflective: reflective{
//			reflected: true,
//		},
//		kind: r.Kind(),
//		ptr:  r.Kind() == reflect.Ptr,
//	}
//}
//
//func New(typx Type) Value {
//	// If it's a primitive, return a Value with the primitive kind.
//	if typx.Primitive() {
//		return fromPrimitiveType(typx.kind, nil, true)
//	}
//
//	// If it's a descriptor, return a Value with the descriptor
//	if typx.hasDescr() {
//		return fromReflectiveType(typx.descriptor, true)
//	}
//
//	// If it's a reflected type, return a Value with the reflected type
//	if typx.hasRefl() {
//		return fromReflectType(typx, true)
//	}
//	return invalidValue()
//}
//
//func fromReflectiveType(proto Prototype, ptr ptrs.Reflective) Value {
//	return Value{
//		reflective: reflective{
//			descriptor: proto,
//		},
//		kind: proto.Kind,
//		ptr:  ptr,
//	}
//}
//
//func fromPrimitiveType(k reflect.Kind, value any, ptr ptrs.Reflective) Value {
//	return Value{
//		kind:   k,
//		interf: value,
//		ptr:    ptr,
//	}
//}
//
//func fromReflectType(typx Type, ptr ptrs.Reflective) Value {
//	newValue := reflect.New(typx.reflect)
//	return Value{
//		reflect: newValue,
//		reflective: reflective{
//			reflected: true,
//		},
//		kind: typx.kind,
//		ptr:  ptr,
//	}
//}
//
//func invalidValue() Value {
//	return Value{
//		kind: reflect.Invalid,
//	}
//}
