// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package graphs

//// Grid creates a grid graph with the given nodes
//// The nodes are arranged in a square grid
////
//// Example:
//// Grid[int, int](Undirected, 0, 1, 2, 3, 4, 5, 6, 7, 8)
////
////	0 - 1 - 2
////	|   |   |
////	3 - 4 - 5
////	|   |   |
////	6 - 7 - 8
//func Grid[G any, N any](t Type, nodes ...G) (Graph[N], error) {
//	// Create the graph with the given nodes and types
//	graph, err := New[G, N](t, nodes...)
//	if err != nil {
//		return nil, err
//	}
//
//	g := any(graph).(SingleEdgedGraph[G])
//	size := int(math.Sqrt(float64(len(nodes))))
//
//	for i := range nodes {
//		// Horizontal connection: Connect to the right neighbor if not in the last column
//		if i%size != size-1 && i+1 < len(nodes) {
//			g.AddEdge(nodes[i], nodes[i+1])
//		}
//
//		// Vertical connection: Connect to the bottom neighbor if not in the last row
//		if i+size < len(nodes) {
//			g.AddEdge(nodes[i], nodes[i+size])
//		}
//	}
//
//	// Return the graph
//	return g.(Graph[N]), nil
//}
