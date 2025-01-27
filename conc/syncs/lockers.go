// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package syncs

import (
	"github.com/andrerrcosta2/gtools/core/domain/constraints/prim"
	"sync"
)

type FGLocker[K prim.Hashable] struct {
	rw   map[K]*sync.RWMutex
	mx   map[K]*sync.Mutex
	lock sync.Mutex
}

// Get or create RWMutex for read/write locking
func (fg *FGLocker[K]) getRWLock(key K) *sync.RWMutex {
	fg.lock.Lock()
	defer fg.lock.Unlock()
	if mtx, exists := fg.rw[key]; exists {
		return mtx
	}
	mtx := &sync.RWMutex{}
	fg.rw[key] = mtx
	return mtx
}

// Get or create Mutex for exclusive locking
func (fg *FGLocker[K]) getMutex(key K) *sync.Mutex {
	fg.lock.Lock()
	defer fg.lock.Unlock()
	if mtx, exists := fg.mx[key]; exists {
		return mtx
	}
	mtx := &sync.Mutex{}
	fg.mx[key] = mtx
	return mtx
}

// LockExc exclusively using Mutex
func (fg *FGLocker[K]) LockExc(key K) {
	fg.getMutex(key).Lock()
}

// UnlockExc exclusive Mutex
func (fg *FGLocker[K]) UnlockExc(key K) {
	fg.getMutex(key).Unlock()
}

// RLock for read using RWMutex
func (fg *FGLocker[K]) RLock(key K) {
	fg.getRWLock(key).RLock()
}

// Unlock after read
func (fg *FGLocker[K]) RUnlock(key K) {
	fg.getRWLock(key).RUnlock()
}

// Lock for write using RWMutex
func (fg *FGLocker[K]) Lock(key K) {
	fg.getRWLock(key).Lock()
}

// Unlock after write
func (fg *FGLocker[K]) Unlock(key K) {
	fg.getRWLock(key).Unlock()
}
