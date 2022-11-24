package graph

import (
	"fmt"
	"testing"
)

func TestGraph(t *testing.T) {
	// TODO: Implement reasonable test for graph (not only just print the nodes and its links)

	// See this great video about graph traversal: https://youtu.be/TIbUeeksXcI
	//
	//            G________B___
	//            /       /    \
	//			 	   /	  ___A___   |
	//		      /    /   |   \ /
	//         F____E    |    C
	//      		\    \___D___/
	//           \      /
	//            \____/
	//            H
	A, B, C, D, E, F, G, H := NewNode("A"), NewNode("B"), NewNode("C"), NewNode("D"), NewNode("E"), NewNode("F"), NewNode("G"), NewNode("H")
	LinkNodes(true, LinkMap[string]{
		A: {{B, 0}, {C, 0}, {D, 0}, {E, 0}},
		F: {{E, 0}, {G, 0}, {H, 0}},
		B: {{G, 0}, {C, 0}},
		D: {{C, 0}, {E, 0}},
		H: {{D, 0}},
	})
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
	// See this great video about Dijkstra: https://youtu.be/2poq1Pt32oE
	//     _______ C _____ E ______ Z
	//    /       / \  6  /    7   /
	//  3/      8/  4\   /10      /11
	//  /       /     \ /        /
	// A _____ B       D ______ F
	//  \  1  /       /    8
	//  5\   /2     2/
	//    \ /       /
	//     S _____ G
	//         4
	S, A, B, C, D, E, F, G, Z := NewNode("S"), NewNode("A"), NewNode("B"), NewNode("C"), NewNode("D"), NewNode("E"), NewNode("F"), NewNode("G"), NewNode("Z")
	LinkNodes(true, LinkMap[string]{
		S: {{A, 5}, {B, 2}, {G, 4}},
		A: {{B, 1}, {C, 3}},
		B: {{C, 8}},
		C: {{D, 4}, {E, 6}},
		D: {{E, 10}, {F, 8}},
		E: {{Z, 7}},
		F: {{Z, 11}},
		G: {{D, 2}},
	})

	route := S.Dijkstra(Z)
	fmt.Printf("%v -> %v\n", route.TotalCosts(), route)
	if route.TotalCosts() != 19 {
		t.Fail()
	}
}

type City struct {
	name      string
	heuristic int
}

func NewCity(name string, heuristic int) *City {
	return &City{name, heuristic}
}

func (c *City) String() string {
	return fmt.Sprintf("%v(%v)", c.name, c.heuristic)
}

func TestAStar(t *testing.T) {
	// TODO: Implement reasonable test for A* Algorithm and double check if this behaves as expected

	Wuerzburg, Frankfurt, Kaiserslautern := NewNode(NewCity("Wuerzburg", 0)), NewNode(NewCity("Frankfurt", 96)), NewNode(NewCity("Kaiserslautern", 158))
	Saarbruecken, Karlsruhe, Heilbronn, Ludwidgshafen := NewNode(NewCity("Saarbruecken", 222)), NewNode(NewCity("Karlsruhe", 140)), NewNode(NewCity("Heilbronn", 87)), NewNode(NewCity("Ludwidgshafen", 108))
	LinkNodes(true, LinkMap[*City]{
		Wuerzburg:      {{Frankfurt, 116}, {Ludwidgshafen, 183}, {Heilbronn, 102}},
		Frankfurt:      {{Kaiserslautern, 103}},
		Ludwidgshafen:  {{Kaiserslautern, 53}},
		Heilbronn:      {{Karlsruhe, 84}},
		Karlsruhe:      {{Saarbruecken, 145}},
		Kaiserslautern: {{Saarbruecken, 70}},
	})
	route := Saarbruecken.AStar(Wuerzburg, func(n *Node[*City]) int {
		return n.Data.heuristic
	})
	fmt.Printf("%v -> %v\n", route.TotalCosts(), route)
}
