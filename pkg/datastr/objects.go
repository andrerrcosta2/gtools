// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package datastr

import "errors"

func NewSentinel() *Sentinel {
	return new(Sentinel)
}

type Sentinel int

func (s *Sentinel) Next() {
	*s++
}

func (s *Sentinel) inc(i int) {
	*s += Sentinel(i)
}

func (s *Sentinel) Prev() {
	*s--
}

func (s *Sentinel) dec(i int) {
	*s -= Sentinel(i)
}

func (s *Sentinel) Walk(val int) {
	*s += Sentinel(val)
}

func (s *Sentinel) Less(other int) bool {
	return int(*s) < other
}

func (s *Sentinel) Equal(other int) bool {
	return int(*s) == other
}

func NewDoubleSentinel(left, right int) *DoubleSentinel {
	return &DoubleSentinel{Sentinel(left), Sentinel(right)}
}

type DoubleSentinel struct {
	Left  Sentinel
	Right Sentinel
}

func (s *DoubleSentinel) Next() {
	s.Left.Next()
	s.Right.Next()
}

func (s *DoubleSentinel) WalkLeft(val int) {
	s.Left.Prev()
	s.Right.Next()
}

func (s *DoubleSentinel) inc(i int) {
	s.Left.inc(i)
	s.Right.inc(i)
}

func (s *DoubleSentinel) Prev() {
	s.Left.Prev()
	s.Right.Prev()
}

func (s *DoubleSentinel) dec(i int) {
	s.Left.dec(i)
	s.Right.dec(i)
}

func NewTripleSentinel(left, middle, right int) *TripleSentinel {
	return &TripleSentinel{Sentinel(left), Sentinel(middle), Sentinel(right)}
}

type TripleSentinel struct {
	Left   Sentinel
	Middle Sentinel
	Right  Sentinel
}

type NSentinel struct {
	sentinels []*Sentinel
}

func (s *NSentinel) Next(sentinels ...int) error {
	for sentinel := range sentinels {
		if sentinel > len(s.sentinels) {
			return errors.New("sentinel out of range")
		}
		s.sentinels[sentinel].Next()
	}
	return nil
}

func (s *NSentinel) inc(inc int, sentinels ...int) error {
	for sentinel := range sentinels {
		if sentinel > len(s.sentinels) || s.sentinels[sentinel] == nil {
			return errors.New("sentinel out of range")
		}
		s.sentinels[sentinel].inc(inc)
	}
	return nil
}

func (s *NSentinel) Prev(sentinels ...int) error {
	for sentinel := range sentinels {
		if sentinel > len(s.sentinels) {
			return errors.New("sentinel out of range")
		}
		s.sentinels[sentinel].Prev()
	}
	return nil
}

func (s *NSentinel) Dec(dec int, sentinels ...int) error {
	for sentinel := range sentinels {
		if sentinel > len(s.sentinels) || s.sentinels[sentinel] == nil {
			return errors.New("sentinel out of range")
		}
		s.sentinels[sentinel].dec(dec)
	}
	return nil
}

func (s *NSentinel) Less(sentinel int, value int) bool {
	if sentinel > len(s.sentinels) {
		return false
	}
	return s.sentinels[sentinel].Less(value)
}

func (s *NSentinel) Equal(sentinel int, value int) bool {
	if sentinel > len(s.sentinels) {
		return false
	}
	return s.sentinels[sentinel].Equal(value)
}

func NewRangedSentinel(min, max int) *RangedSentinel {
	return &RangedSentinel{Sentinel(min), min, max}
}

type RangedSentinel struct {
	value Sentinel
	min   int
	max   int
}

func (s *RangedSentinel) Next() bool {
	if int(s.value) < s.max {
		s.value.Next()
		return true
	}
	return false
}

func (s *RangedSentinel) inc(i int) bool {
	if int(s.value+1) <= s.max {
		s.value.inc(i)
		return true
	}
	return false
}

func (s *RangedSentinel) Prev() bool {
	if int(s.value) > s.min {
		s.value.Prev()
		return true
	}
	return false
}

func (s *RangedSentinel) dec(i int) bool {
	if int(s.value-1) >= s.min {
		s.value.dec(i)
		return true
	}
	return false
}

func NewDoubleRangedSentinel(left, right, min, max int) *DoubleRangedSentinel {
	return &DoubleRangedSentinel{
		RangedSentinel{Sentinel(left), min, max},
		RangedSentinel{Sentinel(right), min, max},
	}
}

type DoubleRangedSentinel struct {
	Left  RangedSentinel
	Right RangedSentinel
}

func NewTripleRangedSentinel(left, middle, right, min, max int) *TripleRangedSentinel {
	return &TripleRangedSentinel{
		RangedSentinel{Sentinel(left), min, max},
		RangedSentinel{Sentinel(middle), min, max},
		RangedSentinel{Sentinel(right), min, max},
	}
}

type TripleRangedSentinel struct {
	Left   RangedSentinel
	Middle RangedSentinel
	Right  RangedSentinel
}
