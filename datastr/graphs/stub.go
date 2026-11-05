// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package graphs

type sortableStub struct{}

func (s sortableStub) Equal(other sortableStub) bool {
	return false
}

func (s sortableStub) Less(other sortableStub) bool {
	return false
}
