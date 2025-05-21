// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package differ

import (
	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"github.com/andrerrcosta2/gtools/core/format/differs"
	"github.com/andrerrcosta2/gtools/core/format/fmx"
	"github.com/andrerrcosta2/gtools/core/format/sprints"
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static/interf"
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static/pointers"
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static/structs"
	"github.com/andrerrcosta2/gtools/gtests/testingtools"
	"github.com/andrerrcosta2/gtools/gtests/testingtools/config/testlogs"
	"github.com/andrerrcosta2/gtools/gtests/testingtools/themes"
	"github.com/andrerrcosta2/gtools/reflect4/internal/_testdata"
	"github.com/andrerrcosta2/gtools/reflect4/internal/types"
	"math"
	"reflect"
	"testing"
	"unsafe"
)

// TestValues_Primitives
func TestValues_Primitives(t *testing.T) {
	zero := indent.Zero()

	// Test 1: Equal integers
	t.Run("Equals integers", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		tc := TestCase{
			name:   "Equals integers",
			a:      reflect.ValueOf(36),
			b:      reflect.ValueOf(36),
			diff:   "",
			equals: true,
			err:    nil,
		}
		testDifferValues(tt, tc)
	})

	// Test 2: Different integers
	t.Run("Different integers", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		tc := TestCase{
			name: "Different integers",
			a:    reflect.ValueOf(36),
			b:    reflect.ValueOf(84),
			diff: differs.Values(zero, sprints.TypedDigit("int", 36),
				sprints.TypedDigit("int", 84)),
			equals: false,
			err:    nil,
		}
		testDifferValues(tt, tc)
	})

	// Test 3: Equals floats
	t.Run("Equals floats", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		tc := TestCase{
			name:   "Equals floats",
			a:      reflect.ValueOf(3.14),
			b:      reflect.ValueOf(3.14),
			diff:   "",
			equals: true,
			err:    nil,
		}
		testDifferValues(tt, tc)
	})

	// Test 4: Different floats
	t.Run("Different floats", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		tc := TestCase{
			name: "Different floats",
			a:    reflect.ValueOf(3.14),
			b:    reflect.ValueOf(6.28),
			diff: differs.Values(zero, sprints.TypedRoundFloat("float64", 3.14),
				sprints.TypedRoundFloat("float64", 6.28)),
			equals: false,
			err:    nil,
		}
		testDifferValues(tt, tc)
	})

	// Test 5: Equals strings
	t.Run("Equals strings", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		tc := TestCase{
			name:   "Equals strings",
			a:      reflect.ValueOf("hello"),
			b:      reflect.ValueOf("hello"),
			diff:   "",
			equals: true,
			err:    nil,
		}
		testDifferValues(tt, tc)
	})

	// Test 6: Different strings
	t.Run("Different strings", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		tc := TestCase{
			name: "Different strings",
			a:    reflect.ValueOf("hello"),
			b:    reflect.ValueOf("world"),
			diff: differs.Message(differs.Strings(zero, 0), differs.Diff(zero, fmx.SRed("hello"),
				"world")),
			equals: false,
			err:    nil,
		}
		testDifferValues(tt, tc)
	})

	// Test 7: Equals booleans
	t.Run("Equals booleans", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		tc := TestCase{
			name:   "Equals booleans",
			a:      reflect.ValueOf(true),
			b:      reflect.ValueOf(true),
			diff:   "",
			equals: true,
			err:    nil,
		}
		testDifferValues(tt, tc)
	})

	// Test 8: Different booleans
	t.Run("Different booleans", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		tc := TestCase{
			name: "Different booleans",
			a:    reflect.ValueOf(true),
			b:    reflect.ValueOf(false),
			diff: differs.Values(zero, sprints.TypedBool(true),
				sprints.TypedBool(false)),
			equals: false,
			err:    nil,
		}
		testDifferValues(tt, tc)
	})

	// Test 9: Integer vs float
	t.Run("Integer vs float", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		tc := TestCase{
			name:   "Integer vs float",
			a:      reflect.ValueOf(36),
			b:      reflect.ValueOf(36.0),
			diff:   zero.Smarkf(differs.TypesMismatch(zero, "int", "float64")),
			equals: false,
			err:    nil,
		}
		testDifferValues(tt, tc)
	})

	// Test 10: Invalid received value
	t.Run("Invalid received value", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		tc := TestCase{
			name:   "Invalid received value",
			a:      reflect.Value{},
			b:      reflect.ValueOf(10),
			diff:   zero.Smarkf(differs.InvalidReceived(zero, "int")),
			equals: false,
			err:    nil,
		}
		testDifferValues(tt, tc)
	})

	// Test 11: Invalid expected value
	t.Run("Invalid expected value", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		tc := TestCase{
			name:   "Invalid expected value",
			a:      reflect.ValueOf(10),
			b:      reflect.Value{},
			diff:   zero.Smarkf(differs.InvalidExpected(zero, "int")),
			equals: false,
			err:    nil,
		}
		testDifferValues(tt, tc)
	})
}

