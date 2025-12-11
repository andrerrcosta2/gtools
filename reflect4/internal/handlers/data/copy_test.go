// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package data

import (
	"reflect"
	"testing"

	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"github.com/andrerrcosta2/gtools/core/testlite/assertlite"
	"github.com/andrerrcosta2/gtools/gtests"
	"github.com/andrerrcosta2/gtools/reflect4/internal"
	"github.com/andrerrcosta2/gtools/reflect4/internal/differ"
	"github.com/andrerrcosta2/gtools/reflect4/internal/equals"
	"github.com/andrerrcosta2/gtools/reflect4/internal/sprint"
	"github.com/andrerrcosta2/gtools/reflect4/op/read"
)

func TestDeepCopy(t *testing.T) {

}

func TestDeepCopyStruct(t *testing.T) {
	t.Run("default write", func(t *testing.T) {
		gtests.Structs.Fuzz().Categories().Refs().All().Each(func(ptr any) {
			value := reflect.ValueOf(ptr)
			cp, err := DeepCopy[internal.Option](value)
			assertlite.NoError(t, err, "error while deep copying:\n%v\n", err)
			if !equals.Deep(cp, value, read.SkipFunctions) {
				diff, _ := differ.Between(cp, value, read.SkipFunctions)
				t.Errorf("test Failed for '%T'!\n%s", ptr, diff)
				t.Logf("value: %s\n", sprint.Of[internal.Option](indent.Zero(), value))
				t.Logf("copy: %s\n", sprint.Of[internal.Option](indent.Zero(), cp))
			}
		})
	})
}
