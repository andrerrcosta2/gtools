// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package mailboxes

import (
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/andrerrcosta2/gtools/conc/core/mailbox"
	"github.com/andrerrcosta2/gtools/core/data/closeable"
	"github.com/andrerrcosta2/gtools/core/testlite"
	"github.com/andrerrcosta2/gtools/core/testlite/assertlite"
)

// TestDisruptor_Sync
//
// Single-goroutine operations on a Disruptor mailbox.
//
// We should assert:
//  1. Send writes to buffer and returns true
//  2. Receive reads from buffer in FIFO order
//  3. Receive on empty buffer blocks until data arrives (wake mechanism)
//  4. Send to closed mailbox returns false
//  5. Close transitions state and prevents further operations
//  6. Double close returns mailbox.ErrClosedTwice
//  7. Seal prevents new sends but allows draining existing elements
//  8. Drain processes multiple messages without blocking
//  9. DrainBlocking waits for data or close
//
// 10. Disruptor requires power-of-two size
func TestDisruptor_Sync(t *testing.T) {

	t.Run("panics on buffer not power of two", func(t *testing.T) {
		assertlite.Panics(t, func() {
			Disruptor[int](3)
		})
	})

	t.Run("send writes to buffer and returns true", func(t *testing.T) {
		mb := Disruptor[int](4)
		assertlite.True(t, mb.Send(42))
	})

	t.Run("receive reads from buffer in FIFO order", func(t *testing.T) {
		mb := Disruptor[int](8)

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
	})

	t.Run("receive blocks until data available via wake mechanism", func(t *testing.T) {
		mb := Disruptor[int](4)

		var received int
		var ok bool
		done := make(chan struct{})

		go func() {
			received, ok = mb.Receive()
			close(done)
		}()

		// Give receiver time to block
		time.Sleep(50 * time.Millisecond)

		// Send data - should trigger wake
		mb.Send(42)

		select {
		case <-done:
			assertlite.True(t, ok)
			assertlite.Equals(t, 42, received)
		case <-time.After(time.Second):
			t.Fatal("receive did not complete after send (wake mechanism failed)")
		}
	})

	t.Run("send to closed mailbox returns false", func(t *testing.T) {
		mb := Disruptor[int](4)

		err := mb.Close()
		assertlite.NoError(t, err)

		assertlite.False(t, mb.Send(42))
	})

	t.Run("double close returns mailbox.ErrClosed", func(t *testing.T) {
		mb := Disruptor[int](4)

		assertlite.NoError(t, mb.Close())
		assertlite.IsErrorOf(t, mailbox.ErrClosedTwice, mb.Close())
	})

	t.Run("seal prevents new sends but allows draining", func(t *testing.T) {
		mb := Disruptor[int](8)

		mb.Send(10)
		mb.Send(20)

		waiter := mb.Shutdown()

		// Cannot send after seal
		assertlite.False(t, mb.Send(30))

		// Can still drain existing elements
		v1, ok1 := mb.Receive()
		assertlite.True(t, ok1)
		assertlite.Equals(t, 10, v1)

		// Waiter should be signaled after drain
		done := make(chan struct{})
		go func() {
			waiter.Wait()
			close(done)
		}()

		v2, ok2 := mb.Receive()
		assertlite.True(t, ok2)
		assertlite.Equals(t, 20, v2)

		lmax, ok := mb.(*lmax[int])
		assertlite.True(t, ok)
		assertlite.True(t, closeable.IsClosed(lmax.state), "expected closed state after drained sealed mailbox, but got %v", lmax.state)
		assertlite.True(t, lmax.isEmpty(), "expected empty buffer after seal, but got "+
			"readSeq=%v, writeSeq=%v", lmax.readSeq, lmax.writeSeq)

		select {
		case <-done:
			// Success
		case <-time.After(time.Second):
			t.Fatal("waiter did not complete after sealed buffer drained")
		}
	})

	t.Run("seal on empty mailbox signals immediately", func(t *testing.T) {
		mb := Disruptor[int](4)

		waiter := mb.Shutdown()

		done := make(chan struct{})
		go func() {
			waiter.Wait()
			close(done)
		}()

		select {
		case <-done:
			// Success - empty mailbox should signal immediately
		case <-time.After(time.Second):
			t.Fatal("waiter did not complete on empty sealed mailbox")
		}

		// Should be closed now
		assertlite.False(t, mb.Send(1))
	})

	t.Run("drain processes multiple messages without blocking", func(t *testing.T) {
		mb := Disruptor[int](16)

		// Send multiple messages
		for i := 0; i < 10; i++ {
			mb.Send(i)
		}

		var received []int
		count := mb.Drain(func(v int) {
			received = append(received, v)
		})

		assertlite.Equals(t, 10, count)
		assertlite.Equals(t, 10, len(received))

		for i := 0; i < 10; i++ {
			assertlite.Equals(t, i, received[i])
		}
	})

	t.Run("drain on empty mailbox returns zero", func(t *testing.T) {
		mb := Disruptor[int](4)

		var received []int
		count := mb.Drain(func(v int) {
			received = append(received, v)
		})

		assertlite.Equals(t, 0, count)
		assertlite.Equals(t, 0, len(received))
	})

	t.Run("drain blocking waits for data then drains", func(t *testing.T) {
		mb := Disruptor[int](16)

		var received []int
		done := make(chan bool)

		go func() {
			result := mb.DrainBlocking(func(v int) {
				received = append(received, v)
			})
			done <- result
		}()

		// Give DrainBlocking time to start waiting
		time.Sleep(50 * time.Millisecond)

		// Send data - should wake DrainBlocking
		for i := 0; i < 5; i++ {
			mb.Send(i)
		}

		select {
		case result := <-done:
			assertlite.True(t, result)
			assertlite.Equals(t, 5, len(received))
			for i := 0; i < 5; i++ {
				assertlite.Equals(t, i, received[i])
			}
		case <-time.After(time.Second):
			t.Fatal("DrainBlocking did not complete after send")
		}
	})

	t.Run("drain blocking returns false when closed while waiting", func(t *testing.T) {
		mb := Disruptor[int](16)

		done := make(chan bool)

		go func() {
			result := mb.DrainBlocking(func(v int) {})
			done <- result
		}()

		// Give DrainBlocking time to start waiting
		time.Sleep(50 * time.Millisecond)

		// Close instead of sending - should unblock with false
		mb.Close()

		select {
		case result := <-done:
			assertlite.False(t, result)
		case <-time.After(time.Second):
			t.Fatal("DrainBlocking did not complete after close")
		}
	})

	t.Run("drain blocking returns false on already closed mailbox", func(t *testing.T) {
		mb := Disruptor[int](16)
		mb.Close()

		result := mb.DrainBlocking(func(v int) {})
		assertlite.False(t, result)
	})

	t.Run("drain blocking processes messages and returns true when data available", func(t *testing.T) {
		mb := Disruptor[int](16)

		// Pre-fill with data
		for i := 0; i < 3; i++ {
			mb.Send(i)
		}

		var received []int
		result := mb.DrainBlocking(func(v int) {
			received = append(received, v)
		})

		assertlite.True(t, result)
		assertlite.EqualSlices(t, []int{0, 1, 2}, received)
	})

	t.Run("drain after seal triggers close", func(t *testing.T) {
		mb := Disruptor[int](8)

		mb.Send(1)
		mb.Send(2)
		mb.Send(3)

		waiter := mb.Shutdown()

		var received []int
		count := mb.Drain(func(v int) {
			received = append(received, v)
		})

		assertlite.Equals(t, 3, count)
		assertlite.EqualSlices(t, []int{1, 2, 3}, received)

		// Wait for close
		waiter.Wait()

		// Verify closed state
		assertlite.False(t, mb.Send(4))
	})

	t.Run("power of two size requirement", func(t *testing.T) {
		// Should not panic with power of two
		assertlite.NoPanics(t, func() {
			Disruptor[int](2)
			Disruptor[int](4)
			Disruptor[int](8)
			Disruptor[int](16)
			Disruptor[int](1024)
		})
	})

	t.Run("close after seal and drain works", func(t *testing.T) {
		mb := Disruptor[int](4)

		mb.Send(1)
		mb.Shutdown()
		mb.Drain(func(v int) {})

		time.Sleep(100 * time.Millisecond)

		// Should already be closed
		assertlite.IsErrorOf(t, mb.Close(), mailbox.ErrClosedTwice)
	})

	t.Run("receive preserves FIFO order with large sequences", func(t *testing.T) {
		mb := Disruptor[int](16)
		const numMessages = 100

		// Send many messages
		go func() {
			for i := 0; i < numMessages; i++ {
				mb.Send(i)
			}
		}()

		// Receive all and verify order
		for i := 0; i < numMessages; i++ {
			v, ok := mb.Receive()
			assertlite.True(t, ok)
			assertlite.Equals(t, i, v)
		}
	})
}

