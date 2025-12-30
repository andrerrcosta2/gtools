// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package assertlite

import (
	"fmt"
	"strings"
	"testing"

	"github.com/andrerrcosta2/gtools/core/testlite/testseed"
)

// TestTrue tests the True function. This test must assert:
//   - If the argument is false, it should return false
//   - If the argument is true, it should return true
//   - It should print a default error or override a message if present
func TestTrue(t *testing.T) {
	t.Run("should return true when condition is true", func(t *testing.T) {
		mock := &mockTesting{}
		result := True(mock, true)
		if !result {
			t.FailNow()
		}
		if mock.Failed() {
			t.Errorf("expected no error logs, got %v", mock.fatal)
		}
	})

	t.Run("should log and return false when condition is false", func(t *testing.T) {
		mock := &mockTesting{}
		result := True(mock, false)
		if result {
			t.FailNow()
		}
		if !mock.Failed() {
			t.Error("expected an error log, but none occurred")
		} else if !strings.Contains(mock.fatal[0], "expected true, but got false") {
			t.Errorf("unexpected error message: %s", mock.fatal[0])
		}
	})

	t.Run("should include custom message when provided", func(t *testing.T) {
		mock := &mockTesting{}
		msg := "custom message"
		True(mock, false, msg)
		if !mock.Failed() {
			t.Error("expected an error log, but none occurred")
		} else if !strings.Contains(mock.fatal[0], msg) {
			t.Errorf("custom message not included: %s", mock.fatal[0])
		}
	})

	t.Run("should format message with arguments", func(t *testing.T) {
		mock := &mockTesting{}
		True(mock, false, "value was %d", 42)
		if !mock.Failed() {
			t.Error("expected an error log, but none occurred")
		} else if !strings.Contains(mock.fatal[0], "value was 42") {
			t.Errorf("formatted message not found: %s", mock.fatal[0])
		}
	})
}

// TestAllTrue tests the AreTrue function. This test must assert
//   - AreTrue returns true when all elements in the slice are true
//   - AreTrue returns false when any element in the slice is false
//   - It fails the test and prints a failure message when any element in the slice is false
func TestAllTrue(t *testing.T) {
	t.Run("should return true when all elements satisfy the predicate", func(t *testing.T) {
		mock := &mockTesting{}
		data := []int{2, 4, 6}
		result := AreTrue(mock, data, func(n int) bool {
			return n%2 == 0
		})
		if !result {
			t.FailNow()
		}
		if mock.Failed() {
			t.Errorf("expected no errors, got %v", mock.fatal)
		}
	})

	t.Run("should return false and log the first failing element", func(t *testing.T) {
		mock := &mockTesting{}
		data := []int{2, 4, 5, 6}
		result := AreTrue(mock, data, func(n int) bool {
			return n%2 == 0
		})
		if result {
			t.FailNow()
		}
		if !mock.Failed() {
			t.Error("expected a failure")
		} else if !strings.Contains(mock.fatal[0], "5") {
			t.Errorf("expected '5' in error message, got: %s", mock.fatal[0])
		}
	})

	t.Run("should use custom message if provided", func(t *testing.T) {
		mock := &mockTesting{}
		data := []int{2, 4, 5, 6}
		msg := "custom message"
		AreTrue(mock, data, func(n int) bool {
			return n%2 == 0
		}, msg)

		if !mock.Failed() {
			t.Error("expected a failure")
		} else if !strings.Contains(mock.fatal[0], msg) {
			t.Errorf("expected '%s' in error message, got: %s", msg, mock.fatal[0])
		}
	})

	t.Run("should format custom message with arguments", func(t *testing.T) {
		mock := &mockTesting{}
		data := []int{2, 4, 5, 6}
		AreTrue(mock, data, func(n int) bool {
			return n%2 == 0
		}, "value was %d", 99)

		if !mock.Failed() {
			t.Error("expected a failure")
		} else if !strings.Contains(mock.fatal[0], "value was 99") {
			t.Errorf("formatted message not found: %s", mock.fatal[0])
		}
	})
}

// TestPanic tests the Panic function. This function tests the following scenarios:
//   - Function panics: should return true, no error logged.
//   - Function doesn't panic: should return false, log default message.
//   - Custom message provided: should be used instead of default.
//   - Formatted message with args: should format correctly.
func TestPanic(t *testing.T) {
	t.Run("should return true when function panics", func(t *testing.T) {
		mock := &mockTesting{}
		result := Panic(mock, func() {
			panic("expected")
		})
		if !result {
			t.FailNow()
		}
		if mock.Failed() {
			t.Errorf("expected no errors, got %v", mock.fatal)
		}
	})

	t.Run("should return false and log error when function does not panic", func(t *testing.T) {
		mock := &mockTesting{}
		result := Panic(mock, func() {
			// do nothing
		})
		if result {
			t.Errorf("Expected a fail, got none")
			t.FailNow()
		}
		if !mock.Failed() {
			t.Error("expected a failure")
		} else if !strings.Contains(mock.fatal[0], "expected a panic, but got none") {
			t.Errorf("unexpected error message: %s", mock.fatal[0])
		}
	})

	t.Run("should use custom message when provided", func(t *testing.T) {
		mock := &mockTesting{}
		msg := "custom panic message"
		Panic(mock, func() {}, msg)

		if !mock.Failed() {
			t.Error("expected a failure")
		} else if !strings.Contains(mock.fatal[0], msg) {
			t.Errorf("custom message not included: %s", mock.fatal[0])
		}
	})

	t.Run("should format message with arguments", func(t *testing.T) {
		mock := &mockTesting{}
		Panic(mock, func() {}, "value was %d", 42)

		if !mock.Failed() {
			t.Error("expected a failure")
		} else if !strings.Contains(mock.fatal[0], "value was 42") {
			t.Errorf("formatted message not found: %s", mock.fatal[0])
		}
	})
}

