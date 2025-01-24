// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package runners

import (
	"github.com/andrerrcosta2/gtools/core/gtools/constraints/prim"
	"github.com/andrerrcosta2/gtools/core/gtools/functions"
	"github.com/andrerrcosta2/gtools/gflux/core/tasks"
)

type Serial[S prim.Serializable] interface {
	Run(runnable functions.Runnable, order S)
}

type Task[T tasks.Runnable] interface {
	Run(runnable T)
}

type Releaser[T comparable] interface {
	Release(release T) error
}

type StepReleaser[S comparable] interface {
	Releaser[S]
	Step(step S, fn functions.Supplier[error])
}
