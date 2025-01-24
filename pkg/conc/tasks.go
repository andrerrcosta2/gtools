// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package conc

import (
	"errors"
	"github.com/andrerrcosta2/gtools/core/gerrors"
	"github.com/andrerrcosta2/gtools/core/gtools/functions"
	"github.com/andrerrcosta2/gtools/core/tasks"
	"sync"
)

func SeqTaskChain[T tasks.Runnable](tasks ...T) tasks.Chain[T] {
	return &seqTaskChain[T]{
		tasks: tasks,
	}
}

type seqTaskChain[T tasks.Runnable] struct {
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

func StepBranch(hash string, fn functions.Supplier[error]) tasks.ReleasableBranchOf[tasks.Runnable] {
	return &stepBranch{
		hash:     hash,
		fn:       fn,
		released: false,
	}
}

type stepBranch struct {
	mtx      sync.RWMutex
	hash     string
	fn       functions.Supplier[error]
	branch   tasks.Runnable
	children []tasks.Runnable
	released bool
}

func (s *stepBranch) AddChild(branch tasks.Runnable) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.children = append(s.children, branch)
}

func (s *stepBranch) Branch() (tasks.Runnable, bool) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return s.branch, s.branch != nil
}

func (s *stepBranch) Children() []tasks.Runnable {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return s.children
}

func (s *stepBranch) Release() {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.released = true
}

func (s *stepBranch) Run() (err error) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	if s.released {
		if err = s.fn(); err != nil {
			return
		} else {
			for _, child := range s.children {
				if err = child.Run(); err != nil {
					return err
				}
			}
		}
	}
	return gerrors.Tagged(errors.New("step is not released"), tasks.ReleaseErrorTag)
}

func (s *stepBranch) Unique() string {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return s.hash
}

//func DependencyTaskTree[D comparable](id D, fn functions.Supplier[error]) runners.Releaser {
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
//	fn         functions.Supplier[error]
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
