// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package tries

import (
	"github.com/andrerrcosta2/gtools/conc/syncs/semaph"
	"github.com/andrerrcosta2/gtools/patterns/grammar"
	"github.com/andrerrcosta2/gtools/patterns/tries/internal/handlers/insert"
	"github.com/andrerrcosta2/gtools/patterns/tries/internal/handlers/remove"
	"github.com/andrerrcosta2/gtools/patterns/tries/internal/handlers/search"
	"github.com/andrerrcosta2/gtools/patterns/tries/internal/nodes"
	"log"
	"sync"
)

// searchSymbol is a wrapper for search.Concurrently. It abstracts the complexity of the synchronization by creating a
// SynchronizableSet and listening to it. It returns the nodes found and whether any node was found.
func searchSymbolsByPattern(from nodes.PatternTrie, pattern string, maxParallels int) (symbols []grammar.Symbol, found bool) {
	wg := &sync.WaitGroup{}

	//wg.Add(1)
	// Start the search concurrently.
	//go search.Concurrently(nodes.NewSearch(from, pattern, &symbols), wg, semaph.Channel(maxParallels))
	search.Concurrently(nodes.NewSearch(from, pattern, &symbols), wg, semaph.Channel(maxParallels))
	// set the synchronization ending point
	wg.Wait()

	// Return the nodes found and whether any node was found.
	return symbols, len(symbols) > 0
}

// addSymbolToTrie inserts a new pattern into the trie starting from the given node
// It traverses the trie by moving forward one node at a time, until it finds a node
// that doesn't have a child with the next character in the pattern. At that point, it
// returns the last node it found, which is the insertion point. Then it deploys the
// symbol into the trie by recursively creating new nodes if necessary
func addSymbolToTrie(t *patternTrie, from nodes.PatternTrie, pattern string, symbol grammar.Symbol) error {
	// insert the symbol
	length, err := insert.Symbol(from, pattern, symbol)
	if err != nil {
		return err
	}

	// increment the size
	t.size++
	t.length += length

	return nil
}

// deleteSymbolFromTrie removes a symbol from the OfPatterns.
// It searches the OfPatterns for the given pattern and removes the associated symbol if found.
// It returns an error if the node cannot be removed.
func deleteSymbolFromTrie(t *patternTrie, from nodes.PatternTrie, pattern string, symbol grammar.Symbol) bool {
	// Search the OfPatterns for the given pattern
	if symbolNode, found := search.NodeByPattern(from, pattern, symbol); found {
		// Remove the associated symbol
		if err, length := remove.Symbol(symbolNode); err != nil {
			log.Printf("warning: %s", err.Error())
			return false
		} else {
			t.size--
			t.length -= length
			return true
		}
	}

	return false
}
