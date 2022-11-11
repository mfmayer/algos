package linkedlist

import (
	"fmt"
	"testing"
)

func TestLinkedList(t *testing.T) {
	var head AnyElement = NewElement(1, 2, 3, 4, 5)
	next := head
	for i := 1; next != nil; i++ {
		if next.Value().(int) != i {
			t.Fail()
		}
		next = next.Next()
	}

	newHead := NewElement(0)
	newHead.Append(head)
	next = newHead
	count := 0
	for i := 0; next != nil; i++ {
		fmt.Println(next.Value())
		if next.Value().(int) != i {
			t.Fail()
		}
		next = next.Next()
		count++
	}
	if count != 6 {
		t.Fail()
	}
}
