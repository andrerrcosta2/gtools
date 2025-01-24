// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package graphs

import (
	"github.com/andrerrcosta2/gtools/core/data/str"
	"github.com/andrerrcosta2/gtools/core/data/str/edges"
	"github.com/andrerrcosta2/gtools/datastr/internal/tests"
	"github.com/andrerrcosta2/gtools/gtests"
	"github.com/andrerrcosta2/gtools/gtests/testingtools"
	"testing"
)

// TestIsCyclicOf tests the isCyclicOf function for various graph structures
func TestIsCyclicOf_UndirectedGraph(t *testing.T) {
	// Helper
	tt := testingtools.LoggersLite(t, gtests.LogOnFailure)
	// Types
	type N = *tests.SortableNode
	type E = edges.SortableSingleTyped[N]

	tcs := []struct {
		name     string
		setup    func() str.GraphOf[N, E]
		expected bool
	}{
		{
			name: "No cycle with two nodes and one edge",
			setup: func() str.GraphOf[N, E] {
				g := UndirectedOf[N]()
				g.AddNode(tests.NewSortableNode("A"))
				g.AddNode(tests.NewSortableNode("B"))
				g.AddEdge(tests.NewSortableNode("A"), tests.NewSortableNode("B")) // Edge A-B
				return g
			},
			expected: false,
		},
		{
			name: "Cycle with three nodes in a triangle",
			setup: func() str.GraphOf[N, E] {
				g := UndirectedOf[N]()
				g.AddNode(tests.NewSortableNode("A"))
				g.AddNode(tests.NewSortableNode("B"))
				g.AddNode(tests.NewSortableNode("C"))
				g.AddEdge(tests.NewSortableNode("A"), tests.NewSortableNode("B")) // Edge A-B
				g.AddEdge(tests.NewSortableNode("B"), tests.NewSortableNode("C")) // Edge B-C
				g.AddEdge(tests.NewSortableNode("C"), tests.NewSortableNode("A")) // Edge C-A
				return g
			},
			expected: true,
		},
		{
			name: "Cycle with four nodes in a square",
			setup: func() str.GraphOf[N, E] {
				g := UndirectedOf[N]()
				g.AddNode(tests.NewSortableNode("A"))
				g.AddNode(tests.NewSortableNode("B"))
				g.AddNode(tests.NewSortableNode("C"))
				g.AddNode(tests.NewSortableNode("D"))
				g.AddEdge(tests.NewSortableNode("A"), tests.NewSortableNode("B")) // Edge A-B
				g.AddEdge(tests.NewSortableNode("B"), tests.NewSortableNode("C")) // Edge B-C
				g.AddEdge(tests.NewSortableNode("C"), tests.NewSortableNode("D")) // Edge C-D
				g.AddEdge(tests.NewSortableNode("D"), tests.NewSortableNode("A")) // Edge D-A
				return g
			},
			expected: true,
		},
		{
			name: "Cycle with four nodes and a diagonal",
			setup: func() str.GraphOf[N, E] {
				g := UndirectedOf[N]()
				g.AddNode(tests.NewSortableNode("A"))
				g.AddNode(tests.NewSortableNode("B"))
				g.AddNode(tests.NewSortableNode("C"))
				g.AddNode(tests.NewSortableNode("D"))
				g.AddEdge(tests.NewSortableNode("A"), tests.NewSortableNode("B")) // Edge A-B
				g.AddEdge(tests.NewSortableNode("B"), tests.NewSortableNode("C")) // Edge B-C
				g.AddEdge(tests.NewSortableNode("C"), tests.NewSortableNode("D")) // Edge C-D
				g.AddEdge(tests.NewSortableNode("D"), tests.NewSortableNode("A")) // Edge D-A
				g.AddEdge(tests.NewSortableNode("A"), tests.NewSortableNode("C")) // Diagonal A-C
				return g
			},
			expected: true,
		},
		{
			name: "Self-loop detection",
			setup: func() str.GraphOf[N, E] {
				g := UndirectedOf[N]()
				g.AddNode(tests.NewSortableNode("A"))
				g.AddEdge(tests.NewSortableNode("A"), tests.NewSortableNode("A")) // Self-loop
				return g
			},
			expected: true,
		},
		{
			name: "Disconnected graph",
			setup: func() str.GraphOf[N, E] {
				g := UndirectedOf[N]()
				g.AddNode(tests.NewSortableNode("A"))
				g.AddNode(tests.NewSortableNode("B"))
				g.AddNode(tests.NewSortableNode("C"))
				g.AddEdge(tests.NewSortableNode("A"), tests.NewSortableNode("B")) // Edge A-B
				// Node C is disconnected
				return g
			},
			expected: false,
		},
		{
			name: "Multiple disconnected components with cycles",
			setup: func() str.GraphOf[N, E] {
				g := UndirectedOf[N]()
				g.AddNode(tests.NewSortableNode("A"))
				g.AddNode(tests.NewSortableNode("B"))
				g.AddNode(tests.NewSortableNode("C"))
				g.AddNode(tests.NewSortableNode("D"))
				g.AddEdge(tests.NewSortableNode("A"), tests.NewSortableNode("B")) // Edge A-B
				g.AddEdge(tests.NewSortableNode("B"), tests.NewSortableNode("C")) // Edge B-C
				g.AddEdge(tests.NewSortableNode("C"), tests.NewSortableNode("A")) // Edge C-A (Cycle)
				g.AddEdge(tests.NewSortableNode("D"), tests.NewSortableNode("D")) // Self-loop on D
				return g
			},
			expected: true,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			graph := tc.setup()
			if got := isCyclicOf[N](graph); got != tc.expected {
				tt.Errorf("isCyclicOf() = %v, want %v\nGraph: %v\n", got, tc.expected, graph)
			}
		})
	}

	tt.PrintLogStack()
}

