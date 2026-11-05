// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package tests

type SortableStub struct{}

func (s SortableStub) Equal(other SortableStub) bool {
	return false
}

func (s SortableStub) Less(other SortableStub) bool {
	return false
}
