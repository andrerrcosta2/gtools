// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package models

import (
	"github.com/andrerrcosta2/gtools/core/format/fmx"
	"github.com/andrerrcosta2/gtools/core/seeders/random"
)

var NestedStructZeroInst = new(NestedStruct)

func NestedStructAsValue(field1 int, field2 string) NestedStruct {
	return NestedStruct{NestedField1: field1, NestedField2: field2}
}

func NestedStructAsRef(field1 int, field2 string) *NestedStruct {
	return &NestedStruct{NestedField1: field1, NestedField2: field2}
}

func NestedStructAsRandValue() NestedStruct {
	return NestedStruct{
		NestedField1: random.SingleOf[int](),
		NestedField2: random.SingleOf[string](),
	}
}

func NestedStructAsRandRef() *NestedStruct {
	return &NestedStruct{
		NestedField1: random.SingleOf[int](),
		NestedField2: random.SingleOf[string](),
	}
}

// NestedStruct specifications:
//   - size 24 bytes
//   - alignment 8
//   - no padding
//
// A string in Go is composed of an address and a length, both of which are 8 bytes each.
// Therefore, the total size of the string field is 16 bytes, but it aligns to 8 bytes.
type NestedStruct struct {
	NestedField1 int    // Size 8 bytes. Offset 0
	NestedField2 string // Size 16 bytes. Offset 8
}

func (a *NestedStruct) String() string {
	return fmx.Sprintf("%d %s", a.NestedField1, a.NestedField2)
}

var ComplexUnsafeCastableBaseZeroInst = new(ComplexUnsafeCastableBase)

func ComplexUnsafeCastableBaseAsRef(nested NestedStruct, id int, metadata map[string]any, score float64) *ComplexUnsafeCastableBase {
	return &ComplexUnsafeCastableBase{
		Nested:   nested,
		ID:       id,
		Metadata: metadata,
		Score:    score,
	}
}

func ComplexUnsafeCastableBaseAsValue(nested NestedStruct, id int, metadata map[string]any, score float64) ComplexUnsafeCastableBase {
	return ComplexUnsafeCastableBase{
		Nested:   nested,
		ID:       id,
		Metadata: metadata,
		Score:    score,
	}
}

func ComplexUnsafeCastableBaseAsRandRef() *ComplexUnsafeCastableBase {
	return &ComplexUnsafeCastableBase{
		Nested:   NestedStructAsRandValue(),
		ID:       random.SingleOf[int](),
		Metadata: random.SingleOf[map[string]any](),
		Score:    random.SingleOf[float64](),
	}
}

func ComplexUnsafeCastableBaseAsRandValue() ComplexUnsafeCastableBase {
	return ComplexUnsafeCastableBase{
		Nested:   NestedStructAsRandValue(),
		ID:       random.SingleOf[int](),
		Metadata: random.SingleOf[map[string]any](),
		Score:    random.SingleOf[float64](),
	}
}

// ComplexUnsafeCastableBase is a base struct for unsafe casts
// Specifications:
//   - size 56 bytes
//   - alignment 8
//   - empty padding 0
type ComplexUnsafeCastableBase struct {
	Nested   NestedStruct   // Size 24 bytes. Offset 0
	ID       int            // Size 8 bytes. Offset 24
	Metadata map[string]any // Size 16 bytes. Offset 32
	Score    float64        // Size 8 bytes. Offset 48
}

var ComplexUnsafeCastableOfZeroInst = new(ComplexUnsafeCastableOf)

func ComplexUnsafeCastableOfAsValue(nested NestedStruct, id int, metadata map[string]any, score float64) ComplexUnsafeCastableOf {
	return ComplexUnsafeCastableOf{
		ComplexUnsafeCastableBase: ComplexUnsafeCastableBase{
			Nested:   nested,
			ID:       id,
			Metadata: metadata,
			Score:    score,
		},
	}
}

