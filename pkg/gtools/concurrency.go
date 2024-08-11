// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package gtools

type Semaphore struct {
	ch chan struct{}
}

func NewSemaphore(maxConcurrent int) *Semaphore {
	return &Semaphore{
		ch: make(chan struct{}, maxConcurrent),
	}
}

func (s *Semaphore) Acq() {
	s.ch <- struct{}{}
}

func (s *Semaphore) Rls() {
	<-s.ch
}
