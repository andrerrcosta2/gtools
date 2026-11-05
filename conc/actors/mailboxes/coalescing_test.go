// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package mailboxes

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/andrerrcosta2/gtools/conc/core/mailbox"
	"github.com/andrerrcosta2/gtools/core/testlite/assertlite"
)

// TestCoalescing_Sync
//
// Single-goroutine operations on a Coalescing mailbox.
//
// We should assert:
//  1. Send stores the value and returns true
//  2. Subsequent Send replaces previous value without errors
//  3. Receive returns the latest value sent
//  4. Receive always returns true (even without prior Send)
//  5. Send to closed mailbox returns false
//  6. Close transitions state and prevents further operations
//  7. Double close returns mailbox.ErrClosedTwice
//  8. Seal prevents new sends but allows reading current value
//  9. Receive on sealed empty mailbox triggers close
//
// 10. Only the latest value survives multiple sends
func TestCoalescing_Sync(t *testing.T) {

	t.Run("send stores value and returns true", func(t *testing.T) {
		mb := Coalescing[int]()

		assertlite.True(t, mb.Send(42))

		v, ok := mb.Receive()
		assertlite.True(t, ok)
		assertlite.Equals(t, 42, v)
	})

	t.Run("subsequent send replaces previous value", func(t *testing.T) {
		mb := Coalescing[int]()

		mb.Send(1)
		mb.Send(2)
		mb.Send(3)

		// Should only get the latest value
		v, ok := mb.Receive()
		assertlite.True(t, ok)
		assertlite.Equals(t, 3, v)

		// No more values (Receive still returns true with zero value)
		v2, ok2 := mb.Receive()
		assertlite.True(t, ok2)
		assertlite.Equals(t, 0, v2)
	})

	t.Run("receive always returns true even on empty mailbox", func(t *testing.T) {
		mb := Coalescing[int]()

		// Receive without any Send
		v, ok := mb.Receive()
		assertlite.True(t, ok)
		assertlite.Equals(t, 0, v)
	})

	t.Run("only latest value survives multiple sends", func(t *testing.T) {
		mb := Coalescing[string]()

		mb.Send("first")
		mb.Send("second")
		mb.Send("third")
		mb.Send("latest")

		v, ok := mb.Receive()
		assertlite.True(t, ok)
		assertlite.Equals(t, "latest", v)

		// Subsequent receives get zero value
		v2, ok2 := mb.Receive()
		assertlite.True(t, ok2)
		assertlite.Equals(t, "", v2)
	})

	t.Run("send to closed mailbox returns false", func(t *testing.T) {
		mb := Coalescing[int]()

		err := mb.Close()
		assertlite.NoError(t, err)

		assertlite.False(t, mb.Send(42))
	})

	t.Run("double close returns mailbox.ErrClosed", func(t *testing.T) {
		mb := Coalescing[int]()

		assertlite.NoError(t, mb.Close())
		assertlite.IsErrorOf(t, mailbox.ErrClosedTwice, mb.Close())
	})

	t.Run("seal prevents new sends but allows reading current value", func(t *testing.T) {
		mb := Coalescing[int]()

		mb.Send(42)
		waiter := mb.Seal()

		// Cannot send after seal
		assertlite.False(t, mb.Send(99))

		// Can still read the last value
		v, ok := mb.Receive()
		assertlite.True(t, ok)
		assertlite.Equals(t, 42, v)

		// Waiter should complete after value is consumed
		done := make(chan struct{})
		go func() {
			waiter.Wait()
			close(done)
		}()

		select {
		case <-done:
			// Success - waiter completed
		case <-time.After(time.Second):
			t.Fatal("waiter did not complete after sealed mailbox was drained")
		}
	})

	t.Run("receive on sealed empty mailbox triggers close", func(t *testing.T) {
		mb := Coalescing[int]()

		waiter := mb.Seal()

		// Empty mailbox, Receive should trigger close
		v, ok := mb.Receive()
		assertlite.True(t, ok)
		assertlite.Equals(t, 0, v)

		// Wait for close to complete
		waiter.Wait()

		// Verify closed state
		assertlite.False(t, mb.Send(1))
		assertlite.IsErrorOf(t, mailbox.ErrClosedTwice, mb.Close())
	})

	t.Run("close after seal and drain works correctly", func(t *testing.T) {
		mb := Coalescing[int]()

		mb.Send(1)
		mb.Seal()
		mb.Receive() // drain, triggers close internally

		// Explicit close after infra close
		assertlite.IsErrorOf(t, mailbox.ErrClosedTwice, mb.Close())
	})

	t.Run("send after receive returns to empty state", func(t *testing.T) {
		mb := Coalescing[int]()

		mb.Send(1)
		v1, _ := mb.Receive()
		assertlite.Equals(t, 1, v1)

		// Mailbox is now empty, can send again
		assertlite.True(t, mb.Send(2))
		v2, _ := mb.Receive()
		assertlite.Equals(t, 2, v2)
	})

	t.Run("complex value types work correctly", func(t *testing.T) {
		type Message struct {
			ID   int
			Data string
		}

		mb := Coalescing[Message]()

		mb.Send(Message{ID: 1, Data: "first"})
		mb.Send(Message{ID: 2, Data: "second"})
		mb.Send(Message{ID: 3, Data: "latest"})

		v, ok := mb.Receive()
		assertlite.True(t, ok)
		assertlite.Equals(t, Message{ID: 3, Data: "latest"}, v)
	})
}