func TestValues_Primitives_EdgeCases(t *testing.T) {
	zero := indent.Zero()

	// Test 1: Zero Integers
	t.Run("Zero Integers", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		tc := TestCase{
			name:   "Zero Integers",
			a:      reflect.ValueOf(0),
			b:      reflect.ValueOf(0),
			diff:   "",
			equals: true,
			err:    nil,
		}
		testDifferValues(tt, tc)
	})

	// Test 2: Negative vs Positive
	t.Run("negative vs positive integers", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		tc := TestCase{
			name:   "negative vs positive integers",
			a:      reflect.ValueOf(-36),
			b:      reflect.ValueOf(36),
			diff:   differs.Values(zero, sprints.TypedDigit("int", -36), sprints.TypedDigit("int", 36)),
			equals: false,
			err:    nil,
		}
		testDifferValues(tt, tc)
	})

	// Test 3: Different max int
	t.Run("different max int", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		tc := TestCase{
			name:   "Different max int",
			a:      reflect.ValueOf(math.MaxInt),
			b:      reflect.ValueOf(math.MaxInt - 1),
			diff:   differs.Values(zero, sprints.TypedDigit("int", math.MaxInt), sprints.TypedDigit("int", math.MaxInt-1)),
			equals: false,
			err:    nil,
		}
		testDifferValues(tt, tc)
	})

	// Test 4: Zero Floats
	t.Run("Zero Floats", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		tc := TestCase{
			name:   "Zero Floats",
			a:      reflect.ValueOf(0.0),
			b:      reflect.ValueOf(0.0),
			diff:   "",
			equals: true,
			err:    nil,
		}
		testDifferValues(tt, tc)
	})

	// Test 5: Negative vs positive floats
	t.Run("negative vs positive floats", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		tc := TestCase{
			name:   "negative vs positive floats",
			a:      reflect.ValueOf(-3.14),
			b:      reflect.ValueOf(3.14),
			diff:   differs.Values(zero, sprints.TypedRoundFloat("float64", -3.14), sprints.TypedRoundFloat("float64", 3.14)),
			equals: false,
			err:    nil,
		}
		testDifferValues(tt, tc)
	})

	// Test 6: NaN floats
	// According to the IEEE 754 standard, NaN is not equal to anything , including itself. This means:
	//  - math.NaN() == math.NaN() // false
	//  - reflect.DeepEqual(math.NaN(), math.NaN()) // false
	t.Run("NaN floats", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		tc := TestCase{
			name:   "NaN floats",
			a:      reflect.ValueOf(math.NaN()),
			b:      reflect.ValueOf(math.NaN()),
			diff:   differs.Values(zero, sprints.Typed("float64", math.NaN()), sprints.Typed("float64", math.NaN())),
			equals: false,
			err:    nil,
		}
		testDifferValues(tt, tc)
	})

	// Test 7: Infinity floats
	t.Run("Infinity floats", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		tc := TestCase{
			name:   "Infinity floats",
			a:      reflect.ValueOf(math.Inf(1)),
			b:      reflect.ValueOf(math.Inf(-1)),
			diff:   differs.Values(zero, sprints.TypedRoundFloat("float64", math.Inf(1)), sprints.TypedRoundFloat("float64", math.Inf(-1))),
			equals: false,
			err:    nil,
		}
		testDifferValues(tt, tc)
	})

	// Test 8: Different max float
	t.Run("Different max float", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		tc := TestCase{
			name:   "Different max float",
			a:      reflect.ValueOf(math.MaxFloat64),
			b:      reflect.ValueOf(math.MaxFloat64 * -1),
			diff:   differs.Values(zero, sprints.TypedRoundFloat("float64", math.MaxFloat64), sprints.TypedRoundFloat("float64", math.MaxFloat64*-1)),
			equals: false,
			err:    nil,
		}
		testDifferValues(tt, tc)
	})

	// Test 9: Zero Complex
	t.Run("Zero Complex", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		tc := TestCase{
			name:   "Zero Complex",
			a:      reflect.ValueOf(complex(0, 0)),
			b:      reflect.ValueOf(complex(0, 0)),
			diff:   "",
			equals: true,
			err:    nil,
		}
		testDifferValues(tt, tc)
	})

	// Test 10: Negative vs positive complex
	t.Run("negative vs positive complex", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		tc := TestCase{
			name:   "negative vs positive complex",
			a:      reflect.ValueOf(complex(-3.14, 3.14)),
			b:      reflect.ValueOf(complex(3.14, -3.14)),
			diff:   differs.Values(zero, sprints.TypedComplex128(complex(-3.14, 3.14)), sprints.TypedComplex128(complex(3.14, -3.14))),
			equals: false,
			err:    nil,
		}
		testDifferValues(tt, tc)
	})

	// Test 11: Empty strings
	t.Run("Empty strings", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		tc := TestCase{
			name:   "Empty strings",
			a:      reflect.ValueOf(""),
			b:      reflect.ValueOf(""),
			diff:   "",
			equals: true,
			err:    nil,
		}
		testDifferValues(tt, tc)
	})

	// Test 12: Whitespace strings
	t.Run("Whitespace strings", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		tc := TestCase{
			name: "Whitespace strings",
			a:    reflect.ValueOf(" "),
			b:    reflect.ValueOf(""),
			diff: differs.Message(differs.Strings(zero, 0), differs.Diff(zero, fmx.SRed("␣"),
				"")),
			equals: false,
			err:    nil,
		}
		testDifferValues(tt, tc)
	})

	// Test 12: Case-sensitive strings
	t.Run("Case-sensitive strings", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		tc := TestCase{
			name: "Case-sensitive strings",
			a:    reflect.ValueOf("Hello"),
			b:    reflect.ValueOf("hello"),
			diff: differs.Message(differs.Strings(zero, 0), differs.Diff(zero, fmx.SRed("Hello"),
				"hello")),
			equals: false,
			err:    nil,
		}
		testDifferValues(tt, tc)
	})

	// Test 13: Equals Unicode strings
	t.Run("Equals Unicode strings", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		tc := TestCase{
			name:   "Equals Unicode strings",
			a:      reflect.ValueOf("こんにちは"),
			b:      reflect.ValueOf("こんにちは"),
			diff:   "",
			equals: true,
			err:    nil,
		}
		testDifferValues(tt, tc)
	})

	// Test 14: Different Unicode strings
	t.Run("Different Unicode strings", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		tc := TestCase{
			name: "Different Unicode strings",
			a:    reflect.ValueOf("こんには"),
			b:    reflect.ValueOf("こんにちは"),
			diff: differs.Message(differs.Strings(zero, 3),
				differs.Diff(zero, "こんに"+fmx.SRedf("は"), "こんにちは")),
			equals: false,
			err:    nil,
		}
		testDifferValues(tt, tc)
	})

	// Test 15: Invalid received value
	t.Run("Invalid received value", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		tc := TestCase{
			name:   "Invalid received value",
			a:      reflect.Value{},
			b:      reflect.ValueOf(36),
			diff:   zero.Smarkf(differs.InvalidReceived(zero, "int")),
			equals: false,
			err:    nil,
		}
		testDifferValues(tt, tc)
	})
}