func ComplexUnsafeCastableOfAsRef(nested NestedStruct, id int, metadata map[string]any, score float64) *ComplexUnsafeCastableOf {
	return &ComplexUnsafeCastableOf{
		ComplexUnsafeCastableBase: ComplexUnsafeCastableBase{
			Nested:   nested,
			ID:       id,
			Metadata: metadata,
			Score:    score,
		},
	}
}

func ComplexUnsafeCastableOfAsRandRef() *ComplexUnsafeCastableOf {
	return &ComplexUnsafeCastableOf{
		ComplexUnsafeCastableBase: ComplexUnsafeCastableBaseAsRandValue(),
	}
}

func ComplexUnsafeCastableOfAsRandValue() ComplexUnsafeCastableOf {
	return ComplexUnsafeCastableOf{
		ComplexUnsafeCastableBase: ComplexUnsafeCastableBaseAsRandValue(),
	}
}

// ComplexUnsafeCastableOf is a struct for unsafe casts
// Specifications:
//   - size 56 bytes
//   - alignment 8
//   - empty padding 0
type ComplexUnsafeCastableOf struct {
	ComplexUnsafeCastableBase
}

var ComplexUnsafeCastableDescriptionZeroInst = new(ComplexUnsafeCastableDescription)

func ComplexUnsafeCastableDescriptionAsRef(nested NestedStruct, id int, metadata map[string]any, score float64, description string) *ComplexUnsafeCastableDescription {
	return &ComplexUnsafeCastableDescription{
		ComplexUnsafeCastableBase: ComplexUnsafeCastableBase{
			Nested:   nested,
			ID:       id,
			Metadata: metadata,
			Score:    score,
		},
		Description: description,
	}
}

func ComplexUnsafeCastableDescriptionAsValue(nested NestedStruct, id int, metadata map[string]any, score float64, description string) ComplexUnsafeCastableDescription {
	return ComplexUnsafeCastableDescription{
		ComplexUnsafeCastableBase: ComplexUnsafeCastableBase{
			Nested:   nested,
			ID:       id,
			Metadata: metadata,
			Score:    score,
		},
		Description: description,
	}
}

func ComplexUnsafeCastableDescriptionAsRandRef() *ComplexUnsafeCastableDescription {
	return &ComplexUnsafeCastableDescription{
		ComplexUnsafeCastableBase: ComplexUnsafeCastableBaseAsRandValue(),
		Description:               random.SingleOf[string](),
	}
}

func ComplexUnsafeCastableDescriptionAsRandValue() ComplexUnsafeCastableDescription {
	return ComplexUnsafeCastableDescription{
		ComplexUnsafeCastableBase: ComplexUnsafeCastableBaseAsRandValue(),
		Description:               random.SingleOf[string](),
	}
}

// ComplexUnsafeCastableDescription is a struct for unsafe casts
// Specifications:
//   - size 80 bytes
//   - alignment 8
//   - empty padding 0
type ComplexUnsafeCastableDescription struct {
	ComplexUnsafeCastableBase        // Size 56 bytes. Offset 0
	Value                     int    // Size 8 bytes. Offset 56
	Description               string // Size 16 bytes. Offset 64
}

// ComplexUnsafeCastableA is a struct for unsafe casts
// Specifications:
//   - size 80 bytes
//   - alignment 8
//   - empty padding 0
type ComplexUnsafeCastableA struct {
	ComplexUnsafeCastableBase          // Size 56 bytes. Offset 0
	Idx                       Uint     // Size 8 bytes. Offset 56
	List                      []String // Size 16 bytes. Offset 64
}

var ComplexUnsafeCastableBZeroInst = new(ComplexUnsafeCastableB)

