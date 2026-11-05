// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package zero

import "github.com/andrerrcosta2/gtools/core/data/str/iterables"

type (
	Pack      struct{}
	CatRefs   struct{}
	CatValues struct{}
)

func (p Pack) Refs() CatRefs {
	return CatRefs{}
}

func (p Pack) Values() CatValues {
	return CatValues{}
}

func (c CatRefs) All() *iterables.Slice[any] {
	return iterables.OfSlice(referencesSet()...)
}

func (c CatRefs) Primitives() *iterables.Slice[any] {
	return iterables.OfSlice(primitiveRefs()...)
}

func (c CatRefs) PrimitiveWithPtrKeys() *iterables.Slice[any] {
	return iterables.OfSlice(primitiveWithPtrKeyRefs()...)
}

func (c CatRefs) PrimitiveWithPtrValues() *iterables.Slice[any] {
	return iterables.OfSlice(primitiveWithPtrValueRefs()...)
}

func (c CatRefs) PrimitiveWithPtrKeysAndValues() *iterables.Slice[any] {
	return iterables.OfSlice(primitiveWithPtrKeyAndValueRefs()...)
}

func (c CatRefs) Edges() *iterables.Slice[any] {
	return iterables.OfSlice(edgeRefs()...)
}

func (c CatValues) All() *iterables.Slice[any] {
	return iterables.OfSlice(valuesSet()...)
}

func (c CatValues) Primitives() *iterables.Slice[any] {
	return iterables.OfSlice(primitiveValues()...)
}

func (c CatValues) PrimitiveWithPtrKeys() *iterables.Slice[any] {
	return iterables.OfSlice(primitiveWithPtrKeyValues()...)
}

func (c CatValues) PrimitiveWithPtrValues() *iterables.Slice[any] {
	return iterables.OfSlice(primitiveWithPtrValueValues()...)
}

func (c CatValues) PrimitiveWithPtrKeysAndValues() *iterables.Slice[any] {
	return iterables.OfSlice(primitiveWithPtrKeyAndValueValues()...)
}

func (c CatValues) Edges() *iterables.Slice[any] {
	return iterables.OfSlice(edgeValues()...)
}
