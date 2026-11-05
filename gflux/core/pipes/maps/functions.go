// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package maps

import (
	"sort"

	"github.com/andrerrcosta2/gtools/core/data/str/dict"
	"github.com/andrerrcosta2/gtools/core/domain/function"
	"github.com/andrerrcosta2/gtools/core/util/typeutil/generics"
)

// Any returns the first key-value pair in a map.
//
// Parameters:
// - m: The input map.
//
// Returns:
// - The first key and value in the map, along with a boolean indicating whether the map is empty.
func Any[M ~map[K]V, K comparable, V any](m M) (key K, value V, ok bool) {
	for k, v := range m {
		return k, v, true
	}
	return
}

// AreSameKeys returns true if the map contains the same keys as the given slice.
//
// Parameters:
// - m: The input map.
// - keys: The slice of keys to check for.
//
// Returns:
// - True if the map contains the same keys as the slice, false otherwise.
func AreSameKeys[M ~map[K]V, K comparable, V any](m M, keys []K) bool {
	if len(m) != len(keys) {
		return false
	}
	return ContainsAllKeys(m, keys)
}

// AreSameValues returns true if the map contains the same values as the given slice.
//
// Parameters:
// - m: The input map.
// - values: The slice of values to check for.
//
// Returns:
// - True if the map contains the same values as the slice, false otherwise.
func AreSameValues[M ~map[K]V, K comparable, V comparable](m M, values []V) bool {
	if len(m) != len(values) {
		return false
	}
	return ContainsAllValues(m, values)
}

// Cast applies a given function to each key-value pair in a map and returns a new slice of the results.
//
// Parameters:
// - m: The input map.
// - f: The function to apply to each key-value pair.
//
// Returns:
// - A new slice with the same length as the input map, but with values obtained by applying the function to each key-value pair.
func Cast[M ~map[K]V, L comparable, X any, K comparable, V any, Z generics.BiTyped[L, X]](m M, zip function.BiFunction[K, V, Z]) []Z {
	cast := make([]Z, 0, len(m))

	for k, v := range m {
		cast = append(cast, zip(k, v))
	}

	return cast
}

// Clone creates a map with the same keys and values
func Clone[M ~map[K]V, K comparable, V any](m M) M {
	out := make(map[K]V, len(m))
	for k, v := range m {
		out[k] = v
	}
	return out

}

// Collect returns a map with keys and values obtained by applying the function f to each entry in the input slice.
// The function f takes an entry of type generics.BiTyped[K, V] and returns a key of type K and a value of type V.
// The input slice entries must have unique keys.
// If there are duplicate keys in the input slice, the last value for each key will be kept in the resulting map.
//
// Parameters:
// - entries: The input slice of entries.
// - f: The function to apply to each entry. It should return a key-value pair.
//
// Returns:
// - A map with keys and values obtained by applying the function to each entry.
func Collect[M ~map[K]V, K comparable, E generics.BiTyped[K, V], V any](entries []E, f function.Function2[E, K, V]) M {
	result := make(M, len(entries))

	for _, entry := range entries {
		k, v := f(entry)

		result[k] = v
	}

	return result
}

// ContainsAllKeys returns true if the map contains all the given keys.
//
// Parameters:
// - m: The input map.
// - keys: The keys to check for.
//
// Returns:
// - True if the map contains all the keys, false otherwise.
func ContainsAllKeys[M ~map[K]V, K comparable, V any](m M, keys []K) bool {
	for _, k := range keys {
		if _, ok := m[k]; !ok {
			return false
		}
	}
	return true
}

// ContainsAllValues returns true if the map contains all the given values.
//
// Parameters:
// - m: The input map.
// - values: The values to check for.
//
// Returns:
// - True if the map contains all the values, false otherwise.
func ContainsAllValues[M ~map[K]V, K comparable, V comparable](m M, values []V) bool {
	for _, v := range values {
		if !ContainsValue(m, v) {
			return false
		}
	}
	return true
}

// ContainsValue returns true if the map contains the given value.
//
// Parameters:
// - m: The input map.
// - v: The value to check for.
//
// Returns:
// - True if the map contains the value, false otherwise.
func ContainsValue[M ~map[K]V, K comparable, V comparable](m M, v V) bool {
	for _, value := range m {
		if value == v {
			return true
		}
	}
	return false
}

func Difference[M ~map[K]V, K comparable, V any](a, b M) M {
	res := make(M)

	// Se a for nil, a diferença é um mapa vazio
	if a == nil {
		return res
	}

	// Se b for nil, a diferença é o próprio a
	if b == nil {
		for k, v := range a {
			res[k] = v
		}
		return res
	}

	// Copia apenas os elementos de a que não estão em b
	for k, v := range a {
		if _, exists := b[k]; !exists {
			res[k] = v
		}
	}

	return res
}

