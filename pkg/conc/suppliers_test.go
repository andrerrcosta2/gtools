// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package conc

import (
	"context"
	"fmt"
	"github.com/andrerrcosta2/gtools/core/durations"
	"github.com/andrerrcosta2/gtools/core/functions/runnables"
	"sync/atomic"
	"testing"
)

func TestBufferedSupplier_BlockedLoop(t *testing.T) {
	supplier := BufferedSupplier[string](5)

	runnables.Before(func() {
		err := supplier.Supply(mediumSupply...)

		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}

		fmt.Printf("supplier on loop\n")
	}, durations.Ms(150))

	err := supplier.Close()

	if err != nil {
		t.Errorf("expected no error, got %s\n", err)
	}
}

func TestBufferedSupplier_SupplySameSize(t *testing.T) {
	var counter atomic.Int32
	supplier := BufferedSupplier[string](len(smallSupply))
	err := supplier.Supply(smallSupply...)
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}

	// consume blocking the main thread

	stream := supplier.Stream()
	for {
		value, ok := <-stream
		if !ok {
			break
		}

		counter.Add(1)
		fmt.Printf("value: %s\n", value)
	}

	if counter.Load() != int32(len(smallSupply)) {
		t.Errorf("expected 5 values, got %d", counter.Load())
	}
}

func TestBufferedSupplier_SupplyAfterClose(t *testing.T) {
	supplier := BufferedSupplier[string](5)
	err := supplier.Close()
	if err != nil {
		t.Errorf("expected no error, got %s\n", err)
	}

	err = supplier.Supply(smallSupply...)
	if err == nil {
		t.Errorf("expected supplier to be closed, got no error\n")
	}
}

func TestCancellableSupplier_SupplyAfterClose(t *testing.T) {
	supplier := CancellableSupplier[string](context.Background(), 5)
	err := supplier.Supply(smallSupply...)
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}

	err = supplier.Close()
	if err != nil {
		t.Errorf("expected no error, got %s\n", err)
	}

	// Shouldn't supply anymore
	err = supplier.Supply("1", "2", "3", "4", "5")
	if err == nil {
		t.Errorf("expected supplier to be closed, got no error\n")
	}
}
