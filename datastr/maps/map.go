// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package maps

import (
	"github.com/andrerrcosta2/gtools/core/data/comparators"
	"github.com/andrerrcosta2/gtools/core/data/str"
)

type StructMap[K any, V any] interface {
	str.Map[K, V]
	Iterator(...comparators.Functional[K]) str.MapIterator[K, V]
}

var _ str.Map[any, any] = (StructMap[any, any])(nil)