// TestValues_Arrays tests the Values function with simple arrays
func TestValues_Arrays(t *testing.T) {
	zero := indent.Zero()

	// Test 1: Equal integer arrays
	t.Run("equal integer arrays", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:   "Equals integer arrays",
			a:      reflect.ValueOf([...]int{1, 2, 3}),
			b:      reflect.ValueOf([...]int{1, 2, 3}),
			diff:   "",
			equals: true,
			err:    nil,
		})
	})

	// Test 2: Equal Float arrays
	t.Run("equal float arrays", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:   "Equals float arrays",
			a:      reflect.ValueOf([...]float64{1.1, 2.2, 3.3}),
			b:      reflect.ValueOf([...]float64{1.1, 2.2, 3.3}),
			diff:   "",
			equals: true,
			err:    nil,
		})
	})

	// Test 3: Equal string arrays
	t.Run("equal string arrays", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:   "Equals string arrays",
			a:      reflect.ValueOf([...]string{"a", "b", "c"}),
			b:      reflect.ValueOf([...]string{"a", "b", "c"}),
			diff:   "",
			equals: true,
			err:    nil,
		})
	})

	// Test 4: Different integer arrays
	t.Run("different integer arrays", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name: "Different integer arrays",
			a:    reflect.ValueOf([...]int{1, 2, 3}),
			b:    reflect.ValueOf([...]int{1, 2, 4}),
			diff: differs.Message(differs.Chain(zero, differs.ArrayElem(zero, 2), differs.Values(zero,
				sprints.TypedDigit("int", 3), sprints.TypedDigit("int", 4))), ""),
			equals: false,
			err:    nil,
		})
	})

	// Test 5: Different float arrays
	t.Run("different float arrays", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name: "Different float arrays",
			a:    reflect.ValueOf([...]float64{1.1, 2.2, 3.3}),
			b:    reflect.ValueOf([...]float64{1.1, 2.2, 3.4}),
			diff: differs.Message(differs.Chain(zero, differs.ArrayElem(zero, 2), differs.Values(zero,
				sprints.Typed("float64", "3.3"), sprints.Typed("float64", "3.4"))), ""),
			equals: false,
			err:    nil,
		})
	})

	// Test 6: Different string arrays
	t.Run("different string arrays", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name: "Different string arrays",
			a:    reflect.ValueOf([...]string{"hello", "world"}),
			b:    reflect.ValueOf([...]string{"hello", "there"}),
			diff: differs.Message(differs.Chain(zero, differs.ArrayElem(zero, 1), differs.Strings(zero, 0)),
				differs.Diff(zero.Inc(), fmx.SRedf("world"), "there")),
			equals: false,
			err:    nil,
		})
	})

	// Test 7: Arrays of different lengths
	t.Run("Arrays of different lengths", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:   "Arrays of different lengths",
			a:      reflect.ValueOf([...]int{1, 2, 3}),
			b:      reflect.ValueOf([...]int{1, 2}),
			diff:   differs.Message(differs.ArrayLenMismatch(zero, 3, 2), ""),
			equals: false,
			err:    nil,
		})
	})

	// Test 8: Empty integer arrays
	t.Run("empty integer arrays", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:   "Empty integer arrays",
			a:      reflect.ValueOf([...]int{}),
			b:      reflect.ValueOf([...]int{}),
			diff:   "",
			equals: true,
			err:    nil,
		})
	})

	// Test 9: empty string arrays
	t.Run("empty string arrays", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:   "Empty string arrays",
			a:      reflect.ValueOf([...]string{}),
			b:      reflect.ValueOf([...]string{}),
			diff:   "",
			equals: true,
			err:    nil,
		})
	})

	// Test 10: Equals nested integer arrays
	t.Run("Equals nested integer arrays", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:   "Equals nested integer arrays",
			a:      reflect.ValueOf([...][]int{{1, 2}, {3, 4}}),
			b:      reflect.ValueOf([...][]int{{1, 2}, {3, 4}}),
			diff:   "",
			equals: true,
			err:    nil,
		})
	})

	// Test 11: Different nested integer arrays
	t.Run("Different nested integer arrays", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name: "Different nested integer arrays",
			a:    reflect.ValueOf([...][]int{{1, 2}, {3, 4}}),
			b:    reflect.ValueOf([...][]int{{1, 2}, {3, 5}}),
			diff: differs.Message(differs.Chain(zero, differs.ArrayElem(zero, 1),
				differs.SliceValues(zero, 1), differs.Values(zero, sprints.Typed("int", "4"),
					sprints.Typed("int", "5")), ""), ""),
			equals: false,
			err:    nil,
		})
	})
}

