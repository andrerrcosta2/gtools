// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package zero

import (
	"testing"

	"github.com/andrerrcosta2/gtools/core/testlite/assertlite"
)

func TestAll(t *testing.T) {
	assertlite.NoPanics(t, func() {
		_ = CatRefs{}.All()
		_ = CatValues{}.All()
	})
}
