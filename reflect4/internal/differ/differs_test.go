// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package differ

import (
	"math"
	"reflect"
	"testing"
	"unsafe"

	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"github.com/andrerrcosta2/gtools/core/format/custom/term"
	"github.com/andrerrcosta2/gtools/core/format/differs"
	"github.com/andrerrcosta2/gtools/core/format/fmx"
	"github.com/andrerrcosta2/gtools/core/format/sprints"
	"github.com/andrerrcosta2/gtools/core/util/typeutil/generics"
	"github.com/andrerrcosta2/gtools/core/util/typeutil/ptrs"
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static/interf"
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static/structs/models"
	"github.com/andrerrcosta2/gtools/gtests/testingtools"
	"github.com/andrerrcosta2/gtools/gtests/testingtools/config/testlogs"
	"github.com/andrerrcosta2/gtools/gtests/testingtools/themes"
	"github.com/andrerrcosta2/gtools/reflect4/internal"
	"github.com/andrerrcosta2/gtools/reflect4/internal/_testdata"
	"github.com/andrerrcosta2/gtools/reflect4/internal/types"
	"github.com/andrerrcosta2/gtools/reflect4/op/compare"
)

// TestBetween_Primitives tests the Between method against primitive values
// Primitive diffs don't change its behaviour against strategies
//
//	This test must assert:
//	- equality by value
//	- pattern messages on differences
func TestBetween_Primitives(t *testing.T) {
	var zero indent.Branch
	var strat []internal.Option

	// Test 1: Equal integers
	t.Run("Equals integers", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		tc := TestCase{
			name:    "Equals integers",
			a:       reflect.ValueOf(36),
			b:       reflect.ValueOf(36),
			options: strat,
			diff:    "",
			equals:  true,
		}
		testDifferValues(tt, tc)
	})

	// Test 2: Different integers
	t.Run("Different integers", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		tc := TestCase{
			name:    "Different integers",
			a:       reflect.ValueOf(36),
			b:       reflect.ValueOf(84),
			options: strat,
			diff: differs.Values(zero, sprints.TypedDigit("int", 36),
				sprints.TypedDigit("int", 84)),
			equals: false,
		}
		testDifferValues(tt, tc)
	})

	// Test 3: Equals floats
	t.Run("Equals floats", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		tc := TestCase{
			name:    "Equals floats",
			a:       reflect.ValueOf(3.14),
			b:       reflect.ValueOf(3.14),
			options: strat,
			diff:    "",
			equals:  true,
		}
		testDifferValues(tt, tc)
	})

	// Test 4: Different floats
	t.Run("Different floats", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		tc := TestCase{
			name:    "Different floats",
			a:       reflect.ValueOf(3.14),
			b:       reflect.ValueOf(6.28),
			options: strat,
			diff: differs.Values(zero, sprints.TypedRoundFloat("float64", 3.14),
				sprints.TypedRoundFloat("float64", 6.28)),
			equals: false,
		}
		testDifferValues(tt, tc)
	})

	// Test 5: Equals strings
	t.Run("Equals strings", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		tc := TestCase{
			name:    "Equals strings",
			a:       reflect.ValueOf("hello"),
			b:       reflect.ValueOf("hello"),
			options: strat,
			diff:    "",
			equals:  true,
		}
		testDifferValues(tt, tc)
	})

	// Test 6: Different strings
	t.Run("Different strings", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		tc := TestCase{
			name:    "Different strings",
			a:       reflect.ValueOf("hello"),
			b:       reflect.ValueOf("world"),
			options: strat,
			diff: differs.Message(differs.Strings(zero, 0),
				differs.Diff(indent.Tab(0), "'"+fmx.SRed("h")+"ello'", "'world'")),
			equals: false,
		}
		testDifferValues(tt, tc)
	})

	// Test 7: Equals booleans
	t.Run("Equals booleans", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		tc := TestCase{
			name:    "Equals booleans",
			a:       reflect.ValueOf(true),
			b:       reflect.ValueOf(true),
			options: strat,
			diff:    "",
			equals:  true,
		}
		testDifferValues(tt, tc)
	})

	// Test 8: Different booleans
	t.Run("Different booleans", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		tc := TestCase{
			name:    "Different booleans",
			a:       reflect.ValueOf(true),
			b:       reflect.ValueOf(false),
			options: strat,
			diff: differs.Values(zero, sprints.TypedBool(true),
				sprints.TypedBool(false)),
			equals: false,
		}
		testDifferValues(tt, tc)
	})

	// Test 9: Integer vs float
	t.Run("Integer vs float", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		tc := TestCase{
			name:    "Integer vs float",
			a:       reflect.ValueOf(36),
			b:       reflect.ValueOf(36.0),
			options: strat,
			diff:    differs.TypesMismatch(zero, "int", "float64"),
			equals:  false,
		}
		testDifferValues(tt, tc)
	})

	// Test 10: Invalid received value
	t.Run("Invalid received value", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		tc := TestCase{
			name:    "Invalid received value",
			a:       reflect.Value{},
			b:       reflect.ValueOf(10),
			options: strat,
			diff:    differs.InvalidReceived(zero, "int"),
			equals:  false,
		}
		testDifferValues(tt, tc)
	})

	// Test 11: Invalid expected value
	t.Run("Invalid expected value", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		tc := TestCase{
			name:    "Invalid expected value",
			a:       reflect.ValueOf(10),
			b:       reflect.Value{},
			options: strat,
			diff:    differs.InvalidExpected(zero, "int"),
			equals:  false,
		}
		testDifferValues(tt, tc)
	})
}

// TestBetween_Primitives tests the Between method against primitive edge cases
// Primitive diffs don't change its behaviour against strategies
//
//	This test must assert:
//	- equality by value
//	- pattern messages on differences
func TestBetween_Primitives_EdgeCases(t *testing.T) {
	var zero indent.Branch
	var zerostr indent.Tab
	var strat []internal.Option

	// Test 1: Zero Integers
	t.Run("Zero Integers", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		tc := TestCase{
			name:    "Zero Integers",
			a:       reflect.ValueOf(0),
			b:       reflect.ValueOf(0),
			options: strat,
			diff:    "",
			equals:  true,
		}
		testDifferValues(tt, tc)
	})

	// Test 2: Negative vs Positive
	t.Run("negative vs positive integers", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		tc := TestCase{
			name:    "negative vs positive integers",
			a:       reflect.ValueOf(-36),
			b:       reflect.ValueOf(36),
			options: strat,
			diff: differs.Values(zero,
				sprints.TypedDigit("int", -36), sprints.TypedDigit("int", 36)),
			equals: false,
		}
		testDifferValues(tt, tc)
	})

	// Test 3: Different max int
	t.Run("different max int", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		tc := TestCase{
			name:    "Different max int",
			a:       reflect.ValueOf(math.MaxInt),
			b:       reflect.ValueOf(math.MaxInt - 1),
			options: strat,
			diff: differs.Values(zero, sprints.TypedDigit("int", math.MaxInt),
				sprints.TypedDigit("int", math.MaxInt-1)),
			equals: false,
		}
		testDifferValues(tt, tc)
	})

	// Test 4: Zero Floats
	t.Run("Zero Floats", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		tc := TestCase{
			name:    "Zero Floats",
			a:       reflect.ValueOf(0.0),
			b:       reflect.ValueOf(0.0),
			options: strat,
			diff:    "",
			equals:  true,
		}
		testDifferValues(tt, tc)
	})

	// Test 5: Negative vs positive floats
	t.Run("negative vs positive floats", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		tc := TestCase{
			name:    "negative vs positive floats",
			a:       reflect.ValueOf(-3.14),
			b:       reflect.ValueOf(3.14),
			options: strat,
			diff: differs.Values(zero, sprints.TypedRoundFloat("float64", -3.14),
				sprints.TypedRoundFloat("float64", 3.14)),
			equals: false,
		}
		testDifferValues(tt, tc)
	})

	// Test 6: NaN floats
	// According to the IEEE 754 standard, NaN is not compare to anything , including itself. This means:
	//  - math.NaN() == math.NaN() // false
	//  - reflect.DeepEqual(math.NaN(), math.NaN()) // false
	t.Run("NaN floats", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		tc := TestCase{
			name:    "NaN floats",
			a:       reflect.ValueOf(math.NaN()),
			b:       reflect.ValueOf(math.NaN()),
			options: strat,
			diff: differs.Values(zero, sprints.Typed("float64", math.NaN()),
				sprints.Typed("float64", math.NaN())),
			equals: false,
		}
		testDifferValues(tt, tc)
	})

	// Test 7: Infinity floats
	t.Run("Infinity floats", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		tc := TestCase{
			name:    "Infinity floats",
			a:       reflect.ValueOf(math.Inf(1)),
			b:       reflect.ValueOf(math.Inf(-1)),
			options: strat,
			diff: differs.Values(zero, sprints.TypedRoundFloat("float64", math.Inf(1)),
				sprints.TypedRoundFloat("float64", math.Inf(-1))),
			equals: false,
		}
		testDifferValues(tt, tc)
	})

	// Test 8: Different max float
	t.Run("Different max float", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		tc := TestCase{
			name:    "Different max float",
			a:       reflect.ValueOf(math.MaxFloat64),
			b:       reflect.ValueOf(math.MaxFloat64 * -1),
			options: strat,
			diff: differs.Values(zero, sprints.TypedRoundFloat("float64", math.MaxFloat64),
				sprints.TypedRoundFloat("float64", math.MaxFloat64*-1)),
			equals: false,
		}
		testDifferValues(tt, tc)
	})

	// Test 9: Zero Complex
	t.Run("Zero Complex", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		tc := TestCase{
			name:    "Zero Complex",
			a:       reflect.ValueOf(complex(0, 0)),
			b:       reflect.ValueOf(complex(0, 0)),
			options: strat,
			diff:    "",
			equals:  true,
		}
		testDifferValues(tt, tc)
	})

	// Test 10: Negative vs positive complex
	t.Run("negative vs positive complex", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		tc := TestCase{
			name:    "negative vs positive complex",
			a:       reflect.ValueOf(complex(-3.14, 3.14)),
			b:       reflect.ValueOf(complex(3.14, -3.14)),
			options: strat,
			diff: differs.Values(zero, sprints.TypedComplex128(complex(-3.14, 3.14)),
				sprints.TypedComplex128(complex(3.14, -3.14))),
			equals: false,
		}
		testDifferValues(tt, tc)
	})

	// Test 11: Empty strings
	t.Run("Empty strings", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		tc := TestCase{
			name:    "Empty strings",
			a:       reflect.ValueOf(""),
			b:       reflect.ValueOf(""),
			options: strat,
			diff:    "",
			equals:  true,
		}
		testDifferValues(tt, tc)
	})

	// Test 12: Whitespace strings
	t.Run("Whitespace strings", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		tc := TestCase{
			name:    "Whitespace strings",
			a:       reflect.ValueOf(" "),
			b:       reflect.ValueOf(""),
			options: strat,
			diff: differs.Message(differs.Strings(zerostr, 0), differs.Diff(zerostr,
				"'"+fmx.SCustom(term.BgBrightRed, " ")+"'",
				"''")),
			equals: false,
		}
		testDifferValues(tt, tc)
	})

	// Test 12: Case-sensitive strings
	t.Run("Case-sensitive strings", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		tc := TestCase{
			name:    "Case-sensitive strings",
			a:       reflect.ValueOf("Hello"),
			b:       reflect.ValueOf("hello"),
			options: strat,
			diff: differs.Message(differs.Strings(zerostr, 0), differs.Diff(zerostr,
				"'"+fmx.SRed("H")+"ello'", "'hello'")),
			equals: false,
		}
		testDifferValues(tt, tc)
	})

	// Test 13: Equals Unicode strings
	t.Run("Equals Unicode strings", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		tc := TestCase{
			name:    "Equals Unicode strings",
			a:       reflect.ValueOf("こんにちは"),
			b:       reflect.ValueOf("こんにちは"),
			options: strat,
			diff:    "",
			equals:  true,
		}
		testDifferValues(tt, tc)
	})

	// Test 14: Different Unicode strings
	t.Run("Different Unicode strings", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		tc := TestCase{
			name:    "Different Unicode strings",
			a:       reflect.ValueOf("こんには"),
			b:       reflect.ValueOf("こんにちは"),
			options: strat,
			diff: differs.Message(differs.Strings(zerostr, 3),
				differs.Diff(zerostr, "'こんに"+fmx.SRedf("は")+"'", "'こんにちは'")),
			equals: false,
		}
		testDifferValues(tt, tc)
	})

	// Test 15: Invalid received value
	t.Run("Invalid received value", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		tc := TestCase{
			name:    "Invalid received value",
			a:       reflect.Value{},
			b:       reflect.ValueOf(36),
			options: strat,
			diff:    differs.InvalidReceived(zero, "int"),
			equals:  false,
		}
		testDifferValues(tt, tc)
	})
}

