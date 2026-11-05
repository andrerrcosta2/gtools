// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package engines

import (
	"errors"

	"github.com/andrerrcosta2/gtools/conc/core/syncs"
	"github.com/andrerrcosta2/gtools/conc/core/syncs/latches"
	"github.com/andrerrcosta2/gtools/conc/core/syncs/waiters"
	"github.com/andrerrcosta2/gtools/conc/timing/schedulers/schedulables"
	"github.com/andrerrcosta2/gtools/core/data/closeable"
	"github.com/andrerrcosta2/gtools/core/domain/contracts/timing/agenda"
	"github.com/andrerrcosta2/gtools/core/domain/contracts/timing/chrono"
	"github.com/andrerrcosta2/gtools/core/domain/function"
	"github.com/andrerrcosta2/gtools/core/domain/function/functions"
	"github.com/andrerrcosta2/gtools/gflux/core/domain/task"
)

var (
	ErrSealed      = errors.New("the engine is sealed")
	ErrOpen        = errors.New("the engine is open")
	ErrClosed      = errors.New("the engine is closed")
	ErrClosedTwice = errors.New("the engine is already closed")
)

type TaskSchedulerRunner[T task.Schedulable[chrono.Time]] interface {
	Schedule(T) error
	RunDue(time chrono.Time) error
	closeable.Terminal
}

// InMemory is a non-concurrent scheduler engine that uses an in memory agenda as task storage
func InMemory(queue agenda.Chrono[schedulables.TaskWrapper], ticker chrono.Ticker) TaskSchedulerRunner[schedulables.TaskWrapper] {
	return &memoSched{
		state: &open{
			scheduler: &memSchedInstance{
				ticker,
				queue,
				latches.OncePair(),
			},
		},
	}
}

type state[I any] interface {
	TaskSchedulerRunner[schedulables.TaskWrapper]
	instance() I
}

type memoSched struct {
	state[*memSchedInstance]
}

func (s *memoSched) Close() error {
	err := s.state.Close()
	if err != nil {
		return err
	}
	s.state = &closed{}
	return nil
}

type memSchedInstance struct {
	ticker chrono.Ticker
	agenda agenda.Chrono[schedulables.TaskWrapper]
	sd     syncs.PairLatch
}

func (i *memSchedInstance) Shutdown() syncs.Waiter {
	shutdown := i.shutdown
	functions.Once(&shutdown)
	return i.sd.Waiter()
}

func (i *memSchedInstance) shutdown() {
	go func() {
		drainAgenda(i.agenda, func(runnable task.Runnable) bool {
			_ = runnable.Run()
			return true
		})
		i.sd.Signal()
	}()
}

type open struct {
	scheduler *memSchedInstance
}

func (s *open) Schedule(t schedulables.TaskWrapper) error {
	s.scheduler.agenda.Push(t)
	return nil
}

func (s *open) RunDue(now chrono.Time) error {
	return runDue(s.scheduler.agenda, now)
}

func (s *open) Shutdown() syncs.Waiter {
	return s.scheduler.Shutdown()
}

func (s *open) Seal() error {
	return nil
}

func (s *open) Unseal() error {
	return ErrOpen
}

func (s *open) Close() error {
	return nil
}

func (s *open) instance() *memSchedInstance {
	return s.scheduler
}

type sealed struct {
	scheduler *memSchedInstance
}

func (s *sealed) Schedule(_ schedulables.TaskWrapper) error {
	return ErrSealed
}

func (s *sealed) RunDue(now chrono.Time) error {
	return runDue(s.scheduler.agenda, now)
}

func (s *sealed) Shutdown() syncs.Waiter {
	return s.scheduler.Shutdown()
}

func (s *sealed) Seal() error {
	return ErrSealed
}

func (s *sealed) Unseal() error {
	return nil
}

func (s *sealed) Close() error {
	return nil
}

func (s *sealed) instance() *memSchedInstance {
	return s.scheduler
}

type closed struct{}

func (s *closed) Schedule(_ schedulables.TaskWrapper) error {
	return ErrClosed
}

func (s *closed) RunDue(t chrono.Time) error {
	return ErrClosed
}

func (s *closed) Drain(_ function.Predicate[task.Runnable]) int {
	return 0
}

func (s *closed) Shutdown() syncs.Waiter {
	return waiters.Complete()
}

func (s *closed) Seal() error {
	return ErrClosed
}

func (s *closed) Unseal() error {
	return ErrClosed
}

func (s *closed) Close() error {
	return ErrClosedTwice
}

func (s *closed) instance() *memSchedInstance {
	return nil
}

func drainAgenda[T task.Schedulable[chrono.Time]](agenda agenda.Chrono[T], fn function.Predicate[task.Runnable]) int {
	return agenda.Drain(func(t T) bool {
		return fn(t)
	})
}

func runDue(agenda agenda.Chrono[schedulables.TaskWrapper], now chrono.Time) error {
	e := make([]error, 0)
	agenda.RangeDue(now, func(t schedulables.TaskWrapper) bool {
		err := t.Run()
		if err != nil {
			e = append(e, err)
		}
		if next, ok := t.Next(now); ok {
			agenda.Push(schedulables.WrapTask(t.Unwrap(), next, t.Next))
		}
		return true
	})
	return errors.Join(e...)
}