func TestValues_Arrays_EdgeCases(t *testing.T) {
	zero := indent.Zero()

	// Test 1: Array with nil values
	t.Run("Array with nil values", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:   "Array with nil values",
			a:      reflect.ValueOf([...]interface{}{1, nil}),
			b:      reflect.ValueOf([...]interface{}{1, nil}),
			diff:   "",
			equals: true,
			err:    nil,
		})
	})

	// Test 2: Array with mixed types
	t.Run("Array with mixed types", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name: "Array with mixed types",
			a:    reflect.ValueOf([...]interface{}{1, "hello"}),
			b:    reflect.ValueOf([...]interface{}{1, "world"}),
			diff: differs.Message(differs.Chain(zero, differs.ArrayElem(zero, 1), differs.InterfaceImpl(zero),
				differs.Strings(zero, 0)), differs.Diff(indent.Tab(2), fmx.SRedf("hello"), "world")),
			equals: false,
			err:    nil,
		})
	})

	// Test 3: Nested array with nil slice
	t.Run("Nested array with nil slice", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:   "Nested array with nil slice",
			a:      reflect.ValueOf([...][]int{{1, 2}, nil}),
			b:      reflect.ValueOf([...][]int{{1, 2}, nil}),
			diff:   "",
			equals: true,
			err:    nil,
		})
	})

	// Test 4: Nested array with empty sub-array
	t.Run("Nested array with empty sub-array", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:   "Nested array with empty sub-array",
			a:      reflect.ValueOf([...][]int{{1, 2}, {}}),
			b:      reflect.ValueOf([...][]int{{1, 2}, {}}),
			diff:   "",
			equals: true,
			err:    nil,
		})
	})

	// Test 5: Nested array with nil sub-array
	t.Run("Nested array with nil sub-array", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:   "Nested array with nil sub-array",
			a:      reflect.ValueOf([...][]int{{1, 2}, nil}),
			b:      reflect.ValueOf([...][]int{{1, 2}, nil}),
			diff:   "",
			equals: true,
			err:    nil,
		})
	})

	// Test 6: Large integer arrays
	t.Run("Large integer arrays", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name: "Large integer arrays",
			a: reflect.ValueOf(func() [1000]int {
				var arr [1000]int
				for i := 0; i < 1000; i++ {
					arr[i] = i
				}
				return arr
			}()),
			b: reflect.ValueOf(func() [1000]int {
				var arr [1000]int
				for i := 0; i < 1000; i++ {
					arr[i] = i
				}
				return arr
			}()),
			diff:   "",
			equals: true,
			err:    nil,
		})
	})

	// Test 7: Different type arrays
	t.Run("Different type arrays", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:   "Different type arrays",
			a:      reflect.ValueOf([...]int{1, 2, 3}),
			b:      reflect.ValueOf([...]float64{1.0, 2.0, 3.0}),
			diff:   differs.Message(zero.Smarkf(differs.ArrayTypesMismatch(zero, "[3]int", "[3]float64")), ""),
			equals: false,
			err:    nil,
		})
	})

	// Test 8: Arrays of any with different values
	t.Run("arrays of any with different values", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name: "Different type arrays",
			a:    reflect.ValueOf([...]any{1, 2, 3}),
			b:    reflect.ValueOf([...]any{1.0, 2.0, 3.0}),
			diff: differs.Message(differs.Chain(zero, differs.ArrayElem(zero, 0),
				differs.InterfaceImpl(zero), differs.TypesMismatch(zero, "int", "float64")), ""),
			equals: false,
			err:    nil,
		})
	})

	// Test 9: Arrays of custom structs
	t.Run("arrays of custom structs", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name: "Array of custom structs",
			a:    reflect.ValueOf([...]_testdata.Person2{{"Alice", 30}, {"Bob", 25}}),
			b:    reflect.ValueOf([...]_testdata.Person2{{"Alice", 30}, {"Charlie", 25}}),
			diff: differs.Message(differs.Chain(zero, differs.ArrayElem(zero, 1),
				differs.StructFields(zero, "Name"), differs.Strings(zero, 0)),
				differs.Diff(indent.Tab(2), fmx.SRedf("Bob"), "Charlie")),
			equals: false,
			err:    nil,
		})
	})

	// Test 10: Arrays of pointers
	t.Run("arrays of pointers", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name: "Array of pointers",
			a:    reflect.ValueOf([...]int{1, 2, 3}),
			b:    reflect.ValueOf([...]int{1, 2, 4}),
			diff: differs.Message(differs.Chain(zero, differs.ArrayElem(zero, 2), differs.Values(zero,
				sprints.TypedDigit("int", 3), sprints.TypedDigit("int", 4))), ""),
			equals: false,
			err:    nil,
		})
	})

	// Test 11: Arrays of uninitialized elements
	t.Run("arrays of uninitialized elements", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name: "Array with uninitialized elements",
			a:    reflect.ValueOf([...]string{"hello", "", "world"}),
			b:    reflect.ValueOf([...]string{"hello", "there", "world"}),
			diff: differs.Message(differs.Chain(zero, differs.ArrayElem(zero, 1), differs.Strings(zero, 0)),
				differs.Diff(indent.Tab(1), fmx.SRedf(""), "there")),
			equals: false,
			err:    nil,
		})
	})

	// Test 12: Arrays of non-comparable elements
	t.Run("arrays of non-comparable elements", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name: "Array with non-comparable elements",
			a:    reflect.ValueOf([...][]int{{1, 2}, {3, 4}}),
			b:    reflect.ValueOf([...][]int{{1, 2}, {3, 5}}),
			diff: differs.Message(differs.Chain(zero, differs.ArrayElem(zero, 1),
				differs.SliceValues(zero, 1), differs.Values(zero, sprints.Typed("int", "4"),
					sprints.Typed("int", "5")), ""), ""),
			equals: false,
			err:    nil,
		})
	})
}

func TestValues_Channels(t *testing.T) {
	zero := indent.Zero()
	t.Run("Unequal unbuffered channels", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:   "Equal unbuffered channels",
			a:      reflect.ValueOf(make(chan int)),
			b:      reflect.ValueOf(make(chan int)),
			diff:   "",
			equals: true,
			err:    nil,
		})
	})
	t.Run("Equal buffered channels", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:   "Equal buffered channels",
			a:      reflect.ValueOf(make(chan int, 5)),
			b:      reflect.ValueOf(make(chan int, 5)),
			diff:   "",
			equals: true,
			err:    nil,
		})
	})

	// Different channels
	t.Run("Different channel types", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:   "Different channel types",
			a:      reflect.ValueOf(make(chan int)),
			b:      reflect.ValueOf(make(chan string)),
			diff:   differs.Message(zero.Smarkf(differs.ChanTypesMismatch(zero, "chan int", "chan string")), ""),
			equals: false,
			err:    nil,
		})
	})
	t.Run("Different buffer sizes", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:   "Different buffer sizes",
			a:      reflect.ValueOf(make(chan int, 5)),
			b:      reflect.ValueOf(make(chan int, 10)),
			diff:   zero.Smarkf(differs.ChanBufSizeMismatch(zero, 5, 10)),
			equals: false,
			err:    nil,
		})
	})

	// Edge cases
	t.Run("Nil channel vs valid channel", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:   "Nil channel vs valid channel",
			a:      reflect.ValueOf((chan int)(nil)),
			b:      reflect.ValueOf(make(chan int)),
			diff:   zero.Smarkf(differs.NilReceived(zero, "chan", "chan int")),
			equals: false,
			err:    nil,
		})
	})
}

func TestValues_Funcs(t *testing.T) {
	zero := indent.Zero()
	t.Run("Both functions are nil", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:   "Both functions are nil",
			a:      reflect.ValueOf((func())(nil)),
			b:      reflect.ValueOf((func())(nil)),
			diff:   "",
			equals: true,
			err:    nil,
		})
	})
	t.Run("Same function reference", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:   "Same function reference",
			a:      reflect.ValueOf(func() {}),
			b:      reflect.ValueOf(func() {}),
			diff:   "",
			equals: true,
			err:    nil,
		})
	})

	// Different functions
	t.Run("Different function types", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:   "Different function types",
			a:      reflect.ValueOf(func(int) {}),
			b:      reflect.ValueOf(func(string) {}),
			diff:   zero.Smark(differs.FuncTypesMismatch(zero, "func(int)", "func(string)")),
			equals: false,
			err:    nil,
		})
	})
	t.Run("One function is nil", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:   "One function is nil, the other is not",
			a:      reflect.ValueOf((func())(nil)),
			b:      reflect.ValueOf(func() {}),
			diff:   zero.Smark(differs.NilReceived(zero, "func", "func()")),
			equals: false,
			err:    nil,
		})
	})

	// Edge cases
	t.Run("Two similar nil functions", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:   "Two similar nil functions",
			a:      reflect.ValueOf((func())(nil)),
			b:      reflect.ValueOf((func())(nil)),
			diff:   "",
			equals: true,
			err:    nil,
		})
	})
	t.Run("Two different nil functions", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:   "Two different nil functions",
			a:      reflect.ValueOf((func())(nil)),
			b:      reflect.ValueOf((func() int)(nil)),
			diff:   zero.Smarkf(differs.BothNils(zero, "func", "func()", "func() int")),
			equals: false,
			err:    nil,
		})
	})
	t.Run("Functions with different signatures", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:   "Functions with different signatures",
			a:      reflect.ValueOf(func(int, string) {}),
			b:      reflect.ValueOf(func(string, int) {}),
			diff:   zero.Smarkf(differs.FuncTypesMismatch(zero, "func(int, string)", "func(string, int)")),
			equals: false,
			err:    nil,
		})
	})
}

