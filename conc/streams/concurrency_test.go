// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package streams

import (
	"context"
	"github.com/andrerrcosta2/gtools/conc/syncs/semaph"
	"github.com/andrerrcosta2/gtools/core/domain/functions/runnables"
	"github.com/andrerrcosta2/gtools/gtests/testingtools"
	"github.com/andrerrcosta2/gtools/gtests/testingtools/config/testlogs"
	"sync"
	"sync/atomic"
	"testing"
)

func TestConsume_Sync_NoDelay(t *testing.T) {
	// helper
	tt := testingtools.LoggersLite(t, testlogs.OnFailure)

	values := []string{"a", "b", "c", "d", "e", "A", "B", "C", "D", "E", "f", "g", "h", "i", "j", "F", "G", "H", "I", "J"}
	stream := make(chan string, 2)

	go func() {
		for _, value := range values {
			stream <- value
			tt.StackLogf("sent %s", value)
		}
		close(stream)
	}()

	func() {
		for {
			value, ok := <-stream
			if !ok {
				break
			}
			tt.StackLogf("received %s", value)
		}
	}()

	tt.PrintLogStack()
}

func TestConsume_Sync_NoDelay_OnTheFlyClosedStream(t *testing.T) {
	// helper
	tt := testingtools.LoggersLite(t, testlogs.OnFailure)

	values := []string{"a", "b", "c", "d", "e", "A", "B", "C", "D", "E", "f", "g", "h", "i", "j", "F", "G", "H", "I", "J"}
	stream := make(chan string, 2)
	var closed atomic.Bool
	var consumes atomic.Int32

	// closeable producer
	go func() {
		for _, value := range values {
			if closed.Load() {
				tt.StackLogf("Closing stream")
				close(stream)
				return
			}
			stream <- value
			tt.StackLogf("sent %s", value)
		}
		close(stream)
	}()

	// synchronous consumer
	func() {
		for {
			value, ok := <-stream
			if !ok {
				tt.StackLogf("value not received")
				return
			}
			go func() {
				if consumes.Add(1) == 3 {
					closed.Store(true)
				}
				tt.StackLogf("received %s", value)
			}()
		}
	}()

	if consumes.Load() >= int32(len(values)) {
		t.Errorf("expected less than %d values, got %d", len(values), consumes.Load())
	} else {
		t.Logf("expected less than %d values, got %d", len(values), consumes.Load())
	}
	tt.StackLogf("Closed? %t\n", closed.Load())
	tt.StackLogf("Consumes: %d\n", consumes.Load())
	tt.PrintLogStack()
}

func TestConsume_Sync_OnTheFlyClosedConsumer(t *testing.T) {
	// helper
	tt := testingtools.LoggersLite(t, testlogs.OnFailure)

	values := []string{"a", "b", "c", "d", "e", "A", "B", "C", "D", "E", "f", "g", "h", "i", "j", "F", "G", "H", "I", "J"}
	stream := HotCloseable[string](3, values...)
	var closed atomic.Bool
	var consumes atomic.Int32

	wg := sync.WaitGroup{}
	semaphore := semaph.SingleRoutine()
	// closeable synchronous consumer
	func() {
		for {
			if closed.Load() {
				tt.StackLogf("Closing stream")
				wg.Wait()
				return
			}
			value, ok := <-stream.Stream()
			if !ok {
				tt.StackLogf("value not received")
				wg.Wait()
				return
			}
			wg.Add(1)

			go runnables.SemaphoredSync(&wg, semaphore, func() {
				if consumes.Add(1) == 3 {
					closed.Store(true)
				}
				tt.StackLogf("received %s", value)
			})
		}
	}()

	if consumes.Load() >= int32(len(values)) {
		t.Errorf("expected less than %d values, got %d", len(values), consumes.Load())
	} else {
		t.Logf("expected less than %d values, got %d", len(values), consumes.Load())
	}
	tt.StackLogf("Closed? %t\n", closed.Load())
	tt.StackLogf("Consumes: %d\n", consumes.Load())
	tt.PrintLogStack()
}

// These tests are dev tests
//func TestConsume_Cancel_Sync(t *testing.T) {
//	// helper
//	tt := testingtools.LoggersLite(t, testlogs.OnFailure)
//
//	values := []string{"a", "b", "c", "d", "e", "A", "B", "C", "D", "E", "f", "g", "h", "i", "j", "F", "G", "H", "I", "J"}
//	stream := make(chan string, 2)
//	var consumes atomic.Int32
//	ctx, cancel := context.WithCancel(context.Background())
//
//	// producer
//	go func() {
//		for _, value := range values {
//			stream <- value
//			tt.StackLogf("sent %s", value)
//		}
//		close(stream)
//	}()
//
//	// consumer
//	func() {
//		for {
//			select {
//			case <-ctx.Done():
//				tt.StackLogf("context canceled")
//				return
//			default:
//				go func() {
//					value, ok := <-stream
//					if !ok {
//						tt.StackLogf("value not received")
//						return
//					}
//					if consumes.Add(1) == 3 {
//						cancel()
//					}
//
//					tt.StackLogf("received %s", value)
//				}()
//			}
//		}
//	}()
//
//	if consumes.Load() >= int32(len(values)) {
//		t.Errorf("expected less than %d values, got %d", len(values), consumes.Load())
//	} else {
//		t.Logf("expected less than %d values, got %d", len(values), consumes.Load())
//	}
//	tt.PrintLogStack()
//}

func TestConsume_Cancel_Async(t *testing.T) {
	// helper
	tt := testingtools.LoggersLite(t, testlogs.OnFailure)

	values := []string{"a", "b", "c", "d", "e", "A", "B", "C", "D", "E", "f", "g", "h", "i", "j", "F", "G", "H", "I", "J"}
	stream := make(chan string, 2)
	var consumes atomic.Int32
	ctx, cancel := context.WithCancel(context.Background())

	runnables.ContextRelease(func() {
		// producer
		go func() {
			for _, value := range values {
				stream <- value
				tt.StackLogf("sent %s", value)
			}
			close(stream)
		}()

		// consumer
		go func() {
			for {
				select {
				case <-ctx.Done():
					tt.StackLogf("context CANCELED!")
					return
				default:
					go func() {
						value, ok := <-stream
						if !ok {
							tt.StackLogf("value not received")
							return
						}
						if consumes.Add(1) == 3 {
							cancel()
						}

						tt.StackLogf("received %s", value)
					}()
				}
			}
		}()

	}, ctx)

	if consumes.Load() >= int32(len(values)) {
		t.Errorf("expected less than %d values, got %d", len(values), consumes.Load())
	} else {
		t.Logf("expected less than %d values, got %d", len(values), consumes.Load())
	}
	tt.PrintLogStack()
}
