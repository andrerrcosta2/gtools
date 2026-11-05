// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package tries

import (
	"strings"
	"testing"

	"github.com/andrerrcosta2/gtools/core/data/str/dict"
	"github.com/andrerrcosta2/gtools/core/data/str/iterables"
	"github.com/andrerrcosta2/gtools/core/testlite"
	"github.com/andrerrcosta2/gtools/core/testlite/assertlite"
	"github.com/andrerrcosta2/gtools/patterns/grammar"
	"github.com/andrerrcosta2/gtools/patterns/symbols"
	"github.com/andrerrcosta2/gtools/patterns/tries/internal/def"
)

// TODO: Most of these tests must be rewriten due to the changes of requirements
// I don't even know how most of then still running...

type E = dict.Entry[K, []S]

// TestPatternTrie_Build tests the creation of a OfPatterns
// It tests if the size and length of the newPatternTrie are the same
// as the dictionary it was built from.
// It also tests if an invalid dictionary returns an error.
func TestPatternTrie_Build(t *testing.T) {

	// building a valid newPatternTrie
	trie, err := OfPatternsFrom(linearOpenSymbolsDictionary, false, 6, 12)

	if err != nil {
		t.Fatalf("error while creating newPatternTrie: %v\n", err)
	}

	if trie.Size() != linearOpenSymbolsDictionary.Size() {
		t.Fatalf("newPatternTrie size expected to be %v but was %v\n", linearOpenSymbolsDictionary.Size(), trie.Size())
	}

	if trie.Length() != 30 {
		t.Errorf("newPatternTrie length expected to be 26 but was %v", trie.Length())
	}

	// building an invalid newPatternTrie
	invalidTrie, err := OfPatternsFrom(invalidDictionary, false, 3, 6)

	t.Logf("invalidTrie: %v", invalidTrie)
	t.Logf("err: %v", err)

	if err == nil {
		t.Errorf("error expected but none received\n")
	}

}

// TestPatternTrie_InsertConcurrently tests the insertion of symbols into a newPatternTrie using concurrency.
// It tests if the size and length of the str.Trie are the same as the dictionary it was built from.
// It also tests if an invalid dictionary returns an error.
func TestPatternTrie_InsertConcurrently(t *testing.T) {
	// create an empty trie
	trie := OfPatterns(false, 0)

	type K = string
	type V = []symbols.Logical

	iterables.OfSliceMap(linearOpenSymbolsDictionary.Entries()...).
		Parallel(func(k K, ss V) {
			for _, s := range ss {
				if err := trie.Insert(k, s); err != nil {
					t.Errorf("error while inserting value: %v\n", err)
				}
			}

		}, 5)

	// check if the size and length of the newPatternTrie are the same as the dictionary
	if trie.Size() != linearOpenSymbolsDictionary.Size() {
		t.Errorf("newPatternTrie size expected to be %v but was %v\n", linearOpenSymbolsDictionary.Size(), trie.Size())
	}

	if trie.Length() != 30 {
		t.Errorf("newPatternTrie length expected to be 26 but was %v", trie.Length())
	}
}

// TestPatternTrie_SearchConcurrently_BasicSearch tests the searchSymbol method of a str.Trie
// using concurrency. It creates a new OfPatterns from a linear dictionary and then
// searches for each entry in the dictionary concurrently. It checks if an error
// occurred and if the size of the newPatternTrie is the same as the dictionary.
// It also checks if the searchSymbol results are the same as the dictionary.
func TestPatternTrie_SearchConcurrently_BasicSearch(t *testing.T) {
	t.Skip("Skipping this test due to changing requirements")
	trie, err := OfPatternsFrom(linearOpenSymbolsDictionary, false, 2, 10)

	if err != nil {
		t.Errorf("error while creating newPatternTrie: %v\n", err)
	}

	if trie.Size() != linearOpenSymbolsDictionary.Size() {
		t.Errorf("newPatternTrie size expected to be %v but was %v\n", linearOpenSymbolsDictionary.Size(), trie.Size())
	}

	iterables.OfSlice(linearOpenSymbolsDictionary.Entries()...).
		Parallel(func(i int, e E) {
			if sym, exists := trie.Search(e.Key()); !exists {
				t.Errorf("expected to searchSymbol %v but was not found\n", e.Key())
				return
			} else {
				// Check if there is only one value found
				if len(sym) != 1 {
					t.Errorf("expected just one value for %v but found %v\n", e.Key(), sym)
					return
				}
				// Check if the value found is the same as the value in the dictionary
				if !e.Value()[0].Equal(sym[0]) {
					t.Errorf("expected %v but found %v\n", e.Value(), sym[0])
				}
			}
		}, 5)
}

