// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package datastr

import "github.com/andrerrcosta2/gtools/core/data"

func NewBranch[K comparable](index int) *Branch[K] {
	return &Branch[K]{
		index:    index,
		sequence: make([]K, 0),
	}
}

type Branch[K comparable] struct {
	index    int
	sequence []K
}

type SequentialBranch[B data.Branchable[B], K any] struct {
	index    int
	sequence []K
}
