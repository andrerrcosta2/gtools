// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package maps

import (
	"errors"
	"fmt"
	"github.com/andrerrcosta2/gtools/pkg/functions"
	"github.com/andrerrcosta2/gtools/pkg/generics"
	"sort"
)

type Entry[K comparable, V any] struct {
	Key   K
	Value V
}

// NewEntry creates a new Entry struct with the given key and value.
//
// Parameters:
// - key: The key of the Entry.
// - value: The value of the Entry.
//
// Returns:
// - A pointer to the newly created Entry.
func NewEntry[K comparable, V any](key K, value V) *Entry[K, V] {
	// Create a new Entry struct with the given key and value.
	return &Entry[K, V]{
		Key:   key,   // Set the key of the Entry.
		Value: value, // Set the value of the Entry.
	}
}

type EntrySet[K comparable, V any] struct {
	entries map[K]*Entry[K, V]
}

// NewEntrySet creates a new EntrySet struct from a slice of Entry structs.
// As the underlying map of entries is a map, the keys of the input slice should be unique.
// As a side effect, these entries are always unsorted.
//
// Parameters:
// - e: a slice of Entry structs.
//
// Returns:
// - A pointer to the newly created EntrySet struct.
func NewEntrySet[K comparable, V any](e []*Entry[K, V]) *EntrySet[K, V] {
	// Create a new EntrySet struct with an empty map of entries.
	entryset := &EntrySet[K, V]{
		entries: make(map[K]*Entry[K, V]),
	}

	// Add each entry from the input slice to the EntrySet's map of entries.
	for _, entry := range e {
		entryset.entries[entry.Key] = entry
	}

	// Return the newly created EntrySet struct.
	return entryset
}

// Add adds one or more entries to the EntrySet.
//
// Parameters:
// - entry: The entries to be added to the EntrySet.
//
// Returns:
// - An error if any entry addition fails.
func (e *EntrySet[K, V]) Add(entry ...*Entry[K, V]) error {
	for _, ent := range entry {
		if err := addEntry[K, V](e, ent.Key, ent.Value); err != nil {
			return err
		}
	}
	return nil
}

// addEntry adds an entry to the EntrySet.
//
// Parameters:
// - set: The EntrySet to add the entry to.
// - key: The key of the entry.
// - value: The value of the entry.
//
// Returns:
// - An error if the key already exists in the EntrySet or if the key is nil.
func addEntry[K comparable, V any](set *EntrySet[K, V], key K, value V) error {
	// Check if the key is nil
	if &key == nil {
		return errors.New("key is required")
	}

	// Check if the key already exists in the EntrySet
	if _, ok := set.entries[key]; !ok {
		// If the key doesn't exist, create a new Entry with the given key and value and add it to the EntrySet
		set.entries[key] = &Entry[K, V]{Key: key, Value: value}
		return nil
	}

	// If the key already exists, return an error
	return fmt.Errorf("key already exists: %v", key)
}

// Keys returns all the keys in the EntrySet.
// The keys of the EntrySet aren't guaranteed to be in any particular order.
//
// Returns:
// - A slice of keys.
func (e *EntrySet[K, V]) Keys() []K {
	// Create a slice with initial capacity equal to the number of entries.
	result := make([]K, 0, len(e.entries))

	// Iterate over all the entries in the EntrySet.
	for _, entry := range e.entries {
		// Append the key of the current entry to the result slice.
		result = append(result, entry.Key)
	}

	// Return the resulting slice.
	return result
}

// Values returns all the values in the EntrySet.
// The values of the EntrySet aren't guaranteed to be in any particular order.
//
// Returns:
// - A slice of values.
func (e *EntrySet[K, V]) Values() []V {
	// Create a slice with initial capacity equal to the number of entries.
	result := make([]V, 0, len(e.entries))

	// Iterate over all the entries in the EntrySet.
	for _, entry := range e.entries {
		// Append the value of the current entry to the result slice.
		result = append(result, entry.Value)
	}

	// Return the resulting slice.
	return result
}

