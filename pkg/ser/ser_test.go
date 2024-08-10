// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package ser

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

type MyStruct struct {
	Field string `json:"field"`
}

func TestUnmarshalSingle(t *testing.T) {
	var s MyStruct
	data := []byte(`{"field": "test"}`)
	err := UnmarshalSingle(&s, data)
	if err != nil {
		t.Fatalf("Expected no error, but got %v", err)
	}
	if s.Field != "test" {
		t.Errorf("Expected field to be 'test', but got '%s'", s.Field)
	}

	// Test unmarshalling into a string
	var str string
	data = []byte(`"hello"`)
	err = UnmarshalSingle(&str, data)
	if err != nil {
		t.Fatalf("Expected no error, but got %v", err)
	}
	if str != `"hello"` {
		t.Errorf("Expected str to be 'hello', but got '%s'", str)
	}

	// Test unmarshalling into an integer
	var num int
	data = []byte(`42`)
	err = UnmarshalSingle(&num, data)
	if err != nil {
		t.Fatalf("Expected no error, but got %v", err)
	}
	if num != 42 {
		t.Errorf("Expected num to be 42, but got %d", num)
	}

	// Test unmarshalling into a void type
	t.Run("unmarshal into void type", func(t *testing.T) {
		type Void struct{}
		var result Void

		err := UnmarshalSingle(&result, []byte(`{}`))
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expected := Void{}
		if result != expected {
			t.Errorf("expected: %v, got: %v", expected, result)
		}
	})

	t.Run("unmarshal array error", func(t *testing.T) {
		jsonData := []byte(`[1, 2, 3]`)
		var result int

		err := UnmarshalSingle(&result, jsonData)
		if err == nil {
			t.Fatal("expected an error due to array input, but got none")
		}

		expectedErr := fmt.Sprintf("unexpected array data for single value unmarshal: %s", string(jsonData))
		if err.Error() != expectedErr {
			t.Errorf("expected error: %v, got: %v", expectedErr, err)
		}
	})
}

func TestUnmarshalInto(t *testing.T) {
	// Test case 1: Unmarshalling into a slice of structs
	t.Run("unmarshal into slice of structs", func(t *testing.T) {
		type MyStruct struct {
			Name  string `json:"name"`
			Age   int    `json:"age"`
			Email string `json:"email"`
		}

		jsonData := []byte(`{"name": "John", "age": 30, "email": "john@example.com"}`)
		var result []MyStruct

		err := UnmarshalInto(&result, jsonData)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expected := []MyStruct{{Name: "John", Age: 30, Email: "john@example.com"}}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("expected: %v, got: %v", expected, result)
		}
	})

	// Test case 2: Unmarshalling into a slice of strings
	t.Run("unmarshal into slice of strings", func(t *testing.T) {
		jsonData := []byte(`"This is a string"`)
		var result []string

		err := UnmarshalInto(&result, jsonData)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expected := []string{`"This is a string"`}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("expected: %v, got: %v", expected, result)
		}
	})

	// Test case 3: Unmarshalling into a slice of ints
	t.Run("unmarshal into slice of ints", func(t *testing.T) {
		jsonData := []byte(`[1, 2, 3]`)
		var result []int

		err := UnmarshalInto(&result, jsonData)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expected := []int{1, 2, 3}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("expected: %v, got: %v", expected, result)
		}
	})

	// Test case 4: Unmarshalling invalid JSON
	t.Run("unmarshal invalid JSON", func(t *testing.T) {
		jsonData := []byte(`{"invalid_json"}`)
		var result []int

		err := UnmarshalInto(&result, jsonData)
		if err == nil {
			t.Fatalf("expected an error, but got none")
		}

		if !strings.Contains(err.Error(), "failed unmarshalling data") {
			t.Errorf("expected error message to contain 'failed unmarshalling data', got: %v", err.Error())
		}
	})

	// Test case 5: Unmarshalling into an unsupported type (e.g., function type)
	t.Run("unmarshal unsupported type", func(t *testing.T) {
		type FuncType func()
		jsonData := []byte(`{}`)

		var result []FuncType

		err := UnmarshalInto(&result, jsonData)
		if err == nil {
			t.Fatalf("expected an error, but got none")
		}

		if !strings.Contains(err.Error(), "cannot unmarshal") {
			t.Errorf("expected error message to contain 'cannot unmarshal', got: %v", err.Error())
		}
	})

	// Test case 6: Unmarshalling into a slice of struct pointers
	t.Run("unmarshal into slice of struct pointers", func(t *testing.T) {
		type MyStruct struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}

		jsonData := []byte(`{"name": "Alice", "age": 25}`)
		var result []*MyStruct

		err := UnmarshalInto(&result, jsonData)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expected := []*MyStruct{{Name: "Alice", Age: 25}}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("expected: %v, got: %v", expected, result)
		}
	})

	// Test case 7: Unmarshalling raw string data into a slice of strings
	t.Run("unmarshal raw string data into slice of strings", func(t *testing.T) {
		rawData := []byte("Just a simple string")
		var result []string

		err := UnmarshalInto(&result, rawData)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expected := []string{"Just a simple string"}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("expected: %v, got: %v", expected, result)
		}
	})
}

