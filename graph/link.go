package graph

import "fmt"

func lessLinkCost[T any](a, b Link[T]) bool {
	return a.Costs() < b.Costs()
}

// Link to a node with link costs
type Link[T any] struct {
	node  *Node[T]
	costs int
}

// Node reutrns linked node
func (l Link[T]) Node() *Node[T] {
	return l.node
}

// Costs reutrns link costs
func (l Link[T]) Costs() int {
	return l.costs
}

// String creates string representation of the link
func (l Link[T]) String() string {
	return fmt.Sprintf("%v(%v)", l.node.Data, l.costs)
}
