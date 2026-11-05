// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package keyers

import (
	"fmt"
	"strconv"
	"unsafe"

	"github.com/andrerrcosta2/gtools/core/domain/constraints"
)

// ComparableTyped hashes a value between two comparable representations
func ComparableTyped[T comparable, H comparable](value T) H {
	str := CmpToString(value)
	return ToComparable[H](str)
}

func CmpToString[T comparable](value T) string {
	switch v := any(value).(type) {
	case string:
		return v
	case int:
		return strconv.Itoa(v)
	case int8:
		return strconv.FormatInt(int64(v), 10)
	case int16:
		return strconv.FormatInt(int64(v), 10)
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case int64:
		return strconv.FormatInt(v, 10)
	case uint:
		return strconv.FormatUint(uint64(v), 10)
	case uint8:
		return strconv.FormatUint(uint64(v), 10)
	case uint16:
		return strconv.FormatUint(uint64(v), 10)
	case uint32:
		return strconv.FormatUint(uint64(v), 10)
	case uint64:
		return strconv.FormatUint(v, 10)
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	default:
		panic(fmt.Sprintf("comparator.StringOrdered type <%T> is not supported", value))
	}
}

// ToComparable converts a string to a value of type H.
func ToComparable[H comparable](str string) H {
	var z H
	switch zero := any(z).(type) {
	case string:
		return any(zero).(H)
	case int:
		v, _ := strconv.Atoi(str)
		return any(v).(H)
	case int8:
		v, _ := strconv.ParseInt(str, 10, 8)
		return any(int8(v)).(H)
	case int16:
		v, _ := strconv.ParseInt(str, 10, 16)
		return any(int16(v)).(H)
	case int32:
		v, _ := strconv.ParseInt(str, 10, 32)
		return any(int32(v)).(H)
	case int64:
		v, _ := strconv.ParseInt(str, 10, 64)
		return any(v).(H)
	case uint:
		v, _ := strconv.ParseUint(str, 10, 64)
		return any(uint(v)).(H)
	case uint8:
		v, _ := strconv.ParseUint(str, 10, 8)
		return any(uint8(v)).(H)
	case uint16:
		v, _ := strconv.ParseUint(str, 10, 16)
		return any(uint16(v)).(H)
	case uint32:
		v, _ := strconv.ParseUint(str, 10, 32)
		return any(uint32(v)).(H)
	case uint64:
		v, _ := strconv.ParseUint(str, 10, 64)
		return any(v).(H)
	case float32:
		v, _ := strconv.ParseFloat(str, 32)
		return any(float32(v)).(H)
	case float64:
		v, _ := strconv.ParseFloat(str, 64)
		return any(v).(H)
	default:
		panic(fmt.Sprintf("unsupported target type: %T", zero))
	}
}

// fingerprint returns a unique string identifier for the given object.
// If the object implements the domain.UniqueOf interface, its unique identifier is returned.
// Otherwise, the object's memory address or its string representation is returned.
// This method doesn't return any error and currently doesn't handle interfaces
// as generic type, so its reliability is the best it can be guaranteed based on what
// is delivered to it.
// That means it has its own view of equality between objects which follows the order:
//
//  1. its own uniqueness implementation
//  2. its natural comparability as a typed representation
//  3. its memory address if it is a pointer
//
// If your context requires a different unique representation, don't use this method.
func fingerprint[T any](v any) string {
	if id, ok := IdentityOf(v); ok {
		return id
	}

	if s, ok := primitiveFingerprint(v); ok {
		return s
	}

	if s, ok := v.(*T); ok {
		return fmt.Sprintf("0x%x", uintptr(unsafe.Pointer(s)))
	}

	if s, ok := v.(fmt.Stringer); ok {
		return s.String()
	}

	if err, ok := v.(error); ok {
		return err.Error()
	}

	return fallbackFingerprint(v)
}

func primitiveFingerprint(v any) (string, bool) {
	switch s := v.(type) {
	case int:
		return fmt.Sprintf("<int>%d", s), true
	case int8:
		return fmt.Sprintf("<int8>%d", s), true
	case int16:
		return fmt.Sprintf("<int16>%d", s), true
	case int32:
		return fmt.Sprintf("<int32>%d", s), true
	case int64:
		return fmt.Sprintf("<int64>%d", s), true
	case uint:
		return fmt.Sprintf("<uint>%d", s), true
	case uint8:
		return fmt.Sprintf("<uint8>%d", s), true
	case uint16:
		return fmt.Sprintf("<uint16>%d", s), true
	case uint32:
		return fmt.Sprintf("<uint32>%d", s), true
	case uint64:
		return fmt.Sprintf("<uint64>%d", s), true
	case float32:
		return fmt.Sprintf("<float32>%g", s), true
	case float64:
		return fmt.Sprintf("<float64>%g", s), true
	case string:
		return fmt.Sprintf("<string>%s", s), true
	case bool:
		return fmt.Sprintf("<bool>%t", s), true
	case complex64:
		return fmt.Sprintf("<complex64>%v", s), true
	case complex128:
		return fmt.Sprintf("<complex128>%v", s), true
	default:
		return "", false
	}
}

func IdentityOf(v any) (string, bool) {
	if id, ok := v.(constraints.Identifiable); ok {
		fmt.Printf("Trying to find identity of %v\n", id)
		return id.ID(), true
	}
	return "", false
}

func fallbackFingerprint(v any) string {
	return fmt.Sprintf("%T:%v", v, v)
}
