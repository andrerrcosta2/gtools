// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package nodes

import (
	"errors"
	"fmt"
	"github.com/andrerrcosta2/gtools/patterns/grammar"
	"github.com/andrerrcosta2/gtools/patterns/symbols"
	"github.com/andrerrcosta2/gtools/patterns/tries/errs"
	"github.com/andrerrcosta2/gtools/patterns/tries/internal/def"
)

var UnknownNodeTypeError = errors.New("unknown node type\n")

// AsPatternTrieNode takes a DoubleLinkedTreeNode and tries to assert it as a PatternTrie.
// If the assertion is successful, it returns the node and nil.
// Otherwise, it returns nil and UnknownNodeTypeError.
func AsPatternTrieNode(node PatternTrie) (PatternTrie, error) {
	// Check if the node is a patternTrieTransitionNode
	if ptn, ok := node.(*transitionNode); ok {
		return ptn, nil
	}
	// Unknown node type
	return nil, UnknownNodeTypeError
}

// IsSymbol checks if the given node is a patternTrieSymbolNode.
// It returns true if the node is a symbol node, false otherwise.
func IsSymbol(node PatternTrie) bool {
	// Check if the node is a patternTrieSymbolNode by using type assertions.
	// If the node is a patternTrieSymbolNode, ok will be true.
	if _, ok := node.(*symbolNode); ok {
		return true
	}
	return false
}

func IsRoot(node PatternTrie) bool {
	return def.Root.Equal(symbols.String(node.Key()))
}

func IsTransition(node PatternTrie) bool {
	_, ok := asTransition(node)
	return ok
}

func GetSymbolFromNode(node PatternTrie) (grammar.Symbol, bool) {
	if s, ok := asSymbol(node); ok {
		return s.symbol, true
	}
	return nil, false
}

func HasSymbol(node PatternTrie, symbol grammar.Symbol) bool {
	if s, ok := asSymbol(node); ok {
		return s.symbol.Equal(symbol)
	}
	return false
}

// asTransition checks if the given node is a patternTrieTransitionNode.
// It returns a pointer to the patternTrieTransitionNode and true if the node is a transition node, nil and false otherwise.
func asTransition(node PatternTrie) (*transitionNode, bool) {
	// Check if the node is a patternTrieTransitionNode by using type assertions.
	// If the node is a patternTrieTransitionNode, ok will be true.
	if transition, ok := node.(*transitionNode); ok {
		return transition, ok
	}
	return nil, false
}

// asSymbol checks if the given node is a patternTrieSymbolNode.
// It returns a pointer to the patternTrieSymbolNode and true if the node is a symbol node, nil and false otherwise.
func asSymbol(node PatternTrie) (*symbolNode, bool) {
	// Check if the node is a patternTrieSymbolNode by using type assertions.
	// If the node is a patternTrieSymbolNode, ok will be true.
	if symbol, ok := node.(*symbolNode); ok {
		return symbol, ok
	}
	return nil, false
}

// HasPlaceholderChild returns the child of the given node that's a Placeholder.
// Returns the child and a boolean indicating whether the child exists.
func HasPlaceholderChild(node PatternTrie) (PatternTrie, bool) {
	// Get the child of the node that's a Placeholder
	return node.GetChild(def.Placeholder.String())
}

// StartsWithPlaceholder returns true if the pattern starts with a Placeholder
func StartsWithPlaceholder(pattern string) bool {
	return symbols.String(pattern).StartsWith(def.Placeholder)
}

// Deref removes a node from a PatternTrie
// It returns the parent node and an error.
// If the node is a symbol node, it will call the DerefSymbol method.
// If the node is a transition node, it will call the DerefTransition method.
// If the node type is unknown, it will return an error.
func Deref(node PatternTrie) (parent PatternTrie, err error) {
	// Check if the node is a symbol node
	if sn, ok := asSymbol(node); ok {
		// Call the DerefSymbol method
		return DerefSymbol(sn)
	} else if tn, ok := asTransition(node); ok {
		// Call the DerefTransition method
		return DerefTransition(tn)
	}
	// If the node type is unknown, return an error
	return nil, UnknownNodeTypeError
}

