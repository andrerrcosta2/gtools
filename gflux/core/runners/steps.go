package runners

import (
	"github.com/andrerrcosta2/gtools/gflux/core/domain/runner"
	"github.com/andrerrcosta2/gtools/gflux/core/domain/stepper/step"
)

type State uint8

const (
	Blocked State = iota
	Empty
)

// StepAndRelease is a runner that stores releasable broadcasters and
// executes the longest executable prefix. A step becomes executable
// only after it has been released. The execution always respects
// the ordering defined by the underlying policy.
//
// Diagram:
//
//	step:     --A(blocked)----B(released)----C(blocked)----|
//	release:               A
//	exec:                  a-b---------------(wait)
//
// If Release is called for a step other than the first, only that step
// is released, but execution is deferred until the first step is released.
//
// If any step execution fails, execution stops immediately and the remaining
// step are discarded.
//
// StepAndRelease is concurrency-safe; calls to Step and Release may be made
// from multiple goroutines.
func StepAndRelease[S step.Releasable[T], T any](policy runner.Policy[S, T]) *SRRunner[S, T] {
	return &SRRunner[S, T]{
		policy: policy,
	}
}

type SRRunner[S step.Releasable[T], T any] struct {
	policy runner.Policy[S, T]
}

// Step registers a new step
func (r *SRRunner[S, T]) Step(step S) Ticket {
	r.policy.Push(step)
	peek, ok := r.policy.Peek()
	if ok && peek.Released() {
		r.run()
	}
	return ticket{
		fn: func() (State, error) {
			return r.release(step)
		},
	}
}

// Release releases a step and advances the chain
func (r *SRRunner[S, T]) release(step S) (State, error) {
	step.Release()
	return r.run()
}

func (r *SRRunner[S, T]) run() (State, error) {
	for {
		next, ok := r.policy.Peek()
		if !ok {
			return Empty, nil
		}

		if !next.Released() {
			return Blocked, nil
		}

		r.policy.Pop()

		if err := next.Exec(); err != nil {
			return r.state(), err
		}
	}
}

func (r *SRRunner[S, T]) state() State {
	if r.policy.Empty() {
		return Empty
	}
	return Blocked
}

type Ticket interface {
	Release() (State, error)
}

type ticket struct {
	fn func() (State, error)
}

func (t ticket) Release() (State, error) {
	return t.fn()
}
