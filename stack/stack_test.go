package stack

import (
	"sync"
	"testing"
)

func TestStack_Test(t *testing.T) {
	s := New()
	if s.Size() != 0 {
		t.Errorf("empty size must equal 0.")
	}
	err := s.Push(12)
	if err != nil {
		t.Error(err)
	}
	if s.Size() != 1 {
		t.Errorf("stack size error.")
	}
	if value, err := s.Top(); err != nil {
		t.Error(err)
	} else if value.(int) != 12 {
		t.Errorf("stack top value error.")
	}
	if value, err := s.Pop(); err != nil {
		t.Error(err)
	} else if value.(int) != 12 {
		t.Errorf("stack pop value error.")
	}
	if s.Size() != 0 {
		t.Errorf("empty size must equal 0.")
	}
}

func TestStack_Mut(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(150)
	s := New()
	go func() {
		for i := 0; i < 100; i++ {
			_ = s.Push(i)
			wg.Done()
		}
	}()
	go func() {
		for i := 0; i < 50; i++ {
			_, _ = s.Pop()
			wg.Done()
		}
	}()
	wg.Wait()
	if s.Size() != 50 {
		t.Error("stack length error")
	}
}
