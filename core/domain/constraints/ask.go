// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package constraints

import (
	"github.com/andrerrcosta2/gtools/core/domain/constraints/prim"
	"reflect"
)

// IsComparable checks if the provided values are from a
// naturally comparable type
//
//	(e.g.): x == z
func IsComparable(values ...any) bool {
	if prim.IsComparable(values) {
		return true
	}
	for _, v := range values {
		t := reflect.TypeOf(v)
		if t.Kind() == reflect.Invalid || !t.Comparable() {
			return false
		}
	}
	return true
}

func IsZeroOf[T any](value any) bool {
	var zero T
	if is, err := prim.Compare(value, zero); err == nil {
		return is
	}

	// Can't compare as primitives, pointers don't matter
	// since it can be nil
	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.Slice, reflect.Map:
		return v.Len() == 0
	case reflect.Struct:
		return reflect.DeepEqual(value, reflect.Zero(v.Type()).Interface())
	default:
		return false
	}
}
