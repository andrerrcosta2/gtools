// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package equals

import (
	"github.com/andrerrcosta2/gtools/core/testlite/assertlite"
	"github.com/andrerrcosta2/gtools/core/util/typeutil/generics"
	"github.com/andrerrcosta2/gtools/reflect4/internal"
	"github.com/andrerrcosta2/gtools/reflect4/op/compare"
	"math"
	"reflect"
	"testing"
)

type (
	MyByte byte

	MyBytes []byte

	Basic struct {
		x int
		y float32
	}

	_Complex struct {
		a int
		b [3]*_Complex
		c *string
		d map[float64]float64
	}

	NotBasic Basic

	self struct{}

	Loop  *Loop
	Loopy any

	Recursive struct {
		x int
		r *Recursive
	}

	structWithSelfPtr struct {
		p *structWithSelfPtr
		s string
	}

	UnexpT struct {
		m map[int]int
	}
)

type DeepEqualTest struct {
	a, b any
	eq   bool
}

// Simple functions for DeepEqual tests.
var (
	fnil1        func()                    // nil.
	fnil2        func()                    // nil.
	intSupplierA = func() int { return 0 } // Not nil.
	intSupplierB = func() int { return 1 } // Not nil
)

var loop1, loop2 Loop
var loopy1, loopy2 Loopy
var cycleMap1, cycleMap2, cycleMap3 map[string]any

func init() {
	loop1 = &loop2
	loop2 = &loop1

	loopy1 = &loopy2
	loopy2 = &loopy1

	cycleMap1 = map[string]any{}
	cycleMap1["cycle"] = cycleMap1
	cycleMap2 = map[string]any{}
	cycleMap2["cycle"] = cycleMap2
	cycleMap3 = map[string]any{}
	cycleMap3["different"] = cycleMap3
}