func TestValues_Interfaces(t *testing.T) {
	zero := indent.Zero()
	ftc36 := reflect.ValueOf(interf.FloatTypeCastable(36))
	btctrue := reflect.ValueOf(interf.BoolTypeCastable(true))

	t.Run("Both interfaces are nil", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:   "Both interfaces are nil",
			a:      reflect.ValueOf((interface{})(nil)),
			b:      reflect.ValueOf((interface{})(nil)),
			diff:   differs.BothInvalid(zero),
			equals: true,
			err:    nil,
		})
	})
	t.Run("Same type and value", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:   "Same type and value",
			a:      reflect.ValueOf(interface{}(36)),
			b:      reflect.ValueOf(interface{}(36)),
			diff:   differs.Empty(zero),
			equals: true,
			err:    nil,
		})
	})

	// Different interfaces
	t.Run("Different types", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name: "Different types",
			a:    ftc36,
			b:    btctrue,
			diff: differs.Message(differs.Chain(zero, differs.Pointers(zero),
				differs.StructTypesMismatch(zero, types.Name(ftc36.Elem().Type()), types.Name(btctrue.Elem().Type()))), ""),
			equals: false,
			err:    nil,
		})
	})
	t.Run("Same type, different values", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name: "Same type, different values",
			a:    reflect.ValueOf(interf.FloatTypeCastable(36)),
			b:    reflect.ValueOf(interf.FloatTypeCastable(38)),
			diff: differs.Message(differs.Chain(zero, differs.Pointers(zero),
				differs.StructFields(zero, "ValueC"), differs.Values(zero, sprints.Typed("float64", 36),
					sprints.Typed("float64", 38))), ""),
			equals: false,
			err:    nil,
		})
	})
	t.Run("One interface is nil", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:   "One interface is nil, the other is not",
			a:      reflect.ValueOf((interface{})(nil)),
			b:      reflect.ValueOf(interface{}(36)),
			diff:   zero.Smarkf(differs.InvalidReceived(zero, "int")),
			equals: false,
			err:    nil,
		})
	})

	// Edge cases
	t.Run("Nil vs nil pointer", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:   "Nil vs nil pointer",
			a:      reflect.ValueOf((interface{})(nil)),
			b:      reflect.ValueOf((*int)(nil)),
			diff:   zero.Smark(differs.InvalidReceived(zero, "*int")),
			equals: false,
			err:    nil,
		})
	})
	t.Run("Interfaces holding structs", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name: "Interfaces holding structs",
			a:    reflect.ValueOf(interface{}(struct{ Name string }{"Alice"})),
			b:    reflect.ValueOf(interface{}(struct{ Name string }{"Bob"})),
			diff: differs.Message(differs.Chain(zero, differs.StructFields(zero, "Name"),
				differs.Strings(zero, 0)), differs.Diff(zero.Inc(), fmx.SRed("Alice"), "Bob")),
			equals: false,
			err:    nil,
		})
	})
	t.Run("Interfaces holding slices", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name: "Interfaces holding slices",
			a:    reflect.ValueOf(interface{}([]int{1, 2, 3})),
			b:    reflect.ValueOf(interface{}([]int{1, 2, 4})),
			diff: differs.Message(differs.Chain(zero, differs.SliceValues(zero, 2),
				differs.Values(zero, sprints.Typed("int", 3), sprints.Typed("int", 4))), ""),
			equals: false,
			err:    nil,
		})
	})
	t.Run("Interfaces holding channels"+
		"", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:   "Interfaces holding channels",
			a:      reflect.ValueOf(interface{}(make(chan int))),
			b:      reflect.ValueOf(interface{}(make(chan string))),
			diff:   zero.Smark(differs.ChanTypesMismatch(zero, "chan int", "chan string")),
			equals: false,
			err:    nil,
		})
	})
}

