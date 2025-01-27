// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package grammar

import (
	"bytes"
	"fmt"
	"github.com/andrerrcosta2/gtools/core/domain/constraints/prim/bins"
	"github.com/andrerrcosta2/gtools/core/domain/functions"
	"regexp"
)

type Symbol interface {
	fmt.Stringer
	// After returns a new symbol that is the sub symbol of the current symbol
	// from the given index to the end of the symbol.
	After(i int) Symbol
	// Append returns a new symbol that is the result of appending the given
	// other symbol to the end of the current symbol.
	Append(other Symbol) Symbol
	// Before returns a new symbol that is the sub symbol of the current symbol
	// from the beginning of the symbol to the given index.
	Before(i int) Symbol
	// Between returns a new symbol that is the sub symbol of the current symbol
	// from the given i index to the given j index.
	Between(i, j int) Symbol
	// Bytes returns the byte representation of the symbol.
	Bytes() []byte
	// CharSet returns a set of characters used in the symbol.
	// It handles string symbols, numbers, and any other type.
	// For numbers, it converts them to their string representation and treats digits as characters.
	// For other types, it ignores them.
	// It also handles composed sequences of characters as a single character.
	// For example, if the symbol is "foo-bar", it will treat it as a single character.
	CharSet() map[rune]struct{}
	// Contains returns true if the given other symbol is a sub symbol of the current symbol.
	// It returns false otherwise.
	Contains(other Symbol) bool
	// CutAll removes all occurrences of the given other symbol from the current symbol.
	// It returns a new symbol with all occurrences of the other symbol removed.
	//
	// For example, if the current symbol is "test-test2" and the given other symbol
	// is "test", it will return a symbol with only "2".
	CutAll(other Symbol) Symbol
	// CutFirst removes the first occurrence of the given other symbol from the current symbol.
	// It returns a new symbol with the first occurrence of the other symbol removed.
	//
	// For example, if the current symbol is "test-test2" and the given other symbol
	CutFirst(symbol Symbol) Symbol
	// CutLast removes the last occurrence of the given other symbol from the current symbol.
	// It returns a new symbol with the last occurrence of the other symbol removed.
	//
	// For example, if the current symbol is "test-test2" and the given other symbol
	// is "test", it will return a symbol with only "test2".
	CutLast(other Symbol) Symbol
	// EndsWith returns true if the given other symbol is a suffix of the current symbol.
	// It returns false otherwise.
	//
	// For example, if the current symbol is "test" and the given other symbol is "t",
	EndsWith(other Symbol) bool
	// Equal returns true if the given other symbol is equal to the current symbol.
	// It returns false otherwise.
	Equal(other Symbol) bool
	// Empty returns a new empty symbol.
	Empty() Symbol
	// ExtractSymbol splits the given string into two parts, the first part being the
	// given symbol and the second part being the rest of the string.
	//
	// It returns the symbol and the rest of the string. If the symbol is longer
	// than the string, it returns an empty string for the symbol and the full
	// string as the rest.
	//
	// If the split symbol is not found, it returns an empty symbol and the entire current symbol.
	ExtractSymbol(split Symbol) (symbol, after Symbol)
	// Format formats the given placeholder symbol with the given format symbol.
	// It replaces any occurrence of the placeholder symbol in the format symbol with the given placeholder symbol.
	// For example, if the format symbol is "hello %s" and the placeholder symbol is "world", it will return a symbol with the value "hello world".
	Format(format Symbol, placeholder Symbol) Symbol
	// IsEmpty returns true if the symbol is empty, false otherwise.
	IsEmpty() bool
	// IndexOf returns the index of the first occurrence of the given other symbol in the current symbol.
	// It returns -1 if the symbol is not found.
	// It takes a Symbol as an argument and returns an int and a boolean.
	IndexOf(other Symbol) (int, bool)
	// LastIndexOf returns the index of the last occurrence of the given other symbol in the current symbol.
	// It returns -1 if the symbol is not found.
	// It takes a Symbol as an argument and returns an int and a boolean.
	LastIndexOf(other Symbol) (int, bool)
	// Len returns the length of the symbol.
	Len() int
	// Map applies a given BiFunction to each character in the symbol and returns a new symbol.
	// It takes a BiFunction that takes two parameters: the index of the character in the symbol and the character itself.
	// It returns a new symbol with the transformed characters.
	// For example, if the symbol is "test-test2" and the BiFunction is a lambda that takes the index and the character and returns the character incremented by 1,
	// it will return a symbol with the value "uvft-uvfu3".
	Map(fn functions.BiFunction[int, byte, byte]) Symbol
	// Remove returns a new symbol that is the result of removing the sub symbol
	// from the given i index to the given j index from the current symbol.
	Remove(i, j int) Symbol
	// ReplaceAll replaces all occurrences of the given target symbol in the current symbol
	// with the given replacement symbol and returns the new symbol.
	// It returns a boolean indicating whether the replacement was performed or not.
	ReplaceAll(target Symbol, replacement Symbol) (Symbol, bool)
	// ReplaceEach replaces all occurrences of a given target symbol with a given replacement symbol
	// within this symbol. This is similar to ReplaceAll, but it takes a slice of replacement symbols,
	// and each replacement symbol is used in order. If there are more occurrences of the target symbol
	ReplaceEach(target Symbol, replacements ...Symbol) (Symbol, bool)
	// ReplaceFirst replaces the first occurrence of the given target symbol in the current symbol
	// with the given replacement symbol and returns the new symbol.
	// It returns a boolean indicating whether the replacement was performed or not.
	ReplaceFirst(target Symbol, replacement Symbol) (Symbol, bool)
	// ReplaceLast replaces the last occurrence of the given target symbol in the current symbol
	// with the given replacement symbol and returns the new symbol.
	// It returns a boolean indicating whether the replacement was performed or not.
	ReplaceLast(target Symbol, replacement Symbol) (Symbol, bool)
	// ReplaceMany replaces all occurrences of a given target symbol with a given replacement symbol
	// within this symbol. This is similar to ReplaceAll, but it takes a slice of target symbols,
	// and each target symbol is used in order. If there are more occurrences of the target symbols
	// than the number of replacement symbols, the replacement symbols are cycled over.
	// It returns a boolean indicating whether the replacement was performed or not.
	ReplaceMany(replacement Symbol, targets ...Symbol) (Symbol, bool)
	// StartsWith returns true if the current symbol starts with the given other symbol.
	// It returns false otherwise.
	StartsWith(other Symbol) bool
}

