// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package ser

import "testing"

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
}
