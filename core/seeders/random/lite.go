// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package random

import (
	"fmt"
	"math"
	"math/big"
	"reflect"
	"time"

	"github.com/andrerrcosta2/gtools/core/data/str/iterables"
	"github.com/andrerrcosta2/gtools/core/domain/constraints/prim/nums/ints"
	"github.com/andrerrcosta2/gtools/core/domain/constraints/prim/nums/uints"
	"github.com/andrerrcosta2/gtools/core/domain/gerrors"
	"github.com/andrerrcosta2/gtools/core/format/fmx"
	"github.com/andrerrcosta2/gtools/core/seeders/random/internal/cat"
	"github.com/andrerrcosta2/gtools/core/seeders/random/internal/lite"
	"github.com/andrerrcosta2/gtools/core/seeders/random/internal/prng"
	"github.com/andrerrcosta2/gtools/core/seeders/random/internal/reflectutils/reflectrand"
	"github.com/andrerrcosta2/gtools/core/util/casters"
	"github.com/andrerrcosta2/gtools/core/util/typeutil/charsets"
)

const (
	ImpossibleConstraint = "IMPOSSIBLE_CONSTRAINT"
)

const (
	MinStringSizeDefaults = 1
	MaxStringSizeDefaults = 512
	MinBytesSizeDefaults  = 1
	MaxBytesSizeDefaults  = 1024
)

// Alphabet returns a slice of length q with random strings of length between min and max.
// Each string is composed of characters randomly chosen from the alphabet.
// If min or max aren't provided, it defaults to a minimum length of 1 and a maximum length
// of 512 chars
// If q is negative, it returns an empty slice.
func Alphabet(q int, minMax ...int) *iterables.Slice[string] {
	if q <= 0 {
		return iterables.OfSlice[string]()
	}

	switch len(minMax) {
	case 0:
		minMax = []int{MinStringSizeDefaults, MaxStringSizeDefaults}
	case 1:
		minMax = []int{minMax[0], MaxStringSizeDefaults}
	default:
		break
	}

	result := make(iterables.Slice[string], q)

	// Generate the random strings
	Int(q, minMax...).EachN(func(i, l int) {
		result[i] = lite.RandString(l, charsets.Alphabet)
	})

	return &result
}

// Alphanumeric returns a slice of length q with random alphanumeric strings between min and max length.
// If min or max aren't provided, it defaults to a minimum length of 1 and a maximum length
// compare to the maximum length of a string.
// It returns empty if q is negative.
func Alphanumeric(q int, minMax ...int) *iterables.Slice[string] {
	if q <= 0 {
		return iterables.OfSlice[string]()
	}

	switch len(minMax) {
	case 0:
		minMax = []int{MinStringSizeDefaults, MaxStringSizeDefaults}
	case 1:
		minMax = []int{minMax[0], MaxStringSizeDefaults}
	default:
		break
	}

	result := make(iterables.Slice[string], q)

	// Generate the random strings
	Int(q, minMax...).EachN(func(i, l int) {
		result[i] = lite.RandString(l, charsets.AlphaNumeric)
	})

	return &result
}

// Any generates a slice of length q with random values of any type.
// It returns empty if q is negative.
//
// This method doesn't generate structs.
func Any(q int) *iterables.Slice[any] {
	if q <= 0 {
		return iterables.OfSlice[any]()
	}
	result := make(iterables.Slice[any], q)
	for i := 0; i < q; i++ {
		result[i] = lite.RandAny()
	}
	return &result
}

// ArrayOf returns an iterables.Slice of length 'q' with random array values.
// It returns empty if 'q' is negative or zero.
func ArrayOf[T any](q int) *iterables.Slice[T] {
	if q <= 0 {
		return iterables.OfSlice[T]()
	}
	var zero T
	t := reflect.TypeOf(zero)
	if t.Kind() != reflect.Array {
		panic(fmx.Sprintf("random:ArrayOf[%T] expected an array, got "+
			"a %v: '%s'", zero, t.Kind(), t.String()))
	}
	result := make(iterables.Slice[T], q)
	for i := 0; i < q; i++ {
		result[i] = lite.RandArrayOf(t).(T)
	}
	return &result
}