// Len returns the number of entries in the EntrySet.
//
// It returns an integer representing the length of the entries map.
func (e *EntrySet[K, V]) Len() int {
	// Return the length of the entries map.
	return len(e.entries)
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

// Map applies a given BiFunction to each key-value pair in a map and returns a new map.
//
// Parameters:
// - m: The input map.
// - f: The BiFunction to apply to each key-value pair. It should return a pointer to an Entry struct.
//
// Returns:
// - A new map with the same keys as the input map, but with values obtained by applying the BiFunction to each key-value pair.
func Map[K comparable, V any, L comparable, X any](m map[K]V, f functions.BiFunction[K, V, *Entry[L, X]]) map[L]X {
	// Create a new map with initial capacity equal to the number of entries in the input map.
	result := make(map[L]X, len(m))

	// Iterate over each key-value pair in the input map.
	for k, v := range m {
		// Apply the BiFunction to the current key-value pair and obtain an Entry pointer.
		entry := f(k, v)

		// Add the key-value pair from the Entry to the result map.
		result[entry.Key] = entry.Value
	}

	// Return the resulting map.
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
func MapEntries[K comparable, V any, L comparable, X any](m map[K]V, f functions.BiFunction[K, V, *Entry[L, X]]) *EntrySet[L, X] {
	// Create a new EntrySet with initial capacity equal to the number of entries in the input map.
	entries := NewEntrySet(make([]*Entry[L, X], 0, len(m)))

	// Iterate over each key-value pair in the input map.
	for k, v := range m {
		// Apply the BiFunction to the current key-value pair and obtain an Entry pointer.
		entry := f(k, v)

		// Add the Entry to the EntrySet.
		entries.Add(entry)
	}

	// Return the resulting EntrySet.
	return entries
}

// MapValues applies a given function to each value in a map and returns a new slice of the results.
//
// Parameters:
// - m: The input map.
// - f: The function to apply to each value.
//
// Returns:
// - A new slice with the same length as the input map, but with values obtained by applying the function to each value.
func MapValues[K comparable, V any, X any](m map[K]V, f functions.Function[V, X]) []X {
	// Create a new slice with initial capacity equal to the number of entries in the input map.
	result := make([]X, 0, len(m))

	// Iterate over each value in the input map.
	for _, v := range m {
		// Apply the function to the current value and append the result to the result slice.
		result = append(result, f(v))
	}

	// Return the resulting slice.
	return result
}

func MapValuesSorted[K comparable, V any, X any](m map[K]V, f functions.Function[V, X], less functions.BiFunction[X, X, bool]) []X {
	// Create a new slice with initial capacity equal to the number of entries in the input map.
	result := make([]X, 0, len(m))

	// Iterate over each value in the input map.
	for _, v := range m {
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

// FlatValues applies a function to each element in a map of slices and returns a new slice with the transformed values.
//
// Parameters:
// - m: The input map with slices of values.
// - f: The function to apply to each value.
//
// Returns:
// - A new slice with the transformed values obtained by applying the function to each element in the map.

func FlatValues[K comparable, V any, X any](m map[K][]V, f functions.Function[V, X]) []X {
	// Create a slice to store the flattened and transformed values.
	var result []X

	// Iterate over each slice of values in the map.
	for _, slice := range m {
		// Apply the function to each element in the slice and append the result to the final slice.
		for _, v := range slice {
			result = append(result, f(v))
		}
	}

	// Return the flattened and transformed slice.
	return result
}

func FlatValuesSorted[K comparable, V any, X any](m map[K][]V, f functions.Function[V, X], less functions.BiFunction[X, X, bool]) []X {
	var result []X
	for _, slice := range m {
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

// MapKeys applies a given function to each key in a map and returns a new slice of the results.
//
// Parameters:
// - m: The input map.
// - f: The function to apply to each key.
//
// Returns:
// - A new slice with the same length as the input map, but with values obtained by applying the function to each key.
func MapKeys[K comparable, V any, X any](m map[K]V, f functions.Function[K, X]) []X {
	// Create a new slice with initial capacity equal to the number of entries in the input map.
	result := make([]X, 0, len(m))

	// Iterate over each key in the input map.
	for k := range m {
		// Apply the function to the current key and append the result to the result slice.
		result = append(result, f(k))
	}

	// Return the resulting slice.
	return result
}

// Cast applies a given function to each key-value pair in a map and returns a new slice of the results.
//
// Parameters:
// - m: The input map.
// - f: The function to apply to each key-value pair.
//
// Returns:
// - A new slice with the same length as the input map, but with values obtained by applying the function to each key-value pair.
func Cast[K comparable, V any, L comparable, X any](m map[K]V, f functions.BiFunction[K, V, generics.BiTypedInterface[L, X]]) []generics.BiTypedInterface[L, X] {
	// Create a new slice with initial capacity equal to the number of entries in the input map.
	cast := make([]generics.BiTypedInterface[L, X], 0, len(m))

	// Iterate over each key-value pair in the input map.
	for k, v := range m {
		// Apply the function to the current key-value pair and append the result to the result slice.
		cast = append(cast, f(k, v))
	}

	// Return the resulting slice.
	return cast
}

// MapWithKeys creates a new map with keys generated by applying the function f to each key in the input keys slice.
// The function f takes a key of type K and returns a key of type L and a value of type V.
// The resulting map has keys of type L and values of type V.
// The function assumes that the keys in the input keys slice are unique.
// If the function f generates duplicate keys, the resulting map will only contain the last value for each key.
func MapWithKeys[K comparable, L comparable, V any](keys []K, f functions.Function2[K, L, V]) map[L]V {
	// Create a new map with initial capacity equal to the number of keys in the input keys slice.
	result := make(map[L]V, len(keys))

	// Iterate over each key in the input keys slice.
	for _, key := range keys {
		// Apply the function f to the current key and obtain a new key of type L and a value of type V.
		newKey, value := f(key)

		// Add the new key-value pair to the resulting map.
		result[newKey] = value
	}

	// Return the resulting map.
	return result
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
func MapWithValues[K comparable, V any, X any](values []V, f functions.Function2[V, K, X]) map[K]X {
	// Create a new map with initial capacity equal to the number of values in the input slice.
	result := make(map[K]X, len(values))

	// Iterate over each value in the input slice.
	for _, value := range values {
		// Apply the function to the current value and obtain a key-value pair.
		k, v := f(value)

		// Add the key-value pair to the result map.
		result[k] = v
	}

	// Return the resulting map.
	return result
}

// Fetch returns a map with keys and values obtained by applying the function f to each entry in the input slice.
// The function f takes an entry of type generics.BiTypedInterface[K, V] and returns a key of type K and a value of type V.
// The input slice entries must have unique keys.
// If there are duplicate keys in the input slice, the last value for each key will be kept in the resulting map.
//
// Parameters:
// - entries: The input slice of entries.
// - f: The function to apply to each entry. It should return a key-value pair.
//
// Returns:
// - A map with keys and values obtained by applying the function to each entry.
func Fetch[K comparable, V any](entries []generics.BiTypedInterface[K, V], f functions.Function2[generics.BiTypedInterface[K, V], K, V]) map[K]V {
	// Create a new map with initial capacity equal to the number of entries in the input slice.
	result := make(map[K]V, len(entries))

	// Iterate over each entry in the input slice.
	for _, entry := range entries {
		// Apply the function f to the current entry and obtain a key of type K and a value of type V.
		k, v := f(entry)

		// Add the key-value pair to the result map.
		result[k] = v
	}

	// Return the resulting map.
	return result
}
