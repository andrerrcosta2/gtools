// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package ranges

import (
	"github.com/andrerrcosta2/gtools/core/internal/assertlite"
	"testing"
)

func TestReverseClosedRange(t *testing.T) {
	tests := []struct {
		name       string
		start, end int
		operations []func(r Range) bool
		expected   []int
	}{
		{
			name:  "Basic reverse traversal",
			start: 10, end: 5,
			operations: []func(r Range) bool{
				func(r Range) bool { return r.Next() },
				func(r Range) bool { return r.Next() },
				func(r Range) bool { return r.Next() },
			},
			expected: []int{8, 7, 6},
		},
		{
			name:  "Traversal reaching end",
			start: 10, end: 5,
			operations: []func(r Range) bool{
				func(r Range) bool { return r.Next() },
				func(r Range) bool { return r.Next() },
				func(r Range) bool { return r.Next() },
				func(r Range) bool { return r.Next() },
			},
			expected: []int{8, 7, 6},
		},
		{
			name:  "Prev method test",
			start: 10, end: 5,
			operations: []func(r Range) bool{
				func(r Range) bool { r.Next(); return r.Prev() },
				func(r Range) bool { r.Next(); return r.Prev() },
			},
			expected: []int{9, 9},
		},
		{
			name:  "Reset and restart",
			start: 10, end: 5,
			operations: []func(r Range) bool{
				func(r Range) bool { r.Next(); return true },
				func(r Range) bool { r.Reset(); return r.Next() },
			},
			expected: []int{8, 8},
		},
		{
			name:  "Snap method test",
			start: 10, end: 5,
			operations: []func(r Range) bool{
				func(r Range) bool { r.Next(); return true },
				func(r Range) bool {
					snapshot := r.Snap().(Range)
					return snapshot.Current() == r.Current()
				},
			},
			expected: []int{8, 8},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := New(ReverseClosed, tt.start, tt.end)
			var results []int
			for _, op := range tt.operations {
				if op(r) {
					results = append(results, r.Current())
				}
			}
			assertlite.ArrayEquals(t, results, tt.expected)
		})
	}
}

func TestReverseClosedRangeSize(t *testing.T) {
	r := New(ReverseClosed, 10, 5)
	expectedSize := 4
	if r.Size() != expectedSize {
		t.Errorf("Size() = %v; want %v", r.Size(), expectedSize)
	}
}

func TestReverseClosedRangeSetRange(t *testing.T) {
	r := New(ReverseClosed, 10, 5)
	xs, xe := 8, 6
	err := r.SetRange(8, 6)
	if err != nil {
		t.Errorf("SetRange() failed: %v", err)
	}

	if r.Start() != xs || r.End() != xe {
		t.Errorf("SetRange() failed; start = %v, end = %v; want start = %v, end = %v", r.Start(), r.End(), xs, xe)
	}

	// Try set an invalid range
	err = r.SetRange(20, 21)
	if err == nil {
		t.Error("SetRange() should have failed")
	}
}

func TestReverseClosedRangeTune(t *testing.T) {
	r := New(ReverseClosed, 10, 5)
	r.Tune(2, 1)
	xs, xe := 11, 7
	if r.Start() != xs || r.End() != xe {
		t.Errorf("Tune() failed; start = %v, end = %v; want start = %v, end = %v", r.Start(), r.End(), xs, xe)
	}
}
