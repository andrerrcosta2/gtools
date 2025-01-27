// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package str

import (
	"github.com/andrerrcosta2/gtools/core/data/comparables"
	"github.com/andrerrcosta2/gtools/core/data/str/nodes"
	"github.com/andrerrcosta2/gtools/core/domain/constraints/prim"
)

type Map[K any, V any] interface {
	// Put adds a new key-value pair to the map.
	//
	// The key parameter is the key to be added to the map.
	// The value parameter is the value associated with the key.
	// No return value.
	Put(key K, value V)
	// Get retrieves the value associated with the given key from the map.
	//
	// The key parameter is the key to be searched in the map.
	// Returns the value associated with the key and a boolean indicating whether
	// the key was found.
	Get(key K) (V, bool)
	// Delete removes the key-value pair associated with the given key from the map.
	//
	// The key parameter is the key to be deleted from the map.
	// No return value.
	Delete(key K)
	// Contains checks if the given key exists in the map.
	//
	// The key parameter is the key to be searched in the map.
	// Returns a boolean indicating whether the key was found.
	Contains(key K) bool
	// Len returns the number of key-value pairs in the map.
	//
	// No parameters.
	// Returns the length of the map as an integer.
	Len() int
	// Clear clears all key-value pairs from the map.
	//
	// No parameters.
	// No return value.
	Clear()
	// Keys returns an array of keys from the map.
	//
	// No parameters.
	// Returns an array of keys of type K.
	Keys() []K
	// Values returns an array of values from the map.
	//
	// No parameters.
	// Returns an array of values of type V.
	Values() []V
	// Iterator returns an iterator for the map.
	//
	// The comparator parameter is an optional parameter that can be used to sort
	// the keys in the map.
	// Returns an iterator for the map.
	Iterator(comparator ...comparables.FunctionalComparator[K]) MapIterator[K, V]
}

type BiMap[K any, V any] interface {
	// Put adds a new key-value pair to the map.
	//
	// The key parameter is the key to be added to the map.
	// The value parameter is the value associated with the key.
	// No return value.
	Put(key K, value V)
	// GetByKey retrieves the value associated with the given key from the map.
	//
	// The key parameter is the key to be searched in the map.
	// Returns the value associated with the key and a boolean indicating whether
	// the key was found.
	GetByKey(key K) (V, bool)
	// GetByValue retrieves the key associated with the given value from the map.
	//
	// The value parameter is the value to be searched in the map.
	// Returns the key associated with the value and a boolean indicating whether
	// the value was found.
	GetByValue(value V) (K, bool)
	// DeleteByKey removes the key-value pair associated with the given key from the map.
	//
	// The key parameter is the key to be deleted from the map.
	// No return value.
	DeleteByKey(key K)
	// DeleteByValue removes the key-value pair associated with the given value from the map.
	//
	// The value parameter is the value to be deleted from the map.
	// No return value.
	DeleteByValue(value V)
	// ContainsKey checks if the given key exists in the map.
	//
	// The key parameter is the key to be searched in the map.
	// Returns a boolean indicating whether the key was found.
	ContainsKey(key K) bool
	// ContainsValue checks if the given value exists in the map.
	//
	// The value parameter is the value to be searched in the map.
	// Returns a boolean indicating whether the value was found.
	ContainsValue(value V) bool
	// Len returns the number of key-value pairs in the map.
	//
	// No parameters.
	// Returns the length of the map as an integer.
	Len() int
	// Clear clears all key-value pairs from the map.
	//
	// No parameters.
	// No return value.
	Clear()
	// Keys returns an array of keys from the map.
	//
	// No parameters.
	// Returns an array of keys of type K.
	Keys() []K
	// Values returns an array of values from the map.
	//
	// No parameters.
	// Returns an array of values of type V.
	Values() []V
	// Iterator returns an iterator for the map.
	//
	// The comparator parameter is an optional parameter that can be used to sort
	// the keys in the map.
	// Returns an iterator for the map.
	Iterator(comparator ...comparables.FunctionalComparator[K]) MapIterator[K, V]
}

// Entry is an interface that represents an entry in a map.
type Entry[K any, V any] interface {
	// Key returns the key of the entry.
	//
	// Returns the key of the entry as a value of type K.
	Key() K
	// Value returns the value of the entry.
	//
	// Returns the value of the entry as a value of type V.
	Value() V
	// String returns a string representation of the entry.
	//
	// Returns a string representation of the entry in the format "key: value".
	String() string
}

var _ nodes.KeyValue[string, any] = (Entry[string, any])(nil)

type Dictionary[K comparable, V any] interface {
	// Put adds a new key-value pair to the dictionary.
	//
	// It takes two parameters: the key of type K and the value of type V.
	// It returns nothing.
	Put(key K, value V)
	// EntrySet returns a map containing all the key-value pairs in the dictionary.
	//
	// It takes no parameters.
	// Returns a map containing all the key-value pairs in the dictionary.
	EntrySet() map[K]V
	// Entries returns a slice of Entry objects representing the key-value pairs in the dictionary.
	//
	// It takes no parameters.
	// Returns a slice of Entry objects representing the key-value pairs in the dictionary.
	Entries() []Entry[K, V]
	// Get returns the value associated with the given key.
	// If the key does not exist in the dictionary, it returns the zero value for the value type and false.
	//
	// Parameters:
	// - key: The key to search for in the dictionary.
	//
	// Returns:
	// - The value associated with the given key, if found.
	// - A boolean indicating whether the key was found.
	Get(key K) (V, bool)
	// IsEmpty checks if the dictionary is empty.
	//
	// Returns true if the dictionary is empty, false otherwise.
	IsEmpty() bool
	// Keys returns a slice of keys in the dictionary.
	//
	// Returns a slice of type []K containing all the keys in the dictionary.
	Keys() []K
	// Size returns the number of key-value pairs in the dictionary.
	//
	// Returns an integer representing the number of key-value pairs in the dictionary.
	Size() int
	// Values returns a slice of all the values in the dictionary.
	//
	// It takes no parameters.
	// Returns a slice of type []V containing all the values in the dictionary.
	Values() []V
}

// MergeDictionaries creates a new dictionary by merging the given dictionaries.
// If two dictionaries share a key, the value from the last dictionary in the list is used.
//
// The function takes a variable number of arguments of type Dictionary[K, V].
// The function returns a pointer to a Dictionary[K, V].
func MergeDictionaries[T prim.Hashable, S any](dictionaries ...Dictionary[T, S]) Dictionary[T, S] {
	if len(dictionaries) == 0 {
		// If the input list of dictionaries is empty, return nil.
		return nil
	}

	// Create a new dictionary by merging all the given dictionaries.
	// Iterate over each dictionary in the list.
	d := dictionaries[0]
	for i := 1; i < len(dictionaries); i++ {
		// Iterate over each key-value pair in the current dictionary.
		for key, value := range dictionaries[i].EntrySet() {
			// Add the key-value pair to the new dictionary.
			// If the key already exists, the old value is replaced by the new value.
			d.Put(key, value)
		}
	}

	// Return the new dictionary.
	return d
}
