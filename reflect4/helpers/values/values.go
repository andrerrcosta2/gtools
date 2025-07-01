// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package values

import "reflect"

func Of[T any](args ...T) []reflect.Value {
	values := make([]reflect.Value, len(args))
	for i, arg := range args {
		values[i] = reflect.ValueOf(arg)
	}
	return values
}

func To[T any](values []reflect.Value) []T {
	args := make([]T, len(values))
	for i, value := range values {
		args[i] = value.Interface()
	}
	return args
}
