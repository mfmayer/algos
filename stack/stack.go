package stack

// Stack with any any elements
type Stack[T any] []T

// NewStack creates and initializes new stack with given values
func NewStack[T any](values ...T) *Stack[T] {
	s := append(Stack[T]{}, values...)
	return &s
}

// Len returns the length of the stack
func (s *Stack[T]) Len() int {
	return len(*s)
}

// Push any value on the stack
func (s *Stack[T]) Push(any ...T) {
	*s = append(*s, any...)
}

// Pop value on top of the stack
func (s *Stack[T]) Pop() T {
	var empty T
	len := s.Len()
	if len <= 0 {
		return empty
	}
	lastIdx := len - 1
	ret := (*s)[lastIdx]
	(*s)[lastIdx] = empty
	*s = (*s)[:lastIdx]
	return ret
}

// Peek on/into the stack (i=0 element on top of the stack) without popping the element from the stack.
func (s *Stack[T]) Peek(i int) any {
	len := s.Len()
	if i < 0 || i > len-1 {
		return nil
	}
	ret := (*s)[(len-1)-i]
	return ret
}
