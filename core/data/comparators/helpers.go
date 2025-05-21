// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package comparators

import (
	"github.com/andrerrcosta2/gtools/core/format/fmx"
	"strconv"
)

// Hash hashes a value between two comparable representations
func Hash[T comparable, H comparable](value T) H {
	str := StringHash(value)
	return ToComparable[H](str)
}

func StringHash[T comparable](value T) string {
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
		panic(fmx.Sprintf("comparators.StringOrdered type <%T> is not supported", value))
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
		panic(fmx.Sprintf("unsupported target type: %T", zero))
	}
}
