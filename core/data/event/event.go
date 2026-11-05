// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package event

import (
	"github.com/andrerrcosta2/gtools/core/domain/constraints/prim"
	"github.com/andrerrcosta2/gtools/core/domain/constraints/prim/nums"
)

type Versioned[S prim.Hashable] interface {
	Version() S
}

type Routed[R nums.Natural] interface {
	Route() R
}
