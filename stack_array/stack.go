package stack_array

import (
	"errors"
	"sync"
)

type Stack struct {
	array []interface{}
	size  uint64
	mutex sync.RWMutex
	lock  bool
}

func New() *Stack {
	return &Stack{make([]interface{}, 0), 0, sync.RWMutex{}, false}
}

func (s *Stack) SetLock(lock bool) {
	s.lock = lock
}

func (s *Stack) Push(value interface{}) error {
	if s.lock {
		s.mutex.Lock()
		defer s.mutex.Unlock()
	}
	s.array = append(s.array, value)
	s.size++
	return nil
}

func (s *Stack) Pop() (interface{}, error) {
	if s.lock {
		s.mutex.Lock()
		defer s.mutex.Unlock()
	}
	if len(s.array) == 0 {
		return 0, errors.New("stack is empty")
	}
	value := s.array[len(s.array)-1]
	s.array = s.array[0 : len(s.array)-1]
	s.size--
	return value, nil
}

func (s *Stack) Top() (interface{}, error) {
	if s.lock {
		s.mutex.Lock()
		defer s.mutex.Unlock()
	}
	if len(s.array) == 0 {
		return 0, errors.New("stack is empty")
	}
	value := s.array[len(s.array)-1]
	return value, nil
}

func (s *Stack) Size() uint64 {
	if s.lock {
		s.mutex.Lock()
		defer s.mutex.Unlock()
	}
	return s.size
}
