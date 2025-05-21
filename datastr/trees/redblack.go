// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package trees

import (
	"github.com/andrerrcosta2/gtools/core/data"
	"github.com/andrerrcosta2/gtools/core/data/comparators"
	"github.com/andrerrcosta2/gtools/core/data/str"
	"github.com/andrerrcosta2/gtools/core/data/str/nodes"
	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"github.com/andrerrcosta2/gtools/core/format/sprints"
	"github.com/andrerrcosta2/gtools/core/generics"
	"github.com/andrerrcosta2/gtools/datastr/trees/internal/redblack"
)

func RedBlack[V any](comparator comparators.Typed[V]) str.RedBlackTree[V] {
	return &redBlack[V]{cmp: comparator}
}

var _ str.RedBlackTree[any] = (*redBlack[any])(nil)

// redBlack is a red-black tree implementation
// A red black tree is a tree with the following properties:
// 1. A node it either red or black
// 2. The root and leaves (leaves are nil) are black
// 3. If a node is red, then both its children are black
// 4. All paths from a node to its NIL descendants contain the same number of black nodes
type redBlack[V any] struct {
	abstract[*redblack.Node[V], V]
	cmp comparators.Typed[V]
}

// Ceiling returns the smallest value greater than or equal to the given value
func (t *redBlack[V]) Ceiling(value V) (V, bool) {
	var candidate V
	found := false

	current := t.Root()
	for current != nil {
		cmp := t.cmp.Compare(value, current.Value())
		if cmp == 0 { // same value
			return current.Value(), true
		} else if cmp < 0 {
			// can be the ceiling
			candidate = current.Value()
			found = true
			current = current.LeftChild()
		} else {
			current = current.RightChild()
		}
	}

	return candidate, found
}

// Clear clears the tree
func (t *redBlack[V]) Clear() {
	t.SetRoot(nil)
	t.SetSize(0)
}

// ColorOf returns the color of the node for a given value
func (t *redBlack[V]) ColorOf(value V) (nodes.RBColor, bool) {
	n, ok := redblack.FindNode[V](t.Root(), value, t.cmp.Compare)
	if ok {
		return n.Color(), true
	}
	return nodes.Black, false
}

func (t *redBlack[V]) Contains(value V) bool {
	_, ok := redblack.FindNode(t.Root(), value, t.cmp.Compare)
	return ok
}

func (t *redBlack[V]) Delete(value V) {
	n, ok := redblack.FindNode[V](t.Root(), value, t.cmp.Compare)
	if !ok {
		return
	}
	redblack.DeleteNode[V](t.Root(), n)
	t.SetSize(t.Size() - 1)
}

// Floor returns the largest value less than or equal to the given value
func (t *redBlack[V]) Floor(value V) (V, bool) {
	var candidate V
	found := false

	current := t.Root()
	for current != nil {
		cmp := t.cmp.Compare(value, current.Value())
		if cmp == 0 { // equals
			return current.Value(), true
		} else if cmp > 0 {
			// can be the floor
			candidate = current.Value()
			found = true
			current = current.RightChild()
		} else {
			current = current.LeftChild()
		}
	}

	return candidate, found
}

func (t *redBlack[V]) Insert(value V) {
	node := redblack.NewNode(value)
	if t.Root() == nil {
		t.SetRoot(node)
		return
	}
	redblack.InsertNode(t.Root(), node, t.cmp.Compare)
	//fmx.Redf("After insert: %s\n", t.Root())
	t.SetSize(t.Size() + 1)
}

func (t *redBlack[V]) IsEmpty() bool {
	return t.Size() == 0
}

func (t *redBlack[V]) Iterator() data.Iterator[V] {
	return nil
}

// Max returns the highest value in the tree
func (t *redBlack[V]) Max() (V, bool) {
	if t.IsEmpty() {
		return generics.Zero[V](), false
	}
	current := t.Root()
	for current.RightChild() != nil {
		current = current.RightChild()
	}
	return current.Value(), true
}

// Min returns the lowest value in the tree
func (t *redBlack[V]) Min() (V, bool) {
	if t.IsEmpty() {
		return generics.Zero[V](), false
	}
	current := t.Root()
	for current.LeftChild() != nil {
		current = current.LeftChild()
	}
	return current.Value(), true
}

func (t *redBlack[V]) String() string {
	tab := indent.Zero()
	root := "*redblack.Node<nil>"
	if t.Root() != nil {
		root = t.Root().Sprint(tab.Inc())
	}
	return sprints.Closedobj(tab, "*trees.redBlack",
		sprints.Fieldf(tab, "root", "%s", root),
	)
}
