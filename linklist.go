package linklist

import "errors"

var (
	ErrQueueEmpty = errors.New("queue empty")
	ErrQueueFull  = errors.New("queue full")
)

type Queue[T any] interface {
	Enqueue(T) error
	Dequeue() (T, error)
	Size() int
}

type node[T any] struct {
	value T
	next  *node[T]
}

type LinkedListQueue[T any] struct {
	head   *node[T]
	tail   *node[T]
	length int
}

// Dequeue implements Queue.
func (q *LinkedListQueue[T]) Dequeue() (T, error) {
	panic("unimplemented")
}

// Enqueue implements Queue.
func (q *LinkedListQueue[T]) Enqueue(T) error {
	panic("unimplemented")
}

// Len implements Queue.
func (q *LinkedListQueue[T]) Size() int {
	return q.length
}

func NewLinkedListQueue[T any]() Queue[T] {
	return &LinkedListQueue[T]{}
}