// TestDisruptor_Concurrent
//
// Multiple goroutines operating on a Disruptor mailbox.
//
// We should assert:
//  1. Multiple producers can send concurrently without data loss
//  2. All sent values are eventually received in FIFO order
//  3. No values are lost or duplicated under high contention
//  4. Seal with concurrent receivers drains correctly
//  5. Close with concurrent senders rejects new sends
//  6. DrainBlocking handles concurrent producers correctly
//  7. Wake mechanism works under concurrent load
func TestDisruptor_Concurrent(t *testing.T) {

	t.Run("multiple producers send concurrently without loss", func(t *testing.T) {
		mb := Disruptor[int](1024)
		const numProducers = 8
		const msgsPerProducer = 1000
		const totalMsgs = numProducers * msgsPerProducer

		var sent sync.WaitGroup
		sent.Add(numProducers)

		// Multiple producers
		for p := 0; p < numProducers; p++ {
			go func(producerID int) {
				defer sent.Done()
				for i := 0; i < msgsPerProducer; i++ {
					mb.Send(producerID*msgsPerProducer + i)
				}
			}(p)
		}

		// Collect all messages
		received := make(map[int]bool)
		var mu sync.Mutex
		var receivedCount atomic.Int32

		go func() {
			for receivedCount.Load() < totalMsgs {
				if v, ok := mb.Receive(); ok {
					mu.Lock()
					received[v] = true
					mu.Unlock()
					receivedCount.Add(1)
				}
			}
		}()

		sent.Wait()

		// Wait for all receives
		timeout := time.After(5 * time.Second)
		ticker := time.NewTicker(100 * time.Millisecond)
		defer ticker.Stop()

		for {
			select {
			case <-timeout:
				t.Fatalf("timeout waiting for messages, received '%d' of '%d'",
					receivedCount.Load(), totalMsgs)
			case <-ticker.C:
				if receivedCount.Load() == totalMsgs {
					goto done
				}
			}
		}
	done:

		assertlite.Equals(t, int32(totalMsgs), receivedCount.Load())
		assertlite.Equals(t, totalMsgs, len(received))
	})

	t.Run("no data race with concurrent send and receive", func(t *testing.T) {
		mb := Disruptor[int](64)
		const numOps = 1000

		var wg sync.WaitGroup
		wg.Add(3)

		// Senders
		for i := 0; i < 2; i++ {
			go func() {
				defer wg.Done()
				for j := 0; j < numOps; j++ {
					mb.Send(j)
				}
			}()
		}

		// Single consumer (MPSC pattern)
		go func() {
			defer wg.Done()
			for j := 0; j < numOps*2; j++ {
				mb.Receive()
			}
		}()

		wg.Wait()

	})

	t.Run("drain blocking with concurrent producers", func(t *testing.T) {
		mb := Disruptor[int](256)
		const numProducers = 4
		const msgsPerProducer = 250

		var produced atomic.Int32
		var consumed atomic.Int32
		var wg sync.WaitGroup

		// Start consumer with DrainBlocking
		wg.Add(1)
		go func() {
			defer wg.Done()
			for consumed.Load() < int32(numProducers*msgsPerProducer) {
				if !mb.DrainBlocking(func(v int) {
					consumed.Add(1)
				}) {
					return
				}
			}
		}()

		// Start producers
		for p := 0; p < numProducers; p++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for i := 0; i < msgsPerProducer; i++ {
					if mb.Send(i) {
						produced.Add(1)
					}
					runtime.Gosched()
				}
			}()
		}

		wg.Wait()

		assertlite.Equals(t, produced.Load(), consumed.Load(),
			"all produced messages must be consumed")
	})

	t.Run("seal drains with concurrent producers", func(t *testing.T) {
		mb := Disruptor[int](256)
		const numProducers = 4
		const initialMsgs = 100

		// Pre-fill with some messages
		for i := 0; i < initialMsgs; i++ {
			mb.Send(i)
		}

		var received atomic.Int32
		var produced atomic.Int32
		produced.Store(int32(initialMsgs))

		// Start producers
		var producers sync.WaitGroup
		producers.Add(numProducers)
		for p := 0; p < numProducers; p++ {
			go func() {
				defer producers.Done()
				for i := 0; i < 50; i++ {
					if mb.Send(i) {
						produced.Add(1)
					} else {
						return
					}
				}
			}()
		}

		// Seal after a brief moment
		time.Sleep(10 * time.Millisecond)
		waiter := mb.Shutdown()

		// Drain all messages
		go func() {
			for {
				if _, ok := mb.Receive(); ok {
					received.Add(1)
				} else {
					return
				}
			}
		}()

		producers.Wait()
		waiter.Wait()

		assertlite.Equals(t, produced.Load(), received.Load(),
			"all produced messages must be received")
	})

	t.Run("close with concurrent senders rejects sends", func(t *testing.T) {
		mb := Disruptor[int](16)

		counter := testlite.Map(t)

		var wg sync.WaitGroup
		wg.Add(4)

		// Start senders
		for i := 0; i < 4; i++ {
			go func() {
				defer wg.Done()
				for j := 0; j < 100; j++ {
					if mb.Send(j) {
						counter.Add(1, "accepted")
					} else {
						counter.Add(1, "rejected")
						return
					}
					runtime.Gosched()
				}
			}()
		}

		// Close after some sends
		time.Sleep(20 * time.Millisecond)
		mb.Close()

		wg.Wait()

		total := counter.Get("accepted", "rejected")
		assertlite.True(t, total > 0, "no operations completed")
		assertlite.True(t, counter.Get("rejected") > 0, "no sends rejected after close")
	})

	t.Run("drain blocking unblocks on close with concurrent producers", func(t *testing.T) {
		mb := Disruptor[int](64)

		var wg sync.WaitGroup
		var drainResults atomic.Int32

		// Multiple DrainBlocking consumers
		for i := 0; i < 3; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				if mb.DrainBlocking(func(v int) {}) {
					drainResults.Add(1)
				}
			}()
		}

		// Close after a delay
		time.Sleep(50 * time.Millisecond)
		mb.Close()

		wg.Wait()

		// Some may have drained data, others should have returned false
		t.Logf("DrainBlocking results: %d successful drains", drainResults.Load())
	})

	t.Run("wake mechanism handles rapid send-drain cycles", func(t *testing.T) {
		mb := Disruptor[int](64)
		const cycles = 100

		var consumed atomic.Int32
		done := make(chan struct{})

		// Consumer using DrainBlocking
		go func() {
			for consumed.Load() < int32(cycles) {
				if !mb.DrainBlocking(func(v int) {
					consumed.Add(1)
				}) {
					return
				}
			}
			close(done)
		}()

		// Producer sending in bursts
		for i := 0; i < cycles; i++ {
			mb.Send(i)
			if i%10 == 0 {
				time.Sleep(time.Microsecond) // Small delay between bursts
			}
		}

		select {
		case <-done:
			assertlite.Equals(t, int32(cycles), consumed.Load())
		case <-time.After(5 * time.Second):
			t.Fatalf("timeout: consumed %d of %d", consumed.Load(), cycles)
		}
	})

	t.Run("backpressure prevents overflow under load", func(t *testing.T) {
		mb := Disruptor[int](8) // Small buffer to force backpressure
		const msgsPerProducer = 1000
		const numProducers = 4

		var sent atomic.Int32
		var wg sync.WaitGroup
		wg.Add(numProducers)

		// Fast producers
		for p := 0; p < numProducers; p++ {
			go func() {
				defer wg.Done()
				for i := 0; i < msgsPerProducer; i++ {
					if mb.Send(i) {
						sent.Add(1)
					}
				}
			}()
		}

		// Slow consumer
		var received atomic.Int32
		go func() {
			for received.Load() < int32(numProducers*msgsPerProducer) {
				if _, ok := mb.Receive(); ok {
					received.Add(1)
					time.Sleep(time.Microsecond) // Slow consumer
				}
			}
		}()

		wg.Wait()

		// Wait for consumer to catch up
		timeout := time.After(10 * time.Second)
		ticker := time.NewTicker(100 * time.Millisecond)
		defer ticker.Stop()

		for {
			select {
			case <-timeout:
				t.Fatalf("timeout: sent %d, received %d", sent.Load(), received.Load())
			case <-ticker.C:
				if received.Load() == int32(numProducers*msgsPerProducer) {
					return
				}
			}
		}
	})
}

