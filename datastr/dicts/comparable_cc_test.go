// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package dicts

import (
	"fmt"
	"slices"
	"sync"
	"testing"
	"time"

	"github.com/andrerrcosta2/gtools/core/data/comparator/comparators"
	"github.com/andrerrcosta2/gtools/datastr/internal/tests"
)

// TestConcComparable_ConcurrentPutGet tests concurrent Put and Get operations
func TestConcComparable_ConcurrentPutGet(t *testing.T) {
	type N = *tests.SortableNode

	m := ConcComparable[N, N]()

	var wg sync.WaitGroup
	numGoroutines := 100
	nodesPerGoroutine := 10

	// store expected values for verification
	expected := sync.Map{}

	// Concurrent writes
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(goroutineID int) {
			defer wg.Done()
			for j := 0; j < nodesPerGoroutine; j++ {
				key := tests.NewSortableNode(fmt.Sprintf("goroutine-%d-key-%d", goroutineID, j))
				value := tests.NewSortableNode(fmt.Sprintf("goroutine-%d-value-%d", goroutineID, j))
				m.Put(key, value)
				expected.Store(key, value)
			}
		}(i)
	}

	wg.Wait()

	// Verify all entries
	expected.Range(func(k, v interface{}) bool {
		key := k.(*tests.SortableNode)
		got, ok := m.Get(key)
		if !ok {
			t.Errorf("Key %v not found after concurrent inserts", key)
			return true
		}
		if got != v.(*tests.SortableNode) {
			t.Errorf("Expected value %v, got %v for key %v", v, got, key)
		}
		return true
	})

	// Verify length
	if m.Len() != numGoroutines*nodesPerGoroutine {
		t.Errorf("Expected %d entries, got %d", numGoroutines*nodesPerGoroutine, m.Len())
	}
}

// TestConcComparable_ConcurrentReadWrite tests concurrent reads and writes
func TestConcComparable_ConcurrentReadWrite(t *testing.T) {
	type N = *tests.SortableNode

	m := ConcComparable[N, N]()

	// Pre-populate the map
	baseKey := tests.NewSortableNode("base-key")
	baseValue := tests.NewSortableNode("base-value")
	m.Put(baseKey, baseValue)

	var wg sync.WaitGroup
	numReaders := 50
	numWriters := 10
	iterations := 100

	// start readers
	for i := 0; i < numReaders; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < iterations; j++ {
				val, ok := m.Get(baseKey)
				if !ok {
					t.Errorf("Reader %d: base key not found at iteration %d", id, j)
					return
				}
				if val.ID() != baseValue.ID() {
					t.Errorf("Reader %d: expected %v, got %v", id, baseValue, val)
					return
				}
				// Small sleep to increase interleaving
				time.Sleep(time.Microsecond)
			}
		}(i)
	}

	// start writers
	for i := 0; i < numWriters; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < iterations; j++ {
				key := tests.NewSortableNode(fmt.Sprintf("writer-%d-key-%d", id, j))
				value := tests.NewSortableNode(fmt.Sprintf("writer-%d-value-%d", id, j))
				m.Put(key, value)
				time.Sleep(time.Microsecond)
			}
		}(i)
	}

	wg.Wait()
}

// TestConcComparable_ConcurrentDelete tests concurrent delete operations
func TestConcComparable_ConcurrentDelete(t *testing.T) {
	type N = *tests.SortableNode

	m := ConcComparable[N, N]()

	// Pre-populate
	numEntries := 100
	keys := make([]N, numEntries)
	for i := 0; i < numEntries; i++ {
		keys[i] = tests.NewSortableNode(fmt.Sprintf("key-%d", i))
		m.Put(keys[i], keys[i])
	}

	var wg sync.WaitGroup
	numDeleters := 10
	deletesPerGoroutine := numEntries / numDeleters

	// Concurrent deletes
	for i := 0; i < numDeleters; i++ {
		wg.Add(1)
		go func(start int) {
			defer wg.Done()
			for j := start; j < start+deletesPerGoroutine && j < len(keys); j++ {
				m.Delete(keys[j])
			}
		}(i * deletesPerGoroutine)
	}

	wg.Wait()

	// Verify all expected keys are deleted
	for i := 0; i < numEntries; i++ {
		if i < numDeleters*deletesPerGoroutine {
			if m.Contains(keys[i]) {
				t.Errorf("Key %v should have been deleted", keys[i])
			}
		}
	}

}

