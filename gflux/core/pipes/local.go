// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package pipes

func semaph(cap int) *semaphore {
	return &semaphore{ch: make(chan int, cap)}
}

type semaphore struct {
	ch chan int
}

func (s semaphore) Acq() {
	s.ch <- 1
}

func (s semaphore) Rls() {
	<-s.ch
}

func (s semaphore) Cap() int {
	return cap(s.ch)
}

func (s semaphore) Rem() int {
	return len(s.ch)
}
