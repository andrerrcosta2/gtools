// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package values

import (
	"reflect"
	"testing"

	"github.com/andrerrcosta2/gtools/core/seeders/random"
	"github.com/andrerrcosta2/gtools/core/testlite/assertlite"
	"github.com/andrerrcosta2/gtools/gtests"
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static/structs/models"
	"github.com/andrerrcosta2/gtools/gtests/testingtools"
	"github.com/andrerrcosta2/gtools/gtests/testingtools/config/testlogs"
	"github.com/andrerrcosta2/gtools/gtests/testingtools/themes"
)

// TestAccessField Tests if the method is able to access a field from a struct
//
//		If the field is exported, it should:
//		- Be able to interfaces
//		- Be able to set if the field is addressable
//	 Non exported fields shouldn't be able to interface nor set.
func TestAccessField(t *testing.T) {
	tt := testingtools.LoggersLite(t, testlogs.OnErrors, themes.Color)

	t.Run("should interface but not set exported unaddressable field", func(t *testing.T) {
		ef := models.ExportedFieldsAsValue(1, "")

		err := AccessFieldByName(reflect.ValueOf(ef), "Unaddressable", func(field reflect.Value) {
			// should interface
			assertlite.NoPanic(t, func() {
				val := field.Interface()
				tt.StackLogf("Unaddressable: %d", val)
				assertlite.Equals(t, val, 1)
			})

			// Shouldn't be able to set an unaddressable field
			assertlite.Panic(t, func() {
				field.Set(reflect.ValueOf(2))
				newVal := field.Interface()
				tt.StackErrorf("unaddressable field was setted: %d", newVal)
				t.Fail()
			})
		})
		assertlite.NoError(t, err)
	})

	t.Run("should interface and set exported reference field", func(t *testing.T) {
		x := models.ExportedFieldsAsRef(1, "test")

		err := AccessFieldByName(reflect.ValueOf(x).Elem(), "Addressable", func(field reflect.Value) {
			assertlite.NoPanic(t, func() {
				// should interface
				val := field.Elem().Interface()
				tt.StackLogf("Addressable: %d", val)
				assertlite.Equals(t, val, "test")

				// the field should be addressable
				field.Elem().Set(reflect.ValueOf("test2"))
				newVal := field.Elem().Interface()
				tt.StackLogf("Addressable after set: %d", newVal)
				assertlite.Equals(t, newVal, "test2")
			})
		})
		assertlite.Equals(t, *x.Addressable, "test2")
		assertlite.NoError(t, err)
	})

	t.Run("should not interface not set unexported unaddressable field", func(t *testing.T) {
		u := models.UnexportedFieldsAsValue(1, "test")

		err := AccessFieldByName(reflect.ValueOf(u), "unaddressable", func(field reflect.Value) {
			assertlite.Panic(t, func() {
				// should panic interface
				val := field.Interface()
				tt.StackErrorf("unexported field retrieved its value: '%d'", val)
				t.Fail()
			})

			assertlite.Panic(t, func() {
				// the field shouldn't be addressable
				field.Set(reflect.ValueOf("2"))
				tt.StackErrorf("unaddressable field was setted: %d", field.Interface())
				t.Fail()
			})
		})
		assertlite.Equals(t, u.Unaddressable(), 1)
		assertlite.NoError(t, err)
	})

	t.Run("should not interface nor set unexported reference field", func(t *testing.T) {
		u := models.UnexportedFieldsAsRef(1, "address")

		err := AccessFieldByName(reflect.ValueOf(u).Elem(), "addressable", func(field reflect.Value) {
			assertlite.Panic(t, func() {
				// should not interface
				val := field.Interface()
				tt.StackErrorf("addressable: %d", val)
				t.Fail()
			})

			assertlite.Panic(t, func() {
				// the field shouldn't be accessible
				field.Set(reflect.ValueOf("test2"))
				tt.StackErrorf("not exported field was set")
				t.Fail()
			})
		})
		assertlite.Equals(t, *u.Addressable(), "address")
		assertlite.NoError(t, err)
	})
}

