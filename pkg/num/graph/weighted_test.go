// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package graph

import (
	"fmt"
	"github.com/andrerrcosta2/gtools/pkg/arrays"
	"github.com/andrerrcosta2/gtools/pkg/testdata"
	"testing"
)

func TestAddingNodesAndEdges(t *testing.T) {
	g := WeightedOrderedOf[testdata.TestNode, int]()
	a := testdata.TestNode("A")
	b := testdata.TestNode("B")
	c := testdata.TestNode("C")
	g.AddNode(a)
	g.AddNode(b)
	g.AddNode(c)
	g.AddEdge(a, b, 10)
	g.AddEdge(b, c, 20)

	nodes := g.Nodes()
	exp := []testdata.TestNode{a, b, c}

	fmt.Println("Nodes from graph:", nodes)
	fmt.Println("Expected nodes:", exp)

	// Check if nodes contain all expected nodes
	if !arrays.ContainsAllBy(nodes, exp, func(n1, n2 testdata.TestNode) bool { return n1.Equal(n2) }) {
		t.Errorf("Nodes test failed. \nGot: %v, \nExpected: %v\n", nodes, exp)
	}

	if len(nodes) != len(exp) {
		t.Errorf("Nodes test failed. \nGot: %v, \nexpected: %v\n", nodes, exp)
		return
	}
}

func TestNodeExistence(t *testing.T) {
	g := WeightedOrderedOf[testdata.TestNode, int]()
	a := testdata.TestNode("A")
	b := testdata.TestNode("B")
	g.AddNode(a)
	g.AddNode(b)

	nodes := g.Nodes()
	exp := []testdata.TestNode{a, b}

	if len(nodes) != len(exp) {
		t.Errorf("Nodes test failed. \nGot: %v, \nexpected: %v\n", nodes, exp)
		return
	}

	if !arrays.ContainsAllBy(nodes, exp, func(n1, n2 testdata.TestNode) bool { return n1.Equal(n2) }) {
		t.Errorf("Nodes test failed. \nGot: %v, \nExpected: %v\n", nodes, exp)
	}
}

func TestEdgeExistenceAndWeight(t *testing.T) {
	g := WeightedOrderedOf[testdata.TestNode, int]()
	a := testdata.TestNode("A")
	b := testdata.TestNode("B")
	c := testdata.TestNode("C")
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
	g := WeightedOrderedOf[testdata.TestNode, int]()
	a := testdata.TestNode("A")
	b := testdata.TestNode("B")
	c := testdata.TestNode("C")
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
	g := WeightedOrderedOf[testdata.TestNode, int]()
	a := testdata.TestNode("A")
	b := testdata.TestNode("B")
	c := testdata.TestNode("C")
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
func contains(slice []testdata.TestNode, item testdata.TestNode) bool {
	for _, elem := range slice {
		if elem.Equal(item) {
			return true
		}
	}
	return false
}
