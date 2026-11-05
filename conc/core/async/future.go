// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package async

type state[T any] struct {
	done  chan struct{}
	value T
	err   error
}

func newFuture[T any]() (*fut[T], *prom[T]) {
	s := &state[T]{done: make(chan struct{})}
	p := &prom[T]{state: s}
	f := &fut[T]{state: s}
	return f, p
}

type fut[T any] struct {
	state *state[T]
}

func (f *fut[T]) Wait() {
	<-f.state.done
}

func (f *fut[T]) Ready() bool {
	select {
	case <-f.state.done:
		return true
	default:
		return false
	}
}

func (f *fut[T]) Try() (T, error, bool) {
	select {
	case <-f.state.done:
		return f.state.value, f.state.err, true
	default:
		var zero T
		return zero, nil, false
	}
}

type prom[T any] struct {
	state *state[T]
}

func (f *prom[T]) accept(v T, err error) {
	f.state.value = v
	f.state.err = err
	close(f.state.done)
}
