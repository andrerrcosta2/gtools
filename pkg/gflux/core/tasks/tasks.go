// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package tasks

import (
	"errors"
	"github.com/andrerrcosta2/gtools/core/gtools/functions"
	"github.com/andrerrcosta2/gtools/core/gtools/gerrors"
)

var ReleaseErrorTag = "release-error"

func NewStep(hash string, fn functions.Supplier[error]) *Step {
	return &Step{
		hash:     hash,
		fn:       fn,
		released: false,
	}
}

type Step struct {
	hash     string
	released bool
	fn       functions.Supplier[error]
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

func NewChunk(hash string, fn ...functions.Supplier[error]) *Chunk {
	return &Chunk{
		hash:     hash,
		released: false,
		fn:       fn,
	}
}

type Chunk struct {
	hash     string
	released bool
	fn       []functions.Supplier[error]
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
	return gerrors.Tagged(errors.New("Chunk is not released"), ReleaseErrorTag)
}

func NewStepBranch(hash string, fn functions.Supplier[error]) *StepBranch {
	return &StepBranch{
		hash:     hash,
		fn:       fn,
		released: false,
	}
}

type StepBranch struct {
	hash     string
	fn       functions.Supplier[error]
	branch   Runnable
	children []Runnable
	released bool
}

func (s *StepBranch) AddChild(branch Runnable) {
	s.children = append(s.children, branch)
}

func (s *StepBranch) Branch() (Runnable, bool) {
	return s.branch, s.branch != nil
}

func (s *StepBranch) Children() []Runnable {
	return s.children
}

func (s *StepBranch) Release() {
	s.released = true
}

func (s *StepBranch) Run() (err error) {
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
	return gerrors.Tagged(errors.New("Step is not released"), ReleaseErrorTag)
}

func (s *StepBranch) Unique() string {
	return s.hash
}
