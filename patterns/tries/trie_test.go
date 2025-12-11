// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package tries

import (
	"github.com/andrerrcosta2/gtools/core/data/str"
	"github.com/andrerrcosta2/gtools/core/data/str/iterables"
	"github.com/andrerrcosta2/gtools/gtests/testingtools"
	"github.com/andrerrcosta2/gtools/gtests/testingtools/config/testlogs"
	"github.com/andrerrcosta2/gtools/patterns/symbols"
	"github.com/andrerrcosta2/gtools/patterns/tries/internal/def"
	"strings"
	"testing"
)

// TODO: Most of these tests must be rewriten due to the changes of requirements
// I don't even know how most of then still running...

// TestPatternTrie_BasicMethods tests the basic methods of a OfPatterns.
func TestPatternTrie_BasicMethods(t *testing.T) {
	// std
	tt := testingtools.LoggersLite(t, testlogs.OnFailure)

	// create an empty OfPatterns
	trie := OfPatterns(false, 0)

	// Check if the OfPatterns is empty
	shouldBeEmpty(tt, trie)

	// Insert some entries
	iterables.OfSlice([]str.Entry[string, symbols.Logical]{
		Entry("abcd", symbols.OpenOf(symbols.String("abcd"))),
		Entry("abce", symbols.OpenOf(symbols.String("abce"))),
		Entry("abcf", symbols.OpenOf(symbols.String("abcf"))),
	}...).
		// Insert all the entries
		Each(func(e str.Entry[string, symbols.Logical]) {
			shouldInsertEntrySuc(tt, trie, e)
		})

	// Check if the OfPatterns is not empty
	shouldNotBeEmpty(tt, trie)
	// Check if the length of the OfPatterns is 6
	shouldHaveLength(tt, trie, 6)
	// Check if the size of the OfPatterns is 3
	shouldHaveSize(tt, trie, 3)

	// Append a new different transition
	shouldInsertEntrySuc(tt, trie, Entry("xyz", symbols.LogicalOf(symbols.String("xyz"), symbols.Open)))

	// Check if the length of the OfPatterns is 9
	// Expects 1 new symbol and 3 new lengths
	shouldHaveLength(tt, trie, 9)
	// Check if the size of the OfPatterns is 4
	shouldHaveSize(tt, trie, 4)

	tt.PrintLogStack()
}

// TestBuildTrie tests the creation of a new OfPatterns from a dictionary.
// It tests if the size and length of the new OfPatterns are the same as the dictionary it was built from.
// It also tests if an invalid dictionary returns an error.
func TestBuildTrie(t *testing.T) {
	// std
	tt := testingtools.LoggersLite(t, testlogs.OnFailure)

	// Build new OfPatterns with a single goroutine
	trie, err := OfPatternsFrom(smallRootChildrenDictionary, false, 2, 10)

	if err != nil {
		tt.Errorf("error while creating OfPatterns: %v\n", err)
	}

	// Check if the size of the OfPatterns is the same as the dictionary
	shouldHaveSize(tt, trie, 3)
	// Check if the length of the OfPatterns is the same as the dictionary
	shouldHaveLength(tt, trie, 3)

	// Search for Root immediate children, shouldn't searchSymbol it against empty key
	shouldNotFindEntry(tt, trie, Entry("", symbols.ChdOnlyOf(def.Root)))

	// Check if the symbols can be found
	if children := shouldFindEntrySuc(tt, trie, Entry("a", symbols.OpenOf(symbols.String("a")))); len(children) != 1 {
		tt.Errorf("Expected 3 children, got %d", len(children))
	}

	if children := shouldFindEntrySuc(tt, trie, Entry("b", symbols.OpenOf(symbols.String("b")))); len(children) != 1 {
		tt.Errorf("Expected 4 children, got %d", len(children))
	}

	if children := shouldFindEntrySuc(tt, trie, Entry("c", symbols.OpenOf(symbols.String("c")))); len(children) != 1 {
		tt.Errorf("Expected 4 children, got %d", len(children))
	}

	tt.StackLogf("%v", trie)

	tt.PrintLogStack()
}