// DerefTransition removes a node from the newPatternTrie
// This method completely removes the trace of the node from the newPatternTrie
// If some measure must be taken to prevent side effects, it should be done before
// this method is called.
// It returns an error if the node has children or the node type is a symbol
func DerefTransition(node *transitionNode) (parent PatternTrie, err error) {
	if node.HasChildren() {
		return node.parent, errs.RemoveNodeWithChildrenNotAllowed
	}
	// register to return
	parent = node.parent
	// Remove the node from its parent
	node.parent.RemoveChild(node.key)

	// ToSet the parent and key of the node to its zero value
	node.children = nil
	node.parent = nil
	return
}

// DerefSymbol removes all references from the given symbol node from the newPatternTrie
// It sets the parent, symbol, root and key of the node to its zero value.
// It also removes the node from its parent's children slice.
func DerefSymbol(node *symbolNode) (parent PatternTrie, err error) {
	if node.HasChildren() {
		return node.parent, errs.RemoveNodeWithChildrenNotAllowed
	}

	// register to return
	parent = node.parent
	// Remove the node from its parent
	node.parent.RemoveChild(node.key)

	// ToSet the parent, symbol, root and key of the node to its zero value
	node.parent = nil
	node.symbol = nil
	node.key = ""
	return
}

// ChangeSymbolToTransition changes a patternTrieSymbolNode to a patternTrieTransitionNode.
// This is used when a symbol is removed from the trie, but the node still has children.
// This method is used to ensure children don't become orphaned.
//
// ### IMPORTANT: This method assumes that the symbol node already exists in the trie.
// This method MUST ENSURE the size of the trie is decreased but its length is not changed.
func ChangeSymbolToTransition(node PatternTrie) error {
	// Create a new transition node with the same key as the symbol node
	transition := Transition(string(node.Key()[0]), node.Parent())

	// Replace the key of the symbol node with the transition node
	err := node.Parent().AddChild(transition.Key(), transition)

	if err != nil {
		// Until here no states were changed, just return the error
		return err
	}

	// Create a track for partial changes
	children := make([]PatternTrie, 0)

	// Unique the transition node as the parent of each child
	for _, child := range node.Children() {
		err = child.SetParent(transition)
		if err != nil {
			// if changes were made, cancel before returning the error
			cancelSymbolToTransitionNode(node, children)
			// If the error is not nil, remove the  return the error
			return err
		}
		children = append(children, child)
	}

	return nil
}

func cancelSymbolToTransitionNode(originalParent PatternTrie, track []PatternTrie) {
	for _, child := range track {
		if err := child.SetParent(originalParent); err != nil {
			panic(fmt.Sprintf("Critical failure: unable to cancel changeSymbolToTransitionNode"+
				"(%s, %s)\n\tDetails: %s", originalParent.Value(), child.Value(), err))
		}
	}
}

// ChangeTransitionToSymbol changes a patternTrieTransitionNode to a patternTrieSymbolNode.
// This is used when a symbol is inserted into the newPatternTrie, but the node already exists.
// This method is used to ensure children don't become orphaned.
//
// ### Important: This method assumes that the symbol node already exists in the newPatternTrie.
// This method MUST ENSURE the size of the trie is increased but its length is not changed.
func ChangeTransitionToSymbol(node PatternTrie, symbol grammar.Symbol) error {
	// Create a new symbol node with the same key as the transition node
	// Ensure success before killing the transition node
	symNode := &symbolNode{
		parent:   node.Parent(),
		key:      string(node.Key()[0]),
		symbol:   symbol,
		children: make(map[string]PatternTrie),
	}

	// Replace the key of the transition node with the symbol node
	err := node.Parent().AddChild(symNode.key, symNode)

	if err != nil {
		// If the error is not nil, return the error
		return err
	}

	// Unique the symbol node as the parent of each child
	for _, child := range node.Children() {
		err = child.SetParent(symNode)
		if err != nil {
			// If the error is not nil, return the error
			return err
		}
		symNode.children[child.Key()] = child
	}

	// Remove the transition node
	return node.die()
}
