package graph_test

import (
	"fmt"

	"github.com/mfmayer/algos/graph"
)

// Example_graph_traversal shows how to build a graph and how to traverse it via DFS & BFS
func Example_graph_traversal() {
	// See this great video about graph traversal: https://youtu.be/TIbUeeksXcI
	//
	//	    G_____A_____H
	//	   /     / \   /
	//    /     /   \ /
	//   F     C_____B
	//    \   /     /
	//     \ /     /
	//      D_____E

	A, B, C, D, E, F, G, H :=
		graph.NewNode("A"), graph.NewNode("B"), graph.NewNode("C"),
		graph.NewNode("D"), graph.NewNode("E"), graph.NewNode("F"),
		graph.NewNode("G"), graph.NewNode("H")

	graph.LinkNodes(true, graph.LinkMap[string]{
		A: {{G, 1}, {C, 2}, {B, 3}, {H, 4}},
		B: {{E, 1}, {C, 2}, {H, 3}},
		C: {{D, 1}},
		D: {{F, 1}, {E, 2}},
		F: {{G, 0}},
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
	visitFunc := func(n *graph.Node[string]) {
		fmt.Println(n)
	}
	fmt.Println("Traversing graph via DFS:")
	A.DepthFirstSearch(visitFunc)

	fmt.Println("Traversing graph via BFS:")
	A.BreadthFirstSearch(visitFunc)
	// Output:
	// A -> [G(1) C(2) B(3) H(4)]
	// B -> [E(1) C(2) H(3) A(3)]
	// C -> [D(1) A(2) B(2)]
	// D -> [F(1) C(1) E(2)]
	// E -> [B(1) D(2)]
	// F -> [G(0) D(1)]
	// G -> [F(0) A(1)]
	// H -> [B(3) A(4)]
	// Traversing graph via DFS:
	// A -> [G(1) C(2) B(3) H(4)]
	// G -> [F(0) A(1)]
	// F -> [G(0) D(1)]
	// D -> [F(1) C(1) E(2)]
	// C -> [D(1) A(2) B(2)]
	// B -> [E(1) C(2) H(3) A(3)]
	// E -> [B(1) D(2)]
	// H -> [B(3) A(4)]
	// Traversing graph via BFS:
	// A -> [G(1) C(2) B(3) H(4)]
	// H -> [B(3) A(4)]
	// B -> [E(1) C(2) H(3) A(3)]
	// C -> [D(1) A(2) B(2)]
	// G -> [F(0) A(1)]
	// E -> [B(1) D(2)]
	// D -> [F(1) C(1) E(2)]
	// F -> [G(0) D(1)]
}

func Example_dijkstra() {
	//
	//  O _____ N _____ J       _______ C _____ E _____ Z
	//   \  ⁷  /    ⁶  / \     /       / \  ⁶  /    ⁷  /
	//  ¹⁰\   /⁵      /⁵  \₄  /³     ₉/  ₄\   /¹⁰     /₁₁
	//     \ /       /     \ /       /     \ /       /
	//      M       I       A _____ B       D _____ F
	//     /       / \     / \  ¹  /       /    ⁷
	//   ₃/       /⁵  \₃  /²  \₅  /²     ₂/
	//   /       /     \ /     \ /       /
	//  L _____ K       H       S _____ G
	//      ⁶                       ⁴

	E, F, C, D, G, B, S, A, H, J, I, K, L, M, N, O, Z :=
		graph.NewNode(NewHNode("E", 7)), graph.NewNode(NewHNode("F", 10)),
		graph.NewNode(NewHNode("C", 11)), graph.NewNode(NewHNode("D", 15)),
		graph.NewNode(NewHNode("G", 20)), graph.NewNode(NewHNode("B", 20)),
		graph.NewNode(NewHNode("S", 22)), graph.NewNode(NewHNode("A", 21)),
		graph.NewNode(NewHNode("H", 23)), graph.NewNode(NewHNode("J", 25)),
		graph.NewNode(NewHNode("I", 26)), graph.NewNode(NewHNode("K", 28)),
		graph.NewNode(NewHNode("L", 30)), graph.NewNode(NewHNode("M", 32)),
		graph.NewNode(NewHNode("N", 28)), graph.NewNode(NewHNode("O", 32)),
		graph.NewNode(NewHNode("Z", 0))

	graph.LinkNodes(true, graph.LinkMap[*HNode]{
		Z: {{E, 7}, {F, 11}},
		D: {{E, 10}, {F, 7}, {C, 4}, {G, 2}},
		B: {{C, 9}, {S, 2}, {A, 1}},
		S: {{G, 4}, {A, 5}},
		A: {{C, 3}, {H, 2}, {J, 4}},
		I: {{J, 5}, {H, 3}, {K, 5}},
		N: {{J, 6}, {M, 0}, {O, 7}},
		L: {{M, 3}, {K, 6}},
	})
	route, steps := S.Dijkstra(Z)
	fmt.Printf("total costs %v: %v (visited nodes: %v)\n", route.TotalCosts(), route, steps)
	// Output:
	// total costs 23: [S(22)(0) G(20)(4) D(15)(2) E(7)(10) Z(0)(7)] (visited nodes: 40)
}

type HNode struct {
	name      string
	heuristic int
}

func NewHNode(name string, heuristic int) *HNode {
	return &HNode{name, heuristic}
}

func (hn *HNode) String() string {
	return fmt.Sprintf("%v(%v)", hn.name, hn.heuristic)
}

func Example_astar() {
	//
	//  O _____ N _____ J       _______ C _____ E _____ Z
	//   \  ⁷  /    ⁶  / \     /       / \  ⁶  /    ⁷  /
	//  ¹⁰\   /⁵      /⁵  \₄  /³     ₉/  ₄\   /¹⁰     /₁₁
	//     \ /       /     \ /       /     \ /       /
	//      M       I       A _____ B       D _____ F
	//     /       / \     / \  ¹  /       /    ⁷
	//   ₃/       /⁵  \₃  /²  \₅  /²     ₂/
	//   /       /     \ /     \ /       /
	//  L _____ K       H       S _____ G
	//      ⁶                       ⁴

	E, F, C, D, G, B, S, A, H, J, I, K, L, M, N, O, Z :=
		graph.NewNode(NewHNode("E", 7)), graph.NewNode(NewHNode("F", 10)),
		graph.NewNode(NewHNode("C", 11)), graph.NewNode(NewHNode("D", 15)),
		graph.NewNode(NewHNode("G", 20)), graph.NewNode(NewHNode("B", 20)),
		graph.NewNode(NewHNode("S", 22)), graph.NewNode(NewHNode("A", 21)),
		graph.NewNode(NewHNode("H", 23)), graph.NewNode(NewHNode("J", 25)),
		graph.NewNode(NewHNode("I", 26)), graph.NewNode(NewHNode("K", 28)),
		graph.NewNode(NewHNode("L", 30)), graph.NewNode(NewHNode("M", 32)),
		graph.NewNode(NewHNode("N", 28)), graph.NewNode(NewHNode("O", 32)),
		graph.NewNode(NewHNode("Z", 0))

	graph.LinkNodes(true, graph.LinkMap[*HNode]{
		Z: {{E, 7}, {F, 11}},
		D: {{E, 10}, {F, 7}, {C, 4}, {G, 2}},
		B: {{C, 9}, {S, 2}, {A, 1}},
		S: {{G, 4}, {A, 5}},
		A: {{C, 3}, {H, 2}, {J, 4}},
		I: {{J, 5}, {H, 3}, {K, 5}},
		N: {{J, 6}, {M, 0}, {O, 7}},
		L: {{M, 3}, {K, 6}},
	})
	route, steps := S.AStar(Z, func(n *graph.Node[*HNode]) int {
		return n.Data.heuristic
	})
	fmt.Printf("total costs %v: %v (visited nodes: %v)\n", route.TotalCosts(), route, steps)
	// Output:
	// total costs 23: [S(22)(0) G(20)(4) D(15)(2) E(7)(10) Z(0)(7)] (visited nodes: 8)
}
