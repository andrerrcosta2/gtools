// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package primitives

import "reflect"

func TypeOf(t reflect.Type) bool {
	return t.Kind() <= reflect.Complex128 || t.Kind() == reflect.String
}

func ValueOf(v reflect.Value) bool {
	return v.Kind() <= reflect.Complex128 || v.Kind() == reflect.String
}
