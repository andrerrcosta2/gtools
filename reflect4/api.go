// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package reflect4

import (
	"github.com/andrerrcosta2/gtools/reflect4/internal/handlers/data"
	"github.com/andrerrcosta2/gtools/reflect4/read"
	"reflect"
)

// CopyOf creates a copy of a given value T.
// If `deep` is true, it performs a deep copy
// Otherwise, it performs a shallow copy.
func CopyOf[T any](v T, deep bool) (T, error) {
	if deep {
		// Perform a deep copy of the interf
		return data.DeepCopy(v)
	}

	// Perform a shallow copy of the interf
	return data.ShallowCopy(v)
}

func DeepEqual(a, b interface{}, opts ...read.Opt) (bool, []string) {
	var opt read.Opt
	for _, o := range opts {
		opt |= o
	}
	var diffs []string
	data.DeepEqual(reflect.ValueOf(a), reflect.ValueOf(b), data.Opt(opt), "", &diffs)
	return len(diffs) == 0, diffs
}
