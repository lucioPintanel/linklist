package linklist

import (
	"container/list"
	"errors"
)

var (
	ErrQueueEmpty = errors.New("queue empty")
)

type customQueue[T any] struct {
	queue *list.List
}

func NewList[T any]() *customQueue[T] {
	return &customQueue[T]{
		queue: list.New(),
	}
}

func (c *customQueue[T]) Enqueue(value T) {
	c.queue.PushBack(value)
}

func (q *customQueue[T]) Size() int {
	return q.queue.Len()
}

func (q *customQueue[T]) Empty() bool {
	return q.Size() == 0
}
