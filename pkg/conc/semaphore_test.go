// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package conc

import (
	"sync"
	"testing"
	"time"
)

// TestSemaphore_AcquireRelease tests the acquire and release functionality of the semaphore.
func TestSemaphore_AcquireRelease(t *testing.T) {
	// Create a semaphore with a capacity of 2
	sem := NewChannelSemaphore(2)

	// Create a wait group to wait for the goroutines to finish
	var wg sync.WaitGroup
	wg.Add(3)

	// Create three goroutines to test the semaphore
	for i := 1; i <= 3; i++ {
		go func(id int) {
			// Mark the goroutine as done when it finishes
			defer wg.Done()

			// Acquire the semaphore
			sem.Acq()
			t.Logf("Goroutine %d acquired the semaphore", id)

			// Simulate work
			time.Sleep(1 * time.Second)

			// Release the semaphore
			sem.Rls()
			t.Logf("Goroutine %d released the semaphore", id)
		}(i)
	}

	// Wait for the goroutines to finish
	wg.Wait()
}

// TestSemaphore_OverCapacity tests the behavior of the semaphore when its capacity is exceeded.
// It verifies that the second goroutine blocks until the first goroutine releases the semaphore.
func TestSemaphore_OverCapacity(t *testing.T) {
	// Create a semaphore with a capacity of 1
	sem := NewChannelSemaphore(1)

	// Create a wait group to wait for the goroutines to finish
	var wg sync.WaitGroup
	wg.Add(2)

	// Start the first goroutine to acquire the semaphore
	go func(id int) {
		// Mark the goroutine as done when it finishes
		defer wg.Done()

		// Acquire the semaphore
		sem.Acq()
		t.Logf("Goroutine %d acquired the semaphore", id)

		// Simulate work
		time.Sleep(2 * time.Second)

		// Release the semaphore
		sem.Rls()
		t.Logf("Goroutine %d released the semaphore", id)
	}(1)

	// Start the second goroutine to test blocking behavior
	go func(id int) {
		// Mark the goroutine as done when it finishes
		defer wg.Done()

		// Ensure this goroutine starts after the first
		time.Sleep(100 * time.Millisecond)

		// Record the start time
		start := time.Now()

		// Acquire the semaphore (this should block until the first goroutine releases the semaphore)
		sem.Acq()

		// Calculate the duration waited
		duration := time.Since(start)

		// Log the result
		t.Logf("Goroutine %d acquired the semaphore after waiting %v", id, duration)

		// I'm not sure how to measure more accurately using time.Time. It very often returns imprecise results.
		if duration < 1*time.Second {
			t.Errorf("Goroutine %d should have waited at least 2 seconds, waited %v", id, duration)
		}

		// Release the semaphore
		sem.Rls()
		t.Logf("Goroutine %d released the semaphore", id)
	}(2)

	// Wait for the goroutines to finish
	wg.Wait()
}

// TestSemaphore_ConcurrentAcquisitions tests the concurrent acquisition of a semaphore.
func TestSemaphore_ConcurrentAcquisitions(t *testing.T) {
	// Create a semaphore with a capacity of 3 to test concurrent acquisitions.
	sem := NewChannelSemaphore(3)

	// Create a wait group to wait for the goroutines to finish.
	var wg sync.WaitGroup
	wg.Add(3) // Add 3 goroutines to the wait group.

	// Create three goroutines to test concurrent semaphore acquisitions.
	for i := 1; i <= 3; i++ {
		go func(id int) {
			// Mark the goroutine as done when it finishes.
			defer wg.Done()
			// Acquire the semaphore.
			sem.Acq()
			t.Logf("Goroutine %d acquired the semaphore", id)
			// Simulate work by sleeping for 1 second.
			time.Sleep(1 * time.Second)
			// Release the semaphore.
			sem.Rls()
			t.Logf("Goroutine %d released the semaphore", id)
		}(i)
	}

	// Wait for the goroutines to finish.
	wg.Wait()
}

// TestSemaphore_Capacity tests the capacity of a semaphore by verifying that two goroutines can acquire it simultaneously.
func TestSemaphore_Capacity(t *testing.T) {
	// Create a semaphore with a capacity of 2
	sem := NewChannelSemaphore(2)

	// Create a wait group to wait for the two goroutines to finish
	var wg sync.WaitGroup
	wg.Add(2)

	// Record the start time to measure the duration of the test
	start := time.Now()

	// Test that two goroutines can acquire the semaphore simultaneously
	for i := 1; i <= 2; i++ {
		go func(id int) {
			// Mark the goroutine as done when it finishes
			defer wg.Done()

			// Acquire the semaphore
			sem.Acq()
			t.Logf("Goroutine %d acquired the semaphore", id)

			// Simulate work by sleeping for 1 second
			time.Sleep(1 * time.Second)

			// Release the semaphore
			sem.Rls()
			t.Logf("Goroutine %d released the semaphore", id)
		}(i)
	}

	// Wait for the two goroutines to finish
	wg.Wait()

	// Check if the test took less than 2 seconds, indicating that the semaphore allowed two simultaneous acquisitions
	if time.Since(start) >= 2*time.Second {
		t.Errorf("Semaphore should have allowed two simultaneous acquisitions")
	}
}

