package graph

import (
	"fmt"
	"testing"
)

func TestGraph(t *testing.T) {
	// TODO: Implement reasonable test for graph (not only just print the nodes and its links)
	A, B, C, D, E, F, G, H := NewNode("A"), NewNode("B"), NewNode("C"), NewNode("D"), NewNode("E"), NewNode("F"), NewNode("G"), NewNode("H")
	A.AddLinksTo(Nodes(B, C, D, E), Bidirectional)
	F.AddLinksTo(Nodes(E, H, G), Bidirectional)
	B.AddLinksTo(Nodes(G, C), Bidirectional)
	D.AddLinksTo(Nodes(C, E), Bidirectional)
	H.AddLinksTo(Nodes(D), Bidirectional)
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(D)
	fmt.Println(E)
	fmt.Println(F)
	fmt.Println(G)
	fmt.Println(H)

	visitFunc := func(n AnyNode) {
		fmt.Println(n)
	}
	fmt.Println("Traversing tree via DFS:")
	A.DepthFirstSearch(visitFunc)
}
