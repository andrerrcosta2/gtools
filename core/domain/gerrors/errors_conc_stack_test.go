// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package gerrors

import (
	"errors"
	"fmt"
	"sync"
	"testing"
)

// TestConcurrentStackable_ConcStackable tests a concurrent stacking
func TestConcurrentStackableError_ConcurrentAccess(t *testing.T) {
	cse := ConcStackableOf(errors.New("base error"))

	done := make(chan bool)
	for i := 0; i < 99; i++ {
		go func(i int) {
			defer func() { done <- true }()
			err := fmt.Errorf("error %d", i)
			cse.Stack(err)
		}(i)
	}

	for i := 0; i < 99; i++ {
		<-done
	}

	if got := cse.Len(); got != 100 {
		t.Fatalf("expected stack length of 100 but got %d", got)
	}

	if cse.IsEmpty() {
		t.Fatal("expected non-empty stack but found empty")
	}
}

func TestConcurrentStackableError_ConcurrentFrom(t *testing.T) {
	t.Run("race", func(t *testing.T) {
		base := errors.New("base")
		cse := ConcStackableOf(base)

		var wg sync.WaitGroup
		wg.Add(100)
		for i := 0; i < 100; i++ {
			go func(i int) {
				defer wg.Done()
				cse = cse.From(fmt.Errorf("e%d", i))
			}(i)
		}
		wg.Wait()

		if cse.Len() == 100 {
			t.Errorf("expected growth in stack after many From calls, but got %d", cse.Len())
		}
	})

	t.Run("invariance", func(t *testing.T) {
		base := errors.New("root")
		cse := ConcStackableOf(base)

		var wg sync.WaitGroup
		for i := 0; i < 200; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				out := cse.From(fmt.Errorf("e%d", i))

				// Assert immediate invariants on the returned object
				if out.Len() == 0 {
					t.Errorf("returned stack should never be empty after .From")
				}
				if out.Unwrap() == nil {
					t.Errorf("returned leaf must not be nil")
				}
			}(i)
		}
		wg.Wait()
	})
}

// TestConcurrentStackable_ConcStackable tests the methods:
//   - ConcStackableOf() expecting it to stack, store, flatten and return the given error
func TestConcurrentStackable_ConcStackable(t *testing.T) {
	// Create with a nil error
	cse := ConcStackableOf(nil)
	if cse == nil {
		t.Fatal("expected non-nil stack but found nil")
	}
	var ass *concStackErr
	if !errors.As(cse, &ass) {
		t.Fatal("expected concStackErr")
	}
	// expect a nil error
	if len(ass.stk.stk) != 0 {
		t.Fatalf("expected empty error stack but got '%v'", ass.stk.stk)
	}

	// Create with a non-nil error
	baseErr := fmt.Errorf("base error")
	cse = ConcStackableOf(baseErr)
	if cse == nil {
		t.Fatal("expected non-nil stack but found nil")
	}
	fmt.Printf("cse: %v\n", cse)
	if !errors.As(cse, &ass) {
		t.Fatal("expected concStackErr")
	}
	// expect the base error
	if !errors.Is(baseErr, ass.stk.stk[0]) {
		t.Fatalf("expected base error but found %v", ass.stk.stk[0])
	}
	// expect a stack length of 0
	if ass.stk.Len() != 1 {
		t.Fatalf("expected stack length of 1 but found %d: %v", ass.stk.Len(), ass.stk)
	}
}

func TestConcurrentStackagleError_Cause(t *testing.T) {
	baseErr := fmt.Errorf("base error")
	cse := ConcStackableOf(baseErr)

	if got := cse.Cause(); !errors.Is(got, baseErr) {
		t.Fatalf("expected base error but got %v", got)
	}
}

// TestConcurrentStackable_ConcStack tests the methods:
//   - ConcStack() expecting it to stack, store, flatten and return the given error
func TestConcurrentStackable_ConcStack(t *testing.T) {
	// Create with a nil error
	cse := ConcStack()
	if cse == nil {
		t.Fatal("expected non-nil stack but found nil")
	}
	var ass *concStackErr
	if !errors.As(cse, &ass) {
		t.Fatal("expected concStackErr")
	}
	// expect a stack length of 0
	if ass.stk.Len() != 0 {
		t.Fatalf("expected stack length of 0 but found %d", len(ass.stk.stk))
	}
}

