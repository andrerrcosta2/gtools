// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package fuzz

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

func (c CatRefs) All() []any {
	return referencesSet()
}

func (c CatRefs) Primitives() []any {
	return primitiveRefs()
}

func (c CatRefs) PrimitiveWithPtrKeys() []any {
	return primitiveWithPtrKeyRefs()
}

func (c CatRefs) PrimitiveWithPtrValues() []any {
	return primitiveWithPtrValueRefs()
}

func (c CatRefs) PrimitiveWithPtrKeysAndValues() []any {
	return primitiveWithPtrKeyAndValueRefs()
}

func (c CatRefs) Edges() []any {
	return edgeRefs()
}

func (c CatValues) All() []any {
	return valuesSet()
}

func (c CatValues) Primitives() []any {
	return primitiveValues()
}

func (c CatValues) PrimitiveWithPtrKeys() []any {
	return primitiveWithPtrKeyValues()
}

func (c CatValues) PrimitiveWithPtrValues() []any {
	return primitiveWithPtrValueValues()
}

func (c CatValues) PrimitiveWithPtrKeysAndValues() []any {
	return primitiveWithPtrKeyAndValueValues()
}

func (c CatValues) Edges() []any {
	return edgeValues()
}