// Bool returns an iterables.Slice of length q with random boolean values.
// It returns empty if q is negative.
func Bool(q int) *iterables.Slice[bool] {
	if q <= 0 {
		return iterables.OfSlice[bool]()
	}
	result := make(iterables.Slice[bool], q)

	for i := range result {
		result[i] = prng.Bool()
	}
	return &result
}

// Bytes return a slice of length q with random byte values.
// It returns empty if q is negative.
func Bytes(q int, minMax ...int) *iterables.Slice[[]byte] {
	if q <= 0 {
		return iterables.OfSlice[[]byte]()
	}

	var m, M int
	switch len(minMax) {
	case 0:
		m, M = MinBytesSizeDefaults, MaxBytesSizeDefaults
	case 1:
		m, M = minMax[0], MaxBytesSizeDefaults
	default:
		m, M = minMax[0], minMax[1]
	}
	if m > M {
		m, M = M, m
	}

	result := make(iterables.Slice[[]byte], q)
	// Generate random bytes from 0 to 255
	for i := 0; i < q; i++ {
		result[i] = prng.Bytes(m, M)
	}
	return &result
}

// Comparable generates a slice of length q with random comparable values.
func Comparable(q int) *iterables.Slice[any] {
	if q <= 0 {
		return iterables.OfSlice[any]()
	}
	result := make(iterables.Slice[any], q)
	for i := 0; i < q; i++ {
		result[i] = lite.RandCmp()
	}
	return &result
}

func Complex64(q int, realMinMaxImagMinMax ...float32) *iterables.Slice[complex64] {
	if q <= 0 {
		return iterables.OfSlice[complex64]()
	}
	result := make(iterables.Slice[complex64], q)
	minMax := realMinMaxImagMinMax
	var rm, rM, im, iM float32
	switch len(realMinMaxImagMinMax) {
	case 0:
		rm, rM, im, iM = -math.MaxFloat32, math.MaxFloat32, -math.MaxFloat32, math.MaxFloat32
	case 1:
		rm, rM, im, iM = minMax[0], math.MaxFloat32, -math.MaxFloat32, math.MaxFloat32
	case 2:
		rm, rM, im, iM = minMax[0], minMax[1], -math.MaxFloat32, math.MaxFloat32
	case 3:
		rm, rM, im, iM = minMax[0], minMax[1], minMax[2], math.MaxFloat32
	default:
		rm, rM, im, iM = minMax[0], minMax[1], minMax[2], minMax[3]
	}

	if rm > rM {
		rm, rM = rM, rm
	}
	if im > iM {
		im, iM = iM, im
	}
	for i := range result {
		result[i] = prng.Complex64(rm, rM, im, iM)
	}
	return &result
}

func Complex128(q int, realMinMaxImagMinMax ...float64) *iterables.Slice[complex128] {
	if q <= 0 {
		return iterables.OfSlice[complex128]()
	}
	result := make(iterables.Slice[complex128], q)

	minMax := realMinMaxImagMinMax
	var rm, rM, im, iM float64
	switch len(realMinMaxImagMinMax) {
	case 0:
		rm, rM, im, iM = -math.MaxFloat64, math.MaxFloat64, -math.MaxFloat64, math.MaxFloat64
	case 1:
		rm, rM, im, iM = minMax[0], math.MaxFloat64, -math.MaxFloat64, math.MaxFloat64
	case 2:
		rm, rM, im, iM = minMax[0], minMax[1], -math.MaxFloat64, math.MaxFloat64
	case 3:
		rm, rM, im, iM = minMax[0], minMax[1], minMax[2], math.MaxFloat64
	default:
		rm, rM, im, iM = minMax[0], minMax[1], minMax[2], minMax[3]
	}
	if rm > rM {
		rm, rM = rM, rm
	}
	if im > iM {
		im, iM = iM, im
	}

	for i := range result {
		result[i] = prng.Complex128(rm, rM, im, iM)
	}
	return &result
}