// NewSymbol takes an interface{} and attempts to convert it to a ByteSymbol.
//
// Symbols are simply sequences of bytes used to identify patterns.
//
// If the input is a string, it is converted to a []byte.
// If the input is already a []byte, it is returned as is.
// If the input is a []byte, it is returned as is.
// If the input is a ByteSymbol, it is returned as is.
// Otherwise, the function tries to marshal the input to JSON and return the resulting []byte.
// If all else fails, an error is returned.
func NewSymbol(s any) (Symbol, error) {
	return newBytesSymbol(s)
}

// newBytesSymbol takes an interface{} and attempts to convert it to a ByteSymbol.
//
// Symbols are simply sequences of bytes used to identify patterns.
//
// If the input is a string, it is converted to a []byte.
// If the input is already a []byte or a ByteSymbol, it is returned as is.
// Otherwise, the function tries to marshal the input to JSON and return the resulting []byte.
// If all else fails, an error is returned.
func newBytesSymbol(s any) (ByteSymbol, error) {
	switch v := s.(type) {
	case string:
		// Convert string to []byte
		return ByteSymbol(v), nil
	case []byte:
		return v, nil // Already []byte or ByteSymbol, return as is
	case ByteSymbol:
		return v, nil
	default:
		return bins.ToBytes(v)
	}
}

// ByteSymbol is a sequence of bytes that represents a pattern.
//
// It is used to identify specific patterns in general expressions.
type ByteSymbol []byte

var _ Symbol = ByteSymbol{}

// After returns a new ByteSymbol which is a subset of the receiver, starting at index `i`.
//
// This is a convenience method for building symbols.
func (s ByteSymbol) After(i int) Symbol {
	if i >= len(s) {
		return s.Empty()
	}
	if i <= 0 {
		return s
	}
	return s[i+1:]
}

// Append returns a new ByteSymbol which is the concatenation of the receiver and the other ByteSymbol.
//
// This is a convenience method for building symbols.
func (s ByteSymbol) Append(other Symbol) Symbol {
	return append(s, other.Bytes()...)
}

