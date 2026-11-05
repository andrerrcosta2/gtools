// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package zero

type (
	Pack       struct{}
	Categories struct{}
	CatRefs    struct{}
	CatValues  struct{}
)

func (s Pack) Categories() Categories {
	return Categories{}
}

func (s Pack) Values() Values {
	return Values{}
}

func (s Pack) Refs() Refs {
	return Refs{}
}

func (s Categories) Refs() CatRefs {
	return CatRefs{}
}

func (s Categories) Values() CatValues {
	return CatValues{}
}

func (s CatRefs) All() []any { return referencesSet() }

func (s CatRefs) Castables() []any { return castableRefs() }

func (s CatRefs) Comparisons() []any { return comparisonRefs() }

func (s CatRefs) DataStructs() []any { return dataStructRefs() }

func (s CatRefs) DeepModels() []any { return deepModelsRefs() }

func (s CatRefs) Impls() []any { return implRefs() }

func (s CatRefs) Models() []any { return modelRefs() }

func (s CatRefs) Named() []any { return namedRefs() }

func (s CatRefs) Simple() []any { return simpleRefs() }

func (s CatValues) All() []any { return valuesSet() }

func (s CatValues) Castables() []any { return castableValues() }

func (s CatValues) Comparisons() []any { return comparisonValues() }

func (s CatValues) DataStructs() []any { return dataStructValues() }

func (s CatValues) DeepModels() []any { return deepModelsValues() }

func (s CatValues) Impls() []any { return implValues() }

func (s CatValues) Models() []any { return modelValues() }

func (s CatValues) Named() []any { return namedValues() }

func (s CatValues) Simple() []any { return simpleValues() }
