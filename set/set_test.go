package set

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	s1 := New()
	s1.Insert("string1")
	if s1.Has("string1") != true {
		t.Errorf("set error type 1")
	}
	if s1.Has("string") != false {
		t.Errorf("set error type 1-")
	}
	s1.Insert("string2")
	s1.Insert("string3")
	s1.Remove("string1")
	if s1.Has("string1") != false {
		t.Errorf("set error type 2")
	}
	if s1.Len() != 2 {
		t.Errorf("set error type 3")
	}

	f := func(value interface{}) {
		fmt.Println(value)
	}
	s1.Do(f)
}
