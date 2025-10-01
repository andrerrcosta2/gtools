// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package models

import (
	"fmt"
	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"github.com/andrerrcosta2/gtools/core/format/fmx"
	"github.com/andrerrcosta2/gtools/core/seeders/random"
)

var ZeroBinaryNodeInst = new(BinaryNode)

func BinaryNodeAsValue(value any, left, right *BinaryNode) BinaryNode {
	return BinaryNode{
		Value: value,
		Left:  left,
		Right: right,
	}
}

func BinaryNodeAsRef(value any, left, right *BinaryNode) *BinaryNode {
	return &BinaryNode{
		Value: value,
		Left:  left,
		Right: right,
	}
}

func BinaryNodeAsRandValue() BinaryNode {
	// Generate random depth for left and right branches
	depths := random.Int(2, 1, 5)

	node := BinaryNode{
		Value: random.SingleOf[any](),
	}

	return *randBinaryNodeAsRef(&node, depths.At(0), depths.At(1))
}

func BinaryNodeAsRandRef() *BinaryNode {
	// Generate random depth for left and right branches
	depths := random.Int(2, 1, 5)

	node := &BinaryNode{
		Value: random.SingleOf[any](),
	}

	return randBinaryNodeAsRef(node, depths.At(0), depths.At(1))
}

func randBinaryNodeAsRef(node *BinaryNode, leftDepth, rightDepth int) *BinaryNode {
	if leftDepth > 0 {
		node.Left = &BinaryNode{
			Value: random.SingleOf[any](),
		}
		randBinaryNodeAsRef(node.Left, leftDepth-1, random.Int(1, 0, rightDepth-1).At(0))
	}

	if rightDepth > 0 {
		node.Right = &BinaryNode{
			Value: random.SingleOf[any](),
		}
		randBinaryNodeAsRef(node.Right, random.Int(1, 0, leftDepth-1).At(0), rightDepth-1)
	}

	return node
}

type BinaryNode struct {
	Left  *BinaryNode
	Right *BinaryNode
	Value any
}

func (n *BinaryNode) Sprint(tab indent.Indentor, visited map[*BinaryNode]bool) string {
	if n == nil {
		return "<nil>"
	}
	if visited[n] {
		return fmx.Sprintf("...<cycle detected at %p>...", n)
	}
	visited[n] = true

	var left, right = "<nil>", "<nil>"
	if n.Left != nil {
		left = n.Left.Sprint(tab.Inc(), visited)
	}
	if n.Right != nil {
		right = n.Right.Sprint(tab.Inc(), visited)
	}

	return fmx.Sprintf("\n%sValue: %v,\n%sLeft: %s,\n%sRight: %s\n",
		tab.String(), n.Value, tab.String(), left, tab.String(), right)
}

func (n *BinaryNode) String() string {
	visited := make(map[*BinaryNode]bool)
	return "BinaryNode{" + n.Sprint(indent.Tab(1), visited) + "}"
}

func LinkedListAsRandValue() LinkedList {
	length := random.Int(1, 1, 10).At(0)

	head := &BinaryNode{
		Value: random.SingleOf[any](),
	}
	current := head

	for i := 1; i < length; i++ {
		next := &BinaryNode{
			Value: random.SingleOf[any](),
		}
		current.Right = next
		current = next
	}

	return LinkedList{
		Head: head,
		Tail: current,
	}
}

func LinkedListAsRandRef() *LinkedList {
	length := random.Int(1, 1, 10).At(0) // Random length between 1 and 10

	head := &BinaryNode{
		Value: random.SingleOf[any](),
	}
	current := head

	for i := 1; i < length; i++ {
		next := &BinaryNode{
			Value: random.SingleOf[any](),
		}
		current.Right = next // Use Right as the "next" pointer in the linked list
		current = next
	}

	return &LinkedList{
		Head: head,
		Tail: current,
	}
}

// CyclicLinkedList creates a double-linked list with a cycle
func CyclicLinkedList(values ...any) *LinkedList {
	list := &LinkedList{}
	var previous *BinaryNode
	for _, value := range values {
		node := &BinaryNode{Value: value}
		if list.Head == nil {
			list.Head = node
		}
		if previous != nil {
			previous.Right = node
			node.Left = previous
		}
		previous = node
	}

	list.Tail = previous
	list.Tail.Right = list.Head
	list.Head.Left = list.Tail
	return list
}

