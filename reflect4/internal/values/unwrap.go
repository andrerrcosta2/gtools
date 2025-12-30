// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package values

import "reflect"

// UnwrapInterfaces unwraps all interfaces from a value
// It panics if the value is invalid
func UnwrapInterfaces(v reflect.Value) reflect.Value {
	for v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	return v
}

// UnwrapPointers unwraps all ptrs from a value
// It panics if the value is invalid
func UnwrapPointers(v reflect.Value) reflect.Value {
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	return v
}

// Unwrap unwraps all ptrs and interfaces until it reaches a value
// It panics if the value is invalid
func Unwrap(v reflect.Value) reflect.Value {
	for v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	return v
}
