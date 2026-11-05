# runtimes

Runtimes is a **time simulation and scheduling tools** for Go, designed to provide:

- Virtualized schedulers and timers for deterministic testing
- Support for delayed, recurring, and cancellable tasks
- Real-time or virtual time execution via a unified runtime

---

## Core Concepts

### Task & Scheduler Hierarchy

The following diagram illustrates the relationships between the core components:

```
                     ┌───────────────────────────┐
                     │      scheduler.Task       │
                     │───────────────────────────│
                     │  (job + context wrapper)  │
                     └─────────────┬─────────────┘
                                   │
                   ┌───────────────┴───────────────┐
                   │           Scheduler           │
                   │───────────────────────────────│
                   │ Schedule(task)                │
                   │ Close()                       │
                   │ IsClosed()                    │
                   └─────────-─┬────────┬──────────┘
                               │        │
                  ┌────────────┘        └────────────┐
                  │                                  │
        ┌─────────▼─────────┐              ┌─────────▼─────────┐
        │    Delayed        │              │      Recurring    │
        │───────────────────│              │───────────────────│
        │ At(time)          │              │ Every(interval)   │
        │ After(duration)   │              │                   │
        └─────────┬─────────┘              └─────────┬─────────┘
                  │                                  │
                  │                                  │
                  └─────────────────▲────────────────┘
                                    │
                         ┌──────────┴──────────┐
                         │     timeScheduler   │
                         │─────────────────────│
                         │ Implements:         │
                         │ Scheduler           │
                         │ Delayed             │
                         │ Recurring           │
                         │ TimerFactory        │
                         └──────────┬──────────┘
                                    │
             ┌──────────────────────┼──────────────────────┐
             │                      │                      │
   ┌─────────▼─────────┐  ┌─────────▼─────────┐  ┌─────────▼─────────┐
   │ Scheduler()       │  │ Delayed()         │  │ Recurring()       │
   │───────────────────│  │───────────────────│  │───────────────────│
   │ return Scheduler  │  │ return Delayed    │  │ return Recurring  │
   └───────────────────┘  └───────────────────┘  └───────────────────┘
             │
   ┌─────────▼──────────┐
   │ TimerFactory()     │
   │────────────────────│
   │ return TimerFactory│
   └────────────────────┘
```

---

## Runtime

Runtimes exposes a `Runtime` that wraps schedulers and timers:

```go
type Runtime interface {
    Now() chrono.Time
    AdvanceTo(t chrono.Time)
    Step() bool

    Scheduler() scheduler.Scheduler
    TimerFactory() scheduler.TimerFactory
}
```

- `Now()` – returns the current time
- `AdvanceTo(t)` – advances virtual time to a given point
- `Step()` – executes the next scheduled step
- `Scheduler()` – returns the scheduler interface
- `TimerFactory()` – returns a timer factory

---

## Interface Constructors

Runtimes provides constructors for each scheduler interface:

```go
sched := testing.Scheduler()      // Full scheduler
delayed := testing.Delayed()      // Delayed-only
recurring := testing.Recurring()  // Recurring-only
timerFactory := testing.TimerFactory() // Timer factory
```

`timeScheduler` implements all interfaces:

- `Scheduler`
- `Delayed`
- `Recurring`
- `TimerFactory`

---

## Example: Virtual Scheduler

```go
sched := testingconc.Scheduler()

executed := false
task := scheduler.NewTask(func(ctx context.Context) error {
    executed = true
    fmt.Println("Task executed")
    return nil
}, context.Background())

// Schedule task after 1s virtual time
_, _ = sched.After(time.Second, *task)

// Advance virtual time
sched.AdvanceTo(time.Unix(0, int64(time.Second)))

// executed == true
```

---

## Example: Recurring Task

```go
rec := testingconc.Recurring()
count := 0

task := scheduler.NewTask(func(ctx context.Context) error {
    count++
    return nil
}, context.Background())

timer, _ := rec.Every(500 * time.Millisecond)

// Advance virtual time multiple broadcasters
sched := testingconc.Scheduler()
sched.AdvanceTo(time.Unix(0, int64(time.Millisecond*1500)))

// count == 3
```

---

## Notes

- Recurring tasks return a `Cancellable` to stop intervals manually.
- `Scheduler.Close()` cancels all pending tasks.
- Interface-specific constructors allow exposing **only required functionality**, improving decoupling.

---

## Repository Structure

```
/runtimes
 ├─ /core
 │   └─ /scheduler         # Core scheduler interfaces
 ├─ /conc
 │   └─ /schedulers        # Real-time implementations
 └─ /testing           # Virtual scheduler & timers for tests
```

---

## Philosophy

Runtimes aims to **fully control time** for deterministic testing while providing a **drop-in replacement** 
for production schedulers. It ensures consistent behavior for delayed, recurring, and cancellable tasks, 
enabling robust testing of asynchronous logic.

