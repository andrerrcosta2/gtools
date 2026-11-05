// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package streams

import (
	"context"
	"fmt"
	"sync/atomic"
	"testing"

	"github.com/andrerrcosta2/gtools/core/domain/contracts/chrono/durations"
	"github.com/andrerrcosta2/gtools/core/domain/function/runnables"
)

func TestBuffered_BlockedLoop(t *testing.T) {
	producer := Producer[string](5)

	runnables.Before(func() {
		err := producer.Produce(mediumSupply...)

		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}

		fmt.Printf("producer on loop\n")
	}, durations.Ms(150).ToStd())

	err := producer.Close()

	if err != nil {
		t.Errorf("expected no error, got %s\n", err)
	}
}

func TestBuffered_SupplySameSize(t *testing.T) {
	var counter atomic.Int32
	producer := Producer[string](len(smallSupply))
	err := producer.Produce(smallSupply...)
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}

	// consume blocking the main thread

	Range(producer.Stream(), func(item string) bool {
		counter.Add(1)
		fmt.Printf("value: %s\n", item)
		return true
	})

	if counter.Load() != int32(len(smallSupply)) {
		t.Errorf("expected 5 values, got %d", counter.Load())
	}
}

func TestBuffered_SupplyAfterClose(t *testing.T) {
	supplier := Producer[string](5)
	err := supplier.Close()
	if err != nil {
		t.Errorf("expected no error, got %s\n", err)
	}

	err = supplier.Produce(smallSupply...)
	if err == nil {
		t.Errorf("expected supplier to be closed, got no error\n")
	}
}

func TestCancellable_SupplyAfterCancel(t *testing.T) {
	background, cancel := context.WithCancel(context.Background())
	producer := CancellableProducer[string](background, 5)
	err := producer.Produce(smallSupply...)
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}

	cancel()

	// Shouldn't supply anymore
	err = producer.Produce("1", "2", "3", "4", "5")
	if err == nil {
		t.Errorf("expected producer to be closed, got no error\n")
	}
}
