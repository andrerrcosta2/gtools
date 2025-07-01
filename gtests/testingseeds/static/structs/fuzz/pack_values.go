// Andre R. R. Costa  github.com/andrerrcosta2  andrerrcosta@gmail.com

package fuzz

import (
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static/structs/models"
)

type Values struct{}

func (s Values) Account() models.Account {
	return models.AccountAsRandValue()
}

func (s Values) Address() models.Address {
	return models.AddressAsRandValue()
}

func (s Values) BinaryNode() models.BinaryNode {
	return models.BinaryNodeAsRandValue()
}

func (s Values) Boolean() models.Boolean {
	return models.BooleanAsRandValue()
}

func (s Values) Channel() models.Channel[any] {
	return models.ChannelAsRandValue[any]()
}

func (s Values) CloserError() models.CloserError {
	return models.CloserErrorAsRandValue()
}

func (s Values) CloserReaderError() models.CloserReaderError {
	return models.CloserReaderErrorAsRandValue()
}

func (s Values) CloserReaderSuccess() models.CloserReaderSuccess {
	return models.CloserReaderSuccessAsRandValue()
}

func (s Values) CloserSuccess() models.CloserSuccess {
	return models.CloserSuccessAsRandValue()
}

func (s Values) CloserReaderWriterError() models.CloserReaderWriterError {
	return models.CloserReaderWriterErrorAsRandValue()
}

func (s Values) CloserReaderWriterSuccess() models.CloserReaderWriterSuccess {
	return models.CloserReaderWriterSuccessAsRandValue()
}

func (s Values) Complex() models.Complex {
	return models.ComplexAsRandValue()
}

func (s Values) ComplexUnsafeCastableB() models.ComplexUnsafeCastableB {
	return models.ComplexUnsafeCastableBAsRandValue()
}

func (s Values) ComplexUnsafeCastableBase() models.ComplexUnsafeCastableBase {
	return models.ComplexUnsafeCastableBaseAsRandValue()
}

func (s Values) ComplexUnsafeCastableC() models.ComplexUnsafeCastableC {
	return models.ComplexUnsafeCastableCAsRandValue()
}

func (s Values) ComplexUnsafeCastableDeref() models.ComplexUnsafeCastableDeref {
	return models.ComplexUnsafeCastableDerefAsRandValue()
}

func (s Values) ComplexUnsafeCastableDescription() models.ComplexUnsafeCastableDescription {
	return models.ComplexUnsafeCastableDescriptionAsRandValue()
}

func (s Values) ComplexUnsafeCastableOf() models.ComplexUnsafeCastableOf {
	return models.ComplexUnsafeCastableOfAsRandValue()
}

func (s Values) ComplexUnsafeCastableReinterpreted() models.ComplexUnsafeCastableReinterpreted {
	return models.ComplexUnsafeCastableReinterpretedAsRandValue()
}

func (s Values) Credential() models.Credential {
	return models.CredentialAsRandValue()
}

func (s Values) DeepBook() models.DeepBook {
	return models.DeepBookAsRandValue()
}

func (s Values) DeepCompany() models.DeepCompany {
	return models.DeepCompanyAsRandValue()
}

func (s Values) DeepGalaxy() models.DeepGalaxy {
	return models.DeepGalaxyAsRandValue()
}

func (s Values) DeepLibrary() models.DeepLibrary {
	return models.DeepLibraryAsRandValue()
}

func (s Values) DeepPackage() models.DeepPackage {
	return models.DeepPackageAsRandValue()
}

func (s Values) DeepPlanet() models.DeepPlanet {
	return models.DeepPlanetAsRandValue()
}

func (s Values) DeepSection() models.DeepSection {
	return models.DeepSectionAsRandValue()
}

func (s Values) DeepShelf() models.DeepShelf {
	return models.DeepShelfAsRandValue()
}

func (s Values) DeepStar() models.DeepStar {
	return models.DeepStarAsRandValue()
}

func (s Values) DeepTruck() models.DeepTruck {
	return models.DeepTruckAsRandValue()
}

func (s Values) DeepUniverse() models.DeepUniverse {
	return models.DeepUniverseAsRandValue()
}

func (s Values) DeepWarehouse() models.DeepWarehouse {
	return models.DeepWarehouseAsRandValue()
}

func (s Values) Empty() models.Empty {
	return models.Empty{}
}

func (s Values) ExportedAddressableFields() models.ExportedAddressableFields {
	return models.ExportedAddressableFieldsAsRandValue()
}

