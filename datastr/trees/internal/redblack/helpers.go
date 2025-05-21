// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package redblack

import (
	"github.com/andrerrcosta2/gtools/core/data/str/nodes"
	"github.com/andrerrcosta2/gtools/core/domain/functions"
	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"github.com/andrerrcosta2/gtools/core/format/sprints"
)

func NewNode[V any](value V, color nodes.RBColor, parentLeftRight ...*Node[V]) *Node[V] {
	n := &Node[V]{
		value: value,
		color: color,
	}
	if len(parentLeftRight) > 0 && parentLeftRight[0] != nil {
		n.parent = parentLeftRight[0]
	}
	if len(parentLeftRight) > 1 && parentLeftRight[1] != nil {
		n.left = parentLeftRight[1]
	}
	if len(parentLeftRight) > 2 && parentLeftRight[2] != nil {
		n.right = parentLeftRight[2]
	}
	return n
}

func Black[V any](value V, parentLeftRight ...*Node[V]) *Node[V] {
	return NewNode(value, nodes.Black, parentLeftRight...)
}

func Red[V any](value V, parentLeftRight ...*Node[V]) *Node[V] {
	return NewNode(value, nodes.Red, parentLeftRight...)
}

func Root[V any](value V, LeftRight ...*Node[V]) *Node[V] {
	n := &Node[V]{
		value: value,
		color: nodes.Black,
	}
	if len(LeftRight) > 0 {
		n.left = LeftRight[0]
	}
	if len(LeftRight) > 1 {
		n.right = LeftRight[1]
	}
	return n
}

type Node[V any] struct {
	value  V
	color  nodes.RBColor
	left   *Node[V]
	right  *Node[V]
	parent *Node[V]
}

func (n *Node[V]) Color() nodes.RBColor {
	return n.color
}

func (n *Node[V]) HasChildren() bool {
	return n.left != nil || n.right != nil
}

func (n *Node[V]) LeftChild() *Node[V] {
	return n.left
}

func (n *Node[V]) Parent() *Node[V] {
	return n.parent
}

func (n *Node[V]) RightChild() *Node[V] {
	return n.right
}

func (n *Node[V]) SetColor(color nodes.RBColor) {
	n.color = color
}

func (n *Node[V]) SetLeftChild(left *Node[V]) {
	n.left = left
}

func (n *Node[V]) SetParent(parent *Node[V]) {
	n.parent = parent
}

func (n *Node[V]) SetRightChild(right *Node[V]) {
	n.right = right
}

func (n *Node[V]) SetValue(value V) {
	n.value = value
}

func (n *Node[V]) Value() V {
	return n.value
}

func (n *Node[V]) String() string {
	return n.Sprint(indent.Zero())
}

func (n *Node[V]) Sprint(tab indent.Tab) string {
	var parent = "*redblack.Node<nil>"
	var left = "*redblack.Node<nil>"
	var right = "*redblack.Node<nil>"
	if n.parent != nil {
		parent = sprints.Addr(n.parent)
	}
	if n.left != nil {
		left = n.left.Sprint(tab.Inc())
	}
	if n.right != nil {
		right = n.right.Sprint(tab.Inc())
	}
	return sprints.Closedobj(tab, "*redblack.Node",
		sprints.Fieldf(tab, "value", "%v", n.value),
		sprints.Fieldf(tab, "color", "%s", n.color),
		sprints.Fieldf(tab, "parent", "%s", parent),
		sprints.Fieldf(tab, "left", "%s", left),
		sprints.Fieldf(tab, "right", "%s", right),
	)
}

// Ceiling returns the ceiling of a value in the tree.
func Ceiling[V any](from *Node[V], value V, compare functions.BiFunction[V, V, int]) (V, bool) {
	var ceil V
	var found bool

	for from != nil {
		cmp := compare(value, from.Value())
		if cmp == 0 { // same value
			return from.Value(), true
		} else if cmp < 0 {
			ceil = from.Value()
			found = true
			from = from.LeftChild()
		} else {
			from = from.RightChild()
		}
	}

	return ceil, found
}

