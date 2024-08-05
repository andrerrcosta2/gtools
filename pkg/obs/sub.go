package obs

import (
	"github.com/andrerrcosta2/gtools/pkg/functions"
	"github.com/google/uuid"
)

type Sub[T any] struct {
	id     string
	obs    Obs[T]
	closed bool
	Close  functions.Supplier[bool]
}

func NewSub[T any](obs Obs[T]) *Sub[T] {
	s := &Sub[T]{
		id:     uuid.New().String(),
		obs:    obs,
		closed: false,
	}

	s.Close = functions.Once(func() bool {
		s.closed = true
		return true
	})

	obs.Rmo(s)
	return s
}

func (s *Sub[T]) Gid() string {
	return s.id
}

func (s *Sub[T]) Uns() {
	if !s.closed {
		s.obs.Rmo(s)
	}
}

func (s *Sub[T]) Closed() bool {
	return s.closed
}
