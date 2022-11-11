package stack

import "github.com/mfmayer/algos"

// Stack with any any elements
type Stack []algos.Any

// NewStack creates and initializes new stack with given values
func NewStack(values ...algos.Any) Stack {
	s := append(Stack{}, values...)
	return s
}

// Len returns the length of the stack
func (s *Stack) Len() int {
	return len(*s)
}

// Push any value on the stack
func (s *Stack) Push(any ...algos.Any) {
	*s = append(*s, any...)
}

// Pop value on top of the stack
func (s *Stack) Pop() algos.Any {
	len := s.Len()
	if len <= 0 {
		return nil
	}
	lastIdx := len - 1
	ret := (*s)[lastIdx]
	(*s)[lastIdx] = nil
	*s = (*s)[:lastIdx]
	return ret
}

// Peek on/into the stack (i=0 element on top of the stack) without popping the element from the stack.
func (s *Stack) Peek(i int) algos.Any {
	len := s.Len()
	if i < 0 || i > len-1 {
		return nil
	}
	ret := (*s)[(len-1)-i]
	return ret
}
