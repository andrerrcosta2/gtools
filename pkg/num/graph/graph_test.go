// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package graph

import (
	"github.com/andrerrcosta2/gtools/pkg/testdata/testsortables"
	"testing"
)

// TestIsCyclicOf tests the isCyclicOf function for various graph structures
func TestIsCyclicOf_UndirectedGraph(t *testing.T) {
	tests := []struct {
		name     string
		setup    func() GraphOf[testsortables.TestNode]
		expected bool
	}{
		{
			name: "No cycle with two nodes and one edge",
			setup: func() GraphOf[testsortables.TestNode] {
				g := UndirectOf[testsortables.TestNode]()
				g.AddNode(testsortables.TestNode("A"))
				g.AddNode(testsortables.TestNode("B"))
				g.AddEdge(testsortables.TestNode("A"), testsortables.TestNode("B")) // Edge A-B
				return g
			},
			expected: false,
		},
		{
			name: "Cycle with three nodes in a triangle",
			setup: func() GraphOf[testsortables.TestNode] {
				g := UndirectOf[testsortables.TestNode]()
				g.AddNode(testsortables.TestNode("A"))
				g.AddNode(testsortables.TestNode("B"))
				g.AddNode(testsortables.TestNode("C"))
				g.AddEdge(testsortables.TestNode("A"), testsortables.TestNode("B")) // Edge A-B
				g.AddEdge(testsortables.TestNode("B"), testsortables.TestNode("C")) // Edge B-C
				g.AddEdge(testsortables.TestNode("C"), testsortables.TestNode("A")) // Edge C-A
				return g
			},
			expected: true,
		},
		{
			name: "Cycle with four nodes in a square",
			setup: func() GraphOf[testsortables.TestNode] {
				g := UndirectOf[testsortables.TestNode]()
				g.AddNode(testsortables.TestNode("A"))
				g.AddNode(testsortables.TestNode("B"))
				g.AddNode(testsortables.TestNode("C"))
				g.AddNode(testsortables.TestNode("D"))
				g.AddEdge(testsortables.TestNode("A"), testsortables.TestNode("B")) // Edge A-B
				g.AddEdge(testsortables.TestNode("B"), testsortables.TestNode("C")) // Edge B-C
				g.AddEdge(testsortables.TestNode("C"), testsortables.TestNode("D")) // Edge C-D
				g.AddEdge(testsortables.TestNode("D"), testsortables.TestNode("A")) // Edge D-A
				return g
			},
			expected: true,
		},
		{
			name: "Cycle with four nodes and a diagonal",
			setup: func() GraphOf[testsortables.TestNode] {
				g := UndirectOf[testsortables.TestNode]()
				g.AddNode(testsortables.TestNode("A"))
				g.AddNode(testsortables.TestNode("B"))
				g.AddNode(testsortables.TestNode("C"))
				g.AddNode(testsortables.TestNode("D"))
				g.AddEdge(testsortables.TestNode("A"), testsortables.TestNode("B")) // Edge A-B
				g.AddEdge(testsortables.TestNode("B"), testsortables.TestNode("C")) // Edge B-C
				g.AddEdge(testsortables.TestNode("C"), testsortables.TestNode("D")) // Edge C-D
				g.AddEdge(testsortables.TestNode("D"), testsortables.TestNode("A")) // Edge D-A
				g.AddEdge(testsortables.TestNode("A"), testsortables.TestNode("C")) // Diagonal A-C
				return g
			},
			expected: true,
		},
		{
			name: "Self-loop detection",
			setup: func() GraphOf[testsortables.TestNode] {
				g := UndirectOf[testsortables.TestNode]()
				g.AddNode(testsortables.TestNode("A"))
				g.AddEdge(testsortables.TestNode("A"), testsortables.TestNode("A")) // Self-loop
				return g
			},
			expected: true,
		},
		{
			name: "Disconnected graph",
			setup: func() GraphOf[testsortables.TestNode] {
				g := UndirectOf[testsortables.TestNode]()
				g.AddNode(testsortables.TestNode("A"))
				g.AddNode(testsortables.TestNode("B"))
				g.AddNode(testsortables.TestNode("C"))
				g.AddEdge(testsortables.TestNode("A"), testsortables.TestNode("B")) // Edge A-B
				// Node C is disconnected
				return g
			},
			expected: false,
		},
		{
			name: "Multiple disconnected components with cycles",
			setup: func() GraphOf[testsortables.TestNode] {
				g := UndirectOf[testsortables.TestNode]()
				g.AddNode(testsortables.TestNode("A"))
				g.AddNode(testsortables.TestNode("B"))
				g.AddNode(testsortables.TestNode("C"))
				g.AddNode(testsortables.TestNode("D"))
				g.AddEdge(testsortables.TestNode("A"), testsortables.TestNode("B")) // Edge A-B
				g.AddEdge(testsortables.TestNode("B"), testsortables.TestNode("C")) // Edge B-C
				g.AddEdge(testsortables.TestNode("C"), testsortables.TestNode("A")) // Edge C-A (Cycle)
				g.AddEdge(testsortables.TestNode("D"), testsortables.TestNode("D")) // Self-loop on D
				return g
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			graph := tt.setup()
			if got := isCyclicOf[testsortables.TestNode](graph); got != tt.expected {
				t.Errorf("isCyclicOf() = %v, want %v\nGraph: %v\n", got, tt.expected, graph)
			}
		})
	}
}

