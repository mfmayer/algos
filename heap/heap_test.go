package heap

import (
	"fmt"
	"testing"
)

func TestHeap(t *testing.T) {
	h := NewMaxHeap(func(a *int, b *int) bool { return *a < *b }, 10, 20, 15, 12, 40, 25, 18, 19)
	h.Push(5, 50)
	h.Push(1)
	h.Push(100)
	root := h.Pop()
	lastRoot := root
	for h.Len() > 0 {
		fmt.Println(root)
		root = h.Pop()
		if root > lastRoot {
			t.Fail()
		}
		lastRoot = root
	}

	sh := NewMinHeap(func(a *string, b *string) bool {
		return (*a)[0] < (*b)[0]
	}, "hallo", "welt")

	for sh.Len() > 0 {
		fmt.Println(sh.Pop())
	}
}
