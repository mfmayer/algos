package linkedlist

import "github.com/mfmayer/algos"

// Element Linked List Node
type Element[T any] struct {
	*algos.Element[T]
	next *Element[T]
}

// NewElement creates new element for each given value, links them and returns the head element
func NewElement[T any](values ...T) (head *Element[T]) {
	if len(values) <= 0 {
		return
	}
	var last *Element[T]
	for _, value := range values {
		element := &Element[T]{
			Element: algos.NewElement(value),
		}
		if last == nil {
			head = element
		} else {
			last.SetNext(element)
		}
		last = element
	}
	return head
}

// Next returns the next element
func (n *Element[T]) Next() *Element[T] {
	return n.next
}

// SetNext expliciely sets next element (can destroy a linked list)
func (n *Element[T]) SetNext(next *Element[T]) {
	n.next = next
}

// Insert element (or linked list of elements) next to this one
func (n *Element[T]) InsertAfter(element *Element[T]) {
	oldNext := n.next
	n.next = element
	for element.Next() != nil {
		element = element.Next()
	}
	element.SetNext(oldNext)
}

// Insert element (or linked list of elements) before to this one
func (n *Element[T]) InsertBefore(element *Element[T]) {
	for element.Next() != nil {
		element = element.Next()
	}
	element.SetNext(n)
}