func TestIsCyclicOf_DirectedGraph(t *testing.T) {
	tests := []struct {
		name     string
		setup    func() GraphOf[testsortables.TestNode]
		expected bool
	}{
		{
			name: "No Cycle",
			setup: func() GraphOf[testsortables.TestNode] {
				g := DigraphOf[testsortables.TestNode]()
				g.AddNode(testsortables.TestNode("A"))
				g.AddNode(testsortables.TestNode("B"))
				g.AddNode(testsortables.TestNode("C"))
				g.AddEdge(testsortables.TestNode("A"), testsortables.TestNode("B"))
				g.AddEdge(testsortables.TestNode("B"), testsortables.TestNode("C"))
				return g
			},
			expected: false,
		},
		{
			name: "Simple Cycle",
			setup: func() GraphOf[testsortables.TestNode] {
				g := DigraphOf[testsortables.TestNode]()
				g.AddNode(testsortables.TestNode("A"))
				g.AddNode(testsortables.TestNode("B"))
				g.AddNode(testsortables.TestNode("C"))
				g.AddEdge(testsortables.TestNode("A"), testsortables.TestNode("B"))
				g.AddEdge(testsortables.TestNode("B"), testsortables.TestNode("C"))
				g.AddEdge(testsortables.TestNode("C"), testsortables.TestNode("A")) // Creates a cycle: A -> B -> C -> A
				return g
			},
			expected: true,
		},
		{
			name: "Self Loop",
			setup: func() GraphOf[testsortables.TestNode] {
				g := DigraphOf[testsortables.TestNode]()
				g.AddNode(testsortables.TestNode("A"))
				g.AddEdge(testsortables.TestNode("A"), testsortables.TestNode("A")) // Self-loop
				return g
			},
			expected: true,
		},
		{
			name: "Multi Component with Cycle",
			setup: func() GraphOf[testsortables.TestNode] {
				g := DigraphOf[testsortables.TestNode]()
				// First component (no cycle)
				g.AddNode(testsortables.TestNode("A"))
				g.AddNode(testsortables.TestNode("B"))
				g.AddEdge(testsortables.TestNode("A"), testsortables.TestNode("B"))

				// Second component (with a cycle)
				g.AddNode(testsortables.TestNode("C"))
				g.AddNode(testsortables.TestNode("D"))
				g.AddNode(testsortables.TestNode("E"))
				g.AddEdge(testsortables.TestNode("C"), testsortables.TestNode("D"))
				g.AddEdge(testsortables.TestNode("D"), testsortables.TestNode("E"))
				g.AddEdge(testsortables.TestNode("E"), testsortables.TestNode("C")) // Creates a cycle: C -> D -> E -> C
				return g
			},
			expected: true,
		},
		{
			name: "Multi Disconnected Acyclic",
			setup: func() GraphOf[testsortables.TestNode] {
				g := DigraphOf[testsortables.TestNode]()

				// First component
				g.AddNode(testsortables.TestNode("A"))
				g.AddNode(testsortables.TestNode("B"))
				g.AddEdge(testsortables.TestNode("A"), testsortables.TestNode("B"))

				// Second component
				g.AddNode(testsortables.TestNode("C"))
				g.AddNode(testsortables.TestNode("D"))
				g.AddEdge(testsortables.TestNode("C"), testsortables.TestNode("D"))
				return g
			},
			expected: false,
		},
		{
			name: "Large Cyclic Graph",
			setup: func() GraphOf[testsortables.TestNode] {
				g := DigraphOf[testsortables.TestNode]()
				// Adding nodes
				nodes := []testsortables.TestNode{"A", "B", "C", "D", "E", "F"}
				for _, node := range nodes {
					g.AddNode(node)
				}

				// Adding edges (with a cycle)
				g.AddEdge(testsortables.TestNode("A"), testsortables.TestNode("B"))
				g.AddEdge(testsortables.TestNode("B"), testsortables.TestNode("C"))
				g.AddEdge(testsortables.TestNode("C"), testsortables.TestNode("D"))
				g.AddEdge(testsortables.TestNode("D"), testsortables.TestNode("E"))
				g.AddEdge(testsortables.TestNode("E"), testsortables.TestNode("F"))
				g.AddEdge(testsortables.TestNode("F"), testsortables.TestNode("C")) // Creates a cycle: C -> D -> E -> F -> C
				return g
			},
			expected: true,
		},
		{
			name: "Single Node No Edges",
			setup: func() GraphOf[testsortables.TestNode] {
				g := DigraphOf[testsortables.TestNode]()
				g.AddNode(testsortables.TestNode("A"))
				return g
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			graph := tt.setup()
			result := isCyclicOf[testsortables.TestNode](graph)
			if result != tt.expected {
				t.Errorf("isCyclicOf() = %v; want %v\nGraph:\n%v\n", result, tt.expected, graph)
			}
		})
	}
}

