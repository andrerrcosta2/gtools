// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package zero

import (
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static/structs/models"
	"slices"
)

func valuesSet() []any {
	return slices.Concat(castableValues(), comparisonValues(), dataStructValues(),
		deepModelsValues(), implValues(), modelValues(), namedValues(), simpleValues())
}

func castableValues() []any {
	return []any{
		new(models.SimpleUnsafeCastableBase), new(models.SimpleUnsafeCastableBool),
		new(models.SimpleUnsafeCastableInt), new(models.SimpleUnsafeCastableString),
		new(models.SimpleUnsafeCastableFloat), new(models.ComplexUnsafeCastableBase),
		new(models.ComplexUnsafeCastableB), new(models.ComplexUnsafeCastableC),
		new(models.ComplexUnsafeCastableOf), new(models.ComplexUnsafeCastableDeref),
		new(models.ComplexUnsafeCastableDescription), new(models.ComplexUnsafeCastableReinterpreted),
	}
}

func comparisonValues() []any {
	return []any{
		new(models.NaturallyComparable), new(models.NaturallyComparableWithMethods),
		new(models.NotComparable),
	}
}

func dataStructValues() []any {
	return []any{
		new(models.BinaryNode), new(models.LinkedList), new(models.NTreeNode),
		new(models.SimpleNode),
	}
}

func deepModelsValues() []any {
	return []any{
		new(models.DeepBook), new(models.DeepCompany), new(models.DeepGalaxy),
		new(models.DeepLibrary), new(models.DeepPackage), new(models.DeepPlanet),
		new(models.DeepSection), new(models.DeepShelf), new(models.DeepStar),
		new(models.DeepTruck), new(models.DeepUniverse), new(models.DeepWarehouse),
	}
}

func implValues() []any {
	return []any{
		new(models.CloserSuccess), new(models.CloserError),
		new(models.CloserReaderSuccess), new(models.CloserReaderError),
		new(models.CloserReaderWriterSuccess), new(models.CloserReaderWriterError),
		new(models.Stringer), new(models.StringerBytes),
		new(models.StringerString),
	}
}

func modelValues() []any {
	return []any{
		new(models.User), new(models.Account), new(models.Address),
		new(models.Credential), new(models.Profile), new(models.Product),
	}
}

func namedValues() []any {
	return []any{
		new(models.Boolean), new(models.Integer), new(models.Uint),
		new(models.Float), new(models.Complex), new(models.String),
		new(models.Slice[any]), new(models.Channel[any]), new(models.Map[any, any]),
		new(models.Function[any]),
	}
}

func simpleValues() []any {
	return []any{
		new(models.Simple), new(models.Public), new(models.Empty),
	}
}
