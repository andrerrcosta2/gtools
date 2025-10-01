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
	"sync"
)

// RedBlack returns a new red-black tree
func RedBlack[V any](comparator comparators.Typed[V], values ...V) str.RedBlackTree[V] {
	if comparator == nil {
		panic("comparator must be provided")
	}
	rb := &redBlack[V]{cmp: comparator}
	for _, value := range values {
		rb.Insert(value)
	}
	return rb
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

// Ceiling returns the smallest value greater than or compare to the given value
func (t *redBlack[V]) Ceiling(value V) (V, bool) {
	return redblack.Ceiling(t.Root(), value, t.cmp.Compare)
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

// Contains returns true if the tree contains the given value
func (t *redBlack[V]) Contains(value V) bool {
	_, ok := redblack.FindNode(t.Root(), value, t.cmp.Compare)
	return ok
}

// Delete deletes a value from the tree
func (t *redBlack[V]) Delete(value V) {
	n, ok := redblack.FindNode[V](t.Root(), value, t.cmp.Compare)
	if !ok {
		return
	}
	root := redblack.DeleteNode[V](t.Root(), n)
	t.SetRoot(root)
	t.SetSize(t.Size() - 1)
}

// Floor returns the largest value less than or compare to the given value
func (t *redBlack[V]) Floor(value V) (V, bool) {
	return redblack.Floor(t.Root(), value, t.cmp.Compare)
}

// Insert inserts a value into the tree
func (t *redBlack[V]) Insert(value V) {
	if t.Root() == nil {
		t.SetRoot(redblack.Black(value))
		t.SetSize(1)
	} else {
		root, ok := redblack.InsertNode[V](t.Root(), redblack.Red(value), t.cmp.Compare)
		if ok {
			t.SetRoot(root) // The root must always be updated
			t.SetSize(t.Size() + 1)
		}
	}
}

// IsEmpty returns true if the tree is empty
func (t *redBlack[V]) IsEmpty() bool {
	return t.Size() == 0
}

// Iterator returns an iterator for the tree
// This iterator reflects live changes to the tree.
// Modifying the tree during iteration may cause undefined behavior.
func (t *redBlack[V]) Iterator() data.Iterator[V] {
	return redBlackIterator(t.cmp, redblack.MinimumFrom(t.Root()))
}

// Max returns the highest value in the tree
func (t *redBlack[V]) Max() (V, bool) {
	if t.IsEmpty() {
		return generics.Zero[V](), false
	}
	return redblack.MaximumFrom(t.Root()).Value(), true
}

// Min returns the lowest value in the tree
func (t *redBlack[V]) Min() (V, bool) {
	if t.IsEmpty() {
		return generics.Zero[V](), false
	}
	return redblack.MinimumFrom(t.Root()).Value(), true
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

func redBlackIterator[V any](cmp comparators.Typed[V], min *redblack.Node[V]) data.Iterator[V] {
	return &rbtIterator[V]{cmp: cmp, it: min}
}

type rbtIterator[V any] struct {
	cmp comparators.Typed[V]
	it  *redblack.Node[V]
}

func (i *rbtIterator[V]) Next() (next V) {
	if i.it == nil {
		return next
	}
	next = i.it.Value()
	i.next()
	return next
}

func (i *rbtIterator[V]) next() {
	if i.it.RightChild() != nil {
		i.it = redblack.MinimumFrom(i.it.RightChild())
		return
	}

	// Walk up until we find a node that is the left child of its parent
	parent := i.it.Parent()
	for parent != nil && i.it == parent.RightChild() {
		i.it = parent
		parent = parent.Parent()
	}
	i.it = parent
}

// HasNext Checks if there's a next element
func (i *rbtIterator[V]) HasNext() bool {
	return i.it != nil
}

func ConcRedBlack[V any](cmp comparators.Typed[V], values ...V) str.RedBlackTree[V] {
	if cmp == nil {
		panic("comparator must be provided")
	}
	rb := &concRedBlack[V]{cmp: cmp}
	for _, value := range values {
		if rb.Root() == nil {
			rb.SetRoot(redblack.Black(value))
			rb.SetSize(1)
		} else {
			root, ok := redblack.InsertNode[V](rb.Root(), redblack.Red(value), cmp.Compare)
			if ok {
				rb.SetRoot(root) // The root must always be updated
				rb.SetSize(rb.Size() + 1)
			}
		}
	}
	return rb
}

type concRedBlack[V any] struct {
	abstract[*redblack.Node[V], V]
	cmp comparators.Typed[V]
	mtx sync.RWMutex
}

// Ceiling returns the smallest value greater than or compare to the given value
func (t *concRedBlack[V]) Ceiling(value V) (V, bool) {
	t.mtx.RLock()
	defer t.mtx.RUnlock()
	return redblack.Ceiling(t.Root(), value, t.cmp.Compare)
}

// Clear clears the tree
func (t *concRedBlack[V]) Clear() {
	t.mtx.Lock()
	defer t.mtx.Unlock()
	t.SetRoot(nil)
	t.SetSize(0)
}

// ColorOf returns the color of the node for a given value
func (t *concRedBlack[V]) ColorOf(value V) (nodes.RBColor, bool) {
	t.mtx.RLock()
	defer t.mtx.RUnlock()
	n, ok := redblack.FindNode[V](t.Root(), value, t.cmp.Compare)
	if ok {
		return n.Color(), true
	}
	return nodes.Black, false
}

// Contains returns true if the tree contains the given value
func (t *concRedBlack[V]) Contains(value V) bool {
	t.mtx.RLock()
	defer t.mtx.RUnlock()
	_, ok := redblack.FindNode(t.Root(), value, t.cmp.Compare)
	return ok
}

// Delete deletes a value from the tree
func (t *concRedBlack[V]) Delete(value V) {
	t.mtx.Lock()
	defer t.mtx.Unlock()
	n, ok := redblack.FindNode[V](t.Root(), value, t.cmp.Compare)
	if !ok {
		return
	}
	root := redblack.DeleteNode[V](t.Root(), n)
	t.SetRoot(root)
	t.SetSize(t.Size() - 1)
}

// Floor returns the largest value less than or compare to the given value
func (t *concRedBlack[V]) Floor(value V) (V, bool) {
	t.mtx.RLock()
	defer t.mtx.RUnlock()
	return redblack.Floor(t.Root(), value, t.cmp.Compare)
}

// Insert inserts a value into the tree
func (t *concRedBlack[V]) Insert(value V) {
	t.mtx.Lock()
	defer t.mtx.Unlock()
	if t.Root() == nil {
		t.SetRoot(redblack.Black(value))
		t.SetSize(1)
	} else {
		root, ok := redblack.InsertNode[V](t.Root(), redblack.Red(value), t.cmp.Compare)
		if ok {
			t.SetRoot(root) // The root must always be updated
			t.SetSize(t.Size() + 1)
		}
	}
}

// IsEmpty returns true if the tree is empty
func (t *concRedBlack[V]) IsEmpty() bool {
	t.mtx.RLock()
	defer t.mtx.RUnlock()
	return t.Size() == 0
}

// Iterator returns an iterator for the tree
// This iterator reflects live changes to the tree.
// Modifying the tree during iteration may cause undefined behavior.
func (t *concRedBlack[V]) Iterator() data.Iterator[V] {
	t.mtx.RLock()
	defer t.mtx.RUnlock()
	return concRedBlackIterator(t.snap())
}

// Max returns the highest value in the tree
func (t *concRedBlack[V]) Max() (V, bool) {
	t.mtx.RLock()
	defer t.mtx.RUnlock()
	if t.IsEmpty() {
		return generics.Zero[V](), false
	}
	return redblack.MaximumFrom(t.Root()).Value(), true
}

// Min returns the lowest value in the tree
func (t *concRedBlack[V]) Min() (V, bool) {
	t.mtx.RLock()
	defer t.mtx.RUnlock()
	if t.IsEmpty() {
		return generics.Zero[V](), false
	}
	return redblack.MinimumFrom(t.Root()).Value(), true
}

func (t *concRedBlack[V]) snap() []V {
	t.mtx.RLock()
	defer t.mtx.RUnlock()

	result := make([]V, 0, t.Size())
	curr := redblack.MinimumFrom(t.Root())
	for curr != nil {
		result = append(result, curr.Value())
		curr = successor(curr)
	}
	return result
}

func successor[V any](n *redblack.Node[V]) *redblack.Node[V] {
	if right := n.RightChild(); right != nil {
		return redblack.MinimumFrom(right)
	}
	parent := n.Parent()
	for parent != nil && n == parent.RightChild() {
		n = parent
		parent = parent.Parent()
	}
	return parent
}

func (t *concRedBlack[V]) String() string {
	t.mtx.RLock()
	defer t.mtx.RUnlock()
	tab := indent.Zero()
	root := "*redblack.Node<nil>"
	if t.Root() != nil {
		root = t.Root().Sprint(tab.Inc())
	}
	return sprints.Closedobj(tab, "*trees.redBlack",
		sprints.Fieldf(tab, "root", "%s", root),
	)
}

func concRedBlackIterator[V any](values []V) data.Iterator[V] {
	return &crbtIterator[V]{values: values}
}

type crbtIterator[V any] struct {
	mtx    sync.RWMutex
	values []V
	curr   int
}

func (i *crbtIterator[V]) Next() (next V) {
	if i.curr < len(i.values) {
		next = i.values[i.curr]
	}
	i.curr++
	return next
}

// HasNext Checks if there's a next element
func (i *crbtIterator[V]) HasNext() bool {
	return i.curr < len(i.values)
}
