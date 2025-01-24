// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package testingtools

import (
	"github.com/andrerrcosta2/gtools/core/gtools"
)

type gtoolsRandSortable struct{}

func (g *gtoolsRandSortable) Aggregable(amount int) []gtools.AggregableOf {
	return randUniqueBytesOf(amount, seedAggregableOf)
}

func (g *gtoolsRandSortable) Comparable(amount int) []gtools.ComparableOf {
	return randUniqueBytesOf(amount, seedComparableOf)
}

func (g *gtoolsRandSortable) PersistentComparable(amount int) []gtools.PersistentComparableOf {
	return randUniqueBytesOf(amount, seedPersistentComparableOf)
}

func (g *gtoolsRandSortable) PersistentSortable(amount int) []gtools.PersistentSortableOf {
	return randUniqueBytesOf(amount, seedPersistentSortableOf)
}

func (g *gtoolsRandSortable) Sortable(amount int) []gtools.SortableOf {
	return randUniqueBytesOf(amount, seedSortableOf)
}

func (g *gtoolsRandSortable) Unique(amount int) []gtools.UniqueOf {
	return randUniqueBytesOf(amount, seedUniqueOf)
}
