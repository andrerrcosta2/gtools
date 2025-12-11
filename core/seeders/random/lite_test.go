// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

//go:build tt

// TODO: These tests lacks coverage
package random

import (
	"reflect"
	"slices"
	"strings"
	"testing"
	"time"
	"unicode"

	"github.com/andrerrcosta2/gtools/core/domain/data"
	"github.com/andrerrcosta2/gtools/core/format/fmx"
	"github.com/andrerrcosta2/gtools/core/seeders/random/internal/reflectutils"
	"github.com/andrerrcosta2/gtools/core/testlite/assertlite"
	"github.com/andrerrcosta2/gtools/core/testlite/testseed"
	"github.com/andrerrcosta2/gtools/core/util/typeutil/charsets"
)

const (
	Distributions = 100_000
	Deviation     = 0.15 // 15% acceptable deviation
	Loops         = 5000
)

// TestAlphabet tests the Alphabet method generation of random string
//
// this test must assert:
//   - all generated values are of kind reflect.String
//   - all strings contains only alphabet characters
//   - the generated string chars are between the deviation limit of the distribution
func TestAlphabet(t *testing.T) {
	t.Run("random alphabet string", func(t *testing.T) {
		Alphabet(Loops, 1, 50).Each(func(s string) {
			assertlite.IsTypeOf[string](t, s)
			ok, diff := charsets.Validate(s, charsets.Alphabet)
			assertlite.True(t, ok, "expected alphabet to be valid but got char %v", diff)
			assertlite.True(t, len(s) >= 1, "expected alphabet to have at least one char")
			assertlite.True(t, len(s) <= 50, "expected alphabet to have 50 chars or less")
		})
	})

	t.Run("char distribution should be under deviation threshold", func(t *testing.T) {
		counts := make(map[rune]int)
		buckets := len(charsets.Alphabet)
		Alphabet(Distributions, 20, 20).Each(func(s string) {
			for _, char := range s {
				counts[char]++
			}
		})
		expected := float64(Distributions*20) / float64(buckets)
		lowerBound := expected * (1 - Deviation)
		upperBound := expected * (1 + Deviation)

		for _, char := range charsets.Alphabet {
			count := float64(counts[char])
			if count < lowerBound || count > upperBound {
				t.Errorf("char '%s': expected around %.0f ±%.0f (%.2f%%), got %d",
					string(char), expected, expected*Deviation, Deviation*100, counts[char])
			}
		}
	})
}

// TestAlphanumeric tests the Alphanumeric method generation of random string
//
// this test must assert:
//   - all generated values are of kind reflect.String
//   - all strings contains only alphanumeric characters
//   - the generated string chars are between the deviation limit of the distribution
func TestAlphanumeric(t *testing.T) {
	t.Run("random alphanumeric string", func(t *testing.T) {
		Alphanumeric(Loops, 1, 50).Each(func(s string) {
			assertlite.IsTypeOf[string](t, s)
			ok, diff := charsets.Validate(s, charsets.AlphaNumeric)
			assertlite.True(t, ok, "expected alphanumeric to be valid but got char %v", diff)
			assertlite.True(t, len(s) >= 1, "expected alphanumeric to have at least one char")
			assertlite.True(t, len(s) <= 50, "expected alphanumeric to have 50 chars or less")
		})
	})

	t.Run("char distribution should be under deviation threshold", func(t *testing.T) {
		counts := make(map[rune]int)
		buckets := len(charsets.AlphaNumeric)
		Alphanumeric(Distributions, 20, 20).Each(func(s string) {
			for _, char := range s {
				counts[char]++
			}
		})
		expected := float64(Distributions*20) / float64(buckets)
		lowerBound := expected * (1 - Deviation)
		upperBound := expected * (1 + Deviation)

		for _, char := range charsets.AlphaNumeric {
			count := float64(counts[char])
			if count < lowerBound || count > upperBound {
				t.Errorf("char '%s': expected around %.0f ±%.0f (%.2f%%), got %d",
					string(char), expected, expected*Deviation, Deviation*100, counts[char])
			}
		}
	})
}

