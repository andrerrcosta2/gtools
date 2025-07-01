// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package zero

import (
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static/structs/models"
)

type Refs struct{}

func (s Refs) Account() *models.Account {
	return new(models.Account)
}

func (s Refs) Address() *models.Address {
	return new(models.Address)
}

func (s Refs) BinaryNode() *models.BinaryNode {
	return new(models.BinaryNode)
}

func (s Refs) Boolean() *models.Boolean {
	return new(models.Boolean)
}

func (s Refs) Channel() *models.Channel[any] {
	return new(models.Channel[any])
}

func (s Refs) CloserError() *models.CloserError {
	return new(models.CloserError)
}

func (s Refs) CloserReaderError() *models.CloserReaderError {
	return new(models.CloserReaderError)
}

func (s Refs) CloserReaderSuccess() *models.CloserReaderSuccess {
	return new(models.CloserReaderSuccess)
}

func (s Refs) CloserSuccess() *models.CloserSuccess {
	return new(models.CloserSuccess)
}

func (s Refs) CloserReaderWriterError() *models.CloserReaderWriterError {
	return new(models.CloserReaderWriterError)
}

func (s Refs) CloserReaderWriterSuccess() *models.CloserReaderWriterSuccess {
	return new(models.CloserReaderWriterSuccess)
}

func (s Refs) Complex() *models.Complex {
	return new(models.Complex)
}

func (s Refs) ComplexUnsafeCastableB() *models.ComplexUnsafeCastableB {
	return new(models.ComplexUnsafeCastableB)
}

func (s Refs) ComplexUnsafeCastableBase() *models.ComplexUnsafeCastableBase {
	return new(models.ComplexUnsafeCastableBase)
}

func (s Refs) ComplexUnsafeCastableC() *models.ComplexUnsafeCastableC {
	return new(models.ComplexUnsafeCastableC)
}

func (s Refs) ComplexUnsafeCastableDeref() *models.ComplexUnsafeCastableDeref {
	return new(models.ComplexUnsafeCastableDeref)
}

func (s Refs) ComplexUnsafeCastableDescription() *models.ComplexUnsafeCastableDescription {
	return new(models.ComplexUnsafeCastableDescription)
}

func (s Refs) ComplexUnsafeCastableOf() *models.ComplexUnsafeCastableOf {
	return new(models.ComplexUnsafeCastableOf)
}

func (s Refs) ComplexUnsafeCastableReinterpreted() *models.ComplexUnsafeCastableReinterpreted {
	return new(models.ComplexUnsafeCastableReinterpreted)
}

func (s Refs) Credential() *models.Credential {
	return new(models.Credential)
}

func (s Refs) DeepBook() *models.DeepBook {
	return new(models.DeepBook)
}

func (s Refs) DeepCompany() *models.DeepCompany {
	return new(models.DeepCompany)
}

func (s Refs) DeepGalaxy() *models.DeepGalaxy {
	return new(models.DeepGalaxy)
}

func (s Refs) DeepLibrary() *models.DeepLibrary {
	return new(models.DeepLibrary)
}

func (s Refs) DeepPackage() *models.DeepPackage {
	return new(models.DeepPackage)
}

func (s Refs) DeepPlanet() *models.DeepPlanet {
	return new(models.DeepPlanet)
}

func (s Refs) DeepSection() *models.DeepSection {
	return new(models.DeepSection)
}

func (s Refs) DeepShelf() *models.DeepShelf {
	return new(models.DeepShelf)
}

func (s Refs) DeepStar() *models.DeepStar {
	return new(models.DeepStar)
}

func (s Refs) DeepTruck() *models.DeepTruck {
	return new(models.DeepTruck)
}

func (s Refs) DeepUniverse() *models.DeepUniverse {
	return new(models.DeepUniverse)
}

func (s Refs) DeepWarehouse() *models.DeepWarehouse {
	return new(models.DeepWarehouse)
}

func (s Refs) Empty() *models.Empty {
	return new(models.Empty)
}

func (s Refs) Float() *models.Float {
	return new(models.Float)
}

func (s Refs) Function() *models.Function[any] {
	return new(models.Function[any])
}

func (s Refs) Integer() *models.Integer {
	return new(models.Integer)
}

func (s Refs) LinkedList() *models.LinkedList {
	return new(models.LinkedList)
}

func (s Refs) Map() *models.Map[string, any] {
	return new(models.Map[string, any])
}

func (s Refs) NaturallyComparable() *models.NaturallyComparable {
	return new(models.NaturallyComparable)
}

func (s Refs) NaturallyComparableWithMethods() *models.NaturallyComparableWithMethods {
	return new(models.NaturallyComparableWithMethods)
}

func (s Refs) NestedStruct() *models.NestedStruct {
	return new(models.NestedStruct)
}

func (s Refs) NotComparable() *models.NotComparable {
	return new(models.NotComparable)
}

func (s Refs) NotComparableWithMethods() *models.NotComparableWithMethods {
	return new(models.NotComparableWithMethods)
}

func (s Refs) NTreeNode() *models.NTreeNode {
	return new(models.NTreeNode)
}

func (s Refs) OneData() *models.OneData {
	return new(models.OneData)
}

func (s Refs) Product() *models.Product {
	return new(models.Product)
}

func (s Refs) Profile() *models.Profile {
	return new(models.Profile)
}

func (s Refs) Public() *models.Public {
	return new(models.Public)
}

func (s Refs) Simple() *models.Simple {
	return new(models.Simple)
}

func (s Refs) SimpleNode() *models.SimpleNode {
	return new(models.SimpleNode)
}

func (s Refs) SimpleUnsafeCastableBase() *models.SimpleUnsafeCastableBase {
	return new(models.SimpleUnsafeCastableBase)
}

func (s Refs) SimpleUnsafeCastableBool() *models.SimpleUnsafeCastableBool {
	return new(models.SimpleUnsafeCastableBool)
}

func (s Refs) SimpleUnsafeCastableFloat() *models.SimpleUnsafeCastableFloat {
	return new(models.SimpleUnsafeCastableFloat)
}

func (s Refs) SimpleUnsafeCastableInt() *models.SimpleUnsafeCastableInt {
	return new(models.SimpleUnsafeCastableInt)
}

func (s Refs) SimpleUnsafeCastableString() *models.SimpleUnsafeCastableString {
	return new(models.SimpleUnsafeCastableString)
}

func (s Refs) Slice() *models.Slice[any] {
	return new(models.Slice[any])
}

func (s Refs) String() *models.String {
	return new(models.String)
}

func (s Refs) Stringer() *models.Stringer {
	return new(models.Stringer)
}

func (s Refs) StringerBytes() *models.StringerBytes {
	return new(models.StringerBytes)
}

func (s Refs) TwoData() *models.TwoData {
	return new(models.TwoData)
}

func (s Refs) Uint() *models.Uint {
	return new(models.Uint)
}

func (s Refs) User() *models.User {
	return new(models.User)
}

func (s Refs) ValuedNotComparable() *models.ValuedNotComparable {
	return new(models.ValuedNotComparable)
}
