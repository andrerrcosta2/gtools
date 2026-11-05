// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package mailbox

import (
	"github.com/andrerrcosta2/gtools/core/domain/function"
)

type Route uint16

type Receiver[E any] interface {
	// Receive listen for events and returns true if any event exists
	Receive() (E, bool)
}

type Sender[E any] interface {
	// Send sends an event to the mailbox a returns true if it has succeeded
	Send(E) bool
}

type Open[E any] interface {
	Receiver[E]
	Sender[E]
}

type Buffered[E any] interface {
	Open[E]
	Cap() int
	Len() int
}

type Drainable[E any] interface {
	Open[E]
	// Drain consumes all currently available messages without blocking.
	Drain(fn function.Predicate[E]) int
	// DrainBlocking waits until work is available, then drains all available messages.
	// Returns false when the mailbox has been closed and no more
	// work can be produced.
	DrainBlocking(fn function.Predicate[E]) bool
}