// TestBetween_Arrays_SingleDepth_DefaultStrategy tests the Between function against arrays
// based on the default Strategy
//
//	this test must assert:
//	1. Arrays are equals only when all the conditions below are met:
//	a: same length
//	a: same type signature
//	a: same element values in depth or both are nil - default strategy
//	2. Retrieve the standard messages on defaultStringDiff
func TestBetween_Arrays_SingleDepth_DefaultStrategy(t *testing.T) {
	var zero indent.Branch
	var zerostr indent.Tab
	var strat []internal.Option

	// Test 1: Equal integer arrays
	t.Run("compare integer arrays", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:    "Equals integer arrays",
			a:       reflect.ValueOf([...]int{1, 2, 3}),
			b:       reflect.ValueOf([...]int{1, 2, 3}),
			options: strat,
			diff:    "",
			equals:  true,
		})
	})

	// Test 2: Equal Float arrays
	t.Run("compare float arrays", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:    "Equals float arrays",
			a:       reflect.ValueOf([...]float64{1.1, 2.2, 3.3}),
			b:       reflect.ValueOf([...]float64{1.1, 2.2, 3.3}),
			options: strat,
			diff:    "",
			equals:  true,
		})
	})

	// Test 3: Equal string arrays
	t.Run("compare string arrays", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:    "Equals string arrays",
			a:       reflect.ValueOf([...]string{"a", "b", "c"}),
			b:       reflect.ValueOf([...]string{"a", "b", "c"}),
			options: strat,
			diff:    "",
			equals:  true,
		})
	})

	// Test 4: Different integer arrays
	t.Run("different integer arrays", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:    "Different integer arrays",
			a:       reflect.ValueOf([...]int{1, 2, 3}),
			b:       reflect.ValueOf([...]int{1, 2, 4}),
			options: strat,
			diff: differs.Append(zero, differs.ArrayElem(zero, 2), differs.Values(zero.Inc(),
				sprints.TypedDigit("int", 3), sprints.TypedDigit("int", 4))),
			equals: false,
		})
	})

	// Test 5: Different float arrays
	t.Run("different float arrays", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:    "Different float arrays",
			a:       reflect.ValueOf([...]float64{1.1, 2.2, 3.3}),
			b:       reflect.ValueOf([...]float64{1.1, 2.2, 3.4}),
			options: strat,
			diff: differs.Message(differs.Chain(differs.ArrayElem(zero, 2), differs.Values(zero.Inc(),
				sprints.Typed("float64", "3.3"), sprints.Typed("float64", "3.4"))), ""),
			equals: false,
		})
	})

	// Test 6: Different string arrays
	t.Run("different string arrays", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:    "Different string arrays",
			a:       reflect.ValueOf([...]string{"hello", "world"}),
			b:       reflect.ValueOf([...]string{"hello", "there"}),
			options: strat,
			diff: differs.Message(differs.Chain(differs.ArrayElem(zero, 1),
				differs.Strings(zero.Inc(), 0)), differs.Diff(zerostr.Inc(), "'"+fmx.SRedf("w")+"orld'",
				"'there'")),
			equals: false,
		})
	})

	// Test 7: Arrays of different lengths
	t.Run("Arrays of different lengths", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:    "Arrays of different lengths",
			a:       reflect.ValueOf([...]int{1, 2, 3}),
			b:       reflect.ValueOf([...]int{1, 2}),
			options: strat,
			diff:    differs.Message(differs.ArrayLenMismatch(zero, 3, 2), ""),
			equals:  false,
		})
	})

	// Test 8: Empty integer arrays
	t.Run("empty integer arrays", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:    "Empty integer arrays",
			a:       reflect.ValueOf([...]int{}),
			b:       reflect.ValueOf([...]int{}),
			options: strat,
			diff:    "",
			equals:  true,
		})
	})

	// Test 9: empty string arrays
	t.Run("empty string arrays", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:    "Empty string arrays",
			a:       reflect.ValueOf([...]string{}),
			b:       reflect.ValueOf([...]string{}),
			options: strat,
			diff:    "",
			equals:  true,
		})
	})

	// Test 10: Equals nested integer arrays
	t.Run("Equals nested integer arrays", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:    "Equals nested integer arrays",
			a:       reflect.ValueOf([...][]int{{1, 2}, {3, 4}}),
			b:       reflect.ValueOf([...][]int{{1, 2}, {3, 4}}),
			options: strat,
			diff:    "",
			equals:  true,
		})
	})

	// Test 11: Different nested integer arrays
	t.Run("Different nested integer arrays", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:    "Different nested integer arrays",
			a:       reflect.ValueOf([...][]int{{1, 2}, {3, 4}}),
			b:       reflect.ValueOf([...][]int{{1, 2}, {3, 5}}),
			options: strat,
			diff: differs.Chain(differs.ArrayElem(zero, 1), differs.SliceValues(zero.Inc(), 1),
				differs.Values(zero.Plus(2), sprints.Typed("int", "4"),
					sprints.Typed("int", "5"))),
			equals: false,
		})
	})
}

// TestBetween_Arrays_EdgeCases_DefaultStrategy tests the Between function against arrays
// based on the default Strategy
//
//	this test must assert:
//	1. Arrays are equals only when all the conditions below are met:
//	a: same length
//	a: same type signature
//	a: same element values in depth or both are nil - default strategy
//	2. Retrieve the standard messages on defaultStringDiff
func TestBetween_Arrays_EdgeCases_DefaultStrategy(t *testing.T) {
	var zero indent.Branch
	var strat []internal.Option

	// Test 1: array with nil values
	t.Run("array with nil values", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:    "array with nil values",
			a:       reflect.ValueOf([...]interface{}{1, nil}),
			b:       reflect.ValueOf([...]interface{}{1, nil}),
			options: strat,
			diff:    "",
			equals:  true,
		})
	})

	// Test 2: array with mixed types
	t.Run("array with mixed types", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:    "array with mixed types",
			a:       reflect.ValueOf([...]interface{}{1, "hello"}),
			b:       reflect.ValueOf([...]interface{}{1, "world"}),
			options: strat,
			diff: differs.Message(differs.Chain(differs.ArrayElem(zero, 1), differs.InterfaceImpl(zero.Inc()),
				differs.Strings(zero.Plus(2), 0)), differs.Diff(indent.Tab(2),
				"'"+fmx.SRedf("h")+"ello'", "'world'")),
			equals: false,
		})
	})

	// Test 3: Nested array with nil slice
	t.Run("Nested array with nil slice", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:    "Nested array with nil slice",
			a:       reflect.ValueOf([...][]int{{1, 2}, nil}),
			b:       reflect.ValueOf([...][]int{{1, 2}, nil}),
			options: strat,
			diff:    "",
			equals:  true,
		})
	})

	// Test 4: Nested array with empty sub-array
	t.Run("Nested array with empty sub-array", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:    "Nested array with empty sub-array",
			a:       reflect.ValueOf([...][]int{{1, 2}, {}}),
			b:       reflect.ValueOf([...][]int{{1, 2}, {}}),
			options: strat,
			diff:    "",
			equals:  true,
		})
	})

	// Test 5: Nested array with nil sub-array
	t.Run("Nested array with nil sub-array", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:    "Nested array with nil sub-array",
			a:       reflect.ValueOf([...][]int{{1, 2}, nil}),
			b:       reflect.ValueOf([...][]int{{1, 2}, nil}),
			options: strat,
			diff:    "",
			equals:  true,
		})
	})

	// Test 6: Large integer arrays
	t.Run("Large integer arrays", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
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
			options: strat,
			diff:    "",
			equals:  true,
		})
	})

	// Test 7: Different type arrays
	t.Run("Different type arrays", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:    "Different type arrays",
			a:       reflect.ValueOf([...]int{1, 2, 3}),
			b:       reflect.ValueOf([...]float64{1.0, 2.0, 3.0}),
			options: strat,
			diff:    differs.Message(differs.ArrayTypesMismatch(zero, "[3]int", "[3]float64"), ""),
			equals:  false,
		})
	})

	// Test 8: Arrays of any with different values
	t.Run("arrays of any with different values", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:    "Different type arrays",
			a:       reflect.ValueOf([...]any{1, 2, 3}),
			b:       reflect.ValueOf([...]any{1.0, 2.0, 3.0}),
			options: strat,
			diff: differs.Chain(differs.ArrayElem(zero, 0),
				differs.InterfaceImpl(zero.Plus(1)), differs.TypesMismatch(zero.Plus(2), "int", "float64")),
			equals: false,
		})
	})

	// Test 9: Arrays of custom structs
	t.Run("arrays of custom structs", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		name := "github.com/andrerrcosta2/gtools/reflect4/internal/_testdata.Person2"
		testDifferValues(tt, TestCase{
			name:    "array of custom structs",
			a:       reflect.ValueOf([...]_testdata.Person2{{"Alice", 30}, {"Bob", 25}}),
			b:       reflect.ValueOf([...]_testdata.Person2{{"Alice", 30}, {"Charlie", 25}}),
			options: strat,
			diff: differs.Message(differs.Chain(differs.ArrayElem(zero, 1),
				differs.StructFields(zero.Plus(1), name, "Name"), differs.Strings(zero.Plus(2), 0)),
				differs.Diff(indent.Tab(2), "'"+fmx.SRedf("B")+"ob'", "'Charlie'")),
			equals: false,
		})
	})

	// Test 10: Arrays of ptrs
	t.Run("arrays of ptrs - different values", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:    "array of ptrs",
			a:       reflect.ValueOf([...]*int{ptrs.New(1), ptrs.New(2), ptrs.New(3)}),
			b:       reflect.ValueOf([...]*int{ptrs.New(1), ptrs.New(2), ptrs.New(4)}),
			options: strat,
			diff: differs.Chain(differs.ArrayElem(zero, 2), differs.PointerValues(zero.Inc()),
				differs.Values(zero.Plus(2), sprints.TypedDigit("int", 3),
					sprints.TypedDigit("int", 4))),
			equals: false,
		})
	})

	// Test 11: Arrays of uninitialized elements
	t.Run("arrays of uninitialized elements", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:    "array with uninitialized elements",
			a:       reflect.ValueOf([...]string{"hello", "", "world"}),
			b:       reflect.ValueOf([...]string{"hello", "there", "world"}),
			options: strat,
			diff: differs.Message(differs.Chain(differs.ArrayElem(zero, 1), differs.Strings(zero.Inc(), 0)),
				differs.Diff(indent.Tab(1), "''", "'there'")),
			equals: false,
		})
	})

	// Test 12: Arrays of non-comparable elements
	t.Run("arrays of non-comparable elements", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:    "array with non-comparable elements",
			a:       reflect.ValueOf([...][]int{{1, 2}, {3, 4}}),
			b:       reflect.ValueOf([...][]int{{1, 2}, {3, 5}}),
			options: strat,
			diff: differs.Chain(differs.ArrayElem(zero, 1),
				differs.SliceValues(zero.Inc(), 1), differs.Values(zero.Plus(2), sprints.Typed("int", "4"),
					sprints.Typed("int", "5"))),
			equals: false,
		})
	})
}

