// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package random

import (
	"fmt"
	"github.com/andrerrcosta2/gtools/core/data/str/iterables"
	"github.com/andrerrcosta2/gtools/core/domain/gerrors"
	"github.com/andrerrcosta2/gtools/core/seeders/random/internal/constr"
	"github.com/andrerrcosta2/gtools/core/seeders/random/internal/prng"
	"github.com/google/uuid"
	"math/big"
	"math/rand"
	"reflect"
	"time"
	"unicode"
	"unicode/utf8"
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

// Of generates a slice of length q with random values of type T.
// This method may use reflection to prototype complex types.
func Of[T any](q int) *iterables.Slice[T] {
	if q <= 0 {
		return iterables.OfSlice[T]()
	}
	var zero T
	typx := reflect.TypeOf(zero)
	result := make(iterables.Slice[T], q)
	for i := 0; i < q; i++ {
		// That seems a little bit overcautious from golang compiler to me.
		// Fortunately the optimizations of golang compiler reduces the overhead
		// here close to zero.
		if rnd, ok := randOf(typx).(T); ok {
			result[i] = rnd
		}
	}
	return &result
}

// SingleOf generates a random value of type T
func SingleOf[T any]() (single T) {
	typx := reflect.TypeOf(single)
	single, ok := randOf(typx).(T)
	if !ok {
		panic("unable to generate a random value of type " + typx.String())
	}
	return
}

// Values returns q random values of the type passed as argument
func Values(value any, q int) *iterables.Slice[any] {
	if q <= 0 {
		return iterables.OfSlice[any]()
	}
	typx := reflect.TypeOf(value)
	result := make(iterables.Slice[any], q)
	for i := 0; i < q; i++ {
		result[i] = randOf(typx)
	}
	return &result
}

func Single(value any) (single any) {
	return randOf(reflect.TypeOf(value))
}

// randomOf generates a single random value
func randOf(t reflect.Type) any {
	if t == nil {
		return nil
	}
	kind := t.Kind()
	switch kind {
	case reflect.Bool:
		return prng.Bool()
	case reflect.Int:
		return prng.Int()
	case reflect.Int8:
		return prng.Int8()
	case reflect.Int16:
		return prng.Int16()
	case reflect.Int32:
		return prng.Int32()
	case reflect.Int64:
		return prng.Int64()
	case reflect.Uint, reflect.Uintptr:
		return prng.Uint()
	case reflect.Uint8:
		return prng.Uint8()
	case reflect.Uint16:
		return prng.Uint16()
	case reflect.Uint32:
		return prng.Uint32()
	case reflect.Uint64:
		return prng.Uint64()
	case reflect.Float32:
		return prng.Float32()
	case reflect.Float64:
		return prng.Float64()
	case reflect.Complex64:
		return prng.Complex64()
	case reflect.Complex128:
		return prng.Complex128()
	case reflect.Array:
		return randOf(t.Elem())
	case reflect.Chan:
		return randOf(t.Elem())
	case reflect.Map:
		// Create an empty map
		mapValue := reflect.MakeMap(t)
		// Generate a random key
		key := randOf(t.Key())
		// Generate a random value
		value := randOf(t.Elem())
		// Add the key-value pair to the map
		mapValue.SetMapIndex(reflect.ValueOf(key), reflect.ValueOf(value))
		return mapValue.Interface()

	case reflect.Ptr:
		value := randOf(t.Elem())
		ptr := reflect.New(t.Elem())
		ptr.Elem().Set(reflect.ValueOf(value))
		return ptr.Interface()
	case reflect.Slice:
		return randOf(t.Elem())
	case reflect.Struct:
		return structOf(t)
	case reflect.Interface:
		// is not directly possible to create an unnamed struct in memory based
		// solely on an interface in Go. This is because Go is a statically
		// typed language, and the type system does not support runtime creation
		// of new types, such as structs, without predefined definitions.
		// Reflection in Go operates on existing types and values; it cannot
		// define new types dynamically at runtime.
		// For that reason this method is most likely to panic when no element
		// is found on interface creation.
		if t.NumMethod() == 0 {
			return prng.Bytes(MinBytesSizeDefaults, MaxBytesSizeDefaults)
		}
		return randOf(t.Elem())

	case reflect.String:
		// TODO: This method lacks constraints control. Finish validation library
		// TODO: Add its interface to the core package
		return randString(prng.Int(1, 50), constr.AlphaNumeric)
	default:
		panic(fmt.Sprintf("gtools:random: nil or invalid type: '%T' passed to be randomized", t))
	}
}

// Int generates a slice of length q with random int values between min and max.
// If min or max are not provided, it defaults to the minimum and maximum int values.
// Returns an empty slice if q is negative or zero.
func Int(q int, minMax ...int) *iterables.Slice[int] {
	if q <= 0 {
		return iterables.OfSlice[int]()
	}
	result := make(iterables.Slice[int], q)
	for i := range result {
		result[i] = prng.Int(minMax...)
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
	for i := range result {
		result[i] = prng.Int8(minMax...)
	}
	return &result
}

func Int16(q int, minMax ...int16) *iterables.Slice[int16] {
	if q <= 0 {
		return iterables.OfSlice[int16]()
	}
	result := make(iterables.Slice[int16], q)
	for i := range result {
		result[i] = prng.Int16(minMax...)
	}
	return &result
}

func Int32(q int, minMax ...int32) *iterables.Slice[int32] {
	if q <= 0 {
		return iterables.OfSlice[int32]()
	}
	result := make(iterables.Slice[int32], q)
	for i := range result {
		result[i] = prng.Int32(minMax...)
	}
	return &result
}

func Int64(q int, minMax ...int64) *iterables.Slice[int64] {
	if q <= 0 {
		return iterables.OfSlice[int64]()
	}
	result := make(iterables.Slice[int64], q)
	for i := range result {
		result[i] = prng.Int64(minMax...)
	}
	return &result
}

func Uint(q int, minMax ...uint) *iterables.Slice[uint] {
	if q <= 0 {
		return iterables.OfSlice[uint]()
	}
	result := make(iterables.Slice[uint], q)

	for i := range result {
		result[i] = prng.Uint(minMax...)
	}
	return &result
}

func Uint8(q int, minMax ...uint8) *iterables.Slice[uint8] {
	if q <= 0 {
		return iterables.OfSlice[uint8]()
	}
	result := make(iterables.Slice[uint8], q)
	for i := range result {
		result[i] = prng.Uint8(minMax...)
	}
	return &result
}

func Uint16(q int, minMax ...uint16) *iterables.Slice[uint16] {
	if q <= 0 {
		return iterables.OfSlice[uint16]()
	}
	result := make(iterables.Slice[uint16], q)
	for i := range result {
		result[i] = prng.Uint16(minMax...)
	}
	return &result
}

func Uint32(q int, minMax ...uint32) *iterables.Slice[uint32] {
	if q <= 0 {
		return iterables.OfSlice[uint32]()
	}
	result := make(iterables.Slice[uint32], q)
	for i := range result {
		result[i] = prng.Uint32(minMax...)
	}
	return &result
}

func Uint64(q int, minMax ...uint64) *iterables.Slice[uint64] {
	if q <= 0 {
		return iterables.OfSlice[uint64]()
	}
	result := make(iterables.Slice[uint64], q)

	for i := range result {
		result[i] = prng.Uint64(minMax...)
	}
	return &result
}

// Float32 generates a slice of length q with random float32 values between min and max.
// If min or max are not provided, it defaults to the minimum and maximum float32 values.
//   - For ranges near the maximum float32 value (MaxFloat32), the precision of float32
//     may cause a loss of variation in generated values. This is due to the inherent
//     limitations of the 23-bit mantissa in the IEEE 754 representation of float32,
//     where the distance between representable values grows with magnitude.
//
// Returns an empty slice if q is negative or zero.
func Float32(q int, minMax ...float32) *iterables.Slice[float32] {
	if q <= 0 {
		return iterables.OfSlice[float32]()
	}
	result := make(iterables.Slice[float32], q)

	for i := range result {
		result[i] = prng.Float32(minMax...)
	}
	return &result
}

// Float64 generates a slice of length q with random float64 values between min and max.
// If min or max are not provided, it defaults to the minimum and maximum float64 values.
//   - Float64 offers significantly higher precision compared to Float32, with a 52-bit
//     mantissa in the IEEE 754 representation. This allows for finer variation in
//     values even for large ranges near MaxFloat64.
//
// Returns an empty slice if q is negative or zero.
func Float64(q int, minMax ...float64) *iterables.Slice[float64] {
	if q <= 0 {
		return iterables.OfSlice[float64]()
	}
	result := make(iterables.Slice[float64], q)

	for i := range result {
		result[i] = prng.Float64(minMax...)
	}
	return &result
}

func Complex64(q int, realMinMaxImagMinMax ...float32) *iterables.Slice[complex64] {
	if q <= 0 {
		return iterables.OfSlice[complex64]()
	}
	result := make(iterables.Slice[complex64], q)
	for i := range result {
		result[i] = prng.Complex64(realMinMaxImagMinMax...)
	}
	return &result
}

func Complex128(q int, realMinMaxImagMinMax ...float64) *iterables.Slice[complex128] {
	if q <= 0 {
		return iterables.OfSlice[complex128]()
	}
	result := make(iterables.Slice[complex128], q)

	for i := range result {
		result[i] = prng.Complex128(realMinMaxImagMinMax...)
	}
	return &result
}

// Bool returns a slice of length q with random boolean values.
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

// String returns a slice of length q with random strings of length between min and max.
// Each string is composed of random runes.
// If min or max are not provided, it defaults to a minimum length of 1 and a maximum length
// equal to the maximum length of a string.
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
		result[i] = string(randRune())
	})

	return &result
}