func TestUnmarshalJson(t *testing.T) {
	// Test case 1: unmarshalling into a simple struct
	t.Run("unmarshal into struct", func(t *testing.T) {
		type MyStruct struct {
			Name  string `json:"name"`
			Age   int    `json:"age"`
			Email string `json:"email"`
		}

		jsonData := []byte(`{"name": "John", "age": 30, "email": "john@example.com"}`)
		result, err := UnmarshalJson[MyStruct](jsonData)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expected := MyStruct{Name: "John", Age: 30, Email: "john@example.com"}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("expected: %v, got: %v", expected, result)
		}
	})

	// Test case 2: unmarshalling into a string
	t.Run("unmarshal into string", func(t *testing.T) {
		jsonData := []byte(`"This is a string"`)
		result, err := UnmarshalJson[string](jsonData)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expected := "This is a string"

		if result != expected {
			t.Errorf("expected: %v, got: %v", expected, result)
		}
	})

	// Test case 3: unmarshalling into a slice
	t.Run("unmarshal into slice", func(t *testing.T) {
		jsonData := []byte(`[1, 2, 3, 4, 5]`)
		result, err := UnmarshalJson[[]int](jsonData)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expected := []int{1, 2, 3, 4, 5}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("expected: %v, got: %v", expected, result)
		}
	})

	// Test case 4: unmarshalling into a map
	t.Run("unmarshal into map", func(t *testing.T) {
		jsonData := []byte(`{"key1": "value1", "key2": "value2"}`)
		result, err := UnmarshalJson[map[string]string](jsonData)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expected := map[string]string{"key1": "value1", "key2": "value2"}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("expected: %v, got: %v", expected, result)
		}
	})

	// Test case 5: unmarshalling invalid JSON
	t.Run("unmarshal invalid JSON", func(t *testing.T) {
		jsonData := []byte(`{"invalid_json"}`)
		_, err := UnmarshalJson[int](jsonData)

		if err == nil {
			t.Fatalf("expected an error, but got none")
		}

		if !strings.Contains(err.Error(), "failed unmarshalling data") {
			t.Errorf("expected error message to contain 'failed unmarshalling data', got: %v", err.Error())
		}
	})

	// Test case 6: unmarshalling into an unsupported type (e.g., function type)
	t.Run("unmarshal unsupported type", func(t *testing.T) {
		type FuncType func()
		jsonData := []byte(`{}`)

		_, err := UnmarshalJson[FuncType](jsonData)

		if err == nil {
			t.Fatalf("expected an error, but got none")
		}

		if !strings.Contains(err.Error(), "cannot unmarshal") {
			t.Errorf("expected error message to contain 'cannot unmarshal', got: %v", err.Error())
		}
	})
}
