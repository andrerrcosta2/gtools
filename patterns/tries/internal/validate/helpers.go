// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package validate

import (
	"github.com/andrerrcosta2/gtools/patterns/grammar"
	"github.com/andrerrcosta2/gtools/patterns/symbols"
	"github.com/andrerrcosta2/gtools/patterns/tries/errs"
	"github.com/andrerrcosta2/gtools/patterns/tries/internal/def"
)

// Symbol verifies if a symbol is valid to be used in a pattern
func Symbol(s grammar.Symbol) error {
	if s.IsEmpty() {
		return errs.EmptySymbol
	}

	for i := 0; i < s.Len(); i++ {
		// Reject patterns with consecutive placeholders
		if grammar.HasSymbolAtAll(s, def.Placeholder, i, i+def.Placeholder.Len()) {
			return errs.ConsecutivePlaceholders(s)
		}
		// Reject patterns containing the Root symbol
		if grammar.HasSymbolAt(s, def.Root, i) {
			return errs.RootSymbolOnPath(s)
		}
	}

	return nil
}

// Symbols verifies if all symbols are valid to be used in a pattern
func Symbols(s ...grammar.Symbol) error {
	// Iterate over all symbols and check if they're valid
	for _, symbol := range s {
		if err := Symbol(symbol); err != nil {
			return err
		}
	}
	return nil
}

// PatternAgainstSymbol verifies if a pattern is valid against a given symbol.
// If the pattern contains the Root symbol, it returns an error.
// If the pattern contains the Placeholder, it must be found in the symbol the same number of times.
// If the Placeholder exists in both, it must be separated by at least one character in both.
func PatternAgainstSymbol(p string, symbol grammar.Symbol) error {
	pattern := symbols.Of(p)

	// Reject patterns containing reserved symbols
	if _, exists := pattern.IndexOf(def.Root); exists {
		return errs.RootSymbolOnPath(symbol)
	}

	// The pattern must be non-empty
	for !pattern.IsEmpty() && !symbol.IsEmpty() {
		// Check if there is a Placeholder in the symbol and in the pattern
		if xs, xp, contains, err := xpsph(symbol, pattern); err == nil {
			if !contains {
				return nil // There's no Placeholder in the symbol.
			}

			// If there is a Placeholder, update the symbol and pattern
			symbol, pattern = xs, xp
		} else {
			return err
		}
	}
	return nil
}

// xpsph asserts if the Placeholder exists in the symbol it also exists in the pattern
// It returns a new value for the symbol and pattern, a boolean indicating if the Placeholder was found, and an error.
// It also checks if a Placeholder doesn't appear consecutively in the symbol and in the pattern.
// If the Placeholder exists in the symbol, it must be found in the pattern the same number of times.
// If the Placeholder is found in both, it must be separated by at least one character in both.
func xpsph(symbol, pattern grammar.Symbol) (xSymbol, xPattern grammar.Symbol, contains bool, err error) {
	// The Placeholder exists in the symbol
	if phValue, rest := symbol.ExtractSymbol(def.Placeholder); !phValue.IsEmpty() {
		contains = true
		// Found the Placeholder in the symbol
		xSymbol = rest
		// Extract the Placeholder from the rest of the symbol
		phValue, rest = rest.ExtractSymbol(def.Placeholder)
		// Reject consecutive placeholders
		if rest.StartsWith(def.Placeholder) {
			return symbols.Empty, symbols.Empty, true, errs.ConsecutivePlaceholders(symbol)
		}

		// If the Placeholder exists on the symbol, it must exist on the pattern
		if phValue, rest = pattern.ExtractSymbol(def.Placeholder); !phValue.IsEmpty() {
			// Reject consecutive placeholders
			if rest.StartsWith(def.Placeholder) {
				return symbols.Empty, symbols.Empty, true, errs.ConsecutivePlaceholders(symbol)
			}

			// Found the Placeholder in the pattern
			xPattern = rest
			return
		}
		// The Placeholder exists on the symbol but wasn't found in the pattern
		err = errs.MismatchedPlaceholders(pattern, symbol)
		// The Placeholder was found on the symbol, but not in the pattern
		return
	}

	// The Placeholder wasn't found in the symbol
	if _, exists := pattern.IndexOf(def.Placeholder); exists {
		// But was found in the pattern
		return symbols.Empty, symbols.Empty, true, errs.MismatchedPlaceholders(pattern, symbol)
	}
	// There are no placeholders in the symbol
	return
}
