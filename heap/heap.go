package heap

type childrenOk int

const (
	NoneOK     childrenOk = 0x0
	OnlyLeftOK childrenOk = 0x1
	BothOK     childrenOk = 0x3
)

func children[T any](parentIdx int, slice []T) (left, right int, ok childrenOk) {
	left = ((parentIdx + 1) * 2) - 1
	right = left + 1
	len := len(slice)
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

func biggerChild[T any](parentIdx int, slice []T, less func(a, b *T) bool) (child int, ok bool) {
	left, right, childrenOK := children(parentIdx, slice)
	switch childrenOK {
	case NoneOK:
		ok = false
		return
	case OnlyLeftOK:
		child = left
		ok = true
		return
	case BothOK:
		if less(&slice[left], &slice[right]) {
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

func parent(child int) (parentIdx int) {
	parentIdx = ((child + 1) / 2) - 1
	return
}

func swap[T any](slice []T, i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func heapifyNode[T any](nodeIdx int, slice []T, less func(a, b *T) bool) {
	if biggerChild, ok := biggerChild(nodeIdx, slice, less); ok {
		if less(&slice[nodeIdx], &slice[biggerChild]) {
			// swap parent with bigger child & heapify downwards
			swap(slice, nodeIdx, biggerChild)
			heapifyNode(biggerChild, slice, less)
		}
	}
}

// Heapify arranges a slice's elements according to a Max Heap
func Heapify[T any](slice []T, less func(a, b *T) bool) {
	for i := len(slice) - 1; i >= 0; i-- {
		heapifyNode(i, slice, less)
	}
}

// HeapifyMin arranges a slice's elements according to a Min Heap
func HeapifyMin[T any](slice []T, less func(a, b *T) bool) {
	more := func(a, b *T) bool {
		return !less(a, b)
	}
	for i := len(slice) - 1; i >= 0; i-- {
		heapifyNode(i, slice, more)
	}
}

// HeapSort a slice (ascending order)
func HeapSort[T any](slice []T, less func(a, b *T) bool) []T {
	Heapify(slice, less)
	sortedSlice := slice
	for len(slice) > 0 {
		last := len(slice) - 1
		swap(slice, 0, last)
		slice = slice[0:last]
		heapifyNode(0, slice, less)
	}
	return sortedSlice
}

// HeapSortMin a slice (descending order)
func HeapSortMin[T any](slice []T, less func(a, b *T) bool) []T {
	more := func(a, b *T) bool {
		return !less(a, b)
	}
	return HeapSort(slice, more)
}

// Heap can be either a min or a max heap. See NewMaxHeao & NewMinHeap for creating a heap
type Heap[T any] struct {
	slice []T
	less  func(a, b *T) bool
}

// NewMaxHeap creates a new max heap
func NewMaxHeap[T any](less func(a, b *T) bool, values ...T) *Heap[T] {
	mh := &Heap[T]{
		slice: append([]T{}, values...),
		less:  less,
	}
	mh.Heapify()
	return mh
}

// NewMinHeap creates a new min heap
func NewMinHeap[T any](less func(a, b *T) bool, values ...T) *Heap[T] {
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

func (h *Heap[T]) Heapify() {
	Heapify(h.slice, h.less)
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
	swap(h.slice, 0, last)
	pop = h.slice[last]
	h.slice = h.slice[0:last]
	heapifyNode(0, h.slice, h.less)
	return
}

func (h *Heap[T]) Push(vs ...T) {
	if h == nil {
		return
	}
	for _, v := range vs {
		h.slice = append(h.slice, v)
		childIdx := len(h.slice) - 1
		parentIdx := parent(childIdx)
		for h.less(&h.slice[parentIdx], &h.slice[childIdx]) {
			swap(h.slice, parentIdx, childIdx)
			if parentIdx == 0 {
				break
			}
			childIdx = parentIdx
			parentIdx = parent(childIdx)
		}
	}
	// h.slice
}

// Sort sorts the heap's slice. Be aware that this breaks the heap. It can only be used as heap again after Heapify has been called for the heap.
func (h *Heap[T]) Sort() (sorted []T) {
	if h == nil || len(h.slice) <= 0 {
		return
	}
	slice := h.slice
	for len(h.slice) > 0 {
		last := len(h.slice) - 1
		swap(h.slice, 0, last)
		h.slice = h.slice[0:last]
		heapifyNode(0, h.slice, h.less)
	}
	h.slice = slice
	return h.slice
}