// TestPatternTrie_SearchConcurrently_PlaceholderSymbols tests the searchSymbol method of a str.Trie
// using concurrency. It creates a new OfPatterns from a dictionary with Placeholder symbols
// and then searches for each entry in the dictionary concurrently, replacing the Placeholder
// with a value. It checks if an error occurred and if the size of the new OfPatterns is the same as the dictionary.
// It also checks if the searchSymbol results are the same as the dictionary.
func TestPatternTrie_SearchConcurrently_PlaceholderSymbols(t *testing.T) {
	t.Skip("Skipping this test due to changing requirements")
	// logger tools

	trie, err := OfPatternsFrom(dictionaryWithPlaceholders, false, 12, 12)

	if err != nil {
		t.Errorf("error while creating newPatternTrie: %v\n", err)
	}

	iterables.
		OfSlice(dictionaryWithPlaceholders.Entries()...).
		// Remove symbols without placeholder
		Filter(func(e E) bool {
			return strings.Contains(e.Key(), def.Placeholder.String())
		}).
		// Replace Placeholder with a value
		// As a logical behavior of a "placeholder-able" trie, we should avoid using
		// as placeholder values any character used as a placeholder delimiter.
		// In the case of the used dictionary we have a rune "m" being used as delimiter
		Map(func(e E) (out E) {
			key, ok := symbols.Of(e.Key()).ReplaceAll(def.Placeholder, symbols.Of("<-soMe-value->"))
			if !ok {
				t.Errorf("unexpected error while replacing placeholder:\n")
			}
			return grammar.SymbolEntry[K, []S, S](key.String(), e.Value()...)
		}).
		// Search concurrently. This is in fact searching in batches because I had to change the
		// dictionary from a linear map to a map of slices
		Parallel(func(i int, e E) {
			for _, v := range e.Value() {
				shouldFindEntrySuc(t, trie, Entry(e.Key(), v))
			}
		}, 5)

}

// TestPatternTrie_DeleteConcurrently tests the deletion of symbols on a trie concurrently.
func TestPatternTrie_DeleteConcurrently(t *testing.T) {
	counter := testlite.Map(t)

	// Create by a linear dictionary of 10 symbols
	trie, err := OfPatternsFrom(linearOpenSymbolsDictionary, false, 2, 10)

	if err != nil {
		t.Errorf("error while creating newPatternTrie: %v\n", err)
	}

	startLength := trie.Length()

	// Print the initial length of the trie
	t.Logf("trie length: %v\n", trie.Length())

	t.Logf("trie: %v\n", trie)

	// Concurrently delete half of the symbols
	iterables.OfSlice(linearOpenSymbolsDictionary.Entries()...).Some(linearOpenSymbolsDictionary.Size()/2).
		// Delete concurrently. This is in fact deleting in batches because I had to change the
		// dictionary from a linear map to a map of slices
		Parallel(func(i int, e E) {
			for _, v := range e.Value() {
				shouldDeleteEntrySuc(t, trie, Entry(e.Key(), v))
				counter.Add(1, "deleted-symbols")
			}
			// that only works because this is a linear dictionary
			counter.Add(len(e.Key()), "deleted-node")
		}, 10)

	// Print the length of the trie after deletion
	t.Logf("trie length after deletion: %v\n", trie.Length())

	// Check if the size of the trie is correct
	deletedSymbols := counter.Get("deleted-symbols")
	assertlite.True(t, trie.Size() == linearOpenSymbolsDictionary.Size()-deletedSymbols, "expected to removeSymbol half of the trie symbols '%v' "+
		"but got after size of '%v'\n", linearOpenSymbolsDictionary.Size()-deletedSymbols, trie.Size())

	// Check if the length of the trie is correct
	deletedNodes := counter.Get("deleted-node")
	assertlite.True(t, trie.Length() == startLength-deletedNodes, "expected to removeSymbol half of the trie node '%v' "+
		"but got after length of '%v'\n", startLength-deletedNodes, trie.Length())

	t.Logf("trie: %v\n", trie)

}

