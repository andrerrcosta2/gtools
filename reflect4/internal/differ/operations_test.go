// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package differ

import (
	"reflect"
	"strings"
	"testing"
	"unsafe"

	"github.com/andrerrcosta2/gtools/core/domain/functions"
	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"github.com/andrerrcosta2/gtools/core/format/differs"
	"github.com/andrerrcosta2/gtools/core/format/fmx"
	"github.com/andrerrcosta2/gtools/core/format/sprints"
	"github.com/andrerrcosta2/gtools/core/util/typeutil/generics"
	"github.com/andrerrcosta2/gtools/core/util/typeutil/ptrs"
	"github.com/andrerrcosta2/gtools/core/util/typeutil/slices"
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static/interf"
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static/structs/models"
	"github.com/andrerrcosta2/gtools/reflect4/internal/standards"
	"github.com/andrerrcosta2/gtools/reflect4/internal/types"
	"github.com/andrerrcosta2/gtools/reflect4/op/compare"
)

// TestDefaultChanDiff tests the method defaultChanDiff
//
//	This test must assert:
//	1. channels are equals only when:
//	- They point to the same underlying address; or
//	- Both have the same signature, direction and capacity; or
//	- Both channels are nil with same signature
func TestDefaultChanDiff(t *testing.T) {
	var zero indent.Branch
	strat := DefaultStrat()

	t.Run("same instances", func(t *testing.T) {
		c := make(chan int, 100)
		diff := defaultChanDiff(zero, reflect.ValueOf(c), reflect.ValueOf(c), strat)
		assertEquals(t, diff)
	})

	t.Run("different instances, equals signature, equals buffer size", func(t *testing.T) {
		a, b := make(chan bool, 100), make(chan bool, 100)
		diff := defaultChanDiff(zero, reflect.ValueOf(a), reflect.ValueOf(b), strat)
		assertEquals(t, diff)
	})

	t.Run("different signatures", func(t *testing.T) {
		a, b := reflect.ValueOf(make(chan bool)), reflect.ValueOf(make(chan int))
		expMsg := differs.ChanElemTypesMismatch(zero,
			types.ValidValueName(a.Type().Elem()), types.ValidValueName(b.Type().Elem()))
		diff := defaultChanDiff(zero, a, b, strat)
		assertNotEquals(t, diff, expMsg, "")
	})

	t.Run("equals signature, different buffer size", func(t *testing.T) {
		a, b := make(chan bool, 100), make(chan bool, 1000)
		expMsg := differs.ChanBufSizeMismatch(zero, 100, 1000)
		diff := defaultChanDiff(zero, reflect.ValueOf(a), reflect.ValueOf(b), strat)
		assertNotEquals(t, diff, expMsg, "")
	})

	t.Run("equals signature, different dir", func(t *testing.T) {
		a, b := reflect.ValueOf(make(<-chan bool, 100)), reflect.ValueOf(make(chan bool, 100))
		diff := defaultChanDiff(zero, a, b, strat)
		expMsg := differs.ChanDirMismatch(zero, a.Type().ChanDir().String(),
			b.Type().ChanDir().String())
		assertNotEquals(t, diff, expMsg, "")
	})
}

// TestDefaultChanDiff_EdgeCases tests the method defaultChanDiff edge cases
//
//	This test must assert:
//	1. channels are equals only when:
//	- They point to the same underlying address; or
//	- Both have the same signature, direction and capacity; or
//	- Both channels are nil with the same signature
func TestDefaultChanDiff_EdgeCases(t *testing.T) {
	var zero indent.Branch
	strat := DefaultStrat()

	t.Run("same nil instance", func(t *testing.T) {
		var c chan int
		diff := defaultChanDiff(zero, reflect.ValueOf(c), reflect.ValueOf(c), strat)
		assertEquals(t, diff)
	})

	t.Run("different instances, both nil", func(t *testing.T) {
		var a, b chan int
		diff := defaultChanDiff(zero, reflect.ValueOf(a), reflect.ValueOf(b), strat)
		assertEquals(t, diff)
	})

	t.Run("one is nil, other is not", func(t *testing.T) {
		var c chan int
		null := reflect.ValueOf(c)
		notNull := reflect.ValueOf(make(chan int))
		// nil received
		t.Run("nil received", func(t *testing.T) {
			diff := defaultChanDiff(zero, null, notNull, strat)
			expMsg := differs.NilReceived(zero, null.Kind().String(), notNull.Type().String())
			assertNotEquals(t, diff, expMsg, "")
		})
		// nil expected
		t.Run("nil expected", func(t *testing.T) {
			diff := defaultChanDiff(zero, notNull, null, strat)
			expMsg := differs.NilExpected(zero, null.Kind().String(), notNull.Type().String())
			assertNotEquals(t, diff, expMsg, "")
		})
	})
}

// TestChanSignatureDiff tests the method defaultChanSignatureDiff
//
//	This test must assert
//	1. two channels are equals if:
//	- both have the same signature, direction and capacity;
func TestChanSignatureDiff(t *testing.T) {
	var zero indent.Branch

	t.Run("same instance", func(t *testing.T) {
		c := make(chan bool)
		a := reflect.ValueOf(c)
		b := reflect.ValueOf(c)
		diff := defaultChanSignatureDiff(zero, a, b)
		assertEquals(t, diff)
	})

	t.Run("same type, direction and buffer size", func(t *testing.T) {
		a := reflect.ValueOf(make(<-chan bool, 3))
		b := reflect.ValueOf(make(<-chan bool, 3))
		diff := defaultChanSignatureDiff(zero, a, b)
		assertEquals(t, diff)
	})

	t.Run("different signatures", func(t *testing.T) {
		a, b := reflect.ValueOf(make(chan bool)), reflect.ValueOf(make(chan int))
		expMsg := differs.ChanElemTypesMismatch(zero,
			types.ValidValueName(a.Type().Elem()), types.ValidValueName(b.Type().Elem()))
		diff := defaultChanSignatureDiff(zero, a, b)
		assertNotEquals(t, diff, expMsg, "")
	})

	t.Run("equals signature, different buffer size", func(t *testing.T) {
		a, b := make(chan bool, 100), make(chan bool, 1000)
		expMsg := differs.ChanBufSizeMismatch(zero, 100, 1000)
		diff := defaultChanSignatureDiff(zero, reflect.ValueOf(a), reflect.ValueOf(b))
		assertNotEquals(t, diff, expMsg, "")
	})

	t.Run("equals signature, different dir", func(t *testing.T) {
		a, b := reflect.ValueOf(make(<-chan bool, 100)), reflect.ValueOf(make(chan bool, 100))
		expMsg := differs.ChanDirMismatch(zero, a.Type().ChanDir().String(),
			b.Type().ChanDir().String())
		diff := defaultChanSignatureDiff(zero, a, b)
		assertNotEquals(t, diff, expMsg, "")
	})
}

// TestChanSignatureDiffEdgeCases tests the method defaultChanSignatureDiff edge cases
//
//	This test must assert
//	1. two channels are equals if:
//	- both have the same signature, direction and capacity;
func TestChanSignatureDiffEdgeCases(t *testing.T) {
	var zero indent.Branch

	t.Run("same nil instance", func(t *testing.T) {
		var c chan int
		diff := defaultChanSignatureDiff(zero, reflect.ValueOf(c), reflect.ValueOf(c))
		assertEquals(t, diff)
	})

	t.Run("different instances, both nil", func(t *testing.T) {
		var a, b chan int
		diff := defaultChanSignatureDiff(zero, reflect.ValueOf(a), reflect.ValueOf(b))
		assertEquals(t, diff)
	})
}

// TestNilChanSignatureDiff tests the method nilChanSignatureDiff edge cases
//
//	This test must assert
//	1. two channels are equals if:
//	- both have the same signature and direction - nil channel capacities are always zero
func TestNilChanSignatureDiff(t *testing.T) {
	var zero indent.Branch

	t.Run("same nil instance", func(t *testing.T) {
		var c chan int
		diff := defaultChanSignatureDiff(zero, reflect.ValueOf(c), reflect.ValueOf(c))
		assertEquals(t, diff)
	})

	t.Run("different instances, equals signatures, both nil", func(t *testing.T) {
		var a, b chan int
		diff := defaultChanSignatureDiff(zero, reflect.ValueOf(a), reflect.ValueOf(b))
		assertEquals(t, diff)
	})

	t.Run("different instances, different element, both nil", func(t *testing.T) {
		var a chan bool
		var b chan int
		aa := reflect.ValueOf(a)
		bb := reflect.ValueOf(b)
		expMsg := differs.NilChanTypesMismatch(zero,
			types.ValidValueName(aa.Type().Elem()), types.ValidValueName(bb.Type().Elem()))
		diff := nilChanSignatureDiff(zero, aa, bb)
		assertNotEquals(t, diff, expMsg, "")
	})

	t.Run("equals signature, different dir, both nil", func(t *testing.T) {
		var aa chan bool
		var bb chan<- bool
		a, b := reflect.ValueOf(aa), reflect.ValueOf(bb)
		expMsg := differs.NilChanDirMismatch(zero, a.Type().ChanDir().String(),
			b.Type().ChanDir().String())
		diff := nilChanSignatureDiff(zero, a, b)
		assertNotEquals(t, diff, expMsg, "")
	})
}

// TestSameChanDiff tests the method sameChanDiff
//
//	This test must assert
//	1. two channels are equals if:
//	- they point to the same underlying address; or
//	- both are nil with the same signature
func TestSameChanDiff(t *testing.T) {
	var zero indent.Branch
	strat := DefaultStrat()

	t.Run("same instance", func(t *testing.T) {
		c := make(chan bool)
		a := reflect.ValueOf(c)
		b := reflect.ValueOf(c)
		diff := sameChanDiff(zero, a, b, strat)
		assertEquals(t, diff)
	})

	t.Run("different instances, equals signature, equals buffer size", func(t *testing.T) {
		a, b := reflect.ValueOf(make(chan bool, 2)), reflect.ValueOf(make(chan bool, 2))
		expMsg := differs.ChanAddressMismatch(zero, sprints.Uintptrf(a.Pointer()), sprints.Uintptrf(b.Pointer()))
		diff := sameChanDiff(zero, a, b, strat)
		assertNotEquals(t, diff, expMsg, "")
	})
}

// TestSameChanDiff_EdgeCases tests the method sameChanDiff on edge cases
//
//	This test must assert
//	1. two channels are equals if:
//	- they point to the same underlying address; or
//	- both are nil with the same signature
func TestSameChanDiff_EdgeCases(t *testing.T) {
	var zero indent.Branch
	strat := DefaultStrat()

	t.Run("same nil instance", func(t *testing.T) {
		var c chan int
		diff := sameChanDiff(zero, reflect.ValueOf(c), reflect.ValueOf(c), strat)
		assertEquals(t, diff)
	})

	t.Run("different instances, equals signatures, both nil", func(t *testing.T) {
		var a, b chan int
		diff := sameChanDiff(zero, reflect.ValueOf(a), reflect.ValueOf(b), strat)
		assertEquals(t, diff)
	})

	t.Run("different instances, different element, both nil", func(t *testing.T) {
		var a chan bool
		var b chan int
		aa := reflect.ValueOf(a)
		bb := reflect.ValueOf(b)
		expMsg := differs.NilChanTypesMismatch(zero,
			types.ValidValueName(aa.Type().Elem()), types.ValidValueName(bb.Type().Elem()))
		diff := sameChanDiff(zero, aa, bb, strat)
		assertNotEquals(t, diff, expMsg, "")
	})

	t.Run("equals signature, different dir, both nil", func(t *testing.T) {
		var aa chan bool
		var bb chan<- bool
		a, b := reflect.ValueOf(aa), reflect.ValueOf(bb)
		expMsg := differs.NilChanDirMismatch(zero, a.Type().ChanDir().String(),
			b.Type().ChanDir().String())
		diff := sameChanDiff(zero, a, b, strat)
		assertNotEquals(t, diff, expMsg, "")
	})

	t.Run("one is nil, other is not", func(t *testing.T) {
		var c chan int
		null := reflect.ValueOf(c)
		notNull := reflect.ValueOf(make(chan int))
		// nil received
		t.Run("nil received", func(t *testing.T) {
			expMsg := differs.NilReceived(zero, null.Kind().String(), notNull.Type().String())
			diff := sameChanDiff(zero, null, notNull, strat)
			assertNotEquals(t, diff, expMsg, "")
		})
		// nil expected
		t.Run("nil expected", func(t *testing.T) {
			expMsg := differs.NilExpected(zero, null.Kind().String(), notNull.Type().String())
			diff := sameChanDiff(zero, notNull, null, strat)
			assertNotEquals(t, diff, expMsg, "")
		})
	})
}