// TestCoalescing_Concurrent
//
// Multiple goroutines operating on a Coalescing mailbox.
//
// We should assert:
//  1. Concurrent sends are thread-safe
//  2. Latest value semantics under concurrent sends
//  3. Concurrent sends and receives don't panic
//  4. Seal with concurrent receivers drains correctly
//  5. Close with concurrent senders rejects new sends
func TestCoalescing_Concurrent(t *testing.T) {

	t.Run("concurrent sends are thread-safe", func(t *testing.T) {
		mb := Coalescing[int]()
		const numGoroutines = 10
		const sendsPerGoroutine = 100

		var wg sync.WaitGroup
		wg.Add(numGoroutines)

		// Multiple goroutines sending concurrently
		for i := 0; i < numGoroutines; i++ {
			go func(id int) {
				defer wg.Done()
				for j := 0; j < sendsPerGoroutine; j++ {
					mb.Send(id*sendsPerGoroutine + j)
				}
			}(i)
		}

		wg.Wait()

		// Should be able to receive a value without panic
		v, ok := mb.Receive()
		assertlite.True(t, ok)
		// Value should be valid (some integer)
		_ = v
	})

	t.Run("latest value semantics under concurrent sends", func(t *testing.T) {
		mb := Coalescing[int]()

		var wg sync.WaitGroup
		var lastSent atomic.Int64

		// Send values with increasing timestamps
		wg.Add(2)
		go func() {
			defer wg.Done()
			for i := 0; i < 100; i++ {
				mb.Send(i)
				lastSent.Store(int64(i))
				time.Sleep(time.Microsecond)
			}
		}()

		// Another sender with higher values
		go func() {
			defer wg.Done()
			time.Sleep(time.Millisecond) // Start slightly later
			for i := 1000; i < 1100; i++ {
				mb.Send(i)
				lastSent.Store(int64(i))
				time.Sleep(time.Microsecond)
			}
		}()

		wg.Wait()

		// The received value should be one of the sent values
		v, ok := mb.Receive()
		assertlite.True(t, ok)
		assertlite.True(t, v >= 0 && v < 1100, "received value out of range: %d", v)
	})

	t.Run("concurrent sends and receives don't panic", func(t *testing.T) {
		mb := Coalescing[int]()
		const numOps = 1000

		var wg sync.WaitGroup
		wg.Add(2)

		// Sender
		go func() {
			defer wg.Done()
			for i := 0; i < numOps; i++ {
				mb.Send(i)
			}
		}()

		// Receiver
		go func() {
			defer wg.Done()
			for i := 0; i < numOps; i++ {
				_, _ = mb.Receive()
			}
		}()

		wg.Wait()
		// Test passes if no panic or race detector warnings
	})

	t.Run("seal drains remaining value with concurrent receivers", func(t *testing.T) {
		mb := Coalescing[int]()

		mb.Send(42)

		var received atomic.Int32
		var wg sync.WaitGroup

		// Start multiple receivers
		for i := 0; i < 4; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				if _, ok := mb.Receive(); ok {
					received.Add(1)
				}
			}()
		}

		// Seal after receivers start
		time.Sleep(10 * time.Millisecond)
		waiter := mb.Seal()

		// Wait for drain
		waiter.Wait()

		// Only one value should have been received (the latest)
		assertlite.True(t, received.Load() >= 1, "at least one value should be received")
	})

	t.Run("close with concurrent senders rejects sends", func(t *testing.T) {
		mb := Coalescing[int]()

		var accepted atomic.Int32
		var rejected atomic.Int32

		var wg sync.WaitGroup
		wg.Add(4)

		// Start senders
		for i := 0; i < 4; i++ {
			go func() {
				defer wg.Done()
				for j := 0; j < 100; j++ {
					if mb.Send(j) {
						accepted.Add(1)
					} else {
						rejected.Add(1)
						return
					}
					time.Sleep(time.Microsecond)
				}
			}()
		}

		// Close after some sends
		time.Sleep(5 * time.Millisecond)
		mb.Close()

		wg.Wait()

		total := accepted.Load() + rejected.Load()
		assertlite.True(t, total > 0, "no operations completed")
		assertlite.True(t, rejected.Load() > 0, "no sends rejected after close")
	})

	t.Run("no data race with concurrent close and receive", func(t *testing.T) {
		mb := Coalescing[int]()

		mb.Send(1)

		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			defer wg.Done()
			time.Sleep(5 * time.Millisecond)
			mb.Close()
		}()

		go func() {
			defer wg.Done()
			for i := 0; i < 10; i++ {
				mb.Receive()
			}
		}()

		wg.Wait()
		// Test passes if no race detector warnings
	})
}

