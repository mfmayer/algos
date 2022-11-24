package ringbuffer_test

import (
	"fmt"

	"github.com/mfmayer/algos/ringbuffer"
)

func Example() {
	// Creates ring buffer (FIFO) with fixed capacity of 5 elements
	r := ringbuffer.NewRingBufferWithFixedCapacity[int](5)
	// Push 5 elements into the ring buffer
	r.Push(1, 2, 3, 4, 5)
	r.Push(6) // pushing another element 6 drops element 1, because ring buffer is full
	for r.Len() > 0 {
		fmt.Printf("%v ", r.Pop())
	}
	// Output: 2 3 4 5 6
}
