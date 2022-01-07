package queue_array

import (
	"errors"
	"sync"
)

type Queue struct {
	array []interface{}
	size  uint64
	mutex sync.RWMutex
	lock  bool
}

func New() *Queue {
	q := &Queue{}
	q.array = make([]interface{}, 0)
	q.size = 0
	q.lock = false
	return q
}

func (q *Queue) SetLock(lock bool) {
	q.lock = lock
}

func (q *Queue) Size() uint64 {
	if q.lock {
		q.mutex.Lock()
		defer q.mutex.Unlock()
	}
	return q.size
}

func (q *Queue) Top() (interface{}, error) {
	if q.size == 0 {
		return nil, errors.New("queue is empty. ")
	}
	return q.array[0], nil
}

func (q *Queue) Push(data interface{}) error {
	if q.lock {
		q.mutex.Lock()
		defer q.mutex.Unlock()
	}
	qData := data
	q.array = append(q.array, qData)
	q.size++
	return nil
}

func (q *Queue) Pop() (interface{}, error) {
	if q.lock {
		q.mutex.Lock()
		defer q.mutex.Unlock()
	}
	if q.size == 0 {
		return nil, errors.New("queue is empty. ")
	}
	q.size--
	head := q.array[0]
	q.array = q.array[1:]
	return head, nil
}
