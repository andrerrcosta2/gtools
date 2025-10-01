// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package dispatcher

import (
	"errors"
	"github.com/andrerrcosta2/gtools/core/domain/functions"
	"github.com/andrerrcosta2/gtools/core/format/fmx"
	"sync"
	"sync/atomic"
	"time"
)

var (
	ClosedDispatcherError = errors.New("dispatcher is closed")
)

type Handler[T any] functions.SafeConsumer[T]
type Listener[T any] functions.BiConsumer[T, error]

type (
	Cold[T any] interface {
		Dispatcher[T]
		Start() // Starts the dispatcher
	}

	Dispatcher[T any] interface {
		Dispatch(item T) error       // Dispatch sends an item to be handled.
		Close()                      // Close gracefully stops the dispatcher.
		IsClosed() bool              // IsClosed checks whether the dispatcher is closed.
		Listen(listener Listener[T]) // Sets a listener to react to dispatch ops
		QueueSize() int
	}

	Strategy[T any] interface {
		Feedback(meta *Meta[T])             // Called after each dispatch attempt with the result.
		BeforeDispatch(meta *Meta[T]) error // Called before dispatching. Can block or influence dispatch.
		OnStart()
		OnStop()
	}
)

func noOpListener[T any](_ T, _ error) {}

type noOpStrategy[T any] struct{}

func (noOpStrategy[T]) BeforeDispatch(_ *Meta[T]) error { return nil }
func (noOpStrategy[T]) Feedback(_ *Meta[T])             {}
func (noOpStrategy[T]) OnStart()                        {}
func (noOpStrategy[T]) OnStop()                         {}

// Hot creates a new always-on hotDispatcher.
// A Hot dispatcher is a dispatcher that starts processing as soon as it is created.
func Hot[T any](handler Handler[T], listener Listener[T], maxConcurrency int) Dispatcher[T] {
	if maxConcurrency <= 0 {
		maxConcurrency = 1
	}

	if listener == nil {
		listener = noOpListener[T]
	}

	d := &hotDispatcher[T]{
		ch:         make(chan T),
		handle:     handler,
		closed:     make(chan struct{}),
		listen:     listener,
		maxWorkers: maxConcurrency,
	}

	d.wg.Add(d.maxWorkers)
	for i := 0; i < d.maxWorkers; i++ {
		go d.run()
	}
	return d
}

type hotDispatcher[T any] struct {
	ch         chan T
	handle     Handler[T]
	listen     Listener[T]
	closed     chan struct{}
	once       sync.Once
	wg         sync.WaitGroup
	maxWorkers int
}

// Dispatch sends an item to be handled.
// Returns false if dispatcher is closed.
func (d *hotDispatcher[T]) Dispatch(item T) error {
	select {
	case d.ch <- item:
		return nil
	case <-d.closed:
		return ClosedDispatcherError
	}
}

// Close gracefully stops the dispatcher.
func (d *hotDispatcher[T]) Close() {
	d.once.Do(func() {
		close(d.closed)
		go func() {
			for item := range d.ch {
				d.listen(item, ClosedDispatcherError)
			}
			close(d.ch)
			d.wg.Wait()
		}()
	})
}

// IsClosed checks whether the dispatcher is closed.
func (d *hotDispatcher[T]) IsClosed() bool {
	select {
	case <-d.closed:
		return true
	default:
		return false
	}
}

func (d *hotDispatcher[T]) Listen(listener Listener[T]) {
	d.listen = listener
}

func (d *hotDispatcher[T]) QueueSize() int {
	return len(d.ch)
}

// run starts the processing loop.
func (d *hotDispatcher[T]) run() {
	defer d.wg.Done()
	for {
		select {
		case item, ok := <-d.ch:
			if !ok {
				return
			}
			var err error
			func() {
				defer func() {
					if r := recover(); r != nil {
						err = fmx.Errorf("handler panic: %v", r)
					}
				}()
				err = d.handle(item)
			}()
			d.listen(item, err)
		case <-d.closed:
			return
		}
	}
}

// WithBackpressure creates a new hot dispatcher with a buffered channel of the given size.
func WithBackpressure[T any](bufferSize int, handler Handler[T], listener Listener[T], maxConcurrency int) Dispatcher[T] {
	if maxConcurrency <= 0 {
		maxConcurrency = 1
	}

	if listener == nil {
		listener = noOpListener[T]
	}

	d := &backPressured[T]{
		hotDispatcher: hotDispatcher[T]{
			ch:         make(chan T, bufferSize),
			handle:     handler,
			listen:     listener,
			closed:     make(chan struct{}),
			maxWorkers: maxConcurrency,
		},
	}

	d.wg.Add(d.maxWorkers)
	for i := 0; i < d.maxWorkers; i++ {
		go d.run()
	}

	return d
}

type backPressured[T any] struct {
	hotDispatcher[T]
}

type ShouldRetry[T any] func(T, error) bool
type Backoff func(retry int) time.Duration

func WithRetry[T any](base Dispatcher[T], shouldRetry ShouldRetry[T], maxRetries int, backoff Backoff) Dispatcher[T] {
	if shouldRetry == nil {
		shouldRetry = func(T, error) bool { return true } // Always retry by default
	}
	if backoff == nil {
		backoff = func(int) time.Duration { return 0 } // No backoff by default
	}
	if maxRetries < 0 {
		maxRetries = 0
	}

	return &retryDispatcher[T]{
		Dispatcher:  base,
		shouldRetry: shouldRetry,
		maxRetries:  maxRetries,
		backoff:     backoff,
	}
}

type retryDispatcher[T any] struct {
	Dispatcher[T]
	shouldRetry func(T, error) bool
	maxRetries  int
	backoff     func(retry int) time.Duration
}

func (d *retryDispatcher[T]) Dispatch(value T) error {
	var err error
	for i := 0; i <= d.maxRetries; i++ {
		err = d.Dispatcher.Dispatch(value)
		if err == nil || !d.shouldRetry(value, err) {
			break
		}
		time.Sleep(d.backoff(i))
	}
	return err
}

func WithAdaptation[T any](base Dispatcher[T], strategy Strategy[T], provider MetaProvider[T]) Dispatcher[T] {
	if provider == nil {
		provider = &defaultMetaProvider[T]{}
	}
	if strategy == nil {
		strategy = noOpStrategy[T]{}
	}
	return &adaptable[T]{
		Dispatcher: base,
		strategy:   strategy,
		provider:   provider,
	}
}

type adaptable[T any] struct {
	Dispatcher[T]
	strategy   Strategy[T]
	provider   MetaProvider[T]
	queueIndex int64
}

func (d *adaptable[T]) Dispatch(value T) error {
	// There is too much to say here...
	// This dispatcher performs better — at least in theory — under highly concurrent environments
	// due to the instrumentation overhead and control logic. also I'm not sure about the tradeoff:
	// in order to support feedback-based strategies, we need introspection.
	// In other words: there is a feedback controller requirement for the feedback controller.
	// it can be solved with asynchronous low priority pipes, but it requires statistics.
	// so I will keep it simple, maybe less than optimal, for now
	idx := atomic.AddInt64(&d.queueIndex, 1) - 1
	meta := d.provider.OnQueued(value, d.Dispatcher.QueueSize(), idx)
	if err := d.strategy.BeforeDispatch(meta); err != nil {
		d.provider.OnFinish(meta, err)
		d.strategy.Feedback(meta)
		return err
	}

	d.provider.OnDispatch(meta)
	err := d.Dispatcher.Dispatch(value)
	d.provider.OnFinish(meta, err)
	d.strategy.Feedback(meta)
	return err
}
