// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package ranges

import "testing"

func TestLeftClosedRange_Current(t *testing.T) {
	// Test 1
	r := New(LeftClosed, 1, 10)
	if r.Current() != 2 {
		t.Errorf("Expected 2, got %d", r.Current())
	}
}

func TestLeftClosedRange_Next(t *testing.T) {
	// Test 1
	r := New(LeftClosed, 1, 10)
	r.Next()
	if r.Current() != 3 {
		t.Errorf("Expected 3, got %d", r.Current())
	}

	// Test 2
	r = New(LeftClosed, 1, 10)
	for i := 0; i < 20; i++ {
		r.Next()
	}
	if r.Current() != 10 {
		t.Errorf("Expected 10, got %d", r.Current())
	}
}

func TestLeftClosedRange_Prev(t *testing.T) {
	// Test 1
	r := New(LeftClosed, 1, 10)
	r.Next()
	if r.Current() != 3 {
		t.Errorf("Expected 3, got %d", r.Current())
	}
	for i := 0; i < 20; i++ {
		r.Prev()
	}
	if r.Current() != 2 {
		t.Errorf("Expected 2, got %d", r.Current())
	}
}

func TestLeftClosedRange_Reset(t *testing.T) {
	// Test 1
	r := New(LeftClosed, 1, 10)
	r.Next()
	if r.Current() != 3 {
		t.Errorf("Expected 3, got %d", r.Current())
	}
	r.Reset()
	if r.Current() != 2 {
		t.Errorf("Expected 2, got %d", r.Current())
	}
}

func TestLeftClosedRange_Size(t *testing.T) {
	// Test 1
	r := New(LeftClosed, 1, 10)
	if r.Size() != 9 {
		t.Errorf("Expected 9, got %d", r.Size())
	}
}
