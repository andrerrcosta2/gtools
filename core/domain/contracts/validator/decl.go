// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package validator

type Typed[T any] interface {
	Validate(T) error
}
