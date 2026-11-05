// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package fuzz

import "github.com/andrerrcosta2/gtools/core/data/str/iterables"

type (
	Pack struct{}
)

func (c Pack) Refs() *iterables.Slice[any] {
	return iterables.OfSlice(referencesSet()...)
}

func (c Pack) Values() *iterables.Slice[any] {
	return iterables.OfSlice(valuesSet()...)
}
