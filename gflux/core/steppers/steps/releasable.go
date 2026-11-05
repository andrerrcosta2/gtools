// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package steps

import (
	"errors"

	"github.com/andrerrcosta2/gtools/core/domain/constraints/prim"
	"github.com/andrerrcosta2/gtools/core/domain/function"
	"github.com/andrerrcosta2/gtools/gflux/core/domain/stepper/step"
)

func Releaser[O prim.Ordered](fn function.Supplier[error], at O) step.Releasable[O] {
	return &releasable[O]{
		fn: fn,
		at: at,
	}
}

// releasable implements step.Step and represents a single step in the chain.
type releasable[O prim.Ordered] struct {
	released bool
	fn       function.Supplier[error]
	at       O // registration order
}

// At returns the registration index of the step
func (n *releasable[O]) At() O { return n.at }

// Exec executes the step if it is released, otherwise returns an error
func (n *releasable[O]) Exec() error {
	if !n.released {
		return errors.New("step not released")
	}
	return n.fn()
}

// Release marks the step as released
func (n *releasable[O]) Release() {
	n.released = true
}

func (n *releasable[O]) Released() bool { return n.released }
