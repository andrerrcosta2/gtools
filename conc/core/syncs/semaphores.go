// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package syncs

type Limiter interface {
	Acq()
	Rls()
}

type Semaphore interface {
	Acq()
	Rls()
	Cap() int
	Len() int
}

type RWSemaphore interface {
	StartR()
	EndR()
	StartW()
	EndW()
	Len() (writers int, readers int)
	Cap() (writers int, readers int)
}
