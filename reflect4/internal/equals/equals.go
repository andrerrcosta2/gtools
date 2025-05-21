// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package equals

import "reflect"

func Values(a, b reflect.Value) (bool, error) {
	if a.Kind() != b.Kind() {
		return false, nil
	}
	return a.Interface() == b.Interface(), nil
}
