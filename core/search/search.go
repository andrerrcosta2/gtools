// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package search

type Search[S ~[]E, E any] interface {
	Search(S, E) (int, bool)
}
