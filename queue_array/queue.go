package queue_array

import (
	"sync"
)

type Queue struct {
	array []interface{}
	size  int
	lock  sync.RWMutex
}

func CreateQueue() *Queue {
	return &Queue{
		size:  0,
		array: make([]interface{}, 0),
		lock:  sync.RWMutex{},
	}
}

func New() *Queue {
	pq := &Queue{}
	pq.array = make([]interface{}, 0)
	pq.size = 0
	return pq
}

func (pq *Queue) Count() int {
	pq.lock.RLock()
	defer pq.lock.RUnlock()
	return pq.size
}

func (pq *Queue) GetHead() interface{} {
	if pq.size == 0 {
		return nil
	}
	return pq.array[0]
}

func (pq *Queue) Push(data interface{}) {
	pq.lock.Lock()
	defer pq.lock.Unlock()
	qData := data
	pq.array = append(pq.array, qData)
	pq.size++
}

func (pq *Queue) Pop() interface{} {
	pq.lock.Lock()
	defer pq.lock.Unlock()
	if pq.size == 0 {
		return nil
	}
	pq.size--
	head := pq.array[0]
	pq.array = pq.array[1:]
	return head
}
