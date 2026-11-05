// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package future

import "github.com/andrerrcosta2/gtools/conc/core/syncs"

type Supplier[T any] func() (T, error)

type Future[T any] interface {
	syncs.Awaitable
	Try() (T, error, bool)
}

type Promise[T any] interface {
	Resolve(T)
	Reject(error)
}
