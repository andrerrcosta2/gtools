// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package ser

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

func Json(s any) string {
	var buf strings.Builder
	enc := json.NewEncoder(&buf)
	enc.SetIndent("", "  ")
	err := enc.Encode(s)
	if err != nil {
		return fmt.Sprintf("Error marshaling to JSON: %v\n", err)
	}
	return buf.String()
}

// UnmarshalInto unmarshals data into the specified type.
// It supports JSON and raw string formats.
func UnmarshalInto[T any](out *[]T, data []byte) error {
	var it T
	var err error

	// Check the type of T to determine the unmarshalling strategy
	switch any(it).(type) {
	case *Void:
		return nil
	case string:
		it, err = unmarshalRawString[T](data)
	default:
		err = json.Unmarshal(data, &it)
	}

	if err != nil {
		return fmt.Errorf("failed unmarshalling data: %v [%s]", err, string(data))
	}
	*out = append(*out, it)
	return nil
}

// unmarshalRawString unmarshals raw string data into the specified type.
func unmarshalRawString[T any](data []byte) (T, error) {
	var it T
	strData := string(data)
	it, ok := any(strData).(T)
	if !ok {
		return it, errors.New("type assertion to string failed")
	}
	return it, nil
}
