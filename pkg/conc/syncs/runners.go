// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package syncs

import (
	"context"
	"github.com/andrerrcosta2/gtools/core/gtools/functions"
	"sync"
)

// DeferredRunner returns a Context that runs the given functions after the context is done.
func DeferredRunner[H comparable]() ContextRunner {
	return &deferredRunner[H]{stack: make(map[context.Context][]functions.Runnable)}
}

type deferredRunner[H comparable] struct {
	mtx   sync.Mutex
	stack map[context.Context][]functions.Runnable
}

// Run adds a runnable function to be executed after the context is done.
// If it is the first function for this context, a goroutine is started to
// wait for the context to finish and then run the deferred tasks.
func (o *deferredRunner[H]) Before(f functions.Runnable, ctx context.Context) {
	o.mtx.Lock()
	defer o.mtx.Unlock()

	// Initialize the stack if needed
	if o.stack == nil {
		o.stack = make(map[context.Context][]functions.Runnable)
	}

	// Append the function to the stack
	o.stack[ctx] = append(o.stack[ctx], f)

	// Only start the scheduler if this is the first function for this context
	if len(o.stack[ctx]) == 1 {
		go o.schedule(ctx)
	}
}

// schedule starts a goroutine that waits for the context to finish and then
// runs the deferred tasks for the given context.
func (o *deferredRunner[H]) schedule(ctx context.Context) {
	// Block until the context is finished
	<-ctx.Done()

	// Before the deferred tasks
	o.run(ctx)
}

// run runs the deferred tasks for the given context.
// It is called after the context is finished.
func (o *deferredRunner[H]) run(ctx context.Context) {
	o.mtx.Lock()
	runnables, ok := o.stack[ctx]
	if ok {
		delete(o.stack, ctx) // Safely remove the context from the stack
	}
	o.mtx.Unlock()

	// If the context is not found in the stack, return
	if !ok {
		return
	}

	// Before the deferred tasks
	for _, runnable := range runnables {
		runnable()
	}
}

// StepReleaseRunner This monster seems unavoidable. It is a synchronizer, orchestrator and scheduler
// all into the same component.  It may feel reasonable to separate the scheduler from the orchestrator
// and the synchronizer, but to synchronize concurrently a real time flattening of branches processed in
// parallel isn't quite a trivial task, is it?
//func StepReleaseRunner[C comparable, D comparable]() runners.Task[tasks.Releasable] {
//	return &stepReleaseRunner[C, D]{
//		dependencies: make(map[D][]tasks.Chain[*tasks.Step[D]]),
//		chains:       make(map[C]tasks.Chain[*tasks.Step[D]]),
//	}
//}
//
//type stepReleaseRunner[C comparable, D comparable] struct {
//	mtx   sync.Mutex
//	steps map[D][]tasks.Chain[*tasks.Step[D]]
//	tasks map[C]tasks.Chain[*tasks.Step[D]]
//}
//
//func (o *stepReleaseRunner[C, D]) Append(name C, ts ...*tasks.Step[D]) {
//	o.mtx.Lock()
//	defer o.mtx.Unlock()
//	if chain, ok := o.chains[name]; ok {
//		chain.Append(ts...)
//	} else {
//		c := tasks.ChainOfCondRunnable[*tasks.Step[D]](ts...)
//		o.chains[name] = c
//		o.dependencies[ts[0].Id] = append(o.dependencies[ts[0].Id], c)
//	}
//}
//
//func (o *stepReleaseRunner[C, D]) Prepend(name C, ts ...*tasks.Named[D]) {
//	o.mtx.Lock()
//	defer o.mtx.Unlock()
//	if chain, ok := o.chains[name]; ok {
//		chain.Prepend(ts...)
//	} else {
//		c := tasks.ChainOf[*tasks.Named[D]](ts...)
//		o.chains[name] = c
//		o.dependencies[ts[0].Name()] = append(o.dependencies[ts[0].Name()], c)
//	}
//}
//
//func (o *stepReleaseRunner[C, D]) Insert(name C, pos int, ts *tasks.Named[D]) error {
//	o.mtx.Lock()
//	defer o.mtx.Unlock()
//	if chain, ok := o.chains[name]; ok {
//		return chain.Insert(pos, ts)
//	} else {
//		if pos != 0 {
//			return errors.New("index out of range")
//		}
//		c := tasks.ChainOf[*tasks.Named[D]](ts)
//		o.chains[name] = c
//		o.dependencies[ts.Name()] = append(o.dependencies[ts.Name()], c)
//		return nil
//	}
//}
//
//func (o *stepReleaseRunner[C, D]) Before(chain C) {
//	for _, f := range o.stack[chain] {
//		f()
//	}
//	delete(o.stack, chain)
//}
