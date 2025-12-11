// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package gerrors

import (
	"errors"
	"fmt"
	"testing"
)

var (
	flatA         = errors.New("flat A")
	flatB         = errors.New("flat B")
	flatC         = errors.New("flat C")
	wrapOne       = fmt.Errorf("wrap one: %w", flatA)
	wrapTwo       = fmt.Errorf("wrap two: %w", wrapOne)
	wrapThree     = fmt.Errorf("wrap three: %w", wrapTwo)
	joinFlatOne   = errors.Join(flatA)
	joinFlatTwo   = errors.Join(flatA, flatB)
	joinFlatThree = errors.Join(flatA, flatB, flatC)
	joinWrapOne   = errors.Join(wrapOne)
	joinWrapTwo   = errors.Join(wrapTwo)
	joinWrapThree = errors.Join(wrapThree)
	joinJoined    = errors.Join(joinFlatOne, joinFlatTwo, joinFlatThree)
	joinComplex   = errors.Join(flatA, wrapTwo, joinFlatThree, joinJoined)
)

var baseCases = []struct {
	name   string
	err    error
	cause  error
	length int
	str    string
}{
	{name: "flatA", err: flatA, cause: flatA, length: 1, str: "flat A"},
	{name: "flatB", err: flatB, cause: flatB, length: 1, str: "flat B"},
	{name: "flatC", err: flatC, cause: flatC, length: 1, str: "flat C"},
	// wrapOne = fmt.Errorf("wrap one: %w", flatA)
	// flatten = [wrapOne -> flatA]
	{name: "wrapOne", err: wrapOne, cause: flatA, length: 1, str: "wrap one: flat A"},

	// wrapTwo = wrapTwo(wrapOne(flatA))
	// flatten = [wrapTwo -> wrapOne -> flatA]
	{name: "wrapTwo", err: wrapTwo, cause: flatA, length: 1, str: "wrap two: wrap one: flat A"},

	// wrapThree = wrapThree(wrapTwo(wrapOne(flatA)))
	// flatten = [wrapThree -> wrapTwo -> wrapOne -> flatA]
	{name: "wrapThree", err: wrapThree, cause: flatA, length: 1, str: "wrap three: wrap two: wrap one: flat A"},

	// Join(flatA)
	// → [flatA]
	{name: "joinFlatOne", err: joinFlatOne, cause: flatA, length: 1, str: "flat A"},

	// Join(flatA, flatB)
	// → [flatA, flatB]
	{name: "joinFlatTwo", err: joinFlatTwo, cause: flatA, length: 2, str: "flat A\nflat B"},

	// Join(flatA, flatB, flatC)
	// → [flatA, flatB, flatC]
	{name: "joinFlatThree", err: joinFlatThree, cause: flatA, length: 3, str: "flat A\nflat B\nflat C"},

	// Join(wrapOne) → wrapOne→flatA
	// flatten = [wrapOne -> flatA]
	{name: "joinWrapOne", err: joinWrapOne, cause: flatA, length: 1, str: wrapOne.Error()},

	// Join(wrapTwo) → wrapTwo→wrapOne→flatA
	// flatten = [wrap two -> wrap one -> flat A]
	{name: "joinWrapTwo", err: joinWrapTwo, cause: flatA, length: 1, str: wrapTwo.Error()},

	// Join(wrapThree)
	// flatten = [wrap three -> wrap two -> wrap one -> flat A]
	{name: "joinWrapThree", err: joinWrapThree, cause: flatA, length: 1, str: wrapThree.Error()},

	// joinJoined = Join(joinFlatOne, joinFlatTwo, joinFlatThree)
	// joinFlatOne   → [flatA]
	// joinFlatTwo   → [flatA, flatB]
	// joinFlatThree → [flatA, flatB, flatC]
	// flatten = [flat A, flat A, flat B, flat A, flat B, flat C]
	{name: "joinJoined", err: joinJoined, cause: flatA, length: 6, str: joinJoined.Error()},

	// joinComplex = Join(flatA, wrapTwo, joinFlatThree, joinJoined)
	//   joinJoined
	//   joinFlatThree
	//   wrapTwo
	//   flatA
	// joinJoined → (flatC, flatB, flatA)
	// joinFlatThree → (flatC*, flatB*, flatA*)
	// wrapTwo → wrapTwo, wrapOne, flatA
	// flatA → flatA*
	// flatten = [flat A, wrap two -> wrap one -> flat A, flat A, flat B, flat C, flat A, flat A, flat B,
	// 		flat A, flat B, flat C]
	{name: "joinComplex", err: joinComplex, cause: flatA, length: 11, str: joinComplex.Error()},
}

