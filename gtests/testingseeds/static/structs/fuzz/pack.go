// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package fuzz

import "github.com/andrerrcosta2/gtools/core/data/str/iterables"

type (
	Pack       struct{}
	Categories struct{}
	CatRefs    struct{}
	CatValues  struct{}
)

func (p Pack) Categories() Categories {
	return Categories{}
}

func (p Pack) Values() Values {
	return Values{}
}

func (p Pack) Refs() Refs {
	return Refs{}
}

func (c Categories) Refs() CatRefs {
	return CatRefs{}
}

func (c Categories) Values() CatValues {
	return CatValues{}
}

func (s CatRefs) All() *iterables.Slice[any] { return iterables.OfSlice[any](referencesSet()) }

func (s CatRefs) Access() *iterables.Slice[any] {
	return iterables.OfSlice[any](accessRefs()...)
}

func (s CatRefs) Castables() *iterables.Slice[any] { return iterables.OfSlice[any](castableRefs()...) }

func (s CatRefs) Comparisons() *iterables.Slice[any] {
	return iterables.OfSlice[any](comparisonRefs()...)
}

func (s CatRefs) DataStructs() *iterables.Slice[any] {
	return iterables.OfSlice[any](dataStructRefs()...)
}

func (s CatRefs) DeepModels() *iterables.Slice[any] {
	return iterables.OfSlice[any](deepModelsRefs()...)
}

func (s CatRefs) Impls() *iterables.Slice[any] { return iterables.OfSlice[any](implRefs()...) }

func (s CatRefs) Models() *iterables.Slice[any] { return iterables.OfSlice[any](modelRefs()...) }

func (s CatRefs) SingleTyped() *iterables.Slice[any] {
	return iterables.OfSlice[any](singleTypeRefs()...)
}

func (s CatRefs) Simple() *iterables.Slice[any] { return iterables.OfSlice[any](simpleRefs()...) }

func (s CatRefs) WithUnexportedFields() *iterables.Slice[any] {
	return iterables.OfSlice[any](withUnexportedFieldRefs()...)
}

func (s CatRefs) OnlyExportedFields() *iterables.Slice[any] {
	return iterables.OfSlice[any](onlyExportedFieldRefs()...)
}

func (s CatValues) All() *iterables.Slice[any] { return iterables.OfSlice[any](valuesSet()...) }

func (s CatValues) Access() *iterables.Slice[any] {
	return iterables.OfSlice[any](accessValues()...)
}

func (s CatValues) Castables() *iterables.Slice[any] {
	return iterables.OfSlice[any](castableValues()...)
}

func (s CatValues) Comparisons() *iterables.Slice[any] {
	return iterables.OfSlice[any](comparisonValues()...)
}

func (s CatValues) DataStructs() *iterables.Slice[any] {
	return iterables.OfSlice[any](dataStructValues()...)
}

func (s CatValues) DeepModels() *iterables.Slice[any] {
	return iterables.OfSlice[any](deepModelsValues()...)
}

func (s CatValues) Impls() *iterables.Slice[any] { return iterables.OfSlice[any](implValues()...) }

func (s CatValues) Models() *iterables.Slice[any] { return iterables.OfSlice[any](modelValues()...) }

func (s CatValues) SingleTypes() *iterables.Slice[any] {
	return iterables.OfSlice[any](singleTypeValues()...)
}

func (s CatValues) Simple() *iterables.Slice[any] { return iterables.OfSlice[any](simpleValues()...) }

func (s CatValues) WithUnexportedFields() *iterables.Slice[any] {
	return iterables.OfSlice[any](withUnexportedFieldValues()...)
}

func (s CatValues) OfExportedFieldsOnly() *iterables.Slice[any] {
	return iterables.OfSlice[any](onlyExportedFieldValues()...)
}
