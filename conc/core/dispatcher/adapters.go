// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package dispatcher

import "time"

type Status uint8

const (
	Submitted Status = iota
	Accepted
	Started
	Finished
	Failed
	Rejected
	Dropped
	Cancelled
)

type Record[T any] struct {
	Value       T
	Index       int64
	QueueDepth  int
	Status      Status
	Error       error
	SubmittedAt time.Time
	AcceptedAt  time.Time
	StartedAt   time.Time
	FinishedAt  time.Time
	Attributes  map[string]any
}

type MetaProvider[T any] interface {
	OnQueued(value T, queueSize int, index int64) *Record[T]
	OnDispatch(meta *Record[T])
	OnFinish(meta *Record[T], err error)
}
