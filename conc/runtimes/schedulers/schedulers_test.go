package schedulers

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/andrerrcosta2/gtools/conc/timing/schedulers/internal/cancellation"
	"github.com/andrerrcosta2/gtools/core/chronos"
	"github.com/andrerrcosta2/gtools/core/domain/contracts/chrono"
	"github.com/andrerrcosta2/gtools/core/domain/contracts/chrono/durations"
	"github.com/andrerrcosta2/gtools/core/domain/contracts/timing/scheduler"
	"github.com/andrerrcosta2/gtools/core/testlite"
	"github.com/andrerrcosta2/gtools/core/testlite/assertlite"
)

func TestScheduler_New(t *testing.T) {
	t.Run("creates scheduler with system clock by default", func(t *testing.T) {
		s := New()
		assertlite.NotNil(t, s, "scheduler should not be nil")
		assertlite.False(t, s.IsClosed(), "new scheduler should not be closed")
	})

	t.Run("creates scheduler with custom clock", func(t *testing.T) {
		virtual := chronos.VirtualClock()
		s := WithClock(virtual)
		assertlite.NotNil(t, s, "scheduler should not be nil")
		assertlite.False(t, s.IsClosed(), "new scheduler should not be closed")
	})

	t.Run("uses system clock when nil provided", func(t *testing.T) {
		s := WithClock(nil)
		assertlite.NotNil(t, s, "scheduler should not be nil")
	})
}

func TestScheduler_At(t *testing.T) {
	t.Run("executes task eventAt specified time", func(t *testing.T) {
		clock := chronos.VirtualClock()
		s := WithClock(clock)
		defer testlite.CloseOrFail(t, s)

		var executed bool
		targetTime := clock.Now().Add(chrono.Hour)

		_, err := s.At(targetTime, scheduler.Task{
			Job: func(_ context.Context) error {
				executed = true
				return nil
			},
			Ctx: context.Background(),
		})
		assertlite.NoError(t, err)

		// Advance to just before target
		s.AdvanceTo(targetTime.Add(-chrono.Nanosecond))
		assertlite.False(t, executed, "task should not execute before time")

		// Advance to target
		assertlite.NoPanics(t, func() {
			s.AdvanceTo(targetTime)
		})

		assertlite.True(t, executed, "task should execute eventAt target time")
	})

	t.Run("returns nil cancellable for nil job", func(t *testing.T) {
		clock := chronos.VirtualClock()
		s := WithClock(clock)
		defer testlite.CloseOrFail(t, s)

		cancellable, err := s.At(clock.Now().Add(chrono.Hour), scheduler.Task{
			Job: nil,
			Ctx: context.Background(),
		})
		assertlite.IsErrorOf(t, err, ErrNilJob)
		assertlite.IsTypeOf[cancellation.NoopHandler](t, cancellable)
	})

	t.Run("cancels task before execution", func(t *testing.T) {
		clock := chronos.VirtualClock()
		s := WithClock(clock)
		defer testlite.CloseOrFail(t, s)

		var executed bool
		targetTime := clock.Now().Add(chrono.Hour)

		cancellable, err := s.At(targetTime, scheduler.Task{
			Job: func(ctx context.Context) error {
				executed = true
				return nil
			},
			Ctx: context.Background(),
		})
		assertlite.NoError(t, err)
		assertlite.NotNil(t, cancellable)

		// Cancel before execution
		cancellable.Cancel()

		s.AdvanceTo(targetTime)
		assertlite.False(t, executed, "cancelled task should not execute")
	})

	t.Run("uses background context when nil provided", func(t *testing.T) {
		clock := chronos.VirtualClock()
		s := WithClock(clock)
		defer testlite.CloseOrFail(t, s)

		var executed bool
		targetTime := clock.Now().Add(chrono.Hour)

		_, err := s.At(targetTime, scheduler.Task{
			Job: func(ctx context.Context) error {
				executed = true
				// Verify context is not nil
				if ctx == nil {
					t.Error("context should not be nil")
				}
				return nil
			},
			Ctx: nil, // Explicitly nil
		})
		assertlite.NoError(t, err)

		s.AdvanceTo(targetTime)
		assertlite.True(t, executed)
	})

	t.Run("returns error when scheduler is closed", func(t *testing.T) {
		clock := chronos.VirtualClock()
		s := WithClock(clock)
		assertlite.NoError(t, s.Close())

		_, err := s.At(clock.Now().Add(chrono.Hour), scheduler.Task{
			Job: func(ctx context.Context) error { return nil },
		})
		assertlite.IsErrorOf(t, err, scheduler.ErrClosedScheduler)
	})

	t.Run("handles multiple task eventAt same time", func(t *testing.T) {
		clock := chronos.VirtualClock()
		s := WithClock(clock)
		defer testlite.CloseOrFail(t, s)

		var counter atomic.Int32
		targetTime := clock.Now().Add(chrono.Hour)

		for i := 0; i < 5; i++ {
			_, err := s.At(targetTime, scheduler.Task{
				Job: func(ctx context.Context) error {
					counter.Add(1)
					return nil
				},
				Ctx: context.Background(),
			})
			assertlite.NoError(t, err)
		}

		s.AdvanceTo(targetTime)
		assertlite.Equals(t, int32(5), counter.Load(), "all 5 task should execute")
	})
}

