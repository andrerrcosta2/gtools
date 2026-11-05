// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package chronos

import (
	"sync"
	"time"

	"github.com/andrerrcosta2/gtools/core/chronos/internal/broadcasters"
	"github.com/andrerrcosta2/gtools/core/domain/contracts/chrono"
)

func VirtualClock() chrono.AdvanceableClock {
	return &virtual{
		now:    chrono.Time(time.Unix(0, 0)),
		ticker: broadcasters.New[chrono.Time](),
	}
}

type virtual struct {
	mu     sync.RWMutex
	now    chrono.Time
	ticker *broadcasters.Map[chrono.Time]
}

// PeriodsSince returns how manu periods have elapsed since the given time
func (c *virtual) PeriodsSince(since chrono.Time, period chrono.Duration) int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	interval := c.Since(since)
	return int(interval / period)
}

func (c *virtual) PeriodsUntil(until chrono.Time, period chrono.Duration) int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	interval := c.Until(until)
	return int(interval / period)
}

func (c *virtual) Now() chrono.Time {
	return c.now
}

func (c *virtual) Since(t chrono.Time) chrono.Duration {
	return c.now.Sub(t)
}

func (c *virtual) Until(t chrono.Time) chrono.Duration {
	return t.Sub(c.now)
}

func (c *virtual) Advance(d chrono.Duration) {
	c.now = c.now.Add(d)
	c.tick(c.now)
}

func (c *virtual) AdvanceTo(t chrono.Time) {
	c.now = t
}

func (c *virtual) tick(time chrono.Time) {
	c.ticker.Publish(time)
}

func (c *virtual) NewTicker(period chrono.Duration) chrono.Ticker {
	id, ch := c.ticker.Subscribe()
	return &virtualTicker{
		period:    period,
		startedAt: c.now,
		id:        id,
		clock:     c,
		ch:        ch,
	}
}

type virtualTicker struct {
	period    chrono.Duration
	startedAt chrono.Time
	out       chan chrono.Time
	id        uint64
	ch        <-chan chrono.Time
	clock     *virtual
}

func (t *virtualTicker) C() <-chan chrono.Time {
	go func() {
		next := t.startedAt.NextAfter(t.clock.now, t.period)
		for now := range t.ch {
			for !next.After(now) {
				t.out <- next
				next = next.Add(t.period)
			}
		}
	}()

	return t.out
}

func (t *virtualTicker) Stop() {
	t.clock.ticker.Unsubscribe(t.id)
}
