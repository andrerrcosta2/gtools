// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package typers

import (
	"testing"
)

func TestGens_Rand(t *testing.T) {
	ints, err := Rand[int](0, 200, 10)
	if err != nil {
		t.Error(err)
	}

	for _, va := range ints {
		if _, ok := any(va).(int); !ok {
			t.Errorf("expected int, got %T", va)
		}
	}

	strs, err := Rand[string](0, 200, 10)
	if err != nil {
		t.Error(err)
	}
	for _, va := range strs {
		if _, ok := any(va).(string); !ok {
			t.Errorf("expected string, got %T", va)
		}
	}
}
