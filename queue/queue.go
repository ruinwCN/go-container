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
}

type Node struct {
	value interface{}
	next  *Node
}

func New() *Queue {
	return &Queue{nil, nil, 0, sync.RWMutex{}}
}

func (s *Queue) Push(value interface{}) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	node := &Node{value: value, next: nil}
	if s.size == 0 {
		s.head = node
		s.tail = node
	} else {
		s.tail.next = node
		s.tail = node
	}
	s.size++
	return nil
}

func (s *Queue) Pop() (interface{}, error) {
	if s.head == nil || s.size == 0 {
		return nil, errors.New("queue is empty. ")
	}
	if s.head == s.tail {
		tmp := &Node{
			value: s.head.value,
			next:  nil,
		}
		s.head = nil
		s.tail = nil
		s.size--
		return tmp, nil
	}
	node := s.head
	s.head = s.head.next
	s.size--
	return node.value, nil
}

func (s *Queue) Top() (interface{}, error) {
	if s.head != nil {
		return s.head.value, nil
	}
	return nil, errors.New("queue is empty. ")
}

func (s *Queue) Size() uint64 {
	return s.size
}
