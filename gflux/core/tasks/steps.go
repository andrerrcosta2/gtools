// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package tasks

import (
	"errors"
	"sync"

	"github.com/andrerrcosta2/gtools/core/domain/function"
	"github.com/andrerrcosta2/gtools/core/domain/gerrors"
	"github.com/andrerrcosta2/gtools/gflux/core/domain/task"
)

var ReleaseErrorTag = "release-error"

func NewStep(hash string, fn function.Supplier[error]) *Step {
	return &Step{
		hash:     hash,
		fn:       fn,
		released: false,
	}
}

type Step struct {
	hash     string
	released bool
	fn       function.Supplier[error]
}

func (s *Step) Unique() string {
	return s.hash
}

func (s *Step) Release() {
	s.released = true
}

func (s *Step) Run() error {
	if s.released {
		return s.fn()
	}
	return gerrors.Tagged(errors.New("step is not released"))
}

func NewChunk(hash string, fn ...function.Supplier[error]) *Chunk {
	return &Chunk{
		hash:     hash,
		released: false,
		fn:       fn,
	}
}

type Chunk struct {
	hash     string
	released bool
	fn       []function.Supplier[error]
}

func (c *Chunk) Unique() string {
	return c.hash
}

func (c *Chunk) Release() {
	c.released = true
}

func (c *Chunk) Run() error {
	if c.released {
		for _, f := range c.fn {
			if err := f(); err != nil {
				return err
			}
		}
		return nil
	}
	return errors.New("Chunk is not released")
}

func StepBranch(hash string, fn function.Supplier[error]) task.ReleasableBranchOf[task.Runnable] {
	return &stepBranch{
		hash:     hash,
		fn:       fn,
		released: false,
	}
}

type stepBranch struct {
	mtx      sync.RWMutex
	hash     string
	fn       function.Supplier[error]
	branch   task.Runnable
	children []task.Runnable
	released bool
}

func (s *stepBranch) AddChild(branch task.Runnable) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.children = append(s.children, branch)
}

func (s *stepBranch) Branch() (task.Runnable, bool) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return s.branch, s.branch != nil
}

func (s *stepBranch) Children() []task.Runnable {
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
	return errors.New("step is not released")
}

func (s *stepBranch) ID() string {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return s.hash
}
