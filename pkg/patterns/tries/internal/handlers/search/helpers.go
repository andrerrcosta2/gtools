// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package search

import (
	"fmt"
	"github.com/andrerrcosta2/gtools/core/functions/runnables"
	"github.com/andrerrcosta2/gtools/core/gtools"
	"github.com/andrerrcosta2/gtools/patterns/grammar"
	"github.com/andrerrcosta2/gtools/patterns/symbols"
	"github.com/andrerrcosta2/gtools/patterns/tries/internal/nodes"
	"sync"
)

// Concurrently searches for the possible nodes in the trie for a given pattern
// the use of parallelism here is justified by the 3 main flows:
//
//  1. searchNextLiteral - considers each depth on the pattern as a literal.
//  2. searchNextPlaceholder - considers each depth on the pattern as a placeholder.
//  3. searchRecursively - considers each symbol found to be forwarded recursively from the root.
func Concurrently(n *nodes.Search, wg *sync.WaitGroup, sem gtools.Semaphore) {
	fmt.Printf("Concurrently: [%s]\n", n.Pattern())
	fmt.Printf("Semaphore: capacity: %d, remaining capacity: %d\n", sem.Cap(), sem.Rem())
	fmt.Printf("WaitGroup: waiting: %+v\n", wg)
	if n.IsEmpty() {
		fmt.Printf("Is Empty\n")
		if n.IsSymbol() {
			fmt.Printf("Is Symbol\n")
			n.Write()
		}
		return
	}

	if !n.IsSymbol() {
		fmt.Printf("Not Symbol search: [%s]\n", n.Pattern())
		childrenOnlySearch(n.Clone(), wg, sem)
	} else {
		fmt.Printf("Symbol search: [%s]\n", n.Pattern())
		searchSymbolNode(n.Clone(), wg, sem)
	}
}

func NodeByPattern(from nodes.PatternTrie, pattern string, symbol grammar.Symbol) (node nodes.PatternTrie, exists bool) {
	if len(pattern) == 0 {
		return
	}

	for i := 0; i < len(pattern); i++ {
		if from, exists = from.GetChild(string(pattern[i])); !exists {
			return
		}
	}

	if nodes.HasSymbol(from, symbol) {
		return from, true
	}

	return
}

func searchSymbolNode(b *nodes.Search, wg *sync.WaitGroup, sem gtools.Semaphore) {
	switch b.Port() {
	case symbols.Open:
		openSearch(b, wg, sem)
	case symbols.ChildrenOnly:
		childrenOnlySearch(b, wg, sem)
	case symbols.RecursiveOnly:
		recursiveOnlySearch(b, wg, sem)
	case symbols.RecursiveToSelf:
		recursiveToSelfSearch(b, wg, sem)
	case symbols.Breakpoint:
		breakpointSearch(b, wg, sem)
	case symbols.ChildrenNegative:
		childrenNegativeSearch(b, wg, sem)
	case symbols.ChildrenExact:
		childrenExactSearch(b, wg, sem)
	case symbols.Closed:
		closedSearch(b)
	}
}

func childrenOnlySearch(b *nodes.Search, wg *sync.WaitGroup, sem gtools.Semaphore) {
	fmt.Printf("childrenOnlySearch: [%s]\n", b.Pattern())
	wg.Add(2)
	// Search for next literal
	go runnables.SemaphoredSync(wg, sem, func() {
		fmt.Printf("childrenOnlySearch:literalSearch: [%s]\n", b.Pattern())
		if next, exists := b.Nmlsn(); exists {
			fmt.Printf("childrenOnlySearch:literalSearch:Concurrently: [%s]\n", next.Pattern())
			Concurrently(next, wg, sem)
		}
	})

	// Search for next placeholder
	// I think this is wrong. I think only symbols should allow placeholders.
	go runnables.SemaphoredSync(wg, sem, func() {
		fmt.Printf("childrenOnlySearch:placeholderSearch: [%s]\n", b.Pattern())
		if next, exists := b.Nmpn(); exists {
			fmt.Printf("childrenOnlySearch:placeholderSearch:Concurrently: [%s]\n", next.Pattern())
			Concurrently(next, wg, sem)
		}
	})
}

func openSearch(b *nodes.Search, wg *sync.WaitGroup, sem gtools.Semaphore) {
	wg.Add(3)
	// Search for next literal
	go runnables.SemaphoredSync(wg, sem, func() {
		if next, exists := b.Nmlsn(); exists {
			Concurrently(next, wg, sem)
		}
	})

	// Search for next placeholder
	go runnables.SemaphoredSync(wg, sem, func() {
		if next, exists := b.Nmpn(); exists {
			Concurrently(next, wg, sem)
		}
	})

	// Search recursively
	go runnables.SemaphoredSync(wg, sem, func() { Concurrently(b, wg, sem) })
}

func recursiveOnlySearch(b *nodes.Search, wg *sync.WaitGroup, sem gtools.Semaphore) {
	wg.Add(1)
	go runnables.SemaphoredSync(wg, sem, func() { Concurrently(b, wg, sem) })
}

func recursiveToSelfSearch(b *nodes.Search, wg *sync.WaitGroup, sem gtools.Semaphore) {
	wg.Add(1)
	go runnables.SemaphoredSync(wg, sem, func() { Concurrently(b, wg, sem) })
}

func breakpointSearch(b *nodes.Search, wg *sync.WaitGroup, sem gtools.Semaphore) {

}

func childrenNegativeSearch(b *nodes.Search, wg *sync.WaitGroup, sem gtools.Semaphore) {

}

func childrenExactSearch(b *nodes.Search, wg *sync.WaitGroup, sem gtools.Semaphore) {
	wg.Add(1)
	// Search for next literal
	go runnables.SemaphoredSync(wg, sem, func() {
		if next, exists := b.Nmlsn(); exists {
			Concurrently(next, wg, sem)
		}
	})
}

func closedSearch(b *nodes.Search) {
	if b.IsEmpty() {
		b.Write()
	}
}