func TestScheduler_After(t *testing.T) {
	t.Run("executes task after duration", func(t *testing.T) {
		clock := chronos.VirtualClock()
		s := WithClock(clock)
		defer testlite.CloseOrFail(t, s)

		var executed bool
		startTime := clock.Now()

		_, err := s.After(chrono.Hour, scheduler.Task{
			Job: func(ctx context.Context) error {
				executed = true
				return nil
			},
			Ctx: context.Background(),
		})
		assertlite.NoError(t, err)

		// Advance partially
		s.AdvanceTo(startTime.Add(chrono.Minute * 30))
		assertlite.False(t, executed, "task should not execute before delay")

		// Advance to completion
		s.AdvanceTo(startTime.Add(chrono.Hour))
		assertlite.True(t, executed, "task should execute after delay")
	})

	t.Run("handles zero duration", func(t *testing.T) {
		clock := chronos.VirtualClock()
		s := WithClock(clock)
		defer testlite.CloseOrFail(t, s)

		var executed bool
		startTime := clock.Now()

		_, err := s.After(0, scheduler.Task{
			Job: func(ctx context.Context) error {
				executed = true
				return nil
			},
			Ctx: context.Background(),
		})
		assertlite.NoError(t, err)

		// Advance by a tiny amount to trigger execution
		s.AdvanceTo(startTime.Add(chrono.Nanosecond))
		assertlite.True(t, executed, "zero duration task should execute")
	})
}

