// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package maps

import (
	"github.com/andrerrcosta2/gtools/pkg/comparables"
	"github.com/andrerrcosta2/gtools/pkg/datastr/iterables"
)

type StructMap[K any, V any] interface {
	Put(key K, value V)
	Get(key K) (V, bool)
	Delete(key K)
	Contains(key K) bool
	Len() int
	Clear()
	Keys() []K
	Values() []V
	Iterator(...comparables.FunctionalComparator[K]) iterables.MapIterator[K, V]
}
