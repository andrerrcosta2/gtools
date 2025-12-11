// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package fuzz

import (
	"testing"

	"github.com/andrerrcosta2/gtools/core/format/fmx"
	"github.com/andrerrcosta2/gtools/core/testlite/assertlite"
)

func TestAll(t *testing.T) {
	assertlite.NoPanic(t, func() {
		_ = CatRefs{}.All().Each(Consume)
		_ = CatValues{}.All().Each(Consume)
	})
}

func Consume(a any) {
	_ = fmx.Sprintf("%+v", a)
}
