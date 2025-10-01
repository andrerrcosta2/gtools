// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package reflect4

import (
	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"github.com/andrerrcosta2/gtools/reflect4/internal/differ"
	"github.com/andrerrcosta2/gtools/reflect4/op"
	"reflect"
)

// Diff returns the difference between two values.
func Diff(a, b any, o ...op.Option) (string, bool, error) {
	return differ.Between(indent.Zero(), reflect.ValueOf(a), reflect.ValueOf(b), differ.NewStrategy(o...))
}
