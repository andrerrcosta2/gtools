// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package tasks

import (
	"github.com/andrerrcosta2/gtools/core/data"
	"github.com/andrerrcosta2/gtools/core/gtools"
)

type Runnable interface {
	Run() error
}

type Releasable interface {
	Runnable
	Release()
}

type ReleasableOf interface {
	gtools.UniqueOf
	Releasable
}

type ReleasableBranch[T Runnable] interface {
	ReleasableOf
	Branchable[T]
}

type ReleasableBranchOf[T Runnable] interface {
	ReleasableOf
	Branchable[T]
}

type Chain[T Runnable] interface {
	Runnable
	Append(tasks ...T)
	Insert(i int, task T) error
	Prepend(tasks ...T)
	Remove(i int)
}

type Linked[T Runnable] interface {
	Next() T
}

type Branchable[T Runnable] interface {
	data.Branchable[T]
	AddChild(t T)
	Children() []T
}
