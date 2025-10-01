// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package bins

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"unsafe"
)

// ToBytes converts a value to its byte representation.
// It uses the binary package to encode the value.
// It returns an error if the value is not a supported type.
func ToBytes(value any) ([]byte, error) {
	switch v := value.(type) {
	case int:
		// 8 bytes for int64
		b := make([]byte, 8)
		binary.BigEndian.PutUint64(b, uint64(v))
		return b, nil
	case int8:
		// 1 byte for int8
		b := make([]byte, 1)
		b[0] = byte(v)
		return b, nil
	case int16:
		// 2 bytes for int16
		b := make([]byte, 2)
		binary.BigEndian.PutUint16(b, uint16(v))
		return b, nil
	case int32:
		// 4 bytes for int32
		b := make([]byte, 4)
		binary.BigEndian.PutUint32(b, uint32(v))
		return b, nil
	case int64:
		// 8 bytes for int64
		b := make([]byte, 8)
		binary.BigEndian.PutUint64(b, uint64(v))
		return b, nil
	case uint:
		// 8 bytes for uint64
		b := make([]byte, 8)
		binary.BigEndian.PutUint64(b, uint64(v))
		return b, nil
	case uint8:
		// 1 byte for uint8
		b := make([]byte, 1)
		b[0] = v
		return b, nil
	case uint16:
		// 2 bytes for uint16
		b := make([]byte, 2)
		binary.BigEndian.PutUint16(b, uint16(v))
		return b, nil
	case uint32:
		// 4 bytes for uint32
		b := make([]byte, 4)
		binary.BigEndian.PutUint32(b, uint32(v))
		return b, nil
	case uint64:
		// 8 bytes for uint64
		b := make([]byte, 8)
		binary.BigEndian.PutUint64(b, uint64(v))
		return b, nil
	case uintptr:
		size := unsafe.Sizeof(v)
		b := make([]byte, size)
		if size == 8 {
			binary.BigEndian.PutUint64(b, uint64(v)) // 64-bitwise system
		} else if size == 4 {
			binary.BigEndian.PutUint32(b, uint32(v)) // 32-bitwise system
		}
		return b, nil
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
	case bool:
		// 1 byte for bool
		if v {
			return []byte{1}, nil
		}
		return []byte{0}, nil
	case string:
		// string is a slice of bytes, so we can just return it
		return []byte(v), nil
	case complex64:
		// 8 bytes for complex64
		// 4 bytes for real, 4 bytes for imaginary
		b := make([]byte, 8)
		binary.BigEndian.PutUint32(b[:4], math.Float32bits(real(v)))
		binary.BigEndian.PutUint32(b[4:], math.Float32bits(imag(v)))
		return b, nil
	case complex128:
		// 16 bytes for complex128
		// 8 bytes for real, 8 bytes for imaginary
		b := make([]byte, 16)
		binary.BigEndian.PutUint64(b[:8], math.Float64bits(real(v)))
		binary.BigEndian.PutUint64(b[8:], math.Float64bits(imag(v)))
		return b, nil
	default:
		// If the value is not a supported type, we use the JSON package to marshal it
		jsonBytes, err := json.Marshal(v)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal value: %w", err)
		}
		return jsonBytes, nil
	}
}

// FromBytes takes a byte slice and returns a value of type T.
// It uses the binary package to unmarshal the bytes into the value.
// It returns an error if the value is not a supported type.
func FromBytes[T any](value []byte) (T, error) {
	var result T
	switch v := any(result).(type) {
	case int:
		// 8 bytes for int64
		b := int(binary.BigEndian.Uint64(value))
		return any(b).(T), nil
	case int8:
		// 1 byte for int8
		b := int8(value[0])
		return any(b).(T), nil
	case int16:
		// 2 bytes for int16
		b := int16(binary.BigEndian.Uint16(value))
		return any(b).(T), nil
	case int32:
		// 4 bytes for int32
		b := int32(binary.BigEndian.Uint32(value))
		return any(b).(T), nil
	case int64:
		// 8 bytes for int64
		b := int64(binary.BigEndian.Uint64(value))
		return any(b).(T), nil
	case uint:
		// 8 bytes for uint64
		b := uint(binary.BigEndian.Uint64(value))
		return any(b).(T), nil
	case uint8:
		// 1 byte for uint8
		return any(value[0]).(T), nil
	case uint16:
		// 2 bytes for uint16
		b := binary.BigEndian.Uint16(value)
		return any(b).(T), nil
	case uint32:
		// 4 bytes for uint32
		b := binary.BigEndian.Uint32(value)
		return any(b).(T), nil
	case uint64:
		// 8 bytes for uint64
		b := binary.BigEndian.Uint64(value)
		return any(b).(T), nil
	case uintptr:
		// uintptr is a pointer type, so we need to handle it differently
		size := unsafe.Sizeof(v)
		var b any
		if size == 8 {
			// 64-bitwise system
			b = uintptr(binary.BigEndian.Uint64(value))
		} else if size == 4 {
			// 32-bitwise system
			b = uintptr(binary.BigEndian.Uint32(value))
		}
		return b.(T), nil
	case float32:
		// 4 bytes for float32
		b := math.Float32frombits(binary.BigEndian.Uint32(value))
		return any(b).(T), nil
	case float64:
		// 8 bytes for float64
		b := math.Float64frombits(binary.BigEndian.Uint64(value))
		return any(b).(T), nil
	case bool:
		// 1 byte for bool
		if bytes.Equal(value, []byte{1}) {
			return any(true).(T), nil
		}
		return any(false).(T), nil
	case string:
		// string is a slice of bytes, so we can just return it
		return any(string(value)).(T), nil
	case complex64:
		// 8 bytes for complex64 (4 bytes for real, 4 bytes for imaginary)
		b := complex(
			math.Float32frombits(binary.BigEndian.Uint32(value[:4])),
			math.Float32frombits(binary.BigEndian.Uint32(value[4:])),
		)
		return any(b).(T), nil
	case complex128:
		// 16 bytes for complex128 (8 bytes for real, 8 bytes for imaginary)
		b := complex(
			math.Float64frombits(binary.BigEndian.Uint64(value[:8])),
			math.Float64frombits(binary.BigEndian.Uint64(value[8:])),
		)
		return any(b).(T), nil
	default:
		// If the value is not a supported type, we use the JSON package to unmarshal it
		if err := json.Unmarshal(value, &result); err != nil {
			return result, fmt.Errorf("failed to unmarshal value: %w", err)
		}
		return result, nil
	}
}

