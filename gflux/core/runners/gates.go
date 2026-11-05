// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package runners

import (
	"sync"
	"sync/atomic"
	"time"

	"github.com/andrerrcosta2/gtools/core/domain/function/runnables"
	"github.com/andrerrcosta2/gtools/gflux/core/domain/runner"
)

func AsyncSerialGate(bufferSize int) runner.Gate {
	g := &asyncSerialGate{
		queue: make(chan func(done func()), bufferSize),
		stop:  make(chan struct{}),
	}

	go g.run()

	return g
}

type asyncSerialGate struct {
	queue chan func(done func())
	stop  chan struct{}
	once  sync.Once
}

func (g *asyncSerialGate) run() {
	for {
		select {
		case task, ok := <-g.queue:
			if !ok {
				return
			}

			done := make(chan struct{})

			task(func() {
				close(done)
			})

			<-done

		case <-g.stop:
			return
		}
	}
}

func (g *asyncSerialGate) Run(task func(done func())) bool {
	select {
	case g.queue <- task:
		return true
	case <-g.stop: // runs when stop is closed
		return false
	}
}

func (g *asyncSerialGate) Stop() {
	g.once.Do(func() {
		close(g.stop)
		close(g.queue)
	})
}

type asyncThrottleGate struct {
	window        time.Duration
	mu            sync.Mutex
	nextAllowedAt time.Time
	stopped       bool
}

func AsyncThrottleGate(window time.Duration) runner.Gate {
	return &asyncThrottleGate{
		window: window,
	}
}

func (g *asyncThrottleGate) Run(task func(_ func())) bool {
	g.mu.Lock()

	if g.stopped {
		g.mu.Unlock()
		return false
	}

	now := time.Now()

	if now.Before(g.nextAllowedAt) {
		g.mu.Unlock()
		return false
	}

	g.nextAllowedAt = now.Add(g.window)
	g.mu.Unlock()

	go task(runnables.Void)

	return true
}

func (g *asyncThrottleGate) Stop() {
	g.mu.Lock()
	g.stopped = true
	g.mu.Unlock()
}

// ConcExclusiveGate returns a concurrency-safe Gate that allows at most one
// active task at a time.
//
// Multiple goroutines may call Exec concurrently, but only the first call
// that successfully acquires the gate is allowed to execute. All other
// concurrent or overlapping calls fail immediately.
//
// No queuing or blocking occurs.
//
// Diagram:
//
//	submit:   --A--B--C--|
//	exec(A):     ----a--|
//	exec(B):          x
//	exec(C):               ----c--|
//
//	output:    ----a----------c--|
//
// Range task receives a `done` callback and must call it to release the gate.
// Once released, a new task may be accepted.
//
// Calling Stop permanently disables the gate, causing all future calls
// to Exec to fail.
func ConcExclusiveGate() runner.Gate {
	return &concExclusiveGate{}
}

type concExclusiveGate struct {
	deny atomic.Bool
}

func (g *concExclusiveGate) Run(fn func(done func())) bool {
	if !g.deny.CompareAndSwap(false, true) {
		return false
	}

	fn(func() {
		g.deny.Store(false)
	})

	return true
}

func (g *concExclusiveGate) Stop() {
	g.deny.Store(true)
}

// ConcSerialGate returns a Gate that serializes execution of asynchronous task.
//
// Multiple goroutines may call Exec concurrently, but task are executed
// strictly one at a time, in FIFO order.
//
// Diagram:
//
//	submit:   --A--B--C--|
//	exec(A):     ----a--|
//	exec(B):              ----b--|
//	exec(C):                       ----c--|
//
//	output:    ----a------b------c--|
//
// Range task receives a `done` callback and must call it to allow the next
// queued task to run. The gate does not spawn goroutines by itself; execution
// advances only when task signal completion.
//
// Calling Stop prevents new task from being accepted and clears any queued
// task. The currently running task is allowed to finish, but no further
// task will execute.
func ConcSerialGate() runner.Gate {
	return &concSerialGate{}
}

type concSerialGate struct {
	mu         sync.Mutex
	queue      []func(done func())
	running    bool
	terminated bool
}

func (g *concSerialGate) Run(task func(done func())) bool {
	g.mu.Lock()

	if g.terminated {
		g.mu.Unlock()
		return false
	}

	g.queue = append(g.queue, task)

	if g.running {
		g.mu.Unlock()
		return true
	}

	g.running = true
	g.mu.Unlock()

	g.runNext()
	return true
}

func (g *concSerialGate) runNext() {
	var task func(done func())

	g.mu.Lock()

	if g.terminated {
		g.queue = nil
		g.running = false
		g.mu.Unlock()
		return
	}

	if len(g.queue) == 0 {
		g.running = false
		g.mu.Unlock()
		return
	}

	task = g.queue[0]
	g.queue = g.queue[1:]
	g.mu.Unlock()

	task(func() {
		g.runNext()
	})
}

