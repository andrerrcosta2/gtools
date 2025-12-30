// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package unwrap4

import (
	"reflect"
	"testing"
)

func TestToValue(t *testing.T) {
	t.Run("nil interface returns error", func(t *testing.T) {
		var v any = nil
		_, err := ToValue(v)
		if err == nil {
			t.Errorf("expected error for nil interface, got nil")
		}
	})

	t.Run("direct primitive value", func(t *testing.T) {
		val, err := ToValue(42)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if val.Kind() != reflect.Int || val.Int() != 42 {
			t.Errorf("expected int 42, got %v (kind=%v)", val, val.Kind())
		}
	})

	t.Run("nil pointer", func(t *testing.T) {
		var p *int = nil
		_, err := ToValue(p)
		if err == nil {
			t.Fatal("expected an error, got nil")
		}
	})

	t.Run("non-nil pointer unwraps to value", func(t *testing.T) {
		x := "hello"
		p := &x
		val, err := ToValue(p)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if val.Kind() != reflect.String || val.String() != "hello" {
			t.Errorf("expected string 'hello', got %v", val)
		}
	})

	t.Run("interface holding nil pointer", func(t *testing.T) {
		var p *string = nil
		var i any = p
		_, err := ToValue(i)
		if err == nil {
			t.Fatal("expected an error, got nil")
		}
	})

	t.Run("interface holding concrete value", func(t *testing.T) {
		var i any = []int{1, 2, 3}
		val, err := ToValue(i)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if val.Kind() != reflect.Slice || val.Len() != 3 {
			t.Errorf("expected slice of len 3, got %v", val)
		}
		if val.Index(0).Int() != 1 {
			t.Errorf("expected first element 1, got %v", val.Index(0))
		}
	})

	t.Run("deeply nested interfaces and pointer4", func(t *testing.T) {
		x := 99
		type Wrapper struct{ Value int }
		w := Wrapper{Value: x}
		p := &w
		ww := &p

		var i1 any = &w // any → *Wrapper
		var i2 any = i1 // any → any → *Wrapper
		var pp **Wrapper = ww
		var i3 any = pp // any → **Wrapper

		// Test i2
		val, err := ToValue(i2)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if val.Kind() != reflect.Struct {
			t.Fatalf("expected struct, got %v", val.Kind())
		}
		if v := val.FieldByName("Value").Int(); v != 99 {
			t.Errorf("expected Value=99, got %d", v)
		}

		// Test i3
		val, err = ToValue(i3)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if val.Kind() != reflect.Struct || val.FieldByName("Value").Int() != 99 {
			t.Errorf("deep pointer unwrapping failed")
		}
	})

	t.Run("zero values are valid", func(t *testing.T) {
		testCases := []any{
			"",
			0,
			false,
			[]int(nil),
			map[string]int(nil),
			struct{}{},
		}

		for _, tc := range testCases {
			val, err := ToValue(tc)
			if err != nil {
				t.Errorf("unexpected error for %#v: %v", tc, err)
				continue
			}
			if !val.IsValid() {
				t.Errorf("invalid reflect.Value for %#v", tc)
			}
			// All should be valid and unwrapped
		}
	})

	t.Run("slice and map (non-pointer) are returned as-is", func(t *testing.T) {
		s := []string{"a"}
		val, err := ToValue(s)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if val.Kind() != reflect.Slice || val.Len() != 1 {
			t.Errorf("slice not handled correctly")
		}

		m := map[string]int{"k": 1}
		val, err = ToValue(m)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if val.Kind() != reflect.Map || val.Len() != 1 {
			t.Errorf("map not handled correctly")
		}
	})
}

func TestToTypeValue(t *testing.T) {
	t.Run("nil interface returns error", func(t *testing.T) {
		var v any = nil
		_, err := ToTypeValue(v)
		if err == nil {
			t.Errorf("expected error for nil interface, got nil")
		}
	})

	t.Run("primitive type", func(t *testing.T) {
		typ, err := ToTypeValue(42)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if typ.Kind() != reflect.Int {
			t.Errorf("expected int type, got %v", typ.Kind())
		}
	})

	t.Run("pointer to primitive unwraps to primitive type", func(t *testing.T) {
		var p *string = nil
		typ, err := ToTypeValue(p)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if typ.Kind() != reflect.String {
			t.Errorf("expected string type, got %v", typ.Kind())
		}
	})

	t.Run("double pointer unwraps fully", func(t *testing.T) {
		var x int = 5
		p := &x
		var pp **int = &p
		typ, err := ToTypeValue(pp)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if typ.Kind() != reflect.Int {
			t.Errorf("expected int, got %v", typ.Kind())
		}
	})

	t.Run("interface holding value unwraps to concrete type", func(t *testing.T) {
		var i any = []int{1, 2}
		typ, err := ToTypeValue(i)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if typ.Kind() != reflect.Slice || typ.Elem().Kind() != reflect.Int {
			t.Errorf("expected []int, got %v", typ)
		}
	})

	t.Run("interface holding pointer unwraps fully", func(t *testing.T) {
		type Person struct{ Name string }
		p := &Person{Name: "Alice"}
		var i any = p
		typ, err := ToTypeValue(i)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if typ.Kind() != reflect.Struct || typ.Name() != "Person" {
			t.Errorf("expected Person struct, got %v", typ)
		}
	})

	t.Run("nested interfaces and pointer4", func(t *testing.T) {
		var x float64 = 3.14
		var i1 any = &x
		var i2 any = i1

		typ, err := ToTypeValue(i2)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if typ.Kind() != reflect.Float64 {
			t.Errorf("expected float64, got %v", typ.Kind())
		}
	})

	t.Run("struct type", func(t *testing.T) {
		type Data struct{ ID int }
		d := Data{ID: 1}
		typ, err := ToTypeValue(d)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if typ.Kind() != reflect.Struct || typ.Name() != "Data" {
			t.Errorf("expected Data struct, got %v", typ)
		}
	})

	t.Run("slice and map type4", func(t *testing.T) {
		// Slice
		typ, err := ToTypeValue([]string{"a"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if typ.Kind() != reflect.Slice || typ.Elem().Kind() != reflect.String {
			t.Errorf("expected []string, got %v", typ)
		}

		// Map
		typ, err = ToTypeValue(map[int]bool{1: true})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if typ.Kind() != reflect.Map {
			t.Errorf("expected map, got %v", typ.Kind())
		}
		if typ.Key().Kind() != reflect.Int || typ.Elem().Kind() != reflect.Bool {
			t.Errorf("expected map[int]bool, got key=%v, elem=%v", typ.Key().Kind(), typ.Elem().Kind())
		}
	})

	t.Run("typed nil pointer is valid", func(t *testing.T) {
		var p *struct{ X int } = nil
		typ, err := ToTypeValue(p)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if typ.Kind() != reflect.Struct {
			t.Errorf("expected struct, got %v", typ.Kind())
		}
	})

	t.Run("zero interface{} holding nil is rejected", func(t *testing.T) {
		var v any
		_, err := ToTypeValue(v) // nil interface
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})
}
