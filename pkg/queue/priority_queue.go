package queue

import (
	"adventofcode/y2024/utils"
	"fmt"
	"sync"
)

type PriorityQueueItem[T any] struct {
	Priority int
	Item     T
}

type PriorityQueue[T any] struct {
	lock    *sync.RWMutex
	size    int
	pqItems []PriorityQueueItem[T]
}

func NewPriorityQueue[T any]() PriorityQueue[T] {
	return PriorityQueue[T]{
		pqItems: []PriorityQueueItem[T]{}, lock: &sync.RWMutex{}, size: 0,
	}
}

func (pq *PriorityQueue[T]) Size() int {
	pq.lock.RLock()
	defer pq.lock.RUnlock()
	return pq.size
}

func (pq *PriorityQueue[T]) Push(val PriorityQueueItem[T]) {
	pq.lock.Lock()
	defer pq.lock.Unlock()

	for index := 0; index < pq.size-1; index++ {
		if pq.pqItems[index].Priority <= val.Priority && val.Priority <= pq.pqItems[index+1].Priority {
			pq.pqItems = append(pq.pqItems[:index+1], append([]PriorityQueueItem[T]{val}, pq.pqItems[index+1:]...)...)
			pq.size++
			return
		}
	}
	pq.pqItems = append(pq.pqItems, val)
	pq.size++
}

func (pq *PriorityQueue[T]) Pop() (PriorityQueueItem[T], error) {
	pq.lock.Lock()
	defer pq.lock.Unlock()
	if pq.size == 0 {
		return utils.GetZero[PriorityQueueItem[T]](), fmt.Errorf("priority queue is empty")
	}
	retval := pq.pqItems[0]
	pq.size--
	pq.pqItems = pq.pqItems[1:]

	return retval, nil
}
