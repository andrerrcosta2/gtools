// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package ranges

import (
	"github.com/andrerrcosta2/gtools/core/testlite/assertlite"
	"testing"
)

func TestReverseRightClosedRange(t *testing.T) {
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
			expected: []int{9, 8, 7},
		},
		{
			name:  "Traversal reaching end",
			start: 10, end: 7,
			operations: []func(r Range) bool{
				func(r Range) bool { return r.Next() },
				func(r Range) bool { return r.Next() },
				func(r Range) bool { return r.Next() },
				func(r Range) bool { return r.Next() },
			},
			expected: []int{9, 8},
		},
		{
			name:  "Prev method test",
			start: 10, end: 5,
			operations: []func(r Range) bool{
				func(r Range) bool { r.Next(); return r.Prev() },
				func(r Range) bool { r.Next(); return r.Prev() },
			},
			expected: []int{10, 10},
		},
		{
			name:  "Reset and restart",
			start: 10, end: 5,
			operations: []func(r Range) bool{
				func(r Range) bool { r.Next(); return true },
				func(r Range) bool { r.Reset(); return r.Next() },
			},
			expected: []int{9, 9},
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
			expected: []int{9, 9},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := New(ReverseRightClosed, tt.start, tt.end)
			var results []int
			for _, op := range tt.operations {
				if op(r) {
					results = append(results, r.Current())
				}
			}
			assertlite.EqualSlices(t, results, tt.expected)
		})
	}
}

func TestReverseRightClosedRangeSize(t *testing.T) {
	r := New(ReverseRightClosed, 10, 5)
	expectedSize := 5
	if r.Size() != expectedSize {
		t.Errorf("Size() = %v; want %v", r.Size(), expectedSize)
	}
}

func TestReverseRightClosedRangeSetRange(t *testing.T) {
	r := New(ReverseRightClosed, 10, 5)
	r.SetRange(20, 15)
	expectedStart, expectedEnd := 20, 15
	if r.Start() != expectedStart || r.End() != expectedEnd {
		t.Errorf("SetRange() failed; start = %v, end = %v; want start = %v, end = %v", r.Start(), r.End(), expectedStart, expectedEnd)
	}
}

func TestReverseRightClosedRangeTune(t *testing.T) {
	r := New(ReverseRightClosed, 10, 5)
	xs, xe := 12, 5
	r.Tune(2, -1)
	if r.Start() != xs || r.End() != xe {
		t.Errorf("Tune() failed; start = %v, end = %v; want start = %v, end = %v", r.Start(), r.End(), xs, xe)
	}
}