// TestConcurrentStackableError_Empty tests the methods:
// - IsEmpty()
// - ConcStack()
// - ConcStackableOf(nil)
// And checks whether the stack is empty.
func TestConcurrentStackableError_Empty(t *testing.T) {
	cse := ConcStackableOf(nil)
	var ass *concStackErr
	errors.As(cse, &ass)
	if !cse.IsEmpty() || len(ass.stk.stk) != 0 {
		t.Fatalf("expected empty stack but found errors")
	}
	cse = ConcStack()
	if !cse.IsEmpty() {
		t.Fatalf("expected empty stack but found errors")
	}
}

// TestConcurrentStackableError_Error tests the defaultErr() method
func TestConcurrentStackableError_Error(t *testing.T) {
	baseErr := fmt.Errorf("base error")
	cse := ConcStackableOf(baseErr)

	expected := "base error"
	if got := cse.Error(); got != expected {
		t.Fatalf("expected %q but got %q", expected, got)
	}
}

func TestConcurrentStackableError_From(t *testing.T) {
	t.Run("initial state", func(t *testing.T) {
		cse := ConcStack()
		baseErr := fmt.Errorf("base error")
		result := cse.From(baseErr)

		if result.Len() != 1 {
			t.Fatalf("expected length 1, got %d", result.Len())
		}

		if !errors.Is(result.Unwrap()[0], baseErr) {
			t.Fatalf("expected base error, got: %v", result.Unwrap())
		}
	})

	t.Run("valid errors", func(t *testing.T) {
		baseErr := fmt.Errorf("base error")
		secondErr := fmt.Errorf("second error")
		cse := ConcStackableOf(baseErr)
		err := cse.From(secondErr)
		if err == nil {
			t.Fatalf("expected create error but got nil")
		}
		if err.Len() != 2 {
			t.Fatalf("expected stack length of 2 but got %d", err.Len())
		}
		if got := cse.Unwrap()[0]; !errors.Is(got, baseErr) {
			t.Fatalf("expected base error but got %v", got)
		}
	})

	t.Run("from nil error", func(t *testing.T) {
		cse := ConcStackableOf(fmt.Errorf("base"))
		result := cse.From(nil)

		// Should return the same error when nil is passed
		if cse.Error() != result.Error() {
			t.Fatal("expected same instance when nil error passed to From")
		}
	})
}

// TestConcurrentStackableError_Is tests the method Is
// it should be true only if the stack of errors contains the same error instances
func TestConcurrentStackableError_Is(t *testing.T) {
	t.Run("equal pointer4, same stack len", func(t *testing.T) {
		baseErr := errors.New("base error")
		cse := ConcStackableOf(baseErr)

		if !cse.Is(baseErr) {
			t.Errorf("expected true when both errors are equal")
		}
	})

	t.Run("equal pointer4, different stack len", func(t *testing.T) {
		baseErr := errors.New("base error")
		secondErr := errors.New("second error")
		cse := ConcStackableOf(baseErr)
		cse.Stack(secondErr)

		if !cse.Is(secondErr) {
			t.Errorf("expected true when target is on the stack")
		}
	})

	t.Run("equal strings, different pointer4", func(t *testing.T) {
		s1 := ConcStackableOf(errors.New("a"))
		s1.Stack(errors.New("b"))

		s2 := ConcStackableOf(errors.New("a"))
		s2.Stack(errors.New("b"))

		if s1.Is(s2) {
			t.Fatal("two equal stacks with different backing errors shouldn't be equal")
		}
	})

	t.Run("empty vs nil", func(t *testing.T) {
		var a = ConcStack()
		if !a.Is(nil) {
			t.Error("empty should be equal nil")
		}
	})

	t.Run("nil vs not nil", func(t *testing.T) {
		var a = ConcStack()
		b := ConcStackableOf(errors.New("x"))
		if a.Is(b) {
			t.Fatal("nil should not equal non-nil")
		}
	})

	t.Run("stack vs leaf", func(t *testing.T) {
		a := errors.New("x")
		s := ConcStackableOf(a)
		if !s.Is(a) {
			t.Fatal("stack top should match leaf error")
		}
	})

	t.Run("multiple errors", func(t *testing.T) {
		a1 := errors.New("a1")
		a2 := errors.New("a2")
		b1 := errors.New("a1")
		b2 := errors.New("a2")

		s1 := ConcStackableOf(a1)
		s1.Stack(a2)

		s2 := ConcStackableOf(b1)
		s2.Stack(b2)

		if s1.Is(s2) {
			t.Fatal("expected different errors because it should compare each error in the stack with " +
				"the outer error")
		}

		s1 = ConcStackableOf(a1)
		s1.Stack(a2)
		s2 = ConcStackableOf(a1)
		s2.Stack(a2)

		if s1.Is(s2) {
			t.Fatal("expected different errors because it should compare each error in the stack with " +
				"the outer error")
		}
	})
}

