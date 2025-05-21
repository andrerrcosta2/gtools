package obs

import "github.com/andrerrcosta2/gtools/core/domain/functions"

type Observer[T any] struct {
	Next     functions.Consumer[T]
	Error    functions.Consumer[error]
	Complete functions.Runnable
}

func NewObserver[T any](nxt func(T), err func(error), cpt func()) *Observer[T] {
	return &Observer[T]{
		Next:     nxt,
		Error:    err,
		Complete: cpt,
	}
}