// Each applies the given BiConsumer function to each key-value pair in the map.
//
// Parameters:
// - m: The map to iterate over.
// - f: The BiConsumer function to apply to each key-value pair.
func Each[M ~map[K]V, K comparable, V any](m M, f function.BiConsumer[K, V]) {
	for k, v := range m {
		f(k, v)
	}
}

// ExceptKeys creates a new map by excluding the specified keys from the input map.
//
// Parameters:
// - m: The input map.
// - keys: The keys to exclude.
//
// Returns:
// - A new map with the specified keys excluded.
func ExceptKeys[M ~map[K]V, K comparable, V any](m M, keys ...K) M {
	// Create a set of keys to exclude for O(1) membership checks.
	excludeSet := make(map[K]struct{})
	for _, key := range keys {
		excludeSet[key] = struct{}{}
	}
	filtered := make(map[K]V)
	for k, v := range m {
		if _, exists := excludeSet[k]; !exists {
			filtered[k] = v
		}
	}
	return filtered
}

// Filter applies a predicate to each value in a map and returns a new map with the entries that pass the predicate.
//
// Parameters:
// - m: The input map.
// - f: The predicate function to apply to each value.
//
// Returns:
// - A new map with the entries that pass the predicate.
func Filter[M ~map[K]V, K comparable, V any](m M, f function.BiPredicate[K, V]) M {
	filtered := make(map[K]V)
	for k, v := range m {
		if f(k, v) {
			filtered[k] = v
		}
	}
	return filtered
}

// FilterIn applies a predicate to each key-value pair in a map and removes the entries that do not pass the predicate.
// This pipe modifies the input map and does not return anything.
//
// Parameters:
// - m: The input map.
// - f: The predicate function to apply to each key-value pair.
func FilterIn[M ~map[K]V, K comparable, V any](m *map[K]V, f function.BiPredicate[K, V]) {
	for k, v := range *m {
		if !f(k, v) {
			delete(*m, k)
		}
	}
}

// FilterKeys applies a predicate to each key in a map and returns a new slice with the keys that pass the predicate.
//
// Parameters:
// - m: The input map.
// - f: The predicate function to apply to each key.
//
// Returns:
// - A new slice with the keys that pass the predicate.
func FilterKeys[M ~map[K]V, K comparable, V any](m M, f function.Predicate[K]) []K {
	var result []K
	for k := range m {
		if f(k) {
			result = append(result, k)
		}
	}
	return result
}

