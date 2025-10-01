// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package str

import (
	"fmt"
	"github.com/andrerrcosta2/gtools/core/data"
	"github.com/andrerrcosta2/gtools/core/data/str/nodes"
)

// Tree represents a generic tree structure, which is a hierarchical data structure
// consisting of nodes connected by edges.
// Each node contains a value and references to its child nodes (if any).

type Tree[V any] interface {
	fmt.Stringer
	// Clear clears the tree
	Clear()
	// IsEmpty checks if the tree is empty.
	IsEmpty() bool
	// Size returns the number of nodes in the tree.
	Size() int
}

// BinaryTree represents a binary tree, where each node has at most two children
// (left and right). Binary trees are foundational for more specialized structures
// like Binary Search Trees (BSTs), AVL Trees, Red-Black Trees, and Heaps.
//
// Binary Trees are suitable For :
//   - Searching, sorting, and range queries
//   - Efficient insertion and deletion ops
//   - Implementing priority queues (via heaps)
//   - Efficiently maintaining sorted data.
//
// In average, binary trees maintains a balance factor of O(log n) in best case for all
// ops andO(n) in worst case.
type BinaryTree[V any] interface {
	Tree[V]
	Insert(value V)        // Insert a value
	Contains(value V) bool // Check if value exists
	Delete(value V)        // Remove a value

	Min() (V, bool)            // Get the smallest value
	Max() (V, bool)            // Get the largest value
	Floor(value V) (V, bool)   // Greatest value ≤ given value
	Ceiling(value V) (V, bool) // Smallest value ≥ given value

	Iterator() data.Iterator[V] // Support traversal
}

// BranchTree represents a multi-branch tree (e.g., k-ary tree or B-tree), where each node
// can have multiple children. This structure is particularly useful for managing large
// datasets efficiently, as it minimizes height, I/O ops and supports bulk ops.
//
// Branch Trees are suitable For :
//   - Managing large datasets
//   - File systems and databases (B-trees)
//   - Hierarchical clustering
//   - Multi-way search trees.
//
// In average, branch trees maintain a performance of O(log_m n) for insertions
// O(1) for searches and O(m^h) for iterations.
type BranchTree[V any, B any, C any] interface {
	Tree[V]
	// Insert a value into the branch tree
	Insert(branch B, value V)
	// GetBranch returns the values in the branch
	GetBranch(B) C
	// GetAllBranches returns the values in all the branches
	GetAllBranches() []C
}

// ColoredTree represents a binary tree where each node has an associated color
// (e.g., red or black).
// This is commonly used in self-balancing binary search trees like Red-Black Trees,
// which maintain logarithmic height by enforcing specific coloring rules.
// Red black trees do not allow duplicate values
//
// Colored Trees are suitable For :
//   - Implementing self-balancing binary search trees
//   - Maintaining sorted data with guaranteed logarithmic performance.
//   - Applications that requires frequent insertions and deletions (e.g., symbol
//     tables, ordered maps).
//
// In average, colored trees maintain a performance of O(log n) for all its ops.
type ColoredTree[V any, C ~uint8] interface {
	BinaryTree[V]
	ColorOf(v V) (color C, exists bool) // Retrieve node's color
}

// IOTrie extends Trie to support multiple outputs per key. It is particularly useful for
// scenarios where a single key maps to multiple values, such as in routing systems or
// autocomplete engines.
//
// IOTrie is suitable For :
//   - Multi-output mappings (e.g., routing tables with multiple destinations).
//   - Systems requiring prefix-based lookups with multiple results.
//   - Implementing autocomplete engines.
//
// In average, IOTries maintain a performance of O(key_length) for insertion and deletion
// ops, and O(L+M) for search ops.
type IOTrie[K comparable, I any, O any] interface {
	Tree[O]
	// Insert a key-value pair into the Trie.
	Insert(key K, value I) error
	// Search for a value by key in the Trie.
	// If the key is not found, it returns an empty slice and false.
	// If the key is found, it returns a slice with the associated value and true.
	Search(key K) ([]O, bool)
	// Delete Remove a key from the Trie
	// returns false if the key is not found
	Delete(key K, value I) bool
	// Length returns the real number of nodes in the Trie
	Length() int
}

// RedBlackTree specializes ColoredTree for Red-Black Trees, a type of self-balancing
// binary search tree. It enforces the following invariants:
//   - Each node is either red or black.
//   - The root is black.
//   - Red nodes cannot have red children.
//   - All paths from a node to its descendant leaves contain the same number of black nodes.
//
// Red-Black Trees are suitable For :
//   - Implementing self-balancing binary search trees
//   - Associative containers like maps and sets.
//   - Real-time systems requiring predictable performance.
//
// In average, red-black trees maintain a performance of O(log n) for all its ops.
type RedBlackTree[V any] interface {
	ColoredTree[V, nodes.RBColor]
}

// Trie represents a prefix tree, a specialized tree structure for storing strings or
// other sequences. Each node represents a character or key fragment, and paths from
// the root to a leaf represent complete keys.
//
// Tries are suitable For :
//   - Implementing prefix search and autocomplete features
//   - IP routing tables
//   - Dictionary implementations
//
// In average, 'tries' maintain a performance of O(key_length) for all its ops.
type Trie[K comparable, V any] interface {
	Tree[V]
	// Insert a key-value pair into the Trie.
	Insert(key K, value V) error
	// Search for a value by key in the Trie.
	// If the key is not found, it returns an empty slice and false.
	// If the key is found, it returns a slice with the associated value and true.
	Search(key K) ([]V, bool)
	// Delete Remove a key from the Trie
	// returns false if the key is not found
	Delete(key K, value V) bool
	// Length returns the real number of nodes in the Trie
	Length() int
}
