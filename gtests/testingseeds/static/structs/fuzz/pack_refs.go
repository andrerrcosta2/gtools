// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package fuzz

import (
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static/structs/models"
)

type Refs struct{}

func (s Refs) Account() *models.Account {
	return models.AccountAsRandRef()
}

func (s Refs) Address() *models.Address {
	return models.AddressAsRandRef()
}

func (s Refs) BinaryNode() *models.BinaryNode {
	return models.BinaryNodeAsRandRef()
}

func (s Refs) Boolean() *models.Boolean {
	return models.BooleanAsRandRef()
}

func (s Refs) Channel() *models.Channel[any] {
	return models.ChannelAsRandRef[any]()
}

func (s Refs) CloserError() *models.CloserError {
	return models.CloserErrorAsRandRef()
}

func (s Refs) CloserReaderError() *models.CloserReaderError {
	return models.CloserReaderErrorAsRandRef()
}

func (s Refs) CloserReaderSuccess() *models.CloserReaderSuccess {
	return models.CloserReaderSuccessAsRandRef()
}

func (s Refs) CloserSuccess() *models.CloserSuccess {
	return models.CloserSuccessAsRandRef()
}

func (s Refs) CloserReaderWriterError() *models.CloserReaderWriterError {
	return models.CloserReaderWriterErrorAsRandRef()
}

func (s Refs) CloserReaderWriterSuccess() *models.CloserReaderWriterSuccess {
	return models.CloserReaderWriterSuccessAsRandRef()
}

func (s Refs) Complex() *models.Complex {
	return models.ComplexAsRandRef()
}

func (s Refs) ComplexUnsafeCastableB() *models.ComplexUnsafeCastableB {
	return models.ComplexUnsafeCastableBAsRandRef()
}

func (s Refs) ComplexUnsafeCastableBase() *models.ComplexUnsafeCastableBase {
	return models.ComplexUnsafeCastableBaseAsRandRef()
}

func (s Refs) ComplexUnsafeCastableC() *models.ComplexUnsafeCastableC {
	return models.ComplexUnsafeCastableCAsRandRef()
}

func (s Refs) ComplexUnsafeCastableDeref() *models.ComplexUnsafeCastableDeref {
	return models.ComplexUnsafeCastableDerefAsRandRef()
}

func (s Refs) ComplexUnsafeCastableDescription() *models.ComplexUnsafeCastableDescription {
	return models.ComplexUnsafeCastableDescriptionAsRandRef()
}

func (s Refs) ComplexUnsafeCastableOf() *models.ComplexUnsafeCastableOf {
	return models.ComplexUnsafeCastableOfAsRandRef()
}

func (s Refs) ComplexUnsafeCastableReinterpreted() *models.ComplexUnsafeCastableReinterpreted {
	return models.ComplexUnsafeCastableReinterpretedAsRandRef()
}

func (s Refs) Credential() *models.Credential {
	return models.CredentialAsRandRef()
}

func (s Refs) DeepBook() *models.DeepBook {
	return models.DeepBookAsRandRef()
}

func (s Refs) DeepCompany() *models.DeepCompany {
	return models.DeepCompanyAsRandRef()
}

func (s Refs) DeepGalaxy() *models.DeepGalaxy {
	return models.DeepGalaxyAsRandRef()
}

func (s Refs) DeepLibrary() *models.DeepLibrary {
	return models.DeepLibraryAsRandRef()
}

func (s Refs) DeepPackage() *models.DeepPackage {
	return models.DeepPackageAsRandRef()
}

func (s Refs) DeepPlanet() *models.DeepPlanet {
	return models.DeepPlanetAsRandRef()
}

func (s Refs) DeepSection() *models.DeepSection {
	return models.DeepSectionAsRandRef()
}

func (s Refs) DeepShelf() *models.DeepShelf {
	return models.DeepShelfAsRandRef()
}

func (s Refs) DeepStar() *models.DeepStar {
	return models.DeepStarAsRandRef()
}

func (s Refs) DeepTruck() *models.DeepTruck {
	return models.DeepTruckAsRandRef()
}

func (s Refs) DeepUniverse() *models.DeepUniverse {
	return models.DeepUniverseAsRandRef()
}

func (s Refs) DeepWarehouse() *models.DeepWarehouse {
	return models.DeepWarehouseAsRandRef()
}

func (s Refs) Empty() *models.Empty {
	return new(models.Empty)
}

func (s Refs) ExportedAddressableFields() *models.ExportedAddressableFields {
	return models.ExportedAddressableFieldsAsRandRef()
}

