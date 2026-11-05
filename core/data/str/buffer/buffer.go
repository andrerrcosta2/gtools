// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package buffer

type Ticket uint64

type Typed[T any] interface {
	Cap() int
	Len() int
}

type Readable[T any] interface {
	Typed[T]
	Read() (T, bool)
}

type Writable[T any] interface {
	Typed[T]
	Write(T) bool
}

type Mutable[T any] interface {
	Readable[T]
	Writable[T]
}

type MPSC[T any] interface {
	Readable[T]
	Writable[T]
	Empty() bool
	Full() bool
}
