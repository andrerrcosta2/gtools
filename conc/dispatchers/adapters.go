// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package dispatchers

import (
	"time"

	"github.com/andrerrcosta2/gtools/conc/core/dispatcher"
)

type defaultMetaProvider[T any] struct{}

func (d *defaultMetaProvider[T]) OnQueued(value T, queueSize int, index int64) *dispatcher.Record[T] {
	return &dispatcher.Record[T]{
		Value:      value,
		StartedAt:  time.Now(),
		QueueDepth: queueSize,
		Index:      index,
		Attributes: make(map[string]any),
	}
}

func (d *defaultMetaProvider[T]) OnDispatch(record *dispatcher.Record[T]) {
	record.StartedAt = time.Now()
}

func (d *defaultMetaProvider[T]) OnFinish(record *dispatcher.Record[T], err error) {
	record.FinishedAt = time.Now()
	record.Error = err
}