func TestScheduler_Every(t *testing.T) {
	t.Run("executes recurring task eventAt interval", func(t *testing.T) {
		clock := chronos.VirtualClock()
		s := WithClock(clock)
		defer testlite.CloseOrFail(t, s)

		var counter atomic.Int32
		startTime := clock.Now()
		interval := chrono.Hour

		_, err := s.Every(interval, scheduler.Task{
			Job: func(ctx context.Context) error {
				counter.Add(1)
				return nil
			},
			Ctx: context.Background(),
		})
		assertlite.NoError(t, err)

		// Advance through multiple intervals
		for i := 1; i <= 5; i++ {
			s.AdvanceTo(startTime.Add(interval * chrono.Duration(i)))
			assertlite.Equals(t, int32(i), counter.Load(), "should execute %d times", i)
		}
	})

	t.Run("normalizes non-positive intervals", func(t *testing.T) {
		clock := chronos.VirtualClock()
		s := WithClock(clock)
		defer testlite.CloseOrFail(t, s)

		var executed bool
		startTime := clock.Now()

		// Test with zero interval
		_, err := s.Every(0, scheduler.Task{
			Job: func(ctx context.Context) error {
				executed = true
				return nil
			},
			Ctx: context.Background(),
		})
		assertlite.NoError(t, err)

		// Should still execute (interval normalized to 1ns)
		s.AdvanceTo(startTime.Add(chrono.Nanosecond))
		assertlite.True(t, executed, "zero interval should normalize to 1ns")

		// Test with negative interval
		executed = false
		startTime = clock.Now()

		_, err = s.Every(-5, scheduler.Task{
			Job: func(ctx context.Context) error {
				executed = true
				return nil
			},
			Ctx: context.Background(),
		})
		assertlite.NoError(t, err)

		s.AdvanceTo(startTime.Add(chrono.Nanosecond))
		assertlite.True(t, executed, "negative interval should normalize to 1ns")
	})

	t.Run("cancels recurring task", func(t *testing.T) {
		clock := chronos.VirtualClock()
		s := WithClock(clock)
		defer testlite.CloseOrFail(t, s)

		var counter atomic.Int32
		startTime := clock.Now()
		interval := chrono.Hour

		cancellable, err := s.Every(interval, scheduler.Task{
			Job: func(ctx context.Context) error {
				counter.Add(1)
				return nil
			},
			Ctx: context.Background(),
		})
		assertlite.NoError(t, err)

		// Execute first interval
		s.AdvanceTo(startTime.Add(interval))
		assertlite.Equals(t, int32(1), counter.Load())

		// Cancel
		cancellable.Cancel()

		// Advance further - should not execute again
		s.AdvanceTo(startTime.Add(interval * 3))
		assertlite.Equals(t, int32(1), counter.Load(), "should not execute after cancellation")
	})

	t.Run("stops on context cancellation", func(t *testing.T) {
		clock := chronos.VirtualClock()
		s := WithClock(clock)
		defer testlite.CloseOrFail(t, s)

		var counter atomic.Int32
		startTime := clock.Now()
		interval := chrono.Hour
		ctx, cancel := context.WithCancel(context.Background())

		_, err := s.Every(interval, scheduler.Task{
			Job: func(ctx context.Context) error {
				counter.Add(1)
				return nil
			},
			Ctx: ctx,
		})
		assertlite.NoError(t, err)

		// Execute first interval
		s.AdvanceTo(startTime.Add(interval))
		assertlite.Equals(t, int32(1), counter.Load())

		// Cancel context
		cancel()

		// Advance further - should not execute again due to context cancellation
		s.AdvanceTo(startTime.Add(interval * 3))
		assertlite.Equals(t, int32(1), counter.Load(), "should not execute after context cancellation")
	})

	t.Run("handles task errors gracefully", func(t *testing.T) {
		clock := chronos.VirtualClock()
		s := WithClock(clock)
		defer testlite.CloseOrFail(t, s)

		var counter atomic.Int32
		startTime := clock.Now()
		interval := chrono.Hour

		_, err := s.Every(interval, scheduler.Task{
			Job: func(ctx context.Context) error {
				counter.Add(1)
				if counter.Load() == 2 {
					return errors.New("simulated error")
				}
				return nil
			},
			Ctx: context.Background(),
		})
		assertlite.NoError(t, err)

		// Should continue executing after error
		for i := 1; i <= 3; i++ {
			s.AdvanceTo(startTime.Add(interval * chrono.Duration(i)))
		}
		assertlite.Equals(t, int32(3), counter.Load(), "should continue after errors")
	})
}

