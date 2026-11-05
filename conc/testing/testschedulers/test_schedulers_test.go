// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package testschedulers

import (
	"context"
	"errors"
	"sync"
	"testing"

	"github.com/andrerrcosta2/gtools/core/chronos"
	"github.com/andrerrcosta2/gtools/core/domain/contracts/chrono"
	"github.com/andrerrcosta2/gtools/core/domain/contracts/timing/scheduler"
	"github.com/andrerrcosta2/gtools/core/testlite"
	"github.com/andrerrcosta2/gtools/core/testlite/assertlite"
)

func TestSchedulerSchedule(t *testing.T) {
	t.Run("should execute task when time advances", func(t *testing.T) {
		s := New(chronos.VirtualClock())
		defer testlite.CloseOrFail(t, s)

		executed := false
		task := scheduler.Task{
			Job: func(ctx context.Context) error {
				executed = true
				return nil
			},
			Ctx: context.Background(),
		}

		_, err := s.At(s.Now(), task)
		assertlite.NoError(t, err)

		// Advance time to execute the task (scheduled at current time)
		s.AdvanceTo(s.Now())

		assertlite.True(t, executed, "task should execute after advancing time")
	})

	t.Run("should handle nil context", func(t *testing.T) {
		s := New(chronos.VirtualClock())
		defer testlite.CloseOrFail(t, s)

		executed := false
		task := scheduler.Task{
			Job: func(ctx context.Context) error {
				executed = true
				return nil
			},
			Ctx: nil,
		}

		_, err := s.At(s.Now(), task)
		assertlite.NoError(t, err)

		s.AdvanceTo(s.Now())

		assertlite.True(t, executed, "task should execute with nil context")
	})

	t.Run("should handle task that returns error", func(t *testing.T) {
		s := New(chronos.VirtualClock())
		defer testlite.CloseOrFail(t, s)

		expectedErr := errors.New("task error")
		task := scheduler.Task{
			Job: func(ctx context.Context) error {
				return expectedErr
			},
			Ctx: context.Background(),
		}

		_, err := s.At(s.Now(), task)
		assertlite.NoError(t, err, "scheduler should not return task error")

		// Should not panic when task returns error
		assertlite.NoPanics(t, func() {
			s.AdvanceTo(s.Now())
		})
	})

	t.Run("should reject task when scheduler is closed", func(t *testing.T) {
		s := New(chronos.VirtualClock())
		testlite.CloseOrFail(t, s)

		task := scheduler.Task{
			Job: func(ctx context.Context) error { return nil },
			Ctx: context.Background(),
		}

		_, err := s.At(s.Now(), task)

		assertlite.IsErrorOf(t, err, scheduler.ErrClosedScheduler)
	})
}

func TestSchedulerClose(t *testing.T) {
	t.Run("should close successfully", func(t *testing.T) {
		s := New(chronos.VirtualClock())

		err := s.Close()
		assertlite.NoError(t, err)
		assertlite.True(t, s.IsClosed(), "scheduler should be closed")
	})

	t.Run("should be idempotent", func(t *testing.T) {
		s := New(chronos.VirtualClock())

		err1 := s.Close()
		err2 := s.Close()

		assertlite.NoError(t, err1)
		assertlite.NoError(t, err2)
		assertlite.True(t, s.IsClosed(), "scheduler should remain closed")
	})

	t.Run("closed scheduler should not execute task", func(t *testing.T) {
		s := New(chronos.VirtualClock())

		executed := false
		task := scheduler.Task{
			Job: func(ctx context.Context) error {
				executed = true
				return nil
			},
			Ctx: context.Background(),
		}

		_, err := s.At(s.Now(), task)

		assertlite.NoError(t, err)

		testlite.CloseOrFail(t, s)
		s.AdvanceTo(s.Now())

		assertlite.False(t, executed, "closed scheduler should not execute task")
	})
}

func TestSchedulerIsClosed(t *testing.T) {
	t.Run("should return false for new scheduler", func(t *testing.T) {
		s := New(chronos.VirtualClock())
		defer testlite.CloseOrFail(t, s)

		assertlite.False(t, s.IsClosed(), "new scheduler should not be closed")
	})

	t.Run("should return true after close", func(t *testing.T) {
		s := New(chronos.VirtualClock())
		testlite.CloseOrFail(t, s)
		assertlite.True(t, s.IsClosed(), "scheduler should be closed after Close()")
	})
}

