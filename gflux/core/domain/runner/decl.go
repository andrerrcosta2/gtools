// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package runner

import (
	"github.com/andrerrcosta2/gtools/core/domain/constraints/prim"
	"github.com/andrerrcosta2/gtools/core/domain/function"
	"github.com/andrerrcosta2/gtools/gflux/core/domain/stepper"
	"github.com/andrerrcosta2/gtools/gflux/core/domain/stepper/step"
	"github.com/andrerrcosta2/gtools/gflux/core/domain/task"
)

type Policy[S step.Step[E], E any] interface {
	Push(S)
	Peek() (S, bool)
	Pop() (S, bool)
	Empty() bool
}

type Serial[S prim.Serializable] interface {
	Run(runnable function.Runnable, order S)
}

type Gate interface {
	// Run attempts to admit a task.
	// Returns false if the gate rejects admission.
	Run(task func(done func())) bool

	// Stop permanently closes the gate for new admissions.
	// Running task may complete.
	Stop()
}

type Task[T task.Runnable] interface {
	Run(runnable T)
}

type Releaser[T comparable] interface {
	Release(release T) error
}

type StepReleaser[S comparable] interface {
	Releaser[S]
	Step(step S, fn function.Supplier[error])
}

type Stepper interface {
	stepper.Stepper
	Run(n int)
	RunUntilIdle()
}
