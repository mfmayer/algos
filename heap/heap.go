package heap

import "github.com/mfmayer/algos"

// Heap can be either a min or a max heap. See NewMaxHeao & NewMinHeap for creating a heap
type Heap[T algos.Any] struct {
	slice []T
	less  func(a, b *T) bool
}

// NewMaxHeap creates a new max heap
func NewMaxHeap[T algos.Any](less func(a, b *T) bool, values ...T) *Heap[T] {
	mh := &Heap[T]{
		slice: append([]T{}, values...),
		less:  less,
	}
	mh.Heapify()
	return mh
}

// NewMinHeap creates a new min heap
func NewMinHeap[T algos.Any](less func(a, b *T) bool, values ...T) *Heap[T] {
	comp := func(a, b *T) bool {
		return !less(a, b)
	}
	mh := &Heap[T]{
		slice: append([]T{}, values...),
		less:  comp,
	}
	mh.Heapify()
	return mh
}

type childrenOk int

const (
	NoneOK     childrenOk = 0x0
	OnlyLeftOK childrenOk = 0x1
	BothOK     childrenOk = 0x3
)

func (h *Heap[T]) children(parent int) (left, right int, ok childrenOk) {
	left = ((parent + 1) * 2) - 1
	right = left + 1
	len := len(h.slice)
	if right < len {
		// if right < len then left is also < len
		ok = BothOK
		return
	}
	if left < len {
		ok = OnlyLeftOK
		return
	}
	return
}

func (h *Heap[T]) parent(child int) (parent int) {
	parent = ((child + 1) / 2) - 1
	return
}

func (h *Heap[T]) biggerChild(parent int) (child int, ok bool) {
	left, right, childrenOK := h.children(parent)
	switch childrenOK {
	case NoneOK:
		ok = false
		return
	case OnlyLeftOK:
		child = left
		ok = true
		return
	case BothOK:
		if h.less(&h.slice[left], &h.slice[right]) {
			child = right
			ok = true
			return
		} else {
			child = left
			ok = true
			return
		}
	}
	return
}

func (h *Heap[T]) swap(i, j int) {
	h.slice[i], h.slice[j] = h.slice[j], h.slice[i]
}

func (h *Heap[T]) heapifyNode(i int) {
	if biggerChild, ok := h.biggerChild(i); ok {
		if h.less(&h.slice[i], &h.slice[biggerChild]) {
			// swap parent with bigger child & heapify downwards
			h.swap(i, biggerChild)
			h.heapifyNode(biggerChild)
		}
	}
}

func (h *Heap[T]) Heapify() {
	for i := len(h.slice) - 1; i >= 0; i-- {
		h.heapifyNode(i)
	}
}

func (h *Heap[T]) Len() int {
	if h == nil {
		return 0
	}
	return len(h.slice)
}

func (h *Heap[T]) Pop() (pop T) {
	if h == nil || len(h.slice) <= 0 {
		return
	}
	last := len(h.slice) - 1
	h.swap(0, last)
	pop = h.slice[last]
	h.slice = h.slice[0:last]
	h.heapifyNode(0)
	return
}

func (h *Heap[T]) Push(vs ...T) {
	if h == nil {
		return
	}
	for _, v := range vs {
		h.slice = append(h.slice, v)
		child := len(h.slice) - 1
		parent := h.parent(child)
		for h.less(&h.slice[parent], &h.slice[child]) {
			h.swap(parent, child)
			if parent == 0 {
				break
			}
			child = parent
			parent = h.parent(child)
		}
	}
	// h.slice
}
