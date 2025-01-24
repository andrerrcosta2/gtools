// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package tests

import "github.com/andrerrcosta2/gtools/core/data/str/iterables"

func NewSortableNode(value string) *SortableNode {
	return &SortableNode{Value: value}
}

func NewSortableNodes(values ...string) *iterables.Slice[*SortableNode] {
	nodes := make(iterables.Slice[*SortableNode], len(values))
	for i, value := range values {
		nodes[i] = NewSortableNode(value)
	}
	return &nodes
}

type SortableNode struct {
	Value string
}

func (n *SortableNode) Unique() string {
	return n.Value
}

func (n *SortableNode) Less(o any) bool {
	if other, ok := o.(*SortableNode); ok {
		return n.Value < other.Value
	}
	return false
}

func (n *SortableNode) Equal(o any) bool {
	if other, ok := o.(*SortableNode); ok {
		return n.Value == other.Value
	}
	return false
}

func (n *SortableNode) String() string {
	return n.Value
}
