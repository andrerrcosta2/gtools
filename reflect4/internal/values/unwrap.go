// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package values

import "reflect"

// UnwrapInterfaces unwraps all interfaces from a value
func UnwrapInterfaces(v reflect.Value) reflect.Value {
	for v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	return v
}

// UnwrapPointers unwraps all pointers from a value
func UnwrapPointers(v reflect.Value) reflect.Value {
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	return v
}

// Unwrap unwraps all pointers and interfaces until it reaches a value
func Unwrap(v reflect.Value) reflect.Value {
	for v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	return v
}
