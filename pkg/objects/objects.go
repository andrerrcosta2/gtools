// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package objects

import "github.com/andrerrcosta2/gtools/pkg/functions"

func Equals[T comparable](values ...T) bool {
	return EqualsBy[T](func(a, b T) bool { return a == b }, values...)
}

func EqualsBy[G any](fn functions.BiFunction[G, G, bool], values ...G) bool {
	for i := 0; i < len(values)-1; i++ {
		if !fn(values[i], values[i+1]) {
			return false
		}
	}
	return true
}

func Uniques[T comparable](values ...T) bool {
	return UniquesBy[T](func(a, b T) bool { return a == b }, values...)
}

func UniquesBy[G any](fn functions.BiFunction[G, G, bool], values ...G) bool {
	for i := 0; i < len(values)-1; i++ {
		for j := i + 1; j < len(values); j++ {
			if fn(values[i], values[j]) {
				return false
			}
		}
	}
	return true
}