// colorOf returns the color of a node. If the node is nil, it returns Black
func colorOf[V any](n *Node[V]) nodes.RBColor {
	if n == nil {
		return nodes.Black
	}
	return n.color
}

// DeleteNode deletes a node from the tree. The deletion is one of the
// most complex operations on red black trees because it involves not only removing a
// node but also ensuring that the tree maintains its properties after the removal.
//
// the deletion must be done in three main steps
//
//  1. find the node to delete. (this was performed on the previous step)
//  2. delete the node:
//     - if the node has no children, just remove it
//     - if the node has one child, replace it with the child
//     - If the node has two children, find its in-order successor (the smallest
//     node in its right subtree). Replace the node’s value with the successor’s
//     value, then delete the successor:
//  3. fix up the red-black properties:
//     - after removing the node, the tree may violate Red-Black properties.
//     to avoid that we must perform "fixup" operations (rotations and recoloring)
//     to restore the properties.
func DeleteNode[V any](root, target *Node[V]) *Node[V] {
	if target == nil {
		return root
	}
	var parent, fixup *Node[V]
	// Case 1: target has no left child → replace it with targetRight
	// (if right is also nil it has no children)
	if target.left == nil {
		root, parent, fixup = halfTransplantRight(root, target)
		// After removing the node, the tree might be violating the red-black properties
		// 1. If the removed node was black, the black height property may be violated.
		// 2. If the removed node was red, no balancing is needed because removing a red node
		// does not affect the black height.
		// 3. if the replacement was a red node and the tree was balanced before,
		// the height can be fixed painting the replacement black
		if fixup != root && target.color == nodes.Black {
			if isRed(fixup) { // quick fixup
				fixup.color = nodes.Black
			} else { // long fixup
				root = deleteFixup(root, parent, fixup == parent.left)
			}
		}
		// Case 2: target has no right child → replace it with targetLeft
	} else if target.right == nil {
		root, parent, fixup = halfTransplantLeft(root, target)
		if fixup != root && target.color == nodes.Black {
			if isRed(fixup) { // quick fixup
				fixup.color = nodes.Black
			} else { // long fixup
				root = deleteFixup(root, parent, fixup == parent.left)
			}
		}
	} else {
		// Case 3: If the node has two children, find its in-order successor (the smallest node in
		// its right subtree). Replace the node’s value with the successor’s value,
		// then delete the successor:
		root, parent, fixup = fullTransplantSuccessor(root, target)
		if fixup != root {
			if isRed(fixup) {
				fixup.color = nodes.Black
			} else {
				root = deleteFixup(root, parent, fixup == parent.left)
			}
		}
	}
	return root
}

func deleteFixup[V any](root, parent *Node[V], left bool) *Node[V] {
	for parent != nil {
		if left {
			sibl := parent.right
			if isRed(sibl) {
				sibl.color = nodes.Black
				parent.color = nodes.Red
				rotateLeft(root, parent)
				// the sibling is the only one can replace the root
				if sibl.parent == nil {
					root = sibl
				}
				sibl = parent.right
			}
			if isNilOrHasOnlyBlackChildren(sibl) {
				if sibl != nil {
					sibl.color = nodes.Red
				}
				// if the parent is red, the loop should break,
				// paint it black and return the root
				if parent.color == nodes.Red {
					parent.color = nodes.Black
					if parent.parent == nil {
						return parent
					}
					return root
				}
				if parent.parent == nil {
					return parent
				}
				left = parent == parent.parent.left
				parent = parent.parent
			} else {
				if isNilOrBlack(sibl.right) { // sibl.left is red
					sibl.left.color = nodes.Black
					sibl.color = nodes.Red
					rotateRight(root, sibl)
					sibl = parent.right
				}
				sibl.color = parent.color
				parent.color = nodes.Black
				sibl.right.color = nodes.Black
				rotateLeft(root, parent)
				if sibl.parent == nil {
					root = sibl
				}
				break
			}
			// right fixup
		} else {
			sibl := parent.left
			if isRed(sibl) {
				sibl.color = nodes.Black
				parent.color = nodes.Red
				rotateRight(root, parent)
				// the sibling is the only one can replace the root
				if sibl.parent == nil {
					root = sibl
				}
				sibl = parent.left
			}
			if isNilOrHasOnlyBlackChildren(sibl) {
				if sibl != nil {
					sibl.color = nodes.Red
				}
				if parent.color == nodes.Red {
					parent.color = nodes.Black
					if parent.parent == nil {
						return parent
					}
					return root
				}
				if parent.parent == nil {
					return parent
				}
				left = parent == parent.parent.left
				parent = parent.parent
			} else {
				if isNilOrBlack(sibl.left) {
					sibl.right.color = nodes.Black
					sibl.color = nodes.Red
					rotateLeft(root, sibl)
					sibl = parent.left
				}
				sibl.color = parent.color
				parent.color = nodes.Black
				sibl.left.color = nodes.Black
				rotateRight(root, parent)
				if sibl.parent == nil {
					root = sibl
				}
				break
			}
		}
	}

	root.color = nodes.Black
	return root
}

