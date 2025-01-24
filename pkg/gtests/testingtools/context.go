// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package testingtools

import (
	"context"
	"github.com/andrerrcosta2/gtools/core/gtools/functions"
	"github.com/andrerrcosta2/gtools/gtests/testingtools/conc/syncs"
)

type asyncContextToolsLite struct {
}

// RunAfterBlocking runs the given function and blocks until the given context is done.
func (t *asyncContextToolsLite) RunAfterBlocking(fn functions.Runnable, ctx context.Context) {
	<-ctx.Done()
	go fn()
}

// RunBlocking runs the given function and blocks until the given context is done.
// It's thread-safe and can be used concurrently.
func (t *asyncContextToolsLite) RunBlocking(fn functions.Runnable, ctx context.Context) {
	go fn()
	<-ctx.Done()
}

// AfterAsync runs the given function and blocks asynchronously until the given context is done.
// This function can be replaced by a simple goroutine and may seem useless, but is useful for testing
// which depends on contexts. It also enforces this test kit with semantic coupling.
func (t *asyncContextToolsLite) AfterAsync(fn functions.Runnable, ctx context.Context) {
	go func() {
		<-ctx.Done()
		fn()
	}()
}

// AsyncBefore runs the given function and blocks asynchronously until the given context is done.
// This function can be replaced by a simple goroutine and may seem useless, but is useful for testing
// which depends on contexts. It also enforces this test kit with semantic coupling.
func (t *asyncContextToolsLite) AsyncBefore(fn functions.Runnable, ctx context.Context) {
	go func() {
		fn()
		<-ctx.Done()
	}()
}

func (t *asyncContextToolsLite) Semaphore(maxConcurrent int) *syncs.ChannelSemaphore {
	return &syncs.ChannelSemaphore{
		ch: make(chan struct{}, maxConcurrent),
	}
}

type concurrentContextToolsLite struct {
}

// After runs the given function and blocks until the given context is done.
func (t *concurrentContextToolsLite) After(fn functions.Runnable, ctx context.Context) {
	<-ctx.Done()
	go fn()
}

// Before runs the given function and blocks until the given context is done.
func (t *concurrentContextToolsLite) Before(fn functions.Runnable, ctx context.Context) {
	go fn()
	<-ctx.Done()
}

func (t *concurrentContextToolsLite) Semaphore(maxConcurrent int) *syncs.ChannelSemaphore {
	return &syncs.ChannelSemaphore{
		ch: make(chan struct{}, maxConcurrent),
	}
}
