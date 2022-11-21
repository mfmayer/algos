package ringbuffer

import (
	"fmt"
	"testing"
)

func TestRingBuffer(t *testing.T) {
	r := NewRingBufferWithFixedCapacity(5, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	if r.PeekLastIn() != 10 {
		t.Fail()
	}
	// 1,2,3,4,5 should have been dropped because of limited capacity (5) and overwrite option
	if r.PeekNextOut() != 6 {
		t.Errorf("Got %v, expected 6", r.PeekNextOut())
	}
	// check if 6..10 is coming in correct order
	for i := 6; i <= 10; i++ {
		v := r.Pop()
		if v != i {
			t.Fail()
		}
		fmt.Println(v)
	}

	r = NewRingBuffer[int]()
	// check if len and buffer size is increased correctly
	r.Push(1, 2, 3, 4, 5)
	if r.Len() != 5 {
		t.Fail()
	}

	r.Push(6)
	if r.Len() != 6 {
		t.Fail()
	}

	r.Push(7, 8, 9, 10)
	if r.Len() != 10 {
		t.Fail()
	}

	for i := 1; i <= 5; i++ {
		if r.Pop() != i {
			t.Fail()
		}
	}

	r.Push(11, 12, 13, 14, 15)
	for i := 6; i <= 10; i++ {
		if r.Pop() != i {
			t.Fail()
		}
	}
	for i := 11; i <= 15; i++ {
		if r.Pop() != i {
			t.Fail()
		}
	}
}
