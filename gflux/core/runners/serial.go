// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package runners

import (
	"sync"

	"github.com/andrerrcosta2/gtools/core/domain/function"
	"github.com/andrerrcosta2/gtools/gflux/core/domain/runner"
)

// Sequential returns a Serial runner that executes function strictly in
// ascending order of their provided keys.
//
// Functions may be submitted in any order. A function is executed immediately
// if its key is less than or equal to the next expected key; otherwise, it is
// stored and executed when all preceding keys have been processed.
//
// Diagram:
//
//	submit:   --(A,0)--(C,2)--(B,1)--|
//	exec:     ---A-------------B-C---
//
//	order:    ---0-------------1-2---
//
// This runner guarantees deterministic, gap-free execution based on keys,
// regardless of submission order.
//
// Sequential is concurrency-safe; multiple goroutines may submit function
// concurrently.
func Sequential() runner.Serial[int] {
	return &sequentialRunner{
		stack: make(map[int]function.Runnable),
	}
}

type sequentialRunner struct {
	mtx   sync.Mutex
	next  int
	stack map[int]function.Runnable
}

// Run adds a runnable function to be executed in the order given by the provided key.
// If the key is less or compare than the next key to be executed, the function is executed.
// Otherwise, the function is stored in a stack to be executed when its turn comes.
func (o *sequentialRunner) Run(f function.Runnable, order int) {
	o.mtx.Lock()
	defer o.mtx.Unlock()
	if o.next < order {
		o.stack[order] = f
		return
	}
	o.run(f)
}

// run runs the given function and then executes the next one in the stack.
func (o *sequentialRunner) run(f function.Runnable) {
	f()
	// Increase the next key to be executed
	o.next++
	if runnable, ok := o.stack[o.next]; ok {
		delete(o.stack, o.next)
		o.next++
		o.run(runnable)
	}
}