// Floor returns the floor of the given value
func Floor[V any](from *Node[V], value V, compare functions.BiFunction[V, V, int]) (V, bool) {
	var floor V
	var found bool

	for from != nil {
		cmp := compare(value, from.Value())
		if cmp == 0 { // equals
			return from.Value(), true
		} else if cmp > 0 {
			floor = from.Value()
			found = true
			from = from.RightChild()
		} else {
			from = from.LeftChild()
		}
	}

	return floor, found
}

func hasOnlyBlackChildren[V any](n *Node[V]) bool {
	return (n.left == nil || n.left.color == nodes.Black) &&
		(n.right == nil || n.right.color == nodes.Black)
}

func isRed[V any](n *Node[V]) bool {
	return !isNilOrBlack(n)
}

func isNilOrBlack[V any](n *Node[V]) bool { return n == nil || n.color == nodes.Black }

func isNilOrHasOnlyBlackChildren[V any](n *Node[V]) bool {
	return n == nil || hasOnlyBlackChildren(n)
}

func FindNode[V any](root *Node[V], value V, cmp functions.BiFunction[V, V, int]) (*Node[V], bool) {
	current := root
	for current != nil {
		c := cmp(value, current.value)
		if c == 0 {
			return current, true
		} else if c < 0 {
			current = current.left
		} else {
			current = current.right
		}
	}
	return nil, false
}

func InsertNode[V any](root, node *Node[V], compare functions.BiFunction[V, V, int]) (*Node[V], bool) {
	// Find the correct position for the new node
	var parent *Node[V]
	var current = root
	var i int

	for current != nil {
		parent = current
		i = compare(node.value, current.value)
		if i < 0 {
			current = current.left
		} else if i > 0 {
			current = current.right
		} else {
			return current, false // don't insert repeated values
		}
	}

	// Attach the new node to its parent
	node.parent = parent
	if parent == nil {
		return node, true // New node becomes the root
	}
	if i < 0 {
		parent.left = node
	} else {
		parent.right = node
	}

	// assert the node starts out red
	node.color = nodes.Red
	// Fix the Red-Black Tree properties
	return insertFixup(root, node), true
}

