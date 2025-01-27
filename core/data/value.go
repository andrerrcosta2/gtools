// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package data

type Async[V any] interface {
	Value() (V, bool)
	Sync() error
}
