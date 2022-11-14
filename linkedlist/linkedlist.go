package linkedlist

import "github.com/mfmayer/algos"

// AnyElement interface for any element in the linked list
type AnyElement interface {
	algos.AnyElement
	// Next returns the next element
	Next() AnyElement
	// SetNext expliciely sets next element (can destroy a linked list)
	SetNext(AnyElement)
	// Append or insert element (or linked list of elements) next to this one
	Append(AnyElement)
}

// Element Linked List Node
type Element struct {
	*algos.Element
	next AnyElement
}

// NewElement creates new element for each given value, links them and returns the head element
func NewElement(values ...algos.Any) (head AnyElement) {
	if len(values) <= 0 {
		return
	}
	var last *Element
	for _, value := range values {
		element := &Element{
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
func (n *Element) Next() AnyElement {
	return n.next
}

// SetNext expliciely sets next element (can destroy a linked list)
func (n *Element) SetNext(next AnyElement) {
	n.next = next
}

// Append or insert element (or linked list of elements) next to this one
func (n *Element) Append(element AnyElement) {
	oldNext := n.next
	n.next = element
	for element.Next() != nil {
		element = element.Next()
	}
	element.SetNext(oldNext)
}
