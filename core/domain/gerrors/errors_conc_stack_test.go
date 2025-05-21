// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package gerrors

import (
	"errors"
	"fmt"
	"testing"
)

// TestConcurrentStackable_ConcStackable tests the methods:
//   - ConcStackable() expecting it to stack, store, flatten and return the given error
func TestConcurrentStackable_ConcStackable(t *testing.T) {
	// Create with a nil error
	cse := ConcStackable(nil)
	if cse == nil {
		t.Fatal("expected non-nil stack but found nil")
	}
	var ass *concStackErr
	errors.As(cse, &ass)
	// expect a nil error
	if ass.err != nil {
		t.Fatalf("expected nil error but found %v", cse.(*concStackErr).err)
	}
	// expect a stack length of 0
	if ass.stk.Len() != 0 {
		t.Fatalf("expected stack length of 0 but found %d", ass.stk.Len())
	}

	// Create with a non-nil error
	baseErr := fmt.Errorf("base error")
	cse = ConcStackable(baseErr)
	if cse == nil {
		t.Fatal("expected non-nil stack but found nil")
	}
	fmt.Printf("cse: %v\n", cse)
	errors.As(cse, &ass)
	// expect the base error
	if !errors.Is(baseErr, ass.err) {
		t.Fatalf("expected base error but found %v", cse.(*concStackErr).err)
	}
	// expect a stack length of 0
	if ass.stk.Len() != 1 {
		t.Fatalf("expected stack length of 1 but found %d: %v", ass.stk.Len(), ass.stk)
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
	errors.As(cse, &ass)
	// expect a nil error
	if ass.err != nil {
		t.Fatalf("expected nil error but found %v", cse.(*concStackErr).err)
	}
	// expect a stack length of 0
	if ass.stk.Len() != 0 {
		t.Fatalf("expected stack length of 0 but found %d", len(ass.stk.stk))
	}
}

// TestConcurrentStackableError_Empty tests the methods:
// - IsEmpty()
// - ConcStack()
// - ConcStackable(nil)
// And checks whether the stack is empty.
func TestConcurrentStackableError_Empty(t *testing.T) {
	cse := ConcStackable(nil)
	var ass *concStackErr
	errors.As(cse, &ass)
	if !cse.IsEmpty() || len(ass.stk.stk) != 0 || ass.err != nil {
		t.Fatalf("expected empty stack but found errors")
	}
	cse = ConcStack()
	if !cse.IsEmpty() {
		t.Fatalf("expected empty stack but found errors")
	}
}

// TestConcurrentStackableError_Error tests the Error() method
func TestConcurrentStackableError_Error(t *testing.T) {
	baseErr := fmt.Errorf("base error")
	cse := ConcStackable(baseErr)

	expected := "base error"
	if got := cse.Error(); got != expected {
		t.Fatalf("expected %q but got %q", expected, got)
	}
}

// TestConcurrentStackableError_Stack tests the methods:
//   - Stack() expecting it stack the given error
//   - Len() expecting it return the number of errors in the stack
//   - Unwrap() expecting it return the first wrapped error
//   - Trace() expecting it return a string with the stack trace
func TestConcurrentStackableError_Stack(t *testing.T) {
	baseErr := fmt.Errorf("base error")
	secondErr := fmt.Errorf("second error")
	cse := ConcStackable(baseErr)

	cse.Stack(secondErr)

	if got := cse.Len(); got != 2 {
		t.Fatalf("expected stack length of 2 but got %d", got)
	}

	if !errors.Is(secondErr, cse.Unwrap()) {
		t.Fatalf("expected base error but got %v", cse.Unwrap())
	}

	if got := cse.Trace(); got == "" {
		t.Fatal("expected non-empty trace")
	}
}

func TestConcurrentStackagleError_Cause(t *testing.T) {
	baseErr := fmt.Errorf("base error")
	cse := ConcStackable(baseErr)

	if got := cse.Cause(); !errors.Is(got, baseErr) {
		t.Fatalf("expected base error but got %v", got)
	}
}

func TestConcurrentStackableError_From(t *testing.T) {
	baseErr := fmt.Errorf("base error")
	secondErr := fmt.Errorf("second error")
	cse := ConcStackable(baseErr)

	err := cse.From(secondErr)
	if err == nil {
		t.Fatalf("expected create error but got nil")
	}

	if got := cse.Len(); got != 2 {
		t.Fatalf("expected stack length of 2 but got %d", got)
	}

	if got := cse.Unwrap(); !errors.Is(got, secondErr) {
		t.Fatalf("expected second error but got %v", got)
	}
}

func TestConcurrentStackableError_ConcurrentAccess(t *testing.T) {
	cse := ConcStackable(fmt.Errorf("base error"))

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

func TestConcurrentStackableError_Trace(t *testing.T) {
	baseErr := fmt.Errorf("base error")
	secondErr := fmt.Errorf("second error")
	cse := ConcStackable(baseErr)

	cse.Stack(secondErr)

	trace := cse.Trace()
	if trace == "" {
		t.Fatal("expected non-empty trace")
	}

	if len(cse.Trace()) == 0 {
		t.Fatal("expected non-empty trace after stacking errors")
	}
}