func TestValues_Maps(t *testing.T) {
	zero := indent.Zero()
	t.Run("Equal maps", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:   "Equal maps",
			a:      reflect.ValueOf(map[string]int{"a": 1, "b": 2}),
			b:      reflect.ValueOf(map[string]int{"a": 1, "b": 2}),
			diff:   differs.Empty(zero),
			equals: true,
			err:    nil,
		})
	})
	t.Run("Both maps are nil", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:   "Both maps are nil",
			a:      reflect.ValueOf((map[string]int)(nil)),
			b:      reflect.ValueOf((map[string]int)(nil)),
			diff:   differs.Empty(zero),
			equals: true,
			err:    nil,
		})
	})

	// Different maps
	t.Run("Different keys", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name: "Different keys",
			a:    reflect.ValueOf(map[string]int{"a": 1, "b": 2}),
			b:    reflect.ValueOf(map[string]int{"a": 1, "c": 2}),
			diff: differs.Message(differs.MapKeys(zero), differs.MapKeysDiff(zero, []string{"c"}, []string{"b"},
				[]string{"b"}, []string{"c"})),
			equals: false,
			err:    nil,
		})
	})
	t.Run("Different values", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name: "Different values",
			a:    reflect.ValueOf(map[string]int{"a": 1, "b": 2}),
			b:    reflect.ValueOf(map[string]int{"a": 1, "b": 3}),
			diff: differs.Chain(zero, differs.MapValue(zero, "b"), differs.Values(zero,
				sprints.Typed("int", 2), sprints.Typed("int", 3))),
			equals: false,
			err:    nil,
		})
	})
	t.Run("Different sizes", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name: "Different sizes",
			a:    reflect.ValueOf(map[string]int{"a": 1, "b": 2}),
			b:    reflect.ValueOf(map[string]int{"a": 1}),
			diff: differs.Message(differs.MapKeys(zero),
				differs.MapKeysDiff(zero, []string{}, []string{"b"}, []string{"b"}, []string{})),
			equals: false,
			err:    nil,
		})
	})
	t.Run("One map is nil", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:   "One map is nil, the other is not",
			a:      reflect.ValueOf((map[string]int)(nil)),
			b:      reflect.ValueOf(map[string]int{"a": 1}),
			diff:   zero.Smarkf(differs.NilReceived(zero, "map", "map[string]int")),
			equals: false,
			err:    nil,
		})
	})

	// Edge cases
	t.Run("Maps with nested slices", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name: "Maps with nested slices",
			a:    reflect.ValueOf(map[string][]int{"a": {1, 2}, "b": {3, 4}}),
			b:    reflect.ValueOf(map[string][]int{"a": {1, 2}, "b": {3, 5}}),
			diff: differs.Message(differs.Chain(zero, differs.MapValue(zero, "b"),
				differs.SliceValues(zero, 1), differs.Values(zero, sprints.Typed("int", 4),
					sprints.Typed("int", 5))), ""),
			equals: false,
			err:    nil,
		})
	})
	t.Run("Maps with non-comparable values", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name: "Maps with non-comparable values",
			a:    reflect.ValueOf(map[string][]int{"a": {1, 2}}),
			b:    reflect.ValueOf(map[string][]int{"a": {1, 3}}),
			diff: differs.Message(differs.Chain(zero, differs.MapValue(zero, "a"),
				differs.SliceValues(zero, 1), differs.Values(zero, sprints.Typed("int", 2),
					sprints.Typed("int", 3))), ""),
			equals: false,
			err:    nil,
		})
	})
	t.Run("Maps with struct values", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name: "Maps with struct values",
			a:    reflect.ValueOf(map[string]struct{ Name string }{"a": {"Alice"}, "b": {"Bob"}}),
			b:    reflect.ValueOf(map[string]struct{ Name string }{"a": {"Alice"}, "b": {"Charlie"}}),
			diff: differs.Message(differs.Chain(zero, differs.MapValue(zero, "b"),
				differs.StructFields(zero, "Name"), differs.Strings(zero, 0)),
				differs.Diff(zero.Plus(2), fmx.SRed("Bob"), "Charlie")),
			equals: false,
			err:    nil,
		})
	})
}

func TestValues_Pointers(t *testing.T) {
	zero := indent.Zero()

	t.Run("", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:   "Both pointers are nil",
			a:      reflect.ValueOf(pointers.Nil[int]()),
			b:      reflect.ValueOf(pointers.Nil[int]()),
			diff:   differs.Empty(zero),
			equals: true,
			err:    nil,
		})
	})
	t.Run("", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:   "Pointers point to the same value",
			a:      reflect.ValueOf(pointers.ZeroIntInst),
			b:      reflect.ValueOf(pointers.ZeroIntInst),
			diff:   differs.Empty(zero),
			equals: true,
			err:    nil,
		})
	})

	// Different pointers
	t.Run("", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name: "Pointers point to different values",
			a:    reflect.ValueOf(pointers.Int(36)),
			b:    reflect.ValueOf(pointers.Int(38)),
			diff: differs.Append(zero, differs.Pointers(zero),
				differs.Values(zero, sprints.Typed("int", 36), sprints.Typed("int", 38))),
			equals: false,
			err:    nil,
		})
	})
	t.Run("", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:   "One pointer is nil, the other is not",
			a:      reflect.ValueOf(pointers.Nil[int]()),
			b:      reflect.ValueOf(pointers.Int(36)),
			diff:   zero.Smarkf(differs.NilReceived(zero, "ptr", "*int")),
			equals: false,
			err:    nil,
		})
	})

	// Edge cases
	t.Run("", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:   "Pointers point to different memory addresses but equal values",
			a:      reflect.ValueOf(pointers.Int(36)),
			b:      reflect.ValueOf(pointers.Int(36)),
			diff:   "",
			equals: true,
			err:    nil,
		})
	})
	t.Run("", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:   "Equal Cyclic pointers",
			a:      reflect.ValueOf(structs.CyclicSimpleNode(36, 38)),
			b:      reflect.ValueOf(structs.CyclicSimpleNode(36, 38)),
			diff:   "",
			equals: true,
			err:    nil,
		})
	})
	t.Run("", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name: "Cyclic pointers with different values",
			a:    reflect.ValueOf(structs.CyclicSimpleNode(36, 38)),
			b:    reflect.ValueOf(structs.CyclicSimpleNode(38, 36)),
			diff: differs.Chain(zero, differs.Pointers(zero), differs.StructFields(zero, "Next"),
				differs.Pointers(zero), differs.StructFields(zero, "Value"), differs.InterfaceImpl(zero),
				differs.Values(zero, sprints.Typed("int", 38), sprints.Typed("int", 36))),
			equals: false,
			err:    nil,
		})
	})
	t.Run("", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name: "Cyclic linked list with different values",
			a:    reflect.ValueOf(structs.CyclicLinkedList("innerA", "innerB", "innerC", "innerD")),
			b:    reflect.ValueOf(structs.CyclicLinkedList("innerA", "innerB", "innerC", "innerE")),
			diff: differs.Message(differs.Chain(zero, differs.Pointers(zero), differs.StructFields(zero, "Head"),
				differs.Pointers(zero), differs.StructFields(zero, "Left"),
				differs.Pointers(zero), differs.StructFields(zero, "Value"),
				differs.InterfaceImpl(zero), differs.Strings(zero, 5)),
				differs.Diff(zero.Plus(7), "inner"+fmx.SRed("D"), "innerE")),
			equals: false,
			err:    nil,
		})
	})
}

