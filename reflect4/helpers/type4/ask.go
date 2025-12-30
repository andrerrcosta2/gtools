// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package type4

import (
	"reflect"

	"github.com/andrerrcosta2/gtools/reflect4/internal/reflect4"
	"github.com/andrerrcosta2/gtools/reflect4/internal/types"
)

// HasDepth returns true if the type has a depth and false if it
// is a direct value
func HasDepth(v any) (bool, error) {
	t := reflect.TypeOf(v)
	if t == nil {
		return false, reflect4.ErrNilInterface("type4.HasDepth")
	}
	if t.Kind() == reflect.Invalid {
		return false, reflect4.ErrInvalidValue
	}
	return types.HasDepth(t), nil
}

// HasRecursiveRef reports whether type t has recursive references.
func HasRecursiveRef(v any) (bool, error) {
	t := reflect.TypeOf(v)
	if t == nil {
		return false, reflect4.ErrNilInterface("type4.HasRecursiveRef")
	}
	if t.Kind() == reflect.Invalid {
		return false, reflect4.ErrInvalidValue
	}
	return types.HasRecursiveRef(t), nil
}
