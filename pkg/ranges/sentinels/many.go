// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sentinels

import "errors"

func Many(sentinels ...int) *NSentinel {
	return &NSentinel{
		sentinels: make([]*Sentinel, len(sentinels)),
	}
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
		s.sentinels[sentinel].Inc(inc)
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
		s.sentinels[sentinel].Dec(dec)
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
