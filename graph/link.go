package graph

import "fmt"

func lessLinkCost[T any](a, b Link[T]) bool {
	return a.Costs < b.Costs
}

// Link to a node with link costs
type Link[T any] struct {
	Node  *Node[T]
	Costs int
}

// String creates string representation of the link
func (l Link[T]) String() string {
	return fmt.Sprintf("%v(%v)", l.Node.Data, l.Costs)
}