// TestEqual_Primitives tests Equals assertion. This test covers:
//   - Equals and Not compare primitives (int, string, bool)
//   - Custom message override
//   - Formatted custom message
//   - Nil handling
//   - Mixed types
func TestEqual_Primitives(t *testing.T) {
	t.Run("should return true when integers are compare", func(t *testing.T) {
		mock := &mockTesting{}
		result := Equals(mock, 42, 42)
		if !result {
			t.FailNow()
		}
		if mock.Failed() {
			t.Errorf("expected no errors, got %v", mock.fatal)
		}
	})

	t.Run("should return false when integers differ", func(t *testing.T) {
		mock := &mockTesting{}
		result := Equals(mock, 42, 99)
		if result {
			t.FailNow()
		}
		if !mock.Failed() {
			t.Error("expected an error")
		} else if !strings.Contains(mock.fatal[0], "a = <int>42, b = <int>99") {
			t.Errorf("unexpected error: %s", mock.fatal[0])
		}
	})

	t.Run("should return true when strings are compare", func(t *testing.T) {
		mock := &mockTesting{}
		result := Equals(mock, "hello", "hello")
		if !result {
			t.FailNow()
		}
		if mock.Failed() {
			t.Errorf("expected no errors, got %v", mock.fatal)
		}
	})

	t.Run("should return false when strings differ", func(t *testing.T) {
		mock := &mockTesting{}
		result := Equals(mock, "hello", "world")
		if result {
			t.FailNow()
		}
		if !mock.Failed() {
			t.Error("expected an error")
		} else if !strings.Contains(mock.fatal[0], "a = <string>hello, b = <string>world") {
			t.Errorf("unexpected error: %s", mock.fatal[0])
		}
	})

	t.Run("should return true when booleans are compare", func(t *testing.T) {
		mock := &mockTesting{}
		result := Equals(mock, true, true)
		if !result {
			t.FailNow()
		}
		if mock.Failed() {
			t.Errorf("expected no errors, got %v", mock.fatal)
		}
	})

	t.Run("should return false when booleans differ", func(t *testing.T) {
		mock := &mockTesting{}
		result := Equals(mock, true, false)
		if result {
			t.FailNow()
		}
		if !mock.Failed() {
			t.Error("expected an error")
		} else if !strings.Contains(mock.fatal[0], "a = <bool>true, b = <bool>false") {
			t.Errorf("unexpected error: %s", mock.fatal[0])
		}
	})

	t.Run("should include custom message if provided", func(t *testing.T) {
		mock := &mockTesting{}
		msg := "custom message"
		Equals(mock, 1, 2, msg)
		if !mock.Failed() {
			t.Error("expected failure")
		} else if !strings.Contains(mock.fatal[0], msg) {
			t.Errorf("custom message not found: %s", mock.fatal[0])
		}
	})

	t.Run("should format message with arguments", func(t *testing.T) {
		mock := &mockTesting{}
		Equals(mock, 1, 2, "values mismatched: expected %d, got %d", 1, 2)
		if !mock.Failed() {
			t.Error("expected failure")
		} else if !strings.Contains(mock.fatal[0], "values mismatched: expected 1, got 2") {
			t.Errorf("formatted message not found: %s", mock.fatal[0])
		}
	})

	t.Run("should return true when both values are nil", func(t *testing.T) {
		mock := &mockTesting{}
		var a, b any = nil, nil
		result := Equals(mock, a, b)
		if !result {
			t.FailNow()
		}
		if mock.Failed() {
			t.Errorf("expected no errors, got %v", mock.fatal)
		}
	})

	t.Run("should return false when one value is nil and the other isn't", func(t *testing.T) {
		mock := &mockTesting{}
		var a any = nil
		var b any = 42
		result := Equals(mock, a, b)
		if result {
			t.FailNow()
		}
		if !mock.Failed() {
			t.Error("expected failure")
		} else if !strings.Contains(mock.fatal[0], "assertion failed. values are not equals") {
			t.Errorf("unexpected error: %s", mock.fatal[0])
		}
	})

	t.Run("should fail when types are different", func(t *testing.T) {
		mock := &mockTesting{}
		Equals(mock, 42, "42")
		if !mock.Failed() {
			t.Error("expected failure")
		} else if !strings.Contains(mock.fatal[0], "assertion failed. values are not equals") {
			t.Errorf("unexpected error: %s", mock.fatal[0])
		}
	})
}

func TestEqual_StructWithEqualMethod(t *testing.T) {
	t.Run("should return true when SortableValue structs are compare", func(t *testing.T) {
		mock := &mockTesting{}
		a := testseed.NewSortableValue("Alice", 30)
		b := testseed.NewSortableValue("Alice", 30)
		result := Equals(mock, a, b)
		if !result {
			t.FailNow()
		}
		if mock.Failed() {
			t.Errorf("expected no errors, got %v", mock.fatal)
		}
	})

	t.Run("should return false when SortableValue structs differ", func(t *testing.T) {
		mock := &mockTesting{}
		a := testseed.NewSortableValue("Alice", 30)
		b := testseed.NewSortableValue("Bob", 30)
		result := Equals(mock, a, b)
		if result {
			t.FailNow()
		}
		if !mock.Failed() {
			t.Error("expected an error")
		} else if !strings.Contains(mock.fatal[0], "a = <testseed.SortableValue>{Alice 30}, "+
			"b = <testseed.SortableValue>{Bob 30}") {
			t.Errorf("unexpected error: %s", mock.fatal[0])
		}
	})
}

func TestEqual_PointerToStructWithEqualMethod(t *testing.T) {
	t.Run("should return true when *SortableRef ptrs are compare", func(t *testing.T) {
		mock := &mockTesting{}
		a := testseed.NewSortableRef("Alice", 30)
		b := testseed.NewSortableRef("Alice", 30)
		result := Equals(mock, a, b)
		if !result {
			t.FailNow()
		}
		if mock.Failed() {
			t.Errorf("expected no errors, got %v", mock.fatal)
		}
	})

	t.Run("should return false when *SortableRef ptrs differ", func(t *testing.T) {
		mock := &mockTesting{}
		a := testseed.NewSortableRef("Alice", 30)
		b := testseed.NewSortableRef("Bob", 30)
		result := Equals(mock, a, b)
		if result {
			t.FailNow()
		}
		if !mock.Failed() {
			t.Error("expected an error")
		} else if !strings.Contains(mock.fatal[0], "a = <*testseed.SortableRef>&{Alice 30}, "+
			"b = <*testseed.SortableRef>&{Bob 30}") {
			t.Errorf("unexpected error: %s", mock.fatal[0])
		}
	})
}

func TestEqual_PointerAndValue(t *testing.T) {
	t.Run("should fail when comparing pointer to value", func(t *testing.T) {
		mock := &mockTesting{}
		a := testseed.NewSortableRef("Alice", 30)
		b := testseed.NewSortableValue("Alice", 30)
		result := Equals(mock, a, b)
		if result {
			t.FailNow()
		}
		if !mock.Failed() {
			t.Error("expected an error")
		} else if !strings.Contains(mock.fatal[0], "a = <*testseed.SortableRef>&{Alice 30}, "+
			"b = <testseed.SortableValue>{Alice 30}") {
			t.Errorf("unexpected error: %s", mock.fatal[0])
		}
	})
}

func TestEqual_ComparableValue(t *testing.T) {
	t.Run("should return true when ComparableValue structs are compare", func(t *testing.T) {
		mock := &mockTesting{}
		a := testseed.NewComparableValue("Alice", 30)
		b := testseed.NewComparableValue("Alice", 30)
		result := Equals(mock, a, b)
		if !result {
			t.FailNow()
		}
		if mock.Failed() {
			t.Errorf("expected no errors, got %v", mock.fatal)
		}
	})

	t.Run("should return false when ComparableValue structs differ", func(t *testing.T) {
		mock := &mockTesting{}
		a := testseed.NewComparableValue("Alice", 30)
		b := testseed.NewComparableValue("Bob", 30)
		result := Equals(mock, a, b)
		if result {
			t.FailNow()
		}
		if !mock.Failed() {
			t.Error("expected an error")
		} else if !strings.Contains(mock.fatal[0], "a = <testseed.ComparableValue>{Alice 30}, "+
			"b = <testseed.ComparableValue>{Bob 30}") {
			t.Errorf("unexpected error: %s", mock.fatal[0])
		}
	})
}