// Float32 generates a slice of length q with random float32 values between min and max.
// If min or max are not provided, it defaults to the minimum and maximum float32 values.
//   - For ranges near the maximum float32 value (MaxFloat32), the precision of float32
//     may cause a loss of variation in generated values. This is due to the inherent
//     limitations of the 23-bitwise mantissa in the IEEE 754 representation of float32,
//     where the distance between representable values grows with magnitude.
//
// Returns an empty slice if q is negative or zero.
func Float32(q int, minMax ...float32) *iterables.Slice[float32] {
	if q <= 0 {
		return iterables.OfSlice[float32]()
	}
	var m, M float32
	switch len(minMax) {
	case 0:
		m, M = -math.MaxFloat32, math.MaxFloat32
	case 1:
		m, M = minMax[0], math.MaxFloat32
	default:
		m, M = minMax[0], minMax[1]
	}
	result := make(iterables.Slice[float32], q)
	if m == M {
		for i := range result {
			result[i] = m
		}
	}
	if m > M {
		m, M = M, m
	}
	for i := range result {
		result[i] = prng.Float32(m, M)
	}
	return &result
}

// Float64 generates a slice of length q with random float64 values between min and max.
// If min or max are not provided, it defaults to the minimum and maximum float64 values.
//   - Float64 offers significantly higher precision compared to Float32, with a 52-bitwise
//     mantissa in the IEEE 754 representation. This allows for finer variation in
//     values even for large ranges near MaxFloat64.
//
// Returns an empty slice if q is negative or zero.
func Float64(q int, minMax ...float64) *iterables.Slice[float64] {
	if q <= 0 {
		return iterables.OfSlice[float64]()
	}

	var m, M float64
	switch len(minMax) {
	case 0:
		m, M = -math.MaxFloat64, math.MaxFloat64
	case 1:
		m, M = minMax[0], math.MaxFloat64
	default:
		m, M = minMax[0], minMax[1]
	}
	result := make(iterables.Slice[float64], q)
	if m == M {
		for i := range result {
			result[i] = m
		}
	}
	if m > M {
		m, M = M, m
	}
	for i := range result {
		result[i] = prng.Float64(m, M)
	}
	return &result
}

// Int generates a slice of length q with random int values between min and max.
// If min or max are not provided, it defaults to the minimum and maximum int values.
// Returns an empty slice if q is negative or zero.
func Int(q int, minMax ...int) *iterables.Slice[int] {
	if q <= 0 {
		return iterables.OfSlice[int]()
	}
	result := make(iterables.Slice[int], q)

	var m, M int
	switch len(minMax) {
	case 0:
		m, M = ints.Min, ints.Max
	case 1:
		m, M = minMax[0], ints.Max
	default:
		m, M = minMax[0], minMax[1]
	}
	if m == M {
		for i := range result {
			result[i] = m
		}
	}
	if m > M {
		m, M = M, m
	}

	for i := range result {
		result[i] = prng.Int(m, M)
	}
	return &result
}

// Int8 generates a slice of length q with random int8 values between min and max.
// If min or max are not provided, it defaults to the minimum and maximum int8 values.
// Returns an empty slice if q is negative or zero.
func Int8(q int, minMax ...int8) *iterables.Slice[int8] {
	if q <= 0 {
		return iterables.OfSlice[int8]()
	}
	result := make(iterables.Slice[int8], q)

	var m, M int8
	switch len(minMax) {
	case 0:
		m, M = ints.Min8, ints.Max8
	case 1:
		m, M = minMax[0], ints.Max8
	default:
		m, M = minMax[0], minMax[1]
	}
	if m == M {
		for i := range result {
			result[i] = m
		}
	}
	if m > M {
		m, M = M, m
	}

	for i := range result {
		result[i] = prng.Int8(m, M)
	}
	return &result
}

func Int16(q int, minMax ...int16) *iterables.Slice[int16] {
	if q <= 0 {
		return iterables.OfSlice[int16]()
	}
	result := make(iterables.Slice[int16], q)

	var m, M int16
	switch len(minMax) {
	case 0:
		m, M = ints.Min16, ints.Max16
	case 1:
		m, M = minMax[0], ints.Max16
	default:
		m, M = minMax[0], minMax[1]
	}
	if m == M {
		for i := range result {
			result[i] = m
		}
	}
	if m > M {
		m, M = M, m
	}

	for i := range result {
		result[i] = prng.Int16(m, M)
	}
	return &result
}

func Int32(q int, minMax ...int32) *iterables.Slice[int32] {
	if q <= 0 {
		return iterables.OfSlice[int32]()
	}
	result := make(iterables.Slice[int32], q)

	var m, M int32
	switch len(minMax) {
	case 0:
		m, M = ints.Min32, ints.Max32
	case 1:
		m, M = minMax[0], ints.Max32
	default:
		m, M = minMax[0], minMax[1]
	}
	if m == M {
		for i := range result {
			result[i] = m
		}
	}
	if m > M {
		m, M = M, m
	}

	for i := range result {
		result[i] = prng.Int32(m, M)
	}
	return &result
}

