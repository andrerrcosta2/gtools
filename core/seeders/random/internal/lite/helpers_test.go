// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package lite

import (
	"github.com/andrerrcosta2/gtools/core/format/fmx"
	"github.com/andrerrcosta2/gtools/core/seeders/random/internal/reflectutils"
	"github.com/andrerrcosta2/gtools/core/seeders/random/internal/reflectutils/reflectrand"
	"github.com/andrerrcosta2/gtools/core/testlite/assertlite"
	"github.com/andrerrcosta2/gtools/core/testlite/testseed"
	"reflect"
	"slices"
	"testing"
)

const (
	Distributions = 100_000
	Deviation     = 0.15 // 15% acceptable deviation
	Loops         = 5000
)

// TestRandAny tests the RandAny method generation of random values
//
// this test must assert:
//   - No nil values
//   - the generated types are between the deviation limit of the distribution
func TestRandAny(t *testing.T) {
	t.Run("shouldn't generate nil values", func(t *testing.T) {
		for i := 0; i < Loops; i++ {
			value := RandAny()
			assertlite.NotNil(t, value, "expecting not nil, but got %v", value)
		}
	})

	t.Run("distribution should be under deviation threshold", func(t *testing.T) {
		counts := make(map[reflect.Kind]int)
		buckets := len(reflectutils.InterfaceKinds)

		for i := 0; i < Distributions; i++ {
			value := RandAny()
			counts[reflect.TypeOf(value).Kind()]++
		}

		expected := float64(Distributions) / float64(buckets)
		lowerBound := expected * (1 - Deviation)
		upperBound := expected * (1 + Deviation)

		for _, kind := range reflectutils.InterfaceKinds {
			count := float64(counts[kind])
			if count < lowerBound || count > upperBound {
				t.Errorf("kind '%v': expected around %.0f ±%.0f (%.2f%%), got %d",
					kind, expected, expected*Deviation, Deviation*100, counts[kind])
			}
		}
	})
}

// TestRandArray tests the RandArray method generation of random Arrays
//
// this test should assert
//   - no nil values are generated
//   - all values are reflect.Array
//   - the generated types are between the deviation limit of the distribution
func TestRandArray(t *testing.T) {
	t.Run("shouldn't generate nil values", func(t *testing.T) {
		for i := 0; i < Loops; i++ {
			value := RandArray()
			assertlite.NotNil(t, value, "expecting not nil, but got %v", value)
			assertlite.True(t, reflect.TypeOf(value).Kind() == reflect.Array, "expecting kind array, "+
				"but got %v", reflect.TypeOf(value))
		}
	})

	t.Run("distribution should be under deviation threshold", func(t *testing.T) {
		counts := make(map[reflect.Kind]int)
		buckets := len(reflectutils.RandomizableKinds)
		for i := 0; i < Distributions; i++ {
			value := RandArray()
			counts[reflect.TypeOf(value).Elem().Kind()]++
		}
		expected := float64(Distributions) / float64(buckets)
		lowerBound := expected * (1 - Deviation)
		upperBound := expected * (1 + Deviation)
		for _, kind := range reflectutils.RandomizableKinds {
			count := float64(counts[kind])
			if count < lowerBound || count > upperBound {
				t.Errorf("kind %v: expected around %.0f ±%.0f (%.2f%%), got %d",
					kind, expected, expected*Deviation, Deviation*100, counts[kind])
			}
		}
	})
}

// TestRandArrayOf tests the method RandArrayOf generation of random arrays based on a given type
//
// this test must assert:
//   - no nil values are generated
//   - all values are of kind reflect.Array
//   - the generated values must have its elements of the same type from its type source
func TestRandArrayOf(t *testing.T) {
	t.Run("shouldn't generate nil values", func(t *testing.T) {
		for i := 0; i < Loops; i++ {
			src := reflectrand.Array()
			value := RandArrayOf(src)
			typ := reflect.TypeOf(value)
			assertlite.NotNil(t, value, "expecting not nil, but got %v", value)
			assertlite.True(t, reflect.TypeOf(value).Kind() == reflect.Array, "expecting kind array, "+
				"but got %v", reflect.TypeOf(value))
			assertlite.True(t, typ.Elem().Kind() == src.Elem().Kind(), "expecting generated array element '%v' "+
				"to be of type '%v'", typ.Elem().Kind(), src.Elem().Kind())
			assertlite.True(t, typ.String() == src.String(), "expecting generated array element '%v' "+
				"to be stringified as '%v'", typ.String(), src.String())
		}
	})

}

