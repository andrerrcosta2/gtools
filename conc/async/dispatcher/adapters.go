// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package dispatcher

import "time"

type Meta[T any] struct {
	Value        T
	Error        error
	QueuedAt     time.Time      // Optional: when it was queued
	DispatchedAt time.Time      // Optional: when actual handler started
	FinishedAt   time.Time      // Optional: when handler returned
	Latency      time.Duration  // = FinishedAt - DispatchedAt
	TimeInQueue  time.Duration  // = DispatchedAt - QueuedAt
	Index        int64          // Optional: useful if order matters
	QueueSize    int            // Snapshot when dequeued
	Context      map[string]any // Optional
}

type MetaProvider[T any] interface {
	OnQueued(value T, queueSize int, index int64) *Meta[T]
	OnDispatch(meta *Meta[T])
	OnFinish(meta *Meta[T], err error)
}

type defaultMetaProvider[T any] struct{}

func (d *defaultMetaProvider[T]) OnQueued(value T, queueSize int, index int64) *Meta[T] {
	return &Meta[T]{
		Value:     value,
		QueuedAt:  time.Now(),
		QueueSize: queueSize,
		Index:     index,
		Context:   make(map[string]any),
	}
}

func (d *defaultMetaProvider[T]) OnDispatch(meta *Meta[T]) {
	meta.DispatchedAt = time.Now()
	meta.TimeInQueue = meta.DispatchedAt.Sub(meta.QueuedAt)
}

func (d *defaultMetaProvider[T]) OnFinish(meta *Meta[T], err error) {
	meta.FinishedAt = time.Now()
	meta.Latency = meta.FinishedAt.Sub(meta.DispatchedAt)
	meta.Error = err
}
