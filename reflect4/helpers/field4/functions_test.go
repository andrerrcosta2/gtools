// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package field4

import (
	"strings"
	"testing"
)

type testStruct struct {
	Name    string
	Age     int
	private bool // unexported
	Value   *int
}

type emptyStruct struct{}

func TestEachExp(t *testing.T) {
	t.Run("non-struct returns error", func(t *testing.T) {
		err := EachExp("not a struct", func(name string, value any) bool {
			t.Errorf("callback should not be called")
			return true
		})
		if err == nil {
			t.Errorf("expected error for non-struct, got nil")
		}
		expErr := ErrNotStruct("field4.EachExp", "string").Error()
		if !strings.Contains(err.Error(), expErr) {
			t.Errorf("expected '%s', got: '%s'", expErr, err.Error())
		}
	})

	t.Run("empty struct calls no fields", func(t *testing.T) {
		called := false
		err := EachExp(emptyStruct{}, func(name string, value any) bool {
			called = true
			return true
		})
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if called {
			t.Errorf("callback should not be called for empty struct")
		}
	})

	t.Run("iterates over exported fields only", func(t *testing.T) {
		val := 42
		obj := testStruct{
			Name:    "Alice",
			Age:     30,
			private: true,
			Value:   &val,
		}

		var calledFields []string
		err := EachExp(obj, func(name string, value any) bool {
			calledFields = append(calledFields, name)
			return true
		})
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		expected := []string{"Name", "Age", "Value"}
		if len(calledFields) != len(expected) {
			t.Errorf("expected %d fields, got %d: %v", len(expected), len(calledFields), calledFields)
		}
		for i, exp := range expected {
			if calledFields[i] != exp {
				t.Errorf("field %d: expected %s, got %s", i, exp, calledFields[i])
			}
		}
	})

	t.Run("stops early when fn returns false", func(t *testing.T) {
		obj := testStruct{Name: "stop", Age: 99, private: false}
		count := 0
		err := EachExp(obj, func(name string, value any) bool {
			count++
			return count < 1 // stop after first field
		})
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if count != 1 {
			t.Errorf("expected 1 call, got %d", count)
		}
	})

	t.Run("handles pointer to struct", func(t *testing.T) {
		obj := &testStruct{Name: "Bob", Age: 25}
		fields := 0
		err := EachExp(obj, func(name string, value any) bool {
			fields++
			return true
		})
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if fields != 3 { // Name, Age, Value (Value is nil, but still exported)
			t.Errorf("expected 3 fields, got %d", fields)
		}
	})

	t.Run("nil input returns error", func(t *testing.T) {
		var nilStruct *testStruct = nil
		err := EachExp(nilStruct, func(name string, value any) bool {
			t.Errorf("callback should not be called")
			return true
		})
		if err == nil {
			t.Error("expected error for nil struct pointer")
		}
	})
}

func TestFromExp(t *testing.T) {
	t.Run("non-struct returns error", func(t *testing.T) {
		_, err := FromExp("not a struct")
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})

	t.Run("nil struct pointer returns error", func(t *testing.T) {
		var s *testStruct
		_, err := FromExp(s)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})

	t.Run("empty struct returns empty slice", func(t *testing.T) {
		fields, err := FromExp(emptyStruct{})
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if len(fields) != 0 {
			t.Errorf("expected empty slice, got %v", fields)
		}
	})

	t.Run("returns exported fields only", func(t *testing.T) {
		val := 42
		obj := testStruct{
			Name:    "Alice",
			Age:     30,
			private: true,
			Value:   &val,
		}

		fields, err := FromExp(obj)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if len(fields) != 3 {
			t.Errorf("expected 3 fields, got %d", len(fields))
		}

		if fields[0] != "Alice" {
			t.Errorf("expected Name, got %v", fields[0])
		}
		if fields[1] != 30 {
			t.Errorf("expected Age, got %v", fields[1])
		}
		if fields[2] != &val {
			t.Errorf("expected Value pointer, got %v", fields[2])
		}
	})

	t.Run("handles pointer to struct", func(t *testing.T) {
		obj := &testStruct{Name: "Bob", Age: 25}

		fields, err := FromExp(obj)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if len(fields) != 3 {
			t.Errorf("expected 3 fields, got %d", len(fields))
		}
	})

	t.Run("nil exported field is included", func(t *testing.T) {
		obj := testStruct{Name: "Nil", Age: 1, Value: nil}

		fields, err := FromExp(obj)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		v, ok := fields[2].(*int)
		if !ok {
			t.Errorf("expected *int, got %T", fields[2])
		}
		if v != nil {
			t.Errorf("expected nil *int, got %v", v)
		}
	})
}

