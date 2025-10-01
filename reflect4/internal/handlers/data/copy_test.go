// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package data

import (
	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"github.com/andrerrcosta2/gtools/core/testlite/assertlite"
	"github.com/andrerrcosta2/gtools/gtests"
	"github.com/andrerrcosta2/gtools/reflect4/internal"
	"github.com/andrerrcosta2/gtools/reflect4/internal/differ"
	"github.com/andrerrcosta2/gtools/reflect4/internal/equals"
	"reflect"
	"testing"
)

func TestDeepCopy(t *testing.T) {

}

func TestDeepCopyStruct(t *testing.T) {
	pointers := gtests.Structs.Fuzz().Categories().Refs().All()
	// Test 2 - Deep clone default write
	t.Run("default write", func(t *testing.T) {
		pointers.Each(func(pointer any) {
			if pointer == nil {
				panic("nil pointer")
			}
			cp, err := DeepCopy[internal.Option](reflect.ValueOf(pointer))
			assertlite.NoError(t, err, "error while deep copying:\n%v\n", err)

			diff, eq, err := differ.Between[internal.Option](indent.Zero(), cp, reflect.ValueOf(pointer))
			assertlite.NoError(t, err, "error on differ between: %v\n", err)
			assertlite.True(t, eq, "deep clone failed\n%s", diff)
		})
	})
}

func TestPlayground(t *testing.T) {
	t.Run("account", func(t *testing.T) {
		account := gtests.Structs.Fuzz().Values().Account()
		target := reflect.ValueOf(&account)
		cp, err := DeepCopy[internal.Option](target)
		assertlite.NoError(t, err, "error while deep copying:\n%v\n", err)
		assertlite.True(t, equals.Deep[internal.Option](cp, target), "Expected equals")

		//diff, eq, err := differ.Between[internal.Option](indent.Zero(), cp, target)
		//assertlite.NoError(t, err, "error on differ between: %v\n", err)
		//assertlite.True(t, eq, "deep clone failed\n%s", diff)
		//assertlite.Equals(t, cp.Interface(), account, "Not Equals: %s", diff)
	})
}
