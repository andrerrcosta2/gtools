// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package routes

import (
	"github.com/andrerrcosta2/gtools/conc/core/mailbox"
	"github.com/andrerrcosta2/gtools/core/data/event"
)

const (
	Default  mailbox.Route = 0
	Buffered mailbox.Route = iota
	Conflated
	Cancellable
)

func IsDefault[T event.Routed[mailbox.Route]](event T) bool {
	return event.Route() == Default
}

func IsBuffered[T event.Routed[mailbox.Route]](event T) bool {
	return event.Route() == Buffered
}

func IsConflated[T event.Routed[mailbox.Route]](event T) bool {
	return event.Route() == Conflated
}

func IsCancellable[T event.Routed[mailbox.Route]](event T) bool {
	return event.Route() == Cancellable
}
