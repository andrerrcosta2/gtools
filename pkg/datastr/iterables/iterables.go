// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package iterables

import (
	"github.com/andrerrcosta2/gtools/pkg/constraints"
	"github.com/andrerrcosta2/gtools/pkg/functions"
	"github.com/andrerrcosta2/gtools/pkg/tuple"
	"math/rand"
)

type Deliverer[T any] interface {
	Deliver(T) []T
}

type Slice[G any] []G

func (s *Slice[G]) Len() int {
	return len(*s)
}

func (s *Slice[G]) Each(fn functions.Consumer[G]) *Slice[G] {
	for _, v := range *s {
		fn(v)
	}
	return s
}

func (s *Slice[G]) EachN(fn functions.BiConsumer[int, G]) *Slice[G] {
	for i, v := range *s {
		fn(i, v)
	}
	return s
}

func (s *Slice[G]) Operation(fn functions.BiConsumer[int, *Slice[G]]) *Slice[G] {
	for i, _ := range *s {
		fn(i, s)
	}
	return s
}

func (s *Slice[G]) Append(v ...G) *Slice[G] {
	*s = append(*s, v...)
	return s
}

func (s *Slice[G]) At(i int) G {
	return (*s)[i]
}

func (s *Slice[G]) Values() []G {
	return *s
}

func (s *Slice[G]) Duplicate() (Slice[G], Slice[G]) {
	work := *s
	newSlice := make(Slice[G], len(work))
	copy(newSlice, work)
	return newSlice, work
}

func (s *Slice[G]) Rand() G {
	return (*s)[rand.Intn(s.Len())]
}

func OfSlice[G any](values ...G) *Slice[G] {
	s := Slice[G](values)
	return &s
}

type Map[G constraints.Ordered, K any] map[G]K

func (t *Map[G, K]) At(k G) K {
	return (*t)[k]
}

func (t *Map[G, K]) Put(k G, v K) *Map[G, K] {
	(*t)[k] = v
	return t
}

func (t *Map[G, K]) Remove(k G) *Map[G, K] {
	delete(*t, k)
	return t
}

func (t *Map[G, K]) Contains(k G) bool {
	_, ok := (*t)[k]
	return ok
}

func (t *Map[G, K]) Len() int {
	return len(*t)
}

func (t *Map[G, K]) Each(fn functions.BiConsumer[G, *K]) *Map[G, K] {
	for k, v := range *t {
		fn(k, &v)
	}
	return t
}

func (t *Map[G, K]) Operation(fn functions.BiConsumer[G, *Map[G, K]]) *Map[G, K] {
	for k, _ := range *t {
		fn(k, t)
	}
	return t
}

func (t *Map[G, K]) Values() []K {
	values := make([]K, 0, len(*t))
	for _, v := range *t {
		values = append(values, v)
	}
	return values
}

func OfMap[G constraints.Ordered, K any](values ...tuple.Pair[G, K]) *Map[G, K] {
	m := Map[G, K]{}
	for _, entry := range values {
		m[entry.First()] = entry.Second()
	}
	return &m
}

type SliceMap[G constraints.Ordered, K any] map[G][]K

func (m *SliceMap[G, K]) Append(k G, v K) *SliceMap[G, K] {
	(*m)[k] = append((*m)[k], v)
	return m
}

func (m *SliceMap[G, K]) Put(k G, v []K) *SliceMap[G, K] {
	(*m)[k] = v
	return m
}

func (m *SliceMap[G, K]) PutIfAbsent(k G, v []K) *SliceMap[G, K] {
	if _, ok := (*m)[k]; !ok {
		(*m)[k] = v
	}
	return m
}

func (m *SliceMap[G, K]) PutOrAppend(k G, v K) *SliceMap[G, K] {
	if m.Contains(k) {
		m.Append(k, v)
	} else {
		m.Put(k, []K{v})
	}
	return m
}

func (m *SliceMap[G, K]) Values() []K {
	result := make([]K, 0, len(*m))
	for _, v := range *m {
		result = append(result, v...)
	}
	return result
}

func (m *SliceMap[G, K]) At(k G) []K {
	return (*m)[k]
}

func (m *SliceMap[G, K]) Remove(k G) *SliceMap[G, K] {
	delete(*m, k)
	return m
}

func (m *SliceMap[G, K]) Contains(k G) bool {
	_, ok := (*m)[k]
	return ok
}

func (m *SliceMap[G, K]) Len() int {
	return len(*m)
}

func (m *SliceMap[G, K]) Each(fn functions.BiConsumer[G, *[]K]) *SliceMap[G, K] {
	for k, v := range *m {
		fn(k, &v)
	}
	return m
}

func (m *SliceMap[G, K]) Operation(fn functions.BiConsumer[G, *SliceMap[G, K]]) *SliceMap[G, K] {
	for k, _ := range *m {
		fn(k, m)
	}
	return m
}
