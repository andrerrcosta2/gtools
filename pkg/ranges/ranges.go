// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package ranges

import "sync"

type RangeImpl struct {
	reverse bool
	current int
	start   int
	end     int
}

func (r *RangeImpl) Next() bool {
	if r.reverse {
		return nextDesc(r)
	}
	return nextAsc(r)
}

func (r *RangeImpl) Prev() bool {
	if r.reverse {
		return prevDesc(r)
	}
	return prevAsc(r)
}

func (r *RangeImpl) Size() int {
	if r.reverse {
		return sizeDesc(r)
	}
	return sizeAsc(r)
}

func (r *RangeImpl) Current() int {
	return r.current
}

func (r *RangeImpl) Reset() {
	reset(r)
}

func (r *RangeImpl) SetRange(start, end int) error {
	if r.reverse {
		err := canSetRangeDesc(r, start, end)
		if err != nil {
			return err
		}
		r.start = start
		r.end = end
		return nil
	}
	err := canSetRangeAsc(r, start, end)
	if err != nil {
		return err
	}
	r.start = start
	r.end = end
	return nil
}

func (r *RangeImpl) Start() int {
	return r.start
}

func (r *RangeImpl) End() int {
	return r.end
}

// Tune this method allows changing the range interval by increasing or decreasing the start and end
func (r *RangeImpl) Tune(start int, end int) {
	r.start += start
	r.end += end
}

func (r *RangeImpl) First() int {
	return r.start
}

func (r *RangeImpl) Last() int {
	return r.end
}

func (r *RangeImpl) SetCurrent(current int) bool {
	if r.reverse {
		if canSetCurrentDesc(r, current) {
			r.current = current
			return true
		}
	} else if canSetCurrentAsc(r, current) {
		r.current = current
		return true
	}
	return false
}

func (r *RangeImpl) IsOnStart() bool {
	return r.current == r.First()
}

func (r *RangeImpl) IsOnEnd() bool {
	return r.current == r.Last()
}

func (r *RangeImpl) RideNext(n int) bool {
	if r.reverse {
		return rideNextDesc(r, n)
	}
	return rideNextAsc(r, n)
}

func (r *RangeImpl) RidePrev(n int) bool {
	if r.reverse {
		return ridePrevDesc(r, n)
	}
	return ridePrevAsc(r, n)
}

func (r *RangeImpl) Snap() Range {
	return &RangeImpl{
		reverse: r.reverse,
		current: r.current,
		start:   r.start,
		end:     r.end,
	}
}

type ConcurrentRangeImpl struct {
	mtx     sync.RWMutex
	reverse bool
	current int
	start   int
	end     int
}

func (r *ConcurrentRangeImpl) Next() bool {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	if r.reverse {
		return nextDesc(r)
	}
	return nextAsc(r)
}

func (r *ConcurrentRangeImpl) Prev() bool {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	if r.reverse {
		return prevDesc(r)
	}
	return prevAsc(r)
}

func (r *ConcurrentRangeImpl) Size() int {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	if r.reverse {
		return sizeDesc(r)
	}
	return sizeAsc(r)
}

func (r *ConcurrentRangeImpl) Current() int {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	return r.current
}

// Reset resets the range
func (r *ConcurrentRangeImpl) Reset() {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	reset(r)
}

func (r *ConcurrentRangeImpl) SetRange(start, end int) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	if r.reverse {
		err := canSetRangeDesc(r, start, end)
		if err != nil {
			return err
		}
		r.start = start
		r.end = end
		return nil
	}
	err := canSetRangeAsc(r, start, end)
	if err != nil {
		return err
	}
	r.start = start
	r.end = end
	return nil
}

func (r *ConcurrentRangeImpl) Start() int {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	return r.start
}

func (r *ConcurrentRangeImpl) End() int {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	return r.end
}

func (r *ConcurrentRangeImpl) Tune(start int, end int) {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	r.start += start
	r.end += end
}

func (r *ConcurrentRangeImpl) First() int {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	return r.start
}

func (r *ConcurrentRangeImpl) Last() int {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	return r.end
}

func (r *ConcurrentRangeImpl) SetCurrent(current int) bool {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	if r.reverse {
		if canSetCurrentDesc(r, current) {
			r.current = current
			return true
		}
	} else if canSetCurrentAsc(r, current) {
		r.current = current
		return true
	}
	return false
}

func (r *ConcurrentRangeImpl) IsOnStart() bool {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	return r.current == r.First()
}

func (r *ConcurrentRangeImpl) IsOnEnd() bool {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	return r.current == r.Last()
}

func (r *ConcurrentRangeImpl) RideNext(n int) bool {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	if r.reverse {
		return rideNextDesc(r, n)
	}
	return rideNextAsc(r, n)
}

func (r *ConcurrentRangeImpl) RidePrev(n int) bool {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	if r.reverse {
		return ridePrevDesc(r, n)
	}
	return ridePrevAsc(r, n)
}

func (r *ConcurrentRangeImpl) Snap() Range {
	r.mtx.RLock()
	defer r.mtx.RUnlock()

	// Create a new mutex for the new instance
	return &ConcurrentRangeImpl{
		reverse: r.reverse,
		mtx:     sync.RWMutex{},
		current: r.current,
		start:   r.start,
		end:     r.end,
	}
}
