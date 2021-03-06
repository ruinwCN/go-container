package stack

type BaseStackInterface interface {
	Push(interface{}) error
	Pop() (interface{}, error)
	Top() (interface{}, error)
	Size() uint64
}