func (s Refs) ExportedFields() *models.ExportedFields {
	return models.ExportedFieldsAsRandRef()
}

func (s Refs) ExportedManyFields() *models.ExportedManyFields {
	return models.ExportedManyFieldsAsRandRef()
}

func (s Refs) ExportedUnaddressableFields() *models.ExportedUnaddressableFields {
	return models.ExportedUnaddressableFieldsAsRandRef()
}

func (s Refs) Float() *models.Float {
	return models.FloatAsRandRef()
}

func (s Refs) Function() *models.Function[any] {
	return new(models.Function[any])
}

func (s Refs) Integer() *models.Integer {
	return models.IntegerAsRandRef()
}

func (s Refs) LinkedList() *models.LinkedList {
	return models.LinkedListAsRandRef()
}

func (s Refs) Map() *models.Map[string, any] {
	return models.MapAsRandRef[string, any]()
}

func (s Refs) MixedExportedUnexportedFields() *models.MixedExportedUnexportedFields {
	return models.MixedExportedUnexportedFieldsAsRandRef()
}

func (s Refs) NaturallyComparable() *models.NaturallyComparable {
	return models.NaturallyComparableAsRandRef()
}

func (s Refs) NaturallyComparableWithMethods() *models.NaturallyComparableWithMethods {
	return models.NaturallyComparableWithMethodsAsRandRef()
}

func (s Refs) NestedStruct() *models.NestedStruct {
	return models.NestedStructAsRandRef()
}

func (s Refs) NotComparable() *models.NotComparable {
	return models.NotComparableAsRandRef()
}

func (s Refs) NotComparableWithMethods() *models.NotComparableWithMethods {
	return models.NotComparableWithMethodsAsRandRef()
}

func (s Refs) NTreeNode() *models.NTreeNode {
	return models.NTreeNodeAsRandRef()
}

func (s Refs) OneData() *models.OneData {
	return models.OneDataAsRandRef()
}

func (s Refs) Product() *models.Product {
	return models.ProductAsRandRef()
}

func (s Refs) Profile() *models.Profile {
	return models.ProfileAsRandRef()
}

func (s Refs) Public() *models.Public {
	return models.PublicAsRandRef()
}

func (s Refs) Simple() *models.Simple {
	return models.SimpleAsRandRef()
}

func (s Refs) SimpleNode() *models.SimpleNode {
	return models.SimpleNodeAsRandRef()
}

func (s Refs) SimpleUnsafeCastableBase() *models.SimpleUnsafeCastableBase {
	return models.SimpleUnsafeCastableBaseAsRandRef()
}

func (s Refs) SimpleUnsafeCastableBool() *models.SimpleUnsafeCastableBool {
	return models.SimpleUnsafeCastableBoolAsRandRef()
}

func (s Refs) SimpleUnsafeCastableFloat() *models.SimpleUnsafeCastableFloat {
	return models.SimpleUnsafeCastableFloatAsRandRef()
}

func (s Refs) SimpleUnsafeCastableInt() *models.SimpleUnsafeCastableInt {
	return models.SimpleUnsafeCastableIntAsRandRef()
}

func (s Refs) SimpleUnsafeCastableString() *models.SimpleUnsafeCastableString {
	return models.SimpleUnsafeCastableStringAsRandRef()
}

func (s Refs) Slice() *models.Slice[any] {
	return models.SliceAsRandRef[any]()
}

func (s Refs) String() *models.String {
	return models.StringAsRandRef()
}

func (s Refs) Stringer() *models.Stringer {
	return models.StringerAsRandRef()
}

func (s Refs) StringerBytes() *models.StringerBytes {
	return models.StringerBytesAsRandRef()
}

func (s Refs) TwoData() *models.TwoData {
	return models.TwoDataAsRandRef()
}

func (s Refs) Uint() *models.Uint {
	return models.UintAsRandRef()
}

func (s Refs) UnexportedAddressableFields() *models.UnexportedAddressableFields {
	return models.UnexportedAddressableFieldsAsRandRef()
}

func (s Refs) UnexportedFields() *models.UnexportedFields {
	return models.UnexportedFieldsAsRandRef()
}

func (s Refs) UnexportedManyFields() *models.UnexportedManyFields {
	return models.UnexportedManyFieldsAsRandRef()
}

func (s Refs) UnexportedUnaddressableFields() *models.UnexportedUnaddressableFields {
	return models.UnexportedUnaddressableFieldsAsRandRef()
}

func (s Refs) User() *models.User {
	return models.UserAsRandRef()
}

func (s Refs) ValuedNotComparable() *models.ValuedNotComparable {
	return models.ValuedNotComparableAsRandRef()
}