func Int64(q int, minMax ...int64) *iterables.Slice[int64] {
	if q <= 0 {
		return iterables.OfSlice[int64]()
	}
	result := make(iterables.Slice[int64], q)

	var m, M int64
	switch len(minMax) {
	case 0:
		m, M = ints.Min64, ints.Max64
	case 1:
		m, M = minMax[0], ints.Max64
	default:
		m, M = minMax[0], minMax[1]
	}
	if m == M {
		for i := range result {
			result[i] = m
		}
	}
	if m > M {
		m, M = M, m
	}

	for i := range result {
		result[i] = prng.Int64(m, M)
	}
	return &result
}

// Kind generates q random reflect.Kind
// It returns empty if q is negative.
func Kind(q int) *iterables.Slice[reflect.Kind] {
	if q <= 0 {
		return iterables.OfSlice[reflect.Kind]()
	}

	result := make(iterables.Slice[reflect.Kind], q)
	for i := 0; i < q; i++ {
		result[i] = reflectrand.Kind()
	}
	return &result
}

// Of generates a slice of length q with random values of type T.
// This method may use reflection to prototype complex types.
func Of[T any](q int) *iterables.Slice[T] {
	if q <= 0 {
		return iterables.OfSlice[T]()
	}
	t := reflect.TypeOf((*T)(nil)).Elem()
	result := make(iterables.Slice[T], q)
	// ⚠️ IMPORTANT CASTING STRATEGY NOTES (DO NOT REMOVE):
	//
	// The use of unsafe casts (UnsafeReferenceOf, UnsafeValueOf) is a TEMPORARY WORKAROUND
	// due to a limitation in lite.RandOf(t): it generates values based on KIND
	// rather than preserving the exact TYPE.
	// This breaks direct type assertions like rdn.(T) for named types (ex. type MyInt int), causing panics.
	//
	// - For cat.Any (T = interface{}): direct assignment via rdn.(T) is safe because
	//   every value implements the empty interface.
	//
	// - For cat.Struct: we assume lite.RandOf(t) correctly constructs a reflect.Value of type t
	//   (including nested fields), so rdn.(T) is safe—even with slices/maps—because
	//   reflect.Value.Interface() returns a properly typed value.
	//
	// - For cat.Reference (ptr, chan, map, func, unsafe.Pointer): we use UnsafeReferenceOf
	//   because RandOf returns a correctly typed *interface{}*, but the value inside is a pointer-like
	//   reference type.
	//  UnsafeReferenceOf extracts the pointer word directly.
	//   ⚠️ This assumes the runtime representation matches expectations.
	//	When the Runtime Representation Might Not Match:
	//	1. Go’s internal interface layout could change
	//		- the emptyInterface struct mimics Go’s internal representation of an empty interface.
	//		- This is not part of the Go spec—it’s an implementation detail of the compiler/runtime.
	//		- While it’s been stable for years (two words: type + data), a future Go version could change it (ex, add
	//		  metadata for generics, GC, or debugging).
	//		- If that happens, ei.word no longer points to the data → silent memory corruption or crashes.
	//	2. misusing it for non-pointer T
	//		- Of[map[string]int](5)  // T = map[string]int → reference type → cat.Reference
	//		```go```
	//		 rdn := lite.RandOf(t)        // returns a map[string]int (as interface{})
	//		  result[i] = UnsafeReferenceOf[map[string]int](rdn)
	//		  return *(*map[string]int)(unsafe.Pointer(&ei.word))
	//		``````
	//		ei.word is already the map value (which is a pointer internally).
	//		You’re taking the address of that pointer (&ei.word) and treating it as if it were the map itself.
	// 		This creates a pointer to a local copy of the map header, not the map itself.
	//		Result: The returned map[string]int is corrupted—it points to stack memory that becomes invalid
	//		after the function returns.
	//
	// - For cat.Value (scalars, arrays, named types like MyInt): we currently use UnsafeValueOf
	//   as a fallback because RandOf returns the UNDERLYING type instead of the NAMED type,
	//   making rdn.(T) panic.
	//  This unsafe cast bypasses type checking by reinterpreting raw memory.
	//   ❗ THIS IS FRAGILE AND UNSAFE FOR TYPES CONTAINING POINTERS (structs with slices/maps).
	//     It only "works" for simple, pointer-free named types (type Port uint16).
	//
	//  FUTURE: Update lite.RandOf(t) to always return a value of EXACT type t
	//   (using reflect.New(t).Elem().Set(...) and kind-based population).
	//   Once that’s done, ALL unsafe casts can be replaced with rdn.(T).
	//
	// Until then, this branching logic is necessary to avoid panics on named types
	// while maintaining compatibility with complex and reference types.
	switch cat.CastMethod(t) {
	case cat.Any:
		for i := 0; i < q; i++ {
			rdn := lite.RandAny()
			result[i] = rdn.(T)
		}
		break
	case cat.Injectable:
		for i := 0; i < q; i++ {
			var zero T
			result[i] = zero
		}
		break
	case cat.Reference:
		for i := 0; i < q; i++ {
			rdn := lite.RandOf(t)
			result[i] = rdn.(T)
			//result[i] = casters.UnsafeReferenceOf[T](rdn)
		}
		break
	case cat.Struct:
		for i := 0; i < q; i++ {
			rdn := lite.RandOf(t)
			result[i] = rdn.(T)
		}
	default:
		// Value types (including named scalars like MyInt)
		// UnsafeValueOf is used ONLY because RandOf returns underlying type, not T
		// ⚠️ Do NOT use for types containing pointers (slice, map, etc.)—memory layout will be corrupted!
		for i := 0; i < q; i++ {
			rdn := lite.RandOf(t)
			result[i] = casters.UnsafeValueOf[T](rdn)
		}
		break
	}
	return &result
}