// TestAny tests the Any method generation of random values of any type
//
// this test must assert:
//   - no nil values are generated
//   - the generated types are between the deviation limit of the distribution
func TestAny(t *testing.T) {
	t.Run("random any value", func(t *testing.T) {
		Any(Loops).EachN(func(i int, a any) {
			assertlite.NotNil(t, a, "nil value at index '%d': %T", i, a)
		})
	})

	t.Run("char distribution should be under deviation threshold", func(t *testing.T) {
		counts := make(map[reflect.Kind]int)
		buckets := len(reflectutils.InterfaceKinds)

		Any(Distributions).Each(func(a any) {
			counts[reflect.TypeOf(a).Kind()]++
		})

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

// TestArray tests the method ArrayOf generation of random values for the given array
func TestArray(t *testing.T) {
	t.Run("random array", func(t *testing.T) {
		const size = 5
		ArrayOf[[10]int](Loops).EachN(func(i int, arr [10]int) {
			assertlite.AreNotNil(t, arr[:], "nil value at index '%d': %T", i, arr)
		})
	})
}

// TestBool tests the Bool method generation of random bool values
//
// this test must assert
//   - all values are reflect.Bool
//   - the generated types are between the deviation limit of the distribution
func TestBool(t *testing.T) {
	t.Run("random bool", func(t *testing.T) {
		count := 0
		Bool(100).Each(func(b bool) {
			assertlite.IsTypeOf[bool](t, b, "Expected bool but got '%T' at index '%d'", b, count)
			count++
		})
		assertlite.True(t, count == 100, "Expected 100 bool, got %d", count)

	})

	t.Run("distribution should be under deviation threshold", func(t *testing.T) {
		counts := make(map[bool]int)
		buckets := 2
		Bool(Distributions).Each(func(b bool) {
			counts[b]++
		})
		expected := float64(Distributions) / float64(buckets)
		lowerBound := expected * (1 - Deviation)
		upperBound := expected * (1 + Deviation)

		for _, kind := range []bool{true, false} {
			count := float64(counts[kind])
			if count < lowerBound || count > upperBound {
				t.Errorf("bool '%v': expected around %.0f ±%.0f (%.2f%%), got %d",
					kind, expected, expected*Deviation, Deviation*100, counts[kind])
			}
		}
	})
}

// TestBytes tests the method Bytes generation of random []byte.
//
//	this test must assert:
//	- no nil values are generated
//	- all generated values are []byte
func TestBytes(t *testing.T) {
	t.Run("Bytes: No size specified", func(t *testing.T) {
		bytes := Bytes(100).Each(func(bytes []byte) {
			assertlite.NotNil(t, bytes, "nil value at index '%d': %T", bytes, bytes)
			assertlite.IsTypeOf[[]byte](t, bytes, "Expected []byte but got '%T' at index '%d'",
				bytes, bytes)
		})
		assertlite.True(t, bytes.Len() == 100, "Expected 100 []byte, got %d", bytes.Len())
	})

	t.Run("Bytes: With size specified", func(t *testing.T) {
		bytes := Bytes(100, 2, 5).Each(func(bytes []byte) {
			assertlite.NotNil(t, bytes, "nil value at index '%d': %T", bytes, bytes)
			assertlite.IsTypeOf[[]byte](t, bytes, "Expected []byte but got '%T' at index '%d'",
				bytes, bytes)
			assertlite.True(t, len(bytes) >= 2 && len(bytes) <= 5, "Expected 2 bytes at least "+
				"and 5 bytes at most, but got %d", len(bytes))
		})
		assertlite.True(t, bytes.Len() == 100, "Expected 100 bytes, got %d", bytes.Len())
	})
}

// TestComparable tests the method Comparable generation of random naturally comparable values.
//
//	this test must assert:
//	- no nil values are generated
//	- all values must be naturally comparable
//	- the generated types are between the deviation limit of the distribution
func TestComparable(t *testing.T) {
	t.Run("random comparable", func(t *testing.T) {
		Comparable(Loops).EachN(func(i int, a any) {
			assertlite.NotNil(t, a, "nil value at index '%d': %T", i, a)
			assertlite.True(t, slices.ContainsFunc(reflectutils.CmpKinds, func(kind reflect.Kind) bool {
				return reflect.TypeOf(a).Kind() == kind
			}))
		})
	})

	t.Run("kind distribution should be under deviation threshold", func(t *testing.T) {
		counts := make(map[reflect.Kind]int)
		buckets := len(reflectutils.CmpKinds)
		Comparable(Distributions).EachN(func(i int, a any) {
			counts[reflect.TypeOf(a).Kind()]++
		})
		expected := float64(Distributions) / float64(buckets)
		lowerBound := expected * (1 - Deviation)
		upperBound := expected * (1 + Deviation)
		for _, kind := range reflectutils.CmpKinds {
			count := float64(counts[kind])
			if count < lowerBound || count > upperBound {
				t.Errorf("bool '%v': expected around %.0f ±%.0f (%.2f%%), got %d",
					kind, expected, expected*Deviation, Deviation*100, counts[kind])
			}
		}
	})
}

// TestComplex64 test the method Complex64 generation of random complex64 values
//
//	this test must assert:
//	- all values are within a given range and a given amount.
func TestComplex64(t *testing.T) {
	t.Run("complex64: No amount specified", func(t *testing.T) {
		c := Complex64(100)
		assertlite.True(t, c.Len() == 100, "Expected 100 complex64, got %d", c.Len())
	})

	t.Run("complex64: between real[-5.5, 7.2] imag[-3.1, 2.8]", func(t *testing.T) {
		realMin, realMax := float32(-5.5), float32(7.2)
		imagMin, imagMax := float32(-3.1), float32(2.8)

		c := Complex64(100, realMin, realMax, imagMin, imagMax).EachN(func(i int, c complex64) {
			r, im := real(c), imag(c)
			assertlite.False(t, r < realMin || r > realMax, "real part out of range at index %d: "+
				"got %v, want between [%v, %v]", i, r, realMin, realMax)
			assertlite.False(t, im < imagMin || im > imagMax, "imag part out of range at index %d: "+
				"got %v, want between [%v, %v]", i, im, imagMin, imagMax)
		})
		assertlite.True(t, c.Len() == 100, "Expected 100 complex64, got %d", c.Len())
	})
}

// TestComplex128 test the method Complex128 generation of random complex128 values
//
//	this test must assert:
//	- all values are within a given range and a given amount.
func TestComplex128(t *testing.T) {
	t.Run("Complex128: No size specified", func(t *testing.T) {
		c := Complex128(100)
		assertlite.True(t, c.Len() == 100, "Expected 100 complex128, got %d", c.Len())
	})

	t.Run("complex128: between real[-5.5, 7.2] imag[-3.1, 2.8]", func(t *testing.T) {
		realMin, realMax := -5.5, 7.2
		imagMin, imagMax := -3.1, 2.8

		c := Complex128(100, realMin, realMax, imagMin, imagMax).EachN(func(i int, c complex128) {
			r, im := real(c), imag(c)
			assertlite.False(t, r < realMin || r > realMax, "real part out of range at index %d: "+
				"got %v, want between [%v, %v]", i, r, realMin, realMax)
			assertlite.False(t, im < imagMin || im > imagMax, "imag part out of range at index %d: "+
				"got %v, want between [%v, %v]", i, im, imagMin, imagMax)
		})
		assertlite.True(t, c.Len() == 100, "Expected 100 complex64, got %d", c.Len())
	})

}

func TestFloat32(t *testing.T) {
	t.Run("Float32: No size specified", func(t *testing.T) {
		floats := Float32(100)
		assertlite.True(t, floats.Len() == 100, "Expected 100 floats, got %d", floats.Len())
	})
}

func TestFloat64(t *testing.T) {
	t.Run("Float64: No size specified", func(t *testing.T) {
		floats := Float64(100)
		assertlite.True(t, floats.Len() == 100, "Expected 100 floats, got %d", floats.Len())
	})
}

func TestInt(t *testing.T) {
	t.Run("Int: No size specified", func(t *testing.T) {
		integers := Int(100)
		assertlite.True(t, integers.Len() == 100, "Expected 100 ints, got %d", integers.Len())
	})

	t.Run("Int: With size specified", func(t *testing.T) {
		integers := Int(100, 10, 200)
		assertlite.True(t, integers.Len() == 100, "Expected 100 ints, got %d", integers.Len())
		assertlite.AreTrue(t, integers.Values(), func(i int) bool {
			return i >= 10 && i <= 200
		},
			"Expected All values to be between 10 and 200, got %v", integers)
	})

	t.Run("Int: With size specified out of range", func(t *testing.T) {
		integers := Int(100, 10, 5)
		assertlite.True(t, integers.Len() == 100, "Expected 100 ints, got %d", integers.Len())
		assertlite.AreTrue(t, integers.Values(), func(i int) bool {
			return i <= 10 && i >= 5
		},
			"Expected All values to be between 10 and 5, got %v", integers)
	})

	t.Run("Int: With only min specified", func(t *testing.T) {
		integers := Int(100, 5000000)
		assertlite.True(t, integers.Len() == 100, "Expected 100 ints, got %d", integers.Len())
		assertlite.AreTrue(t, integers.Values(), func(i int) bool {
			return i >= 5000000
		},
			"Expected All values to be above 5000000, got %v", integers)
	})
}

func TestInt8(t *testing.T) {
	t.Run("Int8: No size specified", func(t *testing.T) {
		integers := Int8(100)
		assertlite.True(t, integers.Len() == 100, "Expected 100 ints, got %d", integers.Len())
	})
	t.Run("Int8: With size specified", func(t *testing.T) {
		integers := Int8(100, 10, 100)
		assertlite.True(t, integers.Len() == 100, "Expected 100 ints, got %d", integers.Len())
		assertlite.AreTrue(t, integers.Values(), func(i int8) bool {
			return i >= 10 && i <= 100
		}, "Expected All values to be between 10 and 200, got %v", integers)
	})
	t.Run("Int8: With size specified out of range", func(t *testing.T) {
		integers := Int8(100, 10, 5)
		assertlite.True(t, integers.Len() == 100, "Expected 100 ints, got %d", integers.Len())
		assertlite.AreTrue(t, integers.Values(), func(i int8) bool {
			return i <= 10 && i >= 5
		},
			"Expected All values to be between 10 and 5, got %v", integers)
	})

	t.Run("Int8: With only min specified", func(t *testing.T) {
		integers := Int8(100, 127)
		assertlite.True(t, integers.Len() == 100, "Expected 100 ints, got %d", integers.Len())
		assertlite.AreTrue(t, integers.Values(), func(i int8) bool {
			return i == 127
		},
			"Expected All values to be 127, got %v", integers)
	})
}

func TestInt16(t *testing.T) {
	t.Run("Int16: No size specified", func(t *testing.T) {
		integers := Int16(100)
		assertlite.True(t, integers.Len() == 100, "Expected 100 ints, got %d", integers.Len())
	})
}

func TestInt32(t *testing.T) {
	t.Run("Int32: No size specified", func(t *testing.T) {
		integers := Int32(100)
		assertlite.True(t, integers.Len() == 100, "Expected 100 ints, got %d", integers.Len())
	})
}

func TestInt64(t *testing.T) {
	t.Run("Int64: No size specified", func(t *testing.T) {
		integers := Int64(100)
		assertlite.True(t, integers.Len() == 100, "Expected 100 ints, got %d", integers.Len())
	})
}

// TestOf tests the random generation of values.
//
// This test must assert that after each call:
//   - All returned slices contain the number of elements passed as argument.
//   - All returned slices contain only the types passed as generic argument.
func TestOf(t *testing.T) {
	t.Run("Of primitives", func(t *testing.T) {
		integers := Of[int](10)
		assertlite.True(t, integers.Len() == 10, "Expected 10 integers, got %d", integers.Len())
		int8s := Of[int8](10)
		assertlite.True(t, int8s.Len() == 10, "Expected 10 int8s, got %d", int8s.Len())
		int16s := Of[int16](10)
		assertlite.True(t, int16s.Len() == 10, "Expected 10 int16s, got %d", int16s.Len())
		int32s := Of[int32](10)
		assertlite.True(t, int32s.Len() == 10, "Expected 10 int32s, got %d", int32s.Len())
		int64s := Of[int64](10)
		assertlite.True(t, int64s.Len() == 10, "Expected 10 int64s, got %d", int64s.Len())
		uints := Of[uint](10)
		assertlite.True(t, uints.Len() == 10, "Expected 10 uints, got %d", uints.Len())
		uint8s := Of[uint8](10)
		assertlite.True(t, uint8s.Len() == 10, "Expected 10 uint8s, got %d", uint8s.Len())
		uint16s := Of[uint16](10)
		assertlite.True(t, uint16s.Len() == 10, "Expected 10 uint16s, got %d", uint16s.Len())
		uint32s := Of[uint32](10)
		assertlite.True(t, uint32s.Len() == 10, "Expected 10 uint32s, got %d", uint32s.Len())
		uint64s := Of[uint64](10)
		assertlite.True(t, uint64s.Len() == 10, "Expected 10 uint64s, got %d", uint64s.Len())
		uintptrs := Of[uintptr](10)
		assertlite.True(t, uintptrs.Len() == 10, "Expected 10 uintptrs, got %d", uintptrs.Len())
		floats := Of[float32](10)
		assertlite.True(t, floats.Len() == 10, "Expected 10 floats, got %d", floats.Len())
		float64s := Of[float64](10)
		assertlite.True(t, float64s.Len() == 10, "Expected 10 float64s, got %d", float64s.Len())
		complex64s := Of[complex64](10)
		assertlite.True(t, complex64s.Len() == 10, "Expected 10 complex64s, got %d", complex64s.Len())
		complex128s := Of[complex128](10)
		assertlite.True(t, complex128s.Len() == 10, "Expected 10 complex128s, got %d", complex128s.Len())
		bools := Of[bool](10)
		assertlite.True(t, bools.Len() == 10, "Expected 10 bools, got %d", bools.Len())
		bytes := Of[byte](10)
		assertlite.True(t, bytes.Len() == 10, "Expected 10 bytes, got %d", bytes.Len())
		runes := Of[rune](10)
		assertlite.True(t, runes.Len() == 10, "Expected 10 runes, got %d", runes.Len())
		ss := Of[string](10)
		assertlite.True(t, ss.Len() == 10, "Expected 10 ss, got %d", ss.Len())
	})

	t.Run("Of complex structures", func(t *testing.T) {
		arrays := Of[[5]int](10)
		assertlite.True(t, arrays.Len() == 10, "Expected 10 arrays, got %d", arrays.Len())
		//fmt.Printf("arrays: %v\n", arrays)
		slices := Of[[]int](10)
		assertlite.True(t, slices.Len() == 10, "Expected 10 slices, got %d", slices.Len())
		//fmt.Printf("slices: %v\n", slices)
		channels := Of[chan int](10)
		assertlite.True(t, channels.Len() == 10, "Expected 10 channels, got %d", channels.Len())
		//fmt.Printf("channels: %v\n", channels)
		maps := Of[map[int]int](10)
		assertlite.True(t, maps.Len() == 10, "Expected 10 maps, got %d", maps.Len())
		//fmt.Printf("maps: %v\n", maps)
	})

	t.Run("Structs", func(t *testing.T) {
		oneMethods := Of[testseed.StructOneData](10)
		assertlite.True(t, oneMethods.Len() == 10, "Expected 10 OneData, got %d", oneMethods.Len())
		twoMethods := Of[testseed.StructTwoData](10)
		assertlite.True(t, twoMethods.Len() == 10, "Expected 10 TwoData, got %d", twoMethods.Len())
		closerReaders := Of[testseed.StructCloserReaderSuccess](10)
		assertlite.True(t, closerReaders.Len() == 10, "Expected 10 CloserReaderSuccess, got %d", closerReaders.Len())
		closerReaderWriters := Of[testseed.StructCloserReaderWriterSuccess](10)
		assertlite.True(t, closerReaderWriters.Len() == 10, "Expected 10 CloserReaderWriterSuccess, got %d", closerReaderWriters.Len())
		namedFloats := Of[testseed.StructFloat](10)
		assertlite.True(t, namedFloats.Len() == 10, "Expected 10 Floats, got %d", namedFloats.Len())
		namedIntegers := Of[testseed.StructInteger](10)
		assertlite.True(t, namedIntegers.Len() == 10, "Expected 10 namedIntegers, got %d", namedIntegers.Len())
		namedStrings := Of[testseed.StructString](10)
		assertlite.True(t, namedStrings.Len() == 10, "Expected 10 Strings, got %d", namedStrings.Len())
		complexUnsafeCastables := Of[testseed.StructComplexUnsafeCastableB](10)
		assertlite.True(t, complexUnsafeCastables.Len() == 10, "Expected 10 ComplexUnsafeCastableB, got %d", complexUnsafeCastables.Len())
	})
}

func TestOf_EdgeCases(t *testing.T) {
	t.Run("Of interface", func(t *testing.T) {
		om := Of[testseed.InterfaceOneMethod](10)
		assertlite.AreTrue(t, om.Values(), func(o testseed.InterfaceOneMethod) bool {
			return o == nil
		},
			"Expected All values to be nil, got %v", om)

		a := Of[any](10)
		assertlite.True(t, a.Len() == 10, "Expected 10 any, got %d", a.Len())
		a.Each(fmx.PrintObj)
		assertlite.AreNotNil(t, a.Values())
	})

	t.Run("Of Struct", func(t *testing.T) {
		ns := Of[testseed.NestedStruct](10)
		assertlite.True(t, ns.Len() == 10, "Expected 10 NestedStruct, got %d", ns.Len())
		assertlite.AreNotNil(t, ns.Values())
		ns.Each(func(val testseed.NestedStruct) {
			assertlite.NoNilFields(t, true, val)
		})

		cc := Of[testseed.ComplexUnsafeCastableBase](10)
		assertlite.True(t, cc.Len() == 10, "Expected 10 ComplexUnsafeCastableBase, got %d", cc.Len())
		assertlite.AreNotNil(t, cc.Values())
		cc.Each(func(val testseed.ComplexUnsafeCastableBase) {
			assertlite.NoNilFields(t, true, val)
		})
	})
}

func TestRune(t *testing.T) {
	t.Run("Rune: No size specified", func(t *testing.T) {
		runes := Rune(100)
		assertlite.True(t, runes.Len() == 100, "Expected 100 runes, got %d", runes.Len())
	})
}

func TestStringMethods(t *testing.T) {
	t.Run("String: No size specified", func(t *testing.T) {
		ss := String(100)
		assertlite.True(t, ss.Len() == 100, "Expected 100 ss, got %d", ss.Len())
	})

	t.Run("Alphanumeric: No size specified", func(t *testing.T) {
		ss := Alphanumeric(100)
		assertlite.True(t, ss.Len() == 100, "Expected 100 ss, got %d", ss.Len())
		assertlite.AreTrue(t, ss.Values(), func(s string) bool {
			for _, r := range s {
				if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
					return false
				}
			}
			return len(s) >= MinStringSizeDefaults && len(s) <= MaxStringSizeDefaults
		}, "Expected All values to be between %d and %d, got %v", MinStringSizeDefaults, MaxStringSizeDefaults, ss)
	})

	t.Run("Alphanumeric: With size specified", func(t *testing.T) {
		ss := Alphanumeric(100, 100, 200)
		assertlite.True(t, ss.Len() == 100, "Expected 100 ss, got %d", ss.Len())
		assertlite.AreTrue(t, ss.Values(), func(s string) bool {
			for _, r := range s {
				if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
					return false
				}
			}
			return len(s) >= 100 && len(s) <= 200
		})
	})

	t.Run("StringOf: No size specified", func(t *testing.T) {
		charset := "aeiou"
		stringsOf := StringOf(100, charset)
		// Assert the number of strings
		assertlite.True(t, stringsOf.Len() == 100, "Expected 100 strings, got %d", stringsOf.Len())

		// Validate that all characters belong to the charset
		valid := func(s string) bool {
			for _, ch := range s {
				if !strings.ContainsRune(charset, ch) {
					return false
				}
			}
			return true
		}
		// Assert all
		assertlite.AreTrue(t, stringsOf.Values(), valid)
	})

	t.Run("StringOf: random amount, custom charset", func(t *testing.T) {
		ss := StringOf(Int(1, 1, 10).At(0), "1234567890", 8, 9).Values()
		assertlite.AreTrue(t, ss, func(s string) bool {
			return len(s) >= 8 && len(s) <= 9
		})
	})
}

func TestStruct(t *testing.T) {
	t.Run("Struct", func(t *testing.T) {
		oneMethods := Struct[testseed.StructOneData](10)
		assertlite.True(t, oneMethods.Len() == 10, "Expected 10 OneData, got %d", oneMethods.Len())

		twoMethods := Struct[testseed.StructTwoData](10)
		assertlite.True(t, twoMethods.Len() == 10, "Expected 10 TwoData, got %d", twoMethods.Len())

		closerReaders := Struct[testseed.StructCloserReaderSuccess](10)
		assertlite.True(t, closerReaders.Len() == 10,
			"Expected 10 CloserReaderSuccess, got %d", closerReaders.Len())

		closerReaderWriters := Struct[testseed.StructCloserReaderWriterSuccess](10)
		assertlite.True(t, closerReaderWriters.Len() == 10,
			"Expected 10 CloserReaderWriterSuccess, got %d", closerReaderWriters.Len())

		// This isn't a struct, is a named type
		assertlite.Panic(t, func() {
			Struct[testseed.StructFloat](10)
		})

		// This isn't a struct, is a named type
		assertlite.Panic(t, func() {
			Struct[testseed.StructInteger](10)
		})

		// This isn't a struct, is a named type
		assertlite.Panic(t, func() {
			Struct[testseed.StructString](10)
		})

		complexUnsafeCastables := Struct[testseed.StructComplexUnsafeCastableB](10)
		assertlite.True(t, complexUnsafeCastables.Len() == 10,
			"Expected 10 ComplexUnsafeCastableB, got %d", complexUnsafeCastables.Len())
	})
}

func TestTimestamp(t *testing.T) {
	t.Run("Timestamp: No interval specified", func(t *testing.T) {
		timestamps := Timestamp(10000)
		assertlite.True(t, timestamps.Len() == 10000, "Expected 100 timestamps, got %d", timestamps.Len())
	})

	t.Run("Timestamp: With interval specified", func(t *testing.T) {
		timestamps := Timestamp(10000, time.Now(), time.Now().Add(time.Hour))
		assertlite.True(t, timestamps.Len() == 10000, "Expected 100 timestamps, got %d", timestamps.Len())
		assertlite.AreTrue(t, timestamps.Values(), func(t time.Time) bool {
			return t.Before(time.Now().Add(time.Hour))
		},
			"Expected All timestamps to be before %v, got %v", time.Now().Add(time.Hour), timestamps)
	})
}

func TestUint(t *testing.T) {
	t.Run("Uint: No size specified", func(t *testing.T) {
		integers := Uint(100)
		assertlite.True(t, integers.Len() == 100, "Expected 100 ints, got %d", integers.Len())
	})
}

func TestUint8(t *testing.T) {
	t.Run("Uint8: No size specified", func(t *testing.T) {
		integers := Uint8(100)
		assertlite.True(t, integers.Len() == 100, "Expected 100 ints, got %d", integers.Len())
	})
}

func TestUint16(t *testing.T) {
	t.Run("Uint16: No size specified", func(t *testing.T) {
		integers := Uint16(100)
		assertlite.True(t, integers.Len() == 100, "Expected 100 ints, got %d", integers.Len())
	})
}

func TestUint32(t *testing.T) {
	t.Run("Uint32: No size specified", func(t *testing.T) {
		integers := Uint32(100)
		assertlite.True(t, integers.Len() == 100, "Expected 100 ints, got %d", integers.Len())
	})
}

func TestUint64(t *testing.T) {
	t.Run("Uint64: No size specified", func(t *testing.T) {
		integers := Uint64(100)
		assertlite.True(t, integers.Len() == 100, "Expected 100 ints, got %d", integers.Len())
	})
}

func TestUuid(t *testing.T) {
	t.Run("Uuid: No size specified", func(t *testing.T) {
		uuids := Uuid(100)
		assertlite.True(t, uuids.Len() == 100, "Expected 100 uuids, got %d", uuids.Len())
	})
}

func TestUniqueByteSlices(t *testing.T) {
	t.Run("UniqueByteSlices: With size specified", func(t *testing.T) {
		// Generate 10000 random byte slices between 2 and 12 bytes with
		// 3 maximum uniqueness retries
		bytes, err := UniqueByteSlices(10000, 3, 2, 12)

		// Ensure the error isn't tagged as IMPOSSIBLE_CONSTRAINT
		assertlite.NotTypeOf[data.Taggable[string]](t, err)

		// If there was no error, assert the length is the same as "q"
		if err == nil {
			assertlite.True(t, bytes.Len() == 10000, "Expected 10000 uuids, got %d", bytes.Len())
		}

		bytesHashes := make(map[string]struct{})
		repeated := make(map[string]int)
		index := 0
		// Assert no repeated byte slices
		assertlite.AreTrue(t, bytes.Values(), func(b []byte) bool {
			if len(b) < 2 || len(b) > 12 {
				t.Errorf("Expected length between 2 and 12, got %d on index %d", len(b), index)
			}
			_, ok := bytesHashes[string(b)]
			bytesHashes[string(b)] = struct{}{}
			if ok {
				repeated[string(b)] = index
			}
			index++
			return !ok
		})
	})
}