// TestDefaultFuncDiff tests the function defaultFuncDiff
//
//	this test must assert:
//	1: the functions are equals if:
//	- both addresses are the same; or
//	- have the same signature - even if the implementations are different
func TestDefaultFuncDiff(t *testing.T) {
	var zero indent.Branch
	strat := DefaultStrat()

	t.Run("same instance", func(t *testing.T) {
		f := func() string { return "hello world" }
		a := reflect.ValueOf(f)
		b := reflect.ValueOf(f)
		diff := defaultFuncDiff(zero, a, b, strat)
		assertEquals(t, diff)
	})

	t.Run("different instances, same signature, different implementations", func(t *testing.T) {
		a := reflect.ValueOf(func() string { return "hello world" })
		b := reflect.ValueOf(func() string { return "hello universe" })
		diff := defaultFuncDiff(zero, a, b, strat)
		assertEquals(t, diff)
	})

	t.Run("different signatures", func(t *testing.T) {
		a := reflect.ValueOf(func() string { return "hello world" })
		b := reflect.ValueOf(func() int { return 10 })
		expMsg := differs.FuncTypesMismatch(zero, types.ValidValueName(a.Type()),
			types.ValidValueName(b.Type()))
		diff := defaultFuncDiff(zero, a, b, strat)
		assertNotEquals(t, diff, expMsg, "")
	})
}

// TestDefaultFuncDiff tests the function defaultFuncDiff against edge cases
//
//	this test must assert:
//	1: the functions are equals if:
//	- both addresses are the same; or
//	- have the same signature - even if the implementations are different
func TestDefaultFuncDiff_EdgeCases(t *testing.T) {
	var zero indent.Branch
	strat := DefaultStrat()

	t.Run("same nil instance", func(t *testing.T) {
		f := ptrs.Nil[functions.Supplier[any]]()
		a := reflect.ValueOf(f)
		b := reflect.ValueOf(f)
		diff := defaultFuncDiff(zero, a, b, strat)
		assertEquals(t, diff)
	})

	t.Run("different instances, same signature, both nil", func(t *testing.T) {
		a := reflect.ValueOf(ptrs.Nil[functions.Function[any, any]]())
		b := reflect.ValueOf(ptrs.Nil[functions.Function[any, any]]())
		diff := defaultFuncDiff(zero, a, b, strat)
		assertEquals(t, diff)
	})

	t.Run("different signatures, both nil", func(t *testing.T) {
		a := reflect.ValueOf(ptrs.Nil[functions.Function[any, any]]())
		b := reflect.ValueOf(ptrs.Nil[functions.Supplier[any]]())
		expMsg := differs.NilFuncTypesMismatch(zero, types.ValidValueName(a.Type()),
			types.ValidValueName(b.Type()))
		diff := defaultFuncDiff(zero, a, b, strat)
		assertNotEquals(t, diff, expMsg, "")
	})

	t.Run("one is nil, other is not", func(t *testing.T) {
		var nilf func()
		null := reflect.ValueOf(nilf)
		notNull := reflect.ValueOf(func() {})
		// nil received
		t.Run("nil received", func(t *testing.T) {
			expMsg := differs.NilReceived(zero, null.Kind().String(), notNull.Type().String())
			diff := defaultFuncDiff(zero, null, notNull, strat)
			assertNotEquals(t, diff, expMsg, "")
		})
		// nil expected
		t.Run("nil expected", func(t *testing.T) {
			expMsg := differs.NilExpected(zero, null.Kind().String(), notNull.Type().String())
			diff := defaultFuncDiff(zero, notNull, null, strat)
			assertNotEquals(t, diff, expMsg, "")
		})
	})
}

func TestFuncSignatureDiff(t *testing.T) {
	var zero indent.Branch

	t.Run("same instance", func(t *testing.T) {
		f := func() string { return "hello world" }
		a := reflect.ValueOf(f)
		b := reflect.ValueOf(f)
		diff := funcSignatureDiff(zero, a, b)
		assertEquals(t, diff)
	})

	t.Run("different instances, same signature, different implementations", func(t *testing.T) {
		a := reflect.ValueOf(func() string { return "hello world" })
		b := reflect.ValueOf(func() string { return "hello universe" })
		diff := funcSignatureDiff(zero, a, b)
		assertEquals(t, diff)
	})

	t.Run("different signatures", func(t *testing.T) {
		a := reflect.ValueOf(func() string { return "hello world" })
		b := reflect.ValueOf(func() int { return 10 })
		expMsg := differs.FuncTypesMismatch(zero, types.ValidValueName(a.Type()),
			types.ValidValueName(b.Type()))
		diff := funcSignatureDiff(zero, a, b)
		assertNotEquals(t, diff, expMsg, "")
	})
}

func TestFuncSignatureDiff_EdgeCases(t *testing.T) {
	var zero indent.Branch

	t.Run("same nil instance", func(t *testing.T) {
		f := ptrs.Nil[functions.Supplier[any]]()
		a := reflect.ValueOf(f)
		b := reflect.ValueOf(f)
		diff := funcSignatureDiff(zero, a, b)
		assertEquals(t, diff)
	})

	t.Run("different instances, same signature, both nil", func(t *testing.T) {
		a := reflect.ValueOf(ptrs.Nil[functions.Function[any, any]]())
		b := reflect.ValueOf(ptrs.Nil[functions.Function[any, any]]())
		diff := funcSignatureDiff(zero, a, b)
		assertEquals(t, diff)
	})

	t.Run("one is nil, other is not", func(t *testing.T) {
		null := reflect.ValueOf(generics.Zero[func()]())
		notNull := reflect.ValueOf(func() {})
		// nil received
		t.Run("nil received", func(t *testing.T) {
			diff := funcSignatureDiff(zero, null, notNull)
			assertEquals(t, diff)
		})
		// nil expected
		t.Run("nil expected", func(t *testing.T) {
			diff := funcSignatureDiff(zero, notNull, null)
			assertEquals(t, diff)
		})
	})
}

func TestNilFuncSignatureDiff(t *testing.T) {
	var zero indent.Branch

	t.Run("same instance", func(t *testing.T) {
		f := func() string { return "hello world" }
		a := reflect.ValueOf(f)
		b := reflect.ValueOf(f)
		diff := nilFuncSignatureDiff(zero, a, b)
		assertEquals(t, diff)
	})

	t.Run("different instances, same signature, different implementations", func(t *testing.T) {
		a := reflect.ValueOf(func() string { return "hello world" })
		b := reflect.ValueOf(func() string { return "hello universe" })
		diff := nilFuncSignatureDiff(zero, a, b)
		assertEquals(t, diff)
	})

	t.Run("different signatures", func(t *testing.T) {
		a := reflect.ValueOf(func() string { return "hello world" })
		b := reflect.ValueOf(func() int { return 10 })
		expMsg := differs.NilFuncTypesMismatch(zero, types.ValidValueName(a.Type()),
			types.ValidValueName(b.Type()))
		diff := nilFuncSignatureDiff(zero, a, b)
		assertNotEquals(t, diff, expMsg, "")
	})
}

func TestNilFuncSignatureDiff_EdgeCases(t *testing.T) {
	var zero indent.Branch

	t.Run("same nil instance", func(t *testing.T) {
		f := ptrs.Nil[functions.Supplier[any]]()
		a := reflect.ValueOf(f)
		b := reflect.ValueOf(f)
		diff := nilFuncSignatureDiff(zero, a, b)
		assertEquals(t, diff)
	})

	t.Run("different instances, same signature, both nil", func(t *testing.T) {
		a := reflect.ValueOf(ptrs.Nil[functions.Function[any, any]]())
		b := reflect.ValueOf(ptrs.Nil[functions.Function[any, any]]())
		diff := nilFuncSignatureDiff(zero, a, b)
		assertEquals(t, diff)
	})

	t.Run("different signatures, both nil", func(t *testing.T) {
		a := reflect.ValueOf(ptrs.Nil[functions.Function[any, any]]())
		b := reflect.ValueOf(ptrs.Nil[functions.Supplier[any]]())
		expMsg := differs.NilFuncTypesMismatch(zero, types.ValidValueName(a.Type()),
			types.ValidValueName(b.Type()))
		diff := nilFuncSignatureDiff(zero, a, b)
		assertNotEquals(t, diff, expMsg, "")
	})

	t.Run("one is nil, other is not", func(t *testing.T) {
		null := reflect.ValueOf(generics.Zero[func()]())
		notNull := reflect.ValueOf(func() {})
		// nil received
		t.Run("nil received", func(t *testing.T) {
			diff := nilFuncSignatureDiff(zero, null, notNull)
			assertEquals(t, diff)
		})
		// nil expected
		t.Run("nil expected", func(t *testing.T) {
			diff := nilFuncSignatureDiff(zero, notNull, null)
			assertEquals(t, diff)
		})
	})
}

// TestSameFuncDiff tests the function sameFuncDiff
//
//	this test must assert:
//	1: the functions are equals if:
//	- both addresses are the same; or
//	- both are nil and have the same signature - even if the implementations are different
func TestSameFuncDiff(t *testing.T) {
	var zero indent.Branch
	strat := DefaultStrat()

	t.Run("same instance", func(t *testing.T) {
		f := func() string { return "hello world" }
		a := reflect.ValueOf(f)
		b := reflect.ValueOf(f)
		diff := sameFuncDiff(zero, a, b, strat)
		assertEquals(t, diff)
	})

	t.Run("different instances, same signature, different implementations", func(t *testing.T) {
		a := reflect.ValueOf(func() string { return "hello world" })
		b := reflect.ValueOf(func() string { return "hello universe" })
		expMsg := differs.FuncAddressMismatch(zero, sprints.Uintptrf(a.Pointer()),
			sprints.Uintptrf(b.Pointer()))
		diff := sameFuncDiff(zero, a, b, strat)
		assertNotEquals(t, diff, expMsg, "")
	})

	t.Run("different signatures", func(t *testing.T) {
		a := reflect.ValueOf(func() string { return "hello world" })
		b := reflect.ValueOf(func() int { return 10 })
		expMsg := differs.FuncAddressMismatch(zero, sprints.Uintptrf(a.Pointer()),
			sprints.Uintptrf(b.Pointer()))
		diff := sameFuncDiff(zero, a, b, strat)
		assertNotEquals(t, diff, expMsg, "")
	})
}