// Reflect generates an iterables.Slice of 'q' reflect.Value from the given reflect.Type.
// It returns empty if 'q' is less than one.
func Reflect(t reflect.Type, q int) *iterables.Slice[reflect.Value] {
	if q <= 0 {
		return iterables.OfSlice[reflect.Value]()
	}

	result := make(iterables.Slice[reflect.Value], q)
	for i := 0; i < q; i++ {
		result[i] = reflectrand.ValueOf(t)
	}
	return &result
}

// Rune generates a slice of random Unicode runes that are valid and graphics of length q.
// It returns a slice of valid runes if q is non-negative.
func Rune(q int) *iterables.Slice[rune] {
	if q <= 0 {
		return iterables.OfSlice[rune]()
	}

	result := make(iterables.Slice[rune], q)
	// Unicode runes go up to 0x10FFFF (1,114,111 in decimal)
	for i := 0; i < q; i++ {
		// Generate a random rune in the full Unicode range (up to 0x10FFFF)
		result[i] = lite.RandRune()
	}
	return &result
}

// Single returns a single random value of the type passed as argument
func Single(value any) (single any) {
	return lite.RandOf(reflect.TypeOf(value))
}

// SingleOf generates a random value of type T
func SingleOf[T any]() (single T) {
	t := reflect.TypeOf((*T)(nil)).Elem()
	switch cat.CastMethod(t) {
	case cat.Any:
		return lite.RandAny().(T)
	case cat.Injectable:
		return single
	case cat.Reference, cat.Struct:
		return lite.RandOf(t).(T)
	default:
		return casters.UnsafeValueOf[T](lite.RandOf(t))
	}
}

// String returns a slice of length q with random strings of length between min and max.
// Each string is composed of random runes.
// If min or max are not provided, it defaults to a minimum length of 1 and a maximum length
//
//	compared to the maximum length of a string.
//
// It returns empty if q is negative.
func String(q int, minMax ...int) *iterables.Slice[string] {
	if q <= 0 {
		return iterables.OfSlice[string]()
	}

	switch len(minMax) {
	case 0:
		minMax = []int{MinStringSizeDefaults, MaxStringSizeDefaults}
	case 1:
		minMax = []int{minMax[0], MaxStringSizeDefaults}
	default:
		break
	}

	result := make(iterables.Slice[string], q)
	// Generate the random strings
	Int(q, minMax...).EachN(func(i, l int) {
		result[i] = string(lite.RandRune())
	})

	return &result
}

