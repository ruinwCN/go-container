package stack_array

import (
	"errors"
	"sync"
)

type Stack struct {
	array []interface{}
	size  uint64
	mutex sync.RWMutex
}

func New() *Stack {
	return &Stack{make([]interface{}, 0), 0, sync.RWMutex{}}
}

func (s *Stack) Push(value interface{}) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.array = append(s.array, value)
	s.size++
	return nil
}

func (s *Stack) Pop() (interface{}, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if len(s.array) == 0 {
		return 0, errors.New("stack is empty")
	}
	value := s.array[len(s.array)-1]
	s.array = s.array[0 : len(s.array)-1]
	s.size--
	return value, nil
}

func (s *Stack) Top() (interface{}, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	if len(s.array) == 0 {
		return 0, errors.New("stack is empty")
	}
	value := s.array[len(s.array)-1]
	return value, nil
}

func (s *Stack) Size() uint64 {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.size
}
