// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package floats

import (
	"encoding/binary"
	"fmt"
	"math"
)

// ToBytes converts a float32 or float64 value to its byte representation.
// It uses the binary package to encode the value.
// It returns an error if the value is not a supported type.
func ToBytes(value any) ([]byte, error) {
	switch v := value.(type) {
	case float32:
		// 4 bytes for float32
		b := make([]byte, 4)
		binary.BigEndian.PutUint32(b, math.Float32bits(v))
		return b, nil
	case float64:
		// 8 bytes for float64
		b := make([]byte, 8)
		binary.BigEndian.PutUint64(b, math.Float64bits(v))
		return b, nil
	default:
		return nil, fmt.Errorf("unsupported type %T", v)
	}
}

// FromBytes converts a byte slice to a float32 or float64 value.
// It uses the binary package to decode the value.
// It returns an error if the value is not a supported type.
func FromBytes(b []byte) (any, error) {
	switch len(b) {
	case 4:
		// 4 bytes for float32
		return math.Float32frombits(binary.BigEndian.Uint32(b)), nil
	case 8:
		// 8 bytes for float64
		return math.Float64frombits(binary.BigEndian.Uint64(b)), nil
	default:
		return nil, fmt.Errorf("invalid byte slice length %d", len(b))
	}
}

func ToFloat32(val any) (float32, error) {
	switch v := val.(type) {
	case float64:
		return float32(v), nil
	case float32:
		return v, nil
	case int:
		return float32(v), nil
	case int8:
		return float32(v), nil
	case int16:
		return float32(v), nil
	case int32:
		return float32(v), nil
	case int64:
		return float32(v), nil
	case uint:
		return float32(v), nil
	case uint8:
		return float32(v), nil
	case uint16:
		return float32(v), nil
	case uint32:
		return float32(v), nil
	case uint64:
		return float32(v), nil
	default:
		return 0, fmt.Errorf("type not convertible to float32: %T", v)
	}
}

func ToFloat64(val any) (float64, error) {
	switch v := val.(type) {
	case float64:
		return v, nil
	case float32:
		return float64(v), nil
	case int:
		return float64(v), nil
	case int8:
		return float64(v), nil
	case int16:
		return float64(v), nil
	case int32:
		return float64(v), nil
	case int64:
		return float64(v), nil
	case uint:
		return float64(v), nil
	case uint8:
		return float64(v), nil
	case uint16:
		return float64(v), nil
	case uint32:
		return float64(v), nil
	case uint64:
		return float64(v), nil
	default:
		return 0, fmt.Errorf("type not convertible to float64: %T", v)
	}
}