// TestPatternTrie_SimpleInsert tests the insertion of symbols into a newPatternTrie.
// It tests if the size and length of the newPatternTrie are the same as the dictionary it was built from.
// It also tests if an invalid dictionary returns an error.
func TestPatternTrie_SimpleInsert(t *testing.T) {
	// std
	tt := testingtools.LoggersLite(t, testlogs.OnFailure)

	trie := OfPatterns(false, 0)

	// Insert 3 symbols
	shouldInsertEntrySuc(tt, trie, Entry("abcd", symbols.OpenOf(symbols.String("abcd"))))
	shouldInsertEntrySuc(tt, trie, Entry("abce", symbols.OpenOf(symbols.String("abce"))))
	shouldInsertEntrySuc(tt, trie, Entry("abcf", symbols.OpenOf(symbols.String("abcf"))))

	// Search for the first symbol, should searchSymbol 1 child
	if c := shouldFindEntrySuc(tt, trie, Entry("abcd", symbols.OpenOf(symbols.String("abcd")))); len(c) != 1 {
		tt.Errorf("Expected 1 children, got %d", len(c))
	} else {
		// Check if the symbol found is the same as the value in the dictionary
		if !symbols.String("abcd").Equal(c[0]) {
			tt.Errorf("Expected 'abcd' to match %s", c[0])
		}
	}

	// Search for "ab" and expect 0 children
	shouldNotFindEntry(tt, trie, Entry("ab", symbols.OpenOf(symbols.String("ab"))))

	// Search for the first symbol, should searchSymbol 1 child
	if c := shouldFindEntrySuc(tt, trie, Entry("abcd", symbols.OpenOf(symbols.String("abcd")))); len(c) != 1 {
		tt.Errorf("Expected 1 children, got %d", len(c))
	} else {
		// Check if the symbol found is the same as the value in the dictionary
		if !symbols.String("abcd").Equal(c[0]) {
			tt.Errorf("Expected 'abcd' to match %s", c[0])
		}
	}

	tt.PrintLogStack()
}

// TestPatternTrie_InsertSymbolWithSymbolChild_EdgeCase01_0 tests the insertion of a symbol with a symbol child into a OfPatterns.
// It tests if the size and length of the OfPatterns are the same as the dictionary it was built from.
// It also tests if an invalid dictionary returns an error.
func TestPatternTrie_InsertSymbolWithSymbolChild_EdgeCase01_0(t *testing.T) {
	t.Skip("Skipping this test due to changing requirements")
	// std
	tt := testingtools.LoggersLite(t, testlogs.OnFailure)

	trie := OfPatterns(false, 0)

	// Insert the symbols in the correct order
	shouldInsertEntrySuc(tt, trie, Entry("abcd", symbols.OpenOf(symbols.String("abcd"))))
	shouldInsertEntrySuc(tt, trie, Entry("abce", symbols.OpenOf(symbols.String("abce"))))
	shouldInsertEntrySuc(tt, trie, Entry("abceg", symbols.OpenOf(symbols.String("abceg"))))

	// Assert OfPatterns size and length
	shouldHaveSize(tt, trie, 3)
	shouldHaveLength(tt, trie, 6)

	// Test if the symbol "abcd" has 1 child
	if c := shouldFindEntrySuc(tt, trie, Entry("abcd", symbols.OpenOf(symbols.String("abcd")))); len(c) != 1 {
		tt.Errorf("Expected 1 child, got %d", len(c))
	}

	// Test if the symbol "abce" has 2 children
	if c := shouldFindEntrySuc(tt, trie, Entry("abce", symbols.OpenOf(symbols.String("abce")))); len(c) != 2 {
		tt.Errorf("Expected 2 children, got %d", len(c))
	}

	// Print the log stack on failure
	tt.PrintLogStack()
}

// TestPatternTrie_Search Certainly this test is on infinite loop.
func TestPatternTrie_Search(t *testing.T) {
	t.Skip("Skipping this test due to changing requirements")
	// std
	tt := testingtools.LoggersLite(t, testlogs.OnFailure)

	trie := OfPatterns(false, 0)

	// Insert the symbols in the correct order
	shouldInsertEntrySuc(tt, trie, Entry("abcd", symbols.OpenOf(symbols.String("abcd"))))
	shouldInsertEntrySuc(tt, trie, Entry("abce", symbols.OpenOf(symbols.String("abce"))))
	shouldInsertEntrySuc(tt, trie, Entry("abceg", symbols.OpenOf(symbols.String("abceg"))))

	// Test if the symbol "abcd" has 1 child
	if c := shouldFindEntrySuc(tt, trie, Entry("abcd", symbols.OpenOf(symbols.String("abcd")))); len(c) != 1 {
		tt.Errorf("Expected 1 child, got %d", len(c))
	}

	// Test if the symbol "abce" has 2 children
	if c := shouldFindEntrySuc(tt, trie, Entry("abce", symbols.OpenOf(symbols.String("abce")))); len(c) != 2 {
		tt.Errorf("Expected 2 children, got %d", len(c))
	}

	// Test a simple recursive searchSymbol
	if c := shouldFindEntrySuc(tt, trie, Entry("abcdabce", symbols.OpenOf(symbols.String("abce")))); len(c) != 2 {
		tt.Errorf("Expected 2 children, got %d", len(c))
	}

	// Print the log stack on failure
	tt.PrintLogStack()
}

