package queue

import (
	"testing"
)

func TestQueue_Push(t *testing.T) {
	q := New()
	_ = q.Push(12)
	_ = q.Push(13)
	if q.Size() != 2 {
		t.Errorf("queue size is: %d not equal 2. ", q.Size())
	}
	if val, err := q.Top(); err != nil {
		t.Error("queue top error. ")
	} else {
		if val.(int) != 12 {
			t.Errorf("queue top value is: %d not equal 12. ", q.Size())
		}
	}
	if val, err := q.Pop(); err != nil {
		t.Error("queue pop error. ")
	} else {
		if val.(int) != 12 {
			t.Errorf("queue pop value is: %d not equal 12. ", q.Size())
		}
	}
	if q.Size() != 1 {
		t.Errorf("queue size is: %d not equal 1. ", q.Size())
	}
}