// TestReadWriteSemaphore_MultipleReaders tests the read-write semaphore with multiple readers.
func TestReadWriteSemaphore_MultipleReaders(t *testing.T) {
	// Create a new read-write semaphore
	rw := NewReadWriteSemaphore()

	// Initialize a wait group to wait for all readers to finish
	var wg sync.WaitGroup

	// Number of readers
	readers := 5

	// Add the number of readers to the wait group
	wg.Add(readers)

	// Start multiple readers
	for i := 0; i < readers; i++ {
		// Start a new reader in a separate goroutine
		go func(id int) {
			// Defer the Done call to ensure it's called when the goroutine exits
			defer wg.Done()

			// Start the read operation
			rw.StartR()
			t.Logf("Reader %d started reading", id)

			// Simulate the read operation
			time.Sleep(100 * time.Millisecond)

			// Log the end of the read operation
			t.Logf("Reader %d finished reading", id)

			// End the read operation
			rw.EndR()
		}(i)
	}

	// Wait for all readers to finish
	wg.Wait()

	// Check if there are any remaining readers
	if rw.readers != 0 {
		// If there are, report an error
		t.Errorf("Expected 0 readers, got %d", rw.readers)
	}
}

// TestReadWriteSemaphore_WriterBlocksUntilReadersDone tests that a writer blocks until all readers are done.
func TestReadWriteSemaphore_WriterBlocksUntilReadersDone(t *testing.T) {
	// Create a new read-write semaphore
	rw := NewReadWriteSemaphore()

	// Initialize a wait group to wait for all readers and the writer to finish
	var wg sync.WaitGroup

	// Number of readers
	readers := 3

	// Add the number of readers and the writer to the wait group
	wg.Add(readers + 1)

	// Start multiple readers
	for i := 0; i < readers; i++ {
		// Start a new reader in a separate goroutine
		go func(id int) {
			// Defer the Done call to ensure it's called when the goroutine exits
			defer wg.Done()

			// Start the read operation
			rw.StartR()
			t.Logf("Reader %d started reading", id)

			// Simulate the read operation
			time.Sleep(200 * time.Millisecond)

			// Log the end of the read operation
			t.Logf("Reader %d finished reading", id)

			// End the read operation
			rw.EndR()
		}(i)
	}

	// Start the writer in a separate goroutine
	go func() {
		// Defer the Done call to ensure it's called when the goroutine exits
		defer wg.Done()

		// Ensure readers start first
		time.Sleep(50 * time.Millisecond)

		// Start the write operation
		rw.StartW()
		t.Log("Writer started writing")

		// End the write operation
		rw.EndW()
		t.Log("Writer finished writing")
	}()

	// Wait for all readers and the writer to finish
	wg.Wait()

	// Check if there are any remaining readers or an active writer
	if rw.readers != 0 {
		t.Errorf("Expected 0 readers, got %d", rw.readers)
	}
	if rw.writer {
		t.Errorf("Expected no active writer")
	}
}

// TestReadWriteSemaphore_ReaderBlocksDuringWrite tests that a reader blocks during a write operation.
func TestReadWriteSemaphore_ReaderBlocksDuringWrite(t *testing.T) {
	// Create a new read-write semaphore
	rw := NewReadWriteSemaphore()

	// Initialize a wait group to wait for the writer and reader to finish
	var wg sync.WaitGroup

	// Add the writer and reader to the wait group
	wg.Add(2)

	// Start the writer in a separate goroutine
	go func() {
		// Defer the Done call to ensure it's called when the goroutine exits
		defer wg.Done()

		// Start the write operation
		rw.StartW()
		t.Log("Writer started writing")

		// Simulate the write operation
		time.Sleep(200 * time.Millisecond)

		// Log the end of the write operation
		t.Log("Writer finished writing")

		// End the write operation
		rw.EndW()
	}()

	// Start the reader in a separate goroutine
	go func() {
		// Defer the Done call to ensure it's called when the goroutine exits
		defer wg.Done()

		// Ensure the writer starts first
		time.Sleep(50 * time.Millisecond)

		// Start the read operation
		rw.StartR()
		t.Log("Reader started reading")

		// End the read operation
		rw.EndR()
		t.Log("Reader finished reading")
	}()

	// Wait for the writer and reader to finish
	wg.Wait()

	// Check if there are any remaining readers or an active writer
	if rw.readers != 0 {
		t.Errorf("Expected 0 readers, got %d", rw.readers)
	}
	if rw.writer {
		t.Errorf("Expected no active writer")
	}
}

// TestReadWriteSemaphore_OnlyOneWriter tests that only one writer can write at a time.
func TestReadWriteSemaphore_OnlyOneWriter(t *testing.T) {
	// Create a new read-write semaphore
	rw := NewReadWriteSemaphore()
	// Initialize a wait group to wait for the writers to finish
	var wg sync.WaitGroup

	// Add the number of writers to the wait group
	wg.Add(2)

	// Start the first writer in a separate goroutine
	go func() {
		// Defer the Done call to ensure it's called when the goroutine exits
		defer wg.Done()
		// Start the write operation
		rw.StartW()
		t.Log("Writer 1 started writing")
		// Simulate the write operation
		time.Sleep(200 * time.Millisecond)
		t.Log("Writer 1 finished writing")
		// End the write operation
		rw.EndW()
	}()

	// Start the second writer in a separate goroutine
	go func() {
		// Defer the Done call to ensure it's called when the goroutine exits
		defer wg.Done()
		// Ensure the first writer starts first
		time.Sleep(50 * time.Millisecond)
		// Start the write operation
		rw.StartW()
		t.Log("Writer 2 started writing")
		// End the write operation
		rw.EndW()
		t.Log("Writer 2 finished writing")
	}()

	// Wait for the writers to finish
	wg.Wait()
	// Check if there are any remaining readers or an active writer
	if rw.readers != 0 {
		t.Errorf("Expected 0 readers, got %d", rw.readers)
	}
	if rw.writer {
		t.Errorf("Expected no active writer")
	}
}
