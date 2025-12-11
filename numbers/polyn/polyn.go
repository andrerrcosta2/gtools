// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package polyn

import (
	"github.com/andrerrcosta2/gtools/numbers"
)

// N evaluates a polynomial of degree `d` with given coefficients
// for a given value of `n`. For example, if coefficients are [a, b, c, d] for a cubic
// polynomial, it evaluates a + bn + cn^2 + dn^3.
func N[T numbers.Real](n int, coefficients ...T) T {
	var result T
	for i, cfc := range coefficients {
		result += cfc * T(Pow(n, i))
	}
	return result
}

// Lin evaluates a linear polynomial for a given value of `n`.
// f(n) = a + b*n
func Lin[T numbers.Real](a, b T, n int) T {
	return a + b*T(n)
}

// Quad evaluates a quadratic polynomial with given coefficients
// for a given value of `n`.
// f(n) = a + b*n + c*n^2
func Quad[T numbers.Real](a, b, c T, n int) T {
	return a + b*T(n) + c*T(n*n)
}

// Cub evaluates a cubic polynomial with given coefficients
// for a given value of `n`.
// f(n) = a + b*n + c*n^2 + d*n^3
func Cub[T numbers.Real](a, b, c, d T, n int) T {
	return a + b*T(n) + c*T(n*n) + d*T(n*n*n)
}

// Pow is a std function to compute power for the polynomial evaluation.
func Pow(base, exp int) int {
	result := 1
	for i := 0; i < exp; i++ {
		result *= base
	}
	return result
}
