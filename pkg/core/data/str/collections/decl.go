// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package collections

import (
	"github.com/andrerrcosta2/gtools/core/gtools/constraints/prim"
)

type Ordered[T prim.Ordered] interface {
	Get(i int) (T, bool)
	Exclude(i int) bool
}
