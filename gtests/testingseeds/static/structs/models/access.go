// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package models

import "github.com/andrerrcosta2/gtools/core/seeders/random"

var ExportedAddressableFieldsZeroInst = new(ExportedAddressableFields)

func ExportedAddressableFieldsAsValue(ai int, af float64, ab bool, as string) ExportedAddressableFields {
	return ExportedAddressableFields{
		AddressableInt:    &ai,
		AddressableFloat:  &af,
		AddressableBool:   &ab,
		AddressableString: &as,
	}
}

func ExportedAddressableFieldsAsRef(ai int, af float64, ab bool, as string) *ExportedAddressableFields {
	r := ExportedAddressableFieldsAsValue(ai, af, ab, as)
	return &r
}

func ExportedAddressableFieldsAsRandValue() ExportedAddressableFields {
	return random.Struct[ExportedAddressableFields](1).At(0)
}

func ExportedAddressableFieldsAsRandRef() *ExportedAddressableFields {
	r := random.Struct[ExportedAddressableFields](1).At(0)
	return &r
}

type ExportedAddressableFields struct {
	AddressableInt    *int
	AddressableFloat  *float64
	AddressableBool   *bool
	AddressableString *string
}

var ExportedFieldsZeroInst = new(ExportedFields)

func ExportedFieldsAsValue(unaddr int, addr string) ExportedFields {
	return ExportedFields{
		Addressable:   &addr,
		Unaddressable: unaddr,
	}
}

func ExportedFieldsAsRef(unaddr int, addr string) *ExportedFields {
	return &ExportedFields{
		Addressable:   &addr,
		Unaddressable: unaddr,
	}
}

func ExportedFieldsAsRandValue() ExportedFields {
	a := random.SingleOf[string]()
	u := random.SingleOf[int]()
	return ExportedFields{
		Addressable:   &a,
		Unaddressable: u,
	}
}

func ExportedFieldsAsRandRef() *ExportedFields {
	a := random.SingleOf[string]()
	u := random.SingleOf[int]()
	return &ExportedFields{
		Addressable:   &a,
		Unaddressable: u,
	}
}

type ExportedFields struct {
	Addressable   *string
	Unaddressable int
}

var ExportedManyFieldsZeroInst = new(ExportedManyFields)

func ExportedManyFieldsAsValue(ai, ui int, ab, ub bool, af, uf float64, as, us string) ExportedManyFields {
	return ExportedManyFields{
		AddressableInt:      &ai,
		UnaddressableInt:    ui,
		AddressableBool:     &ab,
		UnaddressableBool:   ub,
		AddressableFloat:    &af,
		UnaddressableFloat:  uf,
		AddressableString:   &as,
		UnaddressableString: us,
	}
}

func ExportedManyFieldsAsRef(ai, ui int, ab, ub bool, af, uf float64, as, us string) *ExportedManyFields {
	return &ExportedManyFields{
		AddressableInt:      &ai,
		UnaddressableInt:    ui,
		AddressableBool:     &ab,
		UnaddressableBool:   ub,
		AddressableFloat:    &af,
		UnaddressableFloat:  uf,
		AddressableString:   &as,
		UnaddressableString: us,
	}
}

func ExportedManyFieldsAsRandValue() ExportedManyFields {
	return random.Struct[ExportedManyFields](1).At(0)
}

func ExportedManyFieldsAsRandRef() *ExportedManyFields {
	r := random.Struct[ExportedManyFields](1).At(0)
	return &r
}

type ExportedManyFields struct {
	AddressableInt      *int
	AddressableFloat    *float64
	AddressableBool     *bool
	AddressableString   *string
	UnaddressableInt    int
	UnaddressableFloat  float64
	UnaddressableBool   bool
	UnaddressableString string
}

var ExportedUnaddressableFieldsZeroInst = new(ExportedUnaddressableFields)

func ExportedUnaddressableFieldsAsValue(ui int, uf float64, ub bool, us string) ExportedUnaddressableFields {
	return ExportedUnaddressableFields{
		UnaddressableInt:    ui,
		UnaddressableFloat:  uf,
		UnaddressableString: us,
		UnaddressableBool:   ub,
	}
}

func ExportedUnaddressableFieldsAsRef(ui int, uf float64, ub bool, us string) *ExportedUnaddressableFields {
	return &ExportedUnaddressableFields{
		UnaddressableInt:    ui,
		UnaddressableFloat:  uf,
		UnaddressableString: us,
		UnaddressableBool:   ub,
	}
}

