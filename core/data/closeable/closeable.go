// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package closeable

type Terminal interface {
	Close() error
}

type Factory[T Terminal] interface {
	Closed() T
}

type State = uint32

const (
	Open State = iota
	Closed
)

func IsOpen(s State) bool {
	return s == Open
}

func IsClosed(s State) bool {
	return s == Closed
}

func CloseIfIsNot(state *State) bool {
	if *state != Closed {
		*state = Closed
		return true
	}
	return false
}
