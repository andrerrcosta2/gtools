// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package ngb

import (
	"fmt"
	"strings"

	"github.com/andrerrcosta2/gtools/core/data/comparator"
	"github.com/andrerrcosta2/gtools/core/data/str/dict"
	"github.com/andrerrcosta2/gtools/core/data/str/set"
	"github.com/andrerrcosta2/gtools/core/domain/constraints"
	"github.com/andrerrcosta2/gtools/core/sortables/sorters"
	"github.com/andrerrcosta2/gtools/datastr/dicts"
	"github.com/andrerrcosta2/gtools/datastr/internal/tests"
)

type WComparable[N comparable, W any] map[N]W

func (m WComparable[N, W]) Contains(n N) bool {
	return m.Has(n)
}

func (m WComparable[N, W]) Get(n N) (W, bool) {
	v, ok := m[n]
	return v, ok
}

func (m WComparable[N, W]) Iterator(cmp ...comparator.EqComparer[N]) dict.MapIterator[N, W] {
	keys := m.Keys()
	if len(cmp) > 0 {
		sorter := sorters.Quick[[]N](cmp[0])
		sorter.Sort(&keys)
	}
	return &wcmpcomparator[N, W]{}
}

type wcmpcomparator[N comparable, W any] struct {
	m     WComparable[N, W]
	keys  []N
	index int
}

func (w *wcmpcomparator[N, W]) Next() (N, W, bool) {
	if w.index >= len(w.keys) {
		var zeroN N
		var zeroW W
		return zeroN, zeroW, false
	}
	key := w.keys[w.index]
	w.index++
	value := w.m[key]
	return key, value, true
}

func (m WComparable[N, W]) Keys() []N {
	out := make([]N, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	return out
}

func (m WComparable[N, W]) Values() []W {
	out := make([]W, 0, len(m))
	for _, v := range m {
		out = append(out, v)
	}
	return out
}

func (m WComparable[N, W]) Equals(other set.Readable[N]) bool {
	if other.Len() != m.Len() {
		return false
	}
	allMatch := true
	other.Range(func(value N) bool {
		if !m.Has(value) {
			allMatch = false
		}
		return allMatch
	})
	return allMatch
}

func (m WComparable[N, W]) IsEmpty() bool {
	return len(m) == 0
}

func (m WComparable[N, W]) Has(n N) bool {
	_, ok := m[n]
	return ok
}

func (m WComparable[N, W]) Len() int {
	return len(m)
}

func (m WComparable[N, W]) Range(callback func(value N) bool) {
	for key := range m {
		if !callback(key) {
			return
		}
	}
}

func (m WComparable[N, W]) Elements() []N {
	values := make([]N, 0, len(m))
	for key := range m {
		values = append(values, key)
	}
	return values
}

func (m WComparable[N, W]) String() string {
	sb := strings.Builder{}
	sb.WriteString("WComparable {\n")
	for key, value := range m {
		sb.WriteString(fmt.Sprintf("\t%v : %v,\n", key, value))
	}
	sb.WriteString("}\n")
	return sb.String()
}

var _ set.Readable[int] = WComparable[int, float32]{}

func NewWComparableOf[N constraints.Comparable[N], W any]() WComparableOf[N, W] {
	return dicts.ComparableOf[N, W]()
}

type WComparableOf[N constraints.Comparable[N], W any] dict.Mutable[N, W]

func KeySet[N constraints.Comparable[N], W any](dict dict.Mutable[N, W]) set.Readable[N] {
	return &keySet[N, W]{
		dict: dict,
	}
}

type keySet[N constraints.Comparable[N], W any] struct {
	dict dict.Mutable[N, W]
}

func (m *keySet[N, W]) Equals(other set.Readable[N]) bool {
	if other.Len() != m.Len() {
		return false
	}
	allMatch := true
	other.Range(func(value N) bool {
		if !m.Has(value) {
			allMatch = false
		}
		return allMatch
	})
	return allMatch
}

func (m *keySet[N, W]) IsEmpty() bool {
	return m.dict.Len() == 0
}

func (m *keySet[N, W]) Has(n N) bool {
	return m.dict.Contains(n)
}

func (m *keySet[N, W]) Len() int {
	return m.dict.Len()
}

func (m *keySet[N, W]) Range(callback func(value N) bool) {
	it := m.dict.Iterator()
	for key, _, hasNext := it.Next(); hasNext; key, _, hasNext = it.Next() {
		if !callback(key) {
			return
		}
	}
}

func (m *keySet[N, W]) Elements() []N {
	values := make([]N, 0, m.dict.Len())
	it := m.dict.Iterator()
	for key, _, hasNext := it.Next(); hasNext; key, _, hasNext = it.Next() {
		values = append(values, key)
	}
	return values
}

func (m *keySet[N, W]) String() string {
	sb := strings.Builder{}
	sb.WriteString("WComparable {\n")
	it := m.dict.Iterator()
	for key, _, hasNext := it.Next(); hasNext; key, _, hasNext = it.Next() {
		sb.WriteString(fmt.Sprintf("\t%v\n", key))
	}
	sb.WriteString("}\n")
	return sb.String()
}

var _ set.Readable[tests.SortableStub] = (*keySet[tests.SortableStub, string])(nil)