func TestScheduler_EveryFrom(t *testing.T) {
	t.Run("executes from specific start time", func(t *testing.T) {
		clock := chronos.VirtualClock()
		s := WithClock(clock)
		defer testlite.CloseOrFail(t, s)

		var counter atomic.Int32
		startTime := clock.Now()
		firstTick := startTime.Add(chrono.Hour * 2)
		interval := chrono.Hour

		_, err := s.EveryFrom(firstTick, interval, scheduler.Task{
			Job: func(ctx context.Context) error {
				counter.Add(1)
				return nil
			},
			Ctx: context.Background(),
		})
		assertlite.NoError(t, err)

		// Advance to just before first tick
		s.AdvanceTo(firstTick.Add(-chrono.Nanosecond))
		assertlite.Equals(t, int32(0), counter.Load(), "should not execute before first tick")

		// Advance to first tick
		s.AdvanceTo(firstTick)
		assertlite.Equals(t, int32(1), counter.Load(), "should execute eventAt first tick")

		// Advance to next interval
		s.AdvanceTo(firstTick.Add(interval))
		assertlite.Equals(t, int32(2), counter.Load(), "should execute eventAt next interval")
	})

	t.Run("handles first tick in the past", func(t *testing.T) {
		clock := chronos.VirtualClock()
		s := WithClock(clock)
		defer testlite.CloseOrFail(t, s)

		var counter atomic.Int32
		startTime := clock.Now()
		pastTick := startTime.Add(chrono.Hours(-10))
		interval := chrono.Hour

		_, err := s.EveryFrom(pastTick, interval, scheduler.Task{
			Job: func(ctx context.Context) error {
				counter.Add(1)
				return nil
			},
			Ctx: context.Background(),
		})
		assertlite.NoError(t, err)

		// Should execute immediately (or eventAt next interval)
		s.AdvanceTo(startTime.Add(chrono.Nanosecond))
		execs := counter.Load()
		assertlite.Equals(t, int32(1), execs, "should have executed '1' task but got '%d'", execs)
	})
}

func TestScheduler_Close(t *testing.T) {
	t.Run("prevents new scheduling after close", func(t *testing.T) {
		clock := chronos.VirtualClock()
		s := WithClock(clock)

		err := s.Close()
		assertlite.NoError(t, err)
		assertlite.True(t, s.IsClosed())

		// Try to schedule new task
		_, err = s.At(clock.Now().Add(chrono.Hour), scheduler.Task{
			Job: func(ctx context.Context) error { return nil },
		})
		assertlite.IsErrorOf(t, err, scheduler.ErrClosedScheduler)

		_, err = s.After(chrono.Hour, scheduler.Task{
			Job: func(ctx context.Context) error { return nil },
		})
		assertlite.IsErrorOf(t, err, scheduler.ErrClosedScheduler)

		_, err = s.Every(chrono.Hour, scheduler.Task{
			Job: func(ctx context.Context) error { return nil },
		})
		assertlite.IsErrorOf(t, err, scheduler.ErrClosedScheduler)
	})

	t.Run("queued task still execute after close", func(t *testing.T) {
		clock := chronos.VirtualClock()
		s := WithClock(clock)

		var executed bool
		targetTime := clock.Now().Add(chrono.Hour)

		_, err := s.At(targetTime, scheduler.Task{
			Job: func(ctx context.Context) error {
				executed = true
				return nil
			},
			Ctx: context.Background(),
		})
		assertlite.NoError(t, err)

		// Close scheduler
		assertlite.NoError(t, s.Close())

		// Advance to target time
		s.AdvanceTo(targetTime)
		assertlite.True(t, executed, "queued task should still execute after close")
	})
}

func TestScheduler_AdvanceTo(t *testing.T) {
	t.Run("executes multiple task in order", func(t *testing.T) {
		clock := chronos.VirtualClock()
		s := WithClock(clock)
		defer testlite.CloseOrFail(t, s)

		var order []string
		startTime := clock.Now()

		// Schedule task eventAt different times
		_, _ = s.At(startTime.Add(chrono.Hour), scheduler.Task{
			Job: func(ctx context.Context) error {
				order = append(order, "hour")
				return nil
			},
			Ctx: context.Background(),
		})

		_, _ = s.At(startTime.Add(chrono.Minute*30), scheduler.Task{
			Job: func(ctx context.Context) error {
				order = append(order, "30min")
				return nil
			},
			Ctx: context.Background(),
		})

		_, _ = s.At(startTime.Add(chrono.Minute*15), scheduler.Task{
			Job: func(ctx context.Context) error {
				order = append(order, "15min")
				return nil
			},
			Ctx: context.Background(),
		})

		// Advance to execute all
		s.AdvanceTo(startTime.Add(chrono.Hour))

		assertlite.Equals(t, []string{"15min", "30min", "hour"}, order,
			"task should execute in chronological order")
	})

	t.Run("idempotent advancement", func(t *testing.T) {
		clock := chronos.VirtualClock()
		s := WithClock(clock)
		defer testlite.CloseOrFail(t, s)

		var counter atomic.Int32
		targetTime := clock.Now().Add(chrono.Hour)

		_, err := s.At(targetTime, scheduler.Task{
			Job: func(ctx context.Context) error {
				counter.Add(1)
				return nil
			},
			Ctx: context.Background(),
		})
		assertlite.NoError(t, err)

		// Advance multiple times to same target
		s.AdvanceTo(targetTime)
		s.AdvanceTo(targetTime)
		s.AdvanceTo(targetTime)

		assertlite.Equals(t, int32(1), counter.Load(),
			"task should execute only once regardless of multiple advances")
	})
}

