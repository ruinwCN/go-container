package stack

import (
	"errors"
	"sync"
)

type Stack struct {
	top   *Node
	size  uint64
	mutex sync.RWMutex
}

type Node struct {
	value interface{}
	next  *Node
}

func New() *Stack {
	return &Stack{nil, 0, sync.RWMutex{}}
}

func (s *Stack) Push(value interface{}) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	node := &Node{value: value, next: nil}
	node.next = s.top
	s.top = node
	s.size++
	return nil
}

func (s *Stack) Pop() (interface{}, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.top == nil || s.size <= 0 {
		return nil, errors.New("stack is empty")
	}
	node := s.top
	s.top = node.next
	s.size--
	return node.value, nil
}

func (s *Stack) Top() (interface{}, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	if s.top != nil {
		return s.top.value, nil
	}
	return nil, errors.New("top is nil")
}

func (s *Stack) Size() uint64 {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.size
}
