// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package regex

import (
	"github.com/andrerrcosta2/gtools/core/data/str"
	"github.com/andrerrcosta2/gtools/gtests"
	"github.com/andrerrcosta2/gtools/patterns/grammar"
	"github.com/andrerrcosta2/gtools/patterns/symbols"
	"github.com/andrerrcosta2/gtools/patterns/tries"
)

// shouldInsertEntrySuc is a std function that tests the insertion of a single entry into a str.Trie.
// It logs the entry and the trie states before and after the insertion and checks if an error occurred.
func shouldInsertEntrySuc(t gtests.Loggable, trie tries.Pattern, e str.Entry[string, symbols.Logical]) {
	t.Helper()
	// Log the entry
	t.StackLogf("Entry: %v, length: %d", e, len(e.Key()))

	// Insert the entry into the trie
	err := trie.Insert(e.Key(), e.Value())
	if err != nil {
		// Log the error if it occurred
		t.Errorf("error while inserting '%v': %v\n", e.Key(), err)
	}

	t.StackLogf("Entry inserted '%v'", e)

	// Log the trie states for debugging
	t.StackLogf("trie: { size: %v, length: %v }\n", trie.Size(), trie.Length())
}

// shouldNotInsertEntry is a std function that tests the insertion of a single entry into a str.Trie.
// It logs the entry and the trie states before and after the insertion and checks if an error occurred.
func shouldNotInsertEntry(t gtests.Loggable, trie tries.Pattern, e str.Entry[string, symbols.Logical]) {
	t.Helper()
	// Log the entry
	t.StackLogf("Entry: %v, length: %d", e, len(e.Key()))

	// Insert the entry into the trie
	err := trie.Insert(e.Key(), e.Value())
	if err == nil {
		// Log the error if it occurred
		t.Errorf("expected an error but none received\n")
	}

	// Log the trie states for debugging
	t.StackLogf("trie: { size: %v, length: %v }\n", trie.Size(), trie.Length())
}

// shouldDeleteEntrySuc is a std function that tests the deletion of a single entry from a str.Trie.
// It logs the entry and the trie states before and after the deletion and checks if an error occurred.
func shouldDeleteEntrySuc(t gtests.Loggable, trie tries.Pattern, e str.Entry[string, symbols.Logical]) {
	t.Helper()
	// Log the entry
	t.StackLogf("Deleting entry '%v'...", e)

	// Delete the entry from the trie
	if !trie.Delete(e.Key(), e.Value()) {
		// Mark as failed if was not deleted
		t.Errorf("expected to delete '%v' but was not found\n", e.Key())
		// Log the entry for debugging
		t.StackLogf("[Not deleted]: %v", e)
	} else {
		// Log the entry for debugging
		t.StackLogf("Deleted: %v", e)
	}

	// Log the trie states for debugging
	t.StackLogf("trie: { size: %v, length: %v }\n", trie.Size(), trie.Length())
}

// shouldNotDeleteEntry is a std function that tests the deletion of a single entry from a str.Trie.
// It logs the entry and the trie states before and after the deletion and checks if an error occurred.
func shouldNotDeleteEntry(t gtests.Loggable, trie tries.Pattern, e str.Entry[string, symbols.Logical]) {
	t.Helper()
	// Log the entry
	t.StackLogf("Deleting entry '%v'...", e)

	// Delete the entry from the trie
	if trie.Delete(e.Key(), e.Value()) {
		// Mark as failed if was not deleted
		t.Errorf("expected to not delete '%v' but was found\n", e.Key())
		// Log the entry for debugging
		t.StackLogf("[Deleted]: %v", e)
	} else {
		// Log the entry for debugging
		t.StackLogf("Not Deleted: %v", e)
	}

	// Log the trie states for debugging
	t.StackLogf("trie: { size: %v, length: %v }\n", trie.Size(), trie.Length())
}

// shouldFindEntrySuc is a std function that tests the search of a single entry from a str.Trie.
// It logs the entry and the trie states before and after the search and checks if an error occurred.
func shouldFindEntrySuc(t gtests.Loggable, trie tries.Pattern, e str.Entry[string, symbols.Logical]) []grammar.Symbol {
	t.Helper()
	// Log the entry
	t.StackLogf("Entry: %v", e)

	// Search for the entry in the PatternTrie
	sym, exists := trie.Search(e.Key())

	// Check if the entry was found
	if !exists {
		t.Errorf("expected to find '%v' but was not found\n", e.Key())
	}

	// Log the result for debugging
	t.StackLogf("search: %v, result: %v\n", e.Key(), sym)

	return sym
}

