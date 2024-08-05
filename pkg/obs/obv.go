package obs

import "github.com/andrerrcosta2/gtools/pkg/functions"

type Obv[T any] struct {
	Nxt functions.Consumer[T]
	Err functions.Consumer[error]
	Cpt functions.Runnable
}

func NewObv[T any](nxt func(T), err func(error), cpt func()) *Obv[T] {
	return &Obv[T]{
		Nxt: nxt,
		Err: err,
		Cpt: cpt,
	}
}
