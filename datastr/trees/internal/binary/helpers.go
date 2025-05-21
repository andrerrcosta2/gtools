// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package binary

import "github.com/andrerrcosta2/gtools/core/domain/functions"

func NewNode[V any](value V) *Node[V] {
	return &Node[V]{
		value: value,
	}
}

type Node[V any] struct {
	value  V
	parent *Node[V]
	left   *Node[V]
	right  *Node[V]
}

func (n *Node[V]) Parent() *Node[V] {
	return n.parent
}

func (n *Node[V]) SetParent(p *Node[V]) {
	n.parent = p
}

func (n *Node[V]) LeftChild() *Node[V] {
	return n.left
}

func (n *Node[V]) RightChild() *Node[V] {
	return n.right
}

func (n *Node[V]) SetLeftChild(c *Node[V]) {
	n.left = c
}

func (n *Node[V]) SetRightChild(c *Node[V]) {
	n.right = c
}

func (n *Node[V]) Value() V {
	return n.value
}

func DeleteNode[V any](root, n *Node[V]) bool {
	var parent *Node[V]
	current := root

	// Case 1: Node has no children (leaf node)
	if current.left == nil && current.right == nil {
		replaceNode(root, parent, current, nil)
	} else if current.left == nil { // Case 2: Node has only right child
		replaceNode(root, parent, current, current.right)
	} else if current.right == nil { // Case 2: Node has only left child
		replaceNode(root, parent, current, current.left)
	} else { // Case 3: Node has two children
		successorParent := current
		successor := current.right

		// Find the in-order successor (smallest in right subtree)
		for successor.left != nil {
			successorParent = successor
			successor = successor.left
		}

		// Replace current value with successor value
		current.value = successor.value

		// Remove the successor node
		if successorParent.left == successor {
			successorParent.left = successor.right
		} else {
			successorParent.right = successor.right
		}
	}
	return true
}

func FindNode[V any](root *Node[V], value V, compare functions.BiFunction[V, V, int]) *Node[V] {
	current := root
	for current != nil {
		cmp := compare(value, current.value)
		if cmp == 0 {
			return current
		} else if cmp < 0 {
			current = current.left
		} else {
			current = current.right
		}
	}
	return nil
}

func InsertNode[V any](root, n *Node[V], compare functions.BiFunction[V, V, int]) bool {
	current := root
	for {
		cmp := compare(n.value, current.value)
		if cmp < 0 {
			if current.left == nil {
				current.left = n
				return true
			}
			current = current.left
		} else if cmp > 0 {
			if current.right == nil {
				current.right = n
				return true
			}
			current = current.right
		} else {
			// ignoring duplicated values
			return false
		}
	}
}

// Helper function to replace a node
func replaceNode[V any](root, parent, target, replacement *Node[V]) {
	if parent == nil {
		root = replacement
	} else if parent.left == target {
		parent.left = replacement
	} else {
		parent.right = replacement
	}
}