func TestAccessFields(t *testing.T) {
	tt := testingtools.LoggersLite(t, testlogs.OnErrors, themes.Color)

	t.Run("value: exported addressable/unaddressable fields", func(t *testing.T) {
		x := models.ExportedFieldsAsValue(1, "address")
		AccessFields(reflect.ValueOf(x), func(name string, field reflect.Value) {
			if name == "Addressable" {
				assertlite.NoPanic(t, func() {
					// should interface
					value := field.Elem().Interface()
					assertlite.Equals(t, value, "address")

					// should set
					field.Elem().Set(reflect.ValueOf("new"))
				})
			} else {
				assertlite.NoPanic(t, func() {
					// should interface
					u := field.Interface()
					assertlite.Equals(t, u, 1)
				})

				assertlite.Panic(t, func() {
					// shouldn't set unnadressable value
					field.Set(reflect.ValueOf(20))
					tt.StackErrorf("set on unaddressable value: %d", field.Interface())
					tt.Fail()
				})
			}
		})
		assertlite.Equals(t, *x.Addressable, "new")
		assertlite.Equals(t, x.Unaddressable, 1)
	})

	t.Run("reference: exported addressable/unaddressable fields", func(t *testing.T) {
		x := models.ExportedFieldsAsRef(1, "address")
		AccessFields(reflect.ValueOf(x).Elem(), func(name string, field reflect.Value) {
			if name == "Addressable" {
				assertlite.NoPanic(t, func() {
					// should interface
					value := field.Elem().Interface()
					assertlite.Equals(t, value, "address")

					// should set
					field.Elem().Set(reflect.ValueOf("new"))
				})
			} else {
				assertlite.NoPanic(t, func() {
					// should interface
					u := field.Interface()
					assertlite.Equals(t, u, 1)

					// Should set since it came from an addressable struct
					field.Set(reflect.ValueOf(20))
				})
			}
		})
		assertlite.Equals(t, *x.Addressable, "new")
		assertlite.Equals(t, x.Unaddressable, 20)
	})

	t.Run("value: unexported addressable/unaddressable fields", func(t *testing.T) {
		x := models.UnexportedFieldsAsValue(1, "address")
		AccessFields(reflect.ValueOf(x), func(name string, field reflect.Value) {
			if name == "addressable" {
				assertlite.Panic(t, func() {
					// shouldn't interface unexported values
					value := field.Elem().Interface()
					tt.StackErrorf("interface of unexported value: %d", value)
					tt.Fail()
				})

				assertlite.Panic(t, func() {
					// shouldn't set unexported unaddressable value
					field.Elem().Set(reflect.ValueOf("new"))
					tt.StackErrorf("set unexported unaddressable value: %s", field.Elem().String())
					tt.Fail()
				})
			} else {
				assertlite.Panic(t, func() {
					// shouldn't interface unexported values
					value := field.Interface()
					tt.StackErrorf("interface of unexported value: %d", value)
					tt.Fail()
				})

				assertlite.Panic(t, func() {
					// shouldn't set unexported unaddressable value
					field.Set(reflect.ValueOf("new"))
					tt.StackErrorf("set unexported unaddressable value: %s", field.String())
					tt.Fail()
				})
			}
		})
		assertlite.Equals(t, *x.Addressable(), "address")
		assertlite.Equals(t, x.Unaddressable(), 1)
	})

	t.Run("reference: unexported addressable/unaddressable fields", func(t *testing.T) {
		x := models.UnexportedFieldsAsRef(1, "address")
		AccessFields(reflect.ValueOf(x).Elem(), func(name string, field reflect.Value) {
			if name == "addressable" {
				assertlite.Panic(t, func() {
					// shouldn't interface unexported values
					value := field.Elem().Interface()
					tt.StackErrorf("interface of unexported value: %d", value)
					tt.Fail()
				})

				assertlite.Panic(t, func() {
					// shouldn't set unexported unaddressable value
					field.Elem().Set(reflect.ValueOf("new"))
					tt.StackErrorf("set unexported unaddressable value: %s", field.Elem().String())
					tt.Fail()
				})
			} else {
				assertlite.Panic(t, func() {
					// shouldn't interface unexported values
					value := field.Interface()
					tt.StackErrorf("interface of unexported value: %d", value)
					tt.Fail()
				})

				assertlite.Panic(t, func() {
					// shouldn't set unexported unaddressable value
					field.Set(reflect.ValueOf("new"))
					tt.StackErrorf("set unexported unaddressable value: %s", field.String())
					tt.Fail()
				})
			}
		})
		assertlite.Equals(t, *x.Addressable(), "address")
		assertlite.Equals(t, x.Unaddressable(), 1)
	})
}