// TestIsCyclicOf_DirectedGraph tests the isCyclicOf function for various directed graph structures
func TestIsCyclicOf_DirectedGraph(t *testing.T) {
	// Helper
	tt := testingtools.LoggersLite(t, gtests.LogOnFailure)
	// Types
	type N = *tests.SortableNode
	type E = edges.SortableSingleTyped[N]

	tcs := []struct {
		name     string
		setup    func() str.GraphOf[N, E]
		expected bool
	}{
		{
			name: "No Cycle",
			setup: func() str.GraphOf[N, E] {
				g := DigraphOf[N]()
				g.AddNode(tests.NewSortableNode("A"))
				g.AddNode(tests.NewSortableNode("B"))
				g.AddNode(tests.NewSortableNode("C"))
				g.AddEdge(tests.NewSortableNode("A"), tests.NewSortableNode("B"))
				g.AddEdge(tests.NewSortableNode("B"), tests.NewSortableNode("C"))
				return g
			},
			expected: false,
		},
		{
			name: "Simple Cycle",
			setup: func() str.GraphOf[N, E] {
				g := DigraphOf[N]()
				g.AddNode(tests.NewSortableNode("A"))
				g.AddNode(tests.NewSortableNode("B"))
				g.AddNode(tests.NewSortableNode("C"))
				g.AddEdge(tests.NewSortableNode("A"), tests.NewSortableNode("B"))
				g.AddEdge(tests.NewSortableNode("B"), tests.NewSortableNode("C"))
				g.AddEdge(tests.NewSortableNode("C"), tests.NewSortableNode("A")) // Creates a cycle: A -> B -> C -> A
				return g
			},
			expected: true,
		},
		{
			name: "Self Loop",
			setup: func() str.GraphOf[N, E] {
				g := DigraphOf[N]()
				g.AddNode(tests.NewSortableNode("A"))
				g.AddEdge(tests.NewSortableNode("A"), tests.NewSortableNode("A")) // Self-loop
				return g
			},
			expected: true,
		},
		{
			name: "Multi Component with Cycle",
			setup: func() str.GraphOf[N, E] {
				g := DigraphOf[N]()
				// First component (no cycle)
				g.AddNode(tests.NewSortableNode("A"))
				g.AddNode(tests.NewSortableNode("B"))
				g.AddEdge(tests.NewSortableNode("A"), tests.NewSortableNode("B"))

				// Second component (with a cycle)
				g.AddNode(tests.NewSortableNode("C"))
				g.AddNode(tests.NewSortableNode("D"))
				g.AddNode(tests.NewSortableNode("E"))
				g.AddEdge(tests.NewSortableNode("C"), tests.NewSortableNode("D"))
				g.AddEdge(tests.NewSortableNode("D"), tests.NewSortableNode("E"))
				g.AddEdge(tests.NewSortableNode("E"), tests.NewSortableNode("C")) // Creates a cycle: C -> D -> E -> C
				return g
			},
			expected: true,
		},
		{
			name: "Multi Disconnected Acyclic",
			setup: func() str.GraphOf[N, E] {
				g := DigraphOf[N]()

				// First component
				g.AddNode(tests.NewSortableNode("A"))
				g.AddNode(tests.NewSortableNode("B"))
				g.AddEdge(tests.NewSortableNode("A"), tests.NewSortableNode("B"))

				// Second component
				g.AddNode(tests.NewSortableNode("C"))
				g.AddNode(tests.NewSortableNode("D"))
				g.AddEdge(tests.NewSortableNode("C"), tests.NewSortableNode("D"))
				return g
			},
			expected: false,
		},
		{
			name: "Large Cyclic Graph",
			setup: func() str.GraphOf[N, E] {
				g := DigraphOf[N]()
				// Adding nodes
				tests.NewSortableNodes("A", "B", "C", "D", "E", "F").Each(g.AddNode)

				// Adding edges (with a cycle)
				g.AddEdge(tests.NewSortableNode("A"), tests.NewSortableNode("B"))
				g.AddEdge(tests.NewSortableNode("B"), tests.NewSortableNode("C"))
				g.AddEdge(tests.NewSortableNode("C"), tests.NewSortableNode("D"))
				g.AddEdge(tests.NewSortableNode("D"), tests.NewSortableNode("E"))
				g.AddEdge(tests.NewSortableNode("E"), tests.NewSortableNode("F"))
				g.AddEdge(tests.NewSortableNode("F"), tests.NewSortableNode("C")) // Creates a cycle: C -> D -> E -> F -> C
				return g
			},
			expected: true,
		},
		{
			name: "Single Node No Edges",
			setup: func() str.GraphOf[N, E] {
				g := DigraphOf[N]()
				g.AddNode(tests.NewSortableNode("A"))
				return g
			},
			expected: false,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			graph := tc.setup()
			result := isCyclicOf[N](graph)
			if result != tc.expected {
				tt.Errorf("isCyclicOf() = %v; want %v\nGraph:\n%v\n", result, tc.expected, graph)
			}
		})
	}

	tt.PrintLogStack()
}

