// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package iterables

type Deliverer[T any] interface {
	Deliver(T) []T
}