func ComplexUnsafeCastableBAsValue(nested NestedStruct, id int, innerMeta, outerMeta map[string]any, innerScore, outerScore float64) ComplexUnsafeCastableB {
	return ComplexUnsafeCastableB{
		ComplexUnsafeCastableBase: ComplexUnsafeCastableBase{
			Nested:   nested,
			ID:       id,
			Metadata: innerMeta,
			Score:    innerScore,
		},
		Score:    outerScore,
		Metadata: outerMeta,
	}
}

func ComplexUnsafeCastableBAsRef(nested NestedStruct, id int, innerMeta, outerMeta map[string]any,
	innerScore, outerScore float64) *ComplexUnsafeCastableB {
	return &ComplexUnsafeCastableB{
		ComplexUnsafeCastableBase: ComplexUnsafeCastableBase{
			Nested:   nested,
			ID:       id,
			Metadata: innerMeta,
			Score:    innerScore,
		},
		Score:    outerScore,
		Metadata: outerMeta,
	}
}

func ComplexUnsafeCastableBAsRandValue() ComplexUnsafeCastableB {
	return ComplexUnsafeCastableB{
		ComplexUnsafeCastableBase: ComplexUnsafeCastableBaseAsRandValue(),
		Score:                     random.SingleOf[float64](),
		Metadata:                  random.SingleOf[map[string]any](),
	}
}

func ComplexUnsafeCastableBAsRandRef() *ComplexUnsafeCastableB {
	return &ComplexUnsafeCastableB{
		ComplexUnsafeCastableBase: ComplexUnsafeCastableBaseAsRandValue(),
		Score:                     random.SingleOf[float64](),
		Metadata:                  random.SingleOf[map[string]any](),
	}
}

// ComplexUnsafeCastableB is a struct for unsafe casts
// Specifications:
//   - size 80 bytes
//   - alignment 8
//   - empty padding 0
type ComplexUnsafeCastableB struct {
	ComplexUnsafeCastableBase                // Size 56 bytes. Offset 0
	Score                     float64        // Size 8 bytes. Offset 56
	Metadata                  map[string]any // Size 16 bytes. Offset 64
}

var ComplexUnsafeCastableCZeroInst = new(ComplexUnsafeCastableC)

func ComplexUnsafeCastableCAsValue(nested NestedStruct, id int, metadata map[string]any, score float64,
	less uint32) ComplexUnsafeCastableC {
	return ComplexUnsafeCastableC{
		ComplexUnsafeCastableBase: ComplexUnsafeCastableBase{
			Nested:   nested,
			ID:       id,
			Metadata: metadata,
			Score:    score,
		},
		Less: less,
	}
}

func ComplexUnsafeCastableCAsRef(nested NestedStruct, id int, metadata map[string]any, score float64,
	less uint32) *ComplexUnsafeCastableC {
	return &ComplexUnsafeCastableC{
		ComplexUnsafeCastableBase: ComplexUnsafeCastableBase{
			Nested:   nested,
			ID:       id,
			Metadata: metadata,
			Score:    score,
		},
		Less: less,
	}
}

func ComplexUnsafeCastableCAsRandValue() ComplexUnsafeCastableC {
	return ComplexUnsafeCastableC{
		ComplexUnsafeCastableBase: ComplexUnsafeCastableBaseAsRandValue(),
		Less:                      random.SingleOf[uint32](),
	}
}

func ComplexUnsafeCastableCAsRandRef() *ComplexUnsafeCastableC {
	return &ComplexUnsafeCastableC{
		ComplexUnsafeCastableBase: ComplexUnsafeCastableBaseAsRandValue(),
		Less:                      random.SingleOf[uint32](),
	}
}

// ComplexUnsafeCastableC is a struct for unsafe casts
// Specifications:
//   - size 60 bytes
//   - alignment 8
//   - empty padding 0
type ComplexUnsafeCastableC struct {
	ComplexUnsafeCastableBase        // Size 56 bytes. Offset 0
	Less                      uint32 // Size 4 bytes. Offset 56
}

var ComplexUnsafeCastableDerefZeroInst = new(ComplexUnsafeCastableDeref)

