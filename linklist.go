package linklist

import "errors"

var (
	ErrQueueEmpty = errors.New("queue empty")
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
	if q.head != nil {
		if q.head == q.tail {
			q.tail = nil
		}
		oldHead := q.head
		q.head = oldHead.next
		q.length--
		return oldHead.value, nil
	}

	var empty T
	return empty, ErrQueueEmpty
}

// Enqueue implements Queue.
func (q *LinkedListQueue[T]) Enqueue(value T) error {
	newNode := &node[T]{
		value: value,
		next:  nil,
	}

	if q.tail != nil {
		q.tail.next = newNode
	} else if q.head == nil {
		q.head = newNode
	}

	q.tail = newNode
	q.length++

	return nil
}

// Len implements Queue.
func (q *LinkedListQueue[T]) Size() int {
	return q.length
}

func NewLinkedListQueue[T any]() Queue[T] {
	return &LinkedListQueue[T]{}
}
