// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package dicts

import (
	"fmt"

	"github.com/andrerrcosta2/gtools/core/data/str/dict"
	"github.com/andrerrcosta2/gtools/core/domain/constraints/prim"
)

// NewAnyEntry creates a new anyEntry struct with the given key and value.
//
// Parameters:
// - key: The key of the Entry.
// - value: The value of the Entry.
//
// Returns:
// - A pointer to the newly created Entry.
func NewAnyEntry[K any, V any](key K, value V) dict.Entry[K, V] {
	// Create a new Entry struct with the given key and value.
	return &anyEntry[K, V]{
		key:   key,   // ToSet the key of the Entry.
		value: value, // ToSet the value of the Entry.
	}
}

// anyEntry is a struct that implements the Entry interface for any type of key and value.
type anyEntry[K any, V any] struct {
	key   K
	value V
}

func (e *anyEntry[K, V]) Key() K {
	return e.key
}

func (e *anyEntry[K, V]) Value() V {
	return e.value
}

func (e *anyEntry[K, V]) String() string {
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
func NewComparableEntry[K comparable, V any](key K, value V) dict.Entry[K, V] {
	// Create a new ComparableEntry struct with the given key and value.
	return &cmpEntry[K, V]{
		key:   key,   // ToSet the key of the entry.
		value: value, // ToSet the value of the entry.
	}
}

type cmpEntry[K comparable, V any] struct {
	key   K
	value V
}

// Equal in the context of entries you can't have two different entries
// with the same key
func (e *cmpEntry[K, V]) Equal(other *cmpEntry[K, V]) bool {
	return e.key == other.key
}

func (e *cmpEntry[K, V]) Key() K {
	return e.key
}

func (e *cmpEntry[K, V]) Value() V {
	return e.value
}

func (e *cmpEntry[K, V]) String() string {
	return fmt.Sprintf("%v: %v", e.key, e.value)
}

func NewSortableEntry[K prim.Ordered, V any](key K, value V) dict.Entry[K, V] {
	return &sortableEntry[K, V]{
		key:   key,
		value: value,
	}
}

type sortableEntry[K prim.Ordered, V any] struct {
	key   K
	value V
}

func (e *sortableEntry[K, V]) Less(other *cmpEntry[K, V]) bool {
	return e.key < other.key
}

// Equal in the context of entries you can't have two different entries
// with the same key
func (e *sortableEntry[K, V]) Equal(other *cmpEntry[K, V]) bool {
	return e.key == other.key
}

func (e *sortableEntry[K, V]) Key() K {
	return e.key
}

func (e *sortableEntry[K, V]) Value() V {
	return e.value
}

func (e *sortableEntry[K, V]) String() string {
	return fmt.Sprintf("%v: %v", e.key, e.value)
}
