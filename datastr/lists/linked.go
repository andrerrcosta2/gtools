// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package lists

import (
	"fmt"
	"github.com/andrerrcosta2/gtools/core/data/comparators"
	"github.com/andrerrcosta2/gtools/core/data/str"
	"github.com/andrerrcosta2/gtools/core/domain/constraints/prim"
	"github.com/andrerrcosta2/gtools/datastr/nodes"
	"github.com/andrerrcosta2/gtools/gflux/core/pipes/arrays"
	"sync"
)

type doubleLinked[T any] struct {
	_head *nodes.TypedDoubleLinked[T]
	_tail *nodes.TypedDoubleLinked[T]
	_size int
}

func (l *doubleLinked[T]) head() *nodes.TypedDoubleLinked[T] {
	return l._head
}

func (l *doubleLinked[T]) tail() *nodes.TypedDoubleLinked[T] {
	return l._tail
}

func (l *doubleLinked[T]) setHead(head *nodes.TypedDoubleLinked[T]) {
	l._head = head
}

func (l *doubleLinked[T]) setTail(tail *nodes.TypedDoubleLinked[T]) {
	l._tail = tail
}

func (l *doubleLinked[T]) size() int {
	return l._size
}

// HashLinked creates a linked list with a hash table from the given elements.
func HashLinked[T any, H prim.Hashable](cmp comparators.KeyTyped[T, H], e ...T) str.List[T] {
	if len(e) == 0 {
		return &hashLinked[T, H]{
			doubleLinked: doubleLinked[T]{},
			idx:          make(map[H][]*nodes.TypedDoubleLinked[T]),
			cmp:          cmp,
		}
	}

	idx := make(map[H][]*nodes.TypedDoubleLinked[T])
	head := nodes.NewTypedDoubleLinked[T](nil, nil, e[0])
	tail := head
	size := len(e)
	idx[cmp.Hash(e[0])] = append(idx[cmp.Hash(e[0])], head)

	// Iterate from the second element, linking nodes normally
	for i := 1; i < size; i++ {
		newNode := nodes.NewTypedDoubleLinked[T](tail, nil, e[i])
		tail.SetNext(newNode)
		tail = newNode
		idx[cmp.Hash(e[i])] = append(idx[cmp.Hash(e[i])], newNode)
	}

	return &hashLinked[T, H]{
		doubleLinked: doubleLinked[T]{
			_head: head,
			_tail: tail,
			_size: size,
		},
		idx: idx,
		cmp: cmp,
	}
}

type hashLinked[T any, H prim.Hashable] struct {
	doubleLinked[T]
	idx map[H][]*nodes.TypedDoubleLinked[T]
	cmp comparators.KeyTyped[T, H]
}

func (l *hashLinked[T, H]) Add(t T) (index int) {
	// If the list is empty, initialize it
	if l.head() == nil {
		index = 0
		node := nodes.NewTypedDoubleLinked[T](nil, nil, t)
		l.setHead(node)
		l.setTail(node)
		// Increment size
		l._size++
	} else {
		// Append new node at the tail
		index = l.size()
		node := nodes.NewTypedDoubleLinked[T](l.tail(), nil, t)
		l.tail().SetNext(node)
		l.setTail(node)
		// Increment size
		l._size++
	}
	// Add to hash map
	l.addKey(l.hash(t), l.tail())

	return index
}

func (l *hashLinked[T, H]) addKey(key H, node *nodes.TypedDoubleLinked[T]) {
	if l.idx[key] == nil {
		l.idx[key] = []*nodes.TypedDoubleLinked[T]{node}
	} else {
		l.idx[key] = append(l.idx[key], node)
	}
}

func (l *hashLinked[T, H]) Contains(t T) bool {
	_, ok := l.idx[l.hash(t)]
	return ok
}

// Get returns the element at the given i (0-based).
func (l *hashLinked[T, H]) Get(i int) (get T, exists bool) {
	node, ok := findLinkedNodeByIndex[T](l, i)
	if !ok {
		return
	}
	return node.Value(), true
}

