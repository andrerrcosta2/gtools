// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package ngb

import (
	"github.com/andrerrcosta2/gtools/core/data/str/set"
	"github.com/andrerrcosta2/gtools/core/domain/constraints"
)

type Comparable[N comparable] map[N]struct{}

func (n Comparable[N]) Add(neighbors ...N) {
	for _, neighbor := range neighbors {
		n[neighbor] = struct{}{}
	}
}

func (n Comparable[N]) IsEmpty() bool {
	return len(n) == 0
}

func (n Comparable[N]) Equals(other set.Readable[N]) bool {
	if n.Len() != other.Len() {
		return false
	}
	allMatch := true
	other.Range(func(value N) bool {
		if _, ok := n[value]; !ok {
			allMatch = false
		}
		return allMatch
	})
	return allMatch
}

func (n Comparable[N]) Has(value N) bool {
	_, ok := n[value]
	return ok
}

func (n Comparable[N]) Len() int {
	return len(n)
}

func (n Comparable[N]) Range(callback func(value N) bool) {
	for value := range n {
		if !callback(value) {
			return
		}
	}
}

func (n Comparable[N]) Elements() []N {
	values := make([]N, 0, len(n))
	for n := range n {
		values = append(values, n)
	}
	return values
}

type ComparableOf[N constraints.Comparable[N]] set.Mutable[N]
