// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package constraints

import "errors"

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~uintptr | ~float32 | ~float64 | ~string
}

func ToOrdered[T Ordered](value any) (T, error) {
	switch v := value.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr, float32, float64, string:
		var result T
		if value, ok := any(v).(T); ok {
			result = value
		} else {
			return result, errors.New("couldn't convert to T\n")
		}
		return result, nil

	default:
		var zero T
		return zero, errors.New("not an ordered type\n")
	}
}