// TestFields Tests if the method is able to extract all fields from a struct,
//
//	This test should be able to extract only exported fields
func TestFields(t *testing.T) {
	structs := gtests.Structs.Fuzz().Categories().Values()

	t.Run("Structs with only exported fields", func(t *testing.T) {
		structs.OfExportedFieldsOnly().Each(func(seed any) {
			rv := reflect.ValueOf(seed)
			fields := Fields(rv)
			assertlite.True(t, len(fields) == rv.NumField(),
				gtests.ErrorMismatchValues("Number of fields mismatch", len(fields), rv.NumField()))
		})
	})

	t.Run("Structs with unexported fields", func(t *testing.T) {
		structs.WithUnexportedFields().Each(func(seed any) {
			rv := reflect.ValueOf(seed)
			fields := Fields(rv)
			assertlite.False(t, len(fields) == 0,
				gtests.ErrorMismatchValues("Number of fields matched", len(fields), 0))
		})
	})
}

// TestNilFields tests the method NilFields
//
// After each test it should:
//   - Return the number of nullable fields that are nil as key-pair map
//   - Both exported and unexported fields must be tracked
func TestNilFields(t *testing.T) {
	t.Run("no nil fields, half nullables, half unexported", func(t *testing.T) {
		value := models.MixedExportedUnexportedFieldsAsValue("addrx", "addru", "unaddrx", "unaddru")
		nf := NilFields(reflect.ValueOf(value))
		assertlite.True(t, len(nf) == 0)
	})

	t.Run("zero struct, half nullables, half unexported", func(t *testing.T) {
		zero := models.MixedExportedUnexportedFields{}
		nf := NilFields(reflect.ValueOf(zero))
		assertlite.True(t, len(nf) == 2)
		assertlite.AreTrue(t, []string{"AddressableExported", "addressableUnexported"}, func(s string) bool {
			_, ok := nf[s]
			return ok
		})
	})

	t.Run("no nil fields, half nullables, all unexported", func(t *testing.T) {
		value := models.UnexportedFieldsAsValue(1, "addru")
		nf := NilFields(reflect.ValueOf(value))
		assertlite.True(t, len(nf) == 0)
	})

	t.Run("zero struct, half nullables, all unexported", func(t *testing.T) {
		zero := models.UnexportedFields{}
		rv := reflect.ValueOf(zero)
		nf := NilFields(rv)
		assertlite.True(t, len(nf) == rv.NumField()/2)
		assertlite.AreTrue(t, []string{"addressable"}, func(s string) bool {
			_, ok := nf[s]
			return ok
		})
	})

	t.Run("no nil fields, all nullables, all exported", func(t *testing.T) {
		value := models.ExportedAddressableFieldsAsValue(1, 1.0, true, "test")
		nf := NilFields(reflect.ValueOf(value))
		assertlite.True(t, len(nf) == 0)
	})

	t.Run("zero struct, all nullables, all exported", func(t *testing.T) {
		value := models.ExportedAddressableFields{}
		nf := NilFields(reflect.ValueOf(value))
		assertlite.True(t, len(nf) == 4)
	})
}

// TestNoNilFields tests the method NoNilFields
//
// After each test it should:
//   - Return the number of fields that are not nil as key-pair map
//   - Both exported and unexported fields must be tracked
//   - Non-addressable fields must be always retrieved.
func TestNoNilFields(t *testing.T) {
	t.Run("no nil fields, half nullables, half exported", func(t *testing.T) {
		value := models.MixedExportedUnexportedFieldsAsValue("addrx", "addru", "unaddrx", "unaddru")
		nf := NoNilFields(reflect.ValueOf(value))
		assertlite.True(t, len(nf) == 4)
	})

	t.Run("zero struct, half nullables, half exported", func(t *testing.T) {
		value := models.MixedExportedUnexportedFields{}
		nf := NoNilFields(reflect.ValueOf(value))
		assertlite.True(t, len(nf) == 2)
		assertlite.AreTrue(t, []string{"UnaddressableExported", "unaddressableUnexported"}, func(s string) bool {
			_, ok := nf[s]
			return ok
		})
	})

	t.Run("no nil fields, half nullables, all unexported", func(t *testing.T) {
		value := models.UnexportedFieldsAsValue(1, "addru")
		rv := reflect.ValueOf(value)
		nf := NoNilFields(rv)
		assertlite.True(t, len(nf) == rv.NumField())
	})

	t.Run("zero struct, half nullables, all unexported", func(t *testing.T) {
		zero := models.UnexportedFields{}
		rv := reflect.ValueOf(zero)
		nf := NoNilFields(rv)
		assertlite.True(t, len(nf) == rv.NumField()/2)
		assertlite.AreTrue(t, []string{"unaddressable"}, func(s string) bool {
			_, ok := nf[s]
			return ok
		})
	})

	t.Run("no nil fields, all fields as nullables", func(t *testing.T) {
		value := models.ExportedAddressableFieldsAsValue(1, 1.0, true, "test")
		rv := reflect.ValueOf(value)
		nf := NoNilFields(rv)
		assertlite.True(t, len(nf) == rv.NumField(), "expected nil fields map "+
			"to have len compare to '%d' but got '%d'", rv.NumField(), len(nf))
		assertlite.NoNilFields(t, true, value)
	})

	t.Run("zero struct, all fields are nullables", func(t *testing.T) {
		value := models.ExportedAddressableFields{}
		nf := NoNilFields(reflect.ValueOf(value))
		assertlite.True(t, len(nf) == 0, "expected nil fields map "+
			"to have len compare to '0' but got '%d'", len(nf))
		assertlite.AllFieldsAreNil(t, true, value)
	})
}