// TestBetween_Arrays_SingleDepth_SerializableStrategy tests the Between function against
// single depth arrays allowing nil equals empty
//
//	this test must assert:
//	1. Arrays are equals only when all the conditions below are met:
//	a: same length
//	a: same type signature
//	a: same element values in depth
//	b: both are nil or empty - serializable strategy
//	2. Retrieve the standard messages on defaultStringDiff
func TestBetween_Arrays_SingleDepth_SerializableStrategy(t *testing.T) {
	// TODO:
}

// TestBetween_Arrays_EdgeCases_SerializableStrategy tests the Between function against
// single depth arrays allowing nil equals empty
//
//	this test must assert:
//	1. Arrays are equals only when all the conditions below are met:
//	a: same length
//	a: same type signature
//	a: same element values in depth
//	b: both are nil or empty - serializable strategy
//	2. Retrieve the standard messages on defaultStringDiff
func TestBetween_Arrays_EdgeCases_SerializableStrategy(t *testing.T) {
	// TODO:
}

// TestBetween_Arrays_ErrorHandling_DefaultStrategy tests the Between error handling
// against arrays based on all strategies
//
//	this test must assert:
//	1. Arrays are equals only when all the conditions below are met:
//	a: same length
//	a: same type signature
//	a: same element values in depth or both are nil - default strategy
//	2. Retrieve the standard messages on defaultStringDiff
func TestBetween_Arrays_ErrorHandling(t *testing.T) {
	// TODO:
}

// TestBetween_Channels_DefaultStrategy tests the Between function against channels
// based on the default Strategy
//
//	This test must assert:
//	1. channels are equals only when they meet the conditions below:
//	a: both have the same signatures
//	a: both have the same capacity - their values are not compared.
//	a: both are nil - default strategy
//	b: same channel underlying addresses
//	2. retrieve the standard diff message when there are differences
func TestBetween_Channels_DefaultStrategy(t *testing.T) {
	var zero indent.Branch
	var strat []internal.Option

	t.Run("same channel instance", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		c := make(chan int, 100)
		testDifferValues(tt, TestCase{
			name:    "same channel instances",
			a:       reflect.ValueOf(&c),
			b:       reflect.ValueOf(&c),
			options: strat,
			diff:    "",
			equals:  true,
		})
	})

	t.Run("equals unbuffered channels", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:    "Equal unbuffered channels",
			a:       reflect.ValueOf(make(chan int)),
			b:       reflect.ValueOf(make(chan int)),
			options: strat,
			diff:    "",
			equals:  true,
		})
	})

	t.Run("equals buffered channels", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:    "equals buffered channels",
			a:       reflect.ValueOf(make(chan int, 5)),
			b:       reflect.ValueOf(make(chan int, 5)),
			options: strat,
			diff:    "",
			equals:  true,
		})
	})

	// Different channels
	t.Run("different channel types", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:    "Different channel types",
			a:       reflect.ValueOf(make(chan int)),
			b:       reflect.ValueOf(make(chan string)),
			options: strat,
			diff:    differs.ChanElemTypesMismatch(zero, "int", "string"),
			equals:  false,
		})
	})

	t.Run("different buffer sizes, same signature", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:    "different buffer sizes, same signature",
			a:       reflect.ValueOf(make(chan int, 5)),
			b:       reflect.ValueOf(make(chan int, 10)),
			options: strat,
			diff:    differs.ChanBufSizeMismatch(zero, 5, 10),
			equals:  false,
		})
	})
}

// TestBetween_Channels_EdgeCases_DefaultStrategy tests the Between function against edge case channels
// based on the default Strategy
//
//	This test must assert:
//	1. channels are equals only when they meet the conditions below:
//	a: both have the same signatures
//	a: both have the same capacity - their values are not compared.
//	b: same channel underlying addresses
//	c: both are nil
//	2. retrieve the standard diff message when there are differences
func TestBetween_Channels_EdgeCases_DefaultStrategy(t *testing.T) {
	var zero indent.Branch
	var strat []internal.Option

	// Edge cases
	t.Run("nil channel vs valid channel", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		// nil received
		testDifferValues(tt, TestCase{
			name:    "nil channel vs valid channel",
			a:       reflect.ValueOf(generics.Zero[chan int]()),
			b:       reflect.ValueOf(make(chan int)),
			options: strat,
			diff:    differs.NilReceived(zero, "chan", "chan int"),
			equals:  false,
		})
		// nil expected
		testDifferValues(tt, TestCase{
			name:    "valid channel vs nil channel",
			a:       reflect.ValueOf(make(chan int)),
			b:       reflect.ValueOf(generics.Zero[chan int]()),
			options: strat,
			diff:    differs.NilExpected(zero, "chan", "chan int"),
			equals:  false,
		})
	})

	t.Run("nil channels, same signature", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure)
		testDifferValues(tt, TestCase{
			name:    "nil channels, same signature",
			a:       reflect.ValueOf(ptrs.Nil[chan int]()),
			b:       reflect.ValueOf(ptrs.Nil[chan int]()),
			options: strat,
			diff:    "",
			equals:  true,
		})
	})

	// Nil channels capacity is always zero
	t.Run("nil channels, different signatures", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure)
		a := reflect.ValueOf(generics.Zero[chan int]())
		b := reflect.ValueOf(generics.Zero[chan float32]())
		testDifferValues(tt, TestCase{
			name:    "Nil channels, different signatures",
			a:       a,
			b:       b,
			options: strat,
			diff:    differs.NilChanTypesMismatch(zero, "int", "float32"),
			equals:  false,
		})
	})

	// Nil channels capacity is always zero
	t.Run("nil channels, different direction", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure)
		a := reflect.ValueOf(generics.Zero[chan<- int]())
		b := reflect.ValueOf(generics.Zero[<-chan int]())
		testDifferValues(tt, TestCase{
			name:    "Nil channels, different signatures",
			a:       a,
			b:       b,
			options: strat,
			diff:    differs.NilChanDirMismatch(zero, "chan<-", "<-chan"),
			equals:  false,
		})
	})
}

// TestBetween_Channels_IdentityStrategy tests the Between function against channels
// based on the identity Strategy
//
//	This test must assert:
//	1. channels are equals only when they meet the conditions below:
//	a: same channel underlying addresses
//	b: both are nil with the same signature - identity strategy
//	2. retrieve the standard diff message when there are differences
func TestBetween_Channels_IdentityStrategy(t *testing.T) {
	var zero indent.Branch
	strat := []internal.Option{compare.ChanIdentity}

	t.Run("same channel instance", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		c := make(chan int, 100)
		testDifferValues(tt, TestCase{
			name:    "same channel instances",
			a:       reflect.ValueOf(&c),
			b:       reflect.ValueOf(&c),
			options: strat,
			diff:    "",
			equals:  true,
		})
	})

	t.Run("different instances, equals signatures, unbuffered", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		a := reflect.ValueOf(make(chan int))
		b := reflect.ValueOf(make(chan int))
		testDifferValues(tt, TestCase{
			name:    "Equal unbuffered channels",
			a:       a,
			b:       b,
			options: strat,
			diff:    differs.ChanAddressMismatch(zero, sprints.Uintptrf(a.Pointer()), sprints.Uintptrf(b.Pointer())),
			equals:  false,
		})
	})

	t.Run("different instances, equals signatures, equals buffer size", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		a := reflect.ValueOf(make(chan int, 5))
		b := reflect.ValueOf(make(chan int, 5))
		testDifferValues(tt, TestCase{
			name:    "different instances, equals signatures, equals buffer size",
			a:       a,
			b:       b,
			options: strat,
			diff:    differs.ChanAddressMismatch(zero, sprints.Uintptrf(a.Pointer()), sprints.Uintptrf(b.Pointer())),
			equals:  false,
		})
	})

	// Different channels
	t.Run("different instances, different buffer sizes", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		a := reflect.ValueOf(make(chan int, 5))
		b := reflect.ValueOf(make(chan int, 10))
		testDifferValues(tt, TestCase{
			name:    "Different buffer sizes",
			a:       a,
			b:       b,
			options: strat,
			diff:    differs.ChanAddressMismatch(zero, sprints.Uintptrf(a.Pointer()), sprints.Uintptrf(b.Pointer())),
			equals:  false,
		})
	})
}

// TestBetween_Channels_EdgeCases_DefaultStrategy tests the Between function against edge case channels
// based on the default Strategy
//
//	This test must assert:
//	1. channels are equals only when they meet the conditions below:
//	1. channels are equals only when they meet the conditions below:
//	a: same channel underlying addresses
//	b: both are nil with the same signature - identity strategy
//	2. retrieve the standard diff message when there are differences
func TestBetween_Channels_EdgeCases_IdentityStrategy(t *testing.T) {
	var zero indent.Branch
	strat := []internal.Option{compare.ChanIdentity}

	// Edge cases
	t.Run("nil channel vs valid channel", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		nilChan := reflect.ValueOf(generics.Zero[chan int]())
		valChan := reflect.ValueOf(make(chan int))
		// nil received
		testDifferValues(tt, TestCase{
			name:    "nil channel vs valid channel",
			a:       nilChan,
			b:       valChan,
			options: strat,
			diff:    differs.NilReceived(zero, "chan", "chan int"),
			equals:  false,
		})
		// nil expected
		testDifferValues(tt, TestCase{
			name:    "valid channel vs nil channel",
			a:       valChan,
			b:       nilChan,
			options: strat,
			diff:    differs.NilExpected(zero, "chan", "chan int"),
			equals:  false,
		})
	})

	t.Run("nil channels, same signature", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure)
		testDifferValues(tt, TestCase{
			name:    "nil channels, same signature",
			a:       reflect.ValueOf(generics.Zero[chan int]()),
			b:       reflect.ValueOf(generics.Zero[chan int]()),
			options: strat,
			diff:    "",
			equals:  true,
		})
	})

	// Nil channels capacity is always zero
	t.Run("nil channels, different signatures", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure)
		a := reflect.ValueOf(generics.Zero[chan int]())
		b := reflect.ValueOf(generics.Zero[chan float32]())
		testDifferValues(tt, TestCase{
			name:    "Nil channels, different signatures",
			a:       a,
			b:       b,
			options: strat,
			diff:    differs.NilChanTypesMismatch(zero, "int", "float32"),
			equals:  false,
		})
	})
}

