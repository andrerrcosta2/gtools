// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package traverse

import "github.com/andrerrcosta2/gtools/core/domain/function"

func DepthFirst[N comparable](expand function.ExpandEach[N], start N, visit function.Predicate[N]) {
	visited := map[N]bool{}

	var walk func(N) bool
	walk = func(n N) bool {
		if visited[n] {
			return true
		}
		visited[n] = true

		if !visit(n) {
			return false
		}

		return expand(n, func(next N) bool {
			return walk(next)
		})
	}

	walk(start)
}