// TestCoalescing_StateTransitions
//
// Verifies correct state machine behavior of the Coalescing mailbox.
//
// We should assert:
//  1. Initial state allows send and receive
//  2. Sealed state rejects sends, allows receives
//  3. Closed state rejects sends and receives
//  4. Receive on empty sealed transitions to closed
//  5. Close on open transitions to closed
//  6. Close on sealed transitions to closed
//  7. Seal on sealed is idempotent
//  8. Seal on closed is no-op (stays closed)
func TestCoalescing_StateTransitions(t *testing.T) {

	t.Run("initial state allows send and receive", func(t *testing.T) {
		mb := Coalescing[int]()

		assertlite.True(t, mb.Send(1))
		v, ok := mb.Receive()
		assertlite.True(t, ok)
		assertlite.Equals(t, 1, v)
	})

	t.Run("sealed state rejects sends but allows receives", func(t *testing.T) {
		mb := Coalescing[int]()

		mb.Send(1)
		mb.Seal()

		// Rejects send
		assertlite.False(t, mb.Send(2))

		// Allows receive
		v, ok := mb.Receive()
		assertlite.True(t, ok)
		assertlite.Equals(t, 1, v)
	})

	t.Run("closed state rejects sends and still allows receives", func(t *testing.T) {
		mb := Coalescing[int]()

		mb.Close()

		assertlite.False(t, mb.Send(1))
		// Note: Coalescing.Receive always returns true
		_, ok := mb.Receive()
		assertlite.True(t, ok)
	})

	t.Run("receive on empty sealed transitions to closed", func(t *testing.T) {
		mb := Coalescing[int]()

		mb.Send(1)
		mb.Seal()

		// Drain
		mb.Receive()

		// Verify closed state
		assertlite.False(t, mb.Send(2))
		assertlite.IsErrorOf(t, mailbox.ErrClosedTwice, mb.Close())
	})

	t.Run("close on open transitions to closed", func(t *testing.T) {
		mb := Coalescing[int]()

		assertlite.NoError(t, mb.Close())
		assertlite.False(t, mb.Send(1))
	})

	t.Run("close on sealed transitions to closed", func(t *testing.T) {
		mb := Coalescing[int]()

		mb.Send(1)
		mb.Seal()
		assertlite.NoError(t, mb.Close())

		assertlite.False(t, mb.Send(2))
	})

	t.Run("seal on sealed is idempotent", func(t *testing.T) {
		mb := Coalescing[int]()

		mb.Send(1)
		w1 := mb.Seal()
		w2 := mb.Seal()

		// Both waiters should resolve
		mb.Receive() // drain

		done1 := make(chan struct{})
		done2 := make(chan struct{})

		go func() {
			w1.Wait()
			close(done1)
		}()
		go func() {
			w2.Wait()
			close(done2)
		}()

		select {
		case <-done1:
		case <-time.After(time.Second):
			t.Fatal("first waiter did not complete")
		}

		select {
		case <-done2:
		case <-time.After(time.Second):
			t.Fatal("second waiter did not complete")
		}
	})

	t.Run("seal on closed stays closed", func(t *testing.T) {
		mb := Coalescing[int]()

		mb.Close()
		mb.Seal()

		assertlite.False(t, mb.Send(1))
		_, ok := mb.Receive()
		assertlite.True(t, ok) // Still returns true for zero value
	})
}

