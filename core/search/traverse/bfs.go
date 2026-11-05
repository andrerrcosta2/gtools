// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package traverse

import "github.com/andrerrcosta2/gtools/core/domain/function"

func BFSTraverse[N comparable](
	start N,
	expand func(N, func(N) bool) bool,
	visit function.Predicate[N],
) {
	const maxQueue = 1024 // pre-allocated buffer buffer
	queue := make([]N, maxQueue)
	var head, tail int

	visited := map[N]bool{}

	enqueue := func(n N) {
		queue[tail] = n
		tail = (tail + 1) % maxQueue
	}

	dequeue := func() N {
		n := queue[head]
		head = (head + 1) % maxQueue
		return n
	}

	isEmpty := func() bool { return head == tail }

	// start traversal
	visited[start] = true
	enqueue(start)

	for !isEmpty() {
		n := dequeue()

		if !visit(n) {
			break
		}

		// push-based neighbor expansion
		expand(n, func(next N) bool {
			if !visited[next] {
				visited[next] = true
				enqueue(next)
			}
			return true // continue traversal
		})
	}
}
