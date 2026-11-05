// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package fuzz

import (
	"github.com/andrerrcosta2/gtools/core/testlite/assertlite"
	"testing"
)

func TestFuzzing_Values(t *testing.T) {
	Pack{}.Categories().Values().All().Each(func(v any) {
		assertlite.NoNilFields(t, true, v)
	})
}
