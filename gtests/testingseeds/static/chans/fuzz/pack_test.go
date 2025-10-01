// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package fuzz

import (
	"testing"
)

func TestAll(t *testing.T) {
	//assertlite.NoPanic(t, func() {
	_ = CatRefs{}.All()
	_ = CatValues{}.All()
	//})
}
