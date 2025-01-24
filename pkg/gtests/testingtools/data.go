// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package testingtools

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/andrerrcosta2/gtools/core/format"
	"github.com/andrerrcosta2/gtools/core/format/serializations"
	"reflect"
)

type reflectionDataToolsLite struct{}

// DeepCopy uses reflection to duplicate the entire object and all objects it references,
// recursively. This means that any nested or referenced data structures are fully
// cloned, so changes in the deep copy don’t affect the original, and vice versa.
// Each layer of the original data is copied independently.
func (t *reflectionDataToolsLite) DeepCopy(data any) (any, error) {
	if data == nil {
		return nil, nil
	}
	val := reflect.ValueOf(data)
	clone, err := cloneRecursive(val)
	if err != nil {
		return nil, err
	}
	return clone.Interface(), nil
}

func cloneRecursive(val reflect.Value) (reflect.Value, error) {
	switch val.Kind() {
	case reflect.Ptr:
		if val.IsNil() {
			return reflect.Zero(val.Type()), nil // return a zero pointer of the same type
		}
		// Create a new pointer and clone the dereferenced value
		clone := reflect.New(val.Elem().Type())
		clonedVal, err := cloneRecursive(val.Elem())
		if err != nil {
			return reflect.Value{}, err
		}
		clone.Elem().Set(clonedVal)
		return clone, nil

	case reflect.Struct:
		clone := reflect.New(val.Type()).Elem()
		for i := 0; i < val.NumField(); i++ {
			fieldClone, err := cloneRecursive(val.Field(i))
			if err != nil {
				return reflect.Value{}, err
			}
			clone.Field(i).Set(fieldClone)
		}
		return clone, nil

	case reflect.Slice:
		if val.IsNil() {
			return reflect.Zero(val.Type()), nil // return nil slice of the same type
		}
		clone := reflect.MakeSlice(val.Type(), val.Len(), val.Cap())
		for i := 0; i < val.Len(); i++ {
			elemClone, err := cloneRecursive(val.Index(i))
			if err != nil {
				return reflect.Value{}, err
			}
			clone.Index(i).Set(elemClone)
		}
		return clone, nil

	case reflect.Map:
		if val.IsNil() {
			return reflect.Zero(val.Type()), nil // return nil map of the same type
		}
		clone := reflect.MakeMapWithSize(val.Type(), val.Len())
		for _, key := range val.MapKeys() {
			valClone, err := cloneRecursive(val.MapIndex(key))
			if err != nil {
				return reflect.Value{}, err
			}
			keyClone, err := cloneRecursive(key)
			if err != nil {
				return reflect.Value{}, err
			}
			clone.SetMapIndex(keyClone, valClone)
		}
		return clone, nil

	case reflect.Array:
		clone := reflect.New(val.Type()).Elem()
		for i := 0; i < val.Len(); i++ {
			elemClone, err := cloneRecursive(val.Index(i))
			if err != nil {
				return reflect.Value{}, err
			}
			clone.Index(i).Set(elemClone)
		}
		return clone, nil

	default:
		// For basic types we copy directly
		return val, nil
	}
}

// ShallowCopy uses reflection to duplicate only the top-level structure, leaving nested
// or referenced objects shared between the original and the copy. For instance,
// if the object has fields that point to other objects (like slices or pointers),
// a shallow copy would copy only the references themselves, not the actual data they
// point to. Therefore, changes to the shared objects will be reflected in both the
// original and the shallow copy.
func (t *reflectionDataToolsLite) ShallowCopy(value any) (any, error) {
	if value == nil {
		return nil, nil
	}

	val := reflect.ValueOf(value)

	switch val.Kind() {
	case reflect.Ptr:
		// If it's a pointer, create a new pointer and copy the original value directly.
		clone := reflect.New(val.Elem().Type())
		clone.Elem().Set(val.Elem())
		return clone.Interface(), nil

	case reflect.Struct:
		// For structs, copy only the top-level fields.
		clone := reflect.New(val.Type()).Elem()
		for i := 0; i < val.NumField(); i++ {
			clone.Field(i).Set(val.Field(i)) // Copy references as-is
		}
		return clone.Interface(), nil

	case reflect.Slice:
		// For slices, create a new slice with the same elements.
		clone := reflect.MakeSlice(val.Type(), val.Len(), val.Cap())
		reflect.Copy(clone, val) // Shallow copy: copy the references within the slice
		return clone.Interface(), nil

	case reflect.Map:
		// For maps, create a new map but copy only the references to values.
		clone := reflect.MakeMapWithSize(val.Type(), val.Len())
		for _, key := range val.MapKeys() {
			clone.SetMapIndex(key, val.MapIndex(key)) // Copy references as-is
		}
		return clone.Interface(), nil

	case reflect.Array:
		// Arrays are value types in Go, so we just assign it directly.
		clone := reflect.New(val.Type()).Elem()
		reflect.Copy(clone.Slice(0, val.Len()), val.Slice(0, val.Len()))
		return clone.Interface(), nil

	default:
		// For basic types, just return the original value directly.
		return value, nil
	}
}