func TestEqual_ComparableRef(t *testing.T) {
	t.Run("should return true when *ComparableRef ptrs are compare", func(t *testing.T) {
		mock := &mockTesting{}
		a := testseed.NewComparableRef("Alice", 30)
		b := testseed.NewComparableRef("Alice", 30)
		result := Equals(mock, a, b)
		if !result {
			t.FailNow()
		}
		if mock.Failed() {
			t.Errorf("expected no errors, got %v", mock.fatal)
		}
	})

	t.Run("should return false when *ComparableRef ptrs differ", func(t *testing.T) {
		mock := &mockTesting{}
		a := testseed.NewComparableRef("Alice", 30)
		b := testseed.NewComparableRef("Bob", 30)
		result := Equals(mock, a, b)
		if result {
			t.FailNow()
		}
		if !mock.Failed() {
			t.Error("expected an error")
		} else if !strings.Contains(mock.fatal[0], "a = <*testseed.ComparableRef>&{Alice 30}, "+
			"b = <*testseed.ComparableRef>&{Bob 30}") {
			t.Errorf("unexpected error: %s", mock.fatal[0])
		}
	})
}

func TestEqual_CustomMessages(t *testing.T) {
	t.Run("should use custom message when provided", func(t *testing.T) {
		mock := &mockTesting{}
		a := testseed.NewComparableValue("Alice", 30)
		b := testseed.NewComparableValue("Bob", 30)
		msg := "custom message"
		Equals(mock, a, b, msg)

		if !mock.Failed() {
			t.Error("expected failure")
		} else if !strings.Contains(mock.fatal[0], msg) {
			t.Errorf("custom message not found: %s", mock.fatal[0])
		}
	})

	t.Run("should format message with arguments", func(t *testing.T) {
		mock := &mockTesting{}
		a := testseed.NewComparableValue("Alice", 30)
		b := testseed.NewComparableValue("Bob", 30)
		Equals(mock, a, b, "values mismatched: expected %q, got %q", "Alice", "Bob")

		if !mock.Failed() {
			t.Error("expected failure")
		} else if !strings.Contains(mock.fatal[0], "values mismatched: expected \"Alice\", got \"Bob\"") {
			t.Errorf("formatted message not found: %s", mock.fatal[0])
		}
	})
}

func TestEqual_InterfaceSameImplementation(t *testing.T) {
	t.Run("should return true when two Stringers have same implementation and value", func(t *testing.T) {
		mock := &mockTesting{}
		var a testseed.InterfaceStringer = testseed.NewStructStringer("hello")
		var b testseed.InterfaceStringer = testseed.NewStructStringer("hello")

		result := Equals(mock, a, b)
		if !result {
			t.FailNow()
		}
		if mock.Failed() {
			t.Errorf("expected no errors, got %v", mock.fatal)
		}
	})
}

func TestEqual_InterfaceDifferentValues(t *testing.T) {
	t.Run("should return false when two Stringers have same impl but different values", func(t *testing.T) {
		mock := &mockTesting{}
		var a testseed.InterfaceStringer = testseed.NewStructStringer("hello")
		var b testseed.InterfaceStringer = testseed.NewStructStringer("world")

		result := Equals(mock, a, b)
		if result {
			t.FailNow()
		}
		if !mock.Failed() {
			t.Error("expected failure")
		} else if !strings.Contains(mock.fatal[0], "a = <*testseed.StructStringer>hello, "+
			"b = <*testseed.StructStringer>world") {
			t.Errorf("unexpected error: %s", mock.fatal[0])
		}
	})
}

func TestEqual_InterfaceAndConcrete(t *testing.T) {
	t.Run("should return true when InterfaceStringer matches concrete type", func(t *testing.T) {
		mock := &mockTesting{}
		var a testseed.InterfaceStringer = testseed.NewStructStringer("hello")
		b := testseed.NewStructStringer("hello")

		result := Equals(mock, a, b)
		if !result {
			t.FailNow()
		}
		if mock.Failed() {
			t.Errorf("expected no errors, got %v", mock.fatal)
		}
	})
}

func TestEqual_NilInterfaceNil(t *testing.T) {
	t.Run("should return true when nil interface compare nil", func(t *testing.T) {
		mock := &mockTesting{}
		var a testseed.InterfaceStringer = nil
		var b testseed.InterfaceStringer = nil

		result := Equals(mock, a, b)
		if !result {
			t.FailNow()
		}
		if mock.Failed() {
			t.Errorf("expected no errors, got %v", mock.fatal)
		}
	})
}

func TestEqual_InterfaceNilVsNonNil(t *testing.T) {
	t.Run("should return false when one interface is nil and the other is not", func(t *testing.T) {
		mock := &mockTesting{}
		var a testseed.InterfaceStringer = nil
		var b testseed.InterfaceStringer = testseed.NewStructStringer("hello")

		result := Equals(mock, a, b)
		if result {
			t.FailNow()
		}
		if !mock.Failed() {
			t.Error("expected failure")
		} else if !strings.Contains(mock.fatal[0], "assertion failed. values are not equals") {
			t.Errorf("unexpected error: %s", mock.fatal[0])
		}
	})
}

func TestEqual_InterfaceNilVsAnyNil(t *testing.T) {
	t.Run("should return true when interface-nil == any(nil)", func(t *testing.T) {
		mock := &mockTesting{}
		var a testseed.InterfaceStringer = nil
		var b any = nil

		result := Equals(mock, a, b)
		if !result {
			t.FailNow()
		}
		if mock.Failed() {
			t.Errorf("expected no errors, got %v", mock.fatal)
		}
	})
}

func TestEqual_InterfaceWithPointerReceiver(t *testing.T) {
	t.Run("should return true when interface implemented by pointer receiver", func(t *testing.T) {
		mock := &mockTesting{}
		var a testseed.InterfaceCloser = &testseed.StructCloserSuccess{}
		var b testseed.InterfaceCloser = &testseed.StructCloserSuccess{}

		result := Equals(mock, a, b)
		if !result {
			t.FailNow()
		}
		if mock.Failed() {
			t.Errorf("expected no errors, got %v", mock.fatal)
		}
	})
}

