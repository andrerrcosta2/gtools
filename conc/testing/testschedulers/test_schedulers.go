// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package testschedulers

import (
	"sync"
	"sync/atomic"

	"github.com/andrerrcosta2/gtools/conc/timing/engines/sources"
	"github.com/andrerrcosta2/gtools/conc/timing/engines/steps"
	"github.com/andrerrcosta2/gtools/core/domain/contracts/chrono"
	"github.com/andrerrcosta2/gtools/core/domain/contracts/timing/scheduler"
)

func New(clock chrono.AdvanceableClock) *TestScheduler {
	return &TestScheduler{
		src:   sources.OfTimedSteps(),
		clock: clock,
	}
}

func (s *TestScheduler) NewTimer(cb func()) chrono.Timer {
	return &testTimer{
		s:  s,
		cb: cb,
	}
}

type TestScheduler struct {
	mtx    sync.Mutex
	src    sources.Timed
	clock  chrono.AdvanceableClock
	closed bool
}

func (s *TestScheduler) At(at chrono.Time, task scheduler.Task) (scheduler.Cancellable, error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	if s.closed {
		return nil, scheduler.ErrClosedScheduler
	}

	return s.at(at, task)
}

func (s *TestScheduler) After(d chrono.Duration, task scheduler.Task) (scheduler.Cancellable, error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	if s.closed {
		return nil, scheduler.ErrClosedScheduler
	}

	now := s.clock.Now()
	return s.at(now.Add(d), task)
}

func (s *TestScheduler) at(at chrono.Time, task scheduler.Task) (scheduler.Cancellable, error) {
	st := steps.FuncTimed(at, func() error {
		if task.Ctx != nil {
			select {
			case <-task.Ctx.Done():
				return nil
			default:
			}
		}
		_ = task.Job(task.Ctx)
		return nil
	})
	s.src.Push(st)
	return &cancellableStep{
		stopFunc: func() {
			st.Stop()
		},
		st: st,
	}, nil
}

func (s *TestScheduler) AdvanceTo(t chrono.Time) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	if s.closed {
		return
	}

	s.clock.StepTo(t)
	s.stepTo(t)
}

// Every schedules a repeating task at the given interval
func (s *TestScheduler) Every(interval chrono.Duration, task scheduler.Task) (scheduler.Cancellable, error) {
	start := s.clock.Now().Add(interval)
	return s.EveryFrom(start, interval, task)
}

func (s *TestScheduler) EveryFrom(start chrono.Time, interval chrono.Duration, task scheduler.Task) (scheduler.Cancellable, error) {
	if interval <= 0 {
		interval = chrono.Nanosecond
	}

	s.mtx.Lock()
	defer s.mtx.Unlock()
	if s.closed {
		return nil, scheduler.ErrClosedScheduler
	}

	var stopped atomic.Bool
	// recursive generator for recurring broadcasters
	var next func(at chrono.Time) steps.Timed

	next = func(at chrono.Time) steps.Timed {
		return steps.FuncTimedWithNext(
			at,
			func() error {
				if stopped.Load() {
					return nil
				}
				if task.Ctx != nil {
					select {
					case <-task.Ctx.Done():
						return nil
					default:
					}
				}
				return task.Job(task.Ctx)
			},
			func() (steps.Timed, bool) {
				if stopped.Load() {
					return nil, false
				}
				return next(at.Add(interval)), true
			},
		)
	}

	st := next(start)
	s.src.Push(st)

	return &cancellableStep{
		stopFunc: func() {
			stopped.Store(true)
			st.Stop()
		},
		st: st,
	}, nil
}

func (s *TestScheduler) StepTo(target chrono.Time) (executed bool, errs []error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	return s.stepTo(target)
}

func (s *TestScheduler) stepTo(target chrono.Time) (executed bool, errs []error) {
	for {
		next, ok := s.src.Peek()
		if !ok || next.At().After(target) {
			break
		}

		step, _ := s.src.Next()
		if !step.Stopped() {
			err := step.Exec()
			executed = true
			if err != nil {
				errs = append(errs, err)
			}
		}

		if n, ok := step.Next(); ok && !step.Stopped() {
			s.src.Push(n)
		}
	}
	return
}

func (s *TestScheduler) Close() error {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.closed = true
	return nil
}

func (s *TestScheduler) IsClosed() bool {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	return s.closed
}

func (s *TestScheduler) Now() chrono.Time {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	return s.clock.Now()
}

type testTimer struct {
	mtx     sync.Mutex
	st      steps.Timed
	s       *TestScheduler
	cb      func()
	stopped atomic.Bool
}

func (t *testTimer) Cancel() {
	t.Stop()
}

func (t *testTimer) Reset(d chrono.Duration) {
	t.mtx.Lock()
	defer t.mtx.Unlock()

	if t.s == nil {
		return
	}

	t.s.mtx.Lock()
	defer t.s.mtx.Unlock()

	// Stop previous step
	if t.st != nil {
		t.st.Stop()
	}

	t.stopped.Store(false)

	next := t.s.clock.Now().Add(d)

	t.st = steps.FuncTimed(next, func() error {
		if t.stopped.Load() {
			return nil
		}
		if t.cb != nil {
			t.cb()
		}
		return nil
	})

	t.s.src.Push(t.st)
}

func (t *testTimer) Stop() {
	t.mtx.Lock()
	defer t.mtx.Unlock()

	t.stopped.Store(true)

	if t.st != nil {
		t.st.Stop()
	}
}
