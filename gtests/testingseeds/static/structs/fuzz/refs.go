// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package fuzz

import (
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static/structs/models"
	"slices"
)

func referencesSet() []any {
	return slices.Concat(accessRefs(), castableRefs(), comparisonRefs(), dataStructRefs(),
		deepModelsRefs(), implRefs(), modelRefs(), singleTypeRefs(), simpleRefs())
}

func accessRefs() []any {
	return []any{
		models.ExportedAddressableFieldsAsRandRef(), models.ExportedFieldsAsRandRef(),
		models.ExportedManyFieldsAsRandRef(), models.ExportedUnaddressableFieldsAsRandRef(),
		models.MixedExportedUnexportedFieldsAsRandRef(), models.UnexportedAddressableFieldsAsRandRef(),
		models.UnexportedFieldsAsRandRef(), models.UnexportedManyFieldsAsRandRef(),
		models.UnexportedUnaddressableFieldsAsRandRef(),
	}
}

func castableRefs() []any {
	return []any{
		models.SimpleUnsafeCastableBaseAsRandRef(), models.SimpleUnsafeCastableBoolAsRandRef(),
		models.SimpleUnsafeCastableIntAsRandRef(), models.SimpleUnsafeCastableStringAsRandRef(),
		models.SimpleUnsafeCastableFloatAsRandRef(), models.ComplexUnsafeCastableBaseAsRandRef(),
		models.ComplexUnsafeCastableBAsRandRef(), models.ComplexUnsafeCastableCAsRandRef(),
		models.ComplexUnsafeCastableOfAsRandRef(), models.ComplexUnsafeCastableDerefAsRandRef(),
		models.ComplexUnsafeCastableDescriptionAsRandRef(), models.ComplexUnsafeCastableReinterpretedAsRandRef(),
	}
}

func comparisonRefs() []any {
	return []any{
		models.NaturallyComparableAsRandRef(), models.NaturallyComparableWithMethodsAsRandRef(),
		models.NotComparableAsRandRef(),
	}
}

func dataStructRefs() []any {
	return []any{
		models.BinaryNodeAsRandRef(), models.LinkedListAsRandRef(), models.NTreeNodeAsRandRef(),
		models.SimpleNodeAsRandRef(),
	}
}

func deepModelsRefs() []any {
	return []any{
		models.DeepBookAsRandRef(), models.DeepCompanyAsRandRef(), models.DeepGalaxyAsRandRef(),
		models.DeepLibraryAsRandRef(), models.DeepPackageAsRandRef(), models.DeepPlanetAsRandRef(),
		models.DeepSectionAsRandRef(), models.DeepShelfAsRandRef(), models.DeepStarAsRandRef(),
		models.DeepTruckAsRandRef(), models.DeepUniverseAsRandRef(), models.DeepWarehouseAsRandRef(),
	}
}

func implRefs() []any {
	return []any{
		models.CloserSuccessAsRandRef(), models.CloserErrorAsRandRef(),
		models.CloserReaderSuccessAsRandRef(), models.CloserReaderErrorAsRandRef(),
		models.CloserReaderWriterSuccessAsRandRef(), models.CloserReaderWriterErrorAsRandRef(),
		models.StringerAsRandRef(), models.StringerBytesAsRandRef(),
		models.StringerStringAsRandRef(),
	}
}

func modelRefs() []any {
	return []any{
		models.UserAsRandRef(), models.AccountAsRandRef(), models.AddressAsRandRef(),
		models.CredentialAsRandRef(), models.ProfileAsRandRef(), models.ProductAsRandRef(),
	}
}

func singleTypeRefs() []any {
	return []any{
		models.BooleanAsRandRef(), models.IntegerAsRandRef(), models.UintAsRandRef(),
		models.FloatAsRandRef(), models.ComplexAsRandRef(), models.StringAsRandRef(),
		models.SliceAsRandRef[any](), models.ChannelAsRandRef[any](),
		models.MapAsRandRef[string, any](), models.FunctionAsRandRef[any](),
	}
}

func simpleRefs() []any {
	return []any{
		models.SimpleAsRandRef(), models.PublicAsRandRef(), new(models.Empty),
	}
}

func withUnexportedFieldRefs() []any {
	return []any{
		models.NaturallyComparableWithMethodsAsRandRef(), models.NotComparableWithMethodsAsRandRef(),
		models.ValuedNotComparableAsRandRef(), models.BooleanAsRandRef(), models.ChannelAsRandRef[any](),
		models.CloserReaderErrorAsRandRef(), models.CloserReaderSuccessAsRandRef(),
		models.CloserReaderWriterErrorAsRandRef(), models.CloserReaderWriterSuccessAsRandRef(),
		models.ComplexAsRandRef(), models.IntegerAsRandRef(), models.MapAsRandRef[string, any](),
		models.FloatAsRandRef(), models.SliceAsRandRef[any](), models.StringAsRandRef(),
		models.StringAsRandRef(), models.StringerBytesAsRandRef(), models.StringerStringAsRandRef(),
		models.UintAsRandRef(), models.FunctionAsRandRef[any](),
	}
}

func onlyExportedFieldRefs() []any {
	return []any{
		models.AccountAsRandRef(), models.AddressAsRandRef(), models.DeepPlanetAsRandRef(),
		models.DeepStarAsRandRef(), models.DeepGalaxyAsRandRef(), models.DeepUniverseAsRandRef(),
		models.NotComparableAsRandRef(), models.NaturallyComparableAsRandRef(), models.SimpleNodeAsRandRef(),
		models.NTreeNodeAsRandRef(), models.LinkedListAsRandRef(), models.BinaryNodeAsRandRef(),
		models.TwoDataAsRandRef(), models.DeepLibraryAsRandRef(), models.DeepSectionAsRandRef(),
		models.DeepShelfAsRandRef(), models.DeepBookAsRandRef(), models.DeepCompanyAsRandRef(),
		models.DeepWarehouseAsRandRef(), models.DeepTruckAsRandRef(), models.DeepPackageAsRandRef(),
		models.OneDataAsRandRef(), models.ProfileAsRandRef(), models.ProductAsRandRef(),
		models.CredentialAsRandRef(), models.UserAsRandRef(),
	}
}
