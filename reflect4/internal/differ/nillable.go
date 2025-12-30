// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package differ

import (
	"errors"
	"reflect"

	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"github.com/andrerrcosta2/gtools/core/format/differs"
	"github.com/andrerrcosta2/gtools/reflect4/internal/types"
)

var (
	ErrInvalidVTypes   = errors.New("can't diff between invalid type4")
	ErrInvalidReceived = errors.New("couldn't diff because the received type is invalid")
	ErrInvalidExpected = errors.New("couldn't diff because the expected type is invalid")
)

func handleInvalidValues(tab indent.Indentor, a, b reflect.Value) (diff differs.Difference, invalids bool) {
	if !a.IsValid() || !b.IsValid() {
		if !a.IsValid() && !b.IsValid() {
			return differs.DiffError(differs.BothInvalid(tab), ErrInvalidVTypes),
				true
		}
		if a.IsValid() {
			aa := types.ValidValueName(a.Type())
			return differs.DiffError(differs.InvalidExpected(tab, aa),
				ErrInvalidReceived), true
		}
		bb := types.ValidValueName(b.Type())
		return differs.DiffError(differs.InvalidReceived(tab, bb),
			ErrInvalidExpected), true
	}
	return
}
