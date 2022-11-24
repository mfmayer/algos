package linkedlist

import (
	"testing"
)

func TestLinkedList(t *testing.T) {
	l1 := NewElement(1, 2, 3, 4)
	l5 := NewElement(5)
	l6 := NewElement(6, 7, 8, 9)
	l5.InsertBefore(l1)
	l5.InsertAfter(l6)

	next := l1
	for i := 1; next != nil; i++ {
		if next.Data != i {
			t.Fail()
		}
		next = next.Next()
	}
}
