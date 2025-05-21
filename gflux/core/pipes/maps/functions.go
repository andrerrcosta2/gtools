// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package maps

import (
	"github.com/andrerrcosta2/gtools/core/domain/functions"
	"github.com/andrerrcosta2/gtools/core/generics"
	"sort"
)

// Any returns the first key-value pair in a map.
//
// Parameters:
// - m: The input map.
//
// Returns:
// - The first key and value in the map, along with a boolean indicating whether the map is empty.
func Any[K comparable, V any](m *map[K]V) (key K, value V, ok bool) {
	for k, v := range *m {
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
func AreSameKeys[K comparable, V any](m *map[K]V, keys []K) bool {
	if len(*m) != len(keys) {
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
func AreSameValues[K comparable, V comparable](m *map[K]V, values []V) bool {
	if len(*m) != len(values) {
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
func Cast[K comparable, V any, L comparable, X any](m *map[K]V, f functions.BiFunction[K, V, generics.BiTyped[L, X]]) []generics.BiTyped[L, X] {
	// Create a new slice with initial capacity equal to the number of entries in the input map.
	cast := make([]generics.BiTyped[L, X], 0, len(*m))

	// Iterate over each key-value pair in the input map.
	for k, v := range *m {
		// Apply the function to the current key-value pair and append the result to the result slice.
		cast = append(cast, f(k, v))
	}

	// Return the resulting slice.
	return cast
}

// ContainsAllKeys returns true if the map contains all the given keys.
//
// Parameters:
// - m: The input map.
// - keys: The keys to check for.
//
// Returns:
// - True if the map contains all the keys, false otherwise.
func ContainsAllKeys[K comparable, V any](m *map[K]V, keys []K) bool {
	for _, k := range keys {
		if !ContainsKey(m, k) {
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
func ContainsAllValues[K comparable, V comparable](m *map[K]V, values []V) bool {
	for _, v := range values {
		if !ContainsValue(m, v) {
			return false
		}
	}
	return true
}

// ContainsKey returns true if the map contains the given key.
//
// Parameters:
// - m: The input map.
// - k: The key to check for.
//
// Returns:
// - True if the map contains the key, false otherwise.
func ContainsKey[K comparable, V any](m *map[K]V, k K) bool {
	_, ok := (*m)[k]
	return ok
}

// ContainsValue returns true if the map contains the given value.
//
// Parameters:
// - m: The input map.
// - v: The value to check for.
//
// Returns:
// - True if the map contains the value, false otherwise.
func ContainsValue[K comparable, V comparable](m *map[K]V, v V) bool {
	for _, value := range *m {
		if value == v {
			return true
		}
	}
	return false
}

// Each applies the given BiConsumer function to each key-value pair in the map.
//
// Parameters:
// - m: The map to iterate over.
// - f: The BiConsumer function to apply to each key-value pair.
func Each[K comparable, V any](m map[K]V, f functions.BiConsumer[K, V]) {
	// Iterate over each key-value pair in the map.
	for k, v := range m {
		// Apply the BiConsumer function to the current key-value pair.
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
func ExceptKeys[K comparable, V any](m *map[K]V, keys ...K) *map[K]V {
	// Create a set of keys to exclude for O(1) membership checks.
	excludeSet := make(map[K]struct{})
	for _, key := range keys {
		excludeSet[key] = struct{}{}
	}

	// Create a new map to store the filtered entries.
	filtered := make(map[K]V)

	// Iterate over the input map and add only keys not in the exclude set.
	for k, v := range *m {
		if _, exists := excludeSet[k]; !exists {
			filtered[k] = v
		}
	}

	return &filtered
}

// Fetch returns a map with keys and values obtained by applying the function f to each entry in the input slice.
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
func Fetch[E generics.BiTyped[K, V], K comparable, V any](entries []E, f functions.Function2[E, K, V]) *map[K]V {
	// Create a new map with initial capacity equal to the number of entries in the input slice.
	result := make(map[K]V, len(entries))

	// Iterate over each entry in the input slice.
	for _, entry := range entries {
		// Apply the function f to the current entry and obtain a key of type K and a value of type V.
		k, v := f(entry)

		// Addf the key-value pair to the result map.
		result[k] = v
	}

	// Return the resulting map.
	return &result
}

// Filter applies a predicate to each value in a map and returns a new map with the entries that pass the predicate.
//
// Parameters:
// - m: The input map.
// - f: The predicate function to apply to each value.
//
// Returns:
// - A new map with the entries that pass the predicate.
func Filter[K comparable, V any](m *map[K]V, f functions.BiPredicate[K, V]) *map[K]V {
	// Create a new map to store the filtered entries.
	filtered := make(map[K]V)

	// Iterate over each key-value pair in the input map.
	for k, v := range *m {
		// Check if the value passes the predicate.
		if f(k, v) {
			// If it passes, add the key-value pair to the filtered map.
			filtered[k] = v
		}
	}

	// Return the filtered map.
	return &filtered
}

// FilterIn applies a predicate to each key-value pair in a map and removes the entries that do not pass the predicate.
// This pipe modifies the input map and does not return anything.
//
// Parameters:
// - m: The input map.
// - f: The predicate function to apply to each key-value pair.
func FilterIn[K comparable, V any](m *map[K]V, f functions.BiPredicate[K, V]) {
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
func FilterKeys[K comparable, V any](m *map[K]V, f functions.Predicate[K]) []K {
	// Create a slice to store the filtered keys.
	var result []K

	// Iterate over each key in the map.
	for k := range *m {
		// Check if the key passes the predicate.
		if f(k) {
			// If it passes, append it to the result slice.
			result = append(result, k)
		}
	}

	// Return the filtered slice.
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
func FilterValues[K comparable, V any](m *map[K]V, f functions.Predicate[V]) []V {
	// Create a slice to store the filtered values.
	var result []V

	// Iterate over each value in the map.
	for _, v := range *m {
		// Check if the value passes the predicate.
		if f(v) {
			// If it passes, append it to the result slice.
			result = append(result, v)
		}
	}

	// Return the filtered slice.
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

func FlatValues[K comparable, V any, X any](m *map[K][]V, f functions.Function[V, X]) []X {
	// Create a slice to store the flattened and transformed values.
	var result []X

	// Iterate over each slice of values in the map.
	for _, slice := range *m {
		// Apply the function to each element in the slice and append the result to the final slice.
		for _, v := range slice {
			result = append(result, f(v))
		}
	}

	// Return the flattened and transformed slice.
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
func FlatValuesSorted[K comparable, V any, X any](m *map[K][]V, f functions.Function[V, X], less functions.BiFunction[X, X, bool]) []X {
	var result []X
	for _, slice := range *m {
		for _, v := range slice {
			result = append(result, f(v))
		}
	}

	// Sort the result using the provided comparison function
	sort.Slice(result, func(i, j int) bool {
		return less(result[i], result[j])
	})

	return result
}

// GroupBy groups elements of a slice into a map based on a key function.
func GroupBy[K comparable, V any, O ~[]V](values []V, fn functions.Function[V, K]) map[K]O {
	// Initialize the result map.
	result := make(map[K]O)

	// Iterate over each value in the input slice.
	for _, value := range values {
		// Derive the key for the current value.
		key := fn(value)

		// Append the value to the slice corresponding to the key.
		result[key] = append(result[key], value)
	}

	// Return the resulting map.
	return result
}

// MakeWith creates a new map with the specified key type and value type, and populates it with the values
// returned by the provided function.
//
// Parameters:
// - s: The input slice.
// - fn: The function to apply to each element in the slice.
//
// Returns:
// - A new map with the specified key type and value type, populated with the values returned by the
// provided function.
func MakeWith[I any, K comparable, V any](s []I, fn functions.Function2[I, K, V]) map[K]V {
	m := make(map[K]V)
	for _, i := range s {
		k, v := fn(i)
		m[k] = v
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
func MapKeys[K comparable, V any, X any](m *map[K]V, f functions.Function[K, X]) []X {
	// Create a new slice with initial capacity equal to the number of entries in the input map.
	result := make([]X, 0, len(*m))

	// Iterate over each key in the input map.
	for k := range *m {
		// Apply the function to the current key and append the result to the result slice.
		result = append(result, f(k))
	}

	// Return the resulting slice.
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
func MapValues[K comparable, V any, X any](m *map[K]V, f functions.Function[V, X]) []X {
	// Create a new slice with initial capacity equal to the number of entries in the input map.
	result := make([]X, 0, len(*m))

	// Iterate over each value in the input map.
	for _, v := range *m {
		// Apply the function to the current value and append the result to the result slice.
		result = append(result, f(v))
	}

	// Return the resulting slice.
	return result
}

func MapValuesSorted[K comparable, V any, X any](m *map[K]V, f functions.Function[V, X], less functions.BiFunction[X, X, bool]) []X {
	// Create a new slice with initial capacity equal to the number of entries in the input map.
	result := make([]X, 0, len(*m))

	// Iterate over each value in the input map.
	for _, v := range *m {
		// Apply the function to the current value and append the result to the result slice.
		result = append(result, f(v))
	}

	// Sort the result slice using the provided less function.
	sort.Slice(result, func(i, j int) bool {
		return less(result[i], result[j])
	})

	// Return the resulting slice.
	return result
}

// MapWithKeys creates a new map with keys generated by applying the function f to each key in the input keys slice.
// The function f takes a key of type K and returns a key of type L and a value of type V.
// The resulting map has keys of type L and values of type V.
// The function assumes that the keys in the input keys slice are unique.
// If the function f generates duplicate keys, the resulting map will only contain the last value for each key.
func MapWithKeys[K comparable, L comparable, V any](keys []K, f functions.Function2[K, L, V]) *map[L]V {
	// Create a new map with initial capacity equal to the number of keys in the input keys slice.
	result := make(map[L]V, len(keys))

	// Iterate over each key in the input keys slice.
	for _, key := range keys {
		// Apply the function f to the current key and obtain a new key of type L and a value of type V.
		newKey, value := f(key)

		// Addf the new key-value pair to the resulting map.
		result[newKey] = value
	}

	// Return the resulting map.
	return &result
}

// MapWithValues applies a given function to each value in a slice and returns a new map with keys and values obtained by applying the function to each value.
// It's from logical awareness the generated keys for this map can be overridden by the function which is applied to each value.
//
// Parameters:
// - values: The input slice of values.
// - f: The function to apply to each value. It should return a key-value pair.
//
// Returns:
// - A new map with keys and values obtained by applying the function to each value.
func MapWithValues[K comparable, V any, X any](values []V, f functions.Function2[V, K, X]) *map[K]X {
	// Create a new map with initial capacity equal to the number of values in the input slice.
	result := make(map[K]X, len(values))

	// Iterate over each value in the input slice.
	for _, value := range values {
		// Apply the function to the current value and obtain a key-value pair.
		k, v := f(value)

		// Addf the key-value pair to the result map.
		result[k] = v
	}

	// Return the resulting map.
	return &result
}