var deepTests_refl = []DeepEqualTest{
	// Equalities
	{nil, nil, true},
	{1, 1, true},
	{int32(1), int32(1), true},
	{0.5, 0.5, true},
	{float32(0.5), float32(0.5), true},
	{"hello", "hello", true},
	{make(chan int), self{}, true},
	{make([]int, 10), make([]int, 10), true},
	{&[3]int{1, 2, 3}, &[3]int{1, 2, 3}, true},
	{Basic{1, 0.5}, Basic{1, 0.5}, true},
	{error(nil), error(nil), true},
	{map[int]string{1: "one", 2: "two"}, map[int]string{2: "two", 1: "one"}, true},
	{fnil1, fnil2, true},
	{[]byte{1, 2, 3}, []byte{1, 2, 3}, true},
	{[]MyByte{1, 2, 3}, []MyByte{1, 2, 3}, true},
	{MyBytes{1, 2, 3}, MyBytes{1, 2, 3}, true},

	// Inequalities
	{1, 2, false},
	{int32(1), int32(2), false},
	{0.5, 0.6, false},
	{float32(0.5), float32(0.6), false},
	{"hello", "hey", false},
	{make([]int, 10), make([]int, 11), false},
	{make([]int, 3), generics.Zero[[]int](), false},    // empty vs nil
	{make(chan int), make(chan int), false},            // reflect semantics
	{generics.Zero[chan int](), make(chan int), false}, // empty vs nil
	{&[3]int{1, 2, 3}, &[3]int{1, 2, 4}, false},
	{Basic{1, 0.5}, Basic{1, 0.6}, false},
	{Basic{1, 0}, Basic{2, 0}, false},
	{map[int]string{}, generics.Zero[map[int]string](), false}, // empty vs nil
	{map[int]string{1: "one", 3: "two"}, map[int]string{2: "two", 1: "one"}, false},
	{map[int]string{1: "one", 2: "txo"}, map[int]string{2: "two", 1: "one"}, false},
	{map[int]string{1: "one"}, map[int]string{2: "two", 1: "one"}, false},
	{map[int]string{2: "two", 1: "one"}, map[int]string{1: "one"}, false},
	{nil, 1, false},
	{1, nil, false},
	{fnil1, intSupplierA, false},
	{intSupplierA, intSupplierB, false},
	{[][]int{{1}}, [][]int{{2}}, false},
	{&structWithSelfPtr{p: &structWithSelfPtr{s: "a"}}, &structWithSelfPtr{p: &structWithSelfPtr{s: "b"}}, false},

	// Fun with floating point.
	{math.NaN(), math.NaN(), false},
	{&[1]float64{math.NaN()}, &[1]float64{math.NaN()}, false},
	{&[1]float64{math.NaN()}, self{}, true},
	{[]float64{math.NaN()}, []float64{math.NaN()}, false},
	{[]float64{math.NaN()}, self{}, true},
	{map[float64]float64{math.NaN(): 1}, map[float64]float64{1: 2}, false},
	{map[float64]float64{math.NaN(): 1}, self{}, true},

	// Nil vs empty: not the same.
	{[]int{}, []int(nil), false},
	{[]int{}, []int{}, true},
	{[]int(nil), []int(nil), true},
	{map[int]int{}, map[int]int(nil), false},
	{map[int]int{}, map[int]int{}, true},
	{map[int]int(nil), map[int]int(nil), true},

	// Mismatched types
	{1, 1.0, false},
	{int32(1), int64(1), false},
	{0.5, "hello", false},
	{[]int{1, 2, 3}, [3]int{1, 2, 3}, false},
	{&[3]any{1, 2, 4}, &[3]any{1, 2, "s"}, false},
	{Basic{1, 0.5}, NotBasic{1, 0.5}, false},
	{map[uint]string{1: "one", 2: "two"}, map[int]string{2: "two", 1: "one"}, false},
	{[]byte{1, 2, 3}, []MyByte{1, 2, 3}, false},
	{[]MyByte{1, 2, 3}, MyBytes{1, 2, 3}, false},
	{[]byte{1, 2, 3}, MyBytes{1, 2, 3}, false},

	// Possible loops.
	{&loop1, &loop1, true},
	{&loop1, &loop2, true},
	{&loopy1, &loopy1, true},
	{&loopy1, &loopy2, true},
	{&cycleMap1, &cycleMap2, true},
	{&cycleMap1, &cycleMap3, false},
}