func (g *concSerialGate) Stop() {
	g.mu.Lock()
	if g.terminated {
		g.mu.Unlock()
		return
	}

	g.terminated = true
	g.queue = nil
	g.mu.Unlock()
}

// ExclusiveGate returns a Gate that allows at most one active task at a time.
//
// If a task is already running, calls to Exec fail immediately and the task
// is dropped. No queuing or blocking occurs.
//
// Diagram:
//
//	submit:   --A--B--C--|
//	exec(A):     ----a--|
//	exec(B):          x
//	exec(C):               ----c--|
//
//	output:    ----a----------c--|
//
// Range task receives a `done` callback and must call it to release the gate.
// Once released, a new task may be accepted.
//
// Calling Stop permanently disables the gate, causing all future calls
// to Exec to fail.
//
// Note: ExclusiveGate is not concurrency-safe and assumes serialized access
// to Exec and Stop.
func ExclusiveGate() runner.Gate {
	return &exclusiveGate{}
}

type exclusiveGate struct {
	active bool
}

func (g *exclusiveGate) Run(fn func(done func())) bool {
	if g.active {
		return false
	}

	fn(func() {
		g.active = false
	})

	return true
}

func (g *exclusiveGate) Stop() {
	g.active = true
}

// SerialGate returns a Gate that serializes execution of asynchronous task
// without concurrency safety.
//
// Tasks are executed strictly one at a time, in FIFO order, but calls to Exec
// must be externally serialized. Concurrent calls may cause data races.
//
// Diagram:
//
//	submit:   --A--B--C--|
//	exec(A):     ----a--|
//	exec(B):              ----b--|
//	exec(C):                       ----c--|
//
//	output:    ----a------b------c--|
//
// Range task receives a `done` callback and must call it to allow the next
// queued task to run. The gate does not spawn goroutines by itself; execution
// advances only when task signal completion.
//
// Calling Stop prevents new task from being accepted and clears any queued
// task. The currently running task is allowed to finish, but no further
// task will execute.
//
// Note: SerialGate is not concurrency-safe and assumes serialized access
// to Exec and Stop.
func SerialGate() runner.Gate {
	return &concSerialGate{}
}

type serialGate struct {
	queue      []func(done func())
	running    bool
	terminated bool
}

func (g *serialGate) Run(task func(done func())) bool {
	if g.terminated {
		return true
	}

	g.queue = append(g.queue, task)

	if g.running {
		return true
	}

	g.running = true

	g.runNext()
	return true
}

func (g *serialGate) runNext() {
	var task func(done func())

	if g.terminated {
		g.queue = nil
		g.running = false
		return
	}

	if len(g.queue) == 0 {
		g.running = false
		return
	}

	task = g.queue[0]
	g.queue = g.queue[1:]

	task(func() {
		g.runNext()
	})
}

func (g *serialGate) Stop() {
	if g.terminated {
		return
	}
	g.terminated = true
	g.queue = nil
}

// TerminalGate returns a Gate that allows task to run freely until it is stopped.
//
// TerminalGate does not serialize, restrict, or coordinate task execution.
// All calls to Exec succeed and task are executed immediately while the gate
// is active.
//
// Diagram:
//
//	submit:   --A--B--C--|
//	exec:        a  b  c
//
//	output:    --a--b--c--|
//
// Range task receives a `done` callback, but calling it has no effect on the
// gate. The callback exists only to satisfy the Gate contract.
//
// Calling Stop permanently terminates the gate. After termination, all
// future calls to Exec fail.
func TerminalGate() runner.Gate {
	return &terminalGate{}
}

type terminalGate struct {
	terminated atomic.Bool
}

func (g *terminalGate) Run(task func(done func())) bool {
	if g.terminated.Load() {
		return false
	}

	// TerminalGate does NOT serialize or restrict
	task(func() {})
	return true
}

func (g *terminalGate) Stop() {
	g.terminated.Store(true)
}

func ThrottleGate(window time.Duration) runner.Gate {
	return &throttleGate{
		window: window,
	}
}

type throttleGate struct {
	window        time.Duration
	mu            sync.Mutex
	nextAllowedAt time.Time
	stopped       bool
}

func (g *throttleGate) Run(task func(_ func())) bool {
	g.mu.Lock()

	if g.stopped {
		g.mu.Unlock()
		return false
	}

	now := time.Now()

	if now.Before(g.nextAllowedAt) {
		g.mu.Unlock()
		return false
	}

	g.nextAllowedAt = now.Add(g.window)
	g.mu.Unlock()

	task(runnables.Void)
	return true
}

func (g *throttleGate) Stop() {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.stopped = true
}
