// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package streams

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"
)

func BenchmarkStreamNoCheck(b *testing.B) {
	for i := 0; i < b.N; i++ {
		values := generateTestValues(1000)
		ch := make(chan string, 10)
		go func() {
			defer close(ch)
			for _, v := range values {
				ch <- v
				time.Sleep(10 * time.Millisecond)
			}
		}()

		for v := range ch {
			fmt.Print(v)
		}
	}
}

func BenchmarkStreamWithCheck(b *testing.B) {
	for i := 0; i < b.N; i++ {
		values := generateTestValues(1000)
		var closed atomic.Bool
		ch := make(chan string, 10)
		go func() {
			defer close(ch)
			for _, v := range values {
				if closed.Load() {
					return
				}
				ch <- v
				time.Sleep(10 * time.Millisecond)
			}
		}()

		for {
			select {
			case v, ok := <-ch:
				if !ok {
					return
				}
				fmt.Print(v)
			default:
				if closed.Load() {
					return
				}
				time.Sleep(10 * time.Millisecond)
			}
		}
	}
}

func generateTestValues(n int) []string {
	values := make([]string, n)
	for i := 0; i < n; i++ {
		values[i] = string(rune('a' + (i % 26)))
	}
	return values
}
