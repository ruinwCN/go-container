package stack

import (
	"errors"
	"sync"
)

type Stack struct {
	top   *Node
	size  uint64
	mutex sync.RWMutex
	lock  bool
}

type Node struct {
	value interface{}
	next  *Node
}

func New() *Stack {
	return &Stack{nil, 0, sync.RWMutex{}, false}
}

func (s *Stack) SetLock(lock bool) {
	s.lock = lock
}

func (s *Stack) Push(value interface{}) error {
	if s.lock {
		s.mutex.Lock()
		defer s.mutex.Unlock()
	}
	node := &Node{value: value, next: nil}
	node.next = s.top
	s.top = node
	s.size++
	return nil
}

func (s *Stack) Pop() (interface{}, error) {
	if s.lock {
		s.mutex.Lock()
		defer s.mutex.Unlock()
	}
	if s.top == nil || s.size <= 0 {
		return nil, errors.New("stack is empty")
	}
	node := s.top
	s.top = node.next
	s.size--
	return node.value, nil
}

func (s *Stack) Top() (interface{}, error) {
	if s.lock {
		s.mutex.Lock()
		defer s.mutex.Unlock()
	}
	if s.top != nil {
		return s.top.value, nil
	}
	return nil, errors.New("top is nil")
}

func (s *Stack) Size() uint64 {
	if s.lock {
		s.mutex.Lock()
		defer s.mutex.Unlock()
	}
	return s.size
}
