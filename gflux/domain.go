// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package gflux

import (
	"github.com/andrerrcosta2/gtools/gflux/core/tasks"
	"github.com/andrerrcosta2/gtools/gflux/obs"
)

type Obs[T any] obs.Obs[T]

type Sbj[T any] obs.Sbj[T]

type Task tasks.Runnable
