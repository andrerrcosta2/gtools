// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sets

type Set[T any] interface {
	Has(T) bool
	Add(T)
	Remove(T)
	Len() int
	Values() []T
	Clear()
	Equals(Set[T]) bool
}

func Map[T any, S comparable](arr []T, f func(v T) S) Set[S] {
	s := Comparable[S]()
	for _, v := range arr {
		s.Add(f(v))
	}
	return s
}

func NewString(arr ...string) Set[string] {
	return Map(arr, func(v string) string {
		return v
	})
}