func ExportedUnaddressableFieldsAsRandValue() ExportedUnaddressableFields {
	return random.SingleOf[ExportedUnaddressableFields]()
}

func ExportedUnaddressableFieldsAsRandRef() *ExportedUnaddressableFields {
	r := ExportedUnaddressableFieldsAsRandValue()
	return &r
}

type ExportedUnaddressableFields struct {
	UnaddressableInt    int
	UnaddressableFloat  float64
	UnaddressableBool   bool
	UnaddressableString string
}

var MixedExportedUnexportedFieldsZeroInst = new(MixedExportedUnexportedFields)

func MixedExportedUnexportedFieldsAsValue(ae, au, ue, uu string) MixedExportedUnexportedFields {
	return MixedExportedUnexportedFields{
		AddressableExported:     &ae,
		addressableUnexported:   &au,
		UnaddressableExported:   ue,
		unaddressableUnexported: uu,
	}
}

func MixedExportedUnexportedFieldsAsRef(ae, au, ue, uu string) *MixedExportedUnexportedFields {
	return &MixedExportedUnexportedFields{
		AddressableExported:     &ae,
		addressableUnexported:   &au,
		UnaddressableExported:   ue,
		unaddressableUnexported: uu,
	}
}

func MixedExportedUnexportedFieldsAsRandValue() MixedExportedUnexportedFields {
	return random.Struct[MixedExportedUnexportedFields](1).At(0)
}

func MixedExportedUnexportedFieldsAsRandRef() *MixedExportedUnexportedFields {
	r := random.Struct[MixedExportedUnexportedFields](1).At(0)
	return &r
}

type MixedExportedUnexportedFields struct {
	AddressableExported     *string
	addressableUnexported   *string
	UnaddressableExported   string
	unaddressableUnexported string
}

func (m MixedExportedUnexportedFields) AddressableUnexported() *string {
	return m.addressableUnexported
}

func (m MixedExportedUnexportedFields) UnaddressableUnexported() string {
	return m.unaddressableUnexported
}

var UnexportedAddressableFieldsZeroInst = new(UnexportedAddressableFields)

func UnexportedAddressableFieldsAsValue(ai int, af float64, ab bool, as string) UnexportedAddressableFields {
	return UnexportedAddressableFields{
		addressableInt:    &ai,
		addressableFloat:  &af,
		addressableBool:   &ab,
		addressableString: &as,
	}
}

func UnexportedAddressableFieldsAsRef(ai int, af float64, ab bool, as string) *UnexportedAddressableFields {
	return &UnexportedAddressableFields{
		addressableInt:    &ai,
		addressableFloat:  &af,
		addressableBool:   &ab,
		addressableString: &as,
	}
}

func UnexportedAddressableFieldsAsRandValue() UnexportedAddressableFields {
	return random.SingleOf[UnexportedAddressableFields]()
}

func UnexportedAddressableFieldsAsRandRef() *UnexportedAddressableFields {
	r := random.SingleOf[UnexportedAddressableFields]()
	return &r
}

type UnexportedAddressableFields struct {
	addressableInt    *int
	addressableFloat  *float64
	addressableBool   *bool
	addressableString *string
}

func (u UnexportedAddressableFields) AddressableInt() *int {
	return u.addressableInt
}

func (u UnexportedAddressableFields) AddressableFloat() *float64 {
	return u.addressableFloat
}

func (u UnexportedAddressableFields) AddressableBool() *bool {
	return u.addressableBool
}

func (u UnexportedAddressableFields) AddressableString() *string {
	return u.addressableString
}

var UnexportedFieldsZeroInst = new(UnexportedFields)

func UnexportedFieldsAsValue(unaddr int, addr string) UnexportedFields {
	return UnexportedFields{
		addressable:   &addr,
		unaddressable: unaddr,
	}
}

func UnexportedFieldsAsRef(unaddr int, addr string) *UnexportedFields {
	return &UnexportedFields{
		addressable:   &addr,
		unaddressable: unaddr,
	}
}

func UnexportedFieldsAsRandValue() UnexportedFields {
	a := random.SingleOf[string]()
	u := random.SingleOf[int]()
	return UnexportedFields{
		addressable:   &a,
		unaddressable: u,
	}
}