// TestBetween_Functions_DefaultStrategy tests the method Between against functions
// based on the default Strategy
//
//	This test must assert:
//	1. channels are equals only when the conditions below are met:
//	a: both functions have the same signature - input/output
//	b: both functions are nil with the same signature
//	c: both functions have the same underlying address pointer
//	2. Retrieve the standard messages on differences
func TestBetween_Functions_DefaultStrategy(t *testing.T) {
	var zero indent.Branch
	var strat []internal.Option

	// equals
	t.Run("same signature, different instance", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure)
		testDifferValues(tt, TestCase{
			name:    "same signature, different instance",
			a:       reflect.ValueOf(func(int) {}),
			b:       reflect.ValueOf(func(int) {}),
			options: strat,
			diff:    "",
			equals:  true,
		})
	})

	t.Run("Same function reference", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		fn := func() {}
		testDifferValues(tt, TestCase{
			name:    "Same function reference",
			a:       reflect.ValueOf(fn),
			b:       reflect.ValueOf(fn),
			options: strat,
			diff:    "",
			equals:  true,
		})
	})

	// Different functions
	t.Run("different function types", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:    "different function types",
			a:       reflect.ValueOf(func(int) {}),
			b:       reflect.ValueOf(func(string) {}),
			options: strat,
			diff:    differs.FuncTypesMismatch(zero, "func(int)", "func(string)"),
			equals:  false,
		})
	})

	t.Run("one function is nil", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		// nil received
		testDifferValues(tt, TestCase{
			name:    "nil received, valid expected",
			a:       reflect.ValueOf(generics.Zero[func()]()),
			b:       reflect.ValueOf(func() {}),
			options: strat,
			diff:    differs.NilReceived(zero, "func", "func()"),
			equals:  false,
		})
		// nil expected
		testDifferValues(tt, TestCase{
			name:    "valid received, nil expected",
			a:       reflect.ValueOf(func() {}),
			b:       reflect.ValueOf(generics.Zero[func()]()),
			options: strat,
			diff:    differs.NilExpected(zero, "func", "func()"),
			equals:  false,
		})
	})

	t.Run("functions with different signatures", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:    "functions with different signatures",
			a:       reflect.ValueOf(func(int, string) {}),
			b:       reflect.ValueOf(func(string, int) {}),
			options: strat,
			diff:    differs.FuncTypesMismatch(zero, "func(int, string)", "func(string, int)"),
			equals:  false,
		})
	})
}

// TestBetween_Functions_EdgeCases_DefaultStrategy tests the method Between against edge case
// functions based on the default Strategy
//
//	This test must assert:
//	1. channels are equals only when the conditions below are met:
//	a: both functions have the same signature - input/output
//	b: both functions are nil
//	c: both functions have the same underlying address pointer
//	2. Retrieve the standard messages on differences
func TestBetween_Functions_EdgeCases_DefaultStrategy(t *testing.T) {
	var zero indent.Branch
	var strat []internal.Option

	// Edge cases
	t.Run("both nil, same signature", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:    t.Name(),
			a:       reflect.ValueOf((func())(nil)),
			b:       reflect.ValueOf((func())(nil)),
			options: strat,
			diff:    "",
			equals:  true,
		})
	})

	t.Run("both nil, different signature", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:    "Two different nil functions",
			a:       reflect.ValueOf(generics.Zero[func()]()),
			b:       reflect.ValueOf(generics.Zero[func() int]()),
			options: strat,
			diff:    differs.NilFuncTypesMismatch(zero, "func()", "func() int"),
			equals:  false,
		})
	})
}

// TestBetween_Functions_IdentityStrategy tests the method Between against functions
// based on the identity Strategy
//
//	This test must assert:
//	1. channels are equals only when the conditions below are met:
//	a: both functions have the same underlying address pointer
//	b: both functions are nil with the same signature
//	2. Retrieve the standard messages on differences
func TestBetween_Functions_IdentityStrategy(t *testing.T) {
	var zero indent.Branch
	strat := []internal.Option{compare.FuncIdentity}

	// equals
	t.Run("same signature, different instance", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure)
		a := reflect.ValueOf(func(int) {})
		b := reflect.ValueOf(func(int) {})
		testDifferValues(tt, TestCase{
			name:    "same signature, different instance",
			a:       a,
			b:       b,
			options: strat,
			diff: differs.FuncAddressMismatch(zero, sprints.Uintptrf(a.Pointer()),
				sprints.Uintptrf(b.Pointer())),
			equals: false,
		})
	})

	t.Run("Same function reference", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		fn := func() {}
		testDifferValues(tt, TestCase{
			name:    "Same function reference",
			a:       reflect.ValueOf(fn),
			b:       reflect.ValueOf(fn),
			options: strat,
			diff:    "",
			equals:  true,
		})
	})

	// Different functions
	t.Run("different function types", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		a := reflect.ValueOf(func(int) {})
		b := reflect.ValueOf(func(string) {})
		testDifferValues(tt, TestCase{
			name:    "different function types",
			a:       a,
			b:       b,
			options: strat,
			diff: differs.FuncAddressMismatch(zero, sprints.Uintptrf(a.Pointer()),
				sprints.Uintptrf(b.Pointer())),
			equals: false,
		})
	})

	t.Run("one function is nil", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		// nil received
		testDifferValues(tt, TestCase{
			name:    "nil received, valid expected",
			a:       reflect.ValueOf(generics.Zero[func()]()),
			b:       reflect.ValueOf(func() {}),
			options: strat,
			diff:    differs.NilReceived(zero, "func", "func()"),
			equals:  false,
		})
		// nil expected
		testDifferValues(tt, TestCase{
			name:    "valid received, nil expected",
			a:       reflect.ValueOf(func() {}),
			b:       reflect.ValueOf(generics.Zero[func()]()),
			options: strat,
			diff:    differs.NilExpected(zero, "func", "func()"),
			equals:  false,
		})
	})
}

// TestBetween_Functions_EdgeCases_DefaultStrategy tests the method Between against edge case
// functions based on the default Strategy
//
//	This test must assert:
//	1. channels are equals only when the conditions below are met:
//	a: both functions have the same underlying address pointer
//	b: both functions are nil with the same signature
//	2. Retrieve the standard messages on differences
func TestBetween_Functions_EdgeCases_IdentituStrategy(t *testing.T) {
	var zero indent.Branch
	strat := []internal.Option{compare.FuncIdentity}

	// Edge cases
	t.Run("both nil, same signature", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:    t.Name(),
			a:       reflect.ValueOf((func())(nil)),
			b:       reflect.ValueOf((func())(nil)),
			options: strat,
			diff:    "",
			equals:  true,
		})
	})

	t.Run("both nil, different signature", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:    "Two different nil functions",
			a:       reflect.ValueOf(generics.Zero[func()]()),
			b:       reflect.ValueOf(generics.Zero[func() int]()),
			options: strat,
			diff:    differs.NilFuncTypesMismatch(zero, "func()", "func() int"),
			equals:  false,
		})
	})
}

// TestBetween_Interfaces tests the method Between against interfaces
// there is no defaultStringDiff between strategies for interfaces
//
//	This test must assert
//	1. interfaces are equals only when the conditions below are met:
//	a: their signatures are equals
//	a: the element they hold are equals or have the same address
//	b: both are nil with the same signature
//	2. Retrieve the standard messages on differences
func TestBetween_Interfaces(t *testing.T) {
	var zero indent.Branch
	var strat []internal.Option
	boolean := reflect.ValueOf(ptrs.New(interf.BoolTypeCastable(2, true))).Elem()
	float := reflect.ValueOf(ptrs.New(interf.FloatTypeCastable(1, 45))).Elem()

	t.Run("same type and value", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:    t.Name(),
			a:       reflect.ValueOf(ptrs.New[interface{}](36)).Elem(),
			b:       reflect.ValueOf(ptrs.New[interface{}](36)).Elem(),
			options: strat,
			diff:    differs.Empty(zero),
			equals:  true,
		})
	})

	t.Run("same interfaces, different implementations", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:    t.Name(),
			a:       boolean,
			b:       float,
			options: strat,
			diff: differs.Message(differs.Chain(differs.InterfaceImpl(zero), differs.PointerValues(zero.Plus(1)),
				differs.StructTypesMismatch(zero.Plus(2), types.Name(boolean.Elem().Elem().Type()),
					types.Name(float.Elem().Elem().Type()))), ""),
			equals: false,
		})
	})

	t.Run("same interfaces, same impl, different fields", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		snm := "github.com/andrerrcosta2/gtools/gtests/testingseeds/static/structs/models.SimpleUnsafeCastableFloat"
		testDifferValues(tt, TestCase{
			name:    "Same type, different values",
			a:       reflect.ValueOf(ptrs.New(interf.FloatTypeCastable(1, 36))).Elem(),
			b:       reflect.ValueOf(ptrs.New(interf.FloatTypeCastable(1, 38))).Elem(),
			options: strat,
			diff: differs.Chain(differs.InterfaceImpl(zero), differs.PointerValues(zero.Plus(1)),
				differs.StructFields(zero.Plus(2), snm, "ValueC"),
				differs.Values(zero.Plus(3), sprints.Typed("float64", 36),
					sprints.Typed("float64", 38))),
			equals: false,
		})
	})

	t.Run("One interface is nil", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		// nil received
		testDifferValues(tt, TestCase{
			name:    "nil received",
			a:       reflect.ValueOf(ptrs.Nil[interface{}]()).Elem(),
			b:       reflect.ValueOf(ptrs.New(interface{}(36))).Elem(),
			options: strat,
			diff:    differs.InvalidReceived(zero, "interface {}"),
			equals:  false,
		})
		// nil expected
		testDifferValues(tt, TestCase{
			name:    "nil expected",
			a:       reflect.ValueOf(ptrs.New(interface{}(36))).Elem(),
			b:       reflect.ValueOf(ptrs.Nil[interface{}]()).Elem(),
			options: strat,
			diff:    differs.InvalidExpected(zero, "interface {}"),
			equals:  false,
		})
	})

	t.Run("same interfaces, different slices", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:    t.Name(),
			a:       reflect.ValueOf(ptrs.New(interface{}([]int{1, 2, 3}))).Elem(),
			b:       reflect.ValueOf(ptrs.New(interface{}([]int{1, 2, 4}))).Elem(),
			options: strat,
			diff: differs.Chain(differs.InterfaceImpl(zero), differs.SliceValues(zero.Plus(1), 2),
				differs.Values(zero.Plus(2), sprints.Typed("int", 3), sprints.Typed("int", 4))),
			equals: false,
		})
	})

	t.Run("same interfaces, different channel signatures", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:    "Interfaces holding channels",
			a:       reflect.ValueOf(ptrs.New(interface{}(make(chan int)))).Elem(),
			b:       reflect.ValueOf(ptrs.New(interface{}(make(chan string)))).Elem(),
			options: strat,
			diff: differs.InterfaceImpl(zero) + "\n" +
				differs.ChanElemTypesMismatch(zero.Inc(), "int", "string"),
			equals: false,
		})
	})
}

// TestBetween_Interfaces tests the method Between against edge cases of interfaces
// there is no defaultStringDiff between strategies for interfaces
//
//	This test must assert
//	1. interfaces are equals only when the conditions below are met:
//	a: their signatures are equals
//	a: the element they hold are equals or have the same address
//	b: both are nil with the same signature
//	2. Retrieve the standard messages on differences
func TestBetween_Interfaces_EdgeCases(t *testing.T) {
	var zero indent.Branch
	var strat []internal.Option

	t.Run("Both interfaces are nil", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:    "Both interfaces are nil",
			a:       reflect.ValueOf(ptrs.Nil[interface{}]()).Elem(),
			b:       reflect.ValueOf(ptrs.Nil[interface{}]()).Elem(),
			diff:    differs.BothInvalid(zero),
			options: strat,
			equals:  false,
		})
	})

	t.Run("nil vs nil, different interfaces", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:    t.Name(),
			a:       reflect.ValueOf(ptrs.Nil[interface{}]()).Elem(),
			b:       reflect.ValueOf(ptrs.Nil[interf.Closer]()).Elem(),
			options: strat,
			diff:    differs.BothInvalid(zero),
			equals:  false,
		})
	})

	t.Run("generic interfaces, different structs fields", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:    t.Name(),
			a:       reflect.ValueOf(ptrs.New(interface{}(struct{ Name string }{"Alice"}))).Elem(),
			b:       reflect.ValueOf(ptrs.New(interface{}(struct{ Name string }{"Bob"}))).Elem(),
			options: strat,
			diff: differs.Message(differs.Chain(differs.InterfaceImpl(zero),
				differs.StructFields(zero.Plus(1), "struct { Name string }", "Name"), differs.Strings(zero.Plus(2), 0)),
				differs.Diff(indent.Tab(2), "'"+fmx.SRed("A")+"lice'", "'Bob'")),
			equals: false,
		})
	})
}

