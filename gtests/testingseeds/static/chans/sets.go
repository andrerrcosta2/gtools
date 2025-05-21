// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package chans

import "github.com/andrerrcosta2/gtools/gtests/testingseeds/static"

func SetOf() static.SeedSets[Categories, Categories] {
	return &chans{}
}

type chans struct{}

func (c chans) Values() Categories {
	return &values{}
}

func (c chans) Refs() Categories {
	return &refs{}
}

type Categories interface {
	All() []any
	Primitives() []any
	RefOfPrimitives() []any
	SliceOfPrimitives() []any
	SliceOfRefToPrimitives() []any
	Interfaces() []any
	RefToInterfaces() []any
	SliceOfInterfaces() []any
	SliceOfRefToInterfaces() []any
	Structs() []any
	SliceOfStructs() []any
	RefToSliceOfStructs() []any
	Senders() []any
	Receivers() []any
	Edges() []any
}

type values struct{}

func (v values) All() []any {
	return ofValues()
}

func (v values) Primitives() []any {
	return primitives()
}

func (v values) RefOfPrimitives() []any {
	return pointerToPrimitives()
}

func (v values) SliceOfPrimitives() []any {
	return sliceOfPrimitives()
}

func (v values) SliceOfRefToPrimitives() []any {
	return sliceOfPointerToPrimitives()
}

func (v values) Interfaces() []any {
	return interfaces()
}

func (v values) RefToInterfaces() []any {
	return pointerToInterfaces()
}

func (v values) SliceOfInterfaces() []any {
	return sliceOfInterfaces()
}

func (v values) SliceOfRefToInterfaces() []any {
	return sliceOfPointerToInterfaces()
}

func (v values) Structs() []any {
	return ofStructs()
}

func (v values) SliceOfStructs() []any {
	return sliceOfStructs()
}

func (v values) RefToSliceOfStructs() []any {
	return pointerToSliceOfStructs()
}

func (v values) Senders() []any {
	return senders()
}

func (v values) Receivers() []any {
	return receivers()
}

func (v values) Edges() []any {
	return edgeValues()
}

type refs struct{}

func (r refs) All() []any {
	return ofRefs()
}

func (r refs) Primitives() []any {
	return primitivesRefs()
}

func (r refs) RefOfPrimitives() []any {
	return pointerToPrimitiveRefs()
}

func (r refs) SliceOfPrimitives() []any {
	return sliceOfPrimitivesRefs()
}

func (r refs) SliceOfRefToPrimitives() []any {
	return sliceOfRefToPrimitivesRefs()
}

func (r refs) Interfaces() []any {
	return interfaceRefs()
}

func (r refs) RefToInterfaces() []any {
	return pointerToInterfaceRefs()
}

func (r refs) SliceOfInterfaces() []any {
	return sliceOfInterfaceRefs()
}

func (r refs) SliceOfRefToInterfaces() []any {
	return sliceOfRefToInterfaceRefs()
}

func (r refs) Structs() []any {
	return ofStructsRefs()
}

func (r refs) SliceOfStructs() []any {
	return sliceOfStructRefs()
}

func (r refs) RefToSliceOfStructs() []any {
	return pointerToSliceOfStructRefs()
}

func (r refs) Senders() []any {
	return sendersAsRefs()
}

func (r refs) Receivers() []any {
	return receiversAsRefs()
}

func (r refs) Edges() []any {
	return edgeRefs()
}
