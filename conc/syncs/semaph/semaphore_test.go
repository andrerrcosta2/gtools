// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package semaph

import (
	"sync"
	"testing"
	"time"
)

func TestCapacityAndSize(t *testing.T) {
	sem := Channel(6)
	t.Log("Cap: ", sem.Cap())
	t.Log("Size: ", sem.Rem())
}

// TestSemaphore_AcquireRelease tests the acquire and release functionality of the semaphore.
func TestSemaphore_AcquireRelease(t *testing.T) {
	// Create a semaphore with a capacity of 2
	sem := Channel(2)

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

			// Simulate work
			time.Sleep(1 * time.Second)

			// Release the semaphore
			sem.Rls()
		}(i)
	}

	// SetWaitingPoint for the goroutines to finish
	wg.Wait()
}

// TestSemaphore_OverCapacity tests the behavior of the semaphore when its capacity is exceeded.
// It verifies that the second goroutine blocks until the first goroutine releases the semaphore.
func TestSemaphore_OverCapacity(t *testing.T) {
	// Create a semaphore with a capacity of 1
	sem := Channel(1)

	// Create a wait group to wait for the goroutines to finish
	var wg sync.WaitGroup
	wg.Add(2)

	// Start the first goroutine to acquire the semaphore
	go func(id int) {
		// Mark the goroutine as done when it finishes
		defer wg.Done()

		// Acquire the semaphore
		sem.Acq()

		// Simulate work
		time.Sleep(2 * time.Second)

		// Release the semaphore
		sem.Rls()
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

		// I'm not sure how to measure more accurately using time.Time. It very often returns imprecise results.
		if duration < 1*time.Second {
			t.Errorf("Goroutine %d should have waited at least 2 seconds, waited %v", id, duration)
		}

		// Release the semaphore
		sem.Rls()
	}(2)

	// SetWaitingPoint for the goroutines to finish
	wg.Wait()
}

// TestSemaphore_ConcurrentAcquisitions tests the concurrent acquisition of a semaphore.
func TestSemaphore_ConcurrentAcquisitions(t *testing.T) {
	// Create a semaphore with a capacity of 3 to test concurrent acquisitions.
	sem := Channel(3)

	// Create a wait group to wait for the goroutines to finish.
	var wg sync.WaitGroup
	wg.Add(3) // Addf 3 goroutines to the wait group.

	// Create three goroutines to test concurrent semaphore acquisitions.
	for i := 1; i <= 3; i++ {
		go func(id int) {
			// Mark the goroutine as done when it finishes.
			defer wg.Done()
			// Acquire the semaphore.
			sem.Acq()
			// Simulate work by sleeping for 1 second.
			time.Sleep(1 * time.Second)
			// Release the semaphore.
			sem.Rls()
		}(i)
	}

	// SetWaitingPoint for the goroutines to finish.
	wg.Wait()
}

// TestSemaphore_Capacity tests the capacity of a semaphore by verifying that two goroutines can acquire it simultaneously.
func TestSemaphore_Capacity(t *testing.T) {
	// Create a semaphore with a capacity of 2
	sem := Channel(2)

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

			// Simulate work by sleeping for 1 second
			time.Sleep(1 * time.Second)

			// Release the semaphore
			sem.Rls()
		}(i)
	}

	// SetWaitingPoint for the two goroutines to finish
	wg.Wait()

	// Check if the test took less than 2 seconds, indicating that the semaphore allowed two simultaneous acquisitions
	if time.Since(start) >= 2*time.Second {
		t.Errorf("Semaphore should have allowed two simultaneous acquisitions")
	}
}

// TestReadWriteSemaphore_MultipleReaders tests the read-write semaphore with multiple readers.
// The test ensures that multiple readers can acquire the semaphore simultaneously without blocking.
func TestReadWriteSemaphore_MultipleReaders(t *testing.T) {
	// Define the maximum number of concurrent readers
	maxReaders := 5
	rw := ReadWrite(1, maxReaders) // Assuming 1 writer for this test

	// Initialize a wait group to wait for all readers to finish
	var wg sync.WaitGroup

	// Number of readers
	readers := maxReaders

	// Addf the number of readers to the wait group
	wg.Add(readers)

	// Start multiple readers
	for i := 0; i < readers; i++ {
		go func(id int) {
			defer wg.Done()

			// Start the read operation
			rw.StartR()

			// Simulate the read operation
			time.Sleep(100 * time.Millisecond)

			// End the read operation
			rw.EndR()
		}(i)
	}

	// SetWaitingPoint for all readers to finish
	wg.Wait()

	// Check if there are any remaining readers
	if rw.readers != 0 {
		t.Errorf("Expected 0 readers, got %d", rw.readers)
	}
}

