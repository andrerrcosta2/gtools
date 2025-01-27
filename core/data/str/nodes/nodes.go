// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package nodes

import (
	"github.com/andrerrcosta2/gtools/core/data"
	"github.com/andrerrcosta2/gtools/core/domain/constraints/prim"
)

type Node[V any] interface {
	// Value returns the value of the node.
	//
	// Returns the value of the node as a value of type V.
	Value() V
}

type KeyValue[K any, V any] interface {
	Node[V]
	// Key returns the key of the node.
	//
	// Returns the key of the node as a value of type K.
	Key() K
}

type Linked[V any, N Node[V]] interface {
	Node[V]
	data.Linked[N]
}

type LinkedKeyValue[K any, V any, N KeyValue[K, V]] interface {
	KeyValue[K, V]
	data.Linked[N]
}

type DoubleLinked[V any, N Node[V]] interface {
	Node[V]
	data.DoubleLinked[N]
}

type DoubleLinkedKeyValue[K any, V any, N LinkedKeyValue[K, V, N]] interface {
	LinkedKeyValue[K, V, N]
	// Prev returns the previous node in the linked list.
	//
	// Returns the previous node in the linked list as a value of type N.
	Prev() N
}

type Serializable[V any, S prim.Serializable] interface {
	Node[V]
	data.Serializable[S]
}

type Branchable[B any, V any] interface {
	Node[V]
	data.Branchable[B]
}

type SerializableBranchable[B any, V any, S prim.Serializable] interface {
	data.Branchable[B]
	Serializable[V, S]
}

type Sequential[V any] Serializable[V, int]

type SequentialBranchable[B any, V any] SerializableBranchable[B, V, int]
