// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package tries

import (
	"github.com/andrerrcosta2/gtools/core/data/str"
	"github.com/andrerrcosta2/gtools/patterns/grammar"
	"github.com/andrerrcosta2/gtools/patterns/symbols"
	"github.com/andrerrcosta2/gtools/patterns/tries/internal/def"
	"github.com/andrerrcosta2/gtools/patterns/tries/internal/filter"
	"github.com/andrerrcosta2/gtools/patterns/tries/internal/nodes"
	"github.com/andrerrcosta2/gtools/patterns/tries/internal/validate"
	"strings"
	"sync"
)

var (
	PH   = def.Placeholder
	ROOT = def.Root
)

type Pattern str.IOTrie[string, symbols.Logical, grammar.Symbol]

// OfPatterns creates a new empty str.Trie of grammar.ByteSymbol
// It is safe to be used concurrently
func OfPatterns(strict bool, maxParallels int) Pattern {
	if strict {
		return newStrictPatternTrie(maxParallels)
	}
	return newPatternTrie(maxParallels)
}

// newPatternTrie creates a new empty patternTrie.
func newPatternTrie(maxParallel int) *patternTrie {
	if maxParallel < defaultConcurrentInternalOperations {
		return &patternTrie{root: nodes.Transition(def.Root.String(), nil), maxParallel: defaultConcurrentInternalOperations}
	}
	// Starts with a patternTrieTransitionNode
	return &patternTrie{root: nodes.Transition(def.Root.String(), nil), maxParallel: maxParallel}
}

// newStrictPatternTrie creates a new empty patternTrie in strict mode
func newStrictPatternTrie(maxParallel int) *patternTrie {
	if maxParallel < defaultConcurrentInternalOperations {
		return &patternTrie{
			root:        nodes.Transition(def.Root.String(), nil),
			strict:      make(map[string]bool),
			maxParallel: defaultConcurrentInternalOperations,
		}
	}
	// Starts with a patternTrieTransitionNode
	return &patternTrie{
		root:        nodes.Transition(def.Root.String(), nil),
		strict:      make(map[string]bool),
		maxParallel: maxParallel,
	}
}

// patternTrie is a str.Trie composed by symbols.
//
// It can be safely used under concurrent environments.
//
// It is composed of patternTrieNodes, which are either
// patternTrieTransitionNode or patternTrieSymbolNode.
//
// A patternTrieTransitionNode is a node that has a key and
// a map of children nodes. It is used to represent transitions
// between symbols.
//
// A patternTrieSymbolNode is a node that has a key, a value and
// a map of children nodes. It is used to represent symbols.
type patternTrie struct {
	mtx         sync.RWMutex
	root        nodes.PatternTrie
	size        int // Number of symbols in the trie. Must keep track of it
	length      int // Number of nodes in the trie. Must keep track of it
	maxParallel int
	strict      map[string]bool // TODO: Is a set to keep track of inserted values to apply strict rules
}

// Clear is a soft clear which relies on the garbage collector to remove dangling nodes.
// It might not be enough if the algorithm keeps references to any node.
func (t *patternTrie) Clear() {
	t.root = nodes.Transition(def.Root.String(), nil)
	t.size = 0
	t.length = 0
}

// Root returns the Root node of the newPatternTrie.
func (t *patternTrie) Root() nodes.PatternTrie {
	return t.root
}

// Size returns the number of symbols in the newPatternTrie.
func (t *patternTrie) Size() int {
	t.mtx.RLock()
	defer t.mtx.RUnlock()
	return t.size
}

// Length returns the number of nodes in the newPatternTrie.
func (t *patternTrie) Length() int {
	t.mtx.RLock()
	defer t.mtx.RUnlock()
	return t.length
}

// IsEmpty returns true if there are no symbols on the newPatternTrie.
func (t *patternTrie) IsEmpty() bool {
	t.mtx.RLock()
	defer t.mtx.RUnlock()
	return t.size == 0
}

// Insert adds a new pattern to the OfPatterns and associates the given symbol to its key.
// If the pattern already exists, it stores a new symbol into the node.
func (t *patternTrie) Insert(pattern string, symbol symbols.Logical) error {
	// Filter the pattern and the symbol to avoid newPatternTrie corruption
	err := filter.Insertion(pattern, symbol)
	if err != nil {
		return err
	}

	t.mtx.Lock()
	defer t.mtx.Unlock()

	return addSymbolToTrie(t, t.root, pattern, symbol)
}

// Search searches the PatternTrie for the given pattern and returns the associated symbol if found.
// It handles both literal and Placeholder symbols.
//
// Properties:
//  1. Case-sensitive searchSymbol.
//  2. Return immediate children if they're symbols.
//  3. Returns the symbol itself if the given path leads to one.
//  4. If the symbol matches a Placeholder, the symbol is returned with its Placeholder replaced by the placeholder contained in the path.
func (t *patternTrie) Search(s string) ([]grammar.Symbol, bool) {
	// validate the pattern to avoid unnecessary search
	if err := validate.Symbol(symbols.Of(s)); err != nil {
		return nil, false
	}

	t.mtx.RLock()
	defer t.mtx.RUnlock()

	return searchSymbolsByPattern(t.root, s, t.maxParallel)
}

// Delete removes a pattern from the newPatternTrie.
// It backtracks and removes nodes that don't have other children or symbols.
// The symbol must be provided to avoid false-positives
func (t *patternTrie) Delete(pattern string, symbol symbols.Logical) bool {
	// validate the symbol to avoid unnecessary search
	if err := validate.Symbol(symbols.Of(pattern)); err != nil {
		return false
	}

	t.mtx.Lock()
	defer t.mtx.Unlock()

	return deleteSymbolFromTrie(t, t.root, pattern, symbol)
}

func (t *patternTrie) String() string {
	return t.printTrie()
}

// printTrie is a recursive function that prints the trie nodes in a hierarchical manner.
func (t *patternTrie) printTrie() string {
	t.mtx.RLock()
	defer t.mtx.RUnlock()

	var result strings.Builder
	result.WriteString("OfPatterns { Root:")
	result.WriteString(t.root.PrintNode(1))
	result.WriteString("}")
	return result.String()
}

var _ str.Tree[grammar.Symbol] = (*patternTrie)(nil)
var _ str.IOTrie[string, symbols.Logical, grammar.Symbol] = (*patternTrie)(nil)
