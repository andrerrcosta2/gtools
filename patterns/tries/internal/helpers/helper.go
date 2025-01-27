// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package helpers

import (
	"github.com/andrerrcosta2/gtools/patterns/symbols"
	"github.com/andrerrcosta2/gtools/patterns/tries/internal/def"
	"github.com/andrerrcosta2/gtools/patterns/tries/internal/nodes"
)

// NextNodeWithPlaceholderAsChild returns the next placeholder node in the trie from the given node and pattern.
func NextNodeWithPlaceholderAsChild(node nodes.PatternTrie, pattern string) (next nodes.PatternTrie, exists bool, patternLeft string) {
	next = node
	for i := range pattern {
		if next, exists = node.GetChild(def.Placeholder.String()); exists {
			_, after := symbols.String(pattern).ExtractSymbol(def.Placeholder)
			patternLeft = after.String()
			return
		} else if node, exists = node.GetChild(string(pattern[i])); !exists {
			patternLeft = pattern[i:]
			return
		}
	}
	return
}

// ExtractPlaceholderFromTransitionNode extracts the Placeholder value from the given node and pattern.
// It returns the extracted Placeholder value and the remaining part of the pattern, as well as
// the next node in the trie
// The next node is the node that should be traversed based on the remaining part of the pattern.
func ExtractPlaceholderFromTransitionNode(node nodes.PatternTrie, pattern string) (placeholder, rest string, nextNode nodes.PatternTrie) {
	// Iterate over the pattern, starting from the beginning
	for i, char := range pattern {
		// Append the current character to the Placeholder value
		placeholder += string(char)

		// If we're not at the end of the pattern, check if the next character is a child of the current node
		if i+1 < len(pattern) {
			nextChar := string(pattern[i+1])

			// If the next character is a child of the current node, it's the next node in the newPatternTrie
			if child, exists := node.GetChild(nextChar); exists {
				// The remaining part of the pattern is the substring starting from the next character
				rest = pattern[i+1:]
				// The next node is the child node
				nextNode = child
				// Break out of the loop
				break
			}
		}
	}
	// Return the extracted Placeholder value, the remaining part of the pattern,
	// and the next node in the trie if it exists.
	// If no next node exists, the next node is nil and the pattern ends with the placeholder value.
	return
}
