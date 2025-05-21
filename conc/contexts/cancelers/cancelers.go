// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package cancelers

import (
	"context"
	"github.com/andrerrcosta2/gtools/core/domain/gtools"
	"sync"
)

type Conditional interface {
	// When - expects a cancelers.Pipe as strategy
	// to cancel itself.
	When(Pipe)
	// Now - Cancels its owner and all its nested contexts immediately.
	Now()
}

// ConditionalCanceler returns a new canceler that can be used to cancel
// a context and all its nested contexts using the strategy from a
// cancelers.Pipe.
//
// The returned canceler can be used to cancel the context and all its
// nested contexts using the When method.
func ConditionalCanceler(cancelFunc context.CancelFunc, stream func() gtools.Stream[context.Context]) Conditional {
	return &condCancel{
		stream:     stream,
		cancelFunc: cancelFunc,
	}
}

type condCancel struct {
	stream     func() gtools.Stream[context.Context]
	once       sync.Once
	cancelFunc context.CancelFunc
}

// When - Cancels its owner context and its nested contexts
// using the strategy from the cancelers.Pipe.
//
// The given canceler is called with the context.CancelFunc and a slice
// of context.Context. The canceler should cancel the context using the
// context.CancelFunc and the given slice of context.Context.
func (c *condCancel) When(pipe Pipe) {
	pipe.Cancel(c.cancel, c.stream)
}

// Now - Cancels its owner and all its nested contexts immediately.
func (c *condCancel) Now() {
	c.cancel()
}

// cancel - Cancels its owner and all its nested contexts immediately.
func (c *condCancel) cancel() {
	c.once.Do(func() {
		c.cancelFunc()
	})
}

// SynchronizedCanceler returns a new canceler that can be used to cancel
// a context and all its nested contexts synchronously.
//
// The returned canceler can be used to cancel the context and all its
// nested contexts using the When method.
func SynchronizedCanceler(cancelFunc context.CancelFunc, stream func() gtools.Stream[context.Context]) Conditional {
	return &syncCondCancel{
		stream:     stream,
		cancelFunc: cancelFunc,
	}
}

type syncCondCancel struct {
	wg         sync.WaitGroup
	stream     func() gtools.Stream[context.Context]
	once       sync.Once
	cancelFunc context.CancelFunc
}

// When - Cancels its owner context and its nested contexts
// using the strategy from the cancelers.Pipe.
//
// The given canceler is called with the context.CancelFunc and a slice
// of context.Context. The canceler should cancel the context using the
// context.CancelFunc and the given slice of context.Context.
func (c *syncCondCancel) When(pipe Pipe) {
	pipe.Cancel(c.cancel, c.stream)
}

// Now - Cancels its owner and all its nested contexts immediately.
// It's equivalent to calling the When method with a canceler that cancels the context immediately.
func (c *syncCondCancel) Now() {
	c.cancel()
}

func (c *syncCondCancel) cancel() {
	c.once.Do(func() {
		c.cancelFunc()
	})
}

type canceler interface {
	cancel()
	Done() <-chan struct{}
}