func TestPatternTrie_Search_PathChildren(t *testing.T) {
	t.Skip("Skipping this test due to changing requirements")
	// std
	tt := testingtools.LoggersLite(t, testlogs.OnFailure)

	trie := OfPatterns(false, 0)

	// Insert the symbols in the correct order
	shouldInsertEntrySuc(tt, trie, Entry("abcd", symbols.OpenOf(symbols.String("D"))))
	shouldInsertEntrySuc(tt, trie, Entry("abce", symbols.OpenOf(symbols.String("E"))))
	shouldInsertEntrySuc(tt, trie, Entry("abcf", symbols.OpenOf(symbols.String("F"))))
	shouldInsertEntrySuc(tt, trie, Entry("abcg", symbols.OpenOf(symbols.String("N"))))
	shouldInsertEntrySuc(tt, trie, Entry("abch", symbols.OpenOf(symbols.String("H"))))
	shouldInsertEntrySuc(tt, trie, Entry("abci", symbols.OpenOf(symbols.String("I"))))

	// Assert its size and length
	shouldHaveSizeAndLength(tt, trie, 6, 9)

	// Search the common path
	if c := shouldFindEntrySuc(tt, trie, Entry("abc", symbols.OpenOf(symbols.String("abc")))); len(c) != 6 {
		tt.Errorf("Expected 6 children, got %d", len(c))
	}

	// Print the log stack on failure
	tt.PrintLogStack()
}

// TestPatternTrie_SearchRecursiveSymbol_EdgeCase01_0 tests the searchSymbol of a OfPatterns with a recursive symbol.
//
// The test inserts a symbol with a recursive child and a symbol that is a recursive child of itself.
// It tests if the searchSymbol returns the correct children and if the recursive children are correctly set.
func TestPatternTrie_SearchRecursiveSymbol_EdgeCase01_0(t *testing.T) {
	t.Skip("Skipping this test due to changing requirements")
	// std
	tt := testingtools.LoggersLite(t, testlogs.OnFailure)

	trie := OfPatterns(false, 0)

	// Define the symbols
	closeSymbolPath := "{abc}"
	directlyRecursiveSymbolPath := "$"
	simpleSymbolPath := "xcd"

	closeSymbol := symbols.OpenOf(symbols.String("CloseSymbol"))
	directlyRecursiveSymbol := symbols.OpenOf(symbols.String("DirectlyRecursiveSymbol"))
	simpleSymbol := symbols.OpenOf(symbols.String("SimpleSymbol"))

	closeSymbolEntry := Entry(closeSymbolPath, closeSymbol)
	directlyRecursiveSymbolEntry := Entry(directlyRecursiveSymbolPath, directlyRecursiveSymbol)
	simpleSymbolEntry := Entry(simpleSymbolPath, simpleSymbol)

	// Insert the symbols in the correct order
	shouldInsertEntrySuc(tt, trie, closeSymbolEntry)
	shouldInsertEntrySuc(tt, trie, directlyRecursiveSymbolEntry)
	shouldInsertEntrySuc(tt, trie, simpleSymbolEntry)

	// recursive symbols are not found as children even if they are direct children of root
	if c := shouldFindEntrySuc(tt, trie, closeSymbolEntry); len(c) != 1 {
		tt.Errorf("Expected 1 children, got %d", len(c))
		if c[0] != closeSymbol {
			tt.Errorf("Expected %v, got %v", closeSymbol, c[0])
		}
	}

	// Search for its adjacent path
	if c := shouldFindEntrySuc(tt, trie, Entry(closeSymbolPath[:len(closeSymbolPath)-1], closeSymbol)); len(c) != 1 {
		tt.Errorf("Expected 1 children, got %d", len(c))
		if c[0] != closeSymbol {
			tt.Errorf("Expected %v, got %v", closeSymbol, c[0])
		}
	}

	// Search for DirectlyRecursiveSymbol
	if c := shouldFindEntrySuc(tt, trie, directlyRecursiveSymbolEntry); len(c) != 1 {
		tt.Errorf("Expected 1 children, got %d", len(c))
		if c[0] != directlyRecursiveSymbol {
			tt.Errorf("Expected %v, got %v", directlyRecursiveSymbol, c[0])
		}
	}

	// Search for SimpleSymbol
	if c := shouldFindEntrySuc(tt, trie, simpleSymbolEntry); len(c) != 1 {
		tt.Errorf("Expected 1 children, got %d", len(c))
		if c[0] != simpleSymbol {
			tt.Errorf("Expected SimpleSymbol, got %v", c[0])
		}
	}

	// Search for its adjacent path
	if c := shouldFindEntrySuc(tt, trie, Entry(simpleSymbolPath[:len(simpleSymbolPath)-1], simpleSymbol)); len(c) != 1 {
		tt.Errorf("Expected 1 children, got %d", len(c))
		if c[0] != simpleSymbol {
			tt.Errorf("Expected SimpleSymbol, got %v", c[0])
		}
	}

	// Search for recursive symbols
	if c := shouldFindEntrySuc(tt, trie, Entry(strings.Join([]string{directlyRecursiveSymbolPath, closeSymbolPath}, ""), closeSymbol)); len(c) != 1 {
		tt.Errorf("Expected 1 children, got %d", len(c))
		if c[0] != closeSymbol {
			tt.Errorf("Expected CloseSymbol, got %v", c[0])
		}
	}

	// Search for its adjacent path
	if c := shouldFindEntrySuc(tt, trie, Entry(strings.Join([]string{directlyRecursiveSymbolPath, closeSymbolPath[:len(closeSymbolPath)-1]}, ""), closeSymbol)); len(c) != 1 {
		tt.Errorf("Expected 1 children, got %d", len(c))
		if c[0] != closeSymbol {
			tt.Errorf("Expected CloseSymbol, got %v", c[0])
		}
	}

	// Search for multiple recursive symbols
	if c := shouldFindEntrySuc(tt, trie, Entry(strings.Join([]string{directlyRecursiveSymbolPath, closeSymbolPath, simpleSymbolPath}, ""), simpleSymbol)); len(c) != 1 {
		tt.Errorf("Expected 1 children, got %d", len(c))
		if c[0] != simpleSymbol {
			tt.Errorf("Expected CloseSymbol, got %v", c[0])
		}
	}

	// Search for multiple single recursive symbols
	if c := shouldFindEntrySuc(tt, trie, Entry(strings.Join([]string{simpleSymbolPath, simpleSymbolPath, simpleSymbolPath, simpleSymbolPath}, ""), simpleSymbol)); len(c) != 1 {
		tt.Errorf("Expected 1 children, got %d", len(c))
		if c[0] != simpleSymbol {
			tt.Errorf("Expected CloseSymbol, got %v", c[0])
		}
	}

	// Assert not found on recursive non-existent
	shouldNotFindEntry(tt, trie, Entry(strings.Join([]string{directlyRecursiveSymbolPath, "non-existent", closeSymbolPath}, ""), closeSymbol))

	// print logs on failure
	tt.PrintLogStack()
}

