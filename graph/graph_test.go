package graph

import (
	"fmt"
	"testing"
)

func TestGraph(t *testing.T) {
	// TODO: Implement reasonable test for graph (not only just print the nodes and its links)
	A, B, C, D, E, F, G, H := NewNode("A"), NewNode("B"), NewNode("C"), NewNode("D"), NewNode("E"), NewNode("F"), NewNode("G"), NewNode("H")
	A.LinkTo(Nodes(B, C), BidirectionalWithCosts(5), Nodes(D, E), BidirectionalWithCosts(2))
	F.LinkTo(Nodes(E, H, G), BidirectionalWithCosts(1))
	B.LinkTo(Nodes(G, C), BidirectionalWithCosts(1))
	D.LinkTo(Nodes(C, E), BidirectionalWithCosts(1))
	H.LinkTo(Nodes(D), BidirectionalWithCosts(1))
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(D)
	fmt.Println(E)
	fmt.Println(F)
	fmt.Println(G)
	fmt.Println(H)

	// Traverse the graph
	visitFunc := func(n AnyNode) {
		fmt.Println(n)
	}
	fmt.Println("Traversing tree via DFS:")
	A.DepthFirstSearch(visitFunc)

	fmt.Println("Traversing tree via BFS:")
	A.BreadthFirstSearch(visitFunc)
}