func (l *hashLinked[T, H]) HardClear() {
	// Traverse the list from _head to _tail and break all links
	node := l.head()
	for node != nil {
		next := node.Next()
		node.SetNext(nil)
		node.SetPrev(nil)
		node = next
	}
	// Reset _head and _tail ptrs to nil
	l.setHead(nil)
	l.setTail(nil)
	l.idx = make(map[H][]*nodes.TypedDoubleLinked[T])
	l._size = 0
}

func (l *hashLinked[T, H]) hash(t T) H {
	return l.cmp.Hash(t)
}

func (l *hashLinked[T, H]) IndexOf(t T) int {
	return findLinkedNodeByValue[T](l, t, l.cmp.Equals)
}

// Insert a t at the given i, moving the rest of the elements to the right
func (l *hashLinked[T, H]) Insert(i int, t T) bool {
	if l.head() == nil { // empty list
		if i != 0 {
			return false // Can't insert at non-zero index in an empty list
		}
		node := nodes.NewTypedDoubleLinked[T](nil, nil, t)
		l.setHead(node)
		l.setTail(node) // Ensure tail is updated
		l._size++
		l.addKey(l.hash(t), node)
		return true
	}

	if i == 0 { // Insert at the head
		head := l.head()
		node := nodes.NewTypedDoubleLinked[T](nil, head, t)
		head.SetPrev(node)
		l.setHead(node)
		l._size++
		l.addKey(l.hash(t), node)

		// If the list had only one element before, update tail
		if l.tail() == head {
			l.setTail(node.Next()) // Ensure tail stays valid
		}

		return true
	}

	// Find insertion point
	node, ok := findLinkedNodeByIndex[T](l, i)
	if !ok {
		return false
	}

	// Insert in the middle
	nn := nodes.NewTypedDoubleLinked[T](node.Prev(), node, t)
	node.Prev().SetNext(nn)
	node.SetPrev(nn)
	l._size++
	l.addKey(l.hash(t), nn)

	return true
}

func (l *hashLinked[T, H]) IsEmpty() bool {
	return l.size() == 0
}

// Remove all matching ts of the given t from the list
func (l *hashLinked[T, H]) Remove(t T) bool {
	key := l.hash(t)
	indices, exists := l.idx[key]
	if !exists {
		return false
	}
	for _, node := range indices {
		removeLinkedNode[T](l, node)
		if !l.removeKey(key, node) {
			// Panic if the node is not found in the hash map
			panic(fmt.Sprintf("fatal error: corrupt linked list. unable to remove node from hash map. "+
				"no node found with the same address as the removed node (value: %v)", node.Value()))
		}
		l._size--
	}
	return true
}

func (l *hashLinked[T, H]) RemoveAt(i int) (removed T, ok bool) {
	node, exists := findLinkedNodeByIndex[T](l, i)
	if !exists {
		return
	}
	removeLinkedNode[T](l, node)
	if !l.removeKey(l.hash(node.Value()), node) {
		// Panic if the node is not found in the hash map
		panic(fmt.Sprintf("fatal error: corrupt linked list. Unable to remove node from hash map. "+
			"no node found with the same address as the removed node (value: %v)", node.Value()))
	}
	l._size--
	return node.Value(), true
}

func (l *hashLinked[T, H]) removeKey(key H, n *nodes.TypedDoubleLinked[T]) bool {
	nodes, ok := l.idx[key]
	if !ok {
		return false
	}

	// Handle multiple nodes with the same hash (collision)
	for i, node := range nodes {
		if node == n {
			if len(l.idx[key]) == 1 {
				delete(l.idx, key)
			} else {
				l.idx[key] = arrays.RemoveByIndex(l.idx[key], i)
			}
			return true
		}
	}

	return false
}

// Set the t at the given i, replacing the old t
func (l *hashLinked[T, H]) Set(i int, t T) bool {
	node, ok := findLinkedNodeByIndex[T](l, i)
	if !ok {
		return false
	}
	l.removeKey(l.hash(node.Value()), node)
	node.SetValue(t)
	l.addKey(l.hash(t), node)
	return true
}

func (l *hashLinked[T, H]) Size() int {
	return l.size()
}

func (l *hashLinked[T, H]) SoftClear() {
	l.setHead(nil)
	l.setTail(nil)
	l.idx = make(map[H][]*nodes.TypedDoubleLinked[T])
	l._size = 0
}