// Before returns a new ByteSymbol which is a subset of the receiver, up to index `i`.
//
// This is a convenience method for building symbols.
func (s ByteSymbol) Before(i int) Symbol {
	if i <= 0 {
		return s.Empty()
	}
	if i >= len(s) {
		return s
	}
	return s[:i]
}

// Between returns a substring of the symbol, starting at index `i` and ending at `j`.
// If `j` is negative, it is treated as the index from the end of the string.
func (s ByteSymbol) Between(i, j int) Symbol {
	if i > j || i >= s.Len() {
		// If i > j, the result is an empty string
		return s.Empty()
	}
	if j > s.Len() {
		j = s.Len()
	}
	if i < 0 {
		i = 0
	}
	if j < 0 {
		j = 0
	}
	return s[i:j]
}

// Bytes return the symbol as a slice of bytes.
//
// This is a convenience method for working with the symbol.
func (s ByteSymbol) Bytes() []byte {
	return s
}

// CharSet returns a set of all characters in the symbol.
//
// This is a convenience method for working with the symbol.
func (s ByteSymbol) CharSet() map[rune]struct{} {
	// Create a new map to store the result
	var result = make(map[rune]struct{})
	// Iterate over each character in the symbol
	for _, char := range s {
		result[rune(char)] = struct{}{}
	}
	return result
}

// Contains returns true if the receiver contains the other symbol.
//
// This is a convenience method for working with the symbol.
func (s ByteSymbol) Contains(other Symbol) bool {
	return bytes.Contains(s, other.Bytes())
}

// CutAll removes all occurrences of the given symbol from the receiver.
//
// Returns the modified symbol.
func (s ByteSymbol) CutAll(other Symbol) Symbol {
	if other.Len() == 0 {
		return s
	}
	var out = s
	for {
		idx, found := out.IndexOf(other)
		if !found {
			return out
		}
		out = out.Remove(idx, idx+other.Len()).Bytes()
	}
}

// CutFirst removes the first occurrence of the given symbol from the receiver.
//
// Returns the modified symbol.
//
// If the given symbol is not found, the receiver is returned unchanged.
func (s ByteSymbol) CutFirst(symbol Symbol) Symbol {
	if idx, found := s.IndexOf(symbol); found {
		return s.Remove(idx, idx+symbol.Len())
	}
	// If the symbol is not found, return the receiver unchanged.
	return s
}

// CutLast removes the last occurrence of the given symbol from the receiver.
//
// symbol - the symbol to be removed.
// Returns the modified symbol. If the given symbol isn't found, the receiver is returned unchanged.
func (s ByteSymbol) CutLast(symbol Symbol) Symbol {
	if idx, found := s.LastIndexOf(symbol); found {
		return s.Remove(idx, idx+symbol.Len())
	}
	// If the symbol is not found, return the receiver unchanged.
	return s
}

// EndsWith returns true if the receiver ends with the given symbol.
//
// This is a convenience method for working with the symbol.
func (s ByteSymbol) EndsWith(other Symbol) bool {
	return bytes.HasSuffix(s, other.Bytes())
}

// Equal returns true if the symbol is equal to the other symbol.
//
// This is a convenience method for working with the symbol.
func (s ByteSymbol) Equal(other Symbol) bool {
	return bytes.Equal(s, other.Bytes())
}

func (s ByteSymbol) Empty() Symbol {
	return ByteSymbol{}
}

// ExtractSymbol splits the given string into two parts, the first part being the
// given symbol and the second part being the rest of the string.
//
// It returns the symbol and the rest of the string. If the symbol is longer
// than the string, it returns an empty string for the symbol and the full
// string as the rest.
//
// The split is done starting from the beginning of the string, and the symbol
// is matched against the start of the string.
//
// If the symbol is longer than the string, it can't match, so the function
// returns an empty string for the symbol and the full string as the rest.
func (s ByteSymbol) ExtractSymbol(split Symbol) (symbol, after Symbol) {
	// If the symbol is longer than the string, it can't match
	// If the symbol is longer than the string, it can't match
	if idx, found := s.IndexOf(split); !found {
		return s.Empty(), s
	} else {
		// ExtractSymbol the symbol between the split and its rest
		return split, s.After(idx + split.Len() - 1)
	}
}

