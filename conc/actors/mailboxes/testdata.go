// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package mailboxes

import (
	"github.com/andrerrcosta2/gtools/conc/actors/mailboxes/routes"
	"github.com/andrerrcosta2/gtools/conc/core/mailbox"
)

func newRoutedEvent(id, value int, r mailbox.Route) routedEvent {
	return routedEvent{
		id:    id,
		value: value,
		route: r,
	}
}

type routedEvent struct {
	id    int
	route mailbox.Route
	value int
}

func (e routedEvent) Route() mailbox.Route {
	return e.route
}

func (e routedEvent) Value() int {
	return e.value
}

func (e routedEvent) ID() int {
	return e.id
}

func defRoutedEv(id, value int) routedEvent {
	return routedEvent{id: id, value: value}
}

func cancelRoutedEv(id, value int) routedEvent {
	return routedEvent{id: id, value: value, route: routes.Cancellable}
}

type mockBuffer[T any] struct {
	value T
	full  bool
}

func (b *mockBuffer[T]) Write(v T) bool {
	if b.full {
		return false
	}
	b.value = v
	b.full = true
	return true
}

func (b *mockBuffer[T]) Read() (T, bool) {
	if !b.full {
		var zero T
		return zero, false
	}
	v := b.value
	b.full = false
	return v, true
}

func (b *mockBuffer[T]) Len() int {
	if b.full {
		return 1
	}
	return 0
}

func (b *mockBuffer[T]) Cap() int {
	return 1
}
