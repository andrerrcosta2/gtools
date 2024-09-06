// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package ranges

import "errors"

type Range interface {
	// Next tries to increase the range value and returns true if the value was increased
	Next() bool
	// Prev tries to decrease the range value and returns true if the value was decreased
	Prev() bool
	// Size returns the size of the range
	Size() int
	// Current returns the current value
	Current() int
	// Reset resets the range to its start value
	Reset()
	// SetRange sets the range and returns an error if the range is invalid
	SetRange(start int, end int) error
	// Tune this method allows changing the range interval by increasing or decreasing the start and end
	Tune(start int, end int)
	// Start returns the start of the range
	Start() int
	// End returns the end of the range
	End() int
	// Snap returns a copy of the range
	Snap() Range
	// SetCurrent sets the current value and returns true if the value was updated
	SetCurrent(current int) bool
	// First returns the first value
	First() int
	// Last returns the last value
	Last() int
	// IsOnStart returns true if the range is on the start
	IsOnStart() bool
	// IsOnEnd returns true if the range is on the end
	IsOnEnd() bool
	// RideNext advances the range and returns true if the range didn't reach the end
	RideNext(n int) bool
	// RidePrev advances the range and returns true if the range didn't reach the start
	RidePrev(n int) bool
}

// nextAsc advances the range and returns true if the range didn't reach the end
func nextAsc(r Range) bool {
	if r.Current() < r.Last() {
		r.SetCurrent(r.Current() + 1)
		return true
	}
	return false
}

// nextDesc advances the range and returns true if the range didn't reach the end
func nextDesc(r Range) bool {
	if r.Current() > r.Last() {
		r.SetCurrent(r.Current() - 1)
		return true
	}
	return false
}

// prevAsc advances the range and returns true if the range didn't reach the end
func prevAsc(r Range) bool {
	if r.Current() > r.First() {
		r.SetCurrent(r.Current() - 1)
		return true
	}
	return false
}

func sizeAsc(r Range) int {
	return r.Last() - r.First()
}

func sizeDesc(r Range) int {
	return r.First() - r.Last()
}

// prevDesc advances the range and returns true if the range didn't reach the end
func prevDesc(r Range) bool {
	if r.Current() < r.First() {
		r.SetCurrent(r.Current() + 1)
		return true
	}
	return false
}

func rideNextAsc(r Range, n int) bool {
	if r.Current()+n <= r.Last() {
		r.SetCurrent(r.Current() + n)
		return true
	} else {
		r.SetCurrent(r.Last())
	}
	return false
}

func ridePrevAsc(r Range, n int) bool {
	if r.Current()-n >= r.First() {
		r.SetCurrent(r.Current() - n)
		return true
	}
	r.SetCurrent(r.First())
	return false
}

func rideNextDesc(r Range, n int) bool {
	if r.Current()-n >= r.Last() {
		r.SetCurrent(r.Current() - n)
		return true
	}
	r.SetCurrent(r.Last())
	return false
}

func ridePrevDesc(r Range, n int) bool {
	if r.Current()+n <= r.First() {
		r.SetCurrent(r.Current() + n)
		return true
	}
	r.SetCurrent(r.First())
	return false
}

func reset(r Range) {
	r.SetCurrent(r.First())
}

func canSetRangeAsc(r Range, start int, end int) error {
	if start > end {
		return errors.New("start must be less than end")
	}
	if start < r.First() {
		return errors.New("start must be greater than or equal to the first element")
	}
	if end > r.Last() {
		return errors.New("end must be less than or equal to the last element")
	}
	return nil
}

func canSetRangeDesc(r Range, start int, end int) error {
	if start < end {
		return errors.New("start must be greater than end")
	}
	if start > r.First() {
		return errors.New("start must be less than or equal to the first element")
	}
	if end < r.Last() {
		return errors.New("end must be greater than or equal to the last element")
	}
	return nil
}

func canSetCurrentAsc(r Range, current int) bool {
	if current > r.Last() {
		return false
	}
	if current < r.First() {
		return false
	}
	return true
}

func canSetCurrentDesc(r Range, current int) bool {
	if current < r.Last() {
		return false
	}
	if current > r.First() {
		return false
	}
	return true
}

type Type int

const (
	Closed Type = iota
	LeftClosed
	RightClosed
	Open
	ReverseClosed
	ReverseLeftClosed
	ReverseRightClosed
	ReverseOpen
	ConcurrentClosed
	ConcurrentLeftClosed
	ConcurrentRightClosed
	ConcurrentOpen
	ConcurrentReverseClosed
	ConcurrentReverseLeftClosed
	ConcurrentReverseRightClosed
	ConcurrentReverseOpen
)