// insertFixup restores the trees.RedBlack Tree properties after insertion.
func insertFixup[V any](root, n *Node[V]) *Node[V] {
	// When inserting a new node into a Red-Black Tree, it is always initially red.
	// However, this might cause two consecutive red nodes (violating Property 4).
	// The insertFixup function restores the tree's balance by checking if the parent
	// is red and applying rotations and recoloring when necessary.
	// This algorithm only runs if the node was inserted as a child of a red node.
	for n.parent != nil && n.parent.color == nodes.Red {
		// When a parent is red, there are three cases to consider:
		grandparent := n.parent.parent
		if n.parent == grandparent.left {
			uncle := grandparent.right

			// Case 1: When the uncle node is red, the algorithm performs recoloring.
			// assigning the grandparent to n we are propagating the issue upwards because
			// the grandparent might now violate the Red-Black property.
			if uncle != nil && uncle.color == nodes.Red {
				n.parent.color = nodes.Black
				uncle.color = nodes.Black
				grandparent.color = nodes.Red
				n = grandparent
			} else {
				// Case 2: Rotation for Zigzag (Triangle Shape)
				// When the uncle node is black, the algorithm performs rotations
				// If the node is the right child of the parent, it must first ensure the primary
				// property of binary search trees: the left child of a node is less
				// than the parent and the right child is greater than the parent.
				// So we flat it out by making the node the parent and then rotating it to the left.
				// The goal is to transform the tree into a linear structure (a straight line of nodes)
				// so that a single right rotation and recoloring can resolve the violation.
				// If n is the right child of its parent, a left rotation on the parent is needed to
				// make the nodes line up.
				// If n is the left child of its parent, no left rotation is needed.
				// If the parent is the left child of the grandparent and n is the right child of its parent
				// that means we are handling a triangle shaped tree (ofter referred to as a zigzag pattern)
				//       (GP)
				//      /
				//    (P[RED])
				//      \
				//      (n[RED])
				if n == n.parent.right {
					n = n.parent
					// The rotateLeft function, will change the root of the tree, if the node x, is the root.
					// After the rotation, the node that was the parent, becomes the node n.
					root = rotateLeft(root, n)
					// After the rotation, the node that was the parent, becomes the node n.
					//       (GP)
					//      /
					//    (n[RED])
					//    /
					//  (P[RED])
				}
				// Case 3: Right Rotation
				// Now that the structure is properly aligned (straight line),
				// we recolor and perform a final right rotation.
				//     (n[RED])
				//    /     \
				//  (P[RED]) (GP)
				n.parent.color = nodes.Black
				grandparent.color = nodes.Red
				//     (n[BLACK])
				//    /     \
				//  (P[RED]) (GP[RED])
				root = rotateRight(root, grandparent)
			}
		} else {
			// Mirror image of the previous case (if the parent is the right child of grandparent).
			uncle := grandparent.left

			// Case 1: Recoloring
			if uncle != nil && uncle.color == nodes.Red {
				n.parent.color = nodes.Black
				uncle.color = nodes.Black
				grandparent.color = nodes.Red
				n = grandparent
			} else {
				// Case 2: Rotation for Zigzag (Triangle Shape)
				if n == n.parent.left {
					n = n.parent
					root = rotateRight(root, n)
				}

				// Case 3: Left Rotation
				n.parent.color = nodes.Black
				grandparent.color = nodes.Red
				root = rotateLeft(root, grandparent)
			}
		}
	}
	root.color = nodes.Black // Ensure the root is always black
	return root
}

func MaximumFrom[V any](n *Node[V]) *Node[V] {
	if n == nil {
		return nil
	}
	for n.right != nil {
		n = n.right
	}
	return n
}

func MinimumFrom[V any](n *Node[V]) *Node[V] {
	if n == nil {
		return nil
	}
	for n.left != nil {
		n = n.left
	}
	return n
}

// rotateLeft performs a left rotation on the given node.
func rotateLeft[V any](root, x *Node[V]) *Node[V] {
	// If the node being rotated is nil or has no right child, return the root as is
	if x == nil || x.right == nil {
		return root
	}

	// top is the right child of x, which will take x's place
	top := x.right

	// Update child pointers
	x.right = top.left
	if x.right != nil {
		x.right.parent = x
	}

	// Update parent pointers
	top.parent = x.parent
	if x.parent == nil {
		root = top // top becomes the new root
	} else if x == x.parent.left {
		x.parent.left = top
	} else {
		x.parent.right = top
	}

	// Finalize the rotation
	top.left = x
	x.parent = top

	// always set the root color to black
	root.color = nodes.Black
	return root
}

// rotateRight performs a right rotation on the given node.
func rotateRight[V any](root, x *Node[V]) *Node[V] {
	// If the node being rotated is nil or has no left child, return the root as is
	if x == nil || x.left == nil {
		return root
	}

	// top is the left child of x, which will take x's place
	top := x.left

	// Update child pointers
	x.left = top.right
	if x.left != nil {
		x.left.parent = x
	}

	// Update parent pointers
	top.parent = x.parent
	if x.parent == nil {
		root = top // top becomes the new root
	} else if x == x.parent.left {
		x.parent.left = top
	} else {
		x.parent.right = top
	}

	// Finalize the rotation
	top.right = x
	x.parent = top

	// always set the root color to black
	root.color = nodes.Black
	return root
}

