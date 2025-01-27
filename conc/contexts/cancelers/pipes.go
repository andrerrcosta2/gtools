// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package cancelers

import (
	"context"
	"github.com/andrerrcosta2/gtools/core/domain/functions"
	"sync"
)

type Pipe interface {
	// Cancel cancels the context using a particular pipe strategy.
	Cancel(cancel context.CancelFunc, looper functions.Looper[context.Context])
}

// Function returns a new canceler that calls that cancels the context using the given function.
//
// The given function takes two arguments: the context.CancelFunc and a slice of context.Context.
// The function is expected to cancel the context using the context.CancelFunc and the given slice of context.Context.
func Function(fn functions.BiConsumer[context.CancelFunc, functions.Looper[context.Context]]) Pipe {
	return funcCanceler(fn)
}

type funcCanceler functions.BiConsumer[context.CancelFunc, functions.Looper[context.Context]]

func (f funcCanceler) Cancel(cancel context.CancelFunc, looper functions.Looper[context.Context]) {
	f(cancel, looper)
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
func (c *andCanceler) Cancel(cancel context.CancelFunc, _ functions.Looper[context.Context]) {
	go c.cancel(cancel)
}

func (c *andCanceler) cancel(cancel context.CancelFunc) {
	// add the total number of contexts to the WaitGroup
	c.wait.Add(len(c.contexts))

	// Start a goroutine for each context to wait for its completion
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
func (c *orCanceler) Cancel(cancel context.CancelFunc, _ functions.Looper[context.Context]) {
	go c.cancel(cancel)
}

func (c *orCanceler) cancel(cancel context.CancelFunc) {
	// Start a goroutine for each context to wait for its completion
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

// Cancel cancels the context when all the contexts done channels are closed.
// It works by adding the total number of contexts to the WaitGroup and then
// waiting for all contexts to finish. Once all is done, the cancel function
// is called.
func (c *allCanceler) Cancel(cancel context.CancelFunc, looper functions.Looper[context.Context]) {
	go c.cancel(cancel, looper)
}

func (c *allCanceler) cancel(cancel context.CancelFunc, looper functions.Looper[context.Context]) {
	// SetWaitingPoint for all contexts to finish
	for ctx := range looper() {
		// Increment the WaitGroup
		c.wait.Add(1)

		// Start a goroutine to wait for the context to be done
		go func(ctx context.Context) {
			// SetWaitingPoint for the context to be done
			<-ctx.Done()
			// Mark this context as completed
			c.wait.Done()
		}(ctx)
	}

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
func (c *anyCanceler) Cancel(cancel context.CancelFunc, looper functions.Looper[context.Context]) {
	go c.cancel(cancel, looper)
}

func (c *anyCanceler) cancel(cancel context.CancelFunc, looper functions.Looper[context.Context]) {
	// Start a goroutine for each context to wait to check for any completion
	for ctx := range looper() {
		go func(ctx context.Context) {
			// SetWaitingPoint for the context to be done
			<-ctx.Done()
			// Cancel the main context
			c.once.Do(cancel)
		}(ctx)
	}
}

var _ Pipe = &anyCanceler{}