// TestDeepEqual_ReflStrat tests the method Deep based on the native reflect package
// perspective of equality:
//
// Array values are deeply equal when their corresponding elements are deeply equal.
//
// Struct values are deeply equal if their corresponding fields,
// both exported and unexported, are deeply equal.
//
// Func values are deeply equal if both are nil; otherwise they are not deeply equal.
//
// Interface values are deeply equal if they hold deeply equal concrete values.
//
// Map values are deeply equal when all the following are true:
// they are both nil or both non-nil, they have the same length,
// and either they are the same map object or their corresponding keys
// (matched using Go equality) map to deeply equal values.
//
// Pointer values are deeply equal if they are equal using Go's == operator
// or if they point to deeply equal values.
//
// Slice values are deeply equal when all the following are true:
// they are both nil or both non-nil, they have the same length,
// and either they point to the same initial entry of the same underlying array
// (that is, &x[0] == &y[0]) or their corresponding elements (up to length) are deeply equal.
// Note that a non-nil empty slice and a nil slice (for example, []byte{} and []byte(nil))
// are not deeply equal.
//
// Other values - numbers, bools, strings, and channels - are deeply equal
// if they are equal using Go's == operator.
func TestDeepEqual_ReflStrat(t *testing.T) {
	t.Run("general tests", func(t *testing.T) {
		for _, test := range deepTests_refl {
			if test.a == nil || test.b == nil {
				assertlite.True(t, (test.a == test.b) == test.eq, "%v != %v", test.a, test.b)
				continue
			}
			if test.b == (self{}) {
				test.b = test.a
			}
			v1, v2 := reflect.ValueOf(test.a), reflect.ValueOf(test.b)
			eq := Deep(v1, v2, compare.ReflectSemantics)
			assertlite.True(t, eq == test.eq,
				"DeepEqual(%#v, %#v) = %v, want %v", test.a, test.b, eq, test.eq)
		}
	})

	t.Run("complex struct", func(t *testing.T) {
		// equals
		m := make(map[float64]float64)
		stra, strb, strc := "hello", "hello", "helloo"
		a, b := new(_Complex), new(_Complex)
		*a = _Complex{5, [3]*_Complex{a, b, a}, &stra, m}
		*b = _Complex{5, [3]*_Complex{b, a, a}, &strb, m}
		v1, v2 := reflect.ValueOf(a), reflect.ValueOf(b)
		assertlite.True(t, Deep(v1, v2, compare.ReflectSemantics), "Deep(complex same) = false, want true")

		// not equals
		m = make(map[float64]float64)

		a, b = new(_Complex), new(_Complex)
		*a = _Complex{5, [3]*_Complex{a, b, a}, &stra, m}
		*b = _Complex{5, [3]*_Complex{b, a, a}, &strc, m}
		v1, v2 = reflect.ValueOf(a), reflect.ValueOf(b)
		assertlite.False(t, Deep(v1, v2, compare.ReflectSemantics), "Deep(complex different) = true, want false")
	})

	t.Run("unexported map", func(t *testing.T) {
		// Check that DeepEqual can look at unexported fields.
		x1 := UnexpT{map[int]int{1: 2}}
		x2 := UnexpT{map[int]int{1: 2}}
		v1, v2 := reflect.ValueOf(&x1), reflect.ValueOf(&x2)
		assertlite.True(t, Deep(v1, v2, compare.ReflectSemantics), "DeepEqual(x1, x2) = false, want true")

		y1 := UnexpT{map[int]int{2: 3}}
		v3 := reflect.ValueOf(&y1)
		assertlite.False(t, Deep(v1, v3, compare.ReflectSemantics), "DeepEqual(x1, y1) = true, want false")
	})
}

