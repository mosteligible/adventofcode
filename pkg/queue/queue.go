package queue

import (
	"adventofcode/y2024/utils"
	"fmt"
)

type Deque[T any] struct {
	queue []T
	size  int
}

func NewQueue[T any]() *Deque[T] {
	return &(Deque[T]{size: 0, queue: []T{}})
}

func (q Deque[T]) String() string {
	retval := ""
	for _, item := range q.queue {
		retval += fmt.Sprintf("%v ", item)
	}
	return fmt.Sprintf("<%s> | size: %d", retval, q.size)
}

func (q *Deque[T]) Size() int {
	return q.size
}

func (q *Deque[T]) Append(item T) {
	q.queue = append(q.queue, item)
	q.size += 1
}

func (q *Deque[T]) Extend(items []T) {
	q.queue = append(q.queue, items...)
	q.size += len(items)
}

func (q *Deque[T]) AppendLeft(item T) {
	q.queue = append([]T{item}, q.queue...)
	q.size += 1
}

func (q *Deque[T]) PopLeft() (T, error) {
	if q.size == 0 {
		return utils.GetZero[T](), fmt.Errorf("no items in deque")
	}
	retval := q.queue[0]
	if q.size == 1 {
		q.queue = []T{}
	} else {
		q.queue = append([]T{}, q.queue[1:]...)
	}
	q.size -= 1
	return retval, nil
}

func (q *Deque[T]) Pop() (T, error) {
	if q.size == 0 {
		return utils.GetZero[T](), fmt.Errorf("no items in deque")
	}
	retval := q.queue[q.size-1]
	q.queue = q.queue[:q.size-1]
	q.size -= 1
	return retval, nil
}

func (q *Deque[T]) PopN(n int) ([]T, error) {
	if n <= 0 {
		return nil, fmt.Errorf("n has to be positive number, got %d", n)
	}
	if n >= q.size {
		retval := q.queue
		q.queue = []T{}
		q.size = 0
		return retval, nil
	}
	retval := q.queue[q.size-n:]
	q.queue = q.queue[:q.size-n]
	q.size -= n
	return retval, nil
}

func (q *Deque[T]) PopNLeft(n int) ([]T, error) {
	if n <= 0 {
		return nil, fmt.Errorf("n has to be positive number, got %d", n)
	}
	if n >= q.size {
		retval := q.queue
		q.queue = []T{}
		q.size = 0
		return retval, nil
	}
	retval := q.queue[:n]
	q.queue = q.queue[n:]
	return retval, nil
}
