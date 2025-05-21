// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package nodes

import "fmt"

type LinkedUp[N any] interface {
	Parent() N
	SetParent(p N)
}

type BinaryLinkedUp[N any] interface {
	LinkedUp[N]
	LeftChild() N
	RightChild() N
	SetLeftChild(c N)
	SetRightChild(c N)
}

type BinaryTree[V any, N Typed[V]] interface {
	BinaryLinkedUp[N]
	Typed[V]
}

type ColoredTree[V any, N Typed[V], C ~uint8] interface {
	BinaryTree[V, N]
	Color() C
	SetColor(c C)
}

type DoubleLinkedKeyValueTree[K comparable, V any, N KeyValue[K, V]] interface {
	SingleLinkedKeyValueTree[K, V, N]
	Parent() N
}

type DoubleLinkedTree[V any, N SingleLinkedTree[V, N]] interface {
	SingleLinkedTree[V, N]
	Parent() N
}

type RBColor uint8

const (
	Black RBColor = 0
	Red   RBColor = 1
)

func (r RBColor) String() string {
	if r == Black {
		return "Black"
	}
	return "Red"
}

// RedBlackTree extends a standard binary node with color and parent tracking
type RedBlackTree[V any, N Typed[V]] interface {
	ColoredTree[V, N, RBColor]
}

type SingleLinkedKeyValueTree[K comparable, V any, N KeyValue[K, V]] interface {
	fmt.Stringer
	KeyValue[K, V]
	Children() []N
}

type SingleLinkedTree[V any, N Typed[V]] interface {
	fmt.Stringer
	Typed[V]
	// Children returns the children of the node.
	Children() []N
}

type Trie[K comparable, V any, N DoubleLinkedKeyValueTree[K, V, N]] interface {
	DoubleLinkedKeyValueTree[K, V, N]
	// ChildrenKeys is a slice of transition chars
	ChildrenKeys() []K
	// ChildrenValues is a slice of values
	ChildrenValues() []V
}
