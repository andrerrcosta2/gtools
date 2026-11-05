// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package cancelers

import (
	"context"
	"sync"

	"github.com/andrerrcosta2/gtools/core/domain/contracts/stream"
	"github.com/andrerrcosta2/gtools/core/domain/function"
)

type Pipe interface {
	// Cancel cancels the context using a particular pipe strategy.
	Cancel(cancel context.CancelFunc, stream func() stream.Readable[context.Context])
}

// Function returns a new canceler that calls that cancels the context using the given function.
//
// The given function takes two arguments: the context.CancelFunc and a slice of context.Context.
// The function is expected to cancel the context using the context.CancelFunc and the given slice of context.Context.
func Function(fn function.BiConsumer[context.CancelFunc, func() stream.Readable[context.Context]]) Pipe {
	return funcCanceler(fn)
}

type funcCanceler function.BiConsumer[context.CancelFunc, func() stream.Readable[context.Context]]

func (f funcCanceler) Cancel(cancel context.CancelFunc, stream func() stream.Readable[context.Context]) {
	f(cancel, stream)
}

var _ Pipe = funcCanceler(nil)

// And creates a new canceler that cancels the context when all the contexts
// done channels are closed.
// It works by adding the total number of contexts to the WaitGroup and then
// waiting for all contexts to finish. Once all is done, the cancel function
// is called.
func And(contexts ...context.Context) Pipe {
	return &andCanceler{
		contexts: contexts,
	}
}

type andCanceler struct {
	contexts []context.Context
	wait     sync.WaitGroup
}

// Cancel cancels the context when all the contexts done channels are closed.
// It works by adding the total number of contexts to the WaitGroup and then
// waiting for all contexts to finish. Once all is done, the cancel function
// is called.
func (c *andCanceler) Cancel(cancel context.CancelFunc, _ func() stream.Readable[context.Context]) {
	go c.cancel(cancel)
}

func (c *andCanceler) cancel(cancel context.CancelFunc) {
	// add the total number of contexts to the WaitGroup
	c.wait.Add(len(c.contexts))

	// start a goroutine for each context to wait for its completion
	for _, ctx := range c.contexts {
		go func(ctx context.Context) {
			// SetWaitingPoint for the context to be done
			<-ctx.Done()
			// Mark this context as completed
			c.wait.Done()
		}(ctx)
	}

	// SetWaitingPoint for all contexts to be done
	go func() {
		// SetWaitingPoint for all contexts to be done
		c.wait.Wait()
		// Once all contexts are done, cancel the main context
		cancel()
	}()
}

var _ Pipe = &andCanceler{}

// Or creates a new canceler that cancels the context when any of the contexts done channels are closed.
// It works by starting a goroutine for each context to wait for its completion and then call the cancel
// function when any of them are done.
func Or(contexts ...context.Context) Pipe {
	return &orCanceler{
		contexts: contexts,
	}
}

type orCanceler struct {
	contexts []context.Context
	once     sync.Once
}

// Cancel cancels the context when any of the contexts done channels are closed.
// It works by starting a goroutine for each context to wait for its completion
// and then call the cancel function when any of them are done.
func (c *orCanceler) Cancel(cancel context.CancelFunc, _ func() stream.Readable[context.Context]) {
	go c.cancel(cancel)
}

func (c *orCanceler) cancel(cancel context.CancelFunc) {
	// start a goroutine for each context to wait for its completion
	for _, ctx := range c.contexts {
		go func(ctx context.Context) {
			// SetWaitingPoint for the context to be done
			<-ctx.Done()
			// Cancel the main context once any context is done
			c.once.Do(cancel)
		}(ctx)
	}
}

var _ Pipe = &orCanceler{}

// All returns a Pipe that cancels the context when all the contexts done channels are closed.
// It works by adding the total number of contexts to the WaitGroup and then
// waiting for all contexts to finish. Once all is done, the cancel function
// is called.
func All() Pipe {
	return &allCanceler{}
}

type allCanceler struct {
	wait sync.WaitGroup
}

// Cancel cancels the context when all the contexts done were closed.
func (c *allCanceler) Cancel(cancel context.CancelFunc, src func() stream.Readable[context.Context]) {
	go c.cancel(cancel, src)
}

func (c *allCanceler) cancel(cancel context.CancelFunc, src func() stream.Readable[context.Context]) {
	src().Iterator()(func(ctx context.Context) bool {
		c.wait.Add(1)

		// start a goroutine to wait for the context to be done
		go func(ctx context.Context) {
			// SetWaitingPoint for the context to be done
			<-ctx.Done()
			// Mark this context as completed
			c.wait.Done()
		}(ctx)
		return true
	})

	// Once all are done, cancel
	go func() {
		// SetWaitingPoint for all contexts to be done
		c.wait.Wait()
		// Once all contexts are done, cancel the main context
		cancel()
	}()
}

var _ Pipe = &allCanceler{}

// Any returns a Pipe that cancels the context when any of the contexts done channels are closed.
// It works by starting a goroutine for each context to wait for its completion
// and then call the cancel function when any of them are done.
func Any() Pipe {
	return &anyCanceler{}
}

type anyCanceler struct {
	once sync.Once
}

// Cancel cancels the context when any of the contexts done channels are closed.
// It works by starting a goroutine for each context to wait for its completion
// and then call the cancel function when any of them are done.
func (c *anyCanceler) Cancel(cancel context.CancelFunc, stream func() stream.Readable[context.Context]) {
	go c.cancel(cancel, stream)
}

func (c *anyCanceler) cancel(cancel context.CancelFunc, src func() stream.Readable[context.Context]) {
	src().Iterator()(func(ctx context.Context) bool {
		go func(ctx context.Context) {
			<-ctx.Done()
			c.once.Do(cancel) // Cancel the main context
		}(ctx)
		return true
	})
}

var _ Pipe = &anyCanceler{}
