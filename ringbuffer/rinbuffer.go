package ringbuffer

import "github.com/mfmayer/algos"

// RinfBugger (FIFO) with any elements
type RingBuffer struct {
	buffer          []algos.Any
	overwrite       bool
	nextIn, nextOut int
}

type settings struct {
	cap       int
	overwrite bool
	initWith  []algos.Any
}

type option func(*settings)

// WithCapacity option to initialize the rinfbuffer with a specific capacity
func WithCapacity(cap int) option {
	return func(s *settings) {
		s.cap = cap + 1 // the ringbuffer can handle one element less than the underlying slice capacity, therefore +1
	}
}

// WithInitialValues option to initialize the ringubffer with given values
func WithInitialValues(values ...algos.Any) option {
	return func(s *settings) {
		s.initWith = append(s.initWith, values...)
	}
}

// Overwrite option to overwrite values instead of increasing the ring buffers size when more values are inserted
// than availaable capacity
func Overwrite() option {
	return func(s *settings) {
		s.overwrite = true
	}
}

// NewRingBuffer creates and initializes a new ring buffer with given options
func NewRingBuffer(options ...option) *RingBuffer {
	s := settings{
		cap: 8,
	}
	for _, opt := range options {
		opt(&s)
	}
	ringBuffer := &RingBuffer{
		overwrite: s.overwrite,
		buffer:    make([]algos.Any, s.cap),
		nextIn:    0,
		nextOut:   0,
	}
	ringBuffer.Push(s.initWith...)
	return ringBuffer
}

// normalizedNextInOut normlizes nextIn and nextOut indices so that nextOut is always < nextIn
func (r *RingBuffer) normalizedNextInOut() (in, out int) {
	in, out = r.nextIn, r.nextOut
	for out > in {
		in = in + len(r.buffer)
	}
	return
}

// Len returns the number of values in the ring buffer
func (r *RingBuffer) Len() int {
	in, out := r.normalizedNextInOut()
	return in - out
}

// PeekNextOut to see what value would pop next (without popping it)
func (r *RingBuffer) PeekNextOut() (value algos.Any) {
	if r.Len() <= 0 {
		return
	}
	value = r.buffer[r.nextOut]
	return
}

// PeekLastIn to see what value was lastly pushed into the ring buffer
func (r *RingBuffer) PeekLastIn() (value algos.Any) {
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
func (r *RingBuffer) Pop() (value algos.Any) {
	if r.Len() <= 0 {
		return
	}
	value = r.buffer[r.nextOut]
	r.buffer[r.nextOut] = nil
	r.nextOut = (r.nextOut + 1) % len(r.buffer)
	return
}

// Push values into the ring buffer. If overwrite is enabled and no capacity left oldest calues will be overwritten first.
// Otherwise the underlying buffer size will be dynamically increased.
func (r *RingBuffer) Push(values ...algos.Any) {
	for _, value := range values {
		if r.Len() >= len(r.buffer)-1 {
			// element must be dropped or buffer resized
			if r.overwrite {
				// drop element
				r.nextOut = (r.nextOut + 1) % len(r.buffer)
			} else {
				// resize underlying buffer
				buffer := make([]algos.Any, 2*len(r.buffer))
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
