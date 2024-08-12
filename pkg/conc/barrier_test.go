// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package conc

import (
	"github.com/andrerrcosta2/gtools/pkg/gtools"
	"math/rand"
	"sync"
	"testing"
)

func OnThreshold(b gtools.Barrier) bool {
	return b.Count() >= 5
}

func OnOddCount(b gtools.Barrier) bool {
	return b.Count()%2 != 0
}

func TestStepBroadcast_Basic(t *testing.T) {
	reg := NewRegisters(30)
	barrier := StepBroadcast(OnThreshold)
	var wg sync.WaitGroup
	wg.Add(reg.max)

	for i := 0; i < reg.max; i++ {
		go func(id int) {
			defer wg.Done()

			goroutine := NewGoroutine(id, barrier.Count(), 300)
			goroutine.Work()

			barrier.Wait()
			reg.Worker(goroutine)
			goroutine.Work()

			barrier.Wait()
			reg.Worker(goroutine)
			reg.Done(goroutine)

		}(i)
	}

	wg.Wait()

	for _, routine := range reg.finished {
		t.Logf("%s\n", routine.state)
	}

	if len(reg.finished) != reg.max {
		t.Errorf("Expected %d finished goroutines, got %d", reg.max, len(reg.finished))
	}

	if len(reg.working) != 0 {
		t.Errorf("Expected 0 working goroutines, got %d", len(reg.working))
	}
}

func TestStepBroadcast_HighConcurrency(t *testing.T) {
	reg := NewRegisters(2000)
	barrier := StepBroadcast(OnThreshold)
	var wg sync.WaitGroup
	wg.Add(reg.max)

	for i := 0; i < reg.max; i++ {
		go func(id int) {
			defer wg.Done()
			r := rand.Intn(100)
			goroutine := NewGoroutine(id, barrier.Count(), r)
			goroutine.Work()

			barrier.Wait()
			reg.Worker(goroutine)
			goroutine.Work()

			barrier.Wait()
			reg.Worker(goroutine)
			reg.Done(goroutine)

		}(i)
	}

	wg.Wait()

	if len(reg.finished) != reg.max {
		t.Errorf("Expected %d finished goroutines, got %d", reg.max, len(reg.finished))
	}

	if len(reg.working) != 0 {
		t.Errorf("Expected 0 working goroutines, got %d", len(reg.working))
	}
}

func TestStepBroadcast_OddHighConcurrency(t *testing.T) {
	reg := NewRegisters(2000)
	barrier := StepBroadcast(OnOddCount)
	var wg sync.WaitGroup
	wg.Add(reg.max)

	for i := 0; i < reg.max; i++ {
		go func(id int) {
			defer wg.Done()
			r := rand.Intn(100)
			goroutine := NewGoroutine(id, barrier.Count(), r)
			goroutine.Work()

			barrier.Wait()
			reg.Worker(goroutine)
			goroutine.Work()

			barrier.Wait()
			reg.Worker(goroutine)
			reg.Done(goroutine)

		}(i)
	}

	wg.Wait()

	if len(reg.finished) != reg.max {
		t.Errorf("Expected %d finished goroutines, got %d", reg.max, len(reg.finished))
	}

	if len(reg.working) != 0 {
		t.Errorf("Expected 0 working goroutines, got %d", len(reg.working))
	}
}

func TestStepBroadcast_EdgeThreshold(t *testing.T) {
	reg := NewRegisters(2000)
	barrier := StepBroadcast(func(b gtools.Barrier) bool {
		return b.Count() == 1
	})
	var wg sync.WaitGroup
	wg.Add(reg.max)

	for i := 0; i < reg.max; i++ {
		go func(id int) {
			defer wg.Done()
			r := rand.Intn(100)
			goroutine := NewGoroutine(id, barrier.Count(), r)
			goroutine.Work()

			barrier.Wait()
			reg.Worker(goroutine)
			goroutine.Work()

			barrier.Wait()
			reg.Worker(goroutine)
			reg.Done(goroutine)

		}(i)
	}

	wg.Wait()

	if len(reg.finished) != reg.max {
		t.Errorf("Expected %d finished goroutines, got %d", reg.max, len(reg.finished))
	}

	if len(reg.working) != 0 {
		t.Errorf("Expected 0 working goroutines, got %d", len(reg.working))
	}
}

func TestThresholdBarrier_Basic(t *testing.T) {
	reg := NewRegisters(12)
	barrier := ThresholdBarrier(3)
	var wg sync.WaitGroup
	wg.Add(reg.max)

	for i := 0; i < reg.max; i++ {
		go func(id int) {
			defer wg.Done()
			r := rand.Intn(100)
			goroutine := NewGoroutine(id, barrier.Count(), r)
			goroutine.Work()

			barrier.Wait()
			reg.Worker(goroutine)
			goroutine.Work()

			barrier.Wait()
			reg.Worker(goroutine)
			reg.Done(goroutine)

		}(i)
	}

	wg.Wait()

	if len(reg.finished) != reg.max {
		t.Errorf("Expected %d finished goroutines, got %d", reg.max, len(reg.finished))
	}

	if len(reg.working) != 0 {
		t.Errorf("Expected 0 working goroutines, got %d", len(reg.working))
	}
}

func TestThresholdBarrier_HighConcurrency(t *testing.T) {
	reg := NewRegisters(2000)
	barrier := ThresholdBarrier(200)
	var wg sync.WaitGroup
	wg.Add(reg.max)

	for i := 0; i < reg.max; i++ {
		go func(id int) {
			defer wg.Done()
			r := rand.Intn(100)
			goroutine := NewGoroutine(id, barrier.Count(), r)
			goroutine.Work()

			barrier.Wait()
			reg.Worker(goroutine)
			goroutine.Work()

			barrier.Wait()
			reg.Worker(goroutine)
			reg.Done(goroutine)

		}(i)
	}

	wg.Wait()

	if len(reg.finished) != reg.max {
		t.Errorf("Expected %d finished goroutines, got %d", reg.max, len(reg.finished))
	}

	if len(reg.working) != 0 {
		t.Errorf("Expected 0 working goroutines, got %d", len(reg.working))
	}
}