func UnexportedFieldsAsRandRef() *UnexportedFields {
	a := random.SingleOf[string]()
	u := random.SingleOf[int]()
	return &UnexportedFields{
		addressable:   &a,
		unaddressable: u,
	}
}

type UnexportedFields struct {
	addressable   *string
	unaddressable int
}

func (u UnexportedFields) Addressable() *string {
	return u.addressable
}

func (u UnexportedFields) Unaddressable() int {
	return u.unaddressable
}

var UnexportedManyFieldsZeroInst = new(UnexportedManyFields)

func UnexportedManyFieldsAsValue(ab, ub bool, ai, ui int, af, uf float64, as, us string) UnexportedManyFields {
	return UnexportedManyFields{
		addressableBool:     &ab,
		unaddressableBool:   ub,
		addressableInt:      &ai,
		unaddressableInt:    ui,
		addressableFloat:    &af,
		unaddressableFloat:  uf,
		addressableString:   &as,
		unaddressableString: us,
	}
}

func UnexportedManyFieldsAsRef(ab, ub bool, ai, ui int, af, uf float64, as, us string) *UnexportedManyFields {
	return &UnexportedManyFields{
		addressableBool:     &ab,
		unaddressableBool:   ub,
		addressableInt:      &ai,
		unaddressableInt:    ui,
		addressableFloat:    &af,
		unaddressableFloat:  uf,
		addressableString:   &as,
		unaddressableString: us,
	}
}

func UnexportedManyFieldsAsRandValue() UnexportedManyFields {
	return random.Struct[UnexportedManyFields](1).At(0)
}

func UnexportedManyFieldsAsRandRef() *UnexportedManyFields {
	r := random.Struct[UnexportedManyFields](1).At(0)
	return &r
}

type UnexportedManyFields struct {
	addressableBool     *bool
	addressableFloat    *float64
	addressableInt      *int
	addressableString   *string
	unaddressableBool   bool
	unaddressableFloat  float64
	unaddressableInt    int
	unaddressableString string
}

func (u UnexportedManyFields) AddressableBool() *bool {
	return u.addressableBool
}

func (u UnexportedManyFields) AddressableFloat() *float64 {
	return u.addressableFloat
}

func (u UnexportedManyFields) AddressableInt() *int {
	return u.addressableInt
}

func (u UnexportedManyFields) AddressableString() *string {
	return u.addressableString
}

func (u UnexportedManyFields) UnaddressableBool() bool {
	return u.unaddressableBool
}

func (u UnexportedManyFields) UnaddressableFloat() float64 {
	return u.unaddressableFloat
}

func (u UnexportedManyFields) UnaddressableInt() int {
	return u.unaddressableInt
}

func (u UnexportedManyFields) UnaddressableString() string {
	return u.unaddressableString
}

var UnexportedUnaddressableFieldsZeroInst = new(UnexportedUnaddressableFields)

func UnexportedUnaddressableFieldsAsValue(ai int, af float64, ab bool, as string) UnexportedUnaddressableFields {
	return UnexportedUnaddressableFields{
		addressableInt:    ai,
		addressableFloat:  af,
		addressableBool:   ab,
		addressableString: as,
	}
}

func UnexportedUnaddressableFieldsAsRef(ai int, af float64, ab bool, as string) *UnexportedUnaddressableFields {
	return &UnexportedUnaddressableFields{
		addressableInt:    ai,
		addressableFloat:  af,
		addressableBool:   ab,
		addressableString: as,
	}
}

func UnexportedUnaddressableFieldsAsRandValue() UnexportedUnaddressableFields {
	return random.SingleOf[UnexportedUnaddressableFields]()
}

func UnexportedUnaddressableFieldsAsRandRef() *UnexportedUnaddressableFields {
	r := random.SingleOf[UnexportedUnaddressableFields]()
	return &r
}

type UnexportedUnaddressableFields struct {
	addressableInt    int
	addressableFloat  float64
	addressableBool   bool
	addressableString string
}

func (u UnexportedUnaddressableFields) AddressableInt() int {
	return u.addressableInt
}

func (u UnexportedUnaddressableFields) AddressableFloat() float64 {
	return u.addressableFloat
}

func (u UnexportedUnaddressableFields) AddressableBool() bool {
	return u.addressableBool
}

func (u UnexportedUnaddressableFields) AddressableString() string {
	return u.addressableString
}
