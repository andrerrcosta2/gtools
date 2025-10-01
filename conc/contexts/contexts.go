// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package contexts

import (
	"context"
	"github.com/andrerrcosta2/gtools/conc/contexts/cancelers"
	"github.com/andrerrcosta2/gtools/core/domain/gtools"
	"sync"
)

type Cancellable interface {
	// Done returns a channel that is closed when the context is canceled.
	//
	// It's thread-safe and can be used concurrently.
	Done() <-chan struct{}
}

type ConditionallyCanceled interface {
	context.Context
	gtools.Streamable[context.Context]
	// AddChild adds a child context to the list.
	//
	// It's thread-safe and can be used concurrently.
	AddChild(child context.Context)
}

type Synchronizable interface {
	ConditionallyCanceled
	// Add adds the given value be conditionally canceled.
	//
	// It's thread-safe and can be used concurrently.
	Add(add int)
	// Wait blocks until the context is canceled.
	//
	// It's thread-safe and can be used concurrently.
	Wait()
}

type SynchronizableContext interface {
	context.Context
	// Add adds the given value be conditionally canceled.
	//
	// It's thread-safe and can be used concurrently.
	Add(add int)
	// Release decrease the wait group counter.
	//
	// It's thread-safe and can be used concurrently.
	Release()
	// Wait blocks until the context is canceled.
	//
	// It's thread-safe and can be used concurrently.
	Wait()
}

// WithConditionalCancel returns a new context that can be canceled using the
// cancelers.ConditionalCanceler.
//
// The context is created with a canceler that can be used to cancel the context
// conditionally. The canceler can be used to cancel the context when a certain
// condition is met.
func WithConditionalCancel(parent context.Context) (ctx context.Context, cancel cancelers.Conditional) {
	// Return the context and the canceler
	return withConditionalCancel(parent)
}

// withConditionalCancel creates a new context that can be canceled using the
// cancelers.ConditionalCanceler.
//
// The context is created with a canceler that can be used to cancel the context
// conditionally. The canceler can be used to cancel the context when a certain
// condition is met.
func withConditionalCancel(parent context.Context) (ctx *condCanceledCtx, cancel cancelers.Conditional) {
	if parent == nil {
		panic("cannot create context from nil parent")
	}

	// Leveraging the cancelCtx from the standard library to avoid
	// breaking the default cancellation behavior
	cancelCtx, cancelFunc := context.WithCancel(parent)

	// Create a new context with conditional cancel
	cond := &condCanceledCtx{
		Context: cancelCtx,
	}

	// add the context as child if its parent is conditionally canceled
	if s, ok := parent.(ConditionallyCanceled); ok {
		s.AddChild(cond)
	}

	// Create a new canceler
	cancel = cancelers.ConditionalCanceler(cancelFunc, cond.Stream)

	return cond, cancel
}

type condCanceledCtx struct {
	context.Context
	mtx      sync.RWMutex
	add      int
	children map[context.Context]struct{}
}

// AddChild adds a child context to the list
// It replaces the child it already exists
func (c *condCanceledCtx) AddChild(child context.Context) {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	if c.children == nil {
		c.children = make(map[context.Context]struct{})
	}
	c.children[child] = struct{}{}
}

func (c *condCanceledCtx) Stream() gtools.Stream[context.Context] {
	children := make(chan context.Context)
	go func() {
		defer close(children)
		for child := range c.children {
			children <- child
		}
	}()
	return children
}

// Add adds the given value to the context.
func (c *condCanceledCtx) Add(add int) {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	c.add += add
}

// String returns a string representation of the context.
//
// It returns the name of the parent context followed by ".WithConditionalCancel".
func (c *condCanceledCtx) String() string {
	return contextName(c.Context) + ".WithConditionalCancel"
}

var _ ConditionallyCanceled = (*condCanceledCtx)(nil)

// Synchronized creates a new synchronized context that can be used to wait for
// ops to be done and to add delays to the context.
func Synchronized(parent context.Context) (ctx SynchronizableContext, cancel cancelers.Conditional) {
	return synchronized(parent)
}

// synchronized creates a new synchronized context that can be used to wait for
// ops to be done and to add delays to the context.
func synchronized(parent context.Context) (*synchronizedContext, cancelers.Conditional) {
	if parent == nil {
		panic("cannot create context from nil parent")
	}

	// Leveraging the cancelCtx from the standard library to avoid
	// breaking the default cancellation behavior
	cancelCtx, cancelFunc := context.WithCancel(parent)

	// Initialize synchronizedContext
	syncCtx := &synchronizedContext{
		condCanceledCtx: condCanceledCtx{
			Context: cancelCtx,
		},
	}

	// If the parent is conditionally canceled, register this as a child
	if cc, ok := parent.(ConditionallyCanceled); ok {
		cc.AddChild(syncCtx)
	}

	// Create a new canceler
	cancel := cancelers.SynchronizedCanceler(cancelFunc, syncCtx.Stream)

	return syncCtx, cancel
}

type synchronizedContext struct {
	condCanceledCtx
	mtx      sync.RWMutex
	wait     sync.WaitGroup
	children map[context.Context]struct{}
}

// AddChild method to register child contexts
func (c *synchronizedContext) AddChild(ctx context.Context) {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	if c.children == nil {
		c.children = make(map[context.Context]struct{})
	}
	c.children[ctx] = struct{}{}
}

func (c *synchronizedContext) Stream() gtools.Stream[context.Context] {
	children := make(chan context.Context)

	go func() {
		defer close(children)
		for child := range c.children {
			children <- child
		}
	}()

	return children
}

// Add adds the given value to the wait group.
//
// It increments the wait group counter by the given value.
// It's thread-safe and can be used concurrently.
func (c *synchronizedContext) Add(add int) {
	c.wait.Add(add)
}

func (c *synchronizedContext) Release() {
	c.wait.Done()
}

// Wait waits until the wait group counter is zero.
//
// It blocks until the wait group counter is zero.
// It's thread-safe and can be used concurrently.
func (c *synchronizedContext) Wait() {
	c.wait.Wait()
}

// String returns a string representation of the context.
//
// It returns the name of the parent context followed by ".Synchronized".
//
// It's thread-safe and can be used concurrently.
func (c *synchronizedContext) String() string {
	return contextName(c.Context) + ".Synchronized"
}

var _ ConditionallyCanceled = (*synchronizedContext)(nil)
var _ context.Context = (*synchronizedContext)(nil)

// New creates a new context that can be used to wait for ops to be
