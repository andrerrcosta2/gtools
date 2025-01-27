// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package str

type Set[T any] interface {
	Has(T) bool
	Add(T)
	Remove(T)
	Len() int
	Values() []T
	Clear()
	Equals(Set[T]) bool
}