func TestStackOf(t *testing.T) {
	t.Run("single error", func(t *testing.T) {
		st := StackOf(flatA)
		var cast *stackErr
		ok := errors.As(st, &cast)
		if !ok {
			t.Error("cast err")
		}
		if st.Len() != 1 {
			t.Errorf("expected st.Len() == 1, got '%d'", st.Len())
		}
		if len(cast.stk.stk) != 1 {
			t.Errorf("expected len(st.stk.stk) == 1, got '%d'", len(cast.stk.stk))
		}
	})

	t.Run("multiple errors", func(t *testing.T) {
		st := StackOf(flatA, flatB, flatC)
		var cast *stackErr
		ok := errors.As(st, &cast)
		if !ok {
			t.Error("cast err")
		}
		if st.Len() != 3 {
			t.Errorf("expected st.Len() == 3, got '%d'", st.Len())
		}
		if len(cast.stk.stk) != 3 {
			t.Errorf("expected len(st.stk.stk) == 3, got %v", len(cast.stk.stk))
		}
	})

	t.Run("mil errors", func(t *testing.T) {
		st := StackOf(nil, nil)
		var cast *stackErr
		ok := errors.As(st, &cast)
		if !ok {
			t.Error("cast err")
		}
		if st.Len() != 0 {
			t.Errorf("expected st.Len() == 0, got '%d'", st.Len())
		}
		if len(cast.stk.stk) != 0 {
			t.Errorf("expected len(st.stk.stk) == 0, got %v", len(cast.stk.stk))
		}
	})
}

func TestConcStackOf(t *testing.T) {
	t.Run("single error", func(t *testing.T) {
		st := ConcStackOf(flatA)
		var cast *concStackErr
		ok := errors.As(st, &cast)
		if !ok {
			t.Error("cast err")
		}
		if st.Len() != 1 {
			t.Errorf("expected st.Len() == 1, got '%d'", st.Len())
		}
		if len(cast.stk.stk) != 1 {
			t.Errorf("expected len(st.stk.stk) == 1, got '%d'", len(cast.stk.stk))
		}
	})

	t.Run("multiple errors", func(t *testing.T) {
		st := ConcStackOf(flatA, flatB, flatC)
		var cast *concStackErr
		ok := errors.As(st, &cast)
		if !ok {
			t.Error("cast err")
		}
		if st.Len() != 3 {
			t.Errorf("expected st.Len() == 3, got '%d'", st.Len())
		}
		if len(cast.stk.stk) != 3 {
			t.Errorf("expected len(st.stk.stk) == 3, got %v", len(cast.stk.stk))
		}
	})

	t.Run("mil errors", func(t *testing.T) {
		st := ConcStackOf(nil, nil)
		var cast *concStackErr
		ok := errors.As(st, &cast)
		if !ok {
			t.Error("cast err")
		}
		if st.Len() != 0 {
			t.Errorf("expected st.Len() == 0, got '%d'", st.Len())
		}
		if len(cast.stk.stk) != 0 {
			t.Errorf("expected len(st.stk.stk) == 0, got %v", len(cast.stk.stk))
		}
	})
}

func TestStackErr_Cause(t *testing.T) {
	for _, tc := range baseCases {
		t.Run(tc.name, func(t *testing.T) {
			stk := StackOf(tc.err)
			cstk := ConcStackOf(tc.err)
			got := stk.Cause()

			if !errors.Is(got, tc.cause) {
				t.Errorf("Cause() = %v, expected %v",
					got, tc.cause)
			}
			if stk.Len() != tc.length {
				t.Errorf("expected st.Len() == %d, got '%d': %v", tc.length, stk.Len(), stk.Unwrap())
			}

			got = cstk.Cause()
			if !errors.Is(got, tc.cause) {
				t.Errorf("Cause() = %v, expected %v",
					got, tc.cause)
			}
			if cstk.Len() != tc.length {
				t.Errorf("expected st.Len() == %d, got '%d': %v", tc.length, stk.Len(), stk.Unwrap())
			}

		})
	}
}

