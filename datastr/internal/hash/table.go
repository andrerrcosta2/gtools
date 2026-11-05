// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package hash

import (
	"maps"
	"slices"

	"github.com/andrerrcosta2/gtools/core/data/comparator"
	"github.com/andrerrcosta2/gtools/core/domain/function"
)

type Store[H comparable, E any] interface {
	Equals(other Store[H, E], equals ...function.Equals[E]) bool
	HasHash(hash H) bool
	IsEmpty() bool
	Len() int
	foreach(func(H, E) bool)
}

type table[H comparable, E any] interface {
	~map[H]E
}

type Table[H comparable] map[H]struct{}

func (t Table[H]) Clear() {
	clear(t)
}

func (t Table[H]) Clone() Table[H] {
	return maps.Clone(t)
}

func (t Table[H]) Difference(table Table[H]) Table[H] {
	if t.IsEmpty() {
		return make(Table[H])
	}
	if table == nil || table.IsEmpty() {
		return t.Clone()
	}
	res := make(Table[H], t.Len())
	for k := range t {
		if _, exists := table[k]; !exists {
			res[k] = struct{}{}
		}
	}
	return res
}

func (t Table[H]) Equals(other Store[H, struct{}], _ ...function.Equals[struct{}]) bool {
	if t.Len() != other.Len() {
		return false
	}
	allMatch := true
	other.foreach(func(k H, e struct{}) bool {
		if _, exists := t[k]; !exists {
			allMatch = false
			return false
		}
		return true
	})
	return allMatch
}

func (t Table[H]) HasHash(hash H) bool {
	_, ok := t[hash]
	return ok
}

func (t Table[H]) Intersection(table Table[H]) Table[H] {
	if t.IsEmpty() || table == nil || table.IsEmpty() {
		return make(Table[H])
	}
	var m, M = t, table
	if t.Len() > table.Len() {
		m, M = table, t
	}
	res := make(Table[H], len(m))
	for k, v := range m {
		if _, ok := M[k]; ok {
			res[k] = v
		}
	}
	return res
}

func (t Table[H]) IsEmpty() bool {
	return len(t) == 0
}

func (t Table[H]) Len() int {
	return len(t)
}

func (t Table[H]) Merge(table Table[H]) Table[H] {
	res := make(Table[H], t.Len()+table.Len())
	if !t.IsEmpty() {
		res = t.Clone()
	} else {
		return table.Clone()
	}
	if table != nil && !table.IsEmpty() {
		for k := range table {
			res[k] = struct{}{}
		}
	}
	return res
}

func (t Table[H]) foreach(callback func(k H, e struct{}) bool) {
	for k, v := range t {
		if !callback(k, v) {
			return
		}
	}
}

func (t Table[H]) TryAddHash(hash H) bool {
	if !t.HasHash(hash) {
		t[hash] = struct{}{}
		return true
	}
	return false
}

func (t Table[H]) TryRemoveHash(hash H) bool {
	if t.HasHash(hash) {
		delete(t, hash)
		return true
	}
	return false
}

func BucketsOf[H comparable, E any](cmp comparator.EqComparer[E], hint int) Buckets[H, E] {
	return Buckets[H, E]{
		buckets: make(map[H][]E, hint),
		cmp:     cmp,
	}
}

type Buckets[H comparable, E any] struct {
	buckets map[H][]E
	cmp     comparator.EqComparer[E]
}

func (t Buckets[H, E]) Clear() {
	clear(t.buckets)
}

func (t Buckets[H, E]) Clone() Buckets[H, E] {
	return Buckets[H, E]{
		buckets: maps.Clone(t.buckets),
		cmp:     t.cmp,
	}
}

func (t Buckets[H, E]) Difference(other Buckets[H, E]) Buckets[H, E] {
	if t.IsEmpty() {
		return BucketsOf[H, E](t.cmp, 0)
	}
	if other.IsEmpty() {
		return t.Clone()
	}
	res := BucketsOf[H, E](t.cmp, t.Len())
	for k, v := range t.buckets {
		if _, exists := other.buckets[k]; !exists {
			res.buckets[k] = v
		}
	}
	return res
}

