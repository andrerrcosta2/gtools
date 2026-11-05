# conc/core/scheduler

`` defines the core scheduling contracts used across concurrency modules.

It is intentionally small and implementation-agnostic: this package provides interfaces and shared types, not runtime behavior.

## Core Types

- `Job`: executable function shape `func(context.Context) error`
- `Task`: a `Job` plus optional `context.Context`
- `ErrSchedulerClosed`: canonical error for scheduling attempts after closure

## Interfaces

- `Scheduler`
  - `Schedule(task Task) error`
  - `Close() error`
  - `IsClosed() bool`
- `Delayed`
  - `At(at time.Time, task Task) (Cancellable, error)`
  - `After(d time.Duration, task Task) (Cancellable, error)`
- `Recurring`
  - `Every(interval time.Duration, task Task) (Cancellable, error)`
- `Advanceable`
  - `AdvanceTo(now chrono.Time)`
- `Cancellable`
  - `Cancel()`
- `TimerFactory`
  - `NewTimer(cb func()) chrono.Timer`

## Design Intent

- Keep scheduling APIs composable:
  - implementations can expose only what they support (`Scheduler` alone, or `Scheduler + Delayed + Recurring`, etc.)
- Support both real-time and virtual-time engines:
  - `Advanceable` exists for deterministic/manual clock progression
- Keep cancellation generic:
  - `Cancellable` avoids leaking implementation details

## Implementation Guidance

- Treat `Task.Ctx == nil` as `context.Background()` unless there is a strong reason not to.
- Return `ErrSchedulerClosed` for new scheduling requests after `Close()`.
- Make `Cancel()` idempotent.
- Document whether `Close()` drains pending tasks, drops them, or just blocks new scheduling.

## Minimal Usage

```go
package example

import (
	"context"
	"time"

	"github.com/andrerrcosta2/gtools/conc/core/scheduler"
)

func enqueueNow(s scheduler.Scheduler) error {
	return s.Schedule(scheduler.Task{
		Ctx: context.Background(),
		Job: func(ctx context.Context) error {
			// work
			return nil
		},
	})
}

func enqueueLater(d scheduler.Delayed) (scheduler.Cancellable, error) {
	return d.After(250*time.Millisecond, scheduler.Task{
		Job: func(ctx context.Context) error {
			return nil
		},
	})
}
```