var deepTests_default = []DeepEqualTest{
	// Equalities
	{nil, nil, true},
	{1, 1, true},
	{int32(1), int32(1), true},
	{0.5, 0.5, true},
	{float32(0.5), float32(0.5), true},
	{"hello", "hello", true},
	{make(chan int), self{}, true},
	{make(chan int), make(chan int), true}, // defaultChan
	{make([]int, 10), make([]int, 10), true},
	{&[3]int{1, 2, 3}, &[3]int{1, 2, 3}, true},
	{Basic{1, 0.5}, Basic{1, 0.5}, true},
	{error(nil), error(nil), true},
	{map[int]string{1: "one", 2: "two"}, map[int]string{2: "two", 1: "one"}, true},
	{fnil1, fnil2, true},
	{intSupplierA, intSupplierB, true},
	{[]byte{1, 2, 3}, []byte{1, 2, 3}, true},
	{[]MyByte{1, 2, 3}, []MyByte{1, 2, 3}, true},
	{MyBytes{1, 2, 3}, MyBytes{1, 2, 3}, true},

	// Inequalities
	{1, 2, false},
	{int32(1), int32(2), false},
	{0.5, 0.6, false},
	{float32(0.5), float32(0.6), false},
	{"hello", "hey", false},
	{make(chan int), make(chan int8), false},           // defaultChan
	{generics.Zero[chan int](), make(chan int), false}, // empty vs nil
	{make([]int, 10), make([]int, 11), false},
	{[]int{}, generics.Zero[[]int](), false}, // empty vs nil
	{&[3]int{1, 2, 3}, &[3]int{1, 2, 4}, false},
	{Basic{1, 0.5}, Basic{1, 0.6}, false},
	{Basic{1, 0}, Basic{2, 0}, false},
	{map[int]string{}, generics.Zero[map[int]string](), false}, // empty vs nil
	{map[int]string{1: "one", 3: "two"}, map[int]string{2: "two", 1: "one"}, false},
	{map[int]string{1: "one", 2: "txo"}, map[int]string{2: "two", 1: "one"}, false},
	{map[int]string{1: "one"}, map[int]string{2: "two", 1: "one"}, false},
	{map[int]string{2: "two", 1: "one"}, map[int]string{1: "one"}, false},
	{nil, 1, false},
	{1, nil, false},
	{fnil1, intSupplierA, false},
	{[][]int{{1}}, [][]int{{2}}, false},
	{&structWithSelfPtr{p: &structWithSelfPtr{s: "a"}}, &structWithSelfPtr{p: &structWithSelfPtr{s: "b"}}, false},

	// Fun with floating point.
	{math.NaN(), math.NaN(), false},
	{&[1]float64{math.NaN()}, &[1]float64{math.NaN()}, false},
	{&[1]float64{math.NaN()}, self{}, true},
	{[]float64{math.NaN()}, []float64{math.NaN()}, false},
	{[]float64{math.NaN()}, self{}, true},
	{map[float64]float64{math.NaN(): 1}, map[float64]float64{1: 2}, false},
	{map[float64]float64{math.NaN(): 1}, self{}, true},

	// Nil vs empty: not the same.
	{[]int{}, []int(nil), false},
	{[]int{}, []int{}, true},
	{[]int(nil), []int(nil), true},
	{map[int]int{}, map[int]int(nil), false},
	{map[int]int{}, map[int]int{}, true},
	{map[int]int(nil), map[int]int(nil), true},

	// Mismatched types
	{1, 1.0, false},
	{int32(1), int64(1), false},
	{0.5, "hello", false},
	{[]int{1, 2, 3}, [3]int{1, 2, 3}, false},
	{&[3]any{1, 2, 4}, &[3]any{1, 2, "s"}, false},
	{Basic{1, 0.5}, NotBasic{1, 0.5}, false},
	{map[uint]string{1: "one", 2: "two"}, map[int]string{2: "two", 1: "one"}, false},
	{[]byte{1, 2, 3}, []MyByte{1, 2, 3}, false},
	{[]MyByte{1, 2, 3}, MyBytes{1, 2, 3}, false},
	{[]byte{1, 2, 3}, MyBytes{1, 2, 3}, false},

	// Possible loops.
	{&loop1, &loop1, true},
	{&loop1, &loop2, true},
	{&loopy1, &loopy1, true},
	{&loopy1, &loopy2, true},
	{&cycleMap1, &cycleMap2, true},
	{&cycleMap1, &cycleMap3, false},
}

func TestDeepEqual_Default(t *testing.T) {
	t.Run("general tests", func(t *testing.T) {
		for _, test := range deepTests_default {
			if test.a == nil || test.b == nil {
				assertlite.True(t, (test.a == test.b) == test.eq, "%v != %v", test.a, test.b)
				continue
			}
			if test.b == (self{}) {
				test.b = test.a
			}
			v1, v2 := reflect.ValueOf(test.a), reflect.ValueOf(test.b)
			eq := Deep[internal.Option](v1, v2)
			assertlite.True(t, eq == test.eq,
				"DeepEqual(%#v, %#v) = %v, want %v", test.a, test.b, eq, test.eq)
		}
	})

	t.Run("complex struct", func(t *testing.T) {
		// equals
		m := make(map[float64]float64)
		stra, strb, strc := "hello", "hello", "helloo"
		a, b := new(_Complex), new(_Complex)
		*a = _Complex{5, [3]*_Complex{a, b, a}, &stra, m}
		*b = _Complex{5, [3]*_Complex{b, a, a}, &strb, m}
		v1, v2 := reflect.ValueOf(a), reflect.ValueOf(b)
		assertlite.True(t, Deep[internal.Option](v1, v2), "Deep(complex same) = false, want true")

		// not equals
		m = make(map[float64]float64)

		a, b = new(_Complex), new(_Complex)
		*a = _Complex{5, [3]*_Complex{a, b, a}, &stra, m}
		*b = _Complex{5, [3]*_Complex{b, a, a}, &strc, m}
		v1, v2 = reflect.ValueOf(a), reflect.ValueOf(b)
		assertlite.False(t, Deep[internal.Option](v1, v2), "Deep(complex different) = true, want false")
	})

	t.Run("unexported map", func(t *testing.T) {
		// Check that DeepEqual can look at unexported fields.
		x1 := UnexpT{map[int]int{1: 2}}
		x2 := UnexpT{map[int]int{1: 2}}
		v1, v2 := reflect.ValueOf(&x1), reflect.ValueOf(&x2)
		assertlite.True(t, Deep[internal.Option](v1, v2), "DeepEqual(x1, x2) = false, want true")

		y1 := UnexpT{map[int]int{2: 3}}
		v3 := reflect.ValueOf(&y1)
		assertlite.False(t, Deep[internal.Option](v1, v3), "DeepEqual(x1, y1) = true, want false")
	})
}

