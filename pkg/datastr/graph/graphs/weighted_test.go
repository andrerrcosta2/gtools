// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package graphs

import (
	"github.com/andrerrcosta2/gtools/pkg/datastr/arrays"
	"github.com/andrerrcosta2/gtools/pkg/datastr/iterables"
	"github.com/andrerrcosta2/gtools/pkg/nums/graph"
	"github.com/andrerrcosta2/gtools/pkg/nums/progression"
	"github.com/andrerrcosta2/gtools/pkg/testdata/testsortables"
	"testing"
)

func TestAddingNodesAndEdges(t *testing.T) {
	g := WeightedOrderedOf[testsortables.TestNode, int]()

	var expEdges iterables.Slice[*graph.SingleTypedEdge[testsortables.TestNode]]
	weights, err := progression.SquareRoot(1_073_741_824, 2, 30)
	if err != nil {
		t.Error(err)
	}
	testsortables.RandomTestNodes(30, "node").
		Operation(func(i int, n *iterables.Slice[testsortables.TestNode]) {
			g.AddNode(n.At(i))
			// Left child
			if 2*i+1 < n.Len() {
				g.AddEdge(n.At(i), n.At(2*i+1), weights[i])
			}
			// Right child
			if 2*i+2 < n.Len() {
				g.AddEdge(n.At(i), n.At(2*i+2), weights[i])
			}
		})

	expEdges.Each(func(e *graph.SingleTypedEdge[testsortables.TestNode]) {
		if !g.HasEdge(e.From(), e.To()) {
			t.Errorf("Expected edge from %s to %s to exist in the graph", e.From(), e.To())
		}
	})
}

func TestNodeExistence(t *testing.T) {
	g := WeightedOrderedOf[testsortables.TestNode, int]()
	a := testsortables.TestNode("A")
	b := testsortables.TestNode("B")
	g.AddNode(a)
	g.AddNode(b)

	nodes := g.Nodes()
	exp := []testsortables.TestNode{a, b}

	if len(nodes) != len(exp) {
		t.Errorf("Nodes test failed. \nGot: %v, \nexpected: %v\n", nodes, exp)
		return
	}

	if !arrays.ContainsAllBy(&nodes, &exp, func(n1, n2 testsortables.TestNode) bool { return n1.Equal(n2) }) {
		t.Errorf("Nodes test failed. \nGot: %v, \nExpected: %v\n", nodes, exp)
	}
}

func TestEdgeExistenceAndWeight(t *testing.T) {
	g := WeightedOrderedOf[testsortables.TestNode, int]()
	a := testsortables.TestNode("A")
	b := testsortables.TestNode("B")
	c := testsortables.TestNode("C")
	g.AddNode(a)
	g.AddNode(b)
	g.AddNode(c)
	g.AddEdge(a, b, 10)
	g.AddEdge(b, c, 20)

	if !g.HasEdge(a, b) {
		t.Errorf("Edge A-B should exist.")
	}

	if weight, ok := g.Weight(a, b); !ok || weight != 10 {
		t.Errorf("Expected weight of edge A-B to be 10, got %v", weight)
	}

	if weight, ok := g.Weight(b, c); !ok || weight != 20 {
		t.Errorf("Expected weight of edge B-C to be 20, got %v", weight)
	}
}

func TestNeighborsRetrieval(t *testing.T) {
	g := WeightedOrderedOf[testsortables.TestNode, int]()
	a := testsortables.TestNode("A")
	b := testsortables.TestNode("B")
	c := testsortables.TestNode("C")
	g.AddNode(a)
	g.AddNode(b)
	g.AddNode(c)
	g.AddEdge(a, b, 10)
	g.AddEdge(b, c, 20)

	neighbors := g.Neighbors(a)
	if !contains(neighbors, b) {
		t.Errorf("Expected neighbors of A to include B")
	}
}

func TestDisconnectedGraph(t *testing.T) {
	g := WeightedOrderedOf[testsortables.TestNode, int]()
	a := testsortables.TestNode("A")
	b := testsortables.TestNode("B")
	c := testsortables.TestNode("C")
	g.AddNode(a)
	g.AddNode(b)
	g.AddNode(c)
	g.AddEdge(a, b, 10)

	nodes := g.Nodes()
	if !contains(nodes, a) || !contains(nodes, b) || !contains(nodes, c) {
		t.Errorf("Nodes test failed. Got: %v", nodes)
	}

	if g.HasEdge(b, c) {
		t.Errorf("Graph should not have an edge from B to C.")
	}
}

// Utility function to check if a slice contains a specific element
func contains(slice []testsortables.TestNode, item testsortables.TestNode) bool {
	for _, elem := range slice {
		if elem.Equal(item) {
			return true
		}
	}
	return false
}
