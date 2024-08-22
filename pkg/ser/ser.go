// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package ser

import (
	"encoding/json"
	"fmt"
	"github.com/andrerrcosta2/gtools/pkg/arrays"
	"strings"
)

const FailedMessage = "failed unmarshalling data: %v [%s]"

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

	if !arrays.Empty(&data) && data[0] == '[' {
		return fmt.Errorf("unexpected array data for single value unmarshal: %s", string(data))
	}

	// Check the type of T to determine the unmarshalling strategy
	switch any(*out).(type) {
	case *Void:
		return nil
	case string:
		*out = any(string(data)).(T)
	default:
		err = json.Unmarshal(data, out)
	}

	if err != nil {
		return fmt.Errorf(FailedMessage, err, string(data))
	}

	return nil
}

// UnmarshalInto unmarshals data into the specified type.
// It supports JSON and raw string formats.
func UnmarshalInto[T any](out *[]T, data []byte) error {
	var it T
	var err error

	// Check if the data is a JSON array
	if data[0] == '[' {
		var arr []T
		err = json.Unmarshal(data, &arr)
		if err != nil {
			return fmt.Errorf(FailedMessage, err, string(data))
		}
		*out = append(*out, arr...)
		return nil
	}

	// Check the type of T to determine the unmarshalling strategy
	switch any(it).(type) {
	// Do nothing if it's a Void type
	case *Void:
		return nil
	case string:
		it = any(string(data)).(T)
	default:
		err = json.Unmarshal(data, &it)
	}

	if err != nil {
		return fmt.Errorf(FailedMessage, err, string(data))
	}
	*out = append(*out, it)
	return nil
}

func UnmarshalJson[T any](data []byte) (any, error) {
	var it T
	err := json.Unmarshal(data, &it)
	if err != nil {
		return it, fmt.Errorf(FailedMessage, err, string(data))
	}
	return it, nil
}