// TestRandChan tests the RandChan method generation of random channels
//
// this test must assert:
//   - no nil values are generated
//   - all values must be of kind reflect.Chan
//   - the generated types are between the deviation limit of the distribution
func TestRandChan(t *testing.T) {
	t.Run("shouldn't generate nil values", func(t *testing.T) {
		for i := 0; i < Loops; i++ {
			value := RandChan()
			assertlite.NotNil(t, value, "expecting not nil, but got %v", value)
			assertlite.True(t, reflect.TypeOf(value).Kind() == reflect.Chan, "expecting kind channel, "+
				"but got %v", reflect.TypeOf(value))
		}
	})

	t.Run("distribution should be under deviation threshold", func(t *testing.T) {
		counts := make(map[reflect.Kind]int)
		buckets := len(reflectutils.RandomizableKinds)

		for i := 0; i < Distributions; i++ {
			value := RandChan()
			counts[reflect.TypeOf(value).Elem().Kind()]++
		}

		expected := float64(Distributions) / float64(buckets)
		lowerBound := expected * (1 - Deviation)
		upperBound := expected * (1 + Deviation)
		for _, kind := range reflectutils.RandomizableKinds {
			count := float64(counts[kind])
			if count < lowerBound || count > upperBound {
				t.Errorf("kind %v: expected around %.0f ±%.0f (%.2f%%), got %d",
					kind, expected, expected*Deviation, Deviation*100, counts[kind])
			}
		}
	})
}

// TestRandChanOf tests the method RandChanOf generation of random channels from a given type
//
// this test must assert:
//   - no nil values are generated
//   - all values must be of kind reflect.Chan
//   - the generated values must have its elements of the same type from its type source
func TestRandChanOf(t *testing.T) {
	t.Run("shouldn't generate nil values", func(t *testing.T) {
		for i := 0; i < Loops; i++ {
			src := reflectrand.Chan()
			value := RandChanOf(src)
			typ := reflect.TypeOf(value)
			assertlite.NotNil(t, value, "expecting not nil, but got %v", value)
			assertlite.True(t, typ.Kind() == reflect.Chan, "expecting kind channel, "+
				"but got %v", reflect.TypeOf(value))
			assertlite.True(t, typ.Elem().Kind() == src.Elem().Kind(), "expecting kind channel, "+
				"but got %v", typ.Elem().Kind())
			assertlite.True(t, typ.String() == src.String(), "expecting generated channel '%s' to be "+
				"stringified as '%s' ", typ.String(), src.String())
		}
	})
}

// TestRandCmp tests the method RandCmp generation of random comparable values
//
// this test must assert:
//   - no nil values are generated
//   - all values must be comparable
//   - the generated types are between the deviation limit of the distribution
func TestRandCmp(t *testing.T) {
	t.Run("shouldn't generate nil values", func(t *testing.T) {
		for i := 0; i < Loops; i++ {
			value := RandCmp()
			typ := reflect.TypeOf(value)
			assertlite.NotNil(t, value, "expecting not nil, but got %v", value)
			assertlite.True(t, slices.ContainsFunc(reflectutils.CmpKinds, func(kind reflect.Kind) bool {
				return kind == typ.Kind()
			}))
			assertlite.True(t, typ.Comparable(), "expected a comparable type but got %v", typ.Kind())
		}
	})

	t.Run("distribution should be under deviation threshold", func(t *testing.T) {
		counts := make(map[reflect.Kind]int)
		buckets := len(reflectutils.CmpKinds)
		for i := 0; i < Distributions; i++ {
			value := RandCmp()
			counts[reflect.TypeOf(value).Kind()]++
		}
		expected := float64(Distributions) / float64(buckets)
		lowerBound := expected * (1 - Deviation)
		upperBound := expected * (1 + Deviation)
		for _, kind := range reflectutils.CmpKinds {
			count := float64(counts[kind])
			if count < lowerBound || count > upperBound {
				t.Errorf("kind %v: expected around %.0f ±%.0f (%.2f%%), got %d",
					kind, expected, expected*Deviation, Deviation*100, counts[kind])
			}
		}
	})
}

// TestRandFunc tests the method RandFunc generation of random functions
//
// this test must assert:
//   - no nil values are generated
//   - all values are of the kind reflect.Func
func TestRandFunc(t *testing.T) {
	t.Run("shouldn't generate nil values", func(t *testing.T) {
		for i := 0; i < Loops; i++ {
			value := RandFunc()
			typ := reflect.TypeOf(value)
			assertlite.NotNil(t, value, "expecting not nil, but got %v", value)
			assertlite.True(t, typ.Kind() == reflect.Func, "expecting kind function, "+
				"but got %v", typ.Kind())
		}
	})
}

