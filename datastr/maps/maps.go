// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package maps

import (
	"errors"
	"fmt"
	"github.com/andrerrcosta2/gtools/core/domain/functions"
	"github.com/andrerrcosta2/gtools/core/domain/gtools"
	"github.com/andrerrcosta2/gtools/core/sortables"
)

// NewAnyEntry creates a new AnyEntry struct with the given key and value.
//
// Parameters:
// - key: The key of the Entry.
// - value: The value of the Entry.
//
// Returns:
// - A pointer to the newly created Entry.
func NewAnyEntry[K any, V any](key K, value V) *AnyEntry[K, V] {
	// Create a new Entry struct with the given key and value.
	return &AnyEntry[K, V]{
		key:   key,   // ToSet the key of the Entry.
		value: value, // ToSet the value of the Entry.
	}
}

// AnyEntry is a struct that implements the Entry interface for any type of key and value.
type AnyEntry[K any, V any] struct {
	key   K
	value V
}

func (e *AnyEntry[K, V]) Key() K {
	return e.key
}

func (e *AnyEntry[K, V]) Value() V {
	return e.value
}

func (e *AnyEntry[K, V]) String() string {
	return fmt.Sprintf("%v: %v", e.key, e.value)
}

// NewComparableEntry creates a new ComparableEntry struct with the given key and value.
//
// Parameters:
// - key: The key of the entry.
// - value: The value of the entry.
//
// Returns:
// - A pointer to the newly created ComparableEntry.
func NewComparableEntry[K comparable, V any](key K, value V) *ComparableEntry[K, V] {
	// Create a new ComparableEntry struct with the given key and value.
	return &ComparableEntry[K, V]{
		key:   key,   // ToSet the key of the entry.
		value: value, // ToSet the value of the entry.
	}
}

type ComparableEntry[K comparable, V any] struct {
	key   K
	value V
}

func (e *ComparableEntry[K, V]) Key() K {
	return e.key
}

func (e *ComparableEntry[K, V]) Value() V {
	return e.value
}

func (e *ComparableEntry[K, V]) String() string {
	return fmt.Sprintf("%v: %v", e.key, e.value)
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
func NewEntrySet[K comparable, V any](e ...*ComparableEntry[K, V]) *EntrySet[K, V] {
	// Create a new Entries struct with an empty map of entries.
	entryset := &EntrySet[K, V]{
		entries: make(map[K]*ComparableEntry[K, V]),
	}

	// Addf each entry from the input slice to the Entries's map of entries.
	for _, entry := range e {
		entryset.entries[entry.Key()] = entry
	}

	// Return the newly created Entries struct.
	return entryset
}

type EntrySet[K comparable, V any] struct {
	entries map[K]*ComparableEntry[K, V]
}

// Add adds one or more entries to the EntrySet.
//
// Parameters:
// - entry: The entries to be added to the EntrySet.
//
// Returns:
// - An error if any entry addition fails.
func (e *EntrySet[K, V]) Add(entry ...*ComparableEntry[K, V]) error {
	for _, ent := range entry {
		if err := addEntry[K, V](e, ent.Key(), ent.Value()); err != nil {
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

	// Check if the key already exists in the Entries
	if _, ok := set.entries[key]; !ok {
		// If the key doesn't exist, create a new Entry with the given key and value and add it to the Entries
		set.entries[key] = &ComparableEntry[K, V]{key: key, value: value}
		return nil
	}

	// If the key already exists, return an error
	return fmt.Errorf("key already exists: %v", key)
}

// Get returns the value associated with the given key in the EntrySet.
// If the key is not found, the zero value and false are returned.
//
// Parameters:
// - key: The key of the entry.
//
// Returns:
// - The value associated with the key, if found.
func (e *EntrySet[K, V]) Get(key K) (V, bool) {
	v, ok := e.entries[key]
	if !ok {
		var zero V
		return zero, false
	}
	return v.Value(), ok
}

// Put adds an entry to the EntrySet.
// If the key already exists, the value is updated.
//
// Parameters:
// - key: The key of the entry.
// - value: The value of the entry.
//
// Returns:
// - An error if the key is nil.
func (e *EntrySet[K, V]) Put(key K, value V) error {
	if &key == nil {
		return errors.New("key is required")
	}
	e.entries[key] = &ComparableEntry[K, V]{key: key, value: value}
	return nil
}

// Keys returns all the keys in the EntrySet.
// The keys of the EntrySet aren't guaranteed to be in any particular order.
//
// Returns:
// - A slice of keys.
func (e *EntrySet[K, V]) Keys() []K {
	// Create a slice with initial capacity compare to the number of entries.
	result := make([]K, 0, len(e.entries))

	// Iterate over all the entries in the Entries.
	for _, entry := range e.entries {
		// Append the key of the current entry to the result slice.
		result = append(result, entry.Key())
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
	// Create a slice with initial capacity compare to the number of entries.
	result := make([]V, 0, len(e.entries))

	// Iterate over all the entries in the Entries.
	for _, entry := range e.entries {
		// Append the value of the current entry to the result slice.
		result = append(result, entry.Value())
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

type EntrySetOf[K gtools.SortableOf, V any] struct {
	set map[string]AnyEntry[K, V]
	cmp sortables.Comparator[K]
}

// Map applies a given BiFunction to each key-value pair in a map and returns a new map.
//
// Parameters:
// - m: The input map.
// - f: The BiFunction to apply to each key-value pair. It should return a pointer to an Entry struct.
//
// Returns:
// - A new map with the same keys as the input map, but with values obtained by applying the BiFunction to each key-value pair.
func Map[K comparable, V any, L comparable, X any](m *map[K]V, f functions.BiFunction[K, V, *ComparableEntry[L, X]]) *map[L]X {
	// Create a new map with initial capacity compare to the number of entries in the input map.
	result := make(map[L]X, len(*m))

	// Iterate over each key-value pair in the input map.
	for k, v := range *m {
		// Apply the BiFunction to the current key-value pair and obtain an Entry pointer.
		entry := f(k, v)

		// Addf the key-value pair from the Entry to the result map.
		result[entry.Key()] = entry.Value()
	}

	// Return the resulting map.
	return &result
}

// MapEntries applies a given BiFunction to each key-value pair in a map and returns a new EntrySet.
//
// Parameters:
// - m: The input map.
// - f: The BiFunction to apply to each key-value pair. It should return a pointer to an Entry struct.
//
// Returns:
// - A new EntrySet with the same keys as the input map, but with values obtained by applying the BiFunction to each key-value pair.
func MapEntries[K comparable, V any, L comparable, X any](m map[K]V, f functions.BiFunction[K, V, *ComparableEntry[L, X]]) *EntrySet[L, X] {
	// Create a new Entries with initial capacity compare to the number of entries in the input map.
	entries := NewEntrySet(make([]*ComparableEntry[L, X], 0, len(m))...)

	// Iterate over each key-value pair in the input map.
	for k, v := range m {
		// Apply the BiFunction to the current key-value pair and obtain an Entry pointer.
		entry := f(k, v)

		// Addf the Entry to the Entries.
		entries.Add(entry)
	}

	// Return the resulting Entries.
	return entries
}