func TestEqual_InterfaceDifferentImplementations(t *testing.T) {
	t.Run("should return false when interfaces have different "+
		"implementations", func(t *testing.T) {
		mock := &mockTesting{}
		var a testseed.InterfaceCloser = &testseed.StructCloserSuccess{}
		var b testseed.InterfaceCloser = &testseed.StructCloserError{}

		result := Equals(mock, a, b)
		if result {
			t.FailNow()
		}
		if !mock.Failed() {
			t.Error("expected failure")
		} else if !strings.Contains(mock.fatal[0], "a = <*testseed.StructCloserSuccess>&{}, "+
			"b = <*testseed.StructCloserError>&{}") {
			t.Errorf("unexpected error: %s", mock.fatal[0])
		}
	})
}

func TestEqual_InterfaceCustomMessage(t *testing.T) {
	t.Run("should use custom message when provided", func(t *testing.T) {
		mock := &mockTesting{}
		var a testseed.InterfaceStringer = testseed.NewStructStringer("hello")
		var b testseed.InterfaceStringer = testseed.NewStructStringer("world")
		msg := "custom interface comparison failed"

		Equals(mock, a, b, msg)

		if !mock.Failed() {
			t.Error("expected failure")
		} else if !strings.Contains(mock.fatal[0], msg) {
			t.Errorf("custom message not found: %s", mock.fatal[0])
		}
	})
}

func TestNoError(t *testing.T) {
	t.Run("should return true when error is nil", func(t *testing.T) {
		mock := &mockTesting{}
		result := NoError(mock, nil)
		if !result {
			t.FailNow()
		}
		if mock.Failed() {
			t.Errorf("expected no errors, got %v", mock.fatal)
		}
	})

	t.Run("should return false and log error when error is not nil", func(t *testing.T) {
		mock := &mockTesting{}
		err := fmt.Errorf("something went wrong")
		result := NoError(mock, err)
		if result {
			t.FailNow()
		}
		if !mock.Failed() {
			t.Error("expected a failure")
		} else if !strings.Contains(mock.fatal[0], "no error was expected, but got 'something went wrong'") {
			t.Errorf("unexpected error message: %s", mock.fatal[0])
		}
	})

	t.Run("should use custom message when provided", func(t *testing.T) {
		mock := &mockTesting{}
		err := fmt.Errorf("something went wrong")
		msg := "custom error message"
		NoError(mock, err, msg)

		if !mock.Failed() {
			t.Error("expected a failure")
		} else if !strings.Contains(mock.fatal[0], msg) {
			t.Errorf("custom message not included: %s", mock.fatal[0])
		}
	})

	t.Run("should format message with arguments", func(t *testing.T) {
		mock := &mockTesting{}
		err := fmt.Errorf("something went wrong")
		NoError(mock, err, "value was %d", 42)

		if !mock.Failed() {
			t.Error("expected a failure")
		} else if !strings.Contains(mock.fatal[0], "value was 42") {
			t.Errorf("formatted message not found: %s", mock.fatal[0])
		}
	})
}

func TestIsTypeOf_Primitives(t *testing.T) {
	t.Run("should return true when value is int", func(t *testing.T) {
		mock := &mockTesting{}
		result := IsTypeOf[int](mock, 42)
		if !result {
			t.FailNow()
		}
		if mock.Failed() {
			t.Errorf("expected no errors, got %v", mock.fatal)
		}
	})

	t.Run("should return false when value is not int", func(t *testing.T) {
		mock := &mockTesting{}
		result := IsTypeOf[int](mock, "hello")
		if result {
			t.FailNow()
		}
		if !mock.Failed() {
			t.Error("expected failure")
		} else if !strings.Contains(mock.fatal[0], "expected data to be of type 'int', but got 'string'") {
			t.Errorf("unexpected error: %s", mock.fatal[0])
		}
	})

	t.Run("should return true when value is string", func(t *testing.T) {
		mock := &mockTesting{}
		result := IsTypeOf[string](mock, "hello")
		if !result {
			t.FailNow()
		}
		if mock.Failed() {
			t.Errorf("expected no errors, got %v", mock.fatal)
		}
	})

	t.Run("should return false when value is not string", func(t *testing.T) {
		mock := &mockTesting{}
		result := IsTypeOf[string](mock, 42)
		if result {
			t.FailNow()
		}
		if !mock.Failed() {
			t.Error("expected failure")
		} else if !strings.Contains(mock.fatal[0], "expected data to be of type 'string', but got 'int'") {
			t.Errorf("unexpected error: %s", mock.fatal[0])
		}
	})

	t.Run("should return true when value is bool", func(t *testing.T) {
		mock := &mockTesting{}
		result := IsTypeOf[bool](mock, true)
		if !result {
			t.FailNow()
		}
		if mock.Failed() {
			t.Errorf("expected no errors, got %v", mock.fatal)
		}
	})

	t.Run("should return false when value is not bool", func(t *testing.T) {
		mock := &mockTesting{}
		result := IsTypeOf[bool](mock, 42)
		if result {
			t.FailNow()
		}
		if !mock.Failed() {
			t.Error("expected failure")
		} else if !strings.Contains(mock.fatal[0], "expected data to be of type 'bool', but got 'int'") {
			t.Errorf("unexpected error: %s", mock.fatal[0])
		}
	})

	t.Run("should use custom message when provided", func(t *testing.T) {
		mock := &mockTesting{}
		IsTypeOf[int](mock, "hello", "custom error message")
		if !mock.Failed() {
			t.Error("expected failure")
		} else if !strings.Contains(mock.fatal[0], "custom error message") {
			t.Errorf("custom message not found: %s", mock.fatal[0])
		}
	})

	t.Run("should format message with arguments", func(t *testing.T) {
		mock := &mockTesting{}
		IsTypeOf[int](mock, "hello", "value was %d", 99)
		if !mock.Failed() {
			t.Error("expected failure")
		} else if !strings.Contains(mock.fatal[0], "value was 99") {
			t.Errorf("formatted message not found: %s", mock.fatal[0])
		}
	})

	t.Run("should fail when checking type against nil", func(t *testing.T) {
		mock := &mockTesting{}
		var v any = nil
		result := IsTypeOf[int](mock, v)
		if result {
			t.FailNow()
		}
		if !mock.Failed() {
			t.Error("expected failure")
		} else if !strings.Contains(mock.fatal[0], "expected data to be of type 'int', but got '<nil>'") {
			t.Errorf("unexpected error: %s", mock.fatal[0])
		}
	})

	t.Run("should fail when checking int against *int", func(t *testing.T) {
		mock := &mockTesting{}
		var v any = new(int)
		result := IsTypeOf[int](mock, v)
		if result {
			t.FailNow()
		}
		if !mock.Failed() {
			t.Error("expected failure")
		} else if !strings.Contains(mock.fatal[0], "expected data to be of type 'int', but got '*int'") {
			t.Errorf("unexpected error: %s", mock.fatal[0])
		}
	})
}

