// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package gerrors

import (
	"errors"
	"testing"
)

func TestFlattenError(t *testing.T) {
	t.Run("Single error", func(t *testing.T) {
		// Create a single error
		err := errors.New("test error")
		stk, peek, isw := FlattenError(err)
		if isw {
			t.Errorf("Expected isWrapped to be false, got true")
		}
		if len(stk) != 1 {
			t.Errorf("Expected stack length to be 1, got %d", len(stk))
		}
		if !errors.Is(err, peek) {
			t.Errorf("Expected last error to be %v, got %v", err, peek)
		}
	})

	t.Run("Multiple errors", func(t *testing.T) {
		// Create a wrapped error
		werr := errors.New("wrapped error")
		stk := Stackable(werr)
		err := errors.New("nested error")
		stk.Stack(err)
		flat, peek, isw := FlattenError(stk)
		if !isw {
			t.Errorf("Expected isWrapped to be true, got false")
		}
		if len(flat) != 2 {
			t.Errorf("Expected stack length to be 2, got %d", len(flat))
		}
		if werr.Error() != flat[0].Error() {
			t.Errorf("Expected first error to be %v, got %v", werr, flat[0])
		}
		if err.Error() != flat[1].Error() {
			t.Errorf("Expected last error to be %v, got %v", err, flat[1])
		}
		if err.Error() != peek.Error() {
			t.Errorf("Expected peek to be error to be %v, got %v", err, peek)
		}
	})
}