// IsEmpty returns true if the symbol is empty.
//
// This is a convenience method for checking if the symbol is empty.
func (s ByteSymbol) IsEmpty() bool {
	// Check if the length of the symbol is 0
	return len(s) == 0
}

// IndexOf returns the index of the first occurrence of the given symbol in the receiver.
// It returns the index and a boolean indicating whether the symbol was found.
// If the symbol is not found, it returns -1 and false.
func (s ByteSymbol) IndexOf(other Symbol) (int, bool) {
	// Use the bytes.Serial function to search for the given symbol in the receiver
	// If the symbol is found, return the index and true
	// If the symbol is not found, return -1 and false
	idx := bytes.Index(s, other.Bytes())
	if idx != -1 {
		return idx, true
	}
	return -1, false
}

// Format returns a new symbol by replacing all occurrences of the placeholder symbol with the receiver in the given format symbol.
//
// The format symbol is the symbol to be formatted.
// The placeholder symbol is the symbol to be replaced in the format symbol.
// It returns the formatted symbol.
func (s ByteSymbol) Format(format Symbol, placeholder Symbol) Symbol {
	// Use the bytes.ReplaceAll function to replace all occurrences of the placeholder symbol with the receiver in the format symbol
	// The bytes.ReplaceAll function returns a new byte slice with all occurrences of the placeholder symbol replaced with the receiver
	var output ByteSymbol = bytes.ReplaceAll(format.Bytes(), placeholder.Bytes(), s)
	return output
}

// IsValid checks if the symbol matches the given regular expression.
// It returns true if the symbol is valid according to the regular expression, otherwise false.
//
// The regular expression is compiled and the symbol is matched against it.
func (s ByteSymbol) IsValid(rgx string) bool {
	// Compile the regular expression
	re, err := regexp.Compile(rgx)
	if err != nil {
		// If the regular expression is invalid, return false
		return false
	}

	// Check if the symbol matches the regular expression
	return re.Match(s)
}

// LastIndexOf returns the index of the last occurrence of the given symbol in the current symbol.
// If the symbol is not found, it returns -1 and false.
//
// This is a convenience method for working with the symbol.
func (s ByteSymbol) LastIndexOf(other Symbol) (int, bool) {
	// Find the index of the last occurrence of the other symbol
	idx := bytes.LastIndex(s, other.Bytes())
	if idx != -1 {
		// If the symbol is found, return the index and true
		return idx, true
	}
	// If the symbol is not found, return -1 and false
	return -1, false
}

// Len returns the length of the symbol in bytes.
//
// This is a convenience method for working with the symbol.
func (s ByteSymbol) Len() int {
	return len(s)
}

// Map applies the given function to each byte of the symbol and returns a new symbol.
//
// The function takes a byte and returns a new byte.
// The function is applied to each byte of the symbol.
// The resulting bytes are concatenated to form the new symbol.
func (s ByteSymbol) Map(fn functions.BiFunction[int, byte, byte]) Symbol {
	out := make(ByteSymbol, len(s))
	for i, b := range s {
		// Apply the function to the current byte and store the result in the output slice
		out[i] = fn(i, b)
	}
	// Return the new symbol, which is the concatenation of the output slice
	return out
}

func (s ByteSymbol) Remove(i, j int) Symbol {
	if i > j || i >= s.Len() {
		// If i > j, the result is an empty string
		return s.Empty()
	}
	if j > s.Len() {
		j = s.Len()
	}
	if i < 0 {
		i = 0
	}
	if j < 0 {
		j = 0
	}
	var result ByteSymbol = append([]byte{}, s[:i]...)
	result = append(result, s[j:]...)
	return result
}