// TestDisruptor_StateTransitions
//
// Verifies correct state machine behavior of the Disruptor mailbox.
//
// We should assert:
//  1. Initial state allows send and receive
//  2. Sealed state rejects sends, allows receives
//  3. Closed state rejects sends, wake is signaled
//  4. Seal on empty transitions to closed immediately
//  5. Close on open transitions to closed, signals wake
//  6. Close on sealed transitions to closed
//  7. Seal on sealed is idempotent
//  8. Seal on closed is no-op
func TestDisruptor_StateTransitions(t *testing.T) {

	t.Run("initial state allows send and receive", func(t *testing.T) {
		mb := Disruptor[int](4)

		assertlite.True(t, mb.Send(1))
		v, ok := mb.Receive()
		assertlite.True(t, ok)
		assertlite.Equals(t, 1, v)
	})

	t.Run("sealed state rejects sends but allows receives", func(t *testing.T) {
		mb := Disruptor[int](4)

		mb.Send(1)
		mb.Shutdown()

		// Rejects send
		assertlite.False(t, mb.Send(2))

		// Allows receive
		v, ok := mb.Receive()
		assertlite.True(t, ok)
		assertlite.Equals(t, 1, v)
	})

	t.Run("closed state rejects sends", func(t *testing.T) {
		mb := Disruptor[int](4)

		mb.Close()

		assertlite.False(t, mb.Send(1))
	})

	t.Run("close signals wake, unblocking DrainBlocking", func(t *testing.T) {
		mb := Disruptor[int](4)

		done := make(chan bool)

		go func() {
			result := mb.DrainBlocking(func(v int) {})
			done <- result
		}()

		// Give DrainBlocking time to start waiting
		time.Sleep(50 * time.Millisecond)

		// Close should signal wake
		mb.Close()

		select {
		case result := <-done:
			assertlite.False(t, result)
		case <-time.After(time.Second):
			t.Fatal("DrainBlocking was not unblocked by close")
		}
	})

	t.Run("seal on empty transitions to closed", func(t *testing.T) {
		mb := Disruptor[int](4)

		mb.Shutdown()

		// Should be closed immediately
		assertlite.False(t, mb.Send(1))
		assertlite.IsErrorOf(t, mailbox.ErrClosedTwice, mb.Close())
	})

	t.Run("close on open transitions to closed", func(t *testing.T) {
		mb := Disruptor[int](4)

		assertlite.NoError(t, mb.Close())
		assertlite.False(t, mb.Send(1))
	})

	t.Run("close on sealed transitions to closed", func(t *testing.T) {
		mb := Disruptor[int](4)

		mb.Send(1)
		mb.Shutdown()
		assertlite.NoError(t, mb.Close())

		assertlite.False(t, mb.Send(2))
	})

	t.Run("seal on sealed is idempotent", func(t *testing.T) {
		mb := Disruptor[int](4)

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
		mb := Disruptor[int](4)

		mb.Close()
		mb.Shutdown()

		assertlite.False(t, mb.Send(1))
	})
}