func TestScheduleNow(t *testing.T) {
	t.Run("schedules task for immediate execution", func(t *testing.T) {
		clock := chronos.VirtualClock()
		s := WithClock(clock)
		defer testlite.CloseOrFail(t, s)

		var executed bool
		startTime := clock.Now()

		err := ScheduleNow(s, scheduler.Task{
			Job: func(ctx context.Context) error {
				executed = true
				return nil
			},
			Ctx: context.Background(),
		})
		assertlite.NoError(t, err)

		// Advance by a tiny amount to trigger
		s.AdvanceTo(startTime.Add(chrono.Nanosecond))
		assertlite.True(t, executed, "task should execute immediately")
	})

	t.Run("returns error when scheduler closed", func(t *testing.T) {
		clock := chronos.VirtualClock()
		s := WithClock(clock)
		s.Close()

		err := ScheduleNow(s, scheduler.Task{
			Job: func(ctx context.Context) error { return nil },
		})
		assertlite.IsErrorOf(t, err, scheduler.ErrClosedScheduler)
	})
}

func TestCancellableFunc(t *testing.T) {
	t.Run("cancel is idempotent", func(t *testing.T) {
		clock := chronos.VirtualClock()
		s := WithClock(clock)
		defer testlite.CloseOrFail(t, s)

		var executed bool
		cancellable, err := s.After(chrono.Hour, scheduler.Task{
			Job: func(ctx context.Context) error {
				executed = true
				return nil
			},
			Ctx: context.Background(),
		})
		assertlite.NoError(t, err)

		// Cancel multiple times
		cancellable.Cancel()
		cancellable.Cancel()
		cancellable.Cancel()

		s.AdvanceTo(clock.Now().Add(chrono.Hour))
		assertlite.False(t, executed, "task should not execute after cancel")
	})

	t.Run("cancel before schedule", func(t *testing.T) {
		clock := chronos.VirtualClock()
		s := WithClock(clock)
		defer testlite.CloseOrFail(t, s)

		var executed bool
		targetTime := clock.Now().Add(chrono.Hour)

		_, err := s.At(targetTime, scheduler.Task{
			Job: func(ctx context.Context) error {
				executed = true
				return nil
			},
			Ctx: context.Background(),
		})
		assertlite.NoError(t, err)

		// Advance past target without explicit cancel
		s.AdvanceTo(targetTime)
		assertlite.True(t, executed, "task should execute normally")
	})
}

func TestConcurrency(t *testing.T) {
	t.Run("concurrent scheduling and advancement", func(t *testing.T) {
		clock := chronos.VirtualClock()
		s := WithClock(clock)
		defer testlite.CloseOrFail(t, s)

		var counter atomic.Int32
		const numTasks = 100

		// Schedule task concurrently
		var wg sync.WaitGroup
		wg.Add(numTasks)

		for i := range numTasks {
			go func(delay chrono.Duration) {
				defer wg.Done()
				_, err := s.After(delay, scheduler.Task{
					Job: func(ctx context.Context) error {
						counter.Add(1)
						return nil
					},
					Ctx: context.Background(),
				})
				assertlite.NoError(t, err)
			}(durations.Ms(i))
		}

		wg.Wait()

		// Advance to execute all
		s.AdvanceTo(clock.Now().Add(chrono.Hour))
		assertlite.Equals(t, int32(numTasks), counter.Load(),
			"expected '%d' task to be executed, got '%d'", numTasks, counter.Load())
	})
}