// TestRandFuncOf tests the method RandFuncOf generation of random functions from a given type
//
// this test should assert:
//   - No nil values are generated
//   - the values have the same type of its random source type
func TestRandFuncOf(t *testing.T) {
	t.Run("shouldn't generate nil values", func(t *testing.T) {
		for i := 0; i < Loops; i++ {
			src := reflectrand.Func()
			value := RandFuncOf(src)
			typ := reflect.TypeOf(value)
			assertlite.NotNil(t, value, "expecting not nil, but got %v", value)
			assertlite.True(t, typ.Kind() == reflect.Func, "expecting kind function, "+
				"but got %v", typ.Kind())
			assertlite.True(t, src.String() == typ.String(), "expecting generated channel '%s' to be "+
				"stringified as '%s'", typ.String(), src.String())
		}
	})
}

// TestRandGiven tests the method RandGiven generation of random value based on a given slice
//
// this test should assert
//   - the generated types are between the deviation limit of the distribution
func TestRandGiven(t *testing.T) {
	t.Run("distribution should be under deviation threshold", func(t *testing.T) {
		counts := make(map[int]int)
		data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
			11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
		buckets := len(data)
		for i := 0; i < Distributions; i++ {
			value := RandGiven(data)
			counts[value]++
		}
		expected := float64(Distributions) / float64(buckets)
		lowerBound := expected * (1 - Deviation)
		upperBound := expected * (1 + Deviation)

		for _, kind := range data {
			count := float64(counts[kind])
			if count < lowerBound || count > upperBound {
				t.Errorf("kind %v: expected around %.0f ±%.0f (%.2f%%), got %d",
					kind, expected, expected*Deviation, Deviation*100, counts[kind])
			}
		}

	})
}

// TestRandMap tests the method RandMap generation of random map types
//
// This test must assert
//   - no nil values are generated
//   - all values must be of kind reflect.Map
//   - the generated values must have its elements of the same type from its type source
func TestRandMap(t *testing.T) {
	t.Run("shouldn't generate nil values", func(t *testing.T) {
		for i := 0; i < Loops; i++ {
			value := RandMap()
			typ := reflect.TypeOf(value)
			assertlite.NotNil(t, value, "expecting not nil, but got %v", value)
			assertlite.True(t, typ.Kind() == reflect.Map, "expecting kind map, "+
				"but got %v", typ.Kind())
		}
	})

	t.Run("distributions should be under deviation threshold", func(t *testing.T) {
		keys := make(map[reflect.Kind]int)
		values := make(map[reflect.Kind]int)
		keyBuckets := len(reflectutils.CmpKinds)
		valueBuckets := len(reflectutils.RandomizableKinds)

		for i := 0; i < Distributions; i++ {
			value := RandMap()
			kk := reflect.TypeOf(value).Key().Kind()
			vk := reflect.TypeOf(value).Elem().Kind()
			keys[kk]++
			values[vk]++
		}

		expectedKeys := float64(Distributions) / float64(keyBuckets)
		lowerBound := expectedKeys * (1 - Deviation)
		upperBound := expectedKeys * (1 + Deviation)

		for _, kind := range reflectutils.CmpKinds {
			count := float64(keys[kind])
			if count < lowerBound || count > upperBound {
				t.Errorf("key kind '%v': expected around %.0f ±%.0f (%.2f%%), got %d",
					kind, expectedKeys, expectedKeys*Deviation, Deviation*100, keys[kind])
			}
		}

		expectedValues := float64(Distributions) / float64(valueBuckets)
		lowerBound = expectedValues * (1 - Deviation)
		upperBound = expectedValues * (1 + Deviation)

		for _, kind := range reflectutils.RandomizableKinds {
			count := float64(values[kind])
			if count < lowerBound || count > upperBound {
				t.Errorf("value kind '%v': expected around %.0f ±%.0f (%.2f%%), got %d",
					kind, expectedValues, expectedValues*Deviation, Deviation*100, values[kind])
			}
		}
	})
}

// TestRandMapOf tests the method RandMapOf generation of a map from a given type
//
// this test must assert:
//   - no nil values are generated
//   - all values are of the kind reflect.Map
//   - the generated value have the same type of its source
func TestRandMapOf(t *testing.T) {
	t.Run("shouldn't generate nil values", func(t *testing.T) {
		for i := 0; i < Loops; i++ {
			src := reflectrand.Map()
			value := RandMapOf(src)
			typ := reflect.TypeOf(value)

			assertlite.NotNil(t, value, "expecting not nil, but got %v", value)
			assertlite.True(t, typ.Kind() == reflect.Map, "expecting kind map, "+
				"but got %v", typ.Kind())
			assertlite.True(t, src.String() == typ.String(), "expecting generated map '%s' to be "+
				"stringified as '%s'", typ.String(), src.String())
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
