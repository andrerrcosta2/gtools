// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package selection

type Selector[S ~[]E, E any] interface {
	Select(S, int) (E, bool)
}