func TestIsTypeOf_Structs(t *testing.T) {
	t.Run("should return true when struct value matches", func(t *testing.T) {
		mock := &mockTesting{}
		var v any = testseed.StructTwoData{DataA: "Alice", DataB: 30}
		result := IsTypeOf[testseed.StructTwoData](mock, v)
		if !result {
			t.FailNow()
		}
		if mock.Failed() {
			t.Errorf("expected no errors, got %v", mock.fatal)
		}
	})

	t.Run("should return false when struct value does not match", func(t *testing.T) {
		mock := &mockTesting{}
		var v any = testseed.StructTwoData{DataA: "Alice", DataB: 30}
		result := IsTypeOf[testseed.StructOneData](mock, v)
		if result {
			t.FailNow()
		}
		if !mock.Failed() {
			t.Error("expected failure")
		} else if !strings.Contains(mock.fatal[0], "expected data to be of type 'testseed.StructOneData', "+
			"but got 'testseed.StructTwoData'") {
			t.Errorf("unexpected error: %s", mock.fatal[0])
		}
	})

	t.Run("should return true when struct pointer matches", func(t *testing.T) {
		mock := &mockTesting{}
		var v any = &testseed.StructTwoData{DataA: "Alice", DataB: 30}
		result := IsTypeOf[*testseed.StructTwoData](mock, v)
		if !result {
			t.FailNow()
		}
		if mock.Failed() {
			t.Errorf("expected no errors, got %v", mock.fatal)
		}
	})

	t.Run("should return false when pointer type checked against value", func(t *testing.T) {
		mock := &mockTesting{}
		s := testseed.StructOneData{Name: "Alice"}
		result := IsTypeOf[*testseed.StructOneData](mock, s)
		if result {
			t.FailNow()
		}
		if !mock.Failed() {
			t.Error("expected failure")
		} else if !strings.Contains(mock.fatal[0], "expected data to be of type '*testseed.StructOneData', but got 'testseed.StructOneData'") {
			t.Errorf("unexpected error: %s", mock.fatal[0])
		}
	})

	t.Run("should return true when interface implementation matches", func(t *testing.T) {
		mock := &mockTesting{}
		var s testseed.InterfaceStringer = testseed.NewStructStringer("hello")
		result := IsTypeOf[*testseed.StructStringer](mock, s)
		if !result {
			t.FailNow()
		}
		if mock.Failed() {
			t.Errorf("expected no errors, got %v", mock.fatal)
		}
	})

	t.Run("should return false when interface implementation does not match", func(t *testing.T) {
		mock := &mockTesting{}
		var s testseed.InterfaceStringer = testseed.NewStructStringer("hello")
		result := IsTypeOf[*testseed.StructTwoData](mock, s)
		if result {
			t.FailNow()
		}
		if !mock.Failed() {
			t.Error("expected failure")
		} else if !strings.Contains(mock.fatal[0], "*testseed.StructTwoData") {
			t.Errorf("unexpected error: %s", mock.fatal[0])
		}
	})

	t.Run("should fail when checking type on nil interface", func(t *testing.T) {
		mock := &mockTesting{}
		var v any = nil
		result := IsTypeOf[*testseed.StructOneData](mock, v)
		if result {
			t.FailNow()
		}
		if !mock.Failed() {
			t.Error("expected failure")
		} else if !strings.Contains(mock.fatal[0], "<nil>") {
			t.Errorf("unexpected error: %s", mock.fatal[0])
		}
	})

	t.Run("should use custom message when provided", func(t *testing.T) {
		mock := &mockTesting{}
		var s testseed.InterfaceStringer = testseed.NewStructStringer("hello")
		msg := "custom type mismatch"
		IsTypeOf[*testseed.StructTwoData](mock, s, msg)

		if !mock.Failed() {
			t.Error("expected failure")
		} else if !strings.Contains(mock.fatal[0], msg) {
			t.Errorf("custom message not found: %s", mock.fatal[0])
		}
	})

	t.Run("should format message with arguments", func(t *testing.T) {
		mock := &mockTesting{}
		var s testseed.InterfaceStringer = testseed.NewStructStringer("hello")
		IsTypeOf[*testseed.StructTwoData](mock, s, "value was %q", "hello")

		if !mock.Failed() {
			t.Error("expected failure")
		} else if !strings.Contains(mock.fatal[0], "value was \"hello\"") {
			t.Errorf("formatted message not found: %s", mock.fatal[0])
		}
	})
}

func TestIsTypeOf_Interfaces(t *testing.T) {
	t.Run("should return true when struct implements InterfaceStringer", func(t *testing.T) {
		mock := &mockTesting{}
		s := &testseed.StructOneData{Name: "Alice"}
		var _ testseed.InterfaceStringer = (*testseed.StructOneData)(nil)

		result := IsTypeOf[testseed.InterfaceStringer](mock, s)
		if !result {
			t.FailNow()
		}
		if mock.Failed() {
			t.Errorf("expected no errors, got %v", mock.fatal)
		}

		mock.Clear()

		result = IsTypeOf[*testseed.StructOneData](mock, s)
		if !result {
			t.FailNow()
		}
		if mock.Failed() {
			t.Errorf("expected no errors, got %v", mock.fatal)
		}
	})

	t.Run("should return false when struct doesn't implement InterfaceStringer", func(t *testing.T) {
		mock := &mockTesting{}
		s := testseed.StructEmpty{}
		result := IsTypeOf[testseed.InterfaceStringer](mock, s)
		if result {
			t.FailNow()
		}
		if !mock.Failed() {
			t.Error("expected failure")
		} else if !strings.Contains(mock.fatal[0], "InterfaceStringer") {
			t.Errorf("unexpected error: %s", mock.fatal[0])
		}
	})

	t.Run("should fail when checking type on interface-nil", func(t *testing.T) {
		mock := &mockTesting{}
		var v testseed.InterfaceStringer = nil
		result := IsTypeOf[testseed.InterfaceStringer](mock, v)
		if result {
			t.FailNow()
		}
		if !mock.Failed() {
			t.Error("expected failure")
		} else if !strings.Contains(mock.fatal[0], "expected data to be of type 'testseed.InterfaceStringer', but got '<nil>'") {
			t.Errorf("unexpected error: %s", mock.fatal[0])
		}
	})

	t.Run("should fail when checking type on any(nil)", func(t *testing.T) {
		mock := &mockTesting{}
		var v any = nil
		result := IsTypeOf[testseed.InterfaceStringer](mock, v)
		if result {
			t.FailNow()
		}
		if !mock.Failed() {
			t.Error("expected failure")
		} else if !strings.Contains(mock.fatal[0], "<nil>") {
			t.Errorf("unexpected error: %s", mock.fatal[0])
		}
	})

	t.Run("should return false when interface implementation mismatch", func(t *testing.T) {
		mock := &mockTesting{}
		var s testseed.InterfaceStringer = testseed.NewStructStringer("hello")

		// Try to check if it's *StructTwoData (which it's not)
		result := IsTypeOf[*testseed.StructTwoData](mock, s)
		if result {
			t.FailNow()
		}
		if !mock.Failed() {
			t.Error("expected failure")
		} else if !strings.Contains(mock.fatal[0], "*testseed.StructTwoData") {
			t.Errorf("unexpected error: %s", mock.fatal[0])
		}
	})

	t.Run("should use custom message when provided", func(t *testing.T) {
		mock := &mockTesting{}
		var s testseed.InterfaceStringer = testseed.NewStructStringer("hello")
		msg := "custom interface type mismatch"
		IsTypeOf[*testseed.StructTwoData](mock, s, msg)

		if !mock.Failed() {
			t.Error("expected failure")
		} else if !strings.Contains(mock.fatal[0], msg) {
			t.Errorf("custom message not found: %s", mock.fatal[0])
		}
	})

	t.Run("should format message with arguments", func(t *testing.T) {
		mock := &mockTesting{}
		var s testseed.InterfaceStringer = testseed.NewStructStringer("hello")
		IsTypeOf[*testseed.StructTwoData](mock, s, "value was %q", "hello")

		if !mock.Failed() {
			t.Error("expected failure")
		} else if !strings.Contains(mock.fatal[0], "value was \"hello\"") {
			t.Errorf("formatted message not found: %s", mock.fatal[0])
		}
	})
}