func TestIsCyclicOf_WeightedGraph(t *testing.T) {
	tests := []struct {
		name     string
		setup    func() Graph[testsortables.TestNode]
		expected bool
	}{
		{
			name: "No cycle with four nodes in a square",
			setup: func() Graph[testsortables.TestNode] {
				g := WeightedOrderedOf[testsortables.TestNode, int]()
				g.AddNode(testsortables.TestNode("A"))
				g.AddNode(testsortables.TestNode("B"))
				g.AddNode(testsortables.TestNode("C"))
				g.AddNode(testsortables.TestNode("D"))
				g.AddEdge(testsortables.TestNode("A"), testsortables.TestNode("B"), 1) // Edge A-B
				g.AddEdge(testsortables.TestNode("B"), testsortables.TestNode("C"), 1) // Edge B-C
				g.AddEdge(testsortables.TestNode("C"), testsortables.TestNode("D"), 1) // Edge C-D
				g.AddEdge(testsortables.TestNode("D"), testsortables.TestNode("A"), 1) // Edge D-A
				return g
			},
			expected: true,
		},
		{
			name: "No cycle with three nodes in a line",
			setup: func() Graph[testsortables.TestNode] {
				g := WeightedOrderedOf[testsortables.TestNode, int]()
				g.AddNode(testsortables.TestNode("A"))
				g.AddNode(testsortables.TestNode("B"))
				g.AddNode(testsortables.TestNode("C"))
				g.AddEdge(testsortables.TestNode("A"), testsortables.TestNode("B"), 1)
				g.AddEdge(testsortables.TestNode("B"), testsortables.TestNode("C"), 1)
				return g
			},
			expected: false,
		},
		{
			name: "Cycle with weighted edges",
			setup: func() Graph[testsortables.TestNode] {
				g := WeightedOrderedOf[testsortables.TestNode, int]()
				g.AddNode(testsortables.TestNode("A"))
				g.AddNode(testsortables.TestNode("B"))
				g.AddNode(testsortables.TestNode("C"))
				g.AddEdge(testsortables.TestNode("A"), testsortables.TestNode("B"), 1)
				g.AddEdge(testsortables.TestNode("B"), testsortables.TestNode("C"), 1)
				g.AddEdge(testsortables.TestNode("C"), testsortables.TestNode("A"), 1) // Creates a cycle: A -> B -> C -> A
				return g
			},
			expected: true,
		},
		{
			name: "Single node with a self-loop",
			setup: func() Graph[testsortables.TestNode] {
				g := WeightedOrderedOf[testsortables.TestNode, int]()
				g.AddNode(testsortables.TestNode("A"))
				g.AddEdge(testsortables.TestNode("A"), testsortables.TestNode("A"), 1) // Self-loop
				return g
			},
			expected: true,
		},
		{
			name: "Disconnected graph with one component containing a cycle",
			setup: func() Graph[testsortables.TestNode] {
				g := WeightedOrderedOf[testsortables.TestNode, int]()

				// First component (no cycle)
				g.AddNode(testsortables.TestNode("A"))
				g.AddNode(testsortables.TestNode("B"))
				g.AddEdge(testsortables.TestNode("A"), testsortables.TestNode("B"), 1)

				// Second component (with a cycle)
				g.AddNode(testsortables.TestNode("C"))
				g.AddNode(testsortables.TestNode("D"))
				g.AddNode(testsortables.TestNode("E"))
				g.AddEdge(testsortables.TestNode("C"), testsortables.TestNode("D"), 1)
				g.AddEdge(testsortables.TestNode("D"), testsortables.TestNode("E"), 1)
				g.AddEdge(testsortables.TestNode("E"), testsortables.TestNode("C"), 1) // Creates a cycle: C -> D -> E -> C
				return g
			},
			expected: true,
		},
		{
			name: "Disconnected graph with all acyclic components",
			setup: func() Graph[testsortables.TestNode] {
				g := WeightedOrderedOf[testsortables.TestNode, int]()

				// First component
				g.AddNode(testsortables.TestNode("A"))
				g.AddNode(testsortables.TestNode("B"))
				g.AddEdge(testsortables.TestNode("A"), testsortables.TestNode("B"), 1)

				// Second component
				g.AddNode(testsortables.TestNode("C"))
				g.AddNode(testsortables.TestNode("D"))
				g.AddEdge(testsortables.TestNode("C"), testsortables.TestNode("D"), 1)
				return g
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			graph := tt.setup()
			if result := isCyclicOf(graph); result != tt.expected {
				t.Errorf("isCyclicOf() = %v; want %v\nGraph:\n%v\n", result, tt.expected, graph)
			}
		})
	}
}
