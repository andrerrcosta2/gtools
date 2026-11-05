// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package task

import (
	"github.com/andrerrcosta2/gtools/core/domain/constraints"
	"github.com/andrerrcosta2/gtools/core/domain/constraints/prim"
	"github.com/andrerrcosta2/gtools/core/domain/contracts/timing/chrono"
	"github.com/andrerrcosta2/gtools/core/domain/contracts/timing/schedulable"
)

type Scheduler[R Runnable, I any] interface {
	At(R, I) error
	Every(R, I) error
}

type Runnable interface {
	Run()
}

type Schedulable[T constraints.Sortable[T]] interface {
	Runnable
	schedulable.Event[chrono.Time]
}

type Priority[P prim.Ordered] interface {
	Priority() P
}

type Releasable interface {
	Runnable
	Release()
}

type Branchable[T Runnable] interface {
	Branch() (T, bool)
	AddChild(t T)
	Children() []T
}
