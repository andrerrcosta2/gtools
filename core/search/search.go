// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package search

type Search[T any] interface {
	Search([]T, T) int
}