type LinkedList struct {
	Head *BinaryNode
	Tail *BinaryNode
}

func (l *LinkedList) Add(value any) {
	newNode := &BinaryNode{Value: value}

	if l.Tail != nil {
		l.Tail.Right = nil
		l.Head.Left = nil
	}
	if l.Tail != nil {
		l.Tail.Right = newNode
		newNode.Left = l.Tail
	} else {
		l.Head = newNode
	}
	l.Tail = newNode
}

func (l *LinkedList) Remove(value any) {
	if l.Head == nil {
		return
	}

	current := l.Head
	for current != nil {
		if current.Value == value {
			if current == l.Head && current == l.Tail {
				l.Head, l.Tail = nil, nil
				return
			}
			if current == l.Head {
				l.Head = l.Head.Right
				l.Head.Left = nil
			}
			if current == l.Tail {
				l.Tail = l.Tail.Left
				l.Tail.Right = nil
			}
			if current.Left != nil {
				current.Left.Right = current.Right
			}
			if current.Right != nil {
				current.Right.Left = current.Left
			}
			return
		}
		current = current.Right
	}
}

func NTreeNodeAsRandValue() NTreeNode {
	rnd := random.Int(2, 1, 5)
	return *randNTreeNodeAtDepth(0, rnd.At(0), rnd.At(1))
}

func NTreeNodeAsRandRef() *NTreeNode {
	rnd := random.Int(2, 1, 5)
	return randNTreeNodeAtDepth(0, rnd.At(0), rnd.At(1))
}

func randNTreeNodeAtDepth(currentDepth, maxDepth, maxChildren int) *NTreeNode {
	if currentDepth >= maxDepth {
		return &NTreeNode{Nodes: nil}
	}

	numChildren := random.Int(1, 0, maxChildren).At(0)
	nodes := make(map[string]*NTreeNode, numChildren)

	for i := 0; i < numChildren; i++ {
		key := random.Alphanumeric(1, 3, 10).At(0)
		nodes[key] = randNTreeNodeAtDepth(currentDepth+1, maxDepth, maxChildren)
	}

	return &NTreeNode{
		Nodes: nodes,
	}
}

type NTreeNode struct {
	Nodes map[string]*NTreeNode
}

func SimpleNodeAsRandValue() SimpleNode {
	length := random.Int(1, 2, 10).At(0)

	head := &SimpleNode{
		Value: random.SingleOf[any](),
	}
	current := head

	for i := 1; i < length; i++ {
		current.Next = &SimpleNode{
			Value: random.SingleOf[any](),
		}
		current = current.Next
	}

	return *head
}

func SimpleNodeAsRandRef() *SimpleNode {
	length := random.Int(1, 2, 10).At(0)

	head := &SimpleNode{
		Value: random.SingleOf[any](),
	}
	current := head

	for i := 1; i < length; i++ {
		current.Next = &SimpleNode{
			Value: random.SingleOf[any](),
		}
		current = current.Next
	}

	return head
}

// CyclicSimpleNode creates a linked list with a cycle
func CyclicSimpleNode(a, b any) *SimpleNode {
	n1 := &SimpleNode{
		Value: a,
	}
	n2 := &SimpleNode{
		Next:  n1,
		Value: b,
	}
	n1.Next = n2
	return n1
}

type SimpleNode struct {
	Next  *SimpleNode
	Value any
}

func (n *SimpleNode) String() string {
	visited := make(map[*SimpleNode]bool)
	return n.stringHelper(visited)
}

func (n *SimpleNode) stringHelper(visited map[*SimpleNode]bool) string {
	if n == nil {
		return "<nil>"
	}

	if visited[n] {
		return "[cyclic reference]"
	}

	visited[n] = true
	nextStr := n.Next.stringHelper(visited)

	return fmt.Sprintf("SimpleNode{\n\tValue: %v,\n\tNext: %s\n}", n.Value, nextStr)
}