// TestPatternTrie_SearchRecursiveSymbol_EdgeCase01_1 tests the searchSymbol of a OfPatterns with a recursive symbol.
//
// The test inserts a sequence of symbols in the OfPatterns and searches for them.
// It tests if the searchSymbol returns the correct children.
func TestPatternTrie_SearchRecursiveSymbol_EdgeCase01_1(t *testing.T) {
	t.Skip("Skipping this test due to changing requirements")
	// std
	tt := testingtools.LoggersLite(t, testlogs.OnFailure)

	// Create a new OfPatterns
	trie, err := OfPatternsFrom(linearOpenSymbolsDictionary, false, 2, 12)

	if err != nil {
		tt.Errorf("Unexpected error: %s", err)
	}

	// Should searchSymbol a sequence of symbols recursively
	//
	// The OfPatterns should searchSymbol the sequence of symbols because the dictionary used is linear
	if c := shouldFindEntrySuc(tt, trie, Entry("abcdef", symbols.OpenOf(symbols.String("CloseSymbol")))); len(c) != 1 {
		tt.Errorf("Expected 1 children, got %d", c)
	}

	tt.PrintLogStack()
}

// TestPatternTrie_Search_AmbiguousNode_EdgeCase01_0 tests the searchSymbol of edge cases
// An edge case can happen when the OfPatterns is empty
// When the path is ambiguous between the tree internal symbols
// When there are no special tokens between placeholders.
func TestPatternTrie_Search_EmptyTrie_EdgeCase01_0(t *testing.T) {
	// std
	tt := testingtools.LoggersLite(t, testlogs.OnFailure)

	trie := OfPatterns(false, 0)

	// Searching on an empty newPatternTrie, should return nothing
	shouldNotFindEntry(tt, trie, Entry("abc", symbols.OpenOf(symbols.Empty)))

	// Searching for reserved tokens, should return nothing
	shouldNotFindEntry(tt, trie, Entry(def.Root.String(), symbols.OpenOf(symbols.Empty)))

	// Searching for a Placeholder, should return nothing
	shouldNotFindEntry(tt, trie, Entry(def.Placeholder.String(), symbols.OpenOf(symbols.Empty)))

	tt.PrintLogStack()
}

