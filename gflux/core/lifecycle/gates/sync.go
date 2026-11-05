// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package gates

import (
	"sync/atomic"

	"github.com/andrerrcosta2/gtools/gflux/core/domain/lifecycle/gate"
)

func Atomic() gate.Closeable {
	return &atomicGate{}
}

type atomicGate struct {
	closed atomic.Bool
}

func (g *atomicGate) Closed() bool {
	return g.closed.Load()
}

func (g *atomicGate) Close() {
	g.closed.Store(true)
}
