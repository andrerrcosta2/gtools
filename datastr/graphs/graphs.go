// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package graphs

import (
	"github.com/andrerrcosta2/gtools/core/domain/constraints"
	"github.com/andrerrcosta2/gtools/datastr/graphs/internal/ngb"
)

type structGraph[N constraints.Comparable[N]] interface {
	adj() ngb.StructDict[N]
}

type wStructGraph[N constraints.Comparable[N], W any] interface {
	adj() ngb.WStructDict[N, W]
}

type cmpGraph[N comparable] interface {
	adj() ngb.CmpDict[N]
}

type wCmpGraph[N comparable, W any] interface {
	adj() ngb.WCmpDict[N, W]
}
