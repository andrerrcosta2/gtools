// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sources

import (
	"container/heap"

	"github.com/andrerrcosta2/gtools/core/domain/constraints/prim"
	"github.com/andrerrcosta2/gtools/gflux/core/domain/stepper/step"
)

type Heap[T prim.Ordered] struct {
	heap stepHeap[T]
}

func (s *Heap[T]) Next() (step.Step[T], bool) {
	return s.heap.popStep()
}

func (s *Heap[T]) Peek() (step.Step[T], bool) {
	return s.heap.peek()
}

type stepHeap[T prim.Ordered] struct {
	data []step.Step[T]
}

func (h *stepHeap[T]) Len() int { return len(h.data) }

func (h *stepHeap[T]) Less(i, j int) bool {
	return h.data[i].At() < h.data[j].At()
}

func (h *stepHeap[T]) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *stepHeap[T]) Push(x any) {
	h.data = append(h.data, x.(step.Step[T]))
}

func (h *stepHeap[T]) Pop() any {
	n := len(h.data)
	if n == 0 {
		return nil
	}

	x := h.data[n-1]
	h.data = h.data[:n-1]
	return x
}

func (h *stepHeap[T]) pushStep(s step.Step[T]) {
	heap.Push(h, s)
}

func (h *stepHeap[T]) popStep() (step.Step[T], bool) {
	if len(h.data) == 0 {
		return nil, false
	}
	return heap.Pop(h).(step.Step[T]), true
}

func (h *stepHeap[T]) peek() (step.Step[T], bool) {
	if len(h.data) == 0 {
		return nil, false
	}
	return h.data[0], true
}
