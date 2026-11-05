// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package obs

import (
	"github.com/andrerrcosta2/gtools/gflux/obs/internal/observables"
	"github.com/andrerrcosta2/gtools/gflux/obs/internal/observers"
)

type observableAdapter[T any] struct {
	inner observables.Observable[T]
}

func (a observableAdapter[T]) Sub(o *Observer[T]) (Subscription, error) {
	return a.inner.Sub(observerAdapter(o))
}

func observerAdapter[T any](o *Observer[T]) *observers.Observer[T] {
	if o == nil {
		return &observers.Observer[T]{}
	}

	return &observers.Observer[T]{
		Next:     o.Next,
		Error:    o.Error,
		Complete: o.Complete,
	}
}

type subjectAdapter[T any] struct {
	inner observables.Subject[T]
}

func (s subjectAdapter[T]) Sub(o *Observer[T]) (Subscription, error) {
	return s.inner.Sub(observerAdapter(o))
}

func (s subjectAdapter[T]) Next(v T)        { s.inner.Next(v) }
func (s subjectAdapter[T]) Error(err error) { s.inner.Error(err) }
func (s subjectAdapter[T]) Complete()       { s.inner.Complete() }
func (s subjectAdapter[T]) UnsubAll()       { s.inner.UnsubAll() }
