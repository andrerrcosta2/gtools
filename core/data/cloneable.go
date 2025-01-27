// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package data

type Cloneable[T any] interface {
	Clone() T
}

func IsCloneable[T any](data any) bool {
	_, ok := data.(Cloneable[T])
	return ok
}
