// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package tries

import (
	"testing"

	"github.com/andrerrcosta2/gtools/core/data/str/iterables"
	"github.com/andrerrcosta2/gtools/patterns/symbols"
)

// TestHandlers_SearchSymbolsByPattern_LinearOpenSymbols tests the method searchSymbolsByPattern
// using a linear dictionary of open symbols.
func TestHandlers_SearchSymbolsByPattern_LinearOpenSymbols(t *testing.T) {
	// Helper
	// types
	type V = *symbols.Logical
	type K = string

	// Create a trie from a linear dictionary
	trie, err := OfPatternsFrom(linearOpenSymbolsDictionary, false, 0, 12)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	// Search for symbols by pattern
	iterables.OfSliceMap(linearOpenSymbolsDictionary.Entries()...).
		Each(func(k K, v V) {
			HANDLERS_shouldFindNSymbols(t, trie.(*patternTrie).Root(), Entry(k, *v), 1)
		})

}

func TestHandlers_SearchSymbolsByPattern_LinearClosedSymbols(t *testing.T) {
	// Helper
	// types
	type V = *symbols.Logical
	type K = string

	// Create a trie from a linear dictionary
	trie, err := OfPatternsFrom(linearOpenSymbolsDictionary, false, 0, 12)

	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	iterables.OfSliceMap(linearOpenSymbolsDictionary.Entries()...).
		// Assert each entry can be found just once
		Each(func(k K, v V) {
			HANDLERS_shouldFindNSymbols(t, trie.(*patternTrie).Root(), Entry(k, symbols.ClosedOf(*v)), 1)
		})

}