// Alphanumeric returns a slice of length q with random alphanumeric strings between min and max length.
// If min or max are not provided, it defaults to a minimum length of 1 and a maximum length
// equal to the maximum length of a string.
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
		result[i] = randString(l, constr.AlphaNumeric)
	})

	return &result
}

// StringOf returns a slice of length q with random strings of length between min and max.
// Each string is composed of characters randomly chosen from the charset.
// If min or max are not provided, it defaults to a minimum length of 1 and a maximum length
// equal to the maximum length of a string.
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
			result = append(result, "")
		}
		return &result
	}

	// Generate random string lengths between min and max
	Int(q, minMax...).EachN(func(i, l int) {
		result[i] = randString(l, charset)
	})

	return &result
}

func randString(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
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
		result[i] = randTimestamp(from, diff)
	}
	return &result
}

func randTimestamp(from time.Time, diff time.Duration) time.Time {
	randomDuration := time.Duration(rand.Int63n(int64(diff)))
	return from.Add(randomDuration)
}

// Uuid generates a slice of UUID strings of length q
func Uuid(q int) *iterables.Slice[string] {
	if q <= 0 {
		return iterables.OfSlice[string]()
	}

	result := make(iterables.Slice[string], q)
	// Generate UUID strings
	for i := 0; i < q; i++ {
		result[i] = randUuid()
	}
	return &result
}