func (s Values) ExportedFields() models.ExportedFields {
	return models.ExportedFieldsAsRandValue()
}

func (s Values) ExportedManyFields() models.ExportedManyFields {
	return models.ExportedManyFieldsAsRandValue()
}

func (s Values) ExportedUnaddressableFields() models.ExportedUnaddressableFields {
	return models.ExportedUnaddressableFieldsAsRandValue()
}

func (s Values) Float() models.Float {
	return models.FloatAsRandValue()
}

func (s Values) Function() models.Function[any] {
	return models.FunctionAsRandValue[any]()
}

func (s Values) Integer() models.Integer {
	return models.IntegerAsRandValue()
}

func (s Values) LinkedList() models.LinkedList {
	return models.LinkedListAsRandValue()
}

func (s Values) Map() models.Map[string, any] {
	return models.MapAsRandValue[string, any]()
}

func (s Values) MixedExportedUnexportedFields() models.MixedExportedUnexportedFields {
	return models.MixedExportedUnexportedFieldsAsRandValue()
}

func (s Values) NaturallyComparable() models.NaturallyComparable {
	return models.NaturallyComparableAsRandValue()
}

func (s Values) NaturallyComparableWithMethods() models.NaturallyComparableWithMethods {
	return models.NaturallyComparableWithMethodsAsRandValue()
}

func (s Values) NestedStruct() models.NestedStruct {
	return models.NestedStructAsRandValue()
}

func (s Values) NotComparable() models.NotComparable {
	return models.NotComparableAsRandValue()
}

func (s Values) NotComparableWithMethods() models.NotComparableWithMethods {
	return models.NotComparableWithMethodsAsRandValue()
}

func (s Values) NTreeNode() models.NTreeNode {
	return models.NTreeNodeAsRandValue()
}

func (s Values) OneData() models.OneData {
	return models.OneDataAsRandValue()
}

func (s Values) Product() models.Product {
	return models.ProductAsRandValue()
}

func (s Values) Profile() models.Profile {
	return models.ProfileAsRandValue()
}

func (s Values) Public() models.Public {
	return models.PublicAsRandValue()
}

func (s Values) Simple() models.Simple {
	return models.SimpleAsRandValue()
}

func (s Values) SimpleNode() models.SimpleNode {
	return models.SimpleNodeAsRandValue()
}

func (s Values) SimpleUnsafeCastableBase() models.SimpleUnsafeCastableBase {
	return models.SimpleUnsafeCastableBaseAsRandValue()
}

func (s Values) SimpleUnsafeCastableBool() models.SimpleUnsafeCastableBool {
	return models.SimpleUnsafeCastableBoolAsRandValue()
}

func (s Values) SimpleUnsafeCastableFloat() models.SimpleUnsafeCastableFloat {
	return models.SimpleUnsafeCastableFloatAsRandValue()
}

func (s Values) SimpleUnsafeCastableInt() models.SimpleUnsafeCastableInt {
	return models.SimpleUnsafeCastableIntAsRandValue()
}

func (s Values) SimpleUnsafeCastableString() models.SimpleUnsafeCastableString {
	return models.SimpleUnsafeCastableStringAsRandValue()
}

func (s Values) Slice() models.Slice[any] {
	return models.SliceAsRandValue[any]()
}

func (s Values) String() models.String {
	return models.StringAsRandValue()
}

func (s Values) Stringer() models.Stringer {
	return models.StringerAsRandValue()
}

func (s Values) StringerBytes() models.StringerBytes {
	return models.StringerBytesAsRandValue()
}

func (s Values) TwoData() models.TwoData {
	return models.TwoDataAsRandValue()
}

func (s Values) Uint() models.Uint {
	return models.UintAsRandValue()
}

func (s Values) UnexportedAddressableFields() models.UnexportedAddressableFields {
	return models.UnexportedAddressableFieldsAsRandValue()
}

func (s Values) UnexportedFields() models.UnexportedFields {
	return models.UnexportedFieldsAsRandValue()
}

func (s Values) UnexportedManyFields() models.UnexportedManyFields {
	return models.UnexportedManyFieldsAsRandValue()
}

func (s Values) UnexportedUnaddressableFields() models.UnexportedUnaddressableFields {
	return models.UnexportedUnaddressableFieldsAsRandValue()
}

func (s Values) User() models.User {
	return models.UserAsRandValue()
}

func (s Values) ValuedNotComparable() models.ValuedNotComparable {
	return models.ValuedNotComparableAsRandValue()
}