func TestArrayEquals_Primitives(t *testing.T) {
	t.Run("should return true when slices are compare", func(t *testing.T) {
		mock := &mockTesting{}
		a := []int{1, 2, 3}
		b := []int{1, 2, 3}
		result := EqualSlices(mock, a, b)
		if !result {
			t.FailNow()
		}
		if mock.Failed() {
			t.Errorf("expected no errors, got %v", mock.fatal)
		}
	})

	t.Run("should fail when slices have different lengths", func(t *testing.T) {
		mock := &mockTesting{}
		a := []int{1, 2}
		b := []int{1, 2, 3}
		result := EqualSlices(mock, a, b)
		if result {
			t.FailNow()
		}
		if !mock.Failed() {
			t.Error("expected failure")
		} else if !strings.Contains(mock.fatal[0], "expected slices to be equals, but got different lengths") {
			t.Errorf("unexpected error: %s", mock.fatal[0])
		}
	})

	t.Run("should fail when slices have same length but different elements", func(t *testing.T) {
		mock := &mockTesting{}
		a := []int{1, 2, 4}
		b := []int{1, 2, 3}
		result := EqualSlices(mock, a, b)
		if result {
			t.FailNow()
		}
		if !mock.Failed() {
			t.Error("expected failure")
		} else if !strings.Contains(mock.fatal[0], "expected slices to be equals, but got different values at") {
			t.Errorf("unexpected error: %s", mock.fatal[0])
		}
	})

	t.Run("should use custom message when provided", func(t *testing.T) {
		mock := &mockTesting{}
		a := []int{1, 2, 4}
		b := []int{1, 2, 3}
		msg := "custom slice comparison failed"
		EqualSlices(mock, a, b, msg)

		if !mock.Failed() {
			t.Error("expected failure")
		} else if !strings.Contains(mock.fatal[0], msg) {
			t.Errorf("custom message not found: %s", mock.fatal[0])
		}
	})

	t.Run("should format message with arguments", func(t *testing.T) {
		mock := &mockTesting{}
		a := []int{1, 2, 4}
		b := []int{1, 2, 3}
		EqualSlices(mock, a, b, "slice mismatch: expected %v, got %v", b, a)

		if !mock.Failed() {
			t.Error("expected failure")
		} else if !strings.Contains(mock.fatal[0], "slice mismatch: expected [1 2 3], got [1 2 4]") {
			t.Errorf("formatted message not found: %s", mock.fatal[0])
		}
	})

	t.Run("should fail when comparing nil slice and empty slice", func(t *testing.T) {
		mock := &mockTesting{}
		var a []int = nil
		var b = []int{}

		result := EqualSlices(mock, a, b)
		if result {
			t.FailNow()
		}
		if !mock.Failed() {
			t.Error("expected failure")
		} else if !strings.Contains(mock.fatal[0], "one of them is nil") {
			t.Errorf("unexpected error: %s", mock.fatal[0])
		}
	})
}

func TestArrayEquals_StructSlices_Comparable(t *testing.T) {
	t.Run("should return true when struct slices are compare", func(t *testing.T) {
		mock := &mockTesting{}
		a := []testseed.StructOneData{
			{Name: "Alice"},
			{Name: "Bob"},
		}
		b := []testseed.StructOneData{
			{Name: "Alice"},
			{Name: "Bob"},
		}
		result := EqualSlices(mock, a, b)
		if !result {
			t.FailNow()
		}
		if mock.Failed() {
			t.Errorf("expected no errors, got %v", mock.fatal)
		}
	})

	t.Run("should fail when struct slices differ", func(t *testing.T) {
		mock := &mockTesting{}
		a := []testseed.StructOneData{{Name: "Alice"}}
		b := []testseed.StructOneData{{Name: "Bob"}}
		result := EqualSlices(mock, a, b)
		if result {
			t.FailNow()
		}
		if !mock.Failed() {
			t.Error("expected failure")
		} else if !strings.Contains(mock.fatal[0], "assertion failed. expected slices to be equals, "+
			"but got different") {
			t.Errorf("unexpected error: %s", mock.fatal[0])
		}
	})

	t.Run("should use custom message when provided", func(t *testing.T) {
		mock := &mockTesting{}
		a := []testseed.StructOneData{{"Alice"}}
		b := []testseed.StructOneData{{"Bob"}}
		msg := "custom struct slice mismatch"
		EqualSlices(mock, a, b, msg)

		if !mock.Failed() {
			t.Error("expected failure")
		} else if !strings.Contains(mock.fatal[0], msg) {
			t.Errorf("custom message not found: %s", mock.fatal[0])
		}
	})

	t.Run("should format message with arguments", func(t *testing.T) {
		mock := &mockTesting{}
		a := []testseed.StructOneData{{"Alice"}}
		b := []testseed.StructOneData{{"Bob"}}
		EqualSlices(mock, a, b, "value mismatch: expected %q, got %q", b[0], a[0])

		if !mock.Failed() {
			t.Error("expected failure")
		} else if !strings.Contains(mock.fatal[0], "value mismatch: expected {\"Bob\"}, got {\"Alice\"}") {
			t.Errorf("formatted message not found: %s", mock.fatal[0])
		}
	})
}

