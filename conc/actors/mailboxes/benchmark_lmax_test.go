// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package mailboxes

import (
	"runtime"
	"sync"
	"testing"
)

// BenchmarkMPSC compares performance with different configurations
func BenchmarkMPSC(b *testing.B) {
	benchmarks := []struct {
		name      string
		producers int
		size      uint64
		useDrain  bool
	}{
		{"SingleProducer", 1, 1024, false},
		{"SingleProducer_Drain", 1, 1024, true},
		{"4Producers", 4, 1024, false},
		{"4Producers_Drain", 4, 1024, true},
		{"8Producers", 8, 1024, false},
		{"8Producers_Drain", 8, 1024, true},
		{"SmallBuffer", 4, 64, false},
		{"LargeBuffer", 4, 65536, false},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			q := Disruptor[int](bm.size)

			// Pre-warm
			runtime.GC()

			b.ResetTimer()
			b.SetBytes(8) // 8 bytes per message

			var wg sync.WaitGroup

			// Consumer
			wg.Add(1)
			go func() {
				defer wg.Done()
				received := 0
				if bm.useDrain {
					for received < b.N {
						received += q.Drain(func(int) {})
					}
				} else {
					for received < b.N {
						q.Receive()
						received++
					}
				}
			}()

			// Producers
			msgsPerProducer := b.N / bm.producers
			for p := 0; p < bm.producers; p++ {
				wg.Add(1)
				go func() {
					defer wg.Done()
					for i := 0; i < msgsPerProducer; i++ {
						q.Send(i)
					}
				}()
			}

			wg.Wait()
		})
	}
}

func BenchmarkBuffered(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mb := Buffered(&mockBuffer[int]{})

		mb.Send(1)
		wait := mb.Shutdown()
		mb.Receive()
		wait.Wait()
	}
}
