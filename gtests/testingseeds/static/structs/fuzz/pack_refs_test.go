// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package fuzz

import (
	"github.com/andrerrcosta2/gtools/core/testlite/assertlite"
	"testing"
)

func TestFuzzing_Refs(t *testing.T) {
	Pack{}.Categories().Refs().All().Each(func(a any) {
		assertlite.NoNilNonInterfaceFields(t, true, a)
	})
}
