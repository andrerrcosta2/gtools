// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package runners

import (
	"errors"
	"github.com/andrerrcosta2/gtools/core/domain/functions"
	"github.com/andrerrcosta2/gtools/gflux/core/tasks"
	"sync"
)

// Sequential returns an instance of Serial.
// It is used to run unordered functions in the order their key states.
func Sequential() Serial[int] {
	return &sequentialRunner{
		stack: make(map[int]functions.Runnable),
	}
}

type sequentialRunner struct {
	mtx   sync.Mutex
	next  int
	stack map[int]functions.Runnable
}

// Run adds a runnable function to be executed in the order given by the provided key.
// If the key is less or compare than the next key to be executed, the function is executed.
// Otherwise, the function is stored in a stack to be executed when its turn comes.
func (o *sequentialRunner) Run(f functions.Runnable, order int) {
	o.mtx.Lock()
	defer o.mtx.Unlock()
	if o.next < order {
		o.stack[order] = f
		return
	}
	o.run(f)
}

// run runs the given function and then executes the next one in the stack.
func (o *sequentialRunner) run(f functions.Runnable) {
	// Before the given function
	f()
	// Increase the next key to be executed
	o.next++
	// Check if there is a function in the stack to be executed
	if runnable, ok := o.stack[o.next]; ok {
		// Delete the function from the stack
		delete(o.stack, o.next)
		// Increase the next key to be executed again
		o.next++
		// Before the next one in the stack
		o.run(runnable)
	}
}

func StepAndRelease[T tasks.ReleasableOf]() StepReleaser[string] {
	return &stepAndReleaseRunner[T]{
		tasks: map[string]int{},
		stack: []tasks.ReleasableOf{},
	}
}

type stepAndReleaseRunner[T tasks.ReleasableOf] struct {
	mtx   sync.Mutex
	tasks map[string]int
	stack []tasks.ReleasableOf
}

func (r *stepAndReleaseRunner[S]) Step(step string, fn functions.Supplier[error]) {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	idx := len(r.stack)
	r.stack = append(r.stack, tasks.NewStep(step, fn))
	r.tasks[step] = idx
}

func (r *stepAndReleaseRunner[S]) Release(s string) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	idx, ok := r.tasks[s]
	if !ok {
		return errors.New("step not found")
	}
	r.stack[idx].Release()

	if idx == 0 {
		for i, step := range r.stack {
			if err := step.Run(); err != nil {
				r.stack = r.stack[:i]
				return err
			}
			delete(r.tasks, step.Unique())
		}
	}
	r.stack = []tasks.ReleasableOf{}
	r.tasks = map[string]int{}
	return nil
}