func (t Buckets[H, E]) Equals(other Store[H, []E], _ ...function.Equals[E]) bool {
	if t.Len() != other.Len() {
		return false
	}
	allMatch := true
	other.foreach(func(k H, elements []E) bool {
		bucket, exists := t.buckets[k]
		if !exists || !slices.EqualFunc(bucket, elements, t.cmp.Equals) {
			allMatch = false
			return false
		}
		return true
	})
	return allMatch
}

func (t Buckets[H, E]) foreach(callback func(k H, e []E) bool) {
	for k, v := range t.buckets {
		if !callback(k, v) {
			return
		}
	}
}

func (t Buckets[H, E]) GetBucket(hash H) ([]E, bool) {
	bucket, ok := t.buckets[hash]
	return bucket, ok
}

func (t Buckets[H, E]) GetElement(hash H, element E) (E, bool) {
	bucket, ok := t.buckets[hash]
	if !ok {
		return element, false
	}
	for _, e := range bucket {
		if t.cmp.Equals(e, element) {
			return e, true
		}
	}
	return element, false
}

func (t Buckets[H, E]) hasElement(bucket []E, element E) bool {
	for _, e := range bucket {
		if t.cmp.Equals(e, element) {
			return true
		}
	}
	return false
}

func (t Buckets[H, E]) HasHash(hash H) bool {
	_, ok := t.buckets[hash]
	return ok
}

func (t Buckets[H, E]) Intersection(other Buckets[H, E]) Buckets[H, E] {
	if t.IsEmpty() || other.IsEmpty() {
		return BucketsOf[H, E](t.cmp, 0)
	}
	var m, M = t.buckets, other.buckets
	if t.Len() > other.Len() {
		m, M = other.buckets, t.buckets
	}
	res := BucketsOf[H, E](t.cmp, t.Len())

	for hashKey, outerBucket := range m {
		innerBucket, exists := M[hashKey]
		if !exists {
			continue
		}

		if len(outerBucket) > 0 && len(innerBucket) > 0 {
			// Para buckets pequenos busca linear é eficiente
			intersection := make([]E, 0, min(len(outerBucket), len(innerBucket)))
			for _, elem := range outerBucket {
				if t.hasElement(innerBucket, elem) {
					intersection = append(intersection, elem)
				}
			}
			if len(intersection) > 0 {
				res.buckets[hashKey] = intersection
			}
		}
	}

	return res
}

func (t Buckets[H, E]) IsEmpty() bool {
	return len(t.buckets) == 0
}

func (t Buckets[H, E]) Len() int {
	return len(t.buckets)
}

func (t Buckets[H, E]) Merge(other Buckets[H, E]) Buckets[H, E] {
	if other.IsEmpty() {
		return t.Clone()
	}
	if t.IsEmpty() {
		return other.Clone()
	}

	var base, toMerge = t.buckets, other.buckets
	if len(t.buckets) < len(other.buckets) {
		base, toMerge = other.buckets, t.buckets
	}

	res := BucketsOf[H, E](t.cmp, len(base)+len(toMerge))

	// Copia todos os buckets da base para evitar compartilhamento
	for k, v := range base {
		res.buckets[k] = append([]E(nil), v...)
	}

	for hashKey, mergeBucket := range toMerge {
		if len(mergeBucket) == 0 {
			continue
		}

		if existingBucket, exists := res.buckets[hashKey]; exists {
			merged := make([]E, 0, len(existingBucket)+len(mergeBucket))
			merged = append(merged, existingBucket...)
			for _, newElem := range mergeBucket {
				if !t.hasElement(existingBucket, newElem) {
					merged = append(merged, newElem)
				}
			}

			res.buckets[hashKey] = merged
		} else {
			res.buckets[hashKey] = append([]E(nil), mergeBucket...)
		}
	}

	return res
}

func (t Buckets[H, E]) Store(hash H, element E) {
	bucket, ok := t.buckets[hash]
	if !ok {
		t.buckets[hash] = []E{element}
		return
	}
	bucket = append(bucket, element)
}