// TestSameFuncDiff_EdgeCases tests the function sameFuncDiff against edge cases
//
//	this test must assert:
//	1: the functions are equals if:
//	- both addresses are the same; or
//	- both are nil and have the same signature - even if the implementations are different
func TestSameFuncDiff_EdgeCases(t *testing.T) {
	var zero indent.Branch
	strat := DefaultStrat()

	t.Run("same nil instance", func(t *testing.T) {
		f := ptrs.Nil[functions.Supplier[any]]()
		a := reflect.ValueOf(f)
		b := reflect.ValueOf(f)
		diff := sameFuncDiff(zero, a, b, strat)
		assertEquals(t, diff)
	})

	t.Run("different instances, same signature, both nil", func(t *testing.T) {
		a := reflect.ValueOf(ptrs.Nil[functions.Function[any, any]]())
		b := reflect.ValueOf(ptrs.Nil[functions.Function[any, any]]())
		diff := sameFuncDiff(zero, a, b, strat)
		assertEquals(t, diff)
	})

	t.Run("different signatures, both nil", func(t *testing.T) {
		a := reflect.ValueOf(ptrs.Nil[functions.Function[any, any]]())
		b := reflect.ValueOf(ptrs.Nil[functions.Supplier[any]]())
		expMsg := differs.NilFuncTypesMismatch(zero, types.ValidValueName(a.Type()),
			types.ValidValueName(b.Type()))
		diff := sameFuncDiff(zero, a, b, strat)
		assertNotEquals(t, diff, expMsg, "")
	})

	t.Run("one is nil, other is not", func(t *testing.T) {
		null := reflect.ValueOf(generics.Zero[func()]())
		notNull := reflect.ValueOf(func() {})
		// nil received
		t.Run("nil received", func(t *testing.T) {
			expMsg := differs.NilReceived(zero, null.Kind().String(), notNull.Type().String())
			diff := sameFuncDiff(zero, null, notNull, strat)
			assertNotEquals(t, diff, expMsg, "")
		})
		// nil expected
		t.Run("nil expected", func(t *testing.T) {
			expMsg := differs.NilExpected(zero, null.Kind().String(), notNull.Type().String())
			diff := sameFuncDiff(zero, notNull, null, strat)
			assertNotEquals(t, diff, expMsg, "")
		})
	})
}

// TestDifferArrays tests the method diffArrays
//
//	this method should assert
//	1. Arrays are equals only when:
//	- both elements and length are equals; and
//	- all its elements are equals
func TestDifferArrays(t *testing.T) {
	var zero indent.Branch
	strat := DefaultStrat()

	t.Run("same signature", func(t *testing.T) {
		m := [1]int{1}
		a, b := reflect.ValueOf(m), reflect.ValueOf(m)
		diff := diffArrays(zero, a, b, strat)
		assertEquals(t, diff)
	})

	// if fact, arrays with different length are considered different types
	t.Run("same element, different length", func(t *testing.T) {
		var a1 [3]int
		var a2 [2]int
		a, b := reflect.ValueOf(a1), reflect.ValueOf(a2)
		diff := diffArrays(zero, a, b, strat)
		expMsg := differs.ArrayLenMismatch(zero, a.Len(), b.Len())
		assertNotEquals(t, diff, expMsg, "")
	})

	t.Run("different elements, same length", func(t *testing.T) {
		var a1 [3]int
		var a2 [3]int8
		a, b := reflect.ValueOf(a1), reflect.ValueOf(a2)
		diff := diffArrays(zero, a, b, strat)
		expMsg := differs.ArrayTypesMismatch(zero, types.ValidValueName(a.Type()),
			types.ValidValueName(b.Type()))
		assertNotEquals(t, diff, expMsg, "")
	})

	t.Run("same element, same length, different values", func(t *testing.T) {
		a := reflect.ValueOf([3]int8{1, 2, 3})
		b := reflect.ValueOf([3]int8{1, 2, 2})
		diff := diffArrays(zero, a, b, strat)
		expMsg := differs.Append(zero, differs.ArrayElem(zero, 2),
			differs.Values(zero.Inc(), sprints.Typed("int8", 3), sprints.Typed("int8", 2)))
		assertNotEquals(t, diff, expMsg, "")
	})
}

func TestDifferBool(t *testing.T) {
	runPrimTests(t, boolSeed, differBool)
}

func TestDifferComplex(t *testing.T) {
	runPrimTests(t, complexSeed, differComplex)
}

func TestDifferFloat(t *testing.T) {
	runPrimTests(t, floatSeed, differFloat)
}

func TestDifferInt(t *testing.T) {
	runPrimTests(t, intSeed, differInt)
}

// TestDifferInterfaces tests the method differInterfaces
//
//	this test must assert:
//	1. Interfaces are equals only if:
//	A: both interface types and implementations are equals
func TestDifferInterfaces(t *testing.T) {
	var zero indent.Branch
	strat := DefaultStrat()

	t.Run("same", func(t *testing.T) {
		var iface interf.DataCalc[models.Float] = models.FloatAsValue(16)
		a := reflect.ValueOf(&iface).Elem()
		diff := differInterfaces(zero, a, a, strat)
		assertEquals(t, diff)
	})

	t.Run("equals", func(t *testing.T) {
		var i1 interf.DataCalc[models.Float] = models.FloatAsValue(16)
		var i2 interf.DataCalc[models.Float] = models.FloatAsValue(16)
		a := reflect.ValueOf(&i1).Elem()
		b := reflect.ValueOf(&i2).Elem()
		diff := differInterfaces(zero, a, b, strat)
		assertEquals(t, diff)
	})

	t.Run("same interface, different impl", func(t *testing.T) {
		i1 := interf.TwoDataAsTwoMethodsRef("abc", 123)
		i2 := interf.ThreeDataAsTwoMethodsRef("abc", 123, 3.14)
		a := reflect.ValueOf(&i1).Elem()
		b := reflect.ValueOf(&i2).Elem()
		diff := differInterfaces(zero, a, b, strat)
		expMsg := differs.Chain(differs.InterfaceImpl(zero),
			differs.PointerValues(zero.Plus(1)), differs.StructTypesMismatch(zero.Plus(2),
				types.ValidValueName(a.Elem().Elem().Type()), types.ValidValueName(b.Elem().Elem().Type())))
		assertNotEquals(t, diff, expMsg, "")
	})

	t.Run("different types, same impl", func(t *testing.T) {
		i1 := interf.ThreeDataAsTwoMethodsRef("abc", 123, 3.14)
		i2 := interf.ThreeDataAsThreeMethodsRef("abc", 123, 3.14)
		a := reflect.ValueOf(&i1).Elem()
		b := reflect.ValueOf(&i2).Elem()
		diff := differInterfaces(zero, a, b, strat)
		expMsg := differs.InterfaceTypesMismatch(zero, types.ValidValueName(a.Type()),
			types.ValidValueName(b.Type()))
		assertNotEquals(t, diff, expMsg, "")
	})
}

// TestDifferInterfaces tests the method differInterfaces edge cases
//
//	this test must assert:
//	1. Interfaces are equals only if:
//	A: both interface types and implementations are equals
func TestDifferInterfaces_EdgeCases(t *testing.T) {
	var zero indent.Branch
	strat := DefaultStrat()

	t.Run("same nil", func(t *testing.T) {
		x := generics.Zero[interf.OneMethod]()
		a := reflect.ValueOf(&x).Elem()
		b := reflect.ValueOf(&x).Elem()
		diff := differInterfaces(zero, a, b, strat)
		assertEquals(t, diff)
	})

	t.Run("nil, different types", func(t *testing.T) {
		i1 := generics.Zero[interf.OneMethod]()
		i2 := generics.Zero[interf.TwoMethods]()
		a := reflect.ValueOf(&i1).Elem()
		b := reflect.ValueOf(&i2).Elem()
		diff := differInterfaces(zero, a, b, strat)
		expMsg := differs.NilTypesMismatch(zero, "interfaces", types.ValidValueName(a.Type()),
			types.ValidValueName(b.Type()))
		assertNotEquals(t, diff, expMsg, "")
	})

	t.Run("one is nil, other is not", func(t *testing.T) {
		null := generics.Zero[interf.OneMethod]()
		notNull := interf.OneDataAsRef("abc")
		t.Run("nil received", func(t *testing.T) {
			a := reflect.ValueOf(&null).Elem()
			b := reflect.ValueOf(&notNull).Elem()
			diff := differInterfaces(zero, a, b, strat)
			expMsg := differs.NilReceived(zero, "interface", types.ValidValueName(b.Type()))
			assertNotEquals(t, diff, expMsg, "")
		})
		t.Run("nil expected", func(t *testing.T) {
			a := reflect.ValueOf(&notNull).Elem()
			b := reflect.ValueOf(&null).Elem()
			diff := differInterfaces(zero, a, b, strat)
			expMsg := differs.NilExpected(zero, "interface", types.ValidValueName(a.Type()))
			assertNotEquals(t, diff, expMsg, "")
		})
	})
}

func TestDifferPrimitives(t *testing.T) {
	t.Run("bool", func(t *testing.T) {
		runPrimTests(t, boolSeed, differPrimitives)
	})

	t.Run("complex", func(t *testing.T) {
		runPrimTests(t, complexSeed, differPrimitives)
	})

	t.Run("float", func(t *testing.T) {
		runPrimTests(t, floatSeed, differPrimitives)
	})

	t.Run("int", func(t *testing.T) {
		runPrimTests(t, intSeed, differPrimitives)
	})

	t.Run("uint", func(t *testing.T) {
		runPrimTests(t, uintSeed, differPrimitives)
	})
}

// TestDefaultMapDiff tests the method defaultMapDiff
//
//	this method must assert:
//	1. two maps are equals only if:
//	A: they are exact the same instances; or
//	B: they have the same signature; and
//	B: they have the same length; and
//	B: their elements are equals; or
//	C: both types are nil: and
//	C: both types have the same signature
func TestDefaultMapDiff(t *testing.T) {
	var zero indent.Branch
	strat := DefaultStrat()

	t.Run("same instances", func(t *testing.T) {
		x := make(map[string]string)
		a := reflect.ValueOf(x)
		b := reflect.ValueOf(x)
		diff := defaultMapDiff(zero, a, b, strat)
		assertEquals(t, diff)
	})

	t.Run("same signature, same length", func(t *testing.T) {
		a := reflect.ValueOf(make(map[string]string))
		b := reflect.ValueOf(make(map[string]string))
		diff := defaultMapDiff(zero, a, b, strat)
		assertEquals(t, diff)
	})

	t.Run("different signature, same length", func(t *testing.T) {
		a := reflect.ValueOf(make(map[string]string))
		b := reflect.ValueOf(make(map[string]any))
		diff := defaultMapDiff(zero, a, b, strat)
		expMsg := differs.MapTypesMismatch(zero, types.ValidValueName(a.Type()),
			types.ValidValueName(b.Type()))
		assertNotEquals(t, diff, expMsg, "")
	})

	t.Run("same signature, different length", func(t *testing.T) {
		a := reflect.ValueOf(map[string]int{"a": 1, "b": 2})
		b := reflect.ValueOf(map[string]int{"a": 1, "b": 2, "c": 3})
		diff := defaultMapDiff(zero, a, b, strat)
		expDiff := differs.MapKeysDiff(indent.Tab(zero), []string{"c"}, []string{}, []string{}, []string{"c"})
		expMsg := differs.MapLenMismatch(zero, 2, 3)
		assertNotEquals(t, diff, expMsg, expDiff)
	})

	t.Run("same signature, same length, different elements", func(t *testing.T) {
		a := reflect.ValueOf(map[string]int{"a": 1, "b": 2, "c": 3})
		b := reflect.ValueOf(map[string]int{"a": 1, "b": 2, "c": 4})
		expMsg := differs.Append(zero, differs.MapValue(zero, "c"),
			differs.Values(zero.Inc(), sprints.Typed("int", 3), sprints.Typed("int", 4)))
		diff := defaultMapDiff(zero, a, b, strat)
		assertNotEquals(t, diff, expMsg, "")
	})
}