func TestStackErr_Error(t *testing.T) {
	for _, tt := range baseCases {
		t.Run(tt.name, func(t *testing.T) {
			stk := StackOf(tt.err)
			cstk := ConcStackOf(tt.err)

			got := stk.Error()
			if got != tt.str {
				t.Errorf("failed. want %q got %q", tt.str, got)
			}
			got = cstk.Error()
			if got != tt.str {
				t.Errorf("failed. want %q got %q", tt.str, got)
			}
		})
	}
}

func TestStackErr_From(t *testing.T) {
	cases := []struct {
		name string
		prev error // first stack: old
		next error // second: replaces bottom
		want []string
	}{
		{
			name: "prev=flatA, next=flatB",
			prev: flatA,
			next: flatB,
			want: []string{
				"flat B", // from next
				"flat A", // from prev
			},
		},
		{
			name: "prev=wrapTwo, next=flatC",
			prev: wrapTwo, // preserved wrapper
			next: flatC,
			want: []string{
				"flat C",
				"wrap two: wrap one: flat A",
			},
		},
		{
			name: "prev=joinFlatTwo, next=flatC",
			prev: joinFlatTwo, // A,B
			next: flatC,
			want: []string{
				"flat C",
				"flat A", "flat B",
			},
		},
		{
			name: "prev=flatA, next=joinFlatThree",
			prev: flatA,
			next: joinFlatThree, // A,B,C
			want: []string{
				"flat A", "flat B", "flat C",
				"flat A",
			},
		},
		{
			name: "prev=joinFlatThree, next=wrapThree",
			prev: joinFlatThree, // A,B,C
			next: wrapThree,     // preserved wrapper
			want: []string{
				"wrap three: wrap two: wrap one: flat A",
				"flat A", "flat B", "flat C",
			},
		},
		{
			name: "prev=joinJoined, next=flatB",
			prev: joinJoined, // A | A,B | A,B,C
			next: flatB,
			want: []string{
				"flat B",
				"flat A",
				"flat A", "flat B",
				"flat A", "flat B", "flat C",
			},
		},
		{
			name: "prev=flatA, next=joinComplex",
			prev: flatA,
			next: joinComplex,
			want: append(
				// flatten(joinComplex)
				[]string{
					"flat A",
					"wrap two: wrap one: flat A",
					"flat A", "flat B", "flat C",
					"flat A",
					"flat A", "flat B",
					"flat A", "flat B", "flat C",
				},
				// then prev
				"flat A",
			),
		},
		{
			name: "prevEmpty_nextFlatA",
			prev: nil, // using Stack() below
			next: flatA,
			want: []string{
				"flat A",
			},
		},
		{
			name: "prevFlatB_nextNil",
			prev: flatB,
			next: nil,
			want: []string{
				"flat B",
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {

			// --- NON-CONCURRENT CASE ---
			var prevStk *stackErr
			if !errors.As(StackOf(tt.prev), &prevStk) {
				t.Fatal("StackOf(prev) did not return *stackErr")
			}

			var newStk *stackErr
			if !errors.As(prevStk.From(tt.next), &newStk) {
				t.Fatal("prevStk.From(next) did not return *stackErr")
			}

			// --- CONCURRENT CASE ---
			var prevConc *concStackErr
			if !errors.As(ConcStackOf(tt.prev), &prevConc) {
				t.Fatal("ConcStackOf(prev) did not return *concStackErr")
			}

			var newConc *concStackErr
			if !errors.As(prevConc.From(tt.next), &newConc) {
				t.Fatal("prevConc.From(next) did not return *concStackErr")
			}

			// --- EXTRACT BOTH RESULTS ---
			got := make([]string, newStk.stk.Len())
			for i, err := range newStk.stk.stk {
				got[i] = err.Error()
			}

			gotConc := make([]string, newConc.stk.Len())
			for i, err := range newConc.stk.stk {
				gotConc[i] = err.Error()
			}

			// --- ASSERT LENGTH ---
			if len(got) != len(tt.want) {
				t.Fatalf("non-concurrent: len=%d want=%d", len(got), len(tt.want))
			}
			if len(gotConc) != len(tt.want) {
				t.Fatalf("concurrent: len=%d want=%d", len(gotConc), len(tt.want))
			}

			// --- ASSERT CONTENT ---
			for i := range got {
				if got[i] != tt.want[i] {
					t.Fatalf("non-concurrent index %d: got %q want %q", i, got[i], tt.want[i])
				}
			}

			for i := range gotConc {
				if gotConc[i] != tt.want[i] {
					t.Fatalf("concurrent index %d: got %q want %q", i, gotConc[i], tt.want[i])
				}
			}
		})
	}
}

func TestStackErr_Is(t *testing.T) {
	cases := []struct {
		name   string
		prev   error
		target error
		want   bool
	}{
		// --- FLAT vs FLAT ---
		{"flat contains same flat", flatA, flatA, true},
		{"flat contains different flat", flatA, flatB, false},

		// --- FLAT vs WRAP ---
		// Go: errors.Is(flatA, wrapOne) == false
		// Your semantics: flatA does NOT match wrapOne
		{"flat matches wrapped(flatA)", flatA, wrapOne, false},
		{"flat matches deep wrapped(flatA)", flatA, wrapThree, false},

		// --- FLAT vs JOIN ---
		{"flat matches join(flatA)", flatA, joinFlatOne, false}, // important!
		{"flat matches join(flatA, flatB)", flatA, joinFlatTwo, false},

		// --- WRAP vs FLAT ---
		{"wrapOne contains flatA", wrapOne, flatA, true},
		{"wrapTwo contains flatA", wrapTwo, flatA, true},
		{"wrapTwo does not contain flatB", wrapTwo, flatB, false},

		// --- WRAP vs WRAP ---
		{"wrapThree contains wrapTwo", wrapThree, wrapTwo, true},
		{"wrapOne does not match wrapThree", wrapOne, wrapThree, false},

		// --- WRAP vs JOIN ---
		{"wrapTwo matches joinWrapTwo", wrapTwo, joinWrapTwo, false},
		{"wrapTwo matches join(flatA, wrapTwo)", wrapTwo, joinComplex, false},

		// --- JOIN vs FLAT ---
		{"join(flatA) contains flatA", joinFlatOne, flatA, true},
		{"join(flatA, flatB) contains flatB", joinFlatTwo, flatB, true},
		{"join(flatA, flatB) does not contain flatC", joinFlatTwo, flatC, false},

		// --- JOIN vs WRAP ---
		// join(flatA,flatB,flatC) contains flatA, so true
		{"join(flatA, flatB, flatC) matches wrapOne? flatA inside", joinFlatThree, wrapOne, false},
		{"joinWrapTwo contains flatA?", joinWrapTwo, flatA, true},

		// --- JOIN vs JOIN ---
		// YOUR semantics: join != join, even if identical
		{"join(flatA) matches join(flatA)?", joinFlatOne, joinFlatOne, false},
		{"join(flatA,flatB) matches join(flatA)?", joinFlatTwo, joinFlatOne, false},
		{"joinJoined contains flatC", joinJoined, flatC, true},
		{"joinJoined contains wrapTwo?", joinJoined, wrapTwo, false},

		// --- COMPLEX JOIN ---
		{"joinComplex contains flatA", joinComplex, flatA, true},
		{"joinComplex contains wrapTwo", joinComplex, wrapTwo, true},
		{"joinComplex contains wrapThree", joinComplex, wrapThree, false},
		{"joinComplex contains flatC (via joinFlatThree)", joinComplex, flatC, true},

		// --- EDGE CASES ---
		{"empty stack does not match non-nil", nil, flatA, false},
		// YOUR semantics: nil Is nil = true? Or false?
		// Choose one; Go says false, but you previously said you want true.
		{"empty stack matches nil target?", nil, nil, true},
		{"non-empty stack with nil target", wrapOne, nil, false},
		{"stackOf(nil) behaves as empty", StackableOf(nil), flatA, false},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			s := StackableOf(tt.prev)
			got := s.Is(tt.target)

			if got != tt.want {
				t.Errorf(
					"%s: StackableOf(prev).Is(target) = %v, want %v\nprev=%v\ntarget=%v",
					tt.name, got, tt.want, tt.prev, tt.target,
				)
			}

			s = ConcStackableOf(tt.prev)
			got = s.Is(tt.target)

			if got != tt.want {
				t.Errorf(
					"%s: StackableOf(prev).Is(target) = %v, want %v\nprev=%v\ntarget=%v",
					tt.name, got, tt.want, tt.prev, tt.target,
				)
			}
		})
	}
}
