// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package collections

import (
	"github.com/andrerrcosta2/gtools/pkg/constraints"
)

type Ordered[T constraints.Ordered] interface {
	Get(i int) (T, bool)
	Exclude(i int) bool
}
