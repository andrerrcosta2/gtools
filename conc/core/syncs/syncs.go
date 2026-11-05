// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package syncs

import "github.com/andrerrcosta2/gtools/conc/core/channel"

type Awaitable interface {
	Ready() bool
	Wait()
}

type Broadcaster interface {
	Broadcast()
}

type Blocker interface {
	Done() channel.Signal
}

type Completer interface {
	Done()
}

type WorkCounter interface {
	Add(int)
}

type Signaler interface {
	Signal()
}

type Waiter interface {
	Wait()
}

type CountCompleter interface {
	Completer
	WorkCounter
}

type CountWaiter interface {
	Waiter
	WorkCounter
}

type WaitGroup interface {
	WorkCounter
	Waiter
	Completer
}

type Flag interface {
	Signaler
	Blocker
}

type Latch interface {
	Waiter
	Signaler
}

type PairLatch interface {
	Waiter
	Signaler
	Waiter() Waiter
	Signaler() Signaler
}

type BroadcastLatch interface {
	Latch
	Broadcaster
}
