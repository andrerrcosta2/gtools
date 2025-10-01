// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package reflect4

import (
	"github.com/andrerrcosta2/gtools/core/testlite/assertlite"
	"github.com/andrerrcosta2/gtools/gtests"
	"github.com/andrerrcosta2/gtools/reflect4/op/clone"
	"github.com/andrerrcosta2/gtools/reflect4/op/read"
	"testing"
)

func TestCopyOf(t *testing.T) {
	gtests.Structs.Fuzz().Categories().Refs().All().
		Each(func(a any) {
			x, err := DeepCopy(&a, read.SkipUnexportedFields, clone.ChanIdentity)
			assertlite.NoError(t, err)
			assertlite.Equals(t, a, x)
		})
}
