// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sealable

import (
	"sync/atomic"

	"github.com/andrerrcosta2/gtools/conc/core/conc"
	"github.com/andrerrcosta2/gtools/conc/core/syncs"
)

type Terminal interface {
	Shutdown() syncs.Waiter
}

type Gate interface {
	Seal() error
	Unseal() error
}

type Disposable interface {
	Dispose() syncs.Waiter
}

type Factory[T conc.Sealable] interface {
	Closed() T
	Sealed() T
}

type State = uint32

const (
	Open State = iota
	Sealed
	Closed
)

func IsOpen(s State) bool {
	return s == Open
}

func IsOpenAtomic(s *State) bool {
	return atomic.LoadUint32(s) == Open
}

func IsSealed(s State) bool {
	return s == Sealed
}

func IsSealedAtomic(s *State) bool {
	return atomic.LoadUint32(s) == Sealed
}

func IsClosed(s State) bool {
	return s == Closed
}

func IsClosedAtomic(s *State) bool {
	return atomic.LoadUint32(s) == Closed
}

func AtomicSeal(state *State) {
	atomic.StoreUint32(state, Sealed)
}

func AtomicClose(state *State) {
	atomic.StoreUint32(state, Closed)
}

func SealIfOpen(state *State) bool {
	if *state == Open {
		*state = Sealed
		return true
	}
	return false
}

func AtomicSealIfOpen(state *State) bool {
	return atomic.CompareAndSwapUint32(state, Open, Sealed)
}

func CloseIfIsNot(state *State) bool {
	if *state != Closed {
		*state = Closed
		return true
	}
	return false
}

func AtomicCloseIfIsNot(state *State) bool {
	return atomic.SwapUint32(state, Closed) != Closed
}
