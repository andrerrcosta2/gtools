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
		stk := FlattenError(err)
		if len(stk) != 1 {
			t.Errorf("Expected stack length to be 1, got %d", len(stk))
		}
		if !errors.Is(err, stk[0]) {
			t.Errorf("Expected last error to be %v, got %v", err, stk[0])
		}
	})

	t.Run("Multiple errors", func(t *testing.T) {
		// Create a wrapped error
		werr := errors.New("wrapped error")
		stk := StackableOf(werr)
		err := errors.New("nested error")
		stk.Stack(err)
		flat := FlattenError(stk)
		if len(flat) != 2 {
			t.Errorf("Expected stack length to be 2, got %d", len(flat))
		}
		if werr.Error() != flat[0].Error() {
			t.Errorf("Expected first error to be %v, got %v", werr, flat[0])
		}
		if err.Error() != flat[1].Error() {
			t.Errorf("Expected last error to be %v, got %v", err, flat[1])
		}
		if err.Error() != flat[len(flat)-1].Error() {
			t.Errorf("Expected peek to be error to be %v, got %v", err, flat[len(flat)-1])
		}
	})
}
