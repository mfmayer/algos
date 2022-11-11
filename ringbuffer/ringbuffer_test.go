package ringbuffer

import (
	"fmt"
	"testing"
)

func TestRingBuffer(t *testing.T) {
	rb := NewRingBuffer(WithInitialValues(1, 2, 3, 4, 5, 6, 7, 8, 9, 10), Overwrite(), WithCapacity(5))
	if rb.PeekLastIn() != 10 {
		t.Fail()
	}
	// 1,2,3,4,5 should have been dropped because of limited capacity (5) and overwrite option
	if rb.PeekNextOut() != 6 {
		t.Fail()
	}
	// check if 6..10 is coming in correct order
	for i := 6; i <= 10; i++ {
		v := rb.Pop()
		if v != i {
			t.Fail()
		}
		fmt.Println(v)
	}

	rb = NewRingBuffer(WithCapacity(5))
	// check if len and buffer size is increased correctly
	rb.Push(1, 2, 3, 4, 5)
	if rb.Len() != 5 {
		t.Fail()
	}

	rb.Push(6)
	if rb.Len() != 6 {
		t.Fail()
	}

	rb.Push(7, 8, 9, 10)
	if rb.Len() != 10 {
		t.Fail()
	}

	for i := 1; i <= 5; i++ {
		if rb.Pop() != i {
			t.Fail()
		}
	}

	rb.Push(11, 12, 13, 14, 15)
	for i := 6; i <= 10; i++ {
		if rb.Pop() != i {
			t.Fail()
		}
	}
	for i := 11; i <= 15; i++ {
		if rb.Pop() != i {
			t.Fail()
		}
	}
}