func TestSchedulerConcurrent(t *testing.T) {
	t.Run("should handle concurrent scheduling", func(t *testing.T) {
		s := New(chronos.VirtualClock())
		defer testlite.CloseOrFail(t, s)

		const numTasks = 100
		executed := make(chan bool, numTasks)
		var wg sync.WaitGroup
		wg.Add(numTasks)

		for i := 0; i < numTasks; i++ {
			go func() {
				task := scheduler.Task{
					Job: func(ctx context.Context) error {
						executed <- true
						return nil
					},
					Ctx: context.Background(),
				}
				_, err := s.At(s.Now(), task)

				wg.Done()
				assertlite.NoError(t, err)
			}()
		}

		wg.Wait()

		// Advance time to execute all task
		s.AdvanceTo(s.Now())

		// Check that all task executed
		for i := 0; i < numTasks; i++ {
			select {
			case <-executed:
				// Task executed
			default:
				t.Errorf("task %d did not execute", i)
			}
		}
	})
}

func TestDelayedAt(t *testing.T) {
	t.Run("should execute task at absolute time", func(t *testing.T) {
		s := New(chronos.VirtualClock())
		defer testlite.CloseOrFail(t, s)

		executed := false
		task := scheduler.Task{
			Job: func(ctx context.Context) error {
				executed = true
				return nil
			},
			Ctx: context.Background(),
		}

		futureTime := s.Now().Add(chrono.Sec(10))
		_, err := s.At(futureTime, task)
		assertlite.NoError(t, err)

		// Advance partially - should not execute
		s.AdvanceTo(s.Now().Add(chrono.Sec(5)))
		assertlite.False(t, executed, "task should not execute before scheduled time")

		// Advance to scheduled time - should execute
		s.AdvanceTo(futureTime)
		assertlite.True(t, executed, "task should execute at scheduled time")
	})

	t.Run("should return cancellable handle", func(t *testing.T) {
		s := New(chronos.VirtualClock())
		defer testlite.CloseOrFail(t, s)

		futureTime := s.Now().Add(chrono.Hour)
		task := scheduler.Task{
			Job: func(ctx context.Context) error { return nil },
			Ctx: context.Background(),
		}

		cancel, err := s.At(futureTime, task)
		assertlite.NoError(t, err)
		assertlite.NotNil(t, cancel)
		assertlite.IsTypeOf[scheduler.Cancellable](t, cancel)
	})

	t.Run("cancelled task should not execute", func(t *testing.T) {
		s := New(chronos.VirtualClock())
		defer testlite.CloseOrFail(t, s)

		executed := false
		task := scheduler.Task{
			Job: func(ctx context.Context) error {
				executed = true
				return nil
			},
			Ctx: context.Background(),
		}

		futureTime := s.Now().Add(chrono.Sec(10))
		cancel, err := s.At(futureTime, task)
		assertlite.NoError(t, err)

		cancel.Cancel()
		s.AdvanceTo(futureTime)

		assertlite.False(t, executed, "cancelled task should not execute")
	})

	t.Run("should reject when scheduler closed", func(t *testing.T) {
		s := New(chronos.VirtualClock())
		testlite.CloseOrFail(t, s)

		task := scheduler.Task{
			Job: func(ctx context.Context) error { return nil },
			Ctx: context.Background(),
		}

		_, err := s.At(chrono.Now().Add(chrono.Hour), task)
		assertlite.IsErrorOf(t, err, scheduler.ErrClosedScheduler)
	})
}

func TestDelayedAfter(t *testing.T) {
	t.Run("should execute task after duration", func(t *testing.T) {
		s := New(chronos.VirtualClock())
		defer testlite.CloseOrFail(t, s)

		executed := false
		task := scheduler.Task{
			Job: func(ctx context.Context) error {
				executed = true
				return nil
			},
			Ctx: context.Background(),
		}

		_, err := s.After(chrono.Sec(10), task)
		assertlite.NoError(t, err)

		// Advance partially
		s.AdvanceTo(s.Now().Add(chrono.Sec(5)))
		assertlite.False(t, executed, "task should not execute before delay")

		// Advance to completion
		s.AdvanceTo(s.Now().Add(chrono.Sec(10)))
		assertlite.True(t, executed, "task should execute after delay")
	})

	t.Run("should handle zero duration", func(t *testing.T) {
		s := New(chronos.VirtualClock())
		defer testlite.CloseOrFail(t, s)

		executed := false
		task := scheduler.Task{
			Job: func(ctx context.Context) error {
				executed = true
				return nil
			},
			Ctx: context.Background(),
		}

		_, err := s.After(0, task)
		assertlite.NoError(t, err)

		// Zero duration task execute immediately when time advances
		s.AdvanceTo(s.Now())
		assertlite.True(t, executed, "task with zero delay should execute when time advances")
	})
}

