// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package commands

import (
	"github.com/andrerrcosta2/gtools/conc/core/syncs"
	"github.com/andrerrcosta2/gtools/conc/runtimes/schedulers/internal/cancellation"
	"github.com/andrerrcosta2/gtools/core/domain/contracts/timing/chrono"
	"github.com/andrerrcosta2/gtools/gflux/core/domain/task"
)

type SchedulerEvent interface {
	At() chrono.Time
}

func AdvanceTime(to chrono.Time, latch syncs.Latch) AdvanceTimeEvent {
	return AdvanceTimeEvent{
		at:    to,
		latch: latch,
	}
}

type AdvanceTimeEvent struct {
	at    chrono.Time
	latch syncs.Latch
}

func (e AdvanceTimeEvent) At() chrono.Time {
	return e.at
}

func (e AdvanceTimeEvent) Latch() syncs.Latch {
	return e.latch
}

func ScheduleAt(at chrono.Time, task task.Runnable) ScheduleAtEvent {
	return ScheduleAtEvent{
		at:    at,
		task:  task,
		token: cancellation.NewToken(),
	}
}

type ScheduleAtEvent struct {
	at    chrono.Time
	task  task.Runnable
	token *cancellation.Token
}

func (s ScheduleAtEvent) At() chrono.Time {
	return s.at
}

func (s ScheduleAtEvent) Task() task.Runnable {
	return s.task
}

func (s ScheduleAtEvent) Token() *cancellation.Token {
	return s.token
}

func ScheduleAtEvery(at chrono.Time, interval chrono.Duration, task task.Runnable) ScheduleAtEveryEvent {
	return ScheduleAtEveryEvent{
		at:       at,
		interval: interval,
		task:     task,
		token:    cancellation.NewToken(),
	}
}

type ScheduleAtEveryEvent struct {
	at       chrono.Time
	interval chrono.Duration
	task     task.Runnable
	token    *cancellation.Token
}

func (e ScheduleAtEveryEvent) At() chrono.Time {
	return e.at
}

func (e ScheduleAtEveryEvent) Interval() chrono.Duration {
	return e.interval
}

func (e ScheduleAtEveryEvent) Task() task.Runnable {
	return e.task
}

func (e ScheduleAtEveryEvent) Token() *cancellation.Token {
	return e.token
}

func DueTasks(at chrono.Time) DueTasksEvent {
	return DueTasksEvent{
		at: at,
	}
}

type DueTasksEvent struct {
	at chrono.Time
}

func (e DueTasksEvent) At() chrono.Time {
	return e.at
}
