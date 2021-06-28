package stack

import (
	"errors"
)

type BaseStackInterface interface {
	Push(interface{}) error
	Pop() (interface{}, error)
	Top() (interface{}, error)
	Size() uint64
}

type BaseStack struct {
}

func (s *BaseStack) Push(interface{}) error {
	return errors.New("[Push] not found")
}

func (s *BaseStack) Pop() (interface{}, error) {
	return nil, errors.New("[Pop] not found")
}

func (s *BaseStack) Top() (interface{}, error) {
	return nil, errors.New("[Top] not found")
}

func (s *BaseStack) Size() uint64 {
	return 0
}
