package executors

import (
	"sync/atomic"
	"testing"
	"time"

	"github.com/andrerrcosta2/gtools/conc/core/executor"
)

func TestPool_SubmitAndWait(t *testing.T) {
	var count int64
	p := NewPool[int](2, 16, func(v int) {
		atomic.AddInt64(&count, int64(v))
	})

	for i := 0; i < 5; i++ {
		if err := p.Submit(1); err != nil {
			t.Fatalf("unexpected submit error: %v", err)
		}
	}

	if err := p.Close(); err != nil {
		t.Fatalf("unexpected close error: %v", err)
	}
	p.Wait()

	if got := atomic.LoadInt64(&count); got != 5 {
		t.Fatalf("expected 5 processed items, got %d", got)
	}
}

func TestPool_SubmitAfterCloseFails(t *testing.T) {
	p := NewPool[int](1, 1, func(v int) {})
	if err := p.Close(); err != nil {
		t.Fatalf("unexpected close error: %v", err)
	}
	if err := p.Submit(1); err != executor.ErrExecutorClosed {
		t.Fatalf("expected ErrExecutorClosed, got %v", err)
	}
}

func TestPool_CloseIsIdempotentError(t *testing.T) {
	p := NewPool[int](1, 1, func(v int) { time.Sleep(time.Millisecond) })
	if err := p.Close(); err != nil {
		t.Fatalf("unexpected close error: %v", err)
	}
	if err := p.Close(); err != executor.ErrExecutorClosed {
		t.Fatalf("expected ErrExecutorClosed, got %v", err)
	}
	p.Wait()
}