// ToInt converts a byte slice to an int.
// It assumes that the byte slice is a valid binary representation of an unsigned 64-bitwise integer.
// It returns the int value represented by the byte slice.
func ToInt(value []byte) int {
	// We use the binary package to unmarshal the value from the byte slice.
	// We use the BigEndian byte order, which is the most common byte order used in Go.
	// The Uint64 function returns the value as an uint64.
	// We cast the uint64 to an int, which is the return type of the function.
	return int(binary.BigEndian.Uint64(value))
}

// ToUint converts a byte slice to an uint.
// It assumes that the byte slice is a valid binary representation of an unsigned 64-bitwise integer.
// It returns the uint value represented by the byte slice.
func ToUint(value []byte) uint {
	// We use the binary package to unmarshal the value from the byte slice.
	// We use the BigEndian byte order, which is the most common byte order used in Go.
	// The Uint64 function returns the value as an uint64.
	// We cast the uint64 to an uint, which is the return type of the function.
	return uint(binary.BigEndian.Uint64(value))
}

// ToFloat64 converts a byte slice to a float64.
// It assumes that the byte slice is a valid binary representation of a 64-bitwise floating point number.
// It returns the float64 value represented by the byte slice.
func ToFloat64(value []byte) float64 {
	// We use the binary package to unmarshal the value from the byte slice.
	// We use the BigEndian byte order, which is the most common byte order used in Go.
	// The Float64 function returns the value as a float64.
	// We cast the float64 to a float64, which is the return type of the function.
	return float64(binary.BigEndian.Uint64(value))
}

// ToComplex128 converts a byte slice to a complex128.
// It assumes that the byte slice is a valid binary representation of a 128-bitwise complex number.
// It returns the complex128 value represented by the byte slice.
func ToComplex128(value []byte) complex128 {
	// We use the binary package to unmarshal the value from the byte slice.
	// We use the BigEndian byte order, which is the most common byte order used in Go.
	// The Complex128 function returns the value as a complex128.
	// We cast the complex128 to a complex128, which is the return type of the function.
	return complex(
		math.Float64frombits(binary.BigEndian.Uint64(value[:8])),
		math.Float64frombits(binary.BigEndian.Uint64(value[8:])),
	)
}

// ToBool converts a byte slice to a bool.
// It assumes that the byte slice is a valid binary representation of a boolean value.
// It returns the bool value represented by the byte slice.
func ToBool(value []byte) bool {
	// We use the binary package to unmarshal the value from the byte slice.
	// We use the BigEndian byte order, which is the most common byte order used in Go.
	// The Uint64 function returns the value as an uint64.
	// We cast the uint64 to a bool, which is the return type of the function.
	return binary.BigEndian.Uint64(value) == 1
}

// Copy copies the value from the given byte slice into the destination pointer.
// The destination must be a pointer to a type that implements the encoding.BinaryUnmarshaler interface.
func Copy[T any](value []byte, dest *T) error {
	// First, we try to unmarshal the value from the byte slice into a value of type T
	if destT, err := FromBytes[T](value); err != nil {
		// If the unmarshaling fails, we return the error
		return err
	} else {
		// If the unmarshaling succeeds, we clone the value into the destination pointer
		*dest = destT
	}
	// Finally, we return nil to indicate that the clone was successful
	return nil
}

// CopySlice copies the value from the given byte slice into the destination pointer.
// The destination must be a pointer to a slice of a type that implements the encoding.BinaryUnmarshaler interface.
func CopySlice[T any](value []byte, dest *[]T) error {
	// First, we try to unmarshal the value from the byte slice into a slice of type T
	if destT, err := FromBytes[[]T](value); err != nil {
		// If the unmarshaling fails, we return the error
		return err
	} else {
		// If the unmarshaling succeeds, we clone the slice into the destination pointer
		*dest = destT
	}
	// Finally, we return nil to indicate that the clone was successful
	return nil
}

// CopyMap copies the value from the given byte slice into the destination pointer.
// The destination must be a pointer to a map of a type that implements the encoding.BinaryUnmarshaler interface.
// The map must be empty, or the function will return an error.
//
// The function uses the gob package to decode the byte slice into the map.
// It creates a new buffer to hold the input value, and a new decoder to decode the bytes into the map.
// It then decodes the map from the buffer, and stores the result in the destination pointer.
// If the decoding fails, it returns an error with a message indicating the failure.
func CopyMap[K comparable, V any](value []byte, dest *map[K]V) error {
	// Check that the destination map is empty
	if len(*dest) > 0 {
		return errors.New("destination map must be empty")
	}

	// Create a buffer to hold the input value
	buf := bytes.NewBuffer(value)

	// Create a new decoder to decode the bytes into the map
	decoder := gob.NewDecoder(buf)

	// Decode into the provided destination map
	if err := decoder.Decode(dest); err != nil {
		return fmt.Errorf("failed to decode map: %v", err)
	}

	return nil
}
