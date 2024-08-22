// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package testsortables

import (
	"fmt"
	"github.com/andrerrcosta2/gtools/pkg/comparables"
	"github.com/andrerrcosta2/gtools/pkg/datastr/iterables"
	"github.com/andrerrcosta2/gtools/pkg/gtools"
)

// NewTestNode creates a new TestNode
// It implements the gtools.SortableOf interface
// It implements the comparables.Comparator interface
// It implements the fmt.Stringer interface
//
// It takes a string and returns a TestNode
func NewTestNode(s string) TestNode {
	return TestNode(s)
}

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

func NewTestNodeComparator() *TestNodeComparator {
	return &TestNodeComparator{}
}

// TestNodeComparator is a comparator for TestNode
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

var _ comparables.Comparator[TestNode] = (*TestNodeComparator)(nil)

// RandomTestNodes returns a slice of TestNode with the given prefix and length.
// Example: RandomTestNodes(10, "node")
//
// output: node_0, node_1, node_2, node_3, node_4, node_5, node_6, node_7, node_8, node_9
func RandomTestNodes(i int, prefix string) *iterables.Slice[TestNode] {
	out := make(iterables.Slice[TestNode], i)
	for j := 0; j < i; j++ {
		out[j] = TestNode(fmt.Sprintf("%s_%d", prefix, j))
	}
	return &out
}