func TestRecurringEvery(t *testing.T) {
	t.Run("should execute recurring task at each interval", func(t *testing.T) {
		s := New(chronos.VirtualClock())
		defer testlite.CloseOrFail(t, s)

		executionCount := 0
		task := scheduler.Task{
			Job: func(ctx context.Context) error {
				executionCount++
				return nil
			},
			Ctx: context.Background(),
		}

		tenSeconds := chrono.Sec(10)
		_, err := s.Every(tenSeconds, task)
		assertlite.NoError(t, err)

		now := s.Now()

		s.AdvanceTo(now)
		assertlite.True(t, 0 == executionCount, "first execution shouldn't fire at start")

		// First execution after ten seconds
		s.AdvanceTo(now.Add(tenSeconds))
		assertlite.True(t, 1 == executionCount, "second execution after tenSeconds")

		// Second execution after twenty seconds
		s.AdvanceTo(now.Add(2 * tenSeconds))
		assertlite.True(t, 2 == executionCount, "third execution after 2 intervals")
	})

	t.Run("should handle zero interval", func(t *testing.T) {
		s := New(chronos.VirtualClock())
		defer testlite.CloseOrFail(t, s)

		executionCount := 0
		task := scheduler.Task{
			Job: func(ctx context.Context) error {
				executionCount++
				return nil
			},
			Ctx: context.Background(),
		}

		_, err := s.Every(0, task)
		assertlite.NoError(t, err)

		now := s.Now()

		// First execution fires immediately
		s.AdvanceTo(now)
		assertlite.Equals(t, 1, executionCount, "zero interval should fire immediately")

		// Next tick should fire almost immediately (nanosecond later)
		s.AdvanceTo(now.Add(chrono.Nanosecond))
		assertlite.True(t, executionCount >= 2, "zero interval should trigger multiple times quickly")
	})

	t.Run("cancelled recurring task should not execute again", func(t *testing.T) {
		s := New(chronos.VirtualClock())
		defer testlite.CloseOrFail(t, s)

		executionCount := 0
		task := scheduler.Task{
			Job: func(ctx context.Context) error {
				executionCount++
				return nil
			},
			Ctx: context.Background(),
		}

		tenSeconds := chrono.Sec(10)
		cancel, err := s.Every(tenSeconds, task)
		assertlite.NoError(t, err)

		now := s.Now()

		s.AdvanceTo(now)
		assertlite.Equals(t, 0, executionCount, "first execution shouldn't fire immediately")

		// Should execute first
		s.AdvanceTo(now.Add(tenSeconds))
		assertlite.True(t, executionCount == 1, "first execution shouldn happen after first tenSeconds")

		// Cancel and advance again
		cancel.Cancel()
		s.AdvanceTo(now.Add(2 * tenSeconds))
		assertlite.Equals(t, 1, executionCount, "cancelled task should not execute again")
	})
}

