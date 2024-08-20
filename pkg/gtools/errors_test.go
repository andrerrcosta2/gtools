// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package gtools

import (
	"errors"
	"fmt"
	"testing"
)

func TestConcurrentStackableError_Empty(t *testing.T) {
	cse := NewConcurrentStackableError(nil)

	if !cse.Empty() {
		t.Fatalf("expected empty stack but found errors")
	}
}

func TestConcurrentStackableError_Error(t *testing.T) {
	baseErr := fmt.Errorf("base error")
	cse := NewConcurrentStackableError(baseErr)

	expected := "base error"
	if got := cse.Error(); got != expected {
		t.Fatalf("expected %q but got %q", expected, got)
	}
}

func TestConcurrentStackableError_Stack(t *testing.T) {
	baseErr := fmt.Errorf("base error")
	secondErr := fmt.Errorf("second error")
	cse := NewConcurrentStackableError(baseErr)

	cse.Stack(secondErr)

	if got := cse.Len(); got != 2 {
		t.Fatalf("expected stack length of 2 but got %d", got)
	}

	if cse.Unwrap() != secondErr {
		t.Fatalf("expected base error but got %v", cse.Unwrap())
	}

	if got := cse.Trace(); got == "" {
		t.Fatal("expected non-empty trace")
	}
}

func TestConcurrentStackableError_From(t *testing.T) {
	baseErr := fmt.Errorf("base error")
	secondErr := fmt.Errorf("second error")
	cse := NewConcurrentStackableError(baseErr)

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
	cse := NewConcurrentStackableError(fmt.Errorf("base error"))

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

	if cse.Empty() {
		t.Fatal("expected non-empty stack but found empty")
	}
}

func TestConcurrentStackableError_Trace(t *testing.T) {
	baseErr := fmt.Errorf("base error")
	secondErr := fmt.Errorf("second error")
	cse := NewConcurrentStackableError(baseErr)

	cse.Stack(secondErr)

	trace := cse.Trace()
	if trace == "" {
		t.Fatal("expected non-empty trace")
	}

	if len(cse.Trace()) == 0 {
		t.Fatal("expected non-empty trace after stacking errors")
	}
}
