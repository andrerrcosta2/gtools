// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package structs

import "github.com/andrerrcosta2/gtools/gtests/testingseeds/static"

func SetOf() static.SeedSets[Categories, Categories] {
	return &structs{}
}

type Categories interface {
	All() []any
	Impls() []any
	Comparisons() []any
	Deeps() []any
	Nameds() []any
	Models() []any
}

type structs struct{}

func (s *structs) Values() Categories {
	return &structValues{}
}

func (s *structs) Refs() Categories {
	return &structRefs{}
}

type structValues struct{}

func (s *structValues) All() []any { return allValues() }

func (s *structValues) Impls() []any { return ofImplementations() }

func (s *structValues) Comparisons() []any { return ofComparisons() }

func (s *structValues) Deeps() []any { return ofDeeps() }

func (s *structValues) Nameds() []any { return ofNameds() }

func (s *structValues) Models() []any { return ofModels() }

type structRefs struct{}

func (s *structRefs) All() []any {
	return allRefs()
}

func (s *structRefs) Impls() []any { return ofImplementationRefs() }

func (s *structRefs) Comparisons() []any { return ofComparisonRefs() }

func (s *structRefs) Deeps() []any { return ofDeepsRefs() }

func (s *structRefs) Nameds() []any { return ofNamedRefs() }

func (s *structRefs) Models() []any { return ofModelRefs() }
