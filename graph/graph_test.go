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

func TestDijkstra(t *testing.T) {
	S, A, B, C, D, E, F, G, Z := NewNode("S"), NewNode("A"), NewNode("B"), NewNode("C"), NewNode("D"), NewNode("E"), NewNode("F"), NewNode("G"), NewNode("Z")

	S.AddLinksBidir([]Link[string]{{A, 5}, {B, 2}, {G, 4}}...)
	A.AddLinksBidir([]Link[string]{{B, 1}, {C, 3}}...)
	B.AddLinksBidir([]Link[string]{{C, 8}}...)
	C.AddLinksBidir([]Link[string]{{D, 4}, {E, 6}}...)
	D.AddLinksBidir([]Link[string]{{E, 10}, {F, 8}}...)
	E.AddLinksBidir([]Link[string]{{Z, 7}}...)
	F.AddLinksBidir([]Link[string]{{Z, 11}}...)
	G.AddLinksBidir([]Link[string]{{D, 2}}...)

	route := S.Dijkstra(Z)
	fmt.Printf("%v -> %v\n", route.TotalCosts(), route)
	if route.TotalCosts() != 19 {
		t.Fail()
	}
}
