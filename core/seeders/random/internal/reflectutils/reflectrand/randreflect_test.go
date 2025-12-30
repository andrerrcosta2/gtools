// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package reflectrand

import (
	"github.com/andrerrcosta2/gtools/core/seeders/random/internal/prng"
	"github.com/andrerrcosta2/gtools/core/seeders/random/internal/reflectutils"
	"github.com/andrerrcosta2/gtools/core/testlite/assertlite"
	"github.com/andrerrcosta2/gtools/core/testlite/testseed"
	"github.com/andrerrcosta2/gtools/core/util/typeutil/charsets"
	"reflect"
	"slices"
	"testing"
)

const (
	Distributions = 100_000
	Deviation     = 0.15 // 15% acceptable deviation
	Loops         = 5000
)

// TestAnyValue tests the AnyValue function generation of random reflect.Value of random
// types.
//
// This test must assert:
//   - No nil value is generated
//   - No invalid value is generated
func TestAnyValue(t *testing.T) {
	for i := 0; i < Loops; i++ {
		value := AnyValue()
		if reflectutils.IsNullable(value.Type()) {
			assertlite.False(t, value.IsNil())
		}
		assertlite.True(t, value.IsValid())
	}
}

// TestAnyInterfaceValue tests the AnyInterfaceValue function generation of random reflect.Value.
//
// This test must assert:
//   - No nil value is generated
//   - no invalid value is generated
//   - no interface is generated
func TestAnyInterfaceValue(t *testing.T) {
	for i := 0; i < Distributions; i++ {
		value := anyInterfaceValue(reflectutils.TrackerOf(3))

		// Assert that the generated value is not nil
		if reflectutils.IsNullable(value.Type()) {
			assertlite.False(t, value.IsNil())
		}
		assertlite.True(t, value.IsValid())
		assertlite.False(t, value.Kind() == reflect.Interface)
	}
}

// TestArray tests the generation of Array types
//
// this test must assert:
//   - All types have its kind as reflect.Array
func TestArray(t *testing.T) {
	for i := 0; i < Loops; i++ {
		a := Array()
		assertlite.True(t, a.Kind() == reflect.Array)
	}
}

// TestArrayOf tests the generation of array values
//
// this test must assert:
//   - No invalid values
//   - All values must have its kind as reflect.Array
func TestArrayOf(t *testing.T) {
	for i := 0; i < Loops; i++ {
		a := ArrayOf(Array())
		assertlite.True(t, a.IsValid())
		assertlite.True(t, a.Kind() == reflect.Array)
	}
}

// TestChan tests the generation of random channel types
//
// this test must assert:
//   - All types have as its kind a reflect.Chan
func TestChan(t *testing.T) {
	for i := 0; i < Loops; i++ {
		c := Chan()
		assertlite.True(t, c.Kind() == reflect.Chan)
	}
}

// TestChanOf tests the generation of channel values by a given type
//
// this test must assert:
//   - All values have as its kind a reflect.Chan
//   - All values are valid
func TestChanOf(t *testing.T) {
	for i := 0; i < Loops; i++ {
		c := ChanOf(Chan())
		assertlite.True(t, c.IsValid(), "invalid channel generated")
		assertlite.True(t, c.Kind() == reflect.Chan, "expected to generate a channel, but got %v",
			c.Kind())
	}
}

