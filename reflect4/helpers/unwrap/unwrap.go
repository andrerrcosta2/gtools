// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package unwrap

import "reflect"

func FromInterfaces(value any) reflect.Value {
	v := reflect.ValueOf(value)
	for v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	return v
}

func FromPointers(value any) reflect.Value {
	v := reflect.ValueOf(value)
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	return v
}

func ToValue(value any) reflect.Value {
	v := reflect.ValueOf(value)
	for v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	return v
}