// TestDefaultMapDiff tests the method defaultMapDiff for edge cases.
//
//	this method must assert:
//	1. two maps are equals only if:
//	A: they are exact the same instances; or
//	B: they have the same signature; and
//	B: they have the same length; and
//	B: their elements are equals; or
//	C: both types are nil: and
//	C: both types have the same signature
func TestDefaultMapDiff_EdgeCases(t *testing.T) {
	var zero indent.Branch
	strat := DefaultStrat()

	t.Run("same nil instance", func(t *testing.T) {
		var x map[string]string
		a := reflect.ValueOf(x)
		b := reflect.ValueOf(x)
		diff := defaultMapDiff(zero, a, b, strat)
		assertEquals(t, diff)
	})

	t.Run("same signature, nil", func(t *testing.T) {
		a := reflect.ValueOf(generics.Zero[map[string]any]())
		b := reflect.ValueOf(generics.Zero[map[string]any]())
		diff := defaultMapDiff(zero, a, b, strat)
		assertEquals(t, diff)
	})

	t.Run("one nil, one valid", func(t *testing.T) {
		null := generics.Zero[map[string]any]()
		notNull := make(map[string]any)
		t.Run("nil received", func(t *testing.T) {
			a := reflect.ValueOf(null)
			b := reflect.ValueOf(notNull)
			diff := defaultMapDiff(zero, a, b, strat)
			expMsg := differs.NilReceived(zero, "map", types.ValidValueName(b.Type()))
			assertNotEquals(t, diff, expMsg, "")
		})
		t.Run("nil expected", func(t *testing.T) {
			a := reflect.ValueOf(notNull)
			b := reflect.ValueOf(null)
			diff := defaultMapDiff(zero, a, b, strat)
			expMsg := differs.NilExpected(zero, "map", types.ValidValueName(b.Type()))
			assertNotEquals(t, diff, expMsg, "")
		})
	})
}

// TestSerializableMapDiff tests the method serializableMapDiff
//
//	this method must assert:
//	1. two maps are equals only if:
//	A: they are exact the same instances; or
//	B: they have the same signature; and
//	B: they have the same length; and
//	B: their elements are equals; or
//	C: both types are nil or empty: and
//	C: both types have the same signature
func TestSerializableMapDiff(t *testing.T) {
	var zero indent.Branch
	strat := DefaultStrat()

	t.Run("same instances", func(t *testing.T) {
		x := make(map[string]string)
		a := reflect.ValueOf(x)
		b := reflect.ValueOf(x)
		diff := serializableMapDiff(zero, a, b, strat)
		assertEquals(t, diff)
	})

	t.Run("same signature, same length", func(t *testing.T) {
		a := reflect.ValueOf(make(map[string]string))
		b := reflect.ValueOf(make(map[string]string))
		diff := serializableMapDiff(zero, a, b, strat)
		assertEquals(t, diff)
	})

	t.Run("different signature, same length", func(t *testing.T) {
		a := reflect.ValueOf(make(map[string]string))
		b := reflect.ValueOf(make(map[string]any))
		diff := serializableMapDiff(zero, a, b, strat)
		expMsg := differs.MapTypesMismatch(zero, types.ValidValueName(a.Type()),
			types.ValidValueName(b.Type()))
		assertNotEquals(t, diff, expMsg, "")
	})

	t.Run("same signature, different length", func(t *testing.T) {
		a := reflect.ValueOf(map[string]int{"a": 1, "b": 2})
		b := reflect.ValueOf(map[string]int{"a": 1, "b": 2, "c": 3})
		diff := serializableMapDiff(zero, a, b, strat)
		expDiff := differs.MapKeysDiff(indent.Tab(zero), []string{"c"}, []string{}, []string{}, []string{"c"})
		expMsg := differs.MapLenMismatch(zero, 2, 3)
		assertNotEquals(t, diff, expMsg, expDiff)
	})

	t.Run("same signature, same length, different elements", func(t *testing.T) {
		a := reflect.ValueOf(map[string]int{"a": 1, "b": 2, "c": 3})
		b := reflect.ValueOf(map[string]int{"a": 1, "b": 2, "c": 4})
		expMsg := differs.Append(zero, differs.MapValue(zero, "c"),
			differs.Values(zero.Inc(), sprints.Typed("int", 3), sprints.Typed("int", 4)))
		diff := serializableMapDiff(zero, a, b, strat)
		assertNotEquals(t, diff, expMsg, "")
	})
}

// TestSerializableMapDiff tests the method serializableMapDiff for edge cases.
//
//	this method must assert:
//	1. two maps are equals only if:
//	A: they are exact the same instances; or
//	B: they have the same signature; and
//	B: they have the same length; and
//	B: their elements are equals; or
//	C: both types are nil or empty: and
//	C: both types have the same signature
func TestSerializableMapDiff_EdgeCases(t *testing.T) {
	var zero indent.Branch
	strat := DefaultStrat()

	t.Run("same nil instance", func(t *testing.T) {
		var x map[string]string
		a := reflect.ValueOf(x)
		b := reflect.ValueOf(x)
		diff := serializableMapDiff(zero, a, b, strat)
		assertEquals(t, diff)
	})

	t.Run("same signature, nil", func(t *testing.T) {
		a := reflect.ValueOf(generics.Zero[map[string]any]())
		b := reflect.ValueOf(generics.Zero[map[string]any]())
		diff := serializableMapDiff(zero, a, b, strat)
		assertEquals(t, diff)
	})

	t.Run("one nil, one valid", func(t *testing.T) {
		null := generics.Zero[map[string]any]()
		notEmpty := map[string]any{"a": 1, "b": "two"}
		empty := make(map[string]any)

		t.Run("nil received, not empty expected", func(t *testing.T) {
			a := reflect.ValueOf(null)
			b := reflect.ValueOf(notEmpty)
			diff := serializableMapDiff(zero, a, b, strat)
			expMsg := differs.NilReceived(zero, "map", types.ValidValueName(b.Type()))
			assertNotEquals(t, diff, expMsg, "")
		})

		t.Run("nil received, empty expected", func(t *testing.T) {
			a := reflect.ValueOf(null)
			b := reflect.ValueOf(empty)
			diff := serializableMapDiff(zero, a, b, strat)
			assertEquals(t, diff)
		})

		t.Run("not empty received, nil expected", func(t *testing.T) {
			a := reflect.ValueOf(notEmpty)
			b := reflect.ValueOf(null)
			diff := serializableMapDiff(zero, a, b, strat)
			expMsg := differs.NilExpected(zero, "map", types.ValidValueName(a.Type()))
			assertNotEquals(t, diff, expMsg, "")
		})

		t.Run("empty received, nil expected", func(t *testing.T) {
			a := reflect.ValueOf(empty)
			b := reflect.ValueOf(null)
			diff := serializableMapDiff(zero, a, b, strat)
			assertEquals(t, diff)
		})
	})
}

// TestNotNilMapDiffer tests the method notNilMapDiffer
//
//	this method must assert:
//	1. two maps are equals only
//	A: they are exact the same instances; or
//	B: they have the same signature; and
//	B: they have the same length; and
//	B: their elements are equals;
func TestNotNilMapDiffer(t *testing.T) {
	var zero indent.Branch
	strat := DefaultStrat()

	t.Run("same instances", func(t *testing.T) {
		x := make(map[string]string)
		a := reflect.ValueOf(x)
		b := reflect.ValueOf(x)
		diff := notNilMapDiffer(zero, a, b, strat)
		assertEquals(t, diff)
	})

	t.Run("same signature, same length", func(t *testing.T) {
		a := reflect.ValueOf(make(map[string]string))
		b := reflect.ValueOf(make(map[string]string))
		diff := notNilMapDiffer(zero, a, b, strat)
		assertEquals(t, diff)
	})

	t.Run("different signature, same length", func(t *testing.T) {
		a := reflect.ValueOf(make(map[string]string))
		b := reflect.ValueOf(make(map[string]any))
		diff := notNilMapDiffer(zero, a, b, strat)
		expMsg := differs.MapTypesMismatch(zero, types.ValidValueName(a.Type()),
			types.ValidValueName(b.Type()))
		assertNotEquals(t, diff, expMsg, "")
	})

	t.Run("same signature, different length", func(t *testing.T) {
		a := reflect.ValueOf(map[string]int{"a": 1, "b": 2})
		b := reflect.ValueOf(map[string]int{"a": 1, "b": 2, "c": 3})
		diff := notNilMapDiffer(zero, a, b, strat)
		expDiff := differs.MapKeysDiff(indent.Tab(zero), []string{"c"}, []string{}, []string{}, []string{"c"})
		expMsg := differs.MapLenMismatch(zero, 2, 3)
		assertNotEquals(t, diff, expMsg, expDiff)
	})

	t.Run("same signature, same length, different elements", func(t *testing.T) {
		a := reflect.ValueOf(map[string]int{"a": 1, "b": 2, "c": 3})
		b := reflect.ValueOf(map[string]int{"a": 1, "b": 2, "c": 4})
		expMsg := differs.Append(zero, differs.MapValue(zero, "c"),
			differs.Values(zero.Inc(), sprints.Typed("int", 3), sprints.Typed("int", 4)))
		diff := notNilMapDiffer(zero, a, b, strat)
		assertNotEquals(t, diff, expMsg, "")
	})
}

// TestNotNilMapDiffer_EdgeCases tests the method notNilMapDiffer edge cases
//
//	this method must assert:
//	1. two maps are equals only
//	A: they are exact the same instances; or
//	B: they have the same signature; and
//	B: they have the same length; and
//	B: their elements are equals;
func TestNotNilMapDiffer_EdgeCases(t *testing.T) {
	// TODO:
}

// TestDefaultPtrDiff tests the method defaultPtrDiff
//
//	this method must assert:
//	1. two ptrs are equals only if:
//	A: both addresses are the same; or
//	B: both types are equals; and
//	B: both elements are equals; or
//	C: both ptrs are nil; and
//	C: both ptrs have the same type
func TestDefaultPtrDiff(t *testing.T) {
	var zero indent.Branch
	strat := DefaultStrat()

	t.Run("same ptrs", func(t *testing.T) {
		x := ptrs.New(10)
		a := reflect.ValueOf(x)
		b := reflect.ValueOf(x)
		diff := defaultPtrDiff(zero, a, b, strat)
		assertEquals(t, diff)
	})

	t.Run("different addresses, equals types and values", func(t *testing.T) {
		a := reflect.ValueOf(ptrs.New(10))
		b := reflect.ValueOf(ptrs.New(10))
		diff := defaultPtrDiff(zero, a, b, strat)
		assertEquals(t, diff)
	})

	t.Run("equals types, different values", func(t *testing.T) {
		a := reflect.ValueOf(ptrs.New(1))
		b := reflect.ValueOf(ptrs.New(2))
		expMsg := differs.Append(zero, differs.PointerValues(zero),
			differs.Values(zero.Inc(), sprints.Typed("int", 1), sprints.Typed("int", 2)))
		diff := defaultPtrDiff(zero, a, b, strat)
		assertNotEquals(t, diff, expMsg, "")
	})

	t.Run("different types, equals values", func(t *testing.T) {
		a := reflect.ValueOf(ptrs.New[any](1))
		b := reflect.ValueOf(ptrs.New[int](2))
		expMsg := differs.Append(zero, differs.PointerValues(zero),
			differs.TypesMismatch(zero.Inc(), a.Elem().Kind().String(), b.Elem().Kind().String()))
		diff := defaultPtrDiff(zero, a, b, strat)
		assertNotEquals(t, diff, expMsg, "")
	})
}