// ReplaceAll replaces all occurrences of a given target symbol with a given replacement symbol
// within this symbol.
//
// It returns a new symbol, which is the result of replacing all occurrences of the target symbol with
// the replacement symbol. If the target symbol is not found in this symbol, it returns the original
// symbol.
func (s ByteSymbol) ReplaceAll(target Symbol, replacement Symbol) (Symbol, bool) {
	if target.Len() == 0 {
		// If the target or replacement are empty, return the original symbol and false
		return s, false
	}
	// Create a new bytes.Buffer to store the result
	b := bytes.Buffer{}
	// Initialize a boolean to track if any replacements were made
	replaced := false
	// Initialize an index to iterate over the bytes of the symbol
	i := 0
	// Iterate until the end of the symbol
	for i < len(s) {
		// Check if the current byte is the start of the target symbol
		if bytes.HasPrefix(s[i:], target.Bytes()) {
			// If it is, write the replacement symbol to the buffer
			b.Write(replacement.Bytes())
			// Move the index to the end of the target symbol
			i += target.Len()
			// Set the replaced flag to true
			replaced = true
		} else {
			// If it's not, write the current byte to the buffer
			b.WriteByte(s[i])
			// Move the index to the next byte
			i++
		}
	}
	// If no replacements were made, return the original symbol and false
	if !replaced {
		return s, false
	}
	// Otherwise, return the new symbol and true
	var output ByteSymbol = b.Bytes()
	return output, true
}

// ReplaceEach replaces all occurrences of a given target symbol with a given replacement symbol
// within this symbol. This is similar to ReplaceAll, but it takes a slice of replacement symbols,
// and each replacement symbol is used in order. If there are more occurrences of the target symbol
// than there are replacement symbols, the remaining occurrences aren't replaced.
//
// It returns a new symbol, which is the result of replacing all occurrences of the target symbol with
// the replacement symbols. If the target symbol is not found in this symbol, it returns the original
// symbol. If there are not enough replacement symbols, it returns the original symbol.
func (s ByteSymbol) ReplaceEach(target Symbol, replacements ...Symbol) (Symbol, bool) {
	if target.Len() == 0 {
		// If the target or replacement are empty, return the original symbol and false
		return s, false
	}
	// Create a new bytes.Buffer to store the result
	b := bytes.Buffer{}
	// Initialize an index to iterate over the bytes of the symbol
	i := 0
	// Initialize an index to iterate over the replacement symbols
	j := 0
	// Iterate until the end of the symbol
	for i < len(s) {
		// Check if the current byte is the start of the target symbol
		if bytes.HasPrefix(s[i:], target.Bytes()) {
			// If it is, write the next replacement symbol to the buffer
			if j < len(replacements) {
				b.Write(replacements[j].Bytes())
				// Move the index to the end of the target symbol
				i += target.Len()
				// Move to the next replacement symbol
				j++
			} else {
				// If there are not enough replacement symbols, return the original symbol and false
				return s, false
			}
		} else {
			// If it's not, write the current byte to the buffer
			b.WriteByte(s[i])
			// Move the index to the next byte
			i++
		}
	}
	// If all replacement symbols were used, return the new symbol and true
	if j == len(replacements) {
		var output ByteSymbol = b.Bytes()
		return output, true
	}
	// Otherwise, return the original symbol and false
	return s, false
}

// ReplaceFirst replaces the first occurrence of the target symbol with the replacement symbol.
// If the target symbol is not found in this symbol, it returns the original symbol and false.
// If the target symbol is found, it returns the new symbol and true.
func (s ByteSymbol) ReplaceFirst(target Symbol, replacement Symbol) (Symbol, bool) {
	if target.Len() == 0 {
		// If the target or replacement are empty, return the original symbol and false
		return s, false
	}
	// Create a new bytes.Buffer to store the result
	b := bytes.Buffer{}
	// Initialize an index to iterate over the bytes of the symbol
	i := 0
	// Set a flag to track if the target symbol has been replaced
	replaced := false
	// Iterate until the end of the symbol
	for i < len(s) {
		// Check if the current byte is the start of the target symbol
		if !replaced && bytes.HasPrefix(s[i:], target.Bytes()) {
			// If it is, write the replacement symbol to the buffer
			b.Write(replacement.Bytes())
			// Move the index to the end of the target symbol
			i += target.Len()
			// Set the flag to true
			replaced = true
		} else {
			// If it's not, write the current byte to the buffer
			b.WriteByte(s[i])
			// Move the index to the next byte
			i++
		}
	}
	// If the target symbol was replaced, return the new symbol and true
	// Otherwise, return the original symbol and false
	var output ByteSymbol = b.Bytes()
	return output, replaced
}

