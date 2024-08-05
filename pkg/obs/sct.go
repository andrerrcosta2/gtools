package obs

type Sct[T any] interface {
	Sub(obv *Obv[T]) (*Sub[T], error)
}
