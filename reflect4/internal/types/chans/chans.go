// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package chans

import (
	"github.com/andrerrcosta2/gtools/reflect4/internal/reflect4"
	"reflect"
)

// Name returns the name of the channel
func Name(t reflect.Type) (string, error) {
	if t.Kind() != reflect.Chan {
		return "", reflect4.ErrNotChan
	}
	if t.Name() == "" {
		return t.String(), nil
	}
	pkgPath := t.PkgPath()
	if pkgPath != "" {
		return pkgPath + "." + t.Name(), nil
	}
	return t.Name(), nil
}