var deepTests_serializable = []DeepEqualTest{
	// Equalities
	{nil, nil, true},
	{1, 1, true},
	{int32(1), int32(1), true},
	{0.5, 0.5, true},
	{float32(0.5), float32(0.5), true},
	{"hello", "hello", true},
	{make(chan int), self{}, true},
	{make(chan int), make(chan int), true},            // serialChan
	{generics.Zero[chan int](), make(chan int), true}, // empty vs nil
	{make([]int, 10), make([]int, 10), true},
	{[]int{}, generics.Zero[[]int](), true}, // empty vs nil
	{&[3]int{1, 2, 3}, &[3]int{1, 2, 3}, true},
	{Basic{1, 0.5}, Basic{1, 0.5}, true},
	{error(nil), error(nil), true},
	{map[int]string{1: "one", 2: "two"}, map[int]string{2: "two", 1: "one"}, true},
	{map[int]string{}, generics.Zero[map[int]string](), true}, // empty vs nil
	{fnil1, fnil2, true},
	{intSupplierA, intSupplierB, true},
	{[]byte{1, 2, 3}, []byte{1, 2, 3}, true},
	{[]MyByte{1, 2, 3}, []MyByte{1, 2, 3}, true},
	{MyBytes{1, 2, 3}, MyBytes{1, 2, 3}, true},

	// Inequalities
	{1, 2, false},
	{int32(1), int32(2), false},
	{0.5, 0.6, false},
	{float32(0.5), float32(0.6), false},
	{"hello", "hey", false},
	{make(chan int), make(chan int8), false}, // serialChan
	{make([]int, 10), make([]int, 11), false},
	{&[3]int{1, 2, 3}, &[3]int{1, 2, 4}, false},
	{Basic{1, 0.5}, Basic{1, 0.6}, false},
	{Basic{1, 0}, Basic{2, 0}, false},
	{map[int]string{1: "one", 3: "two"}, map[int]string{2: "two", 1: "one"}, false},
	{map[int]string{1: "one", 2: "txo"}, map[int]string{2: "two", 1: "one"}, false},
	{map[int]string{1: "one"}, map[int]string{2: "two", 1: "one"}, false},
	{map[int]string{2: "two", 1: "one"}, map[int]string{1: "one"}, false},
	{nil, 1, false},
	{1, nil, false},
	{fnil1, intSupplierA, false},
	{[][]int{{1}}, [][]int{{2}}, false},
	{&structWithSelfPtr{p: &structWithSelfPtr{s: "a"}}, &structWithSelfPtr{p: &structWithSelfPtr{s: "b"}}, false},

	// Fun with floating point.
	{math.NaN(), math.NaN(), false},
	{&[1]float64{math.NaN()}, &[1]float64{math.NaN()}, false},
	{&[1]float64{math.NaN()}, self{}, true},
	{[]float64{math.NaN()}, []float64{math.NaN()}, false},
	{[]float64{math.NaN()}, self{}, true},
	{map[float64]float64{math.NaN(): 1}, map[float64]float64{1: 2}, false},
	{map[float64]float64{math.NaN(): 1}, self{}, true},

	// Nil vs empty: serial strategy
	{[]int{}, []int(nil), true}, // serializable equals
	{[]int{}, []int{}, true},
	{[]int(nil), []int(nil), true},
	{map[int]int{}, map[int]int(nil), true}, // serializable equals
	{map[int]int{}, map[int]int{}, true},
	{map[int]int(nil), map[int]int(nil), true},

	// Mismatched types
	{1, 1.0, false},
	{int32(1), int64(1), false},
	{0.5, "hello", false},
	{[]int{1, 2, 3}, [3]int{1, 2, 3}, false},
	{&[3]any{1, 2, 4}, &[3]any{1, 2, "s"}, false},
	{Basic{1, 0.5}, NotBasic{1, 0.5}, false},
	{map[uint]string{1: "one", 2: "two"}, map[int]string{2: "two", 1: "one"}, false},
	{[]byte{1, 2, 3}, []MyByte{1, 2, 3}, false},
	{[]MyByte{1, 2, 3}, MyBytes{1, 2, 3}, false},
	{[]byte{1, 2, 3}, MyBytes{1, 2, 3}, false},

	// Possible loops.
	{&loop1, &loop1, true},
	{&loop1, &loop2, true},
	{&loopy1, &loopy1, true},
	{&loopy1, &loopy2, true},
	{&cycleMap1, &cycleMap2, true},
	{&cycleMap1, &cycleMap3, false},
}

