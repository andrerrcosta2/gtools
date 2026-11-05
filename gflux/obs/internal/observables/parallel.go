// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package observables

type parallel[T any] struct {
	rails []Observable[T]
}

//func Parallel[T any](src Observable[T], n int) parallel[T] {
//
//}
