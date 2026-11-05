package mailboxes

import "testing"

func BenchmarkBuffered_Create(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Buffered(&mockBuffer[int]{})
	}
}

func BenchmarkBuffered_Operations(b *testing.B) {
	mb := Buffered(&mockBuffer[int]{})

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		mb.Send(1)
		mb.Receive()
	}
}

func BenchmarkBuffered_Close(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mb := Buffered(&mockBuffer[int]{})

		mb.Close()
	}
}

func BenchmarkBuffered_Seal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mb := Buffered(&mockBuffer[int]{})

		wait := mb.Shutdown()
		wait.Wait()
	}
}

func BenchmarkBuffered_Contention(b *testing.B) {
	mb := Buffered(&mockBuffer[int]{})

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			mb.Send(1)
			mb.Receive()
		}
	})
}