// TestPatternTrie_Search_InvalidNodes_EdgeCase01_1 tests the searchSymbol of invalid nodes
func TestPatternTrie_Search_InvalidNodes_EdgeCase01_1(t *testing.T) {
	t.Skip("Skipping this test due to changing requirements")
	// std
	tt := testingtools.LoggersLite(t, testlogs.OnFailure)
	// create an empty trie
	trie := OfPatterns(false, 0)

	// Create invalid node
	rootSymbol := def.Root.Format(symbols.String("abc%sdef"), symbols.Repl)
	placeholderSymbol := def.Placeholder.Format(symbols.String("abc%sdcf"), symbols.Repl)

	tt.StackLogf("rootSymbol: %v\nplaceholderSymbol: %v\n", rootSymbol, placeholderSymbol)

	// Assert error on insertion
	shouldNotInsertEntry(tt, trie, Entry(rootSymbol.String(), symbols.OpenOf(rootSymbol)))

	// Assert trie sizes
	shouldBeEmpty(tt, trie)

	// The ambiguous nodes should be inserted because the trie on not strict mode should keep your behaviour
	// symbol-agnostic and shouldn't reject delimiters.
	// but a searchSymbol for placeholders shouldn't return anything
	shouldInsertEntrySuc(tt, trie, Entry(placeholderSymbol.String(), symbols.OpenOf(placeholderSymbol)))

	// Assert OfPatterns sizes
	shouldNotBeEmpty(tt, trie)
	shouldHaveSize(tt, trie, 1)
	shouldHaveLength(tt, trie, 7)

	// Search for placeholderSymbol
	shouldNotFindEntry(tt, trie, Entry(placeholderSymbol.String(), symbols.OpenOf(placeholderSymbol)))

	// Search for ambiguous nodes
	// This case triggers an endless Placeholder loop.
	// Could be wise to allow some limit of chars within the same placeholders.
	shouldFindEntrySuc(tt, trie, Entry("abc[Placeholder]dcf", symbols.OpenOf(symbols.String("abc[Placeholder]dcf"))))

	tt.PrintLogStack()
}

// TestPatternTrie_Delete tests the deletion of symbols from a str.Trie.
func TestPatternTrie_Delete(t *testing.T) {
	t.Skip("Skipping this test due to changing requirements")
	// std
	tt := testingtools.LoggersLite(t, testlogs.OnFailure)

	trie := OfPatterns(false, 0)

	// Insert some symbols
	shouldInsertEntrySuc(tt, trie, Entry("abcd", symbols.OpenOf(symbols.String("abcd"))))
	shouldInsertEntrySuc(tt, trie, Entry("abce", symbols.OpenOf(symbols.String("abce"))))
	shouldInsertEntrySuc(tt, trie, Entry("abcf", symbols.OpenOf(symbols.String("abcf"))))

	// Assert the symbols were inserted
	if c := shouldFindEntrySuc(tt, trie, Entry("abc", symbols.OpenOf(symbols.String("abc")))); len(c) != 3 {
		tt.Errorf("Expected 3 children, got %d", len(c))
	}

	// Try to delete a no symbol node
	shouldNotDeleteEntry(tt, trie, Entry("abc", symbols.OpenOf(symbols.String("abc"))))

	// Check if the symbols still on the newPatternTrie
	if c := shouldFindEntrySuc(tt, trie, Entry("abc", symbols.OpenOf(symbols.String("abc")))); len(c) != 3 {
		tt.Errorf("Expected 3 children, got %d", len(c))
	}

	// Delete a symbol node
	shouldDeleteEntrySuc(tt, trie, Entry("abce", symbols.OpenOf(symbols.String("abce"))))

	// Check if the symbols no longer exist on the newPatternTrie
	shouldNotFindEntry(tt, trie, Entry("abce", symbols.OpenOf(symbols.String("abce"))))

	// Check if the symbols still exist on the newPatternTrie
	if c := shouldFindEntrySuc(tt, trie, Entry("abc", symbols.OpenOf(symbols.String("abc")))); len(c) != 2 {
		tt.Errorf("Expected 2 children, got %d", len(c))
	}

	tt.PrintLogStack()
}

// TestPatternTrie_DeleteBacktrack_EdgeCase01_0 tests the edge case of deleting a symbol with backtracking
// Tests removeSymbol a symbol on the path to another symbol and expects the length of the newPatternTrie to still the same
func TestPatternTrie_DeleteBacktrack_EdgeCase01_0(t *testing.T) {
	t.Skip("Skipping this test due to changing requirements")
	// std
	tt := testingtools.LoggersLite(t, testlogs.OnFailure)

	trie := OfPatterns(false, 0)

	// Insert some symbols
	shouldInsertEntrySuc(tt, trie, Entry("abc", symbols.OpenOf(symbols.String("abc"))))
	shouldInsertEntrySuc(tt, trie, Entry("abcd", symbols.OpenOf(symbols.String("abcd"))))

	// Assert the symbols were inserted
	if c := shouldFindEntrySuc(tt, trie, Entry("abc", symbols.OpenOf(symbols.String("abc")))); len(c) != 2 {
		tt.Errorf("Expected 2 children, got %d", len(c))
	}

	// Delete a symbol node
	shouldDeleteEntrySuc(tt, trie, Entry("abcd", symbols.OpenOf(symbols.String("abcd"))))

	// Check if the extra nodes were removed, but "abc" still exists
	if c := shouldFindEntrySuc(tt, trie, Entry("abc", symbols.OpenOf(symbols.String("abc")))); len(c) != 1 {
		tt.Errorf("Expected 1 children, got %d", len(c))
		if !c[0].Equal(symbols.String("abc")) {
			tt.Errorf("Expected symbol 'abc', got %v", c[0])
		}
	}

	tt.PrintLogStack()
}

