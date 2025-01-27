// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package regex

import (
	"github.com/andrerrcosta2/gtools/gtests"
	"github.com/andrerrcosta2/gtools/gtests/testingtools"
	"github.com/andrerrcosta2/gtools/patterns/symbols"
	"github.com/andrerrcosta2/gtools/patterns/tries"
	"testing"
)

// TODO: The trie requirements has changed, so most of these tests must be rewritten
// TestBuildTrie tests the creation of a new RegexTrie from a dictionary.
// It tests if the size and length of the new RegexTrie are the same as the dictionary it was built from.
// It also tests if an invalid dictionary returns an error.
func TestBuildTrie(t *testing.T) {
	t.Skip("Skipping this test due to changing requirements")
	tt := testingtools.LoggersLite(t, gtests.LogOnFailure)

	trie, err := Trie(false, 6, 6)

	if err != nil {
		tt.Errorf("failed to build trie: %v\n", err)
	}

	// Check if the RegexTrie is not empty
	shouldNotBeEmpty(tt, trie)

	// Check if the size and length of the RegexTrie are the same as the dictionary
	shouldHaveSizeAndLength(tt, trie, Dictionary().Size(), 103)

	tt.PrintLogStack()
}

// TODO: rewrite
func TestSearch_AllSymbols(t *testing.T) {
	t.Skip("Skipping this test due to changing requirements")
	tt := testingtools.LoggersLite(t, gtests.LogOnFailure)

	// Get the dictionary instance
	dictionary := Dictionary()

	trie, err := Trie(false, 12, 12)

	if err != nil {
		tt.Errorf("failed to build trie: %v\n", err)
	}

	for key, values := range dictionary.EntrySet() {
		for i := 0; i < len(values); i++ {
			shouldFindEntrySuc(tt, trie, tries.Entry(key, values[i]))
		}
	}

	tt.StackLogf("trie: %v\n", trie)
	tt.PrintLogStack()
}

// TODO: rewrite
func TestSearch_SimpleRegexes(t *testing.T) {
	t.Skip("Skipping this test due to changing requirements")

	tt := testingtools.LoggersLite(t, gtests.LogOnFailure)

	trie, err := Trie(false, 0, 0)

	if err != nil {
		tt.Errorf("failed to build trie: %v\n", err)
	}

	simpleRegexes.Each(func(key string, value string) {
		shouldFindEntrySuc(tt, trie, tries.Entry(value, symbols.OpenOf(symbols.Empty)))
	})

	tt.PrintLogStack()
}
