// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package arrays

import "github.com/andrerrcosta2/gtools/gtests/testingseeds/static"

func SetOf() static.SeedSets[ArraySets, ArraySets] {
	return &arrays{}
}

type arrays struct{}

func (a arrays) Values() ArraySets { return &values{} }

func (a arrays) Refs() ArraySets { return &refs{} }

type ArraySets interface {
	All() []any

	Primitives() []any
	PrimitivesEmpty() []any
	RefToPrimitives() []any

	Interfaces() []any
	RefsOfInterfaces() []any

	Structs() []any
	RefsOfStructs() []any

	Channels() []any
	RefsOfChannels() []any
}

type values struct{}

func (v values) All() []any {
	return ofValues()
}

func (v values) Primitives() []any {
	return primitivesZero()
}

func (v values) PrimitivesEmpty() []any {
	return primitivesEmpty()
}

func (v values) RefToPrimitives() []any {
	return pointerToPrimitive()
}

func (v values) Interfaces() []any {
	return ofInterfaces()
}

func (v values) RefsOfInterfaces() []any {
	return ofPointersToInterfaces()
}

func (v values) Structs() []any {
	return ofStructs()
}

func (v values) RefsOfStructs() []any {
	return ofPointersToStructs()
}

func (v values) Channels() []any {
	return ofChannels()
}

func (v values) RefsOfChannels() []any {
	return ofPointersToChannels()
}

type refs struct{}

func (r refs) All() []any {
	return ofRefs()
}

func (r refs) Primitives() []any {
	return primitivesZeroRefs()
}

func (r refs) PrimitivesEmpty() []any {
	return primitivesEmptyRefs()
}

func (r refs) RefToPrimitives() []any {
	return pointerToPrimitiveRefs()
}

func (r refs) Interfaces() []any {
	return interfacesRefs()
}

func (r refs) RefsOfInterfaces() []any {
	return pointerToInterfacesRefs()
}

func (r refs) Structs() []any {
	return structsRefs()
}

func (r refs) RefsOfStructs() []any {
	return refToStructsRefs()
}

func (r refs) Channels() []any {
	return channelsRefs()
}

func (r refs) RefsOfChannels() []any {
	return refToChannelRefs()
}