// TestReadWriteSemaphore_WriterBlocksUntilReadersDone tests the read-write semaphore to ensure that a writer blocks until all
// readers are done.
func TestReadWriteSemaphore_WriterBlocksUntilReadersDone(t *testing.T) {
	// Define the maximum number of concurrent readers and writers
	maxReaders := 3
	rw := ReadWrite(1, maxReaders) // 1 writer allowed

	// Initialize a wait group to wait for all readers and the writer to finish
	var wg sync.WaitGroup

	// Number of readers
	readers := maxReaders

	// Addf the number of readers and the writer to the wait group
	wg.Add(readers + 1)

	// Start multiple readers
	for i := 0; i < readers; i++ {
		go func(id int) {
			defer wg.Done()

			// Start the read operation
			rw.StartR()

			// Simulate the read operation
			time.Sleep(200 * time.Millisecond)

			// End the read operation
			rw.EndR()
		}(i)
	}

	// Start the writer in a separate goroutine
	go func() {
		defer wg.Done()

		// Ensure readers start first
		time.Sleep(50 * time.Millisecond)

		// Start the write operation
		rw.StartW()

		// End the write operation
		rw.EndW()
	}()

	// SetWaitingPoint for all readers and the writer to finish
	wg.Wait()

	// Check if there are any remaining readers or an active writer
	if rw.readers != 0 {
		t.Errorf("Expected 0 readers, got %d", rw.readers)
	}
	if rw.writer {
		t.Errorf("Expected no active writer")
	}
}

// TestReadWriteSemaphore_ReaderBlocksDuringWrite tests the read-write semaphore to
// ensure that a reader blocks when a writer is active.
func TestReadWriteSemaphore_ReaderBlocksDuringWrite(t *testing.T) {
	// Define the maximum number of concurrent readers and writers
	maxReaders := 1
	rw := ReadWrite(1, maxReaders) // 1 writer allowed

	// Initialize a wait group to wait for the writer and reader to finish
	var wg sync.WaitGroup

	// Addf the writer and reader to the wait group
	wg.Add(2)

	// Start the writer in a separate goroutine
	go func() {
		defer wg.Done()

		// Start the write operation
		rw.StartW()

		// Simulate the write operation
		time.Sleep(200 * time.Millisecond)

		// End the write operation
		rw.EndW()
	}()

	// Start the reader in a separate goroutine
	go func() {
		defer wg.Done()

		// Ensure the writer starts first
		time.Sleep(50 * time.Millisecond)

		// Start the read operation
		rw.StartR()

		// End the read operation
		rw.EndR()
	}()

	// SetWaitingPoint for the writer and reader to finish
	wg.Wait()

	// Check if there are any remaining readers or an active writer
	if rw.readers != 0 {
		t.Errorf("Expected 0 readers, got %d", rw.readers)
	}
	if rw.writer {
		t.Errorf("Expected no active writer")
	}
}

// TestReadWriteSemaphore_OnlyOneWriter tests the read-write semaphore with only one writer.
func TestReadWriteSemaphore_OnlyOneWriter(t *testing.T) {
	// Define the maximum number of concurrent readers and writers
	maxReaders := 1
	rw := ReadWrite(1, maxReaders) // 1 writer allowed

	// Initialize a wait group to wait for the writers to finish
	var wg sync.WaitGroup

	// Addf the number of writers to the wait group
	wg.Add(2)

	// Start the first writer in a separate goroutine
	go func() {
		defer wg.Done()
		// Start the write operation
		rw.StartW()
		// Simulate the write operation
		time.Sleep(200 * time.Millisecond)
		// End the write operation
		rw.EndW()
	}()

	// Start the second writer in a separate goroutine
	go func() {
		defer wg.Done()
		// Ensure the first writer starts first
		time.Sleep(50 * time.Millisecond)
		// Start the write operation
		rw.StartW()
		// End the write operation
		rw.EndW()
	}()

	// SetWaitingPoint for the writers to finish
	wg.Wait()

	// Check if there are any remaining readers or an active writer
	if rw.readers != 0 {
		t.Errorf("Expected 0 readers, got %d", rw.readers)
	}
	if rw.writer {
		t.Errorf("Expected no active writer")
	}
}