// TestBetween_Maps_DefaultStrategy tests the method Between against maps
// based on the default Strategy
//
//	This test must assert
//	1. maps are equals only when the conditions below are met:
//	a: their signatures are equals
//	a: their lengths are equals
//	a: each pair key-value they hold are equals or have the same address
//	b: both are nil with the same signature
//	2. Retrieve the standard messages on differences
func TestBetween_Maps_DefaultStrategy(t *testing.T) {
	var zero indent.Branch
	var strat []internal.Option

	t.Run("same instances", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		m := make(map[string]int)
		testDifferValues(tt, TestCase{
			name:    t.Name(),
			a:       reflect.ValueOf(m),
			b:       reflect.ValueOf(m),
			options: strat,
			diff:    differs.Empty(zero),
			equals:  true,
		})
	})

	t.Run("different instances, same signature, length and values", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:    t.Name(),
			a:       reflect.ValueOf(map[string]int{"a": 1, "b": 2}),
			b:       reflect.ValueOf(map[string]int{"a": 1, "b": 2}),
			options: strat,
			diff:    differs.Empty(zero),
			equals:  true,
		})
	})

	t.Run("different instances, same signatures, different key values", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name: t.Name(),
			a:    reflect.ValueOf(map[string]int{"a": 1, "b": 2}),
			b:    reflect.ValueOf(map[string]int{"a": 1, "c": 2}),
			diff: differs.Message(differs.MapKeys(zero), differs.MapKeysDiff(indent.Tab(0), []string{"<string>c"},
				[]string{"<string>b"}, []string{"<string>b"}, []string{"<string>c"})),
			equals: false,
		})
	})

	t.Run("Different values", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name: "Different values",
			a:    reflect.ValueOf(map[string]int{"a": 1, "b": 2}),
			b:    reflect.ValueOf(map[string]int{"a": 1, "b": 3}),
			diff: differs.Chain(differs.MapValue(zero, "b"), differs.Values(zero.Inc(),
				sprints.Typed("int", 2), sprints.Typed("int", 3))),
			equals: false,
		})
	})

	t.Run("Different sizes", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name: "Different sizes",
			a:    reflect.ValueOf(map[string]int{"a": 1, "b": 2}),
			b:    reflect.ValueOf(map[string]int{"a": 1}),
			diff: differs.Message(differs.MapLenMismatch(zero, 2, 1),
				differs.MapKeysDiff(indent.Tab(0), []string{}, []string{"b"}, []string{"b"}, []string{})),
			equals: false,
		})
	})

	t.Run("One map is nil", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		// nil received
		testDifferValues(tt, TestCase{
			name:   "nil received, valid expected",
			a:      reflect.ValueOf((map[string]int)(nil)),
			b:      reflect.ValueOf(map[string]int{"a": 1}),
			diff:   differs.NilReceived(zero, "map", "map[string]int"),
			equals: false,
		})
		// nil expected
		testDifferValues(tt, TestCase{
			name:   "valid received, nil expected",
			a:      reflect.ValueOf(map[string]int{"a": 1}),
			b:      reflect.ValueOf((map[string]int)(nil)),
			diff:   differs.NilExpected(zero, "map", "map[string]int"),
			equals: false,
		})
	})
}

// TestBetween_Maps_DefaultStrategy tests the method Between against edge case maps
// based on the default Strategy
//
//	This test must assert
//	1. maps are equals only when the conditions below are met:
//	a: their signatures are equals
//	a: their lengths are equals
//	a: each pair key-value they hold are equals or have the same address
//	b: both are nil with the same signature
//	2. Retrieve the standard messages on differences
func TestBetween_Maps_EdgeCases_DefaultStrategy(t *testing.T) {
	var root indent.Branch
	var strat []internal.Option

	t.Run("different instances, both nil", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:    t.Name(),
			a:       reflect.ValueOf(ptrs.Nil[map[string]int]()),
			b:       reflect.ValueOf(ptrs.Nil[map[string]int]()),
			options: strat,
			diff:    differs.Empty(root),
			equals:  true,
		})
	})

	t.Run("same signatures, nil vs empty", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		// nil received, empty expected
		testDifferValues(tt, TestCase{
			name:   "nil received, empty expected",
			a:      reflect.ValueOf(generics.Zero[map[string]int]()),
			b:      reflect.ValueOf(map[string]int{}),
			diff:   differs.NilReceived(root, "map", "map[string]int"),
			equals: false,
		})
		// empty received, nil expected
		testDifferValues(tt, TestCase{
			name:   "empty received, nil expected",
			a:      reflect.ValueOf(map[string]int{}),
			b:      reflect.ValueOf(generics.Zero[map[string]int]()),
			diff:   differs.NilExpected(root, "map", "map[string]int"),
			equals: false,
		})
	})

	t.Run("different instances, different nested slice values as elements", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name: t.Name(),
			a:    reflect.ValueOf(map[string][]int{"a": {1, 2}, "b": {3, 4}}),
			b:    reflect.ValueOf(map[string][]int{"a": {1, 2}, "b": {3, 5}}),
			diff: differs.Chain(differs.MapValue(root, "b"), differs.SliceValues(root.Plus(1), 1),
				differs.Values(root.Plus(2), sprints.Typed("int", 4), sprints.Typed("int", 5))),
			equals: false,
		})
	})

	t.Run("different instances, same signature, different non-comparable values", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name: t.Name(),
			a:    reflect.ValueOf(map[string][]int{"a": {1, 2}}),
			b:    reflect.ValueOf(map[string][]int{"a": {1, 3}}),
			diff: differs.Chain(differs.MapValue(root, "a"), differs.SliceValues(root.Plus(1), 1),
				differs.Values(root.Plus(2), sprints.Typed("int", 2), sprints.Typed("int", 3))),
			equals: false,
		})
	})

	t.Run("different instances, same signature, different struct elements", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name: "maps with struct values",
			a:    reflect.ValueOf(map[string]struct{ Name string }{"a": {"Alice"}, "b": {"Bob"}}),
			b:    reflect.ValueOf(map[string]struct{ Name string }{"a": {"Alice"}, "b": {"Charlie"}}),
			diff: differs.Message(differs.Chain(differs.MapValue(root, "b"),
				differs.StructFields(root.Plus(1), "struct { Name string }", "Name"),
				differs.Strings(root.Plus(2), 0)),
				differs.Diff(indent.Tab(2), "'"+fmx.SRed("B")+"ob'", "'Charlie'")),
			equals: false,
		})
	})
}

// TestBetween_Maps_SerializableStrategy tests the method Between against maps
// based on the serializable Strategy
//
//	This test must assert
//	1. maps are equals only when the conditions below are met:
//	a: their signatures are equals
//	a: their lengths are equals
//	a: each pair key-value they hold are equals or have the same address
//	b: both are nil or empty with the same signature
//	2. Retrieve the standard messages on differences
func TestBetween_Maps_SerializableStrategy(t *testing.T) {
	var root indent.Branch
	strat := []internal.Option{compare.Serializable}

	t.Run("same instances", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		m := make(map[string]int)
		testDifferValues(tt, TestCase{
			name:    t.Name(),
			a:       reflect.ValueOf(m),
			b:       reflect.ValueOf(m),
			options: strat,
			diff:    differs.Empty(root),
			equals:  true,
		})
	})

	t.Run("different instances, same signature, length and values", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:    t.Name(),
			a:       reflect.ValueOf(map[string]int{"a": 1, "b": 2}),
			b:       reflect.ValueOf(map[string]int{"a": 1, "b": 2}),
			options: strat,
			diff:    differs.Empty(root),
			equals:  true,
		})
	})

	t.Run("different instances, same signatures, different key values", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name: t.Name(),
			a:    reflect.ValueOf(map[string]int{"a": 1, "b": 2}),
			b:    reflect.ValueOf(map[string]int{"a": 1, "c": 2}),
			diff: differs.Message(differs.MapKeys(root), differs.MapKeysDiff(indent.Tab(0), []string{"<string>c"},
				[]string{"<string>b"}, []string{"<string>b"}, []string{"<string>c"})),
			equals: false,
		})
	})

	t.Run("Different values", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name: "Different values",
			a:    reflect.ValueOf(map[string]int{"a": 1, "b": 2}),
			b:    reflect.ValueOf(map[string]int{"a": 1, "b": 3}),
			diff: differs.Chain(differs.MapValue(root, "b"), differs.Values(root.Inc(),
				sprints.Typed("int", 2), sprints.Typed("int", 3))),
			equals: false,
		})
	})

	t.Run("Different sizes", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name: "Different sizes",
			a:    reflect.ValueOf(map[string]int{"a": 1, "b": 2}),
			b:    reflect.ValueOf(map[string]int{"a": 1}),
			diff: differs.Message(differs.MapLenMismatch(root, 2, 1),
				differs.MapKeysDiff(indent.Tab(0), []string{}, []string{"b"}, []string{"b"}, []string{})),
			equals: false,
		})
	})

	t.Run("One map is nil", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		// nil received
		testDifferValues(tt, TestCase{
			name:   "nil received, valid expected",
			a:      reflect.ValueOf((map[string]int)(nil)),
			b:      reflect.ValueOf(map[string]int{"a": 1}),
			diff:   differs.NilReceived(root, "map", "map[string]int"),
			equals: false,
		})
		// nil expected
		testDifferValues(tt, TestCase{
			name:   "valid received, nil expected",
			a:      reflect.ValueOf(map[string]int{"a": 1}),
			b:      reflect.ValueOf((map[string]int)(nil)),
			diff:   differs.NilExpected(root, "map", "map[string]int"),
			equals: false,
		})
	})
}

