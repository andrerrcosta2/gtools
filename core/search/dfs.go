// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package search

import (
	"github.com/andrerrcosta2/gtools/core/domain/function"
)

// DepthFirst performs a depth-first search on the graph starting from the given node.
// It returns a slice of visited node in the order they were visited.
func DepthFirst[N comparable](expand function.ExpandEach[N], start, target N) []N {
	visited := map[N]bool{}
	parent := map[N]N{}

	var found bool

	var walk function.Predicate[N]
	walk = func(n N) bool {
		if visited[n] {
			return true
		}
		visited[n] = true

		if n == target {
			found = true
			return false
		}

		return expand(n, func(next N) bool {
			if !visited[next] {
				parent[next] = n
				return walk(next)
			}
			return true
		})
	}

	walk(start)

	if !found {
		return nil
	}

	// reconstruct path (only allocation we keep)
	path := make([]N, 0, 8) // small hint capacity

	for cur := target; ; cur = parent[cur] {
		path = append(path, cur)
		if cur == start {
			break
		}
	}

	// reverse
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return path
}
