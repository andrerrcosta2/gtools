// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package conc

import (
	"fmt"
	"sync"
	"time"
)

type Goroutine struct {
	id      int
	value   int
	payload int
	state   string
	regs    []string
}

func NewGoroutine(id int, value int, payload int) *Goroutine {
	return &Goroutine{
		id:      id,
		value:   value,
		payload: payload,
		state:   "new",
	}
}

func (g *Goroutine) SetState(state string) {
	g.state = state
}

func (g *Goroutine) Register(duration time.Duration) {
	g.regs = append(g.regs, fmt.Sprintf("Goroutine %d registered after %v\n", g.id, duration))
}

func (g *Goroutine) Work() {
	time.Sleep(time.Duration(g.payload) * time.Millisecond)
}

type Registers struct {
	mtx      sync.Mutex
	working  map[int]*Goroutine
	finished map[int]*Goroutine
	start    time.Time
	max      int
}

func (r *Registers) Worker(g *Goroutine) {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	if _, inWorking := r.working[g.id]; inWorking {
		return
	}

	if _, inFinished := r.finished[g.id]; inFinished {
		return
	}

	//fmt.Printf("\nGoroutine %d is working\n", g.id)
	g.SetState("working")
	g.Register(time.Since(r.start))
	r.working[g.id] = g
}

func (r *Registers) Done(g *Goroutine) {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	if _, inWorking := r.working[g.id]; !inWorking {
		return
	}

	delete(r.working, g.id)
	//fmt.Printf("\nGoroutine %d is done\n", g.id)
	g.SetState("finished")
	r.finished[g.id] = g
}

func (r *Registers) Max() int {
	return r.max
}

func NewRegisters(max int) *Registers {
	return &Registers{
		working:  make(map[int]*Goroutine),
		finished: make(map[int]*Goroutine),
		start:    time.Now(),
		max:      max,
	}
}

func (r *Registers) Count() int {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	return len(r.working) + len(r.finished)
}