// TestBetween_Maps_EdgeCases_SerializableStrategy tests the method Between against edge case maps
// based on the serializable Strategy
//
//	This test must assert
//	1. maps are equals only when the conditions below are met:
//	a: their signatures are equals
//	a: their lengths are equals
//	a: each pair key-value they hold are equals or have the same address or
//	b: both are nil or empty with the same signature
//	2. Retrieve the standard messages on differences
func TestBetween_Maps_EdgeCases_SerializableStrategy(t *testing.T) {
	var root indent.Branch
	strat := []internal.Option{compare.Serializable}

	t.Run("same sign, both nil", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:    t.Name(),
			a:       reflect.ValueOf(generics.Zero[map[string]int]()),
			b:       reflect.ValueOf(generics.Zero[map[string]int]()),
			options: strat,
			diff:    differs.Empty(root),
			equals:  true,
		})
	})

	t.Run("different sign, both nil", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:    t.Name(),
			a:       reflect.ValueOf(generics.Zero[map[string]int]()),
			b:       reflect.ValueOf(generics.Zero[map[any]int]()),
			options: strat,
			diff:    differs.NilTypesMismatch(root, "maps", "map[string]int", "map[interface {}]int"),
			equals:  false,
		})
	})

	t.Run("same signatures, nil vs empty", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		// nil received, empty expected
		testDifferValues(tt, TestCase{
			name:    "nil received, empty expected",
			a:       reflect.ValueOf(generics.Zero[map[string]int]()),
			b:       reflect.ValueOf(map[string]int{}),
			diff:    differs.Empty(root),
			options: strat,
			equals:  true,
		})
		// empty received, nil expected
		testDifferValues(tt, TestCase{
			name:    "empty received, nil expected",
			a:       reflect.ValueOf(map[string]int{}),
			b:       reflect.ValueOf(generics.Zero[map[string]int]()),
			diff:    differs.Empty(root),
			options: strat,
			equals:  true,
		})
	})

	t.Run("different instances, different nested slice values as elements", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name: t.Name(),
			a:    reflect.ValueOf(map[string][]int{"a": {1, 2}, "b": {3, 4}}),
			b:    reflect.ValueOf(map[string][]int{"a": {1, 2}, "b": {3, 5}}),
			diff: differs.Chain(differs.MapValue(root, "b"), differs.SliceValues(root.Plus(1), 1),
				differs.Values(root.Plus(2), sprints.Typed("int", 4), sprints.Typed("int", 5))),
			equals: false,
		})
	})
	t.Run("different instances, same signature, different non-comparable values", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name: t.Name(),
			a:    reflect.ValueOf(map[string][]int{"a": {1, 2}}),
			b:    reflect.ValueOf(map[string][]int{"a": {1, 3}}),
			diff: differs.Chain(differs.MapValue(root, "a"), differs.SliceValues(root.Plus(1), 1),
				differs.Values(root.Plus(2), sprints.Typed("int", 2), sprints.Typed("int", 3))),
			equals: false,
		})
	})

	t.Run("different instances, same signature, different struct elements", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name: "maps with struct values",
			a:    reflect.ValueOf(map[string]struct{ Name string }{"a": {"Alice"}, "b": {"Bob"}}),
			b:    reflect.ValueOf(map[string]struct{ Name string }{"a": {"Alice"}, "b": {"Charlie"}}),
			diff: differs.Message(differs.Chain(differs.MapValue(root, "b"),
				differs.StructFields(root.Plus(1), "struct { Name string }", "Name"), differs.Strings(root.Plus(2), 0)),
				differs.Diff(indent.Tab(2), "'"+fmx.SRed("B")+"ob'", "'Charlie'")),
			equals: false,
		})
	})
}

// TestBetween_Pointers tests the method Between against ptrs
// based on the default Strategy
//
//	This test must assert:
//	1. ptrs are equals only when the conditions below are met:
//	a: their elements are equals or
//	b: their addresses are equals or
//	c: both are nil with the same signature
//	2. Retrieve the standard messages on differences
func TestBetween_Pointers_DefaultStrategy(t *testing.T) {
	var root indent.Branch
	var strat []internal.Option

	t.Run("same instances", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		x := ptrs.New(10)
		testDifferValues(tt, TestCase{
			name:    t.Name(),
			a:       reflect.ValueOf(x),
			b:       reflect.ValueOf(x),
			options: strat,
			diff:    differs.Empty(root),
			equals:  true,
		})
	})

	t.Run("different instances, same values", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:   t.Name(),
			a:      reflect.ValueOf(ptrs.New(36)),
			b:      reflect.ValueOf(ptrs.New(36)),
			diff:   differs.Empty(root),
			equals: true,
		})
	})

	t.Run("different instances and values", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name: t.Name(),
			a:    reflect.ValueOf(ptrs.New(36)),
			b:    reflect.ValueOf(ptrs.New(38)),
			diff: differs.Append(root, differs.PointerValues(root),
				differs.Values(root.Inc(), sprints.Typed("int", 36), sprints.Typed("int", 38))),
			equals: false,
		})
	})
}

// TestBetween_Pointers tests the method Between against edge case ptrs
// based on the default Strategy
//
//	This test must assert:
//	1. ptrs are equals only when the conditions below are met:
//	a: their elements are equals or
//	b: their addresses are equals or
//	c: both are nil with the same signature
//	2. Retrieve the standard messages on differences
func TestBetween_Pointers_EdgeCases_DefaultStrategy(t *testing.T) {
	var branch indent.Branch
	var strat []internal.Option

	t.Run("both nil, same signatures", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:    t.Name(),
			a:       reflect.ValueOf(ptrs.Nil[int]()),
			b:       reflect.ValueOf(ptrs.Nil[int]()),
			options: strat,
			diff:    differs.Empty(branch),
			equals:  true,
		})
	})

	t.Run("both nil, different signatures", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		a := reflect.ValueOf(ptrs.Nil[int]())
		b := reflect.ValueOf(ptrs.Nil[string]())
		testDifferValues(tt, TestCase{
			name:    t.Name(),
			a:       a,
			b:       b,
			options: strat,
			diff:    differs.NilPointersSignMismatch(branch, a.Type().String(), b.Type().String()),
			equals:  false,
		})
	})

	t.Run("one is nil, other is valid", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure)
		nilPtr := reflect.ValueOf(ptrs.Nil[int]())
		valPtr := reflect.ValueOf(ptrs.New(10))
		// nil received
		testDifferValues(tt, TestCase{
			name:    "nil received, valid expected",
			a:       nilPtr,
			b:       valPtr,
			options: strat,
			diff:    differs.NilReceived(branch, "ptr", "*int"),
			equals:  false,
		})
		// nil expected
		testDifferValues(tt, TestCase{
			name:    "valid received, nil expected",
			a:       valPtr,
			b:       nilPtr,
			options: strat,
			diff:    differs.NilExpected(branch, "ptr", "*int"),
			equals:  false,
		})
	})

	t.Run("cyclic references, equals values", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:   t.Name(),
			a:      reflect.ValueOf(models.CyclicSimpleNode(36, 38)),
			b:      reflect.ValueOf(models.CyclicSimpleNode(36, 38)),
			diff:   "",
			equals: true,
		})
	})

	t.Run("cyclic references, different values", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		snm := "github.com/andrerrcosta2/gtools/gtests/testingseeds/static/structs/models.SimpleNode"
		testDifferValues(tt, TestCase{
			name: t.Name(),
			a:    reflect.ValueOf(models.CyclicSimpleNode(36, 38)),
			b:    reflect.ValueOf(models.CyclicSimpleNode(38, 36)),
			diff: differs.Chain(differs.PointerValues(branch), differs.StructFields(branch.Plus(1),
				snm, "Next"), differs.PointerValues(branch.Plus(2)),
				differs.StructFields(branch.Plus(3), snm, "Value"),
				differs.InterfaceImpl(branch.Plus(4)),
				differs.Values(branch.Plus(5), sprints.Typed("int", 38),
					sprints.Typed("int", 36))),
			equals: false,
		})
	})

	t.Run("cyclic linked list, equals values", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:   t.Name(),
			a:      reflect.ValueOf(models.CyclicLinkedList("innerA", "innerB", "innerC", "innerD")),
			b:      reflect.ValueOf(models.CyclicLinkedList("innerA", "innerB", "innerC", "innerD")),
			diff:   differs.Empty(branch),
			equals: true,
		})
	})

	t.Run("cyclic linked list, different values", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		sn1 := "github.com/andrerrcosta2/gtools/gtests/testingseeds/static/structs/models.LinkedList"
		sn2 := "github.com/andrerrcosta2/gtools/gtests/testingseeds/static/structs/models.BinaryNode"
		testDifferValues(tt, TestCase{
			name: t.Name(),
			a:    reflect.ValueOf(models.CyclicLinkedList("innerA", "innerB", "innerC", "innerD")),
			b:    reflect.ValueOf(models.CyclicLinkedList("innerA", "innerB", "innerC", "innerE")),
			diff: differs.Message(differs.Chain(differs.PointerValues(branch), differs.StructFields(branch.Plus(1),
				sn1, "Head"), differs.PointerValues(branch.Plus(2)),
				differs.StructFields(branch.Plus(3), sn2,
					"Left"), differs.PointerValues(branch.Plus(4)), differs.StructFields(branch.Plus(5),
					sn2, "Value"), differs.InterfaceImpl(branch.Plus(6)),
				differs.Strings(branch.Plus(7), 5)), differs.Diff(indent.Tab(7),
				"'inner"+fmx.SRed("D")+"'", "'innerE'")),
			equals: false,
		})
	})
}

// TestBetween_Pointers tests the method Between against ptrs
// based on the identity Strategy
//
//	This test must assert:
//	1. ptrs are equals only when the conditions below are met:
//	a: their addresses are equals or
//	b: both are nil with the same signature
//	2. Retrieve the standard messages on differences
func TestBetween_Pointers_IdentityStrategy(t *testing.T) {
	var zero indent.Branch
	strat := []internal.Option{compare.PtrIdentity}

	t.Run("same instances", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		x := ptrs.New(10)
		testDifferValues(tt, TestCase{
			name:    t.Name(),
			a:       reflect.ValueOf(x),
			b:       reflect.ValueOf(x),
			options: strat,
			diff:    differs.Empty(zero),
			equals:  true,
		})
	})

	t.Run("different instances, same values", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		a := reflect.ValueOf(ptrs.New(10))
		b := reflect.ValueOf(ptrs.New(10))
		testDifferValues(tt, TestCase{
			name:    t.Name(),
			a:       a,
			b:       b,
			options: strat,
			diff: differs.PointersAddrMismatch(zero,
				sprints.Uintptrf(a.Pointer()), sprints.Uintptrf(b.Pointer())),
			equals: false,
		})
	})

	t.Run("different instances and values", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		a := reflect.ValueOf(ptrs.New(10))
		b := reflect.ValueOf(ptrs.New(12))
		testDifferValues(tt, TestCase{
			name:    t.Name(),
			a:       a,
			b:       b,
			options: strat,
			diff: differs.PointersAddrMismatch(zero,
				sprints.Uintptrf(a.Pointer()), sprints.Uintptrf(b.Pointer())),
			equals: false,
		})
	})
}

// TestBetween_Pointers tests the method Between against edge case ptrs
// based on the identity Strategy
//
//	This test must assert:
//	1. ptrs are equals only when the conditions below are met:
//	a: their addresses are equals or
//	b: both are nil with the same signature
//	2. Retrieve the standard messages on differences
func TestBetween_Pointers_EdgeCases_IdentityStrategy(t *testing.T) {
	var zero indent.Branch
	strat := []internal.Option{compare.PtrIdentity}

	t.Run("both nil, same signatures", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:    t.Name(),
			a:       reflect.ValueOf(ptrs.Nil[int]()),
			b:       reflect.ValueOf(ptrs.Nil[int]()),
			options: strat,
			diff:    differs.Empty(zero),
			equals:  true,
		})
	})

	t.Run("both nil, different signatures", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		a := reflect.ValueOf(ptrs.Nil[int]())
		b := reflect.ValueOf(ptrs.Nil[string]())
		testDifferValues(tt, TestCase{
			name:    t.Name(),
			a:       a,
			b:       b,
			options: strat,
			diff:    differs.NilPointersSignMismatch(zero, a.Type().String(), b.Type().String()),
			equals:  false,
		})
	})

	t.Run("one is nil, other is valid", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure)
		nilPtr := reflect.ValueOf(ptrs.Nil[int]())
		valPtr := reflect.ValueOf(ptrs.New(10))
		// nil received
		testDifferValues(tt, TestCase{
			name:    "nil received, valid expected",
			a:       nilPtr,
			b:       valPtr,
			options: strat,
			diff:    differs.NilReceived(zero, "ptr", "*int"),
			equals:  false,
		})
		// nil expected
		testDifferValues(tt, TestCase{
			name:    "valid received, nil expected",
			a:       valPtr,
			b:       nilPtr,
			options: strat,
			diff:    differs.NilExpected(zero, "ptr", "*int"),
			equals:  false,
		})
	})

	t.Run("same instances, cyclic references", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		n := models.CyclicSimpleNode(36, 38)
		a := reflect.ValueOf(n)
		b := reflect.ValueOf(n)
		testDifferValues(tt, TestCase{
			name:    t.Name(),
			a:       a,
			b:       b,
			options: strat,
			diff:    differs.Empty(zero),
			equals:  true,
		})
	})

	t.Run("different instances, cyclic references, equals values", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		a := reflect.ValueOf(models.CyclicSimpleNode(36, 38))
		b := reflect.ValueOf(models.CyclicSimpleNode(36, 38))
		testDifferValues(tt, TestCase{
			name:    t.Name(),
			a:       a,
			b:       b,
			options: strat,
			diff: differs.PointersAddrMismatch(zero,
				sprints.Uintptrf(a.Pointer()), sprints.Uintptrf(b.Pointer())),
			equals: false,
		})
	})

	t.Run("different instances, cyclic linked list, equals values", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		a := reflect.ValueOf(models.CyclicLinkedList("innerA", "innerB", "innerC", "innerD"))
		b := reflect.ValueOf(models.CyclicLinkedList("innerA", "innerB", "innerC", "innerD"))
		testDifferValues(tt, TestCase{
			name:    t.Name(),
			a:       a,
			b:       b,
			options: strat,
			diff: differs.PointersAddrMismatch(zero,
				sprints.Uintptrf(a.Pointer()), sprints.Uintptrf(b.Pointer())),
			equals: false,
		})
	})
}

