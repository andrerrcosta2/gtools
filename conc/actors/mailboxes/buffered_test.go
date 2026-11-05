// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package mailboxes

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/andrerrcosta2/gtools/conc/core/mailbox"
	"github.com/andrerrcosta2/gtools/core/testlite/assertlite"
	"github.com/andrerrcosta2/gtools/datastr/testing/recorders"
)

// TestBuffered_Sync
//
// Single-goroutine operations on a Buffered mailbox.
//
// We should assert:
//  1. Send writes to the underlying buffer and returns true
//  2. Receive reads from the underlying buffer in order
//  3. Receive on empty buffer returns zero value and false
//  4. Send to closed mailbox returns false without writing to buffer
//  5. Receive from closed mailbox returns zero value and false
//  6. Close transitions state and prevents further operations
//  7. Double close returns mailbox.ErrClosedTwice
//  8. Seal prevents new sends but allows draining existing elements
//  9. Receive on sealed empty buffer triggers close
//
// 10. Len and Cap delegate to underlying buffer
func TestBuffered_Sync(t *testing.T) {

	t.Run("send writes to buffer and returns true", func(t *testing.T) {
		buf := recorders.NewBuffer[int](4)
		mb := Buffered[int](buf)

		assertlite.True(t, mb.Send(42))
		assertlite.Equals(t, 1, buf.WriteCalls)
		assertlite.EqualSlices(t, []int{42}, buf.Writes)
	})

	t.Run("receive reads from buffer in FIFO order", func(t *testing.T) {
		buf := recorders.NewBuffer[int](4)
		mb := Buffered[int](buf)

		mb.Send(1)
		mb.Send(2)
		mb.Send(3)

		v1, ok1 := mb.Receive()
		assertlite.True(t, ok1)
		assertlite.Equals(t, 1, v1)

		v2, ok2 := mb.Receive()
		assertlite.True(t, ok2)
		assertlite.Equals(t, 2, v2)

		v3, ok3 := mb.Receive()
		assertlite.True(t, ok3)
		assertlite.Equals(t, 3, v3)

		assertlite.Equals(t, 3, buf.ReadCalls)
	})

	t.Run("receive on empty buffer returns zero and false", func(t *testing.T) {
		buf := recorders.NewBuffer[int](4)
		mb := Buffered[int](buf)

		v, ok := mb.Receive()
		assertlite.False(t, ok)
		assertlite.Equals(t, 0, v)
		assertlite.Equals(t, 1, buf.ReadCalls)
	})

	t.Run("send to closed mailbox returns false without writing", func(t *testing.T) {
		buf := recorders.NewBuffer[int](4)
		mb := Buffered[int](buf)

		err := mb.Close()
		assertlite.NoError(t, err)

		writeCallsBefore := buf.WriteCalls
		assertlite.False(t, mb.Send(42))
		assertlite.Equals(t, writeCallsBefore, buf.WriteCalls,
			"buffer.Write must not be called after close")
	})

	t.Run("receive from closed empty mailbox returns zero and false", func(t *testing.T) {
		buf := recorders.NewBuffer[int](4)
		mb := Buffered[int](buf)

		mb.Close()

		v, ok := mb.Receive()
		assertlite.False(t, ok)
		assertlite.Equals(t, 0, v)
	})

	t.Run("double close returns mailbox.ErrClosed", func(t *testing.T) {
		buf := recorders.NewBuffer[int](4)
		mb := Buffered[int](buf)

		assertlite.NoError(t, mb.Close())
		assertlite.IsErrorOf(t, mailbox.ErrClosedTwice, mb.Close())
	})

	t.Run("seal prevents new sends but allows draining", func(t *testing.T) {
		buf := recorders.NewBuffer[int](4)
		mb := Buffered[int](buf)

		mb.Send(10)
		mb.Send(20)

		waiter := mb.Shutdown()

		// Cannot send after seal
		writeCallsBefore := buf.WriteCalls
		assertlite.False(t, mb.Send(30))
		assertlite.Equals(t, writeCallsBefore, buf.WriteCalls,
			"buffer.Write must not be called after seal")

		// Can still drain existing elements
		v1, ok1 := mb.Receive()
		assertlite.True(t, ok1)
		assertlite.Equals(t, 10, v1)

		v2, ok2 := mb.Receive()
		assertlite.True(t, ok2)
		assertlite.Equals(t, 20, v2)

		// Drain completes, latch is signaled
		_, ok3 := mb.Receive()
		assertlite.False(t, ok3)

		// Waiter should be signaled after drain triggers close
		done := make(chan struct{})
		go func() {
			waiter.Wait()
			close(done)
		}()

		select {
		case <-done:
			// Success
		case <-time.After(time.Second):
			t.Fatal("waiter did not complete after sealed buffer drained")
		}
	})

	t.Run("receive on sealed empty buffer triggers close", func(t *testing.T) {
		buf := recorders.NewBuffer[int](4)
		mb := Buffered[int](buf)

		mb.Shutdown()

		// Buffer is empty, Receive should trigger close
		_, ok := mb.Receive()
		assertlite.False(t, ok)

		// Now closed, double close should fail
		assertlite.IsErrorOf(t, mailbox.ErrClosedTwice, mb.Close())
	})

	t.Run("len delegates to underlying buffer", func(t *testing.T) {
		buf := recorders.NewBuffer[int](4)
		mb := Buffered[int](buf)

		assertlite.Equals(t, 0, mb.Len())

		mb.Send(1)
		assertlite.Equals(t, 1, mb.Len())

		mb.Send(2)
		assertlite.Equals(t, 2, mb.Len())

		mb.Receive()
		assertlite.Equals(t, 1, mb.Len())
	})

	t.Run("cap delegates to underlying buffer", func(t *testing.T) {
		buf := recorders.NewBuffer[int](8)
		mb := Buffered[int](buf)

		assertlite.Equals(t, 8, mb.Cap())
	})

	t.Run("close after seal and drain works correctly", func(t *testing.T) {
		buf := recorders.NewBuffer[int](4)
		mb := Buffered(buf)

		start := time.Now()
		mb.Send(1)
		t.Logf("Send: %s", time.Since(start))

		start = time.Now()
		mb.Shutdown()
		t.Logf("Seal: %s", time.Since(start))

		start = time.Now()
		val, ok := mb.Receive()
		t.Logf("Receive: %s", time.Since(start))
		assertlite.True(t, ok)
		assertlite.Equals(t, 1, val)

		start = time.Now()
		closeErr := mb.Close()
		t.Logf("Close: %s", time.Since(start))

		assertlite.IsErrorOf(t, closeErr, mailbox.ErrClosedTwice)
		assertlite.Equals(t, 0, mb.Len())
		assertlite.Equals(t, 4, mb.Cap())
	})

	t.Run("write failure propagates from buffer", func(t *testing.T) {
		buf := recorders.NewBuffer[int](2)
		mb := Buffered[int](buf)

		// Fill to capacity
		assertlite.True(t, mb.Send(1))
		assertlite.True(t, mb.Send(2))

		// Buffer rejects write
		assertlite.False(t, mb.Send(3))
		assertlite.Equals(t, 3, buf.WriteCalls)
		assertlite.EqualSlices(t, []int{1, 2, 3}, buf.Writes)
	})
}