// shouldNotFindEntry is a std function that tests the search of a single entry from a str.Trie.
// It logs the entry and the trie states before and after the search and checks if an error occurred.
func shouldNotFindEntry(t gtests.Loggable, trie tries.Pattern, e str.Entry[string, symbols.Logical]) {
	t.Helper()
	// Log the entry
	t.StackLogf("Entry: %v", e)

	// Search for the entry in the PatternTrie
	if sym, exists := trie.Search(e.Key()); exists {
		// Check if the entry was found
		t.Errorf("expected to not find '%v' but was found\n", e.Key())
		// Log the result for debugging
		t.StackLogf("search: %v, result: %v\n", e.Key(), sym)
	}
}

// shouldBeEmpty is a std function that tests if a str.Trie is empty.
// It logs the trie states before and after the search and checks if an error occurred.
func shouldBeEmpty(t gtests.Loggable, trie tries.Pattern) {
	t.Helper()
	// Log the trie states for debugging
	t.StackLogf("trie: { size: %v, length: %v }", trie.Size(), trie.Length())

	// Check if the PatternTrie is empty
	if !trie.IsEmpty() {
		t.Errorf("expected to be empty but was not\n")
	}

	// Assert length
	if trie.Length() != 0 {
		t.Errorf("expected length to be 0 but was %v\n", trie.Length())
	}

	// Assert size
	if trie.Size() != 0 {
		t.Errorf("expected size to be 0 but was %v\n", trie.Size())
	}
}

// shouldNotBeEmpty is a std function that tests if a str.Trie is not empty.
// It logs the trie states before and after the search and checks if an error occurred.
func shouldNotBeEmpty(t gtests.Loggable, trie tries.Pattern) {
	t.Helper()
	// Log the trie states for debugging
	t.StackLogf("trie: { size: %v, length: %v }", trie.Size(), trie.Length())

	// Check if the PatternTrie is not empty
	if trie.IsEmpty() {
		t.Errorf("expected to not be empty but was\n")
	}

	// Assert length
	if trie.Length() == 0 {
		t.Errorf("expected length to be > 0 but was %v\n", trie.Length())
	}

	// Assert size
	if trie.Size() == 0 {
		t.Errorf("expected size to be > 0 but was %v\n", trie.Size())
	}
}

// shouldHaveSize is a std function that tests if a str.Trie has a certain size.
// It logs the trie states before and after the search and checks if an error occurred.
func shouldHaveSize(t gtests.Loggable, trie tries.Pattern, size int) {
	t.Helper()
	// Log the trie states for debugging
	t.StackLogf("trie: { size: %v, length: %v }", trie.Size(), trie.Length())

	// Assert size
	if trie.Size() != size {
		t.Errorf("expected size to be %v but was %v\n", size, trie.Size())
	}
}

// shouldHaveLength is a std function that tests if a str.Trie has a certain length.
// It logs the trie states before and after the search and checks if an error occurred.
func shouldHaveLength(t gtests.Loggable, trie tries.Pattern, length int) {
	t.Helper()
	// Log the trie states for debugging
	t.StackLogf("trie: { size: %v, length: %v }", trie.Size(), trie.Length())

	// Assert length
	if trie.Length() != length {
		t.Errorf("expected length to be %v but was %v\n", length, trie.Length())
	}
}

// shouldHaveSizeAndLength is a std function that tests if a str.Trie has a certain size and length.
// It logs the trie states before and after the search and checks if an error occurred.
func shouldHaveSizeAndLength(t gtests.Loggable, trie tries.Pattern, size int, length int) {
	t.Helper()
	// Log the trie states for debugging
	t.StackLogf("trie: { size: %v, length: %v }", trie.Size(), trie.Length())

	// Assert size
	if trie.Size() != size {
		t.Errorf("expected size to be %v but was %v\n", size, trie.Size())
	}

	// Assert length
	if trie.Length() != length {
		t.Errorf("expected length to be %v but was %v\n", length, trie.Length())
	}
}