// TestPatternTrie_DeleteBacktrack_EdgeCase01_1 tests the edge case of deleting a symbol with backtracking
// Tests removeSymbol a symbol on the path of another symbol and expects the length of the newPatternTrie to remain the same
// while its size should decrease
func TestPatternTrie_DeleteBacktrack_EdgeCase01_1(t *testing.T) {
	t.Skip("Skipping this test due to changing requirements")
	// std
	tt := testingtools.LoggersLite(t, testlogs.OnFailure)

	trie := OfPatterns(false, 0)

	shouldInsertEntrySuc(tt, trie, Entry("abcd", symbols.OpenOf(symbols.String("abcd"))))
	shouldInsertEntrySuc(tt, trie, Entry("abcdefg", symbols.OpenOf(symbols.String("abcdefg"))))

	// Assert the size and the length of the newPatternTrie
	shouldHaveSizeAndLength(tt, trie, 2, 7)

	// Assert the symbols were inserted
	shouldFindEntrySuc(tt, trie, Entry("abcd", symbols.OpenOf(symbols.String("abcd"))))
	shouldFindEntrySuc(tt, trie, Entry("abcdefg", symbols.OpenOf(symbols.String("abcdefg"))))

	// Delete the first symbol node
	shouldDeleteEntrySuc(tt, trie, Entry("abcd", symbols.OpenOf(symbols.String("abcd"))))

	// Check if the node can be found
	shouldNotFindEntry(tt, trie, Entry("abcd", symbols.OpenOf(symbols.String("abcd"))))

	// Check if the second node still reachable
	shouldFindEntrySuc(tt, trie, Entry("abcdefg", symbols.OpenOf(symbols.String("abcdefg"))))

	// Assert the size of the newPatternTrie has changed but its length has not
	shouldHaveSizeAndLength(tt, trie, 1, 7)

	tt.PrintLogStack()
}

// TestPatternTrie_DeleteBacktrack_EdgeCase01_1 tests the edge case of deleting a symbol with backtracking
// Tests removeSymbol a symbol after the path of another symbol and expects the length of the newPatternTrie to decrease the
// number of nodes its symbol should backtrack
func TestPatternTrie_DeleteBacktrack_EdgeCase01_2(t *testing.T) {
	t.Skip("Skipping this test due to changing requirements")
	// std
	tt := testingtools.LoggersLite(t, testlogs.OnFailure)

	trie := OfPatterns(false, 0)

	shouldInsertEntrySuc(tt, trie, Entry("abcd", symbols.OpenOf(symbols.String("abcd"))))
	shouldInsertEntrySuc(tt, trie, Entry("abcdefg", symbols.OpenOf(symbols.String("abcdefg"))))

	// Assert the size and the length of the newPatternTrie
	shouldHaveSizeAndLength(tt, trie, 2, 7)

	// Assert the symbols were inserted
	if c := shouldFindEntrySuc(tt, trie, Entry("abc", symbols.OpenOf(symbols.String("abc")))); len(c) != 1 {
		tt.Errorf("Expected 1 children, got %d", len(c))
		if !c[0].Equal(symbols.String("abcd")) {
			tt.Errorf("Expected symbol 'abcd', got %v", c[0])
		}
	}

	if c := shouldFindEntrySuc(tt, trie, Entry("abcd", symbols.OpenOf(symbols.String("abcd")))); len(c) != 1 {
		tt.Errorf("Expected 1 children, got %d", len(c))
		if !c[0].Equal(symbols.String("abcd")) {
			tt.Errorf("Expected symbol 'abcd', got %v", c[0])
		}
	}

	if c := shouldFindEntrySuc(tt, trie, Entry("abcdef", symbols.OpenOf(symbols.String("abcdef")))); len(c) != 1 {
		tt.Errorf("Expected 1 children, got %d", len(c))
		if !c[0].Equal(symbols.String("abcdefg")) {
			tt.Errorf("Expected symbol 'abcdefg', got %v", c[0])
		}
	}

	if c := shouldFindEntrySuc(tt, trie, Entry("abcdefg", symbols.OpenOf(symbols.String("abcdefg")))); len(c) != 1 {
		tt.Errorf("Expected 1 children, got %d", len(c))
		if !c[0].Equal(symbols.String("abcdefg")) {
			tt.Errorf("Expected symbol 'abcdefg', got %v", c[0])
		}
	}

	// Delete the second symbol node
	shouldDeleteEntrySuc(tt, trie, Entry("abcdefg", symbols.OpenOf(symbols.String("abcdefg"))))

	// Check if the node can be found
	shouldNotFindEntry(tt, trie, Entry("abcdefg", symbols.OpenOf(symbols.String("abcdefg"))))

	// Check if the first node still reachable
	shouldFindEntrySuc(tt, trie, Entry("abcd", symbols.OpenOf(symbols.String("abcd"))))

	// Assert the size and the length of the newPatternTrie has changed
	shouldHaveSizeAndLength(tt, trie, 1, 4)

	tt.PrintLogStack()
}

