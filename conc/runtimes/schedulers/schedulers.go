package schedulers

import (
	"context"
	"errors"
	"sync/atomic"

	"github.com/andrerrcosta2/gtools/conc/actors/mailboxes"
	"github.com/andrerrcosta2/gtools/conc/core/actor/mailbox"
	"github.com/andrerrcosta2/gtools/conc/core/syncs/latches"
	"github.com/andrerrcosta2/gtools/conc/runtimes/schedulers/internal/cancellation"
	"github.com/andrerrcosta2/gtools/conc/runtimes/schedulers/internal/commands"
	"github.com/andrerrcosta2/gtools/conc/runtimes/schedulers/schedulables"
	"github.com/andrerrcosta2/gtools/conc/timing/schedulers/agendas"
	"github.com/andrerrcosta2/gtools/core/chronos"
	"github.com/andrerrcosta2/gtools/core/domain/contracts/timing/agenda"
	"github.com/andrerrcosta2/gtools/core/domain/contracts/timing/chrono"
	"github.com/andrerrcosta2/gtools/core/domain/contracts/timing/chrono/durations"
	"github.com/andrerrcosta2/gtools/core/domain/contracts/timing/scheduler"
)

var (
	ErrNilJob              = errors.New("nil job")
	ErrClosedScheduler     = errors.New("the scheduler is closed")
	ErrNotAdvanceableClock = errors.New("the clock is not advanceable")
	ErrClockMovedBackwards = errors.New("the clock moved backwards")
)

// NewTask is a convenience constructor for Task.
func NewTask(job Job, ctx context.Context) Task {
	return Task{
		Job: job,
		Ctx: ctx,
	}
}

// Task is the executable unit used by scheduler APIs.
//
// Ctx is optional. Implementations commonly treat nil as context.Background().
type Task struct {
	Ctx context.Context
	Job Job
}

func (t Task) Run() {
	t.Job(t.Ctx)
}

// Job is the minimal executable unit for scheduler and dispatcher components.
type Job func(ctx context.Context)

// New constructs a scheduler backed by the system clock and an infra timed source.
//
// Use this when you want a ready-to-use runtime scheduler without wiring dependencies.
func New(tick chrono.Duration) *Scheduler {
	return newScheduler(chronos.SystemClock(), tick)
}

// WithClock constructs a scheduler with the provided clock;
func WithClock(clock chrono.Clock, tick chrono.Duration) *Scheduler {
	if _, ok := clock.(chrono.AdvanceableClock); ok {
		return newAdvanceable(clock, tick)
	}
	return newScheduler(clock, tick)
}

func newAdvanceable(clock chrono.Clock, tick chrono.Duration) *Scheduler {

	s := &Scheduler{
		mb:     mailboxes.Disruptor[commands.SchedulerEvent](1024),
		clock:  clock,
		agenda: agendas.OfTasks(clock.Now(), durations.Ms(100), 1024),
		ticker: clock.NewTicker(tick),
	}
	s.advance = func(to chrono.Time) error {
		if to.Less(clock.Now()) {
			return ErrClockMovedBackwards
		}
		latch := latches.Unchecked()
		if !s.mb.Send(commands.AdvanceTime(to, latch)) {
			return ErrClosedScheduler
		}
		latch.Wait()
		return nil
	}
	s.startEventListener()
	return s
}

func newScheduler(clock chrono.Clock, tick chrono.Duration) *Scheduler {
	s := &Scheduler{
		mb:      mailboxes.Disruptor[commands.SchedulerEvent](1024),
		clock:   clock,
		agenda:  agendas.OfTasks(clock.Now(), durations.Ms(100), 1024),
		ticker:  clock.NewTicker(tick),
		advance: func(to chrono.Time) error { return ErrNotAdvanceableClock },
	}

	s.startEventListener()
	s.startTicker()
	return s
}

func ScheduleNow(s scheduler.NowProvider[Task], task Task) error {
	_, err := s.At(s.Now(), task)
	return err
}

// Scheduler is a Scheduler task scheduler with delayed and recurring scheduling.
//
// Tasks are pushed to a timed source and executed when AdvanceTo reaches or passes
// their scheduled engine. Cancellation is context-based and idempotent.
type Scheduler struct {
	mb      mailbox.Drainable[commands.SchedulerEvent]
	clock   chrono.Clock
	agenda  agenda.Chrono[schedulables.TaskWrapper]
	ticker  chrono.Ticker
	advance func(chrono.Time) error
	closed  atomic.Bool
}

// AdvanceTo advances the underlying stepper clock and executes all due task.
func (s *Scheduler) AdvanceTo(to chrono.Time) error {
	return s.advance(to)
}

// After schedules a one-shot task to run after d from the current clock engine.
func (s *Scheduler) After(d chrono.Duration, task Task) (scheduler.Cancellable, error) {
	if task.Job == nil {
		return cancellation.Noop(), ErrNilJob
	}
	return s.at(s.clock.Now().Add(d), task)
}