func ComplexUnsafeCastableDerefAsValue(field1 int, field2 string, id int, metadata map[string]any, score float64) ComplexUnsafeCastableDeref {
	return ComplexUnsafeCastableDeref{
		NestedField1: field1,
		NestedField2: field2,
		ID:           id,
		Metadata:     metadata,
		Score:        score,
	}
}

func ComplexUnsafeCastableDerefAsRef(field1 int, field2 string, id int, metadata map[string]any, score float64) *ComplexUnsafeCastableDeref {
	return &ComplexUnsafeCastableDeref{
		NestedField1: field1,
		NestedField2: field2,
		ID:           id,
		Metadata:     metadata,
		Score:        score,
	}
}

func ComplexUnsafeCastableDerefAsRandValue() ComplexUnsafeCastableDeref {
	return ComplexUnsafeCastableDeref{
		NestedField1: random.SingleOf[int](),
		NestedField2: random.SingleOf[string](),
		ID:           random.SingleOf[int](),
		Metadata:     random.SingleOf[map[string]any](),
		Score:        random.SingleOf[float64](),
	}
}

func ComplexUnsafeCastableDerefAsRandRef() *ComplexUnsafeCastableDeref {
	return &ComplexUnsafeCastableDeref{
		NestedField1: random.SingleOf[int](),
		NestedField2: random.SingleOf[string](),
		ID:           random.SingleOf[int](),
		Metadata:     random.SingleOf[map[string]any](),
		Score:        random.SingleOf[float64](),
	}
}

// ComplexUnsafeCastableDeref is a struct for unsafe casts
// Specifications:
//   - size 56 bytes
//   - alignment 8
//   - empty padding 0
type ComplexUnsafeCastableDeref struct {
	NestedField1 int            // Size 8 bytes. Offset 0
	NestedField2 string         // Size 16 bytes. Offset 8
	ID           int            // Size 8 bytes. Offset 24
	Metadata     map[string]any // Size 16 bytes. Offset 32
	Score        float64        // Size 8 bytes. Offset 48
}

var ComplexUnsafeCastableReinterpretedZeroInst = new(ComplexUnsafeCastableReinterpreted)

func ComplexUnsafeCastableReinterpretedAsValue(fieldA float64, fieldB any, fieldC uint, fieldD String, fieldE int) ComplexUnsafeCastableReinterpreted {
	return ComplexUnsafeCastableReinterpreted{
		FieldA: fieldA,
		FieldB: fieldB,
		FieldC: fieldC,
		FieldD: fieldD,
		FieldE: fieldE,
	}
}

func ComplexUnsafeCastableReinterpretedAsRef(fieldA float64, fieldB any, fieldC uint, fieldD String, fieldE int) *ComplexUnsafeCastableReinterpreted {
	return &ComplexUnsafeCastableReinterpreted{
		FieldA: fieldA,
		FieldB: fieldB,
		FieldC: fieldC,
		FieldD: fieldD,
		FieldE: fieldE,
	}
}

func ComplexUnsafeCastableReinterpretedAsRandValue() ComplexUnsafeCastableReinterpreted {
	x := ComplexUnsafeCastableReinterpreted{
		FieldA: random.SingleOf[float64](),
		FieldB: random.SingleOf[any](),
		FieldC: random.SingleOf[uint](),
		FieldD: random.SingleOf[String](),
		FieldE: random.SingleOf[int](),
	}
	return x
}

func ComplexUnsafeCastableReinterpretedAsRandRef() *ComplexUnsafeCastableReinterpreted {
	return &ComplexUnsafeCastableReinterpreted{
		FieldA: random.SingleOf[float64](),
		FieldB: random.SingleOf[any](),
		FieldC: random.SingleOf[uint](),
		FieldD: random.SingleOf[String](),
		FieldE: random.SingleOf[int](),
	}
}