// TestConcComparable_RaceConditions tests for race conditions using Go's race detector
func TestConcComparable_RaceConditions(t *testing.T) {
	type N = *tests.SortableNode

	m := ConcComparable[N, N]()

	var wg sync.WaitGroup
	numGoroutines := 20
	iterations := 50

	// Mixed operations to trigger race conditions
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < iterations; j++ {
				key := tests.NewSortableNode(fmt.Sprintf("shared-key-%d", j%10)) // Shared keys

				switch j % 4 {
				case 0:
					m.Put(key, key)
				case 1:
					m.Get(key)
				case 2:
					m.Contains(key)
				case 3:
					if j%10 == 0 {
						m.Delete(key)
					}
				}
			}
		}(i)
	}

	wg.Wait()
}

// TestConcComparable_ConcurrentClear tests concurrent Clear with other operations
func TestConcComparable_ConcurrentClear(t *testing.T) {
	type N = *tests.SortableNode

	m := ConcComparable[N, N]()

	// Pre-populate
	for i := 0; i < 20; i++ {
		key := tests.NewSortableNode(fmt.Sprintf("key-%d", i))
		m.Put(key, key)
	}

	var wg sync.WaitGroup

	// Writer goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 50; i++ {
			key := tests.NewSortableNode(fmt.Sprintf("writer-key-%d", i))
			m.Put(key, key)
			time.Sleep(time.Millisecond)
		}
	}()

	// Reader goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			m.Len()
			m.Keys()
			m.Values()
			time.Sleep(time.Millisecond)
		}
	}()

	// Clear goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 3; i++ {
			time.Sleep(10 * time.Millisecond)
			m.Clear()
		}
	}()

	wg.Wait()

}

// TestConcComparable_ConcurrentKeysValues tests Keys() and Values() during modifications
func TestConcComparable_ConcurrentKeysValues(t *testing.T) {
	type N = *tests.SortableNode

	m := ConcComparable[N, N]()

	// Pre-populate
	for i := 0; i < 10; i++ {
		key := tests.NewSortableNode(fmt.Sprintf("key-%d", i))
		m.Put(key, key)
	}

	var wg sync.WaitGroup
	errors := make(chan string, 100)

	// Continuous modifications
	done := make(chan bool)
	go func() {
		for i := 0; i < 100; i++ {
			select {
			case <-done:
				return
			default:
				key := tests.NewSortableNode(fmt.Sprintf("dynamic-key-%d", i))
				m.Put(key, key)
				if i%2 == 0 {
					m.Delete(key)
				}
			}
		}
	}()

	// Concurrent Keys() and Values() calls
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				keys := m.Keys()
				values := m.Values()
				// Keys and Values might not match in length due to concurrent modifications
				// But neither should panic
				if len(keys) != len(values) {
					// This is expected during concurrent modifications
					// Just verify no panic occurred
					select {
					case errors <- fmt.Sprintf("Length mismatch: keys=%d, values=%d", len(keys), len(values)):
					default:
					}
				}
			}
		}()
	}

	wg.Wait()
	close(done)
	close(errors)

	// Log any length mismatches (these might be expected)
	for err := range errors {
		t.Logf("Note: %s", err)
	}

}

// TestConcComparable_StressTest performs a stress test with many goroutines
func TestConcComparable_StressTest(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping stress test in short mode")
	}

	type N = *tests.SortableNode

	m := ConcComparable[N, N]()

	var wg sync.WaitGroup
	numGoroutines := 500
	duration := 1 * time.Second

	done := make(chan bool)

	// Launch many goroutines doing mixed operations
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			localKeys := make([]N, 0)

			for {
				select {
				case <-done:
					// Cleanup: delete our keys
					for _, k := range localKeys {
						m.Delete(k)
					}
					return
				default:
					key := tests.NewSortableNode(fmt.Sprintf("stress-%d-%d", id, len(localKeys)))
					value := tests.NewSortableNode(fmt.Sprintf("stress-value-%d", id))

					m.Put(key, value)
					localKeys = append(localKeys, key)

					// Random read
					if len(localKeys) > 0 {
						randomIdx := len(localKeys) - 1
						m.Get(localKeys[randomIdx])
						m.Contains(localKeys[randomIdx])
					}

					m.Len()
				}
			}
		}(i)
	}

	// Let it run for the duration
	time.Sleep(duration)
	close(done)
	wg.Wait()
}