// TestDefaultPtrDiff_EdgeCases tests the method defaultPtrDiff edge cases
//
//	this method must assert:
//	1. two ptrs are equals only if:
//	A: both addresses are the same; or
//	B: both types are equals; and
//	B: both elements are equals; or
//	C: both ptrs are nil; and
//	C: both ptrs have the same type
func TestDefaultPtrDiff_EdgeCases(t *testing.T) {
	var zero indent.Branch
	strat := DefaultStrat()

	t.Run("nil pointer, same type", func(t *testing.T) {
		a := reflect.ValueOf(ptrs.Nil[int]())
		b := reflect.ValueOf(ptrs.Nil[int]())
		diff := defaultPtrDiff(zero, a, b, strat)
		assertEquals(t, diff)
	})

	t.Run("nil ptrs, different type", func(t *testing.T) {
		a := reflect.ValueOf(ptrs.Nil[int]())
		b := reflect.ValueOf(ptrs.Nil[any]())
		diff := defaultPtrDiff(zero, a, b, strat)
		expMsg := differs.NilPointersSignMismatch(zero, types.ValidValueName(a.Type()),
			types.ValidValueName(b.Type()))
		assertNotEquals(t, diff, expMsg, "")
	})

	t.Run("one is nil, other is not", func(t *testing.T) {
		null := ptrs.Nil[int]()
		notNull := ptrs.New[int](10)
		t.Run("nil received", func(t *testing.T) {
			a := reflect.ValueOf(null)
			b := reflect.ValueOf(notNull)
			diff := defaultPtrDiff(zero, a, b, strat)
			expMsg := differs.NilReceived(zero, b.Kind().String(), types.ValidValueName(b.Type()))
			assertNotEquals(t, diff, expMsg, "")
		})
		t.Run("nil expected", func(t *testing.T) {
			a := reflect.ValueOf(notNull)
			b := reflect.ValueOf(null)
			diff := defaultPtrDiff(zero, a, b, strat)
			expMsg := differs.NilExpected(zero, b.Kind().String(), types.ValidValueName(a.Type()))
			assertNotEquals(t, diff, expMsg, "")
		})
	})
}

// TestPtrElemDiff tests the method ptrElemDiff
//
//	this test must assert:
//	1. two ptrs are equals if:
//	- their elements are equals
//	2. If the pointer is nil this test must return an error
func TestPtrElemDiff(t *testing.T) {
	var zero indent.Branch
	strat := DefaultStrat()
	t.Run("same ptrs", func(t *testing.T) {
		x := ptrs.New(10)
		a := reflect.ValueOf(x)
		b := reflect.ValueOf(x)
		diff := ptrElemDiff(zero, a, b, strat)
		assertEquals(t, diff)
	})
	t.Run("different addresses, equals types and values", func(t *testing.T) {
		a := reflect.ValueOf(ptrs.New(10))
		b := reflect.ValueOf(ptrs.New(10))
		diff := ptrElemDiff(zero, a, b, strat)
		assertEquals(t, diff)
	})

	t.Run("equals types, different values", func(t *testing.T) {
		a := reflect.ValueOf(ptrs.New(1))
		b := reflect.ValueOf(ptrs.New(2))
		expMsg := differs.Append(zero, differs.PointerValues(zero),
			differs.Values(zero.Inc(), sprints.Typed("int", 1), sprints.Typed("int", 2)))
		diff := ptrElemDiff(zero, a, b, strat)
		assertNotEquals(t, diff, expMsg, "")
	})

	t.Run("different types, equals values", func(t *testing.T) {
		a := reflect.ValueOf(ptrs.New[any](1))
		b := reflect.ValueOf(ptrs.New[int](2))
		expMsg := differs.Append(zero, differs.PointerValues(zero),
			differs.TypesMismatch(zero.Inc(), a.Elem().Kind().String(), b.Elem().Kind().String()))
		diff := ptrElemDiff(zero, a, b, strat)
		assertNotEquals(t, diff, expMsg, "")
	})
}

// TestPtrElemDiff_EdgeCases tests the method ptrElemDiff edge cases
//
//	this test must assert:
//	1. two ptrs are equals if:
//	- their elements are equals
//	2. If the pointer is nil, this test must return an error - these errors
//	should never happen in real cases, since this method never
//	receives nil elements
func TestPtrElemDiff_EdgeCases(t *testing.T) {
	var zero indent.Branch
	strat := DefaultStrat()

	t.Run("nil pointer, same type", func(t *testing.T) {
		a := reflect.ValueOf(ptrs.Nil[int]())
		b := reflect.ValueOf(ptrs.Nil[int]())
		diff := ptrElemDiff(zero, a, b, strat)
		expMsg := differs.Append(zero, differs.PointerValues(zero),
			differs.BothInvalid(zero.Inc()))
		assertNotEquals(t, diff, expMsg, "")
	})

	t.Run("nil ptrs, different type", func(t *testing.T) {
		a := reflect.ValueOf(ptrs.Nil[int]())
		b := reflect.ValueOf(ptrs.Nil[any]())
		diff := ptrElemDiff(zero, a, b, strat)
		expMsg := differs.Append(zero, differs.PointerValues(zero),
			differs.BothInvalid(zero.Inc()))
		assertNotEquals(t, diff, expMsg, "")
	})

	t.Run("one is nil, other is not", func(t *testing.T) {
		null := ptrs.Nil[int]()
		notNull := ptrs.New[int](10)
		t.Run("nil received", func(t *testing.T) {
			a := reflect.ValueOf(null)
			b := reflect.ValueOf(notNull)
			diff := ptrElemDiff(zero, a, b, strat)
			expMsg := differs.Append(zero, differs.PointerValues(zero),
				differs.InvalidReceived(zero.Inc(), "int"))
			assertNotEquals(t, diff, expMsg, "")
		})
		t.Run("nil expected", func(t *testing.T) {
			a := reflect.ValueOf(notNull)
			b := reflect.ValueOf(null)
			diff := ptrElemDiff(zero, a, b, strat)
			expMsg := differs.Append(zero, differs.PointerValues(zero),
				differs.InvalidExpected(zero.Inc(), "int"))
			assertNotEquals(t, diff, expMsg, "")
		})
	})
}

// TestPSamePtrDiff tests the method samePtrDiff
//
//	this test must assert:
//	1. two ptrs are equals if:
//	A: both ptrs hold the same address; or
//	B: both ptrs are nil; and
//	B: both ptrs have the same type
func TestSamePtrDiff(t *testing.T) {
	var zero indent.Branch
	strat := DefaultStrat()
	t.Run("same ptrs", func(t *testing.T) {
		x := ptrs.New(10)
		a := reflect.ValueOf(x)
		b := reflect.ValueOf(x)
		diff := samePtrDiff(zero, a, b, strat)
		assertEquals(t, diff)
	})
	t.Run("different addresses, equals types and values", func(t *testing.T) {
		a := reflect.ValueOf(ptrs.New(10))
		b := reflect.ValueOf(ptrs.New(10))
		diff := samePtrDiff(zero, a, b, strat)
		spa, spb := sprints.Uintptrf(a.Pointer()), sprints.Uintptrf(b.Pointer())
		expMsg := differs.PointersAddrMismatch(zero, spa, spb)
		assertNotEquals(t, diff, expMsg, "")
	})

	t.Run("equals types, different values", func(t *testing.T) {
		a := reflect.ValueOf(ptrs.New(10))
		b := reflect.ValueOf(ptrs.New(11))
		diff := samePtrDiff(zero, a, b, strat)
		spa, spb := sprints.Uintptrf(a.Pointer()), sprints.Uintptrf(b.Pointer())
		expMsg := differs.PointersAddrMismatch(zero, spa, spb)
		assertNotEquals(t, diff, expMsg, "")
	})

	t.Run("different types, equals values", func(t *testing.T) {
		a := reflect.ValueOf(ptrs.New[any](10))
		b := reflect.ValueOf(ptrs.New[int](10))
		diff := samePtrDiff(zero, a, b, strat)
		spa, spb := sprints.Uintptrf(a.Pointer()), sprints.Uintptrf(b.Pointer())
		expMsg := differs.PointersAddrMismatch(zero, spa, spb)
		assertNotEquals(t, diff, expMsg, "")
	})
}

// TestPSamePtrDiff_EdgeCases tests the method samePtrDiff edge cases
//
//	this test must assert:
//	1. two ptrs are equals if:
//	A: both ptrs hold the same address; or
//	B: both ptrs are nil; and
//	B: both ptrs have the same type
func TestSamePtrDiff_EdgeCases(t *testing.T) {
	var zero indent.Branch
	strat := DefaultStrat()

	t.Run("nil pointer, same type", func(t *testing.T) {
		a := reflect.ValueOf(ptrs.Nil[int]())
		b := reflect.ValueOf(ptrs.Nil[int]())
		diff := samePtrDiff(zero, a, b, strat)
		assertEquals(t, diff)
	})

	t.Run("nil ptrs, different type", func(t *testing.T) {
		a := reflect.ValueOf(ptrs.Nil[int]())
		b := reflect.ValueOf(ptrs.Nil[any]())
		diff := samePtrDiff(zero, a, b, strat)
		sa, sb := types.ValidValueName(a.Type()), types.ValidValueName(b.Type())
		expMsg := differs.NilPointersSignMismatch(zero, sa, sb)
		assertNotEquals(t, diff, expMsg, "")
	})

	t.Run("one is nil, other is not", func(t *testing.T) {
		null := ptrs.Nil[int]()
		notNull := ptrs.New[int](10)
		t.Run("nil received", func(t *testing.T) {
			a := reflect.ValueOf(null)
			b := reflect.ValueOf(notNull)
			diff := samePtrDiff(zero, a, b, strat)
			expMsg := differs.NilReceived(zero, "ptr", types.ValidValueName(b.Type()))
			assertNotEquals(t, diff, expMsg, "")
		})
		t.Run("nil expected", func(t *testing.T) {
			a := reflect.ValueOf(notNull)
			b := reflect.ValueOf(null)
			diff := samePtrDiff(zero, a, b, strat)
			expMsg := differs.NilExpected(zero, "ptr", types.ValidValueName(a.Type()))
			assertNotEquals(t, diff, expMsg, "")
		})
	})
}

