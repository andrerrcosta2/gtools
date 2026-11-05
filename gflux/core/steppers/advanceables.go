// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package steppers

import (
	"github.com/andrerrcosta2/gtools/core/domain/constraints"
	"github.com/andrerrcosta2/gtools/gflux/core/domain/stepper"
	"github.com/andrerrcosta2/gtools/gflux/core/domain/stepper/step"
)

// Advanceable creates a basic stepper.Advanceable over a step source.
//
// Steps are executed in source order. StepTo executes all broadcasters due at or before
// target and then updates AtNow to target.
func Advanceable[S step.Step[T], T constraints.Sortable[T]](src step.Queue[S, T]) stepper.Advanceable[T] {
	return &advanceable[S, T]{
		source: src,
	}
}

type advanceable[S step.Step[T], T constraints.Sortable[T]] struct {
	now    T
	source step.Queue[S, T]
}

// Step executes one pending step and advances Now to that step position.
func (s *advanceable[S, T]) Step() (executed bool, errs error) {
	next, ok := s.source.Next()
	if !ok {
		return false, nil
	}

	s.now = next.At()

	return true, next.Exec()
}

// StepTo executes all broadcasters due at or before target and sets Now to target.
func (s *advanceable[S, T]) StepTo(target T) (executed bool, errs []error) {
	for {
		next, ok := s.source.Peek()
		if !ok || target.Less(next.At()) {
			break
		}

		exec, err := s.Step()
		if exec {
			executed = true
		}
		if err != nil {
			errs = append(errs, err)
		}
	}

	s.now = target
	return executed, errs
}

// Now returns the current logical position.
func (s *advanceable[S, T]) Now() T {
	return s.now
}

// Forwarder creates a stepper.Advanceable for forward-linked broadcasters.
//
// After executing a step, if next.Next() exists, that successor is pushed back
// to the source for future execution.
func Forwarder[
	S interface {
		step.Step[T]
		Next() (S, bool)
	},
	T constraints.Sortable[T]](src step.Queue[S, T]) stepper.Advanceable[T] {
	return &forwarder[S, T]{
		source: src,
	}
}

type forwarder[
	S interface {
		step.Step[T]
		Next() (S, bool)
	},
	T constraints.Sortable[T],
] struct {
	counter T
	source  step.Queue[S, T]
}

// Step executes one pending step, then enqueues its successor if present.
func (s *forwarder[S, T]) Step() (bool, error) {
	next, ok := s.source.Next()
	if !ok {
		return false, nil
	}

	s.counter = next.At()
	err := next.Exec()

	if again, ok := next.Next(); ok {
		s.source.Push(again)
	}

	return true, err
}

// StepTo executes all due broadcasters and sets Now to target.
func (s *forwarder[S, T]) StepTo(target T) (executed bool, errs []error) {
	for {
		next, ok := s.source.Peek()
		if !ok || target.Less(next.At()) {
			break
		}

		exec, err := s.Step()
		if !executed && exec {
			executed = true
		}
		if err != nil {
			errs = append(errs, err)
		}
	}

	s.counter = target
	return
}

// Now returns the current logical position.
func (s *forwarder[S, T]) Now() T {
	return s.counter
}