// TestIsCyclicOf_WeightedGraph tests the isCyclicOf function for various weighted graph structures
func TestIsCyclicOf_WeightedGraph(t *testing.T) {
	// Helper
	tt := testingtools.LoggersLite(t, gtests.LogOnFailure)
	// Types
	type N = *tests.SortableNode
	type W = int
	type E = edges.SortableSingleTypedWeighted[N, W]

	tcs := []struct {
		name     string
		setup    func() str.WOrderedGraphOf[N, W, E]
		expected bool
	}{
		{
			name: "No cycle with four nodes in a square",
			setup: func() str.WOrderedGraphOf[N, W, E] {
				g := WeightedSortableDigraphOf[N, W]()
				g.AddNode(tests.NewSortableNode("A"))
				g.AddNode(tests.NewSortableNode("B"))
				g.AddNode(tests.NewSortableNode("C"))
				g.AddNode(tests.NewSortableNode("D"))
				g.AddEdge(tests.NewSortableNode("A"), tests.NewSortableNode("B"), 1) // Edge A-B
				g.AddEdge(tests.NewSortableNode("B"), tests.NewSortableNode("C"), 1) // Edge B-C
				g.AddEdge(tests.NewSortableNode("C"), tests.NewSortableNode("D"), 1) // Edge C-D
				g.AddEdge(tests.NewSortableNode("D"), tests.NewSortableNode("A"), 1) // Edge D-A
				return g
			},
			expected: true,
		},
		{
			name: "No cycle with three nodes in a line",
			setup: func() str.WOrderedGraphOf[N, W, E] {
				g := WeightedSortableDigraphOf[N, W]()
				g.AddNode(tests.NewSortableNode("A"))
				g.AddNode(tests.NewSortableNode("B"))
				g.AddNode(tests.NewSortableNode("C"))
				g.AddEdge(tests.NewSortableNode("A"), tests.NewSortableNode("B"), 1)
				g.AddEdge(tests.NewSortableNode("B"), tests.NewSortableNode("C"), 1)
				return g
			},
			expected: false,
		},
		{
			name: "Cycle with weighted edges",
			setup: func() str.WOrderedGraphOf[N, W, E] {
				g := WeightedSortableDigraphOf[N, W]()
				g.AddNode(tests.NewSortableNode("A"))
				g.AddNode(tests.NewSortableNode("B"))
				g.AddNode(tests.NewSortableNode("C"))
				g.AddEdge(tests.NewSortableNode("A"), tests.NewSortableNode("B"), 1)
				g.AddEdge(tests.NewSortableNode("B"), tests.NewSortableNode("C"), 1)
				g.AddEdge(tests.NewSortableNode("C"), tests.NewSortableNode("A"), 1) // Creates a cycle: A -> B -> C -> A
				return g
			},
			expected: true,
		},
		{
			name: "Single node with a self-loop",
			setup: func() str.WOrderedGraphOf[N, W, E] {
				g := WeightedSortableDigraphOf[N, W]()
				g.AddNode(tests.NewSortableNode("A"))
				g.AddEdge(tests.NewSortableNode("A"), tests.NewSortableNode("A"), 1) // Self-loop
				return g
			},
			expected: true,
		},
		{
			name: "Disconnected graph with one component containing a cycle",
			setup: func() str.WOrderedGraphOf[N, W, E] {
				g := WeightedSortableDigraphOf[N, W]()

				// First component (no cycle)
				g.AddNode(tests.NewSortableNode("A"))
				g.AddNode(tests.NewSortableNode("B"))
				g.AddEdge(tests.NewSortableNode("A"), tests.NewSortableNode("B"), 1)

				// Second component (with a cycle)
				g.AddNode(tests.NewSortableNode("C"))
				g.AddNode(tests.NewSortableNode("D"))
				g.AddNode(tests.NewSortableNode("E"))
				g.AddEdge(tests.NewSortableNode("C"), tests.NewSortableNode("D"), 1)
				g.AddEdge(tests.NewSortableNode("D"), tests.NewSortableNode("E"), 1)
				g.AddEdge(tests.NewSortableNode("E"), tests.NewSortableNode("C"), 1) // Creates a cycle: C -> D -> E -> C
				return g
			},
			expected: true,
		},
		{
			name: "Disconnected graph with all acyclic components",
			setup: func() str.WOrderedGraphOf[N, W, E] {
				g := WeightedSortableDigraphOf[N, W]()

				// First component
				g.AddNode(tests.NewSortableNode("A"))
				g.AddNode(tests.NewSortableNode("B"))
				g.AddEdge(tests.NewSortableNode("A"), tests.NewSortableNode("B"), 1)

				// Second component
				g.AddNode(tests.NewSortableNode("C"))
				g.AddNode(tests.NewSortableNode("D"))
				g.AddEdge(tests.NewSortableNode("C"), tests.NewSortableNode("D"), 1)
				return g
			},
			expected: false,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			graph := tc.setup()
			if result := isCyclicOf[N](graph); result != tc.expected {
				tt.Errorf("isCyclicOf() = %v; want %v\nGraph:\n%v\n", result, tc.expected, graph)
			}
		})
	}

	tt.PrintLogStack()
}
