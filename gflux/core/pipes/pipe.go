// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

// Package pipes Package pipe: this package will be refactored to use a proper recursive pipeline
package pipes

import (
	"github.com/andrerrcosta2/gtools/core/domain/functions"
	"github.com/andrerrcosta2/gtools/core/domain/functions/closers"
	"strings"
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

// EachAsync call f for each element in arr asynchronously
// It means there is no waiting group to block the caller thread
func EachAsync[T any](arr []T, f functions.Consumer[T]) {
	for _, v := range arr {
		go func(v T) {
			f(v)
		}(v)
	}
}

// EachN calls f for each element in arr
func EachN[T any](arr []T, f functions.BiConsumer[int, T]) {
	for i, v := range arr {
		f(i, v)
	}
}

// EachMap calls a BiConsumer for each key-value pair in a map
func EachMap[K comparable, V any](m map[K]V, f functions.BiConsumer[K, V]) {
	for k, v := range m {
		f(k, v)
	}
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

// FilterMap returns the subset of a mapped arr for each f that returns true
func FilterMap[T any, R any](arr []T, f functions.Function2[T, R, bool]) []R {
	result := make([]R, 0)
	for _, v := range arr {
		r, ok := f(v)
		if ok {
			result = append(result, r)
		}
	}
	return result
}

// FindFirst returns the first element in arr for which f returns true
func FindFirst[T any](arr []T, f functions.Function[T, bool]) (T, bool) {
	for _, v := range arr {
		if f(v) {
			return v, true
		}
	}
	return *new(T), false
}

// FlatMap calls a mapper function for each element in arr
func FlatMap[T any, R any](arr []T, f functions.Function[T, []R]) []R {
	result := make([]R, 0)
	for _, v := range arr {
		result = append(result, f(v)...)
	}
	return result
}

// GoEach calls a consumer for each element in arr using goroutines.
// It uses a wait group to wait for all goroutines to finish before release
// the main thread.
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

// Reduce calls a BiFunction to transform each element in arr, receiving as parameter
// the previous result starting from an optional initial value. If no value is passed
// it uses its zero-value.
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

// SemEach calls a consumer for each element in arr using goroutines.
// It uses a semaphore to limit the number of concurrent operations
func SemEach[T any](arr []T, f functions.Consumer[T], maxConcurrency int) {
	var wg sync.WaitGroup
	var sem = semaph(maxConcurrency)
	wg.Add(len(arr))
	for _, v := range arr {
		go func(v T) {
			sem.Acq()
			defer closers.SyncThread(&wg, sem)
			f(v)
		}(v)
	}
	wg.Wait()
}

// StringsMap calls a mapper function for each element in arr, joining the results
// with a separator. Then it returns a single string with all the results.
//
// Example:
//
//	arr := []int{1, 2, 3}
//	f := func(i int) string { return strconv.Itoa(i) }
//	StringsMap(arr, ", ", f) // "1, 2, 3"
func StringsMap[I any](arr []I, join string, f functions.Function[I, string]) string {
	sb := strings.Builder{}

	for i, v := range arr {
		if i < len(arr)-1 {
			sb.WriteString(f(v) + join)
		}
		sb.WriteString(f(v))
	}

	return sb.String()
}
