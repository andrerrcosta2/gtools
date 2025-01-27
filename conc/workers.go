// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package conc

import (
	"github.com/andrerrcosta2/gtools/core/domain/functions"
	"sync"
)

type WorkerPool[C functions.Consumer[T], T any] struct {
	tasks   chan T
	workers int
	wg      sync.WaitGroup
}

func NewWorkerPool[C functions.Consumer[T], T any](numWorkers int, taskBuffer int) *WorkerPool[C, T] {
	return &WorkerPool[C, T]{
		tasks:   make(chan T, taskBuffer),
		workers: numWorkers,
	}
}

// Start initializes the workers to be ready to receive tasks
func (wp *WorkerPool[C, T]) Start(c C) {
	for i := 0; i < wp.workers; i++ {
		go func() {
			for task := range wp.tasks {
				c(task)
			}
		}()
	}
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

// Wait blocks until all tasks are complete
func (wp *WorkerPool[C, T]) Wait() {
	wp.wg.Wait()
	close(wp.tasks)
}