// TestPatternTrie_PlaceholderNode tests the creation of symbols with placeholders
// While creating symbols with placeholders
// The symbol doesn't have to be the same as its path, but it must be valid:
//
// 1. A value inside a Placeholder can't contain the same chars as its start/end markers because it'd make it impossible
// to be differentiated from a possible path
//
// 2. The characters to define a Placeholder must be the same on its path creation and its symbol creation.
// This newPatternTrie uses the value "\x01" as a Placeholder marker
func TestPatternTrie_PlaceholderNode(t *testing.T) {
	t.Skip("Skipping this test due to changing requirements")
	// std
	tt := testingtools.LoggersLite(t, testlogs.OnFailure)
	trie := OfPatterns(false, 0)

	abPhEf := def.Placeholder.Format(symbols.Of("ab]%s[ef"), symbols.Repl)
	abPh := def.Placeholder.Format(symbols.Of("ab{%s}"), symbols.Repl)
	// Multiple placeholders, just one symbol
	phCommaPh := def.Placeholder.Format(symbols.Of("[%s,%s]"), symbols.Repl)

	// Insert the symbols with placeholders
	shouldInsertEntrySuc(tt, trie, Entry(abPhEf.String(), symbols.OpenOf(abPhEf)))
	shouldInsertEntrySuc(tt, trie, Entry(abPh.String(), symbols.OpenOf(abPh)))
	shouldInsertEntrySuc(tt, trie, Entry(phCommaPh.String(), symbols.OpenOf(phCommaPh)))

	// Assert the size and the length of the newPatternTrie
	shouldHaveSizeAndLength(tt, trie, 3, 15)

	// Search for the symbols with placeholders
	if children, _ := trie.Search("ab]-should-retrieve-with-this-[ef"); len(children) != 1 {
		t.Errorf("Expected 1 child for 'ab]-should-retrieve-with-this-[ef', got %d", len(children))
	} else if !children[0].Equal(symbols.String("ab]-should-retrieve-with-this-[ef")) {
		t.Errorf("Expected 'ab]-should-retrieve-with-this-[ef' to match %s", children[0])
	}

	if children, _ := trie.Search("ab{%foo}"); len(children) != 1 {
		t.Errorf("Expected 1 child for 'ab{%%foo}', got %d", len(children))
	} else if !children[0].Equal(symbols.String("ab{%foo}")) {
		t.Errorf("Expected 'ab]-should-retrieve-with-this-' to match %s", children[0])
	}

	if children, _ := trie.Search("[%10,20]"); len(children) != 1 {
		t.Errorf("Expected 1 child for '[10,20]', got %d", len(children))
	} else if !children[0].Equal(symbols.String("[%10,20]")) {
		t.Errorf("Expected '[10,20]' to match %s", children[0])
	}

	tt.PrintLogStack()
}

// TestPatternTrie_PlaceholderNode_EdgeCase01_0 tests the creation of symbols with placeholders
// This test handles edge cases when the tries should handle errors on Placeholder creations.
func TestPatternTrie_PlaceholderNode_EdgeCase01_0(t *testing.T) {
	t.Skip("Skipping this test due to changing requirements")
	// std
	tt := testingtools.LoggersLite(t, testlogs.OnFailure)

	trie := OfPatterns(false, 0)

	// Create empty symbol
	emptySymbol := symbols.OpenOf(symbols.Empty)
	// Create invalid Placeholder symbol
	invalidPlaceholderWithNoSeparation := symbols.Of(`[\x01\x01]`)
	// Create Placeholder with ambiguous markers
	invalidPlaceholderMarkers := symbols.Of(`abcd\x01\def`)

	// Try to addSymbolToTrie an empty path with a valid symbol
	// This one should return an error because the path is empty
	shouldNotInsertEntry(tt, trie, Entry("", symbols.OpenOf(symbols.String("valid-symbol"))))

	// Try to addSymbolToTrie a non-empty path with an empty symbol
	// This one should return an error because the symbol is empty
	shouldNotInsertEntry(tt, trie, Entry("valid-path", emptySymbol))

	// Insert the symbols with placeholders
	// This one should return an error because the Placeholder has no separation
	shouldNotInsertEntry(tt, trie, Entry(invalidPlaceholderWithNoSeparation.String(), symbols.OpenOf(invalidPlaceholderWithNoSeparation)))

	// This one shouldn't return an error because a OfPatterns should be agnostic to chosen Placeholder markers,
	// But the newPatternTrie should be able to handle errors while it searches for symbols which can be placeholders or transition nodes.
	shouldInsertEntrySuc(tt, trie, Entry(invalidPlaceholderMarkers.String(), symbols.OpenOf(invalidPlaceholderMarkers)))

	tt.PrintLogStack()
}

