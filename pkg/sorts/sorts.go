// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sorts

type Sort[T any] interface {
	Sort(arr *[]T)
}

func shouldSort[T any](arr []T) bool {
	return len(arr) > 1
}
