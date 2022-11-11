package algos

// Any alias for an empty interface
type Any interface{}

// AnyElement interface
type AnyElement interface {
	// Value returns the element's value
	Value() Any
	// SetValue sets the element's value
	SetValue(Any)
}

// Element that contains any value
type Element struct {
	value Any
}

// NewElemenet creates new element with any value
func NewElement(value Any) *Element {
	return &Element{
		value: value,
	}
}

// Value returns element's value
func (e *Element) Value() Any {
	return e.value
}

// SetValue sets element's value
func (e *Element) SetValue(value Any) {
	e.value = value
}
