// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package ngb

import (
	"github.com/andrerrcosta2/gtools/core/data/str/dict"
	"github.com/andrerrcosta2/gtools/core/domain/constraints"
	"github.com/andrerrcosta2/gtools/datastr/dicts"
)

type StructDict[N constraints.Comparable[N]] dicts.Struct[N, ComparableOf[N]]

type WStructDict[N constraints.Comparable[N], W any] dicts.Struct[N, WComparableOf[N, W]]

type CmpDict[N comparable] map[N]Comparable[N]

type WCmpDict[N comparable, W any] dict.Mutable[N, WComparable[N, W]]
