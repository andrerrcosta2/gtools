// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package zero

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

func (c CatRefs) All() []any {
	return referencesSets()
}

func (c CatRefs) Primitives() []any {
	return primitiveRefs()
}

func (c CatRefs) PointerToPrimitives() []any {
	return pointerToPrimitiveValues()
}

func (c CatRefs) Interfaces() []any {
	return interfaceRefs()
}

func (c CatRefs) PointerToInterfaces() []any {
	return pointerToInterfaceRefs()
}

func (c CatRefs) SliceOfInterfaces() []any {
	return sliceOfInterfaceRefs()
}

func (c CatRefs) SliceOfPointersToInterfaces() []any {
	return sliceOfPointerToInterfaceRefs()
}

func (c CatRefs) Structs() []any {
	return structRefs()
}

func (c CatRefs) SliceOfStructs() []any {
	return sliceOfStructRefs()
}

func (c CatRefs) SliceOfPointersToStructs() []any {
	return pointerToSliceOfStructRefs()
}

func (c CatRefs) SliceOfPrimitives() []any {
	return sliceOfPrimitiveRefs()
}

func (c CatRefs) SliceOfPointersToPrimitives() []any {
	return sliceOfPointerToPrimitiveRefs()
}

func (c CatRefs) Senders() []any {
	return senderRefs()
}

func (c CatRefs) Receivers() []any {
	return receiverRefs()
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

func (c CatValues) PointerToPrimitives() []any {
	return pointerToPrimitiveValues()
}

func (c CatValues) Interfaces() []any {
	return interfaceValues()
}

func (c CatValues) PointerToInterfaces() []any {
	return pointerToInterfaceValues()
}

func (c CatValues) SliceOfInterfaces() []any {
	return sliceOfInterfaceValues()
}

func (c CatValues) SliceOfPointersToInterfaces() []any {
	return sliceOfPointerToInterfaceValues()
}

func (c CatValues) Structs() []any {
	return structValues()
}

func (c CatValues) SliceOfStructs() []any {
	return sliceOfStructValues()
}

func (c CatValues) SliceOfPointersToStructs() []any {
	return pointerToSliceOfStructValues()
}

func (c CatValues) SliceOfPrimitives() []any {
	return sliceOfPrimitiveValues()
}

func (c CatValues) SliceOfPointersToPrimitives() []any {
	return sliceOfPointerToPrimitiveValues()
}

func (c CatValues) Senders() []any {
	return senderValues()
}

func (c CatValues) Receivers() []any {
	return receiverValues()
}

func (c CatValues) Edges() []any {
	return edgeValues()
}
