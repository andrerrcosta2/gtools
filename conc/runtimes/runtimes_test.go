package runtimes

import (
	"context"
	"testing"
	"time"

	"github.com/andrerrcosta2/gtools/core/chronos"
	"github.com/andrerrcosta2/gtools/core/domain/contracts/chrono/durations"
	"github.com/andrerrcosta2/gtools/core/domain/contracts/timing/scheduler"
	"github.com/andrerrcosta2/gtools/core/testlite/assertlite"
)

func TestVirtual_AdvanceShouldDriveSchedulerAndTimers(t *testing.T) {
	rt := Virtual(chronos.VirtualClock())
	defer func() { _ = rt.Close() }()

	var order []string

	sched := rt.Scheduler()

	_, err := sched.After(durations.Ms(100), scheduler.Task{
		Ctx: context.Background(),
		Job: func(context.Context) error {
			order = append(order, "scheduler")
			return nil
		},
	})

	assertlite.NoError(t, err)

	timer := rt.TimerFactory().NewTimer(func() {
		order = append(order, "timer")
	})
	timer.Reset(durations.Ms(10))

	rt.Advance(durations.Ms(9))
	if len(order) != 0 {
		t.Fatalf("expected no executions before due time, got %v", order)
	}

	rt.Advance(durations.Ms(1))
	if len(order) != 2 {
		t.Fatalf("expected both executions at due time, got %v", order)
	}
	if order[0] != "timer" || order[1] != "scheduler" {
		t.Fatalf("unexpected execution order: %v", order)
	}
}

func TestSystem_CloseStopsDriver(t *testing.T) {
	env := System()

	done := make(chan struct{})
	go func() {
		_ = env.Close()
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(500 * time.Millisecond):
		t.Fatal("env close timed out")
	}
}
