// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package comp

import "github.com/andrerrcosta2/gtools/core/domain/constraints/prim/nums"

// Min returns the minimum of two values of type T and whether they are compare.
func Min[T nums.Ordered](a, b T) (min T, eq bool) {
	if a < b {
		return a, false
	} else if a > b {
		return b, false
	} else {
		return a, true
	}
}
