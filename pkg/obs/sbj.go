package obs

type Sbj[T any] interface {
	Sub(*Obv[T]) (*Sub[T], error)
	Nxt(T)
	Rmo(*Sub[T])
	Err(error)
	Cpt()
	Uns()
	Cld() bool
}
