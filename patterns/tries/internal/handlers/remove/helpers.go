// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package remove

import (
	"github.com/andrerrcosta2/gtools/patterns/tries/internal/nodes"
)

// Symbol removes a symbol node from the trie or changes to transition node if it has children
// Then it proceeds with a backtrack removal of unnecessary transition nodes.
// It returns an error if the node cannot be removed.
func Symbol(node nodes.PatternTrie) (err error, length int) {
	if node.HasChildren() {
		// the node has children, change it to a transition node
		return nodes.ChangeSymbolToTransition(node), 0
	}

	// mark to be removed
	next := node

	// Loop until the root is reached or a node with children is found
	for {
		next, err = nodes.Deref(next)
		if err != nil {
			return
		}

		length++

		if nodes.IsSymbol(next) || nodes.IsRoot(next) || next.HasChildren() {
			// if the node is a symbol or the root, or has children, break the loop
			break
		}
	}

	return
}
