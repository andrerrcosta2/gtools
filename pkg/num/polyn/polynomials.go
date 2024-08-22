// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package polyn

import "github.com/andrerrcosta2/gtools/pkg/constraints"

// Polynomial evaluates a polynomial of degree `d` with given coefficients
// for a given value of `n`. For example, if coefficients are [a, b, c, d] for a cubic
// polynomial, it evaluates a + bn + cn^2 + dn^3.
func Polynomial[T constraints.Numeric](n int, coefficients ...T) T {
	var result T
	for i, coeff := range coefficients {
		result += T(float64(coeff) * float64(Pow(n, i)))
	}
	return result
}

// Linear evaluates a linear polynomial
// for a given value of `n`.
//
// f(n) = a + b*n
func Linear[T constraints.Numeric](a, b T, n int) T {
	return a + T(float64(b)*float64(n))
}

// Quadratic evaluates a quadratic polynomial with given coefficients
// for a given value of `n`.
//
// f(n) = a + b*n + c*n^2
func Quadratic[T constraints.Numeric](a, b, c T, n int) T {
	return a + T(float64(b)*float64(n)) + T(float64(c)*float64(n*n))
}

// Cubic evaluates a cubic polynomial with given coefficients
// for a given value of `n`.
//
// f(n) = a + b*n + c*n^2 + d*n^3
func Cubic[T constraints.Numeric](a, b, c, d T, n int) T {
	return a + T(float64(b)*float64(n)) + T(float64(c)*float64(n*n)) + T(float64(d)*float64(n*n*n))
}

// Pow is a helper function to compute power for the polynomial evaluation.
func Pow(base, exp int) int {
	result := 1
	for i := 0; i < exp; i++ {
		result *= base
	}
	return result
}