func TestGetExp(t *testing.T) {
	t.Run("should get exported field value correctly", func(t *testing.T) {
		x := 42
		s := manyFields{
			A: 10,
			B: &x,
			C: []int{1, 2, 3},
			D: map[string]int{"a": 1},
			E: inner{X: 7},
			f: "secret", // unexported — should not be accessible
		}

		// Test A
		val, err := GetExp(s, "A")
		if err != nil {
			t.Fatalf("unexpected error getting A: %v", err)
		}
		if a, ok := val.(int); !ok || a != 10 {
			t.Errorf("A: expected 10, got %v (type %T)", val, val)
		}

		// Test B (pointer)
		val, err = GetExp(s, "B")
		if err != nil {
			t.Fatalf("unexpected error getting B: %v", err)
		}
		if b, ok := val.(*int); !ok {
			t.Errorf("B: expected *int, got %T", val)
		} else if *b != 42 {
			t.Errorf("B: expected 42, got %d", *b)
		}

		// Test C (slice)
		val, err = GetExp(s, "C")
		if err != nil {
			t.Fatalf("unexpected error getting C: %v", err)
		}
		if c, ok := val.([]int); !ok {
			t.Errorf("C: expected []int, got %T", val)
		} else if len(c) != 3 || c[0] != 1 {
			t.Errorf("C: expected [1,2,3], got %v", c)
		}

		// Test D (map)
		val, err = GetExp(s, "D")
		if err != nil {
			t.Fatalf("unexpected error getting D: %v", err)
		}
		if d, ok := val.(map[string]int); !ok {
			t.Errorf("D: expected map[string]int, got %T", val)
		} else if d["a"] != 1 {
			t.Errorf("D: expected map[a:1], got %v", d)
		}

		// Test E (nested struct)
		val, err = GetExp(s, "E")
		if err != nil {
			t.Fatalf("unexpected error getting E: %v", err)
		}
		if e, ok := val.(inner); !ok {
			t.Errorf("E: expected inner, got %T", val)
		} else if e.X != 7 {
			t.Errorf("E.X: expected 7, got %d", e.X)
		}
	})

	t.Run("should return error for unexported field", func(t *testing.T) {
		s := manyFields{f: "hidden"}
		_, err := GetExp(s, "f")
		if err == nil {
			t.Errorf("expected error accessing unexported field 'f', got nil")
		}
	})

	t.Run("should return error for unknown field", func(t *testing.T) {
		s := manyFields{A: 1}
		_, err := GetExp(s, "NoSuchField")
		if err == nil {
			t.Errorf("expected error for unknown field, got nil")
		}
	})

	t.Run("non-struct target returns error", func(t *testing.T) {
		_, err := GetExp(123, "A")
		if err == nil {
			t.Errorf("expected error for non-struct input, got nil")
		}
	})

	t.Run("nil input returns error", func(t *testing.T) {
		var s *manyFields = nil
		_, err := GetExp(s, "A")
		if err == nil {
			t.Errorf("expected error for nil input, got nil")
		}
	})

	t.Run("pointer to struct works", func(t *testing.T) {
		x := 99
		s := &manyFields{B: &x}
		val, err := GetExp(s, "B")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if b, ok := val.(*int); !ok {
			t.Errorf("B: expected *int, got %T", val)
		} else if *b != 99 {
			t.Errorf("B: expected 99, got %d", *b)
		}
	})
}

