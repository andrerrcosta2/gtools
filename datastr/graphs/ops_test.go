// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package graphs

import (
	"github.com/andrerrcosta2/gtools/core/data/str"
	"github.com/andrerrcosta2/gtools/core/data/str/edges"
	"github.com/andrerrcosta2/gtools/core/domain/functions"
	"github.com/andrerrcosta2/gtools/gtests/testingtools/gtests"
)

// shouldContainNode tests retrieving a node from the graph and comparing
// it with the expected node. It flags an error if the node is not found in
// the graph.
func shouldContainNode[N any](t gtests.Loggable, g str.Graph[N], n N) {
	t.Helper()
	if !g.HasNode(n) {
		t.Errorf("Expected node '%s' to exist in the graph", n)
	}
}

// shouldContainAllNodesBy tests retrieving all nodes from the graph and
// comparing them with the expected nodes. It flags an error if any node
// from the slice is not found in the graph.
func shouldContainAllNodesBy[N any](t gtests.Loggable, g str.Graph[N], nodes []N, cmp functions.BiPredicate[N, N]) {
	t.Helper()

	t.StackLogf("All Nodes: [%d]%v\n", len(g.Nodes()), g.Nodes())

	for _, node := range nodes {
		if !g.HasNode(node) {
			t.Errorf("Expected node '%s' to exist in the graph", node)
		}
	}
}

func shouldContainNNodes[N any](t gtests.Loggable, g str.Graph[N], q int) {
	t.Helper()
	if len(g.Nodes()) != q {
		t.Errorf("Expected %d nodes, got %d", q, len(g.Nodes()))
	}
}

func shouldAddNodeIfNotExists[N any](t gtests.Loggable, g str.Graph[N], n N) {
	t.Helper()
	g.AddNode(n)
	if !g.HasNode(n) {
		t.Errorf("Expected node '%s' to exist in the graph", n)
	}
}

// shouldNotContainEdge tests retrieving an edge from the graph and comparing
// it with the expected edge. It flags an error if the edge is found in
// the graph.
func shouldNotContainEdge[G any, E edges.Edge[F, T], F any, T any](t gtests.Loggable, g str.EdgedGraph[G, E, F, T], e E) {
	t.Helper()
	if g.HasEdge(e.From(), e.To()) {
		t.Errorf("Expected edge from '%s' to '%s' to exist in the graph", e.From(), e.To())
	}
}

// shouldContainAllEdgesBy tests retrieving all edges from the graph and
// comparing them with the expected edges by a predicate. It flags an error if the edges
// are not found in the graph.
func shouldContainAllEdgesBy[G any, E edges.SortableEdge[F, T], F any, T any](t gtests.Loggable, g str.EdgedGraph[G, E, F, T], edges []E, compare functions.BiPredicate[E, E]) {
	t.Helper()

	// Iterate over each element in the expected slice
	for _, edge := range edges {
		var found bool
		for _, neighbor := range g.Edges() {
			// Check if the current element in the result slice is compare to the current element in the expected slice using the provided comparison function
			if compare(edge, neighbor) {
				t.StackLogf("'%s' found as edge\n", neighbor)
				found = true
				break
			}
		}
		// If the current element in the expected slice isn't found in the result slice, return false
		if !found {
			t.Errorf("'%s' not found as edge\n", edge)
		}
	}

	t.StackLogf("All Edges: [%d]%v\n", len(g.Edges()), g.Edges())
}

// shouldHaveNEdges tests retrieving all edges from the graph and comparing
// their length with the expected length.
// It flags an error if the lengths are different.
func shouldHaveNEdges[G any, E edges.Edge[F, T], F any, T any](t gtests.Loggable, g str.EdgedGraph[G, E, F, T], expLen int) {
	t.Helper()
	if len(g.Edges()) != expLen {
		t.Errorf("Expected %d edges, got %d", expLen, len(g.Edges()))
	}
}

// shouldContainAllNeighborsBy tests retrieving all neighbors from the graph and
// comparing them with the expected neighbors.
// It flags an error if any neighbor from the map is not found in the graph.
func shouldContainAllNeighborsBy[N any](t gtests.Loggable, g str.Graph[N], expected str.Map[N, []N], compare functions.BiPredicate[N, N]) {
	t.Helper()

	// Create an iterator over the expected map
	iter := expected.Iterator()

	// Iterate over each element in the expected map
	for key, exp, hasNext := iter.Next(); hasNext; key, exp, hasNext = iter.Next() {
		// Get the neighbors of the current element in the result slice
		neighbors := g.Neighbors(key)
		t.StackLogf("checking neighbors of '%s': %v\n", key, neighbors)

		if len(exp) != len(neighbors) {
			t.Errorf("expected '%d' neighbors for node '%s', got %d", len(exp), key, len(neighbors))
		}

		// Iterate over each element in the expected slice
		for _, edge := range exp {
			var found bool
			for _, neighbor := range neighbors {
				// Check if the current element in the result slice is compare to the current element in the expected slice using the provided comparison function
				if compare(edge, neighbor) {
					t.StackLogf("'%s' found as neighbor of '%s'\n", neighbor, key)
					found = true
					break
				}
			}
			// If the current element in the expected slice isn't found in the result slice, return false
			if !found {
				t.Errorf("'%s' not found as neighbor of '%s'\n", edge, key)
			}
		}
	}
}