// TestCmpKind tests the CmpKind function generation of random reflect.Kind.
//
// This test must assert:
//   - the generated types are between the deviation limit of the distribution
func TestCmpKind(t *testing.T) {
	counts := make(map[reflect.Kind]int)
	buckets := len(reflectutils.CmpKinds)

	for i := 0; i < Distributions; i++ {
		kind := CmpKind()
		assertlite.True(t, kind != reflect.Invalid)
		counts[kind]++
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
}

// TestCmpType tests the Cmp function generation of random reflect.Type.
//
// This test must assert:
//   - No type is invalid
//   - the generated types are between the deviation limit of the distribution
func TestCmpType(t *testing.T) {
	counts := make(map[reflect.Kind]int)
	buckets := len(reflectutils.CmpKinds)

	for i := 0; i < Distributions; i++ {
		tt := Cmp()
		assertlite.True(t, tt.Kind() != reflect.Invalid)
		counts[tt.Kind()]++
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
}

// TestFunc tests the generation of random func types
//
// this test must assert
//   - All types are reflect.Func
func TestFunc(t *testing.T) {
	for i := 0; i < Loops; i++ {
		fn := Func()
		assertlite.True(t, fn.Kind() == reflect.Func, "Func() should have a Kind of reflect.Func, "+
			"but got %v", fn.Kind())
	}
}

// TestFuncOf tests the generation of a function reflect.Value by a given type
//
// this test must assert
//   - All values are valid
//   - All values are reflect.Func
func TestFuncOf(t *testing.T) {
	for i := 0; i < Loops; i++ {
		fn := FuncOf(Func())
		assertlite.True(t, fn.Kind() == reflect.Func, "Func() should have a Kind of reflect.Func, "+
			"but got %v", fn.Kind())
		assertlite.True(t, fn.IsValid(), "Func() should be valid")
	}
}

// TestInterface tests the method Interface generation of Interface types
func TestInterface(t *testing.T) {
	tt := Interface()
	assertlite.True(t, tt.Kind() == reflect.Interface, "Interface() should have a "+
		"Kind of reflect.Interface, but got %v", tt.Kind())
}

// TestInterfaceOf tests the method InterfaceOf generation of random interface values
//
// This test must assert:
//   - Interfaces with methods returns a zero reflect.Value of the reflect.Interface given type
//   - Interfaces with methods must have an invalid element
//   - Interfaces without methods returns an any value wrapped by an interface{}
//   - Interfaces without methods must have a valid element
func TestInterfaceOf(t *testing.T) {
	// interface with methods should be Zero
	t.Run("interface with methods should be zeroed", func(t *testing.T) {
		typ := reflect.TypeOf((*testseed.InterfaceCloser)(nil)).Elem()

		assertlite.True(t, typ.Kind() == reflect.Interface, "type should be interface, "+
			"but got %v", typ.Kind())
		val := InterfaceOf(typ)
		assertlite.True(t, val.IsValid(), "Interface should be valid")
		assertlite.True(t, val.String() == "<testseed.InterfaceCloser Value>", "Interface should be "+
			"testseed.InterfaceCloser, but is %v", val.String())
		assertlite.True(t, val.IsZero(), "Interface should be zero")

		elem := val.Elem()
		assertlite.True(t, elem.Kind() == reflect.Invalid, "expected invalid kind for zero interface")

	})

	t.Run("interface with no methods should be any value", func(t *testing.T) {
		// interface with no methods should be any non-zero
		typ := reflect.TypeOf((*any)(nil)).Elem()
		assertlite.True(t, typ.Kind() == reflect.Interface, "type should be interface, "+
			"but got %v", typ.Kind())
		val := InterfaceOf(typ)
		assertlite.True(t, val.IsValid(), "Interface should be valid")
		assertlite.True(t, val.String() == "<interface {} Value>", "Interface should be "+
			"testseed.InterfaceCloser, but is %v", val.String())
		assertlite.False(t, val.IsZero(), "Interface should be valid")

		assertlite.NoPanic(t, func() {
			elem := val.Elem()
			assertlite.True(t, elem.IsValid(), "Interface element should be valid")
		}, "expected not to panic calling the element of an any interface")
	})
}

// TestInterfaceType tests the method interfaceType generation of random interface{} types
//
// this test must assert:
//   - the generated types are valid and one of the kinds from reflectutils.InterfaceKinds
//   - the generated types are between the deviation limit of the distribution
func TestInterfaceType(t *testing.T) {
	t.Run("interface types shouldn't be interfaces", func(t *testing.T) {
		for i := 0; i < Loops; i++ {
			typ := interfaceType()
			assertlite.True(t, slices.ContainsFunc(reflectutils.InterfaceKinds, func(kind reflect.Kind) bool {
				return kind == typ.Kind()
			}))
		}
	})

	t.Run("distribution should be under deviation threshold", func(t *testing.T) {
		counts := make(map[reflect.Kind]int)
		buckets := len(reflectutils.InterfaceKinds)

		for i := 0; i < Distributions; i++ {
			kind := Kind()
			counts[kind]++
		}

		expected := float64(Distributions) / float64(buckets)
		lowerBound := expected * (1 - Deviation)
		upperBound := expected * (1 + Deviation)

		for _, kind := range reflectutils.InterfaceKinds {
			count := float64(counts[kind])
			if count < lowerBound || count > upperBound {
				t.Errorf("kind %v: expected around %.0f ±%.0f (%.2f%%), got %d",
					kind, expected, expected*Deviation, Deviation*100, counts[kind])
			}
		}
	})
}

// TestInvalidType tests the InvalidType reflect.Type.
func TestInvalidType(t *testing.T) {
	tt := InvalidType()
	assertlite.True(t, tt == nil)
}

// TestInvalidValue tests the InvalidValue reflect.Value.
func TestInvalidValue(t *testing.T) {
	assertlite.False(t, InvalidValue().IsValid())
}

// TestKind tests the Kind function generation of random reflect.Kind.
//
// This test must assert
//   - Only randomized kinds can be generated (expect reflect.Struct and reflect.Invalid
//   - the generated types are between the deviation limit of the distribution
func TestKind(t *testing.T) {
	counts := make(map[reflect.Kind]int)
	buckets := len(reflectutils.RandomizableKinds)

	for i := 0; i < Distributions; i++ {
		kind := Kind()
		assertlite.True(t, kind != reflect.Invalid, "kind is invalid")
		assertlite.True(t, kind != reflect.Struct, "kind is struct")
		counts[kind]++
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
}

// TestMap tests the Map function generation of Map types
//
// this test must assert
func TestMap(t *testing.T) {
	t.Run("all values should be map", func(t *testing.T) {
		for i := 0; i < Loops; i++ {
			typ := Map()
			assertlite.True(t, typ.Kind() == reflect.Map, "expected kind to be reflect.Map, but is %v",
				typ.Kind())
		}
	})

	t.Run("distributions should be under deviation threshold", func(t *testing.T) {
		keys := make(map[reflect.Kind]int)
		values := make(map[reflect.Kind]int)
		keyBuckets := len(reflectutils.CmpKinds)
		valueBuckets := len(reflectutils.RandomizableKinds)

		for i := 0; i < Distributions; i++ {
			typ := Map()
			kk := typ.Key().Kind()
			vk := typ.Elem().Kind()
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

// TestMapOf tests the method MapOf generation of random map values
//
// this test must assert:
//   - the generated values are all reflect.Map
//   - the generated values are all valid and non-zero
func TestMapOf(t *testing.T) {
	t.Run("all values should be map", func(t *testing.T) {
		for i := 0; i < Loops; i++ {
			v := MapOf(Map())
			assertlite.True(t, v.Kind() == reflect.Map, "expected kind to be reflect.Map, but is %v",
				v.Kind())
			assertlite.True(t, v.IsValid(), "expected kind to be valid")
			assertlite.False(t, v.IsZero(), "expected kind not to be zero")
		}
	})
}

// TestPointer tests the Pointer method generation of random pointer types
//
// this test must assert
//   - All types are reflect.Ptr and have valid element types
func TestPointer(t *testing.T) {
	t.Run("all values should be pointer", func(t *testing.T) {
		typ := Pointer()
		assertlite.True(t, typ.Kind() == reflect.Ptr, "expected kind to be reflect.Ptr, but is %v",
			typ.Kind())
		assertlite.True(t, typ.Elem().Kind() != reflect.Invalid, "expected typ.Elem().Kind() to not "+
			"be reflect.Invalid")
	})

	t.Run("distributions should be under deviation threshold", func(t *testing.T) {
		counts := make(map[reflect.Kind]int)
		buckets := len(reflectutils.RandomizableKinds)

		for i := 0; i < Distributions; i++ {
			typ := Pointer()
			counts[typ.Elem().Kind()]++
		}

		expected := float64(Distributions) / float64(buckets)
		lowerBound := expected * (1 - Deviation)
		upperBound := expected * (1 + Deviation)
		for _, kind := range reflectutils.RandomizableKinds {
			count := float64(counts[kind])
			if count < lowerBound || count > upperBound {
				t.Errorf("kind '%v': expected around %.0f ±%.0f (%.2f%%), got %d",
					kind, expected, expected*Deviation, Deviation*100, counts[kind])
			}
		}
	})
}

// TestPointerOf tests the PointerOf method generation of random pointer values
//
// this test must assert:
//   - All values are valid reflect.Pointer
//   - All values contain a valid reflectutils.RandomizableKinds element
func TestPointerOf(t *testing.T) {
	t.Run("all values should be pointer", func(t *testing.T) {
		v := PointerOf(Pointer())
		assertlite.True(t, v.IsValid(), "expected kind to be valid")
		assertlite.False(t, v.IsZero(), "expected kind not to be zero")
		assertlite.True(t, v.Kind() == reflect.Pointer, "expected kind not to be reflect.Pointer but is "+
			"%v", v.Kind())
		assertlite.True(t, v.Elem().IsValid(), "value.Elem().Kind() is invalid")
		assertlite.True(t, slices.ContainsFunc(reflectutils.RandomizableKinds, func(kind reflect.Kind) bool {
			return kind != v.Elem().Kind()
		}), "expected pointer element to be a randomized kind but it is %v", v.Elem().Kind())
	})
}

// TestSlice tests the Slice method generation of random slice types
//
// this test must assert:
//   - all types are reflect.Slice
//   - the generated types are between the deviation limit of the distribution
func TestSlice(t *testing.T) {
	t.Run("all values should be slice", func(t *testing.T) {
		for i := 0; i < Loops; i++ {
			typ := Slice()
			assertlite.True(t, typ.Kind() == reflect.Slice, "expected kind to be reflect.Slice, "+
				"but is %v", typ.Kind())
		}
	})

	t.Run("distributions should be under deviation threshold", func(t *testing.T) {
		counts := make(map[reflect.Kind]int)
		buckets := len(reflectutils.RandomizableKinds)
		for i := 0; i < Distributions; i++ {
			typ := Slice()
			counts[typ.Elem().Kind()]++
		}
		expected := float64(Distributions) / float64(buckets)
		lowerBound := expected * (1 - Deviation)
		upperBound := expected * (1 + Deviation)

		for _, kind := range reflectutils.RandomizableKinds {
			count := float64(counts[kind])
			if count < lowerBound || count > upperBound {
				t.Errorf("kind '%v': expected around %.0f ±%.0f (%.2f%%), got %d",
					kind, expected, expected*Deviation, Deviation*100, counts[kind])
			}
		}
	})
}

func TestSliceOf(t *testing.T) {
	t.Run("all values should be slice", func(t *testing.T) {
		for i := 0; i < Loops; i++ {
			v := SliceOf(Slice())
			assertlite.True(t, v.IsValid(), "expected kind to be valid")
			assertlite.False(t, v.IsZero(), "expected kind not to be zero")
			assertlite.True(t, v.Kind() == reflect.Slice, "expected kind to be reflect.Slice")
			assertlite.True(t, slices.ContainsFunc(reflectutils.RandomizableKinds, func(kind reflect.Kind) bool {
				return kind != v.Type().Elem().Kind()
			}), "expected slice element to be a randomized kind but it is %v", v.Type().Elem().Kind())
		}
	})
}

// TestString tests the method String generation of string types
//
// this test must assert:
//   - all types are reflect.String
func TestString(t *testing.T) {
	t.Run("all values should be string", func(t *testing.T) {
		for i := 0; i < Loops; i++ {
			typ := String()
			assertlite.True(t, typ.Kind() == reflect.String, "expected kind to be reflect.String, "+
				"but is %v", typ.Kind())
		}
	})
}

// TestString tests the generation of random strings as reflect.Value
//
// this test must assert:
//   - No empty string is generated
//   - All value kinds are reflect.String
//   - The strings have the exact length of the parameter
//   - No strings outside the charset is used
func TestStringOf(t *testing.T) {
	for i := 0; i < Distributions; i++ {
		length := prng.Int(3, 30)
		s := StringOf(length, charsets.AlphaNumeric)
		assertlite.True(t, s.Kind() == reflect.String, "expected a string but got '%v'", s.Kind())
		assertlite.True(t, s.Len() <= 30, "expected string length to be less than 30, but "+
			"got '%d'", s.Len())
		assertlite.True(t, s.Len() >= 3, "expected string length to be higher than 3, but "+
			"got '%d'", s.Len())
		valid, r := charsets.Validate(s.String(), charsets.AlphaNumeric)
		assertlite.True(t, valid, "invalid char found within generated string: '%s'", r)
	}
}

// TestStruct tests the Struct method generation of random structs
//
// this test must assert
//   - No invalid values are generated
//   - All values must be of kind 'reflect.Struct'
//   - No fields - except interface with methods - can be nil
func TestStructOf(t *testing.T) {
	for _, zero := range testseed.ZeroValues() {
		tt := reflect.TypeOf(zero)
		assertlite.True(t, tt.Kind() == reflect.Struct,
			"expected seed to be a struct, but got '%s' for '%T'", tt.Kind(), zero)
		rand := StructOf(tt, 3)
		assertlite.True(t, rand.IsValid(), "invalid random struct")
		assertlite.True(t, rand.Kind() == reflect.Struct,
			"expected random struct, but got '%s'", rand.Kind())
		assertlite.NoNilNonInterfaceFields(t, true, rand.Interface())
	}
}

func TestStructOf_EdgeCase(t *testing.T) {
	tt := reflect.TypeOf(testseed.Account{})
	assertlite.True(t, tt.Kind() == reflect.Struct, "expected seed to be a struct, but "+
		"got '%s' for '%T'", tt.Kind(), testseed.Account{})
	rand := StructOf(tt, 3)
	assertlite.True(t, rand.IsValid(), "invalid random struct")
	assertlite.True(t, rand.Kind() == reflect.Struct, "expected random struct, but got '%s'", rand.Kind())
	assertlite.NoNilNonInterfaceFields(t, true, rand.Interface())
}

// TestType tests the Type function generation of random reflect.Type.
//
// This test must assert:
//   - No type is invalid
//   - the generated types are between the deviation limit of the distribution
func TestType(t *testing.T) {
	t.Run("shouldn't generate any invalid type", func(t *testing.T) {
		for i := 0; i < 500; i++ {
			assertlite.True(t, Type().Kind() != reflect.Invalid)
		}
	})

	t.Run("should distribute under threshold", func(t *testing.T) {
		counts := make(map[reflect.Kind]int)
		buckets := len(reflectutils.RandomizableKinds)

		for i := 0; i < Distributions; i++ {
			kind := Type().Kind()
			counts[kind]++
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

// TestValueOf tests the ValueOf generation of random values by a given type
//
// this test must assert:
//   - both types and values have the same kind
//   - No invalid values are generated
//   - the generated values are between the deviation limit of the distribution
func TestValueOf(t *testing.T) {
	t.Run("shouldn't generate any invalid type", func(t *testing.T) {
		for i := 0; i < Loops; i++ {
			typ := Type()
			v := ValueOf(typ)
			assertlite.True(t, v.IsValid(), "generated invalid value from type: %v", typ.String())
			assertlite.True(t, v.Kind() == typ.Kind(), "mismatched kinds between type '%v' and "+
				"value '%v'", typ.Kind(), v.Kind())
		}
	})

	t.Run("should distribute under threshold", func(t *testing.T) {
		counts := make(map[reflect.Kind]int)
		buckets := len(reflectutils.RandomizableKinds)

		for i := 0; i < Distributions; i++ {
			kind := Type().Kind()
			counts[kind]++
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