// ComplexUnsafeCastableReinterpreted is a struct for unsafe casts
// Specifications:
//   - size 56 bytes
//   - alignment 8
//   - empty padding 0
type ComplexUnsafeCastableReinterpreted struct {
	FieldA float64 // Size 8 bytes. Offset 0
	FieldB any     // Size 16 bytes. Offset 8
	FieldC uint    // Size 8 bytes. Offset 24
	FieldD String  // Size 16 bytes. Offset 32
	FieldE int     // Size 8 bytes. Offset 48
}

var SimpleUnsafeCastableBaseZeroInst = new(SimpleUnsafeCastableBase)

func SimpleUnsafeCastableBaseAsRef(id int) *SimpleUnsafeCastableBase {
	return &SimpleUnsafeCastableBase{
		ID: id,
	}
}

func SimpleUnsafeCastableBaseAsValue(id int) SimpleUnsafeCastableBase {
	return SimpleUnsafeCastableBase{
		ID: id,
	}
}

func SimpleUnsafeCastableBaseAsRandRef() *SimpleUnsafeCastableBase {
	return &SimpleUnsafeCastableBase{
		ID: random.SingleOf[int](),
	}
}

func SimpleUnsafeCastableBaseAsRandValue() SimpleUnsafeCastableBase {
	return SimpleUnsafeCastableBase{
		ID: random.SingleOf[int](),
	}
}

// SimpleUnsafeCastableBase is a base struct for unsafe casts
// Unsafe casts are used to modify the underlying memory of unsafe pointers
type SimpleUnsafeCastableBase struct {
	ID int
}

var SimpleUnsafeCastableBoolZeroInst = new(SimpleUnsafeCastableBool)

func SimpleUnsafeCastableBoolAsValue(id int, value bool) SimpleUnsafeCastableBool {
	return SimpleUnsafeCastableBool{
		SimpleUnsafeCastableBase: SimpleUnsafeCastableBase{
			ID: id,
		},
		ValueD: value,
	}
}

func SimpleUnsafeCastableBoolAsRef(id int, value bool) *SimpleUnsafeCastableBool {
	return &SimpleUnsafeCastableBool{
		SimpleUnsafeCastableBase: SimpleUnsafeCastableBase{
			ID: id,
		},
		ValueD: value,
	}
}

func SimpleUnsafeCastableBoolAsRandRef() *SimpleUnsafeCastableBool {
	return &SimpleUnsafeCastableBool{
		SimpleUnsafeCastableBase: SimpleUnsafeCastableBase{
			ID: random.SingleOf[int](),
		},
		ValueD: random.SingleOf[bool](),
	}
}

func SimpleUnsafeCastableBoolAsRandValue() SimpleUnsafeCastableBool {
	return SimpleUnsafeCastableBool{
		SimpleUnsafeCastableBase: SimpleUnsafeCastableBase{
			ID: random.SingleOf[int](),
		},
		ValueD: random.SingleOf[bool](),
	}
}

type SimpleUnsafeCastableBool struct {
	SimpleUnsafeCastableBase
	ValueD bool
}

var SimpleUnsafeCastableFloatZeroInst = new(SimpleUnsafeCastableFloat)

func SimpleUnsafeCastableFloatAsValue(id int, value float64) SimpleUnsafeCastableFloat {
	return SimpleUnsafeCastableFloat{
		SimpleUnsafeCastableBase: SimpleUnsafeCastableBase{
			ID: id,
		},
		ValueC: value,
	}
}

func SimpleUnsafeCastableFloatAsRef(id int, value float64) *SimpleUnsafeCastableFloat {
	return &SimpleUnsafeCastableFloat{
		SimpleUnsafeCastableBase: SimpleUnsafeCastableBase{
			ID: id,
		},
		ValueC: value,
	}
}

