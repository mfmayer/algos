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

// Element in a linked list
type Element struct {
	*algos.Element
	next AnyElement
}

// NewElement creates new element for each given value, links them and returns the head element
func NewElement(values ...algos.Any) (head *Element) {
	if len(values) <= 0 {
		return
	}
	var lastElement *Element
	for _, value := range values {
		element := &Element{
			Element: algos.NewElement(value),
		}
		if lastElement == nil {
			head = element
		} else {
			lastElement.SetNext(element)
		}
		lastElement = element
	}
	return head
}

// Next returns the next element
func (e *Element) Next() AnyElement {
	return e.next
}

// SetNext expliciely sets next element (can destroy a linked list)
func (e *Element) SetNext(next AnyElement) {
	e.next = next
}

// Append or insert element (or linked list of elements) next to this one
func (e *Element) Append(element AnyElement) {
	oldNext := e.next
	e.next = element
	for element.Next() != nil {
		element = element.Next()
	}
	element.SetNext(oldNext)
}
