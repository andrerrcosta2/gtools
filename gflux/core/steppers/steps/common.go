// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package steps

import (
	"github.com/andrerrcosta2/gtools/gflux/core/domain/stepper/step"
)

// At creates a functional step that executes fn at logical position at.
func At[T any](fn func() error, at T) step.Step[T] {
	return functional[T]{
		exec: fn,
		at:   at,
	}
}

type functional[T any] struct {
	at   T
	exec func() error
}

func (s functional[T]) At() T       { return s.at }
func (s functional[T]) Exec() error { return s.exec() }

// WithKey wraps a step with an identification key.
func WithKey[T any](s step.Step[T], key any) step.Step[T] {
	return identifiable[T]{Step: s, key: key}
}

type identifiable[T any] struct {
	step.Step[T]
	key any
}

func (s identifiable[T]) Key() any { return s.key }

// WithPriority wraps a step with priority metadata.
func WithPriority[T any](s step.Step[T], p int) step.Step[T] {
	return priority[T]{Step: s, p: p}
}

type priority[T any] struct {
	step.Step[T]
	p int
}

func (s priority[T]) Priority() int { return s.p }

// WithRecurrency returns a concrete Recurrent that returns itself from Next()
func WithRecurrency[T any](s step.Step[T], next func(at T) T) *Recurrent[T] {
	return &Recurrent[T]{
		Step: s,
		next: next,
	}
}

type Recurrent[T any] struct {
	step.Step[T]
	next func(at T) T
}

// Next returns *Recurrent[T] - the same type
func (r *Recurrent[T]) Next() (*Recurrent[T], bool) {
	// Create a new step with the same execution function but new time
	return &Recurrent[T]{
		Step: At(r.Exec, r.next(r.At())),
		next: r.next,
	}, true
}

// WithStop wraps a step with stoppable behavior.
func WithStop[T any](s step.Step[T]) step.Stoppable[T] {
	return &stoppable[T]{Step: s}
}

type stoppable[T any] struct {
	step.Step[T]
	stopped bool
}

func (s *stoppable[T]) Stop() {
	s.stopped = true
}

func (s *stoppable[T]) Stopped() bool {
	return s.stopped
}
