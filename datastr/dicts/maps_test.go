// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package dicts

import (
	"testing"
)

func TestNewComparableEntry(t *testing.T) {
	e1 := NewComparableEntry(1, 2)
	if e1.Key() != 1 || e1.Value() != 2 {
		t.Errorf("NewComparableEntry() = %v, want %v", e1, &anyEntry[int, int]{1, 2})
	}
	e2 := NewComparableEntry("key", "value")
	if e2.Key() != "key" || e2.Value() != "value" {
		t.Errorf("NewComparableEntry() = %v, want %v", e2, &anyEntry[string, string]{"key", "value"})
	}
}
