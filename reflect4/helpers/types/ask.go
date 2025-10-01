// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package types

import (
	"github.com/andrerrcosta2/gtools/reflect4"
	"github.com/andrerrcosta2/gtools/reflect4/internal/types"
	"reflect"
)

// HasDepth returns true if the type has a depth and false if it is a direct value
func HasDepth(v any) (bool, error) {
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Invalid {
		return false, reflect4.ErrInvalidValue
	}
	return types.HasDepth(t), nil
}

// HasRecursiveRef reports whether type t has recursive references.
func HasRecursiveRef(v any) (bool, error) {
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Invalid {
		return false, reflect4.ErrInvalidValue
	}
	return types.HasRecursiveRef(t), nil
}
