// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package steps

import (
	"sync"

	"github.com/andrerrcosta2/gtools/conc/core/future"
	"github.com/andrerrcosta2/gtools/gflux/core/domain/stepper/step"
)

type state[T any] struct {
	once  sync.Once
	done  chan struct{}
	value T
	err   error
}

func Future[S step.Step[E], E any](step S) (step.Future[E], future.Promise[E]) {
	state := &state[E]{
		done: make(chan struct{}),
	}
	return &futureStep[S, E]{
			step:  step,
			state: state,
		}, &promiseStep[S, E]{
			state: state,
		}
}

type futureStep[S step.Step[E], E any] struct {
	step  S
	state *state[E]
}

func (f *futureStep[S, E]) At() E {
	return f.step.At()
}

func (f *futureStep[S, E]) Exec() error {
	f.wait()
	return f.step.Exec()
}

func (f *futureStep[S, E]) Wait() {
	f.wait()
}

func (f *futureStep[S, E]) wait() {
	<-f.state.done
}

func (f *futureStep[S, E]) Ready() bool {
	select {
	case <-f.state.done:
		return true
	default:
		return false
	}
}

func (f *futureStep[S, E]) Try() (exec error, rejected error, ready bool) {
	return f.step.Exec(), f.state.err, f.Ready()
}

var _ future.Future[error] = (*futureStep[step.Step[int], int])(nil)

type promiseStep[S step.Step[E], E any] struct {
	state *state[E]
}

func (f *promiseStep[S, E]) Resolve(value E) {
	f.state.value = value
	close(f.state.done)
}

func (f *promiseStep[S, E]) Reject(err error) {
	f.state.err = err
	close(f.state.done)
}