func TestArrayEquals_PointerStructs(t *testing.T) {
	t.Run("should return true when slices of ptrs have compare values", func(t *testing.T) {
		mock := &mockTesting{}
		a := []*testseed.StructOneData{
			{Name: "Alice"},
			{Name: "Bob"},
		}
		b := []*testseed.StructOneData{
			{Name: "Alice"},
			{Name: "Bob"},
		}
		result := EqualSlices(mock, a, b)
		if !result {
			t.Log("Expected result to be true but got false")
			t.FailNow()
		}
		if mock.Failed() {
			t.Errorf("expected no errors, got %v", mock.fatal)
		}
	})

	t.Run("should return false if ptrs values differ", func(t *testing.T) {
		mock := &mockTesting{}

		val1 := testseed.StructOneData{Name: "Alice"}
		val2 := testseed.StructOneData{Name: "Alice2"}

		a := []*testseed.StructOneData{&val1}
		b := []*testseed.StructOneData{&val2}

		result := EqualSlices(mock, a, b)
		if result {
			t.FailNow()
		}
		if !mock.Failed() {
			t.Error("expected failure")
		} else if !strings.Contains(mock.fatal[0], "got different values at") {
			t.Errorf("unexpected error: %s", mock.fatal[0])
		}
	})
}

func TestNotNil_Primitives(t *testing.T) {
	mock := &mockTesting{}

	tests := []struct {
		name  string
		value any
	}{
		{"int", 42},
		{"string", "hello"},
		{"bool", true},
		{"struct", struct{}{}},
		{"array", [1]int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NotNil(mock, tt.value)
			if !result {
				t.FailNow()
			}
			if mock.Failed() {
				t.Errorf("expected no errors, got %v", mock.fatal)
			}
		})
	}
}

func TestNotNil_Structs(t *testing.T) {
	t.Run("should return true for non-nil struct", func(t *testing.T) {
		mock := &mockTesting{}
		s := testseed.StructOneData{Name: "Alice"}
		result := NotNil(mock, s)
		if !result {
			t.FailNow()
		}
		if mock.Failed() {
			t.Errorf("expected no errors, got %v", mock.fatal)
		}
	})

	t.Run("should return true for non-nil pointer", func(t *testing.T) {
		mock := &mockTesting{}
		s := &testseed.StructOneData{Name: "Alice"}
		result := NotNil(mock, s)
		if !result {
			t.FailNow()
		}
		if mock.Failed() {
			t.Errorf("expected no errors, got %v", mock.fatal)
		}
	})

	t.Run("should return false for nil pointer", func(t *testing.T) {
		mock := &mockTesting{}
		var s *testseed.StructOneData = nil
		result := NotNil(mock, s)
		if result {
			t.FailNow()
		}
		if !mock.Failed() {
			t.Error("expected failure")
		} else if !strings.Contains(mock.fatal[0], "assertion failed. target is nil") {
			t.Errorf("unexpected error: %s", mock.fatal[0])
		}
	})
}

func TestNotNil_Interfaces(t *testing.T) {
	t.Run("should return false for interface-nil holding nil pointer", func(t *testing.T) {
		mock := &mockTesting{}
		var val any = (*testseed.StructOneData)(nil)
		result := NotNil(mock, val)
		if result {
			t.FailNow()
		}
		if !mock.Failed() {
			t.Error("expected failure")
		} else if !strings.Contains(mock.fatal[0], "assertion failed. target is nil") {
			t.Errorf("unexpected error: %s", mock.fatal[0])
		}
	})

	t.Run("should return true for interface-nil holding non-nil struct", func(t *testing.T) {
		mock := &mockTesting{}
		var val any = testseed.StructOneData{Name: "Alice"}
		result := NotNil(mock, val)
		if !result {
			t.FailNow()
		}
		if mock.Failed() {
			t.Errorf("expected no errors, got %v", mock.fatal)
		}
	})

	t.Run("should return false for nil map inside interface", func(t *testing.T) {
		mock := &mockTesting{}
		var val any = map[string]int(nil)
		result := NotNil(mock, val)
		if result {
			t.FailNow()
		}
		if !mock.Failed() {
			t.Error("expected failure")
		} else if !strings.Contains(mock.fatal[0], "assertion failed. target is nil") {
			t.Errorf("unexpected error: %s", mock.fatal[0])
		}
	})

	t.Run("should return false for nil slice inside interface", func(t *testing.T) {
		mock := &mockTesting{}
		var val any = []int(nil)
		result := NotNil(mock, val)
		if result {
			t.FailNow()
		}
		if !mock.Failed() {
			t.Error("expected failure")
		} else if !strings.Contains(mock.fatal[0], "assertion failed. target is nil") {
			t.Errorf("unexpected error: %s", mock.fatal[0])
		}
	})

	t.Run("should return true for non-nil interface implementation", func(t *testing.T) {
		mock := &mockTesting{}
		var val testseed.InterfaceStringer = &testseed.StructOneData{Name: "Alice"}
		result := NotNil(mock, val)
		if !result {
			t.FailNow()
		}
		if mock.Failed() {
			t.Errorf("expected no errors, got %v", mock.fatal)
		}
	})

	t.Run("should return false for InterfaceStringer(nil)", func(t *testing.T) {
		mock := &mockTesting{}
		var val testseed.InterfaceStringer = nil
		result := NotNil(mock, val)
		if result {
			t.FailNow()
		}
		if !mock.Failed() {
			t.Error("expected failure")
		} else if !strings.Contains(mock.fatal[0], "assertion failed. target is nil") {
			t.Errorf("unexpected error: %s", mock.fatal[0])
		}
	})

	t.Run("should return false for interface-nil holding nil pointer", func(t *testing.T) {
		mock := &mockTesting{}
		var val testseed.InterfaceStringer = (*testseed.StructOneData)(nil)
		result := NotNil(mock, val)
		if result {
			t.FailNow()
		}
		if !mock.Failed() {
			t.Error("expected failure")
		} else if !strings.Contains(mock.fatal[0], "assertion failed. target is nil") {
			t.Errorf("unexpected error: %s", mock.fatal[0])
		}
	})

	t.Run("should return false for any(nil)", func(t *testing.T) {
		mock := &mockTesting{}
		var val any = nil
		result := NotNil(mock, val)
		if result {
			t.FailNow()
		}
		if !mock.Failed() {
			t.Error("expected failure")
		} else if !strings.Contains(mock.fatal[0], "assertion failed. target is nil") {
			t.Errorf("unexpected error: %s", mock.fatal[0])
		}
	})

	t.Run("should return true for interface implemented by pointer receiver", func(t *testing.T) {
		mock := &mockTesting{}
		var val testseed.InterfaceStringer = &testseed.StructOneData{Name: "Alice"}
		result := NotNil(mock, val)
		if !result {
			t.FailNow()
		}
		if mock.Failed() {
			t.Errorf("expected no errors, got %v", mock.fatal)
		}
	})

	t.Run("should return false for interface-nil with method set", func(t *testing.T) {
		mock := &mockTesting{}
		var val testseed.InterfaceStringer = (*testseed.StructOneData)(nil)
		result := NotNil(mock, val)
		if result {
			t.FailNow()
		}
		if !mock.Failed() {
			t.Error("expected failure")
		} else if !strings.Contains(mock.fatal[0], "assertion failed. target is nil") {
			t.Errorf("unexpected error: %s", mock.fatal[0])
		}
	})

	t.Run("should use custom message when provided", func(t *testing.T) {
		mock := &mockTesting{}
		var val testseed.InterfaceStringer = nil
		msg := "custom interface-nil check failed"
		NotNil(mock, val, msg)

		if !mock.Failed() {
			t.Error("expected failure")
		} else if !strings.Contains(mock.fatal[0], msg) {
			t.Errorf("custom message not found: %s", mock.fatal[0])
		}
	})

	t.Run("should format message with arguments", func(t *testing.T) {
		mock := &mockTesting{}
		var val testseed.InterfaceStringer = nil
		NotNil(mock, val, "value was unexpectedly %s", "nil")

		if !mock.Failed() {
			t.Error("expected failure")
		} else if !strings.Contains(mock.fatal[0], "value was unexpectedly nil") {
			t.Errorf("formatted message not found: %s", mock.fatal[0])
		}
	})
}