// TestDifferStrings_Default tests the method diffStrings based
// on the default strategy
//
//	this test should assert:
//	1. a string should be equals if:
//	- both have the exact same chars at the same exact positions; and
//	- exact same length
func TestDifferStrings_Default(t *testing.T) {
	before, _ := standards.DifferStartIndex, standards.DifferMaxSize
	var zero indent.Tab
	strat := DefaultStrat()
	t.Run("equals strings", func(t *testing.T) {
		a := "lorem ipsum"
		b := "lorem ipsum"
		diff := diffStrings(zero, a, b, strat)
		assertEquals(t, diff)
	})

	t.Run("space difference", func(t *testing.T) {
		a := "loremipsum "
		b := "lorem ipsum"
		diff := diffStrings(zero, a, b, strat)
		expMsg := differs.Strings(zero, 5)
		rec := fmx.Sprintf("'%s'", markDiff(a, 5))
		exp := fmx.Sprintf("'%s'", b)
		expDiff := differs.Diff(zero, rec, exp)
		assertNotEquals(t, diff, expMsg, expDiff)
		//t.Log(expDiff)
	})

	t.Run("trailing newline in A", func(t *testing.T) {
		a := "lorem ipsum\n"
		b := "lorem ipsum"
		diff := diffStrings(zero, a, b, strat)
		expMsg := differs.Strings(zero, 11)
		rec := fmx.Sprintf("'%s'", markDiff(a, 11)[11-before:])
		exp := fmx.Sprintf("'%s'", b[11-before:])
		expDiff := differs.Diff(zero, rec, exp)
		assertNotEquals(t, diff, expMsg, expDiff)
	})

	t.Run("case difference", func(t *testing.T) {
		a := "lorem ipsum"
		b := "Lorem ipsum"
		diff := diffStrings(zero, a, b, strat)
		expMsg := differs.Strings(zero, 0)
		rec := fmx.Sprintf("'%s'", markDiff(a, 0))
		exp := fmx.Sprintf("'%s'", b)
		expDiff := differs.Diff(zero, rec, exp)
		assertNotEquals(t, diff, expMsg, expDiff)
	})

	t.Run("unicode accent difference", func(t *testing.T) {
		a := "coração"
		b := "coracao"
		diff := diffStrings(zero, a, b, strat)
		expMsg := differs.Strings(zero, 4)
		rec := fmx.Sprintf("'%s'", markDiff(a, 4))
		exp := fmx.Sprintf("'%s'", b)
		expDiff := differs.Diff(zero, rec, exp)
		assertNotEquals(t, diff, expMsg, expDiff)
	})

	t.Run("middle character replaced", func(t *testing.T) {
		a := "lorem ipsum"
		b := "lorem xpsum"
		diff := diffStrings(zero, a, b, strat)
		recMsg := differs.Strings(zero, 6)
		rec := fmx.Sprintf("'%s'", markDiff(a, 6)[1:])
		exp := fmx.Sprintf("'%s'", b[1:])
		expDiff := differs.Diff(zero, rec, exp)
		assertNotEquals(t, diff, recMsg, expDiff)
	})

	t.Run("completely different", func(t *testing.T) {
		a := "abc"
		b := "xyz"
		diff := diffStrings(zero, a, b, strat)
		msg := differs.Strings(zero, 0)
		rec := fmx.Sprintf("'%s'", markDiff(a, 0))
		exp := fmx.Sprintf("'%s'", b)
		expDiff := differs.Diff(zero, rec, exp)
		assertNotEquals(t, diff, msg, expDiff)
	})
}

// TestDifferStrings_Default tests the method diffStrings based
// on the default strategy
//
//	this test should assert:
//	1. a string should be equals if:
//	- both have the exact same chars at the same exact positions; and
//	- exact same length
func TestDifferStrings_Default_FormatOpts(t *testing.T) {
	_, _ = standards.DifferStartIndex, standards.DifferMaxSize
	var zero indent.Tab

	t.Run("IgnoreWhitespace only", func(t *testing.T) {
		strat := NewStrategy(compare.IgnoreWhitespace)
		a := "hello   world"
		b := "hello world"
		diff := diffStrings(zero, a, b, strat)
		assertEquals(t, diff)
	})

	t.Run("TrimSpace only", func(t *testing.T) {
		strat := NewStrategy(compare.TrimSpace)
		t.Run("should be equals", func(t *testing.T) {
			a := "   hello world   "
			b := "hello world"
			diff := diffStrings(zero, a, b, strat)
			assertEquals(t, diff)
		})

		t.Run("shouldn't be equals", func(t *testing.T) {
			a := "   helloworld   "
			b := "hello world"
			diff := diffStrings(zero, a, b, strat)
			expMsg := differs.Strings(zero, 5)
			rec := fmx.Sprintf("'%s'", markDiff(strings.TrimSpace(a), 5))
			exp := fmx.Sprintf("'%s'", strings.TrimSpace(b))
			expDiff := differs.Diff(zero, rec, exp)
			assertNotEquals(t, diff, expMsg, expDiff)
		})
	})

	t.Run("IgnoreAccents only", func(t *testing.T) {
		strat := NewStrategy(compare.IgnoreAccents)
		a := "coração"
		b := "coracao"
		diff := diffStrings(zero, a, b, strat)
		assertEquals(t, diff)
	})

	t.Run("IgnoreWhitespace + TrimSpace", func(t *testing.T) {
		strat := NewStrategy(compare.IgnoreWhitespace, compare.TrimSpace)
		a := "   hello     world   "
		b := "hello world"
		diff := diffStrings(zero, a, b, strat)
		assertEquals(t, diff)
	})

	t.Run("IgnoreWhitespace + IgnoreAccents", func(t *testing.T) {
		strat := NewStrategy(compare.IgnoreWhitespace, compare.IgnoreAccents)
		a := "olá   mundo"
		b := "ola mundo"
		diff := diffStrings(zero, a, b, strat)
		assertEquals(t, diff)
	})

	t.Run("TrimSpace + IgnoreAccents", func(t *testing.T) {
		strat := NewStrategy(compare.TrimSpace, compare.IgnoreAccents)
		a := "  coração "
		b := "coracao"
		diff := diffStrings(zero, a, b, strat)
		assertEquals(t, diff)
	})

	t.Run("All format options", func(t *testing.T) {
		strat := NewStrategy(compare.IgnoreWhitespace, compare.TrimSpace, compare.IgnoreAccents)
		a := "  olá     coração  "
		b := "ola coracao"
		diff := diffStrings(zero, a, b, strat)
		assertEquals(t, diff)
	})

	t.Run("Difference remains after all format options", func(t *testing.T) {
		strat := NewStrategy(compare.IgnoreWhitespace, compare.TrimSpace, compare.IgnoreAccents)
		a := "  olá     coração  "
		b := "ola vida"
		diff := diffStrings(zero, a, b, strat)
		expMsg := differs.Strings(zero, 4)
		rd := fmx.Sprintf("'%s'", "ola "+fmx.SRed("c")+"oracao")
		ed := fmx.Sprintf("'%s'", "ola vida")
		expDiff := differs.Diff(zero, rd, ed)
		assertNotEquals(t, diff, expMsg, expDiff)
	})
}

// TestDifferStrings_IgnoreCase tests the method diffStrings ignoring cases
//
//	this test should assert:
//	1. a string should be equals if:
//	- both have the exact same words at the same exact positions; and
//	- exact same length
func TestDifferStrings_IgnoreCase(t *testing.T) {
	_, _ = standards.DifferStartIndex, standards.DifferMaxSize
	var zero indent.Tab
	strat := NewStrategy(compare.IgnoreCase)

	t.Run("equals strings, same case", func(t *testing.T) {
		a := "GoLang Test"
		b := "GoLang Test"
		diff := diffStrings(zero, a, b, strat)
		assertEquals(t, diff)
	})

	t.Run("equals strings, different case", func(t *testing.T) {
		a := "GoLang Test"
		b := "golang test"
		diff := diffStrings(zero, a, b, strat)
		assertEquals(t, diff)
	})

	t.Run("different strings, same case", func(t *testing.T) {
		a := "GoLang Test"
		b := "GoLang Text"
		diff := diffStrings(zero, a, b, strat)
		expMsg := differs.Strings(zero, 9)
		rd := "'ng te" + fmx.SRed("s") + "t'"
		ed := "'ng text'"
		expDiff := differs.Diff(zero, rd, ed)
		assertNotEquals(t, diff, expMsg, expDiff)
	})

	t.Run("different strings, only 1 character case-sensitive match", func(t *testing.T) {
		a := "A"
		b := "a"
		diff := diffStrings(zero, a, b, strat)
		assertEquals(t, diff)
	})

	t.Run("space difference with case ignored", func(t *testing.T) {
		a := "HelloWorld"
		b := "hello world"
		diff := diffStrings(zero, a, b, strat)
		expMsg := differs.Strings(zero, 5)
		rd := "'hello" + fmx.SRed("w") + "orld'"
		ed := "'hello world'"
		expDiff := differs.Diff(zero, rd, ed)
		assertNotEquals(t, diff, expMsg, expDiff)
	})

	t.Run("different length with case ignored", func(t *testing.T) {
		a := "GOLANG"
		b := "golang!"
		diff := diffStrings(zero, a, b, strat)
		expMsg := differs.Strings(zero, 6)
		rd := "'olang'"
		ed := "'olang!'"
		expDiff := differs.Diff(zero, rd, ed)
		assertNotEquals(t, diff, expMsg, expDiff)
	})
}

// TestDifferStrings_IgnoreCase_FormatOpts tests the method diffStrings ignoring cases
//
//	this test should assert:
//	1. a string should be equals if:
//	- both have the exact same words at the same exact positions; and
//	- exact same length
func TestDifferStrings_IgnoreCase_FormatOpts(t *testing.T) {
	_, _ = standards.DifferStartIndex, standards.DifferMaxSize
	var zero indent.Tab

	t.Run("IgnoreCase only", func(t *testing.T) {
		strat := NewStrategy(compare.IgnoreCase)
		a := "Hello World"
		b := "hello world"
		diff := diffStrings(zero, a, b, strat)
		assertEquals(t, diff)
	})

	t.Run("IgnoreCase + IgnoreWhitespace", func(t *testing.T) {
		strat := NewStrategy(compare.IgnoreCase, compare.IgnoreWhitespace)
		a := "Hello   World"
		b := "hello world"
		diff := diffStrings(zero, a, b, strat)
		assertEquals(t, diff)
	})

	t.Run("IgnoreCase + TrimSpace", func(t *testing.T) {
		strat := NewStrategy(compare.IgnoreCase, compare.TrimSpace)
		a := "   Hello World   "
		b := "hello world"
		diff := diffStrings(zero, a, b, strat)
		assertEquals(t, diff)
	})

	t.Run("IgnoreCase + IgnoreAccents", func(t *testing.T) {
		strat := NewStrategy(compare.IgnoreCase, compare.IgnoreAccents)
		a := "CORAÇÃO"
		b := "coracao"
		diff := diffStrings(zero, a, b, strat)
		assertEquals(t, diff)
	})

	t.Run("IgnoreCase + IgnoreWhitespace + TrimSpace", func(t *testing.T) {
		strat := NewStrategy(compare.IgnoreCase, compare.IgnoreWhitespace, compare.TrimSpace)
		a := "   HELLO     WORLD   "
		b := "hello world"
		diff := diffStrings(zero, a, b, strat)
		assertEquals(t, diff)
	})

	t.Run("IgnoreCase + IgnoreWhitespace + IgnoreAccents", func(t *testing.T) {
		strat := NewStrategy(compare.IgnoreCase, compare.IgnoreWhitespace, compare.IgnoreAccents)
		a := "OLÁ   MUNDO"
		b := "ola mundo"
		diff := diffStrings(zero, a, b, strat)
		assertEquals(t, diff)
	})

	t.Run("IgnoreCase + TrimSpace + IgnoreAccents", func(t *testing.T) {
		strat := NewStrategy(compare.IgnoreCase, compare.TrimSpace, compare.IgnoreAccents)
		a := "  CORAÇÃO "
		b := "coracao"
		diff := diffStrings(zero, a, b, strat)
		assertEquals(t, diff)
	})

	t.Run("IgnoreCase + All format options", func(t *testing.T) {
		strat := NewStrategy(compare.IgnoreCase, compare.IgnoreWhitespace, compare.TrimSpace, compare.IgnoreAccents)
		a := "  OLÁ     CORAÇÃO  "
		b := "ola coracao"
		diff := diffStrings(zero, a, b, strat)
		assertEquals(t, diff)
	})

	t.Run("IgnoreCase + All format options + real difference", func(t *testing.T) {
		strat := NewStrategy(compare.IgnoreCase, compare.IgnoreWhitespace, compare.TrimSpace, compare.IgnoreAccents)
		a := "  OLÁ     CORAÇÃO  "
		b := "ola vida"
		diff := diffStrings(zero, a, b, strat)
		expMsg := differs.Strings(zero, 4)
		rd := "'ola " + fmx.SRed("c") + "oracao'"
		ed := "'ola vida'"
		expDiff := differs.Diff(zero, rd, ed)
		assertNotEquals(t, diff, expMsg, expDiff)
	})
}