// TestDisruptor_DrainBlocking
//
// Specific tests for the DrainBlocking mechanism.
//
// We should assert:
//  1. DrainBlocking waits when mailbox is empty
//  2. DrainBlocking is woken by Send
//  3. DrainBlocking is woken by Close
//  4. DrainBlocking drains all available messages
//  5. Multiple DrainBlocking callers are handled correctly
//  6. DrainBlocking after seal drains remaining then closes
func TestDisruptor_DrainBlocking(t *testing.T) {

	t.Run("blocks when empty and wakes on send", func(t *testing.T) {
		mb := Disruptor[int](16)

		blocked := make(chan struct{})
		unblocked := make(chan bool)

		go func() {
			close(blocked)
			result := mb.DrainBlocking(func(v int) {})
			unblocked <- result
		}()

		// Wait for goroutine to block
		<-blocked
		time.Sleep(50 * time.Millisecond)

		// Send should wake it
		mb.Send(42)

		select {
		case result := <-unblocked:
			assertlite.True(t, result)
		case <-time.After(time.Second):
			t.Fatal("DrainBlocking was not woken by send")
		}
	})

	t.Run("wakes on close and returns false", func(t *testing.T) {
		mb := Disruptor[int](16)

		blocked := make(chan struct{})
		unblocked := make(chan bool)

		go func() {
			close(blocked)
			result := mb.DrainBlocking(func(v int) {})
			unblocked <- result
		}()

		// Wait for goroutine to block
		<-blocked
		time.Sleep(50 * time.Millisecond)

		// Close should wake it with false
		mb.Close()

		select {
		case result := <-unblocked:
			assertlite.False(t, result)
		case <-time.After(time.Second):
			t.Fatal("DrainBlocking was not woken by close")
		}
	})

	t.Run("drains all available messages when woken", func(t *testing.T) {
		mb := Disruptor[int](32)

		// Send messages before DrainBlocking
		for i := 0; i < 10; i++ {
			mb.Send(i)
		}

		var received []int
		mb.DrainBlocking(func(v int) {
			received = append(received, v)
		})

		assertlite.Equals(t, 10, len(received))
		for i := 0; i < 10; i++ {
			assertlite.Equals(t, i, received[i])
		}
	})

	t.Run("handles rapid sequence of drain blocking calls", func(t *testing.T) {
		mb := Disruptor[int](32)
		const cycles = 50

		var totalReceived atomic.Int32

		// Consumer loop with DrainBlocking
		go func() {
			for totalReceived.Load() < int32(cycles) {
				if !mb.DrainBlocking(func(v int) {
					totalReceived.Add(1)
				}) {
					return
				}
			}
		}()

		// Send messages with delays
		for i := 0; i < cycles; i++ {
			mb.Send(i)
			time.Sleep(time.Microsecond * 10)
		}

		// Wait for all to be received
		timeout := time.After(5 * time.Second)
		ticker := time.NewTicker(50 * time.Millisecond)
		defer ticker.Stop()

		for {
			select {
			case <-timeout:
				t.Fatalf("timeout: received %d of %d", totalReceived.Load(), cycles)
			case <-ticker.C:
				if totalReceived.Load() == int32(cycles) {
					return
				}
			}
		}
	})

	t.Run("drain blocking after seal drains and signals", func(t *testing.T) {
		mb := Disruptor[int](16)

		mb.Send(1)
		mb.Send(2)
		mb.Send(3)

		waiter := mb.Shutdown()

		var received []int
		result := mb.DrainBlocking(func(v int) {
			received = append(received, v)
		})

		assertlite.True(t, result)
		assertlite.EqualSlices(t, []int{1, 2, 3}, received)

		// Waiter should be signaled
		done := make(chan struct{})
		go func() {
			waiter.Wait()
			close(done)
		}()

		select {
		case <-done:
			// Success
		case <-time.After(time.Second):
			t.Fatal("waiter not signaled after drain blocking on sealed mailbox")
		}
	})

	t.Run("concurrent drain blocking and send operations", func(t *testing.T) {
		mb := Disruptor[int](64)
		const numConsumers = 3
		const numMessages = 100

		var consumed atomic.Int32
		var wg sync.WaitGroup

		// Multiple DrainBlocking consumers
		for i := 0; i < numConsumers; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for consumed.Load() < int32(numMessages) {
					if !mb.DrainBlocking(func(v int) {
						consumed.Add(1)
					}) {
						return
					}
				}
			}()
		}

		// Producer
		for i := 0; i < numMessages; i++ {
			mb.Send(i)
			runtime.Gosched()
		}

		wg.Wait()

		assertlite.Equals(t, int32(numMessages), consumed.Load())
	})
}