func TestRecurringEveryFrom(t *testing.T) {
	t.Run("should execute recurring task at each interval", func(t *testing.T) {
		s := New(chronos.VirtualClock())
		defer testlite.CloseOrFail(t, s)

		executionCount := 0
		task := scheduler.Task{
			Job: func(ctx context.Context) error {
				executionCount++
				return nil
			},
			Ctx: context.Background(),
		}

		interval := chrono.Sec(10)
		start := s.Now().Add(chrono.Sec(5)) // start 5s in the future
		_, err := s.EveryFrom(start, interval, task)
		assertlite.NoError(t, err)

		// Before start, nothing executed
		s.AdvanceTo(s.Now().Add(chrono.Sec(4)))
		assertlite.Equals(t, 0, executionCount, "no execution before start")

		// First execution at start
		s.AdvanceTo(start)
		assertlite.Equals(t, 1, executionCount, "first execution should fire at start")

		// Second execution after one interval
		s.AdvanceTo(start.Add(interval))
		assertlite.Equals(t, 2, executionCount, "second execution after interval")

		// Third execution after two intervals
		s.AdvanceTo(start.Add(2 * interval))
		assertlite.Equals(t, 3, executionCount, "third execution after 2 intervals")
	})

	t.Run("should handle zero interval", func(t *testing.T) {
		s := New(chronos.VirtualClock())
		defer testlite.CloseOrFail(t, s)

		executionCount := 0
		task := scheduler.Task{
			Job: func(ctx context.Context) error {
				executionCount++
				return nil
			},
			Ctx: context.Background(),
		}

		start := s.Now().Add(chrono.Sec(5))
		_, err := s.EveryFrom(start, 0, task)
		assertlite.NoError(t, err)

		// First execution fires at start
		s.AdvanceTo(start)
		assertlite.Equals(t, 1, executionCount, "zero interval should fire at start")

		// Next tick should fire almost immediately (nanosecond later)
		s.AdvanceTo(start.Add(chrono.Nanosecond))
		assertlite.True(t, executionCount >= 2, "zero interval should trigger multiple times quickly")
	})

	t.Run("cancelled recurring task should not execute again", func(t *testing.T) {
		s := New(chronos.VirtualClock())
		defer testlite.CloseOrFail(t, s)

		executionCount := 0
		task := scheduler.Task{
			Job: func(ctx context.Context) error {
				executionCount++
				return nil
			},
			Ctx: context.Background(),
		}

		interval := chrono.Sec(10)
		start := s.Now().Add(chrono.Sec(5))
		cancel, err := s.EveryFrom(start, interval, task)
		assertlite.NoError(t, err)

		// First execution at start
		s.AdvanceTo(start)
		assertlite.Equals(t, 1, executionCount, "first execution should fire at start")

		// Cancel and advance again
		cancel.Cancel()
		s.AdvanceTo(start.Add(interval))
		assertlite.Equals(t, 1, executionCount, "cancelled task should not execute again")
	})
}

func TestRecurringEveryErrorHandling(t *testing.T) {
	t.Run("recurring task continues after returning error", func(t *testing.T) {
		s := New(chronos.VirtualClock())
		defer testlite.CloseOrFail(t, s)

		count := 0
		task := scheduler.NewTask(func(ctx context.Context) error {
			count++
			return errors.New("expected error")
		}, context.Background())

		tenSeconds := chrono.Sec(10)
		_, err := s.Every(tenSeconds, task)
		assertlite.NoError(t, err)

		start := s.Now()

		// First execution (with error)
		s.AdvanceTo(start.Add(tenSeconds))
		assertlite.Equals(t, 1, count, "expected 1 executions, but got '%d'", count)

		// Second execution should still happen
		s.AdvanceTo(start.Add(2 * tenSeconds))
		assertlite.Equals(t, 2, count, "expected 2 executions, but got '%d'", count)
	})

	t.Run("recurring task stops if context cancelled", func(t *testing.T) {
		s := New(chronos.VirtualClock())
		defer testlite.CloseOrFail(t, s)

		ctx, cancel := context.WithCancel(context.Background())
		executionCount := 0

		task := scheduler.Task{
			Job: func(ctx context.Context) error {
				executionCount++
				return nil
			},
			Ctx: ctx,
		}

		interval := chrono.Sec(10)
		_, err := s.Every(interval, task)
		assertlite.NoError(t, err)

		// First execution
		s.AdvanceTo(s.Now().Add(interval))
		assertlite.Equals(t, 1, executionCount, "expected 1 executions, but got '%d'", executionCount)

		// Cancel context and advance
		cancel()
		s.AdvanceTo(s.Now().Add(2 * interval))

		// Should not execute again due to cancelled context
		assertlite.Equals(t, 1, executionCount, "recurring task should stop when context is cancelled")
	})
}

