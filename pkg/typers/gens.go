// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package typers

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"math"
	"math/big"
	"reflect"
)

func Rand[T any](minBytes, maxBytes int, q int) ([]T, error) {
	if minBytes < 0 || maxBytes < minBytes || q < 0 {
		return nil, fmt.Errorf("invalid input parameters")
	}
	result := make([]T, q)
	t := reflect.TypeOf((*T)(nil)).Elem()

	for i := 0; i < q; i++ {
		// Generate random size between minBytes and maxBytes
		size, err := rand.Int(rand.Reader, big.NewInt(int64(maxBytes-minBytes+1)))
		if err != nil {
			return nil, fmt.Errorf("failed to generate random size: %w", err)
		}
		size.Add(size, big.NewInt(int64(minBytes)))

		// Generate random bytes
		randomBytes := make([]byte, size.Int64())
		_, err = rand.Read(randomBytes)
		if err != nil {
			return nil, fmt.Errorf("failed to generate random bytes: %w", err)
		}

		var value T
		switch t.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			intValue := big.NewInt(0).SetBytes(randomBytes).Int64()
			value = reflect.ValueOf(intValue).Convert(t).Interface().(T)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			uintValue := big.NewInt(0).SetBytes(randomBytes).Uint64()
			value = reflect.ValueOf(uintValue).Convert(t).Interface().(T)
		case reflect.Float32, reflect.Float64:
			floatValue := float64(big.NewInt(0).SetBytes(randomBytes).Int64()) / float64(1<<63)
			value = reflect.ValueOf(floatValue).Convert(t).Interface().(T)
		case reflect.String:
			value = reflect.ValueOf(string(randomBytes)).Convert(t).Interface().(T)
		default:
			return nil, fmt.Errorf("unsupported type: %v", t)
		}

		result[i] = value
	}

	return result, nil
}

func RandPtr[T any](minBytes, maxBytes int, q int) ([]*T, error) {
	if minBytes < 0 || maxBytes < minBytes || q < 0 {
		return nil, fmt.Errorf("invalid input parameters")
	}

	result := make([]*T, q)

	for i := 0; i < q; i++ {
		// Generate random size between minBytes and maxBytes
		size, err := rand.Int(rand.Reader, big.NewInt(int64(maxBytes-minBytes+1)))
		if err != nil {
			return nil, fmt.Errorf("failed to generate random size: %w", err)
		}
		size.Add(size, big.NewInt(int64(minBytes)))

		// Generate random bytes
		randomBytes := make([]byte, size.Int64())
		_, err = rand.Read(randomBytes)
		if err != nil {
			return nil, fmt.Errorf("failed to generate random bytes: %w", err)
		}

		var value T
		switch t := any(value).(type) {
		case int, int8, int16, int32, int64:
			intValue := int64(binary.BigEndian.Uint64(append(make([]byte, 8-len(randomBytes)), randomBytes...)))
			value = any(intValue).(T)
		case uint, uint8, uint16, uint32, uint64:
			uintValue := binary.BigEndian.Uint64(append(make([]byte, 8-len(randomBytes)), randomBytes...))
			value = any(uintValue).(T)
		case float32, float64:
			floatValue := float64(binary.BigEndian.Uint64(append(make([]byte, 8-len(randomBytes)), randomBytes...))) / float64(math.MaxUint64)
			value = any(floatValue).(T)
		case string, []byte:
			value = any(string(randomBytes)).(T)
		default:
			return nil, fmt.Errorf("unsupported type: %v", t)
		}
		ptr := new(T)
		*ptr = value
		result[i] = ptr
	}

	return result, nil
}