func TestPatternTrie_EdgeCases_Build_EdgeValidDictionary(t *testing.T) {
	t.Skip("Skipping this test due to changing requirements")

	// build a valid trie
	trie, err := OfPatternsFrom(edgeValidDictionary, false, 5, 10)

	if err != nil {
		t.Errorf("error while creating newPatternTrie: %v\n", err)
	}

	if trie.Size() != edgeValidDictionary.Size() {
		t.Errorf("newPatternTrie size expected to be %v but was %v\n", edgeValidDictionary.Size(), trie.Size())
	}

	if trie.Length() != 262 {
		t.Errorf("newPatternTrie length expected to be %v but was %v", 262, trie.Length())
	}

}

// TestPatternTrie_EdgeCases_InsertConcurrently_EdgeValidDictionary tests the insertion of symbols on a trie concurrently.
// It creates a OfPatterns from a dictionary with all the edge cases a valid dictionary can have,
// and then inserts all the symbols concurrently.
func TestPatternTrie_EdgeCases_InsertConcurrently_EdgeValidDictionary(t *testing.T) {
	counter := testlite.Map(t)
	t.Skip("Skipping this test due to changing requirements")

	t.Log("Starting TestPatternTrie_EdgeCases_Insert_FullValidDictionary")

	// Create an empty trie
	trie := OfPatterns(false, 0)

	// Insert all concurrently
	iterables.OfSlice(edgeValidDictionary.Entries()...).
		Parallel(func(i int, e E) {
			for _, v := range e.Value() {
				shouldInsertEntrySuc(t, trie, Entry(e.Key(), v))
				// Register how many symbols were inserted
				counter.Add(1, "symbol-insertion")
			}
		}, 10)

	t.Logf("Trie:\n%v", trie)

	// Verify if all symbols were inserted
	counter.AssertCounts(edgeValidDictionary.Size(), "symbol-insertion")
	// Verify that the size of the trie is the same as the dictionary
	assertlite.True(t, trie.Size() == edgeValidDictionary.Size(), "expected trie size to be %d but was %d", edgeValidDictionary.Size(), trie.Size())
	// Verify that the length of the trie is the same as the dictionary
	assertlite.True(t, trie.Length() == 262, "expected trie length to be %d but was %d", 262, trie.Length())

}

// TestPatternTrie_EdgeCases_DeleteConcurrently_EdgeValidDictionary tests the deletion of symbols on a trie concurrently.
// It creates a newPatternTrie from a dictionary with all the edge cases a valid dictionary can have,
// and then deletes half of the symbols concurrently.
func TestPatternTrie_EdgeCases_DeleteConcurrently_EdgeValidDictionary(t *testing.T) {
	t.Skip("Skipping this test due to changing requirements")

	t.Log("Starting TestPatternTrie_EdgeCases_Delete_FullValidDictionary")

	// Create by an edge dictionary
	trie, err := OfPatternsFrom(edgeValidDictionary, false, 3, 9)

	if err != nil {
		t.Errorf("unexpected error while creating newPatternTrie: %v\n", err)
	}

	// Delete a few random symbols concurrently
	iterables.OfSlice(edgeValidDictionary.Entries()...).Some(40).
		Parallel(func(i int, e E) {
			// this trie is too irregular to predict the deleted node, this is in fact why we use tries
			// instead of dicts for these task
			for _, v := range e.Value() {
				shouldDeleteEntrySuc(t, trie, Entry(e.Key(), v))
			}
		}, 5)

	assertlite.True(t, trie.Size() == edgeValidDictionary.Size()-40, "expected trie size to be %d but was %d", edgeValidDictionary.Size()-40, trie.Size())

	//TODO: create an utility to determine the length of a trie based on its dictionary

}
