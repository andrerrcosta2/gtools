// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package zero

import "github.com/andrerrcosta2/gtools/core/data/str/iterables"

type (
	Pack      struct{}
	CatRefs   struct{}
	CatValues struct{}
)

func (c Pack) Refs() CatRefs {
	return CatRefs{}
}

func (c Pack) Values() CatValues {
	return CatValues{}
}

func (c CatRefs) All() *iterables.Slice[any] {
	return iterables.OfSlice(referencesSet()...)
}

func (c CatRefs) Primitives() *iterables.Slice[any] {
	return iterables.OfSlice(primitiveRefs()...)
}

func (c CatRefs) PrimitiveSlices() *iterables.Slice[any] {
	return iterables.OfSlice(primitiveSliceRefs()...)
}

func (c CatRefs) PrimitiveStructs() *iterables.Slice[any] {
	return iterables.OfSlice(primitiveStructRefs()...)
}

func (c CatRefs) Channels() *iterables.Slice[any] {
	return iterables.OfSlice(channelRefs()...)
}

func (c CatRefs) Maps() *iterables.Slice[any] {
	return iterables.OfSlice(mapRefs()...)
}

func (c CatValues) All() *iterables.Slice[any] {
	return iterables.OfSlice(valuesSet()...)
}

func (c CatValues) Primitives() *iterables.Slice[any] {
	return iterables.OfSlice(primitiveValues()...)
}

func (c CatValues) PrimitiveSlices() *iterables.Slice[any] {
	return iterables.OfSlice(primitiveSliceValues()...)
}

func (c CatValues) PrimitiveStructs() *iterables.Slice[any] {
	return iterables.OfSlice(primitiveStructValues()...)
}

func (c CatValues) Channels() *iterables.Slice[any] {
	return iterables.OfSlice(channelValues()...)
}

func (c CatValues) Maps() *iterables.Slice[any] {
	return iterables.OfSlice(mapValues()...)
}
