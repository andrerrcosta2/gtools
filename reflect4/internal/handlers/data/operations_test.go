// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package data

import (
	"github.com/andrerrcosta2/gtools/core/format/fmx"
	"github.com/andrerrcosta2/gtools/core/testlite/assertlite"
	"github.com/andrerrcosta2/gtools/gtests"
	"github.com/andrerrcosta2/gtools/reflect4/internal"
	"github.com/andrerrcosta2/gtools/reflect4/internal/equals"
	"reflect"
	"testing"
)

func TestDeepCopyArray(t *testing.T) {
	s := defCopyStrat()
	gtests.Arrays.Fuzz().Values().All().Each(func(arr any) {
		a := reflect.ValueOf(arr)
		b, err := deepCopyArray(a, s)
		assertlite.NoError(t, err, "expected no error, got %v", err)
		assertlite.True(t, equals.Deep[internal.Option](a, b),
			gtests.ErrorDiff("expected data to be deep equal copy",
				fmx.Sobj(a), fmx.Sobj(b)))
	})
}