func (l *hashLinked[T, H]) ToSlice() []T {
	var s []T
	for node := l.head(); node != nil; node = node.Next() {
		s = append(s, node.Value())
	}
	return s
}

var _ LinkedList[int] = (*hashLinked[int, string])(nil)

func ConcHashLinked[T any, H prim.Hashable](cmp comparators.KeyTyped[T, H], e ...T) str.List[T] {
	if len(e) == 0 {
		return &concHashLinked[T, H]{
			doubleLinked: doubleLinked[T]{},
			idx:          make(map[H][]*nodes.TypedDoubleLinked[T]),
			cmp:          cmp,
		}
	}

	idx := make(map[H][]*nodes.TypedDoubleLinked[T])
	head := nodes.NewTypedDoubleLinked[T](nil, nil, e[0])
	tail := head
	size := len(e)
	idx[cmp.Hash(e[0])] = append(idx[cmp.Hash(e[0])], head)

	// Iterate from the second element, linking nodes normally
	for i := 1; i < size; i++ {
		newNode := nodes.NewTypedDoubleLinked[T](tail, nil, e[i])
		tail.SetNext(newNode)
		tail = newNode
		idx[cmp.Hash(e[i])] = append(idx[cmp.Hash(e[i])], newNode)
	}

	return &concHashLinked[T, H]{
		doubleLinked: doubleLinked[T]{
			_head: head,
			_tail: tail,
			_size: size,
		},
		idx: idx,
		cmp: cmp,
	}
}

type concHashLinked[T any, H prim.Hashable] struct {
	doubleLinked[T]
	idx map[H][]*nodes.TypedDoubleLinked[T]
	cmp comparators.KeyTyped[T, H]
	mtx sync.RWMutex
}

func (l *concHashLinked[T, H]) Add(t T) (index int) {
	l.mtx.Lock()
	defer l.mtx.Unlock()
	// If the list is empty, initialize it
	if l.head() == nil {
		index = 0
		node := nodes.NewTypedDoubleLinked[T](nil, nil, t)
		l.setHead(node)
		l.setTail(node)
		// Increment size
		l._size++
	} else {
		// Append new node at the tail
		index = l.size()
		node := nodes.NewTypedDoubleLinked[T](l.tail(), nil, t)
		l.tail().SetNext(node)
		l.setTail(node)
		// Increment size
		l._size++
	}
	// Add key to hash map
	l.addKey(l.hash(t), l.tail())
	return index
}

func (l *concHashLinked[T, H]) addKey(key H, node *nodes.TypedDoubleLinked[T]) {
	if l.idx[key] == nil {
		l.idx[key] = []*nodes.TypedDoubleLinked[T]{node}
	} else {
		l.idx[key] = append(l.idx[key], node)
	}
}

func (l *concHashLinked[T, H]) Contains(t T) bool {
	l.mtx.RLock()
	defer l.mtx.RUnlock()
	_, ok := l.idx[l.hash(t)]
	return ok
}

func (l *concHashLinked[T, H]) Get(i int) (t T, ok bool) {
	l.mtx.RLock()
	defer l.mtx.RUnlock()
	node, ok := findLinkedNodeByIndex[T](l, i)
	if !ok {
		return
	}
	return node.Value(), true
}

func (l *concHashLinked[T, H]) HardClear() {
	l.mtx.Lock()
	defer l.mtx.Unlock()
	// Traverse the list from _head to _tail and break all links
	node := l.head()
	for node != nil {
		next := node.Next()
		node.SetNext(nil)
		node.SetPrev(nil)
		node = next
	}
	// Reset _head and _tail ptrs to nil
	l.setHead(nil)
	l.setTail(nil)
	l.idx = make(map[H][]*nodes.TypedDoubleLinked[T])
	l._size = 0
}

func (l *concHashLinked[T, H]) hash(t T) H {
	return l.cmp.Hash(t)
}

func (l *concHashLinked[T, H]) IndexOf(t T) int {
	l.mtx.RLock()
	defer l.mtx.RUnlock()
	return findLinkedNodeByValue[T](l, t, l.cmp.Equals)
}

