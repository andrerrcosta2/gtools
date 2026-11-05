// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package consume

import (
	"errors"
	"log"

	"github.com/andrerrcosta2/gtools/conc/core/selects"
	"github.com/andrerrcosta2/gtools/conc/core/syncs"
	"github.com/andrerrcosta2/gtools/core/domain/contracts/stream"
	"github.com/andrerrcosta2/gtools/core/domain/function"
	"github.com/andrerrcosta2/gtools/core/io"
)

var ClosedConsumer = errors.New("consumer is closed")
var ErrClosedConsumerOnTheFly = errors.New("consumer was found unexpectedly closed on the fly while consume")
var ErrClosed = errors.New("consumer is already closed")

// CloseableSelect is a std function to be used with goroutines.
// It receives a domain.Streamable, a domain.CloseableConsumer, a WaitGroup, a Semaphore, a channel as signal,
// and a function.BiConsumer[int, T] function.
// It checks if the signal is sent, if so it call the closeable and returns false.
// Otherwise, it tries to consume the value from the channel using the BiConsumer function
func CloseableSelect[T any](
	closeable io.Closeable,
	stream stream.Readable[T],
	leave syncs.Blocker,
	wg syncs.WaitGroup,
	semaphore syncs.Semaphore,
	fn function.Consumer[T], counter int) bool {
	// Return false on signal sent
	return selects.Bool(leave, func() bool {
		// Otherwise consume the value from the channel using the BiConsumer function
		return ParallelCloseableStream(closeable, stream, wg, semaphore, fn)
	})
}

// ParallelCloseableStream consumes a closeable stream asynchronously using semaphores and waiting groups
// // to control access and barriers
func ParallelCloseableStream[T any](
	closeable io.Closeable,
	stream stream.Readable[T],
	wg syncs.WaitGroup,
	semaphore syncs.Semaphore,
	fn function.Consumer[T]) bool {
	// If the consumer is closed, wait for all ops to complete before returning false
	if closeable.IsClosed() {
		log.Printf("\nconsumer is closed, exiting...\n")
		wg.Wait()
		return false
	}
	// Try consuming the value from the channel using the BiConsumer function
	return StreamSynchronously[T](stream, wg, semaphore, fn)
}

// StreamSynchronously consumes a stream asynchronously using semaphores and waiting groups
// to control access and barriers
func StreamSynchronously[T any](
	stream stream.Readable[T],
	wg syncs.WaitGroup,
	sem syncs.Semaphore,
	consume function.Consumer[T]) bool {

	stream.Iterator()(func(value T) bool {
		wg.Add(1)
		sem.Acq()
		go func() {
			defer func() {
				sem.Rls()
				wg.Done()
			}()

			consume(value)
		}()
		return true
	})
	return true
}

func Once()
