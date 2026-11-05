// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package events

type Schedulable interface {
	Schedule(func())
}

func AsyncScheduler() Schedulable {
	return asyncScheduler{}
}

type asyncScheduler struct{}

func (asyncScheduler) Schedule(f func()) {
	go f()
}

func InlineScheduler() Schedulable {
	return inlineScheduler{}
}

type inlineScheduler struct{}

func (inlineScheduler) Schedule(f func()) {
	f()
}