// TestDefaultSliceDiff tests the method defaultSliceDiff
//
//	this test must assert:
//	1. two slices are equals if:
//	A: Both slices must have the same type; and
//	A: Both slices must have the same length and capacity
//	A: Both slices have equals elements; or
//	B: Both slices have the same type, length, capacity and backing array pointer; or
//	C: Both slices are nil and have the same type.
func TestDefaultSliceDiff(t *testing.T) {
	var zero indent.Branch
	s := DefaultStrat()

	t.Run("same slice", func(t *testing.T) {
		x := []int{1, 2, 3}
		a := reflect.ValueOf(x)
		b := reflect.ValueOf(x)
		diff := defaultSliceDiff(zero, a, b, s)
		assertEquals(t, diff)
	})

	t.Run("same type, same length, same capacity, equals elements", func(t *testing.T) {
		a := reflect.ValueOf([]int{1, 2, 3})
		b := reflect.ValueOf([]int{1, 2, 3})
		diff := defaultSliceDiff(zero, a, b, s)
		assertEquals(t, diff)
	})

	t.Run("same type, same length, different capacity, equals elements", func(t *testing.T) {
		a := reflect.ValueOf(slices.Of[[]int](3, 1, 2, 3))
		b := reflect.ValueOf(slices.Of[[]int](4, 1, 2, 3))
		diff := defaultSliceDiff(zero, a, b, s)
		expMsg := differs.SliceCapMismatch(zero, a.Cap(), b.Cap())
		assertNotEquals(t, diff, expMsg, "")
	})

	t.Run("same type, same length, same capacity, not equals elements", func(t *testing.T) {
		a := reflect.ValueOf(slices.Of[[]int](3, 1, 2, 3))
		b := reflect.ValueOf(slices.Of[[]int](3, 1, 2, 4))
		diff := defaultSliceDiff(zero, a, b, s)
		expMsg := differs.SliceValues(zero, 2) + "\n" +
			differs.Values(zero.Inc(), sprints.TypedDigit("int", 3), sprints.TypedDigit("int", 4))
		assertNotEquals(t, diff, expMsg, "")
	})

	t.Run("same type, same capacity, different length", func(t *testing.T) {
		a := reflect.ValueOf(slices.Of[[]int](3, 1, 2, 3))
		b := reflect.ValueOf(slices.Of[[]int](3, 1, 2))
		diff := defaultSliceDiff(zero, a, b, s)
		expMsg := differs.SliceLenMismatch(zero, a.Len(), b.Len())
		assertNotEquals(t, diff, expMsg, "")
	})

	t.Run("different types, same length, capacity and elements", func(t *testing.T) {
		a := reflect.ValueOf(slices.Of[[]int](3, 1, 2, 3))
		b := reflect.ValueOf(slices.Of[[]int8](4, 1, 2, 3))
		diff := defaultSliceDiff(zero, a, b, s)
		expMsg := differs.SliceTypesMismatch(zero, types.ValidValueName(a.Type()),
			types.ValidValueName(b.Type()))
		assertNotEquals(t, diff, expMsg, "")
	})
}

// TestDefaultSliceDiff_EdgeCases tests the method defaultSliceDiff edge cases
//
//	this test must assert:
//	1. two slices are equals if:
//	A: Both slices must have the same type; and
//	A: Both slices must have the same length and capacity
//	A: Both slices have equals elements; or
//	B: Both slices have the same type, length, capacity and backing array pointer; or
//	C: Both slices are nil and have the same type.
func TestDefaultSliceDiff_EdgeCases(t *testing.T) {
	var zero indent.Branch
	s := DefaultStrat()

	t.Run("same type, nil", func(t *testing.T) {
		var x []int = nil
		a := reflect.ValueOf(x)
		b := reflect.ValueOf(x)
		diff := defaultSliceDiff(zero, a, b, s)
		assertEquals(t, diff)
	})

	t.Run("different type, nil", func(t *testing.T) {
		var s1 []int = nil
		var s2 []any = nil
		a := reflect.ValueOf(s1)
		b := reflect.ValueOf(s2)
		diff := defaultSliceDiff(zero, a, b, s)
		expMsg := differs.NilTypesMismatch(zero, "slices", types.ValidValueName(a.Type()),
			types.ValidValueName(b.Type()))
		assertNotEquals(t, diff, expMsg, "")
	})

	t.Run("one is nil, other is valid", func(t *testing.T) {
		var null []int = nil
		var notNull = []int{1, 2, 3}
		var empty = []int{}
		t.Run("nil received, non nil expected", func(t *testing.T) {
			a := reflect.ValueOf(null)
			b := reflect.ValueOf(notNull)
			diff := defaultSliceDiff(zero, a, b, s)
			expMsg := differs.NilReceived(zero, b.Kind().String(), types.ValidValueName(a.Type()))
			assertNotEquals(t, diff, expMsg, "")
		})
		t.Run("nil received, empty expected", func(t *testing.T) {
			a := reflect.ValueOf(null)
			b := reflect.ValueOf(empty)
			diff := defaultSliceDiff(zero, a, b, s)
			expMsg := differs.NilReceived(zero, b.Kind().String(), types.ValidValueName(a.Type()))
			assertNotEquals(t, diff, expMsg, "")
		})
		t.Run("nil expected, non nil received", func(t *testing.T) {
			a := reflect.ValueOf(notNull)
			b := reflect.ValueOf(null)
			diff := defaultSliceDiff(zero, a, b, s)
			expMsg := differs.NilExpected(zero, a.Kind().String(), types.ValidValueName(a.Type()))
			assertNotEquals(t, diff, expMsg, "")
		})
		t.Run("nil expected, empty received", func(t *testing.T) {
			a := reflect.ValueOf(empty)
			b := reflect.ValueOf(null)
			diff := defaultSliceDiff(zero, a, b, s)
			expMsg := differs.NilExpected(zero, a.Kind().String(), types.ValidValueName(a.Type()))
			assertNotEquals(t, diff, expMsg, "")
		})
	})
}

// TestDefaultNonNilSliceDiff tests the method defaultNotNilSliceDiff
//
//	this test must assert:
//	1. two slices are equals if:
//	A: Both slices must have the same type; and
//	A: Both slices must have the same length and capacity
//	A: Both slices have equals elements; or
//	B: Both slices have the same type, length, capacity and backing array pointer; or
//	C: Both slices are nil and have the same type.
func TestDefaultNonNilSliceDiff(t *testing.T) {
	var zero indent.Branch
	s := DefaultStrat()

	t.Run("same slice", func(t *testing.T) {
		x := []int{1, 2, 3}
		a := reflect.ValueOf(x)
		b := reflect.ValueOf(x)
		diff := defaultSliceDiff(zero, a, b, s)
		assertEquals(t, diff)
	})

	t.Run("same type, same length, same capacity, equals elements", func(t *testing.T) {
		a := reflect.ValueOf([]int{1, 2, 3})
		b := reflect.ValueOf([]int{1, 2, 3})
		diff := defaultSliceDiff(zero, a, b, s)
		assertEquals(t, diff)
	})

	t.Run("same type, same length, different capacity, equals elements", func(t *testing.T) {
		a := reflect.ValueOf(slices.Of[[]int](3, 1, 2, 3))
		b := reflect.ValueOf(slices.Of[[]int](4, 1, 2, 3))
		diff := defaultSliceDiff(zero, a, b, s)
		expMsg := differs.SliceCapMismatch(zero, a.Cap(), b.Cap())
		assertNotEquals(t, diff, expMsg, "")
	})

	t.Run("same type, same length, same capacity, not equals elements", func(t *testing.T) {
		a := reflect.ValueOf(slices.Of[[]int](3, 1, 2, 3))
		b := reflect.ValueOf(slices.Of[[]int](3, 1, 2, 4))
		diff := defaultSliceDiff(zero, a, b, s)
		expMsg := differs.SliceValues(zero, 2) + "\n" +
			differs.Values(zero.Inc(), sprints.TypedDigit("int", 3), sprints.TypedDigit("int", 4))
		assertNotEquals(t, diff, expMsg, "")
	})

	t.Run("same type, same capacity, different length", func(t *testing.T) {
		a := reflect.ValueOf(slices.Of[[]int](3, 1, 2, 3))
		b := reflect.ValueOf(slices.Of[[]int](3, 1, 2))
		diff := defaultSliceDiff(zero, a, b, s)
		expMsg := differs.SliceLenMismatch(zero, a.Len(), b.Len())
		assertNotEquals(t, diff, expMsg, "")
	})

	t.Run("different types, same length, capacity and elements", func(t *testing.T) {
		a := reflect.ValueOf(slices.Of[[]int](3, 1, 2, 3))
		b := reflect.ValueOf(slices.Of[[]int8](4, 1, 2, 3))
		diff := defaultSliceDiff(zero, a, b, s)
		expMsg := differs.SliceTypesMismatch(zero, types.ValidValueName(a.Type()),
			types.ValidValueName(b.Type()))
		assertNotEquals(t, diff, expMsg, "")
	})
}

// Tests the method sliceElemDiff
func TestSliceElemDiff(t *testing.T) {
	var zero indent.Branch
	s := DefaultStrat()

	t.Run("same slice", func(t *testing.T) {
		x := []int{1, 2, 3}
		a := reflect.ValueOf(x)
		b := reflect.ValueOf(x)
		diff := sliceElemDiff(zero, a, b, s)
		assertEquals(t, diff)
	})

	t.Run("same type, same length, equals elements", func(t *testing.T) {
		a := reflect.ValueOf([]int{1, 2, 3})
		b := reflect.ValueOf([]int{1, 2, 3})
		diff := sliceElemDiff(zero, a, b, s)
		assertEquals(t, diff)
	})

	t.Run("same type, empty slices", func(t *testing.T) {
		a := reflect.ValueOf([]int{})
		b := reflect.ValueOf([]int{})
		diff := sliceElemDiff(zero, a, b, s)
		assertEquals(t, diff)
	})

	t.Run("same type, same length, same capacity, not equals elements", func(t *testing.T) {
		a := reflect.ValueOf(slices.Of[[]int](3, 1, 2, 3))
		b := reflect.ValueOf(slices.Of[[]int](3, 1, 2, 4))
		diff := defaultSliceDiff(zero, a, b, s)
		expMsg := differs.SliceValues(zero, 2) + "\n" +
			differs.Values(zero.Inc(), sprints.TypedDigit("int", 3), sprints.TypedDigit("int", 4))
		assertNotEquals(t, diff, expMsg, "")
	})
}

