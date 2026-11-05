# conc/schedulers

`` provides a concrete scheduler implementation for:

- one-shot tasks (`Schedule`, `At`, `After`)
- recurring tasks (`Every`)
- resettable timers (`TimerFactory`, `chrono.Timer`)

It is designed to work with the `../../../core/domain/contracts/timing/scheduler` interfaces and `chrono.Clock`.

## What You Get

- `New()`: scheduler with system clock and internal timed source
- `Standalone(clock)`: scheduler with caller-selected clock
- `WithSource(src, clock)`: scheduler bound to a caller-provided timed source
- `TimerFactory(src, clock)`: timer factory that creates resettable timers

## Scheduler Semantics

- `Schedule(task)` schedules at the scheduler's current clock time.
- `At(at, task)` schedules a one-shot task at absolute time `at`.
- `After(d, task)` schedules a one-shot task after delay `d`.
- `Every(interval, task)` schedules recurring execution.
- Cancelling the returned `scheduler.Cancellable` prevents future execution.
- Non-positive `Every` intervals are normalized to `1ns`.
- `Close()` rejects new scheduling with `scheduler.ErrSchedulerClosed`.
- `AdvanceTo(now)` executes all due work up to `now`.

## Timer Semantics

- `NewTimer(cb)` returns a stopped timer.
- `Reset(d)` schedules callback execution at `clock.Now() + d`.
- `Stop()` cancels only the current schedule.
- `Stop()` is not terminal: calling `Reset` later re-arms the timer.
- Calling `Reset` while active cancels the previous schedule and replaces it.

## Basic Example

```go
package main

import (
	"context"
	"time"

	"github.com/andrerrcosta2/gtools/core/domain/contracts/timing/scheduler"
	"github.com/andrerrcosta2/gtools/conc/timing/schedulers"
)

func main() {
	s := schedulers.New()

	_, _ = s.After(250*time.Millisecond, scheduler.Task{
		Ctx: context.Background(),
		Job: func(ctx context.Context) error {
			// do work
			return nil
		},
	})

	_, _ = s.Every(1*time.Second, scheduler.Task{
		Ctx: context.Background(),
		Job: func(ctx context.Context) error {
			// recurring work
			return nil
		},
	})
}
```

## Timer Example

```go
package main

import (
	"time"

	"github.com/andrerrcosta2/gtools/core/domain/contracts/timing/scheduler"
)

func main() {
	var tf scheduler.TimerFactory // provided by your scheduler/engine wiring

	t := tf.NewTimer(func() {
		// fire once
	})

	t.Reset(100 * time.Millisecond)
	t.Stop()                        // cancel current schedule
	t.Reset(200 * time.Millisecond) // re-arm
}
```

## Notes

- `TimerFactory` is a low-level primitive that requires a timed source and clock. In most app code, obtain a `scheduler.TimerFactory` from your runtime/engine wiring rather than constructing it ad hoc.
- For deterministic testing, pair this package with virtual-clock-based test schedulers.
