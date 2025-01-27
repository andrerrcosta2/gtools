// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sorters

type Sorter[T any] interface {
	Sort(arr *[]T)
	SortP(arr *[]*T)
}

func shouldSort[T any](arr []T) bool {
	return len(arr) > 1
}
