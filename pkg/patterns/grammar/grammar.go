// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package grammar

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/andrerrcosta2/gtools/core/data/str"
	"github.com/andrerrcosta2/gtools/core/domain/constraints/prim"
	io2 "github.com/andrerrcosta2/gtools/core/io"
	"github.com/andrerrcosta2/gtools/patterns/symbols"
	"io"
	"os"
)

// Dictionary creates a new empty dictionary of symbols.
func Dictionary[K prim.Hashable, V ~[]S, S symbols.Logical](entries ...str.Entry[K, V]) str.Dictionary[K, V] {
	d := make(SymbolDictionary[K, V, S], len(entries))

	// Iterate over the given entries and add them to the dictionary.
	for _, entry := range entries {
		d.Put(entry.Key(), entry.Value())
	}

	return &d
}

type SymbolDictionary[K prim.Hashable, V ~[]S, S symbols.Logical] map[K]V

func (d *SymbolDictionary[K, V, S]) EntrySet() map[K]V {
	return *d
}

// Put adds a new key-value pair to the dictionary.
// If the key already exists, the old value is replaced by the new value.
func (d *SymbolDictionary[K, V, S]) Put(key K, value V) {
	(*d)[key] = value
}

func (d *SymbolDictionary[K, V, S]) putOrAppend(key K, value S) {
	if existing, ok := (*d)[key]; ok {
		(*d)[key] = append(existing, value)
	} else {
		(*d)[key] = V{value}
	}
}

// Get returns the value associated with the given key.
// If the key does not exist, it returns the zero value for the ByteSymbol type and false.
func (d *SymbolDictionary[K, V, S]) Get(key K) (V, bool) {
	// Get the value associated with the given key
	value, ok := (*d)[key]
	// Return the value and a boolean indicating whether the key exists
	return value, ok
}

// Size returns the number of key-value pairs in the dictionary.
func (d *SymbolDictionary[K, V, S]) Size() int {
	// Return the size of the dictionary
	return len(*d)
}

// IsEmpty returns true if the dictionary is empty, false otherwise.
func (d *SymbolDictionary[K, V, S]) IsEmpty() bool {
	// Return true if the dictionary is empty, false otherwise
	return d.Size() == 0
}

// Entries returns a slice of Entry objects representing the key-value pairs in the dictionary.
func (d *SymbolDictionary[K, V, S]) Entries() []str.Entry[K, V] {
	var entries []str.Entry[K, V]
	// Iterate over each key-value pair in the dictionary
	for key, value := range *d {
		for _, v := range value {
			// Create a new Entry from the current key-value pair
			entry := SymbolEntry[K, V, S](key, V{v})
			// Append the Entry to the slice of Entries
			entries = append(entries, entry)
		}
	}
	// Return the slice of Entries
	return entries
}

// Keys returns a slice of keys in the dictionary.
func (d *SymbolDictionary[K, V, S]) Keys() []K {
	// Create a slice of keys
	var keys []K
	// Iterate over each key-value pair in the dictionary
	for key := range *d {
		// Append the key to the slice of keys
		keys = append(keys, key)
	}
	// Return the slice of keys
	return keys
}

// Values returns a slice of values in the dictionary.
func (d *SymbolDictionary[K, V, S]) Values() []V {
	// Create a slice of values
	var values []V
	// Iterate over each key-value pair in the dictionary
	for _, value := range *d {
		// Append the value to the slice of values
		values = append(values, value)
	}
	// Return the slice of values
	return values
}

// CharSet returns a set of all characters used in the values of the dictionary.
//
// The returned map is a set of runes, where each rune is a character used in at
// least one of the values of the dictionary. The set is represented as a map of
// runes to empty structs, as this is a common idiom in Go for representing a
// set.
//
// This function is useful when you want to know all the characters used in the
// values of the dictionary. For example, you can use it to pre-allocate a
// buffer for parsing a string, or to check if a string contains any characters
// that are not in the dictionary.
func (d *SymbolDictionary[K, V, S]) CharSet() map[rune]struct{} {
	// Create a map to store the characters used in the values of the dictionary
	charSet := make(map[rune]struct{}, len(*d))

	// Iterate over each value in the dictionary
	for _, values := range *d {
		// Iterate over each character in the value
		for _, value := range values {
			// Iterate over each character in the value
			for char := range value.CharSet() {
				// Append the character to the set
				charSet[char] = struct{}{}
			}
		}
	}

	// Return the set of characters
	return charSet
}