// TestPatternTrie_PlaceholderNode_EdgeCase01_1 tests the creation of symbols with placeholders
// This test handles the case when a Placeholder is created successfully with an ambiguous marker
func TestPatternTrie_PlaceholderNode_EdgeCase01_1(t *testing.T) {
	t.Skip("Skipping this test due to changing requirements")
	// std
	tt := testingtools.LoggersLite(t, testlogs.OnFailure)

	trie := OfPatterns(false, 0)

	// Create Placeholder with ambiguous markers
	invalidPlaceholderMarkers := symbols.OpenOf(symbols.String(`abcd\x01\def`))

	// Insert the symbols with placeholders
	// This shouldn't return any error because besides its internal symbols, the newPatternTrie must be symbol-agnostic.
	// But it should be capable to notice ambiguity in its searchSymbol
	shouldInsertEntrySuc(tt, trie, Entry(invalidPlaceholderMarkers.String(), invalidPlaceholderMarkers))

	tt.PrintLogStack()
}

func TestPatternTrie_EdgeCases_Insert_EdgeDictionary(t *testing.T) {
	t.Skip("Skipping this test due to changing requirements")
	// std
	tt := testingtools.LoggableToolsLite(t, testlogs.OnErrors)
	tt.StackLog("Starting TestPatternTrie_EdgeCases_Insert_FullValidDictionary")

	// Create an empty trie
	trie := OfPatterns(false, 0)

	// Insert all the symbols of the valid dictionary concurrently
	iterables.OfSlice(edgeValidDictionary.Entries()...).
		Each(func(e E) {
			for _, v := range e.Value() {
				shouldInsertEntrySuc(tt, trie, Entry(e.Key(), v))
				tt.RegisterCalls(1, "symbol-insertion")
			}
		})

	tt.StackLogf("Trie:\n%v", trie)

	tt.AssertCalls(edgeValidDictionary.Size(), "symbol-insertion")
	tt.Condition(trie.Size() == edgeValidDictionary.Size(), "expected trie size to be %d but was %d", edgeValidDictionary.Size(), trie.Size())
	tt.Condition(trie.Length() == 262, "expected trie length to be %d but was %d", 262, trie.Length())

	tt.PrintLogStack()
}

// TestPatternTrie_EdgeCases_Delete_EdgeDictionary tests the deletion of symbols on a trie.
// It creates a OfPatterns from a dictionary with all the edge cases a valid dictionary can have,
// and then deletes part of the symbols randomly.
// It then checks if an error occurred and its expected size.
// It also checks if the searchSymbol results are the same as the dictionary.
func TestPatternTrie_EdgeCases_Delete_EdgeDictionary(t *testing.T) {
	t.Skip("Skipping this test due to changing requirements")
	// std
	tt := testingtools.LoggableToolsLite(t, testlogs.OnFailure)

	tt.StackLog("Starting TestPatternTrie_EdgeCases_Delete_FullValidDictionary")

	// Create by an edge dictionary
	trie, err := OfPatternsFrom(edgeValidDictionary, false, 12, 20)

	if err != nil {
		tt.Errorf("unexpected error while creating newPatternTrie: %v\n", err)
	}

	tt.StackLogf("Trie before deletion: { size: %d, length: %d }", trie.Size(), trie.Length())

	// Delete a few random symbols concurrently
	iterables.OfSlice(edgeValidDictionary.Entries()...).
		Some(40).
		Each(func(e E) {
			for _, v := range e.Value() {
				shouldDeleteEntrySuc(tt, trie, Entry(e.Key(), v))
			}
		})

	tt.Condition(trie.Size() == edgeValidDictionary.Size()-40, "expected trie size to be %d but was %d", edgeValidDictionary.Size()-40, trie.Size())

	//TODO: create an utility to determine the length of a trie based on its dictionary
	tt.PrintLogStack()
}
