// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package slices

import "github.com/andrerrcosta2/gtools/gtests/testingseeds/static"

func SetOf() static.SeedSets[Categories, Categories] {
	return slicesSet{}
}

type slicesSet struct{}

func (s slicesSet) Values() Categories {
	return values{}
}

func (s slicesSet) Refs() Categories {
	return refs{}
}

type Categories interface {
	All() []any
	PrimitivesEmpty() []any
	PrimitivesNil() []any
	RefToPrimitivesNil() []any
	Interfaces() []any
	RefsOfInterfaces() []any
	Structs() []any
	RefsOfStructs() []any
	Channels() []any
	RefsOfChannels() []any
	Edges() []any
}

type values struct{}

func (v values) All() []any {
	return ofValues()
}

func (v values) PrimitivesEmpty() []any {
	return primitivesEmpty()
}

func (v values) PrimitivesNil() []any {
	return primitivesNil()
}

func (v values) RefToPrimitivesNil() []any {
	return pointerToPrimitiveNil()
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
	return ofPointersOfChannels()
}

func (v values) Edges() []any {
	return edgeValues()
}

type refs struct{}

func (r refs) All() []any {
	return ofRefs()
}

func (r refs) PrimitivesEmpty() []any {
	return primitivesEmptyAsRef()
}

func (r refs) PrimitivesNil() []any {
	return primitivesNilAsRef()
}

func (r refs) RefToPrimitivesNil() []any {
	return pointerToPrimitiveNilAsRef()
}

func (r refs) Interfaces() []any {
	return ofInterfacesAsRef()
}

func (r refs) RefsOfInterfaces() []any {
	return ofPointersToInterfacesAsRef()
}

func (r refs) Structs() []any {
	return ofStructsAsRef()
}

func (r refs) RefsOfStructs() []any {
	return ofPointersToStructsAsRef()
}

func (r refs) Channels() []any {
	return ofChannelsAsRef()
}

func (r refs) RefsOfChannels() []any {
	return ofPointersToChannelsAsRef()
}

func (r refs) Edges() []any {
	return edgeRefs()
}