func TestTimerFactoryNewTimer(t *testing.T) {
	t.Run("should create new timer", func(t *testing.T) {
		s := New(chronos.VirtualClock())
		defer testlite.CloseOrFail(t, s)

		called := false
		timer := s.NewTimer(func() {
			called = true
		})

		assertlite.NotNil(t, timer)
		assertlite.False(t, called, "timer should not fire without reset")
	})

	t.Run("timer should fire after reset with duration", func(t *testing.T) {
		s := New(chronos.VirtualClock())
		defer testlite.CloseOrFail(t, s)

		called := false
		timer := s.NewTimer(func() {
			called = true
		})

		timer.Reset(chrono.Sec(10))

		start := s.Now()

		// Advance partially
		s.AdvanceTo(start.Add(chrono.Sec(5)))
		assertlite.False(t, called, "timer should not fire before duration")

		// Advance to completion
		s.AdvanceTo(start.Add(chrono.Sec(10)))
		assertlite.True(t, called, "timer should fire after duration")
	})

	t.Run("timer should fire with zero duration", func(t *testing.T) {
		s := New(chronos.VirtualClock())
		defer testlite.CloseOrFail(t, s)

		called := false
		timer := s.NewTimer(func() {
			called = true
		})

		timer.Reset(0)

		// Zero duration timers fire immediately when time advances
		s.AdvanceTo(s.Now())
		assertlite.True(t, called, "timer should fire immediately with zero duration")
	})

	t.Run("timer should be stoppable", func(t *testing.T) {
		s := New(chronos.VirtualClock())
		defer testlite.CloseOrFail(t, s)

		called := false
		timer := s.NewTimer(func() {
			called = true
		})

		timer.Reset(chrono.Ms(10))
		timer.Stop()
		s.AdvanceTo(s.Now().Add(chrono.Sec(10)))

		assertlite.False(t, called, "timer should not fire after being stopped")
	})

	t.Run("multiple resets reschedule timer", func(t *testing.T) {
		s := New(chronos.VirtualClock())
		defer testlite.CloseOrFail(t, s)

		count := 0
		timer := s.NewTimer(func() {
			count++
		})

		timer.Reset(chrono.Sec(10))
		timer.Reset(chrono.Sec(20))

		// Advance to first scheduled time - should not fire
		s.AdvanceTo(s.Now().Add(chrono.Sec(10)))
		assertlite.Equals(t, 0, count, "timer should not fire at first scheduled time after reset")

		// Advance to second scheduled time - should fire
		s.AdvanceTo(s.Now().Add(chrono.Sec(20)))
		assertlite.Equals(t, 1, count, "timer should fire at new scheduled time")
	})
}

func TestAdvanceable(t *testing.T) {
	t.Run("should advance to exact time", func(t *testing.T) {
		s := New(chronos.VirtualClock())
		defer testlite.CloseOrFail(t, s)

		executed := false
		task := scheduler.Task{
			Job: func(ctx context.Context) error {
				executed = true
				return nil
			},
			Ctx: context.Background(),
		}

		futureTime := s.Now().Add(chrono.Sec(10))
		_, err := s.At(futureTime, task)
		assertlite.NoError(t, err)

		s.AdvanceTo(futureTime)
		assertlite.True(t, executed)
	})

	t.Run("should advance beyond scheduled time", func(t *testing.T) {
		s := New(chronos.VirtualClock())
		defer testlite.CloseOrFail(t, s)

		executed := false
		task := scheduler.Task{
			Job: func(ctx context.Context) error {
				executed = true
				return nil
			},
			Ctx: context.Background(),
		}

		futureTime := s.Now().Add(chrono.Sec(10))
		_, err := s.At(futureTime, task)
		assertlite.NoError(t, err)

		// Advance beyond scheduled time
		s.AdvanceTo(s.Now().Add(chrono.Sec(15)))
		assertlite.True(t, executed)
	})

	t.Run("should execute multiple task in order", func(t *testing.T) {
		s := New(chronos.VirtualClock())
		defer testlite.CloseOrFail(t, s)

		var executionOrder []int

		for i := 1; i <= 3; i++ {
			scheduleTime := s.Now().Add(chrono.Sec(i * 10))
			task := scheduler.Task{
				Job: func(ctx context.Context) error {
					executionOrder = append(executionOrder, i)
					return nil
				},
				Ctx: context.Background(),
			}
			_, err := s.At(scheduleTime, task)
			assertlite.NoError(t, err)
		}

		// Advance to execute all
		s.AdvanceTo(s.Now().Add(chrono.Sec(30)))

		assertlite.Equals(t, []int{1, 2, 3}, executionOrder, "task should execute in scheduled order")
	})
}

