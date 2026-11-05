// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package syncs

type Barrier interface {
	Waiter
	Completer
}

type ThresholdBarrier interface {
	Wait()
	Count() int
	Threshold() int
}
