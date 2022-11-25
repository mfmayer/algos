package heap

import (
	"fmt"
	"testing"
)

func TestHeap(t *testing.T) {
	// TODO: Implement appropriate test
	h := NewMaxHeap(func(a int, b int) bool { return a < b }, 10, 20, 15, 40, 25, 18, 19)
	fmt.Println(h.slice)
	h.Sort()
	fmt.Println(h.slice)
	h.Heapify()
	fmt.Println(h.slice)
}
