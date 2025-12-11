// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package loop

import (
	"context"
	"errors"
	"fmt"

	"github.com/andrerrcosta2/gtools/core/domain/functions"
	"github.com/andrerrcosta2/gtools/core/format/fmx"
)

// ErrEmptyElements convenience error
var ErrEmptyElements = errors.New("elements slice is empty")

// BF brute force loop, order matters, repetition allowed (cartesian power)
func BF[S ~[]E, E any](ctx context.Context, elems S, min, max, step int, fn functions.Function[S, error]) error {
	if len(elems) == 0 {
		return ErrEmptyElements
	}
	if min < 0 || max < 0 || min > max || step <= 0 {
		return fmx.Errorf("invalid min/max/step: %d/%d/%d", min, max, step)
	}
	for k := min; k <= max; k += step {
		if k == 0 {
			var empty []E
			if err := fn(S(empty)); err != nil {
				return err
			}
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
			}
			continue
		}
		// current tuple
		tuple := make([]E, k)
		var gen func(pos int) error
		gen = func(pos int) error {
			// cancellation check
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
			}
			if pos == k {
				if err := fn(S(tuple)); err != nil {
					return err
				}
				return nil
			}
			for i := 0; i < len(elems); i++ {
				tuple[pos] = elems[i]
				if err := gen(pos + 1); err != nil {
					return err
				}
			}
			return nil
		}
		if err := gen(0); err != nil {
			return err
		}
	}
	return nil
}

// StaticBF convenience wrapper for a single fixed length
func StaticBF[S ~[]E, E any](ctx context.Context, elems S, length int, fn func(combo S) error) error {
	return BF[S, E](ctx, elems, length, length, 1, fn)
}

// BFOrderlessNoRepeat combinations (order doesn't matter), no repetition.
// Generates all subsets of sizes in [min,max]
func BFOrderlessNoRepeat[S ~[]E, E any](ctx context.Context, elems S, min, max, step int, fn func(combo S) error) error {
	n := len(elems)
	if n == 0 {
		return ErrEmptyElements
	}
	if min < 0 || max < 0 || min > max || step <= 0 {
		return fmt.Errorf("invalid min/max/step: %d/%d/%d", min, max, step)
	}
	for k := min; k <= max; k += step {
		if k == 0 {
			var empty []E
			if err := fn(S(empty)); err != nil {
				return err
			}
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
			}
			continue
		}
		indices := make([]int, k)
		// initialize first combination indices 0..k-1
		for i := 0; i < k; i++ {
			indices[i] = i
		}
		for {
			// build combo
			combo := make([]E, k)
			for i := 0; i < k; i++ {
				combo[i] = elems[indices[i]]
			}
			if err := fn(S(combo)); err != nil {
				return err
			}
			// next lexicographic combination
			// find rightmost index we can increment
			j := k - 1
			for j >= 0 && indices[j] == n-k+j {
				j--
			}
			if j < 0 {
				break // finished
			}
			indices[j]++
			for t := j + 1; t < k; t++ {
				indices[t] = indices[t-1] + 1
			}
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
			}
		}
	}
	return nil
}

// BFOrderlessWithRepeat combinations where order doesn't matter but repetition IS allowed
// (multicombinations) -> use non-decreasing index sequences
func BFOrderlessWithRepeat[S ~[]E, E any](ctx context.Context, elems S, min, max, step int, fn func(combo S) error) error {
	n := len(elems)
	if n == 0 {
		return ErrEmptyElements
	}
	if min < 0 || max < 0 || min > max || step <= 0 {
		return fmt.Errorf("invalid min/max/step: %d/%d/%d", min, max, step)
	}
	for k := min; k <= max; k += step {
		if k == 0 {
			var empty []E
			if err := fn(S(empty)); err != nil {
				return err
			}
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
			}
			continue
		}
		indices := make([]int, k) // default zeros -> first multiset [0,0,...,0]
		for {
			// build combo
			combo := make([]E, k)
			for i := 0; i < k; i++ {
				combo[i] = elems[indices[i]]
			}
			if err := fn(S(combo)); err != nil {
				return err
			}
			// increment like adding 1 with non-decreasing constraint
			// find rightmost position that can be incremented (indices[pos] < n-1)
			pos := k - 1
			for pos >= 0 && indices[pos] == n-1 {
				pos--
			}
			if pos < 0 {
				break
			}
			indices[pos]++
			// make tail equal to indices[pos] (to keep non-decreasing)
			for j := pos + 1; j < k; j++ {
				indices[j] = indices[pos]
			}
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
			}
		}
	}
	return nil
}