// TestRideFields tests the method RideFields
//
// This test should assert:
//   - All fields of any struct are visited - exported and unexported.
//   - The method retrieves its name and its reflect.Value
func TestRideFields(t *testing.T) {
	t.Run("unaddressable struct", func(t *testing.T) {
		gtests.Structs.Fuzz().Categories().Values().All().Each(func(seed any) {
			v := reflect.ValueOf(seed)
			count := 0
			RideFields(v, func(idx int, field reflect.Value) bool {
				assertlite.True(t, idx == count, "expected index count to be "+
					"%d but got %d", count, idx)
				assertlite.NotNil(t, field)
				count++
				return true
			})
			assertlite.True(t, count == v.NumField())
		})
	})
}

func TestSetField(t *testing.T) {
	t.Run("unaddressable structs", func(t *testing.T) {
		gtests.Structs.Fuzz().Categories().Values().All().Each(func(seed any) {
			s := reflect.ValueOf(seed)
			// Get field name and type
			for i := 0; i < s.NumField(); i++ {
				fieldType := s.Type().Field(i)
				fieldName := fieldType.Name
				fieldValue := s.Field(i)
				fieldNewValue := random.Reflect(fieldType.Type, 1).At(0)
				if fieldValue.CanAddr() {
					err := SetField(s, fieldName, fieldNewValue)
					assertlite.NoError(t, err)
					assertlite.Equals(t, fieldNewValue, s.FieldByName(fieldName))
				}
			}

		})
	})
}

// TestUnsafeGetAllFields Tests if the method is able to extract all fields from a struct using
// the unsafe package
func TestUnsafeGetAllFields(t *testing.T) {
	structs := gtests.Structs.Fuzz().Categories().Values()

	t.Run("Structs with exported fields", func(t *testing.T) {
		structs.OfExportedFieldsOnly().Each(func(seed any) {
			rv := reflect.ValueOf(seed)
			fields := UnsafeGetAllFields(rv)
			assertlite.True(t, len(fields) == rv.NumField(),
				gtests.ErrorMismatchValues("Number of fields mismatch", len(fields), rv.NumField()))
			for k, v := range fields {
				assertlite.NotNil(t, v, "field '%s' is nil", k)
			}
		})
	})

	t.Run("Structs with unexported fields", func(t *testing.T) {
		structs.WithUnexportedFields().Each(func(seed any) {
			rv := reflect.ValueOf(seed)
			fields := UnsafeGetAllFields(rv)
			assertlite.True(t, len(fields) == rv.NumField(),
				gtests.ErrorMismatchValues("Number of fields mismatch", len(fields), rv.NumField()))
			for k, v := range fields {
				assertlite.NotNil(t, v, "field '%s' is nil", k)
			}
		})
	})
}

// TestUnsafeRideFields Tests if the method is able to ride over all fields from a struct using
// the unsafe package
//
//   - All fields (exported/unexported) should be accessible as unsafe ptrs.
//func TestUnsafeRideFields(t *testing.T) {
//	values := gtests.Structs.Fuzz().Categories().Values()
//	refs := gtests.Structs.Fuzz().Categories().Refs()
//
//	t.Run("Unsafe Ride Fields: values", func(t *testing.T) {
//
//	})
//}
