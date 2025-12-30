// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package type4

import "testing"

func TestHasDepth(t *testing.T) {
	type Person struct {
		Name string
	}

	t.Run("direct struct returns true", func(t *testing.T) {
		v := Person{Name: "Alice"}
		ok, err := HasDepth(v)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !ok {
			t.Errorf("expected true for direct struct, got false")
		}
	})

	t.Run("pointer to struct returns true", func(t *testing.T) {
		p := &Person{Name: "Bob"}
		ok, err := HasDepth(p)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !ok {
			t.Errorf("expected true for *struct, got false")
		}
	})

	t.Run("double pointer to struct returns true", func(t *testing.T) {
		p := &Person{Name: "Charlie"}
		pp := &p
		ok, err := HasDepth(pp)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !ok {
			t.Errorf("expected true for **struct, got false")
		}
	})

	t.Run("interface holding struct returns true", func(t *testing.T) {
		var i any = Person{Name: "Dana"}
		ok, err := HasDepth(i)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !ok {
			t.Errorf("expected true for interface{struct}, got false")
		}
	})

	t.Run("interface holding pointer to struct returns true", func(t *testing.T) {
		var i any = &Person{Name: "Eve"}
		ok, err := HasDepth(i)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !ok {
			t.Errorf("expected true for interface{&struct}, got false")
		}
	})

	t.Run("primitive type returns false", func(t *testing.T) {
		ok, err := HasDepth(42)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if ok {
			t.Errorf("expected false for int, got true")
		}
	})

	t.Run("pointer to primitive returns false", func(t *testing.T) {
		x := 100
		ok, err := HasDepth(&x)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if ok {
			t.Errorf("expected false for *int, got true")
		}
	})

	t.Run("slice returns false", func(t *testing.T) {
		ok, err := HasDepth([]string{"a", "b"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if ok {
			t.Errorf("expected false for slice, got true")
		}
	})

	t.Run("map returns false", func(t *testing.T) {
		ok, err := HasDepth(map[string]int{"k": 1})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if ok {
			t.Errorf("expected false for map, got true")
		}
	})

	t.Run("channel returns false", func(t *testing.T) {
		ch := make(chan int)
		ok, err := HasDepth(ch)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if ok {
			t.Errorf("expected false for channel, got true")
		}
	})

	t.Run("function returns false", func(t *testing.T) {
		fn := func() {}
		ok, err := HasDepth(fn)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if ok {
			t.Errorf("expected false for func, got true")
		}
	})

	t.Run("nil interface returns error", func(t *testing.T) {
		var v any = nil
		_, err := HasDepth(v)
		if err == nil {
			t.Errorf("expected error for nil input, got nil")
		}
	})

	t.Run("nil pointer returns false (not error)", func(t *testing.T) {
		var p *Person = nil
		// reflect.TypeOf(p) is *Person (valid type), not invalid!
		ok, err := HasDepth(p)
		if err != nil {
			t.Fatalf("unexpected error for nil pointer: %v", err)
		}
		// Underlying type after unwrapping: Person → struct → true
		if !ok {
			t.Errorf("expected true for nil *struct (type is still struct), got false")
		}
	})
}

type RecursiveNode struct {
	Value int
	Next  *RecursiveNode
}

type MutualA struct {
	B *MutualB
}
type MutualB struct {
	A *MutualA
}

type NestedRecursive struct {
	Items []NestedRecursive
}

type NonRecursive struct {
	Name string
	Age  int
}

type Wrapper struct {
	Data *NonRecursive
}

type RecursiveViaInterface struct {
	Next interface{}
}

type RecursiveSlice struct {
	Children []*RecursiveSlice
}

type RecursiveChan struct {
	In chan RecursiveChan
}

func TestHasRecursiveRef(t *testing.T) {
	t.Run("direct recursive struct (pointer)", func(t *testing.T) {
		v := RecursiveNode{Value: 1}
		ok, err := HasRecursiveRef(v)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !ok {
			t.Errorf("expected true for RecursiveNode, got false")
		}
	})

	t.Run("mutual recursion", func(t *testing.T) {
		a := MutualA{}
		ok, err := HasRecursiveRef(a)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !ok {
			t.Errorf("expected true for mutual recursion, got false")
		}
	})

	t.Run("recursive via slice", func(t *testing.T) {
		v := NestedRecursive{}
		ok, err := HasRecursiveRef(v)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !ok {
			t.Errorf("expected true for slice-based recursion, got false")
		}
	})

	t.Run("recursive via chan", func(t *testing.T) {
		v := RecursiveChan{}
		ok, err := HasRecursiveRef(v)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !ok {
			t.Errorf("expected true for chan-based recursion, got false")
		}
	})

	t.Run("recursive via slice of pointers", func(t *testing.T) {
		v := RecursiveSlice{}
		ok, err := HasRecursiveRef(v)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !ok {
			t.Errorf("expected true for []*RecursiveSlice, got false")
		}
	})

	t.Run("non-recursive struct", func(t *testing.T) {
		v := NonRecursive{Name: "test"}
		ok, err := HasRecursiveRef(v)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if ok {
			t.Errorf("expected false for non-recursive struct, got true")
		}
	})

	t.Run("primitive types are not recursive", func(t *testing.T) {
		for _, val := range []any{42, "hello", true, []int{1, 2, 3}, map[string]int{}} {
			ok, err := HasRecursiveRef(val)
			if err != nil {
				t.Fatalf("unexpected error for %T: %v", val, err)
			}
			if ok {
				t.Errorf("expected false for %T, got true", val)
			}
		}
	})

	t.Run("interface field does not count as recursion", func(t *testing.T) {
		v := RecursiveViaInterface{}
		ok, err := HasRecursiveRef(v)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if ok {
			t.Errorf("expected false (no static recursion via interface), got true")
		}
	})

	t.Run("nil interface returns error", func(t *testing.T) {
		var v any = nil
		_, err := HasRecursiveRef(v)
		if err == nil {
			t.Errorf("expected error for nil interface, got nil")
		}
	})

	t.Run("nil pointer to recursive struct is still recursive (type-level)", func(t *testing.T) {
		var v *RecursiveNode = nil
		ok, err := HasRecursiveRef(v)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !ok {
			t.Errorf("expected true, got false")
		}
	})
}
