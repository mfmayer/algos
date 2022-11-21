package ringbuffer

// RinfBugger (FIFO) with any elements
type RingBuffer[T any] struct {
	buffer          []T
	overwrite       bool
	nextIn, nextOut int
}

// NewRingBuffer creates and initializes a new ring buffer. The capacity of underlying slice increases dynamically
func NewRingBuffer[T any](values ...T) *RingBuffer[T] {
	capacity := len(values)
	if capacity < 1 {
		capacity = 1
	}
	ringBuffer := &RingBuffer[T]{
		overwrite: false,
		buffer:    make([]T, len(values)+1),
		nextIn:    0,
		nextOut:   0,
	}
	ringBuffer.Push(values...)
	return ringBuffer
}

// NewRingBufferWithFixedCapacity creates and initialized a new ring buffer with a fixed capacaity. Pushing more values into the buffer than it can store results in dropping first inserted values first.
func NewRingBufferWithFixedCapacity[T any](capacity int, values ...T) *RingBuffer[T] {
	if capacity < 1 {
		capacity = 1
	}
	ringBuffer := &RingBuffer[T]{
		overwrite: true,
		buffer:    make([]T, capacity+1),
		nextIn:    0,
		nextOut:   0,
	}
	ringBuffer.Push(values...)
	return ringBuffer
}

// normalizedNextInOut normlizes nextIn and nextOut indices so that nextOut is always < nextIn
func (r *RingBuffer[T]) normalizedNextInOut() (in, out int) {
	in, out = r.nextIn, r.nextOut
	for out > in {
		in = in + len(r.buffer)
	}
	return
}

// Len returns the number of values in the ring buffer
func (r *RingBuffer[T]) Len() int {
	in, out := r.normalizedNextInOut()
	return in - out
}

// PeekNextOut to see what value would pop next (without popping it)
func (r *RingBuffer[T]) PeekNextOut() (value T) {
	if r.Len() <= 0 {
		return
	}
	value = r.buffer[r.nextOut]
	return
}

// PeekLastIn to see what value was lastly pushed into the ring buffer
func (r *RingBuffer[T]) PeekLastIn() (value T) {
	if r.Len() <= 0 {
		return
	}
	lastIn := r.nextIn - 1
	if lastIn < 0 {
		lastIn = len(r.buffer) - 1
	}
	value = r.buffer[lastIn]
	return
}

// Pop next value from ring buffer
func (r *RingBuffer[T]) Pop() (value T) {
	if r.Len() <= 0 {
		return
	}
	value = r.buffer[r.nextOut]
	var empty T
	r.buffer[r.nextOut] = empty
	r.nextOut = (r.nextOut + 1) % len(r.buffer)
	return
}

// Push values into the ring buffer. If overwrite is enabled and no capacity left oldest calues will be overwritten first.
// Otherwise the underlying buffer size will be dynamically increased.
func (r *RingBuffer[T]) Push(values ...T) {
	for _, value := range values {
		if r.Len() >= len(r.buffer)-1 {
			// element must be dropped or buffer resized
			if r.overwrite {
				// drop element
				r.nextOut = (r.nextOut + 1) % len(r.buffer)
			} else {
				// resize underlying buffer
				buffer := make([]T, 2*len(r.buffer))
				i := 0
				for r.Len() > 0 {
					buffer[i] = r.Pop()
					i++
				}
				r.nextIn = i
				r.nextOut = 0
				r.buffer = buffer
			}
		}
		r.buffer[r.nextIn] = value
		r.nextIn = (r.nextIn + 1) % len(r.buffer)
	}
}