// TestBuffered_Concurrent
//
// Multiple goroutines operating on a Buffered mailbox.
//
// We should assert:
//  1. Concurrent sends and receives are thread-safe
//  2. All sent values are eventually received
//  3. No values are lost or duplicated under concurrent access
//  4. Seal with concurrent receivers drains correctly
//  5. Close with concurrent senders rejects new sends
//  6. Latch signals exactly once after drain
func TestBuffered_Concurrent(t *testing.T) {

	t.Run("concurrent sends and receives are thread-safe", func(t *testing.T) {
		buf := recorders.NewBuffer[int](64)
		mb := Buffered[int](buf)
		const numOps = 1000

		var sent atomic.Int32
		var received atomic.Int32
		var sum atomic.Int64

		var wg sync.WaitGroup
		wg.Add(2)

		// Producer
		go func() {
			defer wg.Done()
			for i := 0; i < numOps; i++ {
				if mb.Send(i) {
					sent.Add(1)
				}
			}
		}()

		// Consumer
		go func() {
			defer wg.Done()
			for received.Load() < numOps {
				if v, ok := mb.Receive(); ok {
					received.Add(1)
					sum.Add(int64(v))
				}
			}
		}()

		wg.Wait()

		assertlite.Equals(t, int32(numOps), sent.Load())
		assertlite.Equals(t, int32(numOps), received.Load())

		// Sum of 0..999 = n*(n-1)/2 = 999*1000/2 = 499500
		expected := int64(numOps * (numOps - 1) / 2)
		assertlite.Equals(t, expected, sum.Load())
	})

	t.Run("seal drains remaining elements with concurrent receivers", func(t *testing.T) {
		buf := recorders.NewBuffer[int](128)
		mb := Buffered[int](buf)
		const numItems = 500

		var received atomic.Int32
		var sum atomic.Int64

		// Fill buffer
		for i := 0; i < numItems; i++ {
			mb.Send(i)
		}

		// start consume
		var wg sync.WaitGroup
		for i := 0; i < 4; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for {
					if v, ok := mb.Receive(); ok {
						received.Add(1)
						sum.Add(int64(v))
					} else {
						return
					}
				}
			}()
		}

		// Seal after consume start
		time.Sleep(10 * time.Millisecond)
		waiter := mb.Shutdown()

		// Wait for drain
		waiter.Wait()

		assertlite.Equals(t, int32(numItems), received.Load())

		expected := int64(numItems * (numItems - 1) / 2)
		assertlite.Equals(t, expected, sum.Load())
	})

	t.Run("close with concurrent senders rejects sends", func(t *testing.T) {
		buf := recorders.NewBuffer[int](16)
		mb := Buffered[int](buf)

		var accepted atomic.Int32
		var rejected atomic.Int32

		var wg sync.WaitGroup
		wg.Add(4)

		// start senders
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
					time.Sleep(time.Millisecond)
				}
			}()
		}

		// Close after some sends
		time.Sleep(20 * time.Millisecond)
		mb.Close()

		wg.Wait()

		total := accepted.Load() + rejected.Load()
		assertlite.True(t, total > 0, "no operations completed")
		assertlite.True(t, rejected.Load() > 0, "no sends rejected after close")
	})

	t.Run("no data race with concurrent close and receive", func(t *testing.T) {
		buf := recorders.NewBuffer[int](8)
		mb := Buffered[int](buf)

		mb.Send(1)
		mb.Send(2)

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

	t.Run("latch signals exactly once after seal and drain", func(t *testing.T) {
		buf := recorders.NewBuffer[int](4)
		mb := Buffered[int](buf)

		mb.Send(1)
		mb.Send(2)

		var signals atomic.Int32

		waiter := mb.Shutdown()

		// Drain
		go func() {
			for {
				if _, ok := mb.Receive(); !ok {
					break
				}
			}
		}()

		// Wait for drain
		waiter.Wait()

		// Try to signal again (should be no-op via sync.Once)
		mb.Close()

		assertlite.Equals(t, int32(1), signals.Load(),
			"latch should signal exactly once")
	})

	t.Run("concurrent seal and send does not lose values", func(t *testing.T) {
		buf := recorders.NewBuffer[int](64)
		mb := Buffered[int](buf)

		var sent atomic.Int32
		var received atomic.Int32

		var wg sync.WaitGroup
		wg.Add(2)

		// Sender
		go func() {
			defer wg.Done()
			for i := 0; i < 500; i++ {
				if mb.Send(i) {
					sent.Add(1)
				} else {
					return
				}
			}
		}()

		// Receiver that seals after a delay
		go func() {
			defer wg.Done()
			time.Sleep(5 * time.Millisecond)
			mb.Shutdown()
			for {
				if _, ok := mb.Receive(); ok {
					received.Add(1)
				} else {
					return
				}
			}
		}()

		wg.Wait()

		assertlite.Equals(t, sent.Load(), received.Load(),
			"all sent values must be received")
	})
}

// TestBuffered_StateTransitions
//
// Verifies correct state machine behavior of the mailbox.
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
func TestBuffered_StateTransitions(t *testing.T) {

	t.Run("initial state allows send and receive", func(t *testing.T) {
		buf := recorders.NewBuffer[int](4)
		mb := Buffered[int](buf)

		assertlite.True(t, mb.Send(1))
		v, ok := mb.Receive()
		assertlite.True(t, ok)
		assertlite.Equals(t, 1, v)
	})

	t.Run("sealed state rejects sends but allows receives", func(t *testing.T) {
		buf := recorders.NewBuffer[int](4)
		mb := Buffered[int](buf)

		mb.Send(1)
		mb.Shutdown()

		// Rejects send
		assertlite.False(t, mb.Send(2))

		// Allows receive
		v, ok := mb.Receive()
		assertlite.True(t, ok)
		assertlite.Equals(t, 1, v)
	})

	t.Run("closed state rejects sends and receives", func(t *testing.T) {
		buf := recorders.NewBuffer[int](4)
		mb := Buffered[int](buf)

		mb.Close()

		assertlite.False(t, mb.Send(1))
		_, ok := mb.Receive()
		assertlite.False(t, ok)
	})

	t.Run("receive on empty sealed transitions to closed", func(t *testing.T) {
		buf := recorders.NewBuffer[int](4)
		mb := Buffered[int](buf)

		mb.Send(1)
		mb.Shutdown()

		// Drain
		mb.Receive()

		// Receive on empty sealed triggers close
		_, ok := mb.Receive()
		assertlite.False(t, ok)

		// Verify closed state
		assertlite.False(t, mb.Send(2))
		assertlite.IsErrorOf(t, mailbox.ErrClosedTwice, mb.Close())
	})

	t.Run("close on open transitions to closed", func(t *testing.T) {
		buf := recorders.NewBuffer[int](4)
		mb := Buffered[int](buf)

		assertlite.NoError(t, mb.Close())
		assertlite.False(t, mb.Send(1))
	})

	t.Run("close on sealed transitions to closed", func(t *testing.T) {
		buf := recorders.NewBuffer[int](4)
		mb := Buffered[int](buf)

		mb.Send(1)
		mb.Shutdown()
		assertlite.NoError(t, mb.Close())

		assertlite.False(t, mb.Send(2))
	})

	t.Run("seal on sealed is idempotent", func(t *testing.T) {
		buf := recorders.NewBuffer[int](4)
		mb := Buffered[int](buf)

		mb.Send(1)
		w1 := mb.Shutdown()
		w2 := mb.Shutdown()

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
		buf := recorders.NewBuffer[int](4)
		mb := Buffered[int](buf)

		mb.Close()
		mb.Shutdown()

		assertlite.False(t, mb.Send(1))
		_, ok := mb.Receive()
		assertlite.False(t, ok)
	})
}