// TestSerializableSliceDiff tests the method serializableSliceDiff
//
//	this test must assert:
//	1. two slices are equals if:
//	A: Both slices must have the same type; and
//	A: Both slices must have the same length
//	A: Both slices have equals elements; or
//	B: Both slices have the same type, length, and backing array pointer; or
//	C: Both slices are nil or empty and have the same type.
func TestSerializableSliceDiff(t *testing.T) {
	var zero indent.Branch
	s := DefaultStrat()

	t.Run("same slice", func(t *testing.T) {
		x := []int{1, 2, 3}
		a := reflect.ValueOf(x)
		b := reflect.ValueOf(x)
		diff := serializableSliceDiff(zero, a, b, s)
		assertEquals(t, diff)
	})

	t.Run("same type, same length, same capacity, equals elements", func(t *testing.T) {
		a := reflect.ValueOf([]int{1, 2, 3})
		b := reflect.ValueOf([]int{1, 2, 3})
		diff := serializableSliceDiff(zero, a, b, s)
		assertEquals(t, diff)
	})

	// capacity doesn't matter on serializable iterables
	t.Run("same type, same length, different capacity, equals elements", func(t *testing.T) {
		a := reflect.ValueOf(slices.Of[[]int](3, 1, 2, 3))
		b := reflect.ValueOf(slices.Of[[]int](4, 1, 2, 3))
		diff := serializableSliceDiff(zero, a, b, s)
		assertEquals(t, diff)
	})

	t.Run("same type, same length, same capacity, not equals elements", func(t *testing.T) {
		a := reflect.ValueOf(slices.Of[[]int](3, 1, 2, 3))
		b := reflect.ValueOf(slices.Of[[]int](3, 1, 2, 4))
		diff := serializableSliceDiff(zero, a, b, s)
		expMsg := differs.SliceValues(zero, 2) + "\n" +
			differs.Values(zero.Inc(), sprints.TypedDigit("int", 3), sprints.TypedDigit("int", 4))
		assertNotEquals(t, diff, expMsg, "")
	})

	t.Run("same type, same capacity, different length", func(t *testing.T) {
		a := reflect.ValueOf(slices.Of[[]int](3, 1, 2, 3))
		b := reflect.ValueOf(slices.Of[[]int](3, 1, 2))
		diff := serializableSliceDiff(zero, a, b, s)
		expMsg := differs.SliceLenMismatch(zero, a.Len(), b.Len())
		assertNotEquals(t, diff, expMsg, "")
	})

	t.Run("different types, same length, capacity and elements", func(t *testing.T) {
		a := reflect.ValueOf(slices.Of[[]int](3, 1, 2, 3))
		b := reflect.ValueOf(slices.Of[[]int8](4, 1, 2, 3))
		diff := serializableSliceDiff(zero, a, b, s)
		expMsg := differs.SliceTypesMismatch(zero, types.ValidValueName(a.Type()),
			types.ValidValueName(b.Type()))
		assertNotEquals(t, diff, expMsg, "")
	})
}

// TestSerializableSliceDiff_EdgeCases tests the method serializableSliceDiff edge cases
//
//	this test must assert:
//	1. two slices are equals if:
//	A: Both slices must have the same type; and
//	A: Both slices must have the same length
//	A: Both slices have equals elements; or
//	B: Both slices have the same type, length, and backing array pointer; or
//	C: Both slices are nil or empty and have the same type.
func TestSerializableSliceDiff_EdgeCases(t *testing.T) {
	var zero indent.Branch
	s := DefaultStrat()

	t.Run("same type, nil", func(t *testing.T) {
		var x []int = nil
		a := reflect.ValueOf(x)
		b := reflect.ValueOf(x)
		diff := serializableSliceDiff(zero, a, b, s)
		assertEquals(t, diff)
	})

	t.Run("different type, nil", func(t *testing.T) {
		var s1 []int = nil
		var s2 []any = nil
		a := reflect.ValueOf(s1)
		b := reflect.ValueOf(s2)
		diff := serializableSliceDiff(zero, a, b, s)
		expMsg := differs.NilTypesMismatch(zero, "slices", types.ValidValueName(a.Type()),
			types.ValidValueName(b.Type()))
		assertNotEquals(t, diff, expMsg, "")
	})

	t.Run("one is nil, other is valid", func(t *testing.T) {
		var null []int = nil
		var notNull = []int{1, 2, 3}
		var empty = []int{}
		t.Run("nil received, not nil expected", func(t *testing.T) {
			a := reflect.ValueOf(null)
			b := reflect.ValueOf(notNull)
			diff := serializableSliceDiff(zero, a, b, s)
			expMsg := differs.NilReceived(zero, b.Kind().String(), types.ValidValueName(a.Type()))
			assertNotEquals(t, diff, expMsg, "")
		})
		t.Run("nil received, empty expected", func(t *testing.T) {
			a := reflect.ValueOf(null)
			b := reflect.ValueOf(empty)
			diff := serializableSliceDiff(zero, a, b, s)
			assertEquals(t, diff)
		})
		t.Run("not nil received nil expected", func(t *testing.T) {
			a := reflect.ValueOf(notNull)
			b := reflect.ValueOf(null)
			diff := serializableSliceDiff(zero, a, b, s)
			expMsg := differs.NilExpected(zero, a.Kind().String(), types.ValidValueName(a.Type()))
			assertNotEquals(t, diff, expMsg, "")
		})
		t.Run("emoty received, nil expected", func(t *testing.T) {
			a := reflect.ValueOf(empty)
			b := reflect.ValueOf(null)
			diff := serializableSliceDiff(zero, a, b, s)
			assertEquals(t, diff)
		})
	})
}

// TestSerializableNotNilSliceDiff tests the method serializableSliceDiff
//
//	this test must assert:
//	1. two slices are equals if:
//	A: Both slices must have the same type; and
//	A: Both slices must have the same length
//	A: Both slices have equals elements; or
//	B: Both slices have the same type, length, and backing array pointer
func TestSerializableNotNilSliceDiff(t *testing.T) {
	var zero indent.Branch
	s := DefaultStrat()

	t.Run("same slice", func(t *testing.T) {
		x := []int{1, 2, 3}
		a := reflect.ValueOf(x)
		b := reflect.ValueOf(x)
		diff := serializableNotNilSliceDiff(zero, a, b, s)
		assertEquals(t, diff)
	})

	t.Run("same type, same length, same capacity, equals elements", func(t *testing.T) {
		a := reflect.ValueOf([]int{1, 2, 3})
		b := reflect.ValueOf([]int{1, 2, 3})
		diff := serializableNotNilSliceDiff(zero, a, b, s)
		assertEquals(t, diff)
	})

	// capacity doesn't matter on serializable iterables
	t.Run("same type, same length, different capacity, equals elements", func(t *testing.T) {
		a := reflect.ValueOf(slices.Of[[]int](3, 1, 2, 3))
		b := reflect.ValueOf(slices.Of[[]int](4, 1, 2, 3))
		diff := serializableNotNilSliceDiff(zero, a, b, s)
		assertEquals(t, diff)
	})

	t.Run("same type, same length, same capacity, not equals elements", func(t *testing.T) {
		a := reflect.ValueOf(slices.Of[[]int](3, 1, 2, 3))
		b := reflect.ValueOf(slices.Of[[]int](3, 1, 2, 4))
		diff := serializableNotNilSliceDiff(zero, a, b, s)
		expMsg := differs.SliceValues(zero, 2) + "\n" +
			differs.Values(zero.Inc(), sprints.TypedDigit("int", 3), sprints.TypedDigit("int", 4))
		assertNotEquals(t, diff, expMsg, "")
	})

	t.Run("same type, same capacity, different length", func(t *testing.T) {
		a := reflect.ValueOf(slices.Of[[]int](3, 1, 2, 3))
		b := reflect.ValueOf(slices.Of[[]int](3, 1, 2))
		diff := serializableNotNilSliceDiff(zero, a, b, s)
		expMsg := differs.SliceLenMismatch(zero, a.Len(), b.Len())
		assertNotEquals(t, diff, expMsg, "")
	})

	t.Run("different types, same length, capacity and elements", func(t *testing.T) {
		a := reflect.ValueOf(slices.Of[[]int](3, 1, 2, 3))
		b := reflect.ValueOf(slices.Of[[]int8](4, 1, 2, 3))
		diff := serializableNotNilSliceDiff(zero, a, b, s)
		expMsg := differs.SliceTypesMismatch(zero, types.ValidValueName(a.Type()),
			types.ValidValueName(b.Type()))
		assertNotEquals(t, diff, expMsg, "")
	})
}

// TestDiffStructs_Default tests the method diffStructs
//
//	this test must assert:
//	1. two structs are equals if:
//	A: both structs have the same type; and
//	A: both struct fields are equals;
func TestDiffStructs_Default(t *testing.T) {
	var zero indent.Branch
	s := DefaultStrat()

	t.Run("same type, same fields", func(t *testing.T) {
		a := reflect.ValueOf(models.TwoDataAsValue("abc", 123))
		b := reflect.ValueOf(models.TwoDataAsValue("abc", 123))
		diff := diffStructs(zero, a, b, s)
		assertEquals(t, diff)
	})

	t.Run("same type, different fields", func(t *testing.T) {
		a := reflect.ValueOf(models.TwoDataAsValue("abc", 123))
		b := reflect.ValueOf(models.TwoDataAsValue("abc", 12))
		diff := diffStructs(zero, a, b, s)
		name := "github.com/andrerrcosta2/gtools/gtests/testingseeds/static/structs/models.TwoData"
		expMsg := differs.Append(zero, differs.StructFields(zero, name, "Age"),
			differs.Values(zero.Inc(), sprints.Typed("int", 123), sprints.Typed("int", 12)))
		assertNotEquals(t, diff, expMsg, "")
	})

	t.Run("different types", func(t *testing.T) {
		a := reflect.ValueOf(models.TwoDataAsValue("abc", 123))
		b := reflect.ValueOf(models.OneDataAsValue("abc"))
		diff := diffStructs(zero, a, b, s)
		expMsg := differs.StructTypesMismatch(zero, types.ValidValueName(a.Type()),
			types.ValidValueName(b.Type()))
		assertNotEquals(t, diff, expMsg, "")
	})
}

func TestDifferUint(t *testing.T) {
	runPrimTests(t, uintSeed, differUint)
}

// TestDifferUnsafePointers tests the method defaultUnsafePtrs
//
//	this test must assert:
//	1. two unsafe ptrs are equals if:
//	A: both ptrs hold the same address; or
//	B: both ptrs are nil
func TestDifferUnsafePointers(t *testing.T) {
	var zero indent.Branch
	s := DefaultStrat()

	t.Run("same address", func(t *testing.T) {
		ptr := ptrs.Unsafe(10)
		diff := defaultUnsafePtrs(zero, reflect.ValueOf(ptr), reflect.ValueOf(ptr), s)
		assertEquals(t, diff)
	})

	t.Run("both nil", func(t *testing.T) {
		var p1 unsafe.Pointer
		var p2 unsafe.Pointer
		diff := defaultUnsafePtrs(zero, reflect.ValueOf(p1), reflect.ValueOf(p2), s)
		assertEquals(t, diff)
	})

	t.Run("different addresses", func(t *testing.T) {
		ptr1 := ptrs.Unsafe(10)
		ptr2 := ptrs.Unsafe(10)
		diff := defaultUnsafePtrs(zero, reflect.ValueOf(ptr1), reflect.ValueOf(ptr2), s)
		expMsg := differs.UnsafePointersAddr(zero,
			sprints.Uintptrf(uintptr(ptr1)),
			sprints.Uintptrf(uintptr(ptr2)),
		)
		assertNotEquals(t, diff, expMsg, "")
	})

	t.Run("one nil, one non-nil", func(t *testing.T) {
		null := reflect.ValueOf(generics.Zero[unsafe.Pointer]())
		nonNull := reflect.ValueOf(ptrs.Unsafe(20))
		t.Run("nil received", func(t *testing.T) {
			diff := defaultUnsafePtrs(zero, null, nonNull, s)
			expMsg := differs.NilReceived(zero, "unsafe.pointer",
				sprints.Uintptrf(nonNull.Pointer()))
			assertNotEquals(t, diff, expMsg, "")
		})
		t.Run("nil expected", func(t *testing.T) {
			diff := defaultUnsafePtrs(zero, nonNull, null, s)
			expMsg := differs.NilExpected(zero, "unsafe.pointer",
				sprints.Uintptrf(nonNull.Pointer()))
			assertNotEquals(t, diff, expMsg, "")
		})
	})
}
