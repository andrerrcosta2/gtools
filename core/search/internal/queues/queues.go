// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package queues

import "github.com/andrerrcosta2/gtools/core/domain/constraints/prim/nums"

type NodeID comparable

func NewPqItem[T NodeID, P nums.Ordered](node T, priority P) *PqItem[T, P] {
	return &PqItem[T, P]{
		node:     node,
		priority: priority,
	}
}

type PqItem[T NodeID, P nums.Ordered] struct {
	node     T
	priority P
	index    int
}

func (p *PqItem[T, P]) Node() T {
	return p.node
}

type Priority[T NodeID, P nums.Ordered] []*PqItem[T, P]

func (pq *Priority[T, P]) Push(x any) {
	item := x.(*PqItem[T, P])
	item.index = len(*pq)
	*pq = append(*pq, item)
}

func (pq *Priority[T, P]) Pop() any {
	old := *pq
	n := len(old)

	item := old[n-1]
	old[n-1] = nil

	item.index = -1
	*pq = old[0 : n-1]

	return item
}

func (pq Priority[T, P]) Len() int { return len(pq) }

func (pq Priority[T, P]) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}
func (pq Priority[T, P]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