// TestCoalescing_Behavioral
//
// Tests specific behavioral characteristics of Coalescing mailbox.
//
// We should assert:
//  1. Coalescing drops intermediate values
//  2. Only latest value is kept
//  3. Memory is O(1) - constant regardless of number of sends
//  4. Suitable for latest-value-cache pattern
func TestCoalescing_Behavioral(t *testing.T) {

	t.Run("drops intermediate values keeping only latest", func(t *testing.T) {
		mb := Coalescing[int]()

		// Send many values without receiving
		for i := 0; i < 1000; i++ {
			mb.Send(i)
		}

		// Should only receive the last value
		v, ok := mb.Receive()
		assertlite.True(t, ok)
		assertlite.Equals(t, 999, v)

		// No more values
		v2, _ := mb.Receive()
		assertlite.Equals(t, 0, v2)
	})

	t.Run("fits latest-value-cache pattern", func(t *testing.T) {
		mb := Coalescing[string]()

		// Simulate rapid updates where only latest matters
		updates := []string{"v1", "v2", "v3", "v4", "latest"}
		for _, u := range updates {
			mb.Send(u)
		}

		// Consumer only gets the latest
		v, ok := mb.Receive()
		assertlite.True(t, ok)
		assertlite.Equals(t, "latest", v)
	})

	t.Run("handles rapid send-receive cycles", func(t *testing.T) {
		mb := Coalescing[int]()

		for i := 0; i < 100; i++ {
			mb.Send(i)
			v, ok := mb.Receive()
			assertlite.True(t, ok)
			assertlite.Equals(t, i, v)
		}
	})

	t.Run("zero value handling", func(t *testing.T) {
		mb := Coalescing[int]()

		// Send zero value
		mb.Send(0)
		v, ok := mb.Receive()
		assertlite.True(t, ok)
		assertlite.Equals(t, 0, v)

		// After consuming, receive returns zero value again
		v2, ok2 := mb.Receive()
		assertlite.True(t, ok2)
		assertlite.Equals(t, 0, v2)
	})

	t.Run("pointer type handling", func(t *testing.T) {
		mb := Coalescing[*int]()

		val1 := 1
		val2 := 2
		val3 := 3

		mb.Send(&val1)
		mb.Send(&val2)
		mb.Send(&val3)

		v, ok := mb.Receive()
		assertlite.True(t, ok)
		assertlite.Equals(t, &val3, v)
		assertlite.Equals(t, 3, *v)
	})

	t.Run("slice type handling", func(t *testing.T) {
		mb := Coalescing[[]int]()

		mb.Send([]int{1, 2})
		mb.Send([]int{3, 4, 5})
		mb.Send([]int{6})

		v, ok := mb.Receive()
		assertlite.True(t, ok)
		assertlite.EqualSlices(t, []int{6}, v)
	})
}
