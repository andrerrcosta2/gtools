package obs

type Obs[T any] interface {
	Sct[T]
	Nxt(d T)
	Rmo(sub *Sub[T])
}