func TestDeepEqual_Serializable(t *testing.T) {
	t.Run("general tests", func(t *testing.T) {
		for _, test := range deepTests_serializable {
			if test.a == nil || test.b == nil {
				assertlite.True(t, (test.a == test.b) == test.eq, "%v != %v", test.a, test.b)
				continue
			}
			if test.b == (self{}) {
				test.b = test.a
			}
			v1, v2 := reflect.ValueOf(test.a), reflect.ValueOf(test.b)
			eq := Deep(v1, v2, compare.Serializable)
			assertlite.True(t, eq == test.eq,
				"DeepEqual(%#v, %#v) = %v, want %v", test.a, test.b, eq, test.eq)
		}
	})

	t.Run("complex struct", func(t *testing.T) {
		// equals
		m := make(map[float64]float64)
		stra, strb, strc := "hello", "hello", "helloo"
		a, b := new(_Complex), new(_Complex)
		*a = _Complex{5, [3]*_Complex{a, b, a}, &stra, m}
		*b = _Complex{5, [3]*_Complex{b, a, a}, &strb, m}
		v1, v2 := reflect.ValueOf(a), reflect.ValueOf(b)
		assertlite.True(t, Deep(v1, v2, compare.Serializable), "Deep(complex same) = false, want true")

		// not equals
		m = make(map[float64]float64)

		a, b = new(_Complex), new(_Complex)
		*a = _Complex{5, [3]*_Complex{a, b, a}, &stra, m}
		*b = _Complex{5, [3]*_Complex{b, a, a}, &strc, m}
		v1, v2 = reflect.ValueOf(a), reflect.ValueOf(b)
		assertlite.False(t, Deep(v1, v2, compare.Serializable), "Deep(complex different) = true, want false")
	})

	t.Run("unexported map", func(t *testing.T) {
		// Check that DeepEqual can look at unexported fields.
		x1 := UnexpT{map[int]int{1: 2}}
		x2 := UnexpT{map[int]int{1: 2}}
		v1, v2 := reflect.ValueOf(&x1), reflect.ValueOf(&x2)
		assertlite.True(t, Deep(v1, v2, compare.Serializable), "DeepEqual(x1, x2) = false, want true")

		y1 := UnexpT{map[int]int{2: 3}}
		v3 := reflect.ValueOf(&y1)
		assertlite.False(t, Deep(v1, v3, compare.Serializable), "DeepEqual(x1, y1) = true, want false")
	})
}

