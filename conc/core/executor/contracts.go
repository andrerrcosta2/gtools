package executor

import "errors"

var ErrExecutorClosed = errors.New("executor is closed")

// QueueExecutor accepts items for asynchronous processing.
type QueueExecutor[T any] interface {
	Submit(item T) error
	Close() error
	IsClosed() bool
	Wait()
}
