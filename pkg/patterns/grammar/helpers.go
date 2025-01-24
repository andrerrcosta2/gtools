// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package grammar

import (
	"github.com/andrerrcosta2/gtools/core/gtools/constraints/prim/bins"
	"github.com/andrerrcosta2/gtools/patterns/symbols"
)

// GetSymbolValue takes a symbol and attempts to convert it to a value of type K.
//
// Returns:
// - The value of type K if the conversion was successful.
// - An error if the conversion failed.
func GetSymbolValue[T any](symbol Symbol) (T, error) {
	// Use the bins.FromBytes function to convert the symbol to a value of type K.
	return bins.FromBytes[T](symbol.Bytes())
}

// HasConsecutiveSymbol checks if a symbol has a consecutive occurrence in a pattern.
//
// If the symbol has a consecutive occurrence in the pattern, it returns true.
// If the symbol doesn't have a consecutive occurrence in the pattern, it returns false.
func HasConsecutiveSymbol(pattern, symbol Symbol) bool {
	// Iterate over the pattern
	for i := 0; i < pattern.Len(); i++ {
		// Check if the current index is within the bounds of the pattern
		if i+1 < pattern.Len() {
			// ExtractSymbol the pattern at the current index using the symbol
			// If the symbol is present at the current index, split the rest of the pattern
			if sym, rest := symbol.ExtractSymbol(pattern.After(i)); !sym.IsEmpty() {
				// ExtractSymbol the rest of the pattern again using the symbol
				// If the symbol is present again, the symbol has a consecutive occurrence
				if sym, rest = symbol.ExtractSymbol(rest); !sym.IsEmpty() {
					return true
				}
			}
		}
	}
	// If the symbol doesn't have a consecutive occurrence in the pattern, return false
	return false
}

// HasSymbolAt checks if a symbol is present at a given index in a pattern.
//
// If the symbol is present at the given index, it returns true.
// If the symbol isn't present at the given index, it returns false.
func HasSymbolAt(pattern, symbol Symbol, idx int) bool {
	// Check if the index is within the bounds of the pattern
	if idx >= 0 && idx < pattern.Len() {
		// ExtractSymbol the pattern at the given index using the symbol
		// If the split symbol is not empty, the symbol is present at this index
		if pattern.After(idx).StartsWith(symbol) {
			return true
		}
	}
	return false
}

// HasSymbolAtAny checks if a symbol is present at any of the given indices in a pattern.
//
// If the symbol is present at any of the given indices, it returns true.
// If the symbol is not present at any of the given indices, it returns false.
func HasSymbolAtAny(pattern, symbol Symbol, idx ...int) bool {
	// Iterate over the given indices
	for _, i := range idx {
		// Check if the index is within the bounds of the pattern
		if i < pattern.Len() && i >= 0 {
			// ExtractSymbol the pattern at the given index using the symbol
			// If the split symbol is not empty, the symbol is present at this index
			if pattern.After(i).StartsWith(symbol) {
				return true
			}
		}
	}
	return false
}

// HasSymbolAtAll checks if a symbol is present at all the given indices in a pattern.
//
// If the symbol is present at any of the given indices, it returns true.
// If the symbol isn't present at any of the given indices, it returns false.
func HasSymbolAtAll(pattern, symbol Symbol, idx ...int) bool {
	// Iterate over the given indices
	for _, i := range idx {
		// Check if the index is within the bounds of the pattern
		if i < pattern.Len() && i >= 0 {
			// Check if the pattern starts with the symbol
			if !pattern.After(i).StartsWith(symbol) {
				return false
			}
		} else {
			return false
		}
	}
	// If the symbol is present at any of the given indices, return true
	return true
}

func mapToSymbol[V any, T symbols.Logical](s ...V) (T, error) {
	bs, err := NewSymbol(s)
	return bs.(T), err
}
