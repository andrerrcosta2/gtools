// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package progression

import (
	"github.com/andrerrcosta2/gtools/pkg/constraints"
	"github.com/andrerrcosta2/gtools/pkg/num/polyn"
	"math/rand"
)

// Sequence generates a sequence of length `n` with a given starting value.
func Sequence[T constraints.Numeric](length int) []T {
	sequence := make([]T, length)
	var zero T
	for i := range sequence {
		sequence[i] = T(i) + zero
	}
	return sequence
}

// Arithmetic generates an arithmetic progression sequence of length `n`
// with a given starting value and step size.
//
// f(n) = start + step * n
func Arithmetic[T constraints.Numeric](start T, step T, length int) []T {
	ap := make([]T, length)
	for i := 0; i < length; i++ {
		ap[i] = polyn.Linear(start, step, i)
	}
	return ap
}

// Geometric generates a geometric progression sequence of length `n`
// with a given starting value and ratio.
//
// f(n) = start * ratio^n
func Geometric[T constraints.Numeric](start T, ratio T, length int) []T {
	gp := make([]T, length)
	for i := 0; i < length; i++ {
		if i == 0 {
			gp[i] = start
		} else {
			gp[i] = T(float64(gp[i-1]) * float64(ratio))
		}
	}
	return gp
}

// Fibonacci generates a Fibonacci sequence of length `n`.
//
// f(n) = f(n-1) + f(n-2)
func Fibonacci[T constraints.Numeric](length int) []T {
	fib := make([]T, length)
	fib[0] = 0
	fib[1] = 1
	for i := 2; i < length; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	return fib
}

// Random generates a slice of random numbers of type `T` with the specified range.
func Random[T constraints.Numeric](length int, min T, max T) []T {
	randoms := make([]T, length)

	switch any(min).(type) {
	case int:
		for i := 0; i < length; i++ {
			randoms[i] = T(rand.Intn(int(max)-int(min)+1) + int(min))
		}
	case float64:
		for i := 0; i < length; i++ {
			randoms[i] = T(rand.Float64()*(float64(max)-float64(min)) + float64(min))
		}
	}

	return randoms
}

// Polynomial generates a polynomial progression sequence of length `n`
// with given coefficients.
//
// f(n) = c0 + c1*n + c2*n^2 + ... + ck*n^k
func Polynomial[T constraints.Numeric](length int, coefficients ...T) []T {
	poly := make([]T, length)
	for n := 0; n < length; n++ {
		poly[n] = polyn.Polynomial(n, coefficients...)
	}
	return poly
}

// Quadratic generates a quadratic progression sequence of length `n`
// with coefficients a, b, and c for the polynomial a + bn + cn^2.
func Quadratic[T constraints.Numeric](a, b, c T, length int) []T {
	quad := make([]T, length)
	for n := 0; n < length; n++ {
		quad[n] = polyn.Quadratic(a, b, c, n)
	}
	return quad
}

// Cubic generates a cubic progression sequence of length `n`
// with coefficients a, b, c, and d for the polynomial a + bn + cn^2 + dn^3.
func Cubic[T constraints.Numeric](a, b, c, d T, length int) []T {
	cubic := make([]T, length)
	for n := 0; n < length; n++ {
		cubic[n] = polyn.Cubic(a, b, c, d, n)
	}
	return cubic
}

// SquareRoot generates a square root progression sequence of length `n`
// with a given starting value and ratio.
//
// f(n) = start / ratio^n
func SquareRoot[T constraints.Numeric](start T, ratio T, length int) []T {
	sqrt := make([]T, length)
	for i := 0; i < length; i++ {
		if i == 0 {
			sqrt[i] = start
		} else {
			sqrt[i] = T(float64(sqrt[i-1]) / float64(ratio))
		}
	}
	return sqrt
}
