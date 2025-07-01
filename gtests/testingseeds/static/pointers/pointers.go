// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package pointers

func Nil[T any]() *T {
	return (*T)(nil)
}
