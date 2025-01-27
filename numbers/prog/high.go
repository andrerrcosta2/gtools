// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package prog

import (
	"github.com/andrerrcosta2/gtools/numbers"
	"github.com/andrerrcosta2/gtools/numbers/polyn"
)

// High generates a progression of length `n`
// with polynomial coefficients for an m-th order polynomial.
//
// f(n) = c0 + c1*n + c2*n^2 + ... + ck*n^k
//
// where:
// - ck are the coefficients of the polynomial
// - n is the input value
//
// https://en.wikipedia.org/wiki/Linear_recurrence_relation
func High[T numbers.Real](coefficients []T, length int) []T {
	progression := make([]T, length)

	for n := 0; n < length; n++ {
		progression[n] = polyn.N(n, coefficients...)
	}

	return progression
}
