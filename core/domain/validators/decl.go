// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package validators

type Typed[T any] interface {
	Validate(T) error
}
