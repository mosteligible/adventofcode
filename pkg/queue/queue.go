package queue

import (
	"adventofcode/y2024/utils"
	"fmt"
	"sync"
)

type Queue[T any] []T

type Deque[T any] struct {
	queue  []T
	size   int
	rwlock *sync.RWMutex
}

func NewDeque[T any]() *Deque[T] {
	return &(Deque[T]{size: 0, queue: []T{}, rwlock: &sync.RWMutex{}})
}

func (q Deque[T]) String() string {
	retval := ""
	for _, item := range q.queue {
		retval += fmt.Sprintf("%v ", item)
	}
	return fmt.Sprintf("<%s> | size: %d", retval, q.size)
}

func (q *Deque[T]) Size() int {
	q.rwlock.RLock()
	defer q.rwlock.RUnlock()
	return q.size
}

func (q *Deque[T]) Append(item T) {
	q.rwlock.Lock()
	defer q.rwlock.Unlock()
	q.queue = append(q.queue, item)
	q.size += 1
}

func (q *Deque[T]) Extend(items []T) {
	q.rwlock.Lock()
	defer q.rwlock.Unlock()
	q.queue = append(q.queue, items...)
	q.size += len(items)
}

func (q *Deque[T]) AppendLeft(item T) {
	q.rwlock.Lock()
	defer q.rwlock.Unlock()
	q.queue = append([]T{item}, q.queue...)
	q.size += 1
}

func (q *Deque[T]) PopLeft() (T, error) {
	q.rwlock.Lock()
	defer q.rwlock.Unlock()
	if q.size <= 0 {
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
	q.rwlock.Lock()
	defer q.rwlock.Unlock()
	if q.size <= 0 {
		return utils.GetZero[T](), fmt.Errorf("no items in deque")
	}
	retval := q.queue[q.size-1]
	q.queue = q.queue[:q.size-1]
	q.size -= 1
	return retval, nil
}

func (q *Deque[T]) PopN(n int) ([]T, error) {
	q.rwlock.Lock()
	defer q.rwlock.Unlock()
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
	q.rwlock.Lock()
	defer q.rwlock.Unlock()
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

func (q *Deque[T]) GetN(index, num int) []T {
	retval := []T{}
	right := index + num
	if right >= q.Size() {
		right = q.Size()
	}
	for ; index < right; index++ {
		retval = append(retval, q.queue[index])
	}

	return retval
}
