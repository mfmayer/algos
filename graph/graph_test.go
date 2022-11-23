package graph

import (
	"fmt"
	"testing"
)

func TestGraph(t *testing.T) {
	// TODO: Implement reasonable test for graph (not only just print the nodes and its links)
	A, B, C, D, E, F, G, H := NewNode("A"), NewNode("B"), NewNode("C"), NewNode("D"), NewNode("E"), NewNode("F"), NewNode("G"), NewNode("H")
	A.AddLinksBidir(Links(B, C)...)
	F.AddLinksBidir(Links(E, H, G)...)
	B.AddLinksBidir(Links(G, C)...)
	D.AddLinksBidir(Links(C, E)...)
	H.AddLinksBidir(Links(D)...)
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(D)
	fmt.Println(E)
	fmt.Println(F)
	fmt.Println(G)
	fmt.Println(H)

	// Traverse the graph
	visitFunc := func(n *Node[string]) {
		fmt.Println(n)
	}
	fmt.Println("Traversing tree via DFS:")
	A.DepthFirstSearch(visitFunc)

	fmt.Println("Traversing tree via BFS:")
	A.BreadthFirstSearch(visitFunc)
}
