// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package zero

import (
	"github.com/andrerrcosta2/gtools/core/testlite/assertlite"
	"testing"
)

func TestAll(t *testing.T) {
	assertlite.NoPanic(t, func() {
		_ = Pack{}.Refs()
		_ = Pack{}.Values()
	})
}
