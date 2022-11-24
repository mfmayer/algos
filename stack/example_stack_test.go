package stack_test

import (
	"fmt"

	"github.com/mfmayer/algos/stack"
)

func ExampleStack() {
	// Create stack and initially push some elements -> 10 on top
	s := stack.NewStack(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	s.Push(20)             // pushes 20 on top of the stack
	fmt.Println(s.Pop())   // pop top element (ouputs 20)
	fmt.Println(s.Peek(0)) // peek on top element without removing it from stack (outputs 10)
	// Output:
	// 20
	// 10
}
