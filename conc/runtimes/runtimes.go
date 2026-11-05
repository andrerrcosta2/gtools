// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package runtimes

import (
	"time"

	"github.com/andrerrcosta2/gtools/conc/core/channel/channels"
	"github.com/andrerrcosta2/gtools/conc/core/dispatcher"
	"github.com/andrerrcosta2/gtools/conc/runtimes/schedulers"
	"github.com/andrerrcosta2/gtools/core/chronos"
	"github.com/andrerrcosta2/gtools/core/domain/contracts/timing/chrono"
	"github.com/andrerrcosta2/gtools/core/domain/contracts/timing/scheduler"
)

type Scheduler[T any] interface {
	scheduler.Scheduler[T]
	scheduler.Advanceable
	scheduler.Delayed[T]
	scheduler.Recurring[T]
}

type Runtime[C chrono.Clock] interface {
	Clock() C

	After(d chrono.Duration) <-chan chrono.Time
	NewTicker(d chrono.Duration) chrono.Ticker

	Scheduler() Scheduler[C]

	TimerFactory() scheduler.TimerFactory
}

type Advanceable[C chrono.Clock] interface {
	Runtime[C]
	// Advance advances the clock "d" duration from now
	Advance(d chrono.Duration)
	// AdvanceTo advances the clock to the exact time.
	// if the time is before the current clock it doesn't perform any operation
	AdvanceTo()
}

func New[C chrono.Clock](clock C) Runtime[C] {
	rt := &runtime[C]{
		clock:        clock,
		scheduler:    schedulers.WithClock(clock),
		dispatcher:   nil,
		timerFactory: chronos.TimerFactory(clock),
	}

	return rt
}

type runtime[C chrono.Clock] struct {
	clock      C
	scheduler  Scheduler
	dispatcher dispatcher.Executor[schedulers.Task]
	timerFactory scheduler.TimerFactory
}

func (r *runtime[C]) Clock() C {
	return r.clock
}



func (r *runtime[C]) Scheduler() Scheduler {
	return r.scheduler
}

func (r *runtime[C]) TimerFactory() scheduler.TimerFactory {
	return r.timerFactory
}

func Start[C chrono.Clock](rt *runtime[C], tick time.Duration) (stop func()) {
	flag := channels.Flag()

	go func() {
		ticker := time.NewTicker(tick)
		defer ticker.Stop()

		for {
			select {
			case <-flag.Done():
				return
			case now := <-ticker.C:
				rt.AdvanceTo(chrono.FromStdTime(now))
			}
		}
	}()

	return func() { flag.Signal() }
}

func System() *SystemRuntime {
	clock := chronos.SystemClock()
	rt := New(clock, clock.Now())
	stop := Start(rt, time.Millisecond)
	return &SystemRuntime{rt, stop}
}

type SystemRuntime struct {
	Runtime
	stop func()
}

func (e *SystemRuntime) Close() error {
	e.stop()
	return e.Runtime.Scheduler().Close()
}

type VirtualRuntime struct {
	Runtime *runtime // embed the runtime for convenience
	Clock   chrono.AdvanceableClock
}

func Virtual(clock chrono.AdvanceableClock) *VirtualRuntime {
	rt := &runtime{
		clock:     clock,
		scheduler: schedulers.New(clock)),
		dispatcher:   nil,
		timerFactory: nil,
	}

	rt.scheduler = schedulers.WithSource(src, clock)
	rt.timerFactory = chronos.TimerFactory(src, clock)

	return &VirtualRuntime{
		Runtime: rt,
		Clock:   clock,
	}
}

// Advance the virtual clock and run all broadcasters up to the new time
func (e *VirtualRuntime) Advance(d chrono.Duration) {
	e.Clock.Advance(d)
	e.Runtime.AdvanceTo(e.Clock.Now())
}

// convenience methods
func (e *VirtualRuntime) Now() chrono.Time                     { return e.Runtime.Now() }
func (e *VirtualRuntime) Step() (bool, error)                  { return e.Runtime.Step() }
func (e *VirtualRuntime) Scheduler() Scheduler                 { return e.Runtime.Scheduler() }
func (e *VirtualRuntime) TimerFactory() scheduler.TimerFactory { return e.Runtime.TimerFactory() }

func (e *VirtualRuntime) Close() error {
	return e.Runtime.Scheduler().Close()
}