var deepTests_strict = []DeepEqualTest{
	// Equalities
	{nil, nil, true},
	{1, 1, true},
	{int32(1), int32(1), true},
	{0.5, 0.5, true},
	{float32(0.5), float32(0.5), true},
	{"hello", "hello", true},
	{make(chan int), self{}, true},
	{generics.Zero[chan int](), generics.Zero[chan int](), true}, // identityShallow (nil vs nil)
	{&[3]int{1, 2, 3}, self{}, true},
	{Basic{1, 0.5}, Basic{1, 0.5}, true},
	{error(nil), error(nil), true},
	{map[int]string{1: "one", 2: "two"}, self{}, true},
	{generics.Zero[map[int]any](), generics.Zero[map[int]any](), true}, // identityShallow (nil vs nil)
	{fnil1, fnil2, true},                                               // identityShallow (nil vs nil)
	{intSupplierA, self{}, true},

	// Inequalities
	{1, 2, false},
	{int32(1), int32(2), false},
	{0.5, 0.6, false},
	{float32(0.5), float32(0.6), false},
	{"hello", "hey", false},
	{make(chan int), make(chan int), false},            // identityShallow
	{make(chan int), make(chan int8), false},           // identityShallow
	{make([]int, 10), make([]int, 10), false},          // identityDeep
	{generics.Zero[chan int](), make(chan int), false}, // empty vs nil
	{make([]int, 10), make([]int, 11), false},
	{[]int{}, generics.Zero[[]int](), false},      // empty vs nil
	{[]byte{1, 2, 3}, []byte{1, 2, 3}, false},     // identityDeep
	{[]MyByte{1, 2, 3}, []MyByte{1, 2, 3}, false}, // identityDeep
	{MyBytes{1, 2, 3}, MyBytes{1, 2, 3}, false},   // identityDeep
	{&[3]int{1, 2, 3}, &[3]int{1, 2, 3}, false},   // identityDeep
	{Basic{1, 0.5}, Basic{1, 0.6}, false},
	{Basic{1, 0}, Basic{2, 0}, false},
	{map[int]string{1: "one", 2: "two"}, map[int]string{2: "two", 1: "one"}, false}, // identityDeep
	{map[int]string{1: "one", 3: "two"}, map[int]string{2: "two", 1: "one"}, false},
	{map[int]string{1: "one", 2: "txo"}, map[int]string{2: "two", 1: "one"}, false},
	{map[int]string{1: "one"}, map[int]string{2: "two", 1: "one"}, false},
	{map[int]string{2: "two", 1: "one"}, map[int]string{1: "one"}, false},
	{map[int]string{}, generics.Zero[map[int]string](), false}, // empty vs nil
	{nil, 1, false},
	{1, nil, false},
	{intSupplierA, intSupplierB, false}, // identityShallow
	{fnil1, intSupplierA, false},
	{[][]int{{1}}, [][]int{{2}}, false},
	{&structWithSelfPtr{p: &structWithSelfPtr{s: "a"}}, &structWithSelfPtr{p: &structWithSelfPtr{s: "b"}}, false},

	// Fun with floating point.
	{math.NaN(), math.NaN(), false},
	{&[1]float64{math.NaN()}, &[1]float64{math.NaN()}, false},
	{&[1]float64{math.NaN()}, self{}, true},
	{[]float64{math.NaN()}, []float64{math.NaN()}, false},
	{[]float64{math.NaN()}, self{}, true},
	{map[float64]float64{math.NaN(): 1}, map[float64]float64{1: 2}, false},
	{map[float64]float64{math.NaN(): 1}, self{}, true},

	// Nil vs empty: not the same.
	{[]int{}, []int(nil), false},
	{[]int{}, []int{}, true},       // identitySlice
	{[]int(nil), []int(nil), true}, // identitySlice (nil vs nil)
	{map[int]int{}, map[int]int(nil), false},
	{map[int]int{}, map[int]int{}, false},      // identityDeep
	{map[int]int(nil), map[int]int(nil), true}, // identityDeep (nul vs nil)

	// Mismatched types
	{1, 1.0, false},
	{int32(1), int64(1), false},
	{0.5, "hello", false},
	{[]int{1, 2, 3}, [3]int{1, 2, 3}, false},
	{&[3]any{1, 2, 4}, &[3]any{1, 2, "s"}, false},
	{Basic{1, 0.5}, NotBasic{1, 0.5}, false},
	{map[uint]string{1: "one", 2: "two"}, map[int]string{2: "two", 1: "one"}, false},
	{[]byte{1, 2, 3}, []MyByte{1, 2, 3}, false},
	{[]MyByte{1, 2, 3}, MyBytes{1, 2, 3}, false},
	{[]byte{1, 2, 3}, MyBytes{1, 2, 3}, false},

	// Possible loops.
	{&loop1, &loop1, true},
	{&loop1, &loop2, false},
	{&loopy1, &loopy1, true},
	{&loopy1, &loopy2, false},
	{&cycleMap1, &cycleMap2, false},
	{&cycleMap1, &cycleMap3, false},
}

