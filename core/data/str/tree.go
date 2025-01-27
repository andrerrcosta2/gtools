// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package str

import (
	"fmt"
	"github.com/andrerrcosta2/gtools/core/data/str/nodes"
)

type Tree[V any, N nodes.Node[V]] interface {
	fmt.Stringer
	// Root returns the root node of the tree.
	Root() N
	// Size returns the number of nodes in the tree.
	Size() int
	// IsEmpty checks if the tree is empty.
	IsEmpty() bool
}

type SingleLinkedTreeNode[V any, N nodes.Node[V]] interface {
	fmt.Stringer
	nodes.Node[V]
	// Children returns the children of the node.
	Children() []N
}

type SingleLinkedKeyValueTreeNode[K comparable, V any, N nodes.KeyValue[K, V]] interface {
	fmt.Stringer
	nodes.KeyValue[K, V]
	Children() []N
}

type DoubleLinkedTreeNode[V any, N SingleLinkedTreeNode[V, N]] interface {
	SingleLinkedTreeNode[V, N]
	Parent() N
}

type DoubleLinkedTreeKeyValueNode[K comparable, V any, N nodes.KeyValue[K, V]] interface {
	SingleLinkedKeyValueTreeNode[K, V, N]
	Parent() N
}

type BinaryTree[V any, N nodes.Node[V]] interface {
	Tree[V, N]
	// LeftChild returns the left child of a given node.
	// If the node does not have a left child, it returns nil.
	LeftChild() N

	// RightChild returns the right child of a given node.
	// If the node does not have a right child, it returns nil.
	RightChild() N
	Insert(value V)           // Insert a value into the binary tree
	Search(value V) (N, bool) // Search for a node with the given value
	Delete(value V)
}

type Trie[K comparable, V any, N TrieNode[K, V, N]] interface {
	Tree[V, N]
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

type IOTrie[K comparable, I any, O any, N TrieNode[K, O, N]] interface {
	Tree[O, N]
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

type TrieNode[K comparable, V any, N DoubleLinkedTreeKeyValueNode[K, V, N]] interface {
	DoubleLinkedTreeKeyValueNode[K, V, N]
	// ChildrenKeys is a slice of transition chars
	ChildrenKeys() []K
	// ChildrenValues is a slice of values
	ChildrenValues() []V
}

type BranchTree[V any, N nodes.Node[V], B any, C any] interface {
	Tree[V, N]
	// Insert a value into the branch tree
	Insert(branch B, value V)
	// GetBranch returns the values in the branch
	GetBranch(B) C
	// GetAllBranches returns the values in all the branches
	GetAllBranches() []C
}
