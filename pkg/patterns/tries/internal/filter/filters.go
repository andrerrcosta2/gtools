// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package filter

import (
	"github.com/andrerrcosta2/gtools/patterns/grammar"
	"github.com/andrerrcosta2/gtools/patterns/symbols"
	"github.com/andrerrcosta2/gtools/patterns/tries/errs"
	"github.com/andrerrcosta2/gtools/patterns/tries/internal/def"
	"github.com/andrerrcosta2/gtools/patterns/tries/internal/validate"
)

// Insertion verifies if a pattern is valid to be inserted before or after a symbol in a trie.
// It returns an error if the pattern is invalid.
func Insertion(pattern string, symbol grammar.Symbol) error {
	if len(pattern) == 0 {
		return errs.EmptyEntry
	}
	// Reject patterns containing the Root symbol
	if symbols.String(pattern).Contains(def.Root) {
		return errs.RootSymbolOnPath(symbol)
	}

	// Reject patterns equals Placeholder symbol
	if symbols.String(pattern).Equal(def.Placeholder) {
		return errs.PlaceholderSymbol
	}

	return validate.PatternAgainstSymbol(pattern, symbol)
}
