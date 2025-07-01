// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package lite

import (
	"github.com/andrerrcosta2/gtools/core/format/fmx"
	"github.com/andrerrcosta2/gtools/core/testlite/assertlite"
	"github.com/andrerrcosta2/gtools/core/testlite/testseed"
	"reflect"
	"testing"
)

func TestRandAny(t *testing.T) {
	t.Run("RandAny", func(t *testing.T) {
		for i := 0; i < 1000; i++ {
			value := RandAny()
			assertlite.NotNil(t, value, "expecting not nil, but got %v", value)
		}
	})
}

// TestRandOf tests the private random generation of a single value.
//
// This test must assert that after each call:
//   - No nil values are returned.
//   - The returned value must be of the same type passed as the reflect.Value argument.
func TestRandOf(t *testing.T) {
	t.Run("Primitives", func(t *testing.T) {
		d1 := RandOf(reflect.TypeOf(true))
		assertlite.IsTypeOf[bool](t, d1)

		d2 := RandOf(reflect.TypeOf(0))
		assertlite.IsTypeOf[int](t, d2)

		d3 := RandOf(reflect.TypeOf(int8(0)))
		assertlite.IsTypeOf[int8](t, d3)

		d4 := RandOf(reflect.TypeOf(int16(0)))
		assertlite.IsTypeOf[int16](t, d4)

		d5 := RandOf(reflect.TypeOf(int32(0)))
		assertlite.IsTypeOf[int32](t, d5)

		d6 := RandOf(reflect.TypeOf(int64(0)))
		assertlite.IsTypeOf[int64](t, d6)

		d7 := RandOf(reflect.TypeOf(uint(0)))
		assertlite.IsTypeOf[uint](t, d7)

		d8 := RandOf(reflect.TypeOf(uint8(0)))
		assertlite.IsTypeOf[uint8](t, d8)

		d9 := RandOf(reflect.TypeOf(uint16(0)))
		assertlite.IsTypeOf[uint16](t, d9)

		d10 := RandOf(reflect.TypeOf(uint32(0)))
		assertlite.IsTypeOf[uint32](t, d10)

		d11 := RandOf(reflect.TypeOf(uint64(0)))
		assertlite.IsTypeOf[uint64](t, d11)

		d12 := RandOf(reflect.TypeOf(float32(0)))
		assertlite.IsTypeOf[float32](t, d12)

		d13 := RandOf(reflect.TypeOf(float64(0)))
		assertlite.IsTypeOf[float64](t, d13)

		d14 := RandOf(reflect.TypeOf(complex64(0)))
		assertlite.IsTypeOf[complex64](t, d14)

		d15 := RandOf(reflect.TypeOf(complex128(0)))
		assertlite.IsTypeOf[complex128](t, d15)

		d16 := RandOf(reflect.TypeOf(""))
		assertlite.IsTypeOf[string](t, d16)
	})

	// Values can't be nil, but must check its fields.
	t.Run("Struct Values", func(t *testing.T) {
		d1 := RandOf(reflect.TypeOf(testseed.NewSortableValue("test", 1)))
		assertlite.IsTypeOf[testseed.SortableValue](t, d1)
		assertlite.NoNilFields(t, true, d1)

		d2 := RandOf(reflect.TypeOf(testseed.NewSortableRef("test", 1)))
		assertlite.IsTypeOf[*testseed.SortableRef](t, d2)
		assertlite.NoNilFields(t, true, d2)

		d3 := RandOf(reflect.TypeOf(testseed.NewComparableValue("test", 1)))
		assertlite.IsTypeOf[testseed.ComparableValue](t, d3)
		assertlite.NoNilFields(t, true, d3)

		d4 := RandOf(reflect.TypeOf(testseed.NewComparableRef("test", 1)))
		assertlite.IsTypeOf[*testseed.ComparableRef](t, d4)
		assertlite.NoNilFields(t, true, d4)

		d5 := RandOf(reflect.TypeOf(testseed.NewStructStringer("test")))
		assertlite.IsTypeOf[*testseed.StructStringer](t, d5)
		assertlite.NoNilFields(t, true, d5)
		fmx.Printf("d5: %v", d5)
	})
}
