// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package field4

import (
	"testing"
	"unsafe"
)

func TestUnsafeEach(t *testing.T) {

	t.Run("non-struct returns error", func(t *testing.T) {
		err := UnsafeEach("not a struct", func(name string, value any) bool {
			t.Errorf("callback should not be called")
			return true
		})
		if err == nil {
			t.Fatalf("expected error, got nil")
		}
	})

	t.Run("nil struct pointer returns error", func(t *testing.T) {
		var s *testStruct
		err := UnsafeEach(s, func(name string, value any) bool {
			t.Errorf("callback should not be called")
			return true
		})
		if err == nil {
			t.Fatalf("expected error, got nil")
		}
	})

	t.Run("empty struct calls no fields", func(t *testing.T) {
		called := false
		err := UnsafeEach(emptyStruct{}, func(name string, value any) bool {
			called = true
			return true
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if called {
			t.Fatalf("callback should not be called")
		}
	})

	t.Run("iterates over exported and unexported fields", func(t *testing.T) {
		val := 42
		obj := testStruct{
			Name:    "Alice",
			Age:     30,
			private: true,
			Value:   &val,
		}

		got := make(map[string]any)

		err := UnsafeEach(obj, func(name string, value any) bool {
			got[name] = value
			return true
		})

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if len(got) != 4 {
			t.Fatalf("expected 4 fields, got %d", len(got))
		}

		if got["private"] != true {
			t.Errorf("expected private=true, got %v", got["private"])
		}
	})

	t.Run("stops early when fn returns false", func(t *testing.T) {
		obj := testStruct{Name: "stop", Age: 99, private: true}

		count := 0
		err := UnsafeEach(obj, func(name string, value any) bool {
			count++
			return false
		})

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if count != 1 {
			t.Fatalf("expected 1 call, got %d", count)
		}
	})

	t.Run("handles pointer to struct", func(t *testing.T) {
		obj := &testStruct{Name: "Bob", Age: 25}

		fields := 0
		err := UnsafeEach(obj, func(name string, value any) bool {
			fields++
			return true
		})

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if fields != 4 {
			t.Fatalf("expected 4 fields, got %d", fields)
		}
	})
}

type unsafeTestStruct struct {
	Exported   int
	unexported string
}

func TestUnsafeSet(t *testing.T) {
	t.Run("Exported field", func(t *testing.T) {
		obj := &unsafeTestStruct{Exported: 1}

		err := UnsafeSet(obj, "Exported", 42)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if obj.Exported != 42 {
			t.Errorf("expected Exported=42, got %d", obj.Exported)
		}
	})

	t.Run("unexported field", func(t *testing.T) {
		obj := &unsafeTestStruct{unexported: "old"}

		err := UnsafeSet(obj, "unexported", "new")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if obj.unexported != "new" {
			t.Errorf("expected unexported='new', got %q", obj.unexported)
		}
	})

	t.Run("non-addressable", func(t *testing.T) {
		obj := unsafeTestStruct{Exported: 1}

		err := UnsafeSet(obj, "Exported", 10)
		if err == nil {
			t.Error("expected error for non-addressable target")
		}
	})

	t.Run("field not found", func(t *testing.T) {
		obj := &unsafeTestStruct{}

		err := UnsafeSet(obj, "DoesNotExist", 1)
		if err == nil {
			t.Error("expected error for missing field")
		}
	})

	t.Run("not struct", func(t *testing.T) {
		var x int
		err := UnsafeSet(&x, "Anything", 1)
		if err == nil {
			t.Error("expected error for non-struct target")
		}
	})

	t.Run("type mismatch", func(t *testing.T) {
		obj := &unsafeTestStruct{}

		err := UnsafeSet(obj, "Exported", "not an int")
		if err == nil {
			t.Error("expected type mismatch error")
		}
	})

	t.Run("nil assignment", func(t *testing.T) {
		obj := &unsafeTestStruct{unexported: "x"}

		err := UnsafeSet(obj, "unexported", nil)
		if err == nil {
			t.Error("expected error assigning nil to string")
		}
	})
}

type unsafeGetStruct struct {
	A int
	b int
	P *int
}

func TestUnsafeGet(t *testing.T) {
	t.Run("exported value field", func(t *testing.T) {
		obj := &unsafeGetStruct{A: 10}

		v, err := UnsafeGet(obj, "A")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if v.(int) != 10 {
			t.Errorf("expected 10, got %v", v)
		}
	})

	t.Run("unexported value field", func(t *testing.T) {
		obj := unsafeGetStruct{b: 20}

		v, err := UnsafeGet(obj, "b")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if v.(int) != 20 {
			t.Errorf("expected 20, got %v", v)
		}
	})

	t.Run("pointer field preserves address", func(t *testing.T) {
		x := 42
		obj := unsafeGetStruct{P: &x}

		v, err := UnsafeGet(obj, "P")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		p := v.(*int)

		if p != &x {
			t.Errorf("expected same pointer address")
		}
	})

	t.Run("nil pointer field returns nil", func(t *testing.T) {
		obj := unsafeGetStruct{P: nil}

		get, err := UnsafeGet(obj, "P")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		v, ok := get.(*int)
		if !ok {
			t.Errorf("expected pointer to int")
		}
		if v != nil {
			t.Errorf("expected nil, got '%v'", v)
		}
	})

	t.Run("non-struct target returns error", func(t *testing.T) {
		_, err := UnsafeGet(123, "A")
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})

	t.Run("missing field returns error", func(t *testing.T) {
		obj := unsafeGetStruct{A: 1}
		_, err := UnsafeGet(obj, "DoesNotExist")
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})

	t.Run("non-addressable struct still returns correct value", func(t *testing.T) {
		// Explicitly non-addressable temporary value
		v, err := UnsafeGet(unsafeGetStruct{A: 77}, "A")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if v.(int) != 77 {
			t.Fatalf("expected 77, got %v", v)
		}
	})
}

type inner struct {
	X int
}

type manyFields struct {
	A int
	B *int
	C []int
	D map[string]int
	E inner
	f string // unexported
}

func TestUnsafeGetAll(t *testing.T) {
	t.Run("should get all fields correctly", func(t *testing.T) {
		x := 42
		s := manyFields{
			A: 10,
			B: &x,
			C: []int{1, 2, 3},
			D: map[string]int{"a": 1},
			E: inner{X: 7},
			f: "secret",
		}

		out, err := UnsafeGetAll(s)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		// Check A
		if a, ok := out["A"].(int); !ok || a != 10 {
			t.Errorf("A: expected 10, got %v (type %T)", out["A"], out["A"])
		}

		// Check B (pointer)
		if b, ok := out["B"].(*int); !ok {
			t.Errorf("B: expected *int, got %T", out["B"])
		} else if *b != 42 {
			t.Errorf("B: expected 42, got %d", *b)
		}

		// Check C (slice)
		if c, ok := out["C"].([]int); !ok {
			t.Errorf("C: expected []int, got %T", out["C"])
		} else {
			expectedC := []int{1, 2, 3}
			if len(c) != len(expectedC) {
				t.Errorf("C: length mismatch: expected %v, got %v", expectedC, c)
			} else {
				for i := range c {
					if c[i] != expectedC[i] {
						t.Errorf("C[%d]: expected %d, got %d", i, expectedC[i], c[i])
					}
				}
			}
		}

		// Check D (map)
		if d, ok := out["D"].(map[string]int); !ok {
			t.Errorf("D: expected map[string]int, got %T", out["D"])
		} else {
			if len(d) != 1 || d["a"] != 1 {
				t.Errorf("D: expected map[a:1], got %v", d)
			}
		}

		// Check E (nested struct)
		if e, ok := out["E"].(inner); !ok {
			t.Errorf("E: expected inner, got %T", out["E"])
		} else if e.X != 7 {
			t.Errorf("E.X: expected 7, got %d", e.X)
		}

		// Check unexported field f
		if f, ok := out["f"].(string); !ok {
			t.Errorf("f: expected string, got %T", out["f"])
		} else if f != "secret" {
			t.Errorf("f: expected 'secret', got %q", f)
		}

		// no extra fields
		expectedKeys := []string{"A", "B", "C", "D", "E", "f"}
		if len(out) != len(expectedKeys) {
			t.Errorf("expected %d fields, got %d: %v", len(expectedKeys), len(out), out)
		}
	})

	t.Run("pointer identity", func(t *testing.T) {
		x := 100
		s := manyFields{B: &x}
		out, err := UnsafeGetAll(s)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		b := out["B"]
		value, ok := b.(*int)
		if !ok {
			t.Errorf("B: expected *int, got '%T'", b)
		}
		if value != &x {
			t.Errorf("B: expected %p, got %p", &x, value)
		}
	})

	t.Run("shared backing array", func(t *testing.T) {
		// unaddressable struct should force addressability
		s := manyFields{
			C: []int{1, 2, 3},
		}

		out, err := UnsafeGetAll(s)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		slice := out["C"].([]int)
		slice[0] = 99

		// the original struct should reflect the change
		if 99 != s.C[0] {
			t.Errorf("C[0]: expected %d, got %d", s.C[0], slice[0])
		}
	})

	t.Run("map identity", func(t *testing.T) {
		s := manyFields{
			D: map[string]int{"a": 1},
		}

		out, err := UnsafeGetAll(s)
		if err != nil {
			t.Errorf("unexpected error: '%v'", err)
		}

		m := out["D"].(map[string]int)
		m["b"] = 2

		if 2 != s.D["b"] {
			t.Errorf("D: expected %d, got %d", s.D["b"], m["b"])
		}
	})

	t.Run("not struct", func(t *testing.T) {
		_, err := UnsafeGetAll(123)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})
}

func TestUnsafeGetf(t *testing.T) {
	t.Run("should get specified fields correctly", func(t *testing.T) {
		x := 42
		s := manyFields{
			A: 10,
			B: &x,
			C: []int{1, 2, 3},
			D: map[string]int{"a": 1},
			E: inner{X: 7},
			f: "secret",
		}

		out, err := UnsafeGetf(s, "A", "C", "f")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		// Should only have requested fields
		if len(out) != 3 {
			t.Fatalf("expected 3 fields, got %d: %v", len(out), out)
		}

		// Check A
		if a, ok := out["A"].(int); !ok || a != 10 {
			t.Errorf("A: expected 10, got %v (type %T)", out["A"], out["A"])
		}

		// Check C
		if c, ok := out["C"].([]int); !ok {
			t.Errorf("C: expected []int, got %T", out["C"])
		} else if len(c) != 3 || c[0] != 1 || c[1] != 2 || c[2] != 3 {
			t.Errorf("C: expected [1,2,3], got %v", c)
		}

		// Check unexported f
		if f, ok := out["f"].(string); !ok || f != "secret" {
			t.Errorf("f: expected 'secret', got %q (type %T)", f, out["f"])
		}
	})

	t.Run("should handle nested struct field", func(t *testing.T) {
		s := manyFields{E: inner{X: 99}}
		out, err := UnsafeGetf(s, "E")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if len(out) != 1 {
			t.Fatalf("expected 1 field, got %d", len(out))
		}

		e, ok := out["E"].(inner)
		if !ok {
			t.Fatalf("E: expected inner, got %T", out["E"])
		}
		if e.X != 99 {
			t.Errorf("E.X: expected 99, got %d", e.X)
		}
	})

	t.Run("pointer identity preserved", func(t *testing.T) {
		x := 123
		s := manyFields{B: &x}
		out, err := UnsafeGetf(s, "B")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		b, ok := out["B"].(*int)
		if !ok {
			t.Fatalf("B: expected *int, got %T", out["B"])
		}
		if b != &x {
			t.Errorf("B: pointer identity lost; expected %p, got %p", &x, b)
		}
	})

	t.Run("slice and map share backing data", func(t *testing.T) {
		s := manyFields{
			C: []int{10, 20},
			D: map[string]int{"key": 42},
		}

		out, err := UnsafeGetf(s, "C", "D")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		// Mutate slice via result
		c := out["C"].([]int)
		c[0] = 999
		if s.C[0] != 999 {
			t.Errorf("slice backing array not shared: s.C[0] = %d, expected 999", s.C[0])
		}

		// Mutate map via result
		d := out["D"].(map[string]int)
		d["new"] = 888
		if s.D["new"] != 888 {
			t.Errorf("map not shared: s.D['new'] = %d, expected 888", s.D["new"])
		}
	})

	t.Run("requesting unknown field name", func(t *testing.T) {
		s := manyFields{A: 1}
		out, err := UnsafeGetf(s, "A", "NonExistent", "f")
		if err == nil {
			t.Error("expected not found error but got nil")
		}

		// Should ignore unknown field
		if len(out) != 2 {
			t.Errorf("expected 2 fields (A and f), got %d: %v", len(out), out)
		}
		if _, hasA := out["A"]; !hasA {
			t.Errorf("missing A")
		}
		if _, hasF := out["f"]; !hasF {
			t.Errorf("missing f")
		}
		// "NonExistent" should not be present
		if _, hasBad := out["NonExistent"]; hasBad {
			t.Errorf("unexpected field 'NonExistent' in output")
		}
	})

	t.Run("empty field list returns empty map", func(t *testing.T) {
		s := manyFields{A: 1}
		out, err := UnsafeGetf(s)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if len(out) != 0 {
			t.Errorf("expected empty map, got %v", out)
		}
	})

	t.Run("not a struct", func(t *testing.T) {
		_, err := UnsafeGetf(123, "A")
		if err == nil {
			t.Errorf("expected error for non-struct input, got nil")
		}
	})
}

func TestUnsafeAccess(t *testing.T) {
	t.Run("should access exported field and provide value and pointer", func(t *testing.T) {
		x := 42
		s := &manyFields{B: &x}

		var observedValue any
		var observedPtr unsafe.Pointer

		err := UnsafeAccess(s, "B", func(value any, ptr unsafe.Pointer) {
			observedValue = value
			observedPtr = ptr
		})

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		// Check value
		if val, ok := observedValue.(*int); !ok {
			t.Errorf("value: expected *int, got %T", observedValue)
		} else if *val != 42 {
			t.Errorf("value: expected 42, got %d", *val)
		}

		expectedPtr := unsafe.Pointer(&s.B)
		if observedPtr != expectedPtr {
			t.Errorf("pointer mismatch: expected %p, got %p", expectedPtr, observedPtr)
		}
	})

	t.Run("should access unexported field", func(t *testing.T) {
		s := &manyFields{f: "hidden"}

		var captured string
		var ptrCaptured unsafe.Pointer

		err := UnsafeAccess(s, "f", func(value any, ptr unsafe.Pointer) {
			if str, ok := value.(string); ok {
				captured = str
				ptrCaptured = ptr
			}
		})

		if err != nil {
			t.Fatalf("failed to access unexported field 'f': %v", err)
		}

		if captured != "hidden" {
			t.Errorf("expected 'hidden', got %q", captured)
		}

		// Optional: verify ptr corresponds to s.f by writing back (use with caution)
		// Since it's unsafe, we can demonstrate mutability if needed
		*(*string)(ptrCaptured) = "modified"
		if s.f != "modified" {
			t.Errorf("pointer did not point to original field")
		}
	})

	t.Run("should access nested struct field", func(t *testing.T) {
		s := &manyFields{E: inner{X: 999}}

		var val any
		err := UnsafeAccess(s, "E", func(value any, _ unsafe.Pointer) {
			val = value
		})

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if innerVal, ok := val.(inner); !ok {
			t.Errorf("expected inner, got %T", val)
		} else if innerVal.X != 999 {
			t.Errorf("inner.X: expected 999, got %d", innerVal.X)
		}
	})

	t.Run("unknown field name returns error", func(t *testing.T) {
		s := &manyFields{A: 1}
		err := UnsafeAccess(s, "NoSuchField", func(any, unsafe.Pointer) {})
		if err == nil {
			t.Errorf("expected error for unknown field, got nil")
		}
		// Optionally, check for specific error if you have one
	})

	t.Run("non-struct target returns error", func(t *testing.T) {
		err := UnsafeAccess(123, "A", func(any, unsafe.Pointer) {})
		if err == nil {
			t.Errorf("expected error for non-struct input, got nil")
		}
	})

	t.Run("callback can mutate via pointer (slice example)", func(t *testing.T) {
		s := &manyFields{C: []int{1, 2, 3}}

		err := UnsafeAccess(s, "C", func(value any, ptr unsafe.Pointer) {
			// Reinterpret pointer as *[]int
			slicePtr := (*[]int)(ptr)
			(*slicePtr)[0] = 888
		})

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if s.C[0] != 888 {
			t.Errorf("mutation via pointer failed: expected 888, got %d", s.C[0])
		}
	})
}