// ExtractField allows extraction of a named field from a struct.
func (t *reflectionDataToolsLite) ExtractField(obj interface{}, fieldName string) (interface{}, error) {
	v := reflect.ValueOf(obj)

	// Ensure obj is a pointer to a struct.
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return nil, fmt.Errorf("expected a pointer to a struct")
	}
	v = v.Elem()

	// Get the field by name
	field := v.FieldByName(fieldName)
	if !field.IsValid() {
		return nil, fmt.Errorf("field '%s' not found", fieldName)
	}

	return field.Interface(), nil
}

// InjectField allows injection of a value into a named field in a struct.
func (t *reflectionDataToolsLite) InjectField(obj interface{}, fieldName string, value interface{}) error {
	v := reflect.ValueOf(obj)

	// Ensure obj is a pointer to a struct.
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("expected a pointer to a struct")
	}
	v = v.Elem()

	// Get the field by name.
	field := v.FieldByName(fieldName)
	if !field.IsValid() {
		return fmt.Errorf("field %s not found", fieldName)
	}

	// Ensure the field is settable.
	if !field.CanSet() {
		return fmt.Errorf("field %s is not settable", fieldName)
	}

	// Ensure the value is assignable to the field.
	val := reflect.ValueOf(value)
	if val.Type() != field.Type() {
		return fmt.Errorf("provided value type %s does not match field type %s", val.Type(), field.Type())
	}

	field.Set(val)
	return nil
}

// EqualsBy compares two structs for equality, with options to ignore specific fields.
func (t *reflectionDataToolsLite) EqualsBy(a, b interface{}, ignoreFields ...string) (bool, error) {
	// Ensure both inputs are of the same type and are structs.
	v1 := reflect.ValueOf(a)
	v2 := reflect.ValueOf(b)

	if v1.Type() != v2.Type() {
		return false, fmt.Errorf("type mismatch: %T vs %T", a, b)
	}
	if v1.Kind() != reflect.Struct {
		return false, fmt.Errorf("only struct types are supported")
	}

	// Create a map for quick lookup of ignored fields.
	ignoreMap := make(map[string]struct{}, len(ignoreFields))
	for _, field := range ignoreFields {
		ignoreMap[field] = struct{}{}
	}

	// Iterate through the fields and compare values, ignoring specified fields.
	for i := 0; i < v1.NumField(); i++ {
		fieldType := v1.Type().Field(i)
		if _, ignore := ignoreMap[fieldType.Name]; ignore {
			continue
		}

		// Compare the fields; if they differ, return false.
		if !reflect.DeepEqual(v1.Field(i).Interface(), v2.Field(i).Interface()) {
			return false, nil
		}
	}

	return true, nil
}

// Stringify takes an interface{} and returns a string representation of it in the
// specified format.
func (t *reflectionDataToolsLite) Stringify(data interface{}, format format.Serialization) (string, error) {
	var result []byte
	var err error

	switch format {
	case serializations.JSON:
		// Marshal the data using the built-in json package.
		result, err = json.Marshal(data)
	case serializations.XML:
		// Marshal the data using the built-in xml package.
		result, err = xml.Marshal(data)
	default:
		// If the format is unsupported, return an error.
		return "", fmt.Errorf("unsupported format: %s", format)
	}

	if err != nil {
		// If an error occurred while marshaling, return it.
		return "", err
	}

	// Return the string representation of the marshaled data.
	return string(result), nil
}
