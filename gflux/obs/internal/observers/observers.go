// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package observers

import (
	"github.com/andrerrcosta2/gtools/core/domain/function"
	"github.com/andrerrcosta2/gtools/core/domain/function/consumers"
	"github.com/andrerrcosta2/gtools/core/domain/function/runnables"
	"github.com/andrerrcosta2/gtools/gflux/core/domain/lifecycle/gate"
)

type Observer[T any] struct {
	Next     function.Consumer[T]
	Error    function.Consumer[error]
	Complete function.Runnable
}

func New[T any](nxt func(T), err func(error), cpt func()) *Observer[T] {
	return &Observer[T]{
		Next:     nxt,
		Error:    err,
		Complete: cpt,
	}
}

func NewComplete[T any](cpt func()) *Observer[T] {
	return &Observer[T]{
		Complete: cpt,
	}
}

func NewError[T any](err func(error)) *Observer[T] {
	return &Observer[T]{
		Error: err,
	}
}

func NewNext[T any](nxt func(T)) *Observer[T] {
	return &Observer[T]{
		Next: nxt,
	}
}

func Noop[T any]() *Observer[T] {
	return &Observer[T]{
		Next:     consumers.Void[T],
		Error:    consumers.Void[error],
		Complete: runnables.Void,
	}
}

func Create[T any](o *Observer[T], gate gate.Closeable) *Observer[T] {
	origNext := o.Next
	origErr := o.Error
	origCpt := o.Complete

	if origNext == nil {
		o.Next = consumers.Void[T]
	} else {
		o.Next = func(v T) {
			if gate.Closed() {
				return
			}
			origNext(v)
		}
	}

	if origErr == nil {
		o.Error = consumers.Void[error]
	} else {
		o.Error = func(err error) {
			if !gate.Closed() {
				gate.Close()
				origErr(err)
			}
		}
	}

	if origCpt == nil {
		o.Complete = runnables.Void
	} else {
		o.Complete = func() {
			if !gate.Closed() {
				gate.Close()
				origCpt()
			}
		}
	}

	return o
}

func Normalize[T any](o *Observer[T]) *Observer[T] {
	if o.Next == nil {
		o.Next = consumers.Void[T]
	}
	if o.Error == nil {
		o.Error = consumers.Void[error]
	}
	if o.Complete == nil {
		o.Complete = runnables.Void
	}
	return o
}

func NormalizeNext[T any](o *Observer[T]) *Observer[T] {
	if o.Next == nil {
		o.Next = consumers.Void[T]
	}
	return o
}

func NormalizeError[T any](o *Observer[T]) *Observer[T] {
	if o.Error == nil {
		o.Error = func(err error) {}
	}
	return o
}

func NormalizeComplete[T any](o *Observer[T]) *Observer[T] {
	if o.Complete == nil {
		o.Complete = runnables.Void
	}
	return o
}