func TestConcurrentStackableError_Output(t *testing.T) {
	// Test with empty stack
	cse := ConcStackableOf(nil)
	if cse.Output() != nil {
		t.Fatal("expected nil output for empty stack")
	}

	// Test with non-empty stack
	baseErr := fmt.Errorf("base error")
	cse = ConcStackableOf(baseErr)
	output := cse.Output()
	if output == nil {
		t.Fatal("expected non-nil output for non-empty stack")
	}

	if !errors.Is(output, baseErr) {
		t.Fatalf("expected base error in output, got: %v", output)
	}
}

func TestConcurrentStackableError_Race(t *testing.T) {
	cse := ConcStackableOf(fmt.Errorf("base error"))
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_ = cse.Error()
			_ = cse.Len()
			_ = cse.IsEmpty()
			_ = cse.Unwrap()
			_ = cse.From(fmt.Errorf("e%d", i))
			cse.Stack(fmt.Errorf("concurrent error %d", i))
		}()
	}

	wg.Wait()
}

func TestConcurrentStackableError_ConcurrentIsMethod(t *testing.T) {
	cse := ConcStackableOf(fmt.Errorf("base error"))
	targetErr := fmt.Errorf("target")

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_ = cse.Is(targetErr)
		}()
	}

	wg.Wait()
}

// TestConcurrentStackableError_Stack tests the methods:
//   - Stack() expecting it stack the given error
//   - Len() expecting it return the number of errors in the stack
//   - Unwrap() expecting it return the first wrapped error
//   - Trace() expecting it return a string with the stack trace
func TestConcurrentStackableError_Stack(t *testing.T) {
	baseErr := fmt.Errorf("base error")
	secondErr := fmt.Errorf("second error")
	cse := ConcStackableOf(baseErr)

	cse.Stack(secondErr)

	if got := cse.Len(); got != 2 {
		t.Fatalf("expected stack length of 2 but got %d", got)
	}

	if !errors.Is(secondErr, cse.Unwrap()[1]) {
		t.Fatalf("expected base error but got %v", cse.Unwrap())
	}

	if got := cse.Error(); got == "" {
		t.Fatal("expected non-empty trace")
	}
}

func TestConcurrentStackableError_Trace(t *testing.T) {
	baseErr := fmt.Errorf("base error")
	secondErr := fmt.Errorf("second error")
	cse := ConcStackableOf(baseErr)

	cse.Stack(secondErr)

	trace := cse.Error()
	if trace == "" {
		t.Fatal("expected non-empty trace")
	}

	if len(cse.Error()) == 0 {
		t.Fatal("expected non-empty trace after stacking errors")
	}
}

func TestConcurrentStackableError_Unstack(t *testing.T) {
	baseErr := fmt.Errorf("base error")
	secondErr := fmt.Errorf("second error")
	cse := ConcStackableOf(baseErr)
	cse.Stack(secondErr)

	unstacked := cse.Unwrap()
	if len(unstacked) != 2 {
		t.Fatalf("expected 2 errors in stack, got %d", len(unstacked))
	}

	if !errors.Is(unstacked[0], baseErr) || !errors.Is(unstacked[1], secondErr) {
		t.Fatalf("expected errors in correct order")
	}
}
