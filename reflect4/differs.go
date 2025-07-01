// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package reflect4

import (
	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"github.com/andrerrcosta2/gtools/reflect4/internal/differ"
	"reflect"
)

// Diff returns the difference between two values.
func Diff(tab indent.Tab, a, b interface{}) (string, bool, error) {
	return differ.Between(tab, reflect.ValueOf(a), reflect.ValueOf(b))
}
