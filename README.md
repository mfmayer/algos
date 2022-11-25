# al**go**s

Module, respectively collection of go packages concerning data structures and related *al**go***rithms.

## Data Structues

Currently implemented are

* Stack (LIFO)
* Ring Buffer (FIFO)
* Linked List
* Heap
* Graph

### Stack (LIFO)

```go
func ExampleStack() {
  // Create stack and initially push some elements -> 10 on top
  s := stack.NewStack(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
  s.Push(20)             // pushes 20 on top of the stack
  fmt.Println(s.Pop())   // pop top element (ouputs 20)
  fmt.Println(s.Peek(0)) // peek on top element without removing it from stack (outputs 10)
  // Output:
  // 20
  // 10
}
```

### Ring Buffer (FIFO)

```go
func ExampleNewRingBufferWithFixedCapacity() {
  // Creates ring buffer (FIFO) with fixed capacity of 5 elements
  rFixed := ringbuffer.NewRingBufferWithFixedCapacity[int](5)
  // Push 5 elements into the ring buffer
  rFixed.Push(1, 2, 3, 4, 5)
  rFixed.Push(6) // pushing another element 6 drops element 1, because ring buffer is full
  for rFixed.Len() > 0 {
    fmt.Printf("%v ", rFixed.Pop())
  }
  // Output: 2 3 4 5 6
}
```

> 💡 Using `NewRingBuffer()` instead of `NewRingBufferWithFixedCapacity(capacity int)` creates a ring buffer that dynamically grows and won't drop elements.

### Linked List (singly linked)

```go
func Example() {
  // Create linked list and return head element
  l1 := linkedlist.NewElement(1, 2, 3, 4)
  l5 := linkedlist.NewElement(5)
  l6 := linkedlist.NewElement(6, 7, 8, 9)
  l5.InsertBefore(l1)
  l5.InsertAfter(l6)

  next := l1
  for i := 1; next != nil; i++ {
    fmt.Printf("%v,", next.Data)
    next = next.Next()
  }
  // Output: 1,2,3,4,5,6,7,8,9,
}
```

### Heap

```go
func Example() {
  // Create new heap with providing an appropricate less function and the heap's inital elements
  h := heap.NewMaxHeap(func(a int, b int) bool { return a < b }, 10, 20, 15, 12, 40, 25, 18, 19)

  // Heap then looks like this:
  //        ___40___
  //       /        \
  //     25          19
  //    /  \        /  \
  //  20    10    18    15

  fmt.Println(h.Peek()) // In a max heap the first element is always the biggest
  fmt.Println(h.Sort()) // Sort the heap's slice. This method breaks the heap
  h.Heapify()           // Heapify repairs the heap's order e.g. after calling Sort method
  h.Push(99)
  // Pop all elements until heap is empty
  for h.Len() > 0 {
    e := h.Pop()
    fmt.Printf("%v ", e)
  }
  // Output:
  // 40
  // [10 12 15 18 19 20 25 40]
  // 99 40 25 20 19 18 15 12 10
}

```

### Graph

#### Traversal via DFS & BFS (Depth First Search & Breadth First Search)
```go
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
```

#### Dijkstra Path Search Algorithm

```go
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
```

#### A* Path Search Algorithm

```go
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
```