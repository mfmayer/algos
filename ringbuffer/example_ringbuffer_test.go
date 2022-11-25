package ringbuffer_test

import (
	"fmt"

	"github.com/mfmayer/algos/ringbuffer"
)

func ExampleNewRingBufferWithFixedCapacity() {
	// Creates ring buffer (FIFO) with fixed capacity of 5 elements
	rFixed := ringbuffer.NewRingBufferWithFixedCapacity[int](5)
	// Push 5 elements into the ring buffer
	rFixed.Push(1, 2, 3, 4, 5)
	rFixed.Push(6) // pushing another element 6 drops element 1, because ring buffer is full
	for rFixed.Len() > 0 {
		fmt.Printf("%v ", rFixed.Pop())
	}
	// Output: 2 3 4 5 6
}

func ExampleNewRingBuffer() {
	// Creates ring buffer (FIFO) with fixed capacity of 5 elements
	rFixed := ringbuffer.NewRingBuffer(1, 2, 3)
	// Push 5 elements into the ring buffer
	rFixed.Push(4, 5)
	for rFixed.Len() > 0 {
		fmt.Printf("%v ", rFixed.Pop())
	}
	rFixed.Push(6, 7, 8, 9, 10) // pushing additional elements into the ringbuffer
	for rFixed.Len() > 0 {
		fmt.Printf("%v ", rFixed.Pop())
	}
	// Output: 1 2 3 4 5 6 7 8 9 10
}
