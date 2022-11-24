package graph

// Route is a slice of links
type Route[T any] []Link[T]

// TotalCosts returns route total costs
func (r Route[T]) TotalCosts() int {
	totalCosts := 0
	for _, l := range r {
		totalCosts = totalCosts + l.Costs()
	}
	return totalCosts
}

// LastNode returns route's last linked node
func (r Route[T]) LastNode() *Node[T] {
	return r.LastLink().Node()
}

// LastLink returns route's last link
func (r Route[T]) LastLink() Link[T] {
	return r[len(r)-1]
}
