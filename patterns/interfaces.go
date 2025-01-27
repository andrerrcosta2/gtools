// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package patterns

type Matcher[T any] interface {
	Match(T) bool
}