func (t Buckets[H, E]) TryRemoveElement(hash H, element E, compare function.Equals[E]) bool {
	if bucket, ok := t.GetBucket(hash); ok {
		if len(bucket) == 1 {
			delete(t.buckets, hash)
			return true
		}
		for i, e := range bucket {
			if compare(e, element) {
				bucket = append(bucket[:i], bucket[i+1:]...)
				return true
			}
		}
		return false
	}
	return false
}

func (t Buckets[H, E]) TryRemoveBucket(hash H) bool {
	if t.HasHash(hash) {
		delete(t.buckets, hash)
		return true
	}
	return false
}

type FlatStore[H comparable, V any] map[H]V

func (t FlatStore[H, E]) Clear() {
	clear(t)
}

func (t FlatStore[H, E]) Clone() FlatStore[H, E] {
	return maps.Clone(t)
}

func (t FlatStore[H, E]) Difference(table FlatStore[H, E]) FlatStore[H, E] {
	if t.IsEmpty() {
		return make(FlatStore[H, E])
	}
	if table == nil || table.IsEmpty() {
		return t.Clone()
	}
	res := make(FlatStore[H, E], t.Len())
	for k, v := range t {
		if _, exists := table[k]; !exists {
			res[k] = v
		}
	}
	return res
}

func (t FlatStore[H, E]) Equals(other Store[H, E], equals ...function.Equals[E]) bool {
	if t.Len() != other.Len() {
		return false
	}
	allMatch := true
	if equals == nil || len(equals) == 0 {
		other.foreach(func(h H, _ E) bool {
			allMatch = t.HasHash(h)
			return allMatch
		})
	} else {
		equalsFunc := equals[0]
		other.foreach(func(h H, e E) bool {
			if element, exists := t[h]; exists {
				allMatch = equalsFunc(e, element)
				return allMatch
			}
			allMatch = false
			return false
		})
	}

	return allMatch
}

func (t FlatStore[H, E]) foreach(callback func(k H, e E) bool) {
	for k, v := range t {
		if !callback(k, v) {
			return
		}
	}
}

func (t FlatStore[H, E]) Get(hash H) (E, bool) {
	element, ok := t[hash]
	return element, ok
}

func (t FlatStore[H, E]) HasHash(hash H) bool {
	_, ok := t[hash]
	return ok
}

func (t FlatStore[H, E]) Intersection(table FlatStore[H, E]) FlatStore[H, E] {
	if t.IsEmpty() || table == nil || table.IsEmpty() {
		return make(FlatStore[H, E])
	}
	res := make(FlatStore[H, E], t.Len())
	for k, v := range t {
		if _, exists := table[k]; exists {
			res[k] = v
		}
	}
	return res
}

func (t FlatStore[H, E]) IsEmpty() bool {
	return len(t) == 0
}

func (t FlatStore[H, V]) Keys() []H {
	keys := make([]H, 0, t.Len())
	for key := range t {
		keys = append(keys, key)
	}
	return keys
}

func (t FlatStore[H, E]) Len() int {
	return len(t)
}

func (t FlatStore[H, E]) Merge(table FlatStore[H, E]) FlatStore[H, E] {
	res := make(FlatStore[H, E], t.Len()+table.Len())
	if !t.IsEmpty() {
		res = t.Clone()
	} else {
		return table.Clone()
	}

	if table != nil && !table.IsEmpty() {
		for k, v := range table {
			res[k] = v
		}
	}
	return res
}

func (t FlatStore[H, E]) Store(hash H, element E) {
	t[hash] = element
}

func (t FlatStore[H, E]) TryStore(hash H, element E) bool {
	if !t.HasHash(hash) {
		t[hash] = element
		return true
	}
	return false
}

func (t FlatStore[H, E]) TryRemove(hash H) bool {
	if t.HasHash(hash) {
		delete(t, hash)
		return true
	}
	return false
}

func (t FlatStore[H, E]) Values() []E {
	values := make([]E, 0, t.Len())
	for _, v := range t {
		values = append(values, v)
	}
	return values
}

var _ Store[any, any] = FlatStore[any, any]{}