func TestValues_Slices(t *testing.T) {
	zero := indent.Zero()

	t.Run("Equal slices", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:   "Equal slices",
			a:      reflect.ValueOf([]int{1, 2, 3}),
			b:      reflect.ValueOf([]int{1, 2, 3}),
			diff:   differs.Empty(zero),
			equals: true,
			err:    nil,
		})
	})
	t.Run("Both slices are nil", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:   "Both slices are nil",
			a:      reflect.ValueOf(([]int)(nil)),
			b:      reflect.ValueOf(([]int)(nil)),
			diff:   differs.Empty(zero),
			equals: true,
			err:    nil,
		})
	})

	// Different slices
	t.Run("Different lengths", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:   "Different lengths",
			a:      reflect.ValueOf([]int{1, 2, 3}),
			b:      reflect.ValueOf([]int{1, 2}),
			diff:   zero.Smarkf(differs.SliceLenMismatch(zero, 3, 2)),
			equals: false,
			err:    nil,
		})
	})
	t.Run("Different int types as any", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name: "Different int types as any",
			a:    reflect.ValueOf([]any{1, 2, 3, 4, 5, 6}),
			b:    reflect.ValueOf([]any{int8(1), int8(2), int8(3), int8(4), int8(5), int8(6)}),
			diff: differs.Chain(zero, differs.SliceValues(zero, 1), differs.InterfaceImpl(zero),
				differs.TypesMismatch(zero, "int", "int8")),
			equals: false,
			err:    nil,
		})
	})
	t.Run("Slices with nil elements", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name: "slices with nil elements",
			a:    reflect.ValueOf([]*int{pointers.Int(1), pointers.Int(2), nil, pointers.Int(4), pointers.Int(5), pointers.Int(6)}),
			b:    reflect.ValueOf([]*int{pointers.Int(1), pointers.Int(2), pointers.Int(3), pointers.Int(4), pointers.Int(5), pointers.Int(6)}),
			diff: differs.Chain(zero, differs.SliceValues(zero, 2), differs.NilReceived(zero, "ptr", "*int")),
		})
	})
	t.Run("Same length, different elements", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name: "Same length, different elements",
			a:    reflect.ValueOf([]int{1, 2, 3}),
			b:    reflect.ValueOf([]int{1, 2, 4}),
			diff: differs.Append(zero, differs.SliceValues(zero, 2), differs.Values(zero,
				sprints.TypedDigit("int", 3), sprints.TypedDigit("int", 4))),
			equals: false,
			err:    nil,
		})
	})
	t.Run("One slice is nil", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:   "One slice is nil, the other is not",
			a:      reflect.ValueOf(([]int)(nil)),
			b:      reflect.ValueOf([]int{1, 2, 3}),
			diff:   zero.Smarkf(differs.NilReceived(zero, "slice", "[]int")),
			equals: false,
			err:    nil,
		})
	})

	// Edge cases
	t.Run("Empty slice vs nil slice", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:   "Empty slice vs nil slice",
			a:      reflect.ValueOf([]int{}),
			b:      reflect.ValueOf(([]int)(nil)),
			diff:   zero.Smarkf(differs.NilExpected(zero, "slice", "[]int")),
			equals: false,
			err:    nil,
		})
	})
	t.Run("Slices with nested slices", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name: "Slices with nested slices",
			a:    reflect.ValueOf([][]int{{1, 2}, {3, 4}}),
			b:    reflect.ValueOf([][]int{{1, 2}, {3, 5}}),
			diff: differs.Chain(zero, differs.SliceValues(zero, 1), differs.SliceValues(zero, 1),
				differs.Values(zero, sprints.TypedDigit("int", 4), sprints.TypedDigit("int", 5))),
			equals: false,
			err:    nil,
		})
	})
	t.Run("Slices with structs", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name: "Slices with structs",
			a:    reflect.ValueOf([]struct{ Name string }{{"Alice"}, {"Bob"}}),
			b:    reflect.ValueOf([]struct{ Name string }{{"Alice"}, {"Charlie"}}),
			diff: differs.Message(differs.Chain(zero, differs.SliceValues(zero, 1), differs.StructFields(zero, "Name"),
				differs.Strings(zero, 0)), differs.Diff(zero.Plus(2), fmx.SRedf("Bob"), "Charlie")),
			equals: false,
			err:    nil,
		})
	})
}

func TestValues_Strings(t *testing.T) {
	zero := indent.Zero()

	t.Run("Equal strings", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:   "Equal strings",
			a:      reflect.ValueOf("hello"),
			b:      reflect.ValueOf("hello"),
			diff:   differs.Empty(zero),
			equals: true,
			err:    nil,
		})
	})
	t.Run("Both strings are empty", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:   "Both strings are empty",
			a:      reflect.ValueOf(""),
			b:      reflect.ValueOf(""),
			diff:   differs.Empty(zero),
			equals: true,
			err:    nil,
		})
	})

	// Different strings
	t.Run("Different lengths", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name: "Different lengths",
			a:    reflect.ValueOf("hello"),
			b:    reflect.ValueOf("hell"),
			diff: differs.Message(differs.Strings(zero, 4), differs.Diff(zero, "hell"+fmx.SRedf("o"),
				"hell")),
			equals: false,
			err:    nil,
		})
	})
	t.Run("Same length, different characters", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name: "Same length, different characters",
			a:    reflect.ValueOf("hello"),
			b:    reflect.ValueOf("hallo"),
			diff: differs.Message(differs.Strings(zero, 1), differs.Diff(zero, "h"+
				fmx.SRedf("ello"), "hallo")),
			equals: false,
			err:    nil,
		})
	})
	t.Run("Unicode characters", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name: "Unicode characters",
			a:    reflect.ValueOf("こんにちは"), // Japanese greeting
			b:    reflect.ValueOf("こんばんは"), // Japanese evening greeting
			diff: differs.Message(differs.Strings(zero, 2), differs.Diff(zero, "こん"+
				fmx.SRedf("にちは"), "こんばんは")), // "Character mismatch at index 2: 'に' vs 'ば'",
			equals: false,
			err:    nil,
		})
	})

	// Edge cases
	t.Run("Empty string vs non-empty string", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:   "Empty string vs non-empty string",
			a:      reflect.ValueOf(""),
			b:      reflect.ValueOf("non-empty"),
			diff:   differs.Message(differs.Strings(zero, 0), differs.Diff(zero, fmx.SRedf(""), "non-empty")),
			equals: false,
			err:    nil,
		})
	})
	t.Run("Case sensitivity", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:   "Case sensitivity",
			a:      reflect.ValueOf("Hello"),
			b:      reflect.ValueOf("hello"),
			diff:   differs.Message(differs.Strings(zero, 0), differs.Diff(zero, fmx.SRedf("Hello"), "hello")),
			equals: false,
			err:    nil,
		})
	})
}

