package scheduler

import (
	"errors"

	"github.com/andrerrcosta2/gtools/core/domain/contracts/timing/chrono"
)

// ErrClosedScheduler is returned when scheduling is attempted on a closed scheduler.
var ErrClosedScheduler = errors.New("scheduler is closed")

// Scheduler is the base contract for scheduling systems.
//
// Implementations may also satisfy Delayed and Recurring.
type Scheduler[T any] interface {
	// At schedules task for execution at absolute time at.
	At(at chrono.Time, task T) (Cancellable, error)
	// Close transitions the scheduler to a closed states.
	// Closed schedulers are expected to reject new scheduling operations.
	Close() error
}

type NowProvider[T any] interface {
	Scheduler[T]
	Now() chrono.Time
}

// Delayed schedules one-shot task at absolute or relative times.
type Delayed[T any] interface {
	// After schedules task for execution after delay d.
	After(d chrono.Duration, task T) (Cancellable, error)
}

// Recurring schedules repeatedly executing task.
type Recurring[T any] interface {
	// Every schedules task to run at each interval tick.
	Every(interval chrono.Duration, task T) (Cancellable, error)
	// EveryFrom schedules task starting at a specific time
	EveryFrom(start chrono.Time, interval chrono.Duration, task T) (Cancellable, error)
}

// Advanceable can drive virtual or externally controlled time progression.
type Advanceable interface {
	// AdvanceTo advances execution to now, running due work.
	AdvanceTo(now chrono.Time) error
}

// Cancellable is a handle to cancel scheduled execution.
type Cancellable interface {
	// Cancel prevents future execution for the associated schedule.
	Cancel()
}

// TimerFactory creates resettable timers integrated with scheduler time semantics.
type TimerFactory interface {
	// NewTimer returns a timer bound to the factory/runtime.
	NewTimer(cb func()) chrono.Timer
}
