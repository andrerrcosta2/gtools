// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package reflect4

import "reflect"

// UnwrapInterfaces unwraps all interfaces from a value
// until the result kind is not an interface
func UnwrapInterfaces(v reflect.Value) reflect.Value {
	for v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	return v
}

// UnwrapPointers unwraps all pointers from a value
// until the result kind is not a pointer
func UnwrapPointers(v reflect.Value) reflect.Value {
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	return v
}

// UnwrapIntoValue unwraps all pointers and interfaces
// until it reaches a value
func UnwrapIntoValue(v reflect.Value) reflect.Value {
	for v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	return v
}