// TestDisruptor_Performance
//
// Performance and scalability characteristics of Disruptor mailbox.
//
// We should assert:
//  1. Disruptor handles high throughput without data loss
//  2. False sharing mitigation works (padding)
//  3. Drain is more efficient than individual receives
//  4. DrainBlocking is efficient for consumer loops
//  5. Backpressure works under sustained load
func TestDisruptor_Performance(t *testing.T) {

	if testing.Short() {
		t.Skip("skipping performance test in short mode")
	}

	t.Run("high throughput message processing", func(t *testing.T) {
		mb := Disruptor[int](4096)
		const totalMsgs = 100000
		const numProducers = 8

		var produced atomic.Int64
		var consumed atomic.Int64
		start := time.Now()

		// Producers
		var producers sync.WaitGroup
		producers.Add(numProducers)
		for p := 0; p < numProducers; p++ {
			go func() {
				defer producers.Done()
				for i := 0; i < totalMsgs/numProducers; i++ {
					mb.Send(i)
					produced.Add(1)
				}
			}()
		}

		// Consumer with Drain
		go func() {
			for consumed.Load() < int64(totalMsgs) {
				count := mb.Drain(func(v int) {
					consumed.Add(1)
				})
				if count == 0 {
					runtime.Gosched()
				}
			}
		}()

		producers.Wait()

		// Wait for consumer
		timeout := time.After(10 * time.Second)
		ticker := time.NewTicker(100 * time.Millisecond)
		defer ticker.Stop()

		for {
			select {
			case <-timeout:
				t.Fatal("timeout waiting for consumer")
			case <-ticker.C:
				if consumed.Load() == int64(totalMsgs) {
					duration := time.Since(start)
					throughput := float64(totalMsgs) / duration.Seconds()
					t.Logf("Throughput: %.2f msgs/sec with %d producers",
						throughput, numProducers)
					assertlite.True(t, throughput > 100000,
						"throughput too low: %.2f msgs/sec", throughput)
					return
				}
			}
		}
	})

	t.Run("drain is more efficient than individual receives", func(t *testing.T) {
		mb1 := Disruptor[int](1024)
		mb2 := Disruptor[int](1024)
		const totalMsgs = 50000

		// Fill both mailboxes
		for i := 0; i < totalMsgs; i++ {
			mb1.Send(i)
			mb2.Send(i)
		}

		// Test individual receives
		start1 := time.Now()
		for i := 0; i < totalMsgs; i++ {
			mb1.Receive()
		}
		duration1 := time.Since(start1)

		// Test drain
		start2 := time.Now()
		mb2.Drain(func(v int) {})
		duration2 := time.Since(start2)

		t.Logf("Individual receives: %v", duration1)
		t.Logf("Drain: %v", duration2)
		t.Logf("Drain speedup: %.2fx", float64(duration1)/float64(duration2))

		assertlite.True(t, duration2 < duration1,
			"drain should be faster than individual receives")
	})
}
