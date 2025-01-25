// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package tries

import (
	"context"
	"fmt"
	"github.com/andrerrcosta2/gtools/conc/syncs/semaph"
	"github.com/andrerrcosta2/gtools/core/data/str"
	"github.com/andrerrcosta2/gtools/core/domain/gtools"
	"github.com/andrerrcosta2/gtools/patterns/symbols"
	"sync"
)

// OfPatternsFrom builds a OfPatterns from the given dictionary
//
// It supports both building concurrently using the given semaphore or sequentially.
// If the semaphore is nil or its capacity is 0, the newPatternTrie is built sequentially.
// If the strict flag is true, the newPatternTrie is built in strict mode, meaning it will not
// allow any invalid patterns to be inserted.
func OfPatternsFrom(dictionary str.Dictionary[string, symbols.Logical], strict bool, maxTrieParallelOps, buildInParallel int) (Pattern, error) {
	if buildInParallel > 1 {
		// Build the newPatternTrie concurrently
		if strict {
			return buildConcurrently(newStrictPatternTrie(maxTrieParallelOps), dictionary, semaph.Channel(buildInParallel))
		}
		return buildConcurrently(newPatternTrie(maxTrieParallelOps), dictionary, semaph.Channel(buildInParallel))
	}

	// Build the newPatternTrie sequentially
	t := OfPatterns(strict, maxTrieParallelOps)

	for pattern, symbol := range dictionary.EntrySet() {
		err := t.Insert(pattern, symbol)
		if err != nil {
			return nil, fmt.Errorf("could not build trie from dictionary: %w", err)
		}
	}

	return t, nil
}

// buildConcurrently builds a newPatternTrie from the given dictionary concurrently.
// It returns the final newPatternTrie and an error if it occurs.
// It cancels new insertions of patterns if any error occurs.
func buildConcurrently(trie *patternTrie, dictionary str.Dictionary[string, symbols.Logical], sem gtools.Semaphore) (*patternTrie, error) {
	// Create a cancellable context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Create a wait group to wait for all goroutines to finish
	var wg sync.WaitGroup
	// Create an error channel to receive errors from the goroutines
	errCh := make(chan error, 1) // Buffer size 1 to ensure non-blocking behavior

	// Addf the size of the dictionary to the wait group
	wg.Add(dictionary.Size())

	// Iterate over the dictionary and create a goroutine for each pattern
	for pattern, symbol := range dictionary.EntrySet() {

		go func(pattern string, symbol symbols.Logical) {
			// If an error has already occurred, return early
			select {
			case <-ctx.Done():
				return
			default:
				// Continue if no error has occurred
			}

			// Acquire the semaphore
			sem.Acq()

			// Insert the pattern into the newPatternTrie
			if err := trie.Insert(pattern, symbol); err != nil {
				// Send the error to the channel and cancel the context
				select {
				case errCh <- fmt.Errorf("could not build newPatternTrie from dictionary: %w", err):
					cancel() // Cancel all remaining goroutines
				default:
					// Ignore if an error is already sent
				}
			}

			// Release the semaphore
			sem.Rls()

			// Mark the goroutine as done
			wg.Done()
		}(pattern, symbol)
	}

	// SetWaitingPoint for all goroutines to finish and close the error channel
	go func() {
		wg.Wait()
		close(errCh)
	}()

	// Check if any errors occurred
	if err := <-errCh; err != nil {
		return nil, err
	}

	// Return the final newPatternTrie
	return trie, nil
}
