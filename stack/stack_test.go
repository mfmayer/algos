package stack

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	stack := NewStack(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	for i := 0; i < 5; i++ {
		fmt.Println(stack.Peek(0))
		if stack.Pop() != 10-i {
			t.Fail()
		}
	}
	stack.Push(20)
	stack.Push(30)
	if stack.Peek(0) != 30 || stack.Peek(1) != 20 {
		t.Fail()
	}
	for i := 0; i < 2; i++ {
		peek := stack.Peek(0)
		fmt.Println(peek)
		if stack.Pop() != 30-i*10 {
			t.Fail()
		}
	}
	for i := 0; stack.Len() > 0; i++ {
		fmt.Println(stack.Peek(0))
		if stack.Pop() != 5-i {
			t.Fail()
		}
	}
}