func TestNoNilFields_ExportedOnly(t *testing.T) {
	t.Run("should return true for struct with no nil fields (as value)", func(t *testing.T) {
		mock := &mockTesting{}
		target := testseed.StructWithPointersNoNilFieldsAsValue()
		result := NoNilFields(mock, false, target)
		if !result {
			t.FailNow()
		}
		if mock.Failed() {
			t.Errorf("expected no errors, got %v", mock.fatal)
		}
	})

	t.Run("should return true for struct with no nil fields (as pointer)", func(t *testing.T) {
		mock := &mockTesting{}
		target := testseed.StructWithPointersNoNilFieldsAsRef()
		result := NoNilFields(mock, false, target)
		if !result {
			t.FailNow()
		}
		if mock.Failed() {
			t.Errorf("expected no errors, got %v", mock.fatal)
		}
	})

	t.Run("should fail for struct with all nil fields (as value)", func(t *testing.T) {
		mock := &mockTesting{}
		target := testseed.StructWithPointersAllNilFieldsAsValue()
		result := NoNilFields(mock, false, target)
		if result {
			t.FailNow()
		}
		if !mock.Failed() {
			t.Error("expected failure")
		} else if !strings.Contains(mock.fatal[0], "Name") ||
			!strings.Contains(mock.fatal[0], "Data") ||
			!strings.Contains(mock.fatal[0], "Values") ||
			!strings.Contains(mock.fatal[0], "Config") {
			t.Errorf("unexpected error: %s", mock.fatal[0])
		}
	})

	t.Run("should fail for struct with all nil fields (as pointer)", func(t *testing.T) {
		mock := &mockTesting{}
		target := testseed.StructWithPointersAllNilFieldsAsRef()
		result := NoNilFields(mock, false, target)
		if result {
			t.FailNow()
		}
		if !mock.Failed() {
			t.Error("expected failure")
		} else if !strings.Contains(mock.fatal[0], "Name") {
			t.Errorf("missing field in error: Name")
		}
	})

	t.Run("should fail for struct with some nil fields (as value)", func(t *testing.T) {
		mock := &mockTesting{}
		target := testseed.StructWithPointersSomeNilFieldsAsValue()
		result := NoNilFields(mock, false, target)
		if result {
			t.FailNow()
		}
		if !mock.Failed() {
			t.Error("expected failure")
		} else if !strings.Contains(mock.fatal[0], "Data") {
			t.Errorf("missing field in error: Data")
		}
	})

	t.Run("should fail for struct with some nil fields (as pointer)", func(t *testing.T) {
		mock := &mockTesting{}
		target := testseed.StructWithPointersSomeNilFieldsAsRef()
		result := NoNilFields(mock, false, target)
		if result {
			t.FailNow()
		}
		if !mock.Failed() {
			t.Error("expected failure")
		} else if !strings.Contains(mock.fatal[0], "Values") {
			t.Errorf("missing field in error: Values")
		}
	})

	t.Run("should use custom message when provided", func(t *testing.T) {
		mock := &mockTesting{}
		target := testseed.StructWithPointersAllNilFieldsAsValue()
		msg := "custom nil field check failed"
		NoNilFields(mock, false, target, msg)

		if !mock.Failed() {
			t.Error("expected failure")
		} else if !strings.Contains(mock.fatal[0], msg) {
			t.Errorf("custom message not found: %s", mock.fatal[0])
		}
	})

	t.Run("should format message with arguments", func(t *testing.T) {
		mock := &mockTesting{}
		target := testseed.StructWithPointersAllNilFieldsAsValue()
		NoNilFields(mock, false, target, "struct has nil fields: %v", target)

		if !mock.Failed() {
			t.Error("expected failure")
		} else if !strings.Contains(mock.fatal[0], "nil fields") {
			t.Errorf("formatted message not found: %s", mock.fatal[0])
		}
	})
}

func TestNoNilFields_Unexported(t *testing.T) {
	t.Run("should pass if no unexported nils and checkUnexported is true", func(t *testing.T) {
		mock := &mockTesting{}
		target := testseed.StructWithPointersAndUnexportedsNoNilFieldsAsValue()
		result := NoNilFields(mock, true, target)
		if !result {
			t.FailNow()
		}
		if mock.Failed() {
			t.Errorf("expected no errors, got %v", mock.fatal)
		}
	})

	t.Run("should fail if unexported field is nil and checkUnexported is true", func(t *testing.T) {
		mock := &mockTesting{}
		target := testseed.StructWithPointersAndUnexportedsSomeNilFieldsAsValue()
		result := NoNilFields(mock, true, target)
		if result {
			t.FailNow()
		}
		if !mock.Failed() {
			t.Error("expected failure")
		} else if !strings.Contains(mock.fatal[0], "data") {
			t.Errorf("missing field in error: data")
		}
	})

	t.Run("should fail for all nil unexported/exported fields", func(t *testing.T) {
		mock := &mockTesting{}
		target := testseed.StructWithPointersAndUnexportedsAllNilFieldsAsRef()
		result := NoNilFields(mock, true, target)
		if result {
			t.FailNow()
		}
		if !mock.Failed() {
			t.Error("expected failure")
		} else {
			err := mock.fatal[0]
			if !strings.Contains(err, "name") ||
				!strings.Contains(err, "data") ||
				!strings.Contains(err, "Values") ||
				!strings.Contains(err, "Config") {
				t.Errorf("missing expected nil fields in error: %s", err)
			}
		}
	})

	t.Run("should ignore unexported fields when checkUnexported is false", func(t *testing.T) {
		mock := &mockTesting{}
		target := testseed.StructWithPointersAndUnexportedsSomeNilFieldsAsRef()
		result := NoNilFields(mock, false, target)
		if result {
			t.FailNow()
		}
		if !mock.Failed() {
			t.Error("expected failure")
		} else if strings.Contains(mock.fatal[0], "name") ||
			strings.Contains(mock.fatal[0], "data") {
			t.Error("unexpected unexported field found in error")
		} else if !strings.Contains(mock.fatal[0], "Values") {
			t.Error("expected 'Values' field to be reported as nil")
		}
	})
}
