// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package subscr

import (
	"sync"
	"sync/atomic"

	"github.com/andrerrcosta2/gtools/gflux/core/domain/lifecycle/gate"
	"github.com/andrerrcosta2/gtools/gflux/obs/internal/observers"
)

type (
	// Subscription represents cooperative cancellation.
	//
	// For synchronous cold Observables (e.g. Of, Range),
	// emissions may occur before Sub returns.
	// In such cases, Unsub does not interrupt in-flight emissions.
	//
	// For asynchronous Observables, Unsub actively stops
	// future emissions and releases resources.
	Subscription interface {
		Unsub()
		Closed() bool
	}

	Unsubscribable interface {
		RemoveSub(Subscription)
	}
)

type UnsubscribableFunc func(Subscription)

func (f UnsubscribableFunc) RemoveSub(s Subscription) {
	f(s)
}

type boundSub struct {
	inner Subscription
	gate  gate.Closeable
}

func Bind(inner Subscription, gate gate.Closeable) Subscription {
	if inner == nil {
		return &gateOnlySub{gate: gate}
	}
	return &boundSub{
		inner: inner,
		gate:  gate,
	}
}

func (s *boundSub) Unsub() {
	if !s.gate.Closed() {
		s.gate.Close()
		s.inner.Unsub()
	}
}

func (s *boundSub) Closed() bool {
	return s.gate.Closed()
}

type gateOnlySub struct {
	gate gate.Closeable
}

func (s *gateOnlySub) Unsub() {
	s.gate.Close()
}

func (s *gateOnlySub) Closed() bool {
	return s.gate.Closed()
}

type Slot[T any] struct {
	gate     gate.Closeable
	observer *observers.Observer[T]
}

func (s *Slot[T]) Close() {
	s.gate.Close()
}

func (s *Slot[T]) Closed() bool {
	return s.gate.Closed()
}

func (s *Slot[T]) Observer() *observers.Observer[T] {
	return s.observer
}

type SlotSub[T any] struct {
	slot *Slot[T]
}

func NewSlotSub[T any](obs *observers.Observer[T], gate gate.Closeable) *SlotSub[T] {
	return &SlotSub[T]{
		slot: &Slot[T]{
			gate:     gate,
			observer: obs,
		},
	}
}

func (s *SlotSub[T]) Closed() bool {
	return s.slot.Closed()
}

func (s *SlotSub[T]) Slot() *Slot[T] {
	return s.slot
}

func (s *SlotSub[T]) Unsub() {
	s.slot.Close()
}

func NewDiSub[T any](unsub Unsubscribable) *DiSubscr[T] {
	return &DiSubscr[T]{
		unsub: unsub,
	}
}

type DiSubscr[T any] struct {
	unsub  Unsubscribable
	closed atomic.Bool
}

func (s *DiSubscr[T]) Closed() bool {
	return s.closed.Load()
}

func (s *DiSubscr[T]) Close() {
	s.closed.Store(true)
}

func (s *DiSubscr[T]) Unsub() {
	if s.closed.CompareAndSwap(false, true) && s.unsub != nil {
		s.unsub.RemoveSub(s)
	}
}

func NewNoop() *Noop {
	return &Noop{}
}

type Noop struct {
}

func (n Noop) Unsub() {}

func (n Noop) Closed() bool { return false }

func Func(unsub func()) *FuncSub {
	return &FuncSub{
		closed: atomic.Bool{},
		unsub:  unsub,
	}
}

type FuncSub struct {
	closed atomic.Bool
	unsub  func()
}

func (f *FuncSub) Unsub() {
	if f.closed.CompareAndSwap(false, true) && f.unsub != nil {
		f.unsub()
	}
}

func (f *FuncSub) Closed() bool {
	return f.closed.Load()
}

func (f *FuncSub) Close() {
	f.closed.Store(true)
}

func NewSwitching() *Switching {
	return &Switching{
		mu:     sync.Mutex{},
		sub:    nil,
		closed: false,
	}
}

type Switching struct {
	mu     sync.Mutex
	sub    Subscription
	closed bool
}

func (c *Switching) Switch(next Subscription) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.closed {
		if next != nil {
			next.Unsub()
		}
		return
	}

	prev := c.sub
	c.sub = next

	if prev != nil {
		prev.Unsub()
	}
}

func (c *Switching) Unsub() {
	c.mu.Lock()
	if c.closed {
		c.mu.Unlock()
		return
	}
	c.closed = true
	sub := c.sub
	c.sub = nil
	c.mu.Unlock()

	if sub != nil {
		sub.Unsub()
	}
}

func (c *Switching) Close() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.closed = true
}

func (c *Switching) Closed() bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.closed
}

func SetOf(u ...Subscription) *Set {
	return &Set{
		unsubs: u,
	}
}

type Set struct {
	mu     sync.RWMutex
	unsubs []Subscription
	once   sync.Once
	closed bool
}

func (s *Set) Add(sub Subscription) {
	if sub == nil {
		return
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.closed { // unsubscribe late additions
		sub.Unsub()
		return
	}
	s.unsubs = append(s.unsubs, sub)
}

func (s *Set) Unsub() {
	s.once.Do(func() {
		s.mu.Lock()
		defer s.mu.Unlock()
		s.closed = true
		for _, sub := range s.unsubs {
			sub.Unsub()
		}
	})
}

func (s *Set) Closed() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.closed
}
