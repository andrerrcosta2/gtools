// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package ranges

import "testing"

func TestRightClosedRange_Current(t *testing.T) {
	// Test 1
	r := New(RightClosed, 1, 10)
	if r.Current() != 1 {
		t.Errorf("Expected 1, got %d", r.Current())
	}

	// Test 2
	r.Next()
	if r.Current() != 2 {
		t.Errorf("Expected 2, got %d", r.Current())
	}

	// Test 3
	for i := 0; i < 20; i++ {
		r.Next()
	}
	if r.Current() != 9 {
		t.Errorf("Expected 9, got %d", r.Current())
	}
}

func TestRightClosedRange_Next(t *testing.T) {
	// Test 1
	r := New(RightClosed, 1, 10)
	r.Next()
	if r.Current() != 2 {
		t.Errorf("Expected 2, got %d", r.Current())
	}

	// Test 2
	r = New(RightClosed, 1, 10)
	for i := 0; i < 20; i++ {
		r.Next()
	}
	if r.Current() != 9 {
		t.Errorf("Expected 9, got %d", r.Current())
	}
}

func TestRightClosedRange_Prev(t *testing.T) {
	// Test 1
	r := New(RightClosed, 1, 10)
	r.Next()
	if r.Current() != 2 {
		t.Errorf("Expected 2, got %d", r.Current())
	}
	r.Prev()
	if r.Current() != 1 {
		t.Errorf("Expected 1, got %d", r.Current())
	}
}

func TestRightClosedRange_Reset(t *testing.T) {
	// Test 1
	r := New(RightClosed, 1, 10)
	r.Next()
	if r.Current() != 2 {
		t.Errorf("Expected 2, got %d", r.Current())
	}
	r.Reset()
	if r.Current() != 1 {
		t.Errorf("Expected 1, got %d", r.Current())
	}
}

func TestRightClosedRange_Size(t *testing.T) {
	// Test 1
	r := New(RightClosed, 1, 10)
	if r.Size() != 9 {
		t.Errorf("Expected 10, got %d", r.Size())
	}
}
