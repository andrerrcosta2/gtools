// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

// TODO: These tests lacks coverage
package random

import (
	"github.com/andrerrcosta2/gtools/core/domain/data"
	"github.com/andrerrcosta2/gtools/core/seeders/random/internal/tests/assertlite"
	"github.com/andrerrcosta2/gtools/core/seeders/random/internal/tests/testseed"
	"strings"
	"testing"
	"time"
	"unicode"
)

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
		strings := Of[string](10)
		assertlite.True(t, strings.Len() == 10, "Expected 10 strings, got %d", strings.Len())
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
		oneMethods := Of[testseed.InterfaceOneMethod](10)
		assertlite.AllTrue(t, oneMethods.Values(), func(o testseed.InterfaceOneMethod) bool {
			return o == nil
		},
			"Expected All values to be nil, got %v", oneMethods)
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
		assertlite.AllTrue(t, integers.Values(), func(i int) bool {
			return i >= 10 && i <= 200
		},
			"Expected All values to be between 10 and 200, got %v", integers)
	})

	t.Run("Int: With size specified out of range", func(t *testing.T) {
		integers := Int(100, 10, 5)
		assertlite.True(t, integers.Len() == 100, "Expected 100 ints, got %d", integers.Len())
		assertlite.AllTrue(t, integers.Values(), func(i int) bool {
			return i <= 10 && i >= 5
		},
			"Expected All values to be between 10 and 5, got %v", integers)
	})

	t.Run("Int: With only min specified", func(t *testing.T) {
		integers := Int(100, 5000000)
		assertlite.True(t, integers.Len() == 100, "Expected 100 ints, got %d", integers.Len())
		assertlite.AllTrue(t, integers.Values(), func(i int) bool {
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
		assertlite.AllTrue(t, integers.Values(), func(i int8) bool {
			return i >= 10 && i <= 100
		}, "Expected All values to be between 10 and 200, got %v", integers)
	})
	t.Run("Int8: With size specified out of range", func(t *testing.T) {
		integers := Int8(100, 10, 5)
		assertlite.True(t, integers.Len() == 100, "Expected 100 ints, got %d", integers.Len())
		assertlite.AllTrue(t, integers.Values(), func(i int8) bool {
			return i <= 10 && i >= 5
		},
			"Expected All values to be between 10 and 5, got %v", integers)
	})

	t.Run("Int8: With only min specified", func(t *testing.T) {
		integers := Int8(100, 127)
		assertlite.True(t, integers.Len() == 100, "Expected 100 ints, got %d", integers.Len())
		assertlite.AllTrue(t, integers.Values(), func(i int8) bool {
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

func TestComplex64(t *testing.T) {
	t.Run("Complex64: No size specified", func(t *testing.T) {
		floats := Complex64(100)
		assertlite.True(t, floats.Len() == 100, "Expected 100 floats, got %d", floats.Len())
	})
}

func TestComplex128(t *testing.T) {
	t.Run("Complex128: No size specified", func(t *testing.T) {
		floats := Complex128(100)
		assertlite.True(t, floats.Len() == 100, "Expected 100 floats, got %d", floats.Len())
	})
}

func TestStringMethods(t *testing.T) {
	t.Run("String: No size specified", func(t *testing.T) {
		strings := String(100)
		assertlite.True(t, strings.Len() == 100, "Expected 100 strings, got %d", strings.Len())
	})

	t.Run("Alphanumeric: No size specified", func(t *testing.T) {
		strings := Alphanumeric(100)
		assertlite.True(t, strings.Len() == 100, "Expected 100 strings, got %d", strings.Len())
		assertlite.AllTrue(t, strings.Values(), func(s string) bool {
			for _, r := range s {
				if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
					return false
				}
			}
			return len(s) >= MinStringSizeDefaults && len(s) <= MaxStringSizeDefaults
		}, "Expected All values to be between %d and %d, got %v", MinStringSizeDefaults, MaxStringSizeDefaults, strings)
	})

	t.Run("Alphanumeric: With size specified", func(t *testing.T) {
		strings := Alphanumeric(100, 100, 200)
		assertlite.True(t, strings.Len() == 100, "Expected 100 strings, got %d", strings.Len())
		assertlite.AllTrue(t, strings.Values(), func(s string) bool {
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
		assertlite.AllTrue(t, stringsOf.Values(), valid)
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
		assertlite.AllTrue(t, timestamps.Values(), func(t time.Time) bool {
			return t.Before(time.Now().Add(time.Hour))
		},
			"Expected All timestamps to be before %v, got %v", time.Now().Add(time.Hour), timestamps)
	})
}

func TestBool(t *testing.T) {
	t.Run("Bool: No size specified", func(t *testing.T) {
		bools := Bool(100)
		assertlite.True(t, bools.Len() == 100, "Expected 100 bools, got %d", bools.Len())
	})
}

func TestBytes(t *testing.T) {
	t.Run("Bytes: No size specified", func(t *testing.T) {
		bytes := Bytes(100)
		assertlite.True(t, bytes.Len() == 100, "Expected 100 bytes, got %d", bytes.Len())
		//fmt.Printf("%v\n", bytes)
	})

	t.Run("Bytes: With size specified", func(t *testing.T) {
		bytes := Bytes(100, 2, 5)
		assertlite.True(t, bytes.Len() == 100, "Expected 100 bytes, got %d", bytes.Len())
		//fmt.Printf("%v\n", bytes)
		assertlite.AllTrue(t, bytes.Values(), func(b []byte) bool {
			return len(b) >= 2 && len(b) <= 5
		},
			"Expected All bytes to have between 2 and 5 length, got %v", bytes)
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

		// If there was no error, assert the length is the same as "q'
		if err == nil {
			assertlite.True(t, bytes.Len() == 10000, "Expected 10000 uuids, got %d", bytes.Len())
		}

		bytesHashes := make(map[string]struct{})
		repeated := make(map[string]int)
		index := 0
		// Assert no repeated byte slices
		assertlite.AllTrue(t, bytes.Values(), func(b []byte) bool {
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

func TestRune(t *testing.T) {
	t.Run("Rune: No size specified", func(t *testing.T) {
		runes := Rune(100)
		assertlite.True(t, runes.Len() == 100, "Expected 100 runes, got %d", runes.Len())
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
