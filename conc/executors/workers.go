// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package executors

import (
	"sync"

	"github.com/andrerrcosta2/gtools/conc/core/executor"
	"github.com/andrerrcosta2/gtools/core/domain/function"
)

type WorkerPool[C function.Consumer[T], T any] struct {
	tasks      chan T
	workers    int
	taskBuffer int
	running    executor.QueueExecutor[T]
	startOnce  sync.Once
	wg         sync.WaitGroup
	transfer   sync.WaitGroup
}

// Deprecated: use github.com/andrerrcosta2/gtools/conc/executors.NewPool.
func NewWorkerPool[C function.Consumer[T], T any](numWorkers int, taskBuffer int) *WorkerPool[C, T] {
	return &WorkerPool[C, T]{
		tasks:      make(chan T, taskBuffer),
		workers:    numWorkers,
		taskBuffer: taskBuffer,
	}
}

// Start initializes the workers to be ready to receive task
func (wp *WorkerPool[C, T]) Start(c C) {
	wp.startOnce.Do(func() {
		wp.running = NewPool[T](wp.workers, wp.taskBuffer, func(task T) { c(task) })
		wp.transfer.Add(1)
		go func() {
			defer wp.transfer.Done()
			for task := range wp.tasks {
				_ = wp.running.Submit(task)
			}
			_ = wp.running.Close()
			wp.running.Wait()
		}()
	})
}

// Submit adds a task to the pool
func (wp *WorkerPool[C, T]) Submit(task T) {
	wp.wg.Add(1)
	wp.tasks <- task
}

// Done marks a task as completed
func (wp *WorkerPool[C, T]) Done() {
	wp.wg.Done()
}

// Wait blocks until all task are complete
func (wp *WorkerPool[C, T]) Wait() {
	wp.wg.Wait()
	close(wp.tasks)
	wp.transfer.Wait()
}
