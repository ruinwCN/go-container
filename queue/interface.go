package queue

type BaseQueueInterface interface {
	Push(interface{}) error
	Pop() (interface{}, error)
	Top() (interface{}, error)
	//Tail() (interface{}, error)
	Size() uint64
}