func TestValues_Structs(t *testing.T) {
	zero := indent.Zero()
	t.Run("Equal structs", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:   "Equal structs",
			a:      reflect.ValueOf(struct{ Name string }{"Alice"}),
			b:      reflect.ValueOf(struct{ Name string }{"Alice"}),
			diff:   differs.Empty(zero),
			equals: true,
			err:    nil,
		})
	})
	t.Run("Both structs are nil", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:   "Both structs are nil",
			a:      reflect.ValueOf((*struct{ Name string })(nil)),
			b:      reflect.ValueOf((*struct{ Name string })(nil)),
			diff:   differs.Empty(zero),
			equals: true,
			err:    nil,
		})
	})

	// Different structs
	t.Run("Different field values", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name: "Different field values",
			a:    reflect.ValueOf(struct{ Name string }{"Alice"}),
			b:    reflect.ValueOf(struct{ Name string }{"Bob"}),
			diff: differs.Message(differs.Append(zero, differs.StructFields(zero, "Name"),
				differs.Strings(zero, 0)), differs.Diff(zero.Plus(1), fmx.SRedf("Alice"), "Bob")),
			equals: false,
			err:    nil,
		})
	})
	t.Run("Different field names", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name: "Different field names",
			a:    reflect.ValueOf(struct{ Name string }{"Alice"}),
			b:    reflect.ValueOf(struct{ FullName string }{"Alice"}),
			diff: zero.Smarkf(differs.StructTypesMismatch(zero, "struct { Name string }",
				"struct { FullName string }")),
			equals: false,
			err:    nil,
		})
	})
	t.Run("Different field types", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name: "Different field types",
			a:    reflect.ValueOf(struct{ Age int }{25}),
			b:    reflect.ValueOf(struct{ Age string }{"25"}),
			diff: zero.Smarkf(differs.StructTypesMismatch(zero, "struct { Age int }",
				"struct { Age string }")),
			equals: false,
			err:    nil,
		})
	})
	t.Run("One struct is nil", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:   "One struct is nil, the other is not",
			a:      reflect.ValueOf((*struct{ Name string })(nil)),
			b:      reflect.ValueOf(struct{ Name string }{"Alice"}),
			diff:   zero.Smarkf(differs.TypesMismatch(zero, "*struct { Name string }", "struct { Name string }")),
			equals: false,
			err:    nil,
		})
	})

	// Edge cases
	t.Run("Structs with nested structs", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name: "Structs with nested structs",
			a:    reflect.ValueOf(struct{ Info struct{ Age int } }{Info: struct{ Age int }{25}}),
			b:    reflect.ValueOf(struct{ Info struct{ Age int } }{Info: struct{ Age int }{30}}),
			diff: differs.Chain(zero, differs.StructFields(zero, "Info"),
				differs.StructFields(zero, "Age"), differs.Values(zero, sprints.Typed("int", 25),
					sprints.Typed("int", 30))),
			equals: false,
			err:    nil,
		})
	})
	t.Run("Structs with slices", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name: "Structs with slices",
			a:    reflect.ValueOf(struct{ Numbers []int }{Numbers: []int{1, 2, 3}}),
			b:    reflect.ValueOf(struct{ Numbers []int }{Numbers: []int{1, 2, 4}}),
			diff: differs.Chain(zero, differs.StructFields(zero, "Numbers"), differs.SliceValues(zero, 2),
				differs.Values(zero, sprints.Typed("int", 3), sprints.Typed("int", 4))),
			equals: false,
			err:    nil,
		})
	})
	t.Run("Structs with unexported fields", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name: "Structs with unexported fields",
			a:    reflect.ValueOf(struct{ name string }{"Alice"}),
			b:    reflect.ValueOf(struct{ name string }{"Bob"}),
			diff: differs.Message(differs.Append(zero, differs.StructFields(zero, "name"),
				differs.Strings(zero, 0)), differs.Diff(zero.Plus(1), fmx.SRedf("Alice"), "Bob")),
			equals: false,
			err:    nil,
		})
	})
	t.Run("Structs with cyclic references", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:   "Structs with cyclic references",
			a:      reflect.ValueOf(structs.CyclicSimpleNode(42, 43)),
			b:      reflect.ValueOf(structs.CyclicSimpleNode(42, 43)),
			diff:   differs.Empty(zero),
			equals: true,
			err:    nil,
		})
	})
}

func TestValues_UnsafePointers(t *testing.T) {
	zero := indent.Zero()
	t.Run("", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:   "Equal unsafe.Pointers",
			a:      reflect.ValueOf(pointers.UnsafeOfZeroCloserSuccessInst),
			b:      reflect.ValueOf(pointers.UnsafeOfZeroCloserSuccessInst),
			diff:   differs.Empty(zero),
			equals: true,
			err:    nil,
		})
	})
	t.Run("", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:   "Both unsafe.Pointers are nil",
			a:      reflect.ValueOf(unsafe.Pointer(nil)),
			b:      reflect.ValueOf(unsafe.Pointer(nil)),
			diff:   differs.Empty(zero),
			equals: true,
			err:    nil,
		})
	})

	// Different unsafe.Pointers
	t.Run("", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name: "Different unsafe.Pointers",
			a:    reflect.ValueOf(pointers.Unsafe(structs.ZeroOneDataInst)),
			b:    reflect.ValueOf(pointers.Unsafe(structs.ZeroSimpleInst)),
			diff: zero.Smarkf(differs.UnsafePointersAddr(zero, sprints.UnsafeAddrf(pointers.Unsafe(structs.ZeroOneDataInst)),
				sprints.UnsafeAddrf(pointers.Unsafe(structs.ZeroSimpleInst)))),
			equals: false,
			err:    nil,
		})
	})
	t.Run("", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name: "One unsafe.Pointer is nil, the other is not",
			a:    reflect.ValueOf(unsafe.Pointer(nil)),
			b:    reflect.ValueOf(pointers.UnsafeOfZeroNaturallyComparableInst),
			diff: zero.Smarkf(differs.NilReceived(zero, "unsafe.Pointer", types.Name(reflect.
				ValueOf(pointers.UnsafeOfZeroNaturallyComparableInst).Type()))),
			equals: false,
			err:    nil,
		})
	})

	// Edge cases
	t.Run("", func(t *testing.T) {
		tt := testingtools.LoggersLitetm(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name: "uninitialized pointers",
			a: func() reflect.Value {
				var x *int
				return reflect.ValueOf(unsafe.Pointer(x))
			}(),
			b: func() reflect.Value {
				var y *int
				return reflect.ValueOf(unsafe.Pointer(y))
			}(),
			diff:   differs.Empty(zero),
			equals: true,
			err:    nil,
		})
	})
}