// halfTransplantLeft is a standard transplant that doesn't modify the left/right children
// it receives a node and change its parent to point to its left child.
func halfTransplantLeft[V any](root, target *Node[V]) (newRoot, parent, x *Node[V]) {
	repl := target.left
	if target.parent == nil {
		if repl == nil {
			// if the target is the root and has no left child
			return nil, nil, nil
		}
		repl.parent = nil
		repl.color = nodes.Black
		return repl, nil, repl
	}
	if target == target.parent.left {
		target.parent.left = repl
		if repl != nil {
			repl.parent = target.parent
		}
	} else {
		target.parent.right = repl
		if repl != nil {
			repl.parent = target.parent
		}
	}
	return root, target.parent, repl
}

// halfTransplantRight is a standard transplant that doesn't modify the left/right children
// it receives a node and change its parent to point to its right child.
func halfTransplantRight[V any](root, target *Node[V]) (newRoot, parent, x *Node[V]) {
	repl := target.right
	if target.parent == nil {
		if repl == nil {
			// if the target is the root and has no right child
			return nil, nil, nil
		}
		repl.parent = nil
		repl.color = nodes.Black
		return repl, nil, repl
	}
	if target == target.parent.left {
		target.parent.left = repl
		if repl != nil {
			repl.parent = target.parent
		}
	} else {
		target.parent.right = repl
		if repl != nil {
			repl.parent = target.parent
		}
	}
	return root, target.parent, repl
}

// fullTransplant is a transplant that modifies the whole structure considering
// the target has both left and right children. It means the left and right children
// aren't nil
func fullTransplantLeft[V any](root, target *Node[V]) (newRoot, parent, x *Node[V]) {
	repl := target.left
	if target.parent == nil {
		repl.parent = nil
		repl.color = nodes.Black
		newRoot = repl
	} else {
		newRoot = root
		repl.color = target.color
		repl.parent = target.parent
		if target == target.parent.left {
			target.parent.left = repl
		} else {
			target.parent.right = repl
		}
	}
	repl.right = target.right
	target.right.parent = repl
	return newRoot, target.parent, repl
}

// fullTransplant is a transplant that modifies the whole structure considering
// the target has both left and right children. It means the left and right children
// aren't nil
func fullTransplantRight[V any](root, target *Node[V]) (newRoot, parent, x *Node[V]) {
	repl := target.right
	if target.parent == nil {
		repl.parent = nil
		repl.color = nodes.Black
		newRoot = repl
	} else {
		newRoot = root
		repl.color = target.color
		repl.parent = target.parent
		if target == target.parent.left {
			target.parent.left = repl
		} else {
			target.parent.right = repl
		}
	}
	repl.left = target.left
	target.left.parent = repl
	return newRoot, target.parent, repl
}

func fullTransplantSuccessor[V any](root, target *Node[V]) (newRoot, parent, x *Node[V]) {
	repl := MinimumFrom(target.right)
	if repl.parent != target {
		// repl.right replaces repl position. in this case
		// the parent of the transplanted always exist
		// if the replacement is the direct child of the node to be deleted
		// we need to give its child to the right child of the node to be deleted:
		//        A (target)
		//         \
		//          B
		//         /
		//        C (successor)
		//         \
		//          D
		_, parent, x = halfTransplantRight(root, repl)
		// The tree becomes:
		//         A (target)
		//         \
		//          B
		//         /
		//        D  <-- `C` is now floating and ready to replace A.
		if target.parent != nil {
			repl.color = target.color
			target.parent.right = repl
			newRoot = root
		} else {
			repl.parent = nil
			newRoot = repl
		}
		newRoot.color = nodes.Black
		repl.right = target.right
		repl.right.parent = repl
		repl.left = target.left
		repl.left.parent = repl
		// The tree becomes:
		//        C (replacement)  <-- `C` now replaces `A`
		//       /  \
		//      x   B (fixup parent)
		//          /
		//         D (fixup)
		return newRoot, parent, x
	}
	return fullTransplantRight(root, repl.parent)
}