// At schedules a one-shot task to run eventAt the provided absolute engine.
//
// Returns ErrSchedulerClosed when the scheduler is closed.
// If task.Job is nil, At is a no-op and returns (nil, nil).
// The returned cancellable prevents execution when cancelled before run.
func (s *Scheduler) At(at chrono.Time, task Task) (scheduler.Cancellable, error) {
	if task.Job == nil {
		return cancellation.Noop(), ErrNilJob
	}
	return s.at(at, task)
}

func (s *Scheduler) at(at chrono.Time, task Task) (scheduler.Cancellable, error) {
	baseCtx := task.Ctx
	if baseCtx == nil {
		baseCtx = context.Background()
	}

	ctx, cancel := context.WithCancel(baseCtx)
	task.Ctx = ctx
	ev := commands.ScheduleAt(at, task)

	ok := s.mb.Send(ev)
	if !ok {
		cancel() // this is safe (context cancels downward only)
		return cancellation.Noop(), scheduler.ErrClosedScheduler
	}

	return cancellation.Func(cancel, ev.Token()), nil
}

// Every schedules a recurring task with the given interval.
//
// Non-positive intervals are normalized to 1ns.
// The returned cancellable stops all future runs.
func (s *Scheduler) Every(interval chrono.Duration, task Task) (scheduler.Cancellable, error) {
	if task.Job == nil {
		return cancellation.Noop(), ErrNilJob
	}
	return s.everyFrom(s.clock.Now().Add(interval), interval, task)
}

// EveryFrom schedules a recurring task starting eventAt a specific first time (firstTick),
// then repeats eventEvery interval. Non-positive intervals are normalized to 1ns.
func (s *Scheduler) EveryFrom(firstTick chrono.Time, interval chrono.Duration,
	task Task) (scheduler.Cancellable, error) {
	if task.Job == nil {
		return cancellation.Noop(), ErrNilJob
	}
	return s.everyFrom(firstTick, interval, task)

}

func (s *Scheduler) everyFrom(firstTick chrono.Time, interval chrono.Duration,
	task Task) (scheduler.Cancellable, error) {
	if interval <= 0 {
		interval = chrono.Nanosecond
	}

	baseCtx := task.Ctx
	if baseCtx == nil {
		baseCtx = context.Background()
	}

	ctx, cancel := context.WithCancel(baseCtx)
	start := firstTick.NextAfterOrNow(s.clock.Now(), interval)
	task.Ctx = ctx

	ev := commands.ScheduleAtEvery(
		start,
		interval,
		task,
	)

	ok := s.mb.Send(ev)
	if !ok {
		cancel() // this is safe (context cancels downward only)
		return &cancellation.NoopHandler{}, scheduler.ErrClosedScheduler
	}
	return cancellation.Func(cancel, ev.Token()), nil
}

func (s *Scheduler) startEventListener() {
	go func() {
		for s.mb.DrainBlocking(s.handleEvent) {
		}
	}()
}

func (s *Scheduler) startTicker() {
	go func() {
		for tick := range s.ticker.C() {
			s.mb.Send(commands.DueTasks(tick))
		}
	}()
}

// Close marks the scheduler as closed.
//
// New scheduling requests will fail with ErrSchedulerClosed. Already queued
// task are actively removed by this method.
func (s *Scheduler) Close() error {
	if s.closed.CompareAndSwap(false, true) {
		return nil
	}
	return ErrClosedScheduler
}

func (s *Scheduler) Now() chrono.Time {
	return s.clock.Now()
}

func (s *Scheduler) handleEvent(event commands.SchedulerEvent) bool {
	switch e := event.(type) {
	case commands.ScheduleAtEvent:
		s.eventAt(e)
	case commands.ScheduleAtEveryEvent:
		s.eventEvery(e)
	case commands.AdvanceTimeEvent:
		s.eventAdvance(e)
	case commands.DueTasksEvent:
		s.eventDue(e)

	default:
		panic("schedulers: shouldn't never reach here: unknown command")
	}
	return true
}

func (s *Scheduler) eventAdvance(e commands.AdvanceTimeEvent) {
	s.clock.(chrono.AdvanceableClock).AdvanceTo(e.At())
	runDue(s.agenda, e.At())
	e.Latch().Signal()
}

func (s *Scheduler) eventAt(e commands.ScheduleAtEvent) {
	s.agenda.Push(schedulables.WrapTask(e.Task(), e.At()))
}

// eventEvery is the infra recurring scheduler implementation.
func (s *Scheduler) eventEvery(e commands.ScheduleAtEveryEvent) {
	s.agenda.Push(schedulables.WrapTask(e.Task(), e.At(), func(time chrono.Time) (chrono.Time, bool) {
		return time.Add(e.Interval()), true
	}))
}

func (s *Scheduler) eventDue(e commands.DueTasksEvent) {
	runDue(s.agenda, e.At())
}

func runDue(agenda agenda.Chrono[schedulables.TaskWrapper], now chrono.Time) {
	agenda.RangeDue(now, func(t schedulables.TaskWrapper) bool {
		t.Run()
		if next, ok := t.Next(now); ok {
			agenda.Push(schedulables.WrapTask(t.Unwrap(), next, t.Next))
		}
		return true
	})
}