func TestDeepEqual_Strict(t *testing.T) {
	t.Run("general tests", func(t *testing.T) {
		for _, test := range deepTests_strict {
			if test.a == nil || test.b == nil {
				assertlite.True(t, (test.a == test.b) == test.eq, "%v != %v", test.a, test.b)
				continue
			}
			if test.b == (self{}) {
				test.b = test.a
			}
			v1, v2 := reflect.ValueOf(test.a), reflect.ValueOf(test.b)
			eq := Deep(v1, v2, compare.Strict)
			assertlite.True(t, eq == test.eq,
				"DeepEqual(%#v, %#v) = %v, want %v", test.a, test.b, eq, test.eq)
		}
	})

	t.Run("complex struct", func(t *testing.T) {
		// equals (different addresses)
		m := make(map[float64]float64)
		stra, strb, strc := "hello", "hello", "helloo"
		a, b := new(_Complex), new(_Complex)
		*a = _Complex{5, [3]*_Complex{a, b, a}, &stra, m}
		*b = _Complex{5, [3]*_Complex{b, a, a}, &strb, m}
		v1, v2 := reflect.ValueOf(a), reflect.ValueOf(b)
		assertlite.False(t, Deep(v1, v2, compare.Strict), "Deep(complex same) = false, want true")

		// not equals
		m = make(map[float64]float64)

		a, b = new(_Complex), new(_Complex)
		*a = _Complex{5, [3]*_Complex{a, b, a}, &stra, m}
		*b = _Complex{5, [3]*_Complex{b, a, a}, &strc, m}
		v1, v2 = reflect.ValueOf(a), reflect.ValueOf(b)
		assertlite.False(t, Deep(v1, v2, compare.Strict), "Deep(complex different) = true, want false")
	})

	t.Run("unexported map", func(t *testing.T) {
		// Check that DeepEqual can look at unexported fields.
		x1 := UnexpT{map[int]int{1: 2}}
		x2 := UnexpT{map[int]int{1: 2}}
		v1, v2 := reflect.ValueOf(&x1), reflect.ValueOf(&x2)
		assertlite.False(t, Deep(v1, v2, compare.Strict), "DeepEqual(x1, x2) = false, want true")

		y1 := UnexpT{map[int]int{2: 3}}
		v3 := reflect.ValueOf(&y1)
		assertlite.False(t, Deep(v1, v3, compare.Strict), "DeepEqual(x1, y1) = true, want false")
	})
}
