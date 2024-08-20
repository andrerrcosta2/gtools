// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package testdata

import "github.com/andrerrcosta2/gtools/pkg/gtools"

type TestNode string

func (n TestNode) Equal(other interface{}) bool {
	if otherNode, ok := other.(TestNode); ok {
		return n == otherNode
	}
	return false
}

func (n TestNode) Less(other interface{}) bool {
	return n < other.(TestNode)
}

var _ gtools.SortableOf = (*TestNode)(nil)

type TestNodeComparator struct{}

func (t *TestNodeComparator) Compare(a, b TestNode) int {
	if a.Less(b) {
		return -1
	} else if b.Less(a) {
		return 1
	}
	return 0
}

func (t *TestNodeComparator) Equals(a, b TestNode) bool {
	return a.Equal(b)
}

func NewTestNodeComparator() *TestNodeComparator {
	return &TestNodeComparator{}
}
