// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package chronos

import (
	"github.com/andrerrcosta2/gtools/core/domain/contracts/chrono"
)

type timer struct {
	clock  chrono.Clock
	ticker chrono.Ticker
	cb     func()
}

// Stop cancels the currently scheduled timer callback, if any.
//
// Stop does not permanently disable the timer; a later Reset re-arms it.
func (t *timer) Stop() {
	t.ticker.Stop()
}

// Reset schedules the timer callback to run after d from clock.AtNow().
//
// If there is an active scheduled callback, it is cancelled first. Reset always
// re-arms the timer with a new due time.
func (t *timer) Reset(d chrono.Duration) {
	t.ticker = t.clock.NewTicker(d)
	for range t.ticker.C() {
		t.cb()
	}
}

// TimerFactory builds timers backed by the provided timed source and clock.
func TimerFactory(clock chrono.Clock) chrono.TimerFactory {
	return &timerFactory{
		clock: clock,
	}
}

type timerFactory struct {
	clock chrono.Clock
}

// NewTimer creates a stopped timer bound to the factory source/clock.
//
// The callback executes only after Reset is called and time is advanced to due time.
func (f *timerFactory) NewTimer(cb func()) chrono.Timer {
	return &timer{
		clock: f.clock,
		cb:    cb,
	}
}
