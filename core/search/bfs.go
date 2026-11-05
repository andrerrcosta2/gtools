// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package search

import (
	"github.com/andrerrcosta2/gtools/core/domain/function"
)

func BreathFirst[N comparable](expand function.ExpandEach[N], start, target N) []N {
	visited := map[N]bool{}
	parent := map[N]N{}

	// buffer buffer
	const maxQueue = 1024
	queue := make([]N, maxQueue)
	var head, tail int

	enqueue := func(n N) {
		queue[tail] = n
		tail = (tail + 1) % maxQueue
	}

	dequeue := func() N {
		n := queue[head]
		head = (head + 1) % maxQueue
		return n
	}

	isEmpty := func() bool {
		return head == tail
	}

	// enqueue start
	visited[start] = true
	enqueue(start)

	var found bool

	for !isEmpty() {
		n := dequeue()

		if n == target {
			found = true
			break
		}

		expand(n, func(next N) bool {
			if !visited[next] {
				visited[next] = true
				parent[next] = n
				enqueue(next)
			}
			return true
		})
	}

	if !found {
		return nil
	}

	// reconstruct path
	path := make([]N, 0, 8)
	for cur := target; ; cur = parent[cur] {
		path = append(path, cur)
		if cur == start {
			break
		}
	}

	// reverse path
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return path
}
