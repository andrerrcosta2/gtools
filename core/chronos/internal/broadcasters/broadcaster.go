// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package broadcasters

import "sync"

func New[T any]() *Map[T] {
	return &Map[T]{
		subs: make(map[uint64]chan T),
	}
}

type Map[T any] struct {
	mu   sync.RWMutex
	subs map[uint64]chan T
	next uint64
}

func (b *Map[T]) Subscribe() (uint64, <-chan T) {
	ch := make(chan T, 1)
	b.mu.Lock()
	id := b.next
	b.next++

	b.subs[id] = ch
	b.mu.Unlock()
	return id, ch
}
func (b *Map[T]) Publish(value T) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	for _, ch := range b.subs {
		select {
		case ch <- value:
		default:
		}
	}
}

func (b *Map[T]) Unsubscribe(id uint64) {
	b.mu.Lock()
	defer b.mu.Unlock()
	ch, ok := b.subs[id]
	if !ok {
		return
	}
	close(ch)
	delete(b.subs, id)
}
