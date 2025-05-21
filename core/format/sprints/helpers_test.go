// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sprints

import (
	"errors"
	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"github.com/andrerrcosta2/gtools/core/internal/differlite"
	"testing"
)

func TestAnonymous(t *testing.T) {

}

// TestAppend tests the Append function
//
// after appending any sprintable value it should:
//   - return the early representation added by the next sprint
//     without any extra spaces
func TestAppend(t *testing.T) {
	early := "early"
	after := "after"

	s := Append(early, after)

	if s != early+after {
		t.Errorf("expected: %s, got: %s", early+after, s)
	}
}

// TestAppendErrors tests the AppendErrors function
//
// after appending any error it should:
//   - return a string representation of the errors separated by newlines
func TestAppendErrors(t *testing.T) {
	// Test 1: No errors
	t.Run("No errors", func(t *testing.T) {
		s := AppendErrors()
		if s != "" {
			t.Errorf("expected: %s, got: %s", "", s)
		}
	})

	// Test 2: One error
	t.Run("One error", func(t *testing.T) {
		err := errors.New("error")
		s := AppendErrors(err)
		if s != err.Error()+"\n" {
			t.Errorf("expected: %s, got: %s", err, s)
		}
	})

	// Test 3: Multiple errors
	t.Run("Multiple errors", func(t *testing.T) {
		err1 := errors.New("error1")
		err2 := errors.New("error2")
		err3 := errors.New("error3")
		s := AppendErrors(err1, err2, err3)
		expected := err1.Error() + "\n" + err2.Error() + "\n" + err3.Error() + "\n"
		if s != expected {
			t.Errorf("expected: \n%s, \ngot: \n%s", expected, s)
			t.Log(differlite.Quick(s, expected))
		}
	})
}

// TestKeyValue tests the KeyValue function
//
// after creating a key-value pair it should:
//   - return a string representation of the key-value pair in the format "key value"
//   - print the correct number of tabs
func TestKeyValue(t *testing.T) {
	s := KeyValue(indent.Zero(), "key", "value")
	if s != "key value" {
		t.Errorf("expected: '%q', got: '%q'", "key: value", s)
	}

	s = KeyValue(indent.Tab(1), "key", "value")
	if s != "\tkey value" {
		t.Errorf("expected: '%q', got: '%q'", " key value", s)
	}
}

// TestLKeyValue tests the LKeyValue function
//
// after creating a key-value pair it should:
//   - return a string representation of the key-value pair in the format "key value"
//   - print the correct number of tabs
//   - print a new line before the key-value pair
func TestLKeyValue(t *testing.T) {
	s := LKeyValue(indent.Zero(), "key", "value")
	if s != "\nkey value" {
		t.Errorf("expected: '%q', got: '%q'", "\nkey: value", s)
	}

	s = LKeyValue(indent.Tab(1), "key", "value")
	if s != "\n\tkey value" {
		t.Errorf("expected: '%q', got: '%q'", "\n key value", s)
	}
}

// TestLtKeyValue tests the LtKeyValue function
//
// after creating a key-value pair it should:
//   - return a string representation of the key-value pair in the format "key value"
//   - print the correct number of tabs
//   - print a new line and a tab before the key-value pair
func TestLtKeyValue(t *testing.T) {
	s := LtKeyValue(indent.Zero(), "key", "value")
	if s != "\n\tkey value" {
		t.Errorf("expected: '%q', got: '%q'", "\n\tkey: value", s)
	}

	s = LtKeyValue(indent.Tab(1), "key", "value")
	if s != "\n\t\tkey value" {
		t.Errorf("expected: '%q', got: '%q'", "\n\t key value", s)
	}
}

// TestMaxCharsLeft tests the MaxCharsLeft function
//
// after printing a string it should:
//   - return the whole string if it is less or equal to the max chars
//   - return the truncated string to the left if it is greater than the max chars
//   - indent it correctly according to the number of tabs
func TestMaxCharsLeft(t *testing.T) {
	s := MaxCharsLeft(indent.Zero(), "hello", 5)
	if s != "hello" {
		t.Errorf("expected: '%q', got: '%q'", "hello", s)
	}

	s = MaxCharsLeft(indent.Tab(1), "hello", 5)
	if s != "\thello" {
		t.Errorf("expected: '%q', got: '%q'", "hello", s)
	}

	s = MaxCharsLeft(indent.Zero(), "hello world", 5)
	if s != "hello..." {
		t.Errorf("expected: '%q', got: '%q'", "hello...", s)
	}

	s = MaxCharsLeft(indent.Tab(1), "hello world", 5)
	if s != "\thello..." {
		t.Errorf("expected: '%q', got: '%q'", "hello...", s)
	}

	s = MaxCharsLeft(indent.Zero(), "hello", 1)
	if s != "h..." {
		t.Errorf("expected: '%q', got: '%q'", "h...", s)
	}

	s = MaxCharsLeft(indent.Tab(1), "hello", 1)
	if s != "\th..." {
		t.Errorf("expected: '%q', got: '%q'", "h...", s)
	}
}

// TestMaxCharsRight tests the MaxCharsRight function
//
// after printing a string it should:
//   - return the whole string if it is less or equal to the max chars
//   - return the truncated string to the right if it is greater than the max chars
//   - indent it correctly according to the number of tabs
func TestMaxCharsRight(t *testing.T) {
	s := MaxCharsRight(indent.Zero(), "hello", 5)
	if s != "hello" {
		t.Errorf("expected: '%q', got: '%q'", "hello", s)
	}

	s = MaxCharsRight(indent.Tab(1), "hello", 5)
	if s != "\thello" {
		t.Errorf("expected: '%q', got: '%q'", "hello", s)
	}

	s = MaxCharsRight(indent.Zero(), "hello world", 5)
	if s != "...world" {
		t.Errorf("expected: '%q', got: '%q'", "...world", s)
	}

	s = MaxCharsRight(indent.Tab(1), "hello world", 5)
	if s != "\t...world" {
		t.Errorf("expected: '%q', got: '%q'", "...world", s)
	}

	s = MaxCharsRight(indent.Zero(), "hello world", 1)
	if s != "...d" {
		t.Errorf("expected: '%q', got: '%q'", "...d", s)
	}

	s = MaxCharsRight(indent.Tab(1), "hello world", 1)
	if s != "\t...d" {
		t.Errorf("expected: '%q', got: '%q'", "...d", s)
	}
}
