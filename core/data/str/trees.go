// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package str

import (
	"fmt"
	"github.com/andrerrcosta2/gtools/core/data/str/nodes"
)

type Tree[V any, N nodes.Typed[V]] interface {
	fmt.Stringer
	// Root returns the root node of the tree.
	Root() N
	// Size returns the number of nodes in the tree.
	Size() int
	// IsEmpty checks if the tree is empty.
	IsEmpty() bool
}

type BinaryTree[V any, N nodes.Typed[V]] interface {
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

type BranchTree[V any, N nodes.Typed[V], B any, C any] interface {
	Tree[V, N]
	// Insert a value into the branch tree
	Insert(branch B, value V)
	// GetBranch returns the values in the branch
	GetBranch(B) C
	// GetAllBranches returns the values in all the branches
	GetAllBranches() []C
}

type IOTrie[K comparable, I any, O any, N nodes.Trie[K, O, N]] interface {
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

type RedBlackTree[V any, C ~bool, N nodes.RedBlackTree[V, C, N]] interface {
	BinaryTree[V, N]
	// RotateLeft rotates the tree to the left
	RotateLeft(node N)
	// RotateRight rotates the tree to the right
	RotateRight(node N)
	// InsertFixup fixes the tree after an insertion
	InsertFixup(node N)
	// DeleteFixup fixes the tree after a deletion
	DeleteFixup(node N)
}

type Trie[K comparable, V any, N nodes.Trie[K, V, N]] interface {
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