// TestBetween_Slices_DefaultStrategy tests the method Between against slices
// based on the default Strategy
//
//	This test must assert:
//	1. slices are equals only when the conditions below are met:
//	a: their have same signatures
//	a: they have same length
//	a: they same equals elements or
//	b: both are nil with the same signature
//	2. Retrieve the standard messages on differences
func TestBetween_Slices_DefaultStrategy(t *testing.T) {
	var branch indent.Branch
	var strat []internal.Option

	t.Run("same instances", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		s := make([]int, 3)
		testDifferValues(tt, TestCase{
			name:    t.Name(),
			a:       reflect.ValueOf(s),
			b:       reflect.ValueOf(s),
			options: strat,
			diff:    differs.Empty(branch),
			equals:  true,
		})
	})

	t.Run("different instances, equals slices", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:    t.Name(),
			a:       reflect.ValueOf([]int{1, 2, 3}),
			b:       reflect.ValueOf([]int{1, 2, 3}),
			options: strat,
			diff:    differs.Empty(branch),
			equals:  true,
		})
	})

	t.Run("same signature, different elements", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:    t.Name(),
			a:       reflect.ValueOf([]int{1, 2, 3}),
			b:       reflect.ValueOf([]int{1, 2, 5}),
			options: strat,
			diff: differs.Append(indent.Zero(), differs.SliceValues(branch, 2),
				differs.Values(branch.Inc(), sprints.Typed("int", 3), sprints.Typed("int", 5))),
			equals: false,
		})
	})

	t.Run("same signatures, different lengths", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:   "different lengths",
			a:      reflect.ValueOf([]int{1, 2, 3}),
			b:      reflect.ValueOf([]int{1, 2}),
			diff:   differs.SliceLenMismatch(branch, 3, 2),
			equals: false,
		})
	})

	t.Run("slices with different structs", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name: "slices with structs",
			a:    reflect.ValueOf([]struct{ Name string }{{"Alice"}, {"Bob"}}),
			b:    reflect.ValueOf([]struct{ Name string }{{"Alice"}, {"Charlie"}}),
			diff: differs.Message(differs.Chain(differs.SliceValues(branch, 1),
				differs.StructFields(branch.Plus(1), "struct { Name string }", "Name"),
				differs.Strings(branch.Plus(2), 0)),
				differs.Diff(indent.Tab(2), "'"+fmx.SRedf("B")+"ob'", "'Charlie'")),
			equals: false,
		})
	})
}

// TestBetween_Slices_EdgeCases_DefaultStrategy tests the method Between against edge case slices
// based on the default Strategy
//
//	This test must assert:
//	1. slices are equals only when the conditions below are met:
//	a: their have same signatures
//	a: they have same length
//	a: they same equals elements or
//	b: both are nil with the same signature
//	2. Retrieve the standard messages on differences
func TestBetween_Slices_EdgeCases_DefaultStrategy(t *testing.T) {
	var branch indent.Branch
	var strat []internal.Option

	t.Run("both slices are nil", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:    "Both slices are nil",
			a:       reflect.ValueOf(([]int)(nil)),
			b:       reflect.ValueOf(([]int)(nil)),
			options: strat,
			diff:    differs.Empty(branch),
			equals:  true,
		})
	})

	t.Run("One slice is nil", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:    "One slice is nil, the other is not",
			a:       reflect.ValueOf(([]int)(nil)),
			b:       reflect.ValueOf([]int{1, 2, 3}),
			options: strat,
			diff:    differs.NilReceived(branch, "slice", "[]int"),
			equals:  false,
		})
	})

	t.Run("Empty slice vs nil slice", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:    "Empty slice vs nil slice",
			a:       reflect.ValueOf([]int{}),
			b:       reflect.ValueOf(([]int)(nil)),
			options: strat,
			diff:    differs.NilExpected(branch, "slice", "[]int"),
			equals:  false,
		})
	})

	t.Run("different instances, same amount of nil elements", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:    "slices with nil elements",
			a:       reflect.ValueOf([]*int{ptrs.Nil[int]()}),
			b:       reflect.ValueOf([]*int{ptrs.Nil[int]()}),
			options: strat,
			diff:    differs.Empty(branch),
			equals:  true,
		})
	})

	t.Run("slices with nested slices", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:    "Slices with nested slices",
			a:       reflect.ValueOf([][]int{{1, 2}, {3, 4}}),
			b:       reflect.ValueOf([][]int{{1, 2}, {3, 5}}),
			options: strat,
			diff: differs.Chain(differs.SliceValues(branch, 1), differs.SliceValues(branch.Plus(1), 1),
				differs.Values(branch.Plus(2), sprints.TypedDigit("int", 4),
					sprints.TypedDigit("int", 5))),
			equals: false,
		})
	})

	t.Run("Different int types as any", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:    "Different int types as any",
			a:       reflect.ValueOf([]any{1, 2, 3, 4, 5, 6}),
			b:       reflect.ValueOf([]any{int8(1), int8(2), int8(3), int8(4), int8(5), int8(6)}),
			options: strat,
			diff: differs.Chain(differs.SliceValues(branch, 0), differs.InterfaceImpl(branch.Plus(1)),
				differs.TypesMismatch(branch.Plus(2), "int", "int8")),
			equals: false,
		})
	})
}

// TestBetween_Slices_SerializableStrategy tests the method Between against slices
// based on the default Strategy
//
//	This test must assert:
//	1. slices are equals only when the conditions below are met:
//	a: their have same signatures
//	a: they have same length
//	a: they same equals elements or
//	b: both are nil or empty with the same signature
//	2. Retrieve the standard messages on differences
func TestBetween_Slices_SerializableStrategy(t *testing.T) {
	var branch indent.Branch
	strat := []internal.Option{compare.Serializable}

	t.Run("same instances", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		s := make([]int, 3)
		testDifferValues(tt, TestCase{
			name:    t.Name(),
			a:       reflect.ValueOf(s),
			b:       reflect.ValueOf(s),
			options: strat,
			diff:    differs.Empty(branch),
			equals:  true,
		})
	})

	t.Run("different instances, equals slices", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:    t.Name(),
			a:       reflect.ValueOf([]int{1, 2, 3}),
			b:       reflect.ValueOf([]int{1, 2, 3}),
			options: strat,
			diff:    differs.Empty(branch),
			equals:  true,
		})
	})

	t.Run("same signature, different elements", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:    t.Name(),
			a:       reflect.ValueOf([]int{1, 2, 3}),
			b:       reflect.ValueOf([]int{1, 2, 5}),
			options: strat,
			diff: differs.Append(indent.Zero(), differs.SliceValues(branch, 2),
				differs.Values(branch.Inc(), sprints.Typed("int", 3), sprints.Typed("int", 5))),
			equals: false,
		})
	})

	t.Run("same signatures, different lengths", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:    "Different lengths",
			a:       reflect.ValueOf([]int{1, 2, 3}),
			b:       reflect.ValueOf([]int{1, 2}),
			options: strat,
			diff:    differs.SliceLenMismatch(branch, 3, 2),
			equals:  false,
		})
	})

	t.Run("slices with different structs", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:    "Slices with structs",
			a:       reflect.ValueOf([]struct{ Name string }{{"Alice"}, {"Bob"}}),
			b:       reflect.ValueOf([]struct{ Name string }{{"Alice"}, {"Charlie"}}),
			options: strat,
			diff: differs.Message(differs.Chain(differs.SliceValues(branch, 1),
				differs.StructFields(branch.Plus(1), "struct { Name string }", "Name"), differs.Strings(branch.Plus(2), 0)),
				differs.Diff(indent.Tab(2), "'"+fmx.SRedf("B")+"ob'", "'Charlie'")),
			equals: false,
		})
	})
}

// TestBetween_Slices_EdgeCases_SerializableStrategy tests the method Between against edge case slices
// based on the serializable Strategy
//
//	This test must assert:
//	1. slices are equals only when the conditions below are met:
//	a: their have same signatures
//	a: they have same length
//	a: they same equals elements or
//	b: both are nil or empty with the same signature
//	2. Retrieve the standard messages on differences
func TestBetween_Slices_EdgeCases_SerializableStrategy(t *testing.T) {
	var branch indent.Branch
	strat := []internal.Option{compare.Serializable}

	t.Run("Both slices are nil", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:    "Both slices are nil",
			a:       reflect.ValueOf(([]int)(nil)),
			b:       reflect.ValueOf(([]int)(nil)),
			options: strat,
			diff:    differs.Empty(branch),
			equals:  true,
		})
	})

	t.Run("One slice is nil", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:    "One slice is nil, the other is not",
			a:       reflect.ValueOf(([]int)(nil)),
			b:       reflect.ValueOf([]int{1, 2, 3}),
			options: strat,
			diff:    differs.NilReceived(branch, "slice", "[]int"),
			equals:  false,
		})
	})

	t.Run("Empty slice vs nil slice", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:    "Empty slice vs nil slice",
			a:       reflect.ValueOf([]int{}),
			b:       reflect.ValueOf(([]int)(nil)),
			options: strat,
			diff:    differs.Empty(branch),
			equals:  true,
		})
	})

	t.Run("different instances, same amount of nil elements", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:    "slices with nil elements",
			a:       reflect.ValueOf([]*int{ptrs.Nil[int]()}),
			b:       reflect.ValueOf([]*int{ptrs.Nil[int]()}),
			options: strat,
			diff:    differs.Empty(branch),
			equals:  true,
		})
	})

	t.Run("slices with nested slices", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:    "Slices with nested slices",
			a:       reflect.ValueOf([][]int{{1, 2}, {3, 4}}),
			b:       reflect.ValueOf([][]int{{1, 2}, {3, 5}}),
			options: strat,
			diff: differs.Chain(differs.SliceValues(branch, 1), differs.SliceValues(branch.Plus(1), 1),
				differs.Values(branch.Plus(2), sprints.TypedDigit("int", 4),
					sprints.TypedDigit("int", 5))),
			equals: false,
		})
	})

	t.Run("Different int types as any", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:    "Different int types as any",
			a:       reflect.ValueOf([]any{1, 2, 3, 4, 5, 6}),
			b:       reflect.ValueOf([]any{int8(1), int8(2), int8(3), int8(4), int8(5), int8(6)}),
			options: strat,
			diff: differs.Chain(differs.SliceValues(branch, 0), differs.InterfaceImpl(branch.Plus(1)),
				differs.TypesMismatch(branch.Plus(2), "int", "int8")),
			equals: false,
		})
	})
}

