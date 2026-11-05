// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package graphs

import (
	"github.com/andrerrcosta2/gtools/core/data/str/dict"
	"github.com/andrerrcosta2/gtools/core/data/str/edge"
	"github.com/andrerrcosta2/gtools/core/data/str/graph"
	"github.com/andrerrcosta2/gtools/core/data/str/iterables"
	"github.com/andrerrcosta2/gtools/core/data/str/set"
	"github.com/andrerrcosta2/gtools/core/domain/function"
	"github.com/andrerrcosta2/gtools/core/testlite"
	"github.com/andrerrcosta2/gtools/datastr/internal/tests"
)

// shouldContainNode tests retrieving a node from the graph and comparing
// it with the expected node. It flags an error if the node is not found in
// the graph.
func shouldContainNode[N any](t testlite.LoggerTesting, g graph.Readable[N], n N) {
	t.Helper()
	if !g.HasNode(n) {
		t.Errorf("Expected node '%s' to exist in the graph", n)
	}
}

// shouldContainAllNodesBy tests retrieving all node from the graph and
// comparing them with the expected node. It flags an error if any node
// from the slice is not found in the graph.
func shouldContainAllNodesBy[N any](t testlite.LoggerTesting, g graph.Readable[N], nodes []N, cmp function.BiPredicate[N, N]) {
	t.Helper()

	t.Logf("All Nodes: [%d]%v\n", len(g.Nodes()), g.Nodes())

	for _, node := range nodes {
		if !g.HasNode(node) {
			t.Errorf("Expected node '%s' to exist in the graph", node)
		}
	}
}

func shouldContainNNodes[N any](t testlite.LoggerTesting, g graph.Readable[N], q int) {
	t.Helper()
	if len(g.Nodes()) != q {
		t.Errorf("Expected %d node, got %d", q, len(g.Nodes()))
	}
}

func shouldAddNodeIfNotExists[N any](t testlite.LoggerTesting, g graph.Mutable[N], n N) {
	t.Helper()
	g.AddNode(n)
	if !g.HasNode(n) {
		t.Errorf("Expected node '%s' to exist in the graph", n)
	}
}

// shouldNotContainEdge tests retrieving an edge from the graph and comparing
// it with the expected edge. It flags an error if the edge is found in
// the graph.
func shouldNotContainEdge[E edge.Typed[N], N any](
	t testlite.LoggerTesting,
	g graph.ReadableEdged[E, N],
	e E) {
	t.Helper()
	if g.HasEdge(e.From(), e.To()) {
		t.Errorf("Expected edge from '%s' to '%s' to exist in the graph", e.From(), e.To())
	}
}

// shouldContainAllEdgesBy tests retrieving all edge from the graph and
// comparing them with the expected edge by a predicate. It flags an error if the edge
// are not found in the graph.
func shouldContainAllEdgesBy[E edge.Typed[N], N any](
	t testlite.LoggerTesting,
	g graph.ReadableEdged[E, N],
	exp iterables.Slice[E],
	equals function.Equals[E]) {
	t.Helper()

	eq, mismatches := tests.EqualSlicesFunc(g.Edges(), exp, equals)
	if !eq {
		t.Errorf("Expected all edges to be equal, got: '%v'", mismatches)
	}
}

// shouldHaveNEdges tests retrieving all edge from the graph and comparing
// their length with the expected length.
// It flags an error if the lengths are different.
func shouldHaveNEdges[E edge.Typed[N], N any](
	t testlite.LoggerTesting,
	g graph.ReadableEdged[E, N],
	expLen int) {
	t.Helper()
	if len(g.Edges()) != expLen {
		t.Errorf("Expected %d edge, got %d", expLen, len(g.Edges()))
	}
}

// shouldContainAllNeighborsBy tests retrieving all neighbors from the graph and
// comparing them with the expected neighbors.
// It flags an error if any neighbor from the map is not found in the graph.
func shouldContainAllNeighborsBy[N any](
	t testlite.LoggerTesting,
	g graph.Iterable[N],
	expected dict.Mutable[N, []N],
	compare function.Equals[N]) {
	t.Helper()

	g.Range(func(node N, neighbors set.Readable[N]) bool {
		expNeighbors, exists := expected.Get(node)
		if !exists {
			t.Errorf("Expected node '%s' to exist in the expectation", node)
		}
		if len(expNeighbors) != neighbors.Len() {
			t.Errorf("Expected '%d' neighbors, got' %d'", neighbors.Len(), len(expNeighbors))
		}
		if eq, mismatches := tests.EqualSlicesFunc(expNeighbors, neighbors.Elements(), compare); !eq {
			t.Errorf("Expected node '%v' to contain all neighbors, got '%v'", node, mismatches)
		}
		return true
	})
}
