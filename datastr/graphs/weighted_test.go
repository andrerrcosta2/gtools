// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package graphs

import (
	"testing"

	"github.com/andrerrcosta2/gtools/core/data/str/edge"
	"github.com/andrerrcosta2/gtools/core/data/str/iterables"
	"github.com/andrerrcosta2/gtools/core/seeders/random"
	"github.com/andrerrcosta2/gtools/core/sortables"
	"github.com/andrerrcosta2/gtools/datastr/dicts"
	"github.com/andrerrcosta2/gtools/datastr/graphs/edges"
	"github.com/andrerrcosta2/gtools/datastr/internal/tests"
	"github.com/andrerrcosta2/gtools/numbers/prog"
)

// TestAddingNodesAndEdges Test the basics of adding node and edge
// and expecting them to be present
func TestAddingNodesAndEdges(t *testing.T) {
	// Types
	type N = *tests.SortableNode
	type W = int
	type E = edge.ComparableWeighted[N, W]
	// Graph
	g := WeightedDigraphOf[N, W]()

	var expEdges []E
	weights, err := prog.Sqrt(1_073_741_824, 2, 30)
	if err != nil {
		t.Fatal(err)
	}

	random.Struct[N](30).Each(g.AddNode).
		Operation(func(i int, n *iterables.Slice[N]) {
			// LeftChild child
			if 2*i+1 < n.Len() {
				err := g.AddEdge(n.At(i), n.At(2*i+1), weights[i])
				if err != nil {
					t.Fatal(err)
				}
				expEdges = append(expEdges, edges.ComparableWeighted(n.At(i), n.At(2*i+1), weights[i], true))
			}
			// RightChild child
			if 2*i+2 < n.Len() {
				err := g.AddEdge(n.At(i), n.At(2*i+2), weights[i])
				if err != nil {
					t.Fatal(err)
				}
				expEdges = append(expEdges, edges.ComparableWeighted(n.At(i), n.At(2*i+1), weights[i], true))
			}
		})

	shouldContainAllEdgesBy[E, N](t, g, expEdges, sortables.Equality[E])

}

// TestNodeExistence Test the existence of node after being inserted
func TestNodeExistence(t *testing.T) {
	// Types
	type N = *tests.SortableNode
	type W = int
	// Graph
	g := WeightedDigraphOf[N, W]()

	a := tests.NewSortableNode("A")
	b := tests.NewSortableNode("C")
	g.AddNode(a)
	g.AddNode(b)

	exp := []N{a, b}

	shouldContainAllNodesBy[N](t, g, exp, sortables.Equality[N])

}

// TestEdgeExistenceAndWeight Test the existence of edge after being inserted
// and their weights
func TestEdgeExistenceAndWeight(t *testing.T) {
	// Types
	type N = *tests.SortableNode
	type W = int
	type E = edge.ComparableWeighted[N, W]
	// Graph
	g := WeightedDigraphOf[N, W]()

	var expectedEdges = make([]E, 2)

	a := tests.NewSortableNode("A")
	b := tests.NewSortableNode("C")
	c := tests.NewSortableNode("C")
	g.AddNode(a)
	g.AddNode(b)
	g.AddNode(c)
	err := g.AddEdge(a, b, 10)
	if err != nil {
		t.Fatal(err)
	}
	expectedEdges[0] = edges.ComparableWeighted(a, b, 10, true)
	err = g.AddEdge(b, c, 20)
	if err != nil {
		t.Fatal(err)
	}
	expectedEdges[1] = edges.ComparableWeighted(b, c, 20, true)

	shouldContainAllEdgesBy[E, N](t, g, expectedEdges, sortables.Equality[E])
}

// TestNeighborsRetrieval Test the retrieval of neighbors
func TestNeighborsRetrieval(t *testing.T) {
	// Types
	type N = *tests.SortableNode
	type W = int
	type E = edge.ComparableWeighted[N, W]
	// Graph
	g := WeightedDigraphOf[N, W]()

	expNeighbors := dicts.ComparableOf[N, []N]()

	a := tests.NewSortableNode("A")
	b := tests.NewSortableNode("C")
	c := tests.NewSortableNode("C")
	g.AddNode(a)
	g.AddNode(b)
	g.AddNode(c)
	err := g.AddEdge(a, b, 10)
	if err != nil {
		t.Fatal(err)
	}
	expNeighbors.Put(a, []N{b})
	err = g.AddEdge(b, c, 20)
	if err != nil {
		t.Fatal(err)
	}
	expNeighbors.Put(b, []N{c})

	shouldContainAllNeighborsBy[N](t, g, expNeighbors, sortables.Equality[N])
}

// TestDisconnectedGraph Test a disconnected graph
func TestDisconnectedGraph(t *testing.T) {
	// Types
	type N = *tests.SortableNode
	type W = int
	type E = edge.ComparableWeighted[N, W]
	// Graph
	g := WeightedDigraphOf[N, W]()

	a := tests.NewSortableNode("A")
	b := tests.NewSortableNode("C")
	c := tests.NewSortableNode("C")
	g.AddNode(a)
	g.AddNode(b)
	g.AddNode(c)
	err := g.AddEdge(a, b, 10)
	if err != nil {
		t.Fatal(err)
	}

	if !g.HasEdge(a, b) {
		t.Errorf("Graph should have an edge from A to C.")
	}

	if g.HasEdge(b, c) {
		t.Errorf("Graph should not have an edge from C to C.")
	}
}
