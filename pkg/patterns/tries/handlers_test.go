// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package tries

import (
	"github.com/andrerrcosta2/gtools/core/data/str"
	"github.com/andrerrcosta2/gtools/core/data/str/iterables"
	"github.com/andrerrcosta2/gtools/gtests"
	"github.com/andrerrcosta2/gtools/gtests/testingtools"
	"github.com/andrerrcosta2/gtools/patterns/symbols"
	"testing"
)

// TestHandlers_SearchSymbolsByPattern_LinearOpenSymbols tests the method searchSymbolsByPattern
// using a linear dictionary of open symbols.
func TestHandlers_SearchSymbolsByPattern_LinearOpenSymbols(t *testing.T) {
	// Helper
	tt := testingtools.LoggersLite(t, gtests.LogOnFailure)
	// types
	type S = symbols.Logical
	type K = string
	type E = str.Entry[K, S]

	// Create a trie from a linear dictionary
	trie, err := OfPatternsFrom(linearOpenSymbolsDictionary, false, 0, 12)
	if err != nil {
		tt.Errorf("Unexpected error: %s", err)
	}

	// Search for symbols by pattern
	iterables.OfSlice(linearOpenSymbolsDictionary.Entries()...).
		Each(func(e E) { HANDLERS_shouldFindNSymbols(tt, trie.Root(), e, 1) })

	tt.PrintLogStack()
}

func TestHandlers_SearchSymbolsByPattern_LinearClosedSymbols(t *testing.T) {
	// Helper
	tt := testingtools.LoggersLite(t, gtests.LogOnFailure)
	// types
	type S = symbols.Logical
	type K = string
	type E = str.Entry[K, S]

	// Create a trie from a linear dictionary
	trie, err := OfPatternsFrom(linearOpenSymbolsDictionary, false, 0, 12)

	if err != nil {
		tt.Errorf("Unexpected error: %s", err)
	}

	iterables.OfSlice(linearOpenSymbolsDictionary.Entries()...).
		Map(func(e E) E {
			return Entry(e.Key(), symbols.ClosedOf(e.Value()))
		}).
		// Assert each entry can be found just once
		Each(func(e E) {
			HANDLERS_shouldFindNSymbols(tt, trie.Root(), e, 1)
		})

	tt.PrintLogStack()
}