func TestIntegration(t *testing.T) {
	t.Run("complex scenario with multiple task types", func(t *testing.T) {
		s := New(chronos.VirtualClock())
		defer testlite.CloseOrFail(t, s)

		var results []string

		// One-shot delayed task
		delayedTask := scheduler.NewTask(func(ctx context.Context) error {
			results = append(results, "delayed")
			return nil
		}, context.Background())

		_, err := s.After(chrono.Sec(10), delayedTask)
		assertlite.NoError(t, err)

		// Recurring task
		recurringCount := 0
		recurringTask := scheduler.NewTask(func(ctx context.Context) error {
			recurringCount++
			results = append(results, "recurring")
			return nil
		}, context.Background())

		cancelRecurring, err := s.Every(chrono.Sec(5), recurringTask)
		assertlite.NoError(t, err)

		// Timer task
		timerCalled := false
		timer := s.NewTimer(func() {
			timerCalled = true
			results = append(results, "timer")
		})
		timer.Reset(chrono.Sec(15))

		start := s.Now()

		// At 5s - recurring first execution
		s.AdvanceTo(start.Add(chrono.Sec(5)))
		assertlite.Equals(t, 1, recurringCount)
		assertlite.Contains(t, results, "recurring")

		// At 10s - delayed task and recurring second execution
		s.AdvanceTo(start.Add(chrono.Sec(10)))
		assertlite.Equals(t, 2, recurringCount)
		assertlite.Contains(t, results, "delayed")

		// At 15s - timer and recurring third execution
		s.AdvanceTo(start.Add(chrono.Sec(15)))
		assertlite.Equals(t, 3, recurringCount)
		assertlite.True(t, timerCalled)

		// Cancel recurring and advance
		cancelRecurring.Cancel()
		s.AdvanceTo(start.Add(chrono.Sec(20)))
		assertlite.Equals(t, 3, recurringCount, "recurring task should stop after cancel")
	})

	// Add to TestIntegration complex scenario
	t.Run("context cancellation during complex scenario", func(t *testing.T) {
		s := New(chronos.VirtualClock())
		defer testlite.CloseOrFail(t, s)

		ctx, cancel := context.WithCancel(context.Background())

		executed := make(map[string]bool)

		// Task that respects context
		task := scheduler.Task{
			Job: func(ctx context.Context) error {
				select {
				case <-ctx.Done():
					return ctx.Err()
				default:
					executed["context-aware"] = true
					return nil
				}
			},
			Ctx: ctx,
		}
		_, err := s.After(10*chrono.Second, task)
		assertlite.NoError(t, err)

		// Cancel before execution
		cancel()
		s.AdvanceTo(s.Now().Add(chrono.Sec(10)))

		assertlite.False(t, executed["context-aware"], "cancelled context should prevent execution")
	})
}

func TestSchedulerContextCancellation(t *testing.T) {
	t.Run("should skip task if context cancelled before execution", func(t *testing.T) {
		s := New(chronos.VirtualClock())
		defer testlite.CloseOrFail(t, s)

		ctx, cancel := context.WithCancel(context.Background())
		executed := false

		task := scheduler.Task{
			Job: func(ctx context.Context) error {
				executed = true
				return nil
			},
			Ctx: ctx,
		}

		future := s.Now().Add(chrono.Sec(10))
		_, err := s.At(future, task)
		assertlite.NoError(t, err)

		cancel() // Cancel BEFORE advancing time
		s.AdvanceTo(future)

		assertlite.False(t, executed, "cancelled context should prevent task execution")
	})

	t.Run("should execute task if context cancelled after execution starts", func(t *testing.T) {
		s := New(chronos.VirtualClock())
		defer testlite.CloseOrFail(t, s)

		ctx, cancel := context.WithCancel(context.Background())
		executed := false
		executionStarted := make(chan struct{})

		task := scheduler.Task{
			Job: func(ctx context.Context) error {
				close(executionStarted) // Wait that execution began
				executed = true
				return nil
			},
			Ctx: ctx,
		}

		future := s.Now().Add(chrono.Sec(10))
		_, err := s.At(future, task)
		assertlite.NoError(t, err)

		s.AdvanceTo(future)

		// Wait for execution to start before cancelling
		<-executionStarted
		cancel() // Cancel after execution has begun

		assertlite.True(t, executed, "task should execute even if context cancelled during execution")
	})
}