// TestBetween_Strings tests the method Between against strings
//
//	This test must assert:
//	1. strings are equals only when the conditions below are met:
//	a: they have same length
//	a: they have the exact the sane characters in the same order
//	2. Retrieve the standard messages on differences
func TestBetween_Strings(t *testing.T) {
	var branch indent.Branch
	var strat []internal.Option

	t.Run("Equal strings", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:    "Equal strings",
			a:       reflect.ValueOf("hello"),
			b:       reflect.ValueOf("hello"),
			options: strat,
			diff:    differs.Empty(branch),
			equals:  true,
		})
	})

	t.Run("Both strings are empty", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:    "Both strings are empty",
			a:       reflect.ValueOf(""),
			b:       reflect.ValueOf(""),
			options: strat,
			diff:    differs.Empty(branch),
			equals:  true,
		})
	})

	// Different strings
	t.Run("Different lengths", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:    "Different lengths",
			a:       reflect.ValueOf("hello"),
			b:       reflect.ValueOf("hell"),
			options: strat,
			diff: differs.Message(differs.Strings(branch, 4),
				differs.Diff(indent.Tab(0), "'hell"+fmx.SRedf("o")+"'", "'hell'")),
			equals: false,
		})
	})

	t.Run("Same length, different characters", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:    "Same length, different characters",
			a:       reflect.ValueOf("hello"),
			b:       reflect.ValueOf("hallo"),
			options: strat,
			diff: differs.Message(differs.Strings(branch, 1), differs.Diff(indent.Tab(0), "'h"+
				fmx.SRedf("e")+"llo'", "'hallo'")),
			equals: false,
		})
	})
}

func TestBetween_Strings_EdgeCases(t *testing.T) {
	var branch indent.Branch
	var strat []internal.Option

	t.Run("Unicode characters", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:    "Unicode characters",
			a:       reflect.ValueOf("こんにちは"), // Japanese greeting
			b:       reflect.ValueOf("こんばんは"), // Japanese evening greeting
			options: strat,
			diff: differs.Message(differs.Strings(branch, 2), differs.Diff(indent.Tab(0), "'こん"+
				fmx.SRedf("に")+"ちは'", "'こんばんは'")), // "Character mismatch at index 2: 'に' vs 'ば'",
			equals: false,
		})
	})

	t.Run("Empty string vs non-empty string", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:    "Empty string vs non-empty string",
			a:       reflect.ValueOf(""),
			b:       reflect.ValueOf("non-empty"),
			options: strat,
			diff: differs.Message(differs.Strings(branch, 0),
				differs.Diff(indent.Tab(0), "''", "'non-empty'")),
			equals: false,
		})
	})

	t.Run("Case sensitivity", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:    "Case sensitivity",
			a:       reflect.ValueOf("Hello"),
			b:       reflect.ValueOf("hello"),
			options: strat,
			diff: differs.Message(differs.Strings(branch, 0),
				differs.Diff(indent.Tab(0), "'"+fmx.SRedf("H")+"ello'", "'hello'")),
			equals: false,
		})
	})
}

// TestBetween_Structs_DefaultStrategy tests the method Between against structs
// based on the default Strategy
//
//	This test must assert:
//	1. structs are equals only when the conditions below are met:
//	a: they have the same signature
//	a: all its fields are equals
//	2. Retrieve the standard messages on differences
func TestBetween_Structs_DefaultStrategy(t *testing.T) {
	var branch indent.Branch
	var strat []internal.Option

	t.Run("Equal structs", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:    "Equal structs",
			a:       reflect.ValueOf(struct{ Name string }{"Alice"}),
			b:       reflect.ValueOf(struct{ Name string }{"Alice"}),
			options: strat,
			diff:    differs.Empty(branch),
			equals:  true,
		})
	})

	t.Run("Different field values", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:    "Different field values",
			a:       reflect.ValueOf(struct{ Name string }{"Alice"}),
			b:       reflect.ValueOf(struct{ Name string }{"Bob"}),
			options: strat,
			diff: differs.Message(differs.Append(branch, differs.StructFields(branch, "struct { Name string }",
				"Name"), differs.Strings(branch.Plus(1), 0)),
				differs.Diff(indent.Tab(1), "'"+fmx.SRedf("A")+"lice'", "'Bob'")),
			equals: false,
		})
	})

	t.Run("Different field names", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:    "Different field names",
			a:       reflect.ValueOf(struct{ Name string }{"Alice"}),
			b:       reflect.ValueOf(struct{ FullName string }{"Alice"}),
			options: strat,
			diff: differs.StructTypesMismatch(branch, "struct { Name string }",
				"struct { FullName string }"),
			equals: false,
		})
	})

	t.Run("Different field types", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:    "Different field types",
			a:       reflect.ValueOf(struct{ Age int }{25}),
			b:       reflect.ValueOf(struct{ Age string }{"25"}),
			options: strat,
			diff: differs.StructTypesMismatch(branch, "struct { Age int }",
				"struct { Age string }"),
			equals: false,
		})
	})
}

// TestBetween_Structs_EdgeCases_DefaultStrategy tests the method Between against edge case structs
// based on the default Strategy
//
//	This test must assert:
//	1. structs are equals only when the conditions below are met:
//	a: they have the same signature
//	a: all its fields are equals
//	2. Retrieve the standard messages on differences
func TestBetween_Structs_EdgeCases_DefaultStrategy(t *testing.T) {
	var branch indent.Branch
	var strat []internal.Option

	// Edge cases
	t.Run("same signature, different field values", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:    "structs with nested structs",
			a:       reflect.ValueOf(struct{ Info struct{ Age int } }{Info: struct{ Age int }{25}}),
			b:       reflect.ValueOf(struct{ Info struct{ Age int } }{Info: struct{ Age int }{30}}),
			options: strat,
			diff: differs.Chain(differs.StructFields(branch, "struct { Info struct { Age int } }", "Info"),
				differs.StructFields(branch.Plus(1), "struct { Age int }", "Age"), differs.Values(branch.Plus(2),
					sprints.Typed("int", 25), sprints.Typed("int", 30))),
			equals: false,
		})
	})

	t.Run("same signature, slice fields with different values", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:    "structs with slices",
			a:       reflect.ValueOf(struct{ Numbers []int }{Numbers: []int{1, 2, 3}}),
			b:       reflect.ValueOf(struct{ Numbers []int }{Numbers: []int{1, 2, 4}}),
			options: strat,
			diff: differs.Chain(differs.StructFields(branch, "struct { Numbers []int }", "Numbers"),
				differs.SliceValues(branch.Plus(1), 2), differs.Values(branch.Plus(2),
					sprints.Typed("int", 3), sprints.Typed("int", 4))),
			equals: false,
		})
	})

	t.Run("same signature, unexported fields with different values", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:    "StructsPtr with unexported fields",
			a:       reflect.ValueOf(struct{ name string }{"Alice"}),
			b:       reflect.ValueOf(struct{ name string }{"Bob"}),
			options: strat,
			diff: differs.Message(differs.Chain(differs.StructFields(branch, "struct { name string }", "name"),
				differs.Strings(branch.Plus(1), 0)), differs.Diff(indent.Tab(1),
				"'"+fmx.SRedf("A")+"lice'", "'Bob'")),
			equals: false,
		})
	})

	t.Run("same signature, cyclic references with same field value", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:    "StructsPtr with cyclic references",
			a:       reflect.ValueOf(models.CyclicSimpleNode(42, 43)),
			b:       reflect.ValueOf(models.CyclicSimpleNode(42, 43)),
			options: strat,
			diff:    differs.Empty(branch),
			equals:  true,
		})
	})

	t.Run("complex equal structs", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		account := models.AccountAsRandValue()
		testDifferValues(tt, TestCase{
			name:    "complex structs",
			a:       reflect.ValueOf(account),
			b:       reflect.ValueOf(account),
			options: strat,
			diff:    differs.Empty(branch),
			equals:  true,
		})
	})
}

// TestBetween_Structs_SkipUnexpFields tests the method Between against structs
// based on struct extra related options.
//
//	This test must assert:
//	1. structs are equals only when the conditions below are met:
//	a: they have the same signature
//	a: all its exported fields are equals or
//	b: both fields are unaddressable
//	2. Retrieve the standard messages on differences
func TestBetween_Structs_SkipUnexpFields(t *testing.T) {
	// TODO:
}

// TestBetween_Structs_EdgeCases_SkipUnaddr_SkipUnexpFields tests the method Between against edge case structs
// based on struct extra related options.
//
//	This test must assert:
//	1. structs are equals only when the conditions below are met:
//	a: they have the same signature
//	a: all its fields are equals
//	2. Retrieve the standard messages on differences
func TestBetween_Structs_EdgeCases_SkipUnexpFields(t *testing.T) {
	// TODO:
}

// TestBetween_UnsafePointers tests the method Between against unsafe ptrs
//
//	This test must assert:
//	1. unsafe ptrs are equals only when the conditions below are met:
//	a: they have the same addresses; or
//	b: both are nil
//	2. Retrieve the standard messages on differences
func TestBetween_UnsafePointers(t *testing.T) {
	var zero indent.Branch
	var strat []internal.Option

	t.Run("same addresses", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:    t.Name(),
			a:       reflect.ValueOf(unsafe.Pointer(models.UnsafeCastableStringZeroInst)),
			b:       reflect.ValueOf(unsafe.Pointer(models.UnsafeCastableStringZeroInst)),
			options: strat,
			diff:    differs.Empty(zero),
			equals:  true,
		})
	})

	t.Run("same type, different instances", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		a := reflect.ValueOf(unsafe.Pointer(models.SimpleUnsafeCastableStringAsRef(1, "test")))
		b := reflect.ValueOf(unsafe.Pointer(models.SimpleUnsafeCastableStringAsRef(1, "test")))
		testDifferValues(tt, TestCase{
			name:    t.Name(),
			a:       a,
			b:       b,
			options: strat,
			diff: differs.UnsafePointersAddr(zero, sprints.Uintptrf(a.Pointer()),
				sprints.Uintptrf(b.Pointer())),
			equals: false,
		})
	})
}

// TestBetween_UnsafePointers_EdgeCases tests the method Between against edge case unsafe ptrs
//
//	This test must assert:
//	1. unsafe ptrs are equals only when the conditions below are met:
//	a: they have the same addresses; or
//	b: both are nil
//	2. Retrieve the standard messages on differences
func TestBetween_UnsafePointers_EdgeCases(t *testing.T) {
	var zero indent.Branch
	var strat []internal.Option

	t.Run("", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name:    "both unsafe.PointerValues are nil",
			a:       reflect.ValueOf(unsafe.Pointer(nil)),
			b:       reflect.ValueOf(unsafe.Pointer(nil)),
			options: strat,
			diff:    differs.Empty(zero),
			equals:  true,
		})
	})

	t.Run("one is nil, other is not", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		nonNull := reflect.ValueOf(unsafe.Pointer(models.SimpleUnsafeCastableStringAsRef(1, "test")))
		null := reflect.ValueOf(unsafe.Pointer(nil))
		testDifferValues(tt, TestCase{
			name:    "nil received, valid expected",
			a:       null,
			b:       nonNull,
			options: strat,
			diff:    differs.NilReceived(zero, "unsafe.pointer", sprints.Uintptrf(nonNull.Pointer())),
			equals:  false,
		})
		testDifferValues(tt, TestCase{
			name:    "valid received, nil expected",
			a:       nonNull,
			b:       null,
			options: strat,
			diff:    differs.NilExpected(zero, "unsafe.pointer", sprints.Uintptrf(nonNull.Pointer())),
			equals:  false,
		})
	})

	// This is a case where two unsafe ptrs are nil
	// and the api sees them as equals
	t.Run("uninitialized ptrs, different types", func(t *testing.T) {
		tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
		testDifferValues(tt, TestCase{
			name: "uninitialized ptrs",
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
		})
	})
}
