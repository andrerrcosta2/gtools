// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package insert

import (
	"fmt"
	"github.com/andrerrcosta2/gtools/patterns/grammar"
	"github.com/andrerrcosta2/gtools/patterns/symbols"
	"github.com/andrerrcosta2/gtools/patterns/tries/internal/def"
	"github.com/andrerrcosta2/gtools/patterns/tries/internal/nodes"
)

func Symbol(from nodes.PatternTrie, pattern string, symbol grammar.Symbol) (length int, err error) {
	// Find the insertion point
	point, skip := findInsertionPoint(from, pattern)
	// Deploy the symbol into the trie by recursively creating new nodes if necessary
	return deploySymbol(point, pattern[skip:], symbol)
}

// findInsertionPoint finds the insertion point in the trie from the given starting node
// It traverses the trie by moving forward one node at a time, until it finds a node
// that doesn't have a child with the next character in the pattern. At that point, it
// returns the last node it found, which is the insertion point
func findInsertionPoint(from nodes.PatternTrie, pattern string) (point nodes.PatternTrie, skip int) {
	skip = -1

	// Traverse the trie by moving forward one node at a time, until it finds a node
	// that doesn't have a child with the next character in the pattern
	for exists := from != nil; exists; from, exists, pattern = fnn(from, pattern) {
		point = from
		skip++
	}

	// If the pattern is fully traversed, return the last node
	return
}

// fnn finds the next node from the given starting node
// and returns the next node, whether it exists and how many characters
// to skip in the pattern
func fnn(from nodes.PatternTrie, pattern string) (next nodes.PatternTrie, exists bool, left string) {
	if len(pattern) == 0 {
		return
	}
	// If the pattern starts with a Placeholder, find the Placeholder
	// child of the starting node
	if nodes.StartsWithPlaceholder(pattern) {
		next, exists = nodes.HasPlaceholderChild(from)
		// The skip value is the length of the Placeholder minus 1 since
		// the Placeholder is already accounted for
		left = pattern[def.Placeholder.Len():]
		return
	}
	// Otherwise, try to find the node corresponding to the
	// first character of the pattern
	next, exists = from.GetChild(pattern[0:1])
	left = pattern[1:]
	return
}

// deploySymbol inserts a new symbol by spreading the symbol through the trie.
// If the error isn't nil, it backtracks the trie and returns the error
func deploySymbol(from nodes.PatternTrie, pattern string, symbol grammar.Symbol) (length int, err error) {
	var track []nodes.PatternTrie

	// Deploy the symbol recursively
	track, err = deploySymbolRecursively(from, pattern, symbol, []nodes.PatternTrie{})
	if err != nil {
		// If there was an error during the spreading, backtrack the trie and return the error
		cancelInsertion(track)
		return
	}

	return len(track), nil
}

// deploySymbolRecursively creates a new symbol path from the given string and symbol.
// It returns the final node of the path, the number of nodes created and an error.
// If the error isn't nil, it means that the path couldn't be created.
func deploySymbolRecursively(p nodes.PatternTrie, s string, symbol grammar.Symbol, track []nodes.PatternTrie) ([]nodes.PatternTrie, error) {
	// If the given string matches the Placeholder symbol
	// There is a chance here of having a placeholder as the last character of the string
	// in that case the symbol should store a whole placeholder
	if symbols.String(s).StartsWith(def.Placeholder) {
		// Create a new Placeholder node and extract the remaining string after the Placeholder
		child, rest, err, hasNext := addPlaceholderNode(p, s, symbol)
		if err != nil {
			return track, err
		}
		track = append(track, child)

		if hasNext {
			// Recursively create the symbol path from the remaining string
			return deploySymbolRecursively(child, rest, symbol, track)
		}
	} else if len(s) > 1 {
		// Create a new transition node with the given key and parent
		child, err := addTransitionNode(p, s[0])
		if err != nil {
			return track, err
		}
		track = append(track, child)

		// Recursively create the symbol path from the remaining string
		return deploySymbolRecursively(child, s[1:], symbol, track)
	} else if len(s) == 1 {
		// Create a new symbol node with the given key, parent, and symbol
		child, err := addSymbolNode(p, string(s[0]), symbol)
		if err != nil {
			return track, err
		}
		track = append(track, child)

		// Return the track of created nodes
		return track, nil
	}

	return track, nil // Return track of created nodes
}

// cancelInsertion undoes the insert of a symbol in the trie by removing the nodes created during the insert.
// It is used when there is an error during the insert to backtrack the trie and return the error.
func cancelInsertion(track []nodes.PatternTrie) {
	// Iterate over the nodes created in reverse order
	// and remove the nodes that were created
	for i := len(track) - 1; i > 0; i-- {
		// Dereference the node
		_, err := nodes.Deref(track[i])
		if err != nil {
			// This is a critical failure and the program should panic
			panic(fmt.Sprintf("Critical failure: unable to dereference node while canceling insert. error: %s", err))
		}
	}
}

// addPlaceholderNode adds a new Placeholder node to the given parent node.
// It returns the newly created node, the remaining string after the Placeholder and an error.
// If the error isn't nil, it means that the node couldn't be added.
func addPlaceholderNode(n nodes.PatternTrie, s string, symbol grammar.Symbol) (node nodes.PatternTrie, rest string, err error, hasNext bool) {
	ph, xtr := symbols.String(s).ExtractSymbol(def.Placeholder)

	// if the rest is empty after the placeholder extraction, create a new symbol node
	if xtr.IsEmpty() {
		node, err = addSymbolNode(n, ph.String(), symbol)
	} else {
		// Otherwise create a new transition node with the given key and parent
		node = nodes.Transition(ph.String(), n)
	}

	if err != nil {
		return nil, s, err, false
	}

	// Append the child to the parent node
	// The error while setting the child is the same as while creating the node
	_ = n.AddChild(ph.String(), node)

	// Return the created node, the extracted string the error and if the string has more characters
	return node, xtr.String(), nil, !xtr.IsEmpty()
}

// addTransitionNode adds a new transition node to the given parent node.
// It returns the newly created node and an error.
// If the error isn't nil, it means that the node couldn't be added.
func addTransitionNode(n nodes.PatternTrie, key uint8) (nodes.PatternTrie, error) {
	// Create a new transition node with the given key and parent
	child := nodes.Transition(string(key), n)

	// Append the child to the parent node
	err := n.AddChild(string(key), child)
	if err != nil {
		return nil, err
	}

	return child, nil
}

// addSymbolNode adds a new symbol node to the given parent node.
// It returns the newly created node and an error.
// If the error isn't nil, it means that the node couldn't be added.
func addSymbolNode(n nodes.PatternTrie, key string, symbol grammar.Symbol) (nodes.PatternTrie, error) {
	// Create a new symbol node with the given key, parent, and symbol
	child := nodes.Symbol(key, n, symbol)

	// Append the child to the parent node
	err := n.AddChild(key, child)

	if err != nil {
		// If the error is not nil, return the error
		return nil, err
	}

	// Return the newly created node
	return child, nil
}