// FilterValues applies a predicate to each value in a map and returns a new slice with the values that pass the predicate.
//
// Parameters:
// - m: The input map.
// - f: The predicate function to apply to each value.
//
// Returns:
// - A new slice with the values that pass the predicate.
func FilterValues[M ~map[K]V, K comparable, V any](m M, f function.Predicate[V]) []V {
	var result []V
	for _, v := range m {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

// FlatValues applies a function to each element in a map of slices and returns a new slice with the transformed values.
//
// Parameters:
// - m: The input map with slices of values.
// - f: The function to apply to each value.
//
// Returns:
// - A new slice with the transformed values obtained by applying the function to each element in the map.

func FlatValues[M ~map[K][]V, K comparable, V any, X any](m M, f function.Function[V, X]) []X {
	var result []X
	for _, slice := range m {
		for _, v := range slice {
			result = append(result, f(v))
		}
	}
	return result
}

// FlatValuesSorted applies a function to each element in a map of slices and returns a new sorted slice with the transformed values.
//
// Parameters:
// - m: The input map with slices of values.
// - f: The function to apply to each value.
// - less: The comparison function to use for sorting the result.
//
// Returns:
// - A new sorted slice with the transformed values obtained by applying the function to each element in the map.
func FlatValuesSorted[M ~map[K][]V, K comparable, V any, X any](m M, f function.Function[V, X], less function.Less[X]) []X {
	var result = make([]X, 0, len(m))
	for _, slice := range m {
		for _, v := range slice {
			result = append(result, f(v))
		}
	}
	sort.Slice(result, func(i, j int) bool {
		return less(result[i], result[j])
	})
	return result
}

// GroupBy groups elements of a slice into a map based on a key function.
func GroupBy[M ~map[K]O, O ~[]V, K comparable, V any](values O, fn function.Function[V, K]) M {
	return MakeWithAcc[M](values, func(v V) (K, V) {
		return fn(v), v
	})
}

func Intersect[M ~map[K]V, K comparable, V any](a, b M) M {
	if a == nil || b == nil {
		return make(M)
	}
	if len(a) > len(b) {
		a, b = b, a
	}
	res := make(M, len(a))
	for k, v := range a {
		if _, ok := b[k]; ok {
			res[k] = v
		}
	}
	return res
}

// Keys returns all keys from a map
func Keys[M ~map[K]V, K comparable, V any](m M) []K {
	out := make([]K, len(m), len(m))
	for k := range m {
		out = append(out, k)
	}
	return out
}

// Map applies a given BiFunction to each key-value pair in a map and returns a new map.
//
// Parameters:
// - m: The input map.
// - f: The BiFunction to apply to each key-value pair. It should return a pointer to an Entry struct.
//
// Returns:
// - A new map with the same keys as the input map, but with values obtained by applying the BiFunction to each key-value pair.
func Map[MO ~map[L]X, MI ~map[K]V, K comparable, V any, L comparable, X any](m MI, f function.BiFunction[K, V, dict.Entry[L, X]]) MO {
	result := make(MO, len(m))
	for k, v := range m {
		entry := f(k, v)
		result[entry.Key()] = entry.Value()
	}
	return result
}

// MapEntries applies a given BiFunction to each key-value pair in a map and returns a new EntrySet.
//
// Parameters:
// - m: The input map.
// - f: The BiFunction to apply to each key-value pair. It should return a pointer to an Entry struct.
//
// Returns:
// - A new EntrySet with the same keys as the input map, but with values obtained by applying the BiFunction to each key-value pair.
func MapEntries[M ~map[K]V, K comparable, V any, L comparable, X any](m M,
	f function.BiFunction[K, V, dict.Entry[L, X]]) []dict.Entry[L, X] {
	entries := make([]dict.Entry[L, X], 0, len(m))
	for k, v := range m {
		entries = append(entries, f(k, v))
	}
	return entries
}

// MakeWith creates a new map with the specified key type and value type, and populates it with the values
// returned by the provided function.
// If multiple elements produce the same key, the last value produced by fn replaces the previous one.
func MakeWith[M ~map[K]V, S ~[]I, I any, K comparable, V any](s S, fn function.Function2[I, K, V]) M {
	m := make(M, len(s))
	for _, i := range s {
		k, v := fn(i)
		m[k] = v
	}
	return m
}

// MakeWithAcc creates a new map with the specified key type and value type, and populates it with the values
// returned by the provided function.
func MakeWithAcc[M ~map[K]O, O ~[]V, S ~[]I, I any, K comparable, V any](s S, fn function.Function2[I, K, V]) M {
	m := make(M, len(s))

	for _, i := range s {
		k, v := fn(i)
		m[k] = append(m[k], v)
	}

	return m
}

// MapKeys applies a given function to each key in a map and returns a new slice of the results.
//
// Parameters:
// - m: The input map.
// - f: The function to apply to each key.
//
// Returns:
// - A new slice with the same length as the input map, but with values obtained by applying the function to each key.
func MapKeys[M ~map[K]V, K comparable, V any, X any](m M, f function.Function[K, X]) []X {
	result := make([]X, 0, len(m))
	for k := range m {
		result = append(result, f(k))
	}
	return result
}

// MapValues applies a given function to each value in a map and returns a new slice of the results.
//
// Parameters:
// - m: The input map.
// - f: The function to apply to each value.
//
// Returns:
// - A new slice with the same length as the input map, but with values obtained by applying the function to each value.
func MapValues[M ~map[K]V, K comparable, V any, X any](m M, f function.Function[V, X]) []X {
	result := make([]X, 0, len(m))
	for _, v := range m {
		result = append(result, f(v))
	}
	return result
}

func MapValuesThenSort[M ~map[K]V, K comparable, V any, X any](m M, f function.Function[V, X], less function.Less[X]) []X {
	result := make([]X, 0, len(m))
	for _, v := range m {
		result = append(result, f(v))
	}
	sort.Slice(result, func(i, j int) bool {
		return less(result[i], result[j])
	})
	return result
}

// MapWithKeys creates a new map with keys generated by applying the function f to each key in the input keys slice.
// The function f takes a key of type K and returns a key of type L and a value of type V.
// The resulting map has keys of type L and values of type V.
// The function assumes that the keys in the input keys slice are unique.
// If the function f generates duplicate keys, the resulting map will only contain the last value for each key.
func MapWithKeys[M ~map[L]V, K comparable, L comparable, V any](keys []K, f function.Function2[K, L, V]) M {
	result := make(map[L]V, len(keys))
	for _, key := range keys {
		newKey, value := f(key)
		result[newKey] = value
	}
	return result
}

// Merge faz a junção de elementos entre dois mapas
func Merge[M ~map[K]V, K comparable, V any](a, b M) M {
	res := make(M, len(a)+len(b))

	if a != nil {
		for k, v := range a {
			res[k] = v
		}
	}

	if b != nil {
		for k, v := range b {
			res[k] = v
		}
	}

	return res
}

// Values returns all values from a map
func Values[M ~map[K]V, K comparable, V any](m M) []V {
	out := make([]V, len(m), len(m))
	for _, v := range m {
		out = append(out, v)
	}
	return out
}
