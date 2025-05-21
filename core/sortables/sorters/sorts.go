// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sorters

type Sorter[T any, S ~[]T] interface {
	Sort(arr *S)
}

func shouldSort[T any](arr []T) bool {
	return len(arr) > 1
}
