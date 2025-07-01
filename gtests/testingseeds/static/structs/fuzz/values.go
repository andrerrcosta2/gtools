// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package fuzz

import (
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static/structs/models"
	"slices"
)

func valuesSet() []any {
	return slices.Concat(accessValues(), castableValues(), comparisonValues(), dataStructValues(),
		deepModelsValues(), implValues(), modelValues(), singleTypeValues(), simpleValues())
}

func accessValues() []any {
	return []any{
		models.ExportedAddressableFieldsAsRandValue(), models.ExportedFieldsAsRandValue(),
		models.ExportedManyFieldsAsRandValue(), models.ExportedUnaddressableFieldsAsRandValue(),
		models.MixedExportedUnexportedFieldsAsRandValue(), models.UnexportedAddressableFieldsAsRandValue(),
		models.UnexportedFieldsAsRandValue(), models.UnexportedManyFieldsAsRandValue(),
		models.UnexportedUnaddressableFieldsAsRandValue(),
	}
}

func castableValues() []any {
	return []any{
		models.SimpleUnsafeCastableBaseAsRandValue(), models.SimpleUnsafeCastableBoolAsRandValue(),
		models.SimpleUnsafeCastableIntAsRandValue(), models.SimpleUnsafeCastableStringAsRandValue(),
		models.SimpleUnsafeCastableFloatAsRandValue(), models.ComplexUnsafeCastableBaseAsRandValue(),
		models.ComplexUnsafeCastableBAsRandValue(), models.ComplexUnsafeCastableCAsRandValue(),
		models.ComplexUnsafeCastableOfAsRandValue(), models.ComplexUnsafeCastableDerefAsRandValue(),
		models.ComplexUnsafeCastableDescriptionAsRandValue(), models.ComplexUnsafeCastableReinterpretedAsRandValue(),
	}
}

func comparisonValues() []any {
	return []any{
		models.NaturallyComparableAsRandValue(), models.NaturallyComparableWithMethodsAsRandValue(),
		models.NotComparableAsRandValue(),
	}
}

func dataStructValues() []any {
	return []any{
		models.BinaryNodeAsRandValue(), models.LinkedListAsRandValue(), models.NTreeNodeAsRandValue(),
		models.SimpleNodeAsRandValue(),
	}
}

func deepModelsValues() []any {
	return []any{
		models.DeepBookAsRandValue(), models.DeepCompanyAsRandValue(), models.DeepGalaxyAsRandValue(),
		models.DeepLibraryAsRandValue(), models.DeepPackageAsRandValue(), models.DeepPlanetAsRandValue(),
		models.DeepSectionAsRandValue(), models.DeepShelfAsRandValue(), models.DeepStarAsRandValue(),
		models.DeepTruckAsRandValue(), models.DeepUniverseAsRandValue(), models.DeepWarehouseAsRandValue(),
	}
}

func implValues() []any {
	return []any{
		models.CloserSuccessAsRandValue(), models.CloserErrorAsRandValue(),
		models.CloserReaderSuccessAsRandValue(), models.CloserReaderErrorAsRandValue(),
		models.CloserReaderWriterSuccessAsRandValue(), models.CloserReaderWriterErrorAsRandValue(),
		models.StringerAsRandValue(), models.StringerBytesAsRandValue(),
		models.StringerStringAsRandValue(),
	}
}

func modelValues() []any {
	return []any{
		models.UserAsRandValue(), models.AccountAsRandValue(), models.AddressAsRandValue(),
		models.CredentialAsRandValue(), models.ProfileAsRandValue(), models.ProductAsRandValue(),
	}
}

func singleTypeValues() []any {
	return []any{
		models.BooleanAsRandValue(), models.IntegerAsRandValue(), models.UintAsRandValue(),
		models.FloatAsRandValue(), models.ComplexAsRandValue(), models.StringAsRandValue(),
		models.SliceAsRandValue[any](), models.ChannelAsRandValue[any](),
		models.MapAsRandValue[string, any](), models.FunctionAsRandValue[any](),
	}
}

func simpleValues() []any {
	return []any{
		models.SimpleAsRandValue(), models.PublicAsRandValue(), models.Empty{},
	}
}

func withUnexportedFieldValues() []any {
	return []any{
		models.NaturallyComparableWithMethodsAsRandValue(), models.NotComparableWithMethodsAsRandValue(),
		models.ValuedNotComparableAsRandValue(), models.BooleanAsRandValue(), models.ChannelAsRandValue[any](),
		models.CloserReaderErrorAsRandValue(), models.CloserReaderSuccessAsRandValue(),
		models.CloserReaderWriterErrorAsRandValue(), models.CloserReaderWriterSuccessAsRandValue(),
		models.ComplexAsRandValue(), models.IntegerAsRandValue(), models.MapAsRandValue[string, any](),
		models.FloatAsRandValue(), models.SliceAsRandValue[any](), models.StringAsRandValue(),
		models.StringAsRandValue(), models.StringerBytesAsRandValue(), models.StringerStringAsRandValue(),
		models.UintAsRandValue(), models.FunctionAsRandValue[any](),
	}
}

func onlyExportedFieldValues() []any {
	return []any{
		models.AccountAsRandValue(), models.AddressAsRandValue(), models.DeepPlanetAsRandValue(),
		models.DeepStarAsRandValue(), models.DeepGalaxyAsRandValue(), models.DeepUniverseAsRandValue(),
		models.NotComparableAsRandValue(), models.NaturallyComparableAsRandValue(), models.SimpleNodeAsRandValue(),
		models.NTreeNodeAsRandValue(), models.LinkedListAsRandValue(), models.BinaryNodeAsRandValue(),
		models.TwoDataAsRandValue(), models.DeepLibraryAsRandValue(), models.DeepSectionAsRandValue(),
		models.DeepShelfAsRandValue(), models.DeepBookAsRandValue(), models.DeepCompanyAsRandValue(),
		models.DeepWarehouseAsRandValue(), models.DeepTruckAsRandValue(), models.DeepPackageAsRandValue(),
		models.OneDataAsRandValue(), models.ProfileAsRandValue(), models.ProductAsRandValue(),
		models.CredentialAsRandValue(), models.UserAsRandValue(),
	}
}
