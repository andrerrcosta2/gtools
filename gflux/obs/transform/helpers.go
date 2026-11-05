// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package transform

import (
	"sync/atomic"
)

type termination struct {
	closed     atomic.Bool
	terminated atomic.Bool
}

func (t *termination) Close() {
	t.closed.Store(true)
}

func (t *termination) CloseIfOpen() bool {
	return t.closed.CompareAndSwap(false, true)
}

func (t *termination) IsClosed() bool {
	return t.closed.Load()
}

func (t *termination) IsClosedOrTerminated() bool {
	return t.terminated.Load() || t.closed.Load()
}

func (t *termination) Terminate() {
	t.terminated.Store(true)
}

func (t *termination) TerminateIfActive() bool {
	return t.terminated.CompareAndSwap(false, true)
}

func (t *termination) TerminateAndClose() {
	t.terminated.Store(true)
	t.closed.Store(true)
}

func (t *termination) IsTerminated() bool {
	return t.terminated.Load()
}
