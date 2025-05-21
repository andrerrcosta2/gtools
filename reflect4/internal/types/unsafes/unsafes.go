// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package unsafes

import (
	"github.com/andrerrcosta2/gtools/reflect4/internal/reflect4"
	"reflect"
)

func Name(value reflect.Type) (string, error) {
	if value.Kind() != reflect.UnsafePointer {
		return "", reflect4.ErrNotUnsafePtr
	}
	return value.String(), nil
}
