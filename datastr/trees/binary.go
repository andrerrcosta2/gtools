// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package trees

import (
	"github.com/andrerrcosta2/gtools/core/data"
	"github.com/andrerrcosta2/gtools/core/data/comparators"
	"github.com/andrerrcosta2/gtools/core/data/str"
	"github.com/andrerrcosta2/gtools/datastr/trees/internal/binary"
)

func BinarySearch[V any](cmp comparators.Typed[V]) str.BinaryTree[V] {
	return &binTree[V]{
		cmp: cmp,
	}
}

type binTree[V any] struct {
	abstract[*binary.Node[V], V]
	cmp comparators.Typed[V]
}

func (t *binTree[V]) String() string {
	//TODO implement me
	panic("implement me")
}

func (t *binTree[V]) Clear() {
	t.SetSize(0)
	t.SetRoot(nil)
}

func (t *binTree[V]) Contains(value V) bool {
	return binary.FindNode(t.Root(), value, t.cmp.Compare) != nil
}

func (t *binTree[V]) Delete(value V) {
	n := binary.FindNode(t.Root(), value, t.cmp.Compare)
	if n == nil {
		return
	}
	if binary.DeleteNode(t.Root(), n) {
		t.SetSize(t.Size() - 1)
	}
}

func (t *binTree[V]) Insert(value V) {
	if binary.InsertNode(t.Root(), binary.NewNode(value), t.cmp.Compare) {
		t.SetSize(t.Size() + 1)
	}
}

func (t *binTree[V]) IsEmpty() bool {
	return t.Size() == 0
}

func (t *binTree[V]) Min() (V, bool) {
	if t.IsEmpty() {
		var zero V
		return zero, false
	}

	current := t.Root()
	for current.LeftChild() != nil {
		current = current.LeftChild()
	}
	return current.Value(), true
}

func (t *binTree[V]) Max() (V, bool) {
	if t.IsEmpty() {
		var zero V
		return zero, false
	}

	current := t.Root()
	for current.RightChild() != nil {
		current = current.RightChild()
	}
	return current.Value(), true
}

func (t *binTree[V]) Floor(value V) (V, bool) {
	var result *binary.Node[V]
	current := t.Root()

	for current != nil {
		comp := t.cmp.Compare(value, current.Value())
		if comp == 0 {
			return current.Value(), true
		} else if comp < 0 {
			current = current.LeftChild()
		} else {
			result = current
			current = current.RightChild()
		}
	}

	if result == nil {
		var zero V
		return zero, false
	}
	return result.Value(), true
}

func (t *binTree[V]) Ceiling(value V) (V, bool) {
	var result *binary.Node[V]
	current := t.Root()

	for current != nil {
		comp := t.cmp.Compare(value, current.Value())
		if comp == 0 {
			return current.Value(), true
		} else if comp > 0 {
			current = current.RightChild()
		} else {
			result = current
			current = current.LeftChild()
		}
	}

	if result == nil {
		var zero V
		return zero, false
	}
	return result.Value(), true
}

func (t *binTree[V]) Iterator() data.Iterator[V] {
	//TODO implement me
	panic("implement me")
}
