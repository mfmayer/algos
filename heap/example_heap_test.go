package heap_test

import (
	"fmt"

	"github.com/mfmayer/algos/heap"
)

func Example() {
	// Create new heap with providing an appropricate less function and the heap's inital elements
	h := heap.NewMaxHeap(func(a int, b int) bool { return a < b }, 10, 20, 15, 12, 40, 25, 18, 19)

	// Heap then looks like this:
	//        ___40___
	//       /        \
	//     25          19
	//    /  \        /  \
	//  20    10    18    15

	fmt.Println(h.Peek()) // In a max heap the first element is always the biggest
	fmt.Println(h.Sort()) // Sort the heap's slice. This method breaks the heap
	h.Heapify()           // Heapify repairs the heap's order e.g. after calling Sort method
	h.Push(99)
	// Pop all elements until heap is empty
	for h.Len() > 0 {
		e := h.Pop()
		fmt.Printf("%v ", e)
	}
	// Output:
	// 40
	// [10 12 15 18 19 20 25 40]
	// 99 40 25 20 19 18 15 12 10
}