// TestConcComparable_ConcurrentDuplicateKeys tests concurrent writes with same keys
func TestConcComparable_ConcurrentDuplicateKeys(t *testing.T) {
	type N = *tests.SortableNode

	m := ConcComparable[N, N]()

	sharedKey := tests.NewSortableNode("shared-key")

	var wg sync.WaitGroup
	numGoroutines := 20

	// All goroutines write to the same key
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			value := tests.NewSortableNode(fmt.Sprintf("value-%d", id))
			m.Put(sharedKey, value)
		}(i)
	}

	wg.Wait()

	// Should have exactly one entry for the shared key
	count := 0
	for _, k := range m.Keys() {
		if k.ID() == sharedKey.ID() {
			count++
		}
	}

	if count != 1 {
		t.Errorf("Expected 1 entry for shared key, got %d", count)
	}

	if m.Len() != 1 {
		t.Errorf("Expected map length 1, got %d", m.Len())
	}

}

func TestConcComparable_Iterator(t *testing.T) {

	m := ConcComparable(
		NewAnyEntry("c", "value-c"),
		NewAnyEntry("a", "value-a"),
		NewAnyEntry("b", "value-b"),
	)

	iter := m.Iterator()

	found := make(map[string]string)

	for {
		key, value, ok := iter.Next()
		if !ok {
			break
		}

		found[key] = value
	}

	expected := map[string]string{
		"a": "value-a",
		"b": "value-b",
		"c": "value-c",
	}

	if len(found) != len(expected) {
		t.Errorf("Expected %d entries, got %d", len(expected), len(found))
	}

	for key, expectedValue := range expected {
		value, ok := found[key]

		if !ok {
			t.Errorf("Missing key %q", key)
			continue
		}

		if value != expectedValue {
			t.Errorf("Expected %q for key %q, got %q",
				expectedValue,
				key,
				value,
			)
		}
	}
}

func TestConcComparable_IteratorSorted(t *testing.T) {

	m := ConcComparable(
		NewAnyEntry("c", 3),
		NewAnyEntry("a", 1),
		NewAnyEntry("b", 2),
	)

	iter := m.Iterator(comparators.Ordered[string]())

	var keys []string

	for {
		key, _, ok := iter.Next()
		if !ok {
			break
		}

		keys = append(keys, key)
	}

	expected := []string{"a", "b", "c"}

	if !slices.Equal(keys, expected) {
		t.Errorf("Expected sorted keys %v, got %v", expected, keys)
	}
}

func TestConcComparable_IteratorDoesNotDeadlock(t *testing.T) {

	m := ConcComparable(
		NewAnyEntry("c", 3),
		NewAnyEntry("a", 1),
		NewAnyEntry("b", 2),
	)

	done := make(chan struct{})

	go func() {
		defer close(done)

		iter := m.Iterator()

		for {
			_, _, ok := iter.Next()
			if !ok {
				break
			}
		}
	}()

	select {
	case <-done:
	case <-time.After(time.Second):
		t.Errorf("Iterator appears to be deadlocked")
	}
}

func TestConcComparable_IteratorConcurrentAccess(t *testing.T) {

	m := ConcComparable[string, int]()

	for i := 0; i < 100; i++ {
		m.Put(fmt.Sprintf("key-%d", i), i)
	}

	done := make(chan struct{})

	go func() {
		defer close(done)

		for i := 0; i < 100; i++ {
			iter := m.Iterator()

			for {
				_, _, ok := iter.Next()
				if !ok {
					break
				}
			}
		}
	}()

	for i := 0; i < 100; i++ {
		m.Put(fmt.Sprintf("new-key-%d", i), i)
	}

	select {
	case <-done:
		// terminou
	case <-time.After(time.Second):
		t.Errorf("Concurrent iterator access deadlocked")
	}
}
