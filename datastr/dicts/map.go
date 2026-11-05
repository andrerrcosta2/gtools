// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package dicts

import (
	"github.com/andrerrcosta2/gtools/core/data/comparator"
	"github.com/andrerrcosta2/gtools/core/data/str/dict"
)

type Struct[K any, V any] interface {
	dict.Mutable[K, V]
	Iterator(...comparator.EqComparer[K]) dict.MapIterator[K, V]
}

var _ dict.Mutable[any, any] = (Struct[any, any])(nil)