// StringOf returns a slice of length q with random strings of length between min and max.
// Each string is composed of characters randomly chosen from the charset.
// If min or max are not provided, it defaults to a minimum length of 1 and a maximum length
// compare to the maximum length of a string.
// It returns empty if q is negative.
func StringOf(q int, charset string, minMax ...int) *iterables.Slice[string] {
	if q <= 0 {
		return iterables.OfSlice[string]()
	}

	switch len(minMax) {
	case 0:
		minMax = []int{MinStringSizeDefaults, MaxStringSizeDefaults}
	case 1:
		minMax = []int{minMax[0], MaxStringSizeDefaults}
	default:
		break
	}

	result := make(iterables.Slice[string], q)

	if len(charset) == 0 {
		for i := 0; i < q; i++ {
			result[i] = ""
		}
		return &result
	}

	// Generate random string lengths between min and max
	Int(q, minMax...).EachN(func(i, l int) {
		result[i] = lite.RandString(l, charset)
	})

	return &result
}

// Struct populates a struct of type T with random field values (except for interfaces which requires a type)
func Struct[T any](q int) *iterables.Slice[T] {
	if q <= 0 {
		return iterables.OfSlice[T]()
	}

	result := make(iterables.Slice[T], q)
	// Get the type of the provided value
	var zero T
	iter := reflect.TypeOf(zero)
	// keep the outer type
	t := iter

	// Dereference the type until it's no longer a pointer or struct
	for iter.Kind() == reflect.Ptr {
		iter = iter.Elem()
	}

	if iter.Kind() != reflect.Struct {
		panic(fmt.Sprintf("gtools:random - Non struct kind: '%s' passed to random.Struct: %s\n", t.Kind(), t.String()))
	}

	for i := 0; i < q; i++ {
		// Again, a little bitwise overcautious to me.
		// But the output of this is an interface{}, so golang compiler has no
		// idea the type generated returns a T.
		// It might be worth duplicating the structOf function
		// to avoid type assertions, but this is a really negligible overhead.
		if s, ok := lite.RandOf(t).(T); ok {
			result[i] = s
		}
	}
	return &result
}

// Timestamp returns a slice of length q with random time.Time values between from and to.
// If from is after to, it swaps them.
// It returns empty if q is negative.
func Timestamp(q int, fromTo ...time.Time) *iterables.Slice[time.Time] {
	if q <= 0 {
		return iterables.OfSlice[time.Time]()
	}

	from, to := xtimergn(fromTo...)
	result := make(iterables.Slice[time.Time], q)
	diff := to.Sub(from)

	for i := 0; i < q; i++ {
		// Generate a random duration between 0 and diff
		result[i] = lite.RandTimestamp(from, diff)
	}
	return &result
}

// Uint returns a slice of q random uints.
// If q is negative, it returns an empty slice.
func Uint(q int, minMax ...uint) *iterables.Slice[uint] {
	if q <= 0 {
		return iterables.OfSlice[uint]()
	}
	result := make(iterables.Slice[uint], q)

	var m, M uint
	switch len(minMax) {
	case 0:
		m, M = 0, uints.Max
	case 1:
		m, M = minMax[0], uints.Max
	default:
		m, M = minMax[0], minMax[1]
	}
	if m == M {
		for i := range result {
			result[i] = prng.Uint(m, M)
		}
	}
	if m > M {
		m, M = M, m
	}

	for i := range result {
		result[i] = prng.Uint(m, M)
	}
	return &result
}

// Uint8 returns a slice of q random uint8s.
// If q is negative, it returns an empty slice.
func Uint8(q int, minMax ...uint8) *iterables.Slice[uint8] {
	if q <= 0 {
		return iterables.OfSlice[uint8]()
	}
	result := make(iterables.Slice[uint8], q)
	var m, M uint8
	switch len(minMax) {
	case 0:
		m, M = 0, uints.Max8
	case 1:
		m, M = minMax[0], uints.Max8
	default:
		m, M = minMax[0], minMax[1]
	}
	if m == M {
		for i := range result {
			result[i] = prng.Uint8(m, M)
		}
	}
	if m > M {
		m, M = M, m
	}
	for i := range result {
		result[i] = prng.Uint8(m, M)
	}
	return &result
}

