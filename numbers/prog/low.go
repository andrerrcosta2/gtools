// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package prog

import (
	"fmt"
	"github.com/andrerrcosta2/gtools/numbers"
	"github.com/andrerrcosta2/gtools/numbers/polyn"
	"math/rand"
)

// Seq generates a sequence of length `n` with a given starting value.
func Seq[T numbers.Real](length int) []T {
	sequence := make([]T, length)
	var zero T
	for i := range sequence {
		sequence[i] = T(i) + zero
	}
	return sequence
}

// Arit generates an arithmetic progression sequence of length `n`
// with a given starting value and step size.
//
// f(n) = start + step * n
func Arit[T numbers.Real](start T, step T, length int) ([]T, error) {
	ap := make([]T, length)
	for i := 0; i < length; i++ {
		ap[i] = polyn.Lin(start, step, i)
	}
	return ap, nil
}

// Geom generates a geometric progression sequence of length `n`
// with a given starting value and ratio.
//
// f(n) = start * ratio^n
func Geom[T numbers.Real](start T, ratio T, length int) ([]T, error) {
	gp := make([]T, length)
	for i := 0; i < length; i++ {
		if i == 0 {
			gp[i] = start
		} else {
			gp[i] = T(float64(gp[i-1]) * float64(ratio))
		}
	}
	return gp, nil
}

// Fib generates a Fib sequence of length `n`.
//
// f(n) = f(n-1) + f(n-2)
func Fib[T numbers.Real](length int) (fib []T, err error) {
	if length == 0 {
		return
	} else if length == 1 {
		return []T{0}, nil
	} else {
		fib = make([]T, length)
		fib[0] = 0
		fib[1] = 1
		for i := 2; i < length; i++ {
			fib[i] = fib[i-1] + fib[i-2]
		}
		return
	}
}

// Rand generates a slice of random numbers of type `T` with the specified range.
func Rand[T numbers.Real](length int, min T, max T) ([]T, error) {
	randoms := make([]T, length)

	switch any(min).(type) {
	case int, int8, int16, int32, int64:
		for i := 0; i < length; i++ {
			randoms[i] = T(rand.Intn(int(max)-int(min)+1) + int(min))
		}
	case float64, float32:
		for i := 0; i < length; i++ {
			randoms[i] = T(rand.Float64()*(float64(max)-float64(min)) + float64(min))
		}
	}

	return randoms, nil
}

// Quad generates a quadratic progression sequence of length `n`
// with coefficients a, b, and c for the polynomial a + bn + cn^2.
func Quad[T numbers.Real](a, b, c T, length int) ([]T, error) {
	quad := make([]T, length)
	for n := 0; n < length; n++ {
		quad[n] = polyn.Quad(a, b, c, n)
	}
	return quad, nil
}

// Cub generates a cubic progression sequence of length `n`
// with coefficients a, b, c, and d for the polynomial a + bn + cn^2 + dn^3.
func Cub[T numbers.Real](a, b, c, d T, length int) ([]T, error) {
	cubic := make([]T, length)
	for n := 0; n < length; n++ {
		cubic[n] = polyn.Cub(a, b, c, d, n)
	}
	return cubic, nil
}

// Sqrt generates a square root progression sequence of length `n`
// with a given starting value and ratio.
//
// f(n) = start / ratio^n
func Sqrt[T numbers.Real](start T, ratio T, length int) ([]T, error) {
	if ratio == 0 {
		return nil, fmt.Errorf("ratio cannot be zero")
	}

	sqrt := make([]T, length)
	for i := 0; i < length; i++ {
		if i == 0 {
			sqrt[i] = start
		} else {
			sqrt[i] = T(float64(sqrt[i-1]) / float64(ratio))
		}
	}
	return sqrt, nil
}
