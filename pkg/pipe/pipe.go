// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

// Package pipe: this package will be refactored to use a proper recursive pipeline
package pipe

import (
	"com.github/andrerrcosta2/gtools/pkg/functions"
	"sync"
)

// Map calls f for each element in arr
func Map[T any, R any](arr []T, f functions.Function[T, R]) []R {
	result := make([]R, len(arr))
	for i, v := range arr {
		result[i] = f(v)
	}
	return result
}

// Each call f for each element in arr
func Each[T any](arr []T, f functions.Consumer[T]) {
	for _, v := range arr {
		f(v)
	}
}

// EachN calls f for each element in arr
func EachN[T any](arr []T, f functions.BiConsumer[int, T]) {
	for i, v := range arr {
		f(i, v)
	}
}

func EachMap[K comparable, V any](m map[K]V, f functions.BiConsumer[K, V]) {
	for k, v := range m {
		f(k, v)
	}
}

func GoEach[T any](arr []T, f functions.Consumer[T]) {
	var wg sync.WaitGroup
	wg.Add(len(arr))
	for _, v := range arr {
		go func(v T) {
			defer wg.Done()
			f(v)
		}(v)
	}
	wg.Wait()
}

// Filter returns the subset of arr for which f returns true
func Filter[T any](arr []T, f functions.Function[T, bool]) []T {
	result := make([]T, 0)
	for _, v := range arr {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

// Reduce calls f for each element in arr
func Reduce[T any, R any](arr []T, init R, f functions.BiFunction[R, T, R]) R {
	result := init
	for _, v := range arr {
		result = f(result, v)
	}
	return result
}

// ReduceRight calls f for each element in arr
func ReduceRight[T any, R any](arr []T, init R, f functions.BiFunction[T, R, R]) R {
	result := init
	for i := len(arr) - 1; i >= 0; i-- {
		result = f(arr[i], result)
	}
	return result
}

func FlatMap[T any, R any](arr []T, f functions.Function[T, []R]) []R {
	result := make([]R, 0)
	for _, v := range arr {
		result = append(result, f(v)...)
	}
	return result
}

func FindFirst[T any](arr []T, f functions.Function[T, bool]) (T, bool) {
	for _, v := range arr {
		if f(v) {
			return v, true
		}
	}
	return *new(T), false
}