func TestHas(t *testing.T) {
	t.Run("should return true for existing exported field", func(t *testing.T) {
		s := manyFields{A: 1}
		ok, err := Has(s, "A")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !ok {
			t.Errorf("expected true for field 'A', got false")
		}
	})

	t.Run("should return true for unexported field", func(t *testing.T) {
		s := manyFields{f: "secret"}
		ok, err := Has(s, "f")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !ok {
			t.Errorf("expected true for unexported field 'f', got false")
		}
	})

	t.Run("should return false for non-existent field", func(t *testing.T) {
		s := manyFields{A: 1}
		ok, err := Has(s, "NonExistent")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if ok {
			t.Errorf("expected false for unknown field, got true")
		}
	})

	t.Run("should work with pointer to struct", func(t *testing.T) {
		s := &manyFields{A: 1}
		ok, err := Has(s, "A")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !ok {
			t.Errorf("expected true for field 'A' via pointer, got false")
		}
	})

	t.Run("non-struct input returns error", func(t *testing.T) {
		ok, err := Has(123, "A")
		if err == nil {
			t.Errorf("expected error for non-struct input, got nil")
		}
		if ok {
			t.Errorf("ok should be false when error occurs")
		}
	})

	t.Run("nil struct pointer should return true", func(t *testing.T) {
		var s *manyFields = nil
		ok, err := Has(s, "A")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if !ok {
			t.Errorf("field should be found")
		}
	})

	t.Run("empty struct returns false for any field", func(t *testing.T) {
		type Empty struct{}
		ok, err := Has(Empty{}, "Anything")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if ok {
			t.Errorf("expected false for field on empty struct, got true")
		}
	})
}

