// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package ranges

import "sync"

type openRng struct {
	reverse bool
	current int
	start   int
	end     int
}

func (r *openRng) Next() bool {
	if r.reverse {
		return nextDesc(r)
	}
	return nextAsc(r)
}

func (r *openRng) Prev() bool {
	if r.reverse {
		return prevDesc(r)
	}
	return prevAsc(r)
}

func (r *openRng) Size() int {
	if r.reverse {
		return sizeDesc(r)
	}
	return sizeAsc(r)
}

func (r *openRng) Current() int {
	return r.current
}

func (r *openRng) Reset() {
	reset(r)
}

func (r *openRng) SetRange(start, end int) error {
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

func (r *openRng) Start() int {
	return r.start
}

func (r *openRng) End() int {
	return r.end
}

// Tune this method allows changing the range interval by increasing or decreasing the start and end
func (r *openRng) Tune(start int, end int) {
	r.start += start
	r.end += end
}

func (r *openRng) First() int {
	return r.start
}

func (r *openRng) Last() int {
	return r.end
}

func (r *openRng) SetCurrent(current int) bool {
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

func (r *openRng) IsOnStart() bool {
	return r.current == r.First()
}

func (r *openRng) IsOnEnd() bool {
	return r.current == r.Last()
}

func (r *openRng) RideNext(n int) bool {
	if r.reverse {
		return rideNextDesc(r, n)
	}
	return rideNextAsc(r, n)
}

func (r *openRng) RidePrev(n int) bool {
	if r.reverse {
		return ridePrevDesc(r, n)
	}
	return ridePrevAsc(r, n)
}

func (r *openRng) Snap() Range {
	return &openRng{
		reverse: r.reverse,
		current: r.current,
		start:   r.start,
		end:     r.end,
	}
}

func (r *openRng) DistanceFromStart() int {
	if r.reverse {
		return distanceFromStartDesc(r)
	}
	return distanceFromStartAsc(r)
}

func (r *openRng) DistanceToEnd() int {
	if r.reverse {
		return distanceToEndDesc(r)
	}
	return distanceToEndAsc(r)
}

type concOpenRng struct {
	mtx     sync.RWMutex
	reverse bool
	current int
	start   int
	end     int
}

func (r *concOpenRng) Next() bool {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	if r.reverse {
		return nextDesc(r)
	}
	return nextAsc(r)
}

func (r *concOpenRng) Prev() bool {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	if r.reverse {
		return prevDesc(r)
	}
	return prevAsc(r)
}

func (r *concOpenRng) Size() int {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	if r.reverse {
		return sizeDesc(r)
	}
	return sizeAsc(r)
}

func (r *concOpenRng) Current() int {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	return r.current
}

// Reset resets the range
func (r *concOpenRng) Reset() {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	reset(r)
}

func (r *concOpenRng) SetRange(start, end int) error {
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

func (r *concOpenRng) Start() int {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	return r.start
}

func (r *concOpenRng) End() int {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	return r.end
}

func (r *concOpenRng) Tune(start int, end int) {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	r.start += start
	r.end += end
}

func (r *concOpenRng) First() int {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	return r.start
}

func (r *concOpenRng) Last() int {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	return r.end
}

func (r *concOpenRng) SetCurrent(current int) bool {
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

func (r *concOpenRng) IsOnStart() bool {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	return r.current == r.First()
}

func (r *concOpenRng) IsOnEnd() bool {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	return r.current == r.Last()
}

func (r *concOpenRng) RideNext(n int) bool {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	if r.reverse {
		return rideNextDesc(r, n)
	}
	return rideNextAsc(r, n)
}

func (r *concOpenRng) RidePrev(n int) bool {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	if r.reverse {
		return ridePrevDesc(r, n)
	}
	return ridePrevAsc(r, n)
}

func (r *concOpenRng) Snap() Range {
	r.mtx.RLock()
	defer r.mtx.RUnlock()

	// Create a new mutex for the new instance
	return &concOpenRng{
		reverse: r.reverse,
		mtx:     sync.RWMutex{},
		current: r.current,
		start:   r.start,
		end:     r.end,
	}
}

func (r *concOpenRng) DistanceFromStart() int {
	r.mtx.RLock()
	defer r.mtx.RUnlock()

	if r.reverse {
		return distanceFromStartDesc(r)
	}
	return distanceFromStartAsc(r)
}

func (r *concOpenRng) DistanceToEnd() int {
	r.mtx.RLock()
	defer r.mtx.RUnlock()

	if r.reverse {
		return distanceToEndDesc(r)
	}
	return distanceToEndAsc(r)
}
