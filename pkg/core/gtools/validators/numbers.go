// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package validators

import "github.com/andrerrcosta2/gtools/core/gtools/constraints/prim/nums"

func Range[T nums.Ordered](min, max T) bool {
	return min < max
}
