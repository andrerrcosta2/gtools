// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package reflect4

import (
	"github.com/andrerrcosta2/gtools/core/format/fmx"
	"github.com/andrerrcosta2/gtools/reflect4/internal/handlers/data"
	"github.com/andrerrcosta2/gtools/reflect4/opts/read"
	"reflect"
)

// CopyOf creates a copy of a given value T.
// If `deep` is true, it performs a deep copy
// Otherwise, it performs a shallow copy.
func CopyOf[T any](t T, deep bool) (cp T, err error) {
	var val reflect.Value
	var ok bool
	if deep {
		val, err = data.DeepCopy(reflect.ValueOf(t))
	} else {
		val, err = data.ShallowCopy(reflect.ValueOf(t))
	}
	if err != nil {
		return
	}
	cp, ok = val.Interface().(T)
	if !ok {
		return cp, fmx.Errorf("unable to cast value to %T", t)
	}
	return
}

func DeepEqual(a, b any, opts ...read.Opt) (bool, []string) {
	var opt read.Opt
	for _, o := range opts {
		opt |= o
	}
	var diffs []string
	data.DeepEqual(reflect.ValueOf(a), reflect.ValueOf(b), data.Opt(opt), "", &diffs)
	return len(diffs) == 0, diffs
}
