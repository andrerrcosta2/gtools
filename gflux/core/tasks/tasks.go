// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package tasks

import (
	"errors"
	"sync"

	"github.com/andrerrcosta2/gtools/gflux/core/domain/task"
)

func SeqTaskChain[T task.Runnable](tasks ...T) task.Chain[T] {
	return &seqTaskChain[T]{
		tasks: tasks,
	}
}

type seqTaskChain[T task.Runnable] struct {
	mtx   sync.Mutex
	tasks []T
}

func (t *seqTaskChain[T]) Insert(i int, task T) error {
	t.mtx.Lock()
	defer t.mtx.Unlock()
	if i > len(t.tasks) {
		return errors.New("index out of range")
	}
	t.tasks = append(t.tasks[:i], append([]T{task}, t.tasks[i:]...)...)
	return nil
}

func (t *seqTaskChain[T]) Append(tasks ...T) {
	t.mtx.Lock()
	defer t.mtx.Unlock()
	t.append(tasks...)
}

func (t *seqTaskChain[T]) append(tasks ...T) {
	t.tasks = append(t.tasks, tasks...)
}

func (t *seqTaskChain[T]) Prepend(tasks ...T) {
	t.mtx.Lock()
	defer t.mtx.Unlock()
	t.prepend(tasks...)
}

func (t *seqTaskChain[T]) prepend(tasks ...T) {
	t.tasks = append(tasks, t.tasks...)
}

func (t *seqTaskChain[T]) Remove(i int) {
	t.mtx.Lock()
	defer t.mtx.Unlock()
	t.remove(i)
}

func (t *seqTaskChain[T]) remove(i int) {
	t.tasks = append(t.tasks[:i], t.tasks[i+1:]...)
}

func (t *seqTaskChain[T]) clear() {
	t.tasks = make([]T, 0)
}

func (t *seqTaskChain[T]) Run() error {
	t.mtx.Lock()
	defer t.mtx.Unlock()
	for i, task := range t.tasks {
		if err := task.Run(); err != nil {
			t.tasks = t.tasks[i:]
			return err
		}
	}
	t.clear()
	return nil
}

//func DependencyTaskTree[D comparable](id D, fn function.Supplier[error]) runners.Releaser {
//	return &dependencyTree{
//		id:    id,
//		fn:    fn,
//		chain: make([]*dependencyTree[ID], 0),
//	}
//}
//
//type dependencyTree[D comparable] struct {
//	mtx        sync.Mutex
//	dependency D
//	released   bool
//	fn         function.Supplier[error]
//	chain      []*dependencyTree[D]
//}
//
//func (c *dependencyTree[D]) Release() error {
//	c.mtx.Lock()
//	defer c.mtx.Unlock()
//	c.released = true
//	for _, step := range c.chain {
//		if step.released {
//			// must acknowledge the nature of each error
//			step.run()
//		}
//	}
//}
//
//func (c *dependencyTree[D]) run() error {
//	if c.released {
//		for _, f := range c.fn {
//			if err := f(); err != nil {
//				return err
//			}
//		}
//		return nil
//	}
//	return errors.New("step is not released")
//}
