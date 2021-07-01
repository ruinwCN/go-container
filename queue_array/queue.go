package queue_array

import (
	"errors"
	"sync"
)

type Queue struct {
	array []interface{}
	size  uint64
	lock  sync.RWMutex
}

func New() *Queue {
	pq := &Queue{}
	pq.array = make([]interface{}, 0)
	pq.size = 0
	return pq
}

func (pq *Queue) Size() uint64 {
	pq.lock.RLock()
	defer pq.lock.RUnlock()
	return pq.size
}

func (pq *Queue) Top() (interface{}, error) {
	if pq.size == 0 {
		return nil, errors.New("queue is empty. ")
	}
	return pq.array[0], nil
}

func (pq *Queue) Push(data interface{}) error {
	pq.lock.Lock()
	defer pq.lock.Unlock()
	qData := data
	pq.array = append(pq.array, qData)
	pq.size++
	return nil
}

func (pq *Queue) Pop() (interface{}, error) {
	pq.lock.Lock()
	defer pq.lock.Unlock()
	if pq.size == 0 {
		return nil, errors.New("queue is empty. ")
	}
	pq.size--
	head := pq.array[0]
	pq.array = pq.array[1:]
	return head, nil
}
