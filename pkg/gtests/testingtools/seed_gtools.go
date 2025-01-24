// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package testingtools

import (
	"bytes"
	"github.com/andrerrcosta2/gtools/core/gtools"
)

func seedComparableOf(data []byte) gtools.ComparableOf {
	return &byteComparable{data}
}

type byteComparable struct {
	data []byte
}

func (c *byteComparable) Equal(o any) bool {
	if other, ok := o.(*byteComparable); ok {
		return bytes.Equal(c.data, other.data)
	}
	if other, ok := o.([]byte); ok {
		return bytes.Equal(c.data, other)
	}
	return false
}

var _ gtools.ComparableOf = &byteComparable{}

func seedSortableOf(data []byte) gtools.SortableOf {
	return &byteSortableOf{byteComparable{data}}
}

type byteSortableOf struct {
	byteComparable
}

func (s *byteSortableOf) Less(o any) bool {
	if other, ok := o.(*byteComparable); ok {
		return bytes.Compare(s.data, other.data) < 0
	}
	if other, ok := o.([]byte); ok {
		return bytes.Compare(s.data, other) < 0
	}
	return false
}

var _ gtools.SortableOf = &byteSortableOf{}

func seedUniqueOf(data []byte) gtools.UniqueOf {
	return &byteUniqueOf{data}
}

type byteUniqueOf struct {
	data []byte
}

func (u *byteUniqueOf) Unique() string {
	return string(u.data)
}

var _ gtools.UniqueOf = &byteUniqueOf{}

func seedPersistentSortableOf(data []byte) gtools.PersistentSortableOf {
	return &bytePersistentSortableOf{byteSortableOf{byteComparable{data}}}
}

type bytePersistentSortableOf struct {
	byteSortableOf
}

func (p *bytePersistentSortableOf) Unique() string {
	return string(p.data)
}

var _ gtools.PersistentSortableOf = &bytePersistentSortableOf{}

func seedPersistentComparableOf(data []byte) gtools.PersistentComparableOf {
	return &bytePersistentComparableOf{byteComparable{data}}
}

type bytePersistentComparableOf struct {
	byteComparable
}

func (p *bytePersistentComparableOf) Unique() string {
	return string(p.data)
}

var _ gtools.PersistentComparableOf = &bytePersistentComparableOf{}

func seedAggregableOf(data []byte) gtools.AggregableOf {
	return &byteAggregable{byteComparable{data}}
}

type byteAggregable struct {
	byteComparable
}

func (a *byteAggregable) Add(o gtools.AggregableOf) gtools.AggregableOf {
	if o == nil {
		return a
	}

	if oo, ok := o.(*byteAggregable); ok {
		return &byteAggregable{
			byteComparable: byteComparable{
				data: append(a.data, oo.data...),
			},
		}
	}

	return a
}

var _ gtools.AggregableOf = &byteAggregable{}
