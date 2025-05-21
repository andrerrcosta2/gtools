// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package values

import "reflect"

func CanNil(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.Ptr, reflect.Slice, reflect.Map, reflect.Chan, reflect.Func, reflect.Interface:
		return true
	default:
		return false
	}
}