func SimpleUnsafeCastableFloatAsRandRef() *SimpleUnsafeCastableFloat {
	return &SimpleUnsafeCastableFloat{
		SimpleUnsafeCastableBase: SimpleUnsafeCastableBase{
			ID: random.SingleOf[int](),
		},
		ValueC: random.SingleOf[float64](),
	}
}

func SimpleUnsafeCastableFloatAsRandValue() SimpleUnsafeCastableFloat {
	return SimpleUnsafeCastableFloat{
		SimpleUnsafeCastableBase: SimpleUnsafeCastableBase{
			ID: random.SingleOf[int](),
		},
		ValueC: random.SingleOf[float64](),
	}
}

type SimpleUnsafeCastableFloat struct {
	SimpleUnsafeCastableBase
	ValueC float64
}

var SimpleUnsafeCastableIntZeroInst = new(SimpleUnsafeCastableInt)

func SimpleUnsafeCastableIntAsValue(id, value int) SimpleUnsafeCastableInt {
	return SimpleUnsafeCastableInt{
		SimpleUnsafeCastableBase: SimpleUnsafeCastableBase{
			ID: id,
		},
		ValueB: value,
	}
}

func SimpleUnsafeCastableIntAsRef(id, value int) *SimpleUnsafeCastableInt {
	return &SimpleUnsafeCastableInt{
		SimpleUnsafeCastableBase: SimpleUnsafeCastableBase{
			ID: id,
		},
		ValueB: value,
	}
}

func SimpleUnsafeCastableIntAsRandRef() *SimpleUnsafeCastableInt {
	return &SimpleUnsafeCastableInt{
		SimpleUnsafeCastableBase: SimpleUnsafeCastableBase{
			ID: random.SingleOf[int](),
		},
		ValueB: random.SingleOf[int](),
	}
}

func SimpleUnsafeCastableIntAsRandValue() SimpleUnsafeCastableInt {
	return SimpleUnsafeCastableInt{
		SimpleUnsafeCastableBase: SimpleUnsafeCastableBase{
			ID: random.SingleOf[int](),
		},
		ValueB: random.SingleOf[int](),
	}
}

type SimpleUnsafeCastableInt struct {
	SimpleUnsafeCastableBase
	ValueB int
}

var UnsafeCastableStringZeroInst = new(SimpleUnsafeCastableString)

func SimpleUnsafeCastableStringAsValue(id int, value string) SimpleUnsafeCastableString {
	return SimpleUnsafeCastableString{
		SimpleUnsafeCastableBase: SimpleUnsafeCastableBase{
			ID: id,
		},
		ValueA: value,
	}
}

func SimpleUnsafeCastableStringAsRef(id int, value string) *SimpleUnsafeCastableString {
	return &SimpleUnsafeCastableString{
		SimpleUnsafeCastableBase: SimpleUnsafeCastableBase{
			ID: id,
		},
		ValueA: value,
	}
}

func SimpleUnsafeCastableStringAsRandRef() *SimpleUnsafeCastableString {
	return &SimpleUnsafeCastableString{
		SimpleUnsafeCastableBase: SimpleUnsafeCastableBase{
			ID: random.SingleOf[int](),
		},
		ValueA: random.SingleOf[string](),
	}
}

func SimpleUnsafeCastableStringAsRandValue() SimpleUnsafeCastableString {
	return SimpleUnsafeCastableString{
		SimpleUnsafeCastableBase: SimpleUnsafeCastableBase{
			ID: random.SingleOf[int](),
		},
		ValueA: random.SingleOf[string](),
	}
}

type SimpleUnsafeCastableString struct {
	SimpleUnsafeCastableBase
	ValueA string
}

func (a *SimpleUnsafeCastableString) Type() string { return "String" }
func (b *SimpleUnsafeCastableInt) Type() string    { return "Int" }
func (c *SimpleUnsafeCastableFloat) Type() string  { return "Float" }
func (d *SimpleUnsafeCastableBool) Type() string   { return "Bool" }
