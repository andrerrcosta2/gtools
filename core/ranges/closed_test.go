// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package ranges

import "testing"

func TestClosedRange_Current(t *testing.T) {
	// Test 1
	r := New(Closed, 1, 10)
	if r.Current() != 2 {
		t.Errorf("Expected 2, got %d", r.Current())
	}
}

func TestClosedRange_Next(t *testing.T) {
	// Test 1
	r := New(Closed, 1, 10)
	r.Next()
	if r.Current() != 3 {
		t.Errorf("Expected 3, got %d", r.Current())
	}
	for i := 0; i < 20; i++ {
		r.Next()
	}
	if r.Current() != 9 {
		t.Errorf("Expected 9, got %d", r.Current())
	}
}

func TestClosedRange_Prev(t *testing.T) {
	// Test 1
	r := New(Closed, 1, 10)
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

func TestClosedRange_Reset(t *testing.T) {
	// Test 1
	r := New(Closed, 1, 10)
	r.Next()
	if r.Current() != 3 {
		t.Errorf("Expected 3, got %d", r.Current())
	}
	r.Reset()
	if r.Current() != 2 {
		t.Errorf("Expected 2, got %d", r.Current())
	}
}

func TestClosedRange_Size(t *testing.T) {
	// Test 1
	r := New(Closed, 1, 10)
	if r.Size() != 8 {
		t.Errorf("Expected 8, got %d", r.Size())
	}
}
