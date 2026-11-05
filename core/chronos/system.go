package chronos

import (
	"time"

	"github.com/andrerrcosta2/gtools/core/domain/contracts/chrono"
)

func SystemClock() chrono.Clock {
	return systemClock{}
}

type systemClock struct{}

func (c systemClock) PeriodsUntil(until chrono.Time, period chrono.Duration) int {
	interval := c.Until(until)
	return int(interval / period)
}

func (c systemClock) PeriodsSince(since chrono.Time, period chrono.Duration) int {
	interval := c.Since(since)
	return int(interval / period)
}

func (c systemClock) Now() chrono.Time {
	return chrono.FromStdTime(time.Now())
}

func (c systemClock) Since(t chrono.Time) chrono.Duration {
	return chrono.FromStdTime(time.Now()).Sub(t)
}

func (c systemClock) Until(t chrono.Time) chrono.Duration {
	return t.Sub(chrono.FromStdTime(time.Now()))
}

func (systemClock) After(d chrono.Duration) <-chan chrono.Time {
	ch := make(chan chrono.Time, 1)

	go func() {
		t := <-time.After(d.ToStd())
		ch <- chrono.FromStdTime(t)
	}()

	return ch
}

type systemTicker struct {
	ticker *time.Ticker
	out    chan chrono.Time
}

func (systemClock) NewTicker(d chrono.Duration) chrono.Ticker {
	return &systemTicker{
		ticker: time.NewTicker(d.ToStd()),
	}
}

func (t *systemTicker) C() <-chan chrono.Time {
	go func() {
		for v := range t.ticker.C {
			t.out <- chrono.FromStdTime(v)
		}
	}()
	return t.out
}

func (t *systemTicker) Stop() {
	t.ticker.Stop()
}
