// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sentinels

import "github.com/andrerrcosta2/gtools/core/ranges"

// Int returns an int sentinel with no ranges
func Int(value int) *IntSentinel {
	return &IntSentinel{
		value: value,
	}
}

type IntSentinel struct {
	value int
}

func (s *IntSentinel) Next() bool {
	s.value++
	return true
}

func (s *IntSentinel) Inc(i int) bool {
	s.value += i
	return true
}

func (s *IntSentinel) Prev() bool {
	s.value--
	return true
}

func (s *IntSentinel) Dec(i int) bool {
	s.value -= i
	return true
}

func (s *IntSentinel) Set(i int) bool {
	s.value = i
	return true
}

func (s *IntSentinel) Less(other int) bool {
	return s.value < other
}

func (s *IntSentinel) Equal(other int) bool {
	return s.value == other
}

func (s *IntSentinel) Val() int {
	return s.value
}

var _ Sentinel = (*IntSentinel)(nil)

func Ranged(startValue int, rng ranges.Range) RangedSentinel {
	s := &IntRangedSentinel{
		rng: rng,
	}
	s.rng.SetCurrent(startValue)
	return s
}

type IntRangedSentinel struct {
	rng ranges.Range
}

func (s *IntRangedSentinel) Next() bool {
	return s.rng.Next()
}

func (s *IntRangedSentinel) HasNext() bool {
	return !s.rng.IsOnEnd()
}

func (s *IntRangedSentinel) Inc(i int) bool {
	return s.rng.RideNext(i)
}

func (s *IntRangedSentinel) Prev() bool {
	return s.rng.Prev()
}

func (s *IntRangedSentinel) HasPrev() bool {
	return !s.rng.IsOnStart()
}

func (s *IntRangedSentinel) Dec(i int) bool {
	return s.rng.RidePrev(i)
}

func (s *IntRangedSentinel) Set(i int) bool {
	return s.rng.SetCurrent(i)
}

func (s *IntRangedSentinel) Walk(val int) bool {
	return s.rng.SetCurrent(val)
}

func (s *IntRangedSentinel) Less(other int) bool {
	return s.rng.Current() < other
}

func (s *IntRangedSentinel) Equal(other int) bool {
	return s.rng.Current() == other
}

func (s *IntRangedSentinel) Val() int {
	return s.rng.Current()
}

func (s *IntRangedSentinel) FromStart() int {
	if s.rng.First() > s.rng.Last() {
		return s.rng.First() - s.rng.Current()
	}
	return s.rng.Current() - s.rng.First()
}

func (s *IntRangedSentinel) ToEnd() int {
	if s.rng.First() > s.rng.Last() {
		return s.rng.Last() - s.rng.Current()
	}
	return s.rng.Current() - s.rng.Last()
}

// RangedCloseable is a sentinel that closes itself when it reaches the max value
func RangedCloseable(rng ranges.Range) *RangedCloseableSentinel {
	return &RangedCloseableSentinel{
		rng: rng,
	}
}

type RangedCloseableSentinel struct {
	rng      ranges.Range
	isClosed bool
}

func (s *RangedCloseableSentinel) Next() bool {
	if s.IsOpen() {
		return s.rng.Next()
	}
	return false
}

func (s *RangedCloseableSentinel) HasNext() bool {
	return !s.rng.IsOnEnd()
}

func (s *RangedCloseableSentinel) Inc(i int) bool {
	if s.IsOpen() {
		if !s.rng.RideNext(i) {
			s.Close()
			return false
		}
	}
	return false
}

func (s *RangedCloseableSentinel) Prev() bool {
	if s.IsOpen() {
		return s.rng.Prev()
	}
	return false
}

func (s *RangedCloseableSentinel) HasPrev() bool {
	return !s.rng.IsOnStart()
}

func (s *RangedCloseableSentinel) Dec(i int) bool {
	if s.IsOpen() {
		if s.rng.RidePrev(i) {
			return true
		}
		s.Close()
		return false
	}
	return false
}

func (s *RangedCloseableSentinel) Set(i int) bool {
	return s.rng.SetCurrent(i)
}

func (s *RangedCloseableSentinel) Less(other int) bool {
	return s.rng.Current() < other
}

func (s *RangedCloseableSentinel) Equal(other int) bool {
	return s.rng.Current() == other
}

func (s *RangedCloseableSentinel) Val() int {
	return s.rng.Current()
}

func (s *RangedCloseableSentinel) FromStart() int {
	return s.rng.DistanceFromStart()
}

func (s *RangedCloseableSentinel) ToEnd() int {
	return s.rng.DistanceToEnd()
}

func (s *RangedCloseableSentinel) IsClosed() bool {
	return s.isClosed
}

func (s *RangedCloseableSentinel) IsOpen() bool {
	return !s.isClosed
}

func (s *RangedCloseableSentinel) Open() {
	s.isClosed = false
}

func (s *RangedCloseableSentinel) Close() {
	s.isClosed = true
}

var _ CloseableSentinel = (*RangedCloseableSentinel)(nil)
