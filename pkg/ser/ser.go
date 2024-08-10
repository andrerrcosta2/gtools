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

// UnmarshalSingle unmarshals data into a single value of type T.
// It supports JSON and raw string formats.
// If the type of T is a Void, it returns nil.
// If the type of T is a string, it uses the unmarshalRawString function.
// Otherwise, it uses the json.Unmarshal function.
// It returns an error if the unmarshalling fails.
func UnmarshalSingle[T any](out *T, data []byte) error {
	var err error
	// Check the type of T to determine the unmarshalling strategy
	switch any(*out).(type) {
	case *Void:
		return nil
	case string:
		var str string
		err = json.Unmarshal(data, &str)
		if err == nil {
			*out = any(str).(T)
		}
	default:
		err = json.Unmarshal(data, out)
	}

	if err != nil {
		return fmt.Errorf("failed unmarshalling data: %v [%s]", err, string(data))
	}

	return nil
}

// UnmarshalInto unmarshals data into the specified type.
// It supports JSON and raw string formats.
func UnmarshalInto[T any](out *[]T, data []byte) error {
	var it T
	var err error

	// Check the type of T to determine the unmarshalling strategy
	switch any(it).(type) {
	// Do nothing if it's a Void type
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
