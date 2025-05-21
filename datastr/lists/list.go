// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package lists

import (
	"github.com/andrerrcosta2/gtools/core/data/str"
	"github.com/andrerrcosta2/gtools/datastr/nodes"
)

type LinkedList[T any] interface {
	str.List[T]
	head() *nodes.TypedDoubleLinked[T]
	setHead(head *nodes.TypedDoubleLinked[T])
	setTail(tail *nodes.TypedDoubleLinked[T])
	size() int
	tail() *nodes.TypedDoubleLinked[T]
}

func findLinkedNodeByIndex[T any](l LinkedList[T], index int) (node *nodes.TypedDoubleLinked[T], exists bool) {
	if index < 0 || index >= l.size() {
		return
	}

	var current *nodes.TypedDoubleLinked[T]
	if index < l.size()/2 {
		current = l.head()
		for i := 0; i < index; i++ {
			current = current.Next()
		}
	} else {
		current = l.tail()
		for i := l.size() - 1; i > index; i-- {
			current = current.Prev()
		}
	}
	return current, true
}

func findLinkedNodeByValue[T any](l LinkedList[T], value T, eq func(a, b T) bool) int {
	if l.head() == nil {
		return -1
	}

	current := l.head()
	for i := 0; i < l.size(); i++ {
		if eq(current.Value(), value) {
			return i
		}
		current = current.Next()
	}

	return -1
}

func removeLinkedNode[T any](l LinkedList[T], node *nodes.TypedDoubleLinked[T]) {
	if node.Prev() != nil {
		node.Prev().SetNext(node.Next())
	} else {
		l.setHead(node.Next())
	}

	if node.Next() != nil {
		node.Next().SetPrev(node.Prev())
	} else {
		l.setTail(node.Prev())
	}
}

// This is weird
//func ToSlice[T any](l str.List[T]) []T {
//	if ll, ok := l.(LinkedList[T]); ok {
//		var s []T
//		for node := ll.head(); node != nil; node = node.Next() {
//			s = append(s, node.Value())
//		}
//		return s
//	}
//	return nil
//}