// Uint16 returns a slice of q random uint16s.
// If q is negative, it returns an empty slice.
func Uint16(q int, minMax ...uint16) *iterables.Slice[uint16] {
	if q <= 0 {
		return iterables.OfSlice[uint16]()
	}
	result := make(iterables.Slice[uint16], q)
	var m, M uint16
	switch len(minMax) {
	case 0:
		m, M = 0, uints.Max16
	case 1:
		m, M = minMax[0], uints.Max16
	default:
		m, M = minMax[0], minMax[1]
	}
	if m == M {
		for i := range result {
			result[i] = prng.Uint16(m, M)
		}
	}
	if m > M {
		m, M = M, m
	}
	for i := range result {
		result[i] = prng.Uint16(m, M)
	}
	return &result
}

// Uint32 returns a slice of q random uint32s.
// If q is negative, it returns an empty slice.
func Uint32(q int, minMax ...uint32) *iterables.Slice[uint32] {
	if q <= 0 {
		return iterables.OfSlice[uint32]()
	}
	result := make(iterables.Slice[uint32], q)
	var m, M uint32
	switch len(minMax) {
	case 0:
		m, M = 0, uints.Max32
	case 1:
		m, M = minMax[0], uints.Max32
	default:
		m, M = minMax[0], minMax[1]
	}
	if m == M {
		for i := range result {
			result[i] = prng.Uint32(m, M)
		}
	}
	if m > M {
		m, M = M, m
	}
	for i := range result {
		result[i] = prng.Uint32(m, M)
	}
	return &result
}

// Uint64 returns a slice of q random uint64s.
// If q is negative, it returns an empty slice.
func Uint64(q int, minMax ...uint64) *iterables.Slice[uint64] {
	if q <= 0 {
		return iterables.OfSlice[uint64]()
	}
	result := make(iterables.Slice[uint64], q)
	var m, M uint64
	switch len(minMax) {
	case 0:
		m, M = 0, uints.Max64
	case 1:
		m, M = minMax[0], uints.Max64
	default:
		m, M = minMax[0], minMax[1]
	}
	if m == M {
		for i := range result {
			result[i] = prng.Uint64(m, M)
		}
	}
	if m > M {
		m, M = M, m
	}
	for i := range result {
		result[i] = prng.Uint64(m, M)
	}
	return &result
}

// UniqueByteSlices generates q unique byte slices
// with length between min and max (inclusive).
//
// If q is negative, it returns an empty slice.
func UniqueByteSlices(q int, maxRetries int, minMax ...int) (*iterables.Slice[[]byte], error) {
	if q <= 0 {
		return iterables.OfSlice[[]byte](), nil
	}
	uniqueBytes := make(map[string]struct{})
	result := make(iterables.Slice[[]byte], 0, q)
	m, M := xrng(minMax...)

	maxRetries *= q

	// Ensure the range is large enough for the distribution
	base := big.NewInt(256) // 256 possible values per byte
	exponent := big.NewInt(int64(M - m))
	rangeSize := new(big.Int).Exp(base, exponent, nil)
	if big.NewInt(int64(q)).Cmp(rangeSize) > 0 {
		return nil, gerrors.
			Tagged(fmt.Errorf("impossible to generate '%d' unique values with the given range [%d, %d]\n",
				q, m, M), ImpossibleConstraint)
	}

	for i := 0; i < q && maxRetries > 0; i++ {
		randomBytes := prng.Bytes(m, M)
		key := string(randomBytes)
		if _, ok := uniqueBytes[key]; !ok {
			uniqueBytes[key] = struct{}{}
			result = append(result, randomBytes)
			continue
		}
		i-- // Retry current index
		maxRetries--
	}

	l := result.Len()
	if l < q {
		return &result, fmt.Errorf("%d out of %d unique values generated with the given range [%d, %d]\n", l, q, m, M)
	}

	return &result, nil
}

// Uuid generates a slice of UUID strings of length q
func Uuid(q int) *iterables.Slice[string] {
	if q <= 0 {
		return iterables.OfSlice[string]()
	}

	result := make(iterables.Slice[string], q)
	// Generate UUID strings
	for i := 0; i < q; i++ {
		result[i] = lite.RandUuid()
	}
	return &result
}

// Values returns q random values of the type passed as argument
func Values(value any, q int) *iterables.Slice[any] {
	if q <= 0 {
		return iterables.OfSlice[any]()
	}
	typx := reflect.TypeOf(value)
	result := make(iterables.Slice[any], q)
	for i := 0; i < q; i++ {
		result[i] = lite.RandOf(typx)
	}
	return &result
}