// ReplaceLast replaces the last occurrence of the target symbol with the replacement symbol.
// If the target symbol is not found in this symbol, it returns the original symbol and false.
// If the target symbol is found, it returns the new symbol and true.
func (s ByteSymbol) ReplaceLast(target Symbol, replacement Symbol) (Symbol, bool) {
	if target.Len() == 0 {
		// If the target is empty, return the original symbol and false
		return s, false
	}
	// Count the number of occurrences of the target symbol in the symbol
	occurrences := bytes.Count(s, target.Bytes())
	if occurrences == 0 {
		// If the target symbol is not found, return the original symbol and false
		return s, false
	}

	// Create a new bytes.Buffer to store the result
	b := bytes.Buffer{}
	// Initialize an index to iterate over the bytes of the symbol
	i := 0
	// Initialize a count to track the number of occurrences of the target symbol
	count := 0
	// Iterate until the end of the symbol
	for i < len(s) {
		// Check if the current byte is the start of the target symbol
		if bytes.HasPrefix(s[i:], target.Bytes()) {
			// If it is and it's not the last occurrence, write the target symbol to the buffer
			if count != occurrences-1 {
				b.Write(target.Bytes())
			} else {
				// If it is the last occurrence, write the replacement symbol to the buffer
				b.Write(replacement.Bytes())
			}
			// Move the index to the end of the target symbol
			i += target.Len()
			// Increment the count
			count++
		} else {
			// If it's not, write the current byte to the buffer
			b.WriteByte(s[i])
			// Move the index to the next byte
			i++
		}
	}
	// If the target symbol was replaced, return the new symbol and true
	// Otherwise, return the original symbol and false
	var output ByteSymbol = b.Bytes()
	return output, true
}

// ReplaceMany replaces all occurrences of the target symbols in the given symbol
// with the given replacement symbol.
//
// It returns the new symbol and a boolean indicating if any replacements were made.
//
// If no replacements were made, the original symbol is returned.
func (s ByteSymbol) ReplaceMany(replacement Symbol, targets ...Symbol) (Symbol, bool) {
	// Create a new bytes.Buffer to store the result
	b := bytes.Buffer{}
	// Initialize an index to iterate over the bytes of the symbol
	i := 0
	// Initialize a boolean to track if any replacements were made
	replaced := false
	// Iterate until the end of the symbol
	for i < len(s) {
		// Initialize a boolean to track if any replacements were made in the current iteration
		replacedInCurrent := false
		// Iterate over the targets
		for _, target := range targets {
			// Skip empty target symbols to prevent infinite loops
			if target.IsEmpty() {
				continue
			}
			// Check if the current byte is the start of the target symbol
			if bytes.HasPrefix(s[i:], target.Bytes()) {
				// If it is, write the replacement symbol to the buffer
				b.Write(replacement.Bytes())
				// Move the index to the end of the target symbol
				i += target.Len()
				// Set the boolean to true
				replaced = true
				replacedInCurrent = true
				// Break the loop since we've already replaced the current byte
				break
			}
		}
		// If no replacements were made in the current iteration, write the current byte to the buffer
		if !replacedInCurrent {
			b.WriteByte(s[i])
			// Move the index to the next byte
			i++
		}
	}
	// If no replacements were made, return the original symbol and false
	// Otherwise, return the new symbol and true
	if !replaced {
		return s, false
	}
	var output ByteSymbol = b.Bytes()
	return output, true
}

// StartsWith returns true if the given symbol starts with the given other symbol.
//
// This is a convenience method for working with the symbol.
func (s ByteSymbol) StartsWith(other Symbol) bool {
	return bytes.HasPrefix(s, other.Bytes())
}

// String returns the symbol as a string.
//
// This is a convenience method for working with the symbol.
func (s ByteSymbol) String() string {
	return string(s)
}

// Validate checks if the symbol matches the given regular expression.
// If it doesn't match, it returns an error.
func (s ByteSymbol) Validate(rgx string) error {
	// Compile the regular expression
	re, err := regexp.Compile(rgx)
	if err != nil {
		// Return an error if the regular expression is invalid
		return fmt.Errorf("invalid regular expression: %w", err)
	}

	// Check if the symbol matches the regular expression
	if !re.Match(s) {
		// Return an error if the symbol doesn't match the regular expression
		return fmt.Errorf("symbol '%s' does not match the regular expression '%s'", s, rgx)
	}

	return nil
}