func TestAdvanceToEdgeCases(t *testing.T) {
	t.Run("should handle AdvanceTo with empty heap", func(t *testing.T) {
		s := New(chronos.VirtualClock())
		defer testlite.CloseOrFail(t, s)

		start := s.Now()
		target := start.Add(chrono.Hour)

		assertlite.NoPanics(t, func() {
			s.AdvanceTo(target)
		})

		assertlite.Equals(t, target.UnixNano(), s.Now().UnixNano())
	})

	t.Run("should handle AdvanceTo backwards in time", func(t *testing.T) {
		s := New(chronos.VirtualClock())
		defer testlite.CloseOrFail(t, s)

		initial := s.Now()
		future := initial.Add(chrono.Hours(1))

		// Advance forward first
		s.AdvanceTo(future)
		assertlite.Equals(t, future, s.Now())

		// Advance backwards - should update s.now but not re-execute past task
		executed := false
		task := scheduler.Task{
			Job: func(ctx context.Context) error {
				executed = true
				return nil
			},
			Ctx: context.Background(),
		}
		_, err := s.At(initial.Add(chrono.Min(30)), task)
		assertlite.NoError(t, err)

		s.AdvanceTo(initial) // Go back in time

		assertlite.Equals(t, initial, s.Now())
		assertlite.False(t, executed, "task scheduled in 'future' should not execute when time "+
			"goes backwards")
	})

	t.Run("should execute task scheduled exactly at target time", func(t *testing.T) {
		s := New(chronos.VirtualClock())
		defer testlite.CloseOrFail(t, s)

		executed := false
		target := s.Now().Add(chrono.Sec(10))

		task := scheduler.Task{
			Job: func(ctx context.Context) error {
				executed = true
				return nil
			},
			Ctx: context.Background(),
		}

		_, err := s.At(target, task)
		assertlite.NoError(t, err)

		s.AdvanceTo(target) // Advance exactly to scheduled time

		assertlite.True(t, executed, "task scheduled at exact target time should execute")
	})
}

func TestSchedulerConcurrentAdvanceTo(t *testing.T) {
	t.Run("concurrent AdvanceTo calls should be safe", func(t *testing.T) {
		s := New(chronos.VirtualClock())
		defer testlite.CloseOrFail(t, s)

		const numAdvances = 50
		var wg sync.WaitGroup
		wg.Add(numAdvances)

		// Schedule task at various times
		for i := 0; i < 20; i++ {
			task := scheduler.Task{
				Job: func(ctx context.Context) error { return nil },
				Ctx: context.Background(),
			}
			_, err := s.At(s.Now().Add(chrono.Duration(i)*chrono.Second), task)
			assertlite.NoError(t, err)
		}

		// Launch concurrent AdvanceTo calls with overlapping targets
		for i := 0; i < numAdvances; i++ {
			go func(step int) {
				defer wg.Done()
				target := s.Now().Add(chrono.Sec(step * 2))
				assertlite.NoPanics(t, func() {
					s.AdvanceTo(target)
				})
			}(i)
		}

		wg.Wait()

		// Verify scheduler is still in valid states
		assertlite.NoPanics(t, func() {
			testlite.CloseOrFail(t, s)
		})
	})

	t.Run("concurrent scheduling and advancing should be safe", func(t *testing.T) {
		s := New(chronos.VirtualClock())
		defer testlite.CloseOrFail(t, s)

		var executed sync.Map
		wgSchedule := sync.WaitGroup{}
		wgAdvance := sync.WaitGroup{}

		// Goroutines scheduling task
		for i := 0; i < 20; i++ {
			wgSchedule.Add(1)
			go func(id int) {
				defer wgSchedule.Done()
				task := scheduler.Task{
					Job: func(ctx context.Context) error {
						executed.Store(id, true)
						return nil
					},
					Ctx: context.Background(),
				}
				// Schedule at slightly different times to interleave with advancing
				_, err := s.After(chrono.Ms(10), task)
				assertlite.NoError(t, err)
			}(i)
		}

		wgSchedule.Wait()

		// Goroutines advancing time
		for i := 0; i < 10; i++ {
			wgAdvance.Add(1)
			go func(step int) {
				defer wgAdvance.Done()
				inc := chrono.Ms(step * 5)
				s.AdvanceTo(s.Now().Add(inc))
			}(i)
		}

		wgAdvance.Wait()
		s.AdvanceTo(s.Now().Add(chrono.Second)) // Drain remaining

		// Verify all task eventually executed
		for i := 0; i < 20; i++ {
			val, ok := executed.Load(i)
			assertlite.True(t, ok && val.(bool), "task %d should have executed", i)
		}
	})
}