// LoadFromFile reads a file and populates the SymbolDictionary with symbols
func (d *SymbolDictionary[K, V, S]) LoadFromFile(filePath string, fileType DictionaryFileType) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	// It's only reading, I want to log it anyway
	defer io2.CloseOrLog(file, "grammar.SymbolDictionary.LoadFromFile")

	// Parse the file based on fileType
	switch fileType {
	case JSONDictionary:
		return d.loadFromJSON(file)
	case CSVDictionary:
		return d.loadFromCSV(file)
	default:
		return fmt.Errorf("unsupported file type: %v", fileType)
	}
}

// loadFromJSON load a JSON file into a SymbolDictionary.
// The file should contain a JSON map with strings as keys and symbols as values.
func (d *SymbolDictionary[K, V, S]) loadFromJSON(file *os.File) error {
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(d); err != nil {
		return fmt.Errorf("could not decode json file: %w", err)
	}
	return nil
}

// loadFromCSV load a CSV file into a SymbolDictionary.
// The file should contain a CSV file with two columns, the first column is the key, and the second column is the value.
// The values are expected to be symbols in the form of a string.
func (d *SymbolDictionary[K, V, S]) loadFromCSV(file *os.File) error {
	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("could not read csv file: %w", err)
		}

		// The csv file should have at least two columns, the first column is the key, and the
		// next are the symbol values
		if len(record) == 2 {
			if key, err := prim.ToHashable[K](record[0]); err == nil {
				// Convert the value to a Symbol
				if value, err := mapToSymbol[string, S](record[1:]...); err == nil {
					d.putOrAppend(key, value)
				} else {
					return fmt.Errorf("could not convert value to symbol: %w", err)
				}
			} else {
				return fmt.Errorf("could not convert key to any hashable type: %w", err)
			}
		} else {
			return fmt.Errorf("csv file should have two columns, found %d", len(record))
		}
	}
	return nil
}

var _ str.Dictionary[string, []symbols.Logical] = (*SymbolDictionary[string, []symbols.Logical, symbols.Logical])(nil)

// SymbolEntry creates a new Entry for a SymbolDictionary.
//
// The key of the Entry is of type K, and the value is of type ByteSymbol.
// The function takes two parameters, the key and the value, and returns a new Entry.
//
// The Entry is created using the symbolEntry struct, which contains the key and value.
func SymbolEntry[K prim.Hashable, V ~[]S, S symbols.Logical](key K, value V) str.Entry[K, V] {
	return &symbolEntry[K, V, S]{key: key, value: value}
}

type symbolEntry[K prim.Hashable, V ~[]S, S symbols.Logical] struct {
	key   K
	value V
}

// Key returns the key of the Entry.
//
// The key is of type K.
//
// Returns the key of the Entry.
func (e *symbolEntry[K, V, S]) Key() K {
	return e.key
}

// Value returns the value of the Entry.
//
// The value is of type ByteSymbol.
//
// Returns the value of the Entry.
func (e *symbolEntry[K, V, S]) Value() V {
	return e.value
}

// String returns a string representation of the Entry.
//
// The string representation is a JSON-like string,
// with the key and value enclosed in curly braces,
// separated by a colon, and with the key and value
// separated by a comma.
//
// Returns a string representation of the Entry.
func (e *symbolEntry[K, V, S]) String() string {
	return fmt.Sprintf("{ key: %v, value: %v }", e.key, e.value)
}

var _ str.Entry[string, []symbols.Logical] = (*symbolEntry[string, []symbols.Logical, symbols.Logical])(nil)

type DictionaryFileType int

const (
	JSONDictionary = iota
	CSVDictionary
)

type PatternTrieDictionary SymbolDictionary[string, []symbols.Logical, symbols.Logical]