func randUuid() string {
	return uuid.New().String()
}

// Bytes returns a slice of length q with random byte values.
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

func UniqueByteSlices(q int, maxRetries int, minMax ...int) (*iterables.Slice[[]byte], error) {
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
		result[i] = randRune()
	}
	return &result
}

// randRune generates a random Unicode rune that is valid and graphic.
// It continues to generate runes until one satisfies these conditions.
func randRune() rune {
	for {
		// Generate a random rune within the full Unicode range (up to 0x10FFFF)
		r := rune(prng.Uint(0, 0x10FFFF))
		// Check if the rune is valid and a graphic character (i.e. not a control character like '\n', '\t', etc.)
		if utf8.ValidRune(r) && unicode.IsGraphic(r) {
			return r
		}
	}
}

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
		// Again, a little bit overcautious to me.
		// But the output of this is an interface{}, so golang compiler has no
		// idea the type generated returns a T.
		// It might be worth duplicating the structOf function
		// to avoid type assertions but this is a really negligible overhead.
		if s, ok := randOf(t).(T); ok {
			result[i] = s
		}
	}

	return &result
}

// structOf generates a new instance of the given struct type with random values for its fields.
// It uses reflection to dynamically create an instance and set the fields.
func structOf(t reflect.Type) any {
	// Create a new instance of the struct
	v := reflect.New(t).Elem()

	// Iterate through each field of the struct
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldValue := v.Field(i)

		// Skip unexported fields
		if !fieldValue.CanSet() {
			continue
		}

		// Generate a random value for the field using randOf
		randomValue := randOf(field.Type)
		fieldValue.Set(reflect.ValueOf(randomValue))
	}

	// Return the newly created struct instance with populated fields
	return v.Interface()
}
