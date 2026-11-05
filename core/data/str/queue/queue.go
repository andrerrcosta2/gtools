// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package queue

type Readable[T any] interface {
	Front() (T, bool)
	Back() (T, bool)
	Len() int
}

type Writable[T any] interface {
	PushFront(T)
	PushBack(T)

	PopFront() (T, bool)
	PopBack() (T, bool)

	Clear()
}

type Mutable[T any] interface {
	Readable[T]
	Writable[T]
}

type ReadablePriority[T any] interface {
	Peek() (T, bool)
	Len() int
}

type WritablePriority[T any] interface {
	Push(T)
	Pop() (T, bool)
	Clear()
}

// MutablePriority is a readable and writable priority queue
type MutablePriority[T any] interface {
	ReadablePriority[T]
	WritablePriority[T]
}