func (l *concHashLinked[T, H]) Insert(i int, t T) bool {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	if l.head() == nil { // empty list
		if i != 0 {
			return false // Can't insert at non-zero index in an empty list
		}
		node := nodes.NewTypedDoubleLinked[T](nil, nil, t)
		l.setHead(node)
		l.setTail(node) // Ensure tail is updated
		l._size++
		l.addKey(l.hash(t), node)
		return true
	}

	if i == 0 { // Insert at the head
		head := l.head()
		node := nodes.NewTypedDoubleLinked[T](nil, head, t)
		head.SetPrev(node)
		l.setHead(node)
		l._size++
		l.addKey(l.hash(t), node)

		// If the list had only one element before, update tail
		if l.tail() == head {
			l.setTail(node.Next()) // Ensure tail stays valid
		}

		return true
	}

	// Find insertion point
	node, ok := findLinkedNodeByIndex[T](l, i)
	if !ok {
		return false
	}

	// Insert in the middle
	nn := nodes.NewTypedDoubleLinked[T](node.Prev(), node, t)
	node.Prev().SetNext(nn)
	node.SetPrev(nn)
	l._size++
	l.addKey(l.hash(t), nn)

	return true
}

func (l *concHashLinked[T, H]) IsEmpty() bool {
	return l.size() == 0
}

func (l *concHashLinked[T, H]) Remove(t T) bool {
	l.mtx.Lock()
	defer l.mtx.Unlock()
	key := l.hash(t)
	indices, exists := l.idx[key]
	if !exists {
		return false
	}
	for _, node := range indices {
		removeLinkedNode[T](l, node)
		if !l.removeKey(key, node) {
			// Panic if the node is not found in the hash map
			panic(fmt.Sprintf("fatal error: corrupt linked list. unable to remove node from hash map. "+
				"no node found with the same address as the removed node (value: %v)", node.Value()))
		}
		l._size--
	}
	return true
}

func (l *concHashLinked[T, H]) RemoveAt(i int) (removed T, ok bool) {
	node, exists := findLinkedNodeByIndex[T](l, i)
	if !exists {
		return
	}
	removeLinkedNode[T](l, node)
	if !l.removeKey(l.hash(node.Value()), node) {
		// Panic if the node is not found in the hash map
		panic(fmt.Sprintf("fatal error: corrupt linked list. Unable to remove node from hash map. "+
			"no node found with the same address as the removed node (value: %v)", node.Value()))
	}
	l._size--
	return node.Value(), true
}

func (l *concHashLinked[T, H]) removeKey(key H, n *nodes.TypedDoubleLinked[T]) bool {
	nodes, ok := l.idx[key]
	if !ok {
		return false
	}

	// Handle multiple nodes with the same hash (collision)
	for i, node := range nodes {
		if node == n {
			if len(l.idx[key]) == 1 {
				delete(l.idx, key)
			} else {
				l.idx[key] = arrays.RemoveByIndex(l.idx[key], i)
			}
			return true
		}
	}

	return false
}

func (l *concHashLinked[T, H]) Set(i int, t T) bool {
	l.mtx.Lock()
	defer l.mtx.Unlock()
	node, ok := findLinkedNodeByIndex[T](l, i)
	if !ok {
		return false
	}
	l.removeKey(l.hash(node.Value()), node)
	node.SetValue(t)
	l.addKey(l.hash(t), node)
	return true
}

func (l *concHashLinked[T, H]) Size() int {
	l.mtx.RLock()
	defer l.mtx.RUnlock()
	return l.size()
}

func (l *concHashLinked[T, H]) SoftClear() {
	l.mtx.Lock()
	defer l.mtx.Unlock()
	l.setHead(nil)
	l.setTail(nil)
	l.idx = make(map[H][]*nodes.TypedDoubleLinked[T])
	l._size = 0
}

func (l *concHashLinked[T, H]) ToSlice() []T {
	l.mtx.RLock()
	defer l.mtx.RUnlock()
	var s []T
	for node := l.head(); node != nil; node = node.Next() {
		s = append(s, node.Value())
	}
	return s
}

var _ LinkedList[int] = (*concHashLinked[int, int])(nil)
