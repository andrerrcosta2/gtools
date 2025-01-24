// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package graphs

//func TestPatterns_Grid(t *testing.T) {
//	nodes := progression.Sequence[int](9)
//	graph, err := Grid[int, int](Directed, nodes...)
//	if err != nil {
//		t.Error(err)
//	}
//	fmt.Printf("Graph: %v\n", graph)
//
//	expectedAdjacencyList := map[int][]int{
//		0: {1, 3},
//		1: {0, 2, 4},
//		2: {1, 5},
//		3: {0, 4, 6},
//		4: {1, 3, 5, 7},
//		5: {2, 4, 8},
//		6: {3, 7},
//		7: {4, 6, 8},
//		8: {5, 7},
//	}
//
//	if len(graph.Nodes()) != len(nodes) {
//		t.Errorf("Expected %d nodes, but got %d", len(nodes), len(graph.Nodes()))
//	}
//
//	graphImpl := graph.(*DirectedGraph[int])
//
//	for node, expectedEdges := range expectedAdjacencyList {
//		edges := graphImpl.Neighbors(node)
//		if len(edges) != len(expectedEdges) {
//			t.Errorf("Node %d: Expected %d edges, but got %d", node, len(expectedEdges), len(edges))
//			continue
//		}
//		for _, edge := range expectedEdges {
//			if !arrays.Contains(edges, edge) {
//				t.Errorf("Node %d: Expected edge to %d not found", node, edge)
//			}
//		}
//	}
//}