func TestCancellableIdempotency(t *testing.T) {
	t.Run("Cancel called multiple times should be safe", func(t *testing.T) {
		s := New(chronos.VirtualClock())
		defer testlite.CloseOrFail(t, s)

		future := s.Now().Add(chrono.Hour)
		task := scheduler.Task{
			Job: func(ctx context.Context) error { return nil },
			Ctx: context.Background(),
		}

		cancel, err := s.At(future, task)
		assertlite.NoError(t, err)

		// Cancel multiple times
		assertlite.NoPanics(t, func() {
			cancel.Cancel()
			cancel.Cancel()
			cancel.Cancel()
		})

		s.AdvanceTo(future)
		// Task should not execute (cancelled on first call)
	})

	t.Run("Stop called multiple times on timer should be safe", func(t *testing.T) {
		s := New(chronos.VirtualClock())
		defer testlite.CloseOrFail(t, s)

		called := false
		timer := s.NewTimer(func() { called = true })

		timer.Reset(chrono.Sec(10))

		assertlite.NoPanics(t, func() {
			timer.Stop()
			timer.Stop()
			timer.Stop()
		})

		s.AdvanceTo(s.Now().Add(chrono.Sec(10)))
		assertlite.False(t, called, "timer should not fire after being stopped")
	})

	t.Run("Cancel after execution should be safe", func(t *testing.T) {
		s := New(chronos.VirtualClock())
		defer testlite.CloseOrFail(t, s)

		executed := false
		task := scheduler.Task{
			Job: func(ctx context.Context) error {
				executed = true
				return nil
			},
			Ctx: context.Background(),
		}

		future := s.Now().Add(chrono.Sec(10))
		cancel, err := s.At(future, task)
		assertlite.NoError(t, err)

		// Execute first
		s.AdvanceTo(future)
		assertlite.True(t, executed)

		// Cancel after execution - should not panic
		assertlite.NoPanics(t, func() {
			cancel.Cancel()
		})
	})
}

func TestSchedulerSameTimestamp(t *testing.T) {
	t.Run("should execute all task scheduled at same time", func(t *testing.T) {
		s := New(chronos.VirtualClock())
		defer testlite.CloseOrFail(t, s)

		const numTasks = 10
		executed := make([]bool, numTasks)
		target := s.Now().Add(chrono.Sec(10))

		for i := 0; i < numTasks; i++ {
			idx := i
			task := scheduler.Task{
				Job: func(ctx context.Context) error {
					executed[idx] = true
					return nil
				},
				Ctx: context.Background(),
			}
			_, err := s.At(target, task)
			assertlite.NoError(t, err)
		}

		s.AdvanceTo(target)

		for i := 0; i < numTasks; i++ {
			assertlite.True(t, executed[i], "task %d should have executed", i)
		}
	})

	t.Run("should handle mix of one-shot and recurring at same time (Every)", func(t *testing.T) {
		s := New(chronos.VirtualClock())
		defer testlite.CloseOrFail(t, s)

		oneShotExecuted := false
		count := 0
		start := s.Now()

		// happens on every 5 seconds, starting at now + 5
		_, err := s.Every(chrono.Sec(5), scheduler.Task{
			Job: func(ctx context.Context) error {
				count++
				return nil
			},
			Ctx: context.Background(),
		})
		assertlite.NoError(t, err)

		// Target time 10s ahead
		target := start.Add(chrono.Sec(15))

		// One-shot at target
		_, err = s.At(target, scheduler.Task{
			Job: func(ctx context.Context) error {
				oneShotExecuted = true
				return nil
			},
			Ctx: context.Background(),
		})
		assertlite.NoError(t, err)

		// Advance the clock to target
		s.AdvanceTo(target)

		assertlite.True(t, oneShotExecuted)
		assertlite.Equals(t, 3, count, "recurring should fire every 5s, starting from now")
	})

	t.Run("should handle mix of one-shot and recurring at same time (EveryFrom)", func(t *testing.T) {
		s := New(chronos.VirtualClock())
		defer testlite.CloseOrFail(t, s)

		oneShotExecuted := false
		recurringCount := 0

		target := s.Now().Add(chrono.Sec(10))

		// Recurring starting exactly at target
		_, err := s.EveryFrom(target, chrono.Sec(5), scheduler.Task{
			Job: func(ctx context.Context) error {
				recurringCount++
				return nil
			},
			Ctx: context.Background(),
		})
		assertlite.NoError(t, err)

		// One-shot at same target
		_, err = s.At(target, scheduler.Task{
			Job: func(ctx context.Context) error {
				oneShotExecuted = true
				return nil
			},
			Ctx: context.Background(),
		})
		assertlite.NoError(t, err)

		// Advance clock to target
		s.AdvanceTo(target)

		assertlite.True(t, oneShotExecuted)
		assertlite.Equals(t, 1, recurringCount, "recurring should fire first time at target")
	})
}
