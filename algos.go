package algos

import "fmt"

func Empty[T any]() (e T) {
	return
}

// Element that contains any data
type Element[T any] struct {
	Data T
}

// NewElemenet creates new element with any value
func NewElement[T any](data T) *Element[T] {
	return &Element[T]{
		Data: data,
	}
}

// String returns element's string representation
func (e *Element[T]) String() string {
	return fmt.Sprintf("%v", e.GetData())
}

// GetData returns element's value
func (e *Element[T]) GetData() any {
	return e.Data
}

func (e *Element[T]) SetData(data any) (ok bool) {
	e.Data, ok = data.(T)
	return
}