func TestNames(t *testing.T) {
	t.Run("should return all field names including unexported", func(t *testing.T) {
		s := manyFields{
			A: 1,
			B: new(int),
			C: []int{},
			D: map[string]int{},
			E: inner{},
			f: "secret",
		}

		names, err := Names(s)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expected := []string{"A", "B", "C", "D", "E", "f"}
		if len(names) != len(expected) {
			t.Fatalf("expected %d names, got %d: %v", len(expected), len(names), names)
		}

		// Since field order is guaranteed by Go spec (source order), we can compare directly
		for i, exp := range expected {
			if names[i] != exp {
				t.Errorf("names[%d]: expected %q, got %q", i, exp, names[i])
			}
		}
	})

	t.Run("empty struct returns empty slice", func(t *testing.T) {
		type Empty struct{}
		names, err := Names(Empty{})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(names) != 0 {
			t.Errorf("expected empty slice, got %v", names)
		}
	})

	t.Run("pointer to struct works", func(t *testing.T) {
		s := &manyFields{A: 1}
		names, err := Names(s)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(names) == 0 {
			t.Errorf("expected non-empty field list")
		}
		// At least "A" should be present
		found := false
		for _, name := range names {
			if name == "A" {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("field 'A' not found in names: %v", names)
		}
	})

	t.Run("non-struct input returns error", func(t *testing.T) {
		_, err := Names(123)
		if err == nil {
			t.Errorf("expected error for non-struct input, got nil")
		}
	})

	t.Run("nil struct pointer returns error", func(t *testing.T) {
		var s *manyFields = nil
		_, err := Names(s)
		if err == nil {
			t.Errorf("expected error for nil input, got nil")
		}
	})
}

func TestNil(t *testing.T) {
	t.Run("should return names of nil fields", func(t *testing.T) {
		s := manyFields{
			A: 10,
			B: nil, // *int
			C: nil, // []int
			D: nil, // map[string]int
			E: inner{X: 7},
			f: "secret",
		}

		names, err := Nil(s)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expected := map[string]bool{"B": true, "C": true, "D": true}
		if len(names) != len(expected) {
			t.Fatalf("expected %d nil fields, got %d: %v", len(expected), len(names), names)
		}

		for _, name := range names {
			if !expected[name] {
				t.Errorf("unexpected nil field: %q", name)
			}
			delete(expected, name)
		}
		if len(expected) > 0 {
			t.Errorf("missing nil fields: %v", expected)
		}
	})

	t.Run("no nil fields returns empty slice", func(t *testing.T) {
		x := 42
		s := manyFields{
			A: 10,
			B: &x,
			C: []int{1, 2, 3},
			D: map[string]int{"a": 1},
			E: inner{X: 7},
			f: "secret",
		}

		names, err := Nil(s)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(names) != 0 {
			t.Errorf("expected no nil fields, got: %v", names)
		}
	})

	t.Run("unexported nil field is included", func(t *testing.T) {
		type withUnexported struct {
			Exported   *int
			unexported []string
		}

		s := withUnexported{
			Exported:   nil,
			unexported: nil,
		}

		names, err := Nil(s)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(names) != 2 {
			t.Errorf("expected 2 names, got %v", names)
		}
		if names[0] != "Exported" {
			t.Errorf("expected Exported field, got %q", names[0])
		}
		if names[1] != "unexported" {
			t.Errorf("expected unexported field, got %q", names[1])
		}
	})

	t.Run("non-struct input returns error", func(t *testing.T) {
		_, err := Nil(123)
		if err == nil {
			t.Errorf("expected error for non-struct input, got nil")
		}
	})

	t.Run("nil struct pointer returns error", func(t *testing.T) {
		var s *manyFields = nil
		_, err := Nil(s)
		if err == nil {
			t.Errorf("expected error for nil input, got nil")
		}
	})

	t.Run("struct with zero-value non-nilable fields", func(t *testing.T) {
		s := manyFields{
			A: 0,
			B: new(int),
			C: []int{},
			D: make(map[string]int),
			E: inner{},
			f: "",
		}

		names, err := Nil(s)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(names) != 0 {
			t.Errorf("expected no nil fields, got: %v", names)
		}
	})
}

func TestNotNil(t *testing.T) {
	t.Run("should return only exported non-nil fields", func(t *testing.T) {
		x := 42
		s := manyFields{
			A: 10,
			B: &x,
			C: []int{1, 2, 3},
			D: map[string]int{"a": 1},
			E: inner{X: 7},
			f: "secret",
		}

		fields, err := NotNilExp(s)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		// Expected: A, B, C, D, E → 5 fields
		if len(fields) != 5 {
			t.Fatalf("expected 5 non-nil exported fields, got %d: %v", len(fields), fields)
		}

		if val, ok := fields[0].(int); !ok || val != 10 {
			t.Errorf("fields[0] (A): expected 10, got %v", fields[0])
		}
		if val, ok := fields[1].(*int); !ok || *val != 42 {
			t.Errorf("fields[1] (B): expected *int=42, got %v", fields[1])
		}
		if val, ok := fields[2].([]int); !ok || len(val) != 3 || val[0] != 1 {
			t.Errorf("fields[2] (C): expected [1,2,3], got %v", fields[2])
		}
		if val, ok := fields[3].(map[string]int); !ok || val["a"] != 1 {
			t.Errorf("fields[3] (D): expected map[a:1], got %v", fields[3])
		}
		if val, ok := fields[4].(inner); !ok || val.X != 7 {
			t.Errorf("fields[4] (E): expected inner{X:7}, got %v", fields[4])
		}

		for i, f := range fields {
			if f == "secret" {
				t.Errorf("unexported field 'f' incorrectly included at index %d", i)
			}
		}
	})

	t.Run("nil exported fields are excluded", func(t *testing.T) {
		s := manyFields{
			A: 0,
			B: nil,
			C: nil,
			D: nil,
			E: inner{},
			f: "not nil but unexported",
		}

		fields, err := NotNilExp(s)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		// Included: A (0), E (inner{}) → 2 fields
		// B, C, D are nil → excluded
		// f is unexported → excluded
		if len(fields) != 2 {
			t.Errorf("expected 2 fields (A, E), got %d: %v", len(fields), fields)
		}

		// A should be 0
		if a, ok := fields[0].(int); !ok || a != 0 {
			t.Errorf("A: expected 0, got %v", fields[0])
		}
		// E should be inner{}
		if e, ok := fields[1].(inner); !ok || e.X != 0 {
			t.Errorf("E: expected inner{X:0}, got %v", fields[1])
		}
	})

	t.Run("all exported fields nil: empty result", func(t *testing.T) {
		type testStruct struct {
			Ptr   *int
			Slice []string
			Map   map[string]int
			unexp int // unexported, ignored
		}

		s := testStruct{
			Ptr:   nil,
			Slice: nil,
			Map:   nil,
			unexp: 1,
		}

		fields, err := NotNilExp(s)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(fields) != 0 {
			t.Errorf("expected empty slice (all exported fields nil), got %v", fields)
		}
	})

	t.Run("non-struct input returns error", func(t *testing.T) {
		_, err := NotNilExp(123)
		if err == nil {
			t.Errorf("expected error for non-struct input, got nil")
		}
	})

	t.Run("nil struct pointer returns error", func(t *testing.T) {
		var s *manyFields = nil
		_, err := NotNilExp(s)
		if err == nil {
			t.Errorf("expected error for nil input, got nil")
		}
	})

	t.Run("empty struct returns empty slice", func(t *testing.T) {
		type Empty struct{}
		fields, err := NotNilExp(Empty{})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(fields) != 0 {
			t.Errorf("expected empty slice for empty struct, got %v", fields)
		}
	})
}

func TestSetExo(t *testing.T) {
	t.Run("should set exported field correctly", func(t *testing.T) {
		x := 99
		s := manyFields{
			A: 10,
			B: &x,
			C: []int{1, 2},
			D: map[string]int{"old": 0},
			E: inner{X: 5},
			f: "original", // unexported, should remain unchanged
		}

		// SetExp A
		if err := SetExp(&s, "A", 42); err != nil {
			t.Fatalf("failed to set A: %v", err)
		}
		if s.A != 42 {
			t.Errorf("A: expected 42, got %d", s.A)
		}

		// SetExp B (pointer)
		y := 88
		if err := SetExp(&s, "B", &y); err != nil {
			t.Fatalf("failed to set B: %v", err)
		}
		if *s.B != 88 {
			t.Errorf("B: expected 88, got %d", *s.B)
		}

		// SetExp C (slice)
		newSlice := []int{10, 20, 30}
		if err := SetExp(&s, "C", newSlice); err != nil {
			t.Fatalf("failed to set C: %v", err)
		}
		if len(s.C) != 3 || s.C[0] != 10 {
			t.Errorf("C: expected [10,20,30], got %v", s.C)
		}

		// SetExp D (map)
		newMap := map[string]int{"new": 100}
		if err := SetExp(&s, "D", newMap); err != nil {
			t.Fatalf("failed to set D: %v", err)
		}
		if s.D["new"] != 100 {
			t.Errorf("D: expected map[new:100], got %v", s.D)
		}

		// SetExp E (nested struct)
		newInner := inner{X: 999}
		if err := SetExp(&s, "E", newInner); err != nil {
			t.Fatalf("failed to set E: %v", err)
		}
		if s.E.X != 999 {
			t.Errorf("E.X: expected 999, got %d", s.E.X)
		}

		// Ensure unexported field unchanged
		if s.f != "original" {
			t.Errorf("unexported field 'f' was modified; expected 'original', got %q", s.f)
		}
	})

	t.Run("should return error for unexported field", func(t *testing.T) {
		s := manyFields{f: "old"}
		err := SetExp(&s, "f", "new")
		if err == nil {
			t.Errorf("expected error for unexported field 'f', got nil")
		}
		if s.f != "old" {
			t.Errorf("unexported field was modified despite error")
		}
	})

	t.Run("should return error for unknown field", func(t *testing.T) {
		s := manyFields{A: 1}
		err := SetExp(&s, "NoSuchField", 123)
		if err == nil {
			t.Errorf("expected error for unknown field, got nil")
		}
	})

	t.Run("should return error for type mismatch", func(t *testing.T) {
		s := manyFields{A: 1}
		err := SetExp(&s, "A", "this is a string") // A is int
		if err == nil {
			t.Errorf("expected type mismatch error, got nil")
		}
		if s.A != 1 {
			t.Errorf("field was modified despite type error")
		}
	})

	t.Run("should return error for non-struct target", func(t *testing.T) {
		err := SetExp(123, "A", 1)
		if err == nil {
			t.Errorf("expected error for non-struct input, got nil")
		}
	})

	t.Run("should return error for nil value", func(t *testing.T) {
		s := manyFields{B: new(int)}
		var nilPtr *int = nil
		err := SetExp(&s, "B", nilPtr)
		if err != nil {
			t.Fatalf("setting *int to nil should be allowed (nil is valid for pointer)")
		}
		if s.B != nil {
			t.Errorf("B should be nil after SetExp")
		}
	})

	t.Run("should return error for invalid value (nil to int)", func(t *testing.T) {
		s := manyFields{A: 5}
		// reflect.ValueOf(nil) is invalid
		err := SetExp(&s, "A", nil)
		if err == nil {
			t.Errorf("expected error when setting int field to nil, got nil")
		}
		if s.A != 5 {
			t.Errorf("field was modified despite invalid value")
		}
	})

	t.Run("must use pointer to struct", func(t *testing.T) {
		s := manyFields{A: 1}
		// Pass by value — reflect.Value will be unaddressable
		err := SetExp(s, "A", 99)
		if err == nil {
			t.Errorf("expected error when passing struct by value (non-addressable), got nil")
		} else {
			t.Log(err)
		}
	})
}

type taggedStruct struct {
	ExportedField   string `json:"exported" db:"exp" validate:"required"`
	unexportedField int    `json:"unexported" db:"unexp"`
	NoTagField      bool
	EmptyTagField   float64 `json:"" db:""`
	IgnoredField    string  `other:"value"`
}

func TestTags(t *testing.T) {
	t.Run("should return correct tags for given key", func(t *testing.T) {
		s := taggedStruct{}

		// "json" tags
		jsonTags, err := Tags(s, "json")
		if err != nil {
			t.Fatalf("unexpected error for 'json' tags: %v", err)
		}

		expectedJSON := map[string]string{
			"ExportedField":   "exported",
			"unexportedField": "unexported",
			"NoTagField":      "",
			"EmptyTagField":   "",
			"IgnoredField":    "",
		}

		if len(jsonTags) != len(expectedJSON) {
			t.Fatalf("json: expected %d fields, got %d: %v", len(expectedJSON), len(jsonTags), jsonTags)
		}

		for name, expectedTag := range expectedJSON {
			if actual, ok := jsonTags[name]; !ok {
				t.Errorf("json: missing field '%s'", name)
			} else if actual != expectedTag {
				t.Errorf("json: field '%s' expected tag %q, got %q", name, expectedTag, actual)
			}
		}

		// "db" tags
		dbTags, err := Tags(s, "db")
		if err != nil {
			t.Fatalf("unexpected error for 'db' tags: %v", err)
		}

		expectedDB := map[string]string{
			"ExportedField":   "exp",
			"unexportedField": "unexp",
			"NoTagField":      "",
			"EmptyTagField":   "",
			"IgnoredField":    "",
		}

		if len(dbTags) != len(expectedDB) {
			t.Fatalf("db: expected %d fields, got %d: %v", len(expectedDB), len(dbTags), dbTags)
		}

		for name, expectedTag := range expectedDB {
			if actual, ok := dbTags[name]; !ok {
				t.Errorf("db: missing field '%s'", name)
			} else if actual != expectedTag {
				t.Errorf("db: field '%s' expected tag %q, got %q", name, expectedTag, actual)
			}
		}

		// unknown tag key should return all empty
		otherTags, err := Tags(s, "nonexistent")
		if err != nil {
			t.Fatalf("unexpected error for unknown tag key: %v", err)
		}

		expectedOther := map[string]string{
			"ExportedField":   "",
			"unexportedField": "",
			"NoTagField":      "",
			"EmptyTagField":   "",
			"IgnoredField":    "",
		}

		if len(otherTags) != len(expectedOther) {
			t.Errorf("nonexistent: wrong field count")
		}
		for name := range expectedOther {
			if otherTags[name] != "" {
				t.Errorf("nonexistent: field '%s' should have empty tag, got %q", name, otherTags[name])
			}
		}
	})

	t.Run("includes unexported fields", func(t *testing.T) {
		s := taggedStruct{}
		tags, err := Tags(s, "json")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if _, hasUnexported := tags["unexportedField"]; !hasUnexported {
			t.Errorf("expected unexported field 'unexportedField' in tags, but it's missing")
		}
	})

	t.Run("empty struct returns empty map", func(t *testing.T) {
		type Empty struct{}
		tags, err := Tags(Empty{}, "json")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(tags) != 0 {
			t.Errorf("expected empty tag map, got %v", tags)
		}
	})

	t.Run("pointer to struct works", func(t *testing.T) {
		s := &taggedStruct{}
		tags, err := Tags(s, "json")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(tags) == 0 {
			t.Errorf("expected non-empty tag map from pointer")
		}
		if _, ok := tags["ExportedField"]; !ok {
			t.Errorf("missing 'ExportedField' when passing pointer")
		}
	})

	t.Run("non-struct input returns error", func(t *testing.T) {
		_, err := Tags(123, "json")
		if err == nil {
			t.Errorf("expected error for non-struct input, got nil")
		}
	})

	t.Run("nil struct pointer returns error", func(t *testing.T) {
		var s *taggedStruct = nil
		var nilInterface any = nil
		_, err := Tags(nilInterface, "json")
		if err == nil {
			t.Errorf("expected error for nil input, got nil")
		}

		// test typed nil pointer
		_, err = Tags(s, "json") // s is *taggedStruct = nil
		if err != nil {
			t.Errorf("unexpected error for typed nil pointer: %v", err)
		}
	})
}
