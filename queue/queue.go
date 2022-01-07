package queue

import (
	"errors"
	"sync"
)

type Queue struct {
	head  *Node
	tail  *Node
	size  uint64
	mutex sync.RWMutex
	lock  bool
}

type Node struct {
	value interface{}
	next  *Node
}

func New() *Queue {
	return &Queue{nil, nil, 0, sync.RWMutex{}, false}
}

func (q *Queue) SetLock(lock bool) {
	q.lock = lock
}

func (q *Queue) Push(value interface{}) error {
	if q.lock {
		q.mutex.Lock()
		defer q.mutex.Unlock()
	}
	node := &Node{value: value, next: nil}
	if q.size == 0 {
		q.head = node
		q.tail = node
	} else {
		q.tail.next = node
		q.tail = node
	}
	q.size++
	return nil
}

func (q *Queue) Pop() (interface{}, error) {
	if q.lock {
		q.mutex.Lock()
		defer q.mutex.Unlock()
	}
	if q.head == nil || q.size == 0 {
		return nil, errors.New("queue is empty. ")
	}
	if q.head == q.tail {
		tmp := &Node{
			value: q.head.value,
			next:  nil,
		}
		q.head = nil
		q.tail = nil
		q.size--
		return tmp, nil
	}
	node := q.head
	q.head = q.head.next
	q.size--
	return node.value, nil
}

func (q *Queue) Top() (interface{}, error) {
	if q.head != nil {
		return q.head.value, nil
	}
	return nil, errors.New("queue is empty. ")
}

func (q *Queue) Size() uint64 {
	return q.size
}
